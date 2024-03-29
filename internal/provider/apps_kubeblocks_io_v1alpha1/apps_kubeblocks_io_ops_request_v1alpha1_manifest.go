/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_kubeblocks_io_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &AppsKubeblocksIoOpsRequestV1Alpha1Manifest{}
)

func NewAppsKubeblocksIoOpsRequestV1Alpha1Manifest() datasource.DataSource {
	return &AppsKubeblocksIoOpsRequestV1Alpha1Manifest{}
}

type AppsKubeblocksIoOpsRequestV1Alpha1Manifest struct{}

type AppsKubeblocksIoOpsRequestV1Alpha1ManifestData struct {
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
		BackupSpec *struct {
			BackupMethod     *string `tfsdk:"backup_method" json:"backupMethod,omitempty"`
			BackupName       *string `tfsdk:"backup_name" json:"backupName,omitempty"`
			BackupPolicyName *string `tfsdk:"backup_policy_name" json:"backupPolicyName,omitempty"`
			DeletionPolicy   *string `tfsdk:"deletion_policy" json:"deletionPolicy,omitempty"`
			ParentBackupName *string `tfsdk:"parent_backup_name" json:"parentBackupName,omitempty"`
			RetentionPeriod  *string `tfsdk:"retention_period" json:"retentionPeriod,omitempty"`
		} `tfsdk:"backup_spec" json:"backupSpec,omitempty"`
		Cancel     *bool   `tfsdk:"cancel" json:"cancel,omitempty"`
		ClusterRef *string `tfsdk:"cluster_ref" json:"clusterRef,omitempty"`
		CustomSpec *struct {
			Components *[]struct {
				Name       *string `tfsdk:"name" json:"name,omitempty"`
				Parameters *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"parameters" json:"parameters,omitempty"`
			} `tfsdk:"components" json:"components,omitempty"`
			OpsDefinitionRef   *string `tfsdk:"ops_definition_ref" json:"opsDefinitionRef,omitempty"`
			Parallelism        *string `tfsdk:"parallelism" json:"parallelism,omitempty"`
			ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
		} `tfsdk:"custom_spec" json:"customSpec,omitempty"`
		Expose *[]struct {
			ComponentName *string `tfsdk:"component_name" json:"componentName,omitempty"`
			Services      *[]struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
				Ports       *[]struct {
					AppProtocol *string `tfsdk:"app_protocol" json:"appProtocol,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					NodePort    *int64  `tfsdk:"node_port" json:"nodePort,omitempty"`
					Port        *int64  `tfsdk:"port" json:"port,omitempty"`
					Protocol    *string `tfsdk:"protocol" json:"protocol,omitempty"`
					TargetPort  *string `tfsdk:"target_port" json:"targetPort,omitempty"`
				} `tfsdk:"ports" json:"ports,omitempty"`
				RoleSelector *string            `tfsdk:"role_selector" json:"roleSelector,omitempty"`
				Selector     *map[string]string `tfsdk:"selector" json:"selector,omitempty"`
				ServiceType  *string            `tfsdk:"service_type" json:"serviceType,omitempty"`
			} `tfsdk:"services" json:"services,omitempty"`
			Switch *string `tfsdk:"switch" json:"switch,omitempty"`
		} `tfsdk:"expose" json:"expose,omitempty"`
		HorizontalScaling *[]struct {
			ComponentName *string   `tfsdk:"component_name" json:"componentName,omitempty"`
			Instances     *[]string `tfsdk:"instances" json:"instances,omitempty"`
			Nodes         *[]string `tfsdk:"nodes" json:"nodes,omitempty"`
			Replicas      *int64    `tfsdk:"replicas" json:"replicas,omitempty"`
		} `tfsdk:"horizontal_scaling" json:"horizontalScaling,omitempty"`
		Reconfigure *struct {
			ComponentName  *string `tfsdk:"component_name" json:"componentName,omitempty"`
			Configurations *[]struct {
				Keys *[]struct {
					FileContent *string `tfsdk:"file_content" json:"fileContent,omitempty"`
					Key         *string `tfsdk:"key" json:"key,omitempty"`
					Parameters  *[]struct {
						Key   *string `tfsdk:"key" json:"key,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"parameters" json:"parameters,omitempty"`
				} `tfsdk:"keys" json:"keys,omitempty"`
				Name   *string `tfsdk:"name" json:"name,omitempty"`
				Policy *string `tfsdk:"policy" json:"policy,omitempty"`
			} `tfsdk:"configurations" json:"configurations,omitempty"`
		} `tfsdk:"reconfigure" json:"reconfigure,omitempty"`
		Reconfigures *[]struct {
			ComponentName  *string `tfsdk:"component_name" json:"componentName,omitempty"`
			Configurations *[]struct {
				Keys *[]struct {
					FileContent *string `tfsdk:"file_content" json:"fileContent,omitempty"`
					Key         *string `tfsdk:"key" json:"key,omitempty"`
					Parameters  *[]struct {
						Key   *string `tfsdk:"key" json:"key,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"parameters" json:"parameters,omitempty"`
				} `tfsdk:"keys" json:"keys,omitempty"`
				Name   *string `tfsdk:"name" json:"name,omitempty"`
				Policy *string `tfsdk:"policy" json:"policy,omitempty"`
			} `tfsdk:"configurations" json:"configurations,omitempty"`
		} `tfsdk:"reconfigures" json:"reconfigures,omitempty"`
		Restart *[]struct {
			ComponentName *string `tfsdk:"component_name" json:"componentName,omitempty"`
		} `tfsdk:"restart" json:"restart,omitempty"`
		RestoreFrom *struct {
			Backup *[]struct {
				Ref *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"ref" json:"ref,omitempty"`
			} `tfsdk:"backup" json:"backup,omitempty"`
			PointInTime *struct {
				Ref *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"ref" json:"ref,omitempty"`
				Time *string `tfsdk:"time" json:"time,omitempty"`
			} `tfsdk:"point_in_time" json:"pointInTime,omitempty"`
		} `tfsdk:"restore_from" json:"restoreFrom,omitempty"`
		RestoreSpec *struct {
			BackupName                  *string `tfsdk:"backup_name" json:"backupName,omitempty"`
			EffectiveCommonComponentDef *bool   `tfsdk:"effective_common_component_def" json:"effectiveCommonComponentDef,omitempty"`
			RestoreTimeStr              *string `tfsdk:"restore_time_str" json:"restoreTimeStr,omitempty"`
			VolumeRestorePolicy         *string `tfsdk:"volume_restore_policy" json:"volumeRestorePolicy,omitempty"`
		} `tfsdk:"restore_spec" json:"restoreSpec,omitempty"`
		ScriptSpec *struct {
			ComponentName *string   `tfsdk:"component_name" json:"componentName,omitempty"`
			Image         *string   `tfsdk:"image" json:"image,omitempty"`
			Script        *[]string `tfsdk:"script" json:"script,omitempty"`
			ScriptFrom    *struct {
				ConfigMapRef *[]struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
				SecretRef *[]struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
			} `tfsdk:"script_from" json:"scriptFrom,omitempty"`
			Secret *struct {
				Name        *string `tfsdk:"name" json:"name,omitempty"`
				PasswordKey *string `tfsdk:"password_key" json:"passwordKey,omitempty"`
				UsernameKey *string `tfsdk:"username_key" json:"usernameKey,omitempty"`
			} `tfsdk:"secret" json:"secret,omitempty"`
			Selector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"selector" json:"selector,omitempty"`
		} `tfsdk:"script_spec" json:"scriptSpec,omitempty"`
		Switchover *[]struct {
			ComponentName *string `tfsdk:"component_name" json:"componentName,omitempty"`
			InstanceName  *string `tfsdk:"instance_name" json:"instanceName,omitempty"`
		} `tfsdk:"switchover" json:"switchover,omitempty"`
		TtlSecondsAfterSucceed *int64  `tfsdk:"ttl_seconds_after_succeed" json:"ttlSecondsAfterSucceed,omitempty"`
		TtlSecondsBeforeAbort  *int64  `tfsdk:"ttl_seconds_before_abort" json:"ttlSecondsBeforeAbort,omitempty"`
		Type                   *string `tfsdk:"type" json:"type,omitempty"`
		Upgrade                *struct {
			ClusterVersionRef *string `tfsdk:"cluster_version_ref" json:"clusterVersionRef,omitempty"`
		} `tfsdk:"upgrade" json:"upgrade,omitempty"`
		VerticalScaling *[]map[string]string `tfsdk:"vertical_scaling" json:"verticalScaling,omitempty"`
		VolumeExpansion *[]struct {
			ComponentName        *string `tfsdk:"component_name" json:"componentName,omitempty"`
			VolumeClaimTemplates *[]struct {
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Storage *string `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"volume_claim_templates" json:"volumeClaimTemplates,omitempty"`
		} `tfsdk:"volume_expansion" json:"volumeExpansion,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsKubeblocksIoOpsRequestV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_kubeblocks_io_ops_request_v1alpha1_manifest"
}

func (r *AppsKubeblocksIoOpsRequestV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "OpsRequest is the Schema for the opsrequests API",
		MarkdownDescription: "OpsRequest is the Schema for the opsrequests API",
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
				Description:         "OpsRequestSpec defines the desired state of OpsRequest",
				MarkdownDescription: "OpsRequestSpec defines the desired state of OpsRequest",
				Attributes: map[string]schema.Attribute{
					"backup_spec": schema.SingleNestedAttribute{
						Description:         "Defines how to backup the cluster.",
						MarkdownDescription: "Defines how to backup the cluster.",
						Attributes: map[string]schema.Attribute{
							"backup_method": schema.StringAttribute{
								Description:         "Defines the backup method that is defined in backupPolicy.",
								MarkdownDescription: "Defines the backup method that is defined in backupPolicy.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"backup_name": schema.StringAttribute{
								Description:         "Specifies the name of the backup.",
								MarkdownDescription: "Specifies the name of the backup.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"backup_policy_name": schema.StringAttribute{
								Description:         "Indicates the backupPolicy applied to perform this backup.",
								MarkdownDescription: "Indicates the backupPolicy applied to perform this backup.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"deletion_policy": schema.StringAttribute{
								Description:         "Determines whether the backup contents stored in backup repository should be deleted when the backup custom resource is deleted. Supported values are 'Retain' and 'Delete'. - 'Retain' means that the backup content and its physical snapshot on backup repository are kept. - 'Delete' means that the backup content and its physical snapshot on backup repository are deleted.",
								MarkdownDescription: "Determines whether the backup contents stored in backup repository should be deleted when the backup custom resource is deleted. Supported values are 'Retain' and 'Delete'. - 'Retain' means that the backup content and its physical snapshot on backup repository are kept. - 'Delete' means that the backup content and its physical snapshot on backup repository are deleted.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Delete", "Retain"),
								},
							},

							"parent_backup_name": schema.StringAttribute{
								Description:         "If backupType is incremental, parentBackupName is required.",
								MarkdownDescription: "If backupType is incremental, parentBackupName is required.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"retention_period": schema.StringAttribute{
								Description:         "Determines a duration up to which the backup should be kept. Controller will remove all backups that are older than the RetentionPeriod. For example, RetentionPeriod of '30d' will keep only the backups of last 30 days. Sample duration format:  - years: 2y - months: 6mo - days: 30d - hours: 12h - minutes: 30m  You can also combine the above durations. For example: 30d12h30m. If not set, the backup will be kept forever.",
								MarkdownDescription: "Determines a duration up to which the backup should be kept. Controller will remove all backups that are older than the RetentionPeriod. For example, RetentionPeriod of '30d' will keep only the backups of last 30 days. Sample duration format:  - years: 2y - months: 6mo - days: 30d - hours: 12h - minutes: 30m  You can also combine the above durations. For example: 30d12h30m. If not set, the backup will be kept forever.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cancel": schema.BoolAttribute{
						Description:         "Defines the action to cancel the 'Pending/Creating/Running' opsRequest, supported types: 'VerticalScaling/HorizontalScaling'. Once set to true, this opsRequest will be canceled and modifying this property again will not take effect.",
						MarkdownDescription: "Defines the action to cancel the 'Pending/Creating/Running' opsRequest, supported types: 'VerticalScaling/HorizontalScaling'. Once set to true, this opsRequest will be canceled and modifying this property again will not take effect.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_ref": schema.StringAttribute{
						Description:         "References the cluster object.",
						MarkdownDescription: "References the cluster object.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"custom_spec": schema.SingleNestedAttribute{
						Description:         "Specifies a custom operation as defined by OpsDefinition.",
						MarkdownDescription: "Specifies a custom operation as defined by OpsDefinition.",
						Attributes: map[string]schema.Attribute{
							"components": schema.ListNestedAttribute{
								Description:         "Defines which components need to perform the actions defined by this OpsDefinition. At least one component is required. The components are identified by their name and can be merged or retained.",
								MarkdownDescription: "Defines which components need to perform the actions defined by this OpsDefinition. At least one component is required. The components are identified by their name and can be merged or retained.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Specifies the unique identifier of the cluster component",
											MarkdownDescription: "Specifies the unique identifier of the cluster component",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"parameters": schema.ListNestedAttribute{
											Description:         "Represents the parameters for this operation as declared in the opsDefinition.spec.parametersSchema.",
											MarkdownDescription: "Represents the parameters for this operation as declared in the opsDefinition.spec.parametersSchema.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Specifies the identifier of the parameter as defined in the OpsDefinition.",
														MarkdownDescription: "Specifies the identifier of the parameter as defined in the OpsDefinition.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"value": schema.StringAttribute{
														Description:         "Holds the data associated with the parameter. If the parameter type is an array, the format should be 'v1,v2,v3'.",
														MarkdownDescription: "Holds the data associated with the parameter. If the parameter type is an array, the format should be 'v1,v2,v3'.",
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
								Required: true,
								Optional: false,
								Computed: false,
							},

							"ops_definition_ref": schema.StringAttribute{
								Description:         "Is a reference to an OpsDefinition.",
								MarkdownDescription: "Is a reference to an OpsDefinition.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"parallelism": schema.StringAttribute{
								Description:         "Defines the execution concurrency. By default, all incoming Components will be executed simultaneously. The value can be an absolute number (e.g., 5) or a percentage of desired components (e.g., 10%). The absolute number is calculated from the percentage by rounding up. For instance, if the percentage value is 10% and the components length is 1, the calculated number will be rounded up to 1.",
								MarkdownDescription: "Defines the execution concurrency. By default, all incoming Components will be executed simultaneously. The value can be an absolute number (e.g., 5) or a percentage of desired components (e.g., 10%). The absolute number is calculated from the percentage by rounding up. For instance, if the percentage value is 10% and the components length is 1, the calculated number will be rounded up to 1.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_account_name": schema.StringAttribute{
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

					"expose": schema.ListNestedAttribute{
						Description:         "Defines services the component needs to expose.",
						MarkdownDescription: "Defines services the component needs to expose.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the cluster component.",
									MarkdownDescription: "Specifies the name of the cluster component.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"services": schema.ListNestedAttribute{
									Description:         "A list of services that are to be exposed or removed. If componentNamem is not specified, each 'OpsService' in the list must specify ports and selectors.",
									MarkdownDescription: "A list of services that are to be exposed or removed. If componentNamem is not specified, each 'OpsService' in the list must specify ports and selectors.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Contains cloud provider related parameters if ServiceType is LoadBalancer. More info: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer.",
												MarkdownDescription: "Contains cloud provider related parameters if ServiceType is LoadBalancer. More info: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Specifies the name of the service. This name is used by others to refer to this service (e.g., connection credential). Note: This field cannot be updated.",
												MarkdownDescription: "Specifies the name of the service. This name is used by others to refer to this service (e.g., connection credential). Note: This field cannot be updated.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"ports": schema.ListNestedAttribute{
												Description:         "Lists the ports that are exposed by this service. If not provided, the default Services Ports defined in the ClusterDefinition or ComponentDefinition that are neither of NodePort nor LoadBalancer service type will be used. If there is no corresponding Service defined in the ClusterDefinition or ComponentDefinition, the expose operation will fail. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												MarkdownDescription: "Lists the ports that are exposed by this service. If not provided, the default Services Ports defined in the ClusterDefinition or ComponentDefinition that are neither of NodePort nor LoadBalancer service type will be used. If there is no corresponding Service defined in the ClusterDefinition or ComponentDefinition, the expose operation will fail. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"app_protocol": schema.StringAttribute{
															Description:         "The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either:  * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names).  * Kubernetes-defined prefixed names: * 'kubernetes.io/h2c' - HTTP/2 over cleartext as described in https://www.rfc-editor.org/rfc/rfc7540 * 'kubernetes.io/ws'  - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455 * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455  * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.",
															MarkdownDescription: "The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either:  * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names).  * Kubernetes-defined prefixed names: * 'kubernetes.io/h2c' - HTTP/2 over cleartext as described in https://www.rfc-editor.org/rfc/rfc7540 * 'kubernetes.io/ws'  - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455 * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455  * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the 'name' field in the EndpointPort. Optional if only one ServicePort is defined on this service.",
															MarkdownDescription: "The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the 'name' field in the EndpointPort. Optional if only one ServicePort is defined on this service.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"node_port": schema.Int64Attribute{
															Description:         "The port on each node on which this service is exposed when type is NodePort or LoadBalancer.  Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail.  If not specified, a port will be allocated if this Service requires one.  If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type from NodePort to ClusterIP). More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport",
															MarkdownDescription: "The port on each node on which this service is exposed when type is NodePort or LoadBalancer.  Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail.  If not specified, a port will be allocated if this Service requires one.  If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type from NodePort to ClusterIP). More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.Int64Attribute{
															Description:         "The port that will be exposed by this service.",
															MarkdownDescription: "The port that will be exposed by this service.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"protocol": schema.StringAttribute{
															Description:         "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'. Default is TCP.",
															MarkdownDescription: "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'. Default is TCP.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"target_port": schema.StringAttribute{
															Description:         "Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If this is a string, it will be looked up as a named port in the target Pod's container ports. If this is not specified, the value of the 'port' field is used (an identity map). This field is ignored for services with clusterIP=None, and should be omitted or set equal to the 'port' field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service",
															MarkdownDescription: "Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If this is a string, it will be looked up as a named port in the target Pod's container ports. If this is not specified, the value of the 'port' field is used (an identity map). This field is ignored for services with clusterIP=None, and should be omitted or set equal to the 'port' field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service",
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

											"role_selector": schema.StringAttribute{
												Description:         "Allows you to specify a defined role as a selector for the service, extending the ServiceSpec.Selector.",
												MarkdownDescription: "Allows you to specify a defined role as a selector for the service, extending the ServiceSpec.Selector.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"selector": schema.MapAttribute{
												Description:         "Routes service traffic to pods with label keys and values matching this selector. If empty or not present, the service is assumed to have an external process managing its endpoints, which Kubernetes will not modify. This only applies to types ClusterIP, NodePort, and LoadBalancer and is ignored if type is ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/",
												MarkdownDescription: "Routes service traffic to pods with label keys and values matching this selector. If empty or not present, the service is assumed to have an external process managing its endpoints, which Kubernetes will not modify. This only applies to types ClusterIP, NodePort, and LoadBalancer and is ignored if type is ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"service_type": schema.StringAttribute{
												Description:         "Determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. - 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. - 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. - 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types.",
												MarkdownDescription: "Determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. - 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. - 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. - 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types.",
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

								"switch": schema.StringAttribute{
									Description:         "Controls the expose operation. If set to Enable, the corresponding service will be exposed. Conversely, if set to Disable, the service will be removed.",
									MarkdownDescription: "Controls the expose operation. If set to Enable, the corresponding service will be exposed. Conversely, if set to Disable, the service will be removed.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Enable", "Disable"),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"horizontal_scaling": schema.ListNestedAttribute{
						Description:         "Defines what component need to horizontal scale the specified replicas.",
						MarkdownDescription: "Defines what component need to horizontal scale the specified replicas.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the cluster component.",
									MarkdownDescription: "Specifies the name of the cluster component.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"instances": schema.ListAttribute{
									Description:         "Defines the names of instances that the rsm should prioritize for scale-down operations. If the RsmTransformPolicy is set to ToPod and the expected number of replicas is less than the current number, the list of Instances will be used.  - 'current replicas - expected replicas > len(Instances)': Scale down from the list of Instances priorly, the others will select from NodeAssignment. - 'current replicas - expected replicas < len(Instances)': Scale down from the list of Instances. - 'current replicas - expected replicas < len(Instances)': Scale down from a part of Instances.",
									MarkdownDescription: "Defines the names of instances that the rsm should prioritize for scale-down operations. If the RsmTransformPolicy is set to ToPod and the expected number of replicas is less than the current number, the list of Instances will be used.  - 'current replicas - expected replicas > len(Instances)': Scale down from the list of Instances priorly, the others will select from NodeAssignment. - 'current replicas - expected replicas < len(Instances)': Scale down from the list of Instances. - 'current replicas - expected replicas < len(Instances)': Scale down from a part of Instances.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"nodes": schema.ListAttribute{
									Description:         "Defines the list of nodes where pods can be scheduled during a scale-up operation. If the RsmTransformPolicy is set to ToPod and the expected number of replicas is greater than the current number, the list of Nodes will be used. If the list of Nodes is empty, pods will not be assigned to any specific node. However, if the list of Nodes is populated, pods will be evenly distributed across the nodes in the list during scale-up.",
									MarkdownDescription: "Defines the list of nodes where pods can be scheduled during a scale-up operation. If the RsmTransformPolicy is set to ToPod and the expected number of replicas is greater than the current number, the list of Nodes will be used. If the list of Nodes is empty, pods will not be assigned to any specific node. However, if the list of Nodes is populated, pods will be evenly distributed across the nodes in the list during scale-up.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"replicas": schema.Int64Attribute{
									Description:         "Specifies the number of replicas for the workloads.",
									MarkdownDescription: "Specifies the number of replicas for the workloads.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"reconfigure": schema.SingleNestedAttribute{
						Description:         "Deprecated: replace by reconfigures. Defines the variables that need to input when updating configuration.",
						MarkdownDescription: "Deprecated: replace by reconfigures. Defines the variables that need to input when updating configuration.",
						Attributes: map[string]schema.Attribute{
							"component_name": schema.StringAttribute{
								Description:         "Specifies the name of the cluster component.",
								MarkdownDescription: "Specifies the name of the cluster component.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"configurations": schema.ListNestedAttribute{
								Description:         "Specifies the components that will perform the operation.",
								MarkdownDescription: "Specifies the components that will perform the operation.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"keys": schema.ListNestedAttribute{
											Description:         "Sets the parameters to be updated. It should contain at least one item. The keys are merged and retained during patch operations.",
											MarkdownDescription: "Sets the parameters to be updated. It should contain at least one item. The keys are merged and retained during patch operations.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"file_content": schema.StringAttribute{
														Description:         "Represents the content of the configuration file. This field is used to update the entire content of the file.",
														MarkdownDescription: "Represents the content of the configuration file. This field is used to update the entire content of the file.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"key": schema.StringAttribute{
														Description:         "Represents the unique identifier for the ConfigMap.",
														MarkdownDescription: "Represents the unique identifier for the ConfigMap.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"parameters": schema.ListNestedAttribute{
														Description:         "Defines a list of key-value pairs for a single configuration file. These parameters are used to update the specified configuration settings.",
														MarkdownDescription: "Defines a list of key-value pairs for a single configuration file. These parameters are used to update the specified configuration settings.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "Represents the name of the parameter that is to be updated.",
																	MarkdownDescription: "Represents the name of the parameter that is to be updated.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "Represents the parameter values that are to be updated. If set to nil, the parameter defined by the Key field will be removed from the configuration file.",
																	MarkdownDescription: "Represents the parameter values that are to be updated. If set to nil, the parameter defined by the Key field will be removed from the configuration file.",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"name": schema.StringAttribute{
											Description:         "Specifies the name of the configuration template.",
											MarkdownDescription: "Specifies the name of the configuration template.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
											},
										},

										"policy": schema.StringAttribute{
											Description:         "Defines the upgrade policy for the configuration. This field is optional.",
											MarkdownDescription: "Defines the upgrade policy for the configuration. This field is optional.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("simple", "parallel", "rolling", "autoReload", "operatorSyncUpdate", "dynamicReloadBeginRestart"),
											},
										},
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"reconfigures": schema.ListNestedAttribute{
						Description:         "Defines the variables that need to input when updating configuration.",
						MarkdownDescription: "Defines the variables that need to input when updating configuration.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the cluster component.",
									MarkdownDescription: "Specifies the name of the cluster component.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"configurations": schema.ListNestedAttribute{
									Description:         "Specifies the components that will perform the operation.",
									MarkdownDescription: "Specifies the components that will perform the operation.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"keys": schema.ListNestedAttribute{
												Description:         "Sets the parameters to be updated. It should contain at least one item. The keys are merged and retained during patch operations.",
												MarkdownDescription: "Sets the parameters to be updated. It should contain at least one item. The keys are merged and retained during patch operations.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"file_content": schema.StringAttribute{
															Description:         "Represents the content of the configuration file. This field is used to update the entire content of the file.",
															MarkdownDescription: "Represents the content of the configuration file. This field is used to update the entire content of the file.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"key": schema.StringAttribute{
															Description:         "Represents the unique identifier for the ConfigMap.",
															MarkdownDescription: "Represents the unique identifier for the ConfigMap.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"parameters": schema.ListNestedAttribute{
															Description:         "Defines a list of key-value pairs for a single configuration file. These parameters are used to update the specified configuration settings.",
															MarkdownDescription: "Defines a list of key-value pairs for a single configuration file. These parameters are used to update the specified configuration settings.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "Represents the name of the parameter that is to be updated.",
																		MarkdownDescription: "Represents the name of the parameter that is to be updated.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "Represents the parameter values that are to be updated. If set to nil, the parameter defined by the Key field will be removed from the configuration file.",
																		MarkdownDescription: "Represents the parameter values that are to be updated. If set to nil, the parameter defined by the Key field will be removed from the configuration file.",
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
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": schema.StringAttribute{
												Description:         "Specifies the name of the configuration template.",
												MarkdownDescription: "Specifies the name of the configuration template.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(63),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
												},
											},

											"policy": schema.StringAttribute{
												Description:         "Defines the upgrade policy for the configuration. This field is optional.",
												MarkdownDescription: "Defines the upgrade policy for the configuration. This field is optional.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("simple", "parallel", "rolling", "autoReload", "operatorSyncUpdate", "dynamicReloadBeginRestart"),
												},
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

					"restart": schema.ListNestedAttribute{
						Description:         "Restarts the specified components.",
						MarkdownDescription: "Restarts the specified components.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the cluster component.",
									MarkdownDescription: "Specifies the name of the cluster component.",
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

					"restore_from": schema.SingleNestedAttribute{
						Description:         "Cluster RestoreFrom backup or point in time.",
						MarkdownDescription: "Cluster RestoreFrom backup or point in time.",
						Attributes: map[string]schema.Attribute{
							"backup": schema.ListNestedAttribute{
								Description:         "Refers to the backup name and component name used for restoration. Supports recovery of multiple components.",
								MarkdownDescription: "Refers to the backup name and component name used for restoration. Supports recovery of multiple components.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"ref": schema.SingleNestedAttribute{
											Description:         "Refers to a reference backup that needs to be restored.",
											MarkdownDescription: "Refers to a reference backup that needs to be restored.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Refers to the specific name of the resource.",
													MarkdownDescription: "Refers to the specific name of the resource.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "Refers to the specific namespace of the resource.",
													MarkdownDescription: "Refers to the specific namespace of the resource.",
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

							"point_in_time": schema.SingleNestedAttribute{
								Description:         "Refers to the specific point in time for recovery.",
								MarkdownDescription: "Refers to the specific point in time for recovery.",
								Attributes: map[string]schema.Attribute{
									"ref": schema.SingleNestedAttribute{
										Description:         "Refers to a reference source cluster that needs to be restored.",
										MarkdownDescription: "Refers to a reference source cluster that needs to be restored.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Refers to the specific name of the resource.",
												MarkdownDescription: "Refers to the specific name of the resource.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Refers to the specific namespace of the resource.",
												MarkdownDescription: "Refers to the specific namespace of the resource.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"time": schema.StringAttribute{
										Description:         "Refers to the specific time point for restoration, with UTC as the time zone.",
										MarkdownDescription: "Refers to the specific time point for restoration, with UTC as the time zone.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											validators.DateTime64Validator(),
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

					"restore_spec": schema.SingleNestedAttribute{
						Description:         "Defines how to restore the cluster. Note that this restore operation will roll back cluster services.",
						MarkdownDescription: "Defines how to restore the cluster. Note that this restore operation will roll back cluster services.",
						Attributes: map[string]schema.Attribute{
							"backup_name": schema.StringAttribute{
								Description:         "Specifies the name of the backup.",
								MarkdownDescription: "Specifies the name of the backup.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"effective_common_component_def": schema.BoolAttribute{
								Description:         "Indicates if this backup will be restored for all components which refer to common ComponentDefinition.",
								MarkdownDescription: "Indicates if this backup will be restored for all components which refer to common ComponentDefinition.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"restore_time_str": schema.StringAttribute{
								Description:         "Defines the point in time to restore.",
								MarkdownDescription: "Defines the point in time to restore.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"volume_restore_policy": schema.StringAttribute{
								Description:         "Specifies the volume claim restore policy, support values: [Serial, Parallel]",
								MarkdownDescription: "Specifies the volume claim restore policy, support values: [Serial, Parallel]",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Serial", "Parallel"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"script_spec": schema.SingleNestedAttribute{
						Description:         "Defines the script to be executed.",
						MarkdownDescription: "Defines the script to be executed.",
						Attributes: map[string]schema.Attribute{
							"component_name": schema.StringAttribute{
								Description:         "Specifies the name of the cluster component.",
								MarkdownDescription: "Specifies the name of the cluster component.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"image": schema.StringAttribute{
								Description:         "Specifies the image to be used for the exec command. By default, the image of kubeblocks-datascript is used.",
								MarkdownDescription: "Specifies the image to be used for the exec command. By default, the image of kubeblocks-datascript is used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"script": schema.ListAttribute{
								Description:         "Defines the script to be executed.",
								MarkdownDescription: "Defines the script to be executed.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"script_from": schema.SingleNestedAttribute{
								Description:         "Defines the script to be executed from a configMap or secret.",
								MarkdownDescription: "Defines the script to be executed from a configMap or secret.",
								Attributes: map[string]schema.Attribute{
									"config_map_ref": schema.ListNestedAttribute{
										Description:         "Specifies the configMap that is to be executed.",
										MarkdownDescription: "Specifies the configMap that is to be executed.",
										NestedObject: schema.NestedAttributeObject{
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_ref": schema.ListNestedAttribute{
										Description:         "Specifies the secret that is to be executed.",
										MarkdownDescription: "Specifies the secret that is to be executed.",
										NestedObject: schema.NestedAttributeObject{
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

							"secret": schema.SingleNestedAttribute{
								Description:         "Defines the secret to be used to execute the script. If not specified, the default cluster root credential secret is used.",
								MarkdownDescription: "Defines the secret to be used to execute the script. If not specified, the default cluster root credential secret is used.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Specifies the name of the secret.",
										MarkdownDescription: "Specifies the name of the secret.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtMost(63),
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
										},
									},

									"password_key": schema.StringAttribute{
										Description:         "Used to specify the password part of the secret.",
										MarkdownDescription: "Used to specify the password part of the secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"username_key": schema.StringAttribute{
										Description:         "Used to specify the username part of the secret.",
										MarkdownDescription: "Used to specify the username part of the secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"selector": schema.SingleNestedAttribute{
								Description:         "By default, KubeBlocks will execute the script on the primary pod with role=leader. Exceptions exist, such as Redis, which does not synchronize account information between primary and secondary. In such cases, the script needs to be executed on all pods matching the selector. Indicates the components on which the script is executed.",
								MarkdownDescription: "By default, KubeBlocks will execute the script on the primary pod with role=leader. Exceptions exist, such as Redis, which does not synchronize account information between primary and secondary. In such cases, the script needs to be executed on all pods matching the selector. Indicates the components on which the script is executed.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"switchover": schema.ListNestedAttribute{
						Description:         "Switches over the specified components.",
						MarkdownDescription: "Switches over the specified components.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the cluster component.",
									MarkdownDescription: "Specifies the name of the cluster component.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"instance_name": schema.StringAttribute{
									Description:         "Utilized to designate the candidate primary or leader instance for the switchover process. If assigned '*', it signifies that no specific primary or leader is designated for the switchover, and the switchoverAction defined in 'clusterDefinition.componentDefs[x].switchoverSpec.withoutCandidate' will be executed.  It is mandatory that 'clusterDefinition.componentDefs[x].switchoverSpec.withoutCandidate' is not left blank.  If assigned a valid instance name other than '*', it signifies that a specific candidate primary or leader is designated for the switchover. The value can be retrieved using 'kbcli cluster list-instances', any other value is considered invalid.  In this scenario, the 'switchoverAction' defined in clusterDefinition.componentDefs[x].switchoverSpec.withCandidate will be executed, and it is mandatory that clusterDefinition.componentDefs[x].switchoverSpec.withCandidate is not left blank.",
									MarkdownDescription: "Utilized to designate the candidate primary or leader instance for the switchover process. If assigned '*', it signifies that no specific primary or leader is designated for the switchover, and the switchoverAction defined in 'clusterDefinition.componentDefs[x].switchoverSpec.withoutCandidate' will be executed.  It is mandatory that 'clusterDefinition.componentDefs[x].switchoverSpec.withoutCandidate' is not left blank.  If assigned a valid instance name other than '*', it signifies that a specific candidate primary or leader is designated for the switchover. The value can be retrieved using 'kbcli cluster list-instances', any other value is considered invalid.  In this scenario, the 'switchoverAction' defined in clusterDefinition.componentDefs[x].switchoverSpec.withCandidate will be executed, and it is mandatory that clusterDefinition.componentDefs[x].switchoverSpec.withCandidate is not left blank.",
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

					"ttl_seconds_after_succeed": schema.Int64Attribute{
						Description:         "OpsRequest will be deleted after TTLSecondsAfterSucceed second when OpsRequest.status.phase is Succeed.",
						MarkdownDescription: "OpsRequest will be deleted after TTLSecondsAfterSucceed second when OpsRequest.status.phase is Succeed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ttl_seconds_before_abort": schema.Int64Attribute{
						Description:         "OpsRequest will wait at most TTLSecondsBeforeAbort seconds for start-conditions to be met. If not specified, the default value is 0, which means that the start-conditions must be met immediately.",
						MarkdownDescription: "OpsRequest will wait at most TTLSecondsBeforeAbort seconds for start-conditions to be met. If not specified, the default value is 0, which means that the start-conditions must be met immediately.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"type": schema.StringAttribute{
						Description:         "Defines the operation type.",
						MarkdownDescription: "Defines the operation type.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Upgrade", "VerticalScaling", "VolumeExpansion", "HorizontalScaling", "Restart", "Reconfiguring", "Start", "Stop", "Expose", "Switchover", "DataScript", "Backup", "Restore", "Custom"),
						},
					},

					"upgrade": schema.SingleNestedAttribute{
						Description:         "Specifies the cluster version by specifying clusterVersionRef.",
						MarkdownDescription: "Specifies the cluster version by specifying clusterVersionRef.",
						Attributes: map[string]schema.Attribute{
							"cluster_version_ref": schema.StringAttribute{
								Description:         "A reference to the name of the ClusterVersion.",
								MarkdownDescription: "A reference to the name of the ClusterVersion.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"vertical_scaling": schema.ListAttribute{
						Description:         "Note: Quantity struct can not do immutable check by CEL. Defines what component need to vertical scale the specified compute resources.",
						MarkdownDescription: "Note: Quantity struct can not do immutable check by CEL. Defines what component need to vertical scale the specified compute resources.",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"volume_expansion": schema.ListNestedAttribute{
						Description:         "Note: Quantity struct can not do immutable check by CEL. Defines what component and volumeClaimTemplate need to expand the specified storage.",
						MarkdownDescription: "Note: Quantity struct can not do immutable check by CEL. Defines what component and volumeClaimTemplate need to expand the specified storage.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the cluster component.",
									MarkdownDescription: "Specifies the name of the cluster component.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"volume_claim_templates": schema.ListNestedAttribute{
									Description:         "volumeClaimTemplates specifies the storage size and volumeClaimTemplate name.",
									MarkdownDescription: "volumeClaimTemplates specifies the storage size and volumeClaimTemplate name.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "A reference to the volumeClaimTemplate name from the cluster components.",
												MarkdownDescription: "A reference to the volumeClaimTemplate name from the cluster components.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"storage": schema.StringAttribute{
												Description:         "Specifies the requested storage size for the volume.",
												MarkdownDescription: "Specifies the requested storage size for the volume.",
												Required:            true,
												Optional:            false,
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *AppsKubeblocksIoOpsRequestV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_kubeblocks_io_ops_request_v1alpha1_manifest")

	var model AppsKubeblocksIoOpsRequestV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("apps.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("OpsRequest")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
