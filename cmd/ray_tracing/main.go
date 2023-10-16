package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"sync"
	"time"

	"go-ray-tracing/pkg/camera"
	"go-ray-tracing/pkg/renderer"
	"go-ray-tracing/pkg/scene"
	"go-ray-tracing/pkg/vec3"
)

const (
	width             = 1920
	height            = 1080
	maxRayTraceDepth  = 50
	samplesPerPixel   = 100
	renderWorkerCount = 24
)

func main() {
	camera := camera.NewCamera(vec3.Vec3{X: 12.0, Y: 2.0, Z: -3.0}, vec3.Vec3{}, 25.0, 10.0, width, height)
	scene := scene.GenerateScene()

	yChannel := make(chan int, height)
	for y := 0; y < height; y++ {
		yChannel <- y
	}
	close(yChannel)

	type RowData struct {
		y   int
		row []color.RGBA
	}

	rowsChannel := make(chan RowData, height)
	var waitGroup sync.WaitGroup

	renderStartTime := time.Now()

	for i := 0; i < renderWorkerCount; i++ {
		waitGroup.Add(1)
		go func() {
			for y := range yChannel {
				fmt.Printf("%d / %d (%.2f%%)\n", y+1, height, float64(y+1)*100.0/height)
				rowsChannel <- RowData{y, renderer.RenderRow(y, camera, scene, width, height, samplesPerPixel, maxRayTraceDepth)}
			}

			waitGroup.Done()
		}()
	}

	waitGroup.Wait()

	elapsedRenderTime := time.Since(renderStartTime)
	fmt.Printf("rendered %d samples/pixel with %d workers in %d ms\n",
		samplesPerPixel, renderWorkerCount, elapsedRenderTime.Milliseconds())

	img := image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: width, Y: height}})

	close(rowsChannel)
	for row := range rowsChannel {
		for x, rgb := range row.row {
			img.SetRGBA(x, row.y, rgb)
		}
	}

	saveImage("render.png", img)
}

func saveImage(path string, image image.Image) {
	f, err := os.Create(path)

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
