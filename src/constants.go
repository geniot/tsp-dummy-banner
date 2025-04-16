package main

import (
	"embed"
	"github.com/veandco/go-sdl2/sdl"
)

type ButtonCode = uint8
type ResourceKey = int

const (
	APP_NAME          = "TSP Dummy Placeholder App"
	APP_VERSION       = "0.1"
	TSP_SCREEN_WIDTH  = 1280
	TSP_SCREEN_HEIGHT = 720
)

var (
	//go:embed media/*
	mediaList embed.FS
)

var (
	COLOR_BLACK      = sdl.Color{R: 0, G: 0, B: 0, A: 255}
	COLOR_WHITE      = sdl.Color{R: 255, G: 255, B: 255, A: 255}
	BACKGROUND_COLOR = COLOR_BLACK
)

const (
	RESOURCE_BGR_KEY = ResourceKey(iota)
)

const (
	BUTTON_CODE_MENU  = ButtonCode(5)
	BUTTON_CODE_START = ButtonCode(6)
)
