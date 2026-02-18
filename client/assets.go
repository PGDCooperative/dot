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
		if strings.Contains(path, "/") && !strings.Contains(path, "fonts") {
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
			for _, val1 := range assets {
				if val1 == bufstr {
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
