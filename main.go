package main

import (
	"flag"
	"fmt"
	"image/color"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	var f = flag.String("f", "", "output filename")
	flag.Parse()
	args := flag.Args()
	if len(args) == 1 && flag.NFlag() == 0 {
		msg := args[0]
		qrCode, _ := qr.Encode(msg, qr.M, qr.Auto)
		printConsole(qrCode)
	} else if len(args) == 1 && flag.NFlag() > 0 {
		msg := args[0]
		qrCode, _ := qr.Encode(msg, qr.M, qr.Auto)
		outputFile(qrCode, *f)
	} else {
		flag.Usage()
		os.Exit(1)
	}

}

func outputFile(qrCode barcode.Barcode, filename string) {
	//qrCode, _ = barcode.Scale(qrCode, 30, 30)
	file, _ := os.Create(filename)
	defer file.Close()
	png.Encode(file, qrCode)
	fmt.Println("output:", filename)
}

func printConsole(qrCode barcode.Barcode) {
	const black = "\033[40m  \033[0m"
	const white = "\033[47m  \033[0m"

	rect := qrCode.Bounds()
	fmt.Println()
	for y := rect.Min.Y; y <= rect.Max.Y+1; y++ {
		for x := rect.Min.X; x <= rect.Max.X+1; x++ {
			if x == rect.Min.X || x == rect.Max.X+1 || y == rect.Min.Y || y == rect.Max.Y+1 {
				fmt.Print(white)
			} else if qrCode.At(x-1, y-1) == color.Black {
				fmt.Print(black)
			} else {
				fmt.Print(white)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
