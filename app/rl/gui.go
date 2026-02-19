package main

import (
	"dot/client"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func MainMenu(uistate *client.UIState, buttonstate ButtonState, font rl.Font) {
	rl.DrawTextEx(font, "Divide. Operate. Thrive.",
		rl.NewVector2(20.0, 150.0), 100.0, 1.0, rl.White)
	mousePos := rl.GetMousePosition()

	buttonstate["NewGame"].Draw(mousePos)

	//DrawTextButton(locale.Get("#DOT_NEWGAME"),
	//	rl.NewVector2(70.0, 270.0), 50.0, font, rl.White)
	/*
	   rl.DrawTextEx(font, locale.Get("#DOT_NEWGAME"),

	   	rl.NewVector2(70.0, 250.0), 50.0, 1.0, rl.White)

	   rl.DrawTextEx(font, locale.Get("#DOT_LOADGAME"),

	   	rl.NewVector2(70.0, 350.0), 50.0, 1.0, rl.White)

	   rl.DrawTextEx(font, locale.Get("#DOT_SETTINGS"),

	   	rl.NewVector2(70.0, 450.0), 50.0, 1.0, rl.White)

	   rl.DrawTextEx(font, locale.Get("#DOT_QUITGAME"),

	   	rl.NewVector2(70.0, 550.0), 50.0, 1.0, rl.White)
	*/
}

func SettingsMenu(uistate *client.UIState, buttonstate ButtonState, settings *client.Settings) {

}

func PauseMenu(uistate *client.UIState, buttonstate ButtonState) {

}
