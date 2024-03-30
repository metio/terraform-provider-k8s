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
	_ datasource.DataSource = &CoreKubeadmiralIoClusterOverridePolicyV1Alpha1Manifest{}
)

func NewCoreKubeadmiralIoClusterOverridePolicyV1Alpha1Manifest() datasource.DataSource {
	return &CoreKubeadmiralIoClusterOverridePolicyV1Alpha1Manifest{}
}

type CoreKubeadmiralIoClusterOverridePolicyV1Alpha1Manifest struct{}

type CoreKubeadmiralIoClusterOverridePolicyV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		OverrideRules *[]struct {
			Overriders *struct {
				Annotations *[]struct {
					Operator *string            `tfsdk:"operator" json:"operator,omitempty"`
					Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"annotations" json:"annotations,omitempty"`
				Args *[]struct {
					ContainerName *string   `tfsdk:"container_name" json:"containerName,omitempty"`
					Operator      *string   `tfsdk:"operator" json:"operator,omitempty"`
					Value         *[]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"args" json:"args,omitempty"`
				Command *[]struct {
					ContainerName *string   `tfsdk:"container_name" json:"containerName,omitempty"`
					Operator      *string   `tfsdk:"operator" json:"operator,omitempty"`
					Value         *[]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"command" json:"command,omitempty"`
				Image *[]struct {
					ContainerNames *[]string `tfsdk:"container_names" json:"containerNames,omitempty"`
					ImagePath      *string   `tfsdk:"image_path" json:"imagePath,omitempty"`
					Operations     *[]struct {
						ImageComponent *string `tfsdk:"image_component" json:"imageComponent,omitempty"`
						Operator       *string `tfsdk:"operator" json:"operator,omitempty"`
						Value          *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"operations" json:"operations,omitempty"`
				} `tfsdk:"image" json:"image,omitempty"`
				Jsonpatch *[]struct {
					Operator *string            `tfsdk:"operator" json:"operator,omitempty"`
					Path     *string            `tfsdk:"path" json:"path,omitempty"`
					Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"jsonpatch" json:"jsonpatch,omitempty"`
				Labels *[]struct {
					Operator *string            `tfsdk:"operator" json:"operator,omitempty"`
					Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"overriders" json:"overriders,omitempty"`
			TargetClusters *struct {
				ClusterAffinity *[]struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchFields *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_fields" json:"matchFields,omitempty"`
				} `tfsdk:"cluster_affinity" json:"clusterAffinity,omitempty"`
				ClusterSelector *map[string]string `tfsdk:"cluster_selector" json:"clusterSelector,omitempty"`
				Clusters        *[]string          `tfsdk:"clusters" json:"clusters,omitempty"`
			} `tfsdk:"target_clusters" json:"targetClusters,omitempty"`
		} `tfsdk:"override_rules" json:"overrideRules,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CoreKubeadmiralIoClusterOverridePolicyV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_core_kubeadmiral_io_cluster_override_policy_v1alpha1_manifest"
}

func (r *CoreKubeadmiralIoClusterOverridePolicyV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterOverridePolicy describes the override rules for a resource.",
		MarkdownDescription: "ClusterOverridePolicy describes the override rules for a resource.",
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
					"override_rules": schema.ListNestedAttribute{
						Description:         "OverrideRules specify the override rules. Each rule specifies the overriders and the clusters these overriders should be applied to.",
						MarkdownDescription: "OverrideRules specify the override rules. Each rule specifies the overriders and the clusters these overriders should be applied to.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"overriders": schema.SingleNestedAttribute{
									Description:         "Overriders specify the overriders to be applied in the target clusters.",
									MarkdownDescription: "Overriders specify the overriders to be applied in the target clusters.",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.ListNestedAttribute{
											Description:         "Annotation specifies overriders that apply to the resource annotations.",
											MarkdownDescription: "Annotation specifies overriders that apply to the resource annotations.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"operator": schema.StringAttribute{
														Description:         "Operator specifies the operation. If omitted, defaults to 'overwrite'.",
														MarkdownDescription: "Operator specifies the operation. If omitted, defaults to 'overwrite'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("addIfAbsent", "overwrite", "delete"),
														},
													},

													"value": schema.MapAttribute{
														Description:         "Value is the value(s) that will be applied to annotations/labels of resource. If Operator is 'addIfAbsent', items in Value (empty is not allowed) will be added in annotations/labels. - For 'addIfAbsent' Operator, the keys in Value cannot conflict with annotations/labels. If Operator is 'overwrite', items in Value which match in annotations/labels will be replaced. If Operator is 'delete', items in Value which match in annotations/labels will be deleted.",
														MarkdownDescription: "Value is the value(s) that will be applied to annotations/labels of resource. If Operator is 'addIfAbsent', items in Value (empty is not allowed) will be added in annotations/labels. - For 'addIfAbsent' Operator, the keys in Value cannot conflict with annotations/labels. If Operator is 'overwrite', items in Value which match in annotations/labels will be replaced. If Operator is 'delete', items in Value which match in annotations/labels will be deleted.",
														ElementType:         types.StringType,
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

										"args": schema.ListNestedAttribute{
											Description:         "Args specifies overriders that apply to the container arguments.",
											MarkdownDescription: "Args specifies overriders that apply to the container arguments.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"container_name": schema.StringAttribute{
														Description:         "ContainerName targets the specified container or init container in the pod template.",
														MarkdownDescription: "ContainerName targets the specified container or init container in the pod template.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"operator": schema.StringAttribute{
														Description:         "Operator specifies the operation. If omitted, defaults to 'overwrite'.",
														MarkdownDescription: "Operator specifies the operation. If omitted, defaults to 'overwrite'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("append", "overwrite", "delete"),
														},
													},

													"value": schema.ListAttribute{
														Description:         "Value is the value(s) that will be applied to command/args of ContainerName. If Operator is 'append', items in Value (empty is not allowed) will be appended to command/args. If Operator is 'overwrite', current command/args of ContainerName will be completely replaced by Value. If Operator is 'delete', items in Value that match in command/args will be deleted.",
														MarkdownDescription: "Value is the value(s) that will be applied to command/args of ContainerName. If Operator is 'append', items in Value (empty is not allowed) will be appended to command/args. If Operator is 'overwrite', current command/args of ContainerName will be completely replaced by Value. If Operator is 'delete', items in Value that match in command/args will be deleted.",
														ElementType:         types.StringType,
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

										"command": schema.ListNestedAttribute{
											Description:         "Command specifies overriders that apply to the container commands.",
											MarkdownDescription: "Command specifies overriders that apply to the container commands.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"container_name": schema.StringAttribute{
														Description:         "ContainerName targets the specified container or init container in the pod template.",
														MarkdownDescription: "ContainerName targets the specified container or init container in the pod template.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"operator": schema.StringAttribute{
														Description:         "Operator specifies the operation. If omitted, defaults to 'overwrite'.",
														MarkdownDescription: "Operator specifies the operation. If omitted, defaults to 'overwrite'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("append", "overwrite", "delete"),
														},
													},

													"value": schema.ListAttribute{
														Description:         "Value is the value(s) that will be applied to command/args of ContainerName. If Operator is 'append', items in Value (empty is not allowed) will be appended to command/args. If Operator is 'overwrite', current command/args of ContainerName will be completely replaced by Value. If Operator is 'delete', items in Value that match in command/args will be deleted.",
														MarkdownDescription: "Value is the value(s) that will be applied to command/args of ContainerName. If Operator is 'append', items in Value (empty is not allowed) will be appended to command/args. If Operator is 'overwrite', current command/args of ContainerName will be completely replaced by Value. If Operator is 'delete', items in Value that match in command/args will be deleted.",
														ElementType:         types.StringType,
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

										"image": schema.ListNestedAttribute{
											Description:         "Image specifies the overriders that apply to the image.",
											MarkdownDescription: "Image specifies the overriders that apply to the image.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"container_names": schema.ListAttribute{
														Description:         "ContainerNames are ignored when ImagePath is set. If empty, the image override rule applies to all containers. Otherwise, this override targets the specified container(s) or init container(s) in the pod template.",
														MarkdownDescription: "ContainerNames are ignored when ImagePath is set. If empty, the image override rule applies to all containers. Otherwise, this override targets the specified container(s) or init container(s) in the pod template.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"image_path": schema.StringAttribute{
														Description:         "ImagePath indicates the image path to target. For Example: /spec/template/spec/containers/0/image  If empty, the system will automatically resolve the image path if the resource type is Pod, CronJob, Deployment, StatefulSet, DaemonSet or Job.",
														MarkdownDescription: "ImagePath indicates the image path to target. For Example: /spec/template/spec/containers/0/image  If empty, the system will automatically resolve the image path if the resource type is Pod, CronJob, Deployment, StatefulSet, DaemonSet or Job.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"operations": schema.ListNestedAttribute{
														Description:         "Operations are the specific operations to be performed on ContainerNames or ImagePath.",
														MarkdownDescription: "Operations are the specific operations to be performed on ContainerNames or ImagePath.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"image_component": schema.StringAttribute{
																	Description:         "ImageComponent is the part of the image to override.",
																	MarkdownDescription: "ImageComponent is the part of the image to override.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("Registry", "Repository", "Tag", "Digest"),
																	},
																},

																"operator": schema.StringAttribute{
																	Description:         "Operator specifies the operation. If omitted, defaults to 'overwrite'.",
																	MarkdownDescription: "Operator specifies the operation. If omitted, defaults to 'overwrite'.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("addIfAbsent", "overwrite", "delete"),
																	},
																},

																"value": schema.StringAttribute{
																	Description:         "Value is the value required by the operation. For 'addIfAbsent' Operator, the old value of ImageComponent should be empty, and the Value shouldn't be empty.",
																	MarkdownDescription: "Value is the value required by the operation. For 'addIfAbsent' Operator, the old value of ImageComponent should be empty, and the Value shouldn't be empty.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"jsonpatch": schema.ListNestedAttribute{
											Description:         "JsonPatch specifies overriders in a syntax similar to RFC6902 JSON Patch.",
											MarkdownDescription: "JsonPatch specifies overriders in a syntax similar to RFC6902 JSON Patch.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"operator": schema.StringAttribute{
														Description:         "Operator specifies the operation. If omitted, defaults to 'replace'.",
														MarkdownDescription: "Operator specifies the operation. If omitted, defaults to 'replace'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"path": schema.StringAttribute{
														Description:         "Path is a JSON pointer (RFC 6901) specifying the location within the resource document where the operation is performed. Each key in the path should be prefixed with '/', while '~' and '/' should be escaped as '~0' and '~1' respectively. For example, to add a label 'kubeadmiral.io/label', the path should be '/metadata/labels/kubeadmiral.io~1label'.",
														MarkdownDescription: "Path is a JSON pointer (RFC 6901) specifying the location within the resource document where the operation is performed. Each key in the path should be prefixed with '/', while '~' and '/' should be escaped as '~0' and '~1' respectively. For example, to add a label 'kubeadmiral.io/label', the path should be '/metadata/labels/kubeadmiral.io~1label'.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"value": schema.MapAttribute{
														Description:         "Value is the value(s) required by the operation.",
														MarkdownDescription: "Value is the value(s) required by the operation.",
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

										"labels": schema.ListNestedAttribute{
											Description:         "Label specifies overriders that apply to the resource labels.",
											MarkdownDescription: "Label specifies overriders that apply to the resource labels.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"operator": schema.StringAttribute{
														Description:         "Operator specifies the operation. If omitted, defaults to 'overwrite'.",
														MarkdownDescription: "Operator specifies the operation. If omitted, defaults to 'overwrite'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("addIfAbsent", "overwrite", "delete"),
														},
													},

													"value": schema.MapAttribute{
														Description:         "Value is the value(s) that will be applied to annotations/labels of resource. If Operator is 'addIfAbsent', items in Value (empty is not allowed) will be added in annotations/labels. - For 'addIfAbsent' Operator, the keys in Value cannot conflict with annotations/labels. If Operator is 'overwrite', items in Value which match in annotations/labels will be replaced. If Operator is 'delete', items in Value which match in annotations/labels will be deleted.",
														MarkdownDescription: "Value is the value(s) that will be applied to annotations/labels of resource. If Operator is 'addIfAbsent', items in Value (empty is not allowed) will be added in annotations/labels. - For 'addIfAbsent' Operator, the keys in Value cannot conflict with annotations/labels. If Operator is 'overwrite', items in Value which match in annotations/labels will be replaced. If Operator is 'delete', items in Value which match in annotations/labels will be deleted.",
														ElementType:         types.StringType,
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"target_clusters": schema.SingleNestedAttribute{
									Description:         "TargetClusters selects the clusters in which the overriders in this rule should be applied. If multiple types of selectors are specified, the overall result is the intersection of all of them.",
									MarkdownDescription: "TargetClusters selects the clusters in which the overriders in this rule should be applied. If multiple types of selectors are specified, the overall result is the intersection of all of them.",
									Attributes: map[string]schema.Attribute{
										"cluster_affinity": schema.ListNestedAttribute{
											Description:         "ClusterAffinity selects FederatedClusters by matching their labels and fields against expressions. If multiple terms are specified, their results are ORed.",
											MarkdownDescription: "ClusterAffinity selects FederatedClusters by matching their labels and fields against expressions. If multiple terms are specified, their results are ORed.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"match_expressions": schema.ListNestedAttribute{
														Description:         "A list of cluster selector requirements by cluster labels.",
														MarkdownDescription: "A list of cluster selector requirements by cluster labels.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"operator": schema.StringAttribute{
																	Description:         "ClusterSelectorOperator is the set of operators that can be used in a cluster selector requirement.",
																	MarkdownDescription: "ClusterSelectorOperator is the set of operators that can be used in a cluster selector requirement.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist", "Gt", "Lt"),
																	},
																},

																"values": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
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

													"match_fields": schema.ListNestedAttribute{
														Description:         "A list of cluster selector requirements by cluster fields.",
														MarkdownDescription: "A list of cluster selector requirements by cluster fields.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"operator": schema.StringAttribute{
																	Description:         "ClusterSelectorOperator is the set of operators that can be used in a cluster selector requirement.",
																	MarkdownDescription: "ClusterSelectorOperator is the set of operators that can be used in a cluster selector requirement.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist", "Gt", "Lt"),
																	},
																},

																"values": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
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
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"cluster_selector": schema.MapAttribute{
											Description:         "ClusterSelector selects FederatedClusters by their labels. Empty labels selects all FederatedClusters.",
											MarkdownDescription: "ClusterSelector selects FederatedClusters by their labels. Empty labels selects all FederatedClusters.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"clusters": schema.ListAttribute{
											Description:         "Clusters selects FederatedClusters by their names. Empty Clusters selects all FederatedClusters.",
											MarkdownDescription: "Clusters selects FederatedClusters by their names. Empty Clusters selects all FederatedClusters.",
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

func (r *CoreKubeadmiralIoClusterOverridePolicyV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_core_kubeadmiral_io_cluster_override_policy_v1alpha1_manifest")

	var model CoreKubeadmiralIoClusterOverridePolicyV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("core.kubeadmiral.io/v1alpha1")
	model.Kind = pointer.String("ClusterOverridePolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
