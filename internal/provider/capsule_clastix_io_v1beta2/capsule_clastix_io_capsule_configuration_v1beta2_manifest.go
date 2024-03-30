/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package capsule_clastix_io_v1beta2

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &CapsuleClastixIoCapsuleConfigurationV1Beta2Manifest{}
)

func NewCapsuleClastixIoCapsuleConfigurationV1Beta2Manifest() datasource.DataSource {
	return &CapsuleClastixIoCapsuleConfigurationV1Beta2Manifest{}
}

type CapsuleClastixIoCapsuleConfigurationV1Beta2Manifest struct{}

type CapsuleClastixIoCapsuleConfigurationV1Beta2ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

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

func (r *CapsuleClastixIoCapsuleConfigurationV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_capsule_clastix_io_capsule_configuration_v1beta2_manifest"
}

func (r *CapsuleClastixIoCapsuleConfigurationV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CapsuleConfiguration is the Schema for the Capsule configuration API.",
		MarkdownDescription: "CapsuleConfiguration is the Schema for the Capsule configuration API.",
		Attributes: map[string]schema.Attribute{
			"yaml": schema.StringAttribute{
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
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
						Optional:            true,
						Computed:            false,
						Validators: []validator.Map{
							validators.LabelValidator(),
						},
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
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
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"force_tenant_prefix": schema.BoolAttribute{
						Description:         "Enforces the Tenant owner, during Namespace creation, to name it using the selected Tenant name as prefix, separated by a dash. This is useful to avoid Namespace name collision in a public CaaS environment.",
						MarkdownDescription: "Enforces the Tenant owner, during Namespace creation, to name it using the selected Tenant name as prefix, separated by a dash. This is useful to avoid Namespace name collision in a public CaaS environment.",
						Required:            false,
						Optional:            true,
						Computed:            false,
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
										Optional:            true,
										Computed:            false,
									},

									"denied_regex": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
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
										Optional:            true,
										Computed:            false,
									},

									"denied_regex": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"overrides": schema.SingleNestedAttribute{
						Description:         "Allows to set different name rather than the canonical one for the Capsule configuration objects, such as webhook secret or configurations.",
						MarkdownDescription: "Allows to set different name rather than the canonical one for the Capsule configuration objects, such as webhook secret or configurations.",
						Attributes: map[string]schema.Attribute{
							"tls_secret_name": schema.StringAttribute{
								Description:         "Defines the Secret name used for the webhook server. Must be in the same Namespace where the Capsule Deployment is deployed.",
								MarkdownDescription: "Defines the Secret name used for the webhook server. Must be in the same Namespace where the Capsule Deployment is deployed.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"mutating_webhook_configuration_name": schema.StringAttribute{
								Description:         "Name of the MutatingWebhookConfiguration which contains the dynamic admission controller paths and resources.",
								MarkdownDescription: "Name of the MutatingWebhookConfiguration which contains the dynamic admission controller paths and resources.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"validating_webhook_configuration_name": schema.StringAttribute{
								Description:         "Name of the ValidatingWebhookConfiguration which contains the dynamic admission controller paths and resources.",
								MarkdownDescription: "Name of the ValidatingWebhookConfiguration which contains the dynamic admission controller paths and resources.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"protected_namespace_regex": schema.StringAttribute{
						Description:         "Disallow creation of namespaces, whose name matches this regexp",
						MarkdownDescription: "Disallow creation of namespaces, whose name matches this regexp",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"user_groups": schema.ListAttribute{
						Description:         "Names of the groups for Capsule users.",
						MarkdownDescription: "Names of the groups for Capsule users.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *CapsuleClastixIoCapsuleConfigurationV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_capsule_clastix_io_capsule_configuration_v1beta2_manifest")

	var model CapsuleClastixIoCapsuleConfigurationV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("capsule.clastix.io/v1beta2")
	model.Kind = pointer.String("CapsuleConfiguration")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
