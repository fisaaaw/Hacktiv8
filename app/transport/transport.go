package transport

import "swag-tutor/app/interface/container"

type Tp struct {
	Transport *tp
}

func SetupTransport(container container.Container) *Tp {
	transport := NewTransport(container.Usecase)
	return &Tp{
		Transport: transport,
	}
}