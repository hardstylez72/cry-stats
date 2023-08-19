package api

import (
	"context"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	grpcPort string

	callbacks *Callbacks
}

func (s *GrpcServer) Listen() {
	go func() {
		err := s.listenGrpc()
		if err != nil {
			panic(err)
		}
	}()
	log.Println("Server listening")
}

type InitConfig struct {
	GrpcPort string
}

type Callbacks struct {
	GrpcRegister func(s grpc.ServiceRegistrar)
}

func NewServer(initConfig *InitConfig, callbacks *Callbacks) *GrpcServer {
	if callbacks == nil {
		panic("Nil register callbacks")
	}

	return &GrpcServer{
		grpcPort: initConfig.GrpcPort,

		callbacks: callbacks,
	}
}

func (s *GrpcServer) listenGrpc() error {
	lis, err := (&net.ListenConfig{}).Listen(context.Background(), "tcp", s.grpcPort)
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	if s.callbacks.GrpcRegister != nil {
		s.callbacks.GrpcRegister(server)
	}

	err = server.Serve(lis)
	if err != nil { //&& err != http.ErrServerClosed
		return err
	}
	return nil
}

func corsWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
