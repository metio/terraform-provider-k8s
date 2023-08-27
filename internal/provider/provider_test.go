/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package provider_test

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/metio/terraform-provider-k8s/internal/provider"
	"k8s.io/client-go/dynamic"
)

func providerFactories(client dynamic.Interface) map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		"k8s": providerserver.NewProtocol6WithError(provider.NewWithClient(client)),
	}
}

func providerConfig() string {
	return fmt.Sprintf(`
		provider "k8s" {
			
		}
	`)
}
