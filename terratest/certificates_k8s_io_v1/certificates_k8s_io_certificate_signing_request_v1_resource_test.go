/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package certificates_k8s_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCertificatesK8SIoCertificateSigningRequestV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_certificates_k8s_io_certificate_signing_request_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
