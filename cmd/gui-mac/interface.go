package main

import (
	"context"
	"log"
	"time"

	"github.com/caseymrm/menuet"
	"github.com/samber/lo"
)

func (a _Application) runInterface() error {
	app := menuet.App()
	app.Label = "net.mniak.qisno"
	app.SetMenuState(&menuet.MenuState{
		Title: a.Title,
	})
	app.Children = a.generateMenuItems(app)
	app.RunApplication()
	return nil
}

func (a _Application) ShowMessage(msg string) {
	// TODO: implement show message
	log.Println("show message:", msg)
}

func newContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	return ctx
}

type ItemsHolder struct {
	Items []*menuet.MenuItem
}

func (a _Application) generateMenuItems(menuetApp *menuet.Application) func() []menuet.MenuItem {
	keepAliveChan := make(chan any)
	go monitorStatus(keepAliveChan)

	var holder ItemsHolder
	return func() []menuet.MenuItem {
		keepAliveChan <- true
		if !isOpen {
			whenOpen(a, &holder)
		}
		go func() {
			time.Sleep(time.Second)
			menuetApp.MenuChanged()
		}()
		return lo.Map(holder.Items, func(mi *menuet.MenuItem, i int) menuet.MenuItem {
			return *mi
		})
	}
}

var isOpen bool

func monitorStatus(keepAliveChan chan any) {
	for {
		dl100ms, _ := context.WithTimeout(context.Background(), 1100*time.Millisecond)
		select {
		case <-dl100ms.Done():
			if isOpen {
				isOpen = false
			}
		case <-keepAliveChan:
			if !isOpen {
				isOpen = true
			}
		}
	}
}

func whenOpen(a _Application, holder *ItemsHolder) {
	items := make([]*menuet.MenuItem, 0)
	// items = append(items, menuet.MenuItem{
	// 	Text: "Abriu pela primeira vez",
	// })
	ctx := newContext()
	// vpnItems := a.generateVPNMenuItems(ctx)

	// clockItems := a.generateClockMenuItems(ctx)
	// items = append(items, clockItems...)
	// items = append(items, menuet.MenuItem{Type: menuet.Separator})

	otpItems := a.generateOTPMenuItems(ctx)
	items = append(items, otpItems...)
	// items = append(items, menuet.MenuItem{Type: menuet.Separator})
	// items = append(items, vpnItems...)

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	items[0].Text = "Alterei"
	// }()

	holder.Items = items
}
