package main

import (
	"encoding/json"
	"errors"
	"image/color"
	"io/ioutil"
	"os"
)

type Settings struct {
	BackgroundColor color.RGBA
	ButtonsColor    color.RGBA
	ResultColor     color.RGBA
	ErrorColor      color.RGBA
	FontName        string
	Vsync           bool
	Discord         bool
	MaxFPS          int64
	Resolution      resolution
	ClearLogsAfter  int
}

type resolution struct {
	X, Y float64
}

func LoadConfig() Settings {
	var settings Settings
	data, err := ioutil.ReadFile("config.json")
	if errors.Is(err, os.ErrNotExist) {
		data = CreateConfig()
	} else if err != nil {
		LogError("Error while trying to read config", err.Error(), false)
	}
	LogMessage("Loaded config")
	json.Unmarshal(data, &settings)
	return settings
}

func CreateConfig() []byte {
	cfg, err := os.Create("config.json")
	if err != nil {
		LogError("Error while trying to create config", err.Error(), false)
	}
	_, err = cfg.Write([]byte(defaultConfigSettings))
	if err != nil {
		LogError("Error while trying to overwrite config", err.Error(), false)
	}
	LogMessage("Created new config")
	cfg.Close()
	return []byte(defaultConfigSettings)
}

var defaultConfigSettings = `{
    "BackgroundColor": {"R": 0, "G": 0, "B": 0, "A": 125},
    "ButtonsColor": {"R": 80, "G": 220, "B": 100, "A": 200},
    "ResultColor": {"R": 80, "G": 220, "B": 100, "A": 200},
    "ErrorColor": {"R": 255, "G": 0, "B": 0, "A": 255},
	"FontName": "Default",
    "Vsync": false,
	"Discord": true,
    "MaxFPS": 60,
    "Resolution": {"X": 300.0,"Y": 400.0}
}`
