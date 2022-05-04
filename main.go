package main

import (
	"flag"
	"fmt"
	"image/color"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	msg := args[0]
	qrCode, _ := qr.Encode(msg, qr.M, qr.Auto)
	//qrCode, _ = barcode.Scale(qrCode, 30, 30)
	black := "\033[40m  \033[0m"
	white := "\033[47m  \033[0m"
	printQrCode(qrCode, black, white)
}

func printQrCode(qrCode barcode.Barcode, black, white string) {
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
