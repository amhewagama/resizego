package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"sync"

	"golang.org/x/image/draw"
)

var (
	inputDir     string
	maxWidth      = 1000
	maxHeight     = 1000
	maxWidth_lg   = 2000
	maxHeight_lg  = 2000
	jpegQuality   = 80
	concurrency   = 4
	basename      string
)

func init() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run yourprogram.go inputDir")
		os.Exit(1)
	}

	inputDir = os.Args[1]

    // Check if the directory exists
    if _, err := os.Stat(inputDir); os.IsNotExist(err) {
        fmt.Println("Invalid source dir:", inputDir)
        os.Exit(1)
    }
	
	basename = fmt.Sprintf("lankascape-%s-", inputDir)
}

func main() {

	// Create the output directories for JPEG if they do not exist
	outputDirJPG := fmt.Sprintf("%s-sm", inputDir)
	outputDirJPG_lg := fmt.Sprintf("%s-lg", inputDir)

	createOutputDir(outputDirJPG_lg)
	createOutputDir(outputDirJPG)

	// Read all files in the directory
	files, err := filepath.Glob(filepath.Join(inputDir, "*"))
	if err != nil {
		fmt.Println("Error reading input directory:", err)
		return
	}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, concurrency)

	// Process each file concurrently
	for index, file := range files {
		wg.Add(1)
		semaphore <- struct{}{} // Acquire a semaphore
		go func(filePath string, index int) {
			defer func() {
				<-semaphore // Release the semaphore when done
				wg.Done()
			}()
			processImage(filePath, outputDirJPG_lg, outputDirJPG, index+1)
		}(file, index)
	}

	wg.Wait()
	fmt.Println("Processing completed.")
}

func createOutputDir(outputDir string) {
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.Mkdir(outputDir, os.ModeDir)
		if err != nil {
			fmt.Printf("Error creating %s directory: %v\n", outputDir, err)
		} else {
			fmt.Printf("Created %s directory\n", outputDir)
		}
	}
}

func processImage(filePath, outputDirJPG_lg, outputDirJPG string, index int) {
	// Open the image file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	// Resize the image
	resizedImg := resize(img)

	// Create the output file for JPEG
	outputFilePathJPG := filepath.Join(outputDirJPG, fmt.Sprintf("%s%d.jpg", basename, index))
	outputFileJPG, err := os.Create(outputFilePathJPG)
	if err != nil {
		fmt.Println("Error creating JPEG output file:", err)
		return
	}
	defer outputFileJPG.Close()

	// Encode and save
	err = jpeg.Encode(outputFileJPG, resizedImg, &jpeg.Options{Quality: jpegQuality})
	if err != nil {
		fmt.Println("Error encoding JPEG image:", err)
		return
	}else{
		fmt.Println("Files saved in: ",outputFilePathJPG)
	}

	// Create the output LG
	resizedImg_lg := resize_lg(img)

	outputFilePathJPG_lg := filepath.Join(outputDirJPG_lg, fmt.Sprintf("%s%d.jpg", basename, index))
	outputFileJPG_lg, err := os.Create(outputFilePathJPG_lg)
	if err != nil {
		fmt.Println("Error creating JPEG output file:", err)
		return
	}
	defer outputFileJPG_lg.Close()

	// Encode and save
	err = jpeg.Encode(outputFileJPG_lg, resizedImg_lg, &jpeg.Options{Quality: jpegQuality})
	if err != nil {
		fmt.Println("Error encoding JPEG image:", err)
		return
	}else{
		fmt.Println("Files saved in: ",outputFilePathJPG_lg)
	}	

	//fmt.Printf("Processed: %s\n", filePath)
}

func resize(img image.Image) image.Image {
	// Get the original width and height
	originalWidth := img.Bounds().Dx()
	originalHeight := img.Bounds().Dy()

	// Calculate the new width and height while maintaining the aspect ratio
	var newWidth, newHeight int

	if originalWidth > originalHeight {
		// Landscape image
		newWidth = maxWidth
		newHeight = originalHeight * maxWidth / originalWidth
	} else {
		// Portrait or square image
		newWidth = originalWidth * maxHeight / originalHeight
		newHeight = maxHeight
	}

	// Resize the image using Bilinear interpolation
	resizedImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.BiLinear.Scale(resizedImg, resizedImg.Bounds(), img, img.Bounds(), draw.Over, nil)

	return resizedImg
}

func resize_lg(img image.Image) image.Image {
	// Get the original width and height
	originalWidth := img.Bounds().Dx()
	originalHeight := img.Bounds().Dy()

	// Calculate the new width and height while maintaining the aspect ratio
	var newWidth, newHeight int

	if originalWidth > originalHeight {
		// Landscape image
		newWidth = maxWidth_lg
		newHeight = originalHeight * maxWidth_lg / originalWidth
	} else {
		// Portrait or square image
		newWidth = originalWidth * maxHeight_lg / originalHeight
		newHeight = maxHeight_lg
	}

	// Resize the image using Bilinear interpolation
	resizedImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.BiLinear.Scale(resizedImg, resizedImg.Bounds(), img, img.Bounds(), draw.Over, nil)

	return resizedImg
}
