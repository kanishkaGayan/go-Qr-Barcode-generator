package main

import (

	"bytes"
	"image"
	"image/png"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/skip2/go-qrcode"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("QR Barcode Generator")
	window.Resize(fyne.NewSize(600, 500))

	// Create QR Code tab content
	qrTab := createQRTab(window)

	// Create Barcode tab content
	barcodeTab := createBarcodeTab(window)

	// Create tabs
	tabs := container.NewAppTabs(
		container.NewTabItem("QR Code", qrTab),
		container.NewTabItem("Barcode", barcodeTab),
	)

	window.SetContent(tabs)
	window.ShowAndRun()
}

func createQRTab(window fyne.Window) *fyne.Container {
	// Input field with fixed width
	qrInput := widget.NewEntry()
	qrInput.SetPlaceHolder("Enter text for QR code")
	qrInput.Resize(fyne.NewSize(400, 40))

	// Generate button with fixed width
	generateBtn := widget.NewButton("Generate QR Code", nil)
	generateBtn.Resize(fyne.NewSize(150, 40))

	// Image display with fixed size
	qrImage := canvas.NewImageFromResource(nil)
	qrImage.Resize(fyne.NewSize(300, 300))
	qrImage.FillMode = canvas.ImageFillOriginal
	qrImage.ScaleMode = canvas.ImageScaleSmooth

	// Save button with fixed width (initially hidden)
	saveBtn := widget.NewButton("Save QR Code", nil)
	saveBtn.Resize(fyne.NewSize(150, 40))
	saveBtn.Hide()

	var qrData []byte

	// Generate QR code function
	generateBtn.OnTapped = func() {
		if qrInput.Text == "" {
			dialog.ShowInformation("Error", "Please enter text to generate QR code", window)
			return
		}

		// Generate QR code
		qrCode, err := qrcode.Encode(qrInput.Text, qrcode.Medium, 256)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		qrData = qrCode

		// Convert bytes to image and set it
		img := convertBytesToImage(qrCode)
		if img != nil {
			qrImage.Image = img
			qrImage.Refresh()
			saveBtn.Show()
		} else {
			dialog.ShowError(err, window)
		}
	}

	// Save QR code function
	saveBtn.OnTapped = func() {
		if qrData == nil {
			return
		}

		saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			if writer == nil {
				return
			}

			defer writer.Close()
			_, err = writer.Write(qrData)
			if err != nil {
				dialog.ShowError(err, window)
				return
			}

			dialog.ShowInformation("Success", "QR code saved successfully!", window)
		}, window)

		saveDialog.SetFileName("qrcode.png")
		saveDialog.Show()
	}

	// Layout with fixed positions
	content := container.NewVBox(
		widget.NewLabel("Enter text to generate QR code:"),
		container.NewPadded(qrInput),
		container.NewPadded(generateBtn),
		container.NewPadded(qrImage),
		container.NewPadded(saveBtn),
	)

	return content
}

func createBarcodeTab(window fyne.Window) *fyne.Container {
	// Input field with fixed width
	barcodeInput := widget.NewEntry()
	barcodeInput.SetPlaceHolder("Enter text for barcode")
	barcodeInput.Resize(fyne.NewSize(400, 40))

	// Generate button with fixed width
	generateBtn := widget.NewButton("Generate Barcode", nil)
	generateBtn.Resize(fyne.NewSize(150, 40))

	// Image display with fixed size
	barcodeImage := canvas.NewImageFromResource(nil)
	barcodeImage.Resize(fyne.NewSize(400, 150))
	barcodeImage.FillMode = canvas.ImageFillOriginal
	barcodeImage.ScaleMode = canvas.ImageScaleSmooth

	// Save button with fixed width (initially hidden)
	saveBtn := widget.NewButton("Save Barcode", nil)
	saveBtn.Resize(fyne.NewSize(150, 40))
	saveBtn.Hide()

	var barcodeData image.Image

	// Generate barcode function
	generateBtn.OnTapped = func() {
		if barcodeInput.Text == "" {
			dialog.ShowInformation("Error", "Please enter text to generate barcode", window)
			return
		}

		// Generate Code128 barcode
		barcodeCode, err := code128.Encode(barcodeInput.Text)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		// Scale the barcode
		scaledBarcode, err := barcode.Scale(barcodeCode, 400, 100)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		barcodeData = scaledBarcode
		barcodeImage.Image = scaledBarcode
		barcodeImage.Refresh()
		saveBtn.Show()
	}

	// Save barcode function
	saveBtn.OnTapped = func() {
		if barcodeData == nil {
			return
		}

		saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			if writer == nil {
				return
			}

			defer writer.Close()
			err = png.Encode(writer, barcodeData)
			if err != nil {
				dialog.ShowError(err, window)
				return
			}

			dialog.ShowInformation("Success", "Barcode saved successfully!", window)
		}, window)

		saveDialog.SetFileName("barcode.png")
		saveDialog.Show()
	}

	// Layout with fixed positions
	content := container.NewVBox(
		widget.NewLabel("Enter text to generate barcode:"),
		container.NewPadded(barcodeInput),
		container.NewPadded(generateBtn),
		container.NewPadded(barcodeImage),
		container.NewPadded(saveBtn),
	)

	return content
}

func convertBytesToImage(data []byte) image.Image {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Printf("Error decoding image: %v", err)
		return nil
	}
	return img
}
