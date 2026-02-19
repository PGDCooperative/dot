package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func MainMenu(renderRuntime *RenderRuntime) {
	rl.DrawTextEx(renderRuntime.font, "Divide. Operate. Thrive.",
		rl.NewVector2(20.0, 150.0), 100.0, 1.0, rl.White)

	mousePos := rl.GetMousePosition()
	renderRuntime.buttonstate["NewGame"].Draw(mousePos, func() {

	})
	renderRuntime.buttonstate["LoadGame"].Draw(mousePos, func() {

	})
	renderRuntime.buttonstate["Settings"].Draw(mousePos, func() {

	})
	renderRuntime.buttonstate["QuitGame"].Draw(mousePos, func() {

	})
}

func SettingsMenu(renderRuntime *RenderRuntime) {

}

func PauseMenu(renderRuntime *RenderRuntime) {

}
