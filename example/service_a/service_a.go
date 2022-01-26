package service_a

import (
	"fmt"
	"github.com/wizizm/govent"
)

type ServiceA struct {
}

func (a *ServiceA) Create() error {
	fmt.Println("create！")
	govent.SimpleProducer.Push("topic1", "Test Create!")
	return nil
}

func (a *ServiceA) Update(value string) error {
	fmt.Println("receive msg:" + value)
	return nil
}
