package conector

//Usuario es utilizada para demostrar el funcionamiento con estructuras que
//no contienen campo diferente a ID como identificador para RethinkDB
type Usuario struct {
	Correo string `gorethink:"id"`
	Nombre string
}

//Registrar usa la estructura y no copia llaves generadas porque no se necesitan.
func (u *Usuario) Registrar() error {
	_, err := insert(u)
	if err != nil {
		return err
	}
	return nil
}

//BuscarPorID usará el ID de la estructura para realizar la búsqueda.
func (u *Usuario) BuscarPorID() error {
	err := queryByID(u)
	if err != nil {
		return err
	}
	return nil
}
