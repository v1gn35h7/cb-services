package service

import (
	"math/rand"
	"os"
	"strconv"
	"testing"

	kitlog "github.com/go-kit/log"
	"github.com/v1gn35h7/cb-grpc-server/internal/store"
	"github.com/v1gn35h7/cb-grpc-server/server/cb"
	"github.com/v1gn35h7/cb-grpc-server/server/pb"
)

var srvc service

func init() {
	//Logger setup
	logger := kitlog.NewLogfmtLogger(os.Stderr)

	// Setup In-Memory Store
	_ = store.New(2, 10)

	// CloudBees Train Booking Service
	srvc = New(logger)

}

func bookTicket() (*cb.Booking, error) {
	request := &pb.BookingRequest{
		From:  "London",
		To:    "France",
		Price: 20.00,
		User: &pb.User{
			FirstName: "Bob",
			LastName:  "Colt",
			Email:     "bob@colt.in",
			ID:        rand.Int63(),
		},
	}

	return srvc.BookTicket(*request)

}

func TestBookTicket(t *testing.T) {

	got, err := bookTicket()

	if err != nil {
		t.Error("Booking failed")
	}

	if got == nil {
		t.Error("Booking failed")
	}

}

func TestGetSeatArrngmnt(t *testing.T) {
	request := &pb.SeatArrangmentRequest{Section: "A"}

	_, err := bookTicket()

	if err != nil {
		t.Error("Booking failed")
	}
	got, err := srvc.GetSeatArrangenments(request)

	if err != nil {
		t.Error("Loading Seat arrangment failed")
	}

	if got == nil {
		t.Error("Loading Seat arrangment failed")
	}

}

func TestRemoveUser(t *testing.T) {

	b, err := bookTicket()

	if err != nil {
		t.Error("Booking failed")
	}

	request := &pb.RemoveUserRequest{UserId: strconv.Itoa(int(b.User.ID))}

	got, err := srvc.RemoveUser(request)

	if err != nil {
		t.Error("Loading Seat arrangment failed")
	}

	if got == nil {
		t.Error("Loading Seat arrangment failed")
	}

}
