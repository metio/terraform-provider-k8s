/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package secrets_doppler_com_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &SecretsDopplerComDopplerSecretV1Alpha1Manifest{}
)

func NewSecretsDopplerComDopplerSecretV1Alpha1Manifest() datasource.DataSource {
	return &SecretsDopplerComDopplerSecretV1Alpha1Manifest{}
}

type SecretsDopplerComDopplerSecretV1Alpha1Manifest struct{}

type SecretsDopplerComDopplerSecretV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Config        *string `tfsdk:"config" json:"config,omitempty"`
		Format        *string `tfsdk:"format" json:"format,omitempty"`
		Host          *string `tfsdk:"host" json:"host,omitempty"`
		ManagedSecret *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Type      *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"managed_secret" json:"managedSecret,omitempty"`
		NameTransformer *string `tfsdk:"name_transformer" json:"nameTransformer,omitempty"`
		Processors      *struct {
			AsName *string `tfsdk:"as_name" json:"asName,omitempty"`
			Type   *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"processors" json:"processors,omitempty"`
		Project       *string   `tfsdk:"project" json:"project,omitempty"`
		ResyncSeconds *int64    `tfsdk:"resync_seconds" json:"resyncSeconds,omitempty"`
		Secrets       *[]string `tfsdk:"secrets" json:"secrets,omitempty"`
		TokenSecret   *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"token_secret" json:"tokenSecret,omitempty"`
		VerifyTLS *bool `tfsdk:"verify_tls" json:"verifyTLS,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SecretsDopplerComDopplerSecretV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_secrets_doppler_com_doppler_secret_v1alpha1_manifest"
}

func (r *SecretsDopplerComDopplerSecretV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DopplerSecret is the Schema for the dopplersecrets API",
		MarkdownDescription: "DopplerSecret is the Schema for the dopplersecrets API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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

					"namespace": schema.StringAttribute{
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
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
				Description:         "DopplerSecretSpec defines the desired state of DopplerSecret",
				MarkdownDescription: "DopplerSecretSpec defines the desired state of DopplerSecret",
				Attributes: map[string]schema.Attribute{
					"config": schema.StringAttribute{
						Description:         "The Doppler config",
						MarkdownDescription: "The Doppler config",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"format": schema.StringAttribute{
						Description:         "Format enables the downloading of secrets as a file",
						MarkdownDescription: "Format enables the downloading of secrets as a file",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("json", "dotnet-json", "env", "yaml", "docker"),
						},
					},

					"host": schema.StringAttribute{
						Description:         "The Doppler API host",
						MarkdownDescription: "The Doppler API host",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"managed_secret": schema.SingleNestedAttribute{
						Description:         "The Kubernetes secret where the operator will store and sync the fetched secrets",
						MarkdownDescription: "The Kubernetes secret where the operator will store and sync the fetched secrets",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "The name of the Secret resource",
								MarkdownDescription: "The name of the Secret resource",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the resource being referred to. Ignored if not cluster scoped",
								MarkdownDescription: "Namespace of the resource being referred to. Ignored if not cluster scoped",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "The secret type of the managed secret",
								MarkdownDescription: "The secret type of the managed secret",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Opaque", "kubernetes.io/tls", "kubernetes.io/service-account-token", "kubernetes.io/dockercfg", "kubernetes.io/dockerconfigjson", "kubernetes.io/basic-auth", "kubernetes.io/ssh-auth", "bootstrap.kubernetes.io/token"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"name_transformer": schema.StringAttribute{
						Description:         "The environment variable compatible secrets name transformer to apply",
						MarkdownDescription: "The environment variable compatible secrets name transformer to apply",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("upper-camel", "camel", "lower-snake", "tf-var", "dotnet-env", "lower-kebab"),
						},
					},

					"processors": schema.SingleNestedAttribute{
						Description:         "A list of processors to transform the data during ingestion",
						MarkdownDescription: "A list of processors to transform the data during ingestion",
						Attributes: map[string]schema.Attribute{
							"as_name": schema.StringAttribute{
								Description:         "The mapped name of the field in the managed secret, defaults to the original Doppler secret name for Opaque Kubernetes secrets. If omitted for other types, the value is not copied to the managed secret.",
								MarkdownDescription: "The mapped name of the field in the managed secret, defaults to the original Doppler secret name for Opaque Kubernetes secrets. If omitted for other types, the value is not copied to the managed secret.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "The type of process to be performed, either 'plain' or 'base64'",
								MarkdownDescription: "The type of process to be performed, either 'plain' or 'base64'",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("plain", "base64"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"project": schema.StringAttribute{
						Description:         "The Doppler project",
						MarkdownDescription: "The Doppler project",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resync_seconds": schema.Int64Attribute{
						Description:         "The number of seconds to wait between resyncs",
						MarkdownDescription: "The number of seconds to wait between resyncs",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"secrets": schema.ListAttribute{
						Description:         "A list of secrets to sync from the config",
						MarkdownDescription: "A list of secrets to sync from the config",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_secret": schema.SingleNestedAttribute{
						Description:         "The Kubernetes secret containing the Doppler service token",
						MarkdownDescription: "The Kubernetes secret containing the Doppler service token",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "The name of the Secret resource",
								MarkdownDescription: "The name of the Secret resource",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the resource being referred to. Ignored if not cluster scoped",
								MarkdownDescription: "Namespace of the resource being referred to. Ignored if not cluster scoped",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"verify_tls": schema.BoolAttribute{
						Description:         "Whether or not to verify TLS",
						MarkdownDescription: "Whether or not to verify TLS",
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

func (r *SecretsDopplerComDopplerSecretV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_secrets_doppler_com_doppler_secret_v1alpha1_manifest")

	var model SecretsDopplerComDopplerSecretV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("secrets.doppler.com/v1alpha1")
	model.Kind = pointer.String("DopplerSecret")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
