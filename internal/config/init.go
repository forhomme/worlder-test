package config

import (
	"log"
	"os"
	"path"
	"runtime"
	"strconv"

	"github.com/joho/godotenv"
)

type Configuration struct {
	Port       string   `default:":9270"`
	RootURL    string   `split_words:"true" default:"/engine"`
	LogLevel   string   `split_words:"true" default:"debug"`
	VersionApp string   `split_words:"true" default:"1.0.0"`
	NameApp    string   `split_words:"true" default:"API-Worlder"`
	Database   Database `json:"database"`
	Mqtt       Mqtt     `json:"mqtt"`
}

type Database struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	SSLMode  string `json:"sslmode"`
	LogMode  bool   `json:"logMode"`
	Schema   string `json:"schema"`
}

type Mqtt struct {
	Client   string `json:"client"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var Config Configuration

func init() {
	_, file, _, _ := runtime.Caller(0)
	rootPath := path.Join(file, "..", "..", "..")
	log.Println("Path Env:", rootPath)

	if err := godotenv.Load(rootPath + "/.env"); err != nil {
		log.Fatal("error: failed to load the env file>", err)
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		dbPort = 5432 // default postgres value
	}

	dbLog, err := strconv.ParseBool(os.Getenv("DB_LOGMODE"))
	if err != nil {
		dbLog = true // default log mode
	}

	mqttPort, err := strconv.Atoi(os.Getenv("MQTT_PORT"))
	if err != nil {
		mqttPort = 1883 // default mqtt value
	}

	DatabaseData := Database{
		Name:     os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Port:     dbPort,
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		LogMode:  dbLog,
		Schema:   os.Getenv("DB_SCHEMA"),
	}

	MqttData := Mqtt{
		Client:   os.Getenv("MQTT_CLIENT"),
		Host:     os.Getenv("MQTT_HOST"),
		Port:     mqttPort,
		Username: os.Getenv("MQTT_USERNAME"),
		Password: os.Getenv("MQTT_PASSWORD"),
	}

	Config = Configuration{
		Port:       os.Getenv("PORT"),
		RootURL:    os.Getenv("ROOT_URL"),
		LogLevel:   os.Getenv("LOG_LEVEL"),
		VersionApp: os.Getenv("VERSION_APP"),
		NameApp:    os.Getenv("NAME_APP"),
		Database:   DatabaseData,
		Mqtt:       MqttData,
	}

	log.Println("Config: ", Config)
}
