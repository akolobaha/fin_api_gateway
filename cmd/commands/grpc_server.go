package commands

import (
	"context"
	"fin_api_gateway/db"
	"fin_api_gateway/internal/config"
	"fin_api_gateway/internal/entities"
	pb "fin_api_gateway/pkg/grpc"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net"
)

type tickerServer struct {
	pb.TickersServiceServer
}

type targetServer struct {
	pb.TargetsServiceServer
}

func (s *tickerServer) GetMultipleTickers(ctx context.Context, in *pb.TickersRequest) (*pb.MultipleTickerResponse, error) {
	gDB := new(db.GormDB).Connect()
	securities := entities.Securities{}
	gDB.Find(&securities)

	var tickersResponse []*pb.TickersResponse
	for _, sec := range securities {
		tickersResponse = append(tickersResponse, &pb.TickersResponse{
			Ticker:    sec.Ticker,
			Shortname: sec.Shortname,
			Name:      sec.Secname,
			Exists:    true,
		})
	}

	return &pb.MultipleTickerResponse{Tickers: tickersResponse}, nil
}

func (s *targetServer) GetTargets(ctx context.Context, in *pb.TargetRequest) (*pb.TargetResponse, error) {
	gDB := new(db.GormDB).Connect()
	var userTargets []struct {
		entities.UserTarget
		entities.UserResponse
	}

	err := gDB.Table("user_targets").
		Where("achieved = false").
		Select("user_targets.*, users.*").
		Joins("INNER JOIN users ON users.id = user_targets.user_id").
		Scan(&userTargets).Error

	if err != nil {
		return nil, err
	}

	response := &pb.TargetResponse{}
	for _, target := range userTargets {
		response.Targets = append(response.Targets, &pb.TargetItem{
			Id:              int64(target.UserTarget.ID),
			Ticker:          target.UserTarget.Ticker,
			ValuationRatio:  target.UserTarget.ValuationRatio,
			Value:           target.UserTarget.Value,
			FinancialReport: target.UserTarget.FinancialReport,
			User: &pb.User{
				Id:       target.UserId,
				Name:     target.Name,
				Email:    target.Email,
				Telegram: target.Telegram,
			},
		})
	}

	return response, nil
}

func RunGRPCServer(ctx context.Context, cfg *config.Config) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.GrpcHost, cfg.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTickersServiceServer(s, &tickerServer{})
	pb.RegisterTargetsServiceServer(s, &targetServer{})

	go func() {
		slog.Info(fmt.Sprintf("Server is running on port: %s", cfg.GrpcPort))

		if err := s.Serve(lis); err != nil {
			slog.Error("failed to serve: ", err.Error())
		}
	}()

	slog.Info("Сервер grpc запущен")
}
