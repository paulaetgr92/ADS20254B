package dataSrc

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "hoot"
	dbname := "meu_banco"

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Erro ao abrir conexão com DB: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao conectar ao DB: %v", err)
		return nil, err
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
	return db, nil
}
