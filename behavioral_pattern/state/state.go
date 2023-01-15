package state

import "fmt"

type MarioState int64

const (
	Small MarioState = 1 // 小马里奥
	Super MarioState = 2 // 超级马里奥
	Cape  MarioState = 3 // 斗篷马里奥
	Fire  MarioState = 4 // 火焰马里奥

)

type IMario interface {
	GetState() MarioState
	// 事情的定义
	ObtainMushroom()   // 吃蘑菇
	ObtainCape()       // 获得斗篷
	ObtainFireFlower() // 获得火焰
	MeetMonster()      // 遇到怪物
}

type SmallMario struct {
	machine *MarioStateMachine
}

func NewSmallMario(machine *MarioStateMachine) *SmallMario {
	return &SmallMario{machine: machine}
}

func (s *SmallMario) GetState() MarioState {
	return Small
}

func (s *SmallMario) ObtainMushroom() {
	s.machine.SetCurrentState(NewSuperMario(s.machine))
	s.machine.AddScore(100)
}

func (s *SmallMario) ObtainCape() {
	s.machine.SetCurrentState(NewCapMario(s.machine))
	s.machine.AddScore(200)
}

func (s *SmallMario) ObtainFireFlower() {
	s.machine.SetCurrentState(NewFireMario(s.machine))
	s.machine.AddScore(300)
}

func (s *SmallMario) MeetMonster() {
	fmt.Println("Game over, how about try again.")
}

type SuperMario struct {
	machine *MarioStateMachine
}

func NewSuperMario(machine *MarioStateMachine) *SuperMario {
	return &SuperMario{machine: machine}
}

func (s *SuperMario) GetState() MarioState {
	return Super
}

func (s *SuperMario) ObtainMushroom() {
	// do nothing
}

func (s *SuperMario) ObtainCape() {
	s.machine.SetCurrentState(NewCapMario(s.machine))
	s.machine.AddScore(200)
}

func (s *SuperMario) ObtainFireFlower() {
	s.machine.SetCurrentState(NewFireMario(s.machine))
	s.machine.AddScore(300)
}

func (s *SuperMario) MeetMonster() {
	s.machine.SetCurrentState(NewSmallMario(s.machine))
	s.machine.AddScore(-100)
}

type FireMario struct {
	machine *MarioStateMachine
}

func NewFireMario(machine *MarioStateMachine) *FireMario {
	return &FireMario{machine: machine}
}

func (f *FireMario) GetState() MarioState {
	return Fire
}

func (f *FireMario) ObtainMushroom() {
	// do nothing
}

func (f *FireMario) ObtainCape() {
	f.machine.SetCurrentState(NewCapMario(f.machine))
	f.machine.AddScore(200)
}

func (f *FireMario) ObtainFireFlower() {
	// do nothing
}

func (f *FireMario) MeetMonster() {
	f.machine.SetCurrentState(NewSmallMario(f.machine))
	f.machine.AddScore(-200)
}

type CapeMario struct {
	machine *MarioStateMachine
}

func NewCapMario(machine *MarioStateMachine) *CapeMario {
	return &CapeMario{machine: machine}
}

func (c *CapeMario) GetState() MarioState {
	return Cape
}

func (c *CapeMario) ObtainMushroom() {
	// do nothing
}

func (c *CapeMario) ObtainCape() {
	// do nothing
}

func (c *CapeMario) ObtainFireFlower() {
	// do nothing
}

func (c *CapeMario) MeetMonster() {
	c.machine.SetCurrentState(NewSmallMario(c.machine))
	c.machine.AddScore(-300)
}
