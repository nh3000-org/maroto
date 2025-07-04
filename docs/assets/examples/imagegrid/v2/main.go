package main

import (
	"log"

	"github.com/nh3000-org/maroto/v2/pkg/core"

	"github.com/nh3000-org/maroto/v2"

	"github.com/nh3000-org/maroto/v2/pkg/components/image"

	"github.com/nh3000-org/maroto/v2/pkg/config"
	"github.com/nh3000-org/maroto/v2/pkg/props"
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/imagegridv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/imagegridv2.txt")
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

	m.AddRow(40,
		image.NewFromFileCol(2, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  true,
			Percent: 80,
		}),
		image.NewFromFileCol(4, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  true,
			Percent: 80,
		}),
		image.NewFromFileCol(6, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  true,
			Percent: 80,
		}),
	)

	m.AddRow(40,
		image.NewFromFileCol(2, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  false,
			Percent: 50,
			Left:    10,
		}),
		image.NewFromFileCol(4, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  false,
			Percent: 50,
			Top:     10,
		}),
		image.NewFromFileCol(6, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  false,
			Percent: 50,
			Left:    15,
			Top:     15,
		}),
	)

	m.AddRow(40,
		image.NewFromFileCol(8, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  true,
			Percent: 80,
		}),
		image.NewFromFileCol(4, "docs/assets/images/biplane.jpg", props.Rect{
			Center:  true,
			Percent: 80,
		}),
	)

	m.AddRow(40,
		image.NewFromFileCol(6, "docs/assets/images/frontpage.png", props.Rect{
			Center:  false,
			Percent: 80,
			Top:     5,
			Left:    10,
		}),
		image.NewFromFileCol(4, "docs/assets/images/frontpage.png", props.Rect{
			Center:  false,
			Percent: 80,
			Top:     5,
		}),
		image.NewFromFileCol(2, "docs/assets/images/frontpage.png", props.Rect{
			Center:  false,
			Percent: 80,
			Left:    5,
		}),
	)

	m.AddRow(40,
		image.NewFromFileCol(6, "docs/assets/images/frontpage.png", props.Rect{
			Center:  true,
			Percent: 50,
		}),
		image.NewFromFileCol(4, "docs/assets/images/frontpage.png", props.Rect{
			Center:  true,
			Percent: 50,
		}),
		image.NewFromFileCol(2, "docs/assets/images/frontpage.png", props.Rect{
			Center:  true,
			Percent: 50,
		}),
	)

	m.AddRow(40,
		image.NewFromFileCol(4, "docs/assets/images/frontpage.png", props.Rect{
			Center:  true,
			Percent: 80,
		}),
		image.NewFromFileCol(8, "docs/assets/images/frontpage.png", props.Rect{
			Center:  true,
			Percent: 80,
		}),
	)
	m.AddAutoRow(
		image.NewFromFileCol(4, "docs/assets/images/frontpage.png", props.Rect{
			Center:             true,
			Percent:            20,
			JustReferenceWidth: true,
		}),
		image.NewFromFileCol(8, "docs/assets/images/frontpage.png", props.Rect{
			Center:             true,
			Percent:            30,
			JustReferenceWidth: true,
		}),
	)

	m.AddAutoRow(
		image.NewFromFileCol(2, "docs/assets/images/biplane.jpg", props.Rect{
			Center:             false,
			Percent:            50,
			Left:               10,
			JustReferenceWidth: true,
		}),
		image.NewFromFileCol(4, "docs/assets/images/biplane.jpg", props.Rect{
			Center:             false,
			Percent:            50,
			Top:                10,
			JustReferenceWidth: true,
		}),
		image.NewFromFileCol(6, "docs/assets/images/biplane.jpg", props.Rect{
			Center:             false,
			Percent:            50,
			Left:               15,
			Top:                15,
			JustReferenceWidth: true,
		}),
	)
	return m
}
