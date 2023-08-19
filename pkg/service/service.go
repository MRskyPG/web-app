package service

import "github.com/MRskyPG/web-app"

type WebPositions interface {
	InsertPosition(pos *web.WorkingPosition)
	GetPosition(posID int) (web.WorkingPosition, error)
	UpdatePosition(posID int, pos web.WorkingPosition)
	DeletePosition(posID int)
	GetAllPositions() map[int]web.WorkingPosition
}

type WebStaff interface {
	InsertStaff(s *web.Staff)
	GetAllStaff() map[int]web.Staff
	UpdateStaff(sID int, s web.Staff)
	DeleteStaff(sID int)
	GetStaff(sID int) (web.Staff, error)
}

type WebClient interface {
	InsertClient(c *web.Client)
	GetAllClients() map[int]web.Client
	UpdateClient(cID int, c web.Client)
	DeleteClient(cID int)
	GetClient(cID int) (web.Client, error)
}

type WebOrder interface {
	InsertOrder(o *web.Order)
	GetAllOrders() map[int]web.Order
	UpdateOrder(oID int, c web.Order)
	DeleteOrder(oID int)
	GetOrder(oID int) (web.Order, error)
}

type Service interface {
	WebPositions
	WebStaff
	WebClient
	WebOrder
}

type MapService struct {
	*MemoryPositions
	*MemoryStaff
	*MemoryClients
	*MemoryOrders
}

func NewMapService() *MapService {
	return &MapService{
		MemoryPositions: NewMemoryPositions(),
		MemoryStaff:     NewMemoryStaff(),
		MemoryClients:   NewMemoryClients(),
		MemoryOrders:    NewMemoryOrders(),
	}
}
