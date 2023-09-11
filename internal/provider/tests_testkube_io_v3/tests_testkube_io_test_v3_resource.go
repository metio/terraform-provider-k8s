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
)

var (
	_ resource.Resource                = &TestsTestkubeIoTestV3Resource{}
	_ resource.ResourceWithConfigure   = &TestsTestkubeIoTestV3Resource{}
	_ resource.ResourceWithImportState = &TestsTestkubeIoTestV3Resource{}
)

func NewTestsTestkubeIoTestV3Resource() resource.Resource {
	return &TestsTestkubeIoTestV3Resource{}
}

type TestsTestkubeIoTestV3Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type TestsTestkubeIoTestV3ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Content *struct {
			Data       *string `tfsdk:"data" json:"data,omitempty"`
			Repository *struct {
				AuthType          *string `tfsdk:"auth_type" json:"authType,omitempty"`
				Branch            *string `tfsdk:"branch" json:"branch,omitempty"`
				CertificateSecret *string `tfsdk:"certificate_secret" json:"certificateSecret,omitempty"`
				Commit            *string `tfsdk:"commit" json:"commit,omitempty"`
				Path              *string `tfsdk:"path" json:"path,omitempty"`
				TokenSecret       *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"token_secret" json:"tokenSecret,omitempty"`
				Type           *string `tfsdk:"type" json:"type,omitempty"`
				Uri            *string `tfsdk:"uri" json:"uri,omitempty"`
				UsernameSecret *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"username_secret" json:"usernameSecret,omitempty"`
				WorkingDir *string `tfsdk:"working_dir" json:"workingDir,omitempty"`
			} `tfsdk:"repository" json:"repository,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
			Uri  *string `tfsdk:"uri" json:"uri,omitempty"`
		} `tfsdk:"content" json:"content,omitempty"`
		Description      *string `tfsdk:"description" json:"description,omitempty"`
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
			Command                  *[]string `tfsdk:"command" json:"command,omitempty"`
			CronJobTemplate          *string   `tfsdk:"cron_job_template" json:"cronJobTemplate,omitempty"`
			CronJobTemplateReference *string   `tfsdk:"cron_job_template_reference" json:"cronJobTemplateReference,omitempty"`
			EnvConfigMaps            *[]struct {
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
			IsVariablesFileUploaded  *bool              `tfsdk:"is_variables_file_uploaded" json:"isVariablesFileUploaded,omitempty"`
			JobTemplate              *string            `tfsdk:"job_template" json:"jobTemplate,omitempty"`
			JobTemplateReference     *string            `tfsdk:"job_template_reference" json:"jobTemplateReference,omitempty"`
			Name                     *string            `tfsdk:"name" json:"name,omitempty"`
			Namespace                *string            `tfsdk:"namespace" json:"namespace,omitempty"`
			NegativeTest             *bool              `tfsdk:"negative_test" json:"negativeTest,omitempty"`
			Number                   *int64             `tfsdk:"number" json:"number,omitempty"`
			PostRunScript            *string            `tfsdk:"post_run_script" json:"postRunScript,omitempty"`
			PreRunScript             *string            `tfsdk:"pre_run_script" json:"preRunScript,omitempty"`
			PvcTemplate              *string            `tfsdk:"pvc_template" json:"pvcTemplate,omitempty"`
			PvcTemplateReference     *string            `tfsdk:"pvc_template_reference" json:"pvcTemplateReference,omitempty"`
			ScraperTemplate          *string            `tfsdk:"scraper_template" json:"scraperTemplate,omitempty"`
			ScraperTemplateReference *string            `tfsdk:"scraper_template_reference" json:"scraperTemplateReference,omitempty"`
			SecretEnvs               *map[string]string `tfsdk:"secret_envs" json:"secretEnvs,omitempty"`
			Sync                     *bool              `tfsdk:"sync" json:"sync,omitempty"`
			TestSecretUUID           *string            `tfsdk:"test_secret_uuid" json:"testSecretUUID,omitempty"`
			TestSuiteName            *string            `tfsdk:"test_suite_name" json:"testSuiteName,omitempty"`
			TestSuiteSecretUUID      *string            `tfsdk:"test_suite_secret_uuid" json:"testSuiteSecretUUID,omitempty"`
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
			VariablesFile *string `tfsdk:"variables_file" json:"variablesFile,omitempty"`
		} `tfsdk:"execution_request" json:"executionRequest,omitempty"`
		Name     *string   `tfsdk:"name" json:"name,omitempty"`
		Schedule *string   `tfsdk:"schedule" json:"schedule,omitempty"`
		Source   *string   `tfsdk:"source" json:"source,omitempty"`
		Type     *string   `tfsdk:"type" json:"type,omitempty"`
		Uploads  *[]string `tfsdk:"uploads" json:"uploads,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TestsTestkubeIoTestV3Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_tests_testkube_io_test_v3"
}

