/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package devices_kubeedge_io_v1alpha2

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
	_ datasource.DataSource              = &DevicesKubeedgeIoDeviceModelV1Alpha2DataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesKubeedgeIoDeviceModelV1Alpha2DataSource{}
)

func NewDevicesKubeedgeIoDeviceModelV1Alpha2DataSource() datasource.DataSource {
	return &DevicesKubeedgeIoDeviceModelV1Alpha2DataSource{}
}

type DevicesKubeedgeIoDeviceModelV1Alpha2DataSource struct {
	kubernetesClient dynamic.Interface
}

type DevicesKubeedgeIoDeviceModelV1Alpha2DataSourceData struct {
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
		Properties *[]struct {
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Name        *string `tfsdk:"name" json:"name,omitempty"`
			Type        *struct {
				Boolean *struct {
					AccessMode   *string `tfsdk:"access_mode" json:"accessMode,omitempty"`
					DefaultValue *bool   `tfsdk:"default_value" json:"defaultValue,omitempty"`
				} `tfsdk:"boolean" json:"boolean,omitempty"`
				Bytes *struct {
					AccessMode *string `tfsdk:"access_mode" json:"accessMode,omitempty"`
				} `tfsdk:"bytes" json:"bytes,omitempty"`
				Double *struct {
					AccessMode   *string  `tfsdk:"access_mode" json:"accessMode,omitempty"`
					DefaultValue *float64 `tfsdk:"default_value" json:"defaultValue,omitempty"`
					Maximum      *float64 `tfsdk:"maximum" json:"maximum,omitempty"`
					Minimum      *float64 `tfsdk:"minimum" json:"minimum,omitempty"`
					Unit         *string  `tfsdk:"unit" json:"unit,omitempty"`
				} `tfsdk:"double" json:"double,omitempty"`
				Float *struct {
					AccessMode   *string  `tfsdk:"access_mode" json:"accessMode,omitempty"`
					DefaultValue *float64 `tfsdk:"default_value" json:"defaultValue,omitempty"`
					Maximum      *float64 `tfsdk:"maximum" json:"maximum,omitempty"`
					Minimum      *float64 `tfsdk:"minimum" json:"minimum,omitempty"`
					Unit         *string  `tfsdk:"unit" json:"unit,omitempty"`
				} `tfsdk:"float" json:"float,omitempty"`
				Int *struct {
					AccessMode   *string `tfsdk:"access_mode" json:"accessMode,omitempty"`
					DefaultValue *int64  `tfsdk:"default_value" json:"defaultValue,omitempty"`
					Maximum      *int64  `tfsdk:"maximum" json:"maximum,omitempty"`
					Minimum      *int64  `tfsdk:"minimum" json:"minimum,omitempty"`
					Unit         *string `tfsdk:"unit" json:"unit,omitempty"`
				} `tfsdk:"int" json:"int,omitempty"`
				String *struct {
					AccessMode   *string `tfsdk:"access_mode" json:"accessMode,omitempty"`
					DefaultValue *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
				} `tfsdk:"string" json:"string,omitempty"`
			} `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"properties" json:"properties,omitempty"`
		Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DevicesKubeedgeIoDeviceModelV1Alpha2DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_devices_kubeedge_io_device_model_v1alpha2"
}

func (r *DevicesKubeedgeIoDeviceModelV1Alpha2DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DeviceModel is the Schema for the device model API",
		MarkdownDescription: "DeviceModel is the Schema for the device model API",
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
				Description:         "DeviceModelSpec defines the model / template for a device.It is a blueprint which describes the device capabilities and access mechanism via property visitors.",
				MarkdownDescription: "DeviceModelSpec defines the model / template for a device.It is a blueprint which describes the device capabilities and access mechanism via property visitors.",
				Attributes: map[string]schema.Attribute{
					"properties": schema.ListNestedAttribute{
						Description:         "Required: List of device properties.",
						MarkdownDescription: "Required: List of device properties.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"description": schema.StringAttribute{
									Description:         "The device property description.",
									MarkdownDescription: "The device property description.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "Required: The device property name.",
									MarkdownDescription: "Required: The device property name.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"type": schema.SingleNestedAttribute{
									Description:         "Required: PropertyType represents the type and data validation of the property.",
									MarkdownDescription: "Required: PropertyType represents the type and data validation of the property.",
									Attributes: map[string]schema.Attribute{
										"boolean": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"access_mode": schema.StringAttribute{
													Description:         "Required: Access mode of property, ReadWrite or ReadOnly.",
													MarkdownDescription: "Required: Access mode of property, ReadWrite or ReadOnly.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"default_value": schema.BoolAttribute{
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

										"bytes": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"access_mode": schema.StringAttribute{
													Description:         "Required: Access mode of property, ReadWrite or ReadOnly.",
													MarkdownDescription: "Required: Access mode of property, ReadWrite or ReadOnly.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"double": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"access_mode": schema.StringAttribute{
													Description:         "Required: Access mode of property, ReadWrite or ReadOnly.",
													MarkdownDescription: "Required: Access mode of property, ReadWrite or ReadOnly.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"default_value": schema.Float64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"maximum": schema.Float64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"minimum": schema.Float64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"unit": schema.StringAttribute{
													Description:         "The unit of the property",
													MarkdownDescription: "The unit of the property",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"float": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"access_mode": schema.StringAttribute{
													Description:         "Required: Access mode of property, ReadWrite or ReadOnly.",
													MarkdownDescription: "Required: Access mode of property, ReadWrite or ReadOnly.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"default_value": schema.Float64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"maximum": schema.Float64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"minimum": schema.Float64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"unit": schema.StringAttribute{
													Description:         "The unit of the property",
													MarkdownDescription: "The unit of the property",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"int": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"access_mode": schema.StringAttribute{
													Description:         "Required: Access mode of property, ReadWrite or ReadOnly.",
													MarkdownDescription: "Required: Access mode of property, ReadWrite or ReadOnly.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"default_value": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"maximum": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"minimum": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"unit": schema.StringAttribute{
													Description:         "The unit of the property",
													MarkdownDescription: "The unit of the property",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"string": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"access_mode": schema.StringAttribute{
													Description:         "Required: Access mode of property, ReadWrite or ReadOnly.",
													MarkdownDescription: "Required: Access mode of property, ReadWrite or ReadOnly.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"default_value": schema.StringAttribute{
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

					"protocol": schema.StringAttribute{
						Description:         "Required for DMI: Protocol name used by the device.",
						MarkdownDescription: "Required for DMI: Protocol name used by the device.",
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

func (r *DevicesKubeedgeIoDeviceModelV1Alpha2DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *DevicesKubeedgeIoDeviceModelV1Alpha2DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_devices_kubeedge_io_device_model_v1alpha2")

	var data DevicesKubeedgeIoDeviceModelV1Alpha2DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "devices.kubeedge.io", Version: "v1alpha2", Resource: "DeviceModel"}).
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

	var readResponse DevicesKubeedgeIoDeviceModelV1Alpha2DataSourceData
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
	data.ApiVersion = pointer.String("devices.kubeedge.io/v1alpha2")
	data.Kind = pointer.String("DeviceModel")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
