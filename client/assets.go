package client

import (
	_ "embed"
	"io/fs"
	"path/filepath"
	"strings"
)

//go:embed assets.txt
var requiredlist []byte

func GetAssetsList(folder string) ([]string, error) {
	var assets []string
	filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
		if strings.Contains(path, "/") {
			assets = append(assets, path)
		}
		return nil
	})
	return assets, nil
}
