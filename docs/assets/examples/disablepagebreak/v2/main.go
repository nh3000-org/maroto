package main

import (
	"log"
	"os"

	"github.com/nh3000-org/maroto/v2/pkg/consts/extension"

	"github.com/nh3000-org/maroto/v2"
	"github.com/nh3000-org/maroto/v2/pkg/config"
	"github.com/nh3000-org/maroto/v2/pkg/consts/orientation"

	"github.com/nh3000-org/maroto/v2/pkg/components/page"
	"github.com/nh3000-org/maroto/v2/pkg/core"
)

func main() {
	backgroundImage := "docs/assets/images/certificate.png"
	m := GetMaroto(backgroundImage)
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/disablepagebreakv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/disablepagebreakv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto(image string) core.Maroto {
	bytes, err := os.ReadFile(image)
	if err != nil {
		log.Fatal(err)
	}
	b := config.NewBuilder().
		WithTopMargin(0).
		WithRightMargin(0).
		WithLeftMargin(0).
		WithDimensions(361.8, 203.2).
		WithDisableAutoPageBreak(true).
		WithOrientation(orientation.Horizontal).
		WithMaxGridSize(20).
		WithBackgroundImage(bytes, extension.Png).
		Build()

	b.Margins.Bottom = 0

	mrt := maroto.New(b)
	m := maroto.NewMetricsDecorator(mrt)

	m.AddPages(page.New())
	return m
}
