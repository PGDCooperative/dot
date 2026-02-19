package main

import (
	"dot/client"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ButtonInterface interface {
	Draw(mousePos rl.Vector2)
}

type ButtonState map[string]ButtonInterface

type Button struct {
	pos   rl.Vector2
	color rl.Color
}

type TextButton struct {
	Button
	text string
	size float32
	font rl.Font
}

func GenButtonState(locale client.Locale, font *rl.Font) ButtonState {
	state := make(ButtonState)
	state["NewGame"] = &TextButton{
		Button: Button{rl.NewVector2(70.0, 270.0), rl.White},
		text:   locale.Get("#DOT_NEWGAME"),
		size:   50.0,
		font:   *font,
	}
	return state
}

func (b *TextButton) Draw(mousePos rl.Vector2) {
	rl.DrawTextEx(b.font, b.text, b.pos, b.size, 5.0, b.color)
	bounds := rl.NewRectangle(b.pos.X, b.pos.Y, 29.0*float32(len(b.text)), 50.0)
	btnState := 0
	if rl.CheckCollisionPointRec(mousePos, bounds) {
		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			btnState = 2
		} else {
			btnState = 1
		}
	}
	if btnState == 1 {
		b.color = rl.ColorLerp(b.color, rl.NewColor(b.color.R, b.color.G, b.color.B, 80), 0.5)
	} else {
		b.color = rl.ColorLerp(b.color, rl.NewColor(b.color.R, b.color.G, b.color.B, 255), 0.5)
	}
}
