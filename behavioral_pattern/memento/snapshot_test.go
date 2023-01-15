package memento

import "testing"

func TestSnapshot(t *testing.T) {
	input := NewInputText()

	input.Append("hello, ")
	input.Append("this is me. ")

	t.Log(input.GetText())

	snapshot := input.CreateSnapshot()
	t.Log(snapshot.GetText())

	input.Append("undo")
	t.Log(input.GetText())

	input.RestoreFromSnapshot(snapshot)
	t.Log(input.GetText())
}
