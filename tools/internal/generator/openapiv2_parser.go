/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"fmt"
	"github.com/pb33f/libopenapi"
	v2high "github.com/pb33f/libopenapi/datamodel/high/v2"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ParseOpenAPIv2Files(basePath string) []v2high.Swagger {
	schemas := make([]v2high.Swagger, 0)

	err := filepath.WalkDir(basePath, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && strings.HasSuffix(path, ".json") {
			fileContent, readErr := os.ReadFile(path)
			if readErr != nil {
				return readErr
			}
			document, docErr := libopenapi.NewDocument(fileContent)
			if docErr != nil {
				return docErr
			}
			docModel, errors := document.BuildV2Model()
			if len(errors) > 0 {
				for i := range errors {
					fmt.Printf("error: %e\n", errors[i])
				}
				panic(fmt.Sprintf("cannot create v3 model from document: %d errors reported", len(errors)))
			}
			schemas = append(schemas, docModel.Model)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return schemas
}
