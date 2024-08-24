package main

type board struct { // just 1-100 boxes
	snakesPosition map[int]int
	LadderPosition map[int]int
	fullBoard      [][]uint
}

func (ctx *board) nextPosition(presentPostion int, move int) int {
	next, ok := ctx.snakesPosition[presentPostion+move]
	if ok {
		return ctx.nextPosition(next, 0)
	}

	next, ok = ctx.LadderPosition[presentPostion+move]
	if ok {
		return ctx.nextPosition(next, 0)
	}
	return presentPostion + move
}
