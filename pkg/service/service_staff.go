package service

import (
	"errors"
	"github.com/MRskyPG/web-app"
	"sync"
)

type MemoryPositions struct {
	counter int
	data    map[int]web.WorkingPosition
	sync.Mutex
}

type MemoryStaff struct {
	counter int
	data    map[int]web.Staff
	sync.Mutex
}

func NewMemoryPositions() *MemoryPositions {
	return &MemoryPositions{
		counter: 1,
		data:    make(map[int]web.WorkingPosition),
	}
}

func NewMemoryStaff() *MemoryStaff {
	return &MemoryStaff{
		counter: 1,
		data:    make(map[int]web.Staff),
	}
}

func (mp *MapService) InsertPosition(pos *web.WorkingPosition) {
	mp.MemoryPositions.Lock()

	//Вставляем содержимое в мапу по ключу mp.counter = 1, 2...
	pos.PositionID = mp.MemoryPositions.counter
	mp.MemoryPositions.data[pos.PositionID] = *pos
	mp.MemoryPositions.counter++
	mp.MemoryPositions.Unlock()
}

func (mp *MapService) GetPosition(posID int) (web.WorkingPosition, error) {
	mp.MemoryPositions.Lock()
	defer mp.MemoryPositions.Unlock()

	position, ok := mp.MemoryPositions.data[posID]
	if !ok {
		return position, errors.New("Position not found.")
	}
	return position, nil
}

func (mp *MapService) UpdatePosition(posID int, pos web.WorkingPosition) {
	mp.MemoryPositions.Lock()
	defer mp.MemoryPositions.Unlock()
	mp.MemoryPositions.data[posID] = pos
}

func (mp *MapService) DeletePosition(posID int) {
	mp.MemoryPositions.Lock()
	defer mp.MemoryPositions.Unlock()
	delete(mp.MemoryPositions.data, posID)
}

func (mp *MapService) GetAllPositions() map[int]web.WorkingPosition {
	mp.MemoryPositions.Lock()
	defer mp.MemoryPositions.Unlock()
	return mp.MemoryPositions.data
}

func (mp *MapService) InsertStaff(s *web.Staff) {
	mp.MemoryStaff.Lock()

	//Вставляем содержимое в мапу по ключу mp.counter = 1, 2...
	s.StaffID = mp.MemoryStaff.counter
	mp.MemoryStaff.data[s.StaffID] = *s
	mp.MemoryStaff.counter++
	mp.MemoryStaff.Unlock()
}

func (mp *MapService) GetAllStaff() map[int]web.Staff {
	mp.MemoryStaff.Lock()
	defer mp.MemoryStaff.Unlock()
	return mp.MemoryStaff.data
}

func (mp *MapService) GetStaff(sID int) (web.Staff, error) {
	mp.MemoryStaff.Lock()
	defer mp.MemoryStaff.Unlock()

	staff, ok := mp.MemoryStaff.data[sID]
	if !ok {
		return staff, errors.New("Position not found.")
	}
	return staff, nil
}

func (mp *MapService) UpdateStaff(sID int, s web.Staff) {
	mp.MemoryStaff.Lock()
	defer mp.MemoryStaff.Unlock()
	mp.MemoryStaff.data[sID] = s
}

func (mp *MapService) DeleteStaff(sID int) {
	mp.MemoryStaff.Lock()
	defer mp.MemoryStaff.Unlock()
	delete(mp.MemoryStaff.data, sID)
}
