package core
import (
	"database/sql"
	"fmt"  
	_ "github.com/lib/pq"
	)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "passamogh123"
	dbname   = "olivia"
)

  func oliviasend(id string, command string, response string)string{
	var status string
	var err error
	var db *sql.DB
	if db == nil {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err = sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
		panic(err)
	}
	}
	  sqlStatement := `SELECT cmd, cmd FROM responses WHERE cmd=$1;`

		var cmd string
		var idk string
	  var norow bool = false

	  row := db.QueryRow(sqlStatement, command)
	  switch err = row.Scan(&cmd, &idk); err {
	  case sql.ErrNoRows:
		  norow = true
		  status = "New command created!"
	  case nil:
		status = "Command already exists!"
	  default:
		panic(err)
	  }
	  if (norow ==true){
		sqlStatement := `
		INSERT INTO responses (user_id, cmd, res)
		VALUES ($1, $2, $3)
		RETURNING cmd`
		  err = db.QueryRow(sqlStatement, id, command, response).Scan(&cmd)
		  if err != nil {
			panic(err)
		  }
	  }
	  return status
  }

  func oliviadelete( command string )string{
	var status string
	var err error
	var db *sql.DB
	if db == nil {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err = sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
		panic(err)
	}
	}
	  sqlStatement := `SELECT cmd, cmd FROM responses WHERE cmd=$1;`

	  var cmd string
	  var idk string
	  var norow bool = false

	  row := db.QueryRow(sqlStatement, command)
	  switch err = row.Scan(&cmd, &idk); err {
	  case sql.ErrNoRows:
		  status = "Command does not exist!"
	  case nil:
		norow = true
		status = "Command deleted!"
	  default:
		panic(err)
	  }
	  if (norow ==true){
		sqlStatement := `
		DELETE FROM responses
		WHERE cmd = $1;`
		_, err = db.Exec(sqlStatement, command)
		if err != nil {
		  panic(err)
		}
	  }
	  return status
  }

  func oliviafetch( command string )string{
	var response string = "Command not assigned"
	var err error
	var db *sql.DB
	if db == nil {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err = sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
		panic(err)
	}
	}
	  var res string
		var cmd string
		var rows *sql.Rows
	  rows, err = db.Query("SELECT cmd, res FROM responses WHERE cmd=$1", command)
	  if err != nil {
		panic(err)
	  }
	  defer rows.Close()
	  for rows.Next() {
		err = rows.Scan(&cmd, &res)
		if err != nil {
		  panic(err)
		}
		response=res
	  }
	  err = rows.Err()
	  if err != nil {
		panic(err)
	  }
	  return response
	}
	
	func oliviainsult() string {
	var db *sql.DB
	if db == nil {
	var err error
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err = sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
		panic(err)
	}
	}
	
	//max int = 531
	randomval := Uint32n(530)
	randomval = randomval + 1
	
	  var res string
		var id int
		var err error
		var rows *sql.Rows
		err = db.Ping()
		if err != nil {
			panic(err)
			fmt.Println("WORks till here")
		}
		rows, err = db.Query("SELECT res, id FROM insults WHERE id=$1", randomval)
	  if err != nil {
		panic(err)
		}
	  defer rows.Close()
	  for rows.Next() {
		err = rows.Scan(&res, &id)
		if err != nil {
		  panic(err)
		}
	  }
	  err = rows.Err()
	  if err != nil {
		panic(err)
	  }
	  return res
	}
  