package gohotdraw

import (
	_ "fmt"
)

type DrawingEditor interface {
	FigureSelectionListener
	GetView() DrawingView
	SetView(view DrawingView)
	SetTool(tool Tool)
	GetTool() Tool
	ToolDone()
}

type DefaultDrawingEditor struct {
	view DrawingView
	tool Tool
}

func NewDefaultDrawingEditor() *DefaultDrawingEditor {
	editor := &DefaultDrawingEditor{}
	return editor
}

func (this *DefaultDrawingEditor) GetView() DrawingView {
	return this.view
}

func (this *DefaultDrawingEditor) SetView(view DrawingView) {
	this.view = view
	this.view.AddFigureSelectionListener(this)
}

func (this *DefaultDrawingEditor) GetTool() Tool {
	return this.tool
}

func (this *DefaultDrawingEditor) SetTool(tool Tool) {
	this.tool = tool
}

func (this *DefaultDrawingEditor) ToolDone() {
	//do nothing, could be changed to set default tool
}

func (this *DefaultDrawingEditor) FigureSelectionChanged(view DrawingView) {
	view.Repaint()
}
