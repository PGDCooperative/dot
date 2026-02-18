package main

import (
	"dot/client"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	gameWidth       = 1920
	gameHeight      = 1080
	gameAspectRatio = float32(gameWidth) / float32(gameHeight)
)

func InitializeWindow(width int, height int, fullscreen bool) {
	rl.InitWindow(int32(width), int32(height), "D.O.T.")
	if fullscreen {
		rl.ToggleFullscreen()
	} else {
		rl.SetWindowState(rl.FlagWindowResizable)
	}
}

func RenderLoop(rlassets RLAssets, settings client.Settings, locale client.Locale) error {
	uistate := client.UIState{MainMenu: true}

	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(0.0, 0.0, 0.5)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraOrthographic

	//font := rl.LoadFont("assets/fonts/opensans.ttf")
	font := rl.LoadFontEx("assets/fonts/opensans.ttf", 100, nil, 250)

	for !rl.WindowShouldClose() {
		screenWidth := rl.GetScreenWidth()
		screenHeight := rl.GetScreenHeight()
		screenAspectRatio := float32(screenWidth) / float32(screenHeight)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.BeginMode3D(camera)
		if uistate.MainMenu {
			rl.DrawBillboard(camera,
				rlassets["assets/background.png"],
				rl.NewVector3(0.0, 0.0, 0.0), 45.0*PreserveAspectRatio(screenAspectRatio), rl.White)
		}
		rl.EndMode3D()
		if uistate.MainMenu {
			MainMenu(&uistate, locale, &font)
		}
		if uistate.SettingsMenu {
			SettingsMenu(&uistate, &settings, locale, &font)
		}
		if uistate.PauseMenu {
			PauseMenu(&uistate, locale, &font)
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
	return nil
}

func PreserveAspectRatio(screenAspectRatio float32) float32 {
	if screenAspectRatio > gameAspectRatio {
		return screenAspectRatio / gameAspectRatio
	}
	return 1.0
}
