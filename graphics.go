package gohotdraw

import (
	"fmt"
	"os"
	"x-go-binding.googlecode.com/hg/xgb"
)

const (
	DEFAULT_X      int16  = 50
	DEFAULT_Y      int16  = 50
	DEFAULT_WIDTH  uint16 = 400
	DEFAULT_HEIGHT uint16 = 350
	DEFAULT_EVENTS uint32 = 
		xgb.EventMaskKeyRelease | 
		xgb.EventMaskKeyPress | 
		xgb.EventMaskButtonRelease | 
		xgb.EventMaskButtonPress | 
		xgb.EventMaskButtonMotion | 
		xgb.EventMaskExposure
)

type Graphics interface {
	AddInputListener(l InputListener)
	RemoveInputListener(l InputListener)

	SetFGColor(color Color)
	//SetBGColor(color Color)
	ShowWindow()
	SetWindowBackground(color Color)
	SetWindowTitle(title string)
	StartListening()
	GetWindowSize() *Dimension

	DrawRect(x, y, width, height int)
	DrawRectFromRect(rect *Rectangle)
	DrawBorder(x, y, width, height int)
	DrawBorderFromRect(rect *Rectangle)
	DrawBorderedRect(x, y, width, height int)
	DrawBorderedRectFromRect(rect *Rectangle)
	DrawBorderFromRectDirectly(rect *Rectangle)
	
	Repaint(srcX, srcY, destX, destY , width , height int)
}

type XGBGraphics struct {
	listeners  *Set
	connection *xgb.Conn
	winId      xgb.Id
	pixmapId   xgb.Id
	contextId  xgb.Id
}

func NewDefaultXGBGraphics() *XGBGraphics {
	return NewXGBGraphics(DEFAULT_X, DEFAULT_Y, DEFAULT_WIDTH, DEFAULT_HEIGHT)
}

func NewXGBGraphics(x, y int16, width, height uint16) *XGBGraphics {
	g := &XGBGraphics{}
	g.listeners = NewSet()
	g.initialize(x, y, width, height)
	g.SetEventListenTypes([]uint32{DEFAULT_EVENTS})
	return g
}

func (this *XGBGraphics) initialize(x, y int16, width, height uint16) {
	this.connection = this.createConnection()
	this.winId = this.createWindow(x, y, width, height)
	this.pixmapId = this.createPixmap(width, height)
	this.contextId = this.createContext()
}

func (this *XGBGraphics) createConnection() *xgb.Conn {
	connection, err := xgb.Dial(os.Getenv("DISPLAY"))
	if err != nil {
		fmt.Printf("cannot connect: %v\n", err)
		os.Exit(1)
	}
	return connection
}

func (this *XGBGraphics) createWindow(x, y int16, width, height uint16) xgb.Id {
	if this.connection != nil {
		winId := this.connection.NewId()
		this.connection.CreateWindow(
			this.connection.DefaultScreen().RootDepth,
			winId, 
			this.connection.DefaultScreen().Root, 
			x, 
			y, 
			width, 
			height, 
			0, 
			xgb.WindowClassCopyFromParent, 
			this.connection.DefaultScreen().RootVisual, 
			0, 
			nil)
		return winId
	}
	return 0
}

func (this *XGBGraphics) createPixmap(width, height uint16) xgb.Id {
	if this.connection != nil {
		pixmapId := this.connection.NewId()
		this.connection.CreatePixmap(
			this.connection.DefaultScreen().RootDepth, 
			pixmapId, 
			this.winId, 
			width, 
			height)
		return pixmapId
	}
	return 0
}

func (this *XGBGraphics) createContext() xgb.Id {
	if this.connection != nil && this.winId != 0 {
		contextId := this.connection.NewId()
		this.connection.CreateGC(contextId, this.winId, 0, nil)
		return contextId
	}
	return 0
}

func (this *XGBGraphics) ShowWindow() {
	this.connection.MapWindow(this.winId)
}

func (this *XGBGraphics) HideWindow() {
	this.connection.UnmapWindow(this.winId)
}

