package main

import "math/rand"

type player struct {
	ID       uint
	position int
	status   string // finished playing notstarted
}

func (ctx *player) play() int {
	return rand.Intn(6) + 1
}
