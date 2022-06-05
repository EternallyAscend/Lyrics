package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"lyrics/pkg/fyneGUI/toolBar"
	"lyrics/pkg/lyrics/project"
	"lyrics/pkg/player/fmod"
	"strconv"
	"time"
)

type Controller struct {
	window  fyne.Window
	player  *fmod.Player
	object  fyne.CanvasObject
	current *toolBar.TextDescription
	length  *toolBar.TextDescription
	speed   *toolBar.Float64TextController
	volume  *toolBar.Float64ProgressController
}

func GenerateController(window fyne.Window, player *fmod.Player) *Controller {
	return &Controller{
		window: window,
		player: player,
		object: nil,
	}
}

func (that *Controller) UpdateLength(ms uint32) {
	that.length.Update(project.TransferMillSecondToString(int64(ms)))
}

func (that *Controller) UpdateCurrent(ms uint32) {
	that.current.Update(project.TransferMillSecondToString(int64(ms)))
}

func (that *Controller) Combine() {
	that.current = toolBar.GenerateTextDescriptionPointer("00:00:00.000")
	that.length = toolBar.GenerateTextDescriptionPointer("00:00:00.000")
	that.speed = toolBar.GenerateFloat64TextControllerPointer(1,
		theme.MoveDownIcon(),
		theme.MoveUpIcon(),
		0.2,
		func(result float64) bool {
			result, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", result), 64)
			return 0.4 <= result && result <= 4.1
		}, func(value float64) string {
			return fmt.Sprintf("%.1fx", value)
		}, func(value float64) {
			// 调整播放速度
			that.player.SetFrequency(value)
		})

	that.volume = toolBar.GenerateFloat64ProgressControllerPointer(1,
		theme.VolumeDownIcon(),
		theme.VolumeUpIcon(),
		0.1,
		func(result float64) bool {
			return 0 <= result && result <= 1.0
		}, func(value float64) {
			// 调整播放器音量
			that.player.SetVolume(float32(value))
		})

	that.object = widget.NewToolbar(

		// 播放控制
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
			go that.player.Play()
			// 更新GUI
			go func() {
				// 等待播放
				time.Sleep(time.Millisecond * UpdatePositionTimeWaitMS)
				// 监视播放状态，更新当前播放时间
				for that.player.GetPlaying() {
					time.Sleep(time.Millisecond * UpdatePositionTimeGapMS)
					that.UpdateCurrent(that.player.GetPosition())
				}
			}()
		}),
		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
			that.player.Pause()
		}),
		widget.NewToolbarAction(theme.MediaReplayIcon(), func() {
			that.player.SetPosition(0)
			that.UpdateCurrent(0)
		}),

		// 当前位置 & 总长度
		widget.NewToolbarSpacer(),
		toolBar.NewTextDescriptionByPointer(that.current),
		toolBar.NewTextDescription("-"),
		toolBar.NewTextDescriptionByPointer(that.length),

		// 跳转
		widget.NewToolbarSpacer(),
		toolBar.NewButton("Goto", theme.SettingsIcon(), func() {
			log.Println("Goto position.")
		}),
		// 速度控制
		toolBar.NewButton("", theme.MediaReplayIcon(), func() {
			// that.player.SetFrequency(1)
			that.speed.SetValue(1)
		}),
		toolBar.NewFloat64TextControllerByPointer(that.speed),
		//toolBar.NewFloat64TextController(1,
		//	theme.MoveDownIcon(),
		//	theme.MoveUpIcon(),
		//	0.2,
		//	func(result float64) bool {
		//		return 0.2 <= result && result <= 16
		//	}, func(value float64) string {
		//		return fmt.Sprintf("%.1fx", value)
		//	}, func(value float64) {
		//		// 调整播放速度
		//		that.player.SetFrequency(float32(value))
		//	}),
		// 音量控制
		toolBar.NewButton("", theme.VolumeMuteIcon(), func() {
			if that.player.GetVolume() > 0 {
				// that.player.SetVolume(0)
				that.volume.SetValue(0)
			} else {
				// that.player.SetVolume(1)
				that.volume.SetValue(1)
			}
		}),
		toolBar.NewFloat64ProgressControllerByPointer(that.volume),
		//toolBar.NewFloat64ProgressController(1,
		//	theme.VolumeDownIcon(),
		//	theme.VolumeUpIcon(),
		//	0.1,
		//	func(result float64) bool {
		//		return 0 <= result && result <= 1
		//	}, func(value float64) {
		//		// 调整播放器音量
		//		that.player.SetVolume(float32(value))
		//	}),
		//toolBar.NewFigure(FmodLogoPath),
	)
}
