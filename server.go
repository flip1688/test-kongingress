package main

import (
	"fmt"
	log "github.com/golang/glog"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2"
	"net/http"
	"time"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	e.GET("/sub-path", func(c echo.Context) error {
		return c.String(http.StatusOK, "this is sub path")
	})
	e.Logger.Fatal(e.Start(":1323"))
	seesion, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{"mongo-node-1.bs-db"},
		Database: "bslot",
		Username: "bslot",
		Password: "Addlink123!",
		Timeout:  time.Second * 4,
	})
	if err != nil {
		log.Fatalf("Mongodb init: %+v\n", err)
	}
	mgoinfo, err := seesion.BuildInfo()
	if err != nil {
		log.Fatalf("Mongodb get info: %+v\n", err)
	}
	fmt.Printf("Mongo INFO: %+v\n", mgoinfo)
}
