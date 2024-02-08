package cli

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-logr/logr"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/v1gn35h7/cb-grpc-client/internal/pb"
	cpb "github.com/v1gn35h7/cb-grpc-client/pb"

	"github.com/v1gn35h7/cb-grpc-client/pkg/logging"
	"google.golang.org/grpc"
)

func NewReceiptCommand() *cobra.Command {
	var receiptCmd = &cobra.Command{
		Use:   "receipt",
		Short: "Load booking receipt for user",
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
			getReceipt(c, logger)

		},
	}

	receiptCmd.PersistentFlags().Int64Var(&userID, "userID", 0, "User ID")

	return receiptCmd
}

func getReceipt(c *grpc.ClientConn, logger logr.Logger) {
	client := cpb.NewCbServiceClient(c)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1000)
	defer cancel()

	r, err := client.GetReceipt(ctx, &cpb.ReceiptRequest{UserID: strconv.Itoa(int(userID))})

	if err != nil {
		logger.Error(err, "Api failed")
	} else {
		logger.Info("Response from CloudBees gRPC server")
		fmt.Println("Booking details")
		fmt.Println("--------------------------------------")
		fmt.Println(r)
	}
}
