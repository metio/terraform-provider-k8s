/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1alpha1

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
	_ datasource.DataSource = &InfrastructureClusterXK8SIoKubevirtMachineTemplateV1Alpha1Manifest{}
)

func NewInfrastructureClusterXK8SIoKubevirtMachineTemplateV1Alpha1Manifest() datasource.DataSource {
	return &InfrastructureClusterXK8SIoKubevirtMachineTemplateV1Alpha1Manifest{}
}

type InfrastructureClusterXK8SIoKubevirtMachineTemplateV1Alpha1Manifest struct{}

type InfrastructureClusterXK8SIoKubevirtMachineTemplateV1Alpha1ManifestData struct {
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
		Template *struct {
			Spec *struct {
				InfraClusterSecretRef *struct {
					ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
					Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
					Name            *string `tfsdk:"name" json:"name,omitempty"`
					Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
					ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
					Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
				} `tfsdk:"infra_cluster_secret_ref" json:"infraClusterSecretRef,omitempty"`
				ProviderID                   *string `tfsdk:"provider_id" json:"providerID,omitempty"`
				VirtualMachineBootstrapCheck *struct {
					CheckStrategy *string `tfsdk:"check_strategy" json:"checkStrategy,omitempty"`
				} `tfsdk:"virtual_machine_bootstrap_check" json:"virtualMachineBootstrapCheck,omitempty"`
				VirtualMachineTemplate *struct {
					Metadata *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
					Spec     *struct {
						DataVolumeTemplates *[]struct {
							ApiVersion *string            `tfsdk:"api_version" json:"apiVersion,omitempty"`
							Kind       *string            `tfsdk:"kind" json:"kind,omitempty"`
							Metadata   *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
							Spec       *struct {
								Checkpoints *[]struct {
									Current  *string `tfsdk:"current" json:"current,omitempty"`
									Previous *string `tfsdk:"previous" json:"previous,omitempty"`
								} `tfsdk:"checkpoints" json:"checkpoints,omitempty"`
								ContentType       *string `tfsdk:"content_type" json:"contentType,omitempty"`
								FinalCheckpoint   *bool   `tfsdk:"final_checkpoint" json:"finalCheckpoint,omitempty"`
								Preallocation     *bool   `tfsdk:"preallocation" json:"preallocation,omitempty"`
								PriorityClassName *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
								Pvc               *struct {
									AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
									DataSource  *struct {
										ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
										Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
										Name     *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"data_source" json:"dataSource,omitempty"`
									DataSourceRef *struct {
										ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
										Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
										Name      *string `tfsdk:"name" json:"name,omitempty"`
										Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
									} `tfsdk:"data_source_ref" json:"dataSourceRef,omitempty"`
									Resources *struct {
										Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
										Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
									} `tfsdk:"resources" json:"resources,omitempty"`
									Selector *struct {
										MatchExpressions *[]struct {
											Key      *string   `tfsdk:"key" json:"key,omitempty"`
											Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
											Values   *[]string `tfsdk:"values" json:"values,omitempty"`
										} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
										MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
									} `tfsdk:"selector" json:"selector,omitempty"`
									StorageClassName          *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
									VolumeAttributesClassName *string `tfsdk:"volume_attributes_class_name" json:"volumeAttributesClassName,omitempty"`
									VolumeMode                *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
									VolumeName                *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
								} `tfsdk:"pvc" json:"pvc,omitempty"`
								Source *struct {
									Blank *map[string]string `tfsdk:"blank" json:"blank,omitempty"`
									Gcs   *struct {
										SecretRef *string `tfsdk:"secret_ref" json:"secretRef,omitempty"`
										Url       *string `tfsdk:"url" json:"url,omitempty"`
									} `tfsdk:"gcs" json:"gcs,omitempty"`
									Http *struct {
										CertConfigMap      *string   `tfsdk:"cert_config_map" json:"certConfigMap,omitempty"`
										ExtraHeaders       *[]string `tfsdk:"extra_headers" json:"extraHeaders,omitempty"`
										SecretExtraHeaders *[]string `tfsdk:"secret_extra_headers" json:"secretExtraHeaders,omitempty"`
										SecretRef          *string   `tfsdk:"secret_ref" json:"secretRef,omitempty"`
										Url                *string   `tfsdk:"url" json:"url,omitempty"`
									} `tfsdk:"http" json:"http,omitempty"`
									Imageio *struct {
										CertConfigMap *string `tfsdk:"cert_config_map" json:"certConfigMap,omitempty"`
										DiskId        *string `tfsdk:"disk_id" json:"diskId,omitempty"`
										SecretRef     *string `tfsdk:"secret_ref" json:"secretRef,omitempty"`
										Url           *string `tfsdk:"url" json:"url,omitempty"`
									} `tfsdk:"imageio" json:"imageio,omitempty"`
									Pvc *struct {
										Name      *string `tfsdk:"name" json:"name,omitempty"`
										Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
									} `tfsdk:"pvc" json:"pvc,omitempty"`
									Registry *struct {
										CertConfigMap *string `tfsdk:"cert_config_map" json:"certConfigMap,omitempty"`
										ImageStream   *string `tfsdk:"image_stream" json:"imageStream,omitempty"`
										PullMethod    *string `tfsdk:"pull_method" json:"pullMethod,omitempty"`
										SecretRef     *string `tfsdk:"secret_ref" json:"secretRef,omitempty"`
										Url           *string `tfsdk:"url" json:"url,omitempty"`
									} `tfsdk:"registry" json:"registry,omitempty"`
									S3 *struct {
										CertConfigMap *string `tfsdk:"cert_config_map" json:"certConfigMap,omitempty"`
										SecretRef     *string `tfsdk:"secret_ref" json:"secretRef,omitempty"`
										Url           *string `tfsdk:"url" json:"url,omitempty"`
									} `tfsdk:"s3" json:"s3,omitempty"`
									Snapshot *struct {
										Name      *string `tfsdk:"name" json:"name,omitempty"`
										Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
									} `tfsdk:"snapshot" json:"snapshot,omitempty"`
									Upload *map[string]string `tfsdk:"upload" json:"upload,omitempty"`
									Vddk   *struct {
										BackingFile  *string `tfsdk:"backing_file" json:"backingFile,omitempty"`
										InitImageURL *string `tfsdk:"init_image_url" json:"initImageURL,omitempty"`
										SecretRef    *string `tfsdk:"secret_ref" json:"secretRef,omitempty"`
										Thumbprint   *string `tfsdk:"thumbprint" json:"thumbprint,omitempty"`
										Url          *string `tfsdk:"url" json:"url,omitempty"`
										Uuid         *string `tfsdk:"uuid" json:"uuid,omitempty"`
									} `tfsdk:"vddk" json:"vddk,omitempty"`
								} `tfsdk:"source" json:"source,omitempty"`
								SourceRef *struct {
									Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
									Name      *string `tfsdk:"name" json:"name,omitempty"`
									Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
								} `tfsdk:"source_ref" json:"sourceRef,omitempty"`
								Storage *struct {
									AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
									DataSource  *struct {
										ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
										Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
										Name     *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"data_source" json:"dataSource,omitempty"`
									DataSourceRef *struct {
										ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
										Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
										Name      *string `tfsdk:"name" json:"name,omitempty"`
										Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
									} `tfsdk:"data_source_ref" json:"dataSourceRef,omitempty"`
									Resources *struct {
										Claims *[]struct {
											Name *string `tfsdk:"name" json:"name,omitempty"`
										} `tfsdk:"claims" json:"claims,omitempty"`
										Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
										Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
									} `tfsdk:"resources" json:"resources,omitempty"`
									Selector *struct {
										MatchExpressions *[]struct {
											Key      *string   `tfsdk:"key" json:"key,omitempty"`
											Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
											Values   *[]string `tfsdk:"values" json:"values,omitempty"`
										} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
										MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
									} `tfsdk:"selector" json:"selector,omitempty"`
									StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
									VolumeMode       *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
									VolumeName       *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
								} `tfsdk:"storage" json:"storage,omitempty"`
							} `tfsdk:"spec" json:"spec,omitempty"`
							Status *map[string]string `tfsdk:"status" json:"status,omitempty"`
						} `tfsdk:"data_volume_templates" json:"dataVolumeTemplates,omitempty"`
						Instancetype *struct {
							InferFromVolume              *string `tfsdk:"infer_from_volume" json:"inferFromVolume,omitempty"`
							InferFromVolumeFailurePolicy *string `tfsdk:"infer_from_volume_failure_policy" json:"inferFromVolumeFailurePolicy,omitempty"`
							Kind                         *string `tfsdk:"kind" json:"kind,omitempty"`
							Name                         *string `tfsdk:"name" json:"name,omitempty"`
							RevisionName                 *string `tfsdk:"revision_name" json:"revisionName,omitempty"`
						} `tfsdk:"instancetype" json:"instancetype,omitempty"`
						Preference *struct {
							InferFromVolume              *string `tfsdk:"infer_from_volume" json:"inferFromVolume,omitempty"`
							InferFromVolumeFailurePolicy *string `tfsdk:"infer_from_volume_failure_policy" json:"inferFromVolumeFailurePolicy,omitempty"`
							Kind                         *string `tfsdk:"kind" json:"kind,omitempty"`
							Name                         *string `tfsdk:"name" json:"name,omitempty"`
							RevisionName                 *string `tfsdk:"revision_name" json:"revisionName,omitempty"`
						} `tfsdk:"preference" json:"preference,omitempty"`
						RunStrategy *string `tfsdk:"run_strategy" json:"runStrategy,omitempty"`
						Running     *bool   `tfsdk:"running" json:"running,omitempty"`
						Template    *struct {
							Metadata *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
							Spec     *struct {
								AccessCredentials *[]struct {
									SshPublicKey *struct {
										PropagationMethod *struct {
											ConfigDrive    *map[string]string `tfsdk:"config_drive" json:"configDrive,omitempty"`
											NoCloud        *map[string]string `tfsdk:"no_cloud" json:"noCloud,omitempty"`
											QemuGuestAgent *struct {
												Users *[]string `tfsdk:"users" json:"users,omitempty"`
											} `tfsdk:"qemu_guest_agent" json:"qemuGuestAgent,omitempty"`
										} `tfsdk:"propagation_method" json:"propagationMethod,omitempty"`
										Source *struct {
											Secret *struct {
												SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
											} `tfsdk:"secret" json:"secret,omitempty"`
										} `tfsdk:"source" json:"source,omitempty"`
									} `tfsdk:"ssh_public_key" json:"sshPublicKey,omitempty"`
									UserPassword *struct {
										PropagationMethod *struct {
											QemuGuestAgent *map[string]string `tfsdk:"qemu_guest_agent" json:"qemuGuestAgent,omitempty"`
										} `tfsdk:"propagation_method" json:"propagationMethod,omitempty"`
										Source *struct {
											Secret *struct {
												SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
											} `tfsdk:"secret" json:"secret,omitempty"`
										} `tfsdk:"source" json:"source,omitempty"`
									} `tfsdk:"user_password" json:"userPassword,omitempty"`
								} `tfsdk:"access_credentials" json:"accessCredentials,omitempty"`
								Affinity *struct {
									NodeAffinity *struct {
										PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
											Preference *struct {
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
											} `tfsdk:"preference" json:"preference,omitempty"`
											Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
										} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
										RequiredDuringSchedulingIgnoredDuringExecution *struct {
											NodeSelectorTerms *[]struct {
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
											} `tfsdk:"node_selector_terms" json:"nodeSelectorTerms,omitempty"`
										} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
									} `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
									PodAffinity *struct {
										PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
											PodAffinityTerm *struct {
												LabelSelector *struct {
													MatchExpressions *[]struct {
														Key      *string   `tfsdk:"key" json:"key,omitempty"`
														Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
														Values   *[]string `tfsdk:"values" json:"values,omitempty"`
													} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
													MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
												} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
												MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
												MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
												NamespaceSelector *struct {
													MatchExpressions *[]struct {
														Key      *string   `tfsdk:"key" json:"key,omitempty"`
														Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
														Values   *[]string `tfsdk:"values" json:"values,omitempty"`
													} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
													MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
												} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
												Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
												TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
											} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
											Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
										} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
										RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
											LabelSelector *struct {
												MatchExpressions *[]struct {
													Key      *string   `tfsdk:"key" json:"key,omitempty"`
													Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
													Values   *[]string `tfsdk:"values" json:"values,omitempty"`
												} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
												MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
											} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
											MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
											MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
											NamespaceSelector *struct {
												MatchExpressions *[]struct {
													Key      *string   `tfsdk:"key" json:"key,omitempty"`
													Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
													Values   *[]string `tfsdk:"values" json:"values,omitempty"`
												} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
												MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
											} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
											Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
											TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
										} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
									} `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
									PodAntiAffinity *struct {
										PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
											PodAffinityTerm *struct {
												LabelSelector *struct {
													MatchExpressions *[]struct {
														Key      *string   `tfsdk:"key" json:"key,omitempty"`
														Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
														Values   *[]string `tfsdk:"values" json:"values,omitempty"`
													} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
													MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
												} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
												MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
												MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
												NamespaceSelector *struct {
													MatchExpressions *[]struct {
														Key      *string   `tfsdk:"key" json:"key,omitempty"`
														Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
														Values   *[]string `tfsdk:"values" json:"values,omitempty"`
													} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
													MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
												} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
												Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
												TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
											} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
											Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
										} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
										RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
											LabelSelector *struct {
												MatchExpressions *[]struct {
													Key      *string   `tfsdk:"key" json:"key,omitempty"`
													Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
													Values   *[]string `tfsdk:"values" json:"values,omitempty"`
												} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
												MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
											} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
											MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
											MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
											NamespaceSelector *struct {
												MatchExpressions *[]struct {
													Key      *string   `tfsdk:"key" json:"key,omitempty"`
													Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
													Values   *[]string `tfsdk:"values" json:"values,omitempty"`
												} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
												MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
											} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
											Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
											TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
										} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
									} `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
								} `tfsdk:"affinity" json:"affinity,omitempty"`
								Architecture *string `tfsdk:"architecture" json:"architecture,omitempty"`
								DnsConfig    *struct {
									Nameservers *[]string `tfsdk:"nameservers" json:"nameservers,omitempty"`
									Options     *[]struct {
										Name  *string `tfsdk:"name" json:"name,omitempty"`
										Value *string `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"options" json:"options,omitempty"`
									Searches *[]string `tfsdk:"searches" json:"searches,omitempty"`
								} `tfsdk:"dns_config" json:"dnsConfig,omitempty"`
								DnsPolicy *string `tfsdk:"dns_policy" json:"dnsPolicy,omitempty"`
								Domain    *struct {
									Chassis *struct {
										Asset        *string `tfsdk:"asset" json:"asset,omitempty"`
										Manufacturer *string `tfsdk:"manufacturer" json:"manufacturer,omitempty"`
										Serial       *string `tfsdk:"serial" json:"serial,omitempty"`
										Sku          *string `tfsdk:"sku" json:"sku,omitempty"`
										Version      *string `tfsdk:"version" json:"version,omitempty"`
									} `tfsdk:"chassis" json:"chassis,omitempty"`
									Clock *struct {
										Timer *struct {
											Hpet *struct {
												Present    *bool   `tfsdk:"present" json:"present,omitempty"`
												TickPolicy *string `tfsdk:"tick_policy" json:"tickPolicy,omitempty"`
											} `tfsdk:"hpet" json:"hpet,omitempty"`
											Hyperv *struct {
												Present *bool `tfsdk:"present" json:"present,omitempty"`
											} `tfsdk:"hyperv" json:"hyperv,omitempty"`
											Kvm *struct {
												Present *bool `tfsdk:"present" json:"present,omitempty"`
											} `tfsdk:"kvm" json:"kvm,omitempty"`
											Pit *struct {
												Present    *bool   `tfsdk:"present" json:"present,omitempty"`
												TickPolicy *string `tfsdk:"tick_policy" json:"tickPolicy,omitempty"`
											} `tfsdk:"pit" json:"pit,omitempty"`
											Rtc *struct {
												Present    *bool   `tfsdk:"present" json:"present,omitempty"`
												TickPolicy *string `tfsdk:"tick_policy" json:"tickPolicy,omitempty"`
												Track      *string `tfsdk:"track" json:"track,omitempty"`
											} `tfsdk:"rtc" json:"rtc,omitempty"`
										} `tfsdk:"timer" json:"timer,omitempty"`
										Timezone *string `tfsdk:"timezone" json:"timezone,omitempty"`
										Utc      *struct {
											OffsetSeconds *int64 `tfsdk:"offset_seconds" json:"offsetSeconds,omitempty"`
										} `tfsdk:"utc" json:"utc,omitempty"`
									} `tfsdk:"clock" json:"clock,omitempty"`
									Cpu *struct {
										Cores                 *int64 `tfsdk:"cores" json:"cores,omitempty"`
										DedicatedCpuPlacement *bool  `tfsdk:"dedicated_cpu_placement" json:"dedicatedCpuPlacement,omitempty"`
										Features              *[]struct {
											Name   *string `tfsdk:"name" json:"name,omitempty"`
											Policy *string `tfsdk:"policy" json:"policy,omitempty"`
										} `tfsdk:"features" json:"features,omitempty"`
										IsolateEmulatorThread *bool   `tfsdk:"isolate_emulator_thread" json:"isolateEmulatorThread,omitempty"`
										MaxSockets            *int64  `tfsdk:"max_sockets" json:"maxSockets,omitempty"`
										Model                 *string `tfsdk:"model" json:"model,omitempty"`
										Numa                  *struct {
											GuestMappingPassthrough *map[string]string `tfsdk:"guest_mapping_passthrough" json:"guestMappingPassthrough,omitempty"`
										} `tfsdk:"numa" json:"numa,omitempty"`
										Realtime *struct {
											Mask *string `tfsdk:"mask" json:"mask,omitempty"`
										} `tfsdk:"realtime" json:"realtime,omitempty"`
										Sockets *int64 `tfsdk:"sockets" json:"sockets,omitempty"`
										Threads *int64 `tfsdk:"threads" json:"threads,omitempty"`
									} `tfsdk:"cpu" json:"cpu,omitempty"`
									Devices *struct {
										AutoattachGraphicsDevice *bool              `tfsdk:"autoattach_graphics_device" json:"autoattachGraphicsDevice,omitempty"`
										AutoattachInputDevice    *bool              `tfsdk:"autoattach_input_device" json:"autoattachInputDevice,omitempty"`
										AutoattachMemBalloon     *bool              `tfsdk:"autoattach_mem_balloon" json:"autoattachMemBalloon,omitempty"`
										AutoattachPodInterface   *bool              `tfsdk:"autoattach_pod_interface" json:"autoattachPodInterface,omitempty"`
										AutoattachSerialConsole  *bool              `tfsdk:"autoattach_serial_console" json:"autoattachSerialConsole,omitempty"`
										AutoattachVSOCK          *bool              `tfsdk:"autoattach_vsock" json:"autoattachVSOCK,omitempty"`
										BlockMultiQueue          *bool              `tfsdk:"block_multi_queue" json:"blockMultiQueue,omitempty"`
										ClientPassthrough        *map[string]string `tfsdk:"client_passthrough" json:"clientPassthrough,omitempty"`
										DisableHotplug           *bool              `tfsdk:"disable_hotplug" json:"disableHotplug,omitempty"`
										Disks                    *[]struct {
											BlockSize *struct {
												Custom *struct {
													Logical  *int64 `tfsdk:"logical" json:"logical,omitempty"`
													Physical *int64 `tfsdk:"physical" json:"physical,omitempty"`
												} `tfsdk:"custom" json:"custom,omitempty"`
												MatchVolume *struct {
													Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
												} `tfsdk:"match_volume" json:"matchVolume,omitempty"`
											} `tfsdk:"block_size" json:"blockSize,omitempty"`
											BootOrder *int64  `tfsdk:"boot_order" json:"bootOrder,omitempty"`
											Cache     *string `tfsdk:"cache" json:"cache,omitempty"`
											Cdrom     *struct {
												Bus      *string `tfsdk:"bus" json:"bus,omitempty"`
												Readonly *bool   `tfsdk:"readonly" json:"readonly,omitempty"`
												Tray     *string `tfsdk:"tray" json:"tray,omitempty"`
											} `tfsdk:"cdrom" json:"cdrom,omitempty"`
											DedicatedIOThread *bool `tfsdk:"dedicated_io_thread" json:"dedicatedIOThread,omitempty"`
											Disk              *struct {
												Bus        *string `tfsdk:"bus" json:"bus,omitempty"`
												PciAddress *string `tfsdk:"pci_address" json:"pciAddress,omitempty"`
												Readonly   *bool   `tfsdk:"readonly" json:"readonly,omitempty"`
											} `tfsdk:"disk" json:"disk,omitempty"`
											ErrorPolicy *string `tfsdk:"error_policy" json:"errorPolicy,omitempty"`
											Io          *string `tfsdk:"io" json:"io,omitempty"`
											Lun         *struct {
												Bus         *string `tfsdk:"bus" json:"bus,omitempty"`
												Readonly    *bool   `tfsdk:"readonly" json:"readonly,omitempty"`
												Reservation *bool   `tfsdk:"reservation" json:"reservation,omitempty"`
											} `tfsdk:"lun" json:"lun,omitempty"`
											Name      *string `tfsdk:"name" json:"name,omitempty"`
											Serial    *string `tfsdk:"serial" json:"serial,omitempty"`
											Shareable *bool   `tfsdk:"shareable" json:"shareable,omitempty"`
											Tag       *string `tfsdk:"tag" json:"tag,omitempty"`
										} `tfsdk:"disks" json:"disks,omitempty"`
										DownwardMetrics *map[string]string `tfsdk:"downward_metrics" json:"downwardMetrics,omitempty"`
										Filesystems     *[]struct {
											Name     *string            `tfsdk:"name" json:"name,omitempty"`
											Virtiofs *map[string]string `tfsdk:"virtiofs" json:"virtiofs,omitempty"`
										} `tfsdk:"filesystems" json:"filesystems,omitempty"`
										Gpus *[]struct {
											DeviceName        *string `tfsdk:"device_name" json:"deviceName,omitempty"`
											Name              *string `tfsdk:"name" json:"name,omitempty"`
											Tag               *string `tfsdk:"tag" json:"tag,omitempty"`
											VirtualGPUOptions *struct {
												Display *struct {
													Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
													RamFB   *struct {
														Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
													} `tfsdk:"ram_fb" json:"ramFB,omitempty"`
												} `tfsdk:"display" json:"display,omitempty"`
											} `tfsdk:"virtual_gpu_options" json:"virtualGPUOptions,omitempty"`
										} `tfsdk:"gpus" json:"gpus,omitempty"`
										HostDevices *[]struct {
											DeviceName *string `tfsdk:"device_name" json:"deviceName,omitempty"`
											Name       *string `tfsdk:"name" json:"name,omitempty"`
											Tag        *string `tfsdk:"tag" json:"tag,omitempty"`
										} `tfsdk:"host_devices" json:"hostDevices,omitempty"`
										Inputs *[]struct {
											Bus  *string `tfsdk:"bus" json:"bus,omitempty"`
											Name *string `tfsdk:"name" json:"name,omitempty"`
											Type *string `tfsdk:"type" json:"type,omitempty"`
										} `tfsdk:"inputs" json:"inputs,omitempty"`
										Interfaces *[]struct {
											AcpiIndex *int64 `tfsdk:"acpi_index" json:"acpiIndex,omitempty"`
											Binding   *struct {
												Name *string `tfsdk:"name" json:"name,omitempty"`
											} `tfsdk:"binding" json:"binding,omitempty"`
											BootOrder   *int64             `tfsdk:"boot_order" json:"bootOrder,omitempty"`
											Bridge      *map[string]string `tfsdk:"bridge" json:"bridge,omitempty"`
											DhcpOptions *struct {
												BootFileName   *string   `tfsdk:"boot_file_name" json:"bootFileName,omitempty"`
												NtpServers     *[]string `tfsdk:"ntp_servers" json:"ntpServers,omitempty"`
												PrivateOptions *[]struct {
													Option *int64  `tfsdk:"option" json:"option,omitempty"`
													Value  *string `tfsdk:"value" json:"value,omitempty"`
												} `tfsdk:"private_options" json:"privateOptions,omitempty"`
												TftpServerName *string `tfsdk:"tftp_server_name" json:"tftpServerName,omitempty"`
											} `tfsdk:"dhcp_options" json:"dhcpOptions,omitempty"`
											MacAddress *string            `tfsdk:"mac_address" json:"macAddress,omitempty"`
											Macvtap    *map[string]string `tfsdk:"macvtap" json:"macvtap,omitempty"`
											Masquerade *map[string]string `tfsdk:"masquerade" json:"masquerade,omitempty"`
											Model      *string            `tfsdk:"model" json:"model,omitempty"`
											Name       *string            `tfsdk:"name" json:"name,omitempty"`
											Passt      *map[string]string `tfsdk:"passt" json:"passt,omitempty"`
											PciAddress *string            `tfsdk:"pci_address" json:"pciAddress,omitempty"`
											Ports      *[]struct {
												Name     *string `tfsdk:"name" json:"name,omitempty"`
												Port     *int64  `tfsdk:"port" json:"port,omitempty"`
												Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
											} `tfsdk:"ports" json:"ports,omitempty"`
											Slirp *map[string]string `tfsdk:"slirp" json:"slirp,omitempty"`
											Sriov *map[string]string `tfsdk:"sriov" json:"sriov,omitempty"`
											State *string            `tfsdk:"state" json:"state,omitempty"`
											Tag   *string            `tfsdk:"tag" json:"tag,omitempty"`
										} `tfsdk:"interfaces" json:"interfaces,omitempty"`
										LogSerialConsole           *bool              `tfsdk:"log_serial_console" json:"logSerialConsole,omitempty"`
										NetworkInterfaceMultiqueue *bool              `tfsdk:"network_interface_multiqueue" json:"networkInterfaceMultiqueue,omitempty"`
										Rng                        *map[string]string `tfsdk:"rng" json:"rng,omitempty"`
										Sound                      *struct {
											Model *string `tfsdk:"model" json:"model,omitempty"`
											Name  *string `tfsdk:"name" json:"name,omitempty"`
										} `tfsdk:"sound" json:"sound,omitempty"`
										Tpm *struct {
											Persistent *bool `tfsdk:"persistent" json:"persistent,omitempty"`
										} `tfsdk:"tpm" json:"tpm,omitempty"`
										UseVirtioTransitional *bool `tfsdk:"use_virtio_transitional" json:"useVirtioTransitional,omitempty"`
										Watchdog              *struct {
											I6300esb *struct {
												Action *string `tfsdk:"action" json:"action,omitempty"`
											} `tfsdk:"i6300esb" json:"i6300esb,omitempty"`
											Name *string `tfsdk:"name" json:"name,omitempty"`
										} `tfsdk:"watchdog" json:"watchdog,omitempty"`
									} `tfsdk:"devices" json:"devices,omitempty"`
									Features *struct {
										Acpi *struct {
											Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
										} `tfsdk:"acpi" json:"acpi,omitempty"`
										Apic *struct {
											Enabled        *bool `tfsdk:"enabled" json:"enabled,omitempty"`
											EndOfInterrupt *bool `tfsdk:"end_of_interrupt" json:"endOfInterrupt,omitempty"`
										} `tfsdk:"apic" json:"apic,omitempty"`
										Hyperv *struct {
											Evmcs *struct {
												Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
											} `tfsdk:"evmcs" json:"evmcs,omitempty"`
											Frequencies *struct {
												Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
											} `tfsdk:"frequencies" json:"frequencies,omitempty"`
											Ipi *struct {
												Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
											} `tfsdk:"ipi" json:"ipi,omitempty"`
											Reenlightenment *struct {
												Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
											} `tfsdk:"reenlightenment" json:"reenlightenment,omitempty"`
											Relaxed *struct {
												Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
											} `tfsdk:"relaxed" json:"relaxed,omitempty"`
											Reset *struct {
												Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
											} `tfsdk:"reset" json:"reset,omitempty"`
											Runtime *struct {
												Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
											} `tfsdk:"runtime" json:"runtime,omitempty"`
											Spinlocks *struct {
												Enabled   *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
												Spinlocks *int64 `tfsdk:"spinlocks" json:"spinlocks,omitempty"`
											} `tfsdk:"spinlocks" json:"spinlocks,omitempty"`
											Synic *struct {
												Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
											} `tfsdk:"synic" json:"synic,omitempty"`
											Synictimer *struct {
												Direct *struct {
													Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
												} `tfsdk:"direct" json:"direct,omitempty"`
												Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
											} `tfsdk:"synictimer" json:"synictimer,omitempty"`
											Tlbflush *struct {
												Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
											} `tfsdk:"tlbflush" json:"tlbflush,omitempty"`
											Vapic *struct {
												Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
											} `tfsdk:"vapic" json:"vapic,omitempty"`
											Vendorid *struct {
												Enabled  *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
												Vendorid *string `tfsdk:"vendorid" json:"vendorid,omitempty"`
											} `tfsdk:"vendorid" json:"vendorid,omitempty"`
											Vpindex *struct {
												Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
											} `tfsdk:"vpindex" json:"vpindex,omitempty"`
										} `tfsdk:"hyperv" json:"hyperv,omitempty"`
										Kvm *struct {
											Hidden *bool `tfsdk:"hidden" json:"hidden,omitempty"`
										} `tfsdk:"kvm" json:"kvm,omitempty"`
										Pvspinlock *struct {
											Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
										} `tfsdk:"pvspinlock" json:"pvspinlock,omitempty"`
										Smm *struct {
											Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
										} `tfsdk:"smm" json:"smm,omitempty"`
									} `tfsdk:"features" json:"features,omitempty"`
									Firmware *struct {
										Acpi *struct {
											SlicNameRef *string `tfsdk:"slic_name_ref" json:"slicNameRef,omitempty"`
										} `tfsdk:"acpi" json:"acpi,omitempty"`
										Bootloader *struct {
											Bios *struct {
												UseSerial *bool `tfsdk:"use_serial" json:"useSerial,omitempty"`
											} `tfsdk:"bios" json:"bios,omitempty"`
											Efi *struct {
												Persistent *bool `tfsdk:"persistent" json:"persistent,omitempty"`
												SecureBoot *bool `tfsdk:"secure_boot" json:"secureBoot,omitempty"`
											} `tfsdk:"efi" json:"efi,omitempty"`
										} `tfsdk:"bootloader" json:"bootloader,omitempty"`
										KernelBoot *struct {
											Container *struct {
												Image           *string `tfsdk:"image" json:"image,omitempty"`
												ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
												ImagePullSecret *string `tfsdk:"image_pull_secret" json:"imagePullSecret,omitempty"`
												InitrdPath      *string `tfsdk:"initrd_path" json:"initrdPath,omitempty"`
												KernelPath      *string `tfsdk:"kernel_path" json:"kernelPath,omitempty"`
											} `tfsdk:"container" json:"container,omitempty"`
											KernelArgs *string `tfsdk:"kernel_args" json:"kernelArgs,omitempty"`
										} `tfsdk:"kernel_boot" json:"kernelBoot,omitempty"`
										Serial *string `tfsdk:"serial" json:"serial,omitempty"`
										Uuid   *string `tfsdk:"uuid" json:"uuid,omitempty"`
									} `tfsdk:"firmware" json:"firmware,omitempty"`
									IoThreadsPolicy *string `tfsdk:"io_threads_policy" json:"ioThreadsPolicy,omitempty"`
									LaunchSecurity  *struct {
										Sev *struct {
											Attestation *map[string]string `tfsdk:"attestation" json:"attestation,omitempty"`
											DhCert      *string            `tfsdk:"dh_cert" json:"dhCert,omitempty"`
											Policy      *struct {
												EncryptedState *bool `tfsdk:"encrypted_state" json:"encryptedState,omitempty"`
											} `tfsdk:"policy" json:"policy,omitempty"`
											Session *string `tfsdk:"session" json:"session,omitempty"`
										} `tfsdk:"sev" json:"sev,omitempty"`
									} `tfsdk:"launch_security" json:"launchSecurity,omitempty"`
									Machine *struct {
										Type *string `tfsdk:"type" json:"type,omitempty"`
									} `tfsdk:"machine" json:"machine,omitempty"`
									Memory *struct {
										Guest     *string `tfsdk:"guest" json:"guest,omitempty"`
										Hugepages *struct {
											PageSize *string `tfsdk:"page_size" json:"pageSize,omitempty"`
										} `tfsdk:"hugepages" json:"hugepages,omitempty"`
										MaxGuest *string `tfsdk:"max_guest" json:"maxGuest,omitempty"`
									} `tfsdk:"memory" json:"memory,omitempty"`
									Resources *struct {
										Limits                  *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
										OvercommitGuestOverhead *bool              `tfsdk:"overcommit_guest_overhead" json:"overcommitGuestOverhead,omitempty"`
										Requests                *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
									} `tfsdk:"resources" json:"resources,omitempty"`
								} `tfsdk:"domain" json:"domain,omitempty"`
								EvictionStrategy *string `tfsdk:"eviction_strategy" json:"evictionStrategy,omitempty"`
								Hostname         *string `tfsdk:"hostname" json:"hostname,omitempty"`
								LivenessProbe    *struct {
									Exec *struct {
										Command *[]string `tfsdk:"command" json:"command,omitempty"`
									} `tfsdk:"exec" json:"exec,omitempty"`
									FailureThreshold *int64             `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
									GuestAgentPing   *map[string]string `tfsdk:"guest_agent_ping" json:"guestAgentPing,omitempty"`
									HttpGet          *struct {
										Host        *string `tfsdk:"host" json:"host,omitempty"`
										HttpHeaders *[]struct {
											Name  *string `tfsdk:"name" json:"name,omitempty"`
											Value *string `tfsdk:"value" json:"value,omitempty"`
										} `tfsdk:"http_headers" json:"httpHeaders,omitempty"`
										Path   *string `tfsdk:"path" json:"path,omitempty"`
										Port   *string `tfsdk:"port" json:"port,omitempty"`
										Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
									} `tfsdk:"http_get" json:"httpGet,omitempty"`
									InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
									PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
									SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
									TcpSocket           *struct {
										Host *string `tfsdk:"host" json:"host,omitempty"`
										Port *string `tfsdk:"port" json:"port,omitempty"`
									} `tfsdk:"tcp_socket" json:"tcpSocket,omitempty"`
									TimeoutSeconds *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
								} `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
								Networks *[]struct {
									Multus *struct {
										Default     *bool   `tfsdk:"default" json:"default,omitempty"`
										NetworkName *string `tfsdk:"network_name" json:"networkName,omitempty"`
									} `tfsdk:"multus" json:"multus,omitempty"`
									Name *string `tfsdk:"name" json:"name,omitempty"`
									Pod  *struct {
										VmIPv6NetworkCIDR *string `tfsdk:"vm_i_pv6_network_cidr" json:"vmIPv6NetworkCIDR,omitempty"`
										VmNetworkCIDR     *string `tfsdk:"vm_network_cidr" json:"vmNetworkCIDR,omitempty"`
									} `tfsdk:"pod" json:"pod,omitempty"`
								} `tfsdk:"networks" json:"networks,omitempty"`
								NodeSelector      *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
								PriorityClassName *string            `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
								ReadinessProbe    *struct {
									Exec *struct {
										Command *[]string `tfsdk:"command" json:"command,omitempty"`
									} `tfsdk:"exec" json:"exec,omitempty"`
									FailureThreshold *int64             `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
									GuestAgentPing   *map[string]string `tfsdk:"guest_agent_ping" json:"guestAgentPing,omitempty"`
									HttpGet          *struct {
										Host        *string `tfsdk:"host" json:"host,omitempty"`
										HttpHeaders *[]struct {
											Name  *string `tfsdk:"name" json:"name,omitempty"`
											Value *string `tfsdk:"value" json:"value,omitempty"`
										} `tfsdk:"http_headers" json:"httpHeaders,omitempty"`
										Path   *string `tfsdk:"path" json:"path,omitempty"`
										Port   *string `tfsdk:"port" json:"port,omitempty"`
										Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
									} `tfsdk:"http_get" json:"httpGet,omitempty"`
									InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
									PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
									SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
									TcpSocket           *struct {
										Host *string `tfsdk:"host" json:"host,omitempty"`
										Port *string `tfsdk:"port" json:"port,omitempty"`
									} `tfsdk:"tcp_socket" json:"tcpSocket,omitempty"`
									TimeoutSeconds *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
								} `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
								SchedulerName                 *string `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
								StartStrategy                 *string `tfsdk:"start_strategy" json:"startStrategy,omitempty"`
								Subdomain                     *string `tfsdk:"subdomain" json:"subdomain,omitempty"`
								TerminationGracePeriodSeconds *int64  `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
								Tolerations                   *[]struct {
									Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
									Key               *string `tfsdk:"key" json:"key,omitempty"`
									Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
									TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
									Value             *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"tolerations" json:"tolerations,omitempty"`
								TopologySpreadConstraints *[]struct {
									LabelSelector *struct {
										MatchExpressions *[]struct {
											Key      *string   `tfsdk:"key" json:"key,omitempty"`
											Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
											Values   *[]string `tfsdk:"values" json:"values,omitempty"`
										} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
										MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
									MatchLabelKeys     *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
									MaxSkew            *int64    `tfsdk:"max_skew" json:"maxSkew,omitempty"`
									MinDomains         *int64    `tfsdk:"min_domains" json:"minDomains,omitempty"`
									NodeAffinityPolicy *string   `tfsdk:"node_affinity_policy" json:"nodeAffinityPolicy,omitempty"`
									NodeTaintsPolicy   *string   `tfsdk:"node_taints_policy" json:"nodeTaintsPolicy,omitempty"`
									TopologyKey        *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
									WhenUnsatisfiable  *string   `tfsdk:"when_unsatisfiable" json:"whenUnsatisfiable,omitempty"`
								} `tfsdk:"topology_spread_constraints" json:"topologySpreadConstraints,omitempty"`
								Volumes *[]struct {
									CloudInitConfigDrive *struct {
										NetworkData          *string `tfsdk:"network_data" json:"networkData,omitempty"`
										NetworkDataBase64    *string `tfsdk:"network_data_base64" json:"networkDataBase64,omitempty"`
										NetworkDataSecretRef *struct {
											Name *string `tfsdk:"name" json:"name,omitempty"`
										} `tfsdk:"network_data_secret_ref" json:"networkDataSecretRef,omitempty"`
										SecretRef *struct {
											Name *string `tfsdk:"name" json:"name,omitempty"`
										} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
										UserData       *string `tfsdk:"user_data" json:"userData,omitempty"`
										UserDataBase64 *string `tfsdk:"user_data_base64" json:"userDataBase64,omitempty"`
									} `tfsdk:"cloud_init_config_drive" json:"cloudInitConfigDrive,omitempty"`
									CloudInitNoCloud *struct {
										NetworkData          *string `tfsdk:"network_data" json:"networkData,omitempty"`
										NetworkDataBase64    *string `tfsdk:"network_data_base64" json:"networkDataBase64,omitempty"`
										NetworkDataSecretRef *struct {
											Name *string `tfsdk:"name" json:"name,omitempty"`
										} `tfsdk:"network_data_secret_ref" json:"networkDataSecretRef,omitempty"`
										SecretRef *struct {
											Name *string `tfsdk:"name" json:"name,omitempty"`
										} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
										UserData       *string `tfsdk:"user_data" json:"userData,omitempty"`
										UserDataBase64 *string `tfsdk:"user_data_base64" json:"userDataBase64,omitempty"`
									} `tfsdk:"cloud_init_no_cloud" json:"cloudInitNoCloud,omitempty"`
									ConfigMap *struct {
										Name        *string `tfsdk:"name" json:"name,omitempty"`
										Optional    *bool   `tfsdk:"optional" json:"optional,omitempty"`
										VolumeLabel *string `tfsdk:"volume_label" json:"volumeLabel,omitempty"`
									} `tfsdk:"config_map" json:"configMap,omitempty"`
									ContainerDisk *struct {
										Image           *string `tfsdk:"image" json:"image,omitempty"`
										ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
										ImagePullSecret *string `tfsdk:"image_pull_secret" json:"imagePullSecret,omitempty"`
										Path            *string `tfsdk:"path" json:"path,omitempty"`
									} `tfsdk:"container_disk" json:"containerDisk,omitempty"`
									DataVolume *struct {
										Hotpluggable *bool   `tfsdk:"hotpluggable" json:"hotpluggable,omitempty"`
										Name         *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"data_volume" json:"dataVolume,omitempty"`
									DownwardAPI *struct {
										Fields *[]struct {
											FieldRef *struct {
												ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
												FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
											} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
											Mode             *int64  `tfsdk:"mode" json:"mode,omitempty"`
											Path             *string `tfsdk:"path" json:"path,omitempty"`
											ResourceFieldRef *struct {
												ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
												Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
												Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
											} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
										} `tfsdk:"fields" json:"fields,omitempty"`
										VolumeLabel *string `tfsdk:"volume_label" json:"volumeLabel,omitempty"`
									} `tfsdk:"downward_api" json:"downwardAPI,omitempty"`
									DownwardMetrics *map[string]string `tfsdk:"downward_metrics" json:"downwardMetrics,omitempty"`
									EmptyDisk       *struct {
										Capacity *string `tfsdk:"capacity" json:"capacity,omitempty"`
									} `tfsdk:"empty_disk" json:"emptyDisk,omitempty"`
									Ephemeral *struct {
										PersistentVolumeClaim *struct {
											ClaimName *string `tfsdk:"claim_name" json:"claimName,omitempty"`
											ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
										} `tfsdk:"persistent_volume_claim" json:"persistentVolumeClaim,omitempty"`
									} `tfsdk:"ephemeral" json:"ephemeral,omitempty"`
									HostDisk *struct {
										Capacity *string `tfsdk:"capacity" json:"capacity,omitempty"`
										Path     *string `tfsdk:"path" json:"path,omitempty"`
										Shared   *bool   `tfsdk:"shared" json:"shared,omitempty"`
										Type     *string `tfsdk:"type" json:"type,omitempty"`
									} `tfsdk:"host_disk" json:"hostDisk,omitempty"`
									MemoryDump *struct {
										ClaimName    *string `tfsdk:"claim_name" json:"claimName,omitempty"`
										Hotpluggable *bool   `tfsdk:"hotpluggable" json:"hotpluggable,omitempty"`
										ReadOnly     *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
									} `tfsdk:"memory_dump" json:"memoryDump,omitempty"`
									Name                  *string `tfsdk:"name" json:"name,omitempty"`
									PersistentVolumeClaim *struct {
										ClaimName    *string `tfsdk:"claim_name" json:"claimName,omitempty"`
										Hotpluggable *bool   `tfsdk:"hotpluggable" json:"hotpluggable,omitempty"`
										ReadOnly     *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
									} `tfsdk:"persistent_volume_claim" json:"persistentVolumeClaim,omitempty"`
									Secret *struct {
										Optional    *bool   `tfsdk:"optional" json:"optional,omitempty"`
										SecretName  *string `tfsdk:"secret_name" json:"secretName,omitempty"`
										VolumeLabel *string `tfsdk:"volume_label" json:"volumeLabel,omitempty"`
									} `tfsdk:"secret" json:"secret,omitempty"`
									ServiceAccount *struct {
										ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
									} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
									Sysprep *struct {
										ConfigMap *struct {
											Name *string `tfsdk:"name" json:"name,omitempty"`
										} `tfsdk:"config_map" json:"configMap,omitempty"`
										Secret *struct {
											Name *string `tfsdk:"name" json:"name,omitempty"`
										} `tfsdk:"secret" json:"secret,omitempty"`
									} `tfsdk:"sysprep" json:"sysprep,omitempty"`
								} `tfsdk:"volumes" json:"volumes,omitempty"`
							} `tfsdk:"spec" json:"spec,omitempty"`
						} `tfsdk:"template" json:"template,omitempty"`
					} `tfsdk:"spec" json:"spec,omitempty"`
				} `tfsdk:"virtual_machine_template" json:"virtualMachineTemplate,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoKubevirtMachineTemplateV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_kubevirt_machine_template_v1alpha1_manifest"
}

func (r *InfrastructureClusterXK8SIoKubevirtMachineTemplateV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "KubevirtMachineTemplate is the Schema for the kubevirtmachinetemplates API.",
		MarkdownDescription: "KubevirtMachineTemplate is the Schema for the kubevirtmachinetemplates API.",
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
				Description:         "KubevirtMachineTemplateSpec defines the desired state of KubevirtMachineTemplate.",
				MarkdownDescription: "KubevirtMachineTemplateSpec defines the desired state of KubevirtMachineTemplate.",
				Attributes: map[string]schema.Attribute{
					"template": schema.SingleNestedAttribute{
						Description:         "KubevirtMachineTemplateResource describes the data needed to create a KubevirtMachine from a template.",
						MarkdownDescription: "KubevirtMachineTemplateResource describes the data needed to create a KubevirtMachine from a template.",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "Spec is the specification of the desired behavior of the machine.",
								MarkdownDescription: "Spec is the specification of the desired behavior of the machine.",
								Attributes: map[string]schema.Attribute{
									"infra_cluster_secret_ref": schema.SingleNestedAttribute{
										Description:         "InfraClusterSecretRef is a reference to a secret with a kubeconfig for external cluster used for infra.When nil, this defaults to the value present in the KubevirtCluster object's spec associated with this machine.",
										MarkdownDescription: "InfraClusterSecretRef is a reference to a secret with a kubeconfig for external cluster used for infra.When nil, this defaults to the value present in the KubevirtCluster object's spec associated with this machine.",
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

									"provider_id": schema.StringAttribute{
										Description:         "ProviderID TBD what to use for Kubevirt",
										MarkdownDescription: "ProviderID TBD what to use for Kubevirt",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"virtual_machine_bootstrap_check": schema.SingleNestedAttribute{
										Description:         "BootstrapCheckSpec defines how the CAPK controller is checking CAPI Sentinel file inside the VM.",
										MarkdownDescription: "BootstrapCheckSpec defines how the CAPK controller is checking CAPI Sentinel file inside the VM.",
										Attributes: map[string]schema.Attribute{
											"check_strategy": schema.StringAttribute{
												Description:         "CheckStrategy describes how CAPK controller will validate a successful CAPI bootstrap.Following specified method, CAPK will try to retrieve the state of the CAPI Sentinel file from the VM.Possible values are: 'none' or 'ssh' (default is 'ssh') and this value is validated by apiserver.",
												MarkdownDescription: "CheckStrategy describes how CAPK controller will validate a successful CAPI bootstrap.Following specified method, CAPK will try to retrieve the state of the CAPI Sentinel file from the VM.Possible values are: 'none' or 'ssh' (default is 'ssh') and this value is validated by apiserver.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("none", "ssh"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"virtual_machine_template": schema.SingleNestedAttribute{
										Description:         "VirtualMachineTemplateSpec defines the desired state of the kubevirt VM.",
										MarkdownDescription: "VirtualMachineTemplateSpec defines the desired state of the kubevirt VM.",
										Attributes: map[string]schema.Attribute{
											"metadata": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"spec": schema.SingleNestedAttribute{
												Description:         "VirtualMachineSpec contains the VirtualMachine specification.",
												MarkdownDescription: "VirtualMachineSpec contains the VirtualMachine specification.",
												Attributes: map[string]schema.Attribute{
													"data_volume_templates": schema.ListNestedAttribute{
														Description:         "dataVolumeTemplates is a list of dataVolumes that the VirtualMachineInstance template can reference.DataVolumes in this list are dynamically created for the VirtualMachine and are tied to the VirtualMachine's life-cycle.",
														MarkdownDescription: "dataVolumeTemplates is a list of dataVolumes that the VirtualMachineInstance template can reference.DataVolumes in this list are dynamically created for the VirtualMachine and are tied to the VirtualMachine's life-cycle.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"api_version": schema.StringAttribute{
																	Description:         "APIVersion defines the versioned schema of this representation of an object.Servers should convert recognized schemas to the latest internal value, andmay reject unrecognized values.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
																	MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object.Servers should convert recognized schemas to the latest internal value, andmay reject unrecognized values.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"kind": schema.StringAttribute{
																	Description:         "Kind is a string value representing the REST resource this object represents.Servers may infer this from the endpoint the client submits requests to.Cannot be updated.In CamelCase.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																	MarkdownDescription: "Kind is a string value representing the REST resource this object represents.Servers may infer this from the endpoint the client submits requests to.Cannot be updated.In CamelCase.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"metadata": schema.MapAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"spec": schema.SingleNestedAttribute{
																	Description:         "DataVolumeSpec contains the DataVolume specification.",
																	MarkdownDescription: "DataVolumeSpec contains the DataVolume specification.",
																	Attributes: map[string]schema.Attribute{
																		"checkpoints": schema.ListNestedAttribute{
																			Description:         "Checkpoints is a list of DataVolumeCheckpoints, representing stages in a multistage import.",
																			MarkdownDescription: "Checkpoints is a list of DataVolumeCheckpoints, representing stages in a multistage import.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"current": schema.StringAttribute{
																						Description:         "Current is the identifier of the snapshot created for this checkpoint.",
																						MarkdownDescription: "Current is the identifier of the snapshot created for this checkpoint.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"previous": schema.StringAttribute{
																						Description:         "Previous is the identifier of the snapshot from the previous checkpoint.",
																						MarkdownDescription: "Previous is the identifier of the snapshot from the previous checkpoint.",
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

																		"content_type": schema.StringAttribute{
																			Description:         "DataVolumeContentType options: 'kubevirt', 'archive'",
																			MarkdownDescription: "DataVolumeContentType options: 'kubevirt', 'archive'",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("kubevirt", "archive"),
																			},
																		},

																		"final_checkpoint": schema.BoolAttribute{
																			Description:         "FinalCheckpoint indicates whether the current DataVolumeCheckpoint is the final checkpoint.",
																			MarkdownDescription: "FinalCheckpoint indicates whether the current DataVolumeCheckpoint is the final checkpoint.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"preallocation": schema.BoolAttribute{
																			Description:         "Preallocation controls whether storage for DataVolumes should be allocated in advance.",
																			MarkdownDescription: "Preallocation controls whether storage for DataVolumes should be allocated in advance.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"priority_class_name": schema.StringAttribute{
																			Description:         "PriorityClassName for Importer, Cloner and Uploader pod",
																			MarkdownDescription: "PriorityClassName for Importer, Cloner and Uploader pod",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"pvc": schema.SingleNestedAttribute{
																			Description:         "PVC is the PVC specification",
																			MarkdownDescription: "PVC is the PVC specification",
																			Attributes: map[string]schema.Attribute{
																				"access_modes": schema.ListAttribute{
																					Description:         "accessModes contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
																					MarkdownDescription: "accessModes contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"data_source": schema.SingleNestedAttribute{
																					Description:         "dataSource field can be used to specify either:* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)* An existing PVC (PersistentVolumeClaim)If the provisioner or an external controller can support the specified data source,it will create a new volume based on the contents of the specified data source.When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
																					MarkdownDescription: "dataSource field can be used to specify either:* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)* An existing PVC (PersistentVolumeClaim)If the provisioner or an external controller can support the specified data source,it will create a new volume based on the contents of the specified data source.When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
																					Attributes: map[string]schema.Attribute{
																						"api_group": schema.StringAttribute{
																							Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
																							MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
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

																				"data_source_ref": schema.SingleNestedAttribute{
																					Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-emptyvolume is desired. This may be any object from a non-empty API group (noncore object) or a PersistentVolumeClaim object.When this field is specified, volume binding will only succeed if the type ofthe specified object matches some installed volume populator or dynamicprovisioner.This field will replace the functionality of the dataSource field and as suchif both fields are non-empty, they must have the same value. For backwardscompatibility, when namespace isn't specified in dataSourceRef,both fields (dataSource and dataSourceRef) will be set to the samevalue automatically if one of them is empty and the other is non-empty.When namespace is specified in dataSourceRef,dataSource isn't set to the same value and must be empty.There are three important differences between dataSource and dataSourceRef:* While dataSource only allows two specific types of objects, dataSourceRef  allows any non-core object, as well as PersistentVolumeClaim objects.* While dataSource ignores disallowed values (dropping them), dataSourceRef  preserves all values, and generates an error if a disallowed value is  specified.* While dataSource only allows local objects, dataSourceRef allows objects  in any namespaces.(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																					MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-emptyvolume is desired. This may be any object from a non-empty API group (noncore object) or a PersistentVolumeClaim object.When this field is specified, volume binding will only succeed if the type ofthe specified object matches some installed volume populator or dynamicprovisioner.This field will replace the functionality of the dataSource field and as suchif both fields are non-empty, they must have the same value. For backwardscompatibility, when namespace isn't specified in dataSourceRef,both fields (dataSource and dataSourceRef) will be set to the samevalue automatically if one of them is empty and the other is non-empty.When namespace is specified in dataSourceRef,dataSource isn't set to the same value and must be empty.There are three important differences between dataSource and dataSourceRef:* While dataSource only allows two specific types of objects, dataSourceRef  allows any non-core object, as well as PersistentVolumeClaim objects.* While dataSource ignores disallowed values (dropping them), dataSourceRef  preserves all values, and generates an error if a disallowed value is  specified.* While dataSource only allows local objects, dataSourceRef allows objects  in any namespaces.(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																					Attributes: map[string]schema.Attribute{
																						"api_group": schema.StringAttribute{
																							Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
																							MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
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

																						"namespace": schema.StringAttribute{
																							Description:         "Namespace is the namespace of resource being referencedNote that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																							MarkdownDescription: "Namespace is the namespace of resource being referencedNote that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
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
																					Description:         "resources represents the minimum resources the volume should have.If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirementsthat are lower than previous value but must still be higher than capacity recorded in thestatus field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
																					MarkdownDescription: "resources represents the minimum resources the volume should have.If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirementsthat are lower than previous value but must still be higher than capacity recorded in thestatus field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
																					Attributes: map[string]schema.Attribute{
																						"limits": schema.MapAttribute{
																							Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																							MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"requests": schema.MapAttribute{
																							Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																							MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

																				"selector": schema.SingleNestedAttribute{
																					Description:         "selector is a label query over volumes to consider for binding.",
																					MarkdownDescription: "selector is a label query over volumes to consider for binding.",
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
																										Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																										MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"values": schema.ListAttribute{
																										Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																										MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																							Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																				"storage_class_name": schema.StringAttribute{
																					Description:         "storageClassName is the name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																					MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"volume_attributes_class_name": schema.StringAttribute{
																					Description:         "volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim.If specified, the CSI driver will create or update the volume with the attributes definedin the corresponding VolumeAttributesClass. This has a different purpose than storageClassName,it can be changed after the claim is created. An empty string value means that no VolumeAttributesClasswill be applied to the claim but it's not allowed to reset this field to empty string once it is set.If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClasswill be set by the persistentvolume controller if it exists.If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will beset to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resourceexists.More info: https://kubernetes.io/docs/concepts/storage/volume-attributes-classes/(Alpha) Using this field requires the VolumeAttributesClass feature gate to be enabled.",
																					MarkdownDescription: "volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim.If specified, the CSI driver will create or update the volume with the attributes definedin the corresponding VolumeAttributesClass. This has a different purpose than storageClassName,it can be changed after the claim is created. An empty string value means that no VolumeAttributesClasswill be applied to the claim but it's not allowed to reset this field to empty string once it is set.If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClasswill be set by the persistentvolume controller if it exists.If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will beset to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resourceexists.More info: https://kubernetes.io/docs/concepts/storage/volume-attributes-classes/(Alpha) Using this field requires the VolumeAttributesClass feature gate to be enabled.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"volume_mode": schema.StringAttribute{
																					Description:         "volumeMode defines what type of volume is required by the claim.Value of Filesystem is implied when not included in claim spec.",
																					MarkdownDescription: "volumeMode defines what type of volume is required by the claim.Value of Filesystem is implied when not included in claim spec.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"volume_name": schema.StringAttribute{
																					Description:         "volumeName is the binding reference to the PersistentVolume backing this claim.",
																					MarkdownDescription: "volumeName is the binding reference to the PersistentVolume backing this claim.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"source": schema.SingleNestedAttribute{
																			Description:         "Source is the src of the data for the requested DataVolume",
																			MarkdownDescription: "Source is the src of the data for the requested DataVolume",
																			Attributes: map[string]schema.Attribute{
																				"blank": schema.MapAttribute{
																					Description:         "DataVolumeBlankImage provides the parameters to create a new raw blank image for the PVC",
																					MarkdownDescription: "DataVolumeBlankImage provides the parameters to create a new raw blank image for the PVC",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"gcs": schema.SingleNestedAttribute{
																					Description:         "DataVolumeSourceGCS provides the parameters to create a Data Volume from an GCS source",
																					MarkdownDescription: "DataVolumeSourceGCS provides the parameters to create a Data Volume from an GCS source",
																					Attributes: map[string]schema.Attribute{
																						"secret_ref": schema.StringAttribute{
																							Description:         "SecretRef provides the secret reference needed to access the GCS source",
																							MarkdownDescription: "SecretRef provides the secret reference needed to access the GCS source",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"url": schema.StringAttribute{
																							Description:         "URL is the url of the GCS source",
																							MarkdownDescription: "URL is the url of the GCS source",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"http": schema.SingleNestedAttribute{
																					Description:         "DataVolumeSourceHTTP can be either an http or https endpoint, with an optional basic auth user name and password, and an optional configmap containing additional CAs",
																					MarkdownDescription: "DataVolumeSourceHTTP can be either an http or https endpoint, with an optional basic auth user name and password, and an optional configmap containing additional CAs",
																					Attributes: map[string]schema.Attribute{
																						"cert_config_map": schema.StringAttribute{
																							Description:         "CertConfigMap is a configmap reference, containing a Certificate Authority(CA) public key, and a base64 encoded pem certificate",
																							MarkdownDescription: "CertConfigMap is a configmap reference, containing a Certificate Authority(CA) public key, and a base64 encoded pem certificate",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"extra_headers": schema.ListAttribute{
																							Description:         "ExtraHeaders is a list of strings containing extra headers to include with HTTP transfer requests",
																							MarkdownDescription: "ExtraHeaders is a list of strings containing extra headers to include with HTTP transfer requests",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"secret_extra_headers": schema.ListAttribute{
																							Description:         "SecretExtraHeaders is a list of Secret references, each containing an extra HTTP header that may include sensitive information",
																							MarkdownDescription: "SecretExtraHeaders is a list of Secret references, each containing an extra HTTP header that may include sensitive information",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"secret_ref": schema.StringAttribute{
																							Description:         "SecretRef A Secret reference, the secret should contain accessKeyId (user name) base64 encoded, and secretKey (password) also base64 encoded",
																							MarkdownDescription: "SecretRef A Secret reference, the secret should contain accessKeyId (user name) base64 encoded, and secretKey (password) also base64 encoded",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"url": schema.StringAttribute{
																							Description:         "URL is the URL of the http(s) endpoint",
																							MarkdownDescription: "URL is the URL of the http(s) endpoint",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"imageio": schema.SingleNestedAttribute{
																					Description:         "DataVolumeSourceImageIO provides the parameters to create a Data Volume from an imageio source",
																					MarkdownDescription: "DataVolumeSourceImageIO provides the parameters to create a Data Volume from an imageio source",
																					Attributes: map[string]schema.Attribute{
																						"cert_config_map": schema.StringAttribute{
																							Description:         "CertConfigMap provides a reference to the CA cert",
																							MarkdownDescription: "CertConfigMap provides a reference to the CA cert",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"disk_id": schema.StringAttribute{
																							Description:         "DiskID provides id of a disk to be imported",
																							MarkdownDescription: "DiskID provides id of a disk to be imported",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"secret_ref": schema.StringAttribute{
																							Description:         "SecretRef provides the secret reference needed to access the ovirt-engine",
																							MarkdownDescription: "SecretRef provides the secret reference needed to access the ovirt-engine",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"url": schema.StringAttribute{
																							Description:         "URL is the URL of the ovirt-engine",
																							MarkdownDescription: "URL is the URL of the ovirt-engine",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"pvc": schema.SingleNestedAttribute{
																					Description:         "DataVolumeSourcePVC provides the parameters to create a Data Volume from an existing PVC",
																					MarkdownDescription: "DataVolumeSourcePVC provides the parameters to create a Data Volume from an existing PVC",
																					Attributes: map[string]schema.Attribute{
																						"name": schema.StringAttribute{
																							Description:         "The name of the source PVC",
																							MarkdownDescription: "The name of the source PVC",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"namespace": schema.StringAttribute{
																							Description:         "The namespace of the source PVC",
																							MarkdownDescription: "The namespace of the source PVC",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"registry": schema.SingleNestedAttribute{
																					Description:         "DataVolumeSourceRegistry provides the parameters to create a Data Volume from an registry source",
																					MarkdownDescription: "DataVolumeSourceRegistry provides the parameters to create a Data Volume from an registry source",
																					Attributes: map[string]schema.Attribute{
																						"cert_config_map": schema.StringAttribute{
																							Description:         "CertConfigMap provides a reference to the Registry certs",
																							MarkdownDescription: "CertConfigMap provides a reference to the Registry certs",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"image_stream": schema.StringAttribute{
																							Description:         "ImageStream is the name of image stream for import",
																							MarkdownDescription: "ImageStream is the name of image stream for import",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"pull_method": schema.StringAttribute{
																							Description:         "PullMethod can be either 'pod' (default import), or 'node' (node docker cache based import)",
																							MarkdownDescription: "PullMethod can be either 'pod' (default import), or 'node' (node docker cache based import)",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"secret_ref": schema.StringAttribute{
																							Description:         "SecretRef provides the secret reference needed to access the Registry source",
																							MarkdownDescription: "SecretRef provides the secret reference needed to access the Registry source",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"url": schema.StringAttribute{
																							Description:         "URL is the url of the registry source (starting with the scheme: docker, oci-archive)",
																							MarkdownDescription: "URL is the url of the registry source (starting with the scheme: docker, oci-archive)",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"s3": schema.SingleNestedAttribute{
																					Description:         "DataVolumeSourceS3 provides the parameters to create a Data Volume from an S3 source",
																					MarkdownDescription: "DataVolumeSourceS3 provides the parameters to create a Data Volume from an S3 source",
																					Attributes: map[string]schema.Attribute{
																						"cert_config_map": schema.StringAttribute{
																							Description:         "CertConfigMap is a configmap reference, containing a Certificate Authority(CA) public key, and a base64 encoded pem certificate",
																							MarkdownDescription: "CertConfigMap is a configmap reference, containing a Certificate Authority(CA) public key, and a base64 encoded pem certificate",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"secret_ref": schema.StringAttribute{
																							Description:         "SecretRef provides the secret reference needed to access the S3 source",
																							MarkdownDescription: "SecretRef provides the secret reference needed to access the S3 source",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"url": schema.StringAttribute{
																							Description:         "URL is the url of the S3 source",
																							MarkdownDescription: "URL is the url of the S3 source",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"snapshot": schema.SingleNestedAttribute{
																					Description:         "DataVolumeSourceSnapshot provides the parameters to create a Data Volume from an existing VolumeSnapshot",
																					MarkdownDescription: "DataVolumeSourceSnapshot provides the parameters to create a Data Volume from an existing VolumeSnapshot",
																					Attributes: map[string]schema.Attribute{
																						"name": schema.StringAttribute{
																							Description:         "The name of the source VolumeSnapshot",
																							MarkdownDescription: "The name of the source VolumeSnapshot",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"namespace": schema.StringAttribute{
																							Description:         "The namespace of the source VolumeSnapshot",
																							MarkdownDescription: "The namespace of the source VolumeSnapshot",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"upload": schema.MapAttribute{
																					Description:         "DataVolumeSourceUpload provides the parameters to create a Data Volume by uploading the source",
																					MarkdownDescription: "DataVolumeSourceUpload provides the parameters to create a Data Volume by uploading the source",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"vddk": schema.SingleNestedAttribute{
																					Description:         "DataVolumeSourceVDDK provides the parameters to create a Data Volume from a Vmware source",
																					MarkdownDescription: "DataVolumeSourceVDDK provides the parameters to create a Data Volume from a Vmware source",
																					Attributes: map[string]schema.Attribute{
																						"backing_file": schema.StringAttribute{
																							Description:         "BackingFile is the path to the virtual hard disk to migrate from vCenter/ESXi",
																							MarkdownDescription: "BackingFile is the path to the virtual hard disk to migrate from vCenter/ESXi",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"init_image_url": schema.StringAttribute{
																							Description:         "InitImageURL is an optional URL to an image containing an extracted VDDK library, overrides v2v-vmware config map",
																							MarkdownDescription: "InitImageURL is an optional URL to an image containing an extracted VDDK library, overrides v2v-vmware config map",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"secret_ref": schema.StringAttribute{
																							Description:         "SecretRef provides a reference to a secret containing the username and password needed to access the vCenter or ESXi host",
																							MarkdownDescription: "SecretRef provides a reference to a secret containing the username and password needed to access the vCenter or ESXi host",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"thumbprint": schema.StringAttribute{
																							Description:         "Thumbprint is the certificate thumbprint of the vCenter or ESXi host",
																							MarkdownDescription: "Thumbprint is the certificate thumbprint of the vCenter or ESXi host",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"url": schema.StringAttribute{
																							Description:         "URL is the URL of the vCenter or ESXi host with the VM to migrate",
																							MarkdownDescription: "URL is the URL of the vCenter or ESXi host with the VM to migrate",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"uuid": schema.StringAttribute{
																							Description:         "UUID is the UUID of the virtual machine that the backing file is attached to in vCenter/ESXi",
																							MarkdownDescription: "UUID is the UUID of the virtual machine that the backing file is attached to in vCenter/ESXi",
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

																		"source_ref": schema.SingleNestedAttribute{
																			Description:         "SourceRef is an indirect reference to the source of data for the requested DataVolume",
																			MarkdownDescription: "SourceRef is an indirect reference to the source of data for the requested DataVolume",
																			Attributes: map[string]schema.Attribute{
																				"kind": schema.StringAttribute{
																					Description:         "The kind of the source reference, currently only 'DataSource' is supported",
																					MarkdownDescription: "The kind of the source reference, currently only 'DataSource' is supported",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"name": schema.StringAttribute{
																					Description:         "The name of the source reference",
																					MarkdownDescription: "The name of the source reference",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"namespace": schema.StringAttribute{
																					Description:         "The namespace of the source reference, defaults to the DataVolume namespace",
																					MarkdownDescription: "The namespace of the source reference, defaults to the DataVolume namespace",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"storage": schema.SingleNestedAttribute{
																			Description:         "Storage is the requested storage specification",
																			MarkdownDescription: "Storage is the requested storage specification",
																			Attributes: map[string]schema.Attribute{
																				"access_modes": schema.ListAttribute{
																					Description:         "AccessModes contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
																					MarkdownDescription: "AccessModes contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"data_source": schema.SingleNestedAttribute{
																					Description:         "This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) * An existing custom resource that implements data population (Alpha) In order to use custom resource types that implement data population, the AnyVolumeDataSource feature gate must be enabled. If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source.If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",
																					MarkdownDescription: "This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) * An existing custom resource that implements data population (Alpha) In order to use custom resource types that implement data population, the AnyVolumeDataSource feature gate must be enabled. If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source.If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",
																					Attributes: map[string]schema.Attribute{
																						"api_group": schema.StringAttribute{
																							Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
																							MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
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

																				"data_source_ref": schema.SingleNestedAttribute{
																					Description:         "Specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner.This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty.There are two important differences between DataSource and DataSourceRef:* While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects.* While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified.(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
																					MarkdownDescription: "Specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner.This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty.There are two important differences between DataSource and DataSourceRef:* While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects.* While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified.(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
																					Attributes: map[string]schema.Attribute{
																						"api_group": schema.StringAttribute{
																							Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
																							MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
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

																						"namespace": schema.StringAttribute{
																							Description:         "Namespace is the namespace of resource being referencedNote that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																							MarkdownDescription: "Namespace is the namespace of resource being referencedNote that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
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
																					Description:         "Resources represents the minimum resources the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
																					MarkdownDescription: "Resources represents the minimum resources the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
																					Attributes: map[string]schema.Attribute{
																						"claims": schema.ListNestedAttribute{
																							Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																							MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"name": schema.StringAttribute{
																										Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
																										MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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
																							Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																							MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"requests": schema.MapAttribute{
																							Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																							MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

																				"selector": schema.SingleNestedAttribute{
																					Description:         "A label query over volumes to consider for binding.",
																					MarkdownDescription: "A label query over volumes to consider for binding.",
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
																										Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																										MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"values": schema.ListAttribute{
																										Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																										MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																							Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																				"storage_class_name": schema.StringAttribute{
																					Description:         "Name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																					MarkdownDescription: "Name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"volume_mode": schema.StringAttribute{
																					Description:         "volumeMode defines what type of volume is required by the claim.Value of Filesystem is implied when not included in claim spec.",
																					MarkdownDescription: "volumeMode defines what type of volume is required by the claim.Value of Filesystem is implied when not included in claim spec.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"volume_name": schema.StringAttribute{
																					Description:         "VolumeName is the binding reference to the PersistentVolume backing this claim.",
																					MarkdownDescription: "VolumeName is the binding reference to the PersistentVolume backing this claim.",
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

																"status": schema.MapAttribute{
																	Description:         "DataVolumeTemplateDummyStatus is here simply for backwards compatibility witha previous API.",
																	MarkdownDescription: "DataVolumeTemplateDummyStatus is here simply for backwards compatibility witha previous API.",
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

													"instancetype": schema.SingleNestedAttribute{
														Description:         "InstancetypeMatcher references a instancetype that is used to fill fields in Template",
														MarkdownDescription: "InstancetypeMatcher references a instancetype that is used to fill fields in Template",
														Attributes: map[string]schema.Attribute{
															"infer_from_volume": schema.StringAttribute{
																Description:         "InferFromVolume lists the name of a volume that should be used to infer or discover the instancetypeto be used through known annotations on the underlying resource. Once applied to the InstancetypeMatcherthis field is removed.",
																MarkdownDescription: "InferFromVolume lists the name of a volume that should be used to infer or discover the instancetypeto be used through known annotations on the underlying resource. Once applied to the InstancetypeMatcherthis field is removed.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"infer_from_volume_failure_policy": schema.StringAttribute{
																Description:         "InferFromVolumeFailurePolicy controls what should happen on failure when inferring the instancetype.Allowed values are: 'RejectInferFromVolumeFailure' and 'IgnoreInferFromVolumeFailure'.If not specified, 'RejectInferFromVolumeFailure' is used by default.",
																MarkdownDescription: "InferFromVolumeFailurePolicy controls what should happen on failure when inferring the instancetype.Allowed values are: 'RejectInferFromVolumeFailure' and 'IgnoreInferFromVolumeFailure'.If not specified, 'RejectInferFromVolumeFailure' is used by default.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kind": schema.StringAttribute{
																Description:         "Kind specifies which instancetype resource is referenced.Allowed values are: 'VirtualMachineInstancetype' and 'VirtualMachineClusterInstancetype'.If not specified, 'VirtualMachineClusterInstancetype' is used by default.",
																MarkdownDescription: "Kind specifies which instancetype resource is referenced.Allowed values are: 'VirtualMachineInstancetype' and 'VirtualMachineClusterInstancetype'.If not specified, 'VirtualMachineClusterInstancetype' is used by default.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name is the name of the VirtualMachineInstancetype or VirtualMachineClusterInstancetype",
																MarkdownDescription: "Name is the name of the VirtualMachineInstancetype or VirtualMachineClusterInstancetype",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"revision_name": schema.StringAttribute{
																Description:         "RevisionName specifies a ControllerRevision containing a specific copy of theVirtualMachineInstancetype or VirtualMachineClusterInstancetype to be used. This is initiallycaptured the first time the instancetype is applied to the VirtualMachineInstance.",
																MarkdownDescription: "RevisionName specifies a ControllerRevision containing a specific copy of theVirtualMachineInstancetype or VirtualMachineClusterInstancetype to be used. This is initiallycaptured the first time the instancetype is applied to the VirtualMachineInstance.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"preference": schema.SingleNestedAttribute{
														Description:         "PreferenceMatcher references a set of preference that is used to fill fields in Template",
														MarkdownDescription: "PreferenceMatcher references a set of preference that is used to fill fields in Template",
														Attributes: map[string]schema.Attribute{
															"infer_from_volume": schema.StringAttribute{
																Description:         "InferFromVolume lists the name of a volume that should be used to infer or discover the preferenceto be used through known annotations on the underlying resource. Once applied to the PreferenceMatcherthis field is removed.",
																MarkdownDescription: "InferFromVolume lists the name of a volume that should be used to infer or discover the preferenceto be used through known annotations on the underlying resource. Once applied to the PreferenceMatcherthis field is removed.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"infer_from_volume_failure_policy": schema.StringAttribute{
																Description:         "InferFromVolumeFailurePolicy controls what should happen on failure when preference the instancetype.Allowed values are: 'RejectInferFromVolumeFailure' and 'IgnoreInferFromVolumeFailure'.If not specified, 'RejectInferFromVolumeFailure' is used by default.",
																MarkdownDescription: "InferFromVolumeFailurePolicy controls what should happen on failure when preference the instancetype.Allowed values are: 'RejectInferFromVolumeFailure' and 'IgnoreInferFromVolumeFailure'.If not specified, 'RejectInferFromVolumeFailure' is used by default.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kind": schema.StringAttribute{
																Description:         "Kind specifies which preference resource is referenced.Allowed values are: 'VirtualMachinePreference' and 'VirtualMachineClusterPreference'.If not specified, 'VirtualMachineClusterPreference' is used by default.",
																MarkdownDescription: "Kind specifies which preference resource is referenced.Allowed values are: 'VirtualMachinePreference' and 'VirtualMachineClusterPreference'.If not specified, 'VirtualMachineClusterPreference' is used by default.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name is the name of the VirtualMachinePreference or VirtualMachineClusterPreference",
																MarkdownDescription: "Name is the name of the VirtualMachinePreference or VirtualMachineClusterPreference",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"revision_name": schema.StringAttribute{
																Description:         "RevisionName specifies a ControllerRevision containing a specific copy of theVirtualMachinePreference or VirtualMachineClusterPreference to be used. This isinitially captured the first time the instancetype is applied to the VirtualMachineInstance.",
																MarkdownDescription: "RevisionName specifies a ControllerRevision containing a specific copy of theVirtualMachinePreference or VirtualMachineClusterPreference to be used. This isinitially captured the first time the instancetype is applied to the VirtualMachineInstance.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_strategy": schema.StringAttribute{
														Description:         "Running state indicates the requested running state of the VirtualMachineInstancemutually exclusive with Running",
														MarkdownDescription: "Running state indicates the requested running state of the VirtualMachineInstancemutually exclusive with Running",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"running": schema.BoolAttribute{
														Description:         "Running controls whether the associatied VirtualMachineInstance is created or notMutually exclusive with RunStrategy",
														MarkdownDescription: "Running controls whether the associatied VirtualMachineInstance is created or notMutually exclusive with RunStrategy",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"template": schema.SingleNestedAttribute{
														Description:         "Template is the direct specification of VirtualMachineInstance",
														MarkdownDescription: "Template is the direct specification of VirtualMachineInstance",
														Attributes: map[string]schema.Attribute{
															"metadata": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"spec": schema.SingleNestedAttribute{
																Description:         "VirtualMachineInstance Spec contains the VirtualMachineInstance specification.",
																MarkdownDescription: "VirtualMachineInstance Spec contains the VirtualMachineInstance specification.",
																Attributes: map[string]schema.Attribute{
																	"access_credentials": schema.ListNestedAttribute{
																		Description:         "Specifies a set of public keys to inject into the vm guest",
																		MarkdownDescription: "Specifies a set of public keys to inject into the vm guest",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"ssh_public_key": schema.SingleNestedAttribute{
																					Description:         "SSHPublicKey represents the source and method of applying a ssh publickey into a guest virtual machine.",
																					MarkdownDescription: "SSHPublicKey represents the source and method of applying a ssh publickey into a guest virtual machine.",
																					Attributes: map[string]schema.Attribute{
																						"propagation_method": schema.SingleNestedAttribute{
																							Description:         "PropagationMethod represents how the public key is injected into the vm guest.",
																							MarkdownDescription: "PropagationMethod represents how the public key is injected into the vm guest.",
																							Attributes: map[string]schema.Attribute{
																								"config_drive": schema.MapAttribute{
																									Description:         "ConfigDrivePropagation means that the ssh public keys are injectedinto the VM using metadata using the configDrive cloud-init provider",
																									MarkdownDescription: "ConfigDrivePropagation means that the ssh public keys are injectedinto the VM using metadata using the configDrive cloud-init provider",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"no_cloud": schema.MapAttribute{
																									Description:         "NoCloudPropagation means that the ssh public keys are injectedinto the VM using metadata using the noCloud cloud-init provider",
																									MarkdownDescription: "NoCloudPropagation means that the ssh public keys are injectedinto the VM using metadata using the noCloud cloud-init provider",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"qemu_guest_agent": schema.SingleNestedAttribute{
																									Description:         "QemuGuestAgentAccessCredentailPropagation means ssh public keys aredynamically injected into the vm at runtime via the qemu guest agent.This feature requires the qemu guest agent to be running within the guest.",
																									MarkdownDescription: "QemuGuestAgentAccessCredentailPropagation means ssh public keys aredynamically injected into the vm at runtime via the qemu guest agent.This feature requires the qemu guest agent to be running within the guest.",
																									Attributes: map[string]schema.Attribute{
																										"users": schema.ListAttribute{
																											Description:         "Users represents a list of guest users that should have the ssh public keysadded to their authorized_keys file.",
																											MarkdownDescription: "Users represents a list of guest users that should have the ssh public keysadded to their authorized_keys file.",
																											ElementType:         types.StringType,
																											Required:            true,
																											Optional:            false,
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

																						"source": schema.SingleNestedAttribute{
																							Description:         "Source represents where the public keys are pulled from",
																							MarkdownDescription: "Source represents where the public keys are pulled from",
																							Attributes: map[string]schema.Attribute{
																								"secret": schema.SingleNestedAttribute{
																									Description:         "Secret means that the access credential is pulled from a kubernetes secret",
																									MarkdownDescription: "Secret means that the access credential is pulled from a kubernetes secret",
																									Attributes: map[string]schema.Attribute{
																										"secret_name": schema.StringAttribute{
																											Description:         "SecretName represents the name of the secret in the VMI's namespace",
																											MarkdownDescription: "SecretName represents the name of the secret in the VMI's namespace",
																											Required:            true,
																											Optional:            false,
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
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"user_password": schema.SingleNestedAttribute{
																					Description:         "UserPassword represents the source and method for applying a guest user'spassword",
																					MarkdownDescription: "UserPassword represents the source and method for applying a guest user'spassword",
																					Attributes: map[string]schema.Attribute{
																						"propagation_method": schema.SingleNestedAttribute{
																							Description:         "propagationMethod represents how the user passwords are injected into the vm guest.",
																							MarkdownDescription: "propagationMethod represents how the user passwords are injected into the vm guest.",
																							Attributes: map[string]schema.Attribute{
																								"qemu_guest_agent": schema.MapAttribute{
																									Description:         "QemuGuestAgentAccessCredentailPropagation means passwords aredynamically injected into the vm at runtime via the qemu guest agent.This feature requires the qemu guest agent to be running within the guest.",
																									MarkdownDescription: "QemuGuestAgentAccessCredentailPropagation means passwords aredynamically injected into the vm at runtime via the qemu guest agent.This feature requires the qemu guest agent to be running within the guest.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},
																							},
																							Required: true,
																							Optional: false,
																							Computed: false,
																						},

																						"source": schema.SingleNestedAttribute{
																							Description:         "Source represents where the user passwords are pulled from",
																							MarkdownDescription: "Source represents where the user passwords are pulled from",
																							Attributes: map[string]schema.Attribute{
																								"secret": schema.SingleNestedAttribute{
																									Description:         "Secret means that the access credential is pulled from a kubernetes secret",
																									MarkdownDescription: "Secret means that the access credential is pulled from a kubernetes secret",
																									Attributes: map[string]schema.Attribute{
																										"secret_name": schema.StringAttribute{
																											Description:         "SecretName represents the name of the secret in the VMI's namespace",
																											MarkdownDescription: "SecretName represents the name of the secret in the VMI's namespace",
																											Required:            true,
																											Optional:            false,
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

																	"affinity": schema.SingleNestedAttribute{
																		Description:         "If affinity is specifies, obey all the affinity rules",
																		MarkdownDescription: "If affinity is specifies, obey all the affinity rules",
																		Attributes: map[string]schema.Attribute{
																			"node_affinity": schema.SingleNestedAttribute{
																				Description:         "Describes node affinity scheduling rules for the pod.",
																				MarkdownDescription: "Describes node affinity scheduling rules for the pod.",
																				Attributes: map[string]schema.Attribute{
																					"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																						Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node matches the corresponding matchExpressions; thenode(s) with the highest sum are the most preferred.",
																						MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node matches the corresponding matchExpressions; thenode(s) with the highest sum are the most preferred.",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"preference": schema.SingleNestedAttribute{
																									Description:         "A node selector term, associated with the corresponding weight.",
																									MarkdownDescription: "A node selector term, associated with the corresponding weight.",
																									Attributes: map[string]schema.Attribute{
																										"match_expressions": schema.ListNestedAttribute{
																											Description:         "A list of node selector requirements by node's labels.",
																											MarkdownDescription: "A list of node selector requirements by node's labels.",
																											NestedObject: schema.NestedAttributeObject{
																												Attributes: map[string]schema.Attribute{
																													"key": schema.StringAttribute{
																														Description:         "The label key that the selector applies to.",
																														MarkdownDescription: "The label key that the selector applies to.",
																														Required:            true,
																														Optional:            false,
																														Computed:            false,
																													},

																													"operator": schema.StringAttribute{
																														Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																														MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																														Required:            true,
																														Optional:            false,
																														Computed:            false,
																													},

																													"values": schema.ListAttribute{
																														Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																														MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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

																										"match_fields": schema.ListNestedAttribute{
																											Description:         "A list of node selector requirements by node's fields.",
																											MarkdownDescription: "A list of node selector requirements by node's fields.",
																											NestedObject: schema.NestedAttributeObject{
																												Attributes: map[string]schema.Attribute{
																													"key": schema.StringAttribute{
																														Description:         "The label key that the selector applies to.",
																														MarkdownDescription: "The label key that the selector applies to.",
																														Required:            true,
																														Optional:            false,
																														Computed:            false,
																													},

																													"operator": schema.StringAttribute{
																														Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																														MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																														Required:            true,
																														Optional:            false,
																														Computed:            false,
																													},

																													"values": schema.ListAttribute{
																														Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																														MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
																									},
																									Required: true,
																									Optional: false,
																									Computed: false,
																								},

																								"weight": schema.Int64Attribute{
																									Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
																									MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
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

																					"required_during_scheduling_ignored_during_execution": schema.SingleNestedAttribute{
																						Description:         "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to an update), the systemmay or may not try to eventually evict the pod from its node.",
																						MarkdownDescription: "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to an update), the systemmay or may not try to eventually evict the pod from its node.",
																						Attributes: map[string]schema.Attribute{
																							"node_selector_terms": schema.ListNestedAttribute{
																								Description:         "Required. A list of node selector terms. The terms are ORed.",
																								MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"match_expressions": schema.ListNestedAttribute{
																											Description:         "A list of node selector requirements by node's labels.",
																											MarkdownDescription: "A list of node selector requirements by node's labels.",
																											NestedObject: schema.NestedAttributeObject{
																												Attributes: map[string]schema.Attribute{
																													"key": schema.StringAttribute{
																														Description:         "The label key that the selector applies to.",
																														MarkdownDescription: "The label key that the selector applies to.",
																														Required:            true,
																														Optional:            false,
																														Computed:            false,
																													},

																													"operator": schema.StringAttribute{
																														Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																														MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																														Required:            true,
																														Optional:            false,
																														Computed:            false,
																													},

																													"values": schema.ListAttribute{
																														Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																														MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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

																										"match_fields": schema.ListNestedAttribute{
																											Description:         "A list of node selector requirements by node's fields.",
																											MarkdownDescription: "A list of node selector requirements by node's fields.",
																											NestedObject: schema.NestedAttributeObject{
																												Attributes: map[string]schema.Attribute{
																													"key": schema.StringAttribute{
																														Description:         "The label key that the selector applies to.",
																														MarkdownDescription: "The label key that the selector applies to.",
																														Required:            true,
																														Optional:            false,
																														Computed:            false,
																													},

																													"operator": schema.StringAttribute{
																														Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																														MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																														Required:            true,
																														Optional:            false,
																														Computed:            false,
																													},

																													"values": schema.ListAttribute{
																														Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																														MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"pod_affinity": schema.SingleNestedAttribute{
																				Description:         "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
																				MarkdownDescription: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
																				Attributes: map[string]schema.Attribute{
																					"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																						Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
																						MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"pod_affinity_term": schema.SingleNestedAttribute{
																									Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																									MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																									Attributes: map[string]schema.Attribute{
																										"label_selector": schema.SingleNestedAttribute{
																											Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
																											MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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
																																Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																																MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																																Required:            true,
																																Optional:            false,
																																Computed:            false,
																															},

																															"values": schema.ListAttribute{
																																Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																										"match_label_keys": schema.ListAttribute{
																											Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both matchLabelKeys and labelSelector.Also, matchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																											MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both matchLabelKeys and labelSelector.Also, matchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																											ElementType:         types.StringType,
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"mismatch_label_keys": schema.ListAttribute{
																											Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.Also, mismatchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																											MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.Also, mismatchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																											ElementType:         types.StringType,
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"namespace_selector": schema.SingleNestedAttribute{
																											Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
																											MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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
																																Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																																MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																																Required:            true,
																																Optional:            false,
																																Computed:            false,
																															},

																															"values": schema.ListAttribute{
																																Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																										"namespaces": schema.ListAttribute{
																											Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																											MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																											ElementType:         types.StringType,
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"topology_key": schema.StringAttribute{
																											Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																											MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																											Required:            true,
																											Optional:            false,
																											Computed:            false,
																										},
																									},
																									Required: true,
																									Optional: false,
																									Computed: false,
																								},

																								"weight": schema.Int64Attribute{
																									Description:         "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
																									MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
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

																					"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																						Description:         "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
																						MarkdownDescription: "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"label_selector": schema.SingleNestedAttribute{
																									Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
																									MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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
																														Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																														MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																														Required:            true,
																														Optional:            false,
																														Computed:            false,
																													},

																													"values": schema.ListAttribute{
																														Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																											Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																											MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																								"match_label_keys": schema.ListAttribute{
																									Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both matchLabelKeys and labelSelector.Also, matchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																									MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both matchLabelKeys and labelSelector.Also, matchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"mismatch_label_keys": schema.ListAttribute{
																									Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.Also, mismatchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																									MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.Also, mismatchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"namespace_selector": schema.SingleNestedAttribute{
																									Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
																									MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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
																														Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																														MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																														Required:            true,
																														Optional:            false,
																														Computed:            false,
																													},

																													"values": schema.ListAttribute{
																														Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																											Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																											MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																								"namespaces": schema.ListAttribute{
																									Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																									MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"topology_key": schema.StringAttribute{
																									Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																									MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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

																			"pod_anti_affinity": schema.SingleNestedAttribute{
																				Description:         "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
																				MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
																				Attributes: map[string]schema.Attribute{
																					"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																						Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe anti-affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling anti-affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
																						MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe anti-affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling anti-affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"pod_affinity_term": schema.SingleNestedAttribute{
																									Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																									MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																									Attributes: map[string]schema.Attribute{
																										"label_selector": schema.SingleNestedAttribute{
																											Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
																											MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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
																																Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																																MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																																Required:            true,
																																Optional:            false,
																																Computed:            false,
																															},

																															"values": schema.ListAttribute{
																																Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																										"match_label_keys": schema.ListAttribute{
																											Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both matchLabelKeys and labelSelector.Also, matchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																											MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both matchLabelKeys and labelSelector.Also, matchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																											ElementType:         types.StringType,
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"mismatch_label_keys": schema.ListAttribute{
																											Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.Also, mismatchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																											MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.Also, mismatchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																											ElementType:         types.StringType,
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"namespace_selector": schema.SingleNestedAttribute{
																											Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
																											MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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
																																Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																																MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																																Required:            true,
																																Optional:            false,
																																Computed:            false,
																															},

																															"values": schema.ListAttribute{
																																Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																										"namespaces": schema.ListAttribute{
																											Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																											MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																											ElementType:         types.StringType,
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"topology_key": schema.StringAttribute{
																											Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																											MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																											Required:            true,
																											Optional:            false,
																											Computed:            false,
																										},
																									},
																									Required: true,
																									Optional: false,
																									Computed: false,
																								},

																								"weight": schema.Int64Attribute{
																									Description:         "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
																									MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
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

																					"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																						Description:         "If the anti-affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the anti-affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
																						MarkdownDescription: "If the anti-affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the anti-affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"label_selector": schema.SingleNestedAttribute{
																									Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
																									MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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
																														Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																														MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																														Required:            true,
																														Optional:            false,
																														Computed:            false,
																													},

																													"values": schema.ListAttribute{
																														Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																											Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																											MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																								"match_label_keys": schema.ListAttribute{
																									Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both matchLabelKeys and labelSelector.Also, matchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																									MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both matchLabelKeys and labelSelector.Also, matchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"mismatch_label_keys": schema.ListAttribute{
																									Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.Also, mismatchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																									MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.Also, mismatchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"namespace_selector": schema.SingleNestedAttribute{
																									Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
																									MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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
																														Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																														MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																														Required:            true,
																														Optional:            false,
																														Computed:            false,
																													},

																													"values": schema.ListAttribute{
																														Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																											Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																											MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																								"namespaces": schema.ListAttribute{
																									Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																									MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"topology_key": schema.StringAttribute{
																									Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																									MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"architecture": schema.StringAttribute{
																		Description:         "Specifies the architecture of the vm guest you are attempting to run. Defaults to the compiled architecture of the KubeVirt components",
																		MarkdownDescription: "Specifies the architecture of the vm guest you are attempting to run. Defaults to the compiled architecture of the KubeVirt components",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"dns_config": schema.SingleNestedAttribute{
																		Description:         "Specifies the DNS parameters of a pod.Parameters specified here will be merged to the generated DNSconfiguration based on DNSPolicy.",
																		MarkdownDescription: "Specifies the DNS parameters of a pod.Parameters specified here will be merged to the generated DNSconfiguration based on DNSPolicy.",
																		Attributes: map[string]schema.Attribute{
																			"nameservers": schema.ListAttribute{
																				Description:         "A list of DNS name server IP addresses.This will be appended to the base nameservers generated from DNSPolicy.Duplicated nameservers will be removed.",
																				MarkdownDescription: "A list of DNS name server IP addresses.This will be appended to the base nameservers generated from DNSPolicy.Duplicated nameservers will be removed.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"options": schema.ListNestedAttribute{
																				Description:         "A list of DNS resolver options.This will be merged with the base options generated from DNSPolicy.Duplicated entries will be removed. Resolution options given in Optionswill override those that appear in the base DNSPolicy.",
																				MarkdownDescription: "A list of DNS resolver options.This will be merged with the base options generated from DNSPolicy.Duplicated entries will be removed. Resolution options given in Optionswill override those that appear in the base DNSPolicy.",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"name": schema.StringAttribute{
																							Description:         "Required.",
																							MarkdownDescription: "Required.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"value": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
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

																			"searches": schema.ListAttribute{
																				Description:         "A list of DNS search domains for host-name lookup.This will be appended to the base search paths generated from DNSPolicy.Duplicated search paths will be removed.",
																				MarkdownDescription: "A list of DNS search domains for host-name lookup.This will be appended to the base search paths generated from DNSPolicy.Duplicated search paths will be removed.",
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

																	"dns_policy": schema.StringAttribute{
																		Description:         "Set DNS policy for the pod.Defaults to 'ClusterFirst'.Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'.DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy.To have DNS options set along with hostNetwork, you have to specify DNS policyexplicitly to 'ClusterFirstWithHostNet'.",
																		MarkdownDescription: "Set DNS policy for the pod.Defaults to 'ClusterFirst'.Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'.DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy.To have DNS options set along with hostNetwork, you have to specify DNS policyexplicitly to 'ClusterFirstWithHostNet'.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"domain": schema.SingleNestedAttribute{
																		Description:         "Specification of the desired behavior of the VirtualMachineInstance on the host.",
																		MarkdownDescription: "Specification of the desired behavior of the VirtualMachineInstance on the host.",
																		Attributes: map[string]schema.Attribute{
																			"chassis": schema.SingleNestedAttribute{
																				Description:         "Chassis specifies the chassis info passed to the domain.",
																				MarkdownDescription: "Chassis specifies the chassis info passed to the domain.",
																				Attributes: map[string]schema.Attribute{
																					"asset": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"manufacturer": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"serial": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"sku": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"version": schema.StringAttribute{
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

																			"clock": schema.SingleNestedAttribute{
																				Description:         "Clock sets the clock and timers of the vmi.",
																				MarkdownDescription: "Clock sets the clock and timers of the vmi.",
																				Attributes: map[string]schema.Attribute{
																					"timer": schema.SingleNestedAttribute{
																						Description:         "Timer specifies whih timers are attached to the vmi.",
																						MarkdownDescription: "Timer specifies whih timers are attached to the vmi.",
																						Attributes: map[string]schema.Attribute{
																							"hpet": schema.SingleNestedAttribute{
																								Description:         "HPET (High Precision Event Timer) - multiple timers with periodic interrupts.",
																								MarkdownDescription: "HPET (High Precision Event Timer) - multiple timers with periodic interrupts.",
																								Attributes: map[string]schema.Attribute{
																									"present": schema.BoolAttribute{
																										Description:         "Enabled set to false makes sure that the machine type or a preset can't add the timer.Defaults to true.",
																										MarkdownDescription: "Enabled set to false makes sure that the machine type or a preset can't add the timer.Defaults to true.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"tick_policy": schema.StringAttribute{
																										Description:         "TickPolicy determines what happens when QEMU misses a deadline for injecting a tick to the guest.One of 'delay', 'catchup', 'merge', 'discard'.",
																										MarkdownDescription: "TickPolicy determines what happens when QEMU misses a deadline for injecting a tick to the guest.One of 'delay', 'catchup', 'merge', 'discard'.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},
																								},
																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"hyperv": schema.SingleNestedAttribute{
																								Description:         "Hyperv (Hypervclock) - lets guests read the host’s wall clock time (paravirtualized). For windows guests.",
																								MarkdownDescription: "Hyperv (Hypervclock) - lets guests read the host’s wall clock time (paravirtualized). For windows guests.",
																								Attributes: map[string]schema.Attribute{
																									"present": schema.BoolAttribute{
																										Description:         "Enabled set to false makes sure that the machine type or a preset can't add the timer.Defaults to true.",
																										MarkdownDescription: "Enabled set to false makes sure that the machine type or a preset can't add the timer.Defaults to true.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},
																								},
																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"kvm": schema.SingleNestedAttribute{
																								Description:         "KVM 	(KVM clock) - lets guests read the host’s wall clock time (paravirtualized). For linux guests.",
																								MarkdownDescription: "KVM 	(KVM clock) - lets guests read the host’s wall clock time (paravirtualized). For linux guests.",
																								Attributes: map[string]schema.Attribute{
																									"present": schema.BoolAttribute{
																										Description:         "Enabled set to false makes sure that the machine type or a preset can't add the timer.Defaults to true.",
																										MarkdownDescription: "Enabled set to false makes sure that the machine type or a preset can't add the timer.Defaults to true.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},
																								},
																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"pit": schema.SingleNestedAttribute{
																								Description:         "PIT (Programmable Interval Timer) - a timer with periodic interrupts.",
																								MarkdownDescription: "PIT (Programmable Interval Timer) - a timer with periodic interrupts.",
																								Attributes: map[string]schema.Attribute{
																									"present": schema.BoolAttribute{
																										Description:         "Enabled set to false makes sure that the machine type or a preset can't add the timer.Defaults to true.",
																										MarkdownDescription: "Enabled set to false makes sure that the machine type or a preset can't add the timer.Defaults to true.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"tick_policy": schema.StringAttribute{
																										Description:         "TickPolicy determines what happens when QEMU misses a deadline for injecting a tick to the guest.One of 'delay', 'catchup', 'discard'.",
																										MarkdownDescription: "TickPolicy determines what happens when QEMU misses a deadline for injecting a tick to the guest.One of 'delay', 'catchup', 'discard'.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},
																								},
																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"rtc": schema.SingleNestedAttribute{
																								Description:         "RTC (Real Time Clock) - a continuously running timer with periodic interrupts.",
																								MarkdownDescription: "RTC (Real Time Clock) - a continuously running timer with periodic interrupts.",
																								Attributes: map[string]schema.Attribute{
																									"present": schema.BoolAttribute{
																										Description:         "Enabled set to false makes sure that the machine type or a preset can't add the timer.Defaults to true.",
																										MarkdownDescription: "Enabled set to false makes sure that the machine type or a preset can't add the timer.Defaults to true.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"tick_policy": schema.StringAttribute{
																										Description:         "TickPolicy determines what happens when QEMU misses a deadline for injecting a tick to the guest.One of 'delay', 'catchup'.",
																										MarkdownDescription: "TickPolicy determines what happens when QEMU misses a deadline for injecting a tick to the guest.One of 'delay', 'catchup'.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"track": schema.StringAttribute{
																										Description:         "Track the guest or the wall clock.",
																										MarkdownDescription: "Track the guest or the wall clock.",
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

																					"timezone": schema.StringAttribute{
																						Description:         "Timezone sets the guest clock to the specified timezone.Zone name follows the TZ environment variable format (e.g. 'America/New_York').",
																						MarkdownDescription: "Timezone sets the guest clock to the specified timezone.Zone name follows the TZ environment variable format (e.g. 'America/New_York').",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"utc": schema.SingleNestedAttribute{
																						Description:         "UTC sets the guest clock to UTC on each boot. If an offset is specified,guest changes to the clock will be kept during reboots and are not reset.",
																						MarkdownDescription: "UTC sets the guest clock to UTC on each boot. If an offset is specified,guest changes to the clock will be kept during reboots and are not reset.",
																						Attributes: map[string]schema.Attribute{
																							"offset_seconds": schema.Int64Attribute{
																								Description:         "OffsetSeconds specifies an offset in seconds, relative to UTC. If set,guest changes to the clock will be kept during reboots and not reset.",
																								MarkdownDescription: "OffsetSeconds specifies an offset in seconds, relative to UTC. If set,guest changes to the clock will be kept during reboots and not reset.",
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

																			"cpu": schema.SingleNestedAttribute{
																				Description:         "CPU allow specified the detailed CPU topology inside the vmi.",
																				MarkdownDescription: "CPU allow specified the detailed CPU topology inside the vmi.",
																				Attributes: map[string]schema.Attribute{
																					"cores": schema.Int64Attribute{
																						Description:         "Cores specifies the number of cores inside the vmi.Must be a value greater or equal 1.",
																						MarkdownDescription: "Cores specifies the number of cores inside the vmi.Must be a value greater or equal 1.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"dedicated_cpu_placement": schema.BoolAttribute{
																						Description:         "DedicatedCPUPlacement requests the scheduler to place the VirtualMachineInstance on a nodewith enough dedicated pCPUs and pin the vCPUs to it.",
																						MarkdownDescription: "DedicatedCPUPlacement requests the scheduler to place the VirtualMachineInstance on a nodewith enough dedicated pCPUs and pin the vCPUs to it.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"features": schema.ListNestedAttribute{
																						Description:         "Features specifies the CPU features list inside the VMI.",
																						MarkdownDescription: "Features specifies the CPU features list inside the VMI.",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"name": schema.StringAttribute{
																									Description:         "Name of the CPU feature",
																									MarkdownDescription: "Name of the CPU feature",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"policy": schema.StringAttribute{
																									Description:         "Policy is the CPU feature attribute which can have the following attributes:force    - The virtual CPU will claim the feature is supported regardless of it being supported by host CPU.require  - Guest creation will fail unless the feature is supported by the host CPU or the hypervisor is able to emulate it.optional - The feature will be supported by virtual CPU if and only if it is supported by host CPU.disable  - The feature will not be supported by virtual CPU.forbid   - Guest creation will fail if the feature is supported by host CPU.Defaults to require",
																									MarkdownDescription: "Policy is the CPU feature attribute which can have the following attributes:force    - The virtual CPU will claim the feature is supported regardless of it being supported by host CPU.require  - Guest creation will fail unless the feature is supported by the host CPU or the hypervisor is able to emulate it.optional - The feature will be supported by virtual CPU if and only if it is supported by host CPU.disable  - The feature will not be supported by virtual CPU.forbid   - Guest creation will fail if the feature is supported by host CPU.Defaults to require",
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

																					"isolate_emulator_thread": schema.BoolAttribute{
																						Description:         "IsolateEmulatorThread requests one more dedicated pCPU to be allocated for the VMI to placethe emulator thread on it.",
																						MarkdownDescription: "IsolateEmulatorThread requests one more dedicated pCPU to be allocated for the VMI to placethe emulator thread on it.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"max_sockets": schema.Int64Attribute{
																						Description:         "MaxSockets specifies the maximum amount of sockets that canbe hotplugged",
																						MarkdownDescription: "MaxSockets specifies the maximum amount of sockets that canbe hotplugged",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"model": schema.StringAttribute{
																						Description:         "Model specifies the CPU model inside the VMI.List of available models https://github.com/libvirt/libvirt/tree/master/src/cpu_map.It is possible to specify special cases like 'host-passthrough' to get the same CPU as the nodeand 'host-model' to get CPU closest to the node one.Defaults to host-model.",
																						MarkdownDescription: "Model specifies the CPU model inside the VMI.List of available models https://github.com/libvirt/libvirt/tree/master/src/cpu_map.It is possible to specify special cases like 'host-passthrough' to get the same CPU as the nodeand 'host-model' to get CPU closest to the node one.Defaults to host-model.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"numa": schema.SingleNestedAttribute{
																						Description:         "NUMA allows specifying settings for the guest NUMA topology",
																						MarkdownDescription: "NUMA allows specifying settings for the guest NUMA topology",
																						Attributes: map[string]schema.Attribute{
																							"guest_mapping_passthrough": schema.MapAttribute{
																								Description:         "GuestMappingPassthrough will create an efficient guest topology based on host CPUs exclusively assigned to a pod.The created topology ensures that memory and CPUs on the virtual numa nodes never cross boundaries of host numa nodes.",
																								MarkdownDescription: "GuestMappingPassthrough will create an efficient guest topology based on host CPUs exclusively assigned to a pod.The created topology ensures that memory and CPUs on the virtual numa nodes never cross boundaries of host numa nodes.",
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

																					"realtime": schema.SingleNestedAttribute{
																						Description:         "Realtime instructs the virt-launcher to tune the VMI for lower latency, optional for real time workloads",
																						MarkdownDescription: "Realtime instructs the virt-launcher to tune the VMI for lower latency, optional for real time workloads",
																						Attributes: map[string]schema.Attribute{
																							"mask": schema.StringAttribute{
																								Description:         "Mask defines the vcpu mask expression that defines which vcpus are used for realtime. Format matches libvirt's expressions.Example: '0-3,^1','0,2,3','2-3'",
																								MarkdownDescription: "Mask defines the vcpu mask expression that defines which vcpus are used for realtime. Format matches libvirt's expressions.Example: '0-3,^1','0,2,3','2-3'",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},
																						},
																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"sockets": schema.Int64Attribute{
																						Description:         "Sockets specifies the number of sockets inside the vmi.Must be a value greater or equal 1.",
																						MarkdownDescription: "Sockets specifies the number of sockets inside the vmi.Must be a value greater or equal 1.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"threads": schema.Int64Attribute{
																						Description:         "Threads specifies the number of threads inside the vmi.Must be a value greater or equal 1.",
																						MarkdownDescription: "Threads specifies the number of threads inside the vmi.Must be a value greater or equal 1.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"devices": schema.SingleNestedAttribute{
																				Description:         "Devices allows adding disks, network interfaces, and others",
																				MarkdownDescription: "Devices allows adding disks, network interfaces, and others",
																				Attributes: map[string]schema.Attribute{
																					"autoattach_graphics_device": schema.BoolAttribute{
																						Description:         "Whether to attach the default graphics device or not.VNC will not be available if set to false. Defaults to true.",
																						MarkdownDescription: "Whether to attach the default graphics device or not.VNC will not be available if set to false. Defaults to true.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"autoattach_input_device": schema.BoolAttribute{
																						Description:         "Whether to attach an Input Device.Defaults to false.",
																						MarkdownDescription: "Whether to attach an Input Device.Defaults to false.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"autoattach_mem_balloon": schema.BoolAttribute{
																						Description:         "Whether to attach the Memory balloon device with default period.Period can be adjusted in virt-config.Defaults to true.",
																						MarkdownDescription: "Whether to attach the Memory balloon device with default period.Period can be adjusted in virt-config.Defaults to true.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"autoattach_pod_interface": schema.BoolAttribute{
																						Description:         "Whether to attach a pod network interface. Defaults to true.",
																						MarkdownDescription: "Whether to attach a pod network interface. Defaults to true.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"autoattach_serial_console": schema.BoolAttribute{
																						Description:         "Whether to attach the default virtio-serial console or not.Serial console access will not be available if set to false. Defaults to true.",
																						MarkdownDescription: "Whether to attach the default virtio-serial console or not.Serial console access will not be available if set to false. Defaults to true.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"autoattach_vsock": schema.BoolAttribute{
																						Description:         "Whether to attach the VSOCK CID to the VM or not.VSOCK access will be available if set to true. Defaults to false.",
																						MarkdownDescription: "Whether to attach the VSOCK CID to the VM or not.VSOCK access will be available if set to true. Defaults to false.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"block_multi_queue": schema.BoolAttribute{
																						Description:         "Whether or not to enable virtio multi-queue for block devices.Defaults to false.",
																						MarkdownDescription: "Whether or not to enable virtio multi-queue for block devices.Defaults to false.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"client_passthrough": schema.MapAttribute{
																						Description:         "To configure and access client devices such as redirecting USB",
																						MarkdownDescription: "To configure and access client devices such as redirecting USB",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"disable_hotplug": schema.BoolAttribute{
																						Description:         "DisableHotplug disabled the ability to hotplug disks.",
																						MarkdownDescription: "DisableHotplug disabled the ability to hotplug disks.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"disks": schema.ListNestedAttribute{
																						Description:         "Disks describes disks, cdroms and luns which are connected to the vmi.",
																						MarkdownDescription: "Disks describes disks, cdroms and luns which are connected to the vmi.",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"block_size": schema.SingleNestedAttribute{
																									Description:         "If specified, the virtual disk will be presented with the given block sizes.",
																									MarkdownDescription: "If specified, the virtual disk will be presented with the given block sizes.",
																									Attributes: map[string]schema.Attribute{
																										"custom": schema.SingleNestedAttribute{
																											Description:         "CustomBlockSize represents the desired logical and physical block size for a VM disk.",
																											MarkdownDescription: "CustomBlockSize represents the desired logical and physical block size for a VM disk.",
																											Attributes: map[string]schema.Attribute{
																												"logical": schema.Int64Attribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            true,
																													Optional:            false,
																													Computed:            false,
																												},

																												"physical": schema.Int64Attribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            true,
																													Optional:            false,
																													Computed:            false,
																												},
																											},
																											Required: false,
																											Optional: true,
																											Computed: false,
																										},

																										"match_volume": schema.SingleNestedAttribute{
																											Description:         "Represents if a feature is enabled or disabled.",
																											MarkdownDescription: "Represents if a feature is enabled or disabled.",
																											Attributes: map[string]schema.Attribute{
																												"enabled": schema.BoolAttribute{
																													Description:         "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																													MarkdownDescription: "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
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

																								"boot_order": schema.Int64Attribute{
																									Description:         "BootOrder is an integer value > 0, used to determine ordering of boot devices.Lower values take precedence.Each disk or interface that has a boot order must have a unique value.Disks without a boot order are not tried if a disk with a boot order exists.",
																									MarkdownDescription: "BootOrder is an integer value > 0, used to determine ordering of boot devices.Lower values take precedence.Each disk or interface that has a boot order must have a unique value.Disks without a boot order are not tried if a disk with a boot order exists.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"cache": schema.StringAttribute{
																									Description:         "Cache specifies which kvm disk cache mode should be used.Supported values are: CacheNone, CacheWriteThrough.",
																									MarkdownDescription: "Cache specifies which kvm disk cache mode should be used.Supported values are: CacheNone, CacheWriteThrough.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"cdrom": schema.SingleNestedAttribute{
																									Description:         "Attach a volume as a cdrom to the vmi.",
																									MarkdownDescription: "Attach a volume as a cdrom to the vmi.",
																									Attributes: map[string]schema.Attribute{
																										"bus": schema.StringAttribute{
																											Description:         "Bus indicates the type of disk device to emulate.supported values: virtio, sata, scsi.",
																											MarkdownDescription: "Bus indicates the type of disk device to emulate.supported values: virtio, sata, scsi.",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"readonly": schema.BoolAttribute{
																											Description:         "ReadOnly.Defaults to true.",
																											MarkdownDescription: "ReadOnly.Defaults to true.",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"tray": schema.StringAttribute{
																											Description:         "Tray indicates if the tray of the device is open or closed.Allowed values are 'open' and 'closed'.Defaults to closed.",
																											MarkdownDescription: "Tray indicates if the tray of the device is open or closed.Allowed values are 'open' and 'closed'.Defaults to closed.",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},
																									},
																									Required: false,
																									Optional: true,
																									Computed: false,
																								},

																								"dedicated_io_thread": schema.BoolAttribute{
																									Description:         "dedicatedIOThread indicates this disk should have an exclusive IO Thread.Enabling this implies useIOThreads = true.Defaults to false.",
																									MarkdownDescription: "dedicatedIOThread indicates this disk should have an exclusive IO Thread.Enabling this implies useIOThreads = true.Defaults to false.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"disk": schema.SingleNestedAttribute{
																									Description:         "Attach a volume as a disk to the vmi.",
																									MarkdownDescription: "Attach a volume as a disk to the vmi.",
																									Attributes: map[string]schema.Attribute{
																										"bus": schema.StringAttribute{
																											Description:         "Bus indicates the type of disk device to emulate.supported values: virtio, sata, scsi, usb.",
																											MarkdownDescription: "Bus indicates the type of disk device to emulate.supported values: virtio, sata, scsi, usb.",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"pci_address": schema.StringAttribute{
																											Description:         "If specified, the virtual disk will be placed on the guests pci address with the specified PCI address. For example: 0000:81:01.10",
																											MarkdownDescription: "If specified, the virtual disk will be placed on the guests pci address with the specified PCI address. For example: 0000:81:01.10",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"readonly": schema.BoolAttribute{
																											Description:         "ReadOnly.Defaults to false.",
																											MarkdownDescription: "ReadOnly.Defaults to false.",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},
																									},
																									Required: false,
																									Optional: true,
																									Computed: false,
																								},

																								"error_policy": schema.StringAttribute{
																									Description:         "If specified, it can change the default error policy (stop) for the disk",
																									MarkdownDescription: "If specified, it can change the default error policy (stop) for the disk",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"io": schema.StringAttribute{
																									Description:         "IO specifies which QEMU disk IO mode should be used.Supported values are: native, default, threads.",
																									MarkdownDescription: "IO specifies which QEMU disk IO mode should be used.Supported values are: native, default, threads.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"lun": schema.SingleNestedAttribute{
																									Description:         "Attach a volume as a LUN to the vmi.",
																									MarkdownDescription: "Attach a volume as a LUN to the vmi.",
																									Attributes: map[string]schema.Attribute{
																										"bus": schema.StringAttribute{
																											Description:         "Bus indicates the type of disk device to emulate.supported values: virtio, sata, scsi.",
																											MarkdownDescription: "Bus indicates the type of disk device to emulate.supported values: virtio, sata, scsi.",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"readonly": schema.BoolAttribute{
																											Description:         "ReadOnly.Defaults to false.",
																											MarkdownDescription: "ReadOnly.Defaults to false.",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"reservation": schema.BoolAttribute{
																											Description:         "Reservation indicates if the disk needs to support the persistent reservation for the SCSI disk",
																											MarkdownDescription: "Reservation indicates if the disk needs to support the persistent reservation for the SCSI disk",
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
																									Description:         "Name is the device name",
																									MarkdownDescription: "Name is the device name",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"serial": schema.StringAttribute{
																									Description:         "Serial provides the ability to specify a serial number for the disk device.",
																									MarkdownDescription: "Serial provides the ability to specify a serial number for the disk device.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"shareable": schema.BoolAttribute{
																									Description:         "If specified the disk is made sharable and multiple write from different VMs are permitted",
																									MarkdownDescription: "If specified the disk is made sharable and multiple write from different VMs are permitted",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"tag": schema.StringAttribute{
																									Description:         "If specified, disk address and its tag will be provided to the guest via config drive metadata",
																									MarkdownDescription: "If specified, disk address and its tag will be provided to the guest via config drive metadata",
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

																					"downward_metrics": schema.MapAttribute{
																						Description:         "DownwardMetrics creates a virtio serials for exposing the downward metrics to the vmi.",
																						MarkdownDescription: "DownwardMetrics creates a virtio serials for exposing the downward metrics to the vmi.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"filesystems": schema.ListNestedAttribute{
																						Description:         "Filesystems describes filesystem which is connected to the vmi.",
																						MarkdownDescription: "Filesystems describes filesystem which is connected to the vmi.",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"name": schema.StringAttribute{
																									Description:         "Name is the device name",
																									MarkdownDescription: "Name is the device name",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"virtiofs": schema.MapAttribute{
																									Description:         "Virtiofs is supported",
																									MarkdownDescription: "Virtiofs is supported",
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

																					"gpus": schema.ListNestedAttribute{
																						Description:         "Whether to attach a GPU device to the vmi.",
																						MarkdownDescription: "Whether to attach a GPU device to the vmi.",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"device_name": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"name": schema.StringAttribute{
																									Description:         "Name of the GPU device as exposed by a device plugin",
																									MarkdownDescription: "Name of the GPU device as exposed by a device plugin",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"tag": schema.StringAttribute{
																									Description:         "If specified, the virtual network interface address and its tag will be provided to the guest via config drive",
																									MarkdownDescription: "If specified, the virtual network interface address and its tag will be provided to the guest via config drive",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"virtual_gpu_options": schema.SingleNestedAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Attributes: map[string]schema.Attribute{
																										"display": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"enabled": schema.BoolAttribute{
																													Description:         "Enabled determines if a display addapter backed by a vGPU should be enabled or disabled on the guest.Defaults to true.",
																													MarkdownDescription: "Enabled determines if a display addapter backed by a vGPU should be enabled or disabled on the guest.Defaults to true.",
																													Required:            false,
																													Optional:            true,
																													Computed:            false,
																												},

																												"ram_fb": schema.SingleNestedAttribute{
																													Description:         "Enables a boot framebuffer, until the guest OS loads a real GPU driverDefaults to true.",
																													MarkdownDescription: "Enables a boot framebuffer, until the guest OS loads a real GPU driverDefaults to true.",
																													Attributes: map[string]schema.Attribute{
																														"enabled": schema.BoolAttribute{
																															Description:         "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																															MarkdownDescription: "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
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

																					"host_devices": schema.ListNestedAttribute{
																						Description:         "Whether to attach a host device to the vmi.",
																						MarkdownDescription: "Whether to attach a host device to the vmi.",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"device_name": schema.StringAttribute{
																									Description:         "DeviceName is the resource name of the host device exposed by a device plugin",
																									MarkdownDescription: "DeviceName is the resource name of the host device exposed by a device plugin",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"name": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"tag": schema.StringAttribute{
																									Description:         "If specified, the virtual network interface address and its tag will be provided to the guest via config drive",
																									MarkdownDescription: "If specified, the virtual network interface address and its tag will be provided to the guest via config drive",
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

																					"inputs": schema.ListNestedAttribute{
																						Description:         "Inputs describe input devices",
																						MarkdownDescription: "Inputs describe input devices",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"bus": schema.StringAttribute{
																									Description:         "Bus indicates the bus of input device to emulate.Supported values: virtio, usb.",
																									MarkdownDescription: "Bus indicates the bus of input device to emulate.Supported values: virtio, usb.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"name": schema.StringAttribute{
																									Description:         "Name is the device name",
																									MarkdownDescription: "Name is the device name",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"type": schema.StringAttribute{
																									Description:         "Type indicated the type of input device.Supported values: tablet.",
																									MarkdownDescription: "Type indicated the type of input device.Supported values: tablet.",
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

																					"interfaces": schema.ListNestedAttribute{
																						Description:         "Interfaces describe network interfaces which are added to the vmi.",
																						MarkdownDescription: "Interfaces describe network interfaces which are added to the vmi.",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"acpi_index": schema.Int64Attribute{
																									Description:         "If specified, the ACPI index is used to provide network interface device naming, that is stable across changesin PCI addresses assigned to the device.This value is required to be unique across all devices and be between 1 and (16*1024-1).",
																									MarkdownDescription: "If specified, the ACPI index is used to provide network interface device naming, that is stable across changesin PCI addresses assigned to the device.This value is required to be unique across all devices and be between 1 and (16*1024-1).",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"binding": schema.SingleNestedAttribute{
																									Description:         "Binding specifies the binding plugin that will be used to connect the interface to the guest.It provides an alternative to InterfaceBindingMethod.version: 1alphav1",
																									MarkdownDescription: "Binding specifies the binding plugin that will be used to connect the interface to the guest.It provides an alternative to InterfaceBindingMethod.version: 1alphav1",
																									Attributes: map[string]schema.Attribute{
																										"name": schema.StringAttribute{
																											Description:         "Name references to the binding name as denined in the kubevirt CR.version: 1alphav1",
																											MarkdownDescription: "Name references to the binding name as denined in the kubevirt CR.version: 1alphav1",
																											Required:            true,
																											Optional:            false,
																											Computed:            false,
																										},
																									},
																									Required: false,
																									Optional: true,
																									Computed: false,
																								},

																								"boot_order": schema.Int64Attribute{
																									Description:         "BootOrder is an integer value > 0, used to determine ordering of boot devices.Lower values take precedence.Each interface or disk that has a boot order must have a unique value.Interfaces without a boot order are not tried.",
																									MarkdownDescription: "BootOrder is an integer value > 0, used to determine ordering of boot devices.Lower values take precedence.Each interface or disk that has a boot order must have a unique value.Interfaces without a boot order are not tried.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"bridge": schema.MapAttribute{
																									Description:         "InterfaceBridge connects to a given network via a linux bridge.",
																									MarkdownDescription: "InterfaceBridge connects to a given network via a linux bridge.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"dhcp_options": schema.SingleNestedAttribute{
																									Description:         "If specified the network interface will pass additional DHCP options to the VMI",
																									MarkdownDescription: "If specified the network interface will pass additional DHCP options to the VMI",
																									Attributes: map[string]schema.Attribute{
																										"boot_file_name": schema.StringAttribute{
																											Description:         "If specified will pass option 67 to interface's DHCP server",
																											MarkdownDescription: "If specified will pass option 67 to interface's DHCP server",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"ntp_servers": schema.ListAttribute{
																											Description:         "If specified will pass the configured NTP server to the VM via DHCP option 042.",
																											MarkdownDescription: "If specified will pass the configured NTP server to the VM via DHCP option 042.",
																											ElementType:         types.StringType,
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"private_options": schema.ListNestedAttribute{
																											Description:         "If specified will pass extra DHCP options for private use, range: 224-254",
																											MarkdownDescription: "If specified will pass extra DHCP options for private use, range: 224-254",
																											NestedObject: schema.NestedAttributeObject{
																												Attributes: map[string]schema.Attribute{
																													"option": schema.Int64Attribute{
																														Description:         "Option is an Integer value from 224-254Required.",
																														MarkdownDescription: "Option is an Integer value from 224-254Required.",
																														Required:            true,
																														Optional:            false,
																														Computed:            false,
																													},

																													"value": schema.StringAttribute{
																														Description:         "Value is a String value for the Option providedRequired.",
																														MarkdownDescription: "Value is a String value for the Option providedRequired.",
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

																										"tftp_server_name": schema.StringAttribute{
																											Description:         "If specified will pass option 66 to interface's DHCP server",
																											MarkdownDescription: "If specified will pass option 66 to interface's DHCP server",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},
																									},
																									Required: false,
																									Optional: true,
																									Computed: false,
																								},

																								"mac_address": schema.StringAttribute{
																									Description:         "Interface MAC address. For example: de:ad:00:00:be:af or DE-AD-00-00-BE-AF.",
																									MarkdownDescription: "Interface MAC address. For example: de:ad:00:00:be:af or DE-AD-00-00-BE-AF.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"macvtap": schema.MapAttribute{
																									Description:         "Deprecated, please refer to Kubevirt user guide for alternatives.",
																									MarkdownDescription: "Deprecated, please refer to Kubevirt user guide for alternatives.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"masquerade": schema.MapAttribute{
																									Description:         "InterfaceMasquerade connects to a given network using netfilter rules to nat the traffic.",
																									MarkdownDescription: "InterfaceMasquerade connects to a given network using netfilter rules to nat the traffic.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"model": schema.StringAttribute{
																									Description:         "Interface model.One of: e1000, e1000e, ne2k_pci, pcnet, rtl8139, virtio.Defaults to virtio.TODO:(ihar) switch to enums once opengen-api supports them. See: https://github.com/kubernetes/kube-openapi/issues/51",
																									MarkdownDescription: "Interface model.One of: e1000, e1000e, ne2k_pci, pcnet, rtl8139, virtio.Defaults to virtio.TODO:(ihar) switch to enums once opengen-api supports them. See: https://github.com/kubernetes/kube-openapi/issues/51",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"name": schema.StringAttribute{
																									Description:         "Logical name of the interface as well as a reference to the associated networks.Must match the Name of a Network.",
																									MarkdownDescription: "Logical name of the interface as well as a reference to the associated networks.Must match the Name of a Network.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"passt": schema.MapAttribute{
																									Description:         "Deprecated, please refer to Kubevirt user guide for alternatives.",
																									MarkdownDescription: "Deprecated, please refer to Kubevirt user guide for alternatives.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"pci_address": schema.StringAttribute{
																									Description:         "If specified, the virtual network interface will be placed on the guests pci address with the specified PCI address. For example: 0000:81:01.10",
																									MarkdownDescription: "If specified, the virtual network interface will be placed on the guests pci address with the specified PCI address. For example: 0000:81:01.10",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"ports": schema.ListNestedAttribute{
																									Description:         "List of ports to be forwarded to the virtual machine.",
																									MarkdownDescription: "List of ports to be forwarded to the virtual machine.",
																									NestedObject: schema.NestedAttributeObject{
																										Attributes: map[string]schema.Attribute{
																											"name": schema.StringAttribute{
																												Description:         "If specified, this must be an IANA_SVC_NAME and unique within the pod. Eachnamed port in a pod must have a unique name. Name for the port that can bereferred to by services.",
																												MarkdownDescription: "If specified, this must be an IANA_SVC_NAME and unique within the pod. Eachnamed port in a pod must have a unique name. Name for the port that can bereferred to by services.",
																												Required:            false,
																												Optional:            true,
																												Computed:            false,
																											},

																											"port": schema.Int64Attribute{
																												Description:         "Number of port to expose for the virtual machine.This must be a valid port number, 0 < x < 65536.",
																												MarkdownDescription: "Number of port to expose for the virtual machine.This must be a valid port number, 0 < x < 65536.",
																												Required:            true,
																												Optional:            false,
																												Computed:            false,
																											},

																											"protocol": schema.StringAttribute{
																												Description:         "Protocol for port. Must be UDP or TCP.Defaults to 'TCP'.",
																												MarkdownDescription: "Protocol for port. Must be UDP or TCP.Defaults to 'TCP'.",
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

																								"slirp": schema.MapAttribute{
																									Description:         "InterfaceSlirp connects to a given network using QEMU user networking mode.",
																									MarkdownDescription: "InterfaceSlirp connects to a given network using QEMU user networking mode.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"sriov": schema.MapAttribute{
																									Description:         "InterfaceSRIOV connects to a given network by passing-through an SR-IOV PCI device via vfio.",
																									MarkdownDescription: "InterfaceSRIOV connects to a given network by passing-through an SR-IOV PCI device via vfio.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"state": schema.StringAttribute{
																									Description:         "State represents the requested operational state of the interface.The (only) value supported is 'absent', expressing a request to remove the interface.",
																									MarkdownDescription: "State represents the requested operational state of the interface.The (only) value supported is 'absent', expressing a request to remove the interface.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"tag": schema.StringAttribute{
																									Description:         "If specified, the virtual network interface address and its tag will be provided to the guest via config drive",
																									MarkdownDescription: "If specified, the virtual network interface address and its tag will be provided to the guest via config drive",
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

																					"log_serial_console": schema.BoolAttribute{
																						Description:         "Whether to log the auto-attached default serial console or not.Serial console logs will be collect to a file and then streamed from a named 'guest-console-log'.Not relevant if autoattachSerialConsole is disabled.Defaults to cluster wide setting on VirtualMachineOptions.",
																						MarkdownDescription: "Whether to log the auto-attached default serial console or not.Serial console logs will be collect to a file and then streamed from a named 'guest-console-log'.Not relevant if autoattachSerialConsole is disabled.Defaults to cluster wide setting on VirtualMachineOptions.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"network_interface_multiqueue": schema.BoolAttribute{
																						Description:         "If specified, virtual network interfaces configured with a virtio bus will also enable the vhost multiqueue feature for network devices. The number of queues created depends on additional factors of the VirtualMachineInstance, like the number of guest CPUs.",
																						MarkdownDescription: "If specified, virtual network interfaces configured with a virtio bus will also enable the vhost multiqueue feature for network devices. The number of queues created depends on additional factors of the VirtualMachineInstance, like the number of guest CPUs.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"rng": schema.MapAttribute{
																						Description:         "Whether to have random number generator from host",
																						MarkdownDescription: "Whether to have random number generator from host",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"sound": schema.SingleNestedAttribute{
																						Description:         "Whether to emulate a sound device.",
																						MarkdownDescription: "Whether to emulate a sound device.",
																						Attributes: map[string]schema.Attribute{
																							"model": schema.StringAttribute{
																								Description:         "We only support ich9 or ac97.If SoundDevice is not set: No sound card is emulated.If SoundDevice is set but Model is not: ich9",
																								MarkdownDescription: "We only support ich9 or ac97.If SoundDevice is not set: No sound card is emulated.If SoundDevice is set but Model is not: ich9",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"name": schema.StringAttribute{
																								Description:         "User's defined name for this sound device",
																								MarkdownDescription: "User's defined name for this sound device",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},
																						},
																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"tpm": schema.SingleNestedAttribute{
																						Description:         "Whether to emulate a TPM device.",
																						MarkdownDescription: "Whether to emulate a TPM device.",
																						Attributes: map[string]schema.Attribute{
																							"persistent": schema.BoolAttribute{
																								Description:         "Persistent indicates the state of the TPM device should be kept accross rebootsDefaults to false",
																								MarkdownDescription: "Persistent indicates the state of the TPM device should be kept accross rebootsDefaults to false",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},
																						},
																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"use_virtio_transitional": schema.BoolAttribute{
																						Description:         "Fall back to legacy virtio 0.9 support if virtio bus is selected on devices.This is helpful for old machines like CentOS6 or RHEL6 whichdo not understand virtio_non_transitional (virtio 1.0).",
																						MarkdownDescription: "Fall back to legacy virtio 0.9 support if virtio bus is selected on devices.This is helpful for old machines like CentOS6 or RHEL6 whichdo not understand virtio_non_transitional (virtio 1.0).",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"watchdog": schema.SingleNestedAttribute{
																						Description:         "Watchdog describes a watchdog device which can be added to the vmi.",
																						MarkdownDescription: "Watchdog describes a watchdog device which can be added to the vmi.",
																						Attributes: map[string]schema.Attribute{
																							"i6300esb": schema.SingleNestedAttribute{
																								Description:         "i6300esb watchdog device.",
																								MarkdownDescription: "i6300esb watchdog device.",
																								Attributes: map[string]schema.Attribute{
																									"action": schema.StringAttribute{
																										Description:         "The action to take. Valid values are poweroff, reset, shutdown.Defaults to reset.",
																										MarkdownDescription: "The action to take. Valid values are poweroff, reset, shutdown.Defaults to reset.",
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
																								Description:         "Name of the watchdog.",
																								MarkdownDescription: "Name of the watchdog.",
																								Required:            true,
																								Optional:            false,
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

																			"features": schema.SingleNestedAttribute{
																				Description:         "Features like acpi, apic, hyperv, smm.",
																				MarkdownDescription: "Features like acpi, apic, hyperv, smm.",
																				Attributes: map[string]schema.Attribute{
																					"acpi": schema.SingleNestedAttribute{
																						Description:         "ACPI enables/disables ACPI inside the guest.Defaults to enabled.",
																						MarkdownDescription: "ACPI enables/disables ACPI inside the guest.Defaults to enabled.",
																						Attributes: map[string]schema.Attribute{
																							"enabled": schema.BoolAttribute{
																								Description:         "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																								MarkdownDescription: "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},
																						},
																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"apic": schema.SingleNestedAttribute{
																						Description:         "Defaults to the machine type setting.",
																						MarkdownDescription: "Defaults to the machine type setting.",
																						Attributes: map[string]schema.Attribute{
																							"enabled": schema.BoolAttribute{
																								Description:         "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																								MarkdownDescription: "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"end_of_interrupt": schema.BoolAttribute{
																								Description:         "EndOfInterrupt enables the end of interrupt notification in the guest.Defaults to false.",
																								MarkdownDescription: "EndOfInterrupt enables the end of interrupt notification in the guest.Defaults to false.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},
																						},
																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"hyperv": schema.SingleNestedAttribute{
																						Description:         "Defaults to the machine type setting.",
																						MarkdownDescription: "Defaults to the machine type setting.",
																						Attributes: map[string]schema.Attribute{
																							"evmcs": schema.SingleNestedAttribute{
																								Description:         "EVMCS Speeds up L2 vmexits, but disables other virtualization features. Requires vapic.Defaults to the machine type setting.",
																								MarkdownDescription: "EVMCS Speeds up L2 vmexits, but disables other virtualization features. Requires vapic.Defaults to the machine type setting.",
																								Attributes: map[string]schema.Attribute{
																									"enabled": schema.BoolAttribute{
																										Description:         "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										MarkdownDescription: "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},
																								},
																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"frequencies": schema.SingleNestedAttribute{
																								Description:         "Frequencies improves the TSC clock source handling for Hyper-V on KVM.Defaults to the machine type setting.",
																								MarkdownDescription: "Frequencies improves the TSC clock source handling for Hyper-V on KVM.Defaults to the machine type setting.",
																								Attributes: map[string]schema.Attribute{
																									"enabled": schema.BoolAttribute{
																										Description:         "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										MarkdownDescription: "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},
																								},
																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"ipi": schema.SingleNestedAttribute{
																								Description:         "IPI improves performances in overcommited environments. Requires vpindex.Defaults to the machine type setting.",
																								MarkdownDescription: "IPI improves performances in overcommited environments. Requires vpindex.Defaults to the machine type setting.",
																								Attributes: map[string]schema.Attribute{
																									"enabled": schema.BoolAttribute{
																										Description:         "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										MarkdownDescription: "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},
																								},
																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"reenlightenment": schema.SingleNestedAttribute{
																								Description:         "Reenlightenment enables the notifications on TSC frequency changes.Defaults to the machine type setting.",
																								MarkdownDescription: "Reenlightenment enables the notifications on TSC frequency changes.Defaults to the machine type setting.",
																								Attributes: map[string]schema.Attribute{
																									"enabled": schema.BoolAttribute{
																										Description:         "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										MarkdownDescription: "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},
																								},
																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"relaxed": schema.SingleNestedAttribute{
																								Description:         "Relaxed instructs the guest OS to disable watchdog timeouts.Defaults to the machine type setting.",
																								MarkdownDescription: "Relaxed instructs the guest OS to disable watchdog timeouts.Defaults to the machine type setting.",
																								Attributes: map[string]schema.Attribute{
																									"enabled": schema.BoolAttribute{
																										Description:         "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										MarkdownDescription: "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},
																								},
																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"reset": schema.SingleNestedAttribute{
																								Description:         "Reset enables Hyperv reboot/reset for the vmi. Requires synic.Defaults to the machine type setting.",
																								MarkdownDescription: "Reset enables Hyperv reboot/reset for the vmi. Requires synic.Defaults to the machine type setting.",
																								Attributes: map[string]schema.Attribute{
																									"enabled": schema.BoolAttribute{
																										Description:         "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										MarkdownDescription: "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},
																								},
																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"runtime": schema.SingleNestedAttribute{
																								Description:         "Runtime improves the time accounting to improve scheduling in the guest.Defaults to the machine type setting.",
																								MarkdownDescription: "Runtime improves the time accounting to improve scheduling in the guest.Defaults to the machine type setting.",
																								Attributes: map[string]schema.Attribute{
																									"enabled": schema.BoolAttribute{
																										Description:         "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										MarkdownDescription: "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},
																								},
																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"spinlocks": schema.SingleNestedAttribute{
																								Description:         "Spinlocks allows to configure the spinlock retry attempts.",
																								MarkdownDescription: "Spinlocks allows to configure the spinlock retry attempts.",
																								Attributes: map[string]schema.Attribute{
																									"enabled": schema.BoolAttribute{
																										Description:         "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										MarkdownDescription: "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"spinlocks": schema.Int64Attribute{
																										Description:         "Retries indicates the number of retries.Must be a value greater or equal 4096.Defaults to 4096.",
																										MarkdownDescription: "Retries indicates the number of retries.Must be a value greater or equal 4096.Defaults to 4096.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},
																								},
																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"synic": schema.SingleNestedAttribute{
																								Description:         "SyNIC enables the Synthetic Interrupt Controller.Defaults to the machine type setting.",
																								MarkdownDescription: "SyNIC enables the Synthetic Interrupt Controller.Defaults to the machine type setting.",
																								Attributes: map[string]schema.Attribute{
																									"enabled": schema.BoolAttribute{
																										Description:         "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										MarkdownDescription: "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},
																								},
																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"synictimer": schema.SingleNestedAttribute{
																								Description:         "SyNICTimer enables Synthetic Interrupt Controller Timers, reducing CPU load.Defaults to the machine type setting.",
																								MarkdownDescription: "SyNICTimer enables Synthetic Interrupt Controller Timers, reducing CPU load.Defaults to the machine type setting.",
																								Attributes: map[string]schema.Attribute{
																									"direct": schema.SingleNestedAttribute{
																										Description:         "Represents if a feature is enabled or disabled.",
																										MarkdownDescription: "Represents if a feature is enabled or disabled.",
																										Attributes: map[string]schema.Attribute{
																											"enabled": schema.BoolAttribute{
																												Description:         "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																												MarkdownDescription: "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																												Required:            false,
																												Optional:            true,
																												Computed:            false,
																											},
																										},
																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"enabled": schema.BoolAttribute{
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

																							"tlbflush": schema.SingleNestedAttribute{
																								Description:         "TLBFlush improves performances in overcommited environments. Requires vpindex.Defaults to the machine type setting.",
																								MarkdownDescription: "TLBFlush improves performances in overcommited environments. Requires vpindex.Defaults to the machine type setting.",
																								Attributes: map[string]schema.Attribute{
																									"enabled": schema.BoolAttribute{
																										Description:         "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										MarkdownDescription: "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},
																								},
																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"vapic": schema.SingleNestedAttribute{
																								Description:         "VAPIC improves the paravirtualized handling of interrupts.Defaults to the machine type setting.",
																								MarkdownDescription: "VAPIC improves the paravirtualized handling of interrupts.Defaults to the machine type setting.",
																								Attributes: map[string]schema.Attribute{
																									"enabled": schema.BoolAttribute{
																										Description:         "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										MarkdownDescription: "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},
																								},
																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"vendorid": schema.SingleNestedAttribute{
																								Description:         "VendorID allows setting the hypervisor vendor id.Defaults to the machine type setting.",
																								MarkdownDescription: "VendorID allows setting the hypervisor vendor id.Defaults to the machine type setting.",
																								Attributes: map[string]schema.Attribute{
																									"enabled": schema.BoolAttribute{
																										Description:         "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										MarkdownDescription: "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"vendorid": schema.StringAttribute{
																										Description:         "VendorID sets the hypervisor vendor id, visible to the vmi.String up to twelve characters.",
																										MarkdownDescription: "VendorID sets the hypervisor vendor id, visible to the vmi.String up to twelve characters.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},
																								},
																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"vpindex": schema.SingleNestedAttribute{
																								Description:         "VPIndex enables the Virtual Processor Index to help windows identifying virtual processors.Defaults to the machine type setting.",
																								MarkdownDescription: "VPIndex enables the Virtual Processor Index to help windows identifying virtual processors.Defaults to the machine type setting.",
																								Attributes: map[string]schema.Attribute{
																									"enabled": schema.BoolAttribute{
																										Description:         "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																										MarkdownDescription: "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
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

																					"kvm": schema.SingleNestedAttribute{
																						Description:         "Configure how KVM presence is exposed to the guest.",
																						MarkdownDescription: "Configure how KVM presence is exposed to the guest.",
																						Attributes: map[string]schema.Attribute{
																							"hidden": schema.BoolAttribute{
																								Description:         "Hide the KVM hypervisor from standard MSR based discovery.Defaults to false",
																								MarkdownDescription: "Hide the KVM hypervisor from standard MSR based discovery.Defaults to false",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},
																						},
																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"pvspinlock": schema.SingleNestedAttribute{
																						Description:         "Notify the guest that the host supports paravirtual spinlocks.For older kernels this feature should be explicitly disabled.",
																						MarkdownDescription: "Notify the guest that the host supports paravirtual spinlocks.For older kernels this feature should be explicitly disabled.",
																						Attributes: map[string]schema.Attribute{
																							"enabled": schema.BoolAttribute{
																								Description:         "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																								MarkdownDescription: "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},
																						},
																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"smm": schema.SingleNestedAttribute{
																						Description:         "SMM enables/disables System Management Mode.TSEG not yet implemented.",
																						MarkdownDescription: "SMM enables/disables System Management Mode.TSEG not yet implemented.",
																						Attributes: map[string]schema.Attribute{
																							"enabled": schema.BoolAttribute{
																								Description:         "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
																								MarkdownDescription: "Enabled determines if the feature should be enabled or disabled on the guest.Defaults to true.",
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

																			"firmware": schema.SingleNestedAttribute{
																				Description:         "Firmware.",
																				MarkdownDescription: "Firmware.",
																				Attributes: map[string]schema.Attribute{
																					"acpi": schema.SingleNestedAttribute{
																						Description:         "Information that can be set in the ACPI table",
																						MarkdownDescription: "Information that can be set in the ACPI table",
																						Attributes: map[string]schema.Attribute{
																							"slic_name_ref": schema.StringAttribute{
																								Description:         "SlicNameRef should match the volume name of a secret object. The data in the secret shouldbe a binary blob that follows the ACPI SLIC standard, see:https://learn.microsoft.com/en-us/previous-versions/windows/hardware/design/dn653305(v=vs.85)",
																								MarkdownDescription: "SlicNameRef should match the volume name of a secret object. The data in the secret shouldbe a binary blob that follows the ACPI SLIC standard, see:https://learn.microsoft.com/en-us/previous-versions/windows/hardware/design/dn653305(v=vs.85)",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},
																						},
																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"bootloader": schema.SingleNestedAttribute{
																						Description:         "Settings to control the bootloader that is used.",
																						MarkdownDescription: "Settings to control the bootloader that is used.",
																						Attributes: map[string]schema.Attribute{
																							"bios": schema.SingleNestedAttribute{
																								Description:         "If set (default), BIOS will be used.",
																								MarkdownDescription: "If set (default), BIOS will be used.",
																								Attributes: map[string]schema.Attribute{
																									"use_serial": schema.BoolAttribute{
																										Description:         "If set, the BIOS output will be transmitted over serial",
																										MarkdownDescription: "If set, the BIOS output will be transmitted over serial",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},
																								},
																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"efi": schema.SingleNestedAttribute{
																								Description:         "If set, EFI will be used instead of BIOS.",
																								MarkdownDescription: "If set, EFI will be used instead of BIOS.",
																								Attributes: map[string]schema.Attribute{
																									"persistent": schema.BoolAttribute{
																										Description:         "If set to true, Persistent will persist the EFI NVRAM across reboots.Defaults to false",
																										MarkdownDescription: "If set to true, Persistent will persist the EFI NVRAM across reboots.Defaults to false",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"secure_boot": schema.BoolAttribute{
																										Description:         "If set, SecureBoot will be enabled and the OVMF roms will be swapped forSecureBoot-enabled ones.Requires SMM to be enabled.Defaults to true",
																										MarkdownDescription: "If set, SecureBoot will be enabled and the OVMF roms will be swapped forSecureBoot-enabled ones.Requires SMM to be enabled.Defaults to true",
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

																					"kernel_boot": schema.SingleNestedAttribute{
																						Description:         "Settings to set the kernel for booting.",
																						MarkdownDescription: "Settings to set the kernel for booting.",
																						Attributes: map[string]schema.Attribute{
																							"container": schema.SingleNestedAttribute{
																								Description:         "Container defines the container that containes kernel artifacts",
																								MarkdownDescription: "Container defines the container that containes kernel artifacts",
																								Attributes: map[string]schema.Attribute{
																									"image": schema.StringAttribute{
																										Description:         "Image that contains initrd / kernel files.",
																										MarkdownDescription: "Image that contains initrd / kernel files.",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"image_pull_policy": schema.StringAttribute{
																										Description:         "Image pull policy.One of Always, Never, IfNotPresent.Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.Cannot be updated.More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
																										MarkdownDescription: "Image pull policy.One of Always, Never, IfNotPresent.Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.Cannot be updated.More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"image_pull_secret": schema.StringAttribute{
																										Description:         "ImagePullSecret is the name of the Docker registry secret required to pull the image. The secret must already exist.",
																										MarkdownDescription: "ImagePullSecret is the name of the Docker registry secret required to pull the image. The secret must already exist.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"initrd_path": schema.StringAttribute{
																										Description:         "the fully-qualified path to the ramdisk image in the host OS",
																										MarkdownDescription: "the fully-qualified path to the ramdisk image in the host OS",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"kernel_path": schema.StringAttribute{
																										Description:         "The fully-qualified path to the kernel image in the host OS",
																										MarkdownDescription: "The fully-qualified path to the kernel image in the host OS",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},
																								},
																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"kernel_args": schema.StringAttribute{
																								Description:         "Arguments to be passed to the kernel at boot time",
																								MarkdownDescription: "Arguments to be passed to the kernel at boot time",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},
																						},
																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"serial": schema.StringAttribute{
																						Description:         "The system-serial-number in SMBIOS",
																						MarkdownDescription: "The system-serial-number in SMBIOS",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"uuid": schema.StringAttribute{
																						Description:         "UUID reported by the vmi bios.Defaults to a random generated uid.",
																						MarkdownDescription: "UUID reported by the vmi bios.Defaults to a random generated uid.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"io_threads_policy": schema.StringAttribute{
																				Description:         "Controls whether or not disks will share IOThreads.Omitting IOThreadsPolicy disables use of IOThreads.One of: shared, auto",
																				MarkdownDescription: "Controls whether or not disks will share IOThreads.Omitting IOThreadsPolicy disables use of IOThreads.One of: shared, auto",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"launch_security": schema.SingleNestedAttribute{
																				Description:         "Launch Security setting of the vmi.",
																				MarkdownDescription: "Launch Security setting of the vmi.",
																				Attributes: map[string]schema.Attribute{
																					"sev": schema.SingleNestedAttribute{
																						Description:         "AMD Secure Encrypted Virtualization (SEV).",
																						MarkdownDescription: "AMD Secure Encrypted Virtualization (SEV).",
																						Attributes: map[string]schema.Attribute{
																							"attestation": schema.MapAttribute{
																								Description:         "If specified, run the attestation process for a vmi.",
																								MarkdownDescription: "If specified, run the attestation process for a vmi.",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"dh_cert": schema.StringAttribute{
																								Description:         "Base64 encoded guest owner's Diffie-Hellman key.",
																								MarkdownDescription: "Base64 encoded guest owner's Diffie-Hellman key.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"policy": schema.SingleNestedAttribute{
																								Description:         "Guest policy flags as defined in AMD SEV API specification.Note: due to security reasons it is not allowed to enable guest debugging. Therefore NoDebug flag is not exposed to users and is always true.",
																								MarkdownDescription: "Guest policy flags as defined in AMD SEV API specification.Note: due to security reasons it is not allowed to enable guest debugging. Therefore NoDebug flag is not exposed to users and is always true.",
																								Attributes: map[string]schema.Attribute{
																									"encrypted_state": schema.BoolAttribute{
																										Description:         "SEV-ES is required.Defaults to false.",
																										MarkdownDescription: "SEV-ES is required.Defaults to false.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},
																								},
																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"session": schema.StringAttribute{
																								Description:         "Base64 encoded session blob.",
																								MarkdownDescription: "Base64 encoded session blob.",
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

																			"machine": schema.SingleNestedAttribute{
																				Description:         "Machine type.",
																				MarkdownDescription: "Machine type.",
																				Attributes: map[string]schema.Attribute{
																					"type": schema.StringAttribute{
																						Description:         "QEMU machine type is the actual chipset of the VirtualMachineInstance.",
																						MarkdownDescription: "QEMU machine type is the actual chipset of the VirtualMachineInstance.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"memory": schema.SingleNestedAttribute{
																				Description:         "Memory allow specifying the VMI memory features.",
																				MarkdownDescription: "Memory allow specifying the VMI memory features.",
																				Attributes: map[string]schema.Attribute{
																					"guest": schema.StringAttribute{
																						Description:         "Guest allows to specifying the amount of memory which is visible inside the Guest OS.The Guest must lie between Requests and Limits from the resources section.Defaults to the requested memory in the resources section if not specified.",
																						MarkdownDescription: "Guest allows to specifying the amount of memory which is visible inside the Guest OS.The Guest must lie between Requests and Limits from the resources section.Defaults to the requested memory in the resources section if not specified.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"hugepages": schema.SingleNestedAttribute{
																						Description:         "Hugepages allow to use hugepages for the VirtualMachineInstance instead of regular memory.",
																						MarkdownDescription: "Hugepages allow to use hugepages for the VirtualMachineInstance instead of regular memory.",
																						Attributes: map[string]schema.Attribute{
																							"page_size": schema.StringAttribute{
																								Description:         "PageSize specifies the hugepage size, for x86_64 architecture valid values are 1Gi and 2Mi.",
																								MarkdownDescription: "PageSize specifies the hugepage size, for x86_64 architecture valid values are 1Gi and 2Mi.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},
																						},
																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"max_guest": schema.StringAttribute{
																						Description:         "MaxGuest allows to specify the maximum amount of memory which is visible inside the Guest OS.The delta between MaxGuest and Guest is the amount of memory that can be hot(un)plugged.",
																						MarkdownDescription: "MaxGuest allows to specify the maximum amount of memory which is visible inside the Guest OS.The delta between MaxGuest and Guest is the amount of memory that can be hot(un)plugged.",
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
																				Description:         "Resources describes the Compute Resources required by this vmi.",
																				MarkdownDescription: "Resources describes the Compute Resources required by this vmi.",
																				Attributes: map[string]schema.Attribute{
																					"limits": schema.MapAttribute{
																						Description:         "Limits describes the maximum amount of compute resources allowed.Valid resource keys are 'memory' and 'cpu'.",
																						MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.Valid resource keys are 'memory' and 'cpu'.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"overcommit_guest_overhead": schema.BoolAttribute{
																						Description:         "Don't ask the scheduler to take the guest-management overhead into account. Insteadput the overhead only into the container's memory limit. This can lead to crashes ifall memory is in use on a node. Defaults to false.",
																						MarkdownDescription: "Don't ask the scheduler to take the guest-management overhead into account. Insteadput the overhead only into the container's memory limit. This can lead to crashes ifall memory is in use on a node. Defaults to false.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"requests": schema.MapAttribute{
																						Description:         "Requests is a description of the initial vmi resources.Valid resource keys are 'memory' and 'cpu'.",
																						MarkdownDescription: "Requests is a description of the initial vmi resources.Valid resource keys are 'memory' and 'cpu'.",
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

																	"eviction_strategy": schema.StringAttribute{
																		Description:         "EvictionStrategy describes the strategy to follow when a node drain occurs.The possible options are:- 'None': No action will be taken, according to the specified 'RunStrategy' the VirtualMachine will be restarted or shutdown.- 'LiveMigrate': the VirtualMachineInstance will be migrated instead of being shutdown.- 'LiveMigrateIfPossible': the same as 'LiveMigrate' but only if the VirtualMachine is Live-Migratable, otherwise it will behave as 'None'.- 'External': the VirtualMachineInstance will be protected by a PDB and 'vmi.Status.EvacuationNodeName' will be set on eviction. This is mainly useful for cluster-api-provider-kubevirt (capk) which needs a way for VMI's to be blocked from eviction, yet signal capk that eviction has been called on the VMI so the capk controller can handle tearing the VMI down. Details can be found in the commit description https://github.com/kubevirt/kubevirt/commit/c1d77face705c8b126696bac9a3ee3825f27f1fa.",
																		MarkdownDescription: "EvictionStrategy describes the strategy to follow when a node drain occurs.The possible options are:- 'None': No action will be taken, according to the specified 'RunStrategy' the VirtualMachine will be restarted or shutdown.- 'LiveMigrate': the VirtualMachineInstance will be migrated instead of being shutdown.- 'LiveMigrateIfPossible': the same as 'LiveMigrate' but only if the VirtualMachine is Live-Migratable, otherwise it will behave as 'None'.- 'External': the VirtualMachineInstance will be protected by a PDB and 'vmi.Status.EvacuationNodeName' will be set on eviction. This is mainly useful for cluster-api-provider-kubevirt (capk) which needs a way for VMI's to be blocked from eviction, yet signal capk that eviction has been called on the VMI so the capk controller can handle tearing the VMI down. Details can be found in the commit description https://github.com/kubevirt/kubevirt/commit/c1d77face705c8b126696bac9a3ee3825f27f1fa.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"hostname": schema.StringAttribute{
																		Description:         "Specifies the hostname of the vmiIf not specified, the hostname will be set to the name of the vmi, if dhcp or cloud-init is configured properly.",
																		MarkdownDescription: "Specifies the hostname of the vmiIf not specified, the hostname will be set to the name of the vmi, if dhcp or cloud-init is configured properly.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"liveness_probe": schema.SingleNestedAttribute{
																		Description:         "Periodic probe of VirtualMachineInstance liveness.VirtualmachineInstances will be stopped if the probe fails.Cannot be updated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																		MarkdownDescription: "Periodic probe of VirtualMachineInstance liveness.VirtualmachineInstances will be stopped if the probe fails.Cannot be updated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																		Attributes: map[string]schema.Attribute{
																			"exec": schema.SingleNestedAttribute{
																				Description:         "One and only one of the following should be specified.Exec specifies the action to take, it will be executed on the guest through the qemu-guest-agent.If the guest agent is not available, this probe will fail.",
																				MarkdownDescription: "One and only one of the following should be specified.Exec specifies the action to take, it will be executed on the guest through the qemu-guest-agent.If the guest agent is not available, this probe will fail.",
																				Attributes: map[string]schema.Attribute{
																					"command": schema.ListAttribute{
																						Description:         "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																						MarkdownDescription: "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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

																			"failure_threshold": schema.Int64Attribute{
																				Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded.Defaults to 3. Minimum value is 1.",
																				MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded.Defaults to 3. Minimum value is 1.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"guest_agent_ping": schema.MapAttribute{
																				Description:         "GuestAgentPing contacts the qemu-guest-agent for availability checks.",
																				MarkdownDescription: "GuestAgentPing contacts the qemu-guest-agent for availability checks.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"http_get": schema.SingleNestedAttribute{
																				Description:         "HTTPGet specifies the http request to perform.",
																				MarkdownDescription: "HTTPGet specifies the http request to perform.",
																				Attributes: map[string]schema.Attribute{
																					"host": schema.StringAttribute{
																						Description:         "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
																						MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"http_headers": schema.ListNestedAttribute{
																						Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																						MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"name": schema.StringAttribute{
																									Description:         "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																									MarkdownDescription: "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"value": schema.StringAttribute{
																									Description:         "The header field value",
																									MarkdownDescription: "The header field value",
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

																					"path": schema.StringAttribute{
																						Description:         "Path to access on the HTTP server.",
																						MarkdownDescription: "Path to access on the HTTP server.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"port": schema.StringAttribute{
																						Description:         "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
																						MarkdownDescription: "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"scheme": schema.StringAttribute{
																						Description:         "Scheme to use for connecting to the host.Defaults to HTTP.",
																						MarkdownDescription: "Scheme to use for connecting to the host.Defaults to HTTP.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"initial_delay_seconds": schema.Int64Attribute{
																				Description:         "Number of seconds after the VirtualMachineInstance has started before liveness probes are initiated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																				MarkdownDescription: "Number of seconds after the VirtualMachineInstance has started before liveness probes are initiated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"period_seconds": schema.Int64Attribute{
																				Description:         "How often (in seconds) to perform the probe.Default to 10 seconds. Minimum value is 1.",
																				MarkdownDescription: "How often (in seconds) to perform the probe.Default to 10 seconds. Minimum value is 1.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"success_threshold": schema.Int64Attribute{
																				Description:         "Minimum consecutive successes for the probe to be considered successful after having failed.Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
																				MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed.Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"tcp_socket": schema.SingleNestedAttribute{
																				Description:         "TCPSocket specifies an action involving a TCP port.TCP hooks not yet supportedTODO: implement a realistic TCP lifecycle hook",
																				MarkdownDescription: "TCPSocket specifies an action involving a TCP port.TCP hooks not yet supportedTODO: implement a realistic TCP lifecycle hook",
																				Attributes: map[string]schema.Attribute{
																					"host": schema.StringAttribute{
																						Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																						MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"port": schema.StringAttribute{
																						Description:         "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
																						MarkdownDescription: "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"timeout_seconds": schema.Int64Attribute{
																				Description:         "Number of seconds after which the probe times out.For exec probes the timeout fails the probe but does not terminate the command running on the guest.This means a blocking command can result in an increasing load on the guest.A small buffer will be added to the resulting workload exec probe to compensate for delayscaused by the qemu guest exec mechanism.Defaults to 1 second. Minimum value is 1.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																				MarkdownDescription: "Number of seconds after which the probe times out.For exec probes the timeout fails the probe but does not terminate the command running on the guest.This means a blocking command can result in an increasing load on the guest.A small buffer will be added to the resulting workload exec probe to compensate for delayscaused by the qemu guest exec mechanism.Defaults to 1 second. Minimum value is 1.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"networks": schema.ListNestedAttribute{
																		Description:         "List of networks that can be attached to a vm's virtual interface.",
																		MarkdownDescription: "List of networks that can be attached to a vm's virtual interface.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"multus": schema.SingleNestedAttribute{
																					Description:         "Represents the multus cni network.",
																					MarkdownDescription: "Represents the multus cni network.",
																					Attributes: map[string]schema.Attribute{
																						"default": schema.BoolAttribute{
																							Description:         "Select the default network and add it to themultus-cni.io/default-network annotation.",
																							MarkdownDescription: "Select the default network and add it to themultus-cni.io/default-network annotation.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"network_name": schema.StringAttribute{
																							Description:         "References to a NetworkAttachmentDefinition CRD object. Format:<networkName>, <namespace>/<networkName>. If namespace is notspecified, VMI namespace is assumed.",
																							MarkdownDescription: "References to a NetworkAttachmentDefinition CRD object. Format:<networkName>, <namespace>/<networkName>. If namespace is notspecified, VMI namespace is assumed.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"name": schema.StringAttribute{
																					Description:         "Network name.Must be a DNS_LABEL and unique within the vm.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																					MarkdownDescription: "Network name.Must be a DNS_LABEL and unique within the vm.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"pod": schema.SingleNestedAttribute{
																					Description:         "Represents the stock pod network interface.",
																					MarkdownDescription: "Represents the stock pod network interface.",
																					Attributes: map[string]schema.Attribute{
																						"vm_i_pv6_network_cidr": schema.StringAttribute{
																							Description:         "IPv6 CIDR for the vm network.Defaults to fd10:0:2::/120 if not specified.",
																							MarkdownDescription: "IPv6 CIDR for the vm network.Defaults to fd10:0:2::/120 if not specified.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"vm_network_cidr": schema.StringAttribute{
																							Description:         "CIDR for vm network.Default 10.0.2.0/24 if not specified.",
																							MarkdownDescription: "CIDR for vm network.Default 10.0.2.0/24 if not specified.",
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

																	"node_selector": schema.MapAttribute{
																		Description:         "NodeSelector is a selector which must be true for the vmi to fit on a node.Selector which must match a node's labels for the vmi to be scheduled on that node.More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
																		MarkdownDescription: "NodeSelector is a selector which must be true for the vmi to fit on a node.Selector which must match a node's labels for the vmi to be scheduled on that node.More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"priority_class_name": schema.StringAttribute{
																		Description:         "If specified, indicates the pod's priority.If not specified, the pod priority will be default or zero if there is nodefault.",
																		MarkdownDescription: "If specified, indicates the pod's priority.If not specified, the pod priority will be default or zero if there is nodefault.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"readiness_probe": schema.SingleNestedAttribute{
																		Description:         "Periodic probe of VirtualMachineInstance service readiness.VirtualmachineInstances will be removed from service endpoints if the probe fails.Cannot be updated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																		MarkdownDescription: "Periodic probe of VirtualMachineInstance service readiness.VirtualmachineInstances will be removed from service endpoints if the probe fails.Cannot be updated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																		Attributes: map[string]schema.Attribute{
																			"exec": schema.SingleNestedAttribute{
																				Description:         "One and only one of the following should be specified.Exec specifies the action to take, it will be executed on the guest through the qemu-guest-agent.If the guest agent is not available, this probe will fail.",
																				MarkdownDescription: "One and only one of the following should be specified.Exec specifies the action to take, it will be executed on the guest through the qemu-guest-agent.If the guest agent is not available, this probe will fail.",
																				Attributes: map[string]schema.Attribute{
																					"command": schema.ListAttribute{
																						Description:         "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																						MarkdownDescription: "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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

																			"failure_threshold": schema.Int64Attribute{
																				Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded.Defaults to 3. Minimum value is 1.",
																				MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded.Defaults to 3. Minimum value is 1.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"guest_agent_ping": schema.MapAttribute{
																				Description:         "GuestAgentPing contacts the qemu-guest-agent for availability checks.",
																				MarkdownDescription: "GuestAgentPing contacts the qemu-guest-agent for availability checks.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"http_get": schema.SingleNestedAttribute{
																				Description:         "HTTPGet specifies the http request to perform.",
																				MarkdownDescription: "HTTPGet specifies the http request to perform.",
																				Attributes: map[string]schema.Attribute{
																					"host": schema.StringAttribute{
																						Description:         "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
																						MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"http_headers": schema.ListNestedAttribute{
																						Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																						MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"name": schema.StringAttribute{
																									Description:         "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																									MarkdownDescription: "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"value": schema.StringAttribute{
																									Description:         "The header field value",
																									MarkdownDescription: "The header field value",
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

																					"path": schema.StringAttribute{
																						Description:         "Path to access on the HTTP server.",
																						MarkdownDescription: "Path to access on the HTTP server.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"port": schema.StringAttribute{
																						Description:         "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
																						MarkdownDescription: "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"scheme": schema.StringAttribute{
																						Description:         "Scheme to use for connecting to the host.Defaults to HTTP.",
																						MarkdownDescription: "Scheme to use for connecting to the host.Defaults to HTTP.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"initial_delay_seconds": schema.Int64Attribute{
																				Description:         "Number of seconds after the VirtualMachineInstance has started before liveness probes are initiated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																				MarkdownDescription: "Number of seconds after the VirtualMachineInstance has started before liveness probes are initiated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"period_seconds": schema.Int64Attribute{
																				Description:         "How often (in seconds) to perform the probe.Default to 10 seconds. Minimum value is 1.",
																				MarkdownDescription: "How often (in seconds) to perform the probe.Default to 10 seconds. Minimum value is 1.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"success_threshold": schema.Int64Attribute{
																				Description:         "Minimum consecutive successes for the probe to be considered successful after having failed.Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
																				MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed.Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"tcp_socket": schema.SingleNestedAttribute{
																				Description:         "TCPSocket specifies an action involving a TCP port.TCP hooks not yet supportedTODO: implement a realistic TCP lifecycle hook",
																				MarkdownDescription: "TCPSocket specifies an action involving a TCP port.TCP hooks not yet supportedTODO: implement a realistic TCP lifecycle hook",
																				Attributes: map[string]schema.Attribute{
																					"host": schema.StringAttribute{
																						Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																						MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"port": schema.StringAttribute{
																						Description:         "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
																						MarkdownDescription: "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"timeout_seconds": schema.Int64Attribute{
																				Description:         "Number of seconds after which the probe times out.For exec probes the timeout fails the probe but does not terminate the command running on the guest.This means a blocking command can result in an increasing load on the guest.A small buffer will be added to the resulting workload exec probe to compensate for delayscaused by the qemu guest exec mechanism.Defaults to 1 second. Minimum value is 1.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																				MarkdownDescription: "Number of seconds after which the probe times out.For exec probes the timeout fails the probe but does not terminate the command running on the guest.This means a blocking command can result in an increasing load on the guest.A small buffer will be added to the resulting workload exec probe to compensate for delayscaused by the qemu guest exec mechanism.Defaults to 1 second. Minimum value is 1.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"scheduler_name": schema.StringAttribute{
																		Description:         "If specified, the VMI will be dispatched by specified scheduler.If not specified, the VMI will be dispatched by default scheduler.",
																		MarkdownDescription: "If specified, the VMI will be dispatched by specified scheduler.If not specified, the VMI will be dispatched by default scheduler.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"start_strategy": schema.StringAttribute{
																		Description:         "StartStrategy can be set to 'Paused' if Virtual Machine should be started in paused state.",
																		MarkdownDescription: "StartStrategy can be set to 'Paused' if Virtual Machine should be started in paused state.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"subdomain": schema.StringAttribute{
																		Description:         "If specified, the fully qualified vmi hostname will be '<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>'.If not specified, the vmi will not have a domainname at all. The DNS entry will resolve to the vmi,no matter if the vmi itself can pick up a hostname.",
																		MarkdownDescription: "If specified, the fully qualified vmi hostname will be '<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>'.If not specified, the vmi will not have a domainname at all. The DNS entry will resolve to the vmi,no matter if the vmi itself can pick up a hostname.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"termination_grace_period_seconds": schema.Int64Attribute{
																		Description:         "Grace period observed after signalling a VirtualMachineInstance to stop after which the VirtualMachineInstance is force terminated.",
																		MarkdownDescription: "Grace period observed after signalling a VirtualMachineInstance to stop after which the VirtualMachineInstance is force terminated.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"tolerations": schema.ListNestedAttribute{
																		Description:         "If toleration is specified, obey all the toleration rules.",
																		MarkdownDescription: "If toleration is specified, obey all the toleration rules.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"effect": schema.StringAttribute{
																					Description:         "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
																					MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"key": schema.StringAttribute{
																					Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
																					MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"operator": schema.StringAttribute{
																					Description:         "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
																					MarkdownDescription: "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"toleration_seconds": schema.Int64Attribute{
																					Description:         "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
																					MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"value": schema.StringAttribute{
																					Description:         "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
																					MarkdownDescription: "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

																	"topology_spread_constraints": schema.ListNestedAttribute{
																		Description:         "TopologySpreadConstraints describes how a group of VMIs will be spread across a given topologydomains. K8s scheduler will schedule VMI pods in a way which abides by the constraints.",
																		MarkdownDescription: "TopologySpreadConstraints describes how a group of VMIs will be spread across a given topologydomains. K8s scheduler will schedule VMI pods in a way which abides by the constraints.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"label_selector": schema.SingleNestedAttribute{
																					Description:         "LabelSelector is used to find matching pods.Pods that match this label selector are counted to determine the number of podsin their corresponding topology domain.",
																					MarkdownDescription: "LabelSelector is used to find matching pods.Pods that match this label selector are counted to determine the number of podsin their corresponding topology domain.",
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
																										Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																										MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"values": schema.ListAttribute{
																										Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																										MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																							Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

																				"match_label_keys": schema.ListAttribute{
																					Description:         "MatchLabelKeys is a set of pod label keys to select the pods over whichspreading will be calculated. The keys are used to lookup values from theincoming pod labels, those key-value labels are ANDed with labelSelectorto select the group of existing pods over which spreading will be calculatedfor the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.MatchLabelKeys cannot be set when LabelSelector isn't set.Keys that don't exist in the incoming pod labels willbe ignored. A null or empty list means only match against labelSelector.This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).",
																					MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select the pods over whichspreading will be calculated. The keys are used to lookup values from theincoming pod labels, those key-value labels are ANDed with labelSelectorto select the group of existing pods over which spreading will be calculatedfor the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.MatchLabelKeys cannot be set when LabelSelector isn't set.Keys that don't exist in the incoming pod labels willbe ignored. A null or empty list means only match against labelSelector.This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"max_skew": schema.Int64Attribute{
																					Description:         "MaxSkew describes the degree to which pods may be unevenly distributed.When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted differencebetween the number of matching pods in the target topology and the global minimum.The global minimum is the minimum number of matching pods in an eligible domainor zero if the number of eligible domains is less than MinDomains.For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the samelabelSelector spread as 2/2/1:In this case, the global minimum is 1.| zone1 | zone2 | zone3 ||  P P  |  P P  |   P   |- if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2;scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2)violate MaxSkew(1).- if MaxSkew is 2, incoming pod can be scheduled onto any zone.When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedenceto topologies that satisfy it.It's a required field. Default value is 1 and 0 is not allowed.",
																					MarkdownDescription: "MaxSkew describes the degree to which pods may be unevenly distributed.When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted differencebetween the number of matching pods in the target topology and the global minimum.The global minimum is the minimum number of matching pods in an eligible domainor zero if the number of eligible domains is less than MinDomains.For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the samelabelSelector spread as 2/2/1:In this case, the global minimum is 1.| zone1 | zone2 | zone3 ||  P P  |  P P  |   P   |- if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2;scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2)violate MaxSkew(1).- if MaxSkew is 2, incoming pod can be scheduled onto any zone.When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedenceto topologies that satisfy it.It's a required field. Default value is 1 and 0 is not allowed.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"min_domains": schema.Int64Attribute{
																					Description:         "MinDomains indicates a minimum number of eligible domains.When the number of eligible domains with matching topology keys is less than minDomains,Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed.And when the number of eligible domains with matching topology keys equals or greater than minDomains,this value has no effect on scheduling.As a result, when the number of eligible domains is less than minDomains,scheduler won't schedule more than maxSkew Pods to those domains.If value is nil, the constraint behaves as if MinDomains is equal to 1.Valid values are integers greater than 0.When value is not nil, WhenUnsatisfiable must be DoNotSchedule.For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the samelabelSelector spread as 2/2/2:| zone1 | zone2 | zone3 ||  P P  |  P P  |  P P  |The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0.In this situation, new pod with the same labelSelector cannot be scheduled,because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones,it will violate MaxSkew.",
																					MarkdownDescription: "MinDomains indicates a minimum number of eligible domains.When the number of eligible domains with matching topology keys is less than minDomains,Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed.And when the number of eligible domains with matching topology keys equals or greater than minDomains,this value has no effect on scheduling.As a result, when the number of eligible domains is less than minDomains,scheduler won't schedule more than maxSkew Pods to those domains.If value is nil, the constraint behaves as if MinDomains is equal to 1.Valid values are integers greater than 0.When value is not nil, WhenUnsatisfiable must be DoNotSchedule.For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the samelabelSelector spread as 2/2/2:| zone1 | zone2 | zone3 ||  P P  |  P P  |  P P  |The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0.In this situation, new pod with the same labelSelector cannot be scheduled,because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones,it will violate MaxSkew.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"node_affinity_policy": schema.StringAttribute{
																					Description:         "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelectorwhen calculating pod topology spread skew. Options are:- Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations.- Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.If this value is nil, the behavior is equivalent to the Honor policy.This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
																					MarkdownDescription: "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelectorwhen calculating pod topology spread skew. Options are:- Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations.- Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.If this value is nil, the behavior is equivalent to the Honor policy.This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"node_taints_policy": schema.StringAttribute{
																					Description:         "NodeTaintsPolicy indicates how we will treat node taints when calculatingpod topology spread skew. Options are:- Honor: nodes without taints, along with tainted nodes for which the incoming podhas a toleration, are included.- Ignore: node taints are ignored. All nodes are included.If this value is nil, the behavior is equivalent to the Ignore policy.This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
																					MarkdownDescription: "NodeTaintsPolicy indicates how we will treat node taints when calculatingpod topology spread skew. Options are:- Honor: nodes without taints, along with tainted nodes for which the incoming podhas a toleration, are included.- Ignore: node taints are ignored. All nodes are included.If this value is nil, the behavior is equivalent to the Ignore policy.This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"topology_key": schema.StringAttribute{
																					Description:         "TopologyKey is the key of node labels. Nodes that have a label with this keyand identical values are considered to be in the same topology.We consider each <key, value> as a 'bucket', and try to put balanced numberof pods into each bucket.We define a domain as a particular instance of a topology.Also, we define an eligible domain as a domain whose nodes meet the requirements ofnodeAffinityPolicy and nodeTaintsPolicy.e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology.And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology.It's a required field.",
																					MarkdownDescription: "TopologyKey is the key of node labels. Nodes that have a label with this keyand identical values are considered to be in the same topology.We consider each <key, value> as a 'bucket', and try to put balanced numberof pods into each bucket.We define a domain as a particular instance of a topology.Also, we define an eligible domain as a domain whose nodes meet the requirements ofnodeAffinityPolicy and nodeTaintsPolicy.e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology.And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology.It's a required field.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"when_unsatisfiable": schema.StringAttribute{
																					Description:         "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfythe spread constraint.- DoNotSchedule (default) tells the scheduler not to schedule it.- ScheduleAnyway tells the scheduler to schedule the pod in any location,  but giving higher precedence to topologies that would help reduce the  skew.A constraint is considered 'Unsatisfiable' for an incoming podif and only if every possible node assignment for that pod would violate'MaxSkew' on some topology.For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the samelabelSelector spread as 3/1/1:| zone1 | zone2 | zone3 || P P P |   P   |   P   |If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduledto zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfiesMaxSkew(1). In other words, the cluster can still be imbalanced, but schedulerwon't make it *more* imbalanced.It's a required field.",
																					MarkdownDescription: "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfythe spread constraint.- DoNotSchedule (default) tells the scheduler not to schedule it.- ScheduleAnyway tells the scheduler to schedule the pod in any location,  but giving higher precedence to topologies that would help reduce the  skew.A constraint is considered 'Unsatisfiable' for an incoming podif and only if every possible node assignment for that pod would violate'MaxSkew' on some topology.For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the samelabelSelector spread as 3/1/1:| zone1 | zone2 | zone3 || P P P |   P   |   P   |If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduledto zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfiesMaxSkew(1). In other words, the cluster can still be imbalanced, but schedulerwon't make it *more* imbalanced.It's a required field.",
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

																	"volumes": schema.ListNestedAttribute{
																		Description:         "List of volumes that can be mounted by disks belonging to the vmi.",
																		MarkdownDescription: "List of volumes that can be mounted by disks belonging to the vmi.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"cloud_init_config_drive": schema.SingleNestedAttribute{
																					Description:         "CloudInitConfigDrive represents a cloud-init Config Drive user-data source.The Config Drive data will be added as a disk to the vmi. A proper cloud-init installation is required inside the guest.More info: https://cloudinit.readthedocs.io/en/latest/topics/datasources/configdrive.html",
																					MarkdownDescription: "CloudInitConfigDrive represents a cloud-init Config Drive user-data source.The Config Drive data will be added as a disk to the vmi. A proper cloud-init installation is required inside the guest.More info: https://cloudinit.readthedocs.io/en/latest/topics/datasources/configdrive.html",
																					Attributes: map[string]schema.Attribute{
																						"network_data": schema.StringAttribute{
																							Description:         "NetworkData contains config drive inline cloud-init networkdata.",
																							MarkdownDescription: "NetworkData contains config drive inline cloud-init networkdata.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"network_data_base64": schema.StringAttribute{
																							Description:         "NetworkDataBase64 contains config drive cloud-init networkdata as a base64 encoded string.",
																							MarkdownDescription: "NetworkDataBase64 contains config drive cloud-init networkdata as a base64 encoded string.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"network_data_secret_ref": schema.SingleNestedAttribute{
																							Description:         "NetworkDataSecretRef references a k8s secret that contains config drive networkdata.",
																							MarkdownDescription: "NetworkDataSecretRef references a k8s secret that contains config drive networkdata.",
																							Attributes: map[string]schema.Attribute{
																								"name": schema.StringAttribute{
																									Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																									MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},
																							},
																							Required: false,
																							Optional: true,
																							Computed: false,
																						},

																						"secret_ref": schema.SingleNestedAttribute{
																							Description:         "UserDataSecretRef references a k8s secret that contains config drive userdata.",
																							MarkdownDescription: "UserDataSecretRef references a k8s secret that contains config drive userdata.",
																							Attributes: map[string]schema.Attribute{
																								"name": schema.StringAttribute{
																									Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																									MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},
																							},
																							Required: false,
																							Optional: true,
																							Computed: false,
																						},

																						"user_data": schema.StringAttribute{
																							Description:         "UserData contains config drive inline cloud-init userdata.",
																							MarkdownDescription: "UserData contains config drive inline cloud-init userdata.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"user_data_base64": schema.StringAttribute{
																							Description:         "UserDataBase64 contains config drive cloud-init userdata as a base64 encoded string.",
																							MarkdownDescription: "UserDataBase64 contains config drive cloud-init userdata as a base64 encoded string.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"cloud_init_no_cloud": schema.SingleNestedAttribute{
																					Description:         "CloudInitNoCloud represents a cloud-init NoCloud user-data source.The NoCloud data will be added as a disk to the vmi. A proper cloud-init installation is required inside the guest.More info: http://cloudinit.readthedocs.io/en/latest/topics/datasources/nocloud.html",
																					MarkdownDescription: "CloudInitNoCloud represents a cloud-init NoCloud user-data source.The NoCloud data will be added as a disk to the vmi. A proper cloud-init installation is required inside the guest.More info: http://cloudinit.readthedocs.io/en/latest/topics/datasources/nocloud.html",
																					Attributes: map[string]schema.Attribute{
																						"network_data": schema.StringAttribute{
																							Description:         "NetworkData contains NoCloud inline cloud-init networkdata.",
																							MarkdownDescription: "NetworkData contains NoCloud inline cloud-init networkdata.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"network_data_base64": schema.StringAttribute{
																							Description:         "NetworkDataBase64 contains NoCloud cloud-init networkdata as a base64 encoded string.",
																							MarkdownDescription: "NetworkDataBase64 contains NoCloud cloud-init networkdata as a base64 encoded string.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"network_data_secret_ref": schema.SingleNestedAttribute{
																							Description:         "NetworkDataSecretRef references a k8s secret that contains NoCloud networkdata.",
																							MarkdownDescription: "NetworkDataSecretRef references a k8s secret that contains NoCloud networkdata.",
																							Attributes: map[string]schema.Attribute{
																								"name": schema.StringAttribute{
																									Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																									MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},
																							},
																							Required: false,
																							Optional: true,
																							Computed: false,
																						},

																						"secret_ref": schema.SingleNestedAttribute{
																							Description:         "UserDataSecretRef references a k8s secret that contains NoCloud userdata.",
																							MarkdownDescription: "UserDataSecretRef references a k8s secret that contains NoCloud userdata.",
																							Attributes: map[string]schema.Attribute{
																								"name": schema.StringAttribute{
																									Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																									MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},
																							},
																							Required: false,
																							Optional: true,
																							Computed: false,
																						},

																						"user_data": schema.StringAttribute{
																							Description:         "UserData contains NoCloud inline cloud-init userdata.",
																							MarkdownDescription: "UserData contains NoCloud inline cloud-init userdata.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"user_data_base64": schema.StringAttribute{
																							Description:         "UserDataBase64 contains NoCloud cloud-init userdata as a base64 encoded string.",
																							MarkdownDescription: "UserDataBase64 contains NoCloud cloud-init userdata as a base64 encoded string.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"config_map": schema.SingleNestedAttribute{
																					Description:         "ConfigMapSource represents a reference to a ConfigMap in the same namespace.More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/",
																					MarkdownDescription: "ConfigMapSource represents a reference to a ConfigMap in the same namespace.More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/",
																					Attributes: map[string]schema.Attribute{
																						"name": schema.StringAttribute{
																							Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																							MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"optional": schema.BoolAttribute{
																							Description:         "Specify whether the ConfigMap or it's keys must be defined",
																							MarkdownDescription: "Specify whether the ConfigMap or it's keys must be defined",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"volume_label": schema.StringAttribute{
																							Description:         "The volume label of the resulting disk inside the VMI.Different bootstrapping mechanisms require different values.Typical values are 'cidata' (cloud-init), 'config-2' (cloud-init) or 'OEMDRV' (kickstart).",
																							MarkdownDescription: "The volume label of the resulting disk inside the VMI.Different bootstrapping mechanisms require different values.Typical values are 'cidata' (cloud-init), 'config-2' (cloud-init) or 'OEMDRV' (kickstart).",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"container_disk": schema.SingleNestedAttribute{
																					Description:         "ContainerDisk references a docker image, embedding a qcow or raw disk.More info: https://kubevirt.gitbooks.io/user-guide/registry-disk.html",
																					MarkdownDescription: "ContainerDisk references a docker image, embedding a qcow or raw disk.More info: https://kubevirt.gitbooks.io/user-guide/registry-disk.html",
																					Attributes: map[string]schema.Attribute{
																						"image": schema.StringAttribute{
																							Description:         "Image is the name of the image with the embedded disk.",
																							MarkdownDescription: "Image is the name of the image with the embedded disk.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"image_pull_policy": schema.StringAttribute{
																							Description:         "Image pull policy.One of Always, Never, IfNotPresent.Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.Cannot be updated.More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
																							MarkdownDescription: "Image pull policy.One of Always, Never, IfNotPresent.Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.Cannot be updated.More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"image_pull_secret": schema.StringAttribute{
																							Description:         "ImagePullSecret is the name of the Docker registry secret required to pull the image. The secret must already exist.",
																							MarkdownDescription: "ImagePullSecret is the name of the Docker registry secret required to pull the image. The secret must already exist.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"path": schema.StringAttribute{
																							Description:         "Path defines the path to disk file in the container",
																							MarkdownDescription: "Path defines the path to disk file in the container",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"data_volume": schema.SingleNestedAttribute{
																					Description:         "DataVolume represents the dynamic creation a PVC for this volume as well asthe process of populating that PVC with a disk image.",
																					MarkdownDescription: "DataVolume represents the dynamic creation a PVC for this volume as well asthe process of populating that PVC with a disk image.",
																					Attributes: map[string]schema.Attribute{
																						"hotpluggable": schema.BoolAttribute{
																							Description:         "Hotpluggable indicates whether the volume can be hotplugged and hotunplugged.",
																							MarkdownDescription: "Hotpluggable indicates whether the volume can be hotplugged and hotunplugged.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"name": schema.StringAttribute{
																							Description:         "Name of both the DataVolume and the PVC in the same namespace.After PVC population the DataVolume is garbage collected by default.",
																							MarkdownDescription: "Name of both the DataVolume and the PVC in the same namespace.After PVC population the DataVolume is garbage collected by default.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"downward_api": schema.SingleNestedAttribute{
																					Description:         "DownwardAPI represents downward API about the pod that should populate this volume",
																					MarkdownDescription: "DownwardAPI represents downward API about the pod that should populate this volume",
																					Attributes: map[string]schema.Attribute{
																						"fields": schema.ListNestedAttribute{
																							Description:         "Fields is a list of downward API volume file",
																							MarkdownDescription: "Fields is a list of downward API volume file",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"field_ref": schema.SingleNestedAttribute{
																										Description:         "Required: Selects a field of the pod: only annotations, labels, name, namespace and uid are supported.",
																										MarkdownDescription: "Required: Selects a field of the pod: only annotations, labels, name, namespace and uid are supported.",
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

																									"mode": schema.Int64Attribute{
																										Description:         "Optional: mode bits used to set permissions on this file, must be an octal valuebetween 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																										MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal valuebetween 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"path": schema.StringAttribute{
																										Description:         "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																										MarkdownDescription: "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"resource_field_ref": schema.SingleNestedAttribute{
																										Description:         "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																										MarkdownDescription: "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
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
																								},
																							},
																							Required: false,
																							Optional: true,
																							Computed: false,
																						},

																						"volume_label": schema.StringAttribute{
																							Description:         "The volume label of the resulting disk inside the VMI.Different bootstrapping mechanisms require different values.Typical values are 'cidata' (cloud-init), 'config-2' (cloud-init) or 'OEMDRV' (kickstart).",
																							MarkdownDescription: "The volume label of the resulting disk inside the VMI.Different bootstrapping mechanisms require different values.Typical values are 'cidata' (cloud-init), 'config-2' (cloud-init) or 'OEMDRV' (kickstart).",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"downward_metrics": schema.MapAttribute{
																					Description:         "DownwardMetrics adds a very small disk to VMIs which contains a limited view of host and guestmetrics. The disk content is compatible with vhostmd (https://github.com/vhostmd/vhostmd) and vm-dump-metrics.",
																					MarkdownDescription: "DownwardMetrics adds a very small disk to VMIs which contains a limited view of host and guestmetrics. The disk content is compatible with vhostmd (https://github.com/vhostmd/vhostmd) and vm-dump-metrics.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"empty_disk": schema.SingleNestedAttribute{
																					Description:         "EmptyDisk represents a temporary disk which shares the vmis lifecycle.More info: https://kubevirt.gitbooks.io/user-guide/disks-and-volumes.html",
																					MarkdownDescription: "EmptyDisk represents a temporary disk which shares the vmis lifecycle.More info: https://kubevirt.gitbooks.io/user-guide/disks-and-volumes.html",
																					Attributes: map[string]schema.Attribute{
																						"capacity": schema.StringAttribute{
																							Description:         "Capacity of the sparse disk.",
																							MarkdownDescription: "Capacity of the sparse disk.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"ephemeral": schema.SingleNestedAttribute{
																					Description:         "Ephemeral is a special volume source that 'wraps' specified source and provides copy-on-write image on top of it.",
																					MarkdownDescription: "Ephemeral is a special volume source that 'wraps' specified source and provides copy-on-write image on top of it.",
																					Attributes: map[string]schema.Attribute{
																						"persistent_volume_claim": schema.SingleNestedAttribute{
																							Description:         "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace.Directly attached to the vmi via qemu.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																							MarkdownDescription: "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace.Directly attached to the vmi via qemu.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																							Attributes: map[string]schema.Attribute{
																								"claim_name": schema.StringAttribute{
																									Description:         "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																									MarkdownDescription: "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"read_only": schema.BoolAttribute{
																									Description:         "readOnly Will force the ReadOnly setting in VolumeMounts.Default false.",
																									MarkdownDescription: "readOnly Will force the ReadOnly setting in VolumeMounts.Default false.",
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

																				"host_disk": schema.SingleNestedAttribute{
																					Description:         "HostDisk represents a disk created on the cluster level",
																					MarkdownDescription: "HostDisk represents a disk created on the cluster level",
																					Attributes: map[string]schema.Attribute{
																						"capacity": schema.StringAttribute{
																							Description:         "Capacity of the sparse disk",
																							MarkdownDescription: "Capacity of the sparse disk",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"path": schema.StringAttribute{
																							Description:         "The path to HostDisk image located on the cluster",
																							MarkdownDescription: "The path to HostDisk image located on the cluster",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"shared": schema.BoolAttribute{
																							Description:         "Shared indicate whether the path is shared between nodes",
																							MarkdownDescription: "Shared indicate whether the path is shared between nodes",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"type": schema.StringAttribute{
																							Description:         "Contains information if disk.img exists or should be createdallowed options are 'Disk' and 'DiskOrCreate'",
																							MarkdownDescription: "Contains information if disk.img exists or should be createdallowed options are 'Disk' and 'DiskOrCreate'",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"memory_dump": schema.SingleNestedAttribute{
																					Description:         "MemoryDump is attached to the virt launcher and is populated with a memory dump of the vmi",
																					MarkdownDescription: "MemoryDump is attached to the virt launcher and is populated with a memory dump of the vmi",
																					Attributes: map[string]schema.Attribute{
																						"claim_name": schema.StringAttribute{
																							Description:         "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																							MarkdownDescription: "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"hotpluggable": schema.BoolAttribute{
																							Description:         "Hotpluggable indicates whether the volume can be hotplugged and hotunplugged.",
																							MarkdownDescription: "Hotpluggable indicates whether the volume can be hotplugged and hotunplugged.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"read_only": schema.BoolAttribute{
																							Description:         "readOnly Will force the ReadOnly setting in VolumeMounts.Default false.",
																							MarkdownDescription: "readOnly Will force the ReadOnly setting in VolumeMounts.Default false.",
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
																					Description:         "Volume's name.Must be a DNS_LABEL and unique within the vmi.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																					MarkdownDescription: "Volume's name.Must be a DNS_LABEL and unique within the vmi.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"persistent_volume_claim": schema.SingleNestedAttribute{
																					Description:         "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace.Directly attached to the vmi via qemu.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																					MarkdownDescription: "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace.Directly attached to the vmi via qemu.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																					Attributes: map[string]schema.Attribute{
																						"claim_name": schema.StringAttribute{
																							Description:         "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																							MarkdownDescription: "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"hotpluggable": schema.BoolAttribute{
																							Description:         "Hotpluggable indicates whether the volume can be hotplugged and hotunplugged.",
																							MarkdownDescription: "Hotpluggable indicates whether the volume can be hotplugged and hotunplugged.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"read_only": schema.BoolAttribute{
																							Description:         "readOnly Will force the ReadOnly setting in VolumeMounts.Default false.",
																							MarkdownDescription: "readOnly Will force the ReadOnly setting in VolumeMounts.Default false.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"secret": schema.SingleNestedAttribute{
																					Description:         "SecretVolumeSource represents a reference to a secret data in the same namespace.More info: https://kubernetes.io/docs/concepts/configuration/secret/",
																					MarkdownDescription: "SecretVolumeSource represents a reference to a secret data in the same namespace.More info: https://kubernetes.io/docs/concepts/configuration/secret/",
																					Attributes: map[string]schema.Attribute{
																						"optional": schema.BoolAttribute{
																							Description:         "Specify whether the Secret or it's keys must be defined",
																							MarkdownDescription: "Specify whether the Secret or it's keys must be defined",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"secret_name": schema.StringAttribute{
																							Description:         "Name of the secret in the pod's namespace to use.More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																							MarkdownDescription: "Name of the secret in the pod's namespace to use.More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"volume_label": schema.StringAttribute{
																							Description:         "The volume label of the resulting disk inside the VMI.Different bootstrapping mechanisms require different values.Typical values are 'cidata' (cloud-init), 'config-2' (cloud-init) or 'OEMDRV' (kickstart).",
																							MarkdownDescription: "The volume label of the resulting disk inside the VMI.Different bootstrapping mechanisms require different values.Typical values are 'cidata' (cloud-init), 'config-2' (cloud-init) or 'OEMDRV' (kickstart).",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"service_account": schema.SingleNestedAttribute{
																					Description:         "ServiceAccountVolumeSource represents a reference to a service account.There can only be one volume of this type!More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/",
																					MarkdownDescription: "ServiceAccountVolumeSource represents a reference to a service account.There can only be one volume of this type!More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/",
																					Attributes: map[string]schema.Attribute{
																						"service_account_name": schema.StringAttribute{
																							Description:         "Name of the service account in the pod's namespace to use.More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/",
																							MarkdownDescription: "Name of the service account in the pod's namespace to use.More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"sysprep": schema.SingleNestedAttribute{
																					Description:         "Represents a Sysprep volume source.",
																					MarkdownDescription: "Represents a Sysprep volume source.",
																					Attributes: map[string]schema.Attribute{
																						"config_map": schema.SingleNestedAttribute{
																							Description:         "ConfigMap references a ConfigMap that contains Sysprep answer file named autounattend.xml that should be attached as disk of CDROM type.",
																							MarkdownDescription: "ConfigMap references a ConfigMap that contains Sysprep answer file named autounattend.xml that should be attached as disk of CDROM type.",
																							Attributes: map[string]schema.Attribute{
																								"name": schema.StringAttribute{
																									Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																									MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},
																							},
																							Required: false,
																							Optional: true,
																							Computed: false,
																						},

																						"secret": schema.SingleNestedAttribute{
																							Description:         "Secret references a k8s Secret that contains Sysprep answer file named autounattend.xml that should be attached as disk of CDROM type.",
																							MarkdownDescription: "Secret references a k8s Secret that contains Sysprep answer file named autounattend.xml that should be attached as disk of CDROM type.",
																							Attributes: map[string]schema.Attribute{
																								"name": schema.StringAttribute{
																									Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																									MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *InfrastructureClusterXK8SIoKubevirtMachineTemplateV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infrastructure_cluster_x_k8s_io_kubevirt_machine_template_v1alpha1_manifest")

	var model InfrastructureClusterXK8SIoKubevirtMachineTemplateV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1alpha1")
	model.Kind = pointer.String("KubevirtMachineTemplate")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
