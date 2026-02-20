package main

import (
	"dot/client"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ButtonInterface interface {
	Draw(mousePos rl.Vector2, cb func())
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

func GenButtonState(locale client.Locale, rlfonts RLFonts) ButtonState {
	state := make(ButtonState)
	state["NewGame"] = &TextButton{
		Button: Button{rl.NewVector2(70.0, 270.0), rl.White},
		text:   locale.Get("#DOT_NEWGAME"),
		size:   50.0,
		font:   rlfonts["assets/fonts/opensans.ttf"],
	}

	state["LoadGame"] = &TextButton{
		Button: Button{rl.NewVector2(70.0, 370.0), rl.White},
		text:   locale.Get("#DOT_LOADGAME"),
		size:   50.0,
		font:   rlfonts["assets/fonts/opensans.ttf"],
	}

	state["Settings"] = &TextButton{
		Button: Button{rl.NewVector2(70.0, 470.0), rl.White},
		text:   locale.Get("#DOT_SETTINGS"),
		size:   50.0,
		font:   rlfonts["assets/fonts/opensans.ttf"],
	}

	state["QuitGame"] = &TextButton{
		Button: Button{rl.NewVector2(70.0, 570.0), rl.White},
		text:   locale.Get("#DOT_QUITGAME"),
		size:   50.0,
		font:   rlfonts["assets/fonts/opensans.ttf"],
	}
	return state
}

func (b *TextButton) Draw(mousePos rl.Vector2, cb func()) {
	rl.DrawTextEx(b.font, b.text, b.pos, b.size, 1.0, b.color)
	size := rl.MeasureTextEx(b.font, b.text, b.size, 1.0)
	bounds := rl.NewRectangle(b.pos.X, b.pos.Y, size.X, 50.0)
	btnState := 0
	isPressed := false
	if rl.CheckCollisionPointRec(mousePos, bounds) {
		btnState = 1
		if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			isPressed = true
		}
	} else {
		btnState = 0
	}
	if btnState == 1 {
		b.color = rl.ColorLerp(b.color, rl.NewColor(b.color.R, b.color.G, b.color.B, 150), 0.3)
	} else {
		b.color = rl.ColorLerp(b.color, rl.NewColor(b.color.R, b.color.G, b.color.B, 255), 0.7)
	}
	if isPressed {
		cb()
	}
}
