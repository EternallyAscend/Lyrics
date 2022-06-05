package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"log"
	"lyrics/pkg/player/fmod"
)

type TimeView struct {
	window    fyne.Window
	object    *fyne.Container
	player    *fmod.Player
	time      *canvas.Raster
	timeBasic [][]uint8
	timeData  [][]uint8
}

func GenerateTimeView(window fyne.Window, player *fmod.Player) *TimeView {
	tv := &TimeView{
		window:    window,
		object:    container.NewWithoutLayout(),
		player:    player,
		time:      nil,
		timeBasic: nil,
		timeData:  nil,
	}
	tv.object.Resize(fyne.NewSize(TimeCanvasX, TimeCanvasY))
	tv.object.Add(tv.time)
	tv.DrawTime()
	return tv
}

func (that *TimeView) GenerateWaveData() {
	cursorData := make([]uint8, TimeCanvasX)
	for i := TimeLineStart; i <= TimeLineEnd; i++ {
		log.Println(i)
		cursorData[i] = 128
	}
	that.timeBasic = [][]uint8{
		cursorData,
	}
}

func (that *TimeView) DrawTime() {
	that.object.Remove(that.time)
	that.time = canvas.NewRasterWithPixels(that.CalculateTimeData)
	that.time.Resize(fyne.NewSize(TimeCanvasX, TimeCanvasY))
	that.time.SetMinSize(fyne.NewSize(TimeCanvasX, TimeCanvasY))
	that.time.Move(fyne.NewPos(0, 0))
	that.object.Add(that.time)
}

func (that *TimeView) CalculateTimeData(x, y, w, h int) color.Color {
	if TimeLineStart <= x && x <= TimeLineEnd {
		return color.RGBA{
			R: 128,
			G: 128,
			B: 128,
			A: 0xff,
		}
	}
	if y < h/16 {
		return color.RGBA{
			R: 64,
			G: 64,
			B: 64,
			A: 0xff,
		}
	} else if y < h/8 {
		return color.RGBA{
			R: 0,
			G: 0,
			B: 128,
			A: 0xff,
		}
	} else if y < 3*h/8 {
		if (TimeLineStart-x)%(TimeLineStart/TimeTimeNumberL) == 0 ||
			(x-TimeLineEnd)%(TimeLineStart/TimeTimeNumberR) == 0 {
			return color.RGBA{
				R: 0,
				G: 128,
				B: 0,
				A: 0xff,
			}
		}
		return color.RGBA{
			R: 64,
			G: 64,
			B: 64,
			A: 0xff,
		}
	} else if (h - y) < h/16 {
		return color.RGBA{
			R: 64,
			G: 64,
			B: 64,
			A: 0xff,
		}
	} else if (h - y) < h/8 {
		return color.RGBA{
			R: 128,
			G: 0,
			B: 0,
			A: 0xff,
		}
	} else {
		return color.RGBA{
			R: 0,    // that.timeBasic[0][x],
			G: 0,    // that.timeBasic[0][x],
			B: 0,    // that.timeBasic[0][x],
			A: 0xff, // that.timeBasic[0][x],
		}
	}
}

func testCanvas() *fyne.Container {
	var data [][]uint8
	log.Println(data)

	r := canvas.NewRasterWithPixels(func(x, y, w, h int) color.Color {
		//log.Println(x, y, w, h)
		return color.RGBA{
			R: 64,
			G: 64,
			B: 64,
			A: 0xff,
		}
	})

	c := container.NewWithoutLayout()
	r.Move(fyne.Position{
		X: 0,
		Y: 0,
	})
	c.Resize(fyne.NewSize(TimeCanvasX, TimeCanvasY))
	r.Resize(fyne.NewSize(TimeCanvasX, TimeCanvasY))
	r.SetMinSize(fyne.NewSize(TimeCanvasX, TimeCanvasY))
	c.Add(r)
	return c
}
