package game

import (
	"fmt"
	"gome/system/menu"

	"github.com/hajimehoshi/ebiten/v2"
)

var QuitRequested bool = false

// GameMenuTree defines the full menu
var GameMenuTree = menu.NewSubmenu("Main Menu", []menu.MenuItem{
	// Start Game button
	menu.NewButton("Start Game", func() { fmt.Println("Starting game…") }),

	// Settings submenu
	menu.NewSubmenu("Settings", []menu.MenuItem{
		// Audio submenu
		menu.NewSubmenu("Audio", []menu.MenuItem{
			menu.NewSlider("Master Volume", 0, 100, 75, func(val int) {
				fmt.Println("Master Volume set to", val)
				// Here you could set your audio system volume
				// e.g., audio.SetMasterVolume(float64(val)/100)
			}),
		}),

		// Video submenu
		menu.NewSubmenu("Video", []menu.MenuItem{
			menu.NewSlider("Brightness", 0, 100, 70, func(val int) {
				fmt.Println("Brightness set to", val)
				// You could store brightness in a global or config for your renderer
			}),
			menu.NewEnumItem("Resolution", []string{"800x600", "1280x720", "1920x1080"}, 1, func(val string) {
				fmt.Println("Resolution changed to", val)
				// Apply window size change
				switch val {
				case "800x600":
					ebiten.SetWindowSize(800, 600)
				case "1280x720":
					ebiten.SetWindowSize(1280, 720)
				case "1920x1080":
					ebiten.SetWindowSize(1920, 1080)
				}
			}),
			menu.NewEnumItem("Rendering Mode", []string{"Windowed", "Borderless", "Fullscreen"}, 0, func(val string) {
				fmt.Println("Rendering mode changed to", val)
				switch val {
				case "Windowed":
					ebiten.SetFullscreen(false)
					ebiten.SetWindowDecorated(true)
				case "Borderless":
					ebiten.SetFullscreen(false)
					ebiten.SetWindowDecorated(false)
					w, h := ebiten.Monitor().Size() // get current monitor resolution
					ebiten.SetWindowSize(w, h)
					ebiten.SetWindowPosition(0, 0)
				case "Fullscreen":
					ebiten.SetFullscreen(true)
				}
			}),
		}),
	}),

	// Quit button
	menu.NewButton("Quit", func() {
		fmt.Println("Quitting…")
		QuitRequested = true
	}),
})
