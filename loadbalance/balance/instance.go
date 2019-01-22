package balance

type Instnace struct {
	host string
	port int
}

func NewInstance(host string, port int) *Instnace {
	return &Instnace{
		host: host,
		port: port,
	}
}

func (this *Instnace) GetHost() string {
	return this.host
}

func (this *Instnace) GetPort() int {
	return this.port
}
