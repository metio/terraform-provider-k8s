//go:build k8s

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package k8s

import (
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/getkin/kin-openapi/openapi3"
	"io/fs"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	clientschema "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/scheme"
	"net/url"
	"os"
	"path/filepath"
)

var deserializer = clientschema.Codecs.UniversalDeserializer()

func ParseAllCustomResourceDefinitions() []*apiextensionsv1.CustomResourceDefinition {
	crds := make([]*apiextensionsv1.CustomResourceDefinition, 0)

	err := filepath.WalkDir("../../schemas/crd_v1/", func(path string, d fs.DirEntry, err error) error {
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
		return nil, fmt.Errorf("could not cast to apiextensionsv1.CustomResourceDefinition")
	}

	return crd, nil
}

func ParseKubernetesSwagger() map[string]*openapi3.SchemaRef {
	input, err := os.ReadFile("../../schemas/openapi_v2/io.kubernetes/kubernetes-swagger.json")
	if err != nil {
		panic(err)
	}

	var doc openapi2.T
	err = json.Unmarshal(input, &doc)
	if err != nil {
		panic(err)
	}
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, uri *url.URL) ([]byte, error) {
		return os.ReadFile(uri.Path)
	}
	openapi3.CircularReferenceCounter = 10
	v3, err := openapi2conv.ToV3(&doc)
	if err != nil {
		panic(err)
	}
	err = loader.ResolveRefsIn(v3, &url.URL{Path: filepath.ToSlash("../../schemas/openapi_v2/io.kubernetes/kubernetes-swagger.json")})
	if err != nil {
		panic(err)
	}

	return v3.Components.Schemas
}
