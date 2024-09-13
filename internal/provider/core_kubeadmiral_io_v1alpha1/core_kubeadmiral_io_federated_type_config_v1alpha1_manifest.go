/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package core_kubeadmiral_io_v1alpha1

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
	_ datasource.DataSource = &CoreKubeadmiralIoFederatedTypeConfigV1Alpha1Manifest{}
)

func NewCoreKubeadmiralIoFederatedTypeConfigV1Alpha1Manifest() datasource.DataSource {
	return &CoreKubeadmiralIoFederatedTypeConfigV1Alpha1Manifest{}
}

type CoreKubeadmiralIoFederatedTypeConfigV1Alpha1Manifest struct{}

type CoreKubeadmiralIoFederatedTypeConfigV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AutoMigration *struct {
			Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
		} `tfsdk:"auto_migration" json:"autoMigration,omitempty"`
		Controllers    *[]string `tfsdk:"controllers" json:"controllers,omitempty"`
		PathDefinition *struct {
			AvailableReplicasStatus *string `tfsdk:"available_replicas_status" json:"availableReplicasStatus,omitempty"`
			LabelSelector           *string `tfsdk:"label_selector" json:"labelSelector,omitempty"`
			ReadyReplicasStatus     *string `tfsdk:"ready_replicas_status" json:"readyReplicasStatus,omitempty"`
			ReplicasSpec            *string `tfsdk:"replicas_spec" json:"replicasSpec,omitempty"`
			ReplicasStatus          *string `tfsdk:"replicas_status" json:"replicasStatus,omitempty"`
		} `tfsdk:"path_definition" json:"pathDefinition,omitempty"`
		SourceType *struct {
			Group      *string `tfsdk:"group" json:"group,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			PluralName *string `tfsdk:"plural_name" json:"pluralName,omitempty"`
			Scope      *string `tfsdk:"scope" json:"scope,omitempty"`
			Version    *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"source_type" json:"sourceType,omitempty"`
		StatusAggregation *struct {
			Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
		} `tfsdk:"status_aggregation" json:"statusAggregation,omitempty"`
		StatusCollection *struct {
			Enabled *bool     `tfsdk:"enabled" json:"enabled,omitempty"`
			Fields  *[]string `tfsdk:"fields" json:"fields,omitempty"`
		} `tfsdk:"status_collection" json:"statusCollection,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CoreKubeadmiralIoFederatedTypeConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_core_kubeadmiral_io_federated_type_config_v1alpha1_manifest"
}

func (r *CoreKubeadmiralIoFederatedTypeConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "FederatedTypeConfig specifies an API resource type to federate and various type-specific options.",
		MarkdownDescription: "FederatedTypeConfig specifies an API resource type to federate and various type-specific options.",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"auto_migration": schema.SingleNestedAttribute{
						Description:         "Configuration for AutoMigration. If left empty, the AutoMigration feature will be disabled.",
						MarkdownDescription: "Configuration for AutoMigration. If left empty, the AutoMigration feature will be disabled.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Whether or not to automatically migrate unschedulable pods to a different cluster.",
								MarkdownDescription: "Whether or not to automatically migrate unschedulable pods to a different cluster.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"controllers": schema.ListAttribute{
						Description:         "The controllers that must run before the source object can be propagated to member clusters. Each inner slice specifies a step. Step T must complete before step T+1 can commence. Controllers within each step can execute in parallel.",
						MarkdownDescription: "The controllers that must run before the source object can be propagated to member clusters. Each inner slice specifies a step. Step T must complete before step T+1 can commence. Controllers within each step can execute in parallel.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"path_definition": schema.SingleNestedAttribute{
						Description:         "Defines the paths to various fields in the target object's schema.",
						MarkdownDescription: "Defines the paths to various fields in the target object's schema.",
						Attributes: map[string]schema.Attribute{
							"available_replicas_status": schema.StringAttribute{
								Description:         "Path to a numeric field that reflects the number of available replicas that the object currently has. E.g. 'status.availableReplicas' for Deployment and ReplicaSet.",
								MarkdownDescription: "Path to a numeric field that reflects the number of available replicas that the object currently has. E.g. 'status.availableReplicas' for Deployment and ReplicaSet.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"label_selector": schema.StringAttribute{
								Description:         "Path to a metav1.LabelSelector field that selects the replicas for this object. E.g. 'spec.selector' for Deployment and ReplicaSet.",
								MarkdownDescription: "Path to a metav1.LabelSelector field that selects the replicas for this object. E.g. 'spec.selector' for Deployment and ReplicaSet.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ready_replicas_status": schema.StringAttribute{
								Description:         "Path to a numeric field that reflects the number of ready replicas that the object currently has. E.g. 'status.readyReplicas' for Deployment and ReplicaSet.",
								MarkdownDescription: "Path to a numeric field that reflects the number of ready replicas that the object currently has. E.g. 'status.readyReplicas' for Deployment and ReplicaSet.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replicas_spec": schema.StringAttribute{
								Description:         "Path to a numeric field that indicates the number of replicas that an object can be divided into. E.g. 'spec.replicas' for Deployment and ReplicaSet.",
								MarkdownDescription: "Path to a numeric field that indicates the number of replicas that an object can be divided into. E.g. 'spec.replicas' for Deployment and ReplicaSet.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replicas_status": schema.StringAttribute{
								Description:         "Path to a numeric field that reflects the number of replicas that the object currently has. E.g. 'status.replicas' for Deployment and ReplicaSet.",
								MarkdownDescription: "Path to a numeric field that reflects the number of replicas that the object currently has. E.g. 'status.replicas' for Deployment and ReplicaSet.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"source_type": schema.SingleNestedAttribute{
						Description:         "The API resource type to be federated.",
						MarkdownDescription: "The API resource type to be federated.",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "Group of the resource.",
								MarkdownDescription: "Group of the resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the resource.",
								MarkdownDescription: "Kind of the resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"plural_name": schema.StringAttribute{
								Description:         "Lower-cased plural name of the resource (e.g. configmaps). If not provided, it will be computed by lower-casing the kind and suffixing an 's'.",
								MarkdownDescription: "Lower-cased plural name of the resource (e.g. configmaps). If not provided, it will be computed by lower-casing the kind and suffixing an 's'.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"scope": schema.StringAttribute{
								Description:         "Scope of the resource.",
								MarkdownDescription: "Scope of the resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"version": schema.StringAttribute{
								Description:         "Version of the resource.",
								MarkdownDescription: "Version of the resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"status_aggregation": schema.SingleNestedAttribute{
						Description:         "Configuration for StatusAggregation.",
						MarkdownDescription: "Configuration for StatusAggregation.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Whether or not to enable status aggregation.",
								MarkdownDescription: "Whether or not to enable status aggregation.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"status_collection": schema.SingleNestedAttribute{
						Description:         "Configuration for StatusCollection. If left empty, the StatusCollection feature will be disabled.",
						MarkdownDescription: "Configuration for StatusCollection. If left empty, the StatusCollection feature will be disabled.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Whether or not to enable status collection.",
								MarkdownDescription: "Whether or not to enable status collection.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"fields": schema.ListAttribute{
								Description:         "Contains the fields to be collected during status collection. Each field is a dot separated string that corresponds to its path in the source object's schema. E.g. 'metadata.creationTimestamp'.",
								MarkdownDescription: "Contains the fields to be collected during status collection. Each field is a dot separated string that corresponds to its path in the source object's schema. E.g. 'metadata.creationTimestamp'.",
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *CoreKubeadmiralIoFederatedTypeConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_core_kubeadmiral_io_federated_type_config_v1alpha1_manifest")

	var model CoreKubeadmiralIoFederatedTypeConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("core.kubeadmiral.io/v1alpha1")
	model.Kind = pointer.String("FederatedTypeConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
