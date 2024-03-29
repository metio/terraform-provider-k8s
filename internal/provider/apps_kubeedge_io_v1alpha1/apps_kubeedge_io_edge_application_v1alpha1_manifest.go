/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_kubeedge_io_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &AppsKubeedgeIoEdgeApplicationV1Alpha1Manifest{}
)

func NewAppsKubeedgeIoEdgeApplicationV1Alpha1Manifest() datasource.DataSource {
	return &AppsKubeedgeIoEdgeApplicationV1Alpha1Manifest{}
}

type AppsKubeedgeIoEdgeApplicationV1Alpha1Manifest struct{}

type AppsKubeedgeIoEdgeApplicationV1Alpha1ManifestData struct {
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
		WorkloadScope *struct {
			TargetNodeGroups *[]struct {
				Name       *string `tfsdk:"name" json:"name,omitempty"`
				Overriders *struct {
					ArgsOverriders *[]struct {
						ContainerName *string   `tfsdk:"container_name" json:"containerName,omitempty"`
						Operator      *string   `tfsdk:"operator" json:"operator,omitempty"`
						Value         *[]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"args_overriders" json:"argsOverriders,omitempty"`
					CommandOverriders *[]struct {
						ContainerName *string   `tfsdk:"container_name" json:"containerName,omitempty"`
						Operator      *string   `tfsdk:"operator" json:"operator,omitempty"`
						Value         *[]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"command_overriders" json:"commandOverriders,omitempty"`
					EnvOverriders *[]struct {
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
					} `tfsdk:"env_overriders" json:"envOverriders,omitempty"`
					ImageOverriders *[]struct {
						Component *string `tfsdk:"component" json:"component,omitempty"`
						Operator  *string `tfsdk:"operator" json:"operator,omitempty"`
						Predicate *struct {
							Path *string `tfsdk:"path" json:"path,omitempty"`
						} `tfsdk:"predicate" json:"predicate,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"image_overriders" json:"imageOverriders,omitempty"`
					Replicas            *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
					ResourcesOverriders *[]struct {
						ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
						Value         *struct {
							Claims *[]struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"claims" json:"claims,omitempty"`
							Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
							Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
						} `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"resources_overriders" json:"resourcesOverriders,omitempty"`
				} `tfsdk:"overriders" json:"overriders,omitempty"`
			} `tfsdk:"target_node_groups" json:"targetNodeGroups,omitempty"`
		} `tfsdk:"workload_scope" json:"workloadScope,omitempty"`
		WorkloadTemplate *struct {
			Manifests *[]map[string]string `tfsdk:"manifests" json:"manifests,omitempty"`
		} `tfsdk:"workload_template" json:"workloadTemplate,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsKubeedgeIoEdgeApplicationV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_kubeedge_io_edge_application_v1alpha1_manifest"
}

func (r *AppsKubeedgeIoEdgeApplicationV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "EdgeApplication is the Schema for the edgeapplications API",
		MarkdownDescription: "EdgeApplication is the Schema for the edgeapplications API",
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
				Description:         "Spec represents the desired behavior of EdgeApplication.",
				MarkdownDescription: "Spec represents the desired behavior of EdgeApplication.",
				Attributes: map[string]schema.Attribute{
					"workload_scope": schema.SingleNestedAttribute{
						Description:         "WorkloadScope represents which node groups the workload will be deployed in.",
						MarkdownDescription: "WorkloadScope represents which node groups the workload will be deployed in.",
						Attributes: map[string]schema.Attribute{
							"target_node_groups": schema.ListNestedAttribute{
								Description:         "TargetNodeGroups represents the target node groups of workload to be deployed.",
								MarkdownDescription: "TargetNodeGroups represents the target node groups of workload to be deployed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name represents the name of target node group",
											MarkdownDescription: "Name represents the name of target node group",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"overriders": schema.SingleNestedAttribute{
											Description:         "Overriders represents the override rules that would apply on workload.",
											MarkdownDescription: "Overriders represents the override rules that would apply on workload.",
											Attributes: map[string]schema.Attribute{
												"args_overriders": schema.ListNestedAttribute{
													Description:         "ArgsOverriders represents the rules dedicated to handling container args",
													MarkdownDescription: "ArgsOverriders represents the rules dedicated to handling container args",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"container_name": schema.StringAttribute{
																Description:         "The name of container",
																MarkdownDescription: "The name of container",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "Operator represents the operator which will apply on the command/args.",
																MarkdownDescription: "Operator represents the operator which will apply on the command/args.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("add", "remove"),
																},
															},

															"value": schema.ListAttribute{
																Description:         "Value to be applied to command/args. Items in Value which will be appended after command/args when Operator is 'add'. Items in Value which match in command/args will be deleted when Operator is 'remove'. If Value is empty, then the command/args will remain the same.",
																MarkdownDescription: "Value to be applied to command/args. Items in Value which will be appended after command/args when Operator is 'add'. Items in Value which match in command/args will be deleted when Operator is 'remove'. If Value is empty, then the command/args will remain the same.",
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

												"command_overriders": schema.ListNestedAttribute{
													Description:         "CommandOverriders represents the rules dedicated to handling container command",
													MarkdownDescription: "CommandOverriders represents the rules dedicated to handling container command",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"container_name": schema.StringAttribute{
																Description:         "The name of container",
																MarkdownDescription: "The name of container",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "Operator represents the operator which will apply on the command/args.",
																MarkdownDescription: "Operator represents the operator which will apply on the command/args.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("add", "remove"),
																},
															},

															"value": schema.ListAttribute{
																Description:         "Value to be applied to command/args. Items in Value which will be appended after command/args when Operator is 'add'. Items in Value which match in command/args will be deleted when Operator is 'remove'. If Value is empty, then the command/args will remain the same.",
																MarkdownDescription: "Value to be applied to command/args. Items in Value which will be appended after command/args when Operator is 'add'. Items in Value which match in command/args will be deleted when Operator is 'remove'. If Value is empty, then the command/args will remain the same.",
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

												"env_overriders": schema.ListNestedAttribute{
													Description:         "EnvOverriders will override the env field of the container",
													MarkdownDescription: "EnvOverriders will override the env field of the container",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"container_name": schema.StringAttribute{
																Description:         "The name of container",
																MarkdownDescription: "The name of container",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "Operator represents the operator which will apply on the env.",
																MarkdownDescription: "Operator represents the operator which will apply on the env.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("add", "remove", "replace"),
																},
															},

															"value": schema.ListNestedAttribute{
																Description:         "Value to be applied to env. Must not be empty when operator is 'add' or 'replace'. When the operator is 'remove', the matched value in env will be deleted and only the name of the value will be matched. If Value is empty, then the env will remain the same.",
																MarkdownDescription: "Value to be applied to env. Must not be empty when operator is 'add' or 'replace'. When the operator is 'remove', the matched value in env will be deleted and only the name of the value will be matched. If Value is empty, then the env will remain the same.",
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

												"image_overriders": schema.ListNestedAttribute{
													Description:         "ImageOverriders represents the rules dedicated to handling image overrides.",
													MarkdownDescription: "ImageOverriders represents the rules dedicated to handling image overrides.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"component": schema.StringAttribute{
																Description:         "Component is part of image name. Basically we presume an image can be made of '[registry/]repository[:tag]'. The registry could be: - k8s.gcr.io - fictional.registry.example:10443 The repository could be: - kube-apiserver - fictional/nginx The tag cloud be: - latest - v1.19.1 - @sha256:dbcc1c35ac38df41fd2f5e4130b32ffdb93ebae8b3dbe638c23575912276fc9c",
																MarkdownDescription: "Component is part of image name. Basically we presume an image can be made of '[registry/]repository[:tag]'. The registry could be: - k8s.gcr.io - fictional.registry.example:10443 The repository could be: - kube-apiserver - fictional/nginx The tag cloud be: - latest - v1.19.1 - @sha256:dbcc1c35ac38df41fd2f5e4130b32ffdb93ebae8b3dbe638c23575912276fc9c",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("Registry", "Repository", "Tag"),
																},
															},

															"operator": schema.StringAttribute{
																Description:         "Operator represents the operator which will apply on the image.",
																MarkdownDescription: "Operator represents the operator which will apply on the image.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("add", "remove", "replace"),
																},
															},

															"predicate": schema.SingleNestedAttribute{
																Description:         "Predicate filters images before applying the rule.  Defaults to nil, in that case, the system will automatically detect image fields if the resource type is Pod, ReplicaSet, Deployment or StatefulSet by following rule:   - Pod: /spec/containers/<N>/image   - ReplicaSet: /spec/template/spec/containers/<N>/image   - Deployment: /spec/template/spec/containers/<N>/image   - StatefulSet: /spec/template/spec/containers/<N>/image In addition, all images will be processed if the resource object has more than one containers.  If not nil, only images matches the filters will be processed.",
																MarkdownDescription: "Predicate filters images before applying the rule.  Defaults to nil, in that case, the system will automatically detect image fields if the resource type is Pod, ReplicaSet, Deployment or StatefulSet by following rule:   - Pod: /spec/containers/<N>/image   - ReplicaSet: /spec/template/spec/containers/<N>/image   - Deployment: /spec/template/spec/containers/<N>/image   - StatefulSet: /spec/template/spec/containers/<N>/image In addition, all images will be processed if the resource object has more than one containers.  If not nil, only images matches the filters will be processed.",
																Attributes: map[string]schema.Attribute{
																	"path": schema.StringAttribute{
																		Description:         "Path indicates the path of target field",
																		MarkdownDescription: "Path indicates the path of target field",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": schema.StringAttribute{
																Description:         "Value to be applied to image. Must not be empty when operator is 'add' or 'replace'. Defaults to empty and ignored when operator is 'remove'.",
																MarkdownDescription: "Value to be applied to image. Must not be empty when operator is 'add' or 'replace'. Defaults to empty and ignored when operator is 'remove'.",
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

												"replicas": schema.Int64Attribute{
													Description:         "Replicas will override the replicas field of deployment",
													MarkdownDescription: "Replicas will override the replicas field of deployment",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"resources_overriders": schema.ListNestedAttribute{
													Description:         "ResourcesOverriders will override the resources field of the container",
													MarkdownDescription: "ResourcesOverriders will override the resources field of the container",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"container_name": schema.StringAttribute{
																Description:         "The name of container",
																MarkdownDescription: "The name of container",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"value": schema.SingleNestedAttribute{
																Description:         "Value to be applied to resources. Must not be empty",
																MarkdownDescription: "Value to be applied to resources. Must not be empty",
																Attributes: map[string]schema.Attribute{
																	"claims": schema.ListNestedAttribute{
																		Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																		MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																					MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

																	"limits": schema.MapAttribute{
																		Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																		MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"requests": schema.MapAttribute{
																		Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																		MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"workload_template": schema.SingleNestedAttribute{
						Description:         "WorkloadTemplate contains original templates of resources to be deployed as an EdgeApplication.",
						MarkdownDescription: "WorkloadTemplate contains original templates of resources to be deployed as an EdgeApplication.",
						Attributes: map[string]schema.Attribute{
							"manifests": schema.ListAttribute{
								Description:         "Manifests represent a list of Kubernetes resources to be deployed on the managed node groups.",
								MarkdownDescription: "Manifests represent a list of Kubernetes resources to be deployed on the managed node groups.",
								ElementType:         types.MapType{ElemType: types.StringType},
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

func (r *AppsKubeedgeIoEdgeApplicationV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_kubeedge_io_edge_application_v1alpha1_manifest")

	var model AppsKubeedgeIoEdgeApplicationV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("apps.kubeedge.io/v1alpha1")
	model.Kind = pointer.String("EdgeApplication")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
