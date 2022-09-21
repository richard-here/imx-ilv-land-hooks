package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/richard-here/imx-ilv-land-hooks/user/models"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DbUser, DbPassword, DbHost, DbPort, DbName)
	server.DB, err = gorm.Open(DbDriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", DbDriver)
		log.Fatal("Error: ", err)
	} else {
		fmt.Printf("Connected to %s database", DbDriver)
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Subscription{})

	server.DB.Exec("GRANT SELECT, RELOAD, SHOW DATABASES, REPLICATION SLAVE, " +
		"REPLICATION CLIENT ON *.* TO 'richard'")

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
