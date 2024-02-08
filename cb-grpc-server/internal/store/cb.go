package store

import (
	"sync"

	"github.com/v1gn35h7/cb-grpc-server/server/cb"
	"github.com/v1gn35h7/cb-grpc-server/server/pb"
)

var store *CloudBeesStore

type CloudBeesStore struct {
	MU           *sync.RWMutex
	Users        map[int64]*pb.User     // Users hash table
	Bookings     map[int64]*cb.Booking  // Bookings hash table key: bookingId Value: Booking struct
	Sections     map[string]*cb.Section // Train Sections hash table  key: Section Name Value: Section struct
	UserBookings map[int64]int64        // UserId Index of bookings
}

func New(maxSections int32, seatLimit int32) *CloudBeesStore {
	if store != nil {
		return store
	}

	store = new(CloudBeesStore)
	store.MU = &sync.RWMutex{}
	store.Sections = make(map[string]*cb.Section)
	store.Bookings = make(map[int64]*cb.Booking)
	store.Users = make(map[int64]*pb.User)
	store.UserBookings = make(map[int64]int64)

	for i := 0; i < int(maxSections); i++ {
		s := 97
		k := s + i
		section := &cb.Section{
			Name:           string(k),
			Seats:          make(map[int32]*cb.Seat),
			AvailableSeats: make([]int32, 0),
		}

		for j := 0; j < int(seatLimit); j++ {
			section.AvailableSeats = append(section.AvailableSeats, int32(j))
		}

		section.SetLimit(seatLimit)

		store.Sections[string(k)] = section
	}

	return store
}

func GetStore() *CloudBeesStore {
	if store != nil {
		return store
	}

	return New(2, 10)
}
