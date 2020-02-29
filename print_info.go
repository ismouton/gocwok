package main

import (
	"fmt"

	"github.com/jonas-p/go-shp"
)

func fieldTypeToString(fieldType byte) *string {
	var humanReadableFieldType string

	switch fieldType {
	case 67:
		humanReadableFieldType = "String"
	case 78:
		humanReadableFieldType = "Number"
	}

	return &humanReadableFieldType
}

func printInfo(inputFile *string) {
	shape, _ := shp.OpenZip(*inputFile)
	fields := shape.Fields()

	fmt.Println("Field Information:")
	for _, field := range fields {
		fmt.Printf("    %s:\t%s\n", field, *fieldTypeToString(field.Fieldtype))
	}
}
