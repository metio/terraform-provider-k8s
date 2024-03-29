/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package bmc_tinkerbell_org_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &BmcTinkerbellOrgTaskV1Alpha1Manifest{}
)

func NewBmcTinkerbellOrgTaskV1Alpha1Manifest() datasource.DataSource {
	return &BmcTinkerbellOrgTaskV1Alpha1Manifest{}
}

type BmcTinkerbellOrgTaskV1Alpha1Manifest struct{}

type BmcTinkerbellOrgTaskV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		Connection *struct {
			AuthSecretRef *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"auth_secret_ref" json:"authSecretRef,omitempty"`
			Host            *string `tfsdk:"host" json:"host,omitempty"`
			InsecureTLS     *bool   `tfsdk:"insecure_tls" json:"insecureTLS,omitempty"`
			Port            *int64  `tfsdk:"port" json:"port,omitempty"`
			ProviderOptions *struct {
				IntelAMT *struct {
					Port *int64 `tfsdk:"port" json:"port,omitempty"`
				} `tfsdk:"intel_amt" json:"intelAMT,omitempty"`
				Ipmitool *struct {
					CipherSuite *string `tfsdk:"cipher_suite" json:"cipherSuite,omitempty"`
					Port        *int64  `tfsdk:"port" json:"port,omitempty"`
				} `tfsdk:"ipmitool" json:"ipmitool,omitempty"`
				Redfish *struct {
					Port *int64 `tfsdk:"port" json:"port,omitempty"`
				} `tfsdk:"redfish" json:"redfish,omitempty"`
				Rpc *struct {
					ConsumerURL  *string `tfsdk:"consumer_url" json:"consumerURL,omitempty"`
					Experimental *struct {
						CustomRequestPayload *string `tfsdk:"custom_request_payload" json:"customRequestPayload,omitempty"`
						DotPath              *string `tfsdk:"dot_path" json:"dotPath,omitempty"`
					} `tfsdk:"experimental" json:"experimental,omitempty"`
					Hmac *struct {
						PrefixSigDisabled *bool              `tfsdk:"prefix_sig_disabled" json:"prefixSigDisabled,omitempty"`
						Secrets           *map[string]string `tfsdk:"secrets" json:"secrets,omitempty"`
					} `tfsdk:"hmac" json:"hmac,omitempty"`
					LogNotificationsDisabled *bool `tfsdk:"log_notifications_disabled" json:"logNotificationsDisabled,omitempty"`
					Request                  *struct {
						HttpContentType *string              `tfsdk:"http_content_type" json:"httpContentType,omitempty"`
						HttpMethod      *string              `tfsdk:"http_method" json:"httpMethod,omitempty"`
						StaticHeaders   *map[string][]string `tfsdk:"static_headers" json:"staticHeaders,omitempty"`
						TimestampFormat *string              `tfsdk:"timestamp_format" json:"timestampFormat,omitempty"`
						TimestampHeader *string              `tfsdk:"timestamp_header" json:"timestampHeader,omitempty"`
					} `tfsdk:"request" json:"request,omitempty"`
					Signature *struct {
						AppendAlgoToHeaderDisabled *bool     `tfsdk:"append_algo_to_header_disabled" json:"appendAlgoToHeaderDisabled,omitempty"`
						HeaderName                 *string   `tfsdk:"header_name" json:"headerName,omitempty"`
						IncludedPayloadHeaders     *[]string `tfsdk:"included_payload_headers" json:"includedPayloadHeaders,omitempty"`
					} `tfsdk:"signature" json:"signature,omitempty"`
				} `tfsdk:"rpc" json:"rpc,omitempty"`
			} `tfsdk:"provider_options" json:"providerOptions,omitempty"`
		} `tfsdk:"connection" json:"connection,omitempty"`
		Task *struct {
			OneTimeBootDeviceAction *struct {
				Device  *[]string `tfsdk:"device" json:"device,omitempty"`
				EfiBoot *bool     `tfsdk:"efi_boot" json:"efiBoot,omitempty"`
			} `tfsdk:"one_time_boot_device_action" json:"oneTimeBootDeviceAction,omitempty"`
			PowerAction        *string `tfsdk:"power_action" json:"powerAction,omitempty"`
			VirtualMediaAction *struct {
				Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
				MediaURL *string `tfsdk:"media_url" json:"mediaURL,omitempty"`
			} `tfsdk:"virtual_media_action" json:"virtualMediaAction,omitempty"`
		} `tfsdk:"task" json:"task,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *BmcTinkerbellOrgTaskV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_bmc_tinkerbell_org_task_v1alpha1_manifest"
}

