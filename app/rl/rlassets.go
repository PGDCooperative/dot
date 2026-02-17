package main

import (
	"dot/client"
	"errors"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type RLAssets map[string]rl.Texture2D

func Preload(folder string) (RLAssets, error) {
	assetslist, err := client.GetAssetsList(folder)
	if err != nil {
		return nil, err
	}
	rlassets := make(RLAssets)
	for _, val := range assetslist {
		rlassets[val] = rl.LoadTexture(val)
		if rlassets[val].ID <= 0 {
			return nil, errors.New("")
		}
	}
	return rlassets, nil
}
