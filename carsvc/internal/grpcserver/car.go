package grpcserver

import (
	"carsvc/internal/grpcapi"
	"context"
)

func (s *grpcserver) GetCarData(ctx context.Context, req *grpcapi.CarDataRequest) (*grpcapi.CarDataResponse, error) {

	cars := []*grpcapi.Car{
		{Make: "Honda", Model: "Civic", Color: "Blue"},
		{Make: "Mercedes", Model: "Cla", Color: "Black"},
		{Make: "Ford", Model: "Focus", Color: "Gray"},
	}

	res := grpcapi.CarDataResponse{
		User: req.User,
		Cars: cars,
	}

	return &res, nil
}
