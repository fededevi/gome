package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var BrightnessShader *ebiten.Shader

var brightnessShaderSrc = `
package main

var BRIGHTNESS float

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
    return vec4(color.r * BRIGHTNESS, color.g * BRIGHTNESS, color.b * BRIGHTNESS, color.a)
}
`

func LoadBrightnessShader() {
	var err error
	BrightnessShader, err = ebiten.NewShader([]byte(brightnessShaderSrc))
	if err != nil {
		panic("Failed to load brightness shader: " + err.Error())
	}
}
