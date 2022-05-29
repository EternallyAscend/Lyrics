package mainWindow

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
	"lyrics/pkg/lyricsMaker/config"
	"lyrics/pkg/lyricsMaker/mainWindow/body"
	"time"
)

type MainWindow struct {
	window     fyne.Window
	minX       float32
	minY       float32
	menu       *Menu
	controller *Controller
	body       *body.Body
	statusBar  *StatusBar
	player     string
	media      string
	mediaPath  string
	lyrics     []string
	timeline   string
	cue        string
}

func GenerateMainWindow(daemon fyne.App, title string) *MainWindow {
	mainWindow := &MainWindow{
		window:    daemon.NewWindow(title),
		minX:      config.MainWindowX,
		minY:      config.MainWindowY,
		body:      body.GenerateBody(),
		statusBar: GenerateStatusBar(),
		media:     config.DefaultMediaName,
		mediaPath: config.DefaultMediaPath,
	}
	mainWindow.menu = GenerateMenu(mainWindow)
	mainWindow.controller = GenerateController(mainWindow)
	return mainWindow
}

func (that *MainWindow) Load() {
	// 调整窗口大小
	that.window.Resize(fyne.NewSize(that.minX, that.minY))
	// 设置为主窗口
	that.window.SetMaster()

	that.menu.Combine()
	that.controller.Combine()
	that.body.Combine()
	that.statusBar.Combine()

	//widget.NewButton("Export", func() {
	//	that.UpdateStatus("Export file \"" + that.file + "\".")
	//}

	// 设置主窗口布局
	that.window.SetContent(container.NewVBox(
		//that.menu.object,
		// 顶部菜单栏，左侧歌词列表，右侧功能按钮列表
		container.NewBorder(that.menu.object, nil, widget.NewLabel(config.Lyrics), widget.NewButton(config.Export, func() {
			that.UpdateStatus(config.ExportFile + " \"" + that.media + "\".")
		}), that.body.Object),
		// 控制栏和状态栏置于底部
		layout.NewSpacer(),
		that.controller.object,
		that.statusBar.object,
	))

	that.window.SetCloseIntercept(func() {
		closeDialog := dialog.NewConfirm(config.CloseDialogTitle, config.CloseDialogMessage, func(confirmed bool) {
			if confirmed {
				that.window.Close()
			}
		}, that.window)
		closeDialog.Show()
	})
	that.window.SetOnClosed(func() {
		log.Printf("Lyrics Maker closed at %s.", time.Now().String()[0:19])
	})

	// 运行
	that.window.ShowAndRun()
}

func (that *MainWindow) UpdateStatus(status string) {
	that.statusBar.UpdateStatus(status)
}

func (that *MainWindow) OpenMedia(file fyne.URI) {
	// 更新当前打开文件
	that.media = file.Name()
	that.mediaPath = file.Path()
	// TODO 更新Body
	// 更新状态栏
	that.UpdateStatus(config.OpenFile + file.Path())
}
