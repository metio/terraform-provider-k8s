/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_gitlab_com_v1beta2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &AppsGitlabComRunnerV1Beta2DataSource{}
	_ datasource.DataSourceWithConfigure = &AppsGitlabComRunnerV1Beta2DataSource{}
)

func NewAppsGitlabComRunnerV1Beta2DataSource() datasource.DataSource {
	return &AppsGitlabComRunnerV1Beta2DataSource{}
}

type AppsGitlabComRunnerV1Beta2DataSource struct {
	kubernetesClient dynamic.Interface
}

type AppsGitlabComRunnerV1Beta2DataSourceData struct {
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

func (r *AppsGitlabComRunnerV1Beta2DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_gitlab_com_runner_v1beta2"
}

func (r *AppsGitlabComRunnerV1Beta2DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
								Optional:            false,
								Computed:            true,
							},

							"credentials": schema.StringAttribute{
								Description:         "Credentials secret contains 'accountName' and 'privateKey' used to authenticate against Azure blob storage",
								MarkdownDescription: "Credentials secret contains 'accountName' and 'privateKey' used to authenticate against Azure blob storage",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"storage_domain": schema.StringAttribute{
								Description:         "The domain name of the Azure blob storage e.g. blob.core.windows.net",
								MarkdownDescription: "The domain name of the Azure blob storage e.g. blob.core.windows.net",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"build_image": schema.StringAttribute{
						Description:         "The name of the default image to use to run build jobs, when none is specified",
						MarkdownDescription: "The name of the default image to use to run build jobs, when none is specified",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ca": schema.StringAttribute{
						Description:         "Name of tls secret containing the custom certificate authority (CA) certificates",
						MarkdownDescription: "Name of tls secret containing the custom certificate authority (CA) certificates",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cache_path": schema.StringAttribute{
						Description:         "Path defines the Runner Cache path",
						MarkdownDescription: "Path defines the Runner Cache path",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cache_shared": schema.BoolAttribute{
						Description:         "Enable sharing of cache between Runners",
						MarkdownDescription: "Enable sharing of cache between Runners",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cache_type": schema.StringAttribute{
						Description:         "Type of cache used for Runner artifacts Options are: gcs, s3, azure",
						MarkdownDescription: "Type of cache used for Runner artifacts Options are: gcs, s3, azure",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"clone_url": schema.StringAttribute{
						Description:         "If specified, overrides the default URL used to clone or fetch the Git ref",
						MarkdownDescription: "If specified, overrides the default URL used to clone or fetch the Git ref",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"concurrent": schema.Int64Attribute{
						Description:         "Option to limit the number of jobs globally that can run concurrently. The operator sets this to 10, if not specified",
						MarkdownDescription: "Option to limit the number of jobs globally that can run concurrently. The operator sets this to 10, if not specified",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"config": schema.StringAttribute{
						Description:         "allow user to provide configmap name containing the user provided config.toml",
						MarkdownDescription: "allow user to provide configmap name containing the user provided config.toml",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"env": schema.StringAttribute{
						Description:         "Accepts configmap name. Provides user mechanism to inject environment variables in the GitLab Runner pod via the key value pairs in the ConfigMap",
						MarkdownDescription: "Accepts configmap name. Provides user mechanism to inject environment variables in the GitLab Runner pod via the key value pairs in the ConfigMap",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"gcs": schema.SingleNestedAttribute{
						Description:         "options used to setup GCS (Google Container Storage) as GitLab Runner Cache",
						MarkdownDescription: "options used to setup GCS (Google Container Storage) as GitLab Runner Cache",
						Attributes: map[string]schema.Attribute{
							"bucket": schema.StringAttribute{
								Description:         "Name of the bucket in which the cache will be stored",
								MarkdownDescription: "Name of the bucket in which the cache will be stored",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"credentials": schema.StringAttribute{
								Description:         "contains the GCS 'access-id' and 'private-key'",
								MarkdownDescription: "contains the GCS 'access-id' and 'private-key'",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"credentials_file": schema.StringAttribute{
								Description:         "Takes GCS credentials file, 'keys.json'",
								MarkdownDescription: "Takes GCS credentials file, 'keys.json'",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"gitlab_url": schema.StringAttribute{
						Description:         "The fully qualified domain name for the GitLab instance. For example, https://gitlab.example.com",
						MarkdownDescription: "The fully qualified domain name for the GitLab instance. For example, https://gitlab.example.com",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"helper_image": schema.StringAttribute{
						Description:         "If specified, overrides the default GitLab Runner helper image",
						MarkdownDescription: "If specified, overrides the default GitLab Runner helper image",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"image_pull_policy": schema.StringAttribute{
						Description:         "ImagePullPolicy sets the Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
						MarkdownDescription: "ImagePullPolicy sets the Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"interval": schema.Int64Attribute{
						Description:         "Option to define the number of seconds between checks for new jobs. This is set to a default of 30s by operator if not set",
						MarkdownDescription: "Option to define the number of seconds between checks for new jobs. This is set to a default of 30s by operator if not set",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"locked": schema.BoolAttribute{
						Description:         "Specify whether the runner should be locked to a specific project. Defaults to false.",
						MarkdownDescription: "Specify whether the runner should be locked to a specific project. Defaults to false.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"pod_spec": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name is the name given to the custom Pod Spec",
									MarkdownDescription: "Name is the name given to the custom Pod Spec",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"patch": schema.StringAttribute{
									Description:         "A JSON or YAML format string that describes the changes which must be applied to the final PodSpec object before it is generated. You cannot set the patch_path and patch in the same pod_spec configuration, otherwise an error occurs.",
									MarkdownDescription: "A JSON or YAML format string that describes the changes which must be applied to the final PodSpec object before it is generated. You cannot set the patch_path and patch in the same pod_spec configuration, otherwise an error occurs.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"patch_file": schema.StringAttribute{
									Description:         "Path to the file that defines the changes to apply to the final PodSpec object before it is generated. The file must be a JSON or YAML file. You cannot set the patch_path and patch in the same pod_spec configuration, otherwise an error occurs.",
									MarkdownDescription: "Path to the file that defines the changes to apply to the final PodSpec object before it is generated. The file must be a JSON or YAML file. You cannot set the patch_path and patch in the same pod_spec configuration, otherwise an error occurs.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"patch_type": schema.StringAttribute{
									Description:         "The strategy the runner uses to apply the specified changes to the PodSpec object generated by GitLab Runner. The accepted values are merge, json, and strategic (default value).",
									MarkdownDescription: "The strategy the runner uses to apply the specified changes to the PodSpec object generated by GitLab Runner. The accepted values are merge, json, and strategic (default value).",
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

					"protected": schema.BoolAttribute{
						Description:         "Specify whether the runner should only run protected branches. Defaults to false.",
						MarkdownDescription: "Specify whether the runner should only run protected branches. Defaults to false.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"run_untagged": schema.BoolAttribute{
						Description:         "Specify if jobs without tags should be run. If not specified, runner will default to true if no tags were specified. In other case it will default to false.",
						MarkdownDescription: "Specify if jobs without tags should be run. If not specified, runner will default to true if no tags were specified. In other case it will default to false.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"runner_image": schema.StringAttribute{
						Description:         "If specified, overrides the default GitLab Runner image. Default is the Runner image the operator was bundled with.",
						MarkdownDescription: "If specified, overrides the default GitLab Runner image. Default is the Runner image the operator was bundled with.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"s3": schema.SingleNestedAttribute{
						Description:         "options used to setup S3 object store as GitLab Runner Cache",
						MarkdownDescription: "options used to setup S3 object store as GitLab Runner Cache",
						Attributes: map[string]schema.Attribute{
							"bucket": schema.StringAttribute{
								Description:         "Name of the bucket in which the cache will be stored",
								MarkdownDescription: "Name of the bucket in which the cache will be stored",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"credentials": schema.StringAttribute{
								Description:         "Name of the secret containing the 'accesskey' and 'secretkey' used to access the object storage",
								MarkdownDescription: "Name of the secret containing the 'accesskey' and 'secretkey' used to access the object storage",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"insecure": schema.BoolAttribute{
								Description:         "Use insecure connections or HTTP",
								MarkdownDescription: "Use insecure connections or HTTP",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"location": schema.StringAttribute{
								Description:         "Name of the S3 region in use",
								MarkdownDescription: "Name of the S3 region in use",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"server": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"serviceaccount": schema.StringAttribute{
						Description:         "allow user to override service account used by GitLab Runner",
						MarkdownDescription: "allow user to override service account used by GitLab Runner",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"tags": schema.StringAttribute{
						Description:         "List of comma separated tags to be applied to the runner More info: https://docs.gitlab.com/ee/ci/runners/#use-tags-to-limit-the-number-of-jobs-using-the-runner",
						MarkdownDescription: "List of comma separated tags to be applied to the runner More info: https://docs.gitlab.com/ee/ci/runners/#use-tags-to-limit-the-number-of-jobs-using-the-runner",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"token": schema.StringAttribute{
						Description:         "Name of secret containing the 'runner-registration-token' key used to register the runner",
						MarkdownDescription: "Name of secret containing the 'runner-registration-token' key used to register the runner",
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
	}
}

func (r *AppsGitlabComRunnerV1Beta2DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *AppsGitlabComRunnerV1Beta2DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_apps_gitlab_com_runner_v1beta2")

	var data AppsGitlabComRunnerV1Beta2DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "apps.gitlab.com", Version: "v1beta2", Resource: "runners"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
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

	var readResponse AppsGitlabComRunnerV1Beta2DataSourceData
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
	data.ApiVersion = pointer.String("apps.gitlab.com/v1beta2")
	data.Kind = pointer.String("Runner")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
