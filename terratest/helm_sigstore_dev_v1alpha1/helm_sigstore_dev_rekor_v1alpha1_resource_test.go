/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package helm_sigstore_dev_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHelmSigstoreDevRekorV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_helm_sigstore_dev_rekor_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
