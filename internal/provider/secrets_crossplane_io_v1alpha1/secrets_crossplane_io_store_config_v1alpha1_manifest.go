/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package secrets_crossplane_io_v1alpha1

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
	_ datasource.DataSource = &SecretsCrossplaneIoStoreConfigV1Alpha1Manifest{}
)

func NewSecretsCrossplaneIoStoreConfigV1Alpha1Manifest() datasource.DataSource {
	return &SecretsCrossplaneIoStoreConfigV1Alpha1Manifest{}
}

type SecretsCrossplaneIoStoreConfigV1Alpha1Manifest struct{}

type SecretsCrossplaneIoStoreConfigV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		DefaultScope *string `tfsdk:"default_scope" json:"defaultScope,omitempty"`
		Kubernetes   *struct {
			Auth *struct {
				Env *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"env" json:"env,omitempty"`
				Fs *struct {
					Path *string `tfsdk:"path" json:"path,omitempty"`
				} `tfsdk:"fs" json:"fs,omitempty"`
				SecretRef *struct {
					Key       *string `tfsdk:"key" json:"key,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				Source *string `tfsdk:"source" json:"source,omitempty"`
			} `tfsdk:"auth" json:"auth,omitempty"`
		} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
		Plugin *struct {
			ConfigRef *struct {
				ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
				Name       *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"config_ref" json:"configRef,omitempty"`
			Endpoint *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
		} `tfsdk:"plugin" json:"plugin,omitempty"`
		Type *string `tfsdk:"type" json:"type,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SecretsCrossplaneIoStoreConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_secrets_crossplane_io_store_config_v1alpha1_manifest"
}

func (r *SecretsCrossplaneIoStoreConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "A StoreConfig configures how Crossplane controllers should store connection details in an external secret store.",
		MarkdownDescription: "A StoreConfig configures how Crossplane controllers should store connection details in an external secret store.",
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
				Description:         "A StoreConfigSpec defines the desired state of a StoreConfig.",
				MarkdownDescription: "A StoreConfigSpec defines the desired state of a StoreConfig.",
				Attributes: map[string]schema.Attribute{
					"default_scope": schema.StringAttribute{
						Description:         "DefaultScope used for scoping secrets for 'cluster-scoped' resources. If store type is 'Kubernetes', this would mean the default namespace to store connection secrets for cluster scoped resources. In case of 'Vault', this would be used as the default parent path. Typically, should be set as Crossplane installation namespace.",
						MarkdownDescription: "DefaultScope used for scoping secrets for 'cluster-scoped' resources. If store type is 'Kubernetes', this would mean the default namespace to store connection secrets for cluster scoped resources. In case of 'Vault', this would be used as the default parent path. Typically, should be set as Crossplane installation namespace.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"kubernetes": schema.SingleNestedAttribute{
						Description:         "Kubernetes configures a Kubernetes secret store. If the 'type' is 'Kubernetes' but no config provided, in cluster config will be used.",
						MarkdownDescription: "Kubernetes configures a Kubernetes secret store. If the 'type' is 'Kubernetes' but no config provided, in cluster config will be used.",
						Attributes: map[string]schema.Attribute{
							"auth": schema.SingleNestedAttribute{
								Description:         "Credentials used to connect to the Kubernetes API.",
								MarkdownDescription: "Credentials used to connect to the Kubernetes API.",
								Attributes: map[string]schema.Attribute{
									"env": schema.SingleNestedAttribute{
										Description:         "Env is a reference to an environment variable that contains credentials that must be used to connect to the provider.",
										MarkdownDescription: "Env is a reference to an environment variable that contains credentials that must be used to connect to the provider.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name is the name of an environment variable.",
												MarkdownDescription: "Name is the name of an environment variable.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"fs": schema.SingleNestedAttribute{
										Description:         "Fs is a reference to a filesystem location that contains credentials that must be used to connect to the provider.",
										MarkdownDescription: "Fs is a reference to a filesystem location that contains credentials that must be used to connect to the provider.",
										Attributes: map[string]schema.Attribute{
											"path": schema.StringAttribute{
												Description:         "Path is a filesystem path.",
												MarkdownDescription: "Path is a filesystem path.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_ref": schema.SingleNestedAttribute{
										Description:         "A SecretRef is a reference to a secret key that contains the credentials that must be used to connect to the provider.",
										MarkdownDescription: "A SecretRef is a reference to a secret key that contains the credentials that must be used to connect to the provider.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the secret.",
												MarkdownDescription: "Name of the secret.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the secret.",
												MarkdownDescription: "Namespace of the secret.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"source": schema.StringAttribute{
										Description:         "Source of the credentials.",
										MarkdownDescription: "Source of the credentials.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("None", "Secret", "Environment", "Filesystem"),
										},
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

					"plugin": schema.SingleNestedAttribute{
						Description:         "Plugin configures External secret store as a plugin.",
						MarkdownDescription: "Plugin configures External secret store as a plugin.",
						Attributes: map[string]schema.Attribute{
							"config_ref": schema.SingleNestedAttribute{
								Description:         "ConfigRef contains store config reference info.",
								MarkdownDescription: "ConfigRef contains store config reference info.",
								Attributes: map[string]schema.Attribute{
									"api_version": schema.StringAttribute{
										Description:         "APIVersion of the referenced config.",
										MarkdownDescription: "APIVersion of the referenced config.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"kind": schema.StringAttribute{
										Description:         "Kind of the referenced config.",
										MarkdownDescription: "Kind of the referenced config.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referenced config.",
										MarkdownDescription: "Name of the referenced config.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"endpoint": schema.StringAttribute{
								Description:         "Endpoint is the endpoint of the gRPC server.",
								MarkdownDescription: "Endpoint is the endpoint of the gRPC server.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"type": schema.StringAttribute{
						Description:         "Type configures which secret store to be used. Only the configuration block for this store will be used and others will be ignored if provided. Default is Kubernetes.",
						MarkdownDescription: "Type configures which secret store to be used. Only the configuration block for this store will be used and others will be ignored if provided. Default is Kubernetes.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Kubernetes", "Vault", "Plugin"),
						},
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *SecretsCrossplaneIoStoreConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_secrets_crossplane_io_store_config_v1alpha1_manifest")

	var model SecretsCrossplaneIoStoreConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("secrets.crossplane.io/v1alpha1")
	model.Kind = pointer.String("StoreConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
