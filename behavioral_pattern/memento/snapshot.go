package memento

type Snapshot struct {
	text string
}

func NewSnapshot(text string) *Snapshot {
	return &Snapshot{text: text}
}

func (s *Snapshot) GetText() string {
	return s.text
}
