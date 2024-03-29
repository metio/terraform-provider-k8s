/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package config_openshift_io_v1

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
	_ datasource.DataSource = &ConfigOpenshiftIoClusterVersionV1Manifest{}
)

func NewConfigOpenshiftIoClusterVersionV1Manifest() datasource.DataSource {
	return &ConfigOpenshiftIoClusterVersionV1Manifest{}
}

type ConfigOpenshiftIoClusterVersionV1Manifest struct{}

type ConfigOpenshiftIoClusterVersionV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Capabilities *struct {
			AdditionalEnabledCapabilities *[]string `tfsdk:"additional_enabled_capabilities" json:"additionalEnabledCapabilities,omitempty"`
			BaselineCapabilitySet         *string   `tfsdk:"baseline_capability_set" json:"baselineCapabilitySet,omitempty"`
		} `tfsdk:"capabilities" json:"capabilities,omitempty"`
		Channel       *string `tfsdk:"channel" json:"channel,omitempty"`
		ClusterID     *string `tfsdk:"cluster_id" json:"clusterID,omitempty"`
		DesiredUpdate *struct {
			Architecture *string `tfsdk:"architecture" json:"architecture,omitempty"`
			Force        *bool   `tfsdk:"force" json:"force,omitempty"`
			Image        *string `tfsdk:"image" json:"image,omitempty"`
			Version      *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"desired_update" json:"desiredUpdate,omitempty"`
		Overrides *[]struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Unmanaged *bool   `tfsdk:"unmanaged" json:"unmanaged,omitempty"`
		} `tfsdk:"overrides" json:"overrides,omitempty"`
		Upstream *string `tfsdk:"upstream" json:"upstream,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ConfigOpenshiftIoClusterVersionV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_config_openshift_io_cluster_version_v1_manifest"
}

