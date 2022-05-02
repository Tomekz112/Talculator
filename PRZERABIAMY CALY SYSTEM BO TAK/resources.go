package main

import (
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

func LoadButtonSprites() []*pixel.Sprite {
	spritesheet, err := LoadPicture("Resources/Buttons.png")
	if err != nil {
		LogError("Couldn't load button sprites", err.Error(), false)
	}
	var buttonFrames []pixel.Rect
	for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += 40 {
		for y := spritesheet.Bounds().Max.Y; y > spritesheet.Bounds().Min.Y; y -= 64 {
			buttonFrames = append(buttonFrames, pixel.R(x, y, x+40, y-64))
			y--
		}
		x++
	}
	var sprites []*pixel.Sprite
	for i := 0; i < len(buttonFrames); i++ {
		sprites = append(sprites, pixel.NewSprite(spritesheet, buttonFrames[i]))
	}
	return sprites
}

func LoadSprite(path string) *pixel.Sprite {
	pic, err := LoadPicture(path)
	if err != nil {
		LogError("Couldn't load "+path+" sprite", err.Error(), false)
	}
	return pixel.NewSprite(pic, pic.Bounds())
}
func Dummy() *pixel.Sprite {
	picture, err := LoadPicture("Resources/Hitbox.png")
	if err != nil {
		LogError("Couldn't load button sprite", err.Error(), false)
	}
	sprite := pixel.NewSprite(picture, picture.Bounds())
	return sprite
}

func LoadFont(fontName string, size float64) (font.Face, error) {
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
