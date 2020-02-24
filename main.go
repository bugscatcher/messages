package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"github.com/bugscatcher/messages/application"
	"github.com/bugscatcher/messages/config"
	"github.com/bugscatcher/messages/grpcHelper"
	msg "github.com/bugscatcher/messages/server/grpc/messages"
	"github.com/bugscatcher/messages/services"
	"net"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("Read config")
	}

	app, err := application.New(&conf)
	if err != nil {
		log.Fatal().Err(err).Msg("Create app")
	}

	defer app.Close()

	go startGRPCServer(&conf.PublicGRPCServer, app)

}

func startGRPCServer(grpcConf *grpcHelper.ServerConf, app *application.App) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcConf.Port))
	if err != nil {
		log.Fatal().Err(err).Msgf("Listen GRPC port: %s", grpcConf.Addr())
	}
	s := grpc.NewServer()

	handler := msg.New(app)
	services.RegisterMessagesServiceServer(s, handler)
	if err := s.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("Serve")
	}
	log.Info().Msgf("GRPC server started %s", grpcConf.Addr())
}
