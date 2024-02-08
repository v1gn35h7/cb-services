package cb

type Booking struct {
	User    User
	ID      string
	From    string
	To      string
	Price   float64
	Section string
	SeatNo  int32
}
