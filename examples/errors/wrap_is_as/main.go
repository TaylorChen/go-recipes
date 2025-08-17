package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

type ParseError struct{ Offset int }

func (e *ParseError) Error() string { return fmt.Sprintf("parse error at %d", e.Offset) }

func fetchRecord() error { return ErrNotFound }

func serviceCall() error { return fmt.Errorf("service failed: %w", fetchRecord()) }

func parse() error { return &ParseError{Offset: 42} }

func wrapParse() error { return fmt.Errorf("wrapping: %w", parse()) }

func main() {
	if err := serviceCall(); err != nil {
		fmt.Println("got error:", err)
		if errors.Is(err, ErrNotFound) {
			fmt.Println("errors.Is matched ErrNotFound")
		}
	}

	if err := wrapParse(); err != nil {
		var pe *ParseError
		if errors.As(err, &pe) {
			fmt.Println("errors.As extracted ParseError with offset:", pe.Offset)
		}
	}
}
