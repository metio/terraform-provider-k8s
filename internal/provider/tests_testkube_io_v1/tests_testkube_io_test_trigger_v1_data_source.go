/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package tests_testkube_io_v1

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
	_ datasource.DataSource              = &TestsTestkubeIoTestTriggerV1DataSource{}
	_ datasource.DataSourceWithConfigure = &TestsTestkubeIoTestTriggerV1DataSource{}
)

func NewTestsTestkubeIoTestTriggerV1DataSource() datasource.DataSource {
	return &TestsTestkubeIoTestTriggerV1DataSource{}
}

type TestsTestkubeIoTestTriggerV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type TestsTestkubeIoTestTriggerV1DataSourceData struct {
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
		Action        *string `tfsdk:"action" json:"action,omitempty"`
		ConditionSpec *struct {
			Conditions *[]struct {
				Reason *string `tfsdk:"reason" json:"reason,omitempty"`
				Status *string `tfsdk:"status" json:"status,omitempty"`
				Ttl    *int64  `tfsdk:"ttl" json:"ttl,omitempty"`
				Type   *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"conditions" json:"conditions,omitempty"`
			Delay   *int64 `tfsdk:"delay" json:"delay,omitempty"`
			Timeout *int64 `tfsdk:"timeout" json:"timeout,omitempty"`
		} `tfsdk:"condition_spec" json:"conditionSpec,omitempty"`
		Delay     *string `tfsdk:"delay" json:"delay,omitempty"`
		Event     *string `tfsdk:"event" json:"event,omitempty"`
		Execution *string `tfsdk:"execution" json:"execution,omitempty"`
		ProbeSpec *struct {
			Delay  *int64 `tfsdk:"delay" json:"delay,omitempty"`
			Probes *[]struct {
				Headers *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
				Host    *string            `tfsdk:"host" json:"host,omitempty"`
				Path    *string            `tfsdk:"path" json:"path,omitempty"`
				Port    *int64             `tfsdk:"port" json:"port,omitempty"`
				Scheme  *string            `tfsdk:"scheme" json:"scheme,omitempty"`
			} `tfsdk:"probes" json:"probes,omitempty"`
			Timeout *int64 `tfsdk:"timeout" json:"timeout,omitempty"`
		} `tfsdk:"probe_spec" json:"probeSpec,omitempty"`
		Resource         *string `tfsdk:"resource" json:"resource,omitempty"`
		ResourceSelector *struct {
			LabelSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"resource_selector" json:"resourceSelector,omitempty"`
		TestSelector *struct {
			LabelSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"test_selector" json:"testSelector,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TestsTestkubeIoTestTriggerV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_tests_testkube_io_test_trigger_v1"
}

func (r *TestsTestkubeIoTestTriggerV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TestTrigger is the Schema for the testtriggers API",
		MarkdownDescription: "TestTrigger is the Schema for the testtriggers API",
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
				Description:         "TestTriggerSpec defines the desired state of TestTrigger",
				MarkdownDescription: "TestTriggerSpec defines the desired state of TestTrigger",
				Attributes: map[string]schema.Attribute{
					"action": schema.StringAttribute{
						Description:         "Action represents what needs to be executed for selected Execution",
						MarkdownDescription: "Action represents what needs to be executed for selected Execution",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"condition_spec": schema.SingleNestedAttribute{
						Description:         "What resource conditions should be matched",
						MarkdownDescription: "What resource conditions should be matched",
						Attributes: map[string]schema.Attribute{
							"conditions": schema.ListNestedAttribute{
								Description:         "list of test trigger conditions",
								MarkdownDescription: "list of test trigger conditions",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"reason": schema.StringAttribute{
											Description:         "test trigger condition reason",
											MarkdownDescription: "test trigger condition reason",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"status": schema.StringAttribute{
											Description:         "TestTriggerConditionStatuses defines condition statuses for test triggers",
											MarkdownDescription: "TestTriggerConditionStatuses defines condition statuses for test triggers",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"ttl": schema.Int64Attribute{
											Description:         "duration in seconds in the past from current time when the condition is still valid",
											MarkdownDescription: "duration in seconds in the past from current time when the condition is still valid",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"type": schema.StringAttribute{
											Description:         "test trigger condition",
											MarkdownDescription: "test trigger condition",
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

							"delay": schema.Int64Attribute{
								Description:         "duration in seconds the test trigger waits between condition check",
								MarkdownDescription: "duration in seconds the test trigger waits between condition check",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"timeout": schema.Int64Attribute{
								Description:         "duration in seconds the test trigger waits for conditions, until its stopped",
								MarkdownDescription: "duration in seconds the test trigger waits for conditions, until its stopped",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"delay": schema.StringAttribute{
						Description:         "Delay is a duration string which specifies how long should the test be delayed after a trigger is matched",
						MarkdownDescription: "Delay is a duration string which specifies how long should the test be delayed after a trigger is matched",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"event": schema.StringAttribute{
						Description:         "On which Event for a Resource should an Action be triggered",
						MarkdownDescription: "On which Event for a Resource should an Action be triggered",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"execution": schema.StringAttribute{
						Description:         "Execution identifies for which test execution should an Action be executed",
						MarkdownDescription: "Execution identifies for which test execution should an Action be executed",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"probe_spec": schema.SingleNestedAttribute{
						Description:         "What resource probes should be matched",
						MarkdownDescription: "What resource probes should be matched",
						Attributes: map[string]schema.Attribute{
							"delay": schema.Int64Attribute{
								Description:         "duration in seconds the test trigger waits between probes",
								MarkdownDescription: "duration in seconds the test trigger waits between probes",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"probes": schema.ListNestedAttribute{
								Description:         "list of test trigger probes",
								MarkdownDescription: "list of test trigger probes",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"headers": schema.MapAttribute{
											Description:         "test trigger condition probe headers to submit",
											MarkdownDescription: "test trigger condition probe headers to submit",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"host": schema.StringAttribute{
											Description:         "test trigger condition probe host, default is pod ip or service name",
											MarkdownDescription: "test trigger condition probe host, default is pod ip or service name",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"path": schema.StringAttribute{
											Description:         "test trigger condition probe path to check, default is /",
											MarkdownDescription: "test trigger condition probe path to check, default is /",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"port": schema.Int64Attribute{
											Description:         "test trigger condition probe port to connect",
											MarkdownDescription: "test trigger condition probe port to connect",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"scheme": schema.StringAttribute{
											Description:         "test trigger condition probe scheme to connect to host, default is http",
											MarkdownDescription: "test trigger condition probe scheme to connect to host, default is http",
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

							"timeout": schema.Int64Attribute{
								Description:         "duration in seconds the test trigger waits for probes, until its stopped",
								MarkdownDescription: "duration in seconds the test trigger waits for probes, until its stopped",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"resource": schema.StringAttribute{
						Description:         "For which Resource do we monitor Event which triggers an Action on certain conditions",
						MarkdownDescription: "For which Resource do we monitor Event which triggers an Action on certain conditions",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"resource_selector": schema.SingleNestedAttribute{
						Description:         "ResourceSelector identifies which Kubernetes Objects should be watched",
						MarkdownDescription: "ResourceSelector identifies which Kubernetes Objects should be watched",
						Attributes: map[string]schema.Attribute{
							"label_selector": schema.SingleNestedAttribute{
								Description:         "LabelSelector is used to identify a group of Kubernetes Objects based on their metadata labels",
								MarkdownDescription: "LabelSelector is used to identify a group of Kubernetes Objects based on their metadata labels",
								Attributes: map[string]schema.Attribute{
									"match_expressions": schema.ListNestedAttribute{
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "key is the label key that the selector applies to.",
													MarkdownDescription: "key is the label key that the selector applies to.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"operator": schema.StringAttribute{
													Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"values": schema.ListAttribute{
													Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

									"match_labels": schema.MapAttribute{
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"name": schema.StringAttribute{
								Description:         "Name selector is used to identify a Kubernetes Object based on the metadata name",
								MarkdownDescription: "Name selector is used to identify a Kubernetes Object based on the metadata name",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the Kubernetes object",
								MarkdownDescription: "Namespace of the Kubernetes object",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"test_selector": schema.SingleNestedAttribute{
						Description:         "TestSelector identifies on which Testkube Kubernetes Objects an Action should be taken",
						MarkdownDescription: "TestSelector identifies on which Testkube Kubernetes Objects an Action should be taken",
						Attributes: map[string]schema.Attribute{
							"label_selector": schema.SingleNestedAttribute{
								Description:         "LabelSelector is used to identify a group of Kubernetes Objects based on their metadata labels",
								MarkdownDescription: "LabelSelector is used to identify a group of Kubernetes Objects based on their metadata labels",
								Attributes: map[string]schema.Attribute{
									"match_expressions": schema.ListNestedAttribute{
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "key is the label key that the selector applies to.",
													MarkdownDescription: "key is the label key that the selector applies to.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"operator": schema.StringAttribute{
													Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"values": schema.ListAttribute{
													Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

									"match_labels": schema.MapAttribute{
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"name": schema.StringAttribute{
								Description:         "Name selector is used to identify a Kubernetes Object based on the metadata name",
								MarkdownDescription: "Name selector is used to identify a Kubernetes Object based on the metadata name",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the Kubernetes object",
								MarkdownDescription: "Namespace of the Kubernetes object",
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
	}
}

func (r *TestsTestkubeIoTestTriggerV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *TestsTestkubeIoTestTriggerV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_tests_testkube_io_test_trigger_v1")

	var data TestsTestkubeIoTestTriggerV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "tests.testkube.io", Version: "v1", Resource: "TestTrigger"}).
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

	var readResponse TestsTestkubeIoTestTriggerV1DataSourceData
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
	data.ApiVersion = pointer.String("tests.testkube.io/v1")
	data.Kind = pointer.String("TestTrigger")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
