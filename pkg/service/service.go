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

type Service interface {
	WebPositions
	WebStaff
}

type MapService struct {
	*MemoryPositions
	*MemoryStaff
}

func NewMapService() *MapService {
	return &MapService{
		MemoryPositions: NewMemoryPositions(),
		MemoryStaff:     NewMemoryStaff(),
	}
}
