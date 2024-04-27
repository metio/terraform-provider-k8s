/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package devices_kubeedge_io_v1beta1

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
	_ datasource.DataSource = &DevicesKubeedgeIoDeviceModelV1Beta1Manifest{}
)

func NewDevicesKubeedgeIoDeviceModelV1Beta1Manifest() datasource.DataSource {
	return &DevicesKubeedgeIoDeviceModelV1Beta1Manifest{}
}

type DevicesKubeedgeIoDeviceModelV1Beta1Manifest struct{}

type DevicesKubeedgeIoDeviceModelV1Beta1ManifestData struct {
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
		Properties *[]struct {
			AccessMode  *string `tfsdk:"access_mode" json:"accessMode,omitempty"`
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Maximum     *string `tfsdk:"maximum" json:"maximum,omitempty"`
			Minimum     *string `tfsdk:"minimum" json:"minimum,omitempty"`
			Name        *string `tfsdk:"name" json:"name,omitempty"`
			Type        *string `tfsdk:"type" json:"type,omitempty"`
			Unit        *string `tfsdk:"unit" json:"unit,omitempty"`
		} `tfsdk:"properties" json:"properties,omitempty"`
		Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DevicesKubeedgeIoDeviceModelV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_devices_kubeedge_io_device_model_v1beta1_manifest"
}

func (r *DevicesKubeedgeIoDeviceModelV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DeviceModel is the Schema for the device model API",
		MarkdownDescription: "DeviceModel is the Schema for the device model API",
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
				Description:         "DeviceModelSpec defines the model for a device.It is a blueprint which describes the device capabilities and access mechanism via property visitors.",
				MarkdownDescription: "DeviceModelSpec defines the model for a device.It is a blueprint which describes the device capabilities and access mechanism via property visitors.",
				Attributes: map[string]schema.Attribute{
					"properties": schema.ListNestedAttribute{
						Description:         "Required: List of device properties.",
						MarkdownDescription: "Required: List of device properties.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"access_mode": schema.StringAttribute{
									Description:         "Required: Access mode of property, ReadWrite or ReadOnly.",
									MarkdownDescription: "Required: Access mode of property, ReadWrite or ReadOnly.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("ReadWrite", "ReadOnly"),
									},
								},

								"description": schema.StringAttribute{
									Description:         "The device property description.",
									MarkdownDescription: "The device property description.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"maximum": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"minimum": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Required: The device property name. Note: If you need to use the built-in stream data processing function, you need to define Name as saveFrame or saveVideo",
									MarkdownDescription: "Required: The device property name. Note: If you need to use the built-in stream data processing function, you need to define Name as saveFrame or saveVideo",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"type": schema.StringAttribute{
									Description:         "Required: Type of device property, ENUM: INT,FLOAT,DOUBLE,STRING,BOOLEAN,BYTES,STREAM",
									MarkdownDescription: "Required: Type of device property, ENUM: INT,FLOAT,DOUBLE,STRING,BOOLEAN,BYTES,STREAM",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("INT", "FLOAT", "DOUBLE", "STRING", "BOOLEAN", "BYTES", "STREAM"),
									},
								},

								"unit": schema.StringAttribute{
									Description:         "The unit of the property",
									MarkdownDescription: "The unit of the property",
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

					"protocol": schema.StringAttribute{
						Description:         "Required: Protocol name used by the device.",
						MarkdownDescription: "Required: Protocol name used by the device.",
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

func (r *DevicesKubeedgeIoDeviceModelV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_devices_kubeedge_io_device_model_v1beta1_manifest")

	var model DevicesKubeedgeIoDeviceModelV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("devices.kubeedge.io/v1beta1")
	model.Kind = pointer.String("DeviceModel")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
