package client

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func GetAssetsList(folder string) []string {
	var assets []string
	filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
		if strings.Contains(path, "/") {
			assets = append(assets, path)
		}
		return nil
	})
	return assets
}
