package main

import (
	"crud_golang_tuto_1/config"
	"crud_golang_tuto_1/controllers"
	"crud_golang_tuto_1/routes"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//const (
//	host     = "127.0.0.1"
//	port     = 3306
//	user     = "henry"
//	password = "11nhri04p"
//	dbname   = "base1"
//	sslmode  = "disable"
//)

const (
	Port = ":8090"
)

//	func CheckError(err error) {
//		if err != nil {
//			panic(err)
//		}
//	}

func main() {
	// On relie le fichier css
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	templateCache, err := controllers.CreateTemplateCache()
	if err != nil {
		panic(err)
	}
	var appConfig config.Config
	appConfig.TemplateCache = templateCache
	appConfig.Port = ":8070"
	controllers.CreateTemplates(&appConfig)

	// Définir les urls de notre application web
	routes.Web()
	fmt.Printf("http://localhost%v , Cliquez sur le lien pour lancer le navigateur", Port)
	http.ListenAndServe(Port, nil)
}
