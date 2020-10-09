package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/sparkidea25/helium/client"
	"github.com/sparkidea25/helium/db"
)

func initDB() {
	db.Init()
	db.Get().AutoMigrate(&client.Client{})
}

func main() {

	e := echo.New()
	e.GET("/", index)
	e.Logger.Fatal(e.Start(":3000"))
}

func index(c echo.Context) error {
	cli := client.Create("Jasper")
	fmt.Println(cli)
	return c.String(http.StatusOK, "Hello World")
}
