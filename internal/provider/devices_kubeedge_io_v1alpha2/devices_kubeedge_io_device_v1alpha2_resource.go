/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package devices_kubeedge_io_v1alpha2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"strings"
	"time"
)

var (
	_ resource.Resource                = &DevicesKubeedgeIoDeviceV1Alpha2Resource{}
	_ resource.ResourceWithConfigure   = &DevicesKubeedgeIoDeviceV1Alpha2Resource{}
	_ resource.ResourceWithImportState = &DevicesKubeedgeIoDeviceV1Alpha2Resource{}
)

func NewDevicesKubeedgeIoDeviceV1Alpha2Resource() resource.Resource {
	return &DevicesKubeedgeIoDeviceV1Alpha2Resource{}
}

type DevicesKubeedgeIoDeviceV1Alpha2Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type DevicesKubeedgeIoDeviceV1Alpha2ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitForUpsert  types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete  types.Object `tfsdk:"wait_for_delete" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Data *struct {
			DataProperties *[]struct {
				Metadata     *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
				PropertyName *string            `tfsdk:"property_name" json:"propertyName,omitempty"`
			} `tfsdk:"data_properties" json:"dataProperties,omitempty"`
			DataTopic *string `tfsdk:"data_topic" json:"dataTopic,omitempty"`
		} `tfsdk:"data" json:"data,omitempty"`
		DeviceModelRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"device_model_ref" json:"deviceModelRef,omitempty"`
		NodeSelector *struct {
			NodeSelectorTerms *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchFields *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_fields" json:"matchFields,omitempty"`
			} `tfsdk:"node_selector_terms" json:"nodeSelectorTerms,omitempty"`
		} `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		PropertyVisitors *[]struct {
			Bluetooth *struct {
				CharacteristicUUID *string `tfsdk:"characteristic_uuid" json:"characteristicUUID,omitempty"`
				DataConverter      *struct {
					EndIndex          *int64 `tfsdk:"end_index" json:"endIndex,omitempty"`
					OrderOfOperations *[]struct {
						OperationType  *string  `tfsdk:"operation_type" json:"operationType,omitempty"`
						OperationValue *float64 `tfsdk:"operation_value" json:"operationValue,omitempty"`
					} `tfsdk:"order_of_operations" json:"orderOfOperations,omitempty"`
					ShiftLeft  *int64 `tfsdk:"shift_left" json:"shiftLeft,omitempty"`
					ShiftRight *int64 `tfsdk:"shift_right" json:"shiftRight,omitempty"`
					StartIndex *int64 `tfsdk:"start_index" json:"startIndex,omitempty"`
				} `tfsdk:"data_converter" json:"dataConverter,omitempty"`
				DataWrite *map[string]string `tfsdk:"data_write" json:"dataWrite,omitempty"`
			} `tfsdk:"bluetooth" json:"bluetooth,omitempty"`
			CollectCycle       *int64 `tfsdk:"collect_cycle" json:"collectCycle,omitempty"`
			CustomizedProtocol *struct {
				ConfigData   *map[string]string `tfsdk:"config_data" json:"configData,omitempty"`
				ProtocolName *string            `tfsdk:"protocol_name" json:"protocolName,omitempty"`
			} `tfsdk:"customized_protocol" json:"customizedProtocol,omitempty"`
			CustomizedValues *map[string]string `tfsdk:"customized_values" json:"customizedValues,omitempty"`
			Modbus           *struct {
				IsRegisterSwap *bool    `tfsdk:"is_register_swap" json:"isRegisterSwap,omitempty"`
				IsSwap         *bool    `tfsdk:"is_swap" json:"isSwap,omitempty"`
				Limit          *int64   `tfsdk:"limit" json:"limit,omitempty"`
				Offset         *int64   `tfsdk:"offset" json:"offset,omitempty"`
				Register       *string  `tfsdk:"register" json:"register,omitempty"`
				Scale          *float64 `tfsdk:"scale" json:"scale,omitempty"`
			} `tfsdk:"modbus" json:"modbus,omitempty"`
			Opcua *struct {
				BrowseName *string `tfsdk:"browse_name" json:"browseName,omitempty"`
				NodeID     *string `tfsdk:"node_id" json:"nodeID,omitempty"`
			} `tfsdk:"opcua" json:"opcua,omitempty"`
			PropertyName *string `tfsdk:"property_name" json:"propertyName,omitempty"`
			ReportCycle  *int64  `tfsdk:"report_cycle" json:"reportCycle,omitempty"`
		} `tfsdk:"property_visitors" json:"propertyVisitors,omitempty"`
		Protocol *struct {
			Bluetooth *struct {
				MacAddress *string `tfsdk:"mac_address" json:"macAddress,omitempty"`
			} `tfsdk:"bluetooth" json:"bluetooth,omitempty"`
			Common *struct {
				CollectRetryTimes *int64  `tfsdk:"collect_retry_times" json:"collectRetryTimes,omitempty"`
				CollectTimeout    *int64  `tfsdk:"collect_timeout" json:"collectTimeout,omitempty"`
				CollectType       *string `tfsdk:"collect_type" json:"collectType,omitempty"`
				Com               *struct {
					BaudRate   *int64  `tfsdk:"baud_rate" json:"baudRate,omitempty"`
					DataBits   *int64  `tfsdk:"data_bits" json:"dataBits,omitempty"`
					Parity     *string `tfsdk:"parity" json:"parity,omitempty"`
					SerialPort *string `tfsdk:"serial_port" json:"serialPort,omitempty"`
					StopBits   *int64  `tfsdk:"stop_bits" json:"stopBits,omitempty"`
				} `tfsdk:"com" json:"com,omitempty"`
				CommType         *string            `tfsdk:"comm_type" json:"commType,omitempty"`
				CustomizedValues *map[string]string `tfsdk:"customized_values" json:"customizedValues,omitempty"`
				ReconnRetryTimes *int64             `tfsdk:"reconn_retry_times" json:"reconnRetryTimes,omitempty"`
				ReconnTimeout    *int64             `tfsdk:"reconn_timeout" json:"reconnTimeout,omitempty"`
				Tcp              *struct {
					Ip   *string `tfsdk:"ip" json:"ip,omitempty"`
					Port *int64  `tfsdk:"port" json:"port,omitempty"`
				} `tfsdk:"tcp" json:"tcp,omitempty"`
			} `tfsdk:"common" json:"common,omitempty"`
			CustomizedProtocol *struct {
				ConfigData   *map[string]string `tfsdk:"config_data" json:"configData,omitempty"`
				ProtocolName *string            `tfsdk:"protocol_name" json:"protocolName,omitempty"`
			} `tfsdk:"customized_protocol" json:"customizedProtocol,omitempty"`
			Modbus *struct {
				SlaveID *int64 `tfsdk:"slave_id" json:"slaveID,omitempty"`
			} `tfsdk:"modbus" json:"modbus,omitempty"`
			Opcua *struct {
				Certificate    *string `tfsdk:"certificate" json:"certificate,omitempty"`
				Password       *string `tfsdk:"password" json:"password,omitempty"`
				PrivateKey     *string `tfsdk:"private_key" json:"privateKey,omitempty"`
				SecurityMode   *string `tfsdk:"security_mode" json:"securityMode,omitempty"`
				SecurityPolicy *string `tfsdk:"security_policy" json:"securityPolicy,omitempty"`
				Timeout        *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
				Url            *string `tfsdk:"url" json:"url,omitempty"`
				UserName       *string `tfsdk:"user_name" json:"userName,omitempty"`
			} `tfsdk:"opcua" json:"opcua,omitempty"`
		} `tfsdk:"protocol" json:"protocol,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DevicesKubeedgeIoDeviceV1Alpha2Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_devices_kubeedge_io_device_v1alpha2"
}

