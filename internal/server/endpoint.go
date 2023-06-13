package server

type Endpoint struct {
	version string
	server  *Server
}

func NewEndpoint(version string, srv *Server) *Endpoint {
	return &Endpoint{
		version: version,
		server:  srv,
	}
}

func (e *Endpoint) Init() {

	// эндпоинт для добавления данных в очередь
	e.server.engine.PUT("/:key", e.server.controller.Append)
	// эндпоинт для чтения из очереди
	e.server.engine.GET("/:key", e.server.controller.Pull)

}
