package main

import (
	"fmt"
	"math/rand"
)

type gameMaster struct {
	status  string
	players []player
	board   board
	turn    uint
}

func (ctx *gameMaster) getStatus() {
	fmt.Print(ctx.status)
}

func (ctx *gameMaster) setStatus(status string) {
	ctx.status = status
}

func (ctx *gameMaster) printPlayers() {
	for i := range 10 {
		for j := range 10 {
			fmt.Print(ctx.board.fullBoard[i][j])
		}
		fmt.Println()
	}
}

func (ctx *gameMaster) playerPosition(playerID uint) int {
	for _, player := range ctx.players {
		if player.ID == playerID {
			return player.position
		}
	}
	return 0
}

func (ctx *gameMaster) setPlayerPosition(playerID uint, newPosition int) {
	for ind, player := range ctx.players {
		if player.ID == playerID {
			if newPosition <= len(ctx.board.fullBoard)*len(ctx.board.fullBoard[0]) {
				if newPosition == len(ctx.board.fullBoard)*len(ctx.board.fullBoard[0]) {
					ctx.players[ind].status = "finished"
				}
				ctx.players[ind].position = newPosition
			}
		}
	}
}

func (ctx *gameMaster) playerPlayed(playerID uint, playerRole int) {
	playerNextPosition := ctx.board.nextPosition(ctx.playerPosition(playerID), playerRole)
	ctx.setPlayerPosition(playerID, playerNextPosition)
}

func (ctx *gameMaster) setSnakes(head []int, tail []int) map[int]int {
	var M map[int]int = make(map[int]int)
	for ind, _ := range head {
		M[head[ind]] = tail[ind]
	}
	return M
}
func (ctx *gameMaster) setLadders(head []int, tail []int) map[int]int {
	var M map[int]int = make(map[int]int)
	for ind, _ := range head {
		M[head[ind]] = tail[ind]
	}
	return M
}

func (ctx *gameMaster) nextTurn() uint {
	// should now work when all compleated
	var turnIndex int
	for ind, _ := range ctx.players {
		if ctx.players[ind].ID == ctx.turn {
			turnIndex = ind
			break
		}
	}
	for {
		turnIndex = (turnIndex + 1) % (len(ctx.players))
		if ctx.players[turnIndex].status == "playing" {
			return ctx.players[turnIndex].ID
		}
	}
}

func (ctx *gameMaster) startGame() {
	for {
		ctx.playerPlayed(ctx.turn, rand.Intn(3)) // need to change
		ctx.turn = ctx.nextTurn()
		for _, val := range ctx.players {
			fmt.Print(val.position, " ")
		}
		fmt.Println()
	}
}

func main() {
	// set up
	noPlayers := 6
	master := &gameMaster{
		players: make([]player, noPlayers),
		status:  "notStarted",
	}
	master.board = board{
		snakesPosition: master.setSnakes([]int{3}, []int{7}),
		LadderPosition: master.setLadders([]int{7, 4}, []int{5, 3}),
		fullBoard:      make([][]uint, 10),
	}
	for i := range len(master.board.fullBoard) {
		master.board.fullBoard[i] = make([]uint, len(master.board.fullBoard))
	}

	master.turn = 0
	for ind := range noPlayers {
		master.players[ind] = player{
			ID:       uint(ind),
			status:   "playing",
			position: 1,
		}
	}
	master.startGame()
}
