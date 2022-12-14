package main

import (
	"context"
	"fmt"
	"go-port/config"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func main() {
	// メインプロセスの処理をrun関数で処理させて、mainは結果を受けとるだけにする
	// Goの実装パターンでよく使われる模様
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}

// ここにはサーバーの起動を別で記載する
// main関数とは別に記載するのはGoではよくあること
func run(ctx context.Context) error {

	cfg, err := config.New()
	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}

	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)

	mux, cleanup, err := NewMux(ctx, cfg)
	if err != nil {
		return err
	}
	defer cleanup()
	s := NewServer(l, mux)
	return s.Run(ctx)

}
