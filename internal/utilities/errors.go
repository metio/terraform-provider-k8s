/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package utilities

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"net/http"
)

func GetNamespacedResourceError(err error, name string, namespace string) diag.Diagnostic {
	if diagnostic := notFoundGetError(err, fmt.Sprintf("The requested resource cannot be found. "+
		"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
		"Namespace: %s\n"+
		"Name: %s", namespace, name)); diagnostic != nil {
		return diagnostic
	}
	return genericGetError(err)
}

func GetResourceError(err error, name string) diag.Diagnostic {
	if diagnostic := notFoundGetError(err, fmt.Sprintf("The requested resource cannot be found. "+
		"Make sure that it does exist in your cluster and you have set the correct name configured.\n\n"+
		"Name: %s", name)); diagnostic != nil {
		return diagnostic
	}
	return genericGetError(err)
}

func notFoundGetError(err error, description string) diag.Diagnostic {
	var statusError *k8sErrors.StatusError
	if errors.As(err, &statusError) {
		if statusError.Status().Code == http.StatusNotFound {
			return diag.NewErrorDiagnostic("Unable to find resource", description)
		}
	}
	return nil
}

func genericGetError(err error) diag.ErrorDiagnostic {
	return diag.NewErrorDiagnostic(
		"Unable to GET resource",
		fmt.Sprintf("An unexpected error occurred while reading the resource. "+
			"Please report this issue to the provider developers.\n\n"+
			"GET Error (%T): %s", err, err.Error()),
	)
}

func MarshalJsonError(err error) diag.ErrorDiagnostic {
	return diag.NewErrorDiagnostic(
		"Unable to marshal response",
		"Please report this issue to the provider developers.\n\n"+
			"Marshal Error: "+err.Error(),
	)
}

func JsonUnmarshalError(err error) diag.ErrorDiagnostic {
	return diag.NewErrorDiagnostic(
		"Unable to unmarshal resource",
		"An unexpected error occurred while parsing the resource read response. "+
			"Please report this issue to the provider developers.\n\n"+
			"JSON Error: "+err.Error(),
	)
}

func JsonMarshalError(err error) diag.ErrorDiagnostic {
	return diag.NewErrorDiagnostic(
		"Unable to marshal resource",
		"An unexpected error occurred while marshalling the resource. "+
			"Please report this issue to the provider developers.\n\n"+
			"JSON Error: "+err.Error(),
	)
}

func MarshalYamlError(err error) diag.ErrorDiagnostic {
	return diag.NewErrorDiagnostic(
		"Unable to marshal resource",
		"An unexpected error occurred while marshalling the resource. "+
			"Please report this issue to the provider developers.\n\n"+
			"YAML Error: "+err.Error(),
	)
}

func OfflineProviderError() diag.ErrorDiagnostic {
	return diag.NewErrorDiagnostic(
		"Provider in Offline Mode",
		"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
			"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
	)
}

func UnexpectedDataSourceDataError(data any) diag.ErrorDiagnostic {
	return diag.NewErrorDiagnostic(
		"Unexpected Data Source Configure Type",
		fmt.Sprintf("Expected *utilities.DataSourceData, got: %T. Please report this issue to the provider developers.", data),
	)
}

func UnexpectedResourceDataError(data any) diag.ErrorDiagnostic {
	return diag.NewErrorDiagnostic(
		"Unexpected Resource Configure Type",
		fmt.Sprintf("Expected *utilities.ResourceData, got: %T. Please report this issue to the provider developers.", data),
	)
}

func IsDeletionError(err error) bool {
	return err != nil && !k8sErrors.IsNotFound(err) && !k8sErrors.IsGone(err)
}

func IsNotFound(err error) bool {
	return err != nil && k8sErrors.IsNotFound(err)
}

func PatchError(err error) diag.ErrorDiagnostic {
	return diag.NewErrorDiagnostic(
		"Unable to PATCH resource",
		fmt.Sprintf("An unexpected error occurred while updating the resource. "+
			"Please report this issue to the provider developers.\n\n"+
			"PATCH Error (%T): %s", err, err.Error()),
	)
}

func DeleteError(err error) diag.ErrorDiagnostic {
	return diag.NewErrorDiagnostic(
		"Unable to DELETE resource",
		fmt.Sprintf("An unexpected error occurred while deleting the resource. "+
			"Please report this issue to the provider developers.\n\n"+
			"DELETE Error (%T): %s", err, err.Error()),
	)
}

func WaitTimeoutExceeded() diag.ErrorDiagnostic {
	return diag.NewErrorDiagnostic(
		"Wait Timeout Exceeded",
		"The allocated maximum wait time was exceeded. Your resource might still be deleted, but it was "+
			"not removed from Terraform state yet. Re-run 'terraform apply' and optionally increase the "+
			"'timeout' parameter to wait a longer period of time.",
	)
}
