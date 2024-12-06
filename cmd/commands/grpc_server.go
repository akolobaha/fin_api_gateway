package commands

import (
	"context"
	"fin_api_gateway/db"
	"fin_api_gateway/internal/config"
	"fin_api_gateway/internal/entities"
	pb "fin_api_gateway/pkg/grpc"
	"fmt"
	"google.golang.org/grpc"
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
	gDB := &db.GormDB{}
	if err := gDB.Connect(); err != nil {
		slog.Error("Could not connect to database: ", "error", err.Error())
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			slog.Error("Error closing database connection: ", "error", err.Error())
		}
	}()
	if err := gDB.Connect(); err != nil {
		slog.Error("Could not connect to database: ", "error", err.Error())
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			slog.Error("Error closing database connection: ", "error", err.Error())
		}
	}()

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
	gDB := &db.GormDB{}
	if err := gDB.Connect(); err != nil {
		slog.Error("Could not connect to database: ", "error", err.Error())
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			slog.Error("Error closing database connection: ", "error", err.Error())
		}
	}()

	var userTargets []struct {
		entities.UserTarget
		entities.UserResponse
	}

	q := gDB.Table("user_targets").
		Where("achieved = false").
		Select("user_targets.*, users.*").
		Joins("INNER JOIN users ON users.id = user_targets.user_id")

	if in.GetTicker() != "" {
		q.Where("user_targets.ticker = ?", in.Ticker)
	}

	err := q.Scan(&userTargets).Error

	if err != nil {
		return nil, err
	}

	response := &pb.TargetResponse{}
	for _, target := range userTargets {
		response.Targets = append(response.Targets, &pb.TargetItem{
			Id:                 int64(target.UserTarget.ID),
			Ticker:             target.UserTarget.Ticker,
			ValuationRatio:     target.UserTarget.ValuationRatio,
			Value:              target.UserTarget.Value,
			FinancialReport:    target.UserTarget.FinancialReport,
			NotificationMethod: target.UserTarget.NotificationMethod,
			User: &pb.User{
				Id:       target.UserResponse.ID,
				Name:     target.UserResponse.Name,
				Email:    target.UserResponse.Email,
				Telegram: target.UserResponse.Telegram,
			},
		})
	}

	return response, nil
}

func (s *targetServer) SetTargetAchieved(ctx context.Context, in *pb.TargetAchievedRequest) (*pb.TargetItem, error) {
	gDB := &db.GormDB{}
	if err := gDB.Connect(); err != nil {
		slog.Error("Could not connect to database: ", "error", err.Error())
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			slog.Error("Error closing database connection: ", "error", err.Error())
		}
	}()
	err := gDB.Model(&entities.UserTarget{}).Where("id = ?", in.GetId()).Update("achieved", in.GetAchieved()).Error
	if err != nil {
		return nil, err
	}

	var target struct {
		entities.UserTarget
		entities.UserResponse
	}

	err = gDB.First(&entities.UserTarget{ID: in.GetId()}).
		Select("user_targets.*, users.*").
		Joins("INNER JOIN users ON users.id = user_targets.user_id").
		Scan(&target).Error

	if err != nil {
		return nil, err
	}

	return &pb.TargetItem{
		Id:                 int64(target.UserTarget.ID),
		Ticker:             target.UserTarget.Ticker,
		ValuationRatio:     target.UserTarget.ValuationRatio,
		Value:              target.UserTarget.Value,
		FinancialReport:    target.UserTarget.FinancialReport,
		NotificationMethod: target.UserTarget.NotificationMethod,
		Achieved:           target.Achieved,
		User: &pb.User{
			Id:       target.UserResponse.ID,
			Name:     target.UserResponse.Name,
			Email:    target.UserResponse.Email,
			Telegram: target.UserResponse.Telegram,
		},
	}, nil
}

func RunGRPCServer(ctx context.Context, cfg *config.Config) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.GrpcHost, cfg.GrpcPort))
	if err != nil {
		slog.Error("failed to listen: %v", "error", err)
	}

	s := grpc.NewServer()
	pb.RegisterTickersServiceServer(s, &tickerServer{})
	pb.RegisterTargetsServiceServer(s, &targetServer{})

	go func() {
		slog.Info(fmt.Sprintf("Server is running on port: %s", cfg.GrpcPort))

		if err := s.Serve(lis); err != nil {
			slog.Error("failed to serve: ", "error", err.Error())
		}
	}()

	slog.Info("Сервер grpc запущен")
}
