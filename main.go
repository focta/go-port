package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// パラメータ受け取り
	http.HandleFunc("/params", handleParams)

	// POSTのみ許可.
	http.HandleFunc("/onlyPost", handleOnlyPost)

	// POSTのみ許可.
	http.HandleFunc("/basicAuth", handleBasicAuth)

	// DBアクセスするパス
	// http.HandleFunc("/dbAccess", handleDbAccess)

	// 8080ポートで起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleOnlyPost(w http.ResponseWriter, r *http.Request) {

	// HTTPメソッドをチェック（POSTのみ許可）
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		w.Write([]byte("POSTだけだよー \n"))
		w.Write([]byte(fmt.Sprintf("userAgent: %s \n", r.UserAgent())))
		w.Write([]byte(fmt.Sprintf("userAgent: %s \n", r.UserAgent())))
		return
	}

	w.Write([]byte("OK"))
}

func handleParams(w http.ResponseWriter, r *http.Request) {

	// クエリパラメータ取得してみる
	fmt.Fprintf(w, "クエリ：%s\n", r.URL.RawQuery)

	// Bodyデータを扱う場合には、事前にパースを行う
	r.ParseForm()

	// Formデータを取得.
	form := r.PostForm
	fmt.Fprintf(w, "フォーム：\n%v\n", form)

	// または、クエリパラメータも含めて全部.
	params := r.Form
	fmt.Fprintf(w, "フォーム2：\n%v\n", params)
}

func handleBasicAuth(w http.ResponseWriter, r *http.Request) {

	// Basic認証のデータ取得.
	username, password, ok := r.BasicAuth()

	// そもそもそんなヘッダーがないなどのエラー.
	if !ok {
		w.Header().Set("WWW-Authenticate", `Basic realm="SECRET AREA"`)
		w.WriteHeader(http.StatusUnauthorized) // 401
		fmt.Fprintf(w, "%d Not authorized.", http.StatusUnauthorized)
		return
	}

	// Basic認証のヘッダーはあるけど、値が不正な場合.
	if username != "my" || password != "secret" {
		w.Header().Set("WWW-Authenticate", `Basic realm="SECRET AREA"`)
		w.WriteHeader(http.StatusUnauthorized) // 401
		fmt.Fprintf(w, "%d Not authorized.", http.StatusUnauthorized)
		return
	}

	// OK
	fmt.Fprint(w, "OK")
}

func Add(a, b int) int {
	return a + b
}

const englishHelloPrefix = "Hello, "

// Hello テスト実行確認
func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + name
}