func (r *BmcTinkerbellOrgTaskV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Task is the Schema for the Task API.",
		MarkdownDescription: "Task is the Schema for the Task API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "TaskSpec defines the desired state of Task.",
				MarkdownDescription: "TaskSpec defines the desired state of Task.",
				Attributes: map[string]schema.Attribute{
					"connection": schema.SingleNestedAttribute{
						Description:         "Connection represents the Machine connectivity information.",
						MarkdownDescription: "Connection represents the Machine connectivity information.",
						Attributes: map[string]schema.Attribute{
							"auth_secret_ref": schema.SingleNestedAttribute{
								Description:         "AuthSecretRef is the SecretReference that contains authentication information of the Machine. The Secret must contain username and password keys. This is optional as it is not required when using the RPC provider.",
								MarkdownDescription: "AuthSecretRef is the SecretReference that contains authentication information of the Machine. The Secret must contain username and password keys. This is optional as it is not required when using the RPC provider.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "name is unique within a namespace to reference a secret resource.",
										MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "namespace defines the space within which the secret name must be unique.",
										MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"host": schema.StringAttribute{
								Description:         "Host is the host IP address or hostname of the Machine.",
								MarkdownDescription: "Host is the host IP address or hostname of the Machine.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"insecure_tls": schema.BoolAttribute{
								Description:         "InsecureTLS specifies trusted TLS connections.",
								MarkdownDescription: "InsecureTLS specifies trusted TLS connections.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "Port is the port number for connecting with the Machine.",
								MarkdownDescription: "Port is the port number for connecting with the Machine.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"provider_options": schema.SingleNestedAttribute{
								Description:         "ProviderOptions contains provider specific options.",
								MarkdownDescription: "ProviderOptions contains provider specific options.",
								Attributes: map[string]schema.Attribute{
									"intel_amt": schema.SingleNestedAttribute{
										Description:         "IntelAMT contains the options to customize the IntelAMT provider.",
										MarkdownDescription: "IntelAMT contains the options to customize the IntelAMT provider.",
										Attributes: map[string]schema.Attribute{
											"port": schema.Int64Attribute{
												Description:         "Port that intelAMT will use for calls.",
												MarkdownDescription: "Port that intelAMT will use for calls.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"ipmitool": schema.SingleNestedAttribute{
										Description:         "IPMITOOL contains the options to customize the Ipmitool provider.",
										MarkdownDescription: "IPMITOOL contains the options to customize the Ipmitool provider.",
										Attributes: map[string]schema.Attribute{
											"cipher_suite": schema.StringAttribute{
												Description:         "CipherSuite that ipmitool will use for calls.",
												MarkdownDescription: "CipherSuite that ipmitool will use for calls.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "Port that ipmitool will use for calls.",
												MarkdownDescription: "Port that ipmitool will use for calls.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"redfish": schema.SingleNestedAttribute{
										Description:         "Redfish contains the options to customize the Redfish provider.",
										MarkdownDescription: "Redfish contains the options to customize the Redfish provider.",
										Attributes: map[string]schema.Attribute{
											"port": schema.Int64Attribute{
												Description:         "Port that redfish will use for calls.",
												MarkdownDescription: "Port that redfish will use for calls.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"rpc": schema.SingleNestedAttribute{
										Description:         "RPC contains the options to customize the RPC provider.",
										MarkdownDescription: "RPC contains the options to customize the RPC provider.",
										Attributes: map[string]schema.Attribute{
											"consumer_url": schema.StringAttribute{
												Description:         "ConsumerURL is the URL where an rpc consumer/listener is running and to which we will send and receive all notifications.",
												MarkdownDescription: "ConsumerURL is the URL where an rpc consumer/listener is running and to which we will send and receive all notifications.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"experimental": schema.SingleNestedAttribute{
												Description:         "Experimental options.",
												MarkdownDescription: "Experimental options.",
												Attributes: map[string]schema.Attribute{
													"custom_request_payload": schema.StringAttribute{
														Description:         "CustomRequestPayload must be in json.",
														MarkdownDescription: "CustomRequestPayload must be in json.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"dot_path": schema.StringAttribute{
														Description:         "DotPath is the path to the json object where the bmclib RequestPayload{} struct will be embedded. For example: object.data.body",
														MarkdownDescription: "DotPath is the path to the json object where the bmclib RequestPayload{} struct will be embedded. For example: object.data.body",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"hmac": schema.SingleNestedAttribute{
												Description:         "HMAC is the options used to create a HMAC signature.",
												MarkdownDescription: "HMAC is the options used to create a HMAC signature.",
												Attributes: map[string]schema.Attribute{
													"prefix_sig_disabled": schema.BoolAttribute{
														Description:         "PrefixSigDisabled determines whether the algorithm will be prefixed to the signature. Example: sha256=abc123",
														MarkdownDescription: "PrefixSigDisabled determines whether the algorithm will be prefixed to the signature. Example: sha256=abc123",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secrets": schema.MapAttribute{
														Description:         "Secrets are a map of algorithms to secrets used for signing.",
														MarkdownDescription: "Secrets are a map of algorithms to secrets used for signing.",
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

											"log_notifications_disabled": schema.BoolAttribute{
												Description:         "LogNotificationsDisabled determines whether responses from rpc consumer/listeners will be logged or not.",
												MarkdownDescription: "LogNotificationsDisabled determines whether responses from rpc consumer/listeners will be logged or not.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"request": schema.SingleNestedAttribute{
												Description:         "Request is the options used to create the rpc HTTP request.",
												MarkdownDescription: "Request is the options used to create the rpc HTTP request.",
												Attributes: map[string]schema.Attribute{
													"http_content_type": schema.StringAttribute{
														Description:         "HTTPContentType is the content type to use for the rpc request notification.",
														MarkdownDescription: "HTTPContentType is the content type to use for the rpc request notification.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"http_method": schema.StringAttribute{
														Description:         "HTTPMethod is the HTTP method to use for the rpc request notification.",
														MarkdownDescription: "HTTPMethod is the HTTP method to use for the rpc request notification.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"static_headers": schema.MapAttribute{
														Description:         "StaticHeaders are predefined headers that will be added to every request.",
														MarkdownDescription: "StaticHeaders are predefined headers that will be added to every request.",
														ElementType:         types.ListType{ElemType: types.StringType},
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timestamp_format": schema.StringAttribute{
														Description:         "TimestampFormat is the time format for the timestamp header.",
														MarkdownDescription: "TimestampFormat is the time format for the timestamp header.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timestamp_header": schema.StringAttribute{
														Description:         "TimestampHeader is the header name that should contain the timestamp. Example: X-BMCLIB-Timestamp",
														MarkdownDescription: "TimestampHeader is the header name that should contain the timestamp. Example: X-BMCLIB-Timestamp",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"signature": schema.SingleNestedAttribute{
												Description:         "Signature is the options used for adding an HMAC signature to an HTTP request.",
												MarkdownDescription: "Signature is the options used for adding an HMAC signature to an HTTP request.",
												Attributes: map[string]schema.Attribute{
													"append_algo_to_header_disabled": schema.BoolAttribute{
														Description:         "AppendAlgoToHeaderDisabled decides whether to append the algorithm to the signature header or not. Example: X-BMCLIB-Signature becomes X-BMCLIB-Signature-256 When set to true, a header will be added for each algorithm. Example: X-BMCLIB-Signature-256 and X-BMCLIB-Signature-512",
														MarkdownDescription: "AppendAlgoToHeaderDisabled decides whether to append the algorithm to the signature header or not. Example: X-BMCLIB-Signature becomes X-BMCLIB-Signature-256 When set to true, a header will be added for each algorithm. Example: X-BMCLIB-Signature-256 and X-BMCLIB-Signature-512",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"header_name": schema.StringAttribute{
														Description:         "HeaderName is the header name that should contain the signature(s). Example: X-BMCLIB-Signature",
														MarkdownDescription: "HeaderName is the header name that should contain the signature(s). Example: X-BMCLIB-Signature",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"included_payload_headers": schema.ListAttribute{
														Description:         "IncludedPayloadHeaders are headers whose values will be included in the signature payload. Example: X-BMCLIB-My-Custom-Header All headers will be deduplicated.",
														MarkdownDescription: "IncludedPayloadHeaders are headers whose values will be included in the signature payload. Example: X-BMCLIB-My-Custom-Header All headers will be deduplicated.",
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

					"task": schema.SingleNestedAttribute{
						Description:         "Task defines the specific action to be performed.",
						MarkdownDescription: "Task defines the specific action to be performed.",
						Attributes: map[string]schema.Attribute{
							"one_time_boot_device_action": schema.SingleNestedAttribute{
								Description:         "OneTimeBootDeviceAction represents a baseboard management one time set boot device operation.",
								MarkdownDescription: "OneTimeBootDeviceAction represents a baseboard management one time set boot device operation.",
								Attributes: map[string]schema.Attribute{
									"device": schema.ListAttribute{
										Description:         "Devices represents the boot devices, in order for setting one time boot. Currently only the first device in the slice is used to set one time boot.",
										MarkdownDescription: "Devices represents the boot devices, in order for setting one time boot. Currently only the first device in the slice is used to set one time boot.",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"efi_boot": schema.BoolAttribute{
										Description:         "EFIBoot instructs the machine to use EFI boot.",
										MarkdownDescription: "EFIBoot instructs the machine to use EFI boot.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"power_action": schema.StringAttribute{
								Description:         "PowerAction represents a baseboard management power operation.",
								MarkdownDescription: "PowerAction represents a baseboard management power operation.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("on", "off", "soft", "status", "cycle", "reset"),
								},
							},

							"virtual_media_action": schema.SingleNestedAttribute{
								Description:         "VirtualMediaAction represents a baseboard management virtual media insert/eject.",
								MarkdownDescription: "VirtualMediaAction represents a baseboard management virtual media insert/eject.",
								Attributes: map[string]schema.Attribute{
									"kind": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"media_url": schema.StringAttribute{
										Description:         "mediaURL represents the URL of the image to be inserted into the virtual media, or empty to eject media.",
										MarkdownDescription: "mediaURL represents the URL of the image to be inserted into the virtual media, or empty to eject media.",
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
						Required: true,
						Optional: false,
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

func (r *BmcTinkerbellOrgTaskV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_bmc_tinkerbell_org_task_v1alpha1_manifest")

	var model BmcTinkerbellOrgTaskV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("bmc.tinkerbell.org/v1alpha1")
	model.Kind = pointer.String("Task")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
