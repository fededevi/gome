package game

import (
	"fmt"
	"gome/system/audio"
	"gome/system/events"
	"gome/system/menu"

	"github.com/hajimehoshi/ebiten/v2"
)

// QuitRequested is used to signal the game should exit
var QuitRequested bool
var Brightness *events.Property[int] = events.NewProperty[int](50)
var OnStartGame events.Event[any]

// Resolution represents a display resolution.
type Resolution struct {
	Label  string
	Width  int
	Height int
}

// AvailableResolutions defines all supported resolutions.
var AvailableResolutions = []Resolution{
	// --- 4:3 (legacy / low-end) ---
	{"640x480 (VGA - 4:3)", 640, 480},
	{"800x600 (SVGA - 4:3)", 800, 600},
	{"1024x768 (XGA - 4:3)", 1024, 768},
	{"1280x960 (SXGA - 4:3)", 1280, 960},
	{"1600x1200 (UXGA - 4:3)", 1600, 1200},

	// --- 16:10 (common on older laptops and monitors) ---
	{"1280x800 (WXGA - 16:10)", 1280, 800},
	{"1440x900 (WXGA+ - 16:10)", 1440, 900},
	{"1680x1050 (WSXGA+ - 16:10)", 1680, 1050},
	{"1920x1200 (WUXGA - 16:10)", 1920, 1200},
	{"2560x1600 (WQXGA - 16:10)", 2560, 1600},

	// --- 16:9 (standard modern monitors and TVs) ---
	{"1280x720 (HD - 16:9)", 1280, 720},
	{"1366x768 (HD Ready - 16:9)", 1366, 768},
	{"1600x900 (HD+ - 16:9)", 1600, 900},
	{"1920x1080 (Full HD - 16:9)", 1920, 1080},
	{"2560x1440 (QHD / 2K - 16:9)", 2560, 1440},
	{"3200x1800 (QHD+ - 16:9)", 3200, 1800},
	{"3840x2160 (4K UHD - 16:9)", 3840, 2160},
	{"5120x2880 (5K - 16:9)", 5120, 2880},
	{"7680x4320 (8K UHD - 16:9)", 7680, 4320},

	// --- Ultrawide 21:9 (gaming / productivity monitors) ---
	{"2560x1080 (UW FHD - 21:9)", 2560, 1080},
	{"3440x1440 (UW QHD - 21:9)", 3440, 1440},
	{"3840x1600 (UW WQHD - 21:9)", 3840, 1600},
	{"5120x2160 (UW 5K - 21:9)", 5120, 2160},

	// --- Super Ultrawide 32:9 (dual-screen aspect) ---
	{"3840x1080 (Dual FHD - 32:9)", 3840, 1080},
	{"5120x1440 (Dual QHD - 32:9)", 5120, 1440},
}

// helper: get string labels for the menu
func resolutionLabels() []string {
	labels := make([]string, len(AvailableResolutions))
	for i, res := range AvailableResolutions {
		labels[i] = res.Label
	}
	return labels
}

// helper: set resolution by label
func setResolution(label string) {
	for _, res := range AvailableResolutions {
		if res.Label == label {
			ebiten.SetWindowSize(res.Width, res.Height)
			fmt.Printf("Resolution changed to %s (%dx%d)\n", res.Label, res.Width, res.Height)
			return
		}
	}
	fmt.Println("Unknown resolution:", label)
}

// CreateGameMenu creates the main game menu
func CreateGameMenu(audioSys *audio.AudioSystem) *menu.Submenu {
	return menu.NewSubmenu("Main Menu", []menu.MenuItem{
		menu.NewButton("Start Game", func() {
			fmt.Println("Starting game menu button pressed…")
			OnStartGame.Emit(nil)
		}),

		menu.NewSubmenu("Settings", []menu.MenuItem{
			menu.NewSubmenu("Audio", []menu.MenuItem{
				menu.NewSlider("Master Volume", 0, 10, 10, func(val int) {
					volume := float64(val) / 10.0
					audioSys.Master = volume
					fmt.Println("Master Volume set to", val)
					audioSys.Play(ClickSound)
				}),
				menu.NewSlider("Effects Volume", 1, 10, int(audioSys.Effects.GetVolume()*10), func(val int) {
					audioSys.Effects.SetVolume(float64(val) / 10.0)
					fmt.Println("Effects Channel Volume set to", val)
					audioSys.Play(ClickSound)
				}),
				menu.NewSlider("Music Volume", 1, 10, int(audioSys.Music.GetVolume()*10), func(val int) {
					audioSys.Music.SetVolume(float64(val) / 10.0)
					fmt.Println("Music Channel Volume set to", val)
					audioSys.Play(ClickSound)
				}),
			}),

			menu.NewSubmenu("Video", []menu.MenuItem{
				menu.NewSlider("Brightness", 0, 100, Brightness.Get(), func(val int) {
					Brightness.Set(val)
					fmt.Println("Brightness set to", val)
					audioSys.Play(ClickSound)
				}),
				menu.NewEnumItem("Resolution", resolutionLabels(), 1, func(val string) {
					audioSys.Play(ClickSound)
					setResolution(val)
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
						w, h := ebiten.Monitor().Size()
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
