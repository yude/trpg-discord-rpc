package main

import (
	"time"
	"log"
	
	// Discord Rich Presence Library
	"github.com/hugolgst/rich-go/client"
	
	// Fyne.io toolkit
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Get initial time for Discord RPC play time
	now := time.Now()
	
	// Set Discord game app ID
	DiscordAppID := "889095426354061332"
	// Login to Discord
	errLogin := client.Login(DiscordAppID)
	if errLogin != nil {
		panic(errLogin)
	}
	// Start RPC activity
	client.SetActivity(client.Activity{
		State:      "プレイ中",
		LargeText:  "trpg-discord-rpc",
		Timestamps: &client.Timestamps{
			Start: &now,
		},
	})

	// Init Fyne.io toolkit
	a := app.New()
	a.Settings().SetTheme(&myTheme{}) // Load theme for CJK font rendering
	w := a.NewWindow("Discord RPC for TRPG")
	w.Resize(fyne.NewSize(300, 200)) // Set window resolution
	
	// Form for scenario title input
	scenario := widget.NewEntry()
	scenario.SetPlaceHolder("プレイ中のシナリオ名をここに入力...")
	scenarioForm := container.NewVBox(scenario, widget.NewButton("決定", func() {
		now = time.Now()
		// Refresh RPC
		if (scenario.Text == "") {
			client.SetActivity(client.Activity{
				State:      "プレイ中",
				LargeText:  "trpg-discord-rpc",
				Timestamps: &client.Timestamps{
					Start: &now,
				},
			})
		} else {
			client.SetActivity(client.Activity{
				State:      "プレイ中",
				Details:    "シナリオ: 「" + scenario.Text + "」",
				LargeText:  "trpg-discord-rpc",
				Timestamps: &client.Timestamps{
				Start: &now,
			},
		})
		}
		log.Println("[INFO] プレイ中のシナリオを", scenario.Text, "に設定しました。\n")
	}))

	// Show Fyne.io toolkit window
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			layout.NewSpacer(),
			widget.NewLabel("シナリオ"),
			scenarioForm,
			layout.NewSpacer(),
		),
	)
	w.ShowAndRun()
}