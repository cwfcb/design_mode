package observer

import "fmt"

type Observer interface {
	Update(message string)
	GetID() string
}

type CustomerObserver struct {
	customID string
}

func NewCustomerObserver(id string) *CustomerObserver {
	return &CustomerObserver{customID: "customer_" + id}
}

func (c *CustomerObserver) Update(message string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.customID, message)
}

func (c *CustomerObserver) GetID() string {
	return c.customID
}

type MerchantObserver struct {
	merchantID string
}

func NewMerchantObserver(id string) *MerchantObserver {
	return &MerchantObserver{merchantID: "merchant_" + id}
}

func (m *MerchantObserver) Update(message string) {
	fmt.Printf("Phone call merchant %s for item %s\n", m.merchantID, message)
}

func (m *MerchantObserver) GetID() string {
	return m.merchantID
}
