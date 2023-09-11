/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package servicebinding_io_v1beta1

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
	_ datasource.DataSource = &ServicebindingIoServiceBindingV1Beta1Manifest{}
)

func NewServicebindingIoServiceBindingV1Beta1Manifest() datasource.DataSource {
	return &ServicebindingIoServiceBindingV1Beta1Manifest{}
}

type ServicebindingIoServiceBindingV1Beta1Manifest struct{}

type ServicebindingIoServiceBindingV1Beta1ManifestData struct {
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
		Env *[]struct {
			Key  *string `tfsdk:"key" json:"key,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"env" json:"env,omitempty"`
		Name     *string `tfsdk:"name" json:"name,omitempty"`
		Provider *string `tfsdk:"provider" json:"provider,omitempty"`
		Service  *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"service" json:"service,omitempty"`
		Type     *string `tfsdk:"type" json:"type,omitempty"`
		Workload *struct {
			ApiVersion *string   `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Containers *[]string `tfsdk:"containers" json:"containers,omitempty"`
			Kind       *string   `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string   `tfsdk:"name" json:"name,omitempty"`
			Selector   *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"selector" json:"selector,omitempty"`
		} `tfsdk:"workload" json:"workload,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ServicebindingIoServiceBindingV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_servicebinding_io_service_binding_v1beta1_manifest"
}

func (r *ServicebindingIoServiceBindingV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ServiceBinding is the Schema for the servicebindings API",
		MarkdownDescription: "ServiceBinding is the Schema for the servicebindings API",
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
				Description:         "ServiceBindingSpec defines the desired state of ServiceBinding",
				MarkdownDescription: "ServiceBindingSpec defines the desired state of ServiceBinding",
				Attributes: map[string]schema.Attribute{
					"env": schema.ListNestedAttribute{
						Description:         "Env is the collection of mappings from Secret entries to environment variables",
						MarkdownDescription: "Env is the collection of mappings from Secret entries to environment variables",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "Key is the key in the Secret that will be exposed",
									MarkdownDescription: "Key is the key in the Secret that will be exposed",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is the name of the environment variable",
									MarkdownDescription: "Name is the name of the environment variable",
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
						Description:         "Name is the name of the service as projected into the workload container.  Defaults to .metadata.name.",
						MarkdownDescription: "Name is the name of the service as projected into the workload container.  Defaults to .metadata.name.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"provider": schema.StringAttribute{
						Description:         "Provider is the provider of the service as projected into the workload container",
						MarkdownDescription: "Provider is the provider of the service as projected into the workload container",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service": schema.SingleNestedAttribute{
						Description:         "Service is a reference to an object that fulfills the ProvisionedService duck type",
						MarkdownDescription: "Service is a reference to an object that fulfills the ProvisionedService duck type",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"type": schema.StringAttribute{
						Description:         "Type is the type of the service as projected into the workload container",
						MarkdownDescription: "Type is the type of the service as projected into the workload container",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"workload": schema.SingleNestedAttribute{
						Description:         "Workload is a reference to an object",
						MarkdownDescription: "Workload is a reference to an object",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"containers": schema.ListAttribute{
								Description:         "Containers describes which containers in a Pod should be bound to",
								MarkdownDescription: "Containers describes which containers in a Pod should be bound to",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"selector": schema.SingleNestedAttribute{
								Description:         "Selector is a query that selects the workload or workloads to bind the service to",
								MarkdownDescription: "Selector is a query that selects the workload or workloads to bind the service to",
								Attributes: map[string]schema.Attribute{
									"match_expressions": schema.ListNestedAttribute{
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "key is the label key that the selector applies to.",
													MarkdownDescription: "key is the label key that the selector applies to.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"values": schema.ListAttribute{
													Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													ElementType:         types.StringType,
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

									"match_labels": schema.MapAttribute{
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

func (r *ServicebindingIoServiceBindingV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_servicebinding_io_service_binding_v1beta1_manifest")

	var model ServicebindingIoServiceBindingV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("servicebinding.io/v1beta1")
	model.Kind = pointer.String("ServiceBinding")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
