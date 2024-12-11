package commands

import (
	"context"
	"fin_api_gateway/db"
	"fin_api_gateway/internal/config"
	"fin_api_gateway/internal/entities"
	"fin_api_gateway/internal/log"
	pb "fin_api_gateway/pkg/grpc"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type tickerServer struct {
	pb.TickersServiceServer
}

type targetServer struct {
	pb.TargetsServiceServer
}

func (s *tickerServer) GetMultipleTickers(ctx context.Context, in *pb.TickersRequest) (*pb.MultipleTickerResponse, error) {
	gDB := &db.Connection{}
	if err := gDB.Connect(); err != nil {
		log.Error("Could not connect to database: ", err)
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			log.Error("Error closing database connection: ", err)
		}
	}()
	if err := gDB.Connect(); err != nil {
		log.Error("Could not connect to database: ", err)
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			log.Error("Error closing database connection: ", err)
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
	gDB := &db.Connection{}
	if err := gDB.Connect(); err != nil {
		log.Error("Could not connect to database: ", err)
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			log.Error("Error closing database connection: ", err)
		}
	}()

	var userTargets []struct {
		entities.UserTarget
		entities.UserResponse
		entities.TgUser
	}

	q := gDB.Table("user_targets").
		Where("achieved = false").
		Select("user_targets.*, users.*, tg_users.*").
		Joins("LEFT JOIN users ON users.id = user_targets.user_id " +
			"LEFT JOIN tg_users ON tg_users.id = user_targets.tg_user_id")

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
				Telegram: fmt.Sprintf("%d", target.TgUser.TelegramUserID),
			},
		})
	}

	return response, nil
}

func (s *targetServer) SetTargetAchieved(ctx context.Context, in *pb.TargetAchievedRequest) (*pb.TargetItem, error) {
	gDB := &db.Connection{}
	if err := gDB.Connect(); err != nil {
		log.Error("Could not connect to database: ", err)
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			log.Error("Error closing database connection: ", err)
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
			Id:    target.UserResponse.ID,
			Name:  target.UserResponse.Name,
			Email: target.UserResponse.Email,
		},
	}, nil
}

func RunGRPCServer(ctx context.Context, cfg *config.Config) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.GrpcHost, cfg.GrpcPort))
	if err != nil {
		log.Error("failed to listen: ", err)
	}

	s := grpc.NewServer()
	pb.RegisterTickersServiceServer(s, &tickerServer{})
	pb.RegisterTargetsServiceServer(s, &targetServer{})

	go func() {
		err := s.Serve(lis)
		if err != nil {
			log.Error("failed to serve: ", err)
			return
		}
	}()

	<-ctx.Done()
	s.GracefulStop()
	log.Info("GRPC server stopped")

}
