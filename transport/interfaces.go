package transport

type Transporter interface {
	Send(data []byte, endpoint string) error
}
