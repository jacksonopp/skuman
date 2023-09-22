package sku

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/jacksonopp/htmx-app/internal/logger"
)

type Item struct {
	Name string
	Sku  string
}

func ParseFile(file multipart.File, header *multipart.FileHeader, err error) ([]Item, error) {
	if err != nil {
		return nil, newParseError(unknown, fmt.Sprintf("error parsing file: %v", err))
	}

	name := strings.Split(header.Filename, ".")
	if name[1] != "csv" {
		return nil, newParseError(fileFormat, "file is not a csv")
	}

	logger.Infof("parsing file: %s.%s", name[0], name[1])

	var buf bytes.Buffer
	defer buf.Reset()

	io.Copy(&buf, file)

	reader := csv.NewReader(&buf)
	columns, err := reader.Read()
	if err != nil {
		return nil, newParseError(unknown, fmt.Sprintf("error reading csv headers: %v", err.Error()))
	}

	if columns[0] != "Name" || columns[1] != "SKU" {
		return nil, newParseError(csvHead, "incorrectly formatted csv head")
	}

	data, err := reader.ReadAll()
	if err != nil {
		return nil, newParseError(unknown, fmt.Sprintf("error reading csv body %v", err.Error()))
	}

	items := []Item{}

	// TODO write data to database
	for _, item := range data {
		items = append(items, Item{
			Name: item[0],
			Sku:  item[1],
		})
	}

	return items, nil
}
