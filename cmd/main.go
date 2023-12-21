package main

import (
	"fmt"
	zlog "go_crud/common/logger"
	"go_crud/internal/server"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var logger = zlog.CreateLogger("mainApp")

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}
	G := server.CreateServer()
	addressPort := strings.Join([]string{os.Getenv("APP_HOST"), os.Getenv("APP_PORT")}, ":")
	logger.LogInfo("⚡️ Running server on %s\n" + addressPort)
	G.Run(addressPort)

}
