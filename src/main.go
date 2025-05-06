package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"time"
)

const (
	//upCode     = 1
	//rightCode  = 2
	//downCode   = 3
	//leftCode   = 4
	//xCode      = 5
	//aCode      = 6
	//bCode      = 7
	//yCode      = 8
	//l1Code     = 9
	//l2Code     = 10
	//r1Code     = 11
	//r2Code     = 12
	//selectCode = 13
	menuCode  = 14
	startCode = 15
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "TrimUI Placeholder App")

	var (
		titleMessage         = "Placeholder App"
		titleMessageFontSize = float32(100)
		titleMessageSpacing  = float32(10)
		titleMessageSize     = rl.MeasureTextEx(rl.GetFontDefault(), titleMessage, float32(titleMessageFontSize), titleMessageSpacing)

		exitMessage         = "MENU+START => Exit"
		exitMessageFontSize = float32(40)
		exitMessageSpacing  = float32(3)
		exitMessagePadding  = float32(10)
		exitMessageSize     = rl.MeasureTextEx(rl.GetFontDefault(), exitMessage, float32(exitMessageFontSize), exitMessageSpacing)

		gamePadId  = int32(0)
		shouldExit = false
	)

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.DrawTextEx(rl.GetFontDefault(), titleMessage, rl.Vector2{X: (screenWidth - titleMessageSize.X) / 2, Y: (screenHeight - titleMessageSize.Y) / 2}, titleMessageFontSize, titleMessageSpacing, rl.DarkBlue)
	rl.DrawTextEx(rl.GetFontDefault(), exitMessage, rl.Vector2{X: screenWidth - exitMessageSize.X - exitMessagePadding, Y: screenHeight - exitMessageSize.Y - exitMessagePadding}, exitMessageFontSize, exitMessageSpacing, rl.DarkBlue)
	rl.EndDrawing()

	for !rl.WindowShouldClose() && !shouldExit {
		rl.PollInputEvents()
		//exit
		if rl.IsKeyPressed(rl.KeyQ) || (rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode)) {
			shouldExit = true
		}
		time.Sleep(1_000_000_000 / 60)
	}
	rl.CloseWindow()
}
