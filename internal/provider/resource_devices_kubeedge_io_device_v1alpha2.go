/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type DevicesKubeedgeIoDeviceV1Alpha2Resource struct{}

var (
	_ resource.Resource = (*DevicesKubeedgeIoDeviceV1Alpha2Resource)(nil)
)

type DevicesKubeedgeIoDeviceV1Alpha2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type DevicesKubeedgeIoDeviceV1Alpha2GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Data *struct {
			DataProperties *[]struct {
				Metadata *map[string]string `tfsdk:"metadata" yaml:"metadata,omitempty"`

				PropertyName *string `tfsdk:"property_name" yaml:"propertyName,omitempty"`
			} `tfsdk:"data_properties" yaml:"dataProperties,omitempty"`

			DataTopic *string `tfsdk:"data_topic" yaml:"dataTopic,omitempty"`
		} `tfsdk:"data" yaml:"data,omitempty"`

		DeviceModelRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"device_model_ref" yaml:"deviceModelRef,omitempty"`

		NodeSelector *struct {
			NodeSelectorTerms *[]struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchFields *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_fields" yaml:"matchFields,omitempty"`
			} `tfsdk:"node_selector_terms" yaml:"nodeSelectorTerms,omitempty"`
		} `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

		PropertyVisitors *[]struct {
			Bluetooth *struct {
				CharacteristicUUID *string `tfsdk:"characteristic_uuid" yaml:"characteristicUUID,omitempty"`

				DataConverter *struct {
					EndIndex *int64 `tfsdk:"end_index" yaml:"endIndex,omitempty"`

					OrderOfOperations *[]struct {
						OperationType *string `tfsdk:"operation_type" yaml:"operationType,omitempty"`

						OperationValue utilities.DynamicNumber `tfsdk:"operation_value" yaml:"operationValue,omitempty"`
					} `tfsdk:"order_of_operations" yaml:"orderOfOperations,omitempty"`

					ShiftLeft *int64 `tfsdk:"shift_left" yaml:"shiftLeft,omitempty"`

					ShiftRight *int64 `tfsdk:"shift_right" yaml:"shiftRight,omitempty"`

					StartIndex *int64 `tfsdk:"start_index" yaml:"startIndex,omitempty"`
				} `tfsdk:"data_converter" yaml:"dataConverter,omitempty"`

				DataWrite *map[string]string `tfsdk:"data_write" yaml:"dataWrite,omitempty"`
			} `tfsdk:"bluetooth" yaml:"bluetooth,omitempty"`

			CollectCycle *int64 `tfsdk:"collect_cycle" yaml:"collectCycle,omitempty"`

			CustomizedProtocol *struct {
				ConfigData utilities.Dynamic `tfsdk:"config_data" yaml:"configData,omitempty"`

				ProtocolName *string `tfsdk:"protocol_name" yaml:"protocolName,omitempty"`
			} `tfsdk:"customized_protocol" yaml:"customizedProtocol,omitempty"`

			CustomizedValues utilities.Dynamic `tfsdk:"customized_values" yaml:"customizedValues,omitempty"`

			Modbus *struct {
				IsRegisterSwap *bool `tfsdk:"is_register_swap" yaml:"isRegisterSwap,omitempty"`

				IsSwap *bool `tfsdk:"is_swap" yaml:"isSwap,omitempty"`

				Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`

				Offset *int64 `tfsdk:"offset" yaml:"offset,omitempty"`

				Register *string `tfsdk:"register" yaml:"register,omitempty"`

				Scale utilities.DynamicNumber `tfsdk:"scale" yaml:"scale,omitempty"`
			} `tfsdk:"modbus" yaml:"modbus,omitempty"`

			Opcua *struct {
				BrowseName *string `tfsdk:"browse_name" yaml:"browseName,omitempty"`

				NodeID *string `tfsdk:"node_id" yaml:"nodeID,omitempty"`
			} `tfsdk:"opcua" yaml:"opcua,omitempty"`

			PropertyName *string `tfsdk:"property_name" yaml:"propertyName,omitempty"`

			ReportCycle *int64 `tfsdk:"report_cycle" yaml:"reportCycle,omitempty"`
		} `tfsdk:"property_visitors" yaml:"propertyVisitors,omitempty"`

		Protocol *struct {
			Bluetooth *struct {
				MacAddress *string `tfsdk:"mac_address" yaml:"macAddress,omitempty"`
			} `tfsdk:"bluetooth" yaml:"bluetooth,omitempty"`

			Common *struct {
				CollectRetryTimes *int64 `tfsdk:"collect_retry_times" yaml:"collectRetryTimes,omitempty"`

				CollectTimeout *int64 `tfsdk:"collect_timeout" yaml:"collectTimeout,omitempty"`

				CollectType *string `tfsdk:"collect_type" yaml:"collectType,omitempty"`

				Com *struct {
					BaudRate *int64 `tfsdk:"baud_rate" yaml:"baudRate,omitempty"`

					DataBits *int64 `tfsdk:"data_bits" yaml:"dataBits,omitempty"`

					Parity *string `tfsdk:"parity" yaml:"parity,omitempty"`

					SerialPort *string `tfsdk:"serial_port" yaml:"serialPort,omitempty"`

					StopBits *int64 `tfsdk:"stop_bits" yaml:"stopBits,omitempty"`
				} `tfsdk:"com" yaml:"com,omitempty"`

				CommType *string `tfsdk:"comm_type" yaml:"commType,omitempty"`

				CustomizedValues utilities.Dynamic `tfsdk:"customized_values" yaml:"customizedValues,omitempty"`

				ReconnRetryTimes *int64 `tfsdk:"reconn_retry_times" yaml:"reconnRetryTimes,omitempty"`

				ReconnTimeout *int64 `tfsdk:"reconn_timeout" yaml:"reconnTimeout,omitempty"`

				Tcp *struct {
					Ip *string `tfsdk:"ip" yaml:"ip,omitempty"`

					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`
				} `tfsdk:"tcp" yaml:"tcp,omitempty"`
			} `tfsdk:"common" yaml:"common,omitempty"`

			CustomizedProtocol *struct {
				ConfigData utilities.Dynamic `tfsdk:"config_data" yaml:"configData,omitempty"`

				ProtocolName *string `tfsdk:"protocol_name" yaml:"protocolName,omitempty"`
			} `tfsdk:"customized_protocol" yaml:"customizedProtocol,omitempty"`

			Modbus *struct {
				SlaveID *int64 `tfsdk:"slave_id" yaml:"slaveID,omitempty"`
			} `tfsdk:"modbus" yaml:"modbus,omitempty"`

			Opcua *struct {
				Certificate *string `tfsdk:"certificate" yaml:"certificate,omitempty"`

				Password *string `tfsdk:"password" yaml:"password,omitempty"`

				PrivateKey *string `tfsdk:"private_key" yaml:"privateKey,omitempty"`

				SecurityMode *string `tfsdk:"security_mode" yaml:"securityMode,omitempty"`

				SecurityPolicy *string `tfsdk:"security_policy" yaml:"securityPolicy,omitempty"`

				Timeout *int64 `tfsdk:"timeout" yaml:"timeout,omitempty"`

				Url *string `tfsdk:"url" yaml:"url,omitempty"`

				UserName *string `tfsdk:"user_name" yaml:"userName,omitempty"`
			} `tfsdk:"opcua" yaml:"opcua,omitempty"`
		} `tfsdk:"protocol" yaml:"protocol,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewDevicesKubeedgeIoDeviceV1Alpha2Resource() resource.Resource {
	return &DevicesKubeedgeIoDeviceV1Alpha2Resource{}
}

func (r *DevicesKubeedgeIoDeviceV1Alpha2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_kubeedge_io_device_v1alpha2"
}

func (r *DevicesKubeedgeIoDeviceV1Alpha2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Device is the Schema for the devices API",
		MarkdownDescription: "Device is the Schema for the devices API",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "DeviceSpec represents a single device instance. It is an instantation of a device model.",
				MarkdownDescription: "DeviceSpec represents a single device instance. It is an instantation of a device model.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"data": {
						Description:         "Data section describe a list of time-series properties which should be processed on edge node.",
						MarkdownDescription: "Data section describe a list of time-series properties which should be processed on edge node.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"data_properties": {
								Description:         "Required: A list of data properties, which are not required to be processed by edgecore",
								MarkdownDescription: "Required: A list of data properties, which are not required to be processed by edgecore",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"metadata": {
										Description:         "Additional metadata like timestamp when the value was reported etc.",
										MarkdownDescription: "Additional metadata like timestamp when the value was reported etc.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"property_name": {
										Description:         "Required: The property name for which should be processed by external apps. This property should be present in the device model.",
										MarkdownDescription: "Required: The property name for which should be processed by external apps. This property should be present in the device model.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"data_topic": {
								Description:         "Topic used by mapper, all data collected from dataProperties should be published to this topic, the default value is $ke/events/device/+/data/update",
								MarkdownDescription: "Topic used by mapper, all data collected from dataProperties should be published to this topic, the default value is $ke/events/device/+/data/update",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"device_model_ref": {
						Description:         "Required: DeviceModelRef is reference to the device model used as a template to create the device instance.",
						MarkdownDescription: "Required: DeviceModelRef is reference to the device model used as a template to create the device instance.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_selector": {
						Description:         "NodeSelector indicates the binding preferences between devices and nodes. Refer to k8s.io/kubernetes/pkg/apis/core NodeSelector for more details",
						MarkdownDescription: "NodeSelector indicates the binding preferences between devices and nodes. Refer to k8s.io/kubernetes/pkg/apis/core NodeSelector for more details",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"node_selector_terms": {
								Description:         "Required. A list of node selector terms. The terms are ORed.",
								MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "A list of node selector requirements by node's labels.",
										MarkdownDescription: "A list of node selector requirements by node's labels.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The label key that the selector applies to.",
												MarkdownDescription: "The label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
												MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"values": {
												Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_fields": {
										Description:         "A list of node selector requirements by node's fields.",
										MarkdownDescription: "A list of node selector requirements by node's fields.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The label key that the selector applies to.",
												MarkdownDescription: "The label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
												MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"values": {
												Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"property_visitors": {
						Description:         "List of property visitors which describe how to access the device properties. PropertyVisitors must unique by propertyVisitor.propertyName.",
						MarkdownDescription: "List of property visitors which describe how to access the device properties. PropertyVisitors must unique by propertyVisitor.propertyName.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"bluetooth": {
								Description:         "Bluetooth represents a set of additional visitor config fields of bluetooth protocol.",
								MarkdownDescription: "Bluetooth represents a set of additional visitor config fields of bluetooth protocol.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"characteristic_uuid": {
										Description:         "Required: Unique ID of the corresponding operation",
										MarkdownDescription: "Required: Unique ID of the corresponding operation",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"data_converter": {
										Description:         "Responsible for converting the data being read from the bluetooth device into a form that is understandable by the platform",
										MarkdownDescription: "Responsible for converting the data being read from the bluetooth device into a form that is understandable by the platform",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"end_index": {
												Description:         "Required: Specifies the end index of incoming byte stream to be considered to convert the data the value specified should be inclusive for example if 3 is specified it includes the third index",
												MarkdownDescription: "Required: Specifies the end index of incoming byte stream to be considered to convert the data the value specified should be inclusive for example if 3 is specified it includes the third index",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"order_of_operations": {
												Description:         "Specifies in what order the operations(which are required to be performed to convert incoming data into understandable form) are performed",
												MarkdownDescription: "Specifies in what order the operations(which are required to be performed to convert incoming data into understandable form) are performed",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"operation_type": {
														Description:         "Required: Specifies the operation to be performed to convert incoming data",
														MarkdownDescription: "Required: Specifies the operation to be performed to convert incoming data",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"operation_value": {
														Description:         "Required: Specifies with what value the operation is to be performed",
														MarkdownDescription: "Required: Specifies with what value the operation is to be performed",

														Type: utilities.DynamicNumberType{},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"shift_left": {
												Description:         "Refers to the number of bits to shift left, if left-shift operation is necessary for conversion",
												MarkdownDescription: "Refers to the number of bits to shift left, if left-shift operation is necessary for conversion",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"shift_right": {
												Description:         "Refers to the number of bits to shift right, if right-shift operation is necessary for conversion",
												MarkdownDescription: "Refers to the number of bits to shift right, if right-shift operation is necessary for conversion",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"start_index": {
												Description:         "Required: Specifies the start index of the incoming byte stream to be considered to convert the data. For example: start-index:2, end-index:3 concatenates the value present at second and third index of the incoming byte stream. If we want to reverse the order we can give it as start-index:3, end-index:2",
												MarkdownDescription: "Required: Specifies the start index of the incoming byte stream to be considered to convert the data. For example: start-index:2, end-index:3 concatenates the value present at second and third index of the incoming byte stream. If we want to reverse the order we can give it as start-index:3, end-index:2",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"data_write": {
										Description:         "Responsible for converting the data coming from the platform into a form that is understood by the bluetooth device For example: 'ON':[1], 'OFF':[0]",
										MarkdownDescription: "Responsible for converting the data coming from the platform into a form that is understood by the bluetooth device For example: 'ON':[1], 'OFF':[0]",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"collect_cycle": {
								Description:         "Define how frequent mapper will collect from device.",
								MarkdownDescription: "Define how frequent mapper will collect from device.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"customized_protocol": {
								Description:         "CustomizedProtocol represents a set of visitor config fields of bluetooth protocol.",
								MarkdownDescription: "CustomizedProtocol represents a set of visitor config fields of bluetooth protocol.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config_data": {
										Description:         "Required: The configData of customized protocol",
										MarkdownDescription: "Required: The configData of customized protocol",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"protocol_name": {
										Description:         "Required: name of customized protocol",
										MarkdownDescription: "Required: name of customized protocol",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"customized_values": {
								Description:         "Customized values for visitor of provided protocols",
								MarkdownDescription: "Customized values for visitor of provided protocols",

								Type: utilities.DynamicType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"modbus": {
								Description:         "Modbus represents a set of additional visitor config fields of modbus protocol.",
								MarkdownDescription: "Modbus represents a set of additional visitor config fields of modbus protocol.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"is_register_swap": {
										Description:         "Indicates whether the high and low register swapped. Defaults to false.",
										MarkdownDescription: "Indicates whether the high and low register swapped. Defaults to false.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"is_swap": {
										Description:         "Indicates whether the high and low byte swapped. Defaults to false.",
										MarkdownDescription: "Indicates whether the high and low byte swapped. Defaults to false.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"limit": {
										Description:         "Required: Limit number of registers to read/write.",
										MarkdownDescription: "Required: Limit number of registers to read/write.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"offset": {
										Description:         "Required: Offset indicates the starting register number to read/write data.",
										MarkdownDescription: "Required: Offset indicates the starting register number to read/write data.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"register": {
										Description:         "Required: Type of register",
										MarkdownDescription: "Required: Type of register",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("CoilRegister", "DiscreteInputRegister", "InputRegister", "HoldingRegister"),
										},
									},

									"scale": {
										Description:         "The scale to convert raw property data into final units. Defaults to 1.0",
										MarkdownDescription: "The scale to convert raw property data into final units. Defaults to 1.0",

										Type: utilities.DynamicNumberType{},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"opcua": {
								Description:         "Opcua represents a set of additional visitor config fields of opc-ua protocol.",
								MarkdownDescription: "Opcua represents a set of additional visitor config fields of opc-ua protocol.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"browse_name": {
										Description:         "The name of opc-ua node",
										MarkdownDescription: "The name of opc-ua node",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"node_id": {
										Description:         "Required: The ID of opc-ua node, e.g. 'ns=1,i=1005'",
										MarkdownDescription: "Required: The ID of opc-ua node, e.g. 'ns=1,i=1005'",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"property_name": {
								Description:         "Required: The device property name to be accessed. This should refer to one of the device properties defined in the device model.",
								MarkdownDescription: "Required: The device property name to be accessed. This should refer to one of the device properties defined in the device model.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"report_cycle": {
								Description:         "Define how frequent mapper will report the value.",
								MarkdownDescription: "Define how frequent mapper will report the value.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"protocol": {
						Description:         "Required: The protocol configuration used to connect to the device.",
						MarkdownDescription: "Required: The protocol configuration used to connect to the device.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"bluetooth": {
								Description:         "Protocol configuration for bluetooth",
								MarkdownDescription: "Protocol configuration for bluetooth",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"mac_address": {
										Description:         "Unique identifier assigned to the device.",
										MarkdownDescription: "Unique identifier assigned to the device.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"common": {
								Description:         "Configuration for protocol common part",
								MarkdownDescription: "Configuration for protocol common part",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"collect_retry_times": {
										Description:         "Define retry times of mapper will collect from device.",
										MarkdownDescription: "Define retry times of mapper will collect from device.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"collect_timeout": {
										Description:         "Define timeout of mapper collect from device.",
										MarkdownDescription: "Define timeout of mapper collect from device.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"collect_type": {
										Description:         "Define collect type, sync or async.",
										MarkdownDescription: "Define collect type, sync or async.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("sync", "async"),
										},
									},

									"com": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"baud_rate": {
												Description:         "Required. BaudRate 115200|57600|38400|19200|9600|4800|2400|1800|1200|600|300|200|150|134|110|75|50",
												MarkdownDescription: "Required. BaudRate 115200|57600|38400|19200|9600|4800|2400|1800|1200|600|300|200|150|134|110|75|50",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.OneOf(115200, 57600, 38400, 19200, 9600, 4800, 2400, 1800, 1200, 600, 300, 200, 150, 134, 110, 75, 50),
												},
											},

											"data_bits": {
												Description:         "Required. Valid values are 8, 7, 6, 5.",
												MarkdownDescription: "Required. Valid values are 8, 7, 6, 5.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.OneOf(8, 7, 6, 5),
												},
											},

											"parity": {
												Description:         "Required. Valid options are 'none', 'even', 'odd'. Defaults to 'none'.",
												MarkdownDescription: "Required. Valid options are 'none', 'even', 'odd'. Defaults to 'none'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("none", "even", "odd"),
												},
											},

											"serial_port": {
												Description:         "Required.",
												MarkdownDescription: "Required.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"stop_bits": {
												Description:         "Required. Bit that stops 1|2",
												MarkdownDescription: "Required. Bit that stops 1|2",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.OneOf(1, 2),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"comm_type": {
										Description:         "Communication type, like tcp client, tcp server or COM",
										MarkdownDescription: "Communication type, like tcp client, tcp server or COM",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"customized_values": {
										Description:         "Customized values for provided protocol",
										MarkdownDescription: "Customized values for provided protocol",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"reconn_retry_times": {
										Description:         "Reconnecting retry times",
										MarkdownDescription: "Reconnecting retry times",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"reconn_timeout": {
										Description:         "Reconnection timeout",
										MarkdownDescription: "Reconnection timeout",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tcp": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ip": {
												Description:         "Required.",
												MarkdownDescription: "Required.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "Required.",
												MarkdownDescription: "Required.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"customized_protocol": {
								Description:         "Configuration for customized protocol",
								MarkdownDescription: "Configuration for customized protocol",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config_data": {
										Description:         "Any config data",
										MarkdownDescription: "Any config data",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"protocol_name": {
										Description:         "Unique protocol name Required.",
										MarkdownDescription: "Unique protocol name Required.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"modbus": {
								Description:         "Protocol configuration for modbus",
								MarkdownDescription: "Protocol configuration for modbus",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"slave_id": {
										Description:         "Required. 0-255",
										MarkdownDescription: "Required. 0-255",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"opcua": {
								Description:         "Protocol configuration for opc-ua",
								MarkdownDescription: "Protocol configuration for opc-ua",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"certificate": {
										Description:         "Certificate for access opc server.",
										MarkdownDescription: "Certificate for access opc server.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"password": {
										Description:         "Password for access opc server.",
										MarkdownDescription: "Password for access opc server.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"private_key": {
										Description:         "PrivateKey for access opc server.",
										MarkdownDescription: "PrivateKey for access opc server.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"security_mode": {
										Description:         "Defaults to 'none'.",
										MarkdownDescription: "Defaults to 'none'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"security_policy": {
										Description:         "Defaults to 'none'.",
										MarkdownDescription: "Defaults to 'none'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"timeout": {
										Description:         "Timeout seconds for the opc server connection.???",
										MarkdownDescription: "Timeout seconds for the opc server connection.???",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"url": {
										Description:         "Required: The URL for opc server endpoint.",
										MarkdownDescription: "Required: The URL for opc server endpoint.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"user_name": {
										Description:         "Username for access opc server.",
										MarkdownDescription: "Username for access opc server.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *DevicesKubeedgeIoDeviceV1Alpha2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_devices_kubeedge_io_device_v1alpha2")

	var state DevicesKubeedgeIoDeviceV1Alpha2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel DevicesKubeedgeIoDeviceV1Alpha2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("devices.kubeedge.io/v1alpha2")
	goModel.Kind = utilities.Ptr("Device")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *DevicesKubeedgeIoDeviceV1Alpha2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_devices_kubeedge_io_device_v1alpha2")
	// NO-OP: All data is already in Terraform state
}

func (r *DevicesKubeedgeIoDeviceV1Alpha2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_devices_kubeedge_io_device_v1alpha2")

	var state DevicesKubeedgeIoDeviceV1Alpha2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel DevicesKubeedgeIoDeviceV1Alpha2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("devices.kubeedge.io/v1alpha2")
	goModel.Kind = utilities.Ptr("Device")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *DevicesKubeedgeIoDeviceV1Alpha2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_devices_kubeedge_io_device_v1alpha2")
	// NO-OP: Terraform removes the state automatically for us
}