func (r *ConfigOpenshiftIoClusterVersionV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterVersion is the configuration for the ClusterVersionOperator. This is where parameters related to automatic updates can be set.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "ClusterVersion is the configuration for the ClusterVersionOperator. This is where parameters related to automatic updates can be set.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "spec is the desired state of the cluster version - the operator will work to ensure that the desired version is applied to the cluster.",
				MarkdownDescription: "spec is the desired state of the cluster version - the operator will work to ensure that the desired version is applied to the cluster.",
				Attributes: map[string]schema.Attribute{
					"capabilities": schema.SingleNestedAttribute{
						Description:         "capabilities configures the installation of optional, core cluster components.  A null value here is identical to an empty object; see the child properties for default semantics.",
						MarkdownDescription: "capabilities configures the installation of optional, core cluster components.  A null value here is identical to an empty object; see the child properties for default semantics.",
						Attributes: map[string]schema.Attribute{
							"additional_enabled_capabilities": schema.ListAttribute{
								Description:         "additionalEnabledCapabilities extends the set of managed capabilities beyond the baseline defined in baselineCapabilitySet.  The default is an empty set.",
								MarkdownDescription: "additionalEnabledCapabilities extends the set of managed capabilities beyond the baseline defined in baselineCapabilitySet.  The default is an empty set.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"baseline_capability_set": schema.StringAttribute{
								Description:         "baselineCapabilitySet selects an initial set of optional capabilities to enable, which can be extended via additionalEnabledCapabilities.  If unset, the cluster will choose a default, and the default may change over time. The current default is vCurrent.",
								MarkdownDescription: "baselineCapabilitySet selects an initial set of optional capabilities to enable, which can be extended via additionalEnabledCapabilities.  If unset, the cluster will choose a default, and the default may change over time. The current default is vCurrent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("None", "v4.11", "v4.12", "v4.13", "v4.14", "v4.15", "vCurrent"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"channel": schema.StringAttribute{
						Description:         "channel is an identifier for explicitly requesting that a non-default set of updates be applied to this cluster. The default channel will be contain stable updates that are appropriate for production clusters.",
						MarkdownDescription: "channel is an identifier for explicitly requesting that a non-default set of updates be applied to this cluster. The default channel will be contain stable updates that are appropriate for production clusters.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_id": schema.StringAttribute{
						Description:         "clusterID uniquely identifies this cluster. This is expected to be an RFC4122 UUID value (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx in hexadecimal values). This is a required field.",
						MarkdownDescription: "clusterID uniquely identifies this cluster. This is expected to be an RFC4122 UUID value (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx in hexadecimal values). This is a required field.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"desired_update": schema.SingleNestedAttribute{
						Description:         "desiredUpdate is an optional field that indicates the desired value of the cluster version. Setting this value will trigger an upgrade (if the current version does not match the desired version). The set of recommended update values is listed as part of available updates in status, and setting values outside that range may cause the upgrade to fail.  Some of the fields are inter-related with restrictions and meanings described here. 1. image is specified, version is specified, architecture is specified. API validation error. 2. image is specified, version is specified, architecture is not specified. You should not do this. version is silently ignored and image is used. 3. image is specified, version is not specified, architecture is specified. API validation error. 4. image is specified, version is not specified, architecture is not specified. image is used. 5. image is not specified, version is specified, architecture is specified. version and desired architecture are used to select an image. 6. image is not specified, version is specified, architecture is not specified. version and current architecture are used to select an image. 7. image is not specified, version is not specified, architecture is specified. API validation error. 8. image is not specified, version is not specified, architecture is not specified. API validation error.  If an upgrade fails the operator will halt and report status about the failing component. Setting the desired update value back to the previous version will cause a rollback to be attempted. Not all rollbacks will succeed.",
						MarkdownDescription: "desiredUpdate is an optional field that indicates the desired value of the cluster version. Setting this value will trigger an upgrade (if the current version does not match the desired version). The set of recommended update values is listed as part of available updates in status, and setting values outside that range may cause the upgrade to fail.  Some of the fields are inter-related with restrictions and meanings described here. 1. image is specified, version is specified, architecture is specified. API validation error. 2. image is specified, version is specified, architecture is not specified. You should not do this. version is silently ignored and image is used. 3. image is specified, version is not specified, architecture is specified. API validation error. 4. image is specified, version is not specified, architecture is not specified. image is used. 5. image is not specified, version is specified, architecture is specified. version and desired architecture are used to select an image. 6. image is not specified, version is specified, architecture is not specified. version and current architecture are used to select an image. 7. image is not specified, version is not specified, architecture is specified. API validation error. 8. image is not specified, version is not specified, architecture is not specified. API validation error.  If an upgrade fails the operator will halt and report status about the failing component. Setting the desired update value back to the previous version will cause a rollback to be attempted. Not all rollbacks will succeed.",
						Attributes: map[string]schema.Attribute{
							"architecture": schema.StringAttribute{
								Description:         "architecture is an optional field that indicates the desired value of the cluster architecture. In this context cluster architecture means either a single architecture or a multi architecture. architecture can only be set to Multi thereby only allowing updates from single to multi architecture. If architecture is set, image cannot be set and version must be set. Valid values are 'Multi' and empty.",
								MarkdownDescription: "architecture is an optional field that indicates the desired value of the cluster architecture. In this context cluster architecture means either a single architecture or a multi architecture. architecture can only be set to Multi thereby only allowing updates from single to multi architecture. If architecture is set, image cannot be set and version must be set. Valid values are 'Multi' and empty.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Multi", ""),
								},
							},

							"force": schema.BoolAttribute{
								Description:         "force allows an administrator to update to an image that has failed verification or upgradeable checks. This option should only be used when the authenticity of the provided image has been verified out of band because the provided image will run with full administrative access to the cluster. Do not use this flag with images that comes from unknown or potentially malicious sources.",
								MarkdownDescription: "force allows an administrator to update to an image that has failed verification or upgradeable checks. This option should only be used when the authenticity of the provided image has been verified out of band because the provided image will run with full administrative access to the cluster. Do not use this flag with images that comes from unknown or potentially malicious sources.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.StringAttribute{
								Description:         "image is a container image location that contains the update. image should be used when the desired version does not exist in availableUpdates or history. When image is set, version is ignored. When image is set, version should be empty. When image is set, architecture cannot be specified.",
								MarkdownDescription: "image is a container image location that contains the update. image should be used when the desired version does not exist in availableUpdates or history. When image is set, version is ignored. When image is set, version should be empty. When image is set, architecture cannot be specified.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"version": schema.StringAttribute{
								Description:         "version is a semantic version identifying the update version. version is ignored if image is specified and required if architecture is specified.",
								MarkdownDescription: "version is a semantic version identifying the update version. version is ignored if image is specified and required if architecture is specified.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"overrides": schema.ListNestedAttribute{
						Description:         "overrides is list of overides for components that are managed by cluster version operator. Marking a component unmanaged will prevent the operator from creating or updating the object.",
						MarkdownDescription: "overrides is list of overides for components that are managed by cluster version operator. Marking a component unmanaged will prevent the operator from creating or updating the object.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group": schema.StringAttribute{
									Description:         "group identifies the API group that the kind is in.",
									MarkdownDescription: "group identifies the API group that the kind is in.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
									Description:         "kind indentifies which object to override.",
									MarkdownDescription: "kind indentifies which object to override.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "name is the component's name.",
									MarkdownDescription: "name is the component's name.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "namespace is the component's namespace. If the resource is cluster scoped, the namespace should be empty.",
									MarkdownDescription: "namespace is the component's namespace. If the resource is cluster scoped, the namespace should be empty.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"unmanaged": schema.BoolAttribute{
									Description:         "unmanaged controls if cluster version operator should stop managing the resources in this cluster. Default: false",
									MarkdownDescription: "unmanaged controls if cluster version operator should stop managing the resources in this cluster. Default: false",
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

					"upstream": schema.StringAttribute{
						Description:         "upstream may be used to specify the preferred update server. By default it will use the appropriate update server for the cluster and region.",
						MarkdownDescription: "upstream may be used to specify the preferred update server. By default it will use the appropriate update server for the cluster and region.",
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
	}
}

func (r *ConfigOpenshiftIoClusterVersionV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_config_openshift_io_cluster_version_v1_manifest")

	var model ConfigOpenshiftIoClusterVersionV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("config.openshift.io/v1")
	model.Kind = pointer.String("ClusterVersion")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
