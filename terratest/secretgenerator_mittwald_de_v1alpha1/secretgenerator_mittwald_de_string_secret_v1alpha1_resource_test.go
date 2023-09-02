/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package secretgenerator_mittwald_de_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSecretgeneratorMittwaldDeStringSecretV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_secretgenerator_mittwald_de_string_secret_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
