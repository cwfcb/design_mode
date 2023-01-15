package flyweight

import "testing"

func TestChess(t *testing.T) {
	boards := make([]*ChessBoard, 100)
	for i := 0; i < 100; i++ {
		board := &ChessBoard{}
		board.Init()
		boards[i] = board
	}

	for i := 0; i < 100; i++ {
		t.Log(boards[i])
	}
}
