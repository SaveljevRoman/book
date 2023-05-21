package main

import (
	"book/pkg/books"
	"book/pkg/common/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	_ = viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)
	db.Init(dbUrl)

	r := gin.Default()
	h := db.Init(dbUrl)

	books.RegisterRoutes(r, h)

	_ = r.Run(port)
}
