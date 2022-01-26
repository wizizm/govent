package handler

import (
	"github.com/wizizm/govent/example/service_a"
	"github.com/wizizm/govent/example/service_b"
)

/**
 * implements Consumer
 *
 */
type ConcreteConsumer1 struct {
}

func (c *ConcreteConsumer1) Execute(value string) {
	s := service_a.ServiceA{}
	_ = s.Update(value)
}

type ConcreteConsumer2 struct {
}

func (c *ConcreteConsumer2) Execute(value string) {
	s := service_b.ServiceB{}
	_ = s.Update(value)
}
