package main

import (
	"os"
	"fmt"
	"time"
	"github.com/b00lduck/raspberry-datalogger-display/framebuffer"
	"github.com/b00lduck/raspberry-datalogger-display/touchscreen"
	"github.com/b00lduck/raspberry-datalogger-display/gui"
	"github.com/b00lduck/raspberry-datalogger-display/gui/pages"
	"github.com/b00lduck/raspberry-datalogger-display/webserver"
)

func main() {

	fmt.Println("START")

	fb := new(framebuffer.Framebuffer)
	fb.Open(os.Args[1])
	defer fb.Close()

	ts := new(touchscreen.Touchscreen)
	ts.Open(os.Args[2])
	defer ts.Close()
	go ts.Run()

	g := gui.NewGui(fb, ts)
	g.SetMainPage(pages.NewMainPage())
	gasPage := pages.NewGasPage()
	g.SetPage("GAS_1", gasPage)
	g.SelectPage("GAS_1")
	go g.Run(&ts.Event)

	ws := webserver.NewWebserver(fb, &ts.Event)
	go ws.Run()

	for {
		g.Process()
		time.Sleep(1 * time.Second)
	}

}