func (this *XGBGraphics) CloseAll() {
	this.connection.FreePixmap(this.pixmapId)
	this.connection.Close()
}

func (this *XGBGraphics) SetWindowTitle(title string) {
	this.setProperty(xgb.AtomWmName, title)
}

func (this *XGBGraphics) setProperty(propertyName xgb.Id, value string) {
	this.connection.ChangeProperty(
		xgb.PropModeReplace, 
		this.winId, 
		propertyName, 
		xgb.AtomString, 
		8, 
		[]byte(value))
}

func (this *XGBGraphics) SetEventListenTypes(eventTypes []uint32) {
	this.setWindowAttributes(xgb.CWEventMask, eventTypes)
}

func (this *XGBGraphics) setWindowAttributes(attribute uint32, values []uint32) {
	this.connection.ChangeWindowAttributes(this.winId, attribute, values)
}

func (this *XGBGraphics) setContextAttributes(attribute uint32, values []uint32) {
	this.connection.ChangeGC(this.contextId, attribute, values)
}

func (this *XGBGraphics) SetWindowBackground(color Color) {
	this.setWindowAttributes(xgb.CWBackPixel, []uint32{uint32(color)})
}

func (this *XGBGraphics) SetFGColor(color Color) {
	this.setContextAttributes(xgb.GCForeground, []uint32{uint32(color)})
}

func (this *XGBGraphics) SetBGColor(color Color) {
	this.setContextAttributes(xgb.GCBackground, []uint32{uint32(color)})
}

func (this *XGBGraphics) GetWindowSize() *Dimension {
	reply, err := this.connection.GetGeometry(this.winId)
	if err != nil {
		fmt.Println("ERROR getting window size")
	}
	return &Dimension{int(reply.Width), int(reply.Height)}
}

func (this *XGBGraphics) DrawRect(x, y, width, height int) {
	rect := this.createRectangle(x, y, width, height)
	this.connection.PolyFillRectangle(this.pixmapId, this.contextId, rect)
}

func (this *XGBGraphics) DrawRectFromRect(rect *Rectangle) {
	this.DrawRect(rect.X, rect.Y, rect.Width, rect.Height)
}

func (this *XGBGraphics) DrawBorder(x, y, width, height int) {
	rect := this.createRectangle(x, y, width, height)
	this.connection.PolyRectangle(this.pixmapId, this.contextId, rect)
}

func (this *XGBGraphics) DrawBorderFromRect(rect *Rectangle) {
	this.DrawBorder(rect.X, rect.Y, rect.Width, rect.Height)
}

func (this *XGBGraphics) DrawBorderFromRectDirectly(rect *Rectangle) {
	xgbRect := this.createRectangle(rect.X, rect.Y, rect.Width, rect.Height)
	this.connection.PolyRectangle(this.winId, this.contextId, xgbRect)
}

func (this *XGBGraphics) DrawBorderedRect(x, y, width, height int) {
	this.DrawRect(x, y, width, height)
	this.SetFGColor(Black)
	this.DrawBorder(x, y, width, height)
}

func (this *XGBGraphics) DrawBorderedRectFromRect(rect *Rectangle) {
	this.DrawBorderedRect(rect.X, rect.Y, rect.Width, rect.Height)
}

func (this *XGBGraphics) createRectangle(x, y, width, height int) []xgb.Rectangle {
	xgbRect := make([]xgb.Rectangle, 1)
	xgbRect[0] = xgb.Rectangle{int16(x), int16(y), uint16(width), uint16(height)}
	return xgbRect
}

func (this *XGBGraphics) repaint(srcX, srcY, destX, destY  int16 , width , height uint16) {
	this.connection.CopyArea(this.pixmapId, this.winId, this.contextId, srcX , srcY , destX , destY , width, height)
}

func (this *XGBGraphics) Repaint(srcX, srcY, destX, destY , width , height int) {
	this.repaint(int16(srcX), int16(srcY), int16(destX), int16(destY) , uint16(width) , uint16(height))
}
