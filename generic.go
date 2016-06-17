package conector

import (
	"errors"
	"reflect"
	"strings"

	r "github.com/dancannon/gorethink"
)

//insert usa la sesión activa para registrar la estructura en la base de datos.
func insert(i interface{}) (keys []string, err error) {
	err = reconectar()
	if err != nil {
		return
	}
	//El nombre de la tabla será el nombre de la estructura
	table := getTable(i)
	resp, err := r.DB(bd).Table(table).Insert(i).RunWrite(session)
	if err != nil {
		if r.IsConflictErr(err) {
			//Llave duplicada
			return
		}
		return
	}
	return resp.GeneratedKeys, err
}

//queryByID usa el ID de la estructura para realizar la búsqueda, debe recibir un apuntador para llenar la estructura
func queryByID(i interface{}) error {
	err := reconectar()
	if err != nil {
		return err
	}
	//El nombre de la tabla será el nombre de la estructura
	table := getTable(i)
	//El ID para la búsqueda es el valor del campo que contiene el tag: gorethink:id
	id, err := getID(i)
	if err != nil {
		//No se pudo obtener el tag
		return err
	}
	resp, err := r.DB(bd).Table(table).Get(id).Run(session)
	if err != nil {
		return err
	}

	err = resp.One(i)
	if err == r.ErrEmptyResult {
		//No se encontró
		return err
	}
	return err
}

//getID usa reflect para obtener el ID de la estructura.
func getID(i interface{}) (string, error) {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get("gorethink")
		tags := strings.Split(tag, ",")
		for i := range tags {
			if tags[i] == "id" {
				return v.Field(i).String(), nil
			}
		}
	}
	return "", errors.New("No se encontró la etiqueta gorethink:id")
}
