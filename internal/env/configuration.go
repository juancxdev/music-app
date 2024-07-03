package env

import (
	"crypto/rsa"
	"encoding/json"
	"log"
	"os"
	"sync"
)

var (
	once   sync.Once
	config = &Configuration{}
)

type KeyStore struct {
	PrivateKey *rsa.PrivateKey `json:"private_key"`
	PublicKey  *rsa.PublicKey  `json:"public_key"`
}

type Configuration struct {
	App App `json:"app"`
	DB  DB  `json:"db"`
}

type App struct {
	ServiceName       string `json:"service_name"`
	Port              int    `json:"port"`
	AllowedDomains    string `json:"allowed_domains"`
	PathLog           string `json:"path_log"`
	LogReviewInterval int    `json:"log_review_interval"`
	LoggerHttp        bool   `json:"logger_http"`
	TLS               bool   `json:"tls"`
	Cert              string `json:"cert"`
	Key               string `json:"key"`
	SwaggerPath       string `json:"swagger_path"`
	RSAPrivateKey     string `json:"rsa_private_key"`
	RSAPublicKey      string `json:"rsa_public_key"`
}

type DB struct {
	Engine   string `json:"engine"`
	Server   string `json:"server"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Instance string `json:"instance"`
	IsSecure bool   `json:"is_secure"`
	SSLMode  string `json:"ssl_mode"`
}

func NewConfiguration() *Configuration {
	fromFile()
	return config
}

// LoadConfiguration lee el archivo configuration.json
// y lo carga en un objeto de la estructura Configuration
func fromFile() {
	once.Do(func() {
		b, err := os.ReadFile("config.json")
		if err != nil {
			log.Fatalf("no se pudo leer el archivo de configuración: %s", err.Error())
		}

		err = json.Unmarshal(b, config)
		if err != nil {
			log.Fatalf("no se pudo parsear el archivo de configuración: %s", err.Error())
		}

		if config.DB.Engine == "" {
			log.Fatal("no se ha cargado la información de configuración")
		}
	})
}
