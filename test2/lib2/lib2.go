package lib2

import (
	"test/test2/lib1"
)

type Lib1 interface {
	Do()
}

type Lib2 struct {
	l1 *lib1.Lib1
}

// Перемещать объявление данного типа запрещено.
type SomeType struct {
	Data string
}

func New(l1 *lib1.Lib1) *Lib2 {
	return &Lib2{
		l1: l1,
	}
}

func (l Lib2) Do() SomeType {
	return l.l1.Do()
}
