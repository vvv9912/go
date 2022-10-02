package database


type database struct {
	id       int
	username string
	password string
	toDO     toDo
}

type toDO struct {
	day string
	toDO string
	//
}
type databasefunc interface {
	readdatabase()
	writedatabase()
}

