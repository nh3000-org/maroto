package main

import (
	"log"

	"github.com/nh3000-org/maroto/v2/pkg/core"

	"github.com/nh3000-org/maroto/v2"

	"github.com/nh3000-org/maroto/v2/pkg/components/text"

	"github.com/nh3000-org/maroto/v2/pkg/consts/align"
	"github.com/nh3000-org/maroto/v2/pkg/consts/fontstyle"

	"github.com/nh3000-org/maroto/v2/pkg/config"
	"github.com/nh3000-org/maroto/v2/pkg/props"
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/headerv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/headerv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto() core.Maroto {
	cfg := config.NewBuilder().
		WithDebug(true).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	err := m.RegisterHeader(text.NewRow(20, "Header", props.Text{
		Size:  10,
		Style: fontstyle.Bold,
		Align: align.Center,
	}))
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 50; i++ {
		m.AddRows(
			text.NewRow(10, "Dummy text", props.Text{
				Size: 8,
			}),
		)
	}

	return m
}
