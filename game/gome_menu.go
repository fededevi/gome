package game

import (
	"fmt"
	"gome/system/audio"
	"gome/system/menu"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	QuitRequested bool
)

func CreateGameMenu(audioSys *audio.AudioSystem) *menu.Submenu {

	return menu.NewSubmenu("Main Menu", []menu.MenuItem{
		menu.NewButton("Start Game", func() {
			fmt.Println("Starting game…")
		}),

		menu.NewSubmenu("Settings", []menu.MenuItem{
			menu.NewSubmenu("Audio", []menu.MenuItem{
				menu.NewSlider("Master Volume", 0, 10, 10, func(val int) {
					volume := float64(val) / 10.0
					audioSys.Master = volume
					fmt.Println("Master Volume set to", val)
					audioSys.Play(ClickSound)
				}),

				// Effects channel volume
				menu.NewSlider("Effects Volume", 1, 10, int(audioSys.Effects.GetVolume()*10), func(val int) {
					audioSys.Effects.SetVolume(float64(val) / 10.0)
					fmt.Println("Effects Channel Volume set to", val)
					audioSys.Play(ClickSound)
				}),

				// Music channel volume
				menu.NewSlider("Music Volume", 1, 10, int(audioSys.Music.GetVolume()*10), func(val int) {
					audioSys.Music.SetVolume(float64(val) / 10.0)
					fmt.Println("Music Channel Volume set to", val)
					audioSys.Play(ClickSound)
				}),
			}),
			menu.NewSubmenu("Video", []menu.MenuItem{
				menu.NewSlider("Brightness", 0, 100, 50, func(val int) {
					fmt.Println("Brightness set to", val)
					audioSys.Play(ClickSound)
				}),
				menu.NewEnumItem("Resolution", []string{"800x600", "1280x720", "1920x1080"}, 1, func(val string) {
					fmt.Println("Resolution changed to", val)
					audioSys.Play(ClickSound)
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
					audioSys.Play(ClickSound)
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

		menu.NewButton("Quit", func() {
			fmt.Println("Quitting…")
			audioSys.Play(ClickSound)
			QuitRequested = true
		}),
	})
}
