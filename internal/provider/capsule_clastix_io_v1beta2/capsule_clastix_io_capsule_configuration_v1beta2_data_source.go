/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package capsule_clastix_io_v1beta2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &CapsuleClastixIoCapsuleConfigurationV1Beta2DataSource{}
	_ datasource.DataSourceWithConfigure = &CapsuleClastixIoCapsuleConfigurationV1Beta2DataSource{}
)

func NewCapsuleClastixIoCapsuleConfigurationV1Beta2DataSource() datasource.DataSource {
	return &CapsuleClastixIoCapsuleConfigurationV1Beta2DataSource{}
}

type CapsuleClastixIoCapsuleConfigurationV1Beta2DataSource struct {
	kubernetesClient dynamic.Interface
}

type CapsuleClastixIoCapsuleConfigurationV1Beta2DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		EnableTLSReconciler *bool `tfsdk:"enable_tls_reconciler" json:"enableTLSReconciler,omitempty"`
		ForceTenantPrefix   *bool `tfsdk:"force_tenant_prefix" json:"forceTenantPrefix,omitempty"`
		NodeMetadata        *struct {
			ForbiddenAnnotations *struct {
				Denied      *[]string `tfsdk:"denied" json:"denied,omitempty"`
				DeniedRegex *string   `tfsdk:"denied_regex" json:"deniedRegex,omitempty"`
			} `tfsdk:"forbidden_annotations" json:"forbiddenAnnotations,omitempty"`
			ForbiddenLabels *struct {
				Denied      *[]string `tfsdk:"denied" json:"denied,omitempty"`
				DeniedRegex *string   `tfsdk:"denied_regex" json:"deniedRegex,omitempty"`
			} `tfsdk:"forbidden_labels" json:"forbiddenLabels,omitempty"`
		} `tfsdk:"node_metadata" json:"nodeMetadata,omitempty"`
		Overrides *struct {
			TLSSecretName                      *string `tfsdk:"tls_secret_name" json:"TLSSecretName,omitempty"`
			MutatingWebhookConfigurationName   *string `tfsdk:"mutating_webhook_configuration_name" json:"mutatingWebhookConfigurationName,omitempty"`
			ValidatingWebhookConfigurationName *string `tfsdk:"validating_webhook_configuration_name" json:"validatingWebhookConfigurationName,omitempty"`
		} `tfsdk:"overrides" json:"overrides,omitempty"`
		ProtectedNamespaceRegex *string   `tfsdk:"protected_namespace_regex" json:"protectedNamespaceRegex,omitempty"`
		UserGroups              *[]string `tfsdk:"user_groups" json:"userGroups,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CapsuleClastixIoCapsuleConfigurationV1Beta2DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_capsule_clastix_io_capsule_configuration_v1beta2"
}

func (r *CapsuleClastixIoCapsuleConfigurationV1Beta2DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CapsuleConfiguration is the Schema for the Capsule configuration API.",
		MarkdownDescription: "CapsuleConfiguration is the Schema for the Capsule configuration API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"metadata": schema.SingleNestedAttribute{
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Optional:            false,
				Computed:            false,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.NameValidator(),
							stringvalidator.LengthAtLeast(1),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "CapsuleConfigurationSpec defines the Capsule configuration.",
				MarkdownDescription: "CapsuleConfigurationSpec defines the Capsule configuration.",
				Attributes: map[string]schema.Attribute{
					"enable_tls_reconciler": schema.BoolAttribute{
						Description:         "Toggles the TLS reconciler, the controller that is able to generate CA and certificates for the webhooks when not using an already provided CA and certificate, or when these are managed externally with Vault, or cert-manager.",
						MarkdownDescription: "Toggles the TLS reconciler, the controller that is able to generate CA and certificates for the webhooks when not using an already provided CA and certificate, or when these are managed externally with Vault, or cert-manager.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"force_tenant_prefix": schema.BoolAttribute{
						Description:         "Enforces the Tenant owner, during Namespace creation, to name it using the selected Tenant name as prefix, separated by a dash. This is useful to avoid Namespace name collision in a public CaaS environment.",
						MarkdownDescription: "Enforces the Tenant owner, during Namespace creation, to name it using the selected Tenant name as prefix, separated by a dash. This is useful to avoid Namespace name collision in a public CaaS environment.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"node_metadata": schema.SingleNestedAttribute{
						Description:         "Allows to set the forbidden metadata for the worker nodes that could be patched by a Tenant. This applies only if the Tenant has an active NodeSelector, and the Owner have right to patch their nodes.",
						MarkdownDescription: "Allows to set the forbidden metadata for the worker nodes that could be patched by a Tenant. This applies only if the Tenant has an active NodeSelector, and the Owner have right to patch their nodes.",
						Attributes: map[string]schema.Attribute{
							"forbidden_annotations": schema.SingleNestedAttribute{
								Description:         "Define the annotations that a Tenant Owner cannot set for their nodes.",
								MarkdownDescription: "Define the annotations that a Tenant Owner cannot set for their nodes.",
								Attributes: map[string]schema.Attribute{
									"denied": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"denied_regex": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"forbidden_labels": schema.SingleNestedAttribute{
								Description:         "Define the labels that a Tenant Owner cannot set for their nodes.",
								MarkdownDescription: "Define the labels that a Tenant Owner cannot set for their nodes.",
								Attributes: map[string]schema.Attribute{
									"denied": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"denied_regex": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"overrides": schema.SingleNestedAttribute{
						Description:         "Allows to set different name rather than the canonical one for the Capsule configuration objects, such as webhook secret or configurations.",
						MarkdownDescription: "Allows to set different name rather than the canonical one for the Capsule configuration objects, such as webhook secret or configurations.",
						Attributes: map[string]schema.Attribute{
							"tls_secret_name": schema.StringAttribute{
								Description:         "Defines the Secret name used for the webhook server. Must be in the same Namespace where the Capsule Deployment is deployed.",
								MarkdownDescription: "Defines the Secret name used for the webhook server. Must be in the same Namespace where the Capsule Deployment is deployed.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"mutating_webhook_configuration_name": schema.StringAttribute{
								Description:         "Name of the MutatingWebhookConfiguration which contains the dynamic admission controller paths and resources.",
								MarkdownDescription: "Name of the MutatingWebhookConfiguration which contains the dynamic admission controller paths and resources.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"validating_webhook_configuration_name": schema.StringAttribute{
								Description:         "Name of the ValidatingWebhookConfiguration which contains the dynamic admission controller paths and resources.",
								MarkdownDescription: "Name of the ValidatingWebhookConfiguration which contains the dynamic admission controller paths and resources.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"protected_namespace_regex": schema.StringAttribute{
						Description:         "Disallow creation of namespaces, whose name matches this regexp",
						MarkdownDescription: "Disallow creation of namespaces, whose name matches this regexp",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"user_groups": schema.ListAttribute{
						Description:         "Names of the groups for Capsule users.",
						MarkdownDescription: "Names of the groups for Capsule users.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *CapsuleClastixIoCapsuleConfigurationV1Beta2DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *CapsuleClastixIoCapsuleConfigurationV1Beta2DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_capsule_clastix_io_capsule_configuration_v1beta2")

	var data CapsuleClastixIoCapsuleConfigurationV1Beta2DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "capsule.clastix.io", Version: "v1beta2", Resource: "capsuleconfigurations"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name configured.\n\n"+
						"Name: %s", data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CapsuleClastixIoCapsuleConfigurationV1Beta2DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("capsule.clastix.io/v1beta2")
	data.Kind = pointer.String("CapsuleConfiguration")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
