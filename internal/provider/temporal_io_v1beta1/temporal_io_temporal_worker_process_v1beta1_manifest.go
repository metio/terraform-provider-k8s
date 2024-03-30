/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package temporal_io_v1beta1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &TemporalIoTemporalWorkerProcessV1Beta1Manifest{}
)

func NewTemporalIoTemporalWorkerProcessV1Beta1Manifest() datasource.DataSource {
	return &TemporalIoTemporalWorkerProcessV1Beta1Manifest{}
}

type TemporalIoTemporalWorkerProcessV1Beta1Manifest struct{}

type TemporalIoTemporalWorkerProcessV1Beta1ManifestData struct {
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
		Builder *struct {
			Attempt       *int64  `tfsdk:"attempt" json:"attempt,omitempty"`
			BuildDir      *string `tfsdk:"build_dir" json:"buildDir,omitempty"`
			BuildRegistry *struct {
				PasswordSecretRef *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"password_secret_ref" json:"passwordSecretRef,omitempty"`
				Repository *string `tfsdk:"repository" json:"repository,omitempty"`
				Username   *string `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"build_registry" json:"buildRegistry,omitempty"`
			Enabled       *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			GitRepository *struct {
				Reference *struct {
					Branch *string `tfsdk:"branch" json:"branch,omitempty"`
				} `tfsdk:"reference" json:"reference,omitempty"`
				Url *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"git_repository" json:"gitRepository,omitempty"`
			Image   *string `tfsdk:"image" json:"image,omitempty"`
			Version *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"builder" json:"builder,omitempty"`
		ClusterRef *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"cluster_ref" json:"clusterRef,omitempty"`
		Image            *string `tfsdk:"image" json:"image,omitempty"`
		ImagePullSecrets *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
		JobTtlSecondsAfterFinished *int64  `tfsdk:"job_ttl_seconds_after_finished" json:"jobTtlSecondsAfterFinished,omitempty"`
		PullPolicy                 *string `tfsdk:"pull_policy" json:"pullPolicy,omitempty"`
		Replicas                   *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
		TemporalNamespace          *string `tfsdk:"temporal_namespace" json:"temporalNamespace,omitempty"`
		Version                    *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TemporalIoTemporalWorkerProcessV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_temporal_io_temporal_worker_process_v1beta1_manifest"
}

func (r *TemporalIoTemporalWorkerProcessV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TemporalWorkerProcess is the Schema for the temporalworkerprocesses API.",
		MarkdownDescription: "TemporalWorkerProcess is the Schema for the temporalworkerprocesses API.",
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
				Description:         "TemporalWorkerProcessSpec defines the desired state of TemporalWorkerProcess.",
				MarkdownDescription: "TemporalWorkerProcessSpec defines the desired state of TemporalWorkerProcess.",
				Attributes: map[string]schema.Attribute{
					"builder": schema.SingleNestedAttribute{
						Description:         "Builder is the configuration for building a TemporalWorkerProcess. THIS FEATURE IS HIGHLY EXPERIMENTAL.",
						MarkdownDescription: "Builder is the configuration for building a TemporalWorkerProcess. THIS FEATURE IS HIGHLY EXPERIMENTAL.",
						Attributes: map[string]schema.Attribute{
							"attempt": schema.Int64Attribute{
								Description:         "BuildAttempt is the build attempt number of a given version",
								MarkdownDescription: "BuildAttempt is the build attempt number of a given version",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"build_dir": schema.StringAttribute{
								Description:         "BuildDir is the location of where the sources will be built.",
								MarkdownDescription: "BuildDir is the location of where the sources will be built.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"build_registry": schema.SingleNestedAttribute{
								Description:         "BuildRegistry specifies how to connect to container registry.",
								MarkdownDescription: "BuildRegistry specifies how to connect to container registry.",
								Attributes: map[string]schema.Attribute{
									"password_secret_ref": schema.SingleNestedAttribute{
										Description:         "PasswordSecret is the reference to the secret holding the docker repo password.",
										MarkdownDescription: "PasswordSecret is the reference to the secret holding the docker repo password.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "Key in the Secret.",
												MarkdownDescription: "Key in the Secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the Secret.",
												MarkdownDescription: "Name of the Secret.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"repository": schema.StringAttribute{
										Description:         "Repository is the fqdn to the image repo.",
										MarkdownDescription: "Repository is the fqdn to the image repo.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"username": schema.StringAttribute{
										Description:         "Username is the username for the container repo.",
										MarkdownDescription: "Username is the username for the container repo.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"enabled": schema.BoolAttribute{
								Description:         "Enabled defines if the operator should build the temporal worker process.",
								MarkdownDescription: "Enabled defines if the operator should build the temporal worker process.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"git_repository": schema.SingleNestedAttribute{
								Description:         "GitRepository specifies how to connect to Git source control.",
								MarkdownDescription: "GitRepository specifies how to connect to Git source control.",
								Attributes: map[string]schema.Attribute{
									"reference": schema.SingleNestedAttribute{
										Description:         "Reference specifies the Git reference to resolve and monitor for changes, defaults to the 'master' branch.",
										MarkdownDescription: "Reference specifies the Git reference to resolve and monitor for changes, defaults to the 'master' branch.",
										Attributes: map[string]schema.Attribute{
											"branch": schema.StringAttribute{
												Description:         "Branch to check out, defaults to 'main' if no other field is defined.",
												MarkdownDescription: "Branch to check out, defaults to 'main' if no other field is defined.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"url": schema.StringAttribute{
										Description:         "URL specifies the Git repository URL, it can be an HTTP/S or SSH address.",
										MarkdownDescription: "URL specifies the Git repository URL, it can be an HTTP/S or SSH address.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^(http|https|ssh)://.*$`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"image": schema.StringAttribute{
								Description:         "Image is the image that will be used to build worker image.",
								MarkdownDescription: "Image is the image that will be used to build worker image.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"version": schema.StringAttribute{
								Description:         "Version is the version of the image that will be used to build worker image.",
								MarkdownDescription: "Version is the version of the image that will be used to build worker image.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_ref": schema.SingleNestedAttribute{
						Description:         "Reference to the temporal cluster the worker will connect to.",
						MarkdownDescription: "Reference to the temporal cluster the worker will connect to.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "The name of the TemporalCluster to reference.",
								MarkdownDescription: "The name of the TemporalCluster to reference.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "The namespace of the TemporalCluster to reference. Defaults to the namespace of the requested resource if omitted.",
								MarkdownDescription: "The namespace of the TemporalCluster to reference. Defaults to the namespace of the requested resource if omitted.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"image": schema.StringAttribute{
						Description:         "Image defines the temporal worker docker image the instance should run.",
						MarkdownDescription: "Image defines the temporal worker docker image the instance should run.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"image_pull_secrets": schema.ListNestedAttribute{
						Description:         "An optional list of references to secrets in the same namespace to use for pulling temporal images from registries.",
						MarkdownDescription: "An optional list of references to secrets in the same namespace to use for pulling temporal images from registries.",
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

					"job_ttl_seconds_after_finished": schema.Int64Attribute{
						Description:         "JobTTLSecondsAfterFinished is amount of time to keep job pods after jobs are completed. Defaults to 300 seconds.",
						MarkdownDescription: "JobTTLSecondsAfterFinished is amount of time to keep job pods after jobs are completed. Defaults to 300 seconds.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},

					"pull_policy": schema.StringAttribute{
						Description:         "Image pull policy for determining how to pull worker process images.",
						MarkdownDescription: "Image pull policy for determining how to pull worker process images.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"replicas": schema.Int64Attribute{
						Description:         "Number of desired replicas. Default to 1.",
						MarkdownDescription: "Number of desired replicas. Default to 1.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},

					"temporal_namespace": schema.StringAttribute{
						Description:         "TemporalNamespace that worker will poll.",
						MarkdownDescription: "TemporalNamespace that worker will poll.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"version": schema.StringAttribute{
						Description:         "Version defines the worker process version.",
						MarkdownDescription: "Version defines the worker process version.",
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

func (r *TemporalIoTemporalWorkerProcessV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_temporal_io_temporal_worker_process_v1beta1_manifest")

	var model TemporalIoTemporalWorkerProcessV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("temporal.io/v1beta1")
	model.Kind = pointer.String("TemporalWorkerProcess")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
