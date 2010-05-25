include $(GOROOT)/src/Make.$(GOARCH)

TARG=github.com/drankfrabbin/gohotdraw

GOFILES=\
	applications.go\
	constants.go\
	drawings.go\
	editors.go\
	events.go\
	figures.go\
	graphics.go\
	graphicsevents.go\
	handles.go\
	listeners.go\
	locators.go\
	painters.go\
	resizehandles.go\
	tools.go\
	util.go\
	views.go\

include $(GOROOT)/src/Make.pkg
