package main

import (
	"os/exec"

	"github.com/faiface/pixel"

	"strconv"

	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

//layout:
//delete ( ) , =
//1 2 3 + -
//4 5 6 * /
//7 8 9 0 ^
var Operators = []string{"⌫", "(", ")", ".", "=", "clr", "1", "2", "3", "+", "-", "π", "4", "5", "6", "*", "/", "√", "7", "8", "9", "0", "^", "³√"}

var Values []float64

var atlas *text.Atlas
var window *pixelgl.Window
var userSettings Settings
var lastWinRes pixel.Vec

func main() {
	LogMessage("Started program")
	userSettings = LoadConfig()
	exec.Command("clearLogs", strconv.Itoa(userSettings.ClearLogsAfter))
	startCfg()
	CurrentExpression = &MathExpression
	pixelgl.Run(run)
}

func run() {
	icon, _ := LoadPicture("Resources/icon.png")
	cfg := pixelgl.WindowConfig{
		Title:                  "Talculator",
		Bounds:                 pixel.R(0, 0, userSettings.Resolution.X, userSettings.Resolution.Y),
		VSync:                  userSettings.Vsync,
		TransparentFramebuffer: true,
		Resizable:              true,
		Icon:                   []pixel.Picture{icon},
	}
	window, _ = pixelgl.NewWindow(cfg)
	window.SetSmooth(true)
	ticker := NewTicker(userSettings.MaxFPS)
	var Input = ""
	for !window.Closed() {
		LoadDiscordRichPresence()
		if window.Bounds().Max != pixel.ZV {
			go controlRes()
			_, _ = ticker.Tick()
			window.Clear(userSettings.BackgroundColor)
			DrawUIElements(window)
			//Handle Input
			Input = window.Typed()
			if window.JustPressed(pixelgl.KeyEnter) {
				HandleInput("=")
			}
			if window.Repeated(pixelgl.KeyBackspace) || window.JustPressed(pixelgl.KeyBackspace) {
				HandleInput("⌫")
			}
			HandleInput(Input)
			//
			// if TextActive {
			// 	textBox.Draw(window, pixel.IM.Moved(window.Bounds().Center().Sub(pixel.V(0, -30*scale.Y))).ScaledXY(window.Bounds().Center().Sub(pixel.V(0, -30*scale.Y)), scale))
			// 	Message.Pos = window.Bounds().Center().Sub(pixel.V(75*scale.X, -50*scale.Y))
			// 	Message.Scale(pixel.V(scale.X/5, scale.Y/5))
			// 	Message.Draw(window, userSettings.ErrorColor)
			// }
		}
		window.Update()
		ticker.Wait()
	}
	OnClose()
}
