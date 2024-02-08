package cli

import (
	"context"
	"fmt"
	"math/rand"
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
	firstName string
	lastName  string
	email     string
)

func NewBookCommand() *cobra.Command {
	var bookCmd = &cobra.Command{
		Use:   "book",
		Short: "Books ticket from London to France",
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
			bookTicket(c, logger)

		},
	}

	bookCmd.PersistentFlags().StringVar(&firstName, "fname", "", "User First name")
	bookCmd.PersistentFlags().StringVar(&lastName, "lname", "", "User Last Name")
	bookCmd.PersistentFlags().StringVar(&email, "email", "", "User Email")

	return bookCmd
}

func bookTicket(c *grpc.ClientConn, logger logr.Logger) {
	client := cpb.NewCbServiceClient(c)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1000)
	defer cancel()

	request := cpb.BookingRequest{
		From:  "London",
		To:    "France",
		Price: 20.00,
		User: &cpb.User{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			ID:        rand.Int63(),
		},
	}
	r, err := client.BookTicket(ctx, &request)
	if err != nil {
		logger.Error(err, "Api failed")
	} else {
		logger.Info("Response from CloudBees gRPC server")
		fmt.Println("Booking Confirmation")
		fmt.Println("--------------------------------------")
		fmt.Println(r)
	}
}
