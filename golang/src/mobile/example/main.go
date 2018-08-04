package main

import (
	"time"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/exp/sprite"
	"golang.org/x/mobile/exp/sprite/clock"
	"golang.org/x/mobile/exp/sprite/glsprite"
	"golang.org/x/mobile/gl"
)

// See http://klabgames.tech.blog.jp.klab.com/archives/1047363250.html

func main() {
	app.Main(func(a app.App) {
		var glctx gl.Context
		var sz size.Event
		for e := range a.Events() {
			switch ev := a.Filter(e).(type) {
			case lifecycle.Event:
				switch ev.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					// App visible.
					glctx, _ = ev.DrawContext.(gl.Context)
					onStart(glctx)
					a.Send(paint.Event{})
				case lifecycle.CrossOff:
					// App no longer visible.
				}
			case size.Event:
				sz = ev

			case paint.Event:
				if glctx == nil || ev.External {
					continue
				}
				onPaint(glctx, sz)
				a.Publish()
				a.Send(paint.Event{}) // keep animating
			}
		}
	})
}

var (
	startTime = time.Now()
	images    *glutil.Images
	eng       sprite.Engine
	scene     *sprite.Node
	game      *Game
)

func onStart(glctx gl.Context) {
	images = glutil.NewImages(glctx)
	eng = glsprite.Engine(images)
	game = NewGame()
	scene = game.Scene(eng)
}

func onStop() {
	eng.Release()
	images.Release()
}

func onPaint(glctx gl.Context, sz size.Event) {
	glctx.ClearColor(1, 1, 1, 1)
	glctx.Clear(gl.COLOR_BUFFER_BIT)
	now := clock.Time(time.Since(startTime) * 60 / time.Second)
	eng.Render(scene, now, sz)
}
