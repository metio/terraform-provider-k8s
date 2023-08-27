/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package testutilities

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic/fake"
	"regexp"
	"testing"
)

func VerifyConfigurationErrors(t *testing.T, blockType string, name string, testCase ConfigurationErrorTestCase) {
	client := fake.NewSimpleDynamicClient(runtime.NewScheme())
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: ProviderFactories(client),
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig() + fmt.Sprintf(`
							%s "%s" "test" {
								%s
							}
						`, blockType, name, testCase.Configuration),
				ExpectError: regexp.MustCompile(testCase.ErrorRegex),
			},
		},
	})
}

func VerifyCannotBeUsedOffline(t *testing.T, blockType string, name string, configuration string) {
	client := fake.NewSimpleDynamicClient(runtime.NewScheme())
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: ProviderFactories(client),
		Steps: []resource.TestStep{
			{
				Config: OfflineProviderConfig() + fmt.Sprintf(`
							%s "%s" "test" {
								%s
							}
						`, blockType, name, configuration),
				ExpectError: regexp.MustCompile("Error: Provider in Offline Mode"),
			},
		},
	})
}
