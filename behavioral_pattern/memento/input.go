package memento

import "strings"

type InputText struct {
	text strings.Builder
}

func NewInputText() *InputText {
	return &InputText{text: strings.Builder{}}
}

func (i *InputText) GetText() string {
	return i.text.String()
}

func (i *InputText) Append(str string) {
	i.text.WriteString(str)
}

func (i *InputText) CreateSnapshot() *Snapshot {
	return NewSnapshot(i.GetText())
}

func (i *InputText) RestoreFromSnapshot(snapshot *Snapshot) {
	i.text.Reset()
	i.text.WriteString(snapshot.GetText())
}
