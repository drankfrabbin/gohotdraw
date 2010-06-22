package gohotdraw

type Painter interface {
	Draw(g Graphics, view DrawingView)
}

type SimpleUpdateStrategy struct{}

func (this *SimpleUpdateStrategy) Draw(g Graphics, view DrawingView) {
	view.Draw(g)
	size := g.GetWindowSize()
	g.Repaint(0,0,0,0,size.Width, size.Height)
}
