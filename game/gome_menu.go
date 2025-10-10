package game

import (
	"fmt"
	"gome/system/menu"
)

var GameMenuTree = menu.NewSubmenu("Main Menu", []menu.MenuItem{
	menu.NewButton("Start Game", func() { fmt.Println("Starting game…") }),
	menu.NewSubmenu("Settings", []menu.MenuItem{
		menu.NewSlider("Audio", 0, 100, 50, func(val int) { fmt.Println("Audio:", val) }),
		menu.NewSlider("Brightness", 0, 100, 75, func(val int) { fmt.Println("Brightness:", val) }),
	}),
	menu.NewButton("Quit", func() { fmt.Println("Quitting…") }),
})
