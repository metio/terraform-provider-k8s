/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package akri_sh_v0

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &AkriShInstanceV0Manifest{}
)

func NewAkriShInstanceV0Manifest() datasource.DataSource {
	return &AkriShInstanceV0Manifest{}
}

type AkriShInstanceV0Manifest struct{}

type AkriShInstanceV0ManifestData struct {
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
		BrokerProperties  *map[string]string `tfsdk:"broker_properties" json:"brokerProperties,omitempty"`
		Capacity          *int64             `tfsdk:"capacity" json:"capacity,omitempty"`
		CdiName           *string            `tfsdk:"cdi_name" json:"cdiName,omitempty"`
		ConfigurationName *string            `tfsdk:"configuration_name" json:"configurationName,omitempty"`
		DeviceUsage       *map[string]string `tfsdk:"device_usage" json:"deviceUsage,omitempty"`
		Nodes             *[]string          `tfsdk:"nodes" json:"nodes,omitempty"`
		Shared            *bool              `tfsdk:"shared" json:"shared,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AkriShInstanceV0Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_akri_sh_instance_v0_manifest"
}

func (r *AkriShInstanceV0Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Auto-generated derived type for InstanceSpec via 'CustomResource'",
		MarkdownDescription: "Auto-generated derived type for InstanceSpec via 'CustomResource'",
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
				Description:         "Defines the information in the Instance CRD An Instance is a specific instance described by a Configuration. For example, a Configuration may describe many cameras, each camera will be represented by a Instance.",
				MarkdownDescription: "Defines the information in the Instance CRD An Instance is a specific instance described by a Configuration. For example, a Configuration may describe many cameras, each camera will be represented by a Instance.",
				Attributes: map[string]schema.Attribute{
					"broker_properties": schema.MapAttribute{
						Description:         "This defines some properties that will be set as environment variables in broker Pods that request the resource this Instance represents. It contains the 'Configuration.broker_properties' from this Instance's Configuration and the 'Device.properties' set by the Discovery Handler that discovered the resource this Instance represents.",
						MarkdownDescription: "This defines some properties that will be set as environment variables in broker Pods that request the resource this Instance represents. It contains the 'Configuration.broker_properties' from this Instance's Configuration and the 'Device.properties' set by the Discovery Handler that discovered the resource this Instance represents.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"capacity": schema.Int64Attribute{
						Description:         "This contains the number of slots for the Instance",
						MarkdownDescription: "This contains the number of slots for the Instance",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"cdi_name": schema.StringAttribute{
						Description:         "This contains the CDI fully qualified name of the device linked to the Instance",
						MarkdownDescription: "This contains the CDI fully qualified name of the device linked to the Instance",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"configuration_name": schema.StringAttribute{
						Description:         "This contains the name of the corresponding Configuration",
						MarkdownDescription: "This contains the name of the corresponding Configuration",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"device_usage": schema.MapAttribute{
						Description:         "This contains a map of capability slots to node names. The number of slots corresponds to the associated Configuration.capacity field. Each slot will either map to an empty string (if the slot has not been claimed) or to a node name (corresponding to the node that has claimed the slot)",
						MarkdownDescription: "This contains a map of capability slots to node names. The number of slots corresponds to the associated Configuration.capacity field. Each slot will either map to an empty string (if the slot has not been claimed) or to a node name (corresponding to the node that has claimed the slot)",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"nodes": schema.ListAttribute{
						Description:         "This contains a list of the nodes that can access this capability instance",
						MarkdownDescription: "This contains a list of the nodes that can access this capability instance",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"shared": schema.BoolAttribute{
						Description:         "This defines whether the capability is to be shared by multiple nodes",
						MarkdownDescription: "This defines whether the capability is to be shared by multiple nodes",
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
	}
}

func (r *AkriShInstanceV0Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_akri_sh_instance_v0_manifest")

	var model AkriShInstanceV0ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("akri.sh/v0")
	model.Kind = pointer.String("Instance")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
