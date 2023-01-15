package observer

type Subject interface {
	registerObservers(observer ...Observer)
	removeObserver(observer Observer)
	notifyObserves(message string)
}

type ConcreteSubject struct {
	observers []Observer
}

func (c *ConcreteSubject) registerObservers(observer ...Observer) {
	c.observers = append(c.observers, observer...)
}

func (c *ConcreteSubject) removeObserver(observer Observer) {
	observers := c.observers
	index := -1
	length := len(observers)
	for i, existObserver := range observers {
		if existObserver.GetID() == observer.GetID() {
			index = i
			break
		}
	}
	if index == -1 {
		return
	}
	observers[index], observers[length-1] = observers[length-1], observers[index]
	c.observers = observers[:length-1]
}

func (c *ConcreteSubject) notifyObserves(message string) {
	for _, observer := range c.observers {
		observer.Update(message)
	}
}
