/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package digitalis_io_v1

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
	_ datasource.DataSource = &DigitalisIoValsSecretV1Manifest{}
)

func NewDigitalisIoValsSecretV1Manifest() datasource.DataSource {
	return &DigitalisIoValsSecretV1Manifest{}
}

type DigitalisIoValsSecretV1Manifest struct{}

type DigitalisIoValsSecretV1ManifestData struct {
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
		Data *struct {
			Encoding *string `tfsdk:"encoding" json:"encoding,omitempty"`
			Ref      *string `tfsdk:"ref" json:"ref,omitempty"`
		} `tfsdk:"data" json:"data,omitempty"`
		Databases *[]struct {
			Driver           *string   `tfsdk:"driver" json:"driver,omitempty"`
			Hosts            *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
			LoginCredentials *struct {
				Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
				PasswordKey *string `tfsdk:"password_key" json:"passwordKey,omitempty"`
				SecretName  *string `tfsdk:"secret_name" json:"secretName,omitempty"`
				UsernameKey *string `tfsdk:"username_key" json:"usernameKey,omitempty"`
			} `tfsdk:"login_credentials" json:"loginCredentials,omitempty"`
			PasswordKey *string `tfsdk:"password_key" json:"passwordKey,omitempty"`
			Port        *int64  `tfsdk:"port" json:"port,omitempty"`
			UserHost    *string `tfsdk:"user_host" json:"userHost,omitempty"`
			UsernameKey *string `tfsdk:"username_key" json:"usernameKey,omitempty"`
		} `tfsdk:"databases" json:"databases,omitempty"`
		Name     *string            `tfsdk:"name" json:"name,omitempty"`
		Template *map[string]string `tfsdk:"template" json:"template,omitempty"`
		Ttl      *int64             `tfsdk:"ttl" json:"ttl,omitempty"`
		Type     *string            `tfsdk:"type" json:"type,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DigitalisIoValsSecretV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_digitalis_io_vals_secret_v1_manifest"
}

func (r *DigitalisIoValsSecretV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ValsSecret is the Schema for the valssecrets API",
		MarkdownDescription: "ValsSecret is the Schema for the valssecrets API",
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
				Description:         "ValsSecretSpec defines the desired state of ValsSecret",
				MarkdownDescription: "ValsSecretSpec defines the desired state of ValsSecret",
				Attributes: map[string]schema.Attribute{
					"data": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"encoding": schema.StringAttribute{
								Description:         "Encoding type for the secret. Only base64 supported. Optional",
								MarkdownDescription: "Encoding type for the secret. Only base64 supported. Optional",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ref": schema.StringAttribute{
								Description:         "Ref value to the secret in the format ref+backend://path https://github.com/helmfile/vals",
								MarkdownDescription: "Ref value to the secret in the format ref+backend://path https://github.com/helmfile/vals",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"databases": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"driver": schema.StringAttribute{
									Description:         "Defines the database type",
									MarkdownDescription: "Defines the database type",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"hosts": schema.ListAttribute{
									Description:         "List of hosts to connect to, they'll be tried in sequence until one succeeds",
									MarkdownDescription: "List of hosts to connect to, they'll be tried in sequence until one succeeds",
									ElementType:         types.StringType,
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"login_credentials": schema.SingleNestedAttribute{
									Description:         "Credentials to access the database",
									MarkdownDescription: "Credentials to access the database",
									Attributes: map[string]schema.Attribute{
										"namespace": schema.StringAttribute{
											Description:         "Optional namespace of the secret, default current namespace",
											MarkdownDescription: "Optional namespace of the secret, default current namespace",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"password_key": schema.StringAttribute{
											Description:         "Key in the secret containing the database username",
											MarkdownDescription: "Key in the secret containing the database username",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"secret_name": schema.StringAttribute{
											Description:         "Name of the secret containing the credentials to be able to log in to the database",
											MarkdownDescription: "Name of the secret containing the credentials to be able to log in to the database",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"username_key": schema.StringAttribute{
											Description:         "Key in the secret containing the database username",
											MarkdownDescription: "Key in the secret containing the database username",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"password_key": schema.StringAttribute{
									Description:         "Key in the secret containing the database username",
									MarkdownDescription: "Key in the secret containing the database username",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"port": schema.Int64Attribute{
									Description:         "Database port number",
									MarkdownDescription: "Database port number",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"user_host": schema.StringAttribute{
									Description:         "Used for MySQL only, the host part for the username",
									MarkdownDescription: "Used for MySQL only, the host part for the username",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"username_key": schema.StringAttribute{
									Description:         "Key in the secret containing the database username",
									MarkdownDescription: "Key in the secret containing the database username",
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

					"name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"template": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ttl": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"type": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
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

func (r *DigitalisIoValsSecretV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_digitalis_io_vals_secret_v1_manifest")

	var model DigitalisIoValsSecretV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("digitalis.io/v1")
	model.Kind = pointer.String("ValsSecret")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
