package micro_service

const DefaultPrefix = "Micro"

type Config struct {
	Name    string
	Version string
	Port    uint
	Tls     *Tls
}

type Tls struct {
	Cert string
	Key  string
}
