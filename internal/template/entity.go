package template

import (
	"context"
	"fmt"
	"io"

	"github.com/dave/jennifer/jen"
)

const PackageDatabase = "database"

type GoFileWriter io.Writer

// WriteEntity generate entities files based on the passed entities configuration and
// writes those to io.Writer.
func WriteEntity(_ context.Context, entityName string, fields map[string]string, writer GoFileWriter) error {
	f := jen.NewFile(PackageDatabase)

	codeFields := make([]jen.Code, 0, len(fields))

	for fName, fType := range fields {
		field := jen.Id(fName)
		switch fType {
		case "string":
			field = field.String()
		case "int":
			field = field.Int64()
		case "datetime":
			field = field.Qual("time", "Time")
		}

		codeFields = append(codeFields, field)
	}

	f.Type().Id(entityName).Struct(codeFields...)

	if err := f.Render(writer); err != nil {
		return fmt.Errorf("render entity: %w", err)
	}

	return nil
}
