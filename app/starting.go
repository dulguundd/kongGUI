package app

type ServiceConfig struct {
	address string
	port    string
}

var serviceconfig ServiceConfig

func getGWConnection() (addr string, port string) {
	addr = "172.30.52.207"
	port = "8000"
	return addr, port
}
