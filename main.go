package main

import (
	"apirest/api"
	"apirest/internal/env"
)

// @title Comfandi API
// @version 1.0
// @description Api para consumir los provedores de información core comfandi, datacretido, sarlaf inspector, desceval y acuario
// @contact.name API Support
// @contact.email info@bjungle.net
// @license.name Software Owner
// @license.url https://www.bjungle.net/terms/licenses
// @host http://127.0.0.1:50050
// @tag.name Core Comfanidi
// @tag.description Metodos del core de Comfandi
// @tag.name DataCredito
// @tag.description Metodos del cliente DataCredito
// @tag.name Sarlaf Inspektor
// @tag.description Metodos del cliente Sarlaf Inspector
// @tag.Name Desceval
// @tag.description Metodos del cliente Desceval
// @tag.Name Comfandi Acuario
// @tag.description Metodos of de Comfandi Acuario
// @BasePath /
func main() {
	c := env.NewConfiguration()

	api.Start(c.App.Port, c.App.ServiceName, c.App.LoggerHttp, c.App.AllowedDomains)
}
