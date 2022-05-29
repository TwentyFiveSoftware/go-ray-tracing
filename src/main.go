package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"sync"
	"time"
)

const Width = 1920
const Height = 1080
const MaxRayTraceDepth = 50
const SamplesPerPixel = 100
const RenderGoroutines = 24

func main() {
	camera := NewCamera(Vec3{12.0, 2.0, -3.0}, Vec3{0.0, 0.0, 0.0}, 25.0, 10.0)
	scene := GenerateScene()

	yChannel := make(chan int, Height)
	for y := 0; y < Height; y++ {
		yChannel <- y
	}
	close(yChannel)

	type RowData struct {
		y   int
		row []color.RGBA
	}

	rowsChannel := make(chan RowData, Height)
	var waitGroup sync.WaitGroup

	renderStartTime := time.Now()

	for i := 0; i < RenderGoroutines; i++ {
		waitGroup.Add(1)
		go func() {
			for y := range yChannel {
				fmt.Printf("%d / %d (%.2f%%)\n", y+1, Height, float64(y+1)*100.0/Height)
				rowsChannel <- RowData{y, RenderRow(y, camera, scene)}
			}

			waitGroup.Done()
		}()
	}

	waitGroup.Wait()

	elapsedRenderTime := time.Since(renderStartTime)
	fmt.Printf("Rendered %d samples/pixel with %d goroutines in %d ms",
		SamplesPerPixel, RenderGoroutines, elapsedRenderTime.Milliseconds())

	img := image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: Width, Y: Height}})

	close(rowsChannel)
	for row := range rowsChannel {
		for x, rgb := range row.row {
			img.SetRGBA(x, row.y, rgb)
		}
	}

	saveImage(img)
}

func saveImage(image image.Image) {
	f, err := os.Create("render.png")

	if err != nil {
		fmt.Println(err)
		return
	}

	err = png.Encode(f, image)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
