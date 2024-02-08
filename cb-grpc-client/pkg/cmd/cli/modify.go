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

var (
	userID int64
	seat   int32
)

func NewModifyCommand() *cobra.Command {
	var modifyCmd = &cobra.Command{
		Use:   "modify",
		Short: "Modify seat arrangment of user",
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

			// Make ticket modification request
			modify(c, logger)

		},
	}

	modifyCmd.PersistentFlags().Int64Var(&userID, "userID", 0, "User ID")
	modifyCmd.PersistentFlags().Int32Var(&seat, "seat", 0, "Seat no ")

	return modifyCmd
}

/*
*	Makes gRPC api call
*   Modify Ticket booking
 */
func modify(c *grpc.ClientConn, logger logr.Logger) {
	client := cpb.NewCbServiceClient(c)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1000)
	defer cancel()

	r, err := client.ModifySeat(ctx, &cpb.ModifySeatRequest{UserID: strconv.Itoa(int(userID)), SeatNO: strconv.Itoa(int(seat))})
	if err != nil {
		logger.Error(err, "Api failed")
	} else {
		logger.Info("Response from CloudBees gRPC server")
		fmt.Println(r)
	}
}
