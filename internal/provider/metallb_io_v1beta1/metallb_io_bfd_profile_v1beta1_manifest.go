/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package metallb_io_v1beta1

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
	_ datasource.DataSource = &MetallbIoBfdprofileV1Beta1Manifest{}
)

func NewMetallbIoBfdprofileV1Beta1Manifest() datasource.DataSource {
	return &MetallbIoBfdprofileV1Beta1Manifest{}
}

type MetallbIoBfdprofileV1Beta1Manifest struct{}

type MetallbIoBfdprofileV1Beta1ManifestData struct {
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
		DetectMultiplier *int64 `tfsdk:"detect_multiplier" json:"detectMultiplier,omitempty"`
		EchoInterval     *int64 `tfsdk:"echo_interval" json:"echoInterval,omitempty"`
		EchoMode         *bool  `tfsdk:"echo_mode" json:"echoMode,omitempty"`
		MinimumTtl       *int64 `tfsdk:"minimum_ttl" json:"minimumTtl,omitempty"`
		PassiveMode      *bool  `tfsdk:"passive_mode" json:"passiveMode,omitempty"`
		ReceiveInterval  *int64 `tfsdk:"receive_interval" json:"receiveInterval,omitempty"`
		TransmitInterval *int64 `tfsdk:"transmit_interval" json:"transmitInterval,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MetallbIoBfdprofileV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_metallb_io_bfd_profile_v1beta1_manifest"
}

func (r *MetallbIoBfdprofileV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "BFDProfile represents the settings of the bfd session that can be optionally associated with a BGP session.",
		MarkdownDescription: "BFDProfile represents the settings of the bfd session that can be optionally associated with a BGP session.",
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
				Description:         "BFDProfileSpec defines the desired state of BFDProfile.",
				MarkdownDescription: "BFDProfileSpec defines the desired state of BFDProfile.",
				Attributes: map[string]schema.Attribute{
					"detect_multiplier": schema.Int64Attribute{
						Description:         "Configures the detection multiplier to determine packet loss. The remote transmission interval will be multiplied by this value to determine the connection loss detection timer.",
						MarkdownDescription: "Configures the detection multiplier to determine packet loss. The remote transmission interval will be multiplied by this value to determine the connection loss detection timer.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(2),
							int64validator.AtMost(255),
						},
					},

					"echo_interval": schema.Int64Attribute{
						Description:         "Configures the minimal echo receive transmission interval that this system is capable of handling in milliseconds. Defaults to 50ms",
						MarkdownDescription: "Configures the minimal echo receive transmission interval that this system is capable of handling in milliseconds. Defaults to 50ms",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(10),
							int64validator.AtMost(60000),
						},
					},

					"echo_mode": schema.BoolAttribute{
						Description:         "Enables or disables the echo transmission mode. This mode is disabled by default, and not supported on multi hops setups.",
						MarkdownDescription: "Enables or disables the echo transmission mode. This mode is disabled by default, and not supported on multi hops setups.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"minimum_ttl": schema.Int64Attribute{
						Description:         "For multi hop sessions only: configure the minimum expected TTL for an incoming BFD control packet.",
						MarkdownDescription: "For multi hop sessions only: configure the minimum expected TTL for an incoming BFD control packet.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
							int64validator.AtMost(254),
						},
					},

					"passive_mode": schema.BoolAttribute{
						Description:         "Mark session as passive: a passive session will not attempt to start the connection and will wait for control packets from peer before it begins replying.",
						MarkdownDescription: "Mark session as passive: a passive session will not attempt to start the connection and will wait for control packets from peer before it begins replying.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"receive_interval": schema.Int64Attribute{
						Description:         "The minimum interval that this system is capable of receiving control packets in milliseconds. Defaults to 300ms.",
						MarkdownDescription: "The minimum interval that this system is capable of receiving control packets in milliseconds. Defaults to 300ms.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(10),
							int64validator.AtMost(60000),
						},
					},

					"transmit_interval": schema.Int64Attribute{
						Description:         "The minimum transmission interval (less jitter) that this system wants to use to send BFD control packets in milliseconds. Defaults to 300ms",
						MarkdownDescription: "The minimum transmission interval (less jitter) that this system wants to use to send BFD control packets in milliseconds. Defaults to 300ms",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(10),
							int64validator.AtMost(60000),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *MetallbIoBfdprofileV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_metallb_io_bfd_profile_v1beta1_manifest")

	var model MetallbIoBfdprofileV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("metallb.io/v1beta1")
	model.Kind = pointer.String("BFDProfile")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
