package main

import (
	"log"

	"github.com/nh3000-org/maroto/v2/pkg/core"

	"github.com/nh3000-org/maroto/v2"

	"github.com/nh3000-org/maroto/v2/pkg/components/line"
	"github.com/nh3000-org/maroto/v2/pkg/consts/linestyle"
	"github.com/nh3000-org/maroto/v2/pkg/consts/orientation"
	"github.com/nh3000-org/maroto/v2/pkg/props"

	"github.com/nh3000-org/maroto/v2/pkg/config"
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/linegridv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/linegridv2.txt")
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
		line.NewCol(2),
		line.NewCol(4),
		line.NewCol(6),
	)

	m.AddRow(40,
		line.NewCol(6),
		line.NewCol(4),
		line.NewCol(2),
	)

	m.AddRow(40,
		line.NewCol(2, props.Line{Thickness: 0.5}),
		line.NewCol(4, props.Line{Color: &props.RedColor}),
		line.NewCol(6, props.Line{Orientation: orientation.Vertical}),
	)

	m.AddRow(40,
		line.NewCol(6, props.Line{OffsetPercent: 50}),
		line.NewCol(4, props.Line{OffsetPercent: 50, Orientation: orientation.Vertical}),
		line.NewCol(2, props.Line{SizePercent: 50}),
	)

	m.AddRow(40,
		line.NewCol(2, props.Line{Style: linestyle.Dashed}),
		line.NewCol(4, props.Line{Color: &props.RedColor, Style: linestyle.Dashed, Thickness: 0.8, Orientation: orientation.Vertical, OffsetPercent: 70, SizePercent: 70}),
		line.NewCol(6, props.Line{Color: &props.RedColor, Style: linestyle.Dashed, Thickness: 0.8, Orientation: orientation.Horizontal, OffsetPercent: 40, SizePercent: 40}),
	)

	m.AddAutoRow(
		line.NewCol(2, props.Line{Style: linestyle.Dashed}),
		line.NewCol(4, props.Line{Color: &props.RedColor, Style: linestyle.Dashed, Thickness: 0.8, Orientation: orientation.Vertical, OffsetPercent: 70, SizePercent: 70}),
		line.NewCol(6, props.Line{Color: &props.RedColor, Style: linestyle.Dashed, Thickness: 0.8, Orientation: orientation.Horizontal, OffsetPercent: 40, SizePercent: 40}),
	)
	return m
}
