package main

import (
	"flag"
	log "github.com/golang/glog"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2"
	"net/http"
	"strings"
	"time"
)

func main() {
	flag.Parse()
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			if strings.HasPrefix(c.Request().RequestURI, "/health") || strings.HasPrefix(c.Request().RequestURI, "/metrics") {
				return true
			}
			return false
		},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	e.GET("/sub-path", func(c echo.Context) error {
		return c.String(http.StatusOK, "this is sub path")
	})
	seesion, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{"mongo-node-1.bs-db:27017,mongo-node-2.bs-db:27017,mongo-node-3.bs-db:27017"},
		Database: "bslot",
		Source:   "bslot",
		Username: "bslot",
		Password: "Addlink123!",
		Timeout:  time.Second * 3,
	})
	if err != nil {
		log.Fatalf("Mongodb init: %+v\n", err)
	}
	mgoinfo, err := seesion.BuildInfo()
	if err != nil {
		log.Fatalf("Mongodb get info: %+v\n", err)
	}
	log.Infof("Mongo INFO: %+v\n", mgoinfo)
	//e.Logger.Fatal(e.Start(":1323"))
}
