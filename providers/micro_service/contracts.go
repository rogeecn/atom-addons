package micro_service

type Service interface {
	Serve() error
	GetEngine() any
}
