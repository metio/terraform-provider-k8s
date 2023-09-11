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
	_ datasource.DataSource              = &TestsTestkubeIoTestExecutionV1DataSource{}
	_ datasource.DataSourceWithConfigure = &TestsTestkubeIoTestExecutionV1DataSource{}
)

func NewTestsTestkubeIoTestExecutionV1DataSource() datasource.DataSource {
	return &TestsTestkubeIoTestExecutionV1DataSource{}
}

type TestsTestkubeIoTestExecutionV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type TestsTestkubeIoTestExecutionV1DataSourceData struct {
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
		ExecutionRequest *struct {
			ActiveDeadlineSeconds *int64    `tfsdk:"active_deadline_seconds" json:"activeDeadlineSeconds,omitempty"`
			Args                  *[]string `tfsdk:"args" json:"args,omitempty"`
			ArgsMode              *string   `tfsdk:"args_mode" json:"argsMode,omitempty"`
			ArtifactRequest       *struct {
				Dirs                   *[]string `tfsdk:"dirs" json:"dirs,omitempty"`
				OmitFolderPerExecution *bool     `tfsdk:"omit_folder_per_execution" json:"omitFolderPerExecution,omitempty"`
				StorageBucket          *string   `tfsdk:"storage_bucket" json:"storageBucket,omitempty"`
				StorageClassName       *string   `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
				VolumeMountPath        *string   `tfsdk:"volume_mount_path" json:"volumeMountPath,omitempty"`
			} `tfsdk:"artifact_request" json:"artifactRequest,omitempty"`
			Command         *[]string `tfsdk:"command" json:"command,omitempty"`
			CronJobTemplate *string   `tfsdk:"cron_job_template" json:"cronJobTemplate,omitempty"`
			EnvConfigMaps   *[]struct {
				MapToVariables *bool   `tfsdk:"map_to_variables" json:"mapToVariables,omitempty"`
				Mount          *bool   `tfsdk:"mount" json:"mount,omitempty"`
				MountPath      *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				Reference      *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"reference" json:"reference,omitempty"`
			} `tfsdk:"env_config_maps" json:"envConfigMaps,omitempty"`
			EnvSecrets *[]struct {
				MapToVariables *bool   `tfsdk:"map_to_variables" json:"mapToVariables,omitempty"`
				Mount          *bool   `tfsdk:"mount" json:"mount,omitempty"`
				MountPath      *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				Reference      *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"reference" json:"reference,omitempty"`
			} `tfsdk:"env_secrets" json:"envSecrets,omitempty"`
			Envs             *map[string]string `tfsdk:"envs" json:"envs,omitempty"`
			ExecutionLabels  *map[string]string `tfsdk:"execution_labels" json:"executionLabels,omitempty"`
			HttpProxy        *string            `tfsdk:"http_proxy" json:"httpProxy,omitempty"`
			HttpsProxy       *string            `tfsdk:"https_proxy" json:"httpsProxy,omitempty"`
			Image            *string            `tfsdk:"image" json:"image,omitempty"`
			ImagePullSecrets *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
			IsVariablesFileUploaded *bool   `tfsdk:"is_variables_file_uploaded" json:"isVariablesFileUploaded,omitempty"`
			JobTemplate             *string `tfsdk:"job_template" json:"jobTemplate,omitempty"`
			Name                    *string `tfsdk:"name" json:"name,omitempty"`
			Namespace               *string `tfsdk:"namespace" json:"namespace,omitempty"`
			NegativeTest            *bool   `tfsdk:"negative_test" json:"negativeTest,omitempty"`
			Number                  *int64  `tfsdk:"number" json:"number,omitempty"`
			PostRunScript           *string `tfsdk:"post_run_script" json:"postRunScript,omitempty"`
			PreRunScript            *string `tfsdk:"pre_run_script" json:"preRunScript,omitempty"`
			RunningContext          *struct {
				Context *string `tfsdk:"context" json:"context,omitempty"`
				Type    *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"running_context" json:"runningContext,omitempty"`
			ScraperTemplate     *string            `tfsdk:"scraper_template" json:"scraperTemplate,omitempty"`
			SecretEnvs          *map[string]string `tfsdk:"secret_envs" json:"secretEnvs,omitempty"`
			Sync                *bool              `tfsdk:"sync" json:"sync,omitempty"`
			TestSecretUUID      *string            `tfsdk:"test_secret_uuid" json:"testSecretUUID,omitempty"`
			TestSuiteName       *string            `tfsdk:"test_suite_name" json:"testSuiteName,omitempty"`
			TestSuiteSecretUUID *string            `tfsdk:"test_suite_secret_uuid" json:"testSuiteSecretUUID,omitempty"`
			Variables           *struct {
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
			VariablesFile *string `tfsdk:"variables_file" json:"variablesFile,omitempty"`
		} `tfsdk:"execution_request" json:"executionRequest,omitempty"`
		Test *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"test" json:"test,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TestsTestkubeIoTestExecutionV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_tests_testkube_io_test_execution_v1"
}

func (r *TestsTestkubeIoTestExecutionV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TestExecution is the Schema for the testexecutions API",
		MarkdownDescription: "TestExecution is the Schema for the testexecutions API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "TestExecutionSpec defines the desired state of TestExecution",
				MarkdownDescription: "TestExecutionSpec defines the desired state of TestExecution",
				Attributes: map[string]schema.Attribute{
					"execution_request": schema.SingleNestedAttribute{
						Description:         "test execution request body",
						MarkdownDescription: "test execution request body",
						Attributes: map[string]schema.Attribute{
							"active_deadline_seconds": schema.Int64Attribute{
								Description:         "Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.",
								MarkdownDescription: "Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"args": schema.ListAttribute{
								Description:         "additional executor binary arguments",
								MarkdownDescription: "additional executor binary arguments",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"args_mode": schema.StringAttribute{
								Description:         "usage mode for arguments",
								MarkdownDescription: "usage mode for arguments",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"artifact_request": schema.SingleNestedAttribute{
								Description:         "artifact request body with test artifacts",
								MarkdownDescription: "artifact request body with test artifacts",
								Attributes: map[string]schema.Attribute{
									"dirs": schema.ListAttribute{
										Description:         "artifact directories for scraping",
										MarkdownDescription: "artifact directories for scraping",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"omit_folder_per_execution": schema.BoolAttribute{
										Description:         "don't use a separate folder for execution artifacts",
										MarkdownDescription: "don't use a separate folder for execution artifacts",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"storage_bucket": schema.StringAttribute{
										Description:         "artifact bucket storage",
										MarkdownDescription: "artifact bucket storage",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"storage_class_name": schema.StringAttribute{
										Description:         "artifact storage class name for container executor",
										MarkdownDescription: "artifact storage class name for container executor",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"volume_mount_path": schema.StringAttribute{
										Description:         "artifact volume mount path for container executor",
										MarkdownDescription: "artifact volume mount path for container executor",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"command": schema.ListAttribute{
								Description:         "executor binary command",
								MarkdownDescription: "executor binary command",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"cron_job_template": schema.StringAttribute{
								Description:         "cron job template extensions",
								MarkdownDescription: "cron job template extensions",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"env_config_maps": schema.ListNestedAttribute{
								Description:         "config map references",
								MarkdownDescription: "config map references",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"map_to_variables": schema.BoolAttribute{
											Description:         "whether we shoud map to variables from resource",
											MarkdownDescription: "whether we shoud map to variables from resource",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mount": schema.BoolAttribute{
											Description:         "whether we shoud mount resource",
											MarkdownDescription: "whether we shoud mount resource",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mount_path": schema.StringAttribute{
											Description:         "where we shoud mount resource",
											MarkdownDescription: "where we shoud mount resource",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"reference": schema.SingleNestedAttribute{
											Description:         "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
											MarkdownDescription: "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"env_secrets": schema.ListNestedAttribute{
								Description:         "secret references",
								MarkdownDescription: "secret references",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"map_to_variables": schema.BoolAttribute{
											Description:         "whether we shoud map to variables from resource",
											MarkdownDescription: "whether we shoud map to variables from resource",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mount": schema.BoolAttribute{
											Description:         "whether we shoud mount resource",
											MarkdownDescription: "whether we shoud mount resource",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mount_path": schema.StringAttribute{
											Description:         "where we shoud mount resource",
											MarkdownDescription: "where we shoud mount resource",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"reference": schema.SingleNestedAttribute{
											Description:         "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
											MarkdownDescription: "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"envs": schema.MapAttribute{
								Description:         "Environment variables passed to executor. Deprecated: use Basic Variables instead",
								MarkdownDescription: "Environment variables passed to executor. Deprecated: use Basic Variables instead",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"execution_labels": schema.MapAttribute{
								Description:         "test execution labels",
								MarkdownDescription: "test execution labels",
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

							"image": schema.StringAttribute{
								Description:         "container executor image",
								MarkdownDescription: "container executor image",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"image_pull_secrets": schema.ListNestedAttribute{
								Description:         "container executor image pull secrets",
								MarkdownDescription: "container executor image pull secrets",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
											MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

							"is_variables_file_uploaded": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
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

							"negative_test": schema.BoolAttribute{
								Description:         "negative test will fail the execution if it is a success and it will succeed if it is a failure",
								MarkdownDescription: "negative test will fail the execution if it is a success and it will succeed if it is a failure",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"number": schema.Int64Attribute{
								Description:         "test execution number",
								MarkdownDescription: "test execution number",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"post_run_script": schema.StringAttribute{
								Description:         "script to run after test execution",
								MarkdownDescription: "script to run after test execution",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"pre_run_script": schema.StringAttribute{
								Description:         "script to run before test execution",
								MarkdownDescription: "script to run before test execution",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"running_context": schema.SingleNestedAttribute{
								Description:         "running context for test or test suite execution",
								MarkdownDescription: "running context for test or test suite execution",
								Attributes: map[string]schema.Attribute{
									"context": schema.StringAttribute{
										Description:         "Context value depending from its type",
										MarkdownDescription: "Context value depending from its type",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"type": schema.StringAttribute{
										Description:         "One of possible context types",
										MarkdownDescription: "One of possible context types",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"scraper_template": schema.StringAttribute{
								Description:         "scraper template extensions",
								MarkdownDescription: "scraper template extensions",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"secret_envs": schema.MapAttribute{
								Description:         "Execution variables passed to executor from secrets. Deprecated: use Secret Variables instead",
								MarkdownDescription: "Execution variables passed to executor from secrets. Deprecated: use Secret Variables instead",
								ElementType:         types.StringType,
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

							"test_secret_uuid": schema.StringAttribute{
								Description:         "test secret uuid",
								MarkdownDescription: "test secret uuid",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"test_suite_name": schema.StringAttribute{
								Description:         "unique test suite name (CRD Test suite name), if it's run as a part of test suite",
								MarkdownDescription: "unique test suite name (CRD Test suite name), if it's run as a part of test suite",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"test_suite_secret_uuid": schema.StringAttribute{
								Description:         "test suite secret uuid, if it's run as a part of test suite",
								MarkdownDescription: "test suite secret uuid, if it's run as a part of test suite",
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

							"variables_file": schema.StringAttribute{
								Description:         "variables file content - need to be in format for particular executor (e.g. postman envs file)",
								MarkdownDescription: "variables file content - need to be in format for particular executor (e.g. postman envs file)",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"test": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "object name",
								MarkdownDescription: "object name",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespace": schema.StringAttribute{
								Description:         "object kubernetes namespace",
								MarkdownDescription: "object kubernetes namespace",
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

func (r *TestsTestkubeIoTestExecutionV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *TestsTestkubeIoTestExecutionV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_tests_testkube_io_test_execution_v1")

	var data TestsTestkubeIoTestExecutionV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "tests.testkube.io", Version: "v1", Resource: "testexecutions"}).
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

	var readResponse TestsTestkubeIoTestExecutionV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("tests.testkube.io/v1")
	data.Kind = pointer.String("TestExecution")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
