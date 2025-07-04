package main

import (
	"log"

	"github.com/nh3000-org/maroto/v2/pkg/core"

	"github.com/nh3000-org/maroto/v2"

	"github.com/nh3000-org/maroto/v2/pkg/consts/orientation"

	"github.com/nh3000-org/maroto/v2/pkg/components/text"
	"github.com/nh3000-org/maroto/v2/pkg/config"
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/orientationv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/orientationv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto() core.Maroto {
	cfg := config.NewBuilder().
		WithOrientation(orientation.Horizontal).
		WithDebug(true).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	m.AddRows(
		text.NewRow(30, "content"),
	)

	return m
}
