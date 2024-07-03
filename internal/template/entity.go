package template

import (
	"context"
	"github.com/dave/jennifer/jen"
	"io"
)

const PackageDatabase = "database"

type GoFileWriter io.Writer

// WriteEntity generate entities files based on the passed entities configuration and
// writes those to io.Writer.
func WriteEntity(ctx context.Context, entity map[string]string, writer GoFileWriter) error {
	f := jen.NewFile()
}
