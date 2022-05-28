package mainWindow

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"lyrics/pkg/fyneGUI/toolBar"
	"lyrics/pkg/lyricsMaker/config"
)

type Menu struct {
	mainWindow  *MainWindow
	object      fyne.CanvasObject
	description *MenuDescription
}

type MenuDescription struct {
	menu        *Menu
	description *widget.Label
}

func (that *MenuDescription) UpdateDescription(description string) {
	that.description.SetText(description)
}

func (that *MenuDescription) ToolbarObject() fyne.CanvasObject {
	return container.NewBorder(nil, nil, nil, nil, that.description)
}

func NewToolbarDescription(menu *Menu, description string) widget.ToolbarItem {
	pointer := &MenuDescription{
		menu:        menu,
		description: widget.NewLabel(description),
	}
	menu.description = pointer
	return pointer
}

func GenerateMenu(parent *MainWindow) *Menu {
	return &Menu{
		mainWindow: parent,
	}
}

func (that *Menu) Combine() {
	that.object = widget.NewToolbar(
		//widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
		//	log.Println("New document")
		//}),
		// 打开文件
		toolBar.NewButton("Open Project", theme.FolderOpenIcon(), func() {}),
		toolBar.NewButton("Open Media", theme.FolderOpenIcon(), func() {
			fileOpen := dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
				if nil != err {
					that.mainWindow.UpdateStatus(err.Error())
					return
				}
				// 处理未选择文件情况
				if nil == closer {
					return
				}
				// 打开音频文件
				for extension := range config.MusicExtension {
					if config.MusicExtension[extension] == closer.URI().Extension() {
						that.mainWindow.OpenFile(closer.URI())
						that.description.UpdateDescription(closer.URI().Name())
						return
					}
				}
				// 打开视频文件
				for extension := range config.VideoExtension {
					if config.VideoExtension[extension] == closer.URI().Extension() {
						that.mainWindow.OpenFile(closer.URI())
						that.description.UpdateDescription(closer.URI().Name())
						return
					}
				}
				// 处理不能打开的文件类型
				that.mainWindow.UpdateStatus("Do not support \"" + closer.URI().Extension() + "\" file.")
			}, that.mainWindow.window)
			fileOpen.Show()
		}),
		toolBar.NewButton("Open Lyrics", theme.FolderOpenIcon(), func() {}),
		toolBar.NewButton("Open Timeline", theme.FolderOpenIcon(), func() {}),

		// 保存文件
		toolBar.NewButton("Save", theme.DocumentSaveIcon(), func() {}),

		// 撤销重做
		//widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentUndoIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentRedoIcon(), func() {}),

		//widget.NewToolbarSeparator(),
		//widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		//widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		//widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),

		//widget.NewToolbarSeparator(),
		widget.NewToolbarSpacer(),
		NewToolbarDescription(that, that.mainWindow.file),
		widget.NewToolbarSpacer(),
		//widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			helpDialog := dialog.NewInformation(config.ApplicationTitle, "Author: EternallyAscend\n"+
				"\nGitHub: https://github.com/EternallyAscend/lyrics\nPage: https://EternallyAscend.github.io/lyrics\n"+
				"\n本软件提供音频播放制作CUE文件、时间轴、SRT字幕和LRC歌词文件的功能。\n"+
				"\nThis software provide functions for making cue, timeline srt and lrc.\n\n"+
				"\nDevelop for Tom Chang.", that.mainWindow.window)
			helpDialog.Show()
		}),
	)
}
