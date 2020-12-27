package models

// "fmt"
// "net/http"

// Config of Database
type Config struct {
	AppPort string `yaml:"app_port" env:"PORT"`
	Host    string `yaml:"host" evb:"HOST"`
	// DBPort     string `yaml:"db_port" env:"DBPORT" env-default:"5432"`
	// DBHost     string `yaml:"db_host" env:"DBHOST" env-default:"localhost"`
	// DBName     string `yaml:"db_name" env:"DBNAME" env-default:"postgres"`
	// DBUser     string `yaml:"db_user" env:"DBUSER" env-default:"postgres"`
	// DBPassword string `yaml:"db_password" env:"DBPASSWORD"`
	Datapath      string `yaml:"datapath" env:"DATAPATH"`
	ReqTimeoutSec int    `yaml:"ReqTimeoutSec" env:"REQTIMEOUTSEC" env-default:"10"`
	GoogleAPIKey  string `yaml:"GoogleApiKey" env:"GoogleApiKey" `
}

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

type WasteFacility struct {
	Title      string
	Address    string
	WasteTypes map[string]struct{}
}
