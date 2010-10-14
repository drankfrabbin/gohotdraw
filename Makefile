include $(GOROOT)/src/Make.inc

TARG=github.com/drankfrabbin/gohotdraw

GOFILES=\
	applications.go\
	colors.go\
	drawings.go\
	editors.go\
	events.go\
	figures.go\
	figuredecoration.go\
	graphics.go\
	graphicsevents.go\
	handles.go\
	listeners.go\
	locators.go\
	painters.go\
	resizehandles.go\
	set.go\
	tools.go\
	util.go\
	views.go\

include $(GOROOT)/src/Make.pkg
