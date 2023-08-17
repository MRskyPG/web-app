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
