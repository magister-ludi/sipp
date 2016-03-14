// First hacky attempt at getting a UI up and running. This currently doesn't do
// anything useful at all.
package main

import (
	"flag"
    "fmt"
    "image"
//    "image/draw"
    "os"

	"gopkg.in/qml.v1"

	"github.com/Causticity/sipp/simage"
	"github.com/Causticity/sipp/stree"
    //"github.com/Causticity/sipp/sgrad"
    //"github.com/Causticity/sipp/shist"
)

var srcName *string
//var k *int
var src simage.SippImage

func main() {
	/*
	// This will become a parameter to the gradient op.
	k = flag.Int("K", 0, "Number of bins to scale the max radius to. "+
						 "The histogram will be 2K+1 bins on a side.\n"+
						 "        This is used only for 16-bit images.\n"+
						 "        If K is omitted, it is computed from "+
						 "the maximum excursion of the gradient.\n"+
						 "        8-bit images always use a 511x511 histogram, "+
						 "as that covers the entire possible space.")

	// This test will move to the gradient op. Specifically, it won't be 
	// manipulable in the UI for 8-bit images, but it will be displayed.
	if src.Bpp() == 8 {
		*k = 255
		fmt.Println("Image is 8-bit. K forced to 255.")
	}
	*/
	srcName = flag.String("in", "", "input image file; must be grayscale png")
	flag.Parse()
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

var appComponent qml.Object

var app *qml.Window
var currentTreeRoot *stree.SippNode

func run() error {
	engine := qml.NewEngine()
	engine.AddImageProvider("thumb", thumbProvider)
	engine.AddImageProvider("src", srcProvider)

	appComponent, err := engine.LoadFile("sippui.qml")
	if err != nil {
		return err
	}

	err = stree.InitTreeComponents(engine)
	if err != nil {
		return err
	}

	app = appComponent.CreateWindow(nil)
	app.On("gotFile", newTree)
	app.Show()

	if len(*srcName) > 0 {
		newTree(*srcName)
	} else {
		app.Call("getFile")
	}

	app.Wait()

	return nil
}

func newTree (url string) {
	currentTreeRoot = stree.NewSippTree(url)
	currentTreeRoot.BuildUI(url)
}

func srcProvider(srcName string, width, height int) image.Image {
	return currentTreeRoot.Src[0]
}

func thumbProvider(srcName string, width, height int) image.Image {
	return currentTreeRoot.Src[0].Thumbnail()
}
