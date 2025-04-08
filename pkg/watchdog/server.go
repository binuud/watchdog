package watchDogServer

import (
	"context"
	"fmt"

	protoV1 "github.com/binuud/watchdog/gen/go/v1/watchdog"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WatchDogGRPCServer struct {
	Service *WatchDogService
	protoV1.UnimplementedWatchDogServer
}

func NewWatchDogGRPCServer(fileName string) *WatchDogGRPCServer {
	service := NewWatchDogService(fileName)
	go func() {
		service.Watcher()
	}()

	return &WatchDogGRPCServer{
		Service: service,
	}

}

func (s *WatchDogGRPCServer) Get(ctx context.Context, in *protoV1.GetRequest) (*protoV1.GetResponse, error) {

	entry := s.Service.GetByNameOrUUID(in.Name, in.Uuid)
	if entry == nil {
		return nil, fmt.Errorf("no entry found for the domain")
	}
	return &protoV1.GetResponse{
		Summary: entry.Summary,
	}, nil

}

func (s *WatchDogGRPCServer) GetDetails(ctx context.Context, in *protoV1.GetDetailsRequest) (*protoV1.GetDetailsResponse, error) {

	entry := s.Service.GetByNameOrUUID(in.Name, in.Uuid)
	if entry == nil {
		return nil, fmt.Errorf("no entry found for the domain")
	}
	return &protoV1.GetDetailsResponse{
		Domain: entry,
	}, nil

}

func (s *WatchDogGRPCServer) ListSummaries(ctx context.Context, in *protoV1.ListSummariesRequest) (*protoV1.ListSummariesResponse, error) {

	if in.Page < 1 {
		in.Page = 1
	}

	fromIndex := (in.Page - 1) * in.PerPage
	toIndex := fromIndex + in.PerPage
	if toIndex > int64((len(s.Service.Data) - 1)) {
		toIndex = int64(len(s.Service.Data))
	}
	summaries := make([]*protoV1.DomainSummary, 0)

	logrus.Infof("Number of items %d Getting[%d:%d]", len(s.Service.Data), fromIndex, toIndex)

	// Check if indices are valid
	if fromIndex < 0 || fromIndex >= toIndex {
		return nil, fmt.Errorf("wrong indexes")
	}

	// Display elements from index i to j

	for i := fromIndex; i < toIndex; i++ {
		summaries = append(summaries, s.Service.Data[i].Summary)
	}

	return &protoV1.ListSummariesResponse{
		Page:      in.Page,
		PerPage:   in.PerPage,
		Summaries: summaries,
	}, nil

}

func (s *WatchDogGRPCServer) Health(ctx context.Context, in *protoV1.HealthRequest) (*protoV1.HealthResponse, error) {
	healthResponse := &protoV1.HealthResponse{
		NumDomains: int64(len(s.Service.Data)),
		CreatedAt:  timestamppb.Now().Seconds,
		Status:     protoV1.HealthResponse_Active,
	}
	return healthResponse, nil
}

func (s *WatchDogGRPCServer) Reload(ctx context.Context, in *protoV1.ReloadRequest) (*protoV1.ReloadResponse, error) {
	s.Service.CheckDomains()
	return &protoV1.ReloadResponse{}, nil
}
