package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func fileExists(filename *string) bool {
	info, err := os.Stat(*filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func main() {
	inputFilePath := flag.String("i", "", "Your input file")
	outputFilePath := flag.String("o", "", "Your output file")
	withFilePath := flag.String("w", "", "Used to find subboundaries in input file")
	addCentroid := flag.Bool("c", false, "Calculates centroid and inserts into output file")

	flag.Parse()

	if *inputFilePath == "" {
		fmt.Fprintf(os.Stderr, "You must provide an input file!\n")
		os.Exit(ErrNoInputFile)
	}

	if *outputFilePath == "" {
		fmt.Fprintf(os.Stderr, "You must provide an output file!\n")
		os.Exit(ErrNoOutputFile)
	}

	if !fileExists(inputFilePath) {
		fmt.Fprintf(os.Stderr, "Input file does not exist. Check your path.!\n")
		os.Exit(ErrInvalidInputFile)
	}

	var withShapes *FeatureCollection
	if *withFilePath != "" {
		var err error

		if !fileExists(withFilePath) {
			fmt.Fprintf(os.Stderr, "With file does not exist. Check your path.!\n")
			os.Exit(ErrInvalidWithFile)
		}

		start := time.Now()
		withShapes, err = BuildFeatures(withFilePath)
		duration := time.Since(start)
		fmt.Printf("With file indexed in %s!\n", duration)

		if err != nil {
			panic("Error opening with file")
		}
	}

	if *addCentroid {
		fmt.Println("Calculating centroids...")
	}

	start := time.Now()
	inputShapes, err := BuildFeatures(inputFilePath)
	duration := time.Since(start)

	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Printf("Input file indexed in %s!\n", duration)

	if err != nil {
		panic("Error opening input file")
	}

	outputShapes := inputShapes.Clone()
	outputShapes.SaveToShapeFile(outputFilePath)

	l0 := len(inputShapes.Features)
	l1 := len(withShapes.Features)

	fmt.Printf("Found %d features in input file & %d in with file!\n", l0, l1)

	os.Exit(0)
}
