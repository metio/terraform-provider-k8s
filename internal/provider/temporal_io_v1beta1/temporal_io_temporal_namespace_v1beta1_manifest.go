/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package temporal_io_v1beta1

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
	_ datasource.DataSource = &TemporalIoTemporalNamespaceV1Beta1Manifest{}
)

func NewTemporalIoTemporalNamespaceV1Beta1Manifest() datasource.DataSource {
	return &TemporalIoTemporalNamespaceV1Beta1Manifest{}
}

type TemporalIoTemporalNamespaceV1Beta1Manifest struct{}

type TemporalIoTemporalNamespaceV1Beta1ManifestData struct {
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
		ActiveClusterName *string `tfsdk:"active_cluster_name" json:"activeClusterName,omitempty"`
		AllowDeletion     *bool   `tfsdk:"allow_deletion" json:"allowDeletion,omitempty"`
		Archival          *struct {
			History *struct {
				EnableRead *bool   `tfsdk:"enable_read" json:"enableRead,omitempty"`
				Enabled    *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				Path       *string `tfsdk:"path" json:"path,omitempty"`
				Paused     *bool   `tfsdk:"paused" json:"paused,omitempty"`
			} `tfsdk:"history" json:"history,omitempty"`
			Visibility *struct {
				EnableRead *bool   `tfsdk:"enable_read" json:"enableRead,omitempty"`
				Enabled    *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				Path       *string `tfsdk:"path" json:"path,omitempty"`
				Paused     *bool   `tfsdk:"paused" json:"paused,omitempty"`
			} `tfsdk:"visibility" json:"visibility,omitempty"`
		} `tfsdk:"archival" json:"archival,omitempty"`
		ClusterRef *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"cluster_ref" json:"clusterRef,omitempty"`
		Clusters          *[]string          `tfsdk:"clusters" json:"clusters,omitempty"`
		Data              *map[string]string `tfsdk:"data" json:"data,omitempty"`
		Description       *string            `tfsdk:"description" json:"description,omitempty"`
		IsGlobalNamespace *bool              `tfsdk:"is_global_namespace" json:"isGlobalNamespace,omitempty"`
		OwnerEmail        *string            `tfsdk:"owner_email" json:"ownerEmail,omitempty"`
		RetentionPeriod   *string            `tfsdk:"retention_period" json:"retentionPeriod,omitempty"`
		SecurityToken     *string            `tfsdk:"security_token" json:"securityToken,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TemporalIoTemporalNamespaceV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_temporal_io_temporal_namespace_v1beta1_manifest"
}

