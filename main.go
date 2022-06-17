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
	"os"
	"strconv"
)

func main() {
	// コマンドラインインターフェースでの実行
	flag.Parse()
	// arg := flag.Arg(0)
	args := flag.Args()
	fmt.Printf("Hello %s\n", args[0])

	if len(args) == 0 {
		return
	} else {
		for i := range args {
			fmt.Printf("key: %s \t", strconv.Itoa(i))
			fmt.Printf("value: %s \n", args[i])
		}
	}

	// TODO
	_, err := os.Create("./hello.txt")
	if err != nil {
		fmt.Printf("failed to create file \n: %v", err)
		return
	}

	fmt.Println("Done!")
}
