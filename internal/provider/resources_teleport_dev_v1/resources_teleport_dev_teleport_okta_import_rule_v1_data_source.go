/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package resources_teleport_dev_v1

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
	_ datasource.DataSource              = &ResourcesTeleportDevTeleportOktaImportRuleV1DataSource{}
	_ datasource.DataSourceWithConfigure = &ResourcesTeleportDevTeleportOktaImportRuleV1DataSource{}
)

func NewResourcesTeleportDevTeleportOktaImportRuleV1DataSource() datasource.DataSource {
	return &ResourcesTeleportDevTeleportOktaImportRuleV1DataSource{}
}

type ResourcesTeleportDevTeleportOktaImportRuleV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ResourcesTeleportDevTeleportOktaImportRuleV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Mappings *[]struct {
			Add_labels *struct {
				Key   *string `tfsdk:"key" json:"key,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"add_labels" json:"add_labels,omitempty"`
			Match *[]struct {
				App_ids            *[]string `tfsdk:"app_ids" json:"app_ids,omitempty"`
				App_name_regexes   *[]string `tfsdk:"app_name_regexes" json:"app_name_regexes,omitempty"`
				Group_ids          *[]string `tfsdk:"group_ids" json:"group_ids,omitempty"`
				Group_name_regexes *[]string `tfsdk:"group_name_regexes" json:"group_name_regexes,omitempty"`
			} `tfsdk:"match" json:"match,omitempty"`
		} `tfsdk:"mappings" json:"mappings,omitempty"`
		Priority *int64 `tfsdk:"priority" json:"priority,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ResourcesTeleportDevTeleportOktaImportRuleV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_resources_teleport_dev_teleport_okta_import_rule_v1"
}

func (r *ResourcesTeleportDevTeleportOktaImportRuleV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "OktaImportRule is the Schema for the oktaimportrules API",
		MarkdownDescription: "OktaImportRule is the Schema for the oktaimportrules API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
				Description:         "OktaImportRule resource definition v1 from Teleport",
				MarkdownDescription: "OktaImportRule resource definition v1 from Teleport",
				Attributes: map[string]schema.Attribute{
					"mappings": schema.ListNestedAttribute{
						Description:         "Mappings is a list of matches that will map match conditions to labels.",
						MarkdownDescription: "Mappings is a list of matches that will map match conditions to labels.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"add_labels": schema.SingleNestedAttribute{
									Description:         "AddLabels specifies which labels to add if any of the previous matches match.",
									MarkdownDescription: "AddLabels specifies which labels to add if any of the previous matches match.",
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"value": schema.StringAttribute{
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

								"match": schema.ListNestedAttribute{
									Description:         "Match is a set of matching rules for this mapping. If any of these match, then the mapping will be applied.",
									MarkdownDescription: "Match is a set of matching rules for this mapping. If any of these match, then the mapping will be applied.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"app_ids": schema.ListAttribute{
												Description:         "AppIDs is a list of app IDs to match against.",
												MarkdownDescription: "AppIDs is a list of app IDs to match against.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"app_name_regexes": schema.ListAttribute{
												Description:         "AppNameRegexes is a list of regexes to match against app names.",
												MarkdownDescription: "AppNameRegexes is a list of regexes to match against app names.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"group_ids": schema.ListAttribute{
												Description:         "GroupIDs is a list of group IDs to match against.",
												MarkdownDescription: "GroupIDs is a list of group IDs to match against.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"group_name_regexes": schema.ListAttribute{
												Description:         "GroupNameRegexes is a list of regexes to match against group names.",
												MarkdownDescription: "GroupNameRegexes is a list of regexes to match against group names.",
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
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"priority": schema.Int64Attribute{
						Description:         "Priority represents the priority of the rule application. Lower numbered rules will be applied first.",
						MarkdownDescription: "Priority represents the priority of the rule application. Lower numbered rules will be applied first.",
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

func (r *ResourcesTeleportDevTeleportOktaImportRuleV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ResourcesTeleportDevTeleportOktaImportRuleV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_resources_teleport_dev_teleport_okta_import_rule_v1")

	var data ResourcesTeleportDevTeleportOktaImportRuleV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "resources.teleport.dev", Version: "v1", Resource: "TeleportOktaImportRule"}).
		Namespace(data.Metadata.Namespace).
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

	var readResponse ResourcesTeleportDevTeleportOktaImportRuleV1DataSourceData
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

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("resources.teleport.dev/v1")
	data.Kind = pointer.String("TeleportOktaImportRule")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}