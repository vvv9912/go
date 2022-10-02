package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"
)

// Test1 принимает на вход список строк и выводит их в одной строке через пробел, возможно, в случайном порядке.
// Смысл этого кода не в конкатенации слов через пробел, а в понимании работы многопоточности.
//
// Вам необходимо исправить код так, чтобы он работал, не удаляя ничего из него, а только внося исправления.
//
// Доп. задание: напишите тест на эту функцию.
//
// Пример вызова функции:
//
//	data1 := []string{"hello", "world", "test", "data", "code"}
//	r1 := Test1(data1)
//	fmt.Println(r1)
func Test1(list []string) string {
	var result string
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, l := range list {

		wg.Add(1)
		go func(l string) {
			defer wg.Done()
			mu.Lock()
			result += l + " "
			mu.Unlock()
		}(l)
	}
	wg.Wait()
	return result
}

// Test2 создаёт файл во временной директории и записывает в него строку из data.
//
// Вам необходимо исправить его так, чтобы он работал верно под любой ОС.
//
// Доп. задание: Можно ли упростить данный код?
//
// Пример вызова функции:
//
//	err := Test2("test2.txt", "hello test2")
func Test2(filename, data string) error {
	dir, err := os.MkdirTemp("", "tmp")

	f, err := os.Create(path.Join(dir, filename))
	if err != nil {
		return err
	}
	defer f.Close()
	// !упростить - использовать os.WriteFile
	_, err = f.Write([]byte(data))
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)
	return nil
}

// Test3 отправляет запрос через прокси и возвращает внешний IP.
//
// Вам необходимо реализовать поддержку прокси.
//
// Пример вызова функции:
//
//	ip, err := Test3("http://user:pass@127.0.0.1:3128/")
func Test3(proxyURL string) (string, error) {
	// TODO: использовать прокси
	proxyURLL, err := url.Parse(proxyURL)
	if err != nil {
		log.Println(err)
	}

	transport := &http.Transport{
		Proxy:           http.ProxyURL(proxyURLL),
		IdleConnTimeout: 30 * time.Second,
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Second * 30,
	}

	resp, err := client.Get("https://httpbin.org/get")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var v struct {
		Origin string
	}
	err = json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		return "", err
	}

	return v.Origin, nil
}

// Test4 оборачивает соединение, для подсчёта принятых/переданных данных.
//
// Вам необходимо реализовать данный функционал с поддержкой многопоточного доступа к данным.
// По желанию, можете написать тест на данную функцию.
//
// Пример вызова функции:
//
//	conn, err := net.Dial(...)
//	...
//	countConn := Test4(conn)
//	_, err = conn.Write([]byte("send test data"))
//	...
//	fmt.Printf("Write: %d\n", countConn.WriteByteCount())
//
//	...
//	_, err := conn.Read(buf)
//	...
//	fmt.Printf("Read: %d\n", countConn.ReadByteCount())
//
//	err = conn.Close()
//	...
//	fmt.Printf("Read: %d\n", countConn.ReadByteCount())
//	fmt.Printf("Write: %d\n", countConn.WriteByteCount())
//
// func read()

type con struct {
	net.Conn
	nRead  uint64
	nWrite uint64
	rmuw   sync.RWMutex
	rmur   sync.RWMutex
	//muw    sync.Mutex
	//	mur    sync.Mutex
}

func Test4(conn net.Conn) Test4Conn {
	cout := &con{}
	cout.Conn = conn
	return cout
}

type Test4Conn interface {
	net.Conn
	ReadByteCount() uint64
	WriteByteCount() uint64
}

func (c *con) ReadByteCount() uint64 {
	return c.nRead
}

func (c *con) WriteByteCount() uint64 {
	return c.nWrite
}

func (c *con) Write(b []byte) (int, error) {

	var n int
	var err error

	n, err = c.Conn.Write(b)
	if err != nil {
		return n, err
	}
	c.rmuw.Lock()
	c.nWrite = uint64(n)
	c.rmuw.Unlock()

	return n, err
}

func (c *con) Read(b []byte) (int, error) {

	var n int
	var err error

	n, err = c.Conn.Read(b)
	if err != nil {
		return n, err
	}
	c.rmur.Lock()
	c.nRead = uint64(n)
	c.rmur.Unlock()

	return n, err
}

func test_test1(test *testing.T) {
	data1 := []string{"hello", "world", "test", "data", "code"}
	var data11 []string
	r1 := Test1(data1)
	var cout int
	for i := 0; i < len(data1); i++ {
		if createWordRegex(data1[i]).MatchString(r1) {
			cout += 1
			r1 = strings.ReplaceAll(r1, data1[i], "")
		} else {
			data11 = append(data11, data1[i])
		}

	}
	if len(data1) != cout {
		test.Errorf("error, word: %s", data11)
	}

}

func test_test3(test *testing.T) {
	_, err := Test3("http://user:pass@127.0.0.1:3128/")
	if err != nil {
		test.Errorf("Error")
	}
}

func test_test4(test *testing.T) {
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		test.Errorf("Error")
		return
	}
	countConn := Test4(conn)
	conn = countConn

	tmp := make([]byte, 4*1024)
	httpRequest := "111" + "Host: golang.org\n\n"
	_, err = conn.Write([]byte(httpRequest))
	if err != nil {
		test.Errorf("Error")
		return
	}
	_, err = conn.Read(tmp)
	if err != nil {
		test.Errorf("Error")
		return
	}

}
func createWordRegex(word string) *regexp.Regexp {
	return regexp.MustCompile(`\b` + regexp.QuoteMeta(word) + `\b`)
}
func main() {
	//start test1
	/*data1 := []string{"hello", "world", "test", "data", "code"}
	r1 := Test1(data1)
	fmt.Println(r1)*/

	//start test2
	/*err := Test2("test2.txt", "hello test2")
	fmt.Println(err)*/

	//start test3
	/*	ip, err := Test3("http://user:pass@127.0.0.1:3128/")
		fmt.Println(ip)
		fmt.Println(err)*/

	//test4
	/*
		conn, err := net.Dial("tcp", "golang.org:80")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()
		countConn := Test4(conn)
		conn = countConn

		tmp := make([]byte, 4*1024)
		httpRequest := "GET / HTTP/1.1\n" +
			"Host: golang.org\n\n"
		_, err = conn.Write([]byte(httpRequest))
		_, err = conn.Read(tmp)
		fmt.Printf("Write: %d\n", countConn.WriteByteCount())
		fmt.Printf("Read: %d\n", countConn.ReadByteCount())*/
}
