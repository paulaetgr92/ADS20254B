package dataSrc

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var db *sql.DB
	var err error

	// Retry para esperar o Postgres ficar pronto
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Printf("Tentativa %d: Erro ao abrir conexão: %v", i+1, err)
		} else if err = db.Ping(); err == nil {
			fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
			return db, nil
		} else {
			log.Printf("Tentativa %d: Erro ao conectar ao DB: %v", i+1, err)
		}
		time.Sleep(3 * time.Second)
	}

	return nil, fmt.Errorf("não foi possível conectar ao banco de dados após várias tentativas: %v", err)
}
