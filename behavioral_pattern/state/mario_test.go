package state

import "testing"

func TestMario(t *testing.T) {
	machine := NewMarioStateMachine()
	t.Log(machine.GetScore()) // 0
	t.Log(machine.GetState()) // 1, represent small

	machine.ObtainMushroom()
	t.Log(machine.GetScore()) // 100
	t.Log(machine.GetState()) // 2, represent huge

	machine.ObtainMushroom()  // do nothing
	t.Log(machine.GetScore()) // 100
	t.Log(machine.GetState()) // 2, represent huge

	machine.ObtainFireFlower()
	t.Log(machine.GetScore()) // 400
	t.Log(machine.GetState()) // 4, represent fire

	machine.MeetMonster()
	t.Log(machine.GetScore()) // 300
	t.Log(machine.GetState()) // 1, represent small

	machine.MeetMonster() // game over
}
