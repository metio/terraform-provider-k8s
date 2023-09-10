/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package tests_testkube_io_v3

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
	_ datasource.DataSource              = &TestsTestkubeIoTestSuiteV3DataSource{}
	_ datasource.DataSourceWithConfigure = &TestsTestkubeIoTestSuiteV3DataSource{}
)

func NewTestsTestkubeIoTestSuiteV3DataSource() datasource.DataSource {
	return &TestsTestkubeIoTestSuiteV3DataSource{}
}

type TestsTestkubeIoTestSuiteV3DataSource struct {
	kubernetesClient dynamic.Interface
}

type TestsTestkubeIoTestSuiteV3DataSourceData struct {
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
		After *[]struct {
			Execute *[]struct {
				Delay *string `tfsdk:"delay" json:"delay,omitempty"`
				Test  *string `tfsdk:"test" json:"test,omitempty"`
			} `tfsdk:"execute" json:"execute,omitempty"`
			StopOnFailure *bool `tfsdk:"stop_on_failure" json:"stopOnFailure,omitempty"`
		} `tfsdk:"after" json:"after,omitempty"`
		Before *[]struct {
			Execute *[]struct {
				Delay *string `tfsdk:"delay" json:"delay,omitempty"`
				Test  *string `tfsdk:"test" json:"test,omitempty"`
			} `tfsdk:"execute" json:"execute,omitempty"`
			StopOnFailure *bool `tfsdk:"stop_on_failure" json:"stopOnFailure,omitempty"`
		} `tfsdk:"before" json:"before,omitempty"`
		Description      *string `tfsdk:"description" json:"description,omitempty"`
		ExecutionRequest *struct {
			CronJobTemplate          *string            `tfsdk:"cron_job_template" json:"cronJobTemplate,omitempty"`
			CronJobTemplateReference *string            `tfsdk:"cron_job_template_reference" json:"cronJobTemplateReference,omitempty"`
			ExecutionLabels          *map[string]string `tfsdk:"execution_labels" json:"executionLabels,omitempty"`
			HttpProxy                *string            `tfsdk:"http_proxy" json:"httpProxy,omitempty"`
			HttpsProxy               *string            `tfsdk:"https_proxy" json:"httpsProxy,omitempty"`
			JobTemplate              *string            `tfsdk:"job_template" json:"jobTemplate,omitempty"`
			JobTemplateReference     *string            `tfsdk:"job_template_reference" json:"jobTemplateReference,omitempty"`
			Labels                   *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Name                     *string            `tfsdk:"name" json:"name,omitempty"`
			Namespace                *string            `tfsdk:"namespace" json:"namespace,omitempty"`
			PvcTemplate              *string            `tfsdk:"pvc_template" json:"pvcTemplate,omitempty"`
			PvcTemplateReference     *string            `tfsdk:"pvc_template_reference" json:"pvcTemplateReference,omitempty"`
			ScraperTemplate          *string            `tfsdk:"scraper_template" json:"scraperTemplate,omitempty"`
			ScraperTemplateReference *string            `tfsdk:"scraper_template_reference" json:"scraperTemplateReference,omitempty"`
			SecretUUID               *string            `tfsdk:"secret_uuid" json:"secretUUID,omitempty"`
			Sync                     *bool              `tfsdk:"sync" json:"sync,omitempty"`
			Timeout                  *int64             `tfsdk:"timeout" json:"timeout,omitempty"`
			Variables                *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Type      *string `tfsdk:"type" json:"type,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					FieldRef *struct {
						ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
						FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
					} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
					ResourceFieldRef *struct {
						ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
						Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
						Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
					} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"variables" json:"variables,omitempty"`
		} `tfsdk:"execution_request" json:"executionRequest,omitempty"`
		Repeats  *int64  `tfsdk:"repeats" json:"repeats,omitempty"`
		Schedule *string `tfsdk:"schedule" json:"schedule,omitempty"`
		Steps    *[]struct {
			Execute *[]struct {
				Delay *string `tfsdk:"delay" json:"delay,omitempty"`
				Test  *string `tfsdk:"test" json:"test,omitempty"`
			} `tfsdk:"execute" json:"execute,omitempty"`
			StopOnFailure *bool `tfsdk:"stop_on_failure" json:"stopOnFailure,omitempty"`
		} `tfsdk:"steps" json:"steps,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TestsTestkubeIoTestSuiteV3DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_tests_testkube_io_test_suite_v3"
}

func (r *TestsTestkubeIoTestSuiteV3DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TestSuite is the Schema for the testsuites API",
		MarkdownDescription: "TestSuite is the Schema for the testsuites API",
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
				Description:         "TestSuiteSpec defines the desired state of TestSuite",
				MarkdownDescription: "TestSuiteSpec defines the desired state of TestSuite",
				Attributes: map[string]schema.Attribute{
					"after": schema.ListNestedAttribute{
						Description:         "After batch steps is list of batch tests which will be sequentially orchestrated for parallel tests in each batch",
						MarkdownDescription: "After batch steps is list of batch tests which will be sequentially orchestrated for parallel tests in each batch",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"execute": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"delay": schema.StringAttribute{
												Description:         "delay duration in time units",
												MarkdownDescription: "delay duration in time units",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"test": schema.StringAttribute{
												Description:         "object name",
												MarkdownDescription: "object name",
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

								"stop_on_failure": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"before": schema.ListNestedAttribute{
						Description:         "Before batch steps is list of batch tests which will be sequentially orchestrated for parallel tests in each batch",
						MarkdownDescription: "Before batch steps is list of batch tests which will be sequentially orchestrated for parallel tests in each batch",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"execute": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"delay": schema.StringAttribute{
												Description:         "delay duration in time units",
												MarkdownDescription: "delay duration in time units",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"test": schema.StringAttribute{
												Description:         "object name",
												MarkdownDescription: "object name",
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

								"stop_on_failure": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"description": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"execution_request": schema.SingleNestedAttribute{
						Description:         "test suite execution request body",
						MarkdownDescription: "test suite execution request body",
						Attributes: map[string]schema.Attribute{
							"cron_job_template": schema.StringAttribute{
								Description:         "cron job template extensions",
								MarkdownDescription: "cron job template extensions",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"cron_job_template_reference": schema.StringAttribute{
								Description:         "name of the template resource",
								MarkdownDescription: "name of the template resource",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"execution_labels": schema.MapAttribute{
								Description:         "execution labels",
								MarkdownDescription: "execution labels",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"http_proxy": schema.StringAttribute{
								Description:         "http proxy for executor containers",
								MarkdownDescription: "http proxy for executor containers",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"https_proxy": schema.StringAttribute{
								Description:         "https proxy for executor containers",
								MarkdownDescription: "https proxy for executor containers",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"job_template": schema.StringAttribute{
								Description:         "job template extensions",
								MarkdownDescription: "job template extensions",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"job_template_reference": schema.StringAttribute{
								Description:         "name of the template resource",
								MarkdownDescription: "name of the template resource",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"labels": schema.MapAttribute{
								Description:         "test suite labels",
								MarkdownDescription: "test suite labels",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "test execution custom name",
								MarkdownDescription: "test execution custom name",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespace": schema.StringAttribute{
								Description:         "test kubernetes namespace ('testkube' when not set)",
								MarkdownDescription: "test kubernetes namespace ('testkube' when not set)",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"pvc_template": schema.StringAttribute{
								Description:         "pvc template extensions",
								MarkdownDescription: "pvc template extensions",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"pvc_template_reference": schema.StringAttribute{
								Description:         "name of the template resource",
								MarkdownDescription: "name of the template resource",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"scraper_template": schema.StringAttribute{
								Description:         "scraper template extensions",
								MarkdownDescription: "scraper template extensions",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"scraper_template_reference": schema.StringAttribute{
								Description:         "name of the template resource",
								MarkdownDescription: "name of the template resource",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"secret_uuid": schema.StringAttribute{
								Description:         "secret uuid",
								MarkdownDescription: "secret uuid",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"sync": schema.BoolAttribute{
								Description:         "whether to start execution sync or async",
								MarkdownDescription: "whether to start execution sync or async",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"timeout": schema.Int64Attribute{
								Description:         "timeout for test suite execution",
								MarkdownDescription: "timeout for test suite execution",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"variables": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "variable name",
										MarkdownDescription: "variable name",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"type": schema.StringAttribute{
										Description:         "variable type",
										MarkdownDescription: "variable type",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"value": schema.StringAttribute{
										Description:         "variable string value",
										MarkdownDescription: "variable string value",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"value_from": schema.SingleNestedAttribute{
										Description:         "or load it from var source",
										MarkdownDescription: "or load it from var source",
										Attributes: map[string]schema.Attribute{
											"config_map_key_ref": schema.SingleNestedAttribute{
												Description:         "Selects a key of a ConfigMap.",
												MarkdownDescription: "Selects a key of a ConfigMap.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"optional": schema.BoolAttribute{
														Description:         "Specify whether the ConfigMap or its key must be defined",
														MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"field_ref": schema.SingleNestedAttribute{
												Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
												MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
														MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"field_path": schema.StringAttribute{
														Description:         "Path of the field to select in the specified API version.",
														MarkdownDescription: "Path of the field to select in the specified API version.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"resource_field_ref": schema.SingleNestedAttribute{
												Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
												MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
												Attributes: map[string]schema.Attribute{
													"container_name": schema.StringAttribute{
														Description:         "Container name: required for volumes, optional for env vars",
														MarkdownDescription: "Container name: required for volumes, optional for env vars",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"divisor": schema.StringAttribute{
														Description:         "Specifies the output format of the exposed resources, defaults to '1'",
														MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"resource": schema.StringAttribute{
														Description:         "Required: resource to select",
														MarkdownDescription: "Required: resource to select",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"secret_key_ref": schema.SingleNestedAttribute{
												Description:         "Selects a key of a secret in the pod's namespace",
												MarkdownDescription: "Selects a key of a secret in the pod's namespace",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"optional": schema.BoolAttribute{
														Description:         "Specify whether the Secret or its key must be defined",
														MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"repeats": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"schedule": schema.StringAttribute{
						Description:         "schedule in cron job format for scheduled test execution",
						MarkdownDescription: "schedule in cron job format for scheduled test execution",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"steps": schema.ListNestedAttribute{
						Description:         "Batch steps is list of batch tests which will be sequentially orchestrated for parallel tests in each batch",
						MarkdownDescription: "Batch steps is list of batch tests which will be sequentially orchestrated for parallel tests in each batch",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"execute": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"delay": schema.StringAttribute{
												Description:         "delay duration in time units",
												MarkdownDescription: "delay duration in time units",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"test": schema.StringAttribute{
												Description:         "object name",
												MarkdownDescription: "object name",
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

								"stop_on_failure": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
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
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *TestsTestkubeIoTestSuiteV3DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *TestsTestkubeIoTestSuiteV3DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_tests_testkube_io_test_suite_v3")

	var data TestsTestkubeIoTestSuiteV3DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "tests.testkube.io", Version: "v3", Resource: "TestSuite"}).
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

	var readResponse TestsTestkubeIoTestSuiteV3DataSourceData
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
	data.ApiVersion = pointer.String("tests.testkube.io/v3")
	data.Kind = pointer.String("TestSuite")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
