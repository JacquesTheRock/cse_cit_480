package util

import (
	"strconv"
	"strings"
	"os"
	"encoding/json"
)

var Config Configuration

type Configuration struct {
	HTMLRoot string
	TemplateRoot string
	DefaultPage string
	IP   string
	Port int64
	TimeFmt string
	ErrorFmt string
	DatasourceFile string
	DatabaseConnection DatabaseConnection
}

func (c *Configuration)GetURL() (string) {
	return c.IP + ":" + strconv.FormatInt(c.Port,10)
}

func (c Configuration)Pretty() (string) {
	return "Config:" +
		"\n\tHTMLRoot = " + c.HTMLRoot +
		"\n\tTemplateRoot = " + c.TemplateRoot +
		"\n\tDefaultPage = " + c.DefaultPage +
		"\n\tIP = " + c.IP +
		"\n\tPort = " + strconv.FormatInt(c.Port,10) +
		"\n\tTimeFmt = " + c.TimeFmt +
		"\n\tErrorFmt = " + c.ErrorFmt +
		"\n\tDatasourceFile = " + c.DatasourceFile;
}

type DatabaseConnection struct {
	Host     string
	Port     int32
	Username string
	Password string
	Dbname   string
	URL      string
}

func GetDatabaseConnectionInfo(filename string) (DatabaseConnection, error) {
	file, err := os.Open(filename)
	if err != nil {
		PrintError(err.Error())
		return DatabaseConnection{}, err
	}

	var dbConfig DatabaseConnection
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&dbConfig)
	if err != nil {
		PrintError(err.Error())
		return DatabaseConnection{}, err
	}
	return dbConfig, nil
}


func ReadConfigurationInfo(filenames []string) (Configuration, error) {
	out := Config

	for i := 0; i < len(filenames); i++ {
		prechange := out
		file, err := os.Open(filenames[i])
		if err != nil {
			PrintError(err.Error())
			continue
		}
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&out)
		if err != nil {
			PrintError(err.Error())
			out = prechange
		}
	}
	out.ErrorFmt = strings.Replace(out.ErrorFmt, "%", "%%",-1)
	out.TimeFmt = strings.Replace(out.TimeFmt, "%", "%%",-1)
	return out,nil
}

func initConfig() {
	Config = Configuration{ HTMLRoot: "html",
		TemplateRoot: "templates",
		Port: 8080,
		DefaultPage: "index.html",
		IP: "",
		TimeFmt: "2006 Jan 2 15:04:05",
		ErrorFmt: "${level} ${time}:\t${msg}\n",
		DatasourceFile: "datasource.json" }
}
