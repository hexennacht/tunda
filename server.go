package tunda

type handler struct {
	redis RedisRepository
}

type TundaServer interface {
	AddHandler() TundaServer
}

func NewTundaServer() TundaServer {
	return &handler{redis: NewRedisRepository()}
}

type Server interface {
	Serve()
}

