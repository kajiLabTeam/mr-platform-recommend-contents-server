package router

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/mr-platform-recommend-contents-server/controller"
	_ "github.com/lib/pq"
)

func Init() {
	f, _ := os.Create("../log/server.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	db, err := sql.Open("postgres", "host=postgres port="+port+" user="+user+" password="+password+" dbname="+dbname+" sslmode=disable")

	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %q", err)
	}

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!!")
	})
	r.POST("/api/recommend/contents", controller.RecommendContents)

	// サーバーの起動状態を表示しながら、ポート8084でサーバーを起動する
	if err := r.Run("0.0.0.0:8000"); err != nil {
		fmt.Println("サーバーの起動に失敗しました:", err)
	} else {
		fmt.Println("サーバーが正常に起動しました")
	}
}
