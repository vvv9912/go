package lib1

import (
	"test/test2/lib2"
)

type Lib1 struct {
}

func New() *Lib1 {
	return &Lib1{}
}

// Заменять возвращаемый тип, запрещено.
func (l Lib1) Do() lib2.SomeType {
	return lib2.SomeType{Data: "data from lib1"}
}
