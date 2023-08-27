package service

import (
	"errors"
	"github.com/MRskyPG/web-app"
	"sync"
)

type MemoryOrders struct {
	counter int
	data    map[int]web.Order
	sync.Mutex
}

func NewMemoryOrders() *MemoryOrders {
	return &MemoryOrders{
		counter: 1,
		data:    make(map[int]web.Order),
	}
}

func (mp *MapService) InsertOrder(o *web.Order) {
	mp.MemoryOrders.Lock()

	o.OrderID = mp.MemoryOrders.counter
	mp.MemoryOrders.data[o.OrderID] = *o
	mp.MemoryOrders.counter++

	//Производим вставку order_id соответствующему клиенту
	mp.MemoryClients.Lock()
	if client, ok := mp.MemoryClients.data[o.ClientID]; ok {
		client.OrderID = o.OrderID
		mp.MemoryClients.data[o.ClientID] = client
	}
	mp.MemoryClients.Unlock()
	mp.MemoryOrders.Unlock()
}

func (mp *MapService) GetOrder(oID int) (web.Order, error) {
	mp.MemoryOrders.Lock()
	defer mp.MemoryOrders.Unlock()

	order, ok := mp.MemoryOrders.data[oID]
	if !ok {
		return order, errors.New("Order not found.")
	}
	return order, nil
}

func (mp *MapService) UpdateOrder(oID int, o web.Order) {
	mp.MemoryOrders.Lock()
	defer mp.MemoryOrders.Unlock()
	mp.MemoryOrders.data[oID] = o

	//Производим вставку order_id соответствующему клиенту
	mp.MemoryClients.Lock()
	if client, ok := mp.MemoryClients.data[o.ClientID]; ok {
		client.OrderID = o.OrderID
		mp.MemoryClients.data[o.ClientID] = client
	}
	mp.MemoryClients.Unlock()
}

func (mp *MapService) DeleteOrder(oID int) {
	mp.MemoryOrders.Lock()
	defer mp.MemoryOrders.Unlock()
	delete(mp.MemoryOrders.data, oID)
}

func (mp *MapService) GetAllOrders() map[int]web.Order {
	mp.MemoryOrders.Lock()
	defer mp.MemoryOrders.Unlock()
	return mp.MemoryOrders.data
}
