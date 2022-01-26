package handler

import (
	"github.com/stretchr/testify/assert"
	"github.com/wizizm/govent"
	"github.com/wizizm/govent/example/service_a"
	"github.com/wizizm/govent/example/service_b"
	"testing"
)

func Test_Notify(t *testing.T) {
	observer1 := new(ConcreteConsumer1)
	observer2 := new(ConcreteConsumer2)
	govent.SimpleProducer.Register("topic1", observer1)
	govent.SimpleProducer.Register("topic2", observer2)

	a := &service_a.ServiceA{}
	var err error
	err = a.Create()
	assert.NoError(t, err)

	b := &service_b.ServiceB{}
	err = b.Create()
	assert.NoError(t, err)

	govent.SimpleProducer.Notify()
}
