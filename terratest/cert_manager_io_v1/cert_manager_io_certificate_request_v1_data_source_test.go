/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package cert_manager_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCertManagerIoCertificateRequestV1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_cert_manager_io_certificate_request_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
