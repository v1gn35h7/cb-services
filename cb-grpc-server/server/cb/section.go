package cb

type Section struct {
	Name           string
	Seats          map[int32]*Seat
	limit          int32
	SeatsBooked    int32
	AvailableSeats []int32
}

func (s *Section) SetLimit(l int32) {
	s.limit = l
}

func (s *Section) Limit() int32 {
	return s.limit
}
