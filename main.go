package main

import (
	"challenge-chapter-3-sesi-2/config"
	"challenge-chapter-3-sesi-2/routers"
)

func main() {
	config.ConnectDB()

	routers.StartServer().Run(":80")
}
