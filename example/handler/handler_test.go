package handler

import (
	"github.com/stretchr/testify/assert"
	"github.com/wizizm/govent/example/service_a"
	"github.com/wizizm/govent/example/service_b"
	"testing"
)

func Test_Notify(t *testing.T) {
	a := &service_a.ServiceA{}
	var err error
	err = a.Create()
	assert.NoError(t, err)

	b := &service_b.ServiceB{}
	err = b.Create()
	assert.NoError(t, err)
}
