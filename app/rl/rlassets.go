package main

import (
	"dot/client"
	"errors"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type RLTextures map[string]rl.Texture2D

type RLFonts map[string]rl.Font

func PreloadTextures() (RLTextures, error) {
	rltextures := make(RLTextures)
	for _, val := range client.RequiredAssets {
		if strings.Contains(val, ".png") {
			rltextures[val] = rl.LoadTexture(val)
			if rltextures[val].ID <= 0 {
				return nil, errors.New("")
			}
		}
	}
	return rltextures, nil
}

func PreloadFonts() (RLFonts, error) {
	rlfonts := make(RLFonts)
	for _, val := range client.RequiredAssets {
		if strings.Contains(val, ".ttf") {
			rlfonts[val] = rl.LoadFontEx(val, 100, nil, 250)
			if rlfonts[val].BaseSize == 0 {
				return nil, errors.New("")
			}
		}
	}
	return rlfonts, nil
}
