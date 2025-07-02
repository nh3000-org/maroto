package main

import (
	"testing"

	"github.com/nh3000-org/maroto/v2/pkg/test"
)

func TestGetMaroto(t *testing.T) {
	// Act
	sut := GetMaroto()

	// Assert
	test.New(t).Assert(sut.GetStructure()).Equals("examples/list.json")
}
