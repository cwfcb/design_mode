package observer

import "testing"

func TestObserver(t *testing.T) {
	customerObserver1 := &CustomerObserver{customID: "1"}
	customerObserver2 := &CustomerObserver{customID: "2"}

	merchantObserver1 := &MerchantObserver{merchantID: "1"}
	merchantObserver2 := &MerchantObserver{merchantID: "2"}

	subject := &ConcreteSubject{}
	subject.registerObservers(customerObserver1, merchantObserver1)
	subject.notifyObserves("Nike Shirt")

	subject.registerObservers(customerObserver2, merchantObserver2)
	subject.notifyObserves("Adidas Shoes")

	subject.removeObserver(merchantObserver2)
	subject.notifyObserves("Puma Pants")
}
