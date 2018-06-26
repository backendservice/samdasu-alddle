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
		Email:       "hodol.kang@gmail.com",
		Departure:   "Seoul",
		Destination: "Osaka",
		Expense:     1000000,
		Duration:    4,
		FromDate:    "20180701",
		ToDate:      "20180831",
	}

	instance := new(server)
	_, err := instance.Register(ctx, req)
	assert.NoError(t, err)
}
