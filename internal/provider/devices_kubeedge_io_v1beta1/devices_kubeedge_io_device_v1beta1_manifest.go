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
	_ datasource.DataSource = &DevicesKubeedgeIoDeviceV1Beta1Manifest{}
)

func NewDevicesKubeedgeIoDeviceV1Beta1Manifest() datasource.DataSource {
	return &DevicesKubeedgeIoDeviceV1Beta1Manifest{}
}

type DevicesKubeedgeIoDeviceV1Beta1Manifest struct{}

type DevicesKubeedgeIoDeviceV1Beta1ManifestData struct {
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
		DeviceModelRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"device_model_ref" json:"deviceModelRef,omitempty"`
		NodeName   *string `tfsdk:"node_name" json:"nodeName,omitempty"`
		Properties *[]struct {
			CollectCycle *int64 `tfsdk:"collect_cycle" json:"collectCycle,omitempty"`
			Desired      *struct {
				Metadata *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
				Value    *string            `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"desired" json:"desired,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			PushMethod *struct {
				DbMethod *struct {
					TDEngine *struct {
						TDEngineClientConfig *struct {
							Addr   *string `tfsdk:"addr" json:"addr,omitempty"`
							DbName *string `tfsdk:"db_name" json:"dbName,omitempty"`
						} `tfsdk:"td_engine_client_config" json:"TDEngineClientConfig,omitempty"`
					} `tfsdk:"td_engine" json:"TDEngine,omitempty"`
					Influxdb2 *struct {
						Influxdb2ClientConfig *struct {
							Bucket *string `tfsdk:"bucket" json:"bucket,omitempty"`
							Org    *string `tfsdk:"org" json:"org,omitempty"`
							Url    *string `tfsdk:"url" json:"url,omitempty"`
						} `tfsdk:"influxdb2_client_config" json:"influxdb2ClientConfig,omitempty"`
						Influxdb2DataConfig *struct {
							FieldKey    *string            `tfsdk:"field_key" json:"fieldKey,omitempty"`
							Measurement *string            `tfsdk:"measurement" json:"measurement,omitempty"`
							Tag         *map[string]string `tfsdk:"tag" json:"tag,omitempty"`
						} `tfsdk:"influxdb2_data_config" json:"influxdb2DataConfig,omitempty"`
					} `tfsdk:"influxdb2" json:"influxdb2,omitempty"`
					Redis *struct {
						RedisClientConfig *struct {
							Addr         *string `tfsdk:"addr" json:"addr,omitempty"`
							Db           *int64  `tfsdk:"db" json:"db,omitempty"`
							MinIdleConns *int64  `tfsdk:"min_idle_conns" json:"minIdleConns,omitempty"`
							Poolsize     *int64  `tfsdk:"poolsize" json:"poolsize,omitempty"`
						} `tfsdk:"redis_client_config" json:"redisClientConfig,omitempty"`
					} `tfsdk:"redis" json:"redis,omitempty"`
				} `tfsdk:"db_method" json:"dbMethod,omitempty"`
				Http *struct {
					HostName    *string `tfsdk:"host_name" json:"hostName,omitempty"`
					Port        *int64  `tfsdk:"port" json:"port,omitempty"`
					RequestPath *string `tfsdk:"request_path" json:"requestPath,omitempty"`
					Timeout     *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"http" json:"http,omitempty"`
				Mqtt *struct {
					Address  *string `tfsdk:"address" json:"address,omitempty"`
					Qos      *int64  `tfsdk:"qos" json:"qos,omitempty"`
					Retained *bool   `tfsdk:"retained" json:"retained,omitempty"`
					Topic    *string `tfsdk:"topic" json:"topic,omitempty"`
				} `tfsdk:"mqtt" json:"mqtt,omitempty"`
			} `tfsdk:"push_method" json:"pushMethod,omitempty"`
			ReportCycle   *int64 `tfsdk:"report_cycle" json:"reportCycle,omitempty"`
			ReportToCloud *bool  `tfsdk:"report_to_cloud" json:"reportToCloud,omitempty"`
			Visitors      *struct {
				ConfigData   *map[string]string `tfsdk:"config_data" json:"configData,omitempty"`
				ProtocolName *string            `tfsdk:"protocol_name" json:"protocolName,omitempty"`
			} `tfsdk:"visitors" json:"visitors,omitempty"`
		} `tfsdk:"properties" json:"properties,omitempty"`
		Protocol *struct {
			ConfigData   *map[string]string `tfsdk:"config_data" json:"configData,omitempty"`
			ProtocolName *string            `tfsdk:"protocol_name" json:"protocolName,omitempty"`
		} `tfsdk:"protocol" json:"protocol,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DevicesKubeedgeIoDeviceV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_devices_kubeedge_io_device_v1beta1_manifest"
}

func (r *DevicesKubeedgeIoDeviceV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Device is the Schema for the devices API",
		MarkdownDescription: "Device is the Schema for the devices API",
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
				Description:         "DeviceSpec represents a single device instance.",
				MarkdownDescription: "DeviceSpec represents a single device instance.",
				Attributes: map[string]schema.Attribute{
					"device_model_ref": schema.SingleNestedAttribute{
						Description:         "Required: DeviceModelRef is reference to the device model used as a template to create the device instance.",
						MarkdownDescription: "Required: DeviceModelRef is reference to the device model used as a template to create the device instance.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_name": schema.StringAttribute{
						Description:         "NodeName is a request to schedule this device onto a specific node. If it is non-empty, the scheduler simply schedules this device onto that node, assuming that it fits resource requirements.",
						MarkdownDescription: "NodeName is a request to schedule this device onto a specific node. If it is non-empty, the scheduler simply schedules this device onto that node, assuming that it fits resource requirements.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"properties": schema.ListNestedAttribute{
						Description:         "List of properties which describe the device properties. properties list item must be unique by properties.Name.",
						MarkdownDescription: "List of properties which describe the device properties. properties list item must be unique by properties.Name.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"collect_cycle": schema.Int64Attribute{
									Description:         "Define how frequent mapper will collect from device.",
									MarkdownDescription: "Define how frequent mapper will collect from device.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"desired": schema.SingleNestedAttribute{
									Description:         "The desired property value.",
									MarkdownDescription: "The desired property value.",
									Attributes: map[string]schema.Attribute{
										"metadata": schema.MapAttribute{
											Description:         "Additional metadata like timestamp when the value was reported etc.",
											MarkdownDescription: "Additional metadata like timestamp when the value was reported etc.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Required: The value for this property.",
											MarkdownDescription: "Required: The value for this property.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"name": schema.StringAttribute{
									Description:         "Required: The device property name to be accessed. It must be unique.",
									MarkdownDescription: "Required: The device property name to be accessed. It must be unique.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"push_method": schema.SingleNestedAttribute{
									Description:         "PushMethod represents the protocol used to push data, please ensure that the mapper can access the destination address.",
									MarkdownDescription: "PushMethod represents the protocol used to push data, please ensure that the mapper can access the destination address.",
									Attributes: map[string]schema.Attribute{
										"db_method": schema.SingleNestedAttribute{
											Description:         "DBMethod represents the method used to push data to database, please ensure that the mapper can access the destination address.",
											MarkdownDescription: "DBMethod represents the method used to push data to database, please ensure that the mapper can access the destination address.",
											Attributes: map[string]schema.Attribute{
												"td_engine": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"td_engine_client_config": schema.SingleNestedAttribute{
															Description:         "tdengineClientConfig of tdengine database",
															MarkdownDescription: "tdengineClientConfig of tdengine database",
															Attributes: map[string]schema.Attribute{
																"addr": schema.StringAttribute{
																	Description:         "addr of tdEngine database",
																	MarkdownDescription: "addr of tdEngine database",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"db_name": schema.StringAttribute{
																	Description:         "dbname of tdEngine database",
																	MarkdownDescription: "dbname of tdEngine database",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"influxdb2": schema.SingleNestedAttribute{
													Description:         "method configuration for database",
													MarkdownDescription: "method configuration for database",
													Attributes: map[string]schema.Attribute{
														"influxdb2_client_config": schema.SingleNestedAttribute{
															Description:         "Config of influx database",
															MarkdownDescription: "Config of influx database",
															Attributes: map[string]schema.Attribute{
																"bucket": schema.StringAttribute{
																	Description:         "Bucket of the user in influx database",
																	MarkdownDescription: "Bucket of the user in influx database",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"org": schema.StringAttribute{
																	Description:         "Org of the user in influx database",
																	MarkdownDescription: "Org of the user in influx database",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"url": schema.StringAttribute{
																	Description:         "Url of influx database",
																	MarkdownDescription: "Url of influx database",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"influxdb2_data_config": schema.SingleNestedAttribute{
															Description:         "config of device data when push to influx database",
															MarkdownDescription: "config of device data when push to influx database",
															Attributes: map[string]schema.Attribute{
																"field_key": schema.StringAttribute{
																	Description:         "FieldKey of the user data",
																	MarkdownDescription: "FieldKey of the user data",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"measurement": schema.StringAttribute{
																	Description:         "Measurement of the user data",
																	MarkdownDescription: "Measurement of the user data",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"tag": schema.MapAttribute{
																	Description:         "the tag of device data",
																	MarkdownDescription: "the tag of device data",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"redis": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"redis_client_config": schema.SingleNestedAttribute{
															Description:         "RedisClientConfig of redis database",
															MarkdownDescription: "RedisClientConfig of redis database",
															Attributes: map[string]schema.Attribute{
																"addr": schema.StringAttribute{
																	Description:         "Addr of Redis database",
																	MarkdownDescription: "Addr of Redis database",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"db": schema.Int64Attribute{
																	Description:         "Db of Redis database",
																	MarkdownDescription: "Db of Redis database",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"min_idle_conns": schema.Int64Attribute{
																	Description:         "MinIdleConns of Redis database",
																	MarkdownDescription: "MinIdleConns of Redis database",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"poolsize": schema.Int64Attribute{
																	Description:         "Poolsize of Redis database",
																	MarkdownDescription: "Poolsize of Redis database",
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
													Required: false,
													Optional: true,
													Computed: false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"http": schema.SingleNestedAttribute{
											Description:         "HTTP Push method configuration for http",
											MarkdownDescription: "HTTP Push method configuration for http",
											Attributes: map[string]schema.Attribute{
												"host_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"request_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"timeout": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"mqtt": schema.SingleNestedAttribute{
											Description:         "MQTT Push method configuration for mqtt",
											MarkdownDescription: "MQTT Push method configuration for mqtt",
											Attributes: map[string]schema.Attribute{
												"address": schema.StringAttribute{
													Description:         "broker address, like mqtt://127.0.0.1:1883",
													MarkdownDescription: "broker address, like mqtt://127.0.0.1:1883",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"qos": schema.Int64Attribute{
													Description:         "qos of mqtt publish param",
													MarkdownDescription: "qos of mqtt publish param",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"retained": schema.BoolAttribute{
													Description:         "Is the message retained",
													MarkdownDescription: "Is the message retained",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"topic": schema.StringAttribute{
													Description:         "publish topic for mqtt",
													MarkdownDescription: "publish topic for mqtt",
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"report_cycle": schema.Int64Attribute{
									Description:         "Define how frequent mapper will report the value.",
									MarkdownDescription: "Define how frequent mapper will report the value.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"report_to_cloud": schema.BoolAttribute{
									Description:         "whether be reported to the cloud",
									MarkdownDescription: "whether be reported to the cloud",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"visitors": schema.SingleNestedAttribute{
									Description:         "Visitors are intended to be consumed by device mappers which connect to devices and collect data / perform actions on the device. Required: Protocol relevant config details about the how to access the device property.",
									MarkdownDescription: "Visitors are intended to be consumed by device mappers which connect to devices and collect data / perform actions on the device. Required: Protocol relevant config details about the how to access the device property.",
									Attributes: map[string]schema.Attribute{
										"config_data": schema.MapAttribute{
											Description:         "Required: The configData of customized protocol",
											MarkdownDescription: "Required: The configData of customized protocol",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"protocol_name": schema.StringAttribute{
											Description:         "Required: name of customized protocol",
											MarkdownDescription: "Required: name of customized protocol",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"protocol": schema.SingleNestedAttribute{
						Description:         "Required: The protocol configuration used to connect to the device.",
						MarkdownDescription: "Required: The protocol configuration used to connect to the device.",
						Attributes: map[string]schema.Attribute{
							"config_data": schema.MapAttribute{
								Description:         "Any config data",
								MarkdownDescription: "Any config data",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"protocol_name": schema.StringAttribute{
								Description:         "Unique protocol name Required.",
								MarkdownDescription: "Unique protocol name Required.",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *DevicesKubeedgeIoDeviceV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_devices_kubeedge_io_device_v1beta1_manifest")

	var model DevicesKubeedgeIoDeviceV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("devices.kubeedge.io/v1beta1")
	model.Kind = pointer.String("Device")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
