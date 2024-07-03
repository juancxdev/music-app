package main

import (
	"music-app/api"
	_ "music-app/docs"
	"music-app/internal/env"
)

// @title Music API REST
// @version 1.0
// @description System manage music
// @termsOfService https://juancx.site/
// @contact.name API Support
// @contact.email juanm.campos@unas.edu.pe
// @host 127.0.0.1:4021
// @BasePath /
func main() {
	e := env.NewConfiguration()
	api.Start(e.App.Port, e.App.ServiceName, e.App.LoggerHttp, e.App.AllowedDomains)
}
