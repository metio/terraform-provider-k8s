/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package events_k8s_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEventsK8SIoEventV1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_events_k8s_io_event_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
