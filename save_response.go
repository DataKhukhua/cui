// cui: http request/response tui
// Copyright 2022 Mario Finelli
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"os"
	"path/filepath"
	"strings"
)

func SaveResponseFile(filePath string, responseBody string) error {

	var absFilePath string = filePath

	if strings.HasPrefix(filePath, "~/") {
		dirname, _ := os.UserHomeDir()
		absFilePath = filepath.Join(dirname, filePath[2:])

	}

	// relative paths not supported

	_, err := os.Stat(absFilePath)

	if err == nil {
		return os.ErrExist
	}

	f, err := os.Create(absFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(responseBody)
	if err != nil {
		return err
	}

	return nil

}

func ReplaceSaveResponseFile(filePath string, responseBody string) error {

	var absFilePath string = filePath

	if strings.HasPrefix(filePath, "~/") {
		dirname, _ := os.UserHomeDir()
		absFilePath = filepath.Join(dirname, filePath[2:])
	}

	f, err := os.OpenFile(absFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(responseBody)

	if err != nil {
		return err
	}

	return nil

}
