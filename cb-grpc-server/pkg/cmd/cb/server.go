package cb

import (
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	kitlog "github.com/go-kit/log"
	"github.com/oklog/oklog/pkg/group"
	"github.com/spf13/cobra"
	"github.com/v1gn35h7/cb-grpc-server/internal/config"
	"github.com/v1gn35h7/cb-grpc-server/internal/store"
	"github.com/v1gn35h7/cb-grpc-server/pkg/logging"
	"github.com/v1gn35h7/cb-grpc-server/server/pb"
	"github.com/v1gn35h7/cb-grpc-server/server/service"
	grpctransport "github.com/v1gn35h7/cb-grpc-server/server/transport/grpc"
	"google.golang.org/grpc"
)

var (
	configPath string
)

func NewCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "CloudBees gRPC Server",
		Short: "Train booking service",
		Long:  "-----------------------------------------------",
		Run: func(cmd *cobra.Command, args []string) {
			//Bootstrap server
			bootStrapServer()
		},
	}

	// Bind cli flags
	rootCmd.PersistentFlags().StringVar(&configPath, "conf", "", "config file path")

	return rootCmd
}

func bootStrapServer() {
	//Logger setup
	logger := kitlog.NewLogfmtLogger(os.Stderr)

	// Init read config
	fmt.Println("Config path set to: ", configPath)
	config.ReadConfig(configPath, logging.Logger())

	// Setup In-Memory Store
	_ = store.New(2, 10)

	// Create CloudBees Train Booking Service instance
	srvc := service.New(logger)

	// Starts group of gorountines and shuts all goroutines if one member fails
	var g group.Group
	{
		// HTTP listener.
		httpServer := &http.Server{
			Handler:      nil,
			Addr:         ":8080",
			WriteTimeout: 30 * time.Second,
			ReadTimeout:  30 * time.Second,
		}
		g.Add(func() error {
			fmt.Println("Cloudbees Http server started")
			logger.Log("ServiceTransport", "HTTP", "Port", "8080")
			fmt.Println("For Profile data Visit: http://localhost:8080/debug/pprof/")

			return httpServer.ListenAndServe()
		}, func(error) {
			// Stop http listener
			httpServer.Close()
		})
	}
	{
		// Start gRPC server
		grpcServer := grpctransport.NewGRPCServer(grpctransport.MakeGrpcEndpoints(&srvc, logger))

		// The gRPC listener mounts the Go kit gRPC server we created.
		grpcListener, err := net.Listen("tcp", "localhost:8082")
		if err != nil {
			logger.Log("transport", "gRPC", "during", "Listen", "err", err)
			os.Exit(1)
		}

		g.Add(func() error {
			fmt.Println("Cloudbees gRPC server started")
			logger.Log("ServiceTransport", "gRPC", "Port", "8082")
			// we add the Go Kit gRPC Interceptor to our gRPC service
			baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
			pb.RegisterCbServiceServer(baseServer, grpcServer)
			return baseServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}
	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received cancel signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	logger.Log("exit", g.Run())

}
