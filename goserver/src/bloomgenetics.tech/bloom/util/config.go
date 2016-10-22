package util

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

var Config Configuration

type Configuration struct {
	HTMLRoot           string
	TemplateRoot       string
	DefaultPage        string
	IP                 string
	ApiPort            int64
	Port               int64
	TimeFmt            string
	ErrorFmt           string
	DatasourceFile     string
	DatabaseConnection DatabaseConnection
	LogLevel           string
}

func (c *Configuration) GetURL() string {
	return c.IP + ":" + strconv.FormatInt(c.Port, 10)
}

func (c *Configuration) GetApiURL() string {
	return c.IP + ":" + strconv.FormatInt(c.ApiPort, 10)
}

func (c Configuration) GetLogLevel() uint {
	return ConvertToLevelID(c.LogLevel)
}

func ConvertToLevelID(level string) uint {
	switch level {
	case "ERROR":
		return 7
	case "WARN":
		return 6
	case "INFO":
		return 5
	case "DEBUG":
		return 4
	}
	return 0
}

func (c Configuration) Pretty() string {
	return "Config:" +
		"\n\tHTMLRoot = " + c.HTMLRoot +
		"\n\tTemplateRoot = " + c.TemplateRoot +
		"\n\tDefaultPage = " + c.DefaultPage +
		"\n\tIP = " + c.IP +
		"\n\tApiPort = " + strconv.FormatInt(c.ApiPort, 10) +
		"\n\tPort = " + strconv.FormatInt(c.Port, 10) +
		"\n\tTimeFmt = " + c.TimeFmt +
		"\n\tErrorFmt = " + c.ErrorFmt +
		"\n\tDatasourceFile = " + c.DatasourceFile +
		"\n\tLogLevel = " + c.LogLevel
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
		PrintDebug(err.Error())
		return DatabaseConnection{}, err
	}

	var dbConfig DatabaseConnection
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&dbConfig)
	if err != nil {
		PrintDebug(err.Error())
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
			PrintDebug(err.Error())
			continue
		}
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&out)
		if err != nil {
			PrintDebug(err.Error())
			out = prechange
		}
	}
	out.ErrorFmt = strings.Replace(out.ErrorFmt, "%", "%%", -1)
	out.TimeFmt = strings.Replace(out.TimeFmt, "%", "%%", -1)
	return out, nil
}

func initConfig() {
	Config = Configuration{HTMLRoot: "html",
		TemplateRoot:   "templates",
		Port:           8080,
		ApiPort:        8081,
		DefaultPage:    "index.html",
		IP:             "",
		TimeFmt:        "2006 Jan 2 15:04:05",
		ErrorFmt:       "${level} ${time}:\t${msg}\n",
		DatasourceFile: "datasource.json"}
}
