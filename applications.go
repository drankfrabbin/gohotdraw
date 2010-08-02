package gohotdraw

import (
	"fmt"
)

const (
	F1_KEY = 67
	F2_KEY = 68
)

type DefaultApp struct {
	g      Graphics
	title  string
	editor DrawingEditor
}

func NewDefaultApp(graphics Graphics, title string) *DefaultApp {
	app := &DefaultApp{}
	app.title = title
	app.g = graphics
	app.g.AddInputListener(app)
	return app
}

func (this *DefaultApp) SetTool(tool Tool) {
	if this.editor.GetTool() != nil && this.editor.GetTool().IsActive() {
		this.editor.GetTool().Deactivate()
	}
	this.editor.SetTool(tool)
	if this.editor.GetTool() != nil {
		this.editor.GetTool().Activate()
	}
}

func (this *DefaultApp) Open() {
	this.setupGraphics()
	this.editor = this.createEditor(this.g)
	this.g.ShowWindow()
	this.g.StartListening()
}

func (this *DefaultApp) createEditor(graphics Graphics) DrawingEditor {
	editor := NewDefaultDrawingEditor()
	editor.SetTool(NewNullTool(editor))
	editor.SetView(this.createView(editor, graphics))
	return editor
}

func (this *DefaultApp) createView(editor DrawingEditor, graphics Graphics) DrawingView {
	view := NewStandardDrawingView()
	view.SetDrawing(NewStandardDrawing())
	view.SetEditor(editor)
	view.SetGraphics(graphics)
	return view
}

func (this *DefaultApp) setupGraphics() {
	this.g.SetWindowBackground(White)
	this.g.SetWindowTitle(this.title)
	this.g.SetFGColor(Black)
}

func (this *DefaultApp) MouseDown(e *MouseEvent) {
	this.editor.GetTool().MouseDown(e)
}

func (this *DefaultApp) MouseUp(e *MouseEvent) {
	this.editor.GetTool().MouseUp(e)
}
func (this *DefaultApp) MouseDrag(e *MouseEvent) {
	this.editor.GetTool().MouseDrag(e)
}


func (this *DefaultApp) KeyDown(e *KeyEvent) {
	//fmt.Println(e.KeyCode)
	if e.KeyCode == F1_KEY {
		fmt.Println("Creation Tool set")
		this.SetTool(NewCreationTool(this.editor, NewRectangleFigure()))
	} else if e.KeyCode == F2_KEY {
		fmt.Println("Selection Tool set")
		this.SetTool(NewSelectionTool(this.editor))
	}
}
func (this *DefaultApp) KeyUp(e *KeyEvent) {
	//this.editor.GetTool().KeyUp(e)
}

func (this *DefaultApp) ExposeHappened(e *ExposeEvent) {
	this.editor.GetView().Repaint()
}