func (r *DevicesKubeedgeIoDeviceV1Alpha2Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Device is the Schema for the devices API",
		MarkdownDescription: "Device is the Schema for the devices API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for_upsert": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
						"poll_interval": schema.StringAttribute{
							Description:         "The length of time to wait before checking again.",
							MarkdownDescription: "The length of time to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("5s"),
						},
					},
				},
			},

			"wait_for_delete": schema.SingleNestedAttribute{
				Description:         "Wait for deletion of resources.",
				MarkdownDescription: "Wait for deletion of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"timeout": schema.StringAttribute{
						Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("30s"),
					},
					"poll_interval": schema.StringAttribute{
						Description:         "The length of time to wait before checking again.",
						MarkdownDescription: "The length of time to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("5s"),
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "DeviceSpec represents a single device instance. It is an instantation of a device model.",
				MarkdownDescription: "DeviceSpec represents a single device instance. It is an instantation of a device model.",
				Attributes: map[string]schema.Attribute{
					"data": schema.SingleNestedAttribute{
						Description:         "Data section describe a list of time-series properties which should be processed on edge node.",
						MarkdownDescription: "Data section describe a list of time-series properties which should be processed on edge node.",
						Attributes: map[string]schema.Attribute{
							"data_properties": schema.ListNestedAttribute{
								Description:         "Required: A list of data properties, which are not required to be processed by edgecore",
								MarkdownDescription: "Required: A list of data properties, which are not required to be processed by edgecore",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"metadata": schema.MapAttribute{
											Description:         "Additional metadata like timestamp when the value was reported etc.",
											MarkdownDescription: "Additional metadata like timestamp when the value was reported etc.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"property_name": schema.StringAttribute{
											Description:         "Required: The property name for which should be processed by external apps. This property should be present in the device model.",
											MarkdownDescription: "Required: The property name for which should be processed by external apps. This property should be present in the device model.",
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

							"data_topic": schema.StringAttribute{
								Description:         "Topic used by mapper, all data collected from dataProperties should be published to this topic, the default value is $ke/events/device/+/data/update",
								MarkdownDescription: "Topic used by mapper, all data collected from dataProperties should be published to this topic, the default value is $ke/events/device/+/data/update",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

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

					"node_selector": schema.SingleNestedAttribute{
						Description:         "NodeSelector indicates the binding preferences between devices and nodes. Refer to k8s.io/kubernetes/pkg/apis/core NodeSelector for more details",
						MarkdownDescription: "NodeSelector indicates the binding preferences between devices and nodes. Refer to k8s.io/kubernetes/pkg/apis/core NodeSelector for more details",
						Attributes: map[string]schema.Attribute{
							"node_selector_terms": schema.ListNestedAttribute{
								Description:         "Required. A list of node selector terms. The terms are ORed.",
								MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"match_expressions": schema.ListNestedAttribute{
											Description:         "A list of node selector requirements by node's labels.",
											MarkdownDescription: "A list of node selector requirements by node's labels.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The label key that the selector applies to.",
														MarkdownDescription: "The label key that the selector applies to.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"operator": schema.StringAttribute{
														Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
														MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"values": schema.ListAttribute{
														Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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

										"match_fields": schema.ListNestedAttribute{
											Description:         "A list of node selector requirements by node's fields.",
											MarkdownDescription: "A list of node selector requirements by node's fields.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The label key that the selector applies to.",
														MarkdownDescription: "The label key that the selector applies to.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"operator": schema.StringAttribute{
														Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
														MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"values": schema.ListAttribute{
														Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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

					"property_visitors": schema.ListNestedAttribute{
						Description:         "List of property visitors which describe how to access the device properties. PropertyVisitors must unique by propertyVisitor.propertyName.",
						MarkdownDescription: "List of property visitors which describe how to access the device properties. PropertyVisitors must unique by propertyVisitor.propertyName.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"bluetooth": schema.SingleNestedAttribute{
									Description:         "Bluetooth represents a set of additional visitor config fields of bluetooth protocol.",
									MarkdownDescription: "Bluetooth represents a set of additional visitor config fields of bluetooth protocol.",
									Attributes: map[string]schema.Attribute{
										"characteristic_uuid": schema.StringAttribute{
											Description:         "Required: Unique ID of the corresponding operation",
											MarkdownDescription: "Required: Unique ID of the corresponding operation",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"data_converter": schema.SingleNestedAttribute{
											Description:         "Responsible for converting the data being read from the bluetooth device into a form that is understandable by the platform",
											MarkdownDescription: "Responsible for converting the data being read from the bluetooth device into a form that is understandable by the platform",
											Attributes: map[string]schema.Attribute{
												"end_index": schema.Int64Attribute{
													Description:         "Required: Specifies the end index of incoming byte stream to be considered to convert the data the value specified should be inclusive for example if 3 is specified it includes the third index",
													MarkdownDescription: "Required: Specifies the end index of incoming byte stream to be considered to convert the data the value specified should be inclusive for example if 3 is specified it includes the third index",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"order_of_operations": schema.ListNestedAttribute{
													Description:         "Specifies in what order the operations(which are required to be performed to convert incoming data into understandable form) are performed",
													MarkdownDescription: "Specifies in what order the operations(which are required to be performed to convert incoming data into understandable form) are performed",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"operation_type": schema.StringAttribute{
																Description:         "Required: Specifies the operation to be performed to convert incoming data",
																MarkdownDescription: "Required: Specifies the operation to be performed to convert incoming data",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"operation_value": schema.Float64Attribute{
																Description:         "Required: Specifies with what value the operation is to be performed",
																MarkdownDescription: "Required: Specifies with what value the operation is to be performed",
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

												"shift_left": schema.Int64Attribute{
													Description:         "Refers to the number of bits to shift left, if left-shift operation is necessary for conversion",
													MarkdownDescription: "Refers to the number of bits to shift left, if left-shift operation is necessary for conversion",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"shift_right": schema.Int64Attribute{
													Description:         "Refers to the number of bits to shift right, if right-shift operation is necessary for conversion",
													MarkdownDescription: "Refers to the number of bits to shift right, if right-shift operation is necessary for conversion",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"start_index": schema.Int64Attribute{
													Description:         "Required: Specifies the start index of the incoming byte stream to be considered to convert the data. For example: start-index:2, end-index:3 concatenates the value present at second and third index of the incoming byte stream. If we want to reverse the order we can give it as start-index:3, end-index:2",
													MarkdownDescription: "Required: Specifies the start index of the incoming byte stream to be considered to convert the data. For example: start-index:2, end-index:3 concatenates the value present at second and third index of the incoming byte stream. If we want to reverse the order we can give it as start-index:3, end-index:2",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"data_write": schema.MapAttribute{
											Description:         "Responsible for converting the data coming from the platform into a form that is understood by the bluetooth device For example: 'ON':[1], 'OFF':[0]",
											MarkdownDescription: "Responsible for converting the data coming from the platform into a form that is understood by the bluetooth device For example: 'ON':[1], 'OFF':[0]",
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

								"collect_cycle": schema.Int64Attribute{
									Description:         "Define how frequent mapper will collect from device.",
									MarkdownDescription: "Define how frequent mapper will collect from device.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"customized_protocol": schema.SingleNestedAttribute{
									Description:         "CustomizedProtocol represents a set of visitor config fields of bluetooth protocol.",
									MarkdownDescription: "CustomizedProtocol represents a set of visitor config fields of bluetooth protocol.",
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

								"customized_values": schema.MapAttribute{
									Description:         "Customized values for visitor of provided protocols",
									MarkdownDescription: "Customized values for visitor of provided protocols",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"modbus": schema.SingleNestedAttribute{
									Description:         "Modbus represents a set of additional visitor config fields of modbus protocol.",
									MarkdownDescription: "Modbus represents a set of additional visitor config fields of modbus protocol.",
									Attributes: map[string]schema.Attribute{
										"is_register_swap": schema.BoolAttribute{
											Description:         "Indicates whether the high and low register swapped. Defaults to false.",
											MarkdownDescription: "Indicates whether the high and low register swapped. Defaults to false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"is_swap": schema.BoolAttribute{
											Description:         "Indicates whether the high and low byte swapped. Defaults to false.",
											MarkdownDescription: "Indicates whether the high and low byte swapped. Defaults to false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"limit": schema.Int64Attribute{
											Description:         "Required: Limit number of registers to read/write.",
											MarkdownDescription: "Required: Limit number of registers to read/write.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"offset": schema.Int64Attribute{
											Description:         "Required: Offset indicates the starting register number to read/write data.",
											MarkdownDescription: "Required: Offset indicates the starting register number to read/write data.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"register": schema.StringAttribute{
											Description:         "Required: Type of register",
											MarkdownDescription: "Required: Type of register",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("CoilRegister", "DiscreteInputRegister", "InputRegister", "HoldingRegister"),
											},
										},

										"scale": schema.Float64Attribute{
											Description:         "The scale to convert raw property data into final units. Defaults to 1.0",
											MarkdownDescription: "The scale to convert raw property data into final units. Defaults to 1.0",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"opcua": schema.SingleNestedAttribute{
									Description:         "Opcua represents a set of additional visitor config fields of opc-ua protocol.",
									MarkdownDescription: "Opcua represents a set of additional visitor config fields of opc-ua protocol.",
									Attributes: map[string]schema.Attribute{
										"browse_name": schema.StringAttribute{
											Description:         "The name of opc-ua node",
											MarkdownDescription: "The name of opc-ua node",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"node_id": schema.StringAttribute{
											Description:         "Required: The ID of opc-ua node, e.g. 'ns=1,i=1005'",
											MarkdownDescription: "Required: The ID of opc-ua node, e.g. 'ns=1,i=1005'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"property_name": schema.StringAttribute{
									Description:         "Required: The device property name to be accessed. This should refer to one of the device properties defined in the device model.",
									MarkdownDescription: "Required: The device property name to be accessed. This should refer to one of the device properties defined in the device model.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"report_cycle": schema.Int64Attribute{
									Description:         "Define how frequent mapper will report the value.",
									MarkdownDescription: "Define how frequent mapper will report the value.",
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

					"protocol": schema.SingleNestedAttribute{
						Description:         "Required: The protocol configuration used to connect to the device.",
						MarkdownDescription: "Required: The protocol configuration used to connect to the device.",
						Attributes: map[string]schema.Attribute{
							"bluetooth": schema.SingleNestedAttribute{
								Description:         "Protocol configuration for bluetooth",
								MarkdownDescription: "Protocol configuration for bluetooth",
								Attributes: map[string]schema.Attribute{
									"mac_address": schema.StringAttribute{
										Description:         "Unique identifier assigned to the device.",
										MarkdownDescription: "Unique identifier assigned to the device.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"common": schema.SingleNestedAttribute{
								Description:         "Configuration for protocol common part",
								MarkdownDescription: "Configuration for protocol common part",
								Attributes: map[string]schema.Attribute{
									"collect_retry_times": schema.Int64Attribute{
										Description:         "Define retry times of mapper will collect from device.",
										MarkdownDescription: "Define retry times of mapper will collect from device.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"collect_timeout": schema.Int64Attribute{
										Description:         "Define timeout of mapper collect from device.",
										MarkdownDescription: "Define timeout of mapper collect from device.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"collect_type": schema.StringAttribute{
										Description:         "Define collect type, sync or async.",
										MarkdownDescription: "Define collect type, sync or async.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("sync", "async"),
										},
									},

									"com": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"baud_rate": schema.Int64Attribute{
												Description:         "Required. BaudRate 115200|57600|38400|19200|9600|4800|2400|1800|1200|600|300|200|150|134|110|75|50",
												MarkdownDescription: "Required. BaudRate 115200|57600|38400|19200|9600|4800|2400|1800|1200|600|300|200|150|134|110|75|50",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.OneOf(115200, 57600, 38400, 19200, 9600, 4800, 2400, 1800, 1200, 600, 300, 200, 150, 134, 110, 75, 50),
												},
											},

											"data_bits": schema.Int64Attribute{
												Description:         "Required. Valid values are 8, 7, 6, 5.",
												MarkdownDescription: "Required. Valid values are 8, 7, 6, 5.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.OneOf(8, 7, 6, 5),
												},
											},

											"parity": schema.StringAttribute{
												Description:         "Required. Valid options are 'none', 'even', 'odd'. Defaults to 'none'.",
												MarkdownDescription: "Required. Valid options are 'none', 'even', 'odd'. Defaults to 'none'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("none", "even", "odd"),
												},
											},

											"serial_port": schema.StringAttribute{
												Description:         "Required.",
												MarkdownDescription: "Required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"stop_bits": schema.Int64Attribute{
												Description:         "Required. Bit that stops 1|2",
												MarkdownDescription: "Required. Bit that stops 1|2",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.OneOf(1, 2),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"comm_type": schema.StringAttribute{
										Description:         "Communication type, like tcp client, tcp server or COM",
										MarkdownDescription: "Communication type, like tcp client, tcp server or COM",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"customized_values": schema.MapAttribute{
										Description:         "Customized values for provided protocol",
										MarkdownDescription: "Customized values for provided protocol",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"reconn_retry_times": schema.Int64Attribute{
										Description:         "Reconnecting retry times",
										MarkdownDescription: "Reconnecting retry times",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"reconn_timeout": schema.Int64Attribute{
										Description:         "Reconnection timeout",
										MarkdownDescription: "Reconnection timeout",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tcp": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"ip": schema.StringAttribute{
												Description:         "Required.",
												MarkdownDescription: "Required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "Required.",
												MarkdownDescription: "Required.",
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

							"customized_protocol": schema.SingleNestedAttribute{
								Description:         "Configuration for customized protocol",
								MarkdownDescription: "Configuration for customized protocol",
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

							"modbus": schema.SingleNestedAttribute{
								Description:         "Protocol configuration for modbus",
								MarkdownDescription: "Protocol configuration for modbus",
								Attributes: map[string]schema.Attribute{
									"slave_id": schema.Int64Attribute{
										Description:         "Required. 0-255",
										MarkdownDescription: "Required. 0-255",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"opcua": schema.SingleNestedAttribute{
								Description:         "Protocol configuration for opc-ua",
								MarkdownDescription: "Protocol configuration for opc-ua",
								Attributes: map[string]schema.Attribute{
									"certificate": schema.StringAttribute{
										Description:         "Certificate for access opc server.",
										MarkdownDescription: "Certificate for access opc server.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"password": schema.StringAttribute{
										Description:         "Password for access opc server.",
										MarkdownDescription: "Password for access opc server.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"private_key": schema.StringAttribute{
										Description:         "PrivateKey for access opc server.",
										MarkdownDescription: "PrivateKey for access opc server.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"security_mode": schema.StringAttribute{
										Description:         "Defaults to 'none'.",
										MarkdownDescription: "Defaults to 'none'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"security_policy": schema.StringAttribute{
										Description:         "Defaults to 'none'.",
										MarkdownDescription: "Defaults to 'none'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"timeout": schema.Int64Attribute{
										Description:         "Timeout seconds for the opc server connection.???",
										MarkdownDescription: "Timeout seconds for the opc server connection.???",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"url": schema.StringAttribute{
										Description:         "Required: The URL for opc server endpoint.",
										MarkdownDescription: "Required: The URL for opc server endpoint.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"user_name": schema.StringAttribute{
										Description:         "Username for access opc server.",
										MarkdownDescription: "Username for access opc server.",
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
		},
	}
}

func (r *DevicesKubeedgeIoDeviceV1Alpha2Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedResourceDataError(request.ProviderData))
	}
}

func (r *DevicesKubeedgeIoDeviceV1Alpha2Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_devices_kubeedge_io_device_v1alpha2")

	var model DevicesKubeedgeIoDeviceV1Alpha2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("devices.kubeedge.io/v1alpha2")
	model.Kind = pointer.String("Device")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "devices.kubeedge.io", Version: "v1alpha2", Resource: "devices"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse DevicesKubeedgeIoDeviceV1Alpha2ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *DevicesKubeedgeIoDeviceV1Alpha2Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_devices_kubeedge_io_device_v1alpha2")

	var data DevicesKubeedgeIoDeviceV1Alpha2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "devices.kubeedge.io", Version: "v1alpha2", Resource: "devices"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse DevicesKubeedgeIoDeviceV1Alpha2ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *DevicesKubeedgeIoDeviceV1Alpha2Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_devices_kubeedge_io_device_v1alpha2")

	var model DevicesKubeedgeIoDeviceV1Alpha2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("devices.kubeedge.io/v1alpha2")
	model.Kind = pointer.String("Device")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "devices.kubeedge.io", Version: "v1alpha2", Resource: "devices"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse DevicesKubeedgeIoDeviceV1Alpha2ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *DevicesKubeedgeIoDeviceV1Alpha2Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_devices_kubeedge_io_device_v1alpha2")

	var data DevicesKubeedgeIoDeviceV1Alpha2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "devices.kubeedge.io", Version: "v1alpha2", Resource: "devices"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "devices.kubeedge.io", Version: "v1alpha2", Resource: "devices"}).
				Namespace(data.Metadata.Namespace).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout == time.Second*0 {
				break
			}
			if time.Now().After(startTime.Add(timeout)) {
				response.Diagnostics.Append(utilities.WaitTimeoutExceeded())
				return
			}
			time.Sleep(pollInterval)
		}
	}
}

func (r *DevicesKubeedgeIoDeviceV1Alpha2Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
