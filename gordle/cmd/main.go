package main

import (
	"os"

	"github.com/AlexTLDR/tiny-projects/gordle/game"
)

func main(){
	g := game.NewGame(os.Stdin)
	g.Play()
}
