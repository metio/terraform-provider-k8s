/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package provider_test

import (
	"bufio"
	"context"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/metio/terraform-provider-k8s/internal/provider"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestProviderDocumentation(t *testing.T) {
	t.SkipNow()
	file, err := os.Open("../../docs/data-sources/external_secrets_io_cluster_secret_store_v1alpha1_manifest.md")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	linkPattern := regexp.MustCompile(`<a id="nestedatt--(?P<DashedPath>.*)"></a>`)
	titlePattern := regexp.MustCompile("### Nested Schema for `(?P<DottedPath>.*)`")
	expectedTitle := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		link := linkPattern.FindStringSubmatch(line)
		title := titlePattern.FindStringSubmatch(line)

		if len(link) > 0 {
			attributeIndex := linkPattern.SubexpIndex("DashedPath")
			attributeValue := link[attributeIndex]
			expectedTitle = strings.ReplaceAll(attributeValue, "--", ".")
		}
		if len(title) > 0 {
			pathIndex := titlePattern.SubexpIndex("DottedPath")
			pathValue := title[pathIndex]
			if expectedTitle != pathValue {
				t.Errorf("Wanted [%s] but got [%s]", expectedTitle, pathValue)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
}

func TestProvider_Schema(t *testing.T) {
	ctx := context.Background()
	schemaRequest := fwprovider.SchemaRequest{}
	schemaResponse := &fwprovider.SchemaResponse{}

	provider.New().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)
	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}
}
