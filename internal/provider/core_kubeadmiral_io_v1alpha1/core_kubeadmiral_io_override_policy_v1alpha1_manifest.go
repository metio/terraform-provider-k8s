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
	_ datasource.DataSource = &CoreKubeadmiralIoOverridePolicyV1Alpha1Manifest{}
)

func NewCoreKubeadmiralIoOverridePolicyV1Alpha1Manifest() datasource.DataSource {
	return &CoreKubeadmiralIoOverridePolicyV1Alpha1Manifest{}
}

type CoreKubeadmiralIoOverridePolicyV1Alpha1Manifest struct{}

type CoreKubeadmiralIoOverridePolicyV1Alpha1ManifestData struct {
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
				Envs *[]struct {
					ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
					Operator      *string `tfsdk:"operator" json:"operator,omitempty"`
					Value         *[]struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Value     *string `tfsdk:"value" json:"value,omitempty"`
						ValueFrom *struct {
							ConfigMapKeyRef *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
							FieldRef *struct {
								ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
								FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
							} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
							ResourceFieldRef *struct {
								ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
								Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
								Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
							} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
							SecretKeyRef *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						} `tfsdk:"value_from" json:"valueFrom,omitempty"`
					} `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"envs" json:"envs,omitempty"`
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

func (r *CoreKubeadmiralIoOverridePolicyV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_core_kubeadmiral_io_override_policy_v1alpha1_manifest"
}

func (r *CoreKubeadmiralIoOverridePolicyV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "OverridePolicy describes the override rules for a resource.",
		MarkdownDescription: "OverridePolicy describes the override rules for a resource.",
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

										"envs": schema.ListNestedAttribute{
											Description:         "Envs specifies overriders that apply to the container envs.",
											MarkdownDescription: "Envs specifies overriders that apply to the container envs.",
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
															stringvalidator.OneOf("addIfAbsent", "overwrite", "delete"),
														},
													},

													"value": schema.ListNestedAttribute{
														Description:         "List of environment variables to set in the container.",
														MarkdownDescription: "List of environment variables to set in the container.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
																	MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
																	MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"value_from": schema.SingleNestedAttribute{
																	Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
																	MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",
																	Attributes: map[string]schema.Attribute{
																		"config_map_key_ref": schema.SingleNestedAttribute{
																			Description:         "Selects a key of a ConfigMap.",
																			MarkdownDescription: "Selects a key of a ConfigMap.",
																			Attributes: map[string]schema.Attribute{
																				"key": schema.StringAttribute{
																					Description:         "The key to select.",
																					MarkdownDescription: "The key to select.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"name": schema.StringAttribute{
																					Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"optional": schema.BoolAttribute{
																					Description:         "Specify whether the ConfigMap or its key must be defined",
																					MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"field_ref": schema.SingleNestedAttribute{
																			Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																			MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																			Attributes: map[string]schema.Attribute{
																				"api_version": schema.StringAttribute{
																					Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																					MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"field_path": schema.StringAttribute{
																					Description:         "Path of the field to select in the specified API version.",
																					MarkdownDescription: "Path of the field to select in the specified API version.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"resource_field_ref": schema.SingleNestedAttribute{
																			Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																			MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																			Attributes: map[string]schema.Attribute{
																				"container_name": schema.StringAttribute{
																					Description:         "Container name: required for volumes, optional for env vars",
																					MarkdownDescription: "Container name: required for volumes, optional for env vars",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"divisor": schema.StringAttribute{
																					Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																					MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"resource": schema.StringAttribute{
																					Description:         "Required: resource to select",
																					MarkdownDescription: "Required: resource to select",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"secret_key_ref": schema.SingleNestedAttribute{
																			Description:         "Selects a key of a secret in the pod's namespace",
																			MarkdownDescription: "Selects a key of a secret in the pod's namespace",
																			Attributes: map[string]schema.Attribute{
																				"key": schema.StringAttribute{
																					Description:         "The key of the secret to select from.  Must be a valid secret key.",
																					MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"name": schema.StringAttribute{
																					Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"optional": schema.BoolAttribute{
																					Description:         "Specify whether the Secret or its key must be defined",
																					MarkdownDescription: "Specify whether the Secret or its key must be defined",
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

func (r *CoreKubeadmiralIoOverridePolicyV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_core_kubeadmiral_io_override_policy_v1alpha1_manifest")

	var model CoreKubeadmiralIoOverridePolicyV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("core.kubeadmiral.io/v1alpha1")
	model.Kind = pointer.String("OverridePolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