func (r *TestsTestkubeIoTestV3Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Test is the Schema for the tests API",
		MarkdownDescription: "Test is the Schema for the tests API",
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

			"wait_for": schema.ListNestedAttribute{
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
				Description:         "TestSpec defines the desired state of Test",
				MarkdownDescription: "TestSpec defines the desired state of Test",
				Attributes: map[string]schema.Attribute{
					"content": schema.SingleNestedAttribute{
						Description:         "test content object",
						MarkdownDescription: "test content object",
						Attributes: map[string]schema.Attribute{
							"data": schema.StringAttribute{
								Description:         "test content body",
								MarkdownDescription: "test content body",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"repository": schema.SingleNestedAttribute{
								Description:         "repository of test content",
								MarkdownDescription: "repository of test content",
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Description:         "auth type for git requests",
										MarkdownDescription: "auth type for git requests",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("basic", "header"),
										},
									},

									"branch": schema.StringAttribute{
										Description:         "branch/tag name for checkout",
										MarkdownDescription: "branch/tag name for checkout",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"certificate_secret": schema.StringAttribute{
										Description:         "git auth certificate secret for private repositories",
										MarkdownDescription: "git auth certificate secret for private repositories",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"commit": schema.StringAttribute{
										Description:         "commit id (sha) for checkout",
										MarkdownDescription: "commit id (sha) for checkout",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"path": schema.StringAttribute{
										Description:         "if needed we can checkout particular path (dir or file) in case of BIG/mono repositories",
										MarkdownDescription: "if needed we can checkout particular path (dir or file) in case of BIG/mono repositories",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"token_secret": schema.SingleNestedAttribute{
										Description:         "Testkube internal reference for secret storage in Kubernetes secrets",
										MarkdownDescription: "Testkube internal reference for secret storage in Kubernetes secrets",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "object key",
												MarkdownDescription: "object key",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "object name",
												MarkdownDescription: "object name",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": schema.StringAttribute{
										Description:         "VCS repository type",
										MarkdownDescription: "VCS repository type",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"uri": schema.StringAttribute{
										Description:         "uri of content file or git directory",
										MarkdownDescription: "uri of content file or git directory",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"username_secret": schema.SingleNestedAttribute{
										Description:         "Testkube internal reference for secret storage in Kubernetes secrets",
										MarkdownDescription: "Testkube internal reference for secret storage in Kubernetes secrets",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "object key",
												MarkdownDescription: "object key",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "object name",
												MarkdownDescription: "object name",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"working_dir": schema.StringAttribute{
										Description:         "if provided we checkout the whole repository and run test from this directory",
										MarkdownDescription: "if provided we checkout the whole repository and run test from this directory",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": schema.StringAttribute{
								Description:         "test type",
								MarkdownDescription: "test type",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("string", "file-uri", "git-file", "git-dir", "git"),
								},
							},

							"uri": schema.StringAttribute{
								Description:         "uri of test content",
								MarkdownDescription: "uri of test content",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"description": schema.StringAttribute{
						Description:         "test description",
						MarkdownDescription: "test description",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"execution_request": schema.SingleNestedAttribute{
						Description:         "test execution request body",
						MarkdownDescription: "test execution request body",
						Attributes: map[string]schema.Attribute{
							"active_deadline_seconds": schema.Int64Attribute{
								Description:         "Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.",
								MarkdownDescription: "Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"args": schema.ListAttribute{
								Description:         "additional executor binary arguments",
								MarkdownDescription: "additional executor binary arguments",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"args_mode": schema.StringAttribute{
								Description:         "usage mode for arguments",
								MarkdownDescription: "usage mode for arguments",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("append", "override"),
								},
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
										Optional:            true,
										Computed:            false,
									},

									"omit_folder_per_execution": schema.BoolAttribute{
										Description:         "don't use a separate folder for execution artifacts",
										MarkdownDescription: "don't use a separate folder for execution artifacts",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage_bucket": schema.StringAttribute{
										Description:         "artifact bucket storage",
										MarkdownDescription: "artifact bucket storage",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage_class_name": schema.StringAttribute{
										Description:         "artifact storage class name for container executor",
										MarkdownDescription: "artifact storage class name for container executor",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"volume_mount_path": schema.StringAttribute{
										Description:         "artifact volume mount path for container executor",
										MarkdownDescription: "artifact volume mount path for container executor",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"command": schema.ListAttribute{
								Description:         "executor binary command",
								MarkdownDescription: "executor binary command",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cron_job_template": schema.StringAttribute{
								Description:         "cron job template extensions",
								MarkdownDescription: "cron job template extensions",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cron_job_template_reference": schema.StringAttribute{
								Description:         "name of the template resource",
								MarkdownDescription: "name of the template resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
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
											Optional:            true,
											Computed:            false,
										},

										"mount": schema.BoolAttribute{
											Description:         "whether we shoud mount resource",
											MarkdownDescription: "whether we shoud mount resource",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mount_path": schema.StringAttribute{
											Description:         "where we shoud mount resource",
											MarkdownDescription: "where we shoud mount resource",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"reference": schema.SingleNestedAttribute{
											Description:         "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
											MarkdownDescription: "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
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
											Optional:            true,
											Computed:            false,
										},

										"mount": schema.BoolAttribute{
											Description:         "whether we shoud mount resource",
											MarkdownDescription: "whether we shoud mount resource",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mount_path": schema.StringAttribute{
											Description:         "where we shoud mount resource",
											MarkdownDescription: "where we shoud mount resource",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"reference": schema.SingleNestedAttribute{
											Description:         "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
											MarkdownDescription: "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"envs": schema.MapAttribute{
								Description:         "Environment variables passed to executor. Deprecated: use Basic Variables instead",
								MarkdownDescription: "Environment variables passed to executor. Deprecated: use Basic Variables instead",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"execution_labels": schema.MapAttribute{
								Description:         "test execution labels",
								MarkdownDescription: "test execution labels",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"http_proxy": schema.StringAttribute{
								Description:         "http proxy for executor containers",
								MarkdownDescription: "http proxy for executor containers",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"https_proxy": schema.StringAttribute{
								Description:         "https proxy for executor containers",
								MarkdownDescription: "https proxy for executor containers",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.StringAttribute{
								Description:         "container executor image",
								MarkdownDescription: "container executor image",
								Required:            false,
								Optional:            true,
								Computed:            false,
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
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"is_variables_file_uploaded": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"job_template": schema.StringAttribute{
								Description:         "job template extensions",
								MarkdownDescription: "job template extensions",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"job_template_reference": schema.StringAttribute{
								Description:         "name of the template resource",
								MarkdownDescription: "name of the template resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "test execution custom name",
								MarkdownDescription: "test execution custom name",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "test kubernetes namespace ('testkube' when not set)",
								MarkdownDescription: "test kubernetes namespace ('testkube' when not set)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"negative_test": schema.BoolAttribute{
								Description:         "negative test will fail the execution if it is a success and it will succeed if it is a failure",
								MarkdownDescription: "negative test will fail the execution if it is a success and it will succeed if it is a failure",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"number": schema.Int64Attribute{
								Description:         "test execution number",
								MarkdownDescription: "test execution number",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"post_run_script": schema.StringAttribute{
								Description:         "script to run after test execution",
								MarkdownDescription: "script to run after test execution",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pre_run_script": schema.StringAttribute{
								Description:         "script to run before test execution",
								MarkdownDescription: "script to run before test execution",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pvc_template": schema.StringAttribute{
								Description:         "pvc template extensions",
								MarkdownDescription: "pvc template extensions",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pvc_template_reference": schema.StringAttribute{
								Description:         "name of the template resource",
								MarkdownDescription: "name of the template resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scraper_template": schema.StringAttribute{
								Description:         "scraper template extensions",
								MarkdownDescription: "scraper template extensions",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scraper_template_reference": schema.StringAttribute{
								Description:         "name of the template resource",
								MarkdownDescription: "name of the template resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_envs": schema.MapAttribute{
								Description:         "Execution variables passed to executor from secrets. Deprecated: use Secret Variables instead",
								MarkdownDescription: "Execution variables passed to executor from secrets. Deprecated: use Secret Variables instead",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sync": schema.BoolAttribute{
								Description:         "whether to start execution sync or async",
								MarkdownDescription: "whether to start execution sync or async",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"test_secret_uuid": schema.StringAttribute{
								Description:         "test secret uuid",
								MarkdownDescription: "test secret uuid",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"test_suite_name": schema.StringAttribute{
								Description:         "unique test suite name (CRD Test suite name), if it's run as a part of test suite",
								MarkdownDescription: "unique test suite name (CRD Test suite name), if it's run as a part of test suite",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"test_suite_secret_uuid": schema.StringAttribute{
								Description:         "test suite secret uuid, if it's run as a part of test suite",
								MarkdownDescription: "test suite secret uuid, if it's run as a part of test suite",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"variables": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "variable name",
										MarkdownDescription: "variable name",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "variable type",
										MarkdownDescription: "variable type",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"value": schema.StringAttribute{
										Description:         "variable string value",
										MarkdownDescription: "variable string value",
										Required:            false,
										Optional:            true,
										Computed:            false,
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
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"optional": schema.BoolAttribute{
														Description:         "Specify whether the ConfigMap or its key must be defined",
														MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_ref": schema.SingleNestedAttribute{
												Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
												MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
														MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"field_path": schema.StringAttribute{
														Description:         "Path of the field to select in the specified API version.",
														MarkdownDescription: "Path of the field to select in the specified API version.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"resource_field_ref": schema.SingleNestedAttribute{
												Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
												MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
												Attributes: map[string]schema.Attribute{
													"container_name": schema.StringAttribute{
														Description:         "Container name: required for volumes, optional for env vars",
														MarkdownDescription: "Container name: required for volumes, optional for env vars",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"divisor": schema.StringAttribute{
														Description:         "Specifies the output format of the exposed resources, defaults to '1'",
														MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"resource": schema.StringAttribute{
														Description:         "Required: resource to select",
														MarkdownDescription: "Required: resource to select",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_key_ref": schema.SingleNestedAttribute{
												Description:         "Selects a key of a secret in the pod's namespace",
												MarkdownDescription: "Selects a key of a secret in the pod's namespace",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"optional": schema.BoolAttribute{
														Description:         "Specify whether the Secret or its key must be defined",
														MarkdownDescription: "Specify whether the Secret or its key must be defined",
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

							"variables_file": schema.StringAttribute{
								Description:         "variables file content - need to be in format for particular executor (e.g. postman envs file)",
								MarkdownDescription: "variables file content - need to be in format for particular executor (e.g. postman envs file)",
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
						Description:         "test name",
						MarkdownDescription: "test name",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"schedule": schema.StringAttribute{
						Description:         "schedule in cron job format for scheduled test execution",
						MarkdownDescription: "schedule in cron job format for scheduled test execution",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"source": schema.StringAttribute{
						Description:         "reference to test source resource",
						MarkdownDescription: "reference to test source resource",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"type": schema.StringAttribute{
						Description:         "test type",
						MarkdownDescription: "test type",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"uploads": schema.ListAttribute{
						Description:         "files to be used from minio uploads",
						MarkdownDescription: "files to be used from minio uploads",
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
	}
}

func (r *TestsTestkubeIoTestV3Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *TestsTestkubeIoTestV3Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_tests_testkube_io_test_v3")

	var model TestsTestkubeIoTestV3ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("tests.testkube.io/v3")
	model.Kind = pointer.String("Test")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
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
		Resource(k8sSchema.GroupVersionResource{Group: "tests.testkube.io", Version: "v3", Resource: "tests"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse TestsTestkubeIoTestV3ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *TestsTestkubeIoTestV3Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_tests_testkube_io_test_v3")

	var data TestsTestkubeIoTestV3ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "tests.testkube.io", Version: "v3", Resource: "tests"}).
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

	var readResponse TestsTestkubeIoTestV3ResourceData
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

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *TestsTestkubeIoTestV3Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_tests_testkube_io_test_v3")

	var model TestsTestkubeIoTestV3ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("tests.testkube.io/v3")
	model.Kind = pointer.String("Test")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
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
		Resource(k8sSchema.GroupVersionResource{Group: "tests.testkube.io", Version: "v3", Resource: "tests"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse TestsTestkubeIoTestV3ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *TestsTestkubeIoTestV3Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_tests_testkube_io_test_v3")

	var data TestsTestkubeIoTestV3ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "tests.testkube.io", Version: "v3", Resource: "tests"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *TestsTestkubeIoTestV3Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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
