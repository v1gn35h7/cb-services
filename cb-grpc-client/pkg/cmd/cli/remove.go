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

var (
	userid string
)

func NewRemoveCommand() *cobra.Command {
	var rmvCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove user from train",
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

			// Make cancel request
			removeUser(c, logger)

		},
	}

	rmvCmd.PersistentFlags().StringVar(&userid, "userID", "", "User ID")

	return rmvCmd
}

/*
*	Makes gRPC api call
*   Cancel booking
 */
func removeUser(c *grpc.ClientConn, logger logr.Logger) {
	client := cpb.NewCbServiceClient(c)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1000)
	defer cancel()

	r, err := client.RemoveUser(ctx, &cpb.RemoveUserRequest{UserId: userid})
	if err != nil {
		logger.Error(err, "Api failed")
	} else {
		logger.Info("Response from CloudBees gRPC server")
		fmt.Println(r)
	}
}
