package game

import "fmt"

// Game main container for a tic tac toe game
type Game struct {
	x     Player
	o     Player
	size  []int
	board []byte
}

func newBoard(size []int) []byte {
	return make([]byte, size[0]*size[1])
}

// NewHumanVsHumanGame Creates a new game with 2 human players
func NewHumanVsHumanGame() *Game {
	p1 := newHumanPlayer("P1", 'X')
	p2 := newHumanPlayer("P2", 'O')
	var size = []int{3, 3}
	return &Game{p1, p2, size, newBoard(size)}
}

func (g *Game) String() string {
	var str string
	for i, cell := range g.board {
		if cell == 0 {
			cell = '_'
		}
		str += string(cell)
		if (i+1)%g.size[0] == 0 {
			str += "\n"
		}
	}
	return str
}

// Update update the game with a new move by player p
func (g *Game) Update(p Player, move int) bool {
	if move < 0 || move >= g.Cells() || g.board[move] != 0 {
		fmt.Printf("Invalid move %v!\n", move)
		return false
	}
	g.board[move] = p.symbol()
	fmt.Println(g)

	return true
}

// Play the game
func (g *Game) Play(winner chan Player) {
	activePlayer := g.x
	go g.x.play()
	go g.o.play()
	fmt.Println(g)
	for {
		valid := false
		var move int
		for !valid {
			fmt.Printf("Sending activation to %s...\n", activePlayer)
			activePlayer.chanGo() <- true
			fmt.Printf("Waiting for %s\n", activePlayer)
			move = <-activePlayer.chanMove()
			valid = g.Update(activePlayer, move)
		}
		fmt.Printf("%v made move %v\n", activePlayer, move)
		win := g.CheckWinner(activePlayer)
		if win {
			winner <- activePlayer
			break
		}
		switch activePlayer {
		case g.x:
			activePlayer = g.o
		default:
			activePlayer = g.x
		}
	}

}

// Cells # of cells in the game board
func (g *Game) Cells() int {
	return g.size[0] * g.size[1]
}

var winPatterns = [...][]byte{
	// Rows
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},

	// Columns
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},

	// Diagonals
	{0, 4, 8},
	{2, 4, 6},
}

// CheckWinner determines if the player p has won the game
func (g *Game) CheckWinner(p Player) bool {
	symbol := p.symbol()

	for i, pattern := range winPatterns {
		winner := true
		for j := 0; j < 3; j++ {
			if g.board[pattern[j]] != symbol {
				winner = false
				continue
			}
		}
		if winner {
			fmt.Printf("%v won by pattern %v\n", p, i)
			return true
		}
	}
	return false
}
