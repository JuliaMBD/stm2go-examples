package main

import (
	"fmt"
	"machine"

	testpackage "github.com/example/testpackage/src"
	"tinygo.org/x/drivers/examples/ili9341/initdisplay"
	"tinygo.org/x/drivers/ili9341"
	"tinygo.org/x/drivers/tone"

	"tinygo.org/x/tinyfont"

	"image/color"

	"tinygo.org/x/tinyfont/freesans"
)

var (
	black = color.RGBA{0x00, 0x00, 0x00, 0xFF}
	white = color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	// red   = color.RGBA{255, 0, 0, 255}
	// blue  = color.RGBA{0, 0, 255, 255}
	// green = color.RGBA{0, 255, 0, 255}
)

func RGBATo565(c color.RGBA) uint16 {
	r, g, b, _ := c.RGBA()
	return uint16((r & 0xF800) + ((g & 0xFC00) >> 5) + ((b & 0xF800) >> 11))
}

type LCDDisplay struct {
	display *ili9341.Device
	buf     []uint16
	width   int16
	height  int16
}

func NewLCDDisplay(w, h int16) *LCDDisplay {
	display := initdisplay.InitDisplay()
	display.FillScreen(white)
	return &LCDDisplay{
		display: display,
		buf:     make([]uint16, w*h),
		width:   w,
		height:  h,
	}
}

func (d *LCDDisplay) FillScreen(c color.RGBA) {
	for i := range d.buf {
		d.buf[i] = RGBATo565(c)
	}
}

func (d *LCDDisplay) SetPixel(x, y int16, c color.RGBA) {
	if x < 0 || y < 0 || d.width < x || d.height < y {
		return
	}
	d.buf[x+y*d.width] = RGBATo565(c)
}

func (d *LCDDisplay) Size() (x, y int16) {
	return d.width, d.height
}

func (d *LCDDisplay) Display() error {
	return d.display.Display()
}

func (d *LCDDisplay) PrintClock(m, s uint8) {
	str := fmt.Sprintf("%02d:%02d", m, s)
	// d.display.FillRectangle(0, 0, 320, 240, white)
	// tinyfont.WriteLine(d.display, &freesans.Bold18pt7b, 0, 135, str, black)
	d.FillScreen(white)
	tinyfont.WriteLine(d, &freesans.Bold24pt7b, 1, 30, str, black)
	d.display.DrawRGBBitmap(0, 0, d.buf, d.width, d.height)
}

func (d *LCDDisplay) Println(s string) {
	d.display.FillRectangle(0, 0, 320, 240, white)
	tinyfont.WriteLine(d.display, &freesans.Bold24pt7b, 0, 135, s, black)
}

type Button struct {
	button machine.Pin
}

func NewButton(b machine.Pin) Button {
	b.Configure(machine.PinConfig{Mode: machine.PinInput})
	return Button{
		button: b,
	}
}

func (b Button) Get() bool {
	return !b.button.Get()
}

// Buzzer

type Buzzer struct {
	speaker tone.Speaker
	song    []tone.Note
	cnt     []uint16
	i       int
	c       uint16
}

func NewBuzzer() *Buzzer {
	bzrPin := machine.WIO_BUZZER
	pwm := machine.TCC0
	speaker, err := tone.New(pwm, bzrPin)
	if err != nil {
		println("failed to configure PWM")
		return nil
	}
	song := []tone.Note{
		tone.C5,
		tone.D5,
		tone.E5,
		tone.F5,
		tone.E5,
		tone.D5,
		tone.C5,
		tone.E5,
		tone.F5,
		tone.G5,
		tone.A5,
		tone.G5,
		tone.F5,
		tone.E5,
	}
	cnt := []uint16{
		200 * testpackage.MSec1,
		200 * testpackage.MSec1,
		200 * testpackage.MSec1,
		200 * testpackage.MSec1,
		200 * testpackage.MSec1,
		200 * testpackage.MSec1,
		400 * testpackage.MSec1,
		200 * testpackage.MSec1,
		200 * testpackage.MSec1,
		200 * testpackage.MSec1,
		200 * testpackage.MSec1,
		200 * testpackage.MSec1,
		200 * testpackage.MSec1,
		400 * testpackage.MSec1,
		200 * testpackage.MSec1,
		200 * testpackage.MSec1,
	}
	return &Buzzer{
		speaker: speaker,
		song:    song,
		cnt:     cnt,
		i:       0,
		c:       0,
	}
}

func (b *Buzzer) Sing() {
	if b.i < len(b.song) {
		if b.c == 0 {
			b.speaker.SetNote(b.song[b.i])
		}
		if b.c == b.cnt[b.i] {
			b.i++
			b.c = 0
		} else {
			b.c++
		}
	} else {
		b.Stop()
	}
}

func (b *Buzzer) Stop() {
	b.c = 0
	b.i = 0
	b.speaker.Stop()
}
