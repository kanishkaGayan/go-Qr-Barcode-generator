# QR & Barcode Generator

A simple desktop application built with Go and Fyne that generates QR codes and barcodes with an intuitive tabbed interface.

## Features

- **QR Code Generation**: Generate QR codes from any text input
- **Barcode Generation**: Create Code128 barcodes from text
- **Tabbed Interface**: Clean UI with separate tabs for QR codes and barcodes
- **Save Functionality**: Save generated codes as PNG images
- **Fixed Layout**: Consistent element sizing for professional appearance

## Screenshots

![QR Barcode Generator](screenshot.png)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go-qr-barcode-generator.git
cd go-qr-barcode-generator
```

2. Install dependencies:
```bash
go mod tidy
```

3. Build the application:
```bash
go build -o qr-barcode-generator qr-barcode-generator.go
```

4. Run the application:
```bash
./qr-barcode-generator
```

## Usage

1. **QR Code Tab**: Enter text and click "Generate QR Code"
2. **Barcode Tab**: Enter text and click "Generate Barcode"
3. Click "Save" to export generated images to your preferred directory

## Dependencies

- [Fyne](https://fyne.io/) - Cross-platform GUI toolkit
- [go-qrcode](https://github.com/skip2/go-qrcode) - QR code generation
- [barcode](https://github.com/boombuler/barcode) - Barcode generation

## Requirements

- Go 1.19 or higher
- Linux/Windows/macOS

## License

MIT License
