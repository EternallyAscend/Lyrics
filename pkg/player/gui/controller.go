package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"lyrics/pkg/fyneGUI/toolBar"
	"time"
)

type Controller struct {
	mainWindow *mainWindow
	object     fyne.CanvasObject
	current    *toolBar.TextDescription
	length     *toolBar.TextDescription
}

func GenerateController(parent *mainWindow) *Controller {
	return &Controller{
		mainWindow: parent,
		object:     nil,
	}
}

func (that *Controller) UpdateLength(length string) {
	that.length.Update(length)
}

func (that *Controller) UpdateCurrent(current string) {
	that.current.Update(current)
}

func (that *Controller) Combine() {
	that.current = toolBar.GenerateTextDescriptionPointer("0:00:00")
	that.length = toolBar.GenerateTextDescriptionPointer("0:00:00")
	that.object = widget.NewToolbar(

		// 当前位置
		toolBar.NewTextDescriptionPointer(that.current),

		// 播放控制
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
			go that.mainWindow.player.Play()
			go func() {
				time.Sleep(time.Millisecond * UpdatePositionTimeGapMS)
				for that.mainWindow.player.GetPlaying() {
					time.Sleep(time.Millisecond * UpdatePositionTimeGapMS)
					that.UpdateCurrent(fmt.Sprintf("%d", that.mainWindow.player.GetPosition()))
				}
			}()
		}),
		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
			go that.mainWindow.player.Pause()
		}),
		widget.NewToolbarAction(theme.MediaReplayIcon(), func() {}),
		toolBar.NewTextDescriptionPointer(that.length),

		// 速度控制
		widget.NewToolbarSpacer(),
		toolBar.NewFloat64TextController(1,
			theme.MoveDownIcon(),
			theme.MoveUpIcon(),
			0.2,
			func(result float64) bool {
				return 0.2 <= result && result <= 16
			}, func(value float64) string {
				return fmt.Sprintf("%.1fx", value)
			}, func() {
				// TODO 调整播放速度
			}),

		// 音量控制
		widget.NewToolbarSpacer(),
		toolBar.NewFloat64ProgressController(1,
			theme.VolumeDownIcon(),
			theme.VolumeUpIcon(),
			0.1,
			func(result float64) bool {
				return 0 <= result && result <= 1
			}, func() {
				// TODO 调整播放器音量
			}),
		//toolBar.NewFigure(FmodLogoPath),
	)
}
