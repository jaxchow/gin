package main

import (
	cli "github.com/jaxchow/gin/cmd/cli"
)

// @title Planner Task API
// @version 1.0
// @description This is a Planner Task server celler server.
// @termsOfService https://www.topgoer.com

// @contact.name www.topgoer.com
// @contact.url https://www.topgoer.com
// @contact.email jaxchow@gmail

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api/v1

func main() {
	cli.Run()
}
