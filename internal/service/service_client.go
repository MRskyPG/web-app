package service

import (
	"errors"
	"github.com/MRskyPG/web-app"
	"sync"
)

type MemoryClients struct {
	counter int
	data    map[int]web.Client
	sync.Mutex
}

func NewMemoryClients() *MemoryClients {
	return &MemoryClients{
		counter: 1,
		data:    make(map[int]web.Client),
	}
}

func (mp *MapService) InsertClient(c *web.Client) {
	mp.MemoryClients.Lock()

	c.ID = mp.MemoryClients.counter
	mp.MemoryClients.data[c.ID] = *c
	mp.MemoryClients.counter++

	//Производим вставку client_id соответствующему заказу
	mp.MemoryOrders.Lock()
	if order, ok := mp.MemoryOrders.data[c.OrderID]; ok {
		order.ClientID = c.ID
		mp.MemoryOrders.data[c.OrderID] = order
	}
	mp.MemoryOrders.Unlock()
	mp.MemoryClients.Unlock()
}

func (mp *MapService) GetClient(cID int) (web.Client, error) {
	mp.MemoryClients.Lock()
	defer mp.MemoryClients.Unlock()

	client, ok := mp.MemoryClients.data[cID]
	if !ok {
		return client, errors.New("Client not found.")
	}
	return client, nil
}

func (mp *MapService) UpdateClient(cID int, c web.Client) {
	mp.MemoryClients.Lock()
	defer mp.MemoryClients.Unlock()

	mp.MemoryClients.data[cID] = c

	//Производим вставку client_id соответствующему заказу
	mp.MemoryOrders.Lock()
	if order, ok := mp.MemoryOrders.data[c.OrderID]; ok {
		order.ClientID = c.ID
		mp.MemoryOrders.data[c.OrderID] = order
	}
	mp.MemoryOrders.Unlock()
}

func (mp *MapService) DeleteClient(cID int) {
	mp.MemoryClients.Lock()
	defer mp.MemoryClients.Unlock()
	delete(mp.MemoryClients.data, cID)
}

func (mp *MapService) GetAllClients() map[int]web.Client {
	mp.MemoryClients.Lock()
	defer mp.MemoryClients.Unlock()
	return mp.MemoryClients.data
}
