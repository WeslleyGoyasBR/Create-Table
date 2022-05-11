package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {

	var conn *pgx.Conn
	var err error

	conn, err = pgx.Connect(context.Background(), os.Getenv("DATATESTE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "impossivel conectar a banco de dados %v\n", err)
		os.Exit(1)
	}
	sql := "CREATE TABLE CONTATO(cod uuid, nome varchar, sobrenome varchar, telefone varchar, cidade varchar)"
	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "erro ao criar tabela, contate suporte %v", err)
		os.Exit(2)

	}
}
