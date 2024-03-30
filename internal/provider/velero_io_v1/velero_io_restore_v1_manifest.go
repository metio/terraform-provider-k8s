/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package velero_io_v1

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
	_ datasource.DataSource = &VeleroIoRestoreV1Manifest{}
)

func NewVeleroIoRestoreV1Manifest() datasource.DataSource {
	return &VeleroIoRestoreV1Manifest{}
}

type VeleroIoRestoreV1Manifest struct{}

type VeleroIoRestoreV1ManifestData struct {
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
		BackupName             *string   `tfsdk:"backup_name" json:"backupName,omitempty"`
		ExcludedNamespaces     *[]string `tfsdk:"excluded_namespaces" json:"excludedNamespaces,omitempty"`
		ExcludedResources      *[]string `tfsdk:"excluded_resources" json:"excludedResources,omitempty"`
		ExistingResourcePolicy *string   `tfsdk:"existing_resource_policy" json:"existingResourcePolicy,omitempty"`
		Hooks                  *struct {
			Resources *[]struct {
				ExcludedNamespaces *[]string `tfsdk:"excluded_namespaces" json:"excludedNamespaces,omitempty"`
				ExcludedResources  *[]string `tfsdk:"excluded_resources" json:"excludedResources,omitempty"`
				IncludedNamespaces *[]string `tfsdk:"included_namespaces" json:"includedNamespaces,omitempty"`
				IncludedResources  *[]string `tfsdk:"included_resources" json:"includedResources,omitempty"`
				LabelSelector      *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				PostHooks *[]struct {
					Exec *struct {
						Command      *[]string `tfsdk:"command" json:"command,omitempty"`
						Container    *string   `tfsdk:"container" json:"container,omitempty"`
						ExecTimeout  *string   `tfsdk:"exec_timeout" json:"execTimeout,omitempty"`
						OnError      *string   `tfsdk:"on_error" json:"onError,omitempty"`
						WaitForReady *bool     `tfsdk:"wait_for_ready" json:"waitForReady,omitempty"`
						WaitTimeout  *string   `tfsdk:"wait_timeout" json:"waitTimeout,omitempty"`
					} `tfsdk:"exec" json:"exec,omitempty"`
					Init *struct {
						InitContainers *map[string]string `tfsdk:"init_containers" json:"initContainers,omitempty"`
						Timeout        *string            `tfsdk:"timeout" json:"timeout,omitempty"`
					} `tfsdk:"init" json:"init,omitempty"`
				} `tfsdk:"post_hooks" json:"postHooks,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
		} `tfsdk:"hooks" json:"hooks,omitempty"`
		IncludeClusterResources *bool     `tfsdk:"include_cluster_resources" json:"includeClusterResources,omitempty"`
		IncludedNamespaces      *[]string `tfsdk:"included_namespaces" json:"includedNamespaces,omitempty"`
		IncludedResources       *[]string `tfsdk:"included_resources" json:"includedResources,omitempty"`
		ItemOperationTimeout    *string   `tfsdk:"item_operation_timeout" json:"itemOperationTimeout,omitempty"`
		LabelSelector           *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
		NamespaceMapping *map[string]string `tfsdk:"namespace_mapping" json:"namespaceMapping,omitempty"`
		OrLabelSelectors *[]struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"or_label_selectors" json:"orLabelSelectors,omitempty"`
		PreserveNodePorts *bool `tfsdk:"preserve_node_ports" json:"preserveNodePorts,omitempty"`
		ResourceModifier  *struct {
			ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
			Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"resource_modifier" json:"resourceModifier,omitempty"`
		RestorePVs    *bool `tfsdk:"restore_p_vs" json:"restorePVs,omitempty"`
		RestoreStatus *struct {
			ExcludedResources *[]string `tfsdk:"excluded_resources" json:"excludedResources,omitempty"`
			IncludedResources *[]string `tfsdk:"included_resources" json:"includedResources,omitempty"`
		} `tfsdk:"restore_status" json:"restoreStatus,omitempty"`
		ScheduleName   *string `tfsdk:"schedule_name" json:"scheduleName,omitempty"`
		UploaderConfig *struct {
			WriteSparseFiles *bool `tfsdk:"write_sparse_files" json:"writeSparseFiles,omitempty"`
		} `tfsdk:"uploader_config" json:"uploaderConfig,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *VeleroIoRestoreV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_velero_io_restore_v1_manifest"
}

func (r *VeleroIoRestoreV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Restore is a Velero resource that represents the application of resources from a Velero backup to a target Kubernetes cluster.",
		MarkdownDescription: "Restore is a Velero resource that represents the application of resources from a Velero backup to a target Kubernetes cluster.",
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
				Description:         "RestoreSpec defines the specification for a Velero restore.",
				MarkdownDescription: "RestoreSpec defines the specification for a Velero restore.",
				Attributes: map[string]schema.Attribute{
					"backup_name": schema.StringAttribute{
						Description:         "BackupName is the unique name of the Velero backup to restore from.",
						MarkdownDescription: "BackupName is the unique name of the Velero backup to restore from.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"excluded_namespaces": schema.ListAttribute{
						Description:         "ExcludedNamespaces contains a list of namespaces that are not included in the restore.",
						MarkdownDescription: "ExcludedNamespaces contains a list of namespaces that are not included in the restore.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"excluded_resources": schema.ListAttribute{
						Description:         "ExcludedResources is a slice of resource names that are not included in the restore.",
						MarkdownDescription: "ExcludedResources is a slice of resource names that are not included in the restore.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"existing_resource_policy": schema.StringAttribute{
						Description:         "ExistingResourcePolicy specifies the restore behavior for the Kubernetes resource to be restored",
						MarkdownDescription: "ExistingResourcePolicy specifies the restore behavior for the Kubernetes resource to be restored",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"hooks": schema.SingleNestedAttribute{
						Description:         "Hooks represent custom behaviors that should be executed during or post restore.",
						MarkdownDescription: "Hooks represent custom behaviors that should be executed during or post restore.",
						Attributes: map[string]schema.Attribute{
							"resources": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"excluded_namespaces": schema.ListAttribute{
											Description:         "ExcludedNamespaces specifies the namespaces to which this hook spec does not apply.",
											MarkdownDescription: "ExcludedNamespaces specifies the namespaces to which this hook spec does not apply.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"excluded_resources": schema.ListAttribute{
											Description:         "ExcludedResources specifies the resources to which this hook spec does not apply.",
											MarkdownDescription: "ExcludedResources specifies the resources to which this hook spec does not apply.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"included_namespaces": schema.ListAttribute{
											Description:         "IncludedNamespaces specifies the namespaces to which this hook spec applies. If empty, it applies to all namespaces.",
											MarkdownDescription: "IncludedNamespaces specifies the namespaces to which this hook spec applies. If empty, it applies to all namespaces.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"included_resources": schema.ListAttribute{
											Description:         "IncludedResources specifies the resources to which this hook spec applies. If empty, it applies to all resources.",
											MarkdownDescription: "IncludedResources specifies the resources to which this hook spec applies. If empty, it applies to all resources.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"label_selector": schema.SingleNestedAttribute{
											Description:         "LabelSelector, if specified, filters the resources to which this hook spec applies.",
											MarkdownDescription: "LabelSelector, if specified, filters the resources to which this hook spec applies.",
											Attributes: map[string]schema.Attribute{
												"match_expressions": schema.ListNestedAttribute{
													Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"values": schema.ListAttribute{
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

												"match_labels": schema.MapAttribute{
													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

										"name": schema.StringAttribute{
											Description:         "Name is the name of this hook.",
											MarkdownDescription: "Name is the name of this hook.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"post_hooks": schema.ListNestedAttribute{
											Description:         "PostHooks is a list of RestoreResourceHooks to execute during and after restoring a resource.",
											MarkdownDescription: "PostHooks is a list of RestoreResourceHooks to execute during and after restoring a resource.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"exec": schema.SingleNestedAttribute{
														Description:         "Exec defines an exec restore hook.",
														MarkdownDescription: "Exec defines an exec restore hook.",
														Attributes: map[string]schema.Attribute{
															"command": schema.ListAttribute{
																Description:         "Command is the command and arguments to execute from within a container after a pod has been restored.",
																MarkdownDescription: "Command is the command and arguments to execute from within a container after a pod has been restored.",
																ElementType:         types.StringType,
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"container": schema.StringAttribute{
																Description:         "Container is the container in the pod where the command should be executed. If not specified, the pod's first container is used.",
																MarkdownDescription: "Container is the container in the pod where the command should be executed. If not specified, the pod's first container is used.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"exec_timeout": schema.StringAttribute{
																Description:         "ExecTimeout defines the maximum amount of time Velero should wait for the hook to complete before considering the execution a failure.",
																MarkdownDescription: "ExecTimeout defines the maximum amount of time Velero should wait for the hook to complete before considering the execution a failure.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"on_error": schema.StringAttribute{
																Description:         "OnError specifies how Velero should behave if it encounters an error executing this hook.",
																MarkdownDescription: "OnError specifies how Velero should behave if it encounters an error executing this hook.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("Continue", "Fail"),
																},
															},

															"wait_for_ready": schema.BoolAttribute{
																Description:         "WaitForReady ensures command will be launched when container is Ready instead of Running.",
																MarkdownDescription: "WaitForReady ensures command will be launched when container is Ready instead of Running.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"wait_timeout": schema.StringAttribute{
																Description:         "WaitTimeout defines the maximum amount of time Velero should wait for the container to be Ready before attempting to run the command.",
																MarkdownDescription: "WaitTimeout defines the maximum amount of time Velero should wait for the container to be Ready before attempting to run the command.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"init": schema.SingleNestedAttribute{
														Description:         "Init defines an init restore hook.",
														MarkdownDescription: "Init defines an init restore hook.",
														Attributes: map[string]schema.Attribute{
															"init_containers": schema.MapAttribute{
																Description:         "InitContainers is list of init containers to be added to a pod during its restore.",
																MarkdownDescription: "InitContainers is list of init containers to be added to a pod during its restore.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"timeout": schema.StringAttribute{
																Description:         "Timeout defines the maximum amount of time Velero should wait for the initContainers to complete.",
																MarkdownDescription: "Timeout defines the maximum amount of time Velero should wait for the initContainers to complete.",
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

					"include_cluster_resources": schema.BoolAttribute{
						Description:         "IncludeClusterResources specifies whether cluster-scoped resources should be included for consideration in the restore. If null, defaults to true.",
						MarkdownDescription: "IncludeClusterResources specifies whether cluster-scoped resources should be included for consideration in the restore. If null, defaults to true.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"included_namespaces": schema.ListAttribute{
						Description:         "IncludedNamespaces is a slice of namespace names to include objects from. If empty, all namespaces are included.",
						MarkdownDescription: "IncludedNamespaces is a slice of namespace names to include objects from. If empty, all namespaces are included.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"included_resources": schema.ListAttribute{
						Description:         "IncludedResources is a slice of resource names to include in the restore. If empty, all resources in the backup are included.",
						MarkdownDescription: "IncludedResources is a slice of resource names to include in the restore. If empty, all resources in the backup are included.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"item_operation_timeout": schema.StringAttribute{
						Description:         "ItemOperationTimeout specifies the time used to wait for RestoreItemAction operations The default value is 4 hour.",
						MarkdownDescription: "ItemOperationTimeout specifies the time used to wait for RestoreItemAction operations The default value is 4 hour.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"label_selector": schema.SingleNestedAttribute{
						Description:         "LabelSelector is a metav1.LabelSelector to filter with when restoring individual objects from the backup. If empty or nil, all objects are included. Optional.",
						MarkdownDescription: "LabelSelector is a metav1.LabelSelector to filter with when restoring individual objects from the backup. If empty or nil, all objects are included. Optional.",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

					"namespace_mapping": schema.MapAttribute{
						Description:         "NamespaceMapping is a map of source namespace names to target namespace names to restore into. Any source namespaces not included in the map will be restored into namespaces of the same name.",
						MarkdownDescription: "NamespaceMapping is a map of source namespace names to target namespace names to restore into. Any source namespaces not included in the map will be restored into namespaces of the same name.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"or_label_selectors": schema.ListNestedAttribute{
						Description:         "OrLabelSelectors is list of metav1.LabelSelector to filter with when restoring individual objects from the backup. If multiple provided they will be joined by the OR operator. LabelSelector as well as OrLabelSelectors cannot co-exist in restore request, only one of them can be used",
						MarkdownDescription: "OrLabelSelectors is list of metav1.LabelSelector to filter with when restoring individual objects from the backup. If multiple provided they will be joined by the OR operator. LabelSelector as well as OrLabelSelectors cannot co-exist in restore request, only one of them can be used",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"match_expressions": schema.ListNestedAttribute{
									Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
									MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"operator": schema.StringAttribute{
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"values": schema.ListAttribute{
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

								"match_labels": schema.MapAttribute{
									Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
									MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

					"preserve_node_ports": schema.BoolAttribute{
						Description:         "PreserveNodePorts specifies whether to restore old nodePorts from backup.",
						MarkdownDescription: "PreserveNodePorts specifies whether to restore old nodePorts from backup.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resource_modifier": schema.SingleNestedAttribute{
						Description:         "ResourceModifier specifies the reference to JSON resource patches that should be applied to resources before restoration.",
						MarkdownDescription: "ResourceModifier specifies the reference to JSON resource patches that should be applied to resources before restoration.",
						Attributes: map[string]schema.Attribute{
							"api_group": schema.StringAttribute{
								Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
								MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind is the type of resource being referenced",
								MarkdownDescription: "Kind is the type of resource being referenced",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of resource being referenced",
								MarkdownDescription: "Name is the name of resource being referenced",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"restore_p_vs": schema.BoolAttribute{
						Description:         "RestorePVs specifies whether to restore all included PVs from snapshot",
						MarkdownDescription: "RestorePVs specifies whether to restore all included PVs from snapshot",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"restore_status": schema.SingleNestedAttribute{
						Description:         "RestoreStatus specifies which resources we should restore the status field. If nil, no objects are included. Optional.",
						MarkdownDescription: "RestoreStatus specifies which resources we should restore the status field. If nil, no objects are included. Optional.",
						Attributes: map[string]schema.Attribute{
							"excluded_resources": schema.ListAttribute{
								Description:         "ExcludedResources specifies the resources to which will not restore the status.",
								MarkdownDescription: "ExcludedResources specifies the resources to which will not restore the status.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"included_resources": schema.ListAttribute{
								Description:         "IncludedResources specifies the resources to which will restore the status. If empty, it applies to all resources.",
								MarkdownDescription: "IncludedResources specifies the resources to which will restore the status. If empty, it applies to all resources.",
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

					"schedule_name": schema.StringAttribute{
						Description:         "ScheduleName is the unique name of the Velero schedule to restore from. If specified, and BackupName is empty, Velero will restore from the most recent successful backup created from this schedule.",
						MarkdownDescription: "ScheduleName is the unique name of the Velero schedule to restore from. If specified, and BackupName is empty, Velero will restore from the most recent successful backup created from this schedule.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"uploader_config": schema.SingleNestedAttribute{
						Description:         "UploaderConfig specifies the configuration for the restore.",
						MarkdownDescription: "UploaderConfig specifies the configuration for the restore.",
						Attributes: map[string]schema.Attribute{
							"write_sparse_files": schema.BoolAttribute{
								Description:         "WriteSparseFiles is a flag to indicate whether write files sparsely or not.",
								MarkdownDescription: "WriteSparseFiles is a flag to indicate whether write files sparsely or not.",
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

func (r *VeleroIoRestoreV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_velero_io_restore_v1_manifest")

	var model VeleroIoRestoreV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("velero.io/v1")
	model.Kind = pointer.String("Restore")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
