package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"lyrics/pkg/player/extensions"
	"lyrics/pkg/player/fmod"
)

type mainWindow struct {
	window     fyne.Window
	minX       float32
	minY       float32
	controller *Controller
	time       string
	image      string
	player     *fmod.Player
}

func GenerateMainWindow(daemon fyne.App, title string) *mainWindow {
	mainWindow := &mainWindow{
		window: daemon.NewWindow(title),
		minX:   MainWindowX,
		minY:   MainWindowY,
		player: fmod.GeneratePlayerFMOD(),
	}
	mainWindow.controller = GenerateController(mainWindow)
	return mainWindow
}

func (that *mainWindow) Start() {
	// 调整窗口大小
	that.window.Resize(fyne.NewSize(that.minX, that.minY))
	// 设置为主窗口
	that.window.SetMaster()

	that.controller.Combine()

	// 设置主窗口布局
	that.window.SetContent(container.NewVBox(
		//that.menu.object,
		container.NewBorder(widget.NewButtonWithIcon("Open Media", theme.FolderOpenIcon(), func() {
			fileOpen := dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
				if nil != err {
					log.Println(err.Error())
					return
				}
				// 处理未选择文件情况
				if nil == closer {
					return
				}
				// 打开音频文件
				for extension := range extensions.MusicSupportList {
					if extensions.MusicSupportList[extension] == closer.URI().Extension() {
						errIn := that.player.LoadMedia(closer.URI().Path())
						if nil != err {
							log.Println(errIn.Error())
						}
						that.controller.UpdateCurrent("0")
						that.controller.UpdateLength(fmt.Sprintf("%d", that.player.GetLength()))
						return
					}
				}
			}, that.window)
			fileOpen.Show()
		}), nil, nil, nil, that.controller.object),
		//that.controller.object,
		// 控制栏和状态栏置于底部
		layout.NewSpacer(),
		//that.controller.object,
	))

	// 关闭主窗口时记录日志
	that.window.SetOnClosed(func() {
		log.Printf("Exit media player.")
	})

	// 运行
	that.window.ShowAndRun()
}

func (that *mainWindow) Close() {
	that.player.Close()
}

type Player struct {
	daemon     fyne.App
	mainWindow *mainWindow
}

func GeneratePlayerClient() *Player {
	daemon := app.New()
	return &Player{
		daemon:     daemon,
		mainWindow: GenerateMainWindow(daemon, MainWindowTitle),
	}
}

func (that *Player) Start() {
	that.mainWindow.Start()
}

func (that *Player) Stop() {
	that.mainWindow.Close()
}
