package conector

import r "github.com/dancannon/gorethink"

var (
	bd         = "test"                                                        //Base de datos
	session    *r.Session                                                      //Guarda la sesión
	connParams = r.ConnectOpts{Address: "localhost", MaxOpen: 50, MaxIdle: 30} //Utilizado para hacer un pool de conexiones
)

//init crea la conexión con el server. Hace panic si no se establece la conexión.
func init() {
	var err error
	session, err = conectar()
	if err != nil {
		panic(err)
	}
}

//conectar abre una conexión a la base de datos.
func conectar() (*r.Session, error) {
	session, err := r.Connect(connParams)
	if err != nil {
		return nil, err
	}
	session.Use(bd)
	return session, nil
}

//reconectar `activa` la conexión a la base de datos.
func reconectar() (err error) {
	if !session.IsConnected() {
		err = session.Reconnect(r.CloseOpts{})
	}
	return
}
