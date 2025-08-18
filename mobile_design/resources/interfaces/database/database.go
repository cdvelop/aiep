package database

import "fmt"

type DataBase struct {
}

func (db DataBase) GetForId(id string) []string {

	return []string{fmt.Sprintf("<data>%s</data>", id)}
}
