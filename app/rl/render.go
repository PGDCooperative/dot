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

type RenderState struct {
	settings          client.Settings
	locale            client.Locale
	rltextures        RLTextures
	rlfonts           RLFonts
	uistate           client.UIState
	buttonstate       ButtonState
	camera            rl.Camera3D
	screenAspectRatio float32
	mousePos          rl.Vector2
	exitWindow        bool
}

func InitializeWindow(width int, height int, fullscreen bool) {
	rl.InitWindow(int32(width), int32(height), "D.O.T.")
	if fullscreen {
		rl.ToggleFullscreen()
	} else {
		rl.SetWindowState(rl.FlagWindowResizable)
	}
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
}

func RenderLoop(renderState RenderState) {
	for !renderState.exitWindow {
		if rl.WindowShouldClose() {
			renderState.exitWindow = true
		}
		renderState.settings.Width = rl.GetScreenWidth()
		renderState.settings.Height = rl.GetScreenHeight()
		if rl.IsKeyReleased(rl.KeyF11) {
			ToggleFullscreen(&renderState.settings.Fullscreen)
		}
		renderState.screenAspectRatio = float32(renderState.settings.Width) / float32(renderState.settings.Height)
		renderState.mousePos = rl.GetMousePosition()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.BeginMode3D(renderState.camera)
		if renderState.uistate["MainMenu"] {
			rl.DrawBillboard(renderState.camera,
				renderState.rltextures["assets/background.png"],
				rl.NewVector3(0.0, 0.0, 0.0), 45.0*PreserveAspectRatio(renderState.screenAspectRatio), rl.White)
		}
		rl.EndMode3D()

		if renderState.uistate["MainMenu"] {
			renderState.MainMenu()
		}
		if renderState.uistate["SettingsMenu"] {
			renderState.SettingsMenu()
		}

		rl.EndDrawing()
	}
	renderState.settings.WriteSettings()
	rl.CloseWindow()
}

func InitCamera() rl.Camera3D {
	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(0.0, 0.0, 0.5)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraOrthographic
	return camera
}

func PreserveAspectRatio(screenAspectRatio float32) float32 {
	if screenAspectRatio > gameAspectRatio {
		return screenAspectRatio / gameAspectRatio
	}
	return 1.0
}

func ToggleFullscreen(fullscreen *bool) {
	rl.ToggleFullscreen()
	*fullscreen = !*fullscreen
	if *fullscreen == false {
		rl.SetWindowState(rl.FlagWindowResizable)
	}
}
