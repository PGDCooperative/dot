package main

import (
	"dot/client"
	"errors"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	gameWidth       = 1920
	gameHeight      = 1080
	gameAspectRatio = float32(gameWidth) / float32(gameHeight)
)

type RenderRuntime struct {
	rlassets          RLAssets
	settings          client.Settings
	locale            client.Locale
	uistate           client.UIState
	buttonstate       ButtonState
	font              rl.Font
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

func RenderLoop(rlassets RLAssets, settings client.Settings, locale client.Locale) error {
	font := rl.LoadFontEx("assets/fonts/opensans.ttf", 100, nil, 250)
	if font.BaseSize == 0 {
		return errors.New("Failed to load fonts!")
	}

	//megastructure
	renderRuntime := RenderRuntime{
		rlassets:    rlassets,
		settings:    settings,
		locale:      locale,
		uistate:     client.GetUIState(),
		buttonstate: GenButtonState(locale, &font),
		font:        font,
		camera:      InitCamera(),
	}

	for !renderRuntime.exitWindow {
		if rl.WindowShouldClose() {
			renderRuntime.exitWindow = true
		}
		renderRuntime.settings.Width = rl.GetScreenWidth()
		renderRuntime.settings.Height = rl.GetScreenHeight()
		if rl.IsKeyReleased(rl.KeyF11) {
			ToggleFullscreen(&renderRuntime.settings.Fullscreen)
		}
		renderRuntime.screenAspectRatio = float32(renderRuntime.settings.Width) / float32(renderRuntime.settings.Height)
		renderRuntime.mousePos = rl.GetMousePosition()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.BeginMode3D(renderRuntime.camera)
		if renderRuntime.uistate["MainMenu"] {
			rl.DrawBillboard(renderRuntime.camera,
				rlassets["assets/background.png"],
				rl.NewVector3(0.0, 0.0, 0.0), 45.0*PreserveAspectRatio(renderRuntime.screenAspectRatio), rl.White)
		}
		rl.EndMode3D()

		if renderRuntime.uistate["MainMenu"] {
			renderRuntime.MainMenu()
		}
		if renderRuntime.uistate["SettingsMenu"] {
			renderRuntime.SettingsMenu()
		}

		rl.EndDrawing()
	}
	err := renderRuntime.settings.WriteSettings()
	if err != nil {
		return err
	}
	rl.CloseWindow()
	return nil
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
