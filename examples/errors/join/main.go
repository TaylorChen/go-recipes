package main

import (
	"errors"
	"fmt"
)

var ErrA = errors.New("A")
var ErrB = errors.New("B")

func main() {
	joined := errors.Join(ErrA, fmt.Errorf("wrap: %w", ErrB))
	fmt.Println("joined:", joined)
	fmt.Println("is ErrA:", errors.Is(joined, ErrA))
	fmt.Println("is ErrB:", errors.Is(joined, ErrB))
}
