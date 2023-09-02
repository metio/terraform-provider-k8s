/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package policy_karmada_io_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &PolicyKarmadaIoOverridePolicyV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &PolicyKarmadaIoOverridePolicyV1Alpha1DataSource{}
)

func NewPolicyKarmadaIoOverridePolicyV1Alpha1DataSource() datasource.DataSource {
	return &PolicyKarmadaIoOverridePolicyV1Alpha1DataSource{}
}

type PolicyKarmadaIoOverridePolicyV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type PolicyKarmadaIoOverridePolicyV1Alpha1DataSourceData struct {
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
		OverrideRules *[]struct {
			Overriders *struct {
				AnnotationsOverrider *[]struct {
					Operator *string            `tfsdk:"operator" json:"operator,omitempty"`
					Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"annotations_overrider" json:"annotationsOverrider,omitempty"`
				ArgsOverrider *[]struct {
					ContainerName *string   `tfsdk:"container_name" json:"containerName,omitempty"`
					Operator      *string   `tfsdk:"operator" json:"operator,omitempty"`
					Value         *[]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"args_overrider" json:"argsOverrider,omitempty"`
				CommandOverrider *[]struct {
					ContainerName *string   `tfsdk:"container_name" json:"containerName,omitempty"`
					Operator      *string   `tfsdk:"operator" json:"operator,omitempty"`
					Value         *[]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"command_overrider" json:"commandOverrider,omitempty"`
				ImageOverrider *[]struct {
					Component *string `tfsdk:"component" json:"component,omitempty"`
					Operator  *string `tfsdk:"operator" json:"operator,omitempty"`
					Predicate *struct {
						Path *string `tfsdk:"path" json:"path,omitempty"`
					} `tfsdk:"predicate" json:"predicate,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"image_overrider" json:"imageOverrider,omitempty"`
				LabelsOverrider *[]struct {
					Operator *string            `tfsdk:"operator" json:"operator,omitempty"`
					Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"labels_overrider" json:"labelsOverrider,omitempty"`
				Plaintext *[]struct {
					Operator *string            `tfsdk:"operator" json:"operator,omitempty"`
					Path     *string            `tfsdk:"path" json:"path,omitempty"`
					Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"plaintext" json:"plaintext,omitempty"`
			} `tfsdk:"overriders" json:"overriders,omitempty"`
			TargetCluster *struct {
				ClusterNames  *[]string `tfsdk:"cluster_names" json:"clusterNames,omitempty"`
				Exclude       *[]string `tfsdk:"exclude" json:"exclude,omitempty"`
				FieldSelector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				} `tfsdk:"field_selector" json:"fieldSelector,omitempty"`
				LabelSelector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
			} `tfsdk:"target_cluster" json:"targetCluster,omitempty"`
		} `tfsdk:"override_rules" json:"overrideRules,omitempty"`
		Overriders *struct {
			AnnotationsOverrider *[]struct {
				Operator *string            `tfsdk:"operator" json:"operator,omitempty"`
				Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"annotations_overrider" json:"annotationsOverrider,omitempty"`
			ArgsOverrider *[]struct {
				ContainerName *string   `tfsdk:"container_name" json:"containerName,omitempty"`
				Operator      *string   `tfsdk:"operator" json:"operator,omitempty"`
				Value         *[]string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"args_overrider" json:"argsOverrider,omitempty"`
			CommandOverrider *[]struct {
				ContainerName *string   `tfsdk:"container_name" json:"containerName,omitempty"`
				Operator      *string   `tfsdk:"operator" json:"operator,omitempty"`
				Value         *[]string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"command_overrider" json:"commandOverrider,omitempty"`
			ImageOverrider *[]struct {
				Component *string `tfsdk:"component" json:"component,omitempty"`
				Operator  *string `tfsdk:"operator" json:"operator,omitempty"`
				Predicate *struct {
					Path *string `tfsdk:"path" json:"path,omitempty"`
				} `tfsdk:"predicate" json:"predicate,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"image_overrider" json:"imageOverrider,omitempty"`
			LabelsOverrider *[]struct {
				Operator *string            `tfsdk:"operator" json:"operator,omitempty"`
				Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"labels_overrider" json:"labelsOverrider,omitempty"`
			Plaintext *[]struct {
				Operator *string            `tfsdk:"operator" json:"operator,omitempty"`
				Path     *string            `tfsdk:"path" json:"path,omitempty"`
				Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"plaintext" json:"plaintext,omitempty"`
		} `tfsdk:"overriders" json:"overriders,omitempty"`
		ResourceSelectors *[]struct {
			ApiVersion    *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind          *string `tfsdk:"kind" json:"kind,omitempty"`
			LabelSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"resource_selectors" json:"resourceSelectors,omitempty"`
		TargetCluster *struct {
			ClusterNames  *[]string `tfsdk:"cluster_names" json:"clusterNames,omitempty"`
			Exclude       *[]string `tfsdk:"exclude" json:"exclude,omitempty"`
			FieldSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			} `tfsdk:"field_selector" json:"fieldSelector,omitempty"`
			LabelSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
		} `tfsdk:"target_cluster" json:"targetCluster,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *PolicyKarmadaIoOverridePolicyV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_policy_karmada_io_override_policy_v1alpha1"
}

func (r *PolicyKarmadaIoOverridePolicyV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "OverridePolicy represents the policy that overrides a group of resources to one or more clusters.",
		MarkdownDescription: "OverridePolicy represents the policy that overrides a group of resources to one or more clusters.",
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
				Description:         "Spec represents the desired behavior of OverridePolicy.",
				MarkdownDescription: "Spec represents the desired behavior of OverridePolicy.",
				Attributes: map[string]schema.Attribute{
					"override_rules": schema.ListNestedAttribute{
						Description:         "OverrideRules defines a collection of override rules on target clusters.",
						MarkdownDescription: "OverrideRules defines a collection of override rules on target clusters.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"overriders": schema.SingleNestedAttribute{
									Description:         "Overriders represents the override rules that would apply on resources",
									MarkdownDescription: "Overriders represents the override rules that would apply on resources",
									Attributes: map[string]schema.Attribute{
										"annotations_overrider": schema.ListNestedAttribute{
											Description:         "AnnotationsOverrider represents the rules dedicated to handling workload annotations",
											MarkdownDescription: "AnnotationsOverrider represents the rules dedicated to handling workload annotations",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"operator": schema.StringAttribute{
														Description:         "Operator represents the operator which will apply on the workload.",
														MarkdownDescription: "Operator represents the operator which will apply on the workload.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"value": schema.MapAttribute{
														Description:         "Value to be applied to annotations/labels of workload. Items in Value which will be appended after annotations/labels when Operator is 'add'. Items in Value which match in annotations/labels will be deleted when Operator is 'remove'. Items in Value which match in annotations/labels will be replaced when Operator is 'replace'.",
														MarkdownDescription: "Value to be applied to annotations/labels of workload. Items in Value which will be appended after annotations/labels when Operator is 'add'. Items in Value which match in annotations/labels will be deleted when Operator is 'remove'. Items in Value which match in annotations/labels will be replaced when Operator is 'replace'.",
														ElementType:         types.StringType,
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

										"args_overrider": schema.ListNestedAttribute{
											Description:         "ArgsOverrider represents the rules dedicated to handling container args",
											MarkdownDescription: "ArgsOverrider represents the rules dedicated to handling container args",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"container_name": schema.StringAttribute{
														Description:         "The name of container",
														MarkdownDescription: "The name of container",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"operator": schema.StringAttribute{
														Description:         "Operator represents the operator which will apply on the command/args.",
														MarkdownDescription: "Operator represents the operator which will apply on the command/args.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"value": schema.ListAttribute{
														Description:         "Value to be applied to command/args. Items in Value which will be appended after command/args when Operator is 'add'. Items in Value which match in command/args will be deleted when Operator is 'remove'. If Value is empty, then the command/args will remain the same.",
														MarkdownDescription: "Value to be applied to command/args. Items in Value which will be appended after command/args when Operator is 'add'. Items in Value which match in command/args will be deleted when Operator is 'remove'. If Value is empty, then the command/args will remain the same.",
														ElementType:         types.StringType,
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

										"command_overrider": schema.ListNestedAttribute{
											Description:         "CommandOverrider represents the rules dedicated to handling container command",
											MarkdownDescription: "CommandOverrider represents the rules dedicated to handling container command",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"container_name": schema.StringAttribute{
														Description:         "The name of container",
														MarkdownDescription: "The name of container",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"operator": schema.StringAttribute{
														Description:         "Operator represents the operator which will apply on the command/args.",
														MarkdownDescription: "Operator represents the operator which will apply on the command/args.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"value": schema.ListAttribute{
														Description:         "Value to be applied to command/args. Items in Value which will be appended after command/args when Operator is 'add'. Items in Value which match in command/args will be deleted when Operator is 'remove'. If Value is empty, then the command/args will remain the same.",
														MarkdownDescription: "Value to be applied to command/args. Items in Value which will be appended after command/args when Operator is 'add'. Items in Value which match in command/args will be deleted when Operator is 'remove'. If Value is empty, then the command/args will remain the same.",
														ElementType:         types.StringType,
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

										"image_overrider": schema.ListNestedAttribute{
											Description:         "ImageOverrider represents the rules dedicated to handling image overrides.",
											MarkdownDescription: "ImageOverrider represents the rules dedicated to handling image overrides.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"component": schema.StringAttribute{
														Description:         "Component is part of image name. Basically we presume an image can be made of '[registry/]repository[:tag]'. The registry could be: - registry.k8s.io - fictional.registry.example:10443 The repository could be: - kube-apiserver - fictional/nginx The tag cloud be: - latest - v1.19.1 - @sha256:dbcc1c35ac38df41fd2f5e4130b32ffdb93ebae8b3dbe638c23575912276fc9c",
														MarkdownDescription: "Component is part of image name. Basically we presume an image can be made of '[registry/]repository[:tag]'. The registry could be: - registry.k8s.io - fictional.registry.example:10443 The repository could be: - kube-apiserver - fictional/nginx The tag cloud be: - latest - v1.19.1 - @sha256:dbcc1c35ac38df41fd2f5e4130b32ffdb93ebae8b3dbe638c23575912276fc9c",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"operator": schema.StringAttribute{
														Description:         "Operator represents the operator which will apply on the image.",
														MarkdownDescription: "Operator represents the operator which will apply on the image.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"predicate": schema.SingleNestedAttribute{
														Description:         "Predicate filters images before applying the rule.  Defaults to nil, in that case, the system will automatically detect image fields if the resource type is Pod, ReplicaSet, Deployment, StatefulSet, DaemonSet or Job by following rule: - Pod: /spec/containers/<N>/image - ReplicaSet: /spec/template/spec/containers/<N>/image - Deployment: /spec/template/spec/containers/<N>/image - DaemonSet: /spec/template/spec/containers/<N>/image - StatefulSet: /spec/template/spec/containers/<N>/image - Job: /spec/template/spec/containers/<N>/image In addition, all images will be processed if the resource object has more than one container.  If not nil, only images matches the filters will be processed.",
														MarkdownDescription: "Predicate filters images before applying the rule.  Defaults to nil, in that case, the system will automatically detect image fields if the resource type is Pod, ReplicaSet, Deployment, StatefulSet, DaemonSet or Job by following rule: - Pod: /spec/containers/<N>/image - ReplicaSet: /spec/template/spec/containers/<N>/image - Deployment: /spec/template/spec/containers/<N>/image - DaemonSet: /spec/template/spec/containers/<N>/image - StatefulSet: /spec/template/spec/containers/<N>/image - Job: /spec/template/spec/containers/<N>/image In addition, all images will be processed if the resource object has more than one container.  If not nil, only images matches the filters will be processed.",
														Attributes: map[string]schema.Attribute{
															"path": schema.StringAttribute{
																Description:         "Path indicates the path of target field",
																MarkdownDescription: "Path indicates the path of target field",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"value": schema.StringAttribute{
														Description:         "Value to be applied to image. Must not be empty when operator is 'add' or 'replace'. Defaults to empty and ignored when operator is 'remove'.",
														MarkdownDescription: "Value to be applied to image. Must not be empty when operator is 'add' or 'replace'. Defaults to empty and ignored when operator is 'remove'.",
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

										"labels_overrider": schema.ListNestedAttribute{
											Description:         "LabelsOverrider represents the rules dedicated to handling workload labels",
											MarkdownDescription: "LabelsOverrider represents the rules dedicated to handling workload labels",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"operator": schema.StringAttribute{
														Description:         "Operator represents the operator which will apply on the workload.",
														MarkdownDescription: "Operator represents the operator which will apply on the workload.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"value": schema.MapAttribute{
														Description:         "Value to be applied to annotations/labels of workload. Items in Value which will be appended after annotations/labels when Operator is 'add'. Items in Value which match in annotations/labels will be deleted when Operator is 'remove'. Items in Value which match in annotations/labels will be replaced when Operator is 'replace'.",
														MarkdownDescription: "Value to be applied to annotations/labels of workload. Items in Value which will be appended after annotations/labels when Operator is 'add'. Items in Value which match in annotations/labels will be deleted when Operator is 'remove'. Items in Value which match in annotations/labels will be replaced when Operator is 'replace'.",
														ElementType:         types.StringType,
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

										"plaintext": schema.ListNestedAttribute{
											Description:         "Plaintext represents override rules defined with plaintext overriders.",
											MarkdownDescription: "Plaintext represents override rules defined with plaintext overriders.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"operator": schema.StringAttribute{
														Description:         "Operator indicates the operation on target field. Available operators are: add, replace and remove.",
														MarkdownDescription: "Operator indicates the operation on target field. Available operators are: add, replace and remove.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"path": schema.StringAttribute{
														Description:         "Path indicates the path of target field",
														MarkdownDescription: "Path indicates the path of target field",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"value": schema.MapAttribute{
														Description:         "Value to be applied to target field. Must be empty when operator is Remove.",
														MarkdownDescription: "Value to be applied to target field. Must be empty when operator is Remove.",
														ElementType:         types.StringType,
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
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"target_cluster": schema.SingleNestedAttribute{
									Description:         "TargetCluster defines restrictions on this override policy that only applies to resources propagated to the matching clusters. nil means matching all clusters.",
									MarkdownDescription: "TargetCluster defines restrictions on this override policy that only applies to resources propagated to the matching clusters. nil means matching all clusters.",
									Attributes: map[string]schema.Attribute{
										"cluster_names": schema.ListAttribute{
											Description:         "ClusterNames is the list of clusters to be selected.",
											MarkdownDescription: "ClusterNames is the list of clusters to be selected.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"exclude": schema.ListAttribute{
											Description:         "ExcludedClusters is the list of clusters to be ignored.",
											MarkdownDescription: "ExcludedClusters is the list of clusters to be ignored.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"field_selector": schema.SingleNestedAttribute{
											Description:         "FieldSelector is a filter to select member clusters by fields. The key(field) of the match expression should be 'provider', 'region', or 'zone', and the operator of the match expression should be 'In' or 'NotIn'. If non-nil and non-empty, only the clusters match this filter will be selected.",
											MarkdownDescription: "FieldSelector is a filter to select member clusters by fields. The key(field) of the match expression should be 'provider', 'region', or 'zone', and the operator of the match expression should be 'In' or 'NotIn'. If non-nil and non-empty, only the clusters match this filter will be selected.",
											Attributes: map[string]schema.Attribute{
												"match_expressions": schema.ListNestedAttribute{
													Description:         "A list of field selector requirements.",
													MarkdownDescription: "A list of field selector requirements.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The label key that the selector applies to.",
																MarkdownDescription: "The label key that the selector applies to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"operator": schema.StringAttribute{
																Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"values": schema.ListAttribute{
																Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																ElementType:         types.StringType,
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
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"label_selector": schema.SingleNestedAttribute{
											Description:         "LabelSelector is a filter to select member clusters by labels. If non-nil and non-empty, only the clusters match this filter will be selected.",
											MarkdownDescription: "LabelSelector is a filter to select member clusters by labels. If non-nil and non-empty, only the clusters match this filter will be selected.",
											Attributes: map[string]schema.Attribute{
												"match_expressions": schema.ListNestedAttribute{
													Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"operator": schema.StringAttribute{
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"values": schema.ListAttribute{
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																ElementType:         types.StringType,
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

												"match_labels": schema.MapAttribute{
													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													ElementType:         types.StringType,
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
									Required: false,
									Optional: false,
									Computed: true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"overriders": schema.SingleNestedAttribute{
						Description:         "Overriders represents the override rules that would apply on resources  Deprecated: This filed is deprecated in v1.0 and please use the OverrideRules instead.",
						MarkdownDescription: "Overriders represents the override rules that would apply on resources  Deprecated: This filed is deprecated in v1.0 and please use the OverrideRules instead.",
						Attributes: map[string]schema.Attribute{
							"annotations_overrider": schema.ListNestedAttribute{
								Description:         "AnnotationsOverrider represents the rules dedicated to handling workload annotations",
								MarkdownDescription: "AnnotationsOverrider represents the rules dedicated to handling workload annotations",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"operator": schema.StringAttribute{
											Description:         "Operator represents the operator which will apply on the workload.",
											MarkdownDescription: "Operator represents the operator which will apply on the workload.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"value": schema.MapAttribute{
											Description:         "Value to be applied to annotations/labels of workload. Items in Value which will be appended after annotations/labels when Operator is 'add'. Items in Value which match in annotations/labels will be deleted when Operator is 'remove'. Items in Value which match in annotations/labels will be replaced when Operator is 'replace'.",
											MarkdownDescription: "Value to be applied to annotations/labels of workload. Items in Value which will be appended after annotations/labels when Operator is 'add'. Items in Value which match in annotations/labels will be deleted when Operator is 'remove'. Items in Value which match in annotations/labels will be replaced when Operator is 'replace'.",
											ElementType:         types.StringType,
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

							"args_overrider": schema.ListNestedAttribute{
								Description:         "ArgsOverrider represents the rules dedicated to handling container args",
								MarkdownDescription: "ArgsOverrider represents the rules dedicated to handling container args",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"container_name": schema.StringAttribute{
											Description:         "The name of container",
											MarkdownDescription: "The name of container",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"operator": schema.StringAttribute{
											Description:         "Operator represents the operator which will apply on the command/args.",
											MarkdownDescription: "Operator represents the operator which will apply on the command/args.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"value": schema.ListAttribute{
											Description:         "Value to be applied to command/args. Items in Value which will be appended after command/args when Operator is 'add'. Items in Value which match in command/args will be deleted when Operator is 'remove'. If Value is empty, then the command/args will remain the same.",
											MarkdownDescription: "Value to be applied to command/args. Items in Value which will be appended after command/args when Operator is 'add'. Items in Value which match in command/args will be deleted when Operator is 'remove'. If Value is empty, then the command/args will remain the same.",
											ElementType:         types.StringType,
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

							"command_overrider": schema.ListNestedAttribute{
								Description:         "CommandOverrider represents the rules dedicated to handling container command",
								MarkdownDescription: "CommandOverrider represents the rules dedicated to handling container command",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"container_name": schema.StringAttribute{
											Description:         "The name of container",
											MarkdownDescription: "The name of container",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"operator": schema.StringAttribute{
											Description:         "Operator represents the operator which will apply on the command/args.",
											MarkdownDescription: "Operator represents the operator which will apply on the command/args.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"value": schema.ListAttribute{
											Description:         "Value to be applied to command/args. Items in Value which will be appended after command/args when Operator is 'add'. Items in Value which match in command/args will be deleted when Operator is 'remove'. If Value is empty, then the command/args will remain the same.",
											MarkdownDescription: "Value to be applied to command/args. Items in Value which will be appended after command/args when Operator is 'add'. Items in Value which match in command/args will be deleted when Operator is 'remove'. If Value is empty, then the command/args will remain the same.",
											ElementType:         types.StringType,
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

							"image_overrider": schema.ListNestedAttribute{
								Description:         "ImageOverrider represents the rules dedicated to handling image overrides.",
								MarkdownDescription: "ImageOverrider represents the rules dedicated to handling image overrides.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"component": schema.StringAttribute{
											Description:         "Component is part of image name. Basically we presume an image can be made of '[registry/]repository[:tag]'. The registry could be: - registry.k8s.io - fictional.registry.example:10443 The repository could be: - kube-apiserver - fictional/nginx The tag cloud be: - latest - v1.19.1 - @sha256:dbcc1c35ac38df41fd2f5e4130b32ffdb93ebae8b3dbe638c23575912276fc9c",
											MarkdownDescription: "Component is part of image name. Basically we presume an image can be made of '[registry/]repository[:tag]'. The registry could be: - registry.k8s.io - fictional.registry.example:10443 The repository could be: - kube-apiserver - fictional/nginx The tag cloud be: - latest - v1.19.1 - @sha256:dbcc1c35ac38df41fd2f5e4130b32ffdb93ebae8b3dbe638c23575912276fc9c",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"operator": schema.StringAttribute{
											Description:         "Operator represents the operator which will apply on the image.",
											MarkdownDescription: "Operator represents the operator which will apply on the image.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"predicate": schema.SingleNestedAttribute{
											Description:         "Predicate filters images before applying the rule.  Defaults to nil, in that case, the system will automatically detect image fields if the resource type is Pod, ReplicaSet, Deployment, StatefulSet, DaemonSet or Job by following rule: - Pod: /spec/containers/<N>/image - ReplicaSet: /spec/template/spec/containers/<N>/image - Deployment: /spec/template/spec/containers/<N>/image - DaemonSet: /spec/template/spec/containers/<N>/image - StatefulSet: /spec/template/spec/containers/<N>/image - Job: /spec/template/spec/containers/<N>/image In addition, all images will be processed if the resource object has more than one container.  If not nil, only images matches the filters will be processed.",
											MarkdownDescription: "Predicate filters images before applying the rule.  Defaults to nil, in that case, the system will automatically detect image fields if the resource type is Pod, ReplicaSet, Deployment, StatefulSet, DaemonSet or Job by following rule: - Pod: /spec/containers/<N>/image - ReplicaSet: /spec/template/spec/containers/<N>/image - Deployment: /spec/template/spec/containers/<N>/image - DaemonSet: /spec/template/spec/containers/<N>/image - StatefulSet: /spec/template/spec/containers/<N>/image - Job: /spec/template/spec/containers/<N>/image In addition, all images will be processed if the resource object has more than one container.  If not nil, only images matches the filters will be processed.",
											Attributes: map[string]schema.Attribute{
												"path": schema.StringAttribute{
													Description:         "Path indicates the path of target field",
													MarkdownDescription: "Path indicates the path of target field",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"value": schema.StringAttribute{
											Description:         "Value to be applied to image. Must not be empty when operator is 'add' or 'replace'. Defaults to empty and ignored when operator is 'remove'.",
											MarkdownDescription: "Value to be applied to image. Must not be empty when operator is 'add' or 'replace'. Defaults to empty and ignored when operator is 'remove'.",
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

							"labels_overrider": schema.ListNestedAttribute{
								Description:         "LabelsOverrider represents the rules dedicated to handling workload labels",
								MarkdownDescription: "LabelsOverrider represents the rules dedicated to handling workload labels",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"operator": schema.StringAttribute{
											Description:         "Operator represents the operator which will apply on the workload.",
											MarkdownDescription: "Operator represents the operator which will apply on the workload.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"value": schema.MapAttribute{
											Description:         "Value to be applied to annotations/labels of workload. Items in Value which will be appended after annotations/labels when Operator is 'add'. Items in Value which match in annotations/labels will be deleted when Operator is 'remove'. Items in Value which match in annotations/labels will be replaced when Operator is 'replace'.",
											MarkdownDescription: "Value to be applied to annotations/labels of workload. Items in Value which will be appended after annotations/labels when Operator is 'add'. Items in Value which match in annotations/labels will be deleted when Operator is 'remove'. Items in Value which match in annotations/labels will be replaced when Operator is 'replace'.",
											ElementType:         types.StringType,
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

							"plaintext": schema.ListNestedAttribute{
								Description:         "Plaintext represents override rules defined with plaintext overriders.",
								MarkdownDescription: "Plaintext represents override rules defined with plaintext overriders.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"operator": schema.StringAttribute{
											Description:         "Operator indicates the operation on target field. Available operators are: add, replace and remove.",
											MarkdownDescription: "Operator indicates the operation on target field. Available operators are: add, replace and remove.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"path": schema.StringAttribute{
											Description:         "Path indicates the path of target field",
											MarkdownDescription: "Path indicates the path of target field",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"value": schema.MapAttribute{
											Description:         "Value to be applied to target field. Must be empty when operator is Remove.",
											MarkdownDescription: "Value to be applied to target field. Must be empty when operator is Remove.",
											ElementType:         types.StringType,
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"resource_selectors": schema.ListNestedAttribute{
						Description:         "ResourceSelectors restricts resource types that this override policy applies to. nil means matching all resources.",
						MarkdownDescription: "ResourceSelectors restricts resource types that this override policy applies to. nil means matching all resources.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"api_version": schema.StringAttribute{
									Description:         "APIVersion represents the API version of the target resources.",
									MarkdownDescription: "APIVersion represents the API version of the target resources.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"kind": schema.StringAttribute{
									Description:         "Kind represents the Kind of the target resources.",
									MarkdownDescription: "Kind represents the Kind of the target resources.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"label_selector": schema.SingleNestedAttribute{
									Description:         "A label query over a set of resources. If name is not empty, labelSelector will be ignored.",
									MarkdownDescription: "A label query over a set of resources. If name is not empty, labelSelector will be ignored.",
									Attributes: map[string]schema.Attribute{
										"match_expressions": schema.ListNestedAttribute{
											Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
											MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"operator": schema.StringAttribute{
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"values": schema.ListAttribute{
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														ElementType:         types.StringType,
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

										"match_labels": schema.MapAttribute{
											Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the target resource. Default is empty, which means selecting all resources.",
									MarkdownDescription: "Name of the target resource. Default is empty, which means selecting all resources.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace of the target resource. Default is empty, which means inherit from the parent object scope.",
									MarkdownDescription: "Namespace of the target resource. Default is empty, which means inherit from the parent object scope.",
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

					"target_cluster": schema.SingleNestedAttribute{
						Description:         "TargetCluster defines restrictions on this override policy that only applies to resources propagated to the matching clusters. nil means matching all clusters.  Deprecated: This filed is deprecated in v1.0 and please use the OverrideRules instead.",
						MarkdownDescription: "TargetCluster defines restrictions on this override policy that only applies to resources propagated to the matching clusters. nil means matching all clusters.  Deprecated: This filed is deprecated in v1.0 and please use the OverrideRules instead.",
						Attributes: map[string]schema.Attribute{
							"cluster_names": schema.ListAttribute{
								Description:         "ClusterNames is the list of clusters to be selected.",
								MarkdownDescription: "ClusterNames is the list of clusters to be selected.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"exclude": schema.ListAttribute{
								Description:         "ExcludedClusters is the list of clusters to be ignored.",
								MarkdownDescription: "ExcludedClusters is the list of clusters to be ignored.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"field_selector": schema.SingleNestedAttribute{
								Description:         "FieldSelector is a filter to select member clusters by fields. The key(field) of the match expression should be 'provider', 'region', or 'zone', and the operator of the match expression should be 'In' or 'NotIn'. If non-nil and non-empty, only the clusters match this filter will be selected.",
								MarkdownDescription: "FieldSelector is a filter to select member clusters by fields. The key(field) of the match expression should be 'provider', 'region', or 'zone', and the operator of the match expression should be 'In' or 'NotIn'. If non-nil and non-empty, only the clusters match this filter will be selected.",
								Attributes: map[string]schema.Attribute{
									"match_expressions": schema.ListNestedAttribute{
										Description:         "A list of field selector requirements.",
										MarkdownDescription: "A list of field selector requirements.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The label key that the selector applies to.",
													MarkdownDescription: "The label key that the selector applies to.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"operator": schema.StringAttribute{
													Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
													MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"values": schema.ListAttribute{
													Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
													MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
													ElementType:         types.StringType,
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
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"label_selector": schema.SingleNestedAttribute{
								Description:         "LabelSelector is a filter to select member clusters by labels. If non-nil and non-empty, only the clusters match this filter will be selected.",
								MarkdownDescription: "LabelSelector is a filter to select member clusters by labels. If non-nil and non-empty, only the clusters match this filter will be selected.",
								Attributes: map[string]schema.Attribute{
									"match_expressions": schema.ListNestedAttribute{
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "key is the label key that the selector applies to.",
													MarkdownDescription: "key is the label key that the selector applies to.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"operator": schema.StringAttribute{
													Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"values": schema.ListAttribute{
													Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													ElementType:         types.StringType,
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

									"match_labels": schema.MapAttribute{
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										ElementType:         types.StringType,
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
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *PolicyKarmadaIoOverridePolicyV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *PolicyKarmadaIoOverridePolicyV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_policy_karmada_io_override_policy_v1alpha1")

	var data PolicyKarmadaIoOverridePolicyV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "policy.karmada.io", Version: "v1alpha1", Resource: "OverridePolicy"}).
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

	var readResponse PolicyKarmadaIoOverridePolicyV1Alpha1DataSourceData
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
	data.ApiVersion = pointer.String("policy.karmada.io/v1alpha1")
	data.Kind = pointer.String("OverridePolicy")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
