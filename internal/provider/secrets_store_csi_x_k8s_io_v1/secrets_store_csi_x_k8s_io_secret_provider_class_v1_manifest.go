/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package secrets_store_csi_x_k8s_io_v1

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
	_ datasource.DataSource = &SecretsStoreCsiXK8SIoSecretProviderClassV1Manifest{}
)

func NewSecretsStoreCsiXK8SIoSecretProviderClassV1Manifest() datasource.DataSource {
	return &SecretsStoreCsiXK8SIoSecretProviderClassV1Manifest{}
}

type SecretsStoreCsiXK8SIoSecretProviderClassV1Manifest struct{}

type SecretsStoreCsiXK8SIoSecretProviderClassV1ManifestData struct {
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
		Parameters    *map[string]string `tfsdk:"parameters" json:"parameters,omitempty"`
		Provider      *string            `tfsdk:"provider" json:"provider,omitempty"`
		SecretObjects *[]struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Data        *[]struct {
				Key        *string `tfsdk:"key" json:"key,omitempty"`
				ObjectName *string `tfsdk:"object_name" json:"objectName,omitempty"`
			} `tfsdk:"data" json:"data,omitempty"`
			Labels     *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			SecretName *string            `tfsdk:"secret_name" json:"secretName,omitempty"`
			Type       *string            `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"secret_objects" json:"secretObjects,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SecretsStoreCsiXK8SIoSecretProviderClassV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_secrets_store_csi_x_k8s_io_secret_provider_class_v1_manifest"
}

func (r *SecretsStoreCsiXK8SIoSecretProviderClassV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SecretProviderClass is the Schema for the secretproviderclasses API",
		MarkdownDescription: "SecretProviderClass is the Schema for the secretproviderclasses API",
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
				Description:         "SecretProviderClassSpec defines the desired state of SecretProviderClass",
				MarkdownDescription: "SecretProviderClassSpec defines the desired state of SecretProviderClass",
				Attributes: map[string]schema.Attribute{
					"parameters": schema.MapAttribute{
						Description:         "Configuration for specific provider",
						MarkdownDescription: "Configuration for specific provider",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"provider": schema.StringAttribute{
						Description:         "Configuration for provider name",
						MarkdownDescription: "Configuration for provider name",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"secret_objects": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"annotations": schema.MapAttribute{
									Description:         "annotations of k8s secret object",
									MarkdownDescription: "annotations of k8s secret object",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"data": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "data field to populate",
												MarkdownDescription: "data field to populate",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"object_name": schema.StringAttribute{
												Description:         "name of the object to sync",
												MarkdownDescription: "name of the object to sync",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"labels": schema.MapAttribute{
									Description:         "labels of K8s secret object",
									MarkdownDescription: "labels of K8s secret object",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"secret_name": schema.StringAttribute{
									Description:         "name of the K8s secret object",
									MarkdownDescription: "name of the K8s secret object",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"type": schema.StringAttribute{
									Description:         "type of K8s secret object",
									MarkdownDescription: "type of K8s secret object",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *SecretsStoreCsiXK8SIoSecretProviderClassV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_secrets_store_csi_x_k8s_io_secret_provider_class_v1_manifest")

	var model SecretsStoreCsiXK8SIoSecretProviderClassV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("secrets-store.csi.x-k8s.io/v1")
	model.Kind = pointer.String("SecretProviderClass")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
