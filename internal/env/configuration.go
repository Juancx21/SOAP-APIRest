package env

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

var (
	once   sync.Once
	config = &configuration{}
)

type configuration struct {
	App   App         `json:"app"`
	SoapS SoapService `json:"soap_s"`
}

type App struct {
	ServiceName       string `json:"service_name"`
	Port              int    `json:"port"`
	AllowedDomains    string `json:"allowed_domains"`
	PathLog           string `json:"path_log"`
	PathXml           string `json:"path_xml"`
	LogReviewInterval int    `json:"log_review_interval"`
	RegisterLog       bool   `json:"register_log"`
	RSAPrivateKey     string `json:"rsa_private_key"`
	RSAPublicKey      string `json:"rsa_public_key"`
	LoggerHttp        bool   `json:"logger_http"`
}

type SoapService struct {
	BaseUrl string `json:"base_url"`
}

func NewConfiguration() *configuration {
	fromFile()
	return config
}

// LoadConfiguration lee el archivo configuration.json
// y lo carga en un objeto de la estructura Configuration
func fromFile() {
	once.Do(func() {
		b, err := ioutil.ReadFile("config.json")
		if err != nil {
			log.Fatalf("no se pudo leer el archivo de configuración: %s", err.Error())
		}

		err = json.Unmarshal(b, config)
		if err != nil {
			log.Fatalf("no se pudo parsear el archivo de configuración: %s", err.Error())
		}
	})
}
