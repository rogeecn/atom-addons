package http

type Route interface{}

type Service interface {
	Serve() error
	GetEngine() interface{}
}
