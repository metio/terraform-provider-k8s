//go:build fetcher

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import (
	"bytes"
	"fmt"
	goyaml "gopkg.in/yaml.v3"
	"io"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	clientschema "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/scheme"
	"k8s.io/cli-runtime/pkg/printers"
	"log"
	"os"
	"path/filepath"
)

var deserializer = clientschema.Codecs.UniversalDeserializer()

func downloadCRDv1(targetDirectory string) {
	temp := createTemporaryDirectory()
	defer os.RemoveAll(temp)

	for _, url := range crdv1Sources {
		log.Printf("downloading [%s]", url)
		file := createTemporaryFile(temp)
		rawUrl := githubRawUrl(url)
		rawUrl = gitlabRawUrl(rawUrl)
		err := downloadFile(file.Name(), rawUrl)
		if err != nil {
			log.Printf("cannot handle [%s] because of: %s", url, err)
			continue
		}
		crds, err := parseCRDv1(file.Name())
		if err != nil {
			log.Printf("cannot handle [%s] because of: %s", url, err)
			continue
		}

		for _, crd := range crds {
			targetFile := fmt.Sprintf("%s/%s/%s/%s.yaml", targetDirectory, crd.Spec.Group, crd.Spec.Versions[0].Name, crd.Spec.Names.Plural)
			writeYaml(crd, targetFile)
		}
	}
}

func createTemporaryFile(baseDirectory string) *os.File {
	file, err := os.CreateTemp(baseDirectory, "crdv1")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func createTemporaryDirectory() string {
	temp, err := os.MkdirTemp("", "terraform-provider-k8s")
	if err != nil {
		log.Fatal(err)
	}
	return temp
}

func parseCRDv1(filePath string) ([]*apiextensionsv1.CustomResourceDefinition, error) {
	crds := make([]*apiextensionsv1.CustomResourceDefinition, 0)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	yamls, err := splitYAML(data)
	if err != nil {
		return nil, err
	}
	for _, resourceYAML := range yamls {
		if len(resourceYAML) == 0 {
			continue
		}

		object, _, err := deserializer.Decode(resourceYAML, nil, nil)
		if err != nil {
			log.Printf("could not decode because of %s", err)
			continue
		}

		crd, ok := object.(*apiextensionsv1.CustomResourceDefinition)
		if !ok {
			log.Print("could not cast to apiextensionsv1.CustomResourceDefinition")
		} else {
			crds = append(crds, splitVersions(crd)...)
		}
	}

	return crds, nil
}

func splitVersions(crd *apiextensionsv1.CustomResourceDefinition) []*apiextensionsv1.CustomResourceDefinition {
	versions := make([]*apiextensionsv1.CustomResourceDefinition, 0)

	for _, version := range crd.Spec.Versions {
		copied := crd.DeepCopy()
		copied.Spec.Versions = []apiextensionsv1.CustomResourceDefinitionVersion{version}
		versions = append(versions, copied)
	}

	return versions
}

func writeYaml(crd *apiextensionsv1.CustomResourceDefinition, destPath string) {
	err := os.MkdirAll(filepath.Dir(destPath), 0777)
	if err != nil {
		log.Printf("could not create directory because of %s", err)
		return
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		log.Printf("could not create file because of %s", err)
		return
	}
	defer outputFile.Close()
	printer := printers.YAMLPrinter{}
	err = printer.PrintObj(crd, outputFile)
	if err != nil {
		log.Fatalf("could not write yaml to %s: %s", destPath, err)
	}
}

func splitYAML(resources []byte) ([][]byte, error) {
	dec := goyaml.NewDecoder(bytes.NewReader(resources))

	var res [][]byte
	for {
		var value interface{}
		err := dec.Decode(&value)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		valueBytes, err := goyaml.Marshal(value)
		if err != nil {
			return nil, err
		}
		res = append(res, valueBytes)
	}
	return res, nil
}
