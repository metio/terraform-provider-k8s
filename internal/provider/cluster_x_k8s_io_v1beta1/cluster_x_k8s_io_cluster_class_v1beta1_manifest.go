/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cluster_x_k8s_io_v1beta1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &ClusterXK8SIoClusterClassV1Beta1Manifest{}
)

func NewClusterXK8SIoClusterClassV1Beta1Manifest() datasource.DataSource {
	return &ClusterXK8SIoClusterClassV1Beta1Manifest{}
}

type ClusterXK8SIoClusterClassV1Beta1Manifest struct{}

type ClusterXK8SIoClusterClassV1Beta1ManifestData struct {
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
		ControlPlane *struct {
			MachineHealthCheck *struct {
				MaxUnhealthy        *string `tfsdk:"max_unhealthy" json:"maxUnhealthy,omitempty"`
				NodeStartupTimeout  *string `tfsdk:"node_startup_timeout" json:"nodeStartupTimeout,omitempty"`
				RemediationTemplate *struct {
					ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
					Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
					Name            *string `tfsdk:"name" json:"name,omitempty"`
					Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
					ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
					Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
				} `tfsdk:"remediation_template" json:"remediationTemplate,omitempty"`
				UnhealthyConditions *[]struct {
					Status  *string `tfsdk:"status" json:"status,omitempty"`
					Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
					Type    *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"unhealthy_conditions" json:"unhealthyConditions,omitempty"`
				UnhealthyRange *string `tfsdk:"unhealthy_range" json:"unhealthyRange,omitempty"`
			} `tfsdk:"machine_health_check" json:"machineHealthCheck,omitempty"`
			MachineInfrastructure *struct {
				Ref *struct {
					ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
					Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
					Name            *string `tfsdk:"name" json:"name,omitempty"`
					Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
					ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
					Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
				} `tfsdk:"ref" json:"ref,omitempty"`
			} `tfsdk:"machine_infrastructure" json:"machineInfrastructure,omitempty"`
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			NamingStrategy *struct {
				Template *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"naming_strategy" json:"namingStrategy,omitempty"`
			NodeDeletionTimeout     *string `tfsdk:"node_deletion_timeout" json:"nodeDeletionTimeout,omitempty"`
			NodeDrainTimeout        *string `tfsdk:"node_drain_timeout" json:"nodeDrainTimeout,omitempty"`
			NodeVolumeDetachTimeout *string `tfsdk:"node_volume_detach_timeout" json:"nodeVolumeDetachTimeout,omitempty"`
			Ref                     *struct {
				ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
				Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
				Name            *string `tfsdk:"name" json:"name,omitempty"`
				Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
				ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
				Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
			} `tfsdk:"ref" json:"ref,omitempty"`
		} `tfsdk:"control_plane" json:"controlPlane,omitempty"`
		Infrastructure *struct {
			Ref *struct {
				ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
				Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
				Name            *string `tfsdk:"name" json:"name,omitempty"`
				Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
				ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
				Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
			} `tfsdk:"ref" json:"ref,omitempty"`
		} `tfsdk:"infrastructure" json:"infrastructure,omitempty"`
		Patches *[]struct {
			Definitions *[]struct {
				JsonPatches *[]struct {
					Op        *string            `tfsdk:"op" json:"op,omitempty"`
					Path      *string            `tfsdk:"path" json:"path,omitempty"`
					Value     *map[string]string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						Template *string `tfsdk:"template" json:"template,omitempty"`
						Variable *string `tfsdk:"variable" json:"variable,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"json_patches" json:"jsonPatches,omitempty"`
				Selector *struct {
					ApiVersion     *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					Kind           *string `tfsdk:"kind" json:"kind,omitempty"`
					MatchResources *struct {
						ControlPlane           *bool `tfsdk:"control_plane" json:"controlPlane,omitempty"`
						InfrastructureCluster  *bool `tfsdk:"infrastructure_cluster" json:"infrastructureCluster,omitempty"`
						MachineDeploymentClass *struct {
							Names *[]string `tfsdk:"names" json:"names,omitempty"`
						} `tfsdk:"machine_deployment_class" json:"machineDeploymentClass,omitempty"`
						MachinePoolClass *struct {
							Names *[]string `tfsdk:"names" json:"names,omitempty"`
						} `tfsdk:"machine_pool_class" json:"machinePoolClass,omitempty"`
					} `tfsdk:"match_resources" json:"matchResources,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
			} `tfsdk:"definitions" json:"definitions,omitempty"`
			Description *string `tfsdk:"description" json:"description,omitempty"`
			EnabledIf   *string `tfsdk:"enabled_if" json:"enabledIf,omitempty"`
			External    *struct {
				DiscoverVariablesExtension *string            `tfsdk:"discover_variables_extension" json:"discoverVariablesExtension,omitempty"`
				GenerateExtension          *string            `tfsdk:"generate_extension" json:"generateExtension,omitempty"`
				Settings                   *map[string]string `tfsdk:"settings" json:"settings,omitempty"`
				ValidateExtension          *string            `tfsdk:"validate_extension" json:"validateExtension,omitempty"`
			} `tfsdk:"external" json:"external,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"patches" json:"patches,omitempty"`
		Variables *[]struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Required *bool   `tfsdk:"required" json:"required,omitempty"`
			Schema   *struct {
				OpenAPIV3Schema *struct {
					AdditionalProperties                 *map[string]string `tfsdk:"additional_properties" json:"additionalProperties,omitempty"`
					Default                              *map[string]string `tfsdk:"default" json:"default,omitempty"`
					Description                          *string            `tfsdk:"description" json:"description,omitempty"`
					Enum                                 *[]string          `tfsdk:"enum" json:"enum,omitempty"`
					Example                              *map[string]string `tfsdk:"example" json:"example,omitempty"`
					ExclusiveMaximum                     *bool              `tfsdk:"exclusive_maximum" json:"exclusiveMaximum,omitempty"`
					ExclusiveMinimum                     *bool              `tfsdk:"exclusive_minimum" json:"exclusiveMinimum,omitempty"`
					Format                               *string            `tfsdk:"format" json:"format,omitempty"`
					Items                                *map[string]string `tfsdk:"items" json:"items,omitempty"`
					MaxItems                             *int64             `tfsdk:"max_items" json:"maxItems,omitempty"`
					MaxLength                            *int64             `tfsdk:"max_length" json:"maxLength,omitempty"`
					Maximum                              *int64             `tfsdk:"maximum" json:"maximum,omitempty"`
					MinItems                             *int64             `tfsdk:"min_items" json:"minItems,omitempty"`
					MinLength                            *int64             `tfsdk:"min_length" json:"minLength,omitempty"`
					Minimum                              *int64             `tfsdk:"minimum" json:"minimum,omitempty"`
					Pattern                              *string            `tfsdk:"pattern" json:"pattern,omitempty"`
					Properties                           *map[string]string `tfsdk:"properties" json:"properties,omitempty"`
					Required                             *[]string          `tfsdk:"required" json:"required,omitempty"`
					Type                                 *string            `tfsdk:"type" json:"type,omitempty"`
					UniqueItems                          *bool              `tfsdk:"unique_items" json:"uniqueItems,omitempty"`
					X_kubernetes_preserve_unknown_fields *bool              `tfsdk:"x_kubernetes_preserve_unknown_fields" json:"x-kubernetes-preserve-unknown-fields,omitempty"`
				} `tfsdk:"open_apiv3_schema" json:"openAPIV3Schema,omitempty"`
			} `tfsdk:"schema" json:"schema,omitempty"`
		} `tfsdk:"variables" json:"variables,omitempty"`
		Workers *struct {
			MachineDeployments *[]struct {
				Class              *string `tfsdk:"class" json:"class,omitempty"`
				FailureDomain      *string `tfsdk:"failure_domain" json:"failureDomain,omitempty"`
				MachineHealthCheck *struct {
					MaxUnhealthy        *string `tfsdk:"max_unhealthy" json:"maxUnhealthy,omitempty"`
					NodeStartupTimeout  *string `tfsdk:"node_startup_timeout" json:"nodeStartupTimeout,omitempty"`
					RemediationTemplate *struct {
						ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
						FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
						Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
						Name            *string `tfsdk:"name" json:"name,omitempty"`
						Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
						ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
						Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
					} `tfsdk:"remediation_template" json:"remediationTemplate,omitempty"`
					UnhealthyConditions *[]struct {
						Status  *string `tfsdk:"status" json:"status,omitempty"`
						Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
						Type    *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"unhealthy_conditions" json:"unhealthyConditions,omitempty"`
					UnhealthyRange *string `tfsdk:"unhealthy_range" json:"unhealthyRange,omitempty"`
				} `tfsdk:"machine_health_check" json:"machineHealthCheck,omitempty"`
				MinReadySeconds *int64 `tfsdk:"min_ready_seconds" json:"minReadySeconds,omitempty"`
				NamingStrategy  *struct {
					Template *string `tfsdk:"template" json:"template,omitempty"`
				} `tfsdk:"naming_strategy" json:"namingStrategy,omitempty"`
				NodeDeletionTimeout     *string `tfsdk:"node_deletion_timeout" json:"nodeDeletionTimeout,omitempty"`
				NodeDrainTimeout        *string `tfsdk:"node_drain_timeout" json:"nodeDrainTimeout,omitempty"`
				NodeVolumeDetachTimeout *string `tfsdk:"node_volume_detach_timeout" json:"nodeVolumeDetachTimeout,omitempty"`
				Strategy                *struct {
					RollingUpdate *struct {
						DeletePolicy   *string `tfsdk:"delete_policy" json:"deletePolicy,omitempty"`
						MaxSurge       *string `tfsdk:"max_surge" json:"maxSurge,omitempty"`
						MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
					} `tfsdk:"rolling_update" json:"rollingUpdate,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"strategy" json:"strategy,omitempty"`
				Template *struct {
					Bootstrap *struct {
						Ref *struct {
							ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
							FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
							Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
							ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
							Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
						} `tfsdk:"ref" json:"ref,omitempty"`
					} `tfsdk:"bootstrap" json:"bootstrap,omitempty"`
					Infrastructure *struct {
						Ref *struct {
							ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
							FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
							Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
							ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
							Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
						} `tfsdk:"ref" json:"ref,omitempty"`
					} `tfsdk:"infrastructure" json:"infrastructure,omitempty"`
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"machine_deployments" json:"machineDeployments,omitempty"`
			MachinePools *[]struct {
				Class           *string   `tfsdk:"class" json:"class,omitempty"`
				FailureDomains  *[]string `tfsdk:"failure_domains" json:"failureDomains,omitempty"`
				MinReadySeconds *int64    `tfsdk:"min_ready_seconds" json:"minReadySeconds,omitempty"`
				NamingStrategy  *struct {
					Template *string `tfsdk:"template" json:"template,omitempty"`
				} `tfsdk:"naming_strategy" json:"namingStrategy,omitempty"`
				NodeDeletionTimeout     *string `tfsdk:"node_deletion_timeout" json:"nodeDeletionTimeout,omitempty"`
				NodeDrainTimeout        *string `tfsdk:"node_drain_timeout" json:"nodeDrainTimeout,omitempty"`
				NodeVolumeDetachTimeout *string `tfsdk:"node_volume_detach_timeout" json:"nodeVolumeDetachTimeout,omitempty"`
				Template                *struct {
					Bootstrap *struct {
						Ref *struct {
							ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
							FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
							Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
							ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
							Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
						} `tfsdk:"ref" json:"ref,omitempty"`
					} `tfsdk:"bootstrap" json:"bootstrap,omitempty"`
					Infrastructure *struct {
						Ref *struct {
							ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
							FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
							Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
							ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
							Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
						} `tfsdk:"ref" json:"ref,omitempty"`
					} `tfsdk:"infrastructure" json:"infrastructure,omitempty"`
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"machine_pools" json:"machinePools,omitempty"`
		} `tfsdk:"workers" json:"workers,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ClusterXK8SIoClusterClassV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cluster_x_k8s_io_cluster_class_v1beta1_manifest"
}

func (r *ClusterXK8SIoClusterClassV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterClass is a template which can be used to create managed topologies.",
		MarkdownDescription: "ClusterClass is a template which can be used to create managed topologies.",
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
				Description:         "ClusterClassSpec describes the desired state of the ClusterClass.",
				MarkdownDescription: "ClusterClassSpec describes the desired state of the ClusterClass.",
				Attributes: map[string]schema.Attribute{
					"control_plane": schema.SingleNestedAttribute{
						Description:         "ControlPlane is a reference to a local struct that holds the detailsfor provisioning the Control Plane for the Cluster.",
						MarkdownDescription: "ControlPlane is a reference to a local struct that holds the detailsfor provisioning the Control Plane for the Cluster.",
						Attributes: map[string]schema.Attribute{
							"machine_health_check": schema.SingleNestedAttribute{
								Description:         "MachineHealthCheck defines a MachineHealthCheck for this ControlPlaneClass.This field is supported if and only if the ControlPlane provider templatereferenced above is Machine based and supports setting replicas.",
								MarkdownDescription: "MachineHealthCheck defines a MachineHealthCheck for this ControlPlaneClass.This field is supported if and only if the ControlPlane provider templatereferenced above is Machine based and supports setting replicas.",
								Attributes: map[string]schema.Attribute{
									"max_unhealthy": schema.StringAttribute{
										Description:         "Any further remediation is only allowed if at most 'MaxUnhealthy' machines selected by'selector' are not healthy.",
										MarkdownDescription: "Any further remediation is only allowed if at most 'MaxUnhealthy' machines selected by'selector' are not healthy.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"node_startup_timeout": schema.StringAttribute{
										Description:         "NodeStartupTimeout allows to set the maximum time for MachineHealthCheckto consider a Machine unhealthy if a corresponding Node isn't associatedthrough a 'Spec.ProviderID' field.The duration set in this field is compared to the greatest of:- Cluster's infrastructure and control plane ready condition timestamp (if and when available)- Machine's infrastructure ready condition timestamp (if and when available)- Machine's metadata creation timestampDefaults to 10 minutes.If you wish to disable this feature, set the value explicitly to 0.",
										MarkdownDescription: "NodeStartupTimeout allows to set the maximum time for MachineHealthCheckto consider a Machine unhealthy if a corresponding Node isn't associatedthrough a 'Spec.ProviderID' field.The duration set in this field is compared to the greatest of:- Cluster's infrastructure and control plane ready condition timestamp (if and when available)- Machine's infrastructure ready condition timestamp (if and when available)- Machine's metadata creation timestampDefaults to 10 minutes.If you wish to disable this feature, set the value explicitly to 0.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"remediation_template": schema.SingleNestedAttribute{
										Description:         "RemediationTemplate is a reference to a remediation templateprovided by an infrastructure provider.This field is completely optional, when filled, the MachineHealthCheck controllercreates a new object from the template referenced and hands off remediation of the machine toa controller that lives outside of Cluster API.",
										MarkdownDescription: "RemediationTemplate is a reference to a remediation templateprovided by an infrastructure provider.This field is completely optional, when filled, the MachineHealthCheck controllercreates a new object from the template referenced and hands off remediation of the machine toa controller that lives outside of Cluster API.",
										Attributes: map[string]schema.Attribute{
											"api_version": schema.StringAttribute{
												Description:         "API version of the referent.",
												MarkdownDescription: "API version of the referent.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"field_path": schema.StringAttribute{
												Description:         "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
												MarkdownDescription: "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"kind": schema.StringAttribute{
												Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
												MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
												MarkdownDescription: "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"resource_version": schema.StringAttribute{
												Description:         "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
												MarkdownDescription: "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"uid": schema.StringAttribute{
												Description:         "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
												MarkdownDescription: "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"unhealthy_conditions": schema.ListNestedAttribute{
										Description:         "UnhealthyConditions contains a list of the conditions that determinewhether a node is considered unhealthy. The conditions are combined in alogical OR, i.e. if any of the conditions is met, the node is unhealthy.",
										MarkdownDescription: "UnhealthyConditions contains a list of the conditions that determinewhether a node is considered unhealthy. The conditions are combined in alogical OR, i.e. if any of the conditions is met, the node is unhealthy.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"status": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
													},
												},

												"timeout": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"unhealthy_range": schema.StringAttribute{
										Description:         "Any further remediation is only allowed if the number of machines selected by 'selector' as not healthyis within the range of 'UnhealthyRange'. Takes precedence over MaxUnhealthy.Eg. '[3-5]' - This means that remediation will be allowed only when:(a) there are at least 3 unhealthy machines (and)(b) there are at most 5 unhealthy machines",
										MarkdownDescription: "Any further remediation is only allowed if the number of machines selected by 'selector' as not healthyis within the range of 'UnhealthyRange'. Takes precedence over MaxUnhealthy.Eg. '[3-5]' - This means that remediation will be allowed only when:(a) there are at least 3 unhealthy machines (and)(b) there are at most 5 unhealthy machines",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^\[[0-9]+-[0-9]+\]$`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"machine_infrastructure": schema.SingleNestedAttribute{
								Description:         "MachineInfrastructure defines the metadata and infrastructure informationfor control plane machines.This field is supported if and only if the control plane provider templatereferenced above is Machine based and supports setting replicas.",
								MarkdownDescription: "MachineInfrastructure defines the metadata and infrastructure informationfor control plane machines.This field is supported if and only if the control plane provider templatereferenced above is Machine based and supports setting replicas.",
								Attributes: map[string]schema.Attribute{
									"ref": schema.SingleNestedAttribute{
										Description:         "Ref is a required reference to a custom resourceoffered by a provider.",
										MarkdownDescription: "Ref is a required reference to a custom resourceoffered by a provider.",
										Attributes: map[string]schema.Attribute{
											"api_version": schema.StringAttribute{
												Description:         "API version of the referent.",
												MarkdownDescription: "API version of the referent.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"field_path": schema.StringAttribute{
												Description:         "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
												MarkdownDescription: "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"kind": schema.StringAttribute{
												Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
												MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
												MarkdownDescription: "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"resource_version": schema.StringAttribute{
												Description:         "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
												MarkdownDescription: "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"uid": schema.StringAttribute{
												Description:         "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
												MarkdownDescription: "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"metadata": schema.SingleNestedAttribute{
								Description:         "Metadata is the metadata applied to the ControlPlane and the Machines of the ControlPlaneif the ControlPlaneTemplate referenced is machine based. If not, it is applied only to theControlPlane.At runtime this metadata is merged with the corresponding metadata from the topology.This field is supported if and only if the control plane provider templatereferenced is Machine based.",
								MarkdownDescription: "Metadata is the metadata applied to the ControlPlane and the Machines of the ControlPlaneif the ControlPlaneTemplate referenced is machine based. If not, it is applied only to theControlPlane.At runtime this metadata is merged with the corresponding metadata from the topology.This field is supported if and only if the control plane provider templatereferenced is Machine based.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: http://kubernetes.io/docs/user-guide/annotations",
										MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: http://kubernetes.io/docs/user-guide/annotations",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: http://kubernetes.io/docs/user-guide/labels",
										MarkdownDescription: "Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: http://kubernetes.io/docs/user-guide/labels",
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

							"naming_strategy": schema.SingleNestedAttribute{
								Description:         "NamingStrategy allows changing the naming pattern used when creating the control plane provider object.",
								MarkdownDescription: "NamingStrategy allows changing the naming pattern used when creating the control plane provider object.",
								Attributes: map[string]schema.Attribute{
									"template": schema.StringAttribute{
										Description:         "Template defines the template to use for generating the name of the ControlPlane object.If not defined, it will fallback to '{{ .cluster.name }}-{{ .random }}'.If the templated string exceeds 63 characters, it will be trimmed to 58 characters and willget concatenated with a random suffix of length 5.The templating mechanism provides the following arguments:* '.cluster.name': The name of the cluster object.* '.random': A random alphanumeric string, without vowels, of length 5.",
										MarkdownDescription: "Template defines the template to use for generating the name of the ControlPlane object.If not defined, it will fallback to '{{ .cluster.name }}-{{ .random }}'.If the templated string exceeds 63 characters, it will be trimmed to 58 characters and willget concatenated with a random suffix of length 5.The templating mechanism provides the following arguments:* '.cluster.name': The name of the cluster object.* '.random': A random alphanumeric string, without vowels, of length 5.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"node_deletion_timeout": schema.StringAttribute{
								Description:         "NodeDeletionTimeout defines how long the controller will attempt to delete the Node that the Machinehosts after the Machine is marked for deletion. A duration of 0 will retry deletion indefinitely.Defaults to 10 seconds.NOTE: This value can be overridden while defining a Cluster.Topology.",
								MarkdownDescription: "NodeDeletionTimeout defines how long the controller will attempt to delete the Node that the Machinehosts after the Machine is marked for deletion. A duration of 0 will retry deletion indefinitely.Defaults to 10 seconds.NOTE: This value can be overridden while defining a Cluster.Topology.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_drain_timeout": schema.StringAttribute{
								Description:         "NodeDrainTimeout is the total amount of time that the controller will spend on draining a node.The default value is 0, meaning that the node can be drained without any time limitations.NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'NOTE: This value can be overridden while defining a Cluster.Topology.",
								MarkdownDescription: "NodeDrainTimeout is the total amount of time that the controller will spend on draining a node.The default value is 0, meaning that the node can be drained without any time limitations.NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'NOTE: This value can be overridden while defining a Cluster.Topology.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_volume_detach_timeout": schema.StringAttribute{
								Description:         "NodeVolumeDetachTimeout is the total amount of time that the controller will spend on waiting for all volumesto be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.NOTE: This value can be overridden while defining a Cluster.Topology.",
								MarkdownDescription: "NodeVolumeDetachTimeout is the total amount of time that the controller will spend on waiting for all volumesto be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.NOTE: This value can be overridden while defining a Cluster.Topology.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ref": schema.SingleNestedAttribute{
								Description:         "Ref is a required reference to a custom resourceoffered by a provider.",
								MarkdownDescription: "Ref is a required reference to a custom resourceoffered by a provider.",
								Attributes: map[string]schema.Attribute{
									"api_version": schema.StringAttribute{
										Description:         "API version of the referent.",
										MarkdownDescription: "API version of the referent.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"field_path": schema.StringAttribute{
										Description:         "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
										MarkdownDescription: "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"kind": schema.StringAttribute{
										Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
										MarkdownDescription: "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resource_version": schema.StringAttribute{
										Description:         "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
										MarkdownDescription: "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"uid": schema.StringAttribute{
										Description:         "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
										MarkdownDescription: "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"infrastructure": schema.SingleNestedAttribute{
						Description:         "Infrastructure is a reference to a provider-specific template that holdsthe details for provisioning infrastructure specific clusterfor the underlying provider.The underlying provider is responsible for the implementationof the template to an infrastructure cluster.",
						MarkdownDescription: "Infrastructure is a reference to a provider-specific template that holdsthe details for provisioning infrastructure specific clusterfor the underlying provider.The underlying provider is responsible for the implementationof the template to an infrastructure cluster.",
						Attributes: map[string]schema.Attribute{
							"ref": schema.SingleNestedAttribute{
								Description:         "Ref is a required reference to a custom resourceoffered by a provider.",
								MarkdownDescription: "Ref is a required reference to a custom resourceoffered by a provider.",
								Attributes: map[string]schema.Attribute{
									"api_version": schema.StringAttribute{
										Description:         "API version of the referent.",
										MarkdownDescription: "API version of the referent.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"field_path": schema.StringAttribute{
										Description:         "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
										MarkdownDescription: "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"kind": schema.StringAttribute{
										Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
										MarkdownDescription: "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resource_version": schema.StringAttribute{
										Description:         "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
										MarkdownDescription: "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"uid": schema.StringAttribute{
										Description:         "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
										MarkdownDescription: "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"patches": schema.ListNestedAttribute{
						Description:         "Patches defines the patches which are applied to customizereferenced templates of a ClusterClass.Note: Patches will be applied in the order of the array.",
						MarkdownDescription: "Patches defines the patches which are applied to customizereferenced templates of a ClusterClass.Note: Patches will be applied in the order of the array.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"definitions": schema.ListNestedAttribute{
									Description:         "Definitions define inline patches.Note: Patches will be applied in the order of the array.Note: Exactly one of Definitions or External must be set.",
									MarkdownDescription: "Definitions define inline patches.Note: Patches will be applied in the order of the array.Note: Exactly one of Definitions or External must be set.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"json_patches": schema.ListNestedAttribute{
												Description:         "JSONPatches defines the patches which should be applied on the templatesmatching the selector.Note: Patches will be applied in the order of the array.",
												MarkdownDescription: "JSONPatches defines the patches which should be applied on the templatesmatching the selector.Note: Patches will be applied in the order of the array.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"op": schema.StringAttribute{
															Description:         "Op defines the operation of the patch.Note: Only 'add', 'replace' and 'remove' are supported.",
															MarkdownDescription: "Op defines the operation of the patch.Note: Only 'add', 'replace' and 'remove' are supported.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"path": schema.StringAttribute{
															Description:         "Path defines the path of the patch.Note: Only the spec of a template can be patched, thus the path has to start with /spec/.Note: For now the only allowed array modifications are 'append' and 'prepend', i.e.:* for op: 'add': only index 0 (prepend) and - (append) are allowed* for op: 'replace' or 'remove': no indexes are allowed",
															MarkdownDescription: "Path defines the path of the patch.Note: Only the spec of a template can be patched, thus the path has to start with /spec/.Note: For now the only allowed array modifications are 'append' and 'prepend', i.e.:* for op: 'add': only index 0 (prepend) and - (append) are allowed* for op: 'replace' or 'remove': no indexes are allowed",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"value": schema.MapAttribute{
															Description:         "Value defines the value of the patch.Note: Either Value or ValueFrom is required for add and replaceoperations. Only one of them is allowed to be set at the same time.Note: We have to use apiextensionsv1.JSON instead of our JSON type,because controller-tools has a hard-coded schema for apiextensionsv1.JSONwhich cannot be produced by another type (unset type field).Ref: https://github.com/kubernetes-sigs/controller-tools/blob/d0e03a142d0ecdd5491593e941ee1d6b5d91dba6/pkg/crd/known_types.go#L106-L111",
															MarkdownDescription: "Value defines the value of the patch.Note: Either Value or ValueFrom is required for add and replaceoperations. Only one of them is allowed to be set at the same time.Note: We have to use apiextensionsv1.JSON instead of our JSON type,because controller-tools has a hard-coded schema for apiextensionsv1.JSONwhich cannot be produced by another type (unset type field).Ref: https://github.com/kubernetes-sigs/controller-tools/blob/d0e03a142d0ecdd5491593e941ee1d6b5d91dba6/pkg/crd/known_types.go#L106-L111",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"value_from": schema.SingleNestedAttribute{
															Description:         "ValueFrom defines the value of the patch.Note: Either Value or ValueFrom is required for add and replaceoperations. Only one of them is allowed to be set at the same time.",
															MarkdownDescription: "ValueFrom defines the value of the patch.Note: Either Value or ValueFrom is required for add and replaceoperations. Only one of them is allowed to be set at the same time.",
															Attributes: map[string]schema.Attribute{
																"template": schema.StringAttribute{
																	Description:         "Template is the Go template to be used to calculate the value.A template can reference variables defined in .spec.variables and builtin variables.Note: The template must evaluate to a valid YAML or JSON value.",
																	MarkdownDescription: "Template is the Go template to be used to calculate the value.A template can reference variables defined in .spec.variables and builtin variables.Note: The template must evaluate to a valid YAML or JSON value.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"variable": schema.StringAttribute{
																	Description:         "Variable is the variable to be used as value.Variable can be one of the variables defined in .spec.variables or a builtin variable.",
																	MarkdownDescription: "Variable is the variable to be used as value.Variable can be one of the variables defined in .spec.variables or a builtin variable.",
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
												Required: true,
												Optional: false,
												Computed: false,
											},

											"selector": schema.SingleNestedAttribute{
												Description:         "Selector defines on which templates the patch should be applied.",
												MarkdownDescription: "Selector defines on which templates the patch should be applied.",
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "APIVersion filters templates by apiVersion.",
														MarkdownDescription: "APIVersion filters templates by apiVersion.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind filters templates by kind.",
														MarkdownDescription: "Kind filters templates by kind.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"match_resources": schema.SingleNestedAttribute{
														Description:         "MatchResources selects templates based on where they are referenced.",
														MarkdownDescription: "MatchResources selects templates based on where they are referenced.",
														Attributes: map[string]schema.Attribute{
															"control_plane": schema.BoolAttribute{
																Description:         "ControlPlane selects templates referenced in .spec.ControlPlane.Note: this will match the controlPlane and also the controlPlanemachineInfrastructure (depending on the kind and apiVersion).",
																MarkdownDescription: "ControlPlane selects templates referenced in .spec.ControlPlane.Note: this will match the controlPlane and also the controlPlanemachineInfrastructure (depending on the kind and apiVersion).",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"infrastructure_cluster": schema.BoolAttribute{
																Description:         "InfrastructureCluster selects templates referenced in .spec.infrastructure.",
																MarkdownDescription: "InfrastructureCluster selects templates referenced in .spec.infrastructure.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"machine_deployment_class": schema.SingleNestedAttribute{
																Description:         "MachineDeploymentClass selects templates referenced in specific MachineDeploymentClasses in.spec.workers.machineDeployments.",
																MarkdownDescription: "MachineDeploymentClass selects templates referenced in specific MachineDeploymentClasses in.spec.workers.machineDeployments.",
																Attributes: map[string]schema.Attribute{
																	"names": schema.ListAttribute{
																		Description:         "Names selects templates by class names.",
																		MarkdownDescription: "Names selects templates by class names.",
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

															"machine_pool_class": schema.SingleNestedAttribute{
																Description:         "MachinePoolClass selects templates referenced in specific MachinePoolClasses in.spec.workers.machinePools.",
																MarkdownDescription: "MachinePoolClass selects templates referenced in specific MachinePoolClasses in.spec.workers.machinePools.",
																Attributes: map[string]schema.Attribute{
																	"names": schema.ListAttribute{
																		Description:         "Names selects templates by class names.",
																		MarkdownDescription: "Names selects templates by class names.",
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

								"description": schema.StringAttribute{
									Description:         "Description is a human-readable description of this patch.",
									MarkdownDescription: "Description is a human-readable description of this patch.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"enabled_if": schema.StringAttribute{
									Description:         "EnabledIf is a Go template to be used to calculate if a patch should be enabled.It can reference variables defined in .spec.variables and builtin variables.The patch will be enabled if the template evaluates to 'true', otherwise it willbe disabled.If EnabledIf is not set, the patch will be enabled per default.",
									MarkdownDescription: "EnabledIf is a Go template to be used to calculate if a patch should be enabled.It can reference variables defined in .spec.variables and builtin variables.The patch will be enabled if the template evaluates to 'true', otherwise it willbe disabled.If EnabledIf is not set, the patch will be enabled per default.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"external": schema.SingleNestedAttribute{
									Description:         "External defines an external patch.Note: Exactly one of Definitions or External must be set.",
									MarkdownDescription: "External defines an external patch.Note: Exactly one of Definitions or External must be set.",
									Attributes: map[string]schema.Attribute{
										"discover_variables_extension": schema.StringAttribute{
											Description:         "DiscoverVariablesExtension references an extension which is called to discover variables.",
											MarkdownDescription: "DiscoverVariablesExtension references an extension which is called to discover variables.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"generate_extension": schema.StringAttribute{
											Description:         "GenerateExtension references an extension which is called to generate patches.",
											MarkdownDescription: "GenerateExtension references an extension which is called to generate patches.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"settings": schema.MapAttribute{
											Description:         "Settings defines key value pairs to be passed to the extensions.Values defined here take precedence over the values defined in thecorresponding ExtensionConfig.",
											MarkdownDescription: "Settings defines key value pairs to be passed to the extensions.Values defined here take precedence over the values defined in thecorresponding ExtensionConfig.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"validate_extension": schema.StringAttribute{
											Description:         "ValidateExtension references an extension which is called to validate the topology.",
											MarkdownDescription: "ValidateExtension references an extension which is called to validate the topology.",
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
									Description:         "Name of the patch.",
									MarkdownDescription: "Name of the patch.",
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

					"variables": schema.ListNestedAttribute{
						Description:         "Variables defines the variables which can be configuredin the Cluster topology and are then used in patches.",
						MarkdownDescription: "Variables defines the variables which can be configuredin the Cluster topology and are then used in patches.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"metadata": schema.SingleNestedAttribute{
									Description:         "Metadata is the metadata of a variable.It can be used to add additional data for higher level tools toa ClusterClassVariable.",
									MarkdownDescription: "Metadata is the metadata of a variable.It can be used to add additional data for higher level tools toa ClusterClassVariable.",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "Annotations is an unstructured key value map that can be used to store andretrieve arbitrary metadata.They are not queryable.",
											MarkdownDescription: "Annotations is an unstructured key value map that can be used to store andretrieve arbitrary metadata.They are not queryable.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"labels": schema.MapAttribute{
											Description:         "Map of string keys and values that can be used to organize and categorize(scope and select) variables.",
											MarkdownDescription: "Map of string keys and values that can be used to organize and categorize(scope and select) variables.",
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
									Description:         "Name of the variable.",
									MarkdownDescription: "Name of the variable.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"required": schema.BoolAttribute{
									Description:         "Required specifies if the variable is required.Note: this applies to the variable as a whole and thus thetop-level object defined in the schema. If nested fields arerequired, this will be specified inside the schema.",
									MarkdownDescription: "Required specifies if the variable is required.Note: this applies to the variable as a whole and thus thetop-level object defined in the schema. If nested fields arerequired, this will be specified inside the schema.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"schema": schema.SingleNestedAttribute{
									Description:         "Schema defines the schema of the variable.",
									MarkdownDescription: "Schema defines the schema of the variable.",
									Attributes: map[string]schema.Attribute{
										"open_apiv3_schema": schema.SingleNestedAttribute{
											Description:         "OpenAPIV3Schema defines the schema of a variable via OpenAPI v3schema. The schema is a subset of the schema used inKubernetes CRDs.",
											MarkdownDescription: "OpenAPIV3Schema defines the schema of a variable via OpenAPI v3schema. The schema is a subset of the schema used inKubernetes CRDs.",
											Attributes: map[string]schema.Attribute{
												"additional_properties": schema.MapAttribute{
													Description:         "AdditionalProperties specifies the schema of values in a map (keys are always strings).NOTE: Can only be set if type is object.NOTE: AdditionalProperties is mutually exclusive with Properties.NOTE: This field uses PreserveUnknownFields and Schemaless,because recursive validation is not possible.",
													MarkdownDescription: "AdditionalProperties specifies the schema of values in a map (keys are always strings).NOTE: Can only be set if type is object.NOTE: AdditionalProperties is mutually exclusive with Properties.NOTE: This field uses PreserveUnknownFields and Schemaless,because recursive validation is not possible.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"default": schema.MapAttribute{
													Description:         "Default is the default value of the variable.NOTE: Can be set for all types.",
													MarkdownDescription: "Default is the default value of the variable.NOTE: Can be set for all types.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "Description is a human-readable description of this variable.",
													MarkdownDescription: "Description is a human-readable description of this variable.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"enum": schema.ListAttribute{
													Description:         "Enum is the list of valid values of the variable.NOTE: Can be set for all types.",
													MarkdownDescription: "Enum is the list of valid values of the variable.NOTE: Can be set for all types.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"example": schema.MapAttribute{
													Description:         "Example is an example for this variable.",
													MarkdownDescription: "Example is an example for this variable.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"exclusive_maximum": schema.BoolAttribute{
													Description:         "ExclusiveMaximum specifies if the Maximum is exclusive.NOTE: Can only be set if type is integer or number.",
													MarkdownDescription: "ExclusiveMaximum specifies if the Maximum is exclusive.NOTE: Can only be set if type is integer or number.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"exclusive_minimum": schema.BoolAttribute{
													Description:         "ExclusiveMinimum specifies if the Minimum is exclusive.NOTE: Can only be set if type is integer or number.",
													MarkdownDescription: "ExclusiveMinimum specifies if the Minimum is exclusive.NOTE: Can only be set if type is integer or number.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"format": schema.StringAttribute{
													Description:         "Format is an OpenAPI v3 format string. Unknown formats are ignored.For a list of supported formats please see: (of the k8s.io/apiextensions-apiserver version we're currently using)https://github.com/kubernetes/apiextensions-apiserver/blob/master/pkg/apiserver/validation/formats.goNOTE: Can only be set if type is string.",
													MarkdownDescription: "Format is an OpenAPI v3 format string. Unknown formats are ignored.For a list of supported formats please see: (of the k8s.io/apiextensions-apiserver version we're currently using)https://github.com/kubernetes/apiextensions-apiserver/blob/master/pkg/apiserver/validation/formats.goNOTE: Can only be set if type is string.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"items": schema.MapAttribute{
													Description:         "Items specifies fields of an array.NOTE: Can only be set if type is array.NOTE: This field uses PreserveUnknownFields and Schemaless,because recursive validation is not possible.",
													MarkdownDescription: "Items specifies fields of an array.NOTE: Can only be set if type is array.NOTE: This field uses PreserveUnknownFields and Schemaless,because recursive validation is not possible.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_items": schema.Int64Attribute{
													Description:         "MaxItems is the max length of an array variable.NOTE: Can only be set if type is array.",
													MarkdownDescription: "MaxItems is the max length of an array variable.NOTE: Can only be set if type is array.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_length": schema.Int64Attribute{
													Description:         "MaxLength is the max length of a string variable.NOTE: Can only be set if type is string.",
													MarkdownDescription: "MaxLength is the max length of a string variable.NOTE: Can only be set if type is string.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"maximum": schema.Int64Attribute{
													Description:         "Maximum is the maximum of an integer or number variable.If ExclusiveMaximum is false, the variable is valid if it is lower than, or equal to, the value of Maximum.If ExclusiveMaximum is true, the variable is valid if it is strictly lower than the value of Maximum.NOTE: Can only be set if type is integer or number.",
													MarkdownDescription: "Maximum is the maximum of an integer or number variable.If ExclusiveMaximum is false, the variable is valid if it is lower than, or equal to, the value of Maximum.If ExclusiveMaximum is true, the variable is valid if it is strictly lower than the value of Maximum.NOTE: Can only be set if type is integer or number.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"min_items": schema.Int64Attribute{
													Description:         "MinItems is the min length of an array variable.NOTE: Can only be set if type is array.",
													MarkdownDescription: "MinItems is the min length of an array variable.NOTE: Can only be set if type is array.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"min_length": schema.Int64Attribute{
													Description:         "MinLength is the min length of a string variable.NOTE: Can only be set if type is string.",
													MarkdownDescription: "MinLength is the min length of a string variable.NOTE: Can only be set if type is string.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"minimum": schema.Int64Attribute{
													Description:         "Minimum is the minimum of an integer or number variable.If ExclusiveMinimum is false, the variable is valid if it is greater than, or equal to, the value of Minimum.If ExclusiveMinimum is true, the variable is valid if it is strictly greater than the value of Minimum.NOTE: Can only be set if type is integer or number.",
													MarkdownDescription: "Minimum is the minimum of an integer or number variable.If ExclusiveMinimum is false, the variable is valid if it is greater than, or equal to, the value of Minimum.If ExclusiveMinimum is true, the variable is valid if it is strictly greater than the value of Minimum.NOTE: Can only be set if type is integer or number.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pattern": schema.StringAttribute{
													Description:         "Pattern is the regex which a string variable must match.NOTE: Can only be set if type is string.",
													MarkdownDescription: "Pattern is the regex which a string variable must match.NOTE: Can only be set if type is string.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"properties": schema.MapAttribute{
													Description:         "Properties specifies fields of an object.NOTE: Can only be set if type is object.NOTE: Properties is mutually exclusive with AdditionalProperties.NOTE: This field uses PreserveUnknownFields and Schemaless,because recursive validation is not possible.",
													MarkdownDescription: "Properties specifies fields of an object.NOTE: Can only be set if type is object.NOTE: Properties is mutually exclusive with AdditionalProperties.NOTE: This field uses PreserveUnknownFields and Schemaless,because recursive validation is not possible.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"required": schema.ListAttribute{
													Description:         "Required specifies which fields of an object are required.NOTE: Can only be set if type is object.",
													MarkdownDescription: "Required specifies which fields of an object are required.NOTE: Can only be set if type is object.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "Type is the type of the variable.Valid values are: object, array, string, integer, number or boolean.",
													MarkdownDescription: "Type is the type of the variable.Valid values are: object, array, string, integer, number or boolean.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"unique_items": schema.BoolAttribute{
													Description:         "UniqueItems specifies if items in an array must be unique.NOTE: Can only be set if type is array.",
													MarkdownDescription: "UniqueItems specifies if items in an array must be unique.NOTE: Can only be set if type is array.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"x_kubernetes_preserve_unknown_fields": schema.BoolAttribute{
													Description:         "XPreserveUnknownFields allows setting fields in a variable objectwhich are not defined in the variable schema. This affects fields recursively,except if nested properties or additionalProperties are specified in the schema.",
													MarkdownDescription: "XPreserveUnknownFields allows setting fields in a variable objectwhich are not defined in the variable schema. This affects fields recursively,except if nested properties or additionalProperties are specified in the schema.",
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

					"workers": schema.SingleNestedAttribute{
						Description:         "Workers describes the worker nodes for the cluster.It is a collection of node types which can be used to createthe worker nodes of the cluster.",
						MarkdownDescription: "Workers describes the worker nodes for the cluster.It is a collection of node types which can be used to createthe worker nodes of the cluster.",
						Attributes: map[string]schema.Attribute{
							"machine_deployments": schema.ListNestedAttribute{
								Description:         "MachineDeployments is a list of machine deployment classes that can be used to createa set of worker nodes.",
								MarkdownDescription: "MachineDeployments is a list of machine deployment classes that can be used to createa set of worker nodes.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"class": schema.StringAttribute{
											Description:         "Class denotes a type of worker node present in the cluster,this name MUST be unique within a ClusterClass and can be referencedin the Cluster to create a managed MachineDeployment.",
											MarkdownDescription: "Class denotes a type of worker node present in the cluster,this name MUST be unique within a ClusterClass and can be referencedin the Cluster to create a managed MachineDeployment.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"failure_domain": schema.StringAttribute{
											Description:         "FailureDomain is the failure domain the machines will be created in.Must match a key in the FailureDomains map stored on the cluster object.NOTE: This value can be overridden while defining a Cluster.Topology using this MachineDeploymentClass.",
											MarkdownDescription: "FailureDomain is the failure domain the machines will be created in.Must match a key in the FailureDomains map stored on the cluster object.NOTE: This value can be overridden while defining a Cluster.Topology using this MachineDeploymentClass.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"machine_health_check": schema.SingleNestedAttribute{
											Description:         "MachineHealthCheck defines a MachineHealthCheck for this MachineDeploymentClass.",
											MarkdownDescription: "MachineHealthCheck defines a MachineHealthCheck for this MachineDeploymentClass.",
											Attributes: map[string]schema.Attribute{
												"max_unhealthy": schema.StringAttribute{
													Description:         "Any further remediation is only allowed if at most 'MaxUnhealthy' machines selected by'selector' are not healthy.",
													MarkdownDescription: "Any further remediation is only allowed if at most 'MaxUnhealthy' machines selected by'selector' are not healthy.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_startup_timeout": schema.StringAttribute{
													Description:         "NodeStartupTimeout allows to set the maximum time for MachineHealthCheckto consider a Machine unhealthy if a corresponding Node isn't associatedthrough a 'Spec.ProviderID' field.The duration set in this field is compared to the greatest of:- Cluster's infrastructure and control plane ready condition timestamp (if and when available)- Machine's infrastructure ready condition timestamp (if and when available)- Machine's metadata creation timestampDefaults to 10 minutes.If you wish to disable this feature, set the value explicitly to 0.",
													MarkdownDescription: "NodeStartupTimeout allows to set the maximum time for MachineHealthCheckto consider a Machine unhealthy if a corresponding Node isn't associatedthrough a 'Spec.ProviderID' field.The duration set in this field is compared to the greatest of:- Cluster's infrastructure and control plane ready condition timestamp (if and when available)- Machine's infrastructure ready condition timestamp (if and when available)- Machine's metadata creation timestampDefaults to 10 minutes.If you wish to disable this feature, set the value explicitly to 0.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"remediation_template": schema.SingleNestedAttribute{
													Description:         "RemediationTemplate is a reference to a remediation templateprovided by an infrastructure provider.This field is completely optional, when filled, the MachineHealthCheck controllercreates a new object from the template referenced and hands off remediation of the machine toa controller that lives outside of Cluster API.",
													MarkdownDescription: "RemediationTemplate is a reference to a remediation templateprovided by an infrastructure provider.This field is completely optional, when filled, the MachineHealthCheck controllercreates a new object from the template referenced and hands off remediation of the machine toa controller that lives outside of Cluster API.",
													Attributes: map[string]schema.Attribute{
														"api_version": schema.StringAttribute{
															Description:         "API version of the referent.",
															MarkdownDescription: "API version of the referent.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"field_path": schema.StringAttribute{
															Description:         "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
															MarkdownDescription: "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"kind": schema.StringAttribute{
															Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
															MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
															MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
															Description:         "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
															MarkdownDescription: "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"resource_version": schema.StringAttribute{
															Description:         "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
															MarkdownDescription: "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"uid": schema.StringAttribute{
															Description:         "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
															MarkdownDescription: "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"unhealthy_conditions": schema.ListNestedAttribute{
													Description:         "UnhealthyConditions contains a list of the conditions that determinewhether a node is considered unhealthy. The conditions are combined in alogical OR, i.e. if any of the conditions is met, the node is unhealthy.",
													MarkdownDescription: "UnhealthyConditions contains a list of the conditions that determinewhether a node is considered unhealthy. The conditions are combined in alogical OR, i.e. if any of the conditions is met, the node is unhealthy.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"status": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},

															"timeout": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"unhealthy_range": schema.StringAttribute{
													Description:         "Any further remediation is only allowed if the number of machines selected by 'selector' as not healthyis within the range of 'UnhealthyRange'. Takes precedence over MaxUnhealthy.Eg. '[3-5]' - This means that remediation will be allowed only when:(a) there are at least 3 unhealthy machines (and)(b) there are at most 5 unhealthy machines",
													MarkdownDescription: "Any further remediation is only allowed if the number of machines selected by 'selector' as not healthyis within the range of 'UnhealthyRange'. Takes precedence over MaxUnhealthy.Eg. '[3-5]' - This means that remediation will be allowed only when:(a) there are at least 3 unhealthy machines (and)(b) there are at most 5 unhealthy machines",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^\[[0-9]+-[0-9]+\]$`), ""),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"min_ready_seconds": schema.Int64Attribute{
											Description:         "Minimum number of seconds for which a newly created machine shouldbe ready.Defaults to 0 (machine will be considered available as soon as itis ready)NOTE: This value can be overridden while defining a Cluster.Topology using this MachineDeploymentClass.",
											MarkdownDescription: "Minimum number of seconds for which a newly created machine shouldbe ready.Defaults to 0 (machine will be considered available as soon as itis ready)NOTE: This value can be overridden while defining a Cluster.Topology using this MachineDeploymentClass.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"naming_strategy": schema.SingleNestedAttribute{
											Description:         "NamingStrategy allows changing the naming pattern used when creating the MachineDeployment.",
											MarkdownDescription: "NamingStrategy allows changing the naming pattern used when creating the MachineDeployment.",
											Attributes: map[string]schema.Attribute{
												"template": schema.StringAttribute{
													Description:         "Template defines the template to use for generating the name of the MachineDeployment object.If not defined, it will fallback to '{{ .cluster.name }}-{{ .machineDeployment.topologyName }}-{{ .random }}'.If the templated string exceeds 63 characters, it will be trimmed to 58 characters and willget concatenated with a random suffix of length 5.The templating mechanism provides the following arguments:* '.cluster.name': The name of the cluster object.* '.random': A random alphanumeric string, without vowels, of length 5.* '.machineDeployment.topologyName': The name of the MachineDeployment topology (Cluster.spec.topology.workers.machineDeployments[].name).",
													MarkdownDescription: "Template defines the template to use for generating the name of the MachineDeployment object.If not defined, it will fallback to '{{ .cluster.name }}-{{ .machineDeployment.topologyName }}-{{ .random }}'.If the templated string exceeds 63 characters, it will be trimmed to 58 characters and willget concatenated with a random suffix of length 5.The templating mechanism provides the following arguments:* '.cluster.name': The name of the cluster object.* '.random': A random alphanumeric string, without vowels, of length 5.* '.machineDeployment.topologyName': The name of the MachineDeployment topology (Cluster.spec.topology.workers.machineDeployments[].name).",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"node_deletion_timeout": schema.StringAttribute{
											Description:         "NodeDeletionTimeout defines how long the controller will attempt to delete the Node that the Machinehosts after the Machine is marked for deletion. A duration of 0 will retry deletion indefinitely.Defaults to 10 seconds.NOTE: This value can be overridden while defining a Cluster.Topology using this MachineDeploymentClass.",
											MarkdownDescription: "NodeDeletionTimeout defines how long the controller will attempt to delete the Node that the Machinehosts after the Machine is marked for deletion. A duration of 0 will retry deletion indefinitely.Defaults to 10 seconds.NOTE: This value can be overridden while defining a Cluster.Topology using this MachineDeploymentClass.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"node_drain_timeout": schema.StringAttribute{
											Description:         "NodeDrainTimeout is the total amount of time that the controller will spend on draining a node.The default value is 0, meaning that the node can be drained without any time limitations.NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'NOTE: This value can be overridden while defining a Cluster.Topology using this MachineDeploymentClass.",
											MarkdownDescription: "NodeDrainTimeout is the total amount of time that the controller will spend on draining a node.The default value is 0, meaning that the node can be drained without any time limitations.NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'NOTE: This value can be overridden while defining a Cluster.Topology using this MachineDeploymentClass.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"node_volume_detach_timeout": schema.StringAttribute{
											Description:         "NodeVolumeDetachTimeout is the total amount of time that the controller will spend on waiting for all volumesto be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.NOTE: This value can be overridden while defining a Cluster.Topology using this MachineDeploymentClass.",
											MarkdownDescription: "NodeVolumeDetachTimeout is the total amount of time that the controller will spend on waiting for all volumesto be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.NOTE: This value can be overridden while defining a Cluster.Topology using this MachineDeploymentClass.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"strategy": schema.SingleNestedAttribute{
											Description:         "The deployment strategy to use to replace existing machines withnew ones.NOTE: This value can be overridden while defining a Cluster.Topology using this MachineDeploymentClass.",
											MarkdownDescription: "The deployment strategy to use to replace existing machines withnew ones.NOTE: This value can be overridden while defining a Cluster.Topology using this MachineDeploymentClass.",
											Attributes: map[string]schema.Attribute{
												"rolling_update": schema.SingleNestedAttribute{
													Description:         "Rolling update config params. Present only ifMachineDeploymentStrategyType = RollingUpdate.",
													MarkdownDescription: "Rolling update config params. Present only ifMachineDeploymentStrategyType = RollingUpdate.",
													Attributes: map[string]schema.Attribute{
														"delete_policy": schema.StringAttribute{
															Description:         "DeletePolicy defines the policy used by the MachineDeployment to identify nodes to delete when downscaling.Valid values are 'Random, 'Newest', 'Oldest'When no value is supplied, the default DeletePolicy of MachineSet is used",
															MarkdownDescription: "DeletePolicy defines the policy used by the MachineDeployment to identify nodes to delete when downscaling.Valid values are 'Random, 'Newest', 'Oldest'When no value is supplied, the default DeletePolicy of MachineSet is used",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("Random", "Newest", "Oldest"),
															},
														},

														"max_surge": schema.StringAttribute{
															Description:         "The maximum number of machines that can be scheduled above thedesired number of machines.Value can be an absolute number (ex: 5) or a percentage ofdesired machines (ex: 10%).This can not be 0 if MaxUnavailable is 0.Absolute number is calculated from percentage by rounding up.Defaults to 1.Example: when this is set to 30%, the new MachineSet can be scaledup immediately when the rolling update starts, such that the totalnumber of old and new machines do not exceed 130% of desiredmachines. Once old machines have been killed, new MachineSet canbe scaled up further, ensuring that total number of machines runningat any time during the update is at most 130% of desired machines.",
															MarkdownDescription: "The maximum number of machines that can be scheduled above thedesired number of machines.Value can be an absolute number (ex: 5) or a percentage ofdesired machines (ex: 10%).This can not be 0 if MaxUnavailable is 0.Absolute number is calculated from percentage by rounding up.Defaults to 1.Example: when this is set to 30%, the new MachineSet can be scaledup immediately when the rolling update starts, such that the totalnumber of old and new machines do not exceed 130% of desiredmachines. Once old machines have been killed, new MachineSet canbe scaled up further, ensuring that total number of machines runningat any time during the update is at most 130% of desired machines.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_unavailable": schema.StringAttribute{
															Description:         "The maximum number of machines that can be unavailable during the update.Value can be an absolute number (ex: 5) or a percentage of desiredmachines (ex: 10%).Absolute number is calculated from percentage by rounding down.This can not be 0 if MaxSurge is 0.Defaults to 0.Example: when this is set to 30%, the old MachineSet can be scaleddown to 70% of desired machines immediately when the rolling updatestarts. Once new machines are ready, old MachineSet can be scaleddown further, followed by scaling up the new MachineSet, ensuringthat the total number of machines available at all timesduring the update is at least 70% of desired machines.",
															MarkdownDescription: "The maximum number of machines that can be unavailable during the update.Value can be an absolute number (ex: 5) or a percentage of desiredmachines (ex: 10%).Absolute number is calculated from percentage by rounding down.This can not be 0 if MaxSurge is 0.Defaults to 0.Example: when this is set to 30%, the old MachineSet can be scaleddown to 70% of desired machines immediately when the rolling updatestarts. Once new machines are ready, old MachineSet can be scaleddown further, followed by scaling up the new MachineSet, ensuringthat the total number of machines available at all timesduring the update is at least 70% of desired machines.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"type": schema.StringAttribute{
													Description:         "Type of deployment. Allowed values are RollingUpdate and OnDelete.The default is RollingUpdate.",
													MarkdownDescription: "Type of deployment. Allowed values are RollingUpdate and OnDelete.The default is RollingUpdate.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("RollingUpdate", "OnDelete"),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"template": schema.SingleNestedAttribute{
											Description:         "Template is a local struct containing a collection of templates for creation ofMachineDeployment objects representing a set of worker nodes.",
											MarkdownDescription: "Template is a local struct containing a collection of templates for creation ofMachineDeployment objects representing a set of worker nodes.",
											Attributes: map[string]schema.Attribute{
												"bootstrap": schema.SingleNestedAttribute{
													Description:         "Bootstrap contains the bootstrap template reference to be usedfor the creation of worker Machines.",
													MarkdownDescription: "Bootstrap contains the bootstrap template reference to be usedfor the creation of worker Machines.",
													Attributes: map[string]schema.Attribute{
														"ref": schema.SingleNestedAttribute{
															Description:         "Ref is a required reference to a custom resourceoffered by a provider.",
															MarkdownDescription: "Ref is a required reference to a custom resourceoffered by a provider.",
															Attributes: map[string]schema.Attribute{
																"api_version": schema.StringAttribute{
																	Description:         "API version of the referent.",
																	MarkdownDescription: "API version of the referent.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"field_path": schema.StringAttribute{
																	Description:         "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
																	MarkdownDescription: "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"kind": schema.StringAttribute{
																	Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																	MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"namespace": schema.StringAttribute{
																	Description:         "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
																	MarkdownDescription: "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"resource_version": schema.StringAttribute{
																	Description:         "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
																	MarkdownDescription: "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"uid": schema.StringAttribute{
																	Description:         "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
																	MarkdownDescription: "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
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
													Required: true,
													Optional: false,
													Computed: false,
												},

												"infrastructure": schema.SingleNestedAttribute{
													Description:         "Infrastructure contains the infrastructure template reference to be usedfor the creation of worker Machines.",
													MarkdownDescription: "Infrastructure contains the infrastructure template reference to be usedfor the creation of worker Machines.",
													Attributes: map[string]schema.Attribute{
														"ref": schema.SingleNestedAttribute{
															Description:         "Ref is a required reference to a custom resourceoffered by a provider.",
															MarkdownDescription: "Ref is a required reference to a custom resourceoffered by a provider.",
															Attributes: map[string]schema.Attribute{
																"api_version": schema.StringAttribute{
																	Description:         "API version of the referent.",
																	MarkdownDescription: "API version of the referent.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"field_path": schema.StringAttribute{
																	Description:         "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
																	MarkdownDescription: "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"kind": schema.StringAttribute{
																	Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																	MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"namespace": schema.StringAttribute{
																	Description:         "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
																	MarkdownDescription: "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"resource_version": schema.StringAttribute{
																	Description:         "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
																	MarkdownDescription: "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"uid": schema.StringAttribute{
																	Description:         "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
																	MarkdownDescription: "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
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
													Required: true,
													Optional: false,
													Computed: false,
												},

												"metadata": schema.SingleNestedAttribute{
													Description:         "Metadata is the metadata applied to the MachineDeployment and the machines of the MachineDeployment.At runtime this metadata is merged with the corresponding metadata from the topology.",
													MarkdownDescription: "Metadata is the metadata applied to the MachineDeployment and the machines of the MachineDeployment.At runtime this metadata is merged with the corresponding metadata from the topology.",
													Attributes: map[string]schema.Attribute{
														"annotations": schema.MapAttribute{
															Description:         "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: http://kubernetes.io/docs/user-guide/annotations",
															MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: http://kubernetes.io/docs/user-guide/annotations",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"labels": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: http://kubernetes.io/docs/user-guide/labels",
															MarkdownDescription: "Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: http://kubernetes.io/docs/user-guide/labels",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"machine_pools": schema.ListNestedAttribute{
								Description:         "MachinePools is a list of machine pool classes that can be used to createa set of worker nodes.",
								MarkdownDescription: "MachinePools is a list of machine pool classes that can be used to createa set of worker nodes.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"class": schema.StringAttribute{
											Description:         "Class denotes a type of machine pool present in the cluster,this name MUST be unique within a ClusterClass and can be referencedin the Cluster to create a managed MachinePool.",
											MarkdownDescription: "Class denotes a type of machine pool present in the cluster,this name MUST be unique within a ClusterClass and can be referencedin the Cluster to create a managed MachinePool.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"failure_domains": schema.ListAttribute{
											Description:         "FailureDomains is the list of failure domains the MachinePool should be attached to.Must match a key in the FailureDomains map stored on the cluster object.NOTE: This value can be overridden while defining a Cluster.Topology using this MachinePoolClass.",
											MarkdownDescription: "FailureDomains is the list of failure domains the MachinePool should be attached to.Must match a key in the FailureDomains map stored on the cluster object.NOTE: This value can be overridden while defining a Cluster.Topology using this MachinePoolClass.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"min_ready_seconds": schema.Int64Attribute{
											Description:         "Minimum number of seconds for which a newly created machine pool shouldbe ready.Defaults to 0 (machine will be considered available as soon as itis ready)NOTE: This value can be overridden while defining a Cluster.Topology using this MachinePoolClass.",
											MarkdownDescription: "Minimum number of seconds for which a newly created machine pool shouldbe ready.Defaults to 0 (machine will be considered available as soon as itis ready)NOTE: This value can be overridden while defining a Cluster.Topology using this MachinePoolClass.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"naming_strategy": schema.SingleNestedAttribute{
											Description:         "NamingStrategy allows changing the naming pattern used when creating the MachinePool.",
											MarkdownDescription: "NamingStrategy allows changing the naming pattern used when creating the MachinePool.",
											Attributes: map[string]schema.Attribute{
												"template": schema.StringAttribute{
													Description:         "Template defines the template to use for generating the name of the MachinePool object.If not defined, it will fallback to '{{ .cluster.name }}-{{ .machinePool.topologyName }}-{{ .random }}'.If the templated string exceeds 63 characters, it will be trimmed to 58 characters and willget concatenated with a random suffix of length 5.The templating mechanism provides the following arguments:* '.cluster.name': The name of the cluster object.* '.random': A random alphanumeric string, without vowels, of length 5.* '.machinePool.topologyName': The name of the MachinePool topology (Cluster.spec.topology.workers.machinePools[].name).",
													MarkdownDescription: "Template defines the template to use for generating the name of the MachinePool object.If not defined, it will fallback to '{{ .cluster.name }}-{{ .machinePool.topologyName }}-{{ .random }}'.If the templated string exceeds 63 characters, it will be trimmed to 58 characters and willget concatenated with a random suffix of length 5.The templating mechanism provides the following arguments:* '.cluster.name': The name of the cluster object.* '.random': A random alphanumeric string, without vowels, of length 5.* '.machinePool.topologyName': The name of the MachinePool topology (Cluster.spec.topology.workers.machinePools[].name).",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"node_deletion_timeout": schema.StringAttribute{
											Description:         "NodeDeletionTimeout defines how long the controller will attempt to delete the Node that the Machinehosts after the Machine Pool is marked for deletion. A duration of 0 will retry deletion indefinitely.Defaults to 10 seconds.NOTE: This value can be overridden while defining a Cluster.Topology using this MachinePoolClass.",
											MarkdownDescription: "NodeDeletionTimeout defines how long the controller will attempt to delete the Node that the Machinehosts after the Machine Pool is marked for deletion. A duration of 0 will retry deletion indefinitely.Defaults to 10 seconds.NOTE: This value can be overridden while defining a Cluster.Topology using this MachinePoolClass.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"node_drain_timeout": schema.StringAttribute{
											Description:         "NodeDrainTimeout is the total amount of time that the controller will spend on draining a node.The default value is 0, meaning that the node can be drained without any time limitations.NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'NOTE: This value can be overridden while defining a Cluster.Topology using this MachinePoolClass.",
											MarkdownDescription: "NodeDrainTimeout is the total amount of time that the controller will spend on draining a node.The default value is 0, meaning that the node can be drained without any time limitations.NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'NOTE: This value can be overridden while defining a Cluster.Topology using this MachinePoolClass.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"node_volume_detach_timeout": schema.StringAttribute{
											Description:         "NodeVolumeDetachTimeout is the total amount of time that the controller will spend on waiting for all volumesto be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.NOTE: This value can be overridden while defining a Cluster.Topology using this MachinePoolClass.",
											MarkdownDescription: "NodeVolumeDetachTimeout is the total amount of time that the controller will spend on waiting for all volumesto be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.NOTE: This value can be overridden while defining a Cluster.Topology using this MachinePoolClass.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.SingleNestedAttribute{
											Description:         "Template is a local struct containing a collection of templates for creation ofMachinePools objects representing a pool of worker nodes.",
											MarkdownDescription: "Template is a local struct containing a collection of templates for creation ofMachinePools objects representing a pool of worker nodes.",
											Attributes: map[string]schema.Attribute{
												"bootstrap": schema.SingleNestedAttribute{
													Description:         "Bootstrap contains the bootstrap template reference to be usedfor the creation of the Machines in the MachinePool.",
													MarkdownDescription: "Bootstrap contains the bootstrap template reference to be usedfor the creation of the Machines in the MachinePool.",
													Attributes: map[string]schema.Attribute{
														"ref": schema.SingleNestedAttribute{
															Description:         "Ref is a required reference to a custom resourceoffered by a provider.",
															MarkdownDescription: "Ref is a required reference to a custom resourceoffered by a provider.",
															Attributes: map[string]schema.Attribute{
																"api_version": schema.StringAttribute{
																	Description:         "API version of the referent.",
																	MarkdownDescription: "API version of the referent.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"field_path": schema.StringAttribute{
																	Description:         "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
																	MarkdownDescription: "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"kind": schema.StringAttribute{
																	Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																	MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"namespace": schema.StringAttribute{
																	Description:         "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
																	MarkdownDescription: "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"resource_version": schema.StringAttribute{
																	Description:         "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
																	MarkdownDescription: "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"uid": schema.StringAttribute{
																	Description:         "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
																	MarkdownDescription: "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
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
													Required: true,
													Optional: false,
													Computed: false,
												},

												"infrastructure": schema.SingleNestedAttribute{
													Description:         "Infrastructure contains the infrastructure template reference to be usedfor the creation of the MachinePool.",
													MarkdownDescription: "Infrastructure contains the infrastructure template reference to be usedfor the creation of the MachinePool.",
													Attributes: map[string]schema.Attribute{
														"ref": schema.SingleNestedAttribute{
															Description:         "Ref is a required reference to a custom resourceoffered by a provider.",
															MarkdownDescription: "Ref is a required reference to a custom resourceoffered by a provider.",
															Attributes: map[string]schema.Attribute{
																"api_version": schema.StringAttribute{
																	Description:         "API version of the referent.",
																	MarkdownDescription: "API version of the referent.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"field_path": schema.StringAttribute{
																	Description:         "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
																	MarkdownDescription: "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"kind": schema.StringAttribute{
																	Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																	MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"namespace": schema.StringAttribute{
																	Description:         "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
																	MarkdownDescription: "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"resource_version": schema.StringAttribute{
																	Description:         "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
																	MarkdownDescription: "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"uid": schema.StringAttribute{
																	Description:         "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
																	MarkdownDescription: "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
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
													Required: true,
													Optional: false,
													Computed: false,
												},

												"metadata": schema.SingleNestedAttribute{
													Description:         "Metadata is the metadata applied to the MachinePool.At runtime this metadata is merged with the corresponding metadata from the topology.",
													MarkdownDescription: "Metadata is the metadata applied to the MachinePool.At runtime this metadata is merged with the corresponding metadata from the topology.",
													Attributes: map[string]schema.Attribute{
														"annotations": schema.MapAttribute{
															Description:         "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: http://kubernetes.io/docs/user-guide/annotations",
															MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: http://kubernetes.io/docs/user-guide/annotations",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"labels": schema.MapAttribute{
															Description:         "Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: http://kubernetes.io/docs/user-guide/labels",
															MarkdownDescription: "Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: http://kubernetes.io/docs/user-guide/labels",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *ClusterXK8SIoClusterClassV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cluster_x_k8s_io_cluster_class_v1beta1_manifest")

	var model ClusterXK8SIoClusterClassV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cluster.x-k8s.io/v1beta1")
	model.Kind = pointer.String("ClusterClass")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
