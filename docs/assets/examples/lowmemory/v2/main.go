package main

import (
	"log"

	"github.com/nh3000-org/maroto/v2/pkg/core"

	"github.com/nh3000-org/maroto/v2"

	"github.com/nh3000-org/maroto/v2/pkg/components/text"

	"github.com/nh3000-org/maroto/v2/pkg/config"
	"github.com/nh3000-org/maroto/v2/pkg/props"
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/lowmemoryv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/lowmemoryv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto() core.Maroto {
	cfg := config.NewBuilder().
		WithSequentialLowMemoryMode(7).
		WithDebug(true).
		WithPageNumber().
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	for i := 0; i < 50; i++ {
		m.AddRows(
			text.NewRow(10, "Dummy text", props.Text{
				Size: 8,
			}),
		)
	}

	return m
}
