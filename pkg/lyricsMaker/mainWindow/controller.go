package mainWindow

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"lyrics/pkg/fyneGUI/toolBar"
	"lyrics/pkg/lyricsMaker/config"
)

type Controller struct {
	mainWindow *MainWindow
	object     fyne.CanvasObject
	volume     string
	legend     string
}

func GenerateController(parent *MainWindow) *Controller {
	return &Controller{
		mainWindow: nil,
		object:     nil,
		volume:     "",
		legend:     "",
	}
}

func (that *Controller) Combine() {
	that.object = widget.NewToolbar(
		// 音量控制
		toolBar.NewIntController(1,
			theme.VolumeDownIcon(),
			theme.VolumeUpIcon(),
			0.1,
			func(result float64) bool {
				return 0 <= result && result <= 1
			}, func() {
				// TODO 调整播放器音量
			}),
		// 播放控制
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {}),
		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {}),
		widget.NewToolbarAction(theme.MediaReplayIcon(), func() {}),

		// 字幕管理
		// 时间点标记
		widget.NewToolbarSpacer(),
		toolBar.NewButton(config.Begin+" "+config.Flag, theme.ContentAddIcon(), func() {}),
		toolBar.NewButton(config.End+" "+config.Flag, theme.ContentAddIcon(), func() {}),
		// 保存和删除
		toolBar.NewButton(config.Add, theme.ConfirmIcon(), func() {}),
		toolBar.NewButton(config.Delete, theme.DeleteIcon(), func() {}),
		widget.NewToolbarSpacer(),

		// 位置调整
		widget.NewToolbarAction(theme.NavigateBackIcon(), func() {}),
		widget.NewToolbarAction(theme.NavigateNextIcon(), func() {}),

		// 缩放调整
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ZoomFitIcon(), func() {}),
		toolBar.NewIntController(1,
			theme.ZoomInIcon(),
			theme.ZoomOutIcon(),
			0.1,
			func(result float64) bool {
				return 0 <= result && result <= 1
			}, func() {
				// TODO 调整播放器音量
			}),
	)
}
