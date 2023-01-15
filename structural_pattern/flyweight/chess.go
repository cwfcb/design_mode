package flyweight

type Color int8

// ChessPieceUnit 享元对象
type ChessPieceUnit struct {
	id    int
	text  string
	color Color
}

var (
	pieces map[int8]*ChessPieceUnit
	Black  Color = 1
	Red    Color = 2
)

func init() {
	pieces = make(map[int8]*ChessPieceUnit)
	pieces[1] = &ChessPieceUnit{1, "车", Black}
	pieces[2] = &ChessPieceUnit{2, "马", Red}
	pieces[2] = &ChessPieceUnit{2, "将", Black}
	pieces[2] = &ChessPieceUnit{2, "帅", Red}
	// ...省略摆放其他棋子的代码...
}

// ChessPieceUnitFactory 生成享元对象的工厂类
type ChessPieceUnitFactory struct {
}

func (c *ChessPieceUnitFactory) GetChessPieceUnit(chessPieceUnitID int8) *ChessPieceUnit {
	return pieces[chessPieceUnitID]
}

// ChessPiece 棋子
type ChessPiece struct {
	unit      *ChessPieceUnit
	positionX int16
	positionY int16
}

// ChessBoard 棋盘
type ChessBoard struct {
	chessPieces map[int16]ChessPiece
}

func (c *ChessBoard) Init() {
	factory := &ChessPieceUnitFactory{}
	chessPieces := make(map[int16]ChessPiece)

	chessPieces[1] = ChessPiece{factory.GetChessPieceUnit(1), 1, 1}
	chessPieces[2] = ChessPiece{factory.GetChessPieceUnit(2), 1, 2}
	chessPieces[3] = ChessPiece{factory.GetChessPieceUnit(3), 2, 1}
	chessPieces[4] = ChessPiece{factory.GetChessPieceUnit(4), 2, 2}
	// ...省略摆放其他棋子的代码...
	c.chessPieces = chessPieces
}
