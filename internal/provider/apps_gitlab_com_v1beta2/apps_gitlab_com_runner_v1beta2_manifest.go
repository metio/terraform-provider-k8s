/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_gitlab_com_v1beta2

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
	_ datasource.DataSource = &AppsGitlabComRunnerV1Beta2Manifest{}
)

func NewAppsGitlabComRunnerV1Beta2Manifest() datasource.DataSource {
	return &AppsGitlabComRunnerV1Beta2Manifest{}
}

type AppsGitlabComRunnerV1Beta2Manifest struct{}

type AppsGitlabComRunnerV1Beta2ManifestData struct {
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
		Azure *struct {
			Container     *string `tfsdk:"container" json:"container,omitempty"`
			Credentials   *string `tfsdk:"credentials" json:"credentials,omitempty"`
			StorageDomain *string `tfsdk:"storage_domain" json:"storageDomain,omitempty"`
		} `tfsdk:"azure" json:"azure,omitempty"`
		BuildImage  *string `tfsdk:"build_image" json:"buildImage,omitempty"`
		Ca          *string `tfsdk:"ca" json:"ca,omitempty"`
		CachePath   *string `tfsdk:"cache_path" json:"cachePath,omitempty"`
		CacheShared *bool   `tfsdk:"cache_shared" json:"cacheShared,omitempty"`
		CacheType   *string `tfsdk:"cache_type" json:"cacheType,omitempty"`
		CloneURL    *string `tfsdk:"clone_url" json:"cloneURL,omitempty"`
		Concurrent  *int64  `tfsdk:"concurrent" json:"concurrent,omitempty"`
		Config      *string `tfsdk:"config" json:"config,omitempty"`
		Env         *string `tfsdk:"env" json:"env,omitempty"`
		Gcs         *struct {
			Bucket          *string `tfsdk:"bucket" json:"bucket,omitempty"`
			Credentials     *string `tfsdk:"credentials" json:"credentials,omitempty"`
			CredentialsFile *string `tfsdk:"credentials_file" json:"credentialsFile,omitempty"`
		} `tfsdk:"gcs" json:"gcs,omitempty"`
		GitlabUrl       *string `tfsdk:"gitlab_url" json:"gitlabUrl,omitempty"`
		HelperImage     *string `tfsdk:"helper_image" json:"helperImage,omitempty"`
		ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
		Interval        *int64  `tfsdk:"interval" json:"interval,omitempty"`
		Locked          *bool   `tfsdk:"locked" json:"locked,omitempty"`
		PodSpec         *[]struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Patch     *string `tfsdk:"patch" json:"patch,omitempty"`
			PatchFile *string `tfsdk:"patch_file" json:"patchFile,omitempty"`
			PatchType *string `tfsdk:"patch_type" json:"patchType,omitempty"`
		} `tfsdk:"pod_spec" json:"podSpec,omitempty"`
		Protected   *bool   `tfsdk:"protected" json:"protected,omitempty"`
		RunUntagged *bool   `tfsdk:"run_untagged" json:"runUntagged,omitempty"`
		RunnerImage *string `tfsdk:"runner_image" json:"runnerImage,omitempty"`
		S3          *struct {
			Bucket      *string `tfsdk:"bucket" json:"bucket,omitempty"`
			Credentials *string `tfsdk:"credentials" json:"credentials,omitempty"`
			Insecure    *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
			Location    *string `tfsdk:"location" json:"location,omitempty"`
			Server      *string `tfsdk:"server" json:"server,omitempty"`
		} `tfsdk:"s3" json:"s3,omitempty"`
		Serviceaccount *string `tfsdk:"serviceaccount" json:"serviceaccount,omitempty"`
		Tags           *string `tfsdk:"tags" json:"tags,omitempty"`
		Token          *string `tfsdk:"token" json:"token,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsGitlabComRunnerV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_gitlab_com_runner_v1beta2_manifest"
}

func (r *AppsGitlabComRunnerV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Runner is the open source project used to run your jobs and send the results back to GitLab",
		MarkdownDescription: "Runner is the open source project used to run your jobs and send the results back to GitLab",
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
				Description:         "Specification of the desired behavior of a GitLab Runner instance",
				MarkdownDescription: "Specification of the desired behavior of a GitLab Runner instance",
				Attributes: map[string]schema.Attribute{
					"azure": schema.SingleNestedAttribute{
						Description:         "options used to setup Azure blob storage as GitLab Runner Cache",
						MarkdownDescription: "options used to setup Azure blob storage as GitLab Runner Cache",
						Attributes: map[string]schema.Attribute{
							"container": schema.StringAttribute{
								Description:         "Name of the Azure container in which the cache will be stored",
								MarkdownDescription: "Name of the Azure container in which the cache will be stored",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"credentials": schema.StringAttribute{
								Description:         "Credentials secret contains 'accountName' and 'privateKey' used to authenticate against Azure blob storage",
								MarkdownDescription: "Credentials secret contains 'accountName' and 'privateKey' used to authenticate against Azure blob storage",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"storage_domain": schema.StringAttribute{
								Description:         "The domain name of the Azure blob storage e.g. blob.core.windows.net",
								MarkdownDescription: "The domain name of the Azure blob storage e.g. blob.core.windows.net",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"build_image": schema.StringAttribute{
						Description:         "The name of the default image to use to run build jobs, when none is specified",
						MarkdownDescription: "The name of the default image to use to run build jobs, when none is specified",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ca": schema.StringAttribute{
						Description:         "Name of tls secret containing the custom certificate authority (CA) certificates",
						MarkdownDescription: "Name of tls secret containing the custom certificate authority (CA) certificates",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cache_path": schema.StringAttribute{
						Description:         "Path defines the Runner Cache path",
						MarkdownDescription: "Path defines the Runner Cache path",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cache_shared": schema.BoolAttribute{
						Description:         "Enable sharing of cache between Runners",
						MarkdownDescription: "Enable sharing of cache between Runners",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cache_type": schema.StringAttribute{
						Description:         "Type of cache used for Runner artifacts Options are: gcs, s3, azure",
						MarkdownDescription: "Type of cache used for Runner artifacts Options are: gcs, s3, azure",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"clone_url": schema.StringAttribute{
						Description:         "If specified, overrides the default URL used to clone or fetch the Git ref",
						MarkdownDescription: "If specified, overrides the default URL used to clone or fetch the Git ref",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"concurrent": schema.Int64Attribute{
						Description:         "Option to limit the number of jobs globally that can run concurrently. The operator sets this to 10, if not specified",
						MarkdownDescription: "Option to limit the number of jobs globally that can run concurrently. The operator sets this to 10, if not specified",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"config": schema.StringAttribute{
						Description:         "allow user to provide configmap name containing the user provided config.toml",
						MarkdownDescription: "allow user to provide configmap name containing the user provided config.toml",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"env": schema.StringAttribute{
						Description:         "Accepts configmap name. Provides user mechanism to inject environment variables in the GitLab Runner pod via the key value pairs in the ConfigMap",
						MarkdownDescription: "Accepts configmap name. Provides user mechanism to inject environment variables in the GitLab Runner pod via the key value pairs in the ConfigMap",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"gcs": schema.SingleNestedAttribute{
						Description:         "options used to setup GCS (Google Container Storage) as GitLab Runner Cache",
						MarkdownDescription: "options used to setup GCS (Google Container Storage) as GitLab Runner Cache",
						Attributes: map[string]schema.Attribute{
							"bucket": schema.StringAttribute{
								Description:         "Name of the bucket in which the cache will be stored",
								MarkdownDescription: "Name of the bucket in which the cache will be stored",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"credentials": schema.StringAttribute{
								Description:         "contains the GCS 'access-id' and 'private-key'",
								MarkdownDescription: "contains the GCS 'access-id' and 'private-key'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"credentials_file": schema.StringAttribute{
								Description:         "Takes GCS credentials file, 'keys.json'",
								MarkdownDescription: "Takes GCS credentials file, 'keys.json'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"gitlab_url": schema.StringAttribute{
						Description:         "The fully qualified domain name for the GitLab instance. For example, https://gitlab.example.com",
						MarkdownDescription: "The fully qualified domain name for the GitLab instance. For example, https://gitlab.example.com",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"helper_image": schema.StringAttribute{
						Description:         "If specified, overrides the default GitLab Runner helper image",
						MarkdownDescription: "If specified, overrides the default GitLab Runner helper image",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_pull_policy": schema.StringAttribute{
						Description:         "ImagePullPolicy sets the Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
						MarkdownDescription: "ImagePullPolicy sets the Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"interval": schema.Int64Attribute{
						Description:         "Option to define the number of seconds between checks for new jobs. This is set to a default of 30s by operator if not set",
						MarkdownDescription: "Option to define the number of seconds between checks for new jobs. This is set to a default of 30s by operator if not set",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"locked": schema.BoolAttribute{
						Description:         "Specify whether the runner should be locked to a specific project. Defaults to false.",
						MarkdownDescription: "Specify whether the runner should be locked to a specific project. Defaults to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pod_spec": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name is the name given to the custom Pod Spec",
									MarkdownDescription: "Name is the name given to the custom Pod Spec",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"patch": schema.StringAttribute{
									Description:         "A JSON or YAML format string that describes the changes which must be applied to the final PodSpec object before it is generated. You cannot set the patch_path and patch in the same pod_spec configuration, otherwise an error occurs.",
									MarkdownDescription: "A JSON or YAML format string that describes the changes which must be applied to the final PodSpec object before it is generated. You cannot set the patch_path and patch in the same pod_spec configuration, otherwise an error occurs.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"patch_file": schema.StringAttribute{
									Description:         "Path to the file that defines the changes to apply to the final PodSpec object before it is generated. The file must be a JSON or YAML file. You cannot set the patch_path and patch in the same pod_spec configuration, otherwise an error occurs.",
									MarkdownDescription: "Path to the file that defines the changes to apply to the final PodSpec object before it is generated. The file must be a JSON or YAML file. You cannot set the patch_path and patch in the same pod_spec configuration, otherwise an error occurs.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"patch_type": schema.StringAttribute{
									Description:         "The strategy the runner uses to apply the specified changes to the PodSpec object generated by GitLab Runner. The accepted values are merge, json, and strategic (default value).",
									MarkdownDescription: "The strategy the runner uses to apply the specified changes to the PodSpec object generated by GitLab Runner. The accepted values are merge, json, and strategic (default value).",
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

					"protected": schema.BoolAttribute{
						Description:         "Specify whether the runner should only run protected branches. Defaults to false.",
						MarkdownDescription: "Specify whether the runner should only run protected branches. Defaults to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"run_untagged": schema.BoolAttribute{
						Description:         "Specify if jobs without tags should be run. If not specified, runner will default to true if no tags were specified. In other case it will default to false.",
						MarkdownDescription: "Specify if jobs without tags should be run. If not specified, runner will default to true if no tags were specified. In other case it will default to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"runner_image": schema.StringAttribute{
						Description:         "If specified, overrides the default GitLab Runner image. Default is the Runner image the operator was bundled with.",
						MarkdownDescription: "If specified, overrides the default GitLab Runner image. Default is the Runner image the operator was bundled with.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"s3": schema.SingleNestedAttribute{
						Description:         "options used to setup S3 object store as GitLab Runner Cache",
						MarkdownDescription: "options used to setup S3 object store as GitLab Runner Cache",
						Attributes: map[string]schema.Attribute{
							"bucket": schema.StringAttribute{
								Description:         "Name of the bucket in which the cache will be stored",
								MarkdownDescription: "Name of the bucket in which the cache will be stored",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"credentials": schema.StringAttribute{
								Description:         "Name of the secret containing the 'accesskey' and 'secretkey' used to access the object storage",
								MarkdownDescription: "Name of the secret containing the 'accesskey' and 'secretkey' used to access the object storage",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"insecure": schema.BoolAttribute{
								Description:         "Use insecure connections or HTTP",
								MarkdownDescription: "Use insecure connections or HTTP",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"location": schema.StringAttribute{
								Description:         "Name of the S3 region in use",
								MarkdownDescription: "Name of the S3 region in use",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"server": schema.StringAttribute{
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

					"serviceaccount": schema.StringAttribute{
						Description:         "allow user to override service account used by GitLab Runner",
						MarkdownDescription: "allow user to override service account used by GitLab Runner",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.StringAttribute{
						Description:         "List of comma separated tags to be applied to the runner More info: https://docs.gitlab.com/ee/ci/runners/#use-tags-to-limit-the-number-of-jobs-using-the-runner",
						MarkdownDescription: "List of comma separated tags to be applied to the runner More info: https://docs.gitlab.com/ee/ci/runners/#use-tags-to-limit-the-number-of-jobs-using-the-runner",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token": schema.StringAttribute{
						Description:         "Name of secret containing the 'runner-registration-token' key used to register the runner",
						MarkdownDescription: "Name of secret containing the 'runner-registration-token' key used to register the runner",
						Required:            true,
						Optional:            false,
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

func (r *AppsGitlabComRunnerV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_gitlab_com_runner_v1beta2_manifest")

	var model AppsGitlabComRunnerV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("apps.gitlab.com/v1beta2")
	model.Kind = pointer.String("Runner")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
