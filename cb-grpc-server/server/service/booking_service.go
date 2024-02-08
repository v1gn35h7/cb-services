package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/v1gn35h7/cb-grpc-server/internal/store"
	"github.com/v1gn35h7/cb-grpc-server/server/cb"
	"github.com/v1gn35h7/cb-grpc-server/server/pb"
)

type bookingService interface {
	BookTicket(pb.BookingRequest) (*cb.Booking, error)
	RemoveUser(*pb.RemoveUserRequest) (*cb.Ruser, error)
	ModifySeat(*pb.ModifySeatRequest) (*cb.Ruser, error)
	GetReceipt(*pb.ReceiptRequest) (*cb.Booking, error)
	GetSeatArrangenments(*pb.SeatArrangmentRequest) (*cb.Section, error)
}

func (srvc *service) BookTicket(r pb.BookingRequest) (*cb.Booking, error) {

	user := r.GetUser()

	if user == nil {
		return nil, fmt.Errorf("Invalid request")
	}

	store := store.GetStore()
	var booking *cb.Booking

	// Linear search for open slot
	// looks up all sections for availabity and allots a seats
	for k, v := range store.Sections {
		if len(v.AvailableSeats) > 0 {
			store.MU.Lock()
			defer store.MU.Unlock()

			seat := new(cb.Seat)

			seat.ID = v.AvailableSeats[0]

			// Pop top from available seats
			v.AvailableSeats = v.AvailableSeats[1:]

			seat.UserId = user.GetID()
			seat.Booked = true

			store.Sections[k].SeatsBooked += 1
			store.Sections[k].Seats[seat.ID] = seat

			bookingID := time.Now().Second()

			booking = &cb.Booking{
				From:    r.From,
				To:      r.To,
				ID:      strconv.Itoa(bookingID),
				Price:   r.Price,
				Section: k,
				SeatNo:  seat.ID,
				User: cb.User{
					ID:        user.ID,
					FirstName: user.FirstName,
					LastName:  user.LastName,
					Email:     user.Email,
				},
			}

			store.Bookings[int64(bookingID)] = booking
			store.UserBookings[user.ID] = int64(bookingID)
			break
		}
	}

	if booking == nil {
		return nil, fmt.Errorf("Failed: Seat not available")
	}

	return booking, nil
}

func (srvc *service) RemoveUser(r *pb.RemoveUserRequest) (*cb.Ruser, error) {
	userID := r.UserId

	if userID == "" {
		return nil, fmt.Errorf("Invalid request")
	}
	store := store.GetStore()

	store.MU.Lock()
	defer store.MU.Unlock()

	uid, _ := strconv.Atoi(userID)

	bookingID := store.UserBookings[int64(uid)]

	booking := store.Bookings[bookingID]

	if booking == nil {
		return nil, fmt.Errorf("No booking found")
	}

	delete(store.Sections[booking.Section].Seats, booking.SeatNo)

	return &cb.Ruser{Status: "OK"}, nil

}

func (srvc *service) ModifySeat(r *pb.ModifySeatRequest) (*cb.Ruser, error) {
	userID := r.UserID
	seatNO := r.SeatNO

	if userID == "" {
		return nil, fmt.Errorf("Invalid request")
	}

	store := store.GetStore()

	uid, _ := strconv.Atoi(userID)
	sno, _ := strconv.Atoi(seatNO)

	bookingID := store.UserBookings[int64(uid)]

	booking := store.Bookings[bookingID]

	if booking == nil {
		return nil, fmt.Errorf("No booking found")
	}

	seat, ok := store.Sections[booking.Section].Seats[int32(sno)]

	store.MU.Lock()
	defer store.MU.Unlock()

	if !ok {
		seat = new(cb.Seat)
		seat.ID = int32(sno)
		seat.UserId = int64(uid)
		seat.Booked = true

		//Delete Old seating allotment
		delete(store.Sections[booking.Section].Seats, booking.SeatNo)

		//Allocate new seat and Update booking details
		store.Sections[booking.Section].Seats[int32(sno)] = seat
		booking.SeatNo = int32(sno)

	} else {
		seat.UserId = int64(uid)

		booking.SeatNo = seat.ID
	}

	return &cb.Ruser{Status: "OK"}, nil

}

func (srvc *service) GetReceipt(r *pb.ReceiptRequest) (*cb.Booking, error) {
	userID := r.UserID

	if userID == "" {
		return nil, fmt.Errorf("Invalid request")
	}
	store := store.GetStore()

	uid, _ := strconv.Atoi(userID)

	bookingID := store.UserBookings[int64(uid)]

	booking := store.Bookings[bookingID]

	if booking == nil {
		return nil, fmt.Errorf("No booking found")
	}

	return booking, nil
}

func (srvc service) GetSeatArrangenments(req *pb.SeatArrangmentRequest) (*cb.Section, error) {
	store := store.GetStore()

	if store == nil {
		return nil, fmt.Errorf("No data found")
	}

	return store.Sections[strings.ToLower(req.Section)], nil
}
