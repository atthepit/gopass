package pass_generator

import (
	"database/sql"
	"errors"
	_ "github.com/ziutek/mymysql/godrv"
)

type S struct {
	Host string
	Name string
	User string
	Pass string
	db   *sql.DB
	err  error
}

var Saver = &S{"localhost", "passwords", "go", "golang", nil, nil}

func (s *S) Get_instance(host, name, user, pass string) (*sql.DB, error) {
	if Saver.db == nil {
		if host == "" {
			host = Saver.Host
		}
		if name == "" {
			name = Saver.Name
		}
		if user == "" {
			user = Saver.User
		}
		if pass == "" {
			pass = Saver.Pass
		}
		connection_string := name + "/" + user + "/" + pass
		Saver.db, Saver.err = sql.Open("mymysql", connection_string)
	}
	return Saver.db, Saver.err
}

func (s *S) Close() error {
	if Saver.db == nil {
		Saver.err = errors.New("La conexión con la Base de Datos no está establecida")
	} else {
		Saver.err = Saver.Close()
	}
	return Saver.err
}
