package proxy

type Proxy struct {
	Ip       string
	Port     string
	Protocol string
	Latency  string

	Level    int
	UpdateAt string
	refer    string
}

func (p *Proxy) IsAvailable() bool {

	return false
}
