/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"encoding/json"
	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/getkin/kin-openapi/openapi3"
	"io/fs"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var openapi3Loader *openapi3.Loader

func init() {
	openapi3Loader = openapi3.NewLoader()
	openapi3Loader.IsExternalRefsAllowed = true
	openapi3Loader.ReadFromURIFunc = func(loader *openapi3.Loader, uri *url.URL) ([]byte, error) {
		return os.ReadFile(uri.Path)
	}
}

func ParseOpenAPIv2Files(basePath string) []map[string]*openapi3.SchemaRef {
	schemas := make([]map[string]*openapi3.SchemaRef, 0)

	err := filepath.WalkDir(basePath, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && strings.HasSuffix(path, ".json") {
			openapiv2, parseErr := parseOpenAPIv2(path)
			if parseErr != nil {
				return parseErr
			}
			openapiv3, conversionErr := convertV2toV3(openapiv2)
			if conversionErr != nil {
				return conversionErr
			}
			resolveErr := resolveReferences(path, openapiv3)
			if resolveErr != nil {
				return resolveErr
			}
			schemas = append(schemas, openapiv3.Components.Schemas)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return schemas
}

func parseOpenAPIv2(filePath string) (*openapi2.T, error) {
	input, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var doc openapi2.T
	err = json.Unmarshal(input, &doc)
	if err != nil {
		return nil, err
	}
	return &doc, nil

}

func resolveReferences(filePath string, v3 *openapi3.T) error {
	err := openapi3Loader.ResolveRefsIn(v3, &url.URL{Path: filePath})
	if err != nil {
		return err
	}
	return nil
}

func convertV2toV3(doc *openapi2.T) (*openapi3.T, error) {
	v3, err := openapi2conv.ToV3(doc)
	if err != nil {
		return nil, err
	}
	return v3, nil
}
