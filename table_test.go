package conector

import "testing"

func TestCreateTables(t *testing.T) {
	created, err := createTables(Usuario{}, Compra{})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%d Tablas creadas: \n %s", len(created), created)
}
