package data

// ParkingRepo contains all handlers
type ParkingRepo interface {
	userRepo
}

type userRepo interface {
	Login()
}
