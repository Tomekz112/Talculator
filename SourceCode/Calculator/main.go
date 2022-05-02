package main

import (
	"fmt"
	"strings"

	"github.com/Knetic/govaluate"
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
var Operators = []string{"<", "(", ")", ".", "=", "clr", "1", "2", "3", "+", "-", "fbr", "4", "5", "6", "*", "/", "√", "7", "8", "9", "0", "^", "³√"}

var Values []float64

var atlas *text.Atlas
var MathExpression string
var window *pixelgl.Window
var expressionText *toxel.ToxelText //NewText("", pixel.V(15, 350), 4, Atlas)
var Buttons = []*toxel.Button{}
var userSettings Settings
var lastWinX float64
var lastWinY float64

func main() {
	GetSessionID()
	LogMessage("Started program, session id: " + strconv.Itoa(SessionID))
	userSettings = LoadConfig()
	face, err := LoadFont(userSettings.FontName, 52)
	if err != nil {
		LogError("LoadFont failed", err.Error(), true)
	}
	atlas = text.NewAtlas(face, text.ASCII, []rune{'√', '³', '∞'})
	expressionText = toxel.NewText("", pixel.V(15, 350), 1, atlas)
	pixelgl.Run(run)
}

func updateText() {
	expressionText.Text = MathExpression
	expressionText.AutoLineBreak(window.Bounds().Max.X)
}

func StandardizeExpression(mExpression string) string {
	mExpression = strings.Replace(mExpression, "^", "**", -1) //replace to the power of symbols
	mExpression = strings.Replace(mExpression, ",", ".", -1)  //replace dots symbols
	mExpression = strings.Replace(mExpression, "∞", "1/0", -1) //replace infinity
	mExpression = strings.Replace(mExpression, "-∞", "-1/0", -1) //replace -infinity
	return mExpression
}

func Calculate(mExpression string) {
	mExpression = StandardizeExpression(mExpression)
	LogMessage("Calculating: " + mExpression)
	expression, err := govaluate.NewEvaluableExpression(mExpression)
	if err != nil {
		return
	}
	result, err := expression.Evaluate(nil)
	if err != nil {
		return
	}
	MathExpression = fmt.Sprintf("%v", result)
	LogMessage("Result: " + MathExpression)
	MathExpression = strings.Replace(MathExpression, "+Inf", "∞", -1)
	MathExpression = strings.Replace(MathExpression, "-Inf", "-∞", -1)
	updateText()
}

func scaleButtons(Buttons []*toxel.Button) {
	position := pixel.V(window.Bounds().Max.X/2-112, window.Bounds().Max.Y/2+25)
	for j := 0; j < len(Buttons); j++ {
		position.X += 50
		if j%6 == 0 {
			position.Y -= window.Bounds().Max.Y / 8
			position.X = pixel.V(window.Bounds().Max.X/2-112, 225).X
		}
		Buttons[j].GameObject.Pos = position
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
	buttonSprite := LoadButtonSprite()
	for i := 0; i < len(Operators); i++ {
		Buttons = append(Buttons, toxel.NewButton(Operators[i], pixel.ZV, 1, atlas, i, buttonSprite, PressButton))
	}
	ticker := NewTicker(userSettings.MaxFPS)
	var Input = ""
	for !window.Closed() {
		if lastWinX != window.Bounds().Max.X || lastWinY != window.Bounds().Max.Y {
			lastWinY = window.Bounds().Max.Y
			lastWinX = window.Bounds().Max.X
			OnResolutionChange()
		}
		_, _ = ticker.Tick()
		window.Clear(userSettings.BackgroundColor)
		for i := 0; i < len(Buttons); i++ {
			Buttons[i].IsPressed(window)
			Buttons[i].Draw(window, userSettings.ButtonsColor)
		}
		expressionText.Draw(window, userSettings.ResultColor)
		Input = window.Typed()
		if window.JustPressed(pixelgl.KeyEnter) {
			HandleInput("=")
		}
		HandleInput(Input)
		window.Update()
		ticker.Wait()
	}
	OnClose()
}
