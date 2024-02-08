package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/v1gn35h7/cb-grpc-client/internal/pb"
	cpb "github.com/v1gn35h7/cb-grpc-client/pb"

	"github.com/v1gn35h7/cb-grpc-client/pkg/logging"
	"google.golang.org/grpc"
)

func NewTestCommand() *cobra.Command {
	var testCmd = &cobra.Command{
		Use:   "test",
		Short: "test cb grpc connection",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			viper.BindPFlag("verbose", cmd.PersistentFlags().Lookup("verbose"))

			// Set-up logger
			logger := logging.Logger()
			logger.Info("Logger initated...")

			// Set-up gRPC client
			fmt.Println("Connecting to server...")
			c := pb.NewGrpcConnection(logger)
			defer c.Close()

			// Make helath request
			testConnection(c, logger)

		},
	}
	return testCmd
}

func testConnection(c *grpc.ClientConn, logger logr.Logger) {
	client := cpb.NewCbServiceClient(c)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1000)
	defer cancel()

	r, err := client.GetHealth(ctx, &cpb.HealthRequest{})
	if err != nil {
		logger.Error(err, "could not send proto message")
	} else {
		logger.Info("Response from CloudBees gRPC server")
		fmt.Println(r)
	}
}
