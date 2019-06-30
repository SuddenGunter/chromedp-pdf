package pdfstore

import "io"

// PdfStore is the interface that handles PDF file saving after printing.
type PdfStore interface {
	io.Writer
}
