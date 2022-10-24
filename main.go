package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func main() {
	err := http.ListenAndServe(
		":18080",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	)
	if err != nil {
		fmt.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}

const englishHelloPrefix = "Hello, "

// Hello テスト実行確認
func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + name
}
