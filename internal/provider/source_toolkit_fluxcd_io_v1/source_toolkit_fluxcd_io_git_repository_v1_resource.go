/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package source_toolkit_fluxcd_io_v1

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
	"regexp"
	"strings"
)

var (
	_ resource.Resource                = &SourceToolkitFluxcdIoGitRepositoryV1Resource{}
	_ resource.ResourceWithConfigure   = &SourceToolkitFluxcdIoGitRepositoryV1Resource{}
	_ resource.ResourceWithImportState = &SourceToolkitFluxcdIoGitRepositoryV1Resource{}
)

func NewSourceToolkitFluxcdIoGitRepositoryV1Resource() resource.Resource {
	return &SourceToolkitFluxcdIoGitRepositoryV1Resource{}
}

type SourceToolkitFluxcdIoGitRepositoryV1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type SourceToolkitFluxcdIoGitRepositoryV1ResourceData struct {
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
		Ignore  *string `tfsdk:"ignore" json:"ignore,omitempty"`
		Include *[]struct {
			FromPath   *string `tfsdk:"from_path" json:"fromPath,omitempty"`
			Repository *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"repository" json:"repository,omitempty"`
			ToPath *string `tfsdk:"to_path" json:"toPath,omitempty"`
		} `tfsdk:"include" json:"include,omitempty"`
		Interval       *string `tfsdk:"interval" json:"interval,omitempty"`
		ProxySecretRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"proxy_secret_ref" json:"proxySecretRef,omitempty"`
		RecurseSubmodules *bool `tfsdk:"recurse_submodules" json:"recurseSubmodules,omitempty"`
		Ref               *struct {
			Branch *string `tfsdk:"branch" json:"branch,omitempty"`
			Commit *string `tfsdk:"commit" json:"commit,omitempty"`
			Name   *string `tfsdk:"name" json:"name,omitempty"`
			Semver *string `tfsdk:"semver" json:"semver,omitempty"`
			Tag    *string `tfsdk:"tag" json:"tag,omitempty"`
		} `tfsdk:"ref" json:"ref,omitempty"`
		SecretRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		Suspend *bool   `tfsdk:"suspend" json:"suspend,omitempty"`
		Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
		Url     *string `tfsdk:"url" json:"url,omitempty"`
		Verify  *struct {
			Mode      *string `tfsdk:"mode" json:"mode,omitempty"`
			SecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"verify" json:"verify,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SourceToolkitFluxcdIoGitRepositoryV1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_source_toolkit_fluxcd_io_git_repository_v1"
}

func (r *SourceToolkitFluxcdIoGitRepositoryV1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "GitRepository is the Schema for the gitrepositories API.",
		MarkdownDescription: "GitRepository is the Schema for the gitrepositories API.",
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
				Description:         "GitRepositorySpec specifies the required configuration to produce an Artifact for a Git repository.",
				MarkdownDescription: "GitRepositorySpec specifies the required configuration to produce an Artifact for a Git repository.",
				Attributes: map[string]schema.Attribute{
					"ignore": schema.StringAttribute{
						Description:         "Ignore overrides the set of excluded patterns in the .sourceignore format (which is the same as .gitignore). If not provided, a default will be used, consult the documentation for your version to find out what those are.",
						MarkdownDescription: "Ignore overrides the set of excluded patterns in the .sourceignore format (which is the same as .gitignore). If not provided, a default will be used, consult the documentation for your version to find out what those are.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"include": schema.ListNestedAttribute{
						Description:         "Include specifies a list of GitRepository resources which Artifacts should be included in the Artifact produced for this GitRepository.",
						MarkdownDescription: "Include specifies a list of GitRepository resources which Artifacts should be included in the Artifact produced for this GitRepository.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"from_path": schema.StringAttribute{
									Description:         "FromPath specifies the path to copy contents from, defaults to the root of the Artifact.",
									MarkdownDescription: "FromPath specifies the path to copy contents from, defaults to the root of the Artifact.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"repository": schema.SingleNestedAttribute{
									Description:         "GitRepositoryRef specifies the GitRepository which Artifact contents must be included.",
									MarkdownDescription: "GitRepositoryRef specifies the GitRepository which Artifact contents must be included.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the referent.",
											MarkdownDescription: "Name of the referent.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"to_path": schema.StringAttribute{
									Description:         "ToPath specifies the path to copy contents to, defaults to the name of the GitRepositoryRef.",
									MarkdownDescription: "ToPath specifies the path to copy contents to, defaults to the name of the GitRepositoryRef.",
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

					"interval": schema.StringAttribute{
						Description:         "Interval at which the GitRepository URL is checked for updates. This interval is approximate and may be subject to jitter to ensure efficient use of resources.",
						MarkdownDescription: "Interval at which the GitRepository URL is checked for updates. This interval is approximate and may be subject to jitter to ensure efficient use of resources.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m|h))+$`), ""),
						},
					},

					"proxy_secret_ref": schema.SingleNestedAttribute{
						Description:         "ProxySecretRef specifies the Secret containing the proxy configuration to use while communicating with the Git server.",
						MarkdownDescription: "ProxySecretRef specifies the Secret containing the proxy configuration to use while communicating with the Git server.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"recurse_submodules": schema.BoolAttribute{
						Description:         "RecurseSubmodules enables the initialization of all submodules within the GitRepository as cloned from the URL, using their default settings.",
						MarkdownDescription: "RecurseSubmodules enables the initialization of all submodules within the GitRepository as cloned from the URL, using their default settings.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ref": schema.SingleNestedAttribute{
						Description:         "Reference specifies the Git reference to resolve and monitor for changes, defaults to the 'master' branch.",
						MarkdownDescription: "Reference specifies the Git reference to resolve and monitor for changes, defaults to the 'master' branch.",
						Attributes: map[string]schema.Attribute{
							"branch": schema.StringAttribute{
								Description:         "Branch to check out, defaults to 'master' if no other field is defined.",
								MarkdownDescription: "Branch to check out, defaults to 'master' if no other field is defined.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"commit": schema.StringAttribute{
								Description:         "Commit SHA to check out, takes precedence over all reference fields.  This can be combined with Branch to shallow clone the branch, in which the commit is expected to exist.",
								MarkdownDescription: "Commit SHA to check out, takes precedence over all reference fields.  This can be combined with Branch to shallow clone the branch, in which the commit is expected to exist.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the reference to check out; takes precedence over Branch, Tag and SemVer.  It must be a valid Git reference: https://git-scm.com/docs/git-check-ref-format#_description Examples: 'refs/heads/main', 'refs/tags/v0.1.0', 'refs/pull/420/head', 'refs/merge-requests/1/head'",
								MarkdownDescription: "Name of the reference to check out; takes precedence over Branch, Tag and SemVer.  It must be a valid Git reference: https://git-scm.com/docs/git-check-ref-format#_description Examples: 'refs/heads/main', 'refs/tags/v0.1.0', 'refs/pull/420/head', 'refs/merge-requests/1/head'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"semver": schema.StringAttribute{
								Description:         "SemVer tag expression to check out, takes precedence over Tag.",
								MarkdownDescription: "SemVer tag expression to check out, takes precedence over Tag.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tag": schema.StringAttribute{
								Description:         "Tag to check out, takes precedence over Branch.",
								MarkdownDescription: "Tag to check out, takes precedence over Branch.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"secret_ref": schema.SingleNestedAttribute{
						Description:         "SecretRef specifies the Secret containing authentication credentials for the GitRepository. For HTTPS repositories the Secret must contain 'username' and 'password' fields for basic auth or 'bearerToken' field for token auth. For SSH repositories the Secret must contain 'identity' and 'known_hosts' fields.",
						MarkdownDescription: "SecretRef specifies the Secret containing authentication credentials for the GitRepository. For HTTPS repositories the Secret must contain 'username' and 'password' fields for basic auth or 'bearerToken' field for token auth. For SSH repositories the Secret must contain 'identity' and 'known_hosts' fields.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"suspend": schema.BoolAttribute{
						Description:         "Suspend tells the controller to suspend the reconciliation of this GitRepository.",
						MarkdownDescription: "Suspend tells the controller to suspend the reconciliation of this GitRepository.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"timeout": schema.StringAttribute{
						Description:         "Timeout for Git operations like cloning, defaults to 60s.",
						MarkdownDescription: "Timeout for Git operations like cloning, defaults to 60s.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m))+$`), ""),
						},
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

					"verify": schema.SingleNestedAttribute{
						Description:         "Verification specifies the configuration to verify the Git commit signature(s).",
						MarkdownDescription: "Verification specifies the configuration to verify the Git commit signature(s).",
						Attributes: map[string]schema.Attribute{
							"mode": schema.StringAttribute{
								Description:         "Mode specifies which Git object(s) should be verified.  The variants 'head' and 'HEAD' both imply the same thing, i.e. verify the commit that the HEAD of the Git repository points to. The variant 'head' solely exists to ensure backwards compatibility.",
								MarkdownDescription: "Mode specifies which Git object(s) should be verified.  The variants 'head' and 'HEAD' both imply the same thing, i.e. verify the commit that the HEAD of the Git repository points to. The variant 'head' solely exists to ensure backwards compatibility.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("head", "HEAD", "Tag", "TagAndHEAD"),
								},
							},

							"secret_ref": schema.SingleNestedAttribute{
								Description:         "SecretRef specifies the Secret containing the public keys of trusted Git authors.",
								MarkdownDescription: "SecretRef specifies the Secret containing the public keys of trusted Git authors.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent.",
										MarkdownDescription: "Name of the referent.",
										Required:            true,
										Optional:            false,
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *SourceToolkitFluxcdIoGitRepositoryV1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *SourceToolkitFluxcdIoGitRepositoryV1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_source_toolkit_fluxcd_io_git_repository_v1")

	var model SourceToolkitFluxcdIoGitRepositoryV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("source.toolkit.fluxcd.io/v1")
	model.Kind = pointer.String("GitRepository")

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
		Resource(k8sSchema.GroupVersionResource{Group: "source.toolkit.fluxcd.io", Version: "v1", Resource: "gitrepositories"}).
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

	var readResponse SourceToolkitFluxcdIoGitRepositoryV1ResourceData
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

func (r *SourceToolkitFluxcdIoGitRepositoryV1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_source_toolkit_fluxcd_io_git_repository_v1")

	var data SourceToolkitFluxcdIoGitRepositoryV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "source.toolkit.fluxcd.io", Version: "v1", Resource: "gitrepositories"}).
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

	var readResponse SourceToolkitFluxcdIoGitRepositoryV1ResourceData
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

func (r *SourceToolkitFluxcdIoGitRepositoryV1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_source_toolkit_fluxcd_io_git_repository_v1")

	var model SourceToolkitFluxcdIoGitRepositoryV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("source.toolkit.fluxcd.io/v1")
	model.Kind = pointer.String("GitRepository")

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
		Resource(k8sSchema.GroupVersionResource{Group: "source.toolkit.fluxcd.io", Version: "v1", Resource: "gitrepositories"}).
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

	var readResponse SourceToolkitFluxcdIoGitRepositoryV1ResourceData
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

func (r *SourceToolkitFluxcdIoGitRepositoryV1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_source_toolkit_fluxcd_io_git_repository_v1")

	var data SourceToolkitFluxcdIoGitRepositoryV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "source.toolkit.fluxcd.io", Version: "v1", Resource: "gitrepositories"}).
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

func (r *SourceToolkitFluxcdIoGitRepositoryV1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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
