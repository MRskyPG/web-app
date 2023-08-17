package service

import "github.com/MRskyPG/web-app"

type WebPositions interface {
	InsertPosition(pos *web.WorkingPosition)
	GetPosition(posID int) (web.WorkingPosition, error)
	UpdatePosition(posID int, pos web.WorkingPosition)
	DeletePosition(posID int)
	GetAllPositions() map[int]web.WorkingPosition
}

type Service interface {
	WebPositions
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
