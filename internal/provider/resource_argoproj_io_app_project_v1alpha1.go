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

type ArgoprojIoAppProjectV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ArgoprojIoAppProjectV1Alpha1Resource)(nil)
)

type ArgoprojIoAppProjectV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ArgoprojIoAppProjectV1Alpha1GoModel struct {
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
		ClusterResourceWhitelist *[]struct {
			Group *string `tfsdk:"group" yaml:"group,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`
		} `tfsdk:"cluster_resource_whitelist" yaml:"clusterResourceWhitelist,omitempty"`

		Description *string `tfsdk:"description" yaml:"description,omitempty"`

		NamespaceResourceBlacklist *[]struct {
			Group *string `tfsdk:"group" yaml:"group,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`
		} `tfsdk:"namespace_resource_blacklist" yaml:"namespaceResourceBlacklist,omitempty"`

		NamespaceResourceWhitelist *[]struct {
			Group *string `tfsdk:"group" yaml:"group,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`
		} `tfsdk:"namespace_resource_whitelist" yaml:"namespaceResourceWhitelist,omitempty"`

		Roles *[]struct {
			Description *string `tfsdk:"description" yaml:"description,omitempty"`

			Groups *[]string `tfsdk:"groups" yaml:"groups,omitempty"`

			JwtTokens *[]struct {
				Exp *int64 `tfsdk:"exp" yaml:"exp,omitempty"`

				Iat *int64 `tfsdk:"iat" yaml:"iat,omitempty"`

				Id *string `tfsdk:"id" yaml:"id,omitempty"`
			} `tfsdk:"jwt_tokens" yaml:"jwtTokens,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Policies *[]string `tfsdk:"policies" yaml:"policies,omitempty"`
		} `tfsdk:"roles" yaml:"roles,omitempty"`

		ClusterResourceBlacklist *[]struct {
			Group *string `tfsdk:"group" yaml:"group,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`
		} `tfsdk:"cluster_resource_blacklist" yaml:"clusterResourceBlacklist,omitempty"`

		OrphanedResources *struct {
			Ignore *[]struct {
				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Group *string `tfsdk:"group" yaml:"group,omitempty"`
			} `tfsdk:"ignore" yaml:"ignore,omitempty"`

			Warn *bool `tfsdk:"warn" yaml:"warn,omitempty"`
		} `tfsdk:"orphaned_resources" yaml:"orphanedResources,omitempty"`

		PermitOnlyProjectScopedClusters *bool `tfsdk:"permit_only_project_scoped_clusters" yaml:"permitOnlyProjectScopedClusters,omitempty"`

		SignatureKeys *[]struct {
			KeyID *string `tfsdk:"key_id" yaml:"keyID,omitempty"`
		} `tfsdk:"signature_keys" yaml:"signatureKeys,omitempty"`

		SourceNamespaces *[]string `tfsdk:"source_namespaces" yaml:"sourceNamespaces,omitempty"`

		SourceRepos *[]string `tfsdk:"source_repos" yaml:"sourceRepos,omitempty"`

		SyncWindows *[]struct {
			ManualSync *bool `tfsdk:"manual_sync" yaml:"manualSync,omitempty"`

			Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

			Schedule *string `tfsdk:"schedule" yaml:"schedule,omitempty"`

			TimeZone *string `tfsdk:"time_zone" yaml:"timeZone,omitempty"`

			Applications *[]string `tfsdk:"applications" yaml:"applications,omitempty"`

			Clusters *[]string `tfsdk:"clusters" yaml:"clusters,omitempty"`

			Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`
		} `tfsdk:"sync_windows" yaml:"syncWindows,omitempty"`

		Destinations *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

			Server *string `tfsdk:"server" yaml:"server,omitempty"`
		} `tfsdk:"destinations" yaml:"destinations,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewArgoprojIoAppProjectV1Alpha1Resource() resource.Resource {
	return &ArgoprojIoAppProjectV1Alpha1Resource{}
}

func (r *ArgoprojIoAppProjectV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_argoproj_io_app_project_v1alpha1"
}

func (r *ArgoprojIoAppProjectV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "AppProject provides a logical grouping of applications, providing controls for: * where the apps may deploy to (cluster whitelist) * what may be deployed (repository whitelist, resource whitelist/blacklist) * who can access these applications (roles, OIDC group claims bindings) * and what they can do (RBAC policies) * automation access to these roles (JWT tokens)",
		MarkdownDescription: "AppProject provides a logical grouping of applications, providing controls for: * where the apps may deploy to (cluster whitelist) * what may be deployed (repository whitelist, resource whitelist/blacklist) * who can access these applications (roles, OIDC group claims bindings) * and what they can do (RBAC policies) * automation access to these roles (JWT tokens)",
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
				Description:         "AppProjectSpec is the specification of an AppProject",
				MarkdownDescription: "AppProjectSpec is the specification of an AppProject",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"cluster_resource_whitelist": {
						Description:         "ClusterResourceWhitelist contains list of whitelisted cluster level resources",
						MarkdownDescription: "ClusterResourceWhitelist contains list of whitelisted cluster level resources",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"group": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"kind": {
								Description:         "",
								MarkdownDescription: "",

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

					"description": {
						Description:         "Description contains optional project description",
						MarkdownDescription: "Description contains optional project description",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"namespace_resource_blacklist": {
						Description:         "NamespaceResourceBlacklist contains list of blacklisted namespace level resources",
						MarkdownDescription: "NamespaceResourceBlacklist contains list of blacklisted namespace level resources",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"group": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"kind": {
								Description:         "",
								MarkdownDescription: "",

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

					"namespace_resource_whitelist": {
						Description:         "NamespaceResourceWhitelist contains list of whitelisted namespace level resources",
						MarkdownDescription: "NamespaceResourceWhitelist contains list of whitelisted namespace level resources",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"group": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"kind": {
								Description:         "",
								MarkdownDescription: "",

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

					"roles": {
						Description:         "Roles are user defined RBAC roles associated with this project",
						MarkdownDescription: "Roles are user defined RBAC roles associated with this project",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"description": {
								Description:         "Description is a description of the role",
								MarkdownDescription: "Description is a description of the role",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"groups": {
								Description:         "Groups are a list of OIDC group claims bound to this role",
								MarkdownDescription: "Groups are a list of OIDC group claims bound to this role",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"jwt_tokens": {
								Description:         "JWTTokens are a list of generated JWT tokens bound to this role",
								MarkdownDescription: "JWTTokens are a list of generated JWT tokens bound to this role",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"exp": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"iat": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"id": {
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

							"name": {
								Description:         "Name is a name for this role",
								MarkdownDescription: "Name is a name for this role",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"policies": {
								Description:         "Policies Stores a list of casbin formatted strings that define access policies for the role in the project",
								MarkdownDescription: "Policies Stores a list of casbin formatted strings that define access policies for the role in the project",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_resource_blacklist": {
						Description:         "ClusterResourceBlacklist contains list of blacklisted cluster level resources",
						MarkdownDescription: "ClusterResourceBlacklist contains list of blacklisted cluster level resources",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"group": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"kind": {
								Description:         "",
								MarkdownDescription: "",

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

					"orphaned_resources": {
						Description:         "OrphanedResources specifies if controller should monitor orphaned resources of apps in this project",
						MarkdownDescription: "OrphanedResources specifies if controller should monitor orphaned resources of apps in this project",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"ignore": {
								Description:         "Ignore contains a list of resources that are to be excluded from orphaned resources monitoring",
								MarkdownDescription: "Ignore contains a list of resources that are to be excluded from orphaned resources monitoring",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"kind": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"group": {
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

							"warn": {
								Description:         "Warn indicates if warning condition should be created for apps which have orphaned resources",
								MarkdownDescription: "Warn indicates if warning condition should be created for apps which have orphaned resources",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"permit_only_project_scoped_clusters": {
						Description:         "PermitOnlyProjectScopedClusters determines whether destinations can only reference clusters which are project-scoped",
						MarkdownDescription: "PermitOnlyProjectScopedClusters determines whether destinations can only reference clusters which are project-scoped",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"signature_keys": {
						Description:         "SignatureKeys contains a list of PGP key IDs that commits in Git must be signed with in order to be allowed for sync",
						MarkdownDescription: "SignatureKeys contains a list of PGP key IDs that commits in Git must be signed with in order to be allowed for sync",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"key_id": {
								Description:         "The ID of the key in hexadecimal notation",
								MarkdownDescription: "The ID of the key in hexadecimal notation",

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

					"source_namespaces": {
						Description:         "SourceNamespaces defines the namespaces application resources are allowed to be created in",
						MarkdownDescription: "SourceNamespaces defines the namespaces application resources are allowed to be created in",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"source_repos": {
						Description:         "SourceRepos contains list of repository URLs which can be used for deployment",
						MarkdownDescription: "SourceRepos contains list of repository URLs which can be used for deployment",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"sync_windows": {
						Description:         "SyncWindows controls when syncs can be run for apps in this project",
						MarkdownDescription: "SyncWindows controls when syncs can be run for apps in this project",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"manual_sync": {
								Description:         "ManualSync enables manual syncs when they would otherwise be blocked",
								MarkdownDescription: "ManualSync enables manual syncs when they would otherwise be blocked",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespaces": {
								Description:         "Namespaces contains a list of namespaces that the window will apply to",
								MarkdownDescription: "Namespaces contains a list of namespaces that the window will apply to",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"schedule": {
								Description:         "Schedule is the time the window will begin, specified in cron format",
								MarkdownDescription: "Schedule is the time the window will begin, specified in cron format",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"time_zone": {
								Description:         "TimeZone of the sync that will be applied to the schedule",
								MarkdownDescription: "TimeZone of the sync that will be applied to the schedule",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"applications": {
								Description:         "Applications contains a list of applications that the window will apply to",
								MarkdownDescription: "Applications contains a list of applications that the window will apply to",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"clusters": {
								Description:         "Clusters contains a list of clusters that the window will apply to",
								MarkdownDescription: "Clusters contains a list of clusters that the window will apply to",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"duration": {
								Description:         "Duration is the amount of time the sync window will be open",
								MarkdownDescription: "Duration is the amount of time the sync window will be open",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kind": {
								Description:         "Kind defines if the window allows or blocks syncs",
								MarkdownDescription: "Kind defines if the window allows or blocks syncs",

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

					"destinations": {
						Description:         "Destinations contains list of destinations available for deployment",
						MarkdownDescription: "Destinations contains list of destinations available for deployment",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name is an alternate way of specifying the target cluster by its symbolic name",
								MarkdownDescription: "Name is an alternate way of specifying the target cluster by its symbolic name",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace specifies the target namespace for the application's resources. The namespace will only be set for namespace-scoped resources that have not set a value for .metadata.namespace",
								MarkdownDescription: "Namespace specifies the target namespace for the application's resources. The namespace will only be set for namespace-scoped resources that have not set a value for .metadata.namespace",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"server": {
								Description:         "Server specifies the URL of the target cluster and must be set to the Kubernetes control plane API",
								MarkdownDescription: "Server specifies the URL of the target cluster and must be set to the Kubernetes control plane API",

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
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *ArgoprojIoAppProjectV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_argoproj_io_app_project_v1alpha1")

	var state ArgoprojIoAppProjectV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ArgoprojIoAppProjectV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("argoproj.io/v1alpha1")
	goModel.Kind = utilities.Ptr("AppProject")

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

func (r *ArgoprojIoAppProjectV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_argoproj_io_app_project_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ArgoprojIoAppProjectV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_argoproj_io_app_project_v1alpha1")

	var state ArgoprojIoAppProjectV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ArgoprojIoAppProjectV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("argoproj.io/v1alpha1")
	goModel.Kind = utilities.Ptr("AppProject")

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

func (r *ArgoprojIoAppProjectV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_argoproj_io_app_project_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
