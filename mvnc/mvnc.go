package main

import ui "github.com/gizak/termui" // <- ui shortcut, optional
//import "fmt"

func main() {
    //fmt.Println("Hello world")
    err := ui.Init()
    if err != nil {
        panic(err)
    }
    defer ui.Close()

    widget0 := ui.NewPar(":PRESS q TO QUIT DEMO")
    widget2 := ui.NewPar(":PRESS q TO QUIT DEMO")
    widget30 := ui.NewPar(":PRESS q TO QUIT DEMO")
    widget31 := ui.NewPar(":PRESS q TO QUIT DEMO")
    widget32 := ui.NewPar(":PRESS q TO QUIT DEMO")
    widget4 := ui.NewPar(":PRESS q TO QUIT DEMO")

    // build
    ui.Body.AddRows(
        ui.NewRow(
            ui.NewCol(12, 0, widget0)),
        ui.NewRow(
            ui.NewCol(3, 0, widget2),
            ui.NewCol(3, 0, widget30, widget31, widget32),
            ui.NewCol(3, 0, widget4)))

    // calculate layout
    ui.Body.Align()

    ui.Render(ui.Body)

    ui.Handle("/sys/kbd/q", func(ui.Event) {
        ui.StopLoop()
    })
    ui.Handle("/sys/kbd/C-c", func(ui.Event) {
        ui.StopLoop()
    })

    ui.Handle("/sys/kbd", func(ui.Event) {
        // handle all other key pressing
        ui.Render(ui.Body)
    })

    // handle a 1s timer
    ui.Handle("/timer/1s", func(e ui.Event) {
        t := e.Data.(ui.EvtTimer)
        // t is a EvtTimer
        if t.Count%2 ==0 {
            // do something
        }
    })
    ui.Handle("/sys/wnd/resize", func(e ui.Event) {
        ui.Body.Width = ui.TermWidth()
        ui.Body.Align()
        ui.Clear()
        ui.Render(ui.Body)
    })


    ui.Loop() // block until StopLoop is called

}
