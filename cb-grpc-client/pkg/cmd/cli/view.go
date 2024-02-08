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
	section string
)

func NewViewCommand() *cobra.Command {
	var viewCmd = &cobra.Command{
		Use:   "view",
		Short: "View seat arrangements",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			viper.BindPFlag("section", cmd.PersistentFlags().Lookup("section"))

			// Set-up logger
			logger := logging.Logger()
			logger.Info("Logger initated...")

			// Set-up gRPC client
			fmt.Println("Connecting to server...")
			c := pb.NewGrpcConnection(logger)
			defer c.Close()

			// Make helath request
			viewSection(c, logger)

		},
	}

	viewCmd.PersistentFlags().StringVar(&section, "section", "", "View booking details of a section")

	return viewCmd
}

func viewSection(c *grpc.ClientConn, logger logr.Logger) {
	client := cpb.NewCbServiceClient(c)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1000)
	defer cancel()

	r, err := client.GetSeatArrangenments(ctx, &cpb.SeatArrangmentRequest{Section: section})
	if err != nil {
		logger.Error(err, "gRPC API failed")
	} else {
		logger.Info("Response from CloudBees gRPC server")

		for _, v := range r.Sections {
			fmt.Println("Section:", v.Name)
			fmt.Println("-------------------------------------------")

			for _, s := range v.Seats {
				fmt.Println("Seat No:", s.Number, " User Id: ", s.UserID)
			}
		}
	}
}