func (r *TemporalIoTemporalNamespaceV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "A TemporalNamespace creates a namespace in the targeted temporal cluster.",
		MarkdownDescription: "A TemporalNamespace creates a namespace in the targeted temporal cluster.",
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
				Description:         "TemporalNamespaceSpec defines the desired state of Namespace.",
				MarkdownDescription: "TemporalNamespaceSpec defines the desired state of Namespace.",
				Attributes: map[string]schema.Attribute{
					"active_cluster_name": schema.StringAttribute{
						Description:         "The name of active Temporal Cluster.Only applicable if the namespace is a global namespace.",
						MarkdownDescription: "The name of active Temporal Cluster.Only applicable if the namespace is a global namespace.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allow_deletion": schema.BoolAttribute{
						Description:         "AllowDeletion makes the controller delete the Temporal namespace if theCRD is deleted.",
						MarkdownDescription: "AllowDeletion makes the controller delete the Temporal namespace if theCRD is deleted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"archival": schema.SingleNestedAttribute{
						Description:         "Archival is a per-namespace archival configuration.If not set, the default cluster configuration is used.",
						MarkdownDescription: "Archival is a per-namespace archival configuration.If not set, the default cluster configuration is used.",
						Attributes: map[string]schema.Attribute{
							"history": schema.SingleNestedAttribute{
								Description:         "History is the config for this namespace history archival.",
								MarkdownDescription: "History is the config for this namespace history archival.",
								Attributes: map[string]schema.Attribute{
									"enable_read": schema.BoolAttribute{
										Description:         "EnableRead allows temporal to read from the archived Event History.",
										MarkdownDescription: "EnableRead allows temporal to read from the archived Event History.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Enabled defines if the archival is enabled by default for all namespacesor for a particular namespace (depends if it's for a TemporalCluster or a TemporalNamespace).",
										MarkdownDescription: "Enabled defines if the archival is enabled by default for all namespacesor for a particular namespace (depends if it's for a TemporalCluster or a TemporalNamespace).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"path": schema.StringAttribute{
										Description:         "Path is ...",
										MarkdownDescription: "Path is ...",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"paused": schema.BoolAttribute{
										Description:         "Paused defines if the archival is paused.",
										MarkdownDescription: "Paused defines if the archival is paused.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"visibility": schema.SingleNestedAttribute{
								Description:         "Visibility is the config for this namespace visibility archival.",
								MarkdownDescription: "Visibility is the config for this namespace visibility archival.",
								Attributes: map[string]schema.Attribute{
									"enable_read": schema.BoolAttribute{
										Description:         "EnableRead allows temporal to read from the archived Event History.",
										MarkdownDescription: "EnableRead allows temporal to read from the archived Event History.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Enabled defines if the archival is enabled by default for all namespacesor for a particular namespace (depends if it's for a TemporalCluster or a TemporalNamespace).",
										MarkdownDescription: "Enabled defines if the archival is enabled by default for all namespacesor for a particular namespace (depends if it's for a TemporalCluster or a TemporalNamespace).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"path": schema.StringAttribute{
										Description:         "Path is ...",
										MarkdownDescription: "Path is ...",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"paused": schema.BoolAttribute{
										Description:         "Paused defines if the archival is paused.",
										MarkdownDescription: "Paused defines if the archival is paused.",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_ref": schema.SingleNestedAttribute{
						Description:         "Reference to the temporal cluster the namespace will be created.",
						MarkdownDescription: "Reference to the temporal cluster the namespace will be created.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "The name of the temporal object to reference.",
								MarkdownDescription: "The name of the temporal object to reference.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "The namespace of the temporal object to reference.Defaults to the namespace of the requested resource if omitted.",
								MarkdownDescription: "The namespace of the temporal object to reference.Defaults to the namespace of the requested resource if omitted.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"clusters": schema.ListAttribute{
						Description:         "List of clusters names to which the namespace can fail over.Only applicable if the namespace is a global namespace.",
						MarkdownDescription: "List of clusters names to which the namespace can fail over.Only applicable if the namespace is a global namespace.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"data": schema.MapAttribute{
						Description:         "Data is a key-value map for any customized purpose.",
						MarkdownDescription: "Data is a key-value map for any customized purpose.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"description": schema.StringAttribute{
						Description:         "Namespace description.",
						MarkdownDescription: "Namespace description.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"is_global_namespace": schema.BoolAttribute{
						Description:         "IsGlobalNamespace defines whether the namespace is a global namespace.",
						MarkdownDescription: "IsGlobalNamespace defines whether the namespace is a global namespace.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"owner_email": schema.StringAttribute{
						Description:         "Namespace owner email.",
						MarkdownDescription: "Namespace owner email.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"retention_period": schema.StringAttribute{
						Description:         "RetentionPeriod to apply on closed workflow executions.",
						MarkdownDescription: "RetentionPeriod to apply on closed workflow executions.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"security_token": schema.StringAttribute{
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
		},
	}
}

func (r *TemporalIoTemporalNamespaceV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_temporal_io_temporal_namespace_v1beta1_manifest")

	var model TemporalIoTemporalNamespaceV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("temporal.io/v1beta1")
	model.Kind = pointer.String("TemporalNamespace")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
