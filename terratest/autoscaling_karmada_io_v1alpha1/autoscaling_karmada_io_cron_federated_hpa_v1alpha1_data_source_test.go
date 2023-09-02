/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package autoscaling_karmada_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAutoscalingKarmadaIoCronFederatedHPAV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_autoscaling_karmada_io_cron_federated_hpa_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}