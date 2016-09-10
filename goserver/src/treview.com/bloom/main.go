package main

import (
	"database/sql"
	//"encoding/json"
	"fmt"
	_ "github.com/lib/pq" //Required for Postgres
	"html/template"
	"net/http"
	"path/filepath"
	"treview.com/bloom/util"
	"strings"
	"io/ioutil"
)

var config *util.Configuration
var database *sql.DB

type PageMeta struct {
	Title  string
	Author string
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	reqPath := strings.Split(r.URL.Path, "/")
	parsedPath := ""
	relocate :=  false
	for index,part := range reqPath {
		if(part != "") {
			parsedPath = filepath.Join(parsedPath,part)
		}
		if index == len(reqPath) - 1 {
			if strings.HasSuffix(part,".html") ||
				strings.HasSuffix(part,".css") ||
				strings.HasSuffix(part,".js") ||
				strings.HasSuffix(part,".png") ||
				strings.HasSuffix(part,".json") ||
				strings.HasSuffix(part, ".ico"){
				continue
			}
			parsedPath =  filepath.Join(parsedPath,config.DefaultPage)
			relocate = true;
		}
	}
	parsedPath = filepath.Clean(parsedPath)
	file := filepath.Clean(config.HTMLRoot + "/" + parsedPath)
	safe := strings.HasPrefix(file,config.HTMLRoot)
	if safe {
		if relocate {
			w.Header().Set("Location", "/" + strings.Replace(parsedPath,"\\","/",-1))
			w.WriteHeader(http.StatusTemporaryRedirect)
			return
		}
		data, err := ioutil.ReadFile(file)
		if err != nil {
			util.PrintError(err.Error())
			fmt.Fprintf(w,"<html>%s</html>\n","404 Page not found")
			return
		}
		fmt.Fprintf(w,"%s",data)
		return
	}
	fmt.Fprintf(w,"<html>%s</html>\n","404 Page not found: Invalid Path")
}

func apiReference(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(config.HTMLRoot + "/api.html")
	t.Execute(w, nil)
}

func init() {
	configFiles := make([]string, 0)
	configFiles = append(configFiles, "./config.json")
	config = &util.Config
	(*config), _ = util.ReadConfigurationInfo(configFiles)
	var err error
	if config.DatasourceFile != "" {
		config.DatabaseConnection, err = util.GetDatabaseConnectionInfo(config.DatasourceFile)
	}
	if err != nil {
		panic(fmt.Sprintf("Cannot read database info located at: %s", config.DatasourceFile))
	}
	if config.DatabaseConnection.URL == "" {
		config.DatabaseConnection.URL = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", config.DatabaseConnection.Username, config.DatabaseConnection.Password, config.DatabaseConnection.Host, config.DatabaseConnection.Dbname)
		fmt.Printf("No URL Supplied, Deriving URL from parameters. %s\n", config.DatabaseConnection)
	} else {
		fmt.Printf("Database Using URL value, ignoring others\n")
	}
	database, err = sql.Open("postgres", config.DatabaseConnection.URL)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println(util.Config.Pretty())
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/api/", apiReference)
	http.ListenAndServe(config.GetURL(), nil)
}
