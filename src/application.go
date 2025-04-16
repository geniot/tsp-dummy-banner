package main

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/tevino/abool/v2"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"os"
	"strconv"
)

type Application struct {
	settings           *Settings
	resources          map[ResourceKey]*SurfTexture
	sdlWindow          *sdl.Window
	sdlRenderer        *sdl.Renderer
	sdlGameController  *sdl.GameController
	font               *ttf.Font
	pressedKeysCodes   mapset.Set[sdl.Keycode]
	pressedButtonCodes mapset.Set[ButtonCode]
	isRunning          *abool.AtomicBool
}

func NewApplication() *Application {
	return &Application{
		pressedKeysCodes:   mapset.NewSet[sdl.Keycode](),
		pressedButtonCodes: mapset.NewSet[ButtonCode](),
		isRunning:          abool.New(),
		resources:          make(map[int]*SurfTexture),
	}
}

func (app *Application) Start(args []string) {
	var err error

	if err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_JOYSTICK | sdl.INIT_GAMECONTROLLER); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	if err := ttf.Init(); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	if app.font, err = ttf.OpenFontRW(LoadMediaFile("pixelberry.ttf"), 1, 40); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	for i := 0; i < sdl.NumJoysticks(); i++ {
		if sdl.IsGameController(i) {
			app.sdlGameController = sdl.GameControllerOpen(i)
		}
	}

	app.settings = NewSettings()
	if app.sdlWindow, err = sdl.CreateWindow(
		APP_NAME+" "+APP_VERSION,
		int32(app.settings.WindowPosX), int32(app.settings.WindowPosY),
		int32(app.settings.WindowWidth), int32(app.settings.WindowHeight),
		uint32(app.settings.WindowState)); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	if app.sdlRenderer, err = sdl.CreateRenderer(app.sdlWindow, -1, sdl.RENDERER_PRESENTVSYNC|sdl.RENDERER_ACCELERATED); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	app.initResources() //should be called after the creation of sdlRenderer

	app.UpdateView()

	app.isRunning.Set()
	for app.isRunning.IsSet() {
		app.UpdateEvents()
		app.UpdatePhysics()

	}
}

func (app *Application) Stop() {
	app.isRunning.UnSet()
	app.settings.Save(app.sdlWindow)
	app.font.Close()
	ttf.Quit()
	sdl.Quit()
}

func (app *Application) UpdateEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {

		case *sdl.ControllerButtonEvent:
			if t.State == sdl.PRESSED {
				println(t.Button)
				app.pressedButtonCodes.Add(t.Button)
			} else {
				app.pressedButtonCodes.Remove(t.Button)
			}
			break

		case *sdl.KeyboardEvent:
			if t.Repeat > 0 {
				break
			}
			if t.State == sdl.PRESSED {
				app.pressedKeysCodes.Add(t.Keysym.Sym)
			} else { // if t.State == sdl.RELEASED {
				app.pressedKeysCodes.Remove(t.Keysym.Sym)
			}
			break

		case *sdl.WindowEvent:
			if t.Event == sdl.WINDOWEVENT_CLOSE {
				app.settings.SaveWindowState(app.sdlWindow)
			}
			break

		case *sdl.QuitEvent:
			app.Stop()
			break
		}
	}
}

func (app *Application) UpdatePhysics() {
	if app.pressedKeysCodes.Contains(sdl.K_q) || (app.pressedButtonCodes.Contains(BUTTON_CODE_MENU) && app.pressedButtonCodes.Contains(BUTTON_CODE_START)) {
		app.Stop()
	}
}

func (app *Application) UpdateView() {
	if err := app.sdlRenderer.SetDrawColorArray(BACKGROUND_COLOR.R, BACKGROUND_COLOR.G, BACKGROUND_COLOR.B, BACKGROUND_COLOR.A); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	if err := app.sdlRenderer.Clear(); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	if err := app.sdlRenderer.Copy(app.resources[RESOURCE_BGR_KEY].T, nil, &sdl.Rect{X: 0, Y: 0, W: int32(app.settings.WindowWidth), H: int32(app.settings.WindowHeight)}); err != nil {
		println(err.Error())
	}

	if !app.pressedButtonCodes.IsEmpty() {
		for val := range app.pressedButtonCodes.Iter() {
			app.drawText(val)
		}
	}

	app.sdlRenderer.Present()
}

func (app *Application) initResources() {
	app.resources[RESOURCE_BGR_KEY] = LoadSurfTexture("bgr.png", app.sdlRenderer)
}

func (app *Application) drawText(val ButtonCode) {
	textSurface, _ := app.font.RenderUTF8Blended(strconv.FormatUint(uint64(val), 10), COLOR_WHITE)
	defer textSurface.Free()
	textTexture, _ := app.sdlRenderer.CreateTextureFromSurface(textSurface)
	app.sdlRenderer.Copy(textTexture, nil,
		&sdl.Rect{X: 0, Y: 0, W: textSurface.W, H: textSurface.H})
	defer textTexture.Destroy()
}
