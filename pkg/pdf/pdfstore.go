package pdf

import "io"

// Store is the interface that handles PDF file save.
type Store interface {
	io.Writer
}
