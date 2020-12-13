package payload

import "strconv"

type Payload struct {
	id          string
	pixKey      string
	description string
	merchant    Merchant
	amount      string
}

type Merchant struct {
	name string
	city string
}

func NewPayload() *Payload {
	return new(Payload)
}

func (p *Payload) setId(id string) *Payload {
	p.id = id
	return p
}

func (p *Payload) setPixKey(key string) *Payload {
	p.pixKey = key
	return p
}

func (p *Payload) setDescription(description string) *Payload {
	p.description = description
	return p
}

func (p *Payload) setMerchantName(name string) *Payload {
	p.merchant.name = name
	return p
}

func (p *Payload) setMerchantCity(city string) *Payload {
	p.merchant.city = city
	return p
}

func (p *Payload) setAmount(amount float64) *Payload {
	p.amount = strconv.FormatFloat(amount, 'f', 2, 64)
	return p
}

func (p *Payload) getId() string {
	return p.id
}

func (p *Payload) getPixKey() string {
	return p.pixKey
}

func (p *Payload) getDescription() string {
	return p.description
}

func (p *Payload) getMerchantName() string {
	return p.merchant.name
}

func (p *Payload) getMerchantCity() string {
	return p.merchant.city
}

func (p *Payload) getAmount() string {
	return p.amount
}
