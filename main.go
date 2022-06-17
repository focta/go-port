// package main

// import (
// 	"log"

// 	"github.com/gin-gonic/gin"
// )

// func main() {

// 	log.Println("start server...")
// 	r := gin.Default()

// 	r.GET("/hello", func(context *gin.Context) {
// 		context.JSON(200, gin.H{
// 			"message": "Hello World!",
// 		})
// 	})

// 	r.GET("/test", func(context *gin.Context) {
// 		context.JSON(200, gin.H{
// 			"message": "yeah!",
// 		})
// 	})

// 	log.Fatal(r.Run())
// }
package main

import (
	"flag" // 追加する
	"fmt"
	"strconv"
)

func main() {
	// コマンドラインインターフェースでの実行
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		return
	} else {
		fmt.Printf("Hello %s\n", args[0])
		for i := range args {
			fmt.Printf("key: %s \t", strconv.Itoa(i))
			fmt.Printf("value: %s \n", args[i])
		}
	}

	// ファイル作成と実行
	arg := flag.Arg(0)

	msg := fmt.Sprintf("Hello %s\n", arg)

	fileWriter(msg)
}
