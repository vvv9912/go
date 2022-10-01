package main

import (
	"fmt"

	"test/test2/lib1"
	"test/test2/lib2"
)

// Вам необходимо решить проблему с "import cycle not allowed" с ограничениями, которые указаны в файлах lib1 и lib2.
// Если возникли сложности с данным заданием, пропускайте его и переходите к заданию test3.
func main() {
	l1 := lib1.New()
	l2 := lib2.New(l1)
	result := l2.Do()
	fmt.Println(result.Data)
}
