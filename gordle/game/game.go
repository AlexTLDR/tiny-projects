package game

import "fmt"

type Game struct {
    // Define fields for the game state
}

// NewGame creates a new game instance.
func NewGame() *Game {
    g := &Game{}
    return g
}

// Play runs the game
func (g *Game) Play() {
    fmt.Println("Welcome to Gordle!")
    fmt.Printf("Enter your guess:\n ")
}
