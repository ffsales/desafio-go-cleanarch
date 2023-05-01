package event

import "time"

type OrderListed struct {
	Name    string
	Payload interface{}
}

func NewOrderListed() *OrderListed {
	return &OrderListed{
		Name: "OrderCreated",
	}
}

func (ol *OrderListed) GetName() string {
	return ol.Name
}

func (ol *OrderListed) GetPayload() interface{} {
	return ol.Payload
}

func (ol *OrderListed) SetPayload(payload interface{}) {
	ol.Payload = payload
}

func (ol *OrderListed) GetDateTime() time.Time {
	return time.Now()
}
