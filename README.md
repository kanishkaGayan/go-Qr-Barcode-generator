# QR & Barcode Generator

A simple desktop application built with Go and Fyne that generates QR codes and barcodes with an intuitive tabbed interface.

## Features

- **QR Code Generation**: Generate QR codes from any text input
- **Barcode Generation**: Create Code128 barcodes from text
- **Tabbed Interface**: Clean UI with separate tabs for QR codes and barcodes
- **Save Functionality**: Save generated codes as PNG images
- **Fixed Layout**: Consistent element sizing for professional appearance

## Screenshots

![]
<img width="789" height="696" alt="Screenshot from 2025-08-27 18-55-01" src="https://github.com/user-attachments/assets/c4f2ce5b-bd9e-4215-a560-b65de5b4c0a7" />
![]
<img width="789" height="696" alt="Screenshot from 2025-08-27 18-55-09" src="https://github.com/user-attachments/assets/6179e472-9405-48f5-ab10-924dfae86cb6" />
![]
<img width="789" height="696" alt="Screenshot from 2025-08-27 18-55-17" src="https://github.com/user-attachments/assets/8316ba30-2965-471d-bbec-2c203f4149cc" />

## Installation

1. Clone the repository:
```bash
git clone https://github.com/kanishkaGayan/go-Qr-Barcode-generator.git
cd go-Qr-Barcode-generator
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
