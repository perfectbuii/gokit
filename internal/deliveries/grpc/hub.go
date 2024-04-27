package deliveries

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/perfectbuii/gokit/internal/models"
	"github.com/perfectbuii/gokit/internal/services"
	"github.com/perfectbuii/gokit/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type hubDelivery struct {
	hubService services.HubService
	pb.UnimplementedHubServiceServer
}

func NewHubDelivery(hubService services.HubService) pb.HubServiceServer {
	return &hubDelivery{
		hubService: hubService,
	}
}

func (d *hubDelivery) CreateHub(ctx context.Context, req *pb.CreateHubRequest) (*pb.CreateHubResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Errorf("validate failed: %w", err).Error())
	}
	if err := d.hubService.Create(ctx, &models.Hub{
		ID:         req.Hub.Id,
		Name:       req.Hub.Name,
		LocationID: req.Hub.LocationId,
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		DeletedAt: sql.NullTime{},
	}); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Errorf("Create: %v", err).Error())
	}
	return &pb.CreateHubResponse{
		Success: true,
	}, nil
}

func (d *hubDelivery) UpdateHub(ctx context.Context, req *pb.UpdateHubRequest) (*pb.UpdateHubResponse, error) {
	if err := d.hubService.Update(ctx, req.Hub.Id, nil); err != nil {
		return nil, err
	}
	return &pb.UpdateHubResponse{
		Success: true,
	}, nil
}

func (d *hubDelivery) GetList(ctx context.Context, req *pb.GetListHubRequest) (*pb.GetListHubResponse, error) {
	_, err := d.hubService.GetList(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	return &pb.GetListHubResponse{
		Data: nil,
	}, nil
}

func (d *hubDelivery) GetHubByID(ctx context.Context, req *pb.GetHubByIDRequest) (*pb.GetHubByIDResponse, error) {
	hub, err := d.hubService.GetByID(ctx, req.GetHubID())
	if err != nil {
		return nil, err
	}
	return &pb.GetHubByIDResponse{
		Data: &pb.Hub{
			Id:         hub.ID,
			Name:       hub.Name,
			LocationId: hub.LocationID,
		},
	}, nil
}
