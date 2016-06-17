package conector

import "time"

//Compra es utilizado para probar el funcionamiento con estructuras que
//contienen ID por defecto para id de RethinkDB.
type Compra struct {
	ID    string `gorethink:"id,omitempty"`
	Fecha time.Time
	Total float64
}

//Registrar utiliza la estructura y copia la llave generada al ID de la estructura.
func (c *Compra) Registrar() error {
	keys, err := insert(c)
	if err != nil {
		return err
	}
	c.ID = keys[0]
	return nil
}

//BuscarPorID usará el ID de la estructura para realizar la búsqueda.
func (c *Compra) BuscarPorID() error {
	err := queryByID(c)
	if err != nil {
		return err
	}
	return nil
}
