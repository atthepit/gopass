package pass_generator

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	db_name           string = "gopass.db"
	db_driver         string = "sqlite3"
	connection_string string = "./" + db_name
)

func Create_DB() {

	db, err := sql.Open(db_driver, connection_string)
	check_err(err)
	defer db.Close()

	sql := `
	CREATE TABLE IF NOT EXISTS users (
  	  name varchar(45) NOT NULL,
  	  master longblob NOT NULL,
  	  salt longblob NOT NULL,
  	  PRIMARY KEY (name)
	);
	`
	_, err = db.Exec(sql)
	check_err(err)

	sql = `
	CREATE TABLE IF NOT EXISTS sites (
		"name" VARCHAR(90) NOT NULL , 
		"password" BLOB NOT NULL , 
		"user" VARCHAR(45) NOT NULL , 
		PRIMARY KEY ("name", "user"),
		FOREIGN KEY (user) REFERENCES users (name) ON DELETE CASCADE ON UPDATE CASCADE
	);`
	_, err = db.Exec(sql)
	check_err(err)
}

func Create_user(user_name string, master_pass, salt []byte) {
	db, err := sql.Open(db_driver, connection_string)
	check_err(err)
	defer db.Close()

	ps, err := db.Prepare("INSERT INTO users (name,master,salt) VALUES (?,?,?);")
	check_err(err)

	_, err = ps.Exec(user_name, master_pass, salt)
	check_err(err)
}

func Delete_user(user_name string) {
	db, err := sql.Open(db_driver, connection_string)
	check_err(err)
	defer db.Close()

	ps, err := db.Prepare("DELETE FROM `sites` WHERE `user`=?")
	check_err(err)
	_, err = ps.Exec(user_name)
	check_err(err)

	ps, err = db.Prepare("DELETE FROM `users` WHERE `name`=?")
	check_err(err)
	_, err = ps.Exec(user_name)
	check_err(err)
}

func Get_user(user_name string) (int, []byte, []byte) {
	db, err := sql.Open(db_driver, connection_string)
	check_err(err)
	defer db.Close()

	rows, err := db.Query("SELECT count(*), master, salt FROM users WHERE users.name like ?", user_name)
	check_err(err)

	rows.Next()
	var num int
	var master, salt []byte
	err = rows.Scan(&num, &master, &salt)
	check_err(err)
	rows.Close()

	return num, master, salt
}

func Create_site(site_name, user_name string, password []byte) {
	db, err := sql.Open(db_driver, connection_string)
	check_err(err)
	defer db.Close()

	ps, err := db.Prepare("INSERT INTO sites (name,password,user) VALUES (?,?,?);")
	check_err(err)

	_, err = ps.Exec(site_name, password, user_name)
	check_err(err)
}

func Get_password(site_name, user_name string) []byte {
	db, err := sql.Open(db_driver, connection_string)
	check_err(err)
	defer db.Close()

	rows, err := db.Query("SELECT password FROM sites WHERE sites.user LIKE ? AND sites.name LIKE ?;", user_name, site_name)
	check_err(err)

	rows.Next()
	var pass []byte
	err = rows.Scan(&pass)
	check_err(err)
	rows.Close()

	return pass
}

func List_sites(user_name string) []string {
	db, err := sql.Open(db_driver, connection_string)
	check_err(err)
	defer db.Close()

	rows, err := db.Query("SELECT count(*), name FROM sites WHERE sites.user LIKE ?", user_name)
	check_err(err)

	rows.Next()
	var name string
	var num int
	rows.Scan(&num, &name)
	names := make([]string, num)
	if num > 0 {
		names[0] = name
		i := 1

		for rows.Next() {
			name = ""
			rows.Scan(&name)
			names[i] = name
			i++
		}
	}
	rows.Close()

	return names
}

func Delete_site(site_name, user_name string) {
	db, err := sql.Open(db_driver, connection_string)
	check_err(err)
	defer db.Close()

	ps, err := db.Prepare("DELETE FROM `sites` WHERE `name`=? and`user`=?")
	check_err(err)

	_, err = ps.Exec(site_name, user_name)
	check_err(err)
}

func Update_password(site_name, user_name string, pass []byte) {
	db, err := sql.Open(db_driver, connection_string)
	check_err(err)
	defer db.Close()

	ps, err := db.Prepare("UPDATE sites SET password=? WHERE user=? AND name=?;")
	check_err(err)

	_, err = ps.Exec(pass, user_name, site_name)
	check_err(err)
}
