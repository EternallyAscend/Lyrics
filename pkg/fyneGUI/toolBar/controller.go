package toolBar

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// 以进度条为显示的可增减控制器

type Float64ProgressController struct {
	object      fyne.CanvasObject
	value       float64
	progressBar *widget.ProgressBar
	down        fyne.Resource
	up          fyne.Resource
	step        float64
	limitCheck  func(result float64) bool
	callback    func(value float64)
}

func GenerateFloat64ProgressControllerPointer(value float64, down fyne.Resource, up fyne.Resource, step float64, limitChecker func(result float64) bool, callback func(value float64)) *Float64ProgressController {
	return &Float64ProgressController{
		object:      nil,
		value:       value,
		progressBar: nil,
		down:        down,
		up:          up,
		step:        step,
		limitCheck:  limitChecker,
		callback:    callback,
	}
}

func (that *Float64ProgressController) SetValue(value float64) {
	if that.limitCheck(value) {
		that.value = value
	}
	that.setValue()
}

func (that *Float64ProgressController) setValue() {
	that.progressBar.SetValue(that.value)
	that.progressBar.Refresh()
	that.callback(that.value)
}

func (that *Float64ProgressController) Up() {
	if that.limitCheck(that.value + that.step) {
		that.value += that.step
	}
	that.setValue()
}

func (that *Float64ProgressController) Down() {
	if that.limitCheck(that.value - that.step) {
		that.value -= that.step
	}
	that.setValue()
}

func (that *Float64ProgressController) ToolbarObject() fyne.CanvasObject {
	that.progressBar = widget.NewProgressBar()
	that.progressBar.SetValue(that.value)

	that.object = container.NewBorder(nil, nil,
		widget.NewButtonWithIcon("", that.down, func() {
			that.Down()
		}),
		widget.NewButtonWithIcon("", that.up, func() {
			that.Up()
		}), that.progressBar)
	return that.object
}

func NewFloat64ProgressController(value float64, down fyne.Resource, up fyne.Resource, step float64, limitChecker func(result float64) bool, callback func(value float64)) widget.ToolbarItem {
	return &Float64ProgressController{
		object:      nil,
		value:       value,
		progressBar: nil,
		down:        down,
		up:          up,
		step:        step,
		limitCheck:  limitChecker,
		callback:    callback,
	}
}

func NewFloat64ProgressControllerByPointer(pointer *Float64ProgressController) widget.ToolbarItem {
	return pointer
}

// 以自定义内容为显示的可增减控制器

type Float64TextController struct {
	object        fyne.CanvasObject
	value         float64
	text          *widget.Label
	down          fyne.Resource
	up            fyne.Resource
	step          float64
	limitCheck    func(result float64) bool
	valueToString func(value float64) string
	callback      func(value float64)
}

func GenerateFloat64TextControllerPointer(value float64, down fyne.Resource, up fyne.Resource, step float64, limitChecker func(result float64) bool, valueToString func(value float64) string, callback func(value float64)) *Float64TextController {
	return &Float64TextController{
		object:        nil,
		value:         value,
		text:          widget.NewLabel(valueToString(value)),
		down:          down,
		up:            up,
		step:          step,
		limitCheck:    limitChecker,
		valueToString: valueToString,
		callback:      callback,
	}
}

func (that *Float64TextController) SetValue(value float64) {
	if that.limitCheck(value) {
		that.value = value
	}
	that.setValue()
}

func (that *Float64TextController) setValue() {
	that.text.SetText(that.valueToString(that.value))
	that.text.Refresh()
	that.callback(that.value)
}

func (that *Float64TextController) Up() {
	if that.limitCheck(that.value + that.step) {
		that.value += that.step
	}
	that.setValue()
}

func (that *Float64TextController) Down() {
	if that.limitCheck(that.value - that.step) {
		that.value -= that.step
	}
	that.setValue()
}

func (that *Float64TextController) ToolbarObject() fyne.CanvasObject {
	that.object = container.NewBorder(nil, nil,
		widget.NewButtonWithIcon("", that.down, func() {
			that.Down()
		}),
		widget.NewButtonWithIcon("", that.up, func() {
			that.Up()
		}), that.text)
	return that.object
}

func NewFloat64TextController(value float64, down fyne.Resource, up fyne.Resource, step float64, limitChecker func(result float64) bool, valueToString func(value float64) string, callback func(value float64)) widget.ToolbarItem {
	return &Float64TextController{
		object:        nil,
		value:         value,
		text:          widget.NewLabel(valueToString(value)),
		down:          down,
		up:            up,
		step:          step,
		limitCheck:    limitChecker,
		valueToString: valueToString,
		callback:      callback,
	}
}

func NewFloat64TextControllerByPointer(pointer *Float64TextController) widget.ToolbarItem {
	return pointer
}
