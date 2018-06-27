package main

import (
	"context"
	"testing"

	pb "github.com/backendservice/samdasu-alddle"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	ctx := context.Background()
	req := &pb.RegisterRequest{
		Id:          "A",
		Email:       "hodol.kang@gmail.com",
		Departure:   "Seoul",
		Destination: "Osaka",
		Expense:     1000000,
		Duration:    4,
		FromDate:    "20180701",
		ToDate:      "20180831",
	}

	instance := new(server)
	assert.Equal(t, 0, len(registeredList))
	_, err1 := instance.Register(ctx, req)
	_, err2 := instance.Register(ctx, req)
	assert.Equal(t, 1, len(registeredList))
	assert.NoError(t, err1)
	assert.NoError(t, err2)
}
func TestUnregister(t *testing.T) {
	ctx := context.Background()
	instance := new(server)
	instance.Register(ctx, &pb.RegisterRequest{Id: "A"})
	instance.Register(ctx, &pb.RegisterRequest{Id: "B"})
	assert.Equal(t, 2, len(registeredList))
	instance.Unregister(ctx, &pb.UnregisterRequest{RegisterId: "A"})
	instance.Unregister(ctx, &pb.UnregisterRequest{RegisterId: "A"})
	assert.Equal(t, 1, len(registeredList))

}
