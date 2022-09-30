/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package terratest

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestResource(t *testing.T) {
	t.SkipNow()
	err := filepath.WalkDir("../examples/resources", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() && d.Name() != "resources" {
			if _, statErr := os.Stat(path + "/resource.tf"); statErr == nil {

				t.Run(d.Name(), func(t *testing.T) {
					terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
						TerraformDir: path,
						NoColor:      true,
					})

					defer cleanupAfterTest(t, path)
					defer func(t *testing.T, options *terraform.Options) {
						_, tfErr := terraform.DestroyE(t, options)
						if tfErr != nil {
							t.Fatal(tfErr)
						}
					}(t, terraformOptions)

					_, tfErr := terraform.InitAndApplyAndIdempotentE(t, terraformOptions)
					if tfErr != nil {
						t.Fatal(tfErr)
					}

					outputMap := terraform.OutputMap(t, terraformOptions, "resources")
					for key, value := range outputMap {
						assert.NotEmpty(t, value, fmt.Sprintf("resource %s.%s did not produce an output", d.Name(), key))
					}
				})

			} else {
				return statErr
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func cleanupAfterTest(t *testing.T, path string) {
	err := os.RemoveAll(path + "/.terraform.lock.hcl")
	if err != nil {
		t.Fatal(err)
	}
	err = os.RemoveAll(path + "/terraform.tfstate")
	if err != nil {
		t.Fatal(err)
	}
	err = os.RemoveAll(path + "/terraform.tfstate.backup")
	if err != nil {
		t.Fatal(err)
	}
	err = os.RemoveAll(path + "/.terraform")
	if err != nil {
		t.Fatal(err)
	}
}
