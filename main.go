package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/sync/errgroup"
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
	// net/httpクラスを利用した
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

// ここにはサーバーの起動を別で記載する
// main関数とは別に記載するのはGoではよくあること
func run(ctx context.Context) error {

	// http.ListenAndServe よりも下記の店でメリットがある模様
	// - shutdown() メソッドがあり、メソッドの呼び出しでサーバー中断が可能(ListenAndServe だとプロセスキルによる強制終了しか選択肢がない)
	s := &http.Server{
		Addr: ":18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}

	// 以下はgolang.org/x/sync の errgroupサブパッケージの利用例
	// errorgroupのインスタンスに対してGorutineを書きくだし、書き下した処理内でエラーが発生したときに err として返送してもらえる模様
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err

		}
		return nil
	})

	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to close: %+v", err)
	}

	return eg.Wait()
}

const englishHelloPrefix = "Hello, "

// Hello テスト実行確認
func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + name
}
