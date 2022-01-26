package service_b

import (
	"fmt"
	"github.com/wizizm/govent"
)

type ServiceB struct {
}

func (a *ServiceB) Update(value string) error {
	fmt.Println("receive msg:" + value)
	return nil
}

func (a *ServiceB) Create() error {
	fmt.Println("createB！")
	govent.SimpleProducer.Push("topic2", "Test CreateB!")
	return nil
}
