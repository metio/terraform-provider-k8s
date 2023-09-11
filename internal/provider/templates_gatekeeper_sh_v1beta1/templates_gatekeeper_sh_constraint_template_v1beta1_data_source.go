/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package templates_gatekeeper_sh_v1beta1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &TemplatesGatekeeperShConstraintTemplateV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &TemplatesGatekeeperShConstraintTemplateV1Beta1DataSource{}
)

func NewTemplatesGatekeeperShConstraintTemplateV1Beta1DataSource() datasource.DataSource {
	return &TemplatesGatekeeperShConstraintTemplateV1Beta1DataSource{}
}

type TemplatesGatekeeperShConstraintTemplateV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type TemplatesGatekeeperShConstraintTemplateV1Beta1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Crd *struct {
			Spec *struct {
				Names *struct {
					Kind       *string   `tfsdk:"kind" json:"kind,omitempty"`
					ShortNames *[]string `tfsdk:"short_names" json:"shortNames,omitempty"`
				} `tfsdk:"names" json:"names,omitempty"`
				Validation *struct {
					LegacySchema    *bool              `tfsdk:"legacy_schema" json:"legacySchema,omitempty"`
					OpenAPIV3Schema *map[string]string `tfsdk:"open_apiv3_schema" json:"openAPIV3Schema,omitempty"`
				} `tfsdk:"validation" json:"validation,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"crd" json:"crd,omitempty"`
		Targets *[]struct {
			Code *[]struct {
				Engine *string            `tfsdk:"engine" json:"engine,omitempty"`
				Source *map[string]string `tfsdk:"source" json:"source,omitempty"`
			} `tfsdk:"code" json:"code,omitempty"`
			Libs   *[]string `tfsdk:"libs" json:"libs,omitempty"`
			Rego   *string   `tfsdk:"rego" json:"rego,omitempty"`
			Target *string   `tfsdk:"target" json:"target,omitempty"`
		} `tfsdk:"targets" json:"targets,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TemplatesGatekeeperShConstraintTemplateV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_templates_gatekeeper_sh_constraint_template_v1beta1"
}

func (r *TemplatesGatekeeperShConstraintTemplateV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ConstraintTemplate is the Schema for the constrainttemplates API",
		MarkdownDescription: "ConstraintTemplate is the Schema for the constrainttemplates API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "ConstraintTemplateSpec defines the desired state of ConstraintTemplate.",
				MarkdownDescription: "ConstraintTemplateSpec defines the desired state of ConstraintTemplate.",
				Attributes: map[string]schema.Attribute{
					"crd": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"names": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"kind": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"short_names": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
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

									"validation": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"legacy_schema": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"open_apiv3_schema": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
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
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"targets": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"code": schema.ListNestedAttribute{
									Description:         "The source code options for the constraint template. 'Rego' can only be specified in one place (either here or in the 'rego' field)",
									MarkdownDescription: "The source code options for the constraint template. 'Rego' can only be specified in one place (either here or in the 'rego' field)",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"engine": schema.StringAttribute{
												Description:         "The engine used to evaluate the code. Example: 'Rego'. Required.",
												MarkdownDescription: "The engine used to evaluate the code. Example: 'Rego'. Required.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"source": schema.MapAttribute{
												Description:         "The source code for the template. Required.",
												MarkdownDescription: "The source code for the template. Required.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"libs": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"rego": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"target": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
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
		},
	}
}

func (r *TemplatesGatekeeperShConstraintTemplateV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *TemplatesGatekeeperShConstraintTemplateV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_templates_gatekeeper_sh_constraint_template_v1beta1")

	var data TemplatesGatekeeperShConstraintTemplateV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "templates.gatekeeper.sh", Version: "v1beta1", Resource: "constrainttemplates"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetResourceError(err, data.Metadata.Name))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse TemplatesGatekeeperShConstraintTemplateV1Beta1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("templates.gatekeeper.sh/v1beta1")
	data.Kind = pointer.String("ConstraintTemplate")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
