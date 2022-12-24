package main

import "github.com/gdamore/tcell/v2"

type InputEvent int

const (
	GoUp InputEvent = iota
	GoDown
	GoHome
	GoEnd
	GoPageUp
	GoPageDown
	GoForward
	GoBack
	GoQuit
	OnResize
)

const KeyInputOffset InputEvent = 256

func (t InputEvent) HasKey() bool {
	return t >= KeyInputOffset
}

func (t InputEvent) Key() rune {
	return rune(t - KeyInputOffset)
}

func KeyInputEvent(key rune) InputEvent {
	return InputEvent(key) + KeyInputOffset
}

type Input struct {
	screen tcell.Screen
}

func NewInput(screen tcell.Screen) *Input {
	return &Input{screen}
}

func (i *Input) PoolEvent() InputEvent {
	for {
		event := i.screen.PollEvent()
		switch event := event.(type) {
		case *tcell.EventResize:
			return OnResize
		case *tcell.EventKey:
			switch event.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				return GoQuit
			case tcell.KeyUp:
				return GoUp
			case tcell.KeyDown:
				return GoDown
			case tcell.KeyHome:
				return GoHome
			case tcell.KeyEnd:
				return GoEnd
			case tcell.KeyPgUp:
				return GoPageUp
			case tcell.KeyPgDn:
				return GoPageDown
			case tcell.KeyEnter, tcell.KeyRight:
				return GoForward
			case tcell.KeyLeft, tcell.KeyBackspace, tcell.KeyBackspace2:
				return GoBack
			}
			key := event.Rune()
			return KeyInputEvent(key)
		}
	}
}
