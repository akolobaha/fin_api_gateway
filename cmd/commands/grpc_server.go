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

type server struct {
	pb.TickersServiceServer
}

func (s *server) GetMultipleTickers(ctx context.Context, in *pb.TickersRequest) (*pb.MultipleTickerResponse, error) {
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

func RunGRPCServer(ctx context.Context, cfg *config.Config) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.GrpcHost, cfg.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTickersServiceServer(s, &server{})

	go func() {
		slog.Info(fmt.Sprintf("Server is running on port: %s", cfg.GrpcPort))

		if err := s.Serve(lis); err != nil {
			slog.Error("failed to serve: ", err.Error())
		}
	}()

	slog.Info("Сервер grpc запущен")
}
