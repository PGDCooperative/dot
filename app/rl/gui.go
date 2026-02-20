package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *RenderState) MainMenu() {
	rl.DrawTextEx(r.rlfonts["assets/fonts/opensans.ttf"], "Divide. Operate. Thrive.",
		rl.NewVector2(20.0, 150.0), 100.0, 1.0, rl.White)

	r.buttonstate["NewGame"].Draw(r.mousePos, func() {

	})
	r.buttonstate["LoadGame"].Draw(r.mousePos, func() {

	})
	r.buttonstate["Settings"].Draw(r.mousePos, func() {

	})
	r.buttonstate["QuitGame"].Draw(r.mousePos, func() {
		r.exitWindow = true
	})
}

func (r *RenderState) SettingsMenu() {

}
