package main

import (
	"MyPkgs/toxel"
	"image/color"
	"strings"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

var scale = pixel.V(1, 1)
var expressionText *toxel.ToxelText //NewText("", pixel.V(15, 350), 4, Atlas)
var Buttons = []*toxel.Button{}

func startCfg() {
	//Load font & font face
	face, err := LoadFont(userSettings.FontName, 52)
	if err != nil {
		LogError("LoadFont failed", err.Error(), true)
	}
	atlas = text.NewAtlas(face, text.ASCII, []rune{'√', '³', '∞', '⌫', 'π'})

	expressionText = toxel.NewText("", pixel.V(15, 350), pixel.V(1, 1), atlas)
	Message = toxel.NewText("", pixel.ZV, pixel.V(1, 1), atlas)

	//Load buttons
	buttonSprites := LoadButtonSprites()
	for i := 0; i < len(Operators); i++ {
		Buttons = append(Buttons, toxel.NewSpriteButton(buttonSprites[i], pixel.ZV, pixel.V(1, 1), i, PressButton))
	}
}

func controlRes() {
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
}

func OnResolutionChange() {
	updateText()
	expressionText.Pos.Y = window.Bounds().Max.Y - 50
	scaleButtons(Buttons)
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

func DrawUIElements(win *pixelgl.Window) {
	for i := 0; i < len(Buttons); i++ {
		Buttons[i].IsPressed(window)
		Buttons[i].DrawColor(window, userSettings.ButtonsColor)
	}
	expressionText.SplitTextDraw(window, []color.Color{userSettings.NumbersColor, userSettings.AddTextColor}, []int{len(expressionText.Text) - len(MathExpression.Brackets())})
}
