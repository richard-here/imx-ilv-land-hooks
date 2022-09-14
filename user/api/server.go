package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/richard-here/imx-ilv-land-hooks/user/api/controllers"
	seed "github.com/richard-here/imx-ilv-land-hooks/user/api/seeder"
)

var server = controllers.Server{}

func Run() {
	var err error
	err = godotenv.Load()

	if err != nil {
		log.Fatalf("Failed loading env: %v", err)
	} else {
		fmt.Println("Env values loaded")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seed.Load(server.DB)

	server.Run(":8080")
}
