package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *RenderRuntime) MainMenu() {
	rl.DrawTextEx(r.font, "Divide. Operate. Thrive.",
		rl.NewVector2(20.0, 150.0), 100.0, 1.0, rl.White)

	mousePos := rl.GetMousePosition()
	r.buttonstate["NewGame"].Draw(mousePos, func() {

	})
	r.buttonstate["LoadGame"].Draw(mousePos, func() {

	})
	r.buttonstate["Settings"].Draw(mousePos, func() {

	})
	r.buttonstate["QuitGame"].Draw(mousePos, func() {
		r.exitWindow = true
	})
}

func (r *RenderRuntime) SettingsMenu() {

}
