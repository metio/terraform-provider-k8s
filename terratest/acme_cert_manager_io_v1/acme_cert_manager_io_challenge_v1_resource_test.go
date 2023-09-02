/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package acme_cert_manager_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAcmeCertManagerIoChallengeV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_acme_cert_manager_io_challenge_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
