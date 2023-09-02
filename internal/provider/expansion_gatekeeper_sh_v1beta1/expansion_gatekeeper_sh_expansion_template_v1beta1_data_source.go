/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package expansion_gatekeeper_sh_v1beta1

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
	_ datasource.DataSource              = &ExpansionGatekeeperShExpansionTemplateV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &ExpansionGatekeeperShExpansionTemplateV1Beta1DataSource{}
)

func NewExpansionGatekeeperShExpansionTemplateV1Beta1DataSource() datasource.DataSource {
	return &ExpansionGatekeeperShExpansionTemplateV1Beta1DataSource{}
}

type ExpansionGatekeeperShExpansionTemplateV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ExpansionGatekeeperShExpansionTemplateV1Beta1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ApplyTo *[]struct {
			Groups   *[]string `tfsdk:"groups" json:"groups,omitempty"`
			Kinds    *[]string `tfsdk:"kinds" json:"kinds,omitempty"`
			Versions *[]string `tfsdk:"versions" json:"versions,omitempty"`
		} `tfsdk:"apply_to" json:"applyTo,omitempty"`
		EnforcementAction *string `tfsdk:"enforcement_action" json:"enforcementAction,omitempty"`
		GeneratedGVK      *struct {
			Group   *string `tfsdk:"group" json:"group,omitempty"`
			Kind    *string `tfsdk:"kind" json:"kind,omitempty"`
			Version *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"generated_gvk" json:"generatedGVK,omitempty"`
		TemplateSource *string `tfsdk:"template_source" json:"templateSource,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ExpansionGatekeeperShExpansionTemplateV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_expansion_gatekeeper_sh_expansion_template_v1beta1"
}

func (r *ExpansionGatekeeperShExpansionTemplateV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ExpansionTemplate is the Schema for the ExpansionTemplate API.",
		MarkdownDescription: "ExpansionTemplate is the Schema for the ExpansionTemplate API.",
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
				Description:         "ExpansionTemplateSpec defines the desired state of ExpansionTemplate.",
				MarkdownDescription: "ExpansionTemplateSpec defines the desired state of ExpansionTemplate.",
				Attributes: map[string]schema.Attribute{
					"apply_to": schema.ListNestedAttribute{
						Description:         "ApplyTo lists the specific groups, versions and kinds of generator resources which will be expanded.",
						MarkdownDescription: "ApplyTo lists the specific groups, versions and kinds of generator resources which will be expanded.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"groups": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"kinds": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"versions": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"enforcement_action": schema.StringAttribute{
						Description:         "EnforcementAction specifies the enforcement action to be used for resources matching the ExpansionTemplate. Specifying an empty value will use the enforcement action specified by the Constraint in violation.",
						MarkdownDescription: "EnforcementAction specifies the enforcement action to be used for resources matching the ExpansionTemplate. Specifying an empty value will use the enforcement action specified by the Constraint in violation.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"generated_gvk": schema.SingleNestedAttribute{
						Description:         "GeneratedGVK specifies the GVK of the resources which the generator resource creates.",
						MarkdownDescription: "GeneratedGVK specifies the GVK of the resources which the generator resource creates.",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"version": schema.StringAttribute{
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

					"template_source": schema.StringAttribute{
						Description:         "TemplateSource specifies the source field on the generator resource to use as the base for expanded resource. For Pod-creating generators, this is usually spec.template",
						MarkdownDescription: "TemplateSource specifies the source field on the generator resource to use as the base for expanded resource. For Pod-creating generators, this is usually spec.template",
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

func (r *ExpansionGatekeeperShExpansionTemplateV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ExpansionGatekeeperShExpansionTemplateV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_expansion_gatekeeper_sh_expansion_template_v1beta1")

	var data ExpansionGatekeeperShExpansionTemplateV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "expansion.gatekeeper.sh", Version: "v1beta1", Resource: "ExpansionTemplate"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
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

	var readResponse ExpansionGatekeeperShExpansionTemplateV1Beta1DataSourceData
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
	data.ApiVersion = pointer.String("expansion.gatekeeper.sh/v1beta1")
	data.Kind = pointer.String("ExpansionTemplate")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
