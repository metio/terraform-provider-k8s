/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infinispan_org_v2alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &InfinispanOrgRestoreV2Alpha1Manifest{}
)

func NewInfinispanOrgRestoreV2Alpha1Manifest() datasource.DataSource {
	return &InfinispanOrgRestoreV2Alpha1Manifest{}
}

type InfinispanOrgRestoreV2Alpha1Manifest struct{}

type InfinispanOrgRestoreV2Alpha1ManifestData struct {
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
		Backup    *string `tfsdk:"backup" json:"backup,omitempty"`
		Cluster   *string `tfsdk:"cluster" json:"cluster,omitempty"`
		Container *struct {
			CliExtraJvmOpts    *string `tfsdk:"cli_extra_jvm_opts" json:"cliExtraJvmOpts,omitempty"`
			Cpu                *string `tfsdk:"cpu" json:"cpu,omitempty"`
			ExtraJvmOpts       *string `tfsdk:"extra_jvm_opts" json:"extraJvmOpts,omitempty"`
			Memory             *string `tfsdk:"memory" json:"memory,omitempty"`
			RouterExtraJvmOpts *string `tfsdk:"router_extra_jvm_opts" json:"routerExtraJvmOpts,omitempty"`
		} `tfsdk:"container" json:"container,omitempty"`
		Resources *struct {
			CacheConfigs *[]string `tfsdk:"cache_configs" json:"cacheConfigs,omitempty"`
			Caches       *[]string `tfsdk:"caches" json:"caches,omitempty"`
			Counters     *[]string `tfsdk:"counters" json:"counters,omitempty"`
			ProtoSchemas *[]string `tfsdk:"proto_schemas" json:"protoSchemas,omitempty"`
			Scripts      *[]string `tfsdk:"scripts" json:"scripts,omitempty"`
			Tasks        *[]string `tfsdk:"tasks" json:"tasks,omitempty"`
			Templates    *[]string `tfsdk:"templates" json:"templates,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfinispanOrgRestoreV2Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infinispan_org_restore_v2alpha1_manifest"
}

func (r *InfinispanOrgRestoreV2Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Restore is the Schema for the restores API",
		MarkdownDescription: "Restore is the Schema for the restores API",
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
				Description:         "BackupSpec defines the desired state of Backup",
				MarkdownDescription: "BackupSpec defines the desired state of Backup",
				Attributes: map[string]schema.Attribute{
					"backup": schema.StringAttribute{
						Description:         "The Infinispan Backup to restore",
						MarkdownDescription: "The Infinispan Backup to restore",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"cluster": schema.StringAttribute{
						Description:         "Infinispan cluster name",
						MarkdownDescription: "Infinispan cluster name",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"container": schema.SingleNestedAttribute{
						Description:         "InfinispanContainerSpec specify resource requirements per container",
						MarkdownDescription: "InfinispanContainerSpec specify resource requirements per container",
						Attributes: map[string]schema.Attribute{
							"cli_extra_jvm_opts": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cpu": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extra_jvm_opts": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"memory": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"router_extra_jvm_opts": schema.StringAttribute{
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

					"resources": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"cache_configs": schema.ListAttribute{
								Description:         "Deprecated and to be removed on subsequent release. Use .Templates instead.",
								MarkdownDescription: "Deprecated and to be removed on subsequent release. Use .Templates instead.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"caches": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"counters": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proto_schemas": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scripts": schema.ListAttribute{
								Description:         "Deprecated and to be removed on subsequent release. Use .Tasks instead.",
								MarkdownDescription: "Deprecated and to be removed on subsequent release. Use .Tasks instead.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tasks": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"templates": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *InfinispanOrgRestoreV2Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infinispan_org_restore_v2alpha1_manifest")

	var model InfinispanOrgRestoreV2Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("infinispan.org/v2alpha1")
	model.Kind = pointer.String("Restore")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
