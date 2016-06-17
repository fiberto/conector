package conector

import (
	"testing"
	"time"
)

func TestUsuario(t *testing.T) {
	//Registro
	u := Usuario{
		Correo: "mario@example.com",
		Nombre: "Mario",
	}
	err := u.Registrar()
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		t.Log("Registrado: ", u)
	}
	//Búsqueda por ID
	u = Usuario{
		Correo: "mario@example.com",
	}
	err = u.BuscarPorID()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log("Encontrado: ", u)
}

func TestCompra(t *testing.T) {
	//Registro
	c := Compra{
		Fecha: time.Now(),
		Total: 100.00,
	}
	err := c.Registrar()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log("Registrado :", c)

	//Búsqueda por ID
	c = Compra{
		ID: c.ID,
	}
	err = c.BuscarPorID()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log("Encontrado: ", c)
}
