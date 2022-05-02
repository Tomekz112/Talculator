package main

import (
	"errors"
	"fmt"
	"image"
	_ "image/png"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"

	"github.com/faiface/pixel"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var SessionID = 1

func GetSessionID() {
	data, err := os.ReadFile("Resources/session.id")
	if errors.Is(err, os.ErrNotExist) {
		_, err = os.Create("Resources/session.id")
		if err != nil {
			LogError("Creating file with session id error", err.Error(), false)
		}
	} else if err != nil {
		LogError("Reading file with session id error", err.Error(), false)
	} else {
		fmt.Println(string(data))

		SessionID, _ = strconv.Atoi(string(data))
		SessionID++
	}
	os.WriteFile("Resources/session.id", []byte(strconv.Itoa(SessionID)), 0644)
}

func LoadButtonSprite() *pixel.Sprite {
	picture, err := LoadPicture("Resources/Button.png")
	if err != nil {
		LogError("Couldn't load button sprite", err.Error(), false)
	}
	sprite := pixel.NewSprite(picture, picture.Bounds())
	return sprite
}

func LoadFont(fontName string,size float64) (font.Face, error) {
	file, err := os.Open("Resources/" + fontName + ".ttf")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}

func LoadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		LogError("Error while trying to open image", err.Error(), false)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		LogError("Error while trying to decode image", err.Error(), false)
	}
	return pixel.PictureDataFromImage(img), nil
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	LogMessage("Allocated Ram: " + strconv.Itoa(int(bToMb(m.Alloc))) + "MiB")
	LogMessage("Total memory obtained from OS: " + strconv.Itoa(int(bToMb(m.Sys))) + "MiB")
	LogMessage("Total garbage collected memory: " + strconv.Itoa(int(m.NumGC)) + "MiB")
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func OnClose() {
	PrintMemUsage()
	LogMessage("Closing program...")
}
