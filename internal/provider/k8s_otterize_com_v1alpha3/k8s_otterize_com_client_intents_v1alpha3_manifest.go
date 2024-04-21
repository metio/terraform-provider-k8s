/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package k8s_otterize_com_v1alpha3

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
	_ datasource.DataSource = &K8SOtterizeComClientIntentsV1Alpha3Manifest{}
)

func NewK8SOtterizeComClientIntentsV1Alpha3Manifest() datasource.DataSource {
	return &K8SOtterizeComClientIntentsV1Alpha3Manifest{}
}

type K8SOtterizeComClientIntentsV1Alpha3Manifest struct{}

type K8SOtterizeComClientIntentsV1Alpha3ManifestData struct {
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
		Calls *[]struct {
			HTTPResources *[]struct {
				Methods *[]string `tfsdk:"methods" json:"methods,omitempty"`
				Path    *string   `tfsdk:"path" json:"path,omitempty"`
			} `tfsdk:"http_resources" json:"HTTPResources,omitempty"`
			AwsActions          *[]string `tfsdk:"aws_actions" json:"awsActions,omitempty"`
			AzureKeyVaultPolicy *struct {
				CertificatePermissions *[]string `tfsdk:"certificate_permissions" json:"certificatePermissions,omitempty"`
				KeyPermissions         *[]string `tfsdk:"key_permissions" json:"keyPermissions,omitempty"`
				SecretPermissions      *[]string `tfsdk:"secret_permissions" json:"secretPermissions,omitempty"`
				StoragePermissions     *[]string `tfsdk:"storage_permissions" json:"storagePermissions,omitempty"`
			} `tfsdk:"azure_key_vault_policy" json:"azureKeyVaultPolicy,omitempty"`
			AzureRoles        *[]string `tfsdk:"azure_roles" json:"azureRoles,omitempty"`
			DatabaseResources *[]struct {
				DatabaseName *string   `tfsdk:"database_name" json:"databaseName,omitempty"`
				Operations   *[]string `tfsdk:"operations" json:"operations,omitempty"`
				Table        *string   `tfsdk:"table" json:"table,omitempty"`
			} `tfsdk:"database_resources" json:"databaseResources,omitempty"`
			GcpPermissions *[]string `tfsdk:"gcp_permissions" json:"gcpPermissions,omitempty"`
			Internet       *struct {
				Domains *[]string `tfsdk:"domains" json:"domains,omitempty"`
				Ips     *[]string `tfsdk:"ips" json:"ips,omitempty"`
				Ports   *[]string `tfsdk:"ports" json:"ports,omitempty"`
			} `tfsdk:"internet" json:"internet,omitempty"`
			KafkaTopics *[]struct {
				Name       *string   `tfsdk:"name" json:"name,omitempty"`
				Operations *[]string `tfsdk:"operations" json:"operations,omitempty"`
			} `tfsdk:"kafka_topics" json:"kafkaTopics,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"calls" json:"calls,omitempty"`
		Service *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"service" json:"service,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *K8SOtterizeComClientIntentsV1Alpha3Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_k8s_otterize_com_client_intents_v1alpha3_manifest"
}

func (r *K8SOtterizeComClientIntentsV1Alpha3Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClientIntents is the Schema for the intents API",
		MarkdownDescription: "ClientIntents is the Schema for the intents API",
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
				Description:         "IntentsSpec defines the desired state of ClientIntents",
				MarkdownDescription: "IntentsSpec defines the desired state of ClientIntents",
				Attributes: map[string]schema.Attribute{
					"calls": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"http_resources": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"methods": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"path": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"aws_actions": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"azure_key_vault_policy": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"certificate_permissions": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key_permissions": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_permissions": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"storage_permissions": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
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

								"azure_roles": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"database_resources": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"database_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"operations": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"table": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
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

								"gcp_permissions": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"internet": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"domains": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ips": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ports": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
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

								"kafka_topics": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"operations": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
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

								"type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("http", "kafka", "database", "aws", "gcp", "azure", "internet"),
									},
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"service": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
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
		},
	}
}

func (r *K8SOtterizeComClientIntentsV1Alpha3Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_k8s_otterize_com_client_intents_v1alpha3_manifest")

	var model K8SOtterizeComClientIntentsV1Alpha3ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("k8s.otterize.com/v1alpha3")
	model.Kind = pointer.String("ClientIntents")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
