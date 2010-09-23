package gohotdraw

import (
	"fmt"
	"os"
	"x-go-binding.googlecode.com/hg/xgb"
)

func (this *XGBGraphics) AddInputListener(l InputListener) {
	this.listeners.Push(l)
}

func (this *XGBGraphics) RemoveInputListener(l InputListener) {
	for i := 0; i < this.listeners.Len(); i++ {
		currentListener := this.listeners.At(i).(InputListener)
		if currentListener == l {
			this.listeners.Delete(i)
			return
		}
	}
}

func (this *XGBGraphics) StartListening() {
	for {
		reply := this.GetEventReply()
		switch xgbEvent := reply.(type) {
		case xgb.ExposeEvent:
			event := &ExposeEvent{}
			event.X = int(xgbEvent.X)
			event.Y = int(xgbEvent.Y)
			event.Width = int(xgbEvent.Width)
			event.Height = int(xgbEvent.Height)
			this.fireExposeHappened(event)
		case xgb.ButtonPressEvent:
			event := &MouseEvent{}
			event.X = int(xgbEvent.EventX)
			event.Y = int(xgbEvent.EventY)
			event.Button = int(xgbEvent.Detail)
			event.KeyModifier = int(xgbEvent.State)
			this.fireMouseDown(event)
		case xgb.ButtonReleaseEvent:
			event := &MouseEvent{}
			event.X = int(xgbEvent.EventX)
			event.Y = int(xgbEvent.EventY)
			event.Button = int(xgbEvent.Detail)
			event.KeyModifier = int(xgbEvent.State)
			this.fireMouseUp(event)
		case xgb.MotionNotifyEvent:
			event := &MouseEvent{}
			event.X = int(xgbEvent.EventX)
			event.Y = int(xgbEvent.EventY)
			event.Button = int(xgbEvent.Detail)
			event.KeyModifier = int(xgbEvent.State)
			this.fireMouseDrag(event)
		case xgb.KeyPressEvent:
			event := &KeyEvent{}
			event.KeyCode = int(xgbEvent.Detail)
			event.KeyModifier = int(xgbEvent.State)
			this.fireKeyDown(event)
		case xgb.KeyReleaseEvent:
			event := &KeyEvent{}
			event.KeyCode = int(xgbEvent.Detail)
			event.KeyModifier = int(xgbEvent.State)
			this.fireKeyUp(event)
		}
	}
}

func (this *XGBGraphics) fireMouseUp(event *MouseEvent) {
	for i := 0; i < this.listeners.Len(); i++ {
		currentListener := this.listeners.At(i).(InputListener)
		currentListener.MouseUp(event)
	}
}

func (this *XGBGraphics) fireMouseDown(event *MouseEvent) {
	for i := 0; i < this.listeners.Len(); i++ {
		currentListener := this.listeners.At(i).(InputListener)
		currentListener.MouseDown(event)
	}
}

func (this *XGBGraphics) fireMouseDrag(event *MouseEvent) {
	for i := 0; i < this.listeners.Len(); i++ {
		currentListener := this.listeners.At(i).(InputListener)
		currentListener.MouseDrag(event)
	}
}

func (this *XGBGraphics) fireKeyUp(event *KeyEvent) {
	for i := 0; i < this.listeners.Len(); i++ {
		currentListener := this.listeners.At(i).(InputListener)
		currentListener.KeyUp(event)
	}
}

func (this *XGBGraphics) fireKeyDown(event *KeyEvent) {
	for i := 0; i < this.listeners.Len(); i++ {
		currentListener := this.listeners.At(i).(InputListener)
		currentListener.KeyDown(event)
	}
}

func (this *XGBGraphics) fireExposeHappened(event *ExposeEvent) {
	for i := 0; i < this.listeners.Len(); i++ {
		currentListener := this.listeners.At(i).(InputListener)
		currentListener.ExposeHappened(event)
	}
}

func (this *XGBGraphics) GetEventReply() xgb.Event {
	reply, err := this.connection.WaitForEvent()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	return reply
}
