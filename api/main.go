package api

import (
	"context"
	"net"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"coginfra/interceptors"
	"coginfra/protos"
	"coginfra/utils"
)

func InitServer(s *Server) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		utils.Logger.Fatalf("[error] failed to listen: %v", err)
	}

	// Create a gRPC server object
	gServer := grpc.NewServer(grpc.UnaryInterceptor(interceptors.APIKeyInterceptor()))
	// Atach document service to gRPC server
	protos.RegisterDocumentServiceServer(gServer, s)

	// Serve gRPC Server
	utils.Logger.Println("[info] Serving gRPC on :50051 ...")
	go func() {
		utils.Logger.Fatalln(gServer.Serve(lis))
	}()

	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		ctx,
		"0.0.0.0:50051",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		utils.Logger.Fatalln("[error] Failed to dial server:", err)
	}

	// log.Println("[debug] hellooowefnrf")

	gwmux := runtime.NewServeMux(
		runtime.WithMetadata(APIKeyAnnotator),
	)
	// Register Document Service
	utils.Logger.Println("[info] Registering Document Service [gRPC-Gateway]")
	err = protos.RegisterDocumentServiceHandler(ctx, gwmux, conn)
	if err != nil {
		utils.Logger.Fatalln("[error] Failed to register gateway:", err)
	}

	handler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods(
			[]string{http.MethodPost, http.MethodGet},
		),
	)(
		gwmux,
	)

	// Serve gRPC-Gateway
	gatewayServer := &http.Server{
		Addr:    s.listenAddress,
		Handler: handler,
	}

	utils.Logger.Println("[info] Serving gRPC-Gateway on http://0.0.0.0:" + s.listenAddress)

	go func() {
		s.HasStarted = true
		err := gatewayServer.ListenAndServe()
		if err != nil {
			utils.Logger.Fatalln("[error] Failed to start gateway server:", err)
		}
	}()
}

func APIKeyAnnotator(_ context.Context, req *http.Request) metadata.MD {
	apiKey := req.Header.Get(interceptors.XAPIKey)
	return metadata.New(map[string]string{
		interceptors.XAPIKey: apiKey,
	})
}
