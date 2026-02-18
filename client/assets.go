package client

import (
	_ "embed"
	"errors"
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
	var bufstr string
	for _, val := range requiredlist {
		// awful, needs rewrite
		if val != '\n' {
			bufstr += string(val)
		}
		if val == '\n' {
			var there bool
			for _, val := range assets {
				if val == bufstr {
					there = true
				}
			}
			if there == false {
				return nil, errors.New("")
			}
			bufstr = ""
		}
	}
	return assets, nil
}
