package gohotdraw

type Color uint32 

const (
    Black         Color = 0x000000
    White         Color = 0xFFFFFF
    Red           Color = 0xFF0000
    Green         Color = 0x00FF00
    Blue          Color = 0x0000FF
    Cyan          Color = 0x00FFFF
    Magenta       Color = 0xFF00FF
    Yellow        Color = 0xFFFF00
    PaleYellow    Color = 0xFFFFAA
    DarkYellow    Color = 0xEEEE9E
    DarkGreen     Color = 0x448844
    PaleGreen     Color = 0xAAFFAA
    MedGreen      Color = 0x88CC88
    DarkBlue      Color = 0x000055
    PaleBlueGreen Color = 0xAAFFFF
    PaleBlue      Color = 0x0000BB
    BlueGreen     Color = 0x008888
    GreyGreen     Color = 0x55AAAA
    PaleGreyGreen Color = 0x9EEEEE
    YellowGreen   Color = 0x99994C
    MedBlue       Color = 0x000099
    GreyBlue      Color = 0x005DBB
    PaleGreyBlue  Color = 0x4993DD
    PurpleBlue    Color = 0x8888CC
    Gray		  Color = (230 << 16 | 230 << 8 | 230 << 0)
    LightGray	  Color = (210 << 16 | 210 << 8 | 210 << 0)

)

func NewColor(red, green, blue uint32) Color {
	redComponent := uint32(red) << 16
	greenComponent := uint32(green) << 8
	blueComponent := uint32(blue) << 0
	rgb := (redComponent | greenComponent | blueComponent)
	return Color(rgb)
}

func (this Color) GetChannels() (r,g,b uint32) {
    color := uint32(this)
    r = color>>16 
    g = (color>>8)&0xFF
    b = (color>>0)&0xFF
    return
}

func (this Color) GetChannelArray() (array []uint32) {
	array = make([]uint32,3)
	array[0],array[1],array[2] = this.GetChannels()
	return
}

