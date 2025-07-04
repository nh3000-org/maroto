package main

import (
	"log"

	"github.com/nh3000-org/maroto/v2"
	"github.com/nh3000-org/maroto/v2/pkg/components/page"
	"github.com/nh3000-org/maroto/v2/pkg/components/text"
	"github.com/nh3000-org/maroto/v2/pkg/config"
	"github.com/nh3000-org/maroto/v2/pkg/core"
)

func main() {
	m := GetMaroto()

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/addpagev2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/addpagev2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto() core.Maroto {
	cfg := config.NewBuilder().
		WithPageNumber().
		WithDebug(true).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	m.AddPages(
		page.New().Add(
			text.NewRow(30, "page1 row1"),
			text.NewRow(30, "page1 row2"),
			text.NewRow(30, "page1 row3"),
			text.NewRow(30, "page1 row4"),
			text.NewRow(30, "page1 row5"),
			text.NewRow(30, "page1 row6"),
			text.NewRow(30, "page1 row7"),
			text.NewRow(30, "page1 row8"),
			text.NewRow(30, "page1 row9"),
		),
		page.New().Add(
			text.NewRow(10, "page2 row1"),
		),
		page.New().Add(
			text.NewRow(10, "page3 row1"),
		),
	)

	return m
}
