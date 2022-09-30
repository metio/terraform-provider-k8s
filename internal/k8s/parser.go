//go:build k8s

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package k8s

import (
	"fmt"
	"io/fs"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	clientschema "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/scheme"
	"os"
	"path/filepath"
)

var deserializer = clientschema.Codecs.UniversalDeserializer()

func ParseAllCustomResourceDefinitions() []*apiextensionsv1.CustomResourceDefinition {
	crds := make([]*apiextensionsv1.CustomResourceDefinition, 0)

	err := filepath.WalkDir("../../custom-resource-definitions/", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			file, fileErr := os.ReadFile(path)
			if fileErr != nil {
				return fmt.Errorf("error reading %s: %v", path, fileErr)
			}
			crd, parseErr := ParseCustomResourceDefinition(file)
			if parseErr != nil {
				return fmt.Errorf("error parsing %s: %v", path, fileErr)
			}
			crds = append(crds, crd)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	return crds
}

func ParseCustomResourceDefinition(data []byte) (*apiextensionsv1.CustomResourceDefinition, error) {
	object, _, err := deserializer.Decode(data, nil, nil)
	if err != nil {
		return nil, err
	}

	crd, ok := object.(*apiextensionsv1.CustomResourceDefinition)
	if !ok {
		return nil, fmt.Errorf("could not cast to CRD")
	}

	return crd, nil
}
