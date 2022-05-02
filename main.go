package main

import (
	"image/color"
	"os/exec"
	"strings"

	"github.com/faiface/pixel"

	"MyPkgs/toxel"

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
var expressionText *toxel.ToxelText //NewText("", pixel.V(15, 350), 4, Atlas)
var Buttons = []*toxel.Button{}
var userSettings Settings
var lastWinRes pixel.Vec

func main() {
	LogMessage("Started program")
	userSettings = LoadConfig()
	exec.Command("clearLogs", strconv.Itoa(userSettings.ClearLogsAfter))
	face, err := LoadFont(userSettings.FontName, 52)
	if err != nil {
		LogError("LoadFont failed", err.Error(), true)
	}
	atlas = text.NewAtlas(face, text.ASCII, []rune{'√', '³', '∞', '⌫', 'π'})
	expressionText = toxel.NewText("", pixel.V(15, 350), pixel.V(1, 1), atlas)
	Message = toxel.NewText("", pixel.ZV, pixel.V(1, 1), atlas)
	pixelgl.Run(run)
}

func updateText() {
	expressionText.Text = MathExpression.ToString()
	expressionText.AutoLineBreak(window.Bounds().Max.X - 20)
	//Check if there are too many characters
	str := strings.SplitAfter(expressionText.Text, "\n")
	if len(str) > 2 {
		expressionText.Text = str[0] + str[1]
		MathExpression.Set(strings.Trim(str[0]+str[1], "\n"))
	}
}

func scaleButtons(Buttons []*toxel.Button) {
	nextX := window.Bounds().Max.X / 6
	nextY := (window.Bounds().Max.Y - 50) / 5
	position := pixel.V(0, window.Bounds().Max.Y-75)
	for j := 0; j < len(Buttons); j++ {
		position.X += nextX
		if j%6 == 0 {
			position.Y -= nextY
			position.X = nextX / 2
		}
		Buttons[j].SetPosition(position)
	}
}

func OnResolutionChange() {
	updateText()
	expressionText.Pos.Y = window.Bounds().Max.Y - 50
	scaleButtons(Buttons)
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
	buttonSprites := LoadButtonSprites()
	// textBox := LoadSprite("Resources/textBox.png")
	for i := 0; i < len(Operators); i++ {
		Buttons = append(Buttons, toxel.NewSpriteButton(buttonSprites[i], pixel.ZV, pixel.V(1, 1), i, PressButton))
	}
	// dummySprite := Dummy()
	ticker := NewTicker(userSettings.MaxFPS)
	var Input = ""
	var scale = pixel.V(1, 1)
	for !window.Closed() {
		if userSettings.Discord {
			LoadDiscordRichPresence()
		}
		if window.Bounds().Max != pixel.ZV {
			go func() {
				if lastWinRes.X != window.Bounds().Max.X || lastWinRes.Y != window.Bounds().Max.Y {
					if window.Bounds().Max.X < 350 {
						window.SetBounds(pixel.Rect{
							Min: pixel.ZV,
							Max: pixel.V(350, window.Bounds().Max.Y),
						})
					}
					if window.Bounds().Max.Y < 380 {
						window.SetBounds(pixel.Rect{
							Min: pixel.ZV,
							Max: pixel.V(window.Bounds().Max.X, 380),
						})
					}
					scale.X = window.Bounds().Max.X / 300
					scale.Y = window.Bounds().Max.Y / 400
					lastWinRes.Y = window.Bounds().Max.Y
					lastWinRes.X = window.Bounds().Max.X
					OnResolutionChange()
				}
			}()
			_, _ = ticker.Tick()
			window.Clear(userSettings.BackgroundColor)
			for i := 0; i < len(Buttons); i++ {
				Buttons[i].IsPressed(window)
				Buttons[i].DrawColor(window, userSettings.ButtonsColor)
			}
			expressionText.SplitTextDraw(window, []color.Color{userSettings.NumbersColor, userSettings.AddTextColor}, []int{len(expressionText.Text) - len(MathExpression.Brackets())})
			Input = window.Typed()
			if window.JustPressed(pixelgl.KeyEnter) {
				HandleInput("=")
			}
			if window.Repeated(pixelgl.KeyBackspace) || window.JustPressed(pixelgl.KeyBackspace) {
				HandleInput("⌫")
			}
			if Input != "" {
				HandleInput(Input)
			}
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

//350 - X
//380 - Y
