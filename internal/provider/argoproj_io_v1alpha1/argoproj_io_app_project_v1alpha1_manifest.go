/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package argoproj_io_v1alpha1

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
	_ datasource.DataSource = &ArgoprojIoAppProjectV1Alpha1Manifest{}
)

func NewArgoprojIoAppProjectV1Alpha1Manifest() datasource.DataSource {
	return &ArgoprojIoAppProjectV1Alpha1Manifest{}
}

type ArgoprojIoAppProjectV1Alpha1Manifest struct{}

type ArgoprojIoAppProjectV1Alpha1ManifestData struct {
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
		ClusterResourceBlacklist *[]struct {
			Group *string `tfsdk:"group" json:"group,omitempty"`
			Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
		} `tfsdk:"cluster_resource_blacklist" json:"clusterResourceBlacklist,omitempty"`
		ClusterResourceWhitelist *[]struct {
			Group *string `tfsdk:"group" json:"group,omitempty"`
			Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
		} `tfsdk:"cluster_resource_whitelist" json:"clusterResourceWhitelist,omitempty"`
		Description  *string `tfsdk:"description" json:"description,omitempty"`
		Destinations *[]struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Server    *string `tfsdk:"server" json:"server,omitempty"`
		} `tfsdk:"destinations" json:"destinations,omitempty"`
		NamespaceResourceBlacklist *[]struct {
			Group *string `tfsdk:"group" json:"group,omitempty"`
			Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
		} `tfsdk:"namespace_resource_blacklist" json:"namespaceResourceBlacklist,omitempty"`
		NamespaceResourceWhitelist *[]struct {
			Group *string `tfsdk:"group" json:"group,omitempty"`
			Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
		} `tfsdk:"namespace_resource_whitelist" json:"namespaceResourceWhitelist,omitempty"`
		OrphanedResources *struct {
			Ignore *[]struct {
				Group *string `tfsdk:"group" json:"group,omitempty"`
				Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"ignore" json:"ignore,omitempty"`
			Warn *bool `tfsdk:"warn" json:"warn,omitempty"`
		} `tfsdk:"orphaned_resources" json:"orphanedResources,omitempty"`
		PermitOnlyProjectScopedClusters *bool `tfsdk:"permit_only_project_scoped_clusters" json:"permitOnlyProjectScopedClusters,omitempty"`
		Roles                           *[]struct {
			Description *string   `tfsdk:"description" json:"description,omitempty"`
			Groups      *[]string `tfsdk:"groups" json:"groups,omitempty"`
			JwtTokens   *[]struct {
				Exp *int64  `tfsdk:"exp" json:"exp,omitempty"`
				Iat *int64  `tfsdk:"iat" json:"iat,omitempty"`
				Id  *string `tfsdk:"id" json:"id,omitempty"`
			} `tfsdk:"jwt_tokens" json:"jwtTokens,omitempty"`
			Name     *string   `tfsdk:"name" json:"name,omitempty"`
			Policies *[]string `tfsdk:"policies" json:"policies,omitempty"`
		} `tfsdk:"roles" json:"roles,omitempty"`
		SignatureKeys *[]struct {
			KeyID *string `tfsdk:"key_id" json:"keyID,omitempty"`
		} `tfsdk:"signature_keys" json:"signatureKeys,omitempty"`
		SourceNamespaces *[]string `tfsdk:"source_namespaces" json:"sourceNamespaces,omitempty"`
		SourceRepos      *[]string `tfsdk:"source_repos" json:"sourceRepos,omitempty"`
		SyncWindows      *[]struct {
			Applications *[]string `tfsdk:"applications" json:"applications,omitempty"`
			Clusters     *[]string `tfsdk:"clusters" json:"clusters,omitempty"`
			Duration     *string   `tfsdk:"duration" json:"duration,omitempty"`
			Kind         *string   `tfsdk:"kind" json:"kind,omitempty"`
			ManualSync   *bool     `tfsdk:"manual_sync" json:"manualSync,omitempty"`
			Namespaces   *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
			Schedule     *string   `tfsdk:"schedule" json:"schedule,omitempty"`
			TimeZone     *string   `tfsdk:"time_zone" json:"timeZone,omitempty"`
		} `tfsdk:"sync_windows" json:"syncWindows,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ArgoprojIoAppProjectV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_argoproj_io_app_project_v1alpha1_manifest"
}

func (r *ArgoprojIoAppProjectV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AppProject provides a logical grouping of applications, providing controls for: * where the apps may deploy to (cluster whitelist) * what may be deployed (repository whitelist, resource whitelist/blacklist) * who can access these applications (roles, OIDC group claims bindings) * and what they can do (RBAC policies) * automation access to these roles (JWT tokens)",
		MarkdownDescription: "AppProject provides a logical grouping of applications, providing controls for: * where the apps may deploy to (cluster whitelist) * what may be deployed (repository whitelist, resource whitelist/blacklist) * who can access these applications (roles, OIDC group claims bindings) * and what they can do (RBAC policies) * automation access to these roles (JWT tokens)",
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
				Description:         "AppProjectSpec is the specification of an AppProject",
				MarkdownDescription: "AppProjectSpec is the specification of an AppProject",
				Attributes: map[string]schema.Attribute{
					"cluster_resource_blacklist": schema.ListNestedAttribute{
						Description:         "ClusterResourceBlacklist contains list of blacklisted cluster level resources",
						MarkdownDescription: "ClusterResourceBlacklist contains list of blacklisted cluster level resources",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"cluster_resource_whitelist": schema.ListNestedAttribute{
						Description:         "ClusterResourceWhitelist contains list of whitelisted cluster level resources",
						MarkdownDescription: "ClusterResourceWhitelist contains list of whitelisted cluster level resources",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"description": schema.StringAttribute{
						Description:         "Description contains optional project description",
						MarkdownDescription: "Description contains optional project description",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"destinations": schema.ListNestedAttribute{
						Description:         "Destinations contains list of destinations available for deployment",
						MarkdownDescription: "Destinations contains list of destinations available for deployment",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name is an alternate way of specifying the target cluster by its symbolic name. This must be set if Server is not set.",
									MarkdownDescription: "Name is an alternate way of specifying the target cluster by its symbolic name. This must be set if Server is not set.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace specifies the target namespace for the application's resources. The namespace will only be set for namespace-scoped resources that have not set a value for .metadata.namespace",
									MarkdownDescription: "Namespace specifies the target namespace for the application's resources. The namespace will only be set for namespace-scoped resources that have not set a value for .metadata.namespace",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"server": schema.StringAttribute{
									Description:         "Server specifies the URL of the target cluster's Kubernetes control plane API. This must be set if Name is not set.",
									MarkdownDescription: "Server specifies the URL of the target cluster's Kubernetes control plane API. This must be set if Name is not set.",
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

					"namespace_resource_blacklist": schema.ListNestedAttribute{
						Description:         "NamespaceResourceBlacklist contains list of blacklisted namespace level resources",
						MarkdownDescription: "NamespaceResourceBlacklist contains list of blacklisted namespace level resources",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"namespace_resource_whitelist": schema.ListNestedAttribute{
						Description:         "NamespaceResourceWhitelist contains list of whitelisted namespace level resources",
						MarkdownDescription: "NamespaceResourceWhitelist contains list of whitelisted namespace level resources",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"orphaned_resources": schema.SingleNestedAttribute{
						Description:         "OrphanedResources specifies if controller should monitor orphaned resources of apps in this project",
						MarkdownDescription: "OrphanedResources specifies if controller should monitor orphaned resources of apps in this project",
						Attributes: map[string]schema.Attribute{
							"ignore": schema.ListNestedAttribute{
								Description:         "Ignore contains a list of resources that are to be excluded from orphaned resources monitoring",
								MarkdownDescription: "Ignore contains a list of resources that are to be excluded from orphaned resources monitoring",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"group": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kind": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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

							"warn": schema.BoolAttribute{
								Description:         "Warn indicates if warning condition should be created for apps which have orphaned resources",
								MarkdownDescription: "Warn indicates if warning condition should be created for apps which have orphaned resources",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"permit_only_project_scoped_clusters": schema.BoolAttribute{
						Description:         "PermitOnlyProjectScopedClusters determines whether destinations can only reference clusters which are project-scoped",
						MarkdownDescription: "PermitOnlyProjectScopedClusters determines whether destinations can only reference clusters which are project-scoped",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"roles": schema.ListNestedAttribute{
						Description:         "Roles are user defined RBAC roles associated with this project",
						MarkdownDescription: "Roles are user defined RBAC roles associated with this project",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"description": schema.StringAttribute{
									Description:         "Description is a description of the role",
									MarkdownDescription: "Description is a description of the role",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"groups": schema.ListAttribute{
									Description:         "Groups are a list of OIDC group claims bound to this role",
									MarkdownDescription: "Groups are a list of OIDC group claims bound to this role",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"jwt_tokens": schema.ListNestedAttribute{
									Description:         "JWTTokens are a list of generated JWT tokens bound to this role",
									MarkdownDescription: "JWTTokens are a list of generated JWT tokens bound to this role",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"exp": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"iat": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"id": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
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

								"name": schema.StringAttribute{
									Description:         "Name is a name for this role",
									MarkdownDescription: "Name is a name for this role",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"policies": schema.ListAttribute{
									Description:         "Policies Stores a list of casbin formatted strings that define access policies for the role in the project",
									MarkdownDescription: "Policies Stores a list of casbin formatted strings that define access policies for the role in the project",
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

					"signature_keys": schema.ListNestedAttribute{
						Description:         "SignatureKeys contains a list of PGP key IDs that commits in Git must be signed with in order to be allowed for sync",
						MarkdownDescription: "SignatureKeys contains a list of PGP key IDs that commits in Git must be signed with in order to be allowed for sync",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key_id": schema.StringAttribute{
									Description:         "The ID of the key in hexadecimal notation",
									MarkdownDescription: "The ID of the key in hexadecimal notation",
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

					"source_namespaces": schema.ListAttribute{
						Description:         "SourceNamespaces defines the namespaces application resources are allowed to be created in",
						MarkdownDescription: "SourceNamespaces defines the namespaces application resources are allowed to be created in",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"source_repos": schema.ListAttribute{
						Description:         "SourceRepos contains list of repository URLs which can be used for deployment",
						MarkdownDescription: "SourceRepos contains list of repository URLs which can be used for deployment",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"sync_windows": schema.ListNestedAttribute{
						Description:         "SyncWindows controls when syncs can be run for apps in this project",
						MarkdownDescription: "SyncWindows controls when syncs can be run for apps in this project",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"applications": schema.ListAttribute{
									Description:         "Applications contains a list of applications that the window will apply to",
									MarkdownDescription: "Applications contains a list of applications that the window will apply to",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"clusters": schema.ListAttribute{
									Description:         "Clusters contains a list of clusters that the window will apply to",
									MarkdownDescription: "Clusters contains a list of clusters that the window will apply to",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"duration": schema.StringAttribute{
									Description:         "Duration is the amount of time the sync window will be open",
									MarkdownDescription: "Duration is the amount of time the sync window will be open",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
									Description:         "Kind defines if the window allows or blocks syncs",
									MarkdownDescription: "Kind defines if the window allows or blocks syncs",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"manual_sync": schema.BoolAttribute{
									Description:         "ManualSync enables manual syncs when they would otherwise be blocked",
									MarkdownDescription: "ManualSync enables manual syncs when they would otherwise be blocked",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"namespaces": schema.ListAttribute{
									Description:         "Namespaces contains a list of namespaces that the window will apply to",
									MarkdownDescription: "Namespaces contains a list of namespaces that the window will apply to",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"schedule": schema.StringAttribute{
									Description:         "Schedule is the time the window will begin, specified in cron format",
									MarkdownDescription: "Schedule is the time the window will begin, specified in cron format",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"time_zone": schema.StringAttribute{
									Description:         "TimeZone of the sync that will be applied to the schedule",
									MarkdownDescription: "TimeZone of the sync that will be applied to the schedule",
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
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ArgoprojIoAppProjectV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_argoproj_io_app_project_v1alpha1_manifest")

	var model ArgoprojIoAppProjectV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("argoproj.io/v1alpha1")
	model.Kind = pointer.String("AppProject")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
