/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

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

type AppsGitlabComRunnerV1Beta2Resource struct{}

var (
	_ resource.Resource = (*AppsGitlabComRunnerV1Beta2Resource)(nil)
)

type AppsGitlabComRunnerV1Beta2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type AppsGitlabComRunnerV1Beta2GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		BuildImage *string `tfsdk:"build_image" yaml:"buildImage,omitempty"`

		Ca *string `tfsdk:"ca" yaml:"ca,omitempty"`

		Env *string `tfsdk:"env" yaml:"env,omitempty"`

		Azure *struct {
			Credentials *string `tfsdk:"credentials" yaml:"credentials,omitempty"`

			StorageDomain *string `tfsdk:"storage_domain" yaml:"storageDomain,omitempty"`

			Container *string `tfsdk:"container" yaml:"container,omitempty"`
		} `tfsdk:"azure" yaml:"azure,omitempty"`

		CloneURL *string `tfsdk:"clone_url" yaml:"cloneURL,omitempty"`

		Config *string `tfsdk:"config" yaml:"config,omitempty"`

		Serviceaccount *string `tfsdk:"serviceaccount" yaml:"serviceaccount,omitempty"`

		CacheShared *bool `tfsdk:"cache_shared" yaml:"cacheShared,omitempty"`

		CacheType *string `tfsdk:"cache_type" yaml:"cacheType,omitempty"`

		GitlabUrl *string `tfsdk:"gitlab_url" yaml:"gitlabUrl,omitempty"`

		ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

		Interval *int64 `tfsdk:"interval" yaml:"interval,omitempty"`

		Locked *bool `tfsdk:"locked" yaml:"locked,omitempty"`

		RunnerImage *string `tfsdk:"runner_image" yaml:"runnerImage,omitempty"`

		S3 *struct {
			Bucket *string `tfsdk:"bucket" yaml:"bucket,omitempty"`

			Credentials *string `tfsdk:"credentials" yaml:"credentials,omitempty"`

			Insecure *bool `tfsdk:"insecure" yaml:"insecure,omitempty"`

			Location *string `tfsdk:"location" yaml:"location,omitempty"`

			Server *string `tfsdk:"server" yaml:"server,omitempty"`
		} `tfsdk:"s3" yaml:"s3,omitempty"`

		CachePath *string `tfsdk:"cache_path" yaml:"cachePath,omitempty"`

		Concurrent *int64 `tfsdk:"concurrent" yaml:"concurrent,omitempty"`

		Gcs *struct {
			Bucket *string `tfsdk:"bucket" yaml:"bucket,omitempty"`

			Credentials *string `tfsdk:"credentials" yaml:"credentials,omitempty"`

			CredentialsFile *string `tfsdk:"credentials_file" yaml:"credentialsFile,omitempty"`
		} `tfsdk:"gcs" yaml:"gcs,omitempty"`

		HelperImage *string `tfsdk:"helper_image" yaml:"helperImage,omitempty"`

		Protected *bool `tfsdk:"protected" yaml:"protected,omitempty"`

		RunUntagged *bool `tfsdk:"run_untagged" yaml:"runUntagged,omitempty"`

		Tags *string `tfsdk:"tags" yaml:"tags,omitempty"`

		Token *string `tfsdk:"token" yaml:"token,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewAppsGitlabComRunnerV1Beta2Resource() resource.Resource {
	return &AppsGitlabComRunnerV1Beta2Resource{}
}

func (r *AppsGitlabComRunnerV1Beta2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apps_gitlab_com_runner_v1beta2"
}

func (r *AppsGitlabComRunnerV1Beta2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Runner is the open source project used to run your jobs and send the results back to GitLab",
		MarkdownDescription: "Runner is the open source project used to run your jobs and send the results back to GitLab",
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
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
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
				Description:         "Specification of the desired behavior of a GitLab Runner instance",
				MarkdownDescription: "Specification of the desired behavior of a GitLab Runner instance",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"build_image": {
						Description:         "The name of the default image to use to run build jobs, when none is specified",
						MarkdownDescription: "The name of the default image to use to run build jobs, when none is specified",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ca": {
						Description:         "Name of tls secret containing the custom certificate authority (CA) certificates",
						MarkdownDescription: "Name of tls secret containing the custom certificate authority (CA) certificates",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"env": {
						Description:         "Accepts configmap name. Provides user mechanism to inject environment variables in the GitLab Runner pod via the key value pairs in the ConfigMap",
						MarkdownDescription: "Accepts configmap name. Provides user mechanism to inject environment variables in the GitLab Runner pod via the key value pairs in the ConfigMap",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"azure": {
						Description:         "options used to setup Azure blob storage as GitLab Runner Cache",
						MarkdownDescription: "options used to setup Azure blob storage as GitLab Runner Cache",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"credentials": {
								Description:         "Credentials secret contains 'accountName' and 'privateKey' used to authenticate against Azure blob storage",
								MarkdownDescription: "Credentials secret contains 'accountName' and 'privateKey' used to authenticate against Azure blob storage",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"storage_domain": {
								Description:         "The domain name of the Azure blob storage e.g. blob.core.windows.net",
								MarkdownDescription: "The domain name of the Azure blob storage e.g. blob.core.windows.net",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"container": {
								Description:         "Name of the Azure container in which the cache will be stored",
								MarkdownDescription: "Name of the Azure container in which the cache will be stored",

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

					"clone_url": {
						Description:         "If specified, overrides the default URL used to clone or fetch the Git ref",
						MarkdownDescription: "If specified, overrides the default URL used to clone or fetch the Git ref",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"config": {
						Description:         "allow user to provide configmap name containing the user provided config.toml",
						MarkdownDescription: "allow user to provide configmap name containing the user provided config.toml",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"serviceaccount": {
						Description:         "allow user to override service account used by GitLab Runner",
						MarkdownDescription: "allow user to override service account used by GitLab Runner",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cache_shared": {
						Description:         "Enable sharing of cache between Runners",
						MarkdownDescription: "Enable sharing of cache between Runners",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cache_type": {
						Description:         "Type of cache used for Runner artifacts Options are: gcs, s3, azure",
						MarkdownDescription: "Type of cache used for Runner artifacts Options are: gcs, s3, azure",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"gitlab_url": {
						Description:         "The fully qualified domain name for the GitLab instance. For example, https://gitlab.example.com",
						MarkdownDescription: "The fully qualified domain name for the GitLab instance. For example, https://gitlab.example.com",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"image_pull_policy": {
						Description:         "ImagePullPolicy sets the Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
						MarkdownDescription: "ImagePullPolicy sets the Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"interval": {
						Description:         "Option to define the number of seconds between checks for new jobs. This is set to a default of 30s by operator if not set",
						MarkdownDescription: "Option to define the number of seconds between checks for new jobs. This is set to a default of 30s by operator if not set",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"locked": {
						Description:         "Specify whether the runner should be locked to a specific project. Defaults to false.",
						MarkdownDescription: "Specify whether the runner should be locked to a specific project. Defaults to false.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"runner_image": {
						Description:         "If specified, overrides the default GitLab Runner image. Default is the Runner image the operator was bundled with.",
						MarkdownDescription: "If specified, overrides the default GitLab Runner image. Default is the Runner image the operator was bundled with.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"s3": {
						Description:         "options used to setup S3 object store as GitLab Runner Cache",
						MarkdownDescription: "options used to setup S3 object store as GitLab Runner Cache",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"bucket": {
								Description:         "Name of the bucket in which the cache will be stored",
								MarkdownDescription: "Name of the bucket in which the cache will be stored",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"credentials": {
								Description:         "Name of the secret containing the 'accesskey' and 'secretkey' used to access the object storage",
								MarkdownDescription: "Name of the secret containing the 'accesskey' and 'secretkey' used to access the object storage",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"insecure": {
								Description:         "Use insecure connections or HTTP",
								MarkdownDescription: "Use insecure connections or HTTP",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"location": {
								Description:         "Name of the S3 region in use",
								MarkdownDescription: "Name of the S3 region in use",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"server": {
								Description:         "",
								MarkdownDescription: "",

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

					"cache_path": {
						Description:         "Path defines the Runner Cache path",
						MarkdownDescription: "Path defines the Runner Cache path",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"concurrent": {
						Description:         "Option to limit the number of jobs globally that can run concurrently. The operator sets this to 10, if not specified",
						MarkdownDescription: "Option to limit the number of jobs globally that can run concurrently. The operator sets this to 10, if not specified",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"gcs": {
						Description:         "options used to setup GCS (Google Container Storage) as GitLab Runner Cache",
						MarkdownDescription: "options used to setup GCS (Google Container Storage) as GitLab Runner Cache",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"bucket": {
								Description:         "Name of the bucket in which the cache will be stored",
								MarkdownDescription: "Name of the bucket in which the cache will be stored",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"credentials": {
								Description:         "contains the GCS 'access-id' and 'private-key'",
								MarkdownDescription: "contains the GCS 'access-id' and 'private-key'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"credentials_file": {
								Description:         "Takes GCS credentials file, 'keys.json'",
								MarkdownDescription: "Takes GCS credentials file, 'keys.json'",

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

					"helper_image": {
						Description:         "If specified, overrides the default GitLab Runner helper image",
						MarkdownDescription: "If specified, overrides the default GitLab Runner helper image",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"protected": {
						Description:         "Specify whether the runner should only run protected branches. Defaults to false.",
						MarkdownDescription: "Specify whether the runner should only run protected branches. Defaults to false.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"run_untagged": {
						Description:         "Specify if jobs without tags should be run. If not specified, runner will default to true if no tags were specified. In other case it will default to false.",
						MarkdownDescription: "Specify if jobs without tags should be run. If not specified, runner will default to true if no tags were specified. In other case it will default to false.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tags": {
						Description:         "List of comma separated tags to be applied to the runner More info: https://docs.gitlab.com/ee/ci/runners/#use-tags-to-limit-the-number-of-jobs-using-the-runner",
						MarkdownDescription: "List of comma separated tags to be applied to the runner More info: https://docs.gitlab.com/ee/ci/runners/#use-tags-to-limit-the-number-of-jobs-using-the-runner",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"token": {
						Description:         "Name of secret containing the 'runner-registration-token' key used to register the runner",
						MarkdownDescription: "Name of secret containing the 'runner-registration-token' key used to register the runner",

						Type: types.StringType,

						Required: true,
						Optional: false,
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

func (r *AppsGitlabComRunnerV1Beta2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_apps_gitlab_com_runner_v1beta2")

	var state AppsGitlabComRunnerV1Beta2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel AppsGitlabComRunnerV1Beta2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("apps.gitlab.com/v1beta2")
	goModel.Kind = utilities.Ptr("Runner")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *AppsGitlabComRunnerV1Beta2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_gitlab_com_runner_v1beta2")
	// NO-OP: All data is already in Terraform state
}

func (r *AppsGitlabComRunnerV1Beta2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_apps_gitlab_com_runner_v1beta2")

	var state AppsGitlabComRunnerV1Beta2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel AppsGitlabComRunnerV1Beta2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("apps.gitlab.com/v1beta2")
	goModel.Kind = utilities.Ptr("Runner")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *AppsGitlabComRunnerV1Beta2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_apps_gitlab_com_runner_v1beta2")
	// NO-OP: Terraform removes the state automatically for us
}
