/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package tests_testkube_io_v1

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
	_ datasource.DataSource = &TestsTestkubeIoTestTriggerV1Manifest{}
)

func NewTestsTestkubeIoTestTriggerV1Manifest() datasource.DataSource {
	return &TestsTestkubeIoTestTriggerV1Manifest{}
}

type TestsTestkubeIoTestTriggerV1Manifest struct{}

type TestsTestkubeIoTestTriggerV1ManifestData struct {
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
		Action            *string `tfsdk:"action" json:"action,omitempty"`
		ConcurrencyPolicy *string `tfsdk:"concurrency_policy" json:"concurrencyPolicy,omitempty"`
		ConditionSpec     *struct {
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
		Disabled  *bool   `tfsdk:"disabled" json:"disabled,omitempty"`
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
			NameRegex *string `tfsdk:"name_regex" json:"nameRegex,omitempty"`
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
			NameRegex *string `tfsdk:"name_regex" json:"nameRegex,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"test_selector" json:"testSelector,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TestsTestkubeIoTestTriggerV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_tests_testkube_io_test_trigger_v1_manifest"
}

func (r *TestsTestkubeIoTestTriggerV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TestTrigger is the Schema for the testtriggers API",
		MarkdownDescription: "TestTrigger is the Schema for the testtriggers API",
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
				Description:         "TestTriggerSpec defines the desired state of TestTrigger",
				MarkdownDescription: "TestTriggerSpec defines the desired state of TestTrigger",
				Attributes: map[string]schema.Attribute{
					"action": schema.StringAttribute{
						Description:         "Action represents what needs to be executed for selected Execution",
						MarkdownDescription: "Action represents what needs to be executed for selected Execution",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("run"),
						},
					},

					"concurrency_policy": schema.StringAttribute{
						Description:         "ConcurrencyPolicy defines concurrency policy for selected Execution",
						MarkdownDescription: "ConcurrencyPolicy defines concurrency policy for selected Execution",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("allow", "forbid", "replace"),
						},
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
											Optional:            true,
											Computed:            false,
										},

										"status": schema.StringAttribute{
											Description:         "TestTriggerConditionStatuses defines condition statuses for test triggers",
											MarkdownDescription: "TestTriggerConditionStatuses defines condition statuses for test triggers",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("True", "False", "Unknown"),
											},
										},

										"ttl": schema.Int64Attribute{
											Description:         "duration in seconds in the past from current time when the condition is still valid",
											MarkdownDescription: "duration in seconds in the past from current time when the condition is still valid",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"type": schema.StringAttribute{
											Description:         "test trigger condition",
											MarkdownDescription: "test trigger condition",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"delay": schema.Int64Attribute{
								Description:         "duration in seconds the test trigger waits between condition check",
								MarkdownDescription: "duration in seconds the test trigger waits between condition check",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"timeout": schema.Int64Attribute{
								Description:         "duration in seconds the test trigger waits for conditions, until its stopped",
								MarkdownDescription: "duration in seconds the test trigger waits for conditions, until its stopped",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"delay": schema.StringAttribute{
						Description:         "Delay is a duration string which specifies how long should the test be delayed after a trigger is matched",
						MarkdownDescription: "Delay is a duration string which specifies how long should the test be delayed after a trigger is matched",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disabled": schema.BoolAttribute{
						Description:         "whether test trigger is disabled",
						MarkdownDescription: "whether test trigger is disabled",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"event": schema.StringAttribute{
						Description:         "On which Event for a Resource should an Action be triggered",
						MarkdownDescription: "On which Event for a Resource should an Action be triggered",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("created", "modified", "deleted", "deployment-scale-update", "deployment-image-update", "deployment-env-update", "deployment-containers-modified", "event-start-test", "event-end-test-success", "event-end-test-failed", "event-end-test-aborted", "event-end-test-timeout", "event-start-testsuite", "event-end-testsuite-success", "event-end-testsuite-failed", "event-end-testsuite-aborted", "event-end-testsuite-timeout", "event-queue-testworkflow", "event-start-testworkflow", "event-end-testworkflow-success", "event-end-testworkflow-failed", "event-end-testworkflow-aborted", "event-created", "event-updated", "event-deleted"),
						},
					},

					"execution": schema.StringAttribute{
						Description:         "Execution identifies for which test execution should an Action be executed",
						MarkdownDescription: "Execution identifies for which test execution should an Action be executed",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("test", "testsuite", "testworkflow"),
						},
					},

					"probe_spec": schema.SingleNestedAttribute{
						Description:         "What resource probes should be matched",
						MarkdownDescription: "What resource probes should be matched",
						Attributes: map[string]schema.Attribute{
							"delay": schema.Int64Attribute{
								Description:         "duration in seconds the test trigger waits between probes",
								MarkdownDescription: "duration in seconds the test trigger waits between probes",
								Required:            false,
								Optional:            true,
								Computed:            false,
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
											Optional:            true,
											Computed:            false,
										},

										"host": schema.StringAttribute{
											Description:         "test trigger condition probe host, default is pod ip or service name",
											MarkdownDescription: "test trigger condition probe host, default is pod ip or service name",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "test trigger condition probe path to check, default is /",
											MarkdownDescription: "test trigger condition probe path to check, default is /",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "test trigger condition probe port to connect",
											MarkdownDescription: "test trigger condition probe port to connect",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"scheme": schema.StringAttribute{
											Description:         "test trigger condition probe scheme to connect to host, default is http",
											MarkdownDescription: "test trigger condition probe scheme to connect to host, default is http",
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

							"timeout": schema.Int64Attribute{
								Description:         "duration in seconds the test trigger waits for probes, until its stopped",
								MarkdownDescription: "duration in seconds the test trigger waits for probes, until its stopped",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"resource": schema.StringAttribute{
						Description:         "For which Resource do we monitor Event which triggers an Action on certain conditions",
						MarkdownDescription: "For which Resource do we monitor Event which triggers an Action on certain conditions",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("pod", "deployment", "statefulset", "daemonset", "service", "ingress", "event", "configmap"),
						},
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
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
													MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"values": schema.ListAttribute{
													Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
													MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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

									"match_labels": schema.MapAttribute{
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

							"name": schema.StringAttribute{
								Description:         "Name selector is used to identify a Kubernetes Object based on the metadata name",
								MarkdownDescription: "Name selector is used to identify a Kubernetes Object based on the metadata name",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name_regex": schema.StringAttribute{
								Description:         "kubernetes resource name regex",
								MarkdownDescription: "kubernetes resource name regex",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the Kubernetes object",
								MarkdownDescription: "Namespace of the Kubernetes object",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
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
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
													MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"values": schema.ListAttribute{
													Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
													MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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

									"match_labels": schema.MapAttribute{
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

							"name": schema.StringAttribute{
								Description:         "Name selector is used to identify a Kubernetes Object based on the metadata name",
								MarkdownDescription: "Name selector is used to identify a Kubernetes Object based on the metadata name",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name_regex": schema.StringAttribute{
								Description:         "kubernetes resource name regex",
								MarkdownDescription: "kubernetes resource name regex",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the Kubernetes object",
								MarkdownDescription: "Namespace of the Kubernetes object",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *TestsTestkubeIoTestTriggerV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_tests_testkube_io_test_trigger_v1_manifest")

	var model TestsTestkubeIoTestTriggerV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("tests.testkube.io/v1")
	model.Kind = pointer.String("TestTrigger")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
