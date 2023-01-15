package command

type command interface {
	execute()
}

type onCommand struct {
	device device
}

func (o *onCommand) execute() {
	o.device.on()
}

type offCommand struct {
	device device
}

func (o *offCommand) execute() {
	o.device.off()
}
