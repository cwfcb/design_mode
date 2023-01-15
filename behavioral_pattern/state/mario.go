package state

type MarioStateMachine struct {
	Score        int64
	CurrentState IMario
}

func NewMarioStateMachine() *MarioStateMachine {
	machine := &MarioStateMachine{}
	smallMario := NewSmallMario(machine)
	machine.CurrentState = smallMario
	return machine
}

func (m *MarioStateMachine) ObtainMushroom() {
	m.CurrentState.ObtainMushroom()
}

func (m *MarioStateMachine) ObtainCape() {
	m.CurrentState.ObtainCape()
}

func (m *MarioStateMachine) ObtainFireFlower() {
	m.CurrentState.ObtainFireFlower()
}

func (m *MarioStateMachine) MeetMonster() {
	m.CurrentState.MeetMonster()
}

func (m *MarioStateMachine) GetScore() int64 {
	return m.Score
}

func (m *MarioStateMachine) GetState() MarioState {
	return m.CurrentState.GetState()
}

func (m *MarioStateMachine) AddScore(score int64) {
	m.Score += score
}

func (m *MarioStateMachine) SetCurrentState(state IMario) {
	m.CurrentState = state
}
