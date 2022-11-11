package api

func Start(port int, app string, loggerHttp bool, allowedOrigins string) {

	r := routes(loggerHttp, allowedOrigins)
	server := newServer(port, app, r)
	server.Start()
}
