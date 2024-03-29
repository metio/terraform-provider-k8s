/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"fmt"
	"io/fs"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	clientschema "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/scheme"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var deserializer = clientschema.Codecs.UniversalDeserializer()

func ParseCRDv1Files(basePath string) []*apiextensionsv1.CustomResourceDefinition {
	crds := make([]*apiextensionsv1.CustomResourceDefinition, 0)

	err := filepath.WalkDir(basePath, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && strings.HasSuffix(path, ".yaml") {
			file, fileErr := os.ReadFile(path)
			if fileErr != nil {
				return fmt.Errorf("error reading %s: %v", path, fileErr)
			}
			crd, parseErr := parseCRDv1(file)
			if parseErr != nil {
				return fmt.Errorf("error parsing %s: %v", path, parseErr)
			}
			crds = append(crds, crd)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return crds
}

func parseCRDv1(data []byte) (*apiextensionsv1.CustomResourceDefinition, error) {
	object, _, err := deserializer.Decode(data, nil, nil)
	if err != nil {
		return nil, err
	}

	crd, ok := object.(*apiextensionsv1.CustomResourceDefinition)
	if !ok {
		return nil, fmt.Errorf("could not cast to apiextensionsv1.CustomResourceDefinition")
	}

	return crd, nil
}
