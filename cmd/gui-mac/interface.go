package main

import (
	"context"
	"log"
	"time"

	"github.com/caseymrm/menuet"
)

func (a _Application) runInterface() error {
	menuet.App().Label = "net.mniak.pismo.menubar"
	menuet.App().SetMenuState(&menuet.MenuState{
		Title: "Pismo",
	})
	menuet.App().Children = a.generateMenuItems
	menuet.App().RunApplication()
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

func (a _Application) generateMenuItems() []menuet.MenuItem {
	items := []menuet.MenuItem{}

	ctx := newContext()
	clockItems := a.generateClockMenuItems(ctx)
	otpItems := a.generateOTPMenuItems(ctx)

	items = append(items, clockItems...)
	items = append(items, menuet.MenuItem{Type: menuet.Separator})
	items = append(items, otpItems...)
	return items
}
