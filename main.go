package main

import (
    "github.com/gizak/termui" // <- ui shortcut, optional
    log "github.com/Sirupsen/logrus"
    //"fmt"
)

func main() {
    log.Info("init")

    err := termui.Init()
    if err != nil {
        panic(err)
    }
    defer termui.Close()

    termui.Handle("/sys/kbd/q", func(termui.Event) {
        termui.StopLoop()
    })
    termui.Handle("/sys/kbd/C-c", func(termui.Event) {
        termui.StopLoop()
    })

    termui.Handle("/sys/wnd/resize", func(e termui.Event) {
        termui.Body.Width = termui.TermWidth()
        termui.Body.Align()
        termui.Clear()
        termui.Render(termui.Body)
    })

    g := termui.NewGauge()
    init_ui(g)

    //"<left>", "<right>"
    termui.Handle("/sys/kbd/<left>", func(termui.Event) {
        g.Percent = g.Percent - 1
        termui.Render(termui.Body)
    })
    termui.Handle("/sys/kbd/<right>", func(termui.Event) {
        g.Percent = g.Percent + 1
        termui.Render(termui.Body)
    })

    termui.Loop() // block until StopLoop is called
}


func init_ui(g* termui.Gauge) {
    widget0 := termui.NewPar("line 1")
    widget0.Border = false
    widget1 := termui.NewPar("line 2")
    widget1.Border = false

    par := termui.NewPar("<> This row has 3 columns\n<- Widgets can be stacked up like left side\n<- Stacked widgets are treated as a single widget")
    par.Height = 5

    status := termui.NewPar("[H] Help  ")
    status.TextBgColor = termui.StringToAttribute("blue")
    //termui.StringToAttribute()
    status.Border = false
    status.Height = 5
    //status.Width = termui.Body.Width

    g.Percent = 50
    g.Width = 50
    g.Height = 3
    g.BorderLabel = "Slim Gauge"
    g.BarColor = termui.ColorRed
    g.PercentColor = termui.ColorBlue


    // calculate layout
    termui.Body.AddRows(
        termui.NewRow(termui.NewCol(12, 0, widget0)),
        termui.NewRow(termui.NewCol(12, 0, widget1)),
        termui.NewRow(termui.NewCol(12, 0, par)),
        termui.NewRow(termui.NewCol(12, 0, g)),
        termui.NewRow(termui.NewCol(12, 0, termui.Hline{}.Buffer())),
        termui.NewRow(termui.NewCol(12, 0, status)))

    termui.Body.Align()
    termui.Render(termui.Body)
}
