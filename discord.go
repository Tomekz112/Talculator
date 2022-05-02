package main

import (
	"time"

	"github.com/hugolgst/rich-go/client"
)

var timeNow = time.Now()

func LoadDiscordRichPresence() {
	err := client.Login("923727747346497546")
	if err != nil {
		userSettings.Discord = false
		LogError("Discord Presence error", err.Error(), false)
		LogMessage("Please ensure that discord is opened!\nTemporarily disabling discord rich presence")
		return
	}

	err = client.SetActivity(client.Activity{
		State:      MathExpression.ToString(),
		Details:    "Calculating something",
		LargeImage: "icondsc",
		Timestamps: &client.Timestamps{
			Start: &timeNow,
		},
		Buttons: []*client.Button{
			&client.Button{
				Label: "GitHub",
				Url:   "https://github.com/Tomekz112",
			},
		},
	})

	if err != nil {
		panic(err)
	}
}
