package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/patra-go/configs"
	docs "github.com/rohanshrestha09/patra-go/docs"
	"github.com/rohanshrestha09/patra-go/ws"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {

	configs.LoadEnv()

	configs.InitializeFirebaseApp()

}

//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
func main() {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	server.SetTrustedProxies([]string{"localhost"})

	server.MaxMultipartMemory = 10 << 20

	hub := ws.NewHub()

	go hub.Run()

	server.GET("/ws/:initiator/:listener", func(ctx *gin.Context) {
		ws.Serve(hub, ctx)
	})

	router := server.Group("/api/v1")

	InitModule(router)

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "Server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := server.Run(":5000")

	if err != nil {
		log.Fatal(err)
	}
}
