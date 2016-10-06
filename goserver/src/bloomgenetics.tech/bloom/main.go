package main

import (
	//"encoding/json"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" //Required for Postgres
	//"html/template"
	"net/http"
	"bloomgenetics.tech/bloom/util"
)

var config *util.Configuration

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
	util.Database, err = sql.Open("postgres", config.DatabaseConnection.URL)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println(util.Config.Pretty())
	router := NewRouter("/v1")
	http.ListenAndServe(config.GetURL(), router)
}
