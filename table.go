package conector

import (
	"fmt"
	"reflect"

	r "github.com/dancannon/gorethink"
)

//createTables usa las estructuras para crear las tablas según sus nombres.
func createTables(tables ...interface{}) (created []string, err error) {
	err = reconectar()
	if err != nil {
		return
	}
	for _, t := range tables {
		name := getTable(t)
		_, err := r.DB(bd).TableCreate(name).RunWrite(session)
		if err != nil {
			fmt.Printf("Error al crear tabla %s : %s\n", name, err.Error())
			continue
		}
		created = append(created, name)
	}
	return
}

func getTable(i interface{}) string {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Type().Name()
}
