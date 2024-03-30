/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ceph_rook_io_v1

import (
	"context"
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
	_ datasource.DataSource = &CephRookIoCephNfsV1Manifest{}
)

func NewCephRookIoCephNfsV1Manifest() datasource.DataSource {
	return &CephRookIoCephNfsV1Manifest{}
}

type CephRookIoCephNfsV1Manifest struct{}

type CephRookIoCephNfsV1ManifestData struct {
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
		Rados *struct {
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Pool      *string `tfsdk:"pool" json:"pool,omitempty"`
		} `tfsdk:"rados" json:"rados,omitempty"`
		Security *struct {
			Kerberos *struct {
				ConfigFiles *struct {
					VolumeSource *struct {
						ConfigMap *struct {
							DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
							Items       *[]struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
								Path *string `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"items" json:"items,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"config_map" json:"configMap,omitempty"`
						EmptyDir *struct {
							Medium    *string `tfsdk:"medium" json:"medium,omitempty"`
							SizeLimit *string `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
						} `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
						HostPath *struct {
							Path *string `tfsdk:"path" json:"path,omitempty"`
							Type *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"host_path" json:"hostPath,omitempty"`
						PersistentVolumeClaim *struct {
							ClaimName *string `tfsdk:"claim_name" json:"claimName,omitempty"`
							ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						} `tfsdk:"persistent_volume_claim" json:"persistentVolumeClaim,omitempty"`
						Projected *struct {
							DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
							Sources     *[]struct {
								ClusterTrustBundle *struct {
									LabelSelector *struct {
										MatchExpressions *[]struct {
											Key      *string   `tfsdk:"key" json:"key,omitempty"`
											Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
											Values   *[]string `tfsdk:"values" json:"values,omitempty"`
										} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
										MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
									Name       *string `tfsdk:"name" json:"name,omitempty"`
									Optional   *bool   `tfsdk:"optional" json:"optional,omitempty"`
									Path       *string `tfsdk:"path" json:"path,omitempty"`
									SignerName *string `tfsdk:"signer_name" json:"signerName,omitempty"`
								} `tfsdk:"cluster_trust_bundle" json:"clusterTrustBundle,omitempty"`
								ConfigMap *struct {
									Items *[]struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
										Path *string `tfsdk:"path" json:"path,omitempty"`
									} `tfsdk:"items" json:"items,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"config_map" json:"configMap,omitempty"`
								DownwardAPI *struct {
									Items *[]struct {
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
									} `tfsdk:"items" json:"items,omitempty"`
								} `tfsdk:"downward_api" json:"downwardAPI,omitempty"`
								Secret *struct {
									Items *[]struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
										Path *string `tfsdk:"path" json:"path,omitempty"`
									} `tfsdk:"items" json:"items,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"secret" json:"secret,omitempty"`
								ServiceAccountToken *struct {
									Audience          *string `tfsdk:"audience" json:"audience,omitempty"`
									ExpirationSeconds *int64  `tfsdk:"expiration_seconds" json:"expirationSeconds,omitempty"`
									Path              *string `tfsdk:"path" json:"path,omitempty"`
								} `tfsdk:"service_account_token" json:"serviceAccountToken,omitempty"`
							} `tfsdk:"sources" json:"sources,omitempty"`
						} `tfsdk:"projected" json:"projected,omitempty"`
						Secret *struct {
							DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
							Items       *[]struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
								Path *string `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"items" json:"items,omitempty"`
							Optional   *bool   `tfsdk:"optional" json:"optional,omitempty"`
							SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
						} `tfsdk:"secret" json:"secret,omitempty"`
					} `tfsdk:"volume_source" json:"volumeSource,omitempty"`
				} `tfsdk:"config_files" json:"configFiles,omitempty"`
				DomainName *string `tfsdk:"domain_name" json:"domainName,omitempty"`
				KeytabFile *struct {
					VolumeSource *struct {
						ConfigMap *struct {
							DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
							Items       *[]struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
								Path *string `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"items" json:"items,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"config_map" json:"configMap,omitempty"`
						EmptyDir *struct {
							Medium    *string `tfsdk:"medium" json:"medium,omitempty"`
							SizeLimit *string `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
						} `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
						HostPath *struct {
							Path *string `tfsdk:"path" json:"path,omitempty"`
							Type *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"host_path" json:"hostPath,omitempty"`
						PersistentVolumeClaim *struct {
							ClaimName *string `tfsdk:"claim_name" json:"claimName,omitempty"`
							ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						} `tfsdk:"persistent_volume_claim" json:"persistentVolumeClaim,omitempty"`
						Projected *struct {
							DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
							Sources     *[]struct {
								ClusterTrustBundle *struct {
									LabelSelector *struct {
										MatchExpressions *[]struct {
											Key      *string   `tfsdk:"key" json:"key,omitempty"`
											Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
											Values   *[]string `tfsdk:"values" json:"values,omitempty"`
										} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
										MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
									Name       *string `tfsdk:"name" json:"name,omitempty"`
									Optional   *bool   `tfsdk:"optional" json:"optional,omitempty"`
									Path       *string `tfsdk:"path" json:"path,omitempty"`
									SignerName *string `tfsdk:"signer_name" json:"signerName,omitempty"`
								} `tfsdk:"cluster_trust_bundle" json:"clusterTrustBundle,omitempty"`
								ConfigMap *struct {
									Items *[]struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
										Path *string `tfsdk:"path" json:"path,omitempty"`
									} `tfsdk:"items" json:"items,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"config_map" json:"configMap,omitempty"`
								DownwardAPI *struct {
									Items *[]struct {
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
									} `tfsdk:"items" json:"items,omitempty"`
								} `tfsdk:"downward_api" json:"downwardAPI,omitempty"`
								Secret *struct {
									Items *[]struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
										Path *string `tfsdk:"path" json:"path,omitempty"`
									} `tfsdk:"items" json:"items,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"secret" json:"secret,omitempty"`
								ServiceAccountToken *struct {
									Audience          *string `tfsdk:"audience" json:"audience,omitempty"`
									ExpirationSeconds *int64  `tfsdk:"expiration_seconds" json:"expirationSeconds,omitempty"`
									Path              *string `tfsdk:"path" json:"path,omitempty"`
								} `tfsdk:"service_account_token" json:"serviceAccountToken,omitempty"`
							} `tfsdk:"sources" json:"sources,omitempty"`
						} `tfsdk:"projected" json:"projected,omitempty"`
						Secret *struct {
							DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
							Items       *[]struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
								Path *string `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"items" json:"items,omitempty"`
							Optional   *bool   `tfsdk:"optional" json:"optional,omitempty"`
							SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
						} `tfsdk:"secret" json:"secret,omitempty"`
					} `tfsdk:"volume_source" json:"volumeSource,omitempty"`
				} `tfsdk:"keytab_file" json:"keytabFile,omitempty"`
				PrincipalName *string `tfsdk:"principal_name" json:"principalName,omitempty"`
			} `tfsdk:"kerberos" json:"kerberos,omitempty"`
			Sssd *struct {
				Sidecar *struct {
					AdditionalFiles *[]struct {
						SubPath      *string `tfsdk:"sub_path" json:"subPath,omitempty"`
						VolumeSource *struct {
							ConfigMap *struct {
								DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
								Items       *[]struct {
									Key  *string `tfsdk:"key" json:"key,omitempty"`
									Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
									Path *string `tfsdk:"path" json:"path,omitempty"`
								} `tfsdk:"items" json:"items,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"config_map" json:"configMap,omitempty"`
							EmptyDir *struct {
								Medium    *string `tfsdk:"medium" json:"medium,omitempty"`
								SizeLimit *string `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
							} `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
							HostPath *struct {
								Path *string `tfsdk:"path" json:"path,omitempty"`
								Type *string `tfsdk:"type" json:"type,omitempty"`
							} `tfsdk:"host_path" json:"hostPath,omitempty"`
							PersistentVolumeClaim *struct {
								ClaimName *string `tfsdk:"claim_name" json:"claimName,omitempty"`
								ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
							} `tfsdk:"persistent_volume_claim" json:"persistentVolumeClaim,omitempty"`
							Projected *struct {
								DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
								Sources     *[]struct {
									ClusterTrustBundle *struct {
										LabelSelector *struct {
											MatchExpressions *[]struct {
												Key      *string   `tfsdk:"key" json:"key,omitempty"`
												Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
												Values   *[]string `tfsdk:"values" json:"values,omitempty"`
											} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
											MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
										} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
										Name       *string `tfsdk:"name" json:"name,omitempty"`
										Optional   *bool   `tfsdk:"optional" json:"optional,omitempty"`
										Path       *string `tfsdk:"path" json:"path,omitempty"`
										SignerName *string `tfsdk:"signer_name" json:"signerName,omitempty"`
									} `tfsdk:"cluster_trust_bundle" json:"clusterTrustBundle,omitempty"`
									ConfigMap *struct {
										Items *[]struct {
											Key  *string `tfsdk:"key" json:"key,omitempty"`
											Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
											Path *string `tfsdk:"path" json:"path,omitempty"`
										} `tfsdk:"items" json:"items,omitempty"`
										Name     *string `tfsdk:"name" json:"name,omitempty"`
										Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
									} `tfsdk:"config_map" json:"configMap,omitempty"`
									DownwardAPI *struct {
										Items *[]struct {
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
										} `tfsdk:"items" json:"items,omitempty"`
									} `tfsdk:"downward_api" json:"downwardAPI,omitempty"`
									Secret *struct {
										Items *[]struct {
											Key  *string `tfsdk:"key" json:"key,omitempty"`
											Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
											Path *string `tfsdk:"path" json:"path,omitempty"`
										} `tfsdk:"items" json:"items,omitempty"`
										Name     *string `tfsdk:"name" json:"name,omitempty"`
										Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
									} `tfsdk:"secret" json:"secret,omitempty"`
									ServiceAccountToken *struct {
										Audience          *string `tfsdk:"audience" json:"audience,omitempty"`
										ExpirationSeconds *int64  `tfsdk:"expiration_seconds" json:"expirationSeconds,omitempty"`
										Path              *string `tfsdk:"path" json:"path,omitempty"`
									} `tfsdk:"service_account_token" json:"serviceAccountToken,omitempty"`
								} `tfsdk:"sources" json:"sources,omitempty"`
							} `tfsdk:"projected" json:"projected,omitempty"`
							Secret *struct {
								DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
								Items       *[]struct {
									Key  *string `tfsdk:"key" json:"key,omitempty"`
									Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
									Path *string `tfsdk:"path" json:"path,omitempty"`
								} `tfsdk:"items" json:"items,omitempty"`
								Optional   *bool   `tfsdk:"optional" json:"optional,omitempty"`
								SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
							} `tfsdk:"secret" json:"secret,omitempty"`
						} `tfsdk:"volume_source" json:"volumeSource,omitempty"`
					} `tfsdk:"additional_files" json:"additionalFiles,omitempty"`
					DebugLevel *int64  `tfsdk:"debug_level" json:"debugLevel,omitempty"`
					Image      *string `tfsdk:"image" json:"image,omitempty"`
					Resources  *struct {
						Claims *[]struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"claims" json:"claims,omitempty"`
						Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
						Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					SssdConfigFile *struct {
						VolumeSource *struct {
							ConfigMap *struct {
								DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
								Items       *[]struct {
									Key  *string `tfsdk:"key" json:"key,omitempty"`
									Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
									Path *string `tfsdk:"path" json:"path,omitempty"`
								} `tfsdk:"items" json:"items,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"config_map" json:"configMap,omitempty"`
							EmptyDir *struct {
								Medium    *string `tfsdk:"medium" json:"medium,omitempty"`
								SizeLimit *string `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
							} `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
							HostPath *struct {
								Path *string `tfsdk:"path" json:"path,omitempty"`
								Type *string `tfsdk:"type" json:"type,omitempty"`
							} `tfsdk:"host_path" json:"hostPath,omitempty"`
							PersistentVolumeClaim *struct {
								ClaimName *string `tfsdk:"claim_name" json:"claimName,omitempty"`
								ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
							} `tfsdk:"persistent_volume_claim" json:"persistentVolumeClaim,omitempty"`
							Projected *struct {
								DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
								Sources     *[]struct {
									ClusterTrustBundle *struct {
										LabelSelector *struct {
											MatchExpressions *[]struct {
												Key      *string   `tfsdk:"key" json:"key,omitempty"`
												Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
												Values   *[]string `tfsdk:"values" json:"values,omitempty"`
											} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
											MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
										} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
										Name       *string `tfsdk:"name" json:"name,omitempty"`
										Optional   *bool   `tfsdk:"optional" json:"optional,omitempty"`
										Path       *string `tfsdk:"path" json:"path,omitempty"`
										SignerName *string `tfsdk:"signer_name" json:"signerName,omitempty"`
									} `tfsdk:"cluster_trust_bundle" json:"clusterTrustBundle,omitempty"`
									ConfigMap *struct {
										Items *[]struct {
											Key  *string `tfsdk:"key" json:"key,omitempty"`
											Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
											Path *string `tfsdk:"path" json:"path,omitempty"`
										} `tfsdk:"items" json:"items,omitempty"`
										Name     *string `tfsdk:"name" json:"name,omitempty"`
										Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
									} `tfsdk:"config_map" json:"configMap,omitempty"`
									DownwardAPI *struct {
										Items *[]struct {
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
										} `tfsdk:"items" json:"items,omitempty"`
									} `tfsdk:"downward_api" json:"downwardAPI,omitempty"`
									Secret *struct {
										Items *[]struct {
											Key  *string `tfsdk:"key" json:"key,omitempty"`
											Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
											Path *string `tfsdk:"path" json:"path,omitempty"`
										} `tfsdk:"items" json:"items,omitempty"`
										Name     *string `tfsdk:"name" json:"name,omitempty"`
										Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
									} `tfsdk:"secret" json:"secret,omitempty"`
									ServiceAccountToken *struct {
										Audience          *string `tfsdk:"audience" json:"audience,omitempty"`
										ExpirationSeconds *int64  `tfsdk:"expiration_seconds" json:"expirationSeconds,omitempty"`
										Path              *string `tfsdk:"path" json:"path,omitempty"`
									} `tfsdk:"service_account_token" json:"serviceAccountToken,omitempty"`
								} `tfsdk:"sources" json:"sources,omitempty"`
							} `tfsdk:"projected" json:"projected,omitempty"`
							Secret *struct {
								DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
								Items       *[]struct {
									Key  *string `tfsdk:"key" json:"key,omitempty"`
									Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
									Path *string `tfsdk:"path" json:"path,omitempty"`
								} `tfsdk:"items" json:"items,omitempty"`
								Optional   *bool   `tfsdk:"optional" json:"optional,omitempty"`
								SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
							} `tfsdk:"secret" json:"secret,omitempty"`
						} `tfsdk:"volume_source" json:"volumeSource,omitempty"`
					} `tfsdk:"sssd_config_file" json:"sssdConfigFile,omitempty"`
				} `tfsdk:"sidecar" json:"sidecar,omitempty"`
			} `tfsdk:"sssd" json:"sssd,omitempty"`
		} `tfsdk:"security" json:"security,omitempty"`
		Server *struct {
			Active        *int64             `tfsdk:"active" json:"active,omitempty"`
			Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			HostNetwork   *bool              `tfsdk:"host_network" json:"hostNetwork,omitempty"`
			Labels        *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			LivenessProbe *struct {
				Disabled *bool `tfsdk:"disabled" json:"disabled,omitempty"`
				Probe    *struct {
					Exec *struct {
						Command *[]string `tfsdk:"command" json:"command,omitempty"`
					} `tfsdk:"exec" json:"exec,omitempty"`
					FailureThreshold *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
					Grpc             *struct {
						Port    *int64  `tfsdk:"port" json:"port,omitempty"`
						Service *string `tfsdk:"service" json:"service,omitempty"`
					} `tfsdk:"grpc" json:"grpc,omitempty"`
					HttpGet *struct {
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
					TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
					TimeoutSeconds                *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				} `tfsdk:"probe" json:"probe,omitempty"`
			} `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
			LogLevel  *string `tfsdk:"log_level" json:"logLevel,omitempty"`
			Placement *struct {
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
				Tolerations *[]struct {
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
			} `tfsdk:"placement" json:"placement,omitempty"`
			PriorityClassName *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
			Resources         *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
		} `tfsdk:"server" json:"server,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CephRookIoCephNfsV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ceph_rook_io_ceph_nfs_v1_manifest"
}

func (r *CephRookIoCephNfsV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CephNFS represents a Ceph NFS",
		MarkdownDescription: "CephNFS represents a Ceph NFS",
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
				Description:         "NFSGaneshaSpec represents the spec of an nfs ganesha server",
				MarkdownDescription: "NFSGaneshaSpec represents the spec of an nfs ganesha server",
				Attributes: map[string]schema.Attribute{
					"rados": schema.SingleNestedAttribute{
						Description:         "RADOS is the Ganesha RADOS specification",
						MarkdownDescription: "RADOS is the Ganesha RADOS specification",
						Attributes: map[string]schema.Attribute{
							"namespace": schema.StringAttribute{
								Description:         "The namespace inside the Ceph pool (set by 'pool') where shared NFS-Ganesha config is stored. This setting is deprecated as it is internally set to the name of the CephNFS.",
								MarkdownDescription: "The namespace inside the Ceph pool (set by 'pool') where shared NFS-Ganesha config is stored. This setting is deprecated as it is internally set to the name of the CephNFS.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pool": schema.StringAttribute{
								Description:         "The Ceph pool used store the shared configuration for NFS-Ganesha daemons. This setting is deprecated, as it is internally required to be '.nfs'.",
								MarkdownDescription: "The Ceph pool used store the shared configuration for NFS-Ganesha daemons. This setting is deprecated, as it is internally required to be '.nfs'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"security": schema.SingleNestedAttribute{
						Description:         "Security allows specifying security configurations for the NFS cluster",
						MarkdownDescription: "Security allows specifying security configurations for the NFS cluster",
						Attributes: map[string]schema.Attribute{
							"kerberos": schema.SingleNestedAttribute{
								Description:         "Kerberos configures NFS-Ganesha to secure NFS client connections with Kerberos.",
								MarkdownDescription: "Kerberos configures NFS-Ganesha to secure NFS client connections with Kerberos.",
								Attributes: map[string]schema.Attribute{
									"config_files": schema.SingleNestedAttribute{
										Description:         "ConfigFiles defines where the Kerberos configuration should be sourced from. Config files will be placed into the '/etc/krb5.conf.rook/' directory.  If this is left empty, Rook will not add any files. This allows you to manage the files yourself however you wish. For example, you may build them into your custom Ceph container image or use the Vault agent injector to securely add the files via annotations on the CephNFS spec (passed to the NFS server pods).  Rook configures Kerberos to log to stderr. We suggest removing logging sections from config files to avoid consuming unnecessary disk space from logging to files.",
										MarkdownDescription: "ConfigFiles defines where the Kerberos configuration should be sourced from. Config files will be placed into the '/etc/krb5.conf.rook/' directory.  If this is left empty, Rook will not add any files. This allows you to manage the files yourself however you wish. For example, you may build them into your custom Ceph container image or use the Vault agent injector to securely add the files via annotations on the CephNFS spec (passed to the NFS server pods).  Rook configures Kerberos to log to stderr. We suggest removing logging sections from config files to avoid consuming unnecessary disk space from logging to files.",
										Attributes: map[string]schema.Attribute{
											"volume_source": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"config_map": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"default_mode": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"items": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"mode": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"optional": schema.BoolAttribute{
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

													"empty_dir": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"medium": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"size_limit": schema.StringAttribute{
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

													"host_path": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"path": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"type": schema.StringAttribute{
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

													"persistent_volume_claim": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"claim_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"read_only": schema.BoolAttribute{
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

													"projected": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"default_mode": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"sources": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"cluster_trust_bundle": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"label_selector": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"match_expressions": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
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
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"values": schema.ListAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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

																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"optional": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"path": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"signer_name": schema.StringAttribute{
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

																		"config_map": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"items": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"mode": schema.Int64Attribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"path": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
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

																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"optional": schema.BoolAttribute{
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

																		"downward_api": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"items": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"field_ref": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"api_version": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"field_path": schema.StringAttribute{
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

																							"mode": schema.Int64Attribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"path": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"resource_field_ref": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"container_name": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"divisor": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"resource": schema.StringAttribute{
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
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"items": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"mode": schema.Int64Attribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"path": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
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

																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"optional": schema.BoolAttribute{
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

																		"service_account_token": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"audience": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"expiration_seconds": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"path": schema.StringAttribute{
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
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"default_mode": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"items": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"mode": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

															"optional": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_name": schema.StringAttribute{
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

									"domain_name": schema.StringAttribute{
										Description:         "DomainName should be set to the Kerberos Realm.",
										MarkdownDescription: "DomainName should be set to the Kerberos Realm.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keytab_file": schema.SingleNestedAttribute{
										Description:         "KeytabFile defines where the Kerberos keytab should be sourced from. The keytab file will be placed into '/etc/krb5.keytab'. If this is left empty, Rook will not add the file. This allows you to manage the 'krb5.keytab' file yourself however you wish. For example, you may build it into your custom Ceph container image or use the Vault agent injector to securely add the file via annotations on the CephNFS spec (passed to the NFS server pods).",
										MarkdownDescription: "KeytabFile defines where the Kerberos keytab should be sourced from. The keytab file will be placed into '/etc/krb5.keytab'. If this is left empty, Rook will not add the file. This allows you to manage the 'krb5.keytab' file yourself however you wish. For example, you may build it into your custom Ceph container image or use the Vault agent injector to securely add the file via annotations on the CephNFS spec (passed to the NFS server pods).",
										Attributes: map[string]schema.Attribute{
											"volume_source": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"config_map": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"default_mode": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"items": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"mode": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"optional": schema.BoolAttribute{
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

													"empty_dir": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"medium": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"size_limit": schema.StringAttribute{
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

													"host_path": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"path": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"type": schema.StringAttribute{
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

													"persistent_volume_claim": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"claim_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"read_only": schema.BoolAttribute{
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

													"projected": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"default_mode": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"sources": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"cluster_trust_bundle": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"label_selector": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"match_expressions": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
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
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"values": schema.ListAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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

																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"optional": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"path": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"signer_name": schema.StringAttribute{
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

																		"config_map": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"items": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"mode": schema.Int64Attribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"path": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
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

																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"optional": schema.BoolAttribute{
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

																		"downward_api": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"items": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"field_ref": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"api_version": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"field_path": schema.StringAttribute{
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

																							"mode": schema.Int64Attribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"path": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"resource_field_ref": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"container_name": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"divisor": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"resource": schema.StringAttribute{
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
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"items": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"mode": schema.Int64Attribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"path": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
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

																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"optional": schema.BoolAttribute{
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

																		"service_account_token": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"audience": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"expiration_seconds": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"path": schema.StringAttribute{
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
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"default_mode": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"items": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"mode": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

															"optional": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_name": schema.StringAttribute{
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

									"principal_name": schema.StringAttribute{
										Description:         "PrincipalName corresponds directly to NFS-Ganesha's NFS_KRB5:PrincipalName config. In practice, this is the service prefix of the principal name. The default is 'nfs'. This value is combined with (a) the namespace and name of the CephNFS (with a hyphen between) and (b) the Realm configured in the user-provided krb5.conf to determine the full principal name: <principalName>/<namespace>-<name>@<realm>. e.g., nfs/rook-ceph-my-nfs@example.net. See https://github.com/nfs-ganesha/nfs-ganesha/wiki/RPCSEC_GSS for more detail.",
										MarkdownDescription: "PrincipalName corresponds directly to NFS-Ganesha's NFS_KRB5:PrincipalName config. In practice, this is the service prefix of the principal name. The default is 'nfs'. This value is combined with (a) the namespace and name of the CephNFS (with a hyphen between) and (b) the Realm configured in the user-provided krb5.conf to determine the full principal name: <principalName>/<namespace>-<name>@<realm>. e.g., nfs/rook-ceph-my-nfs@example.net. See https://github.com/nfs-ganesha/nfs-ganesha/wiki/RPCSEC_GSS for more detail.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"sssd": schema.SingleNestedAttribute{
								Description:         "SSSD enables integration with System Security Services Daemon (SSSD). SSSD can be used to provide user ID mapping from a number of sources. See https://sssd.io for more information about the SSSD project.",
								MarkdownDescription: "SSSD enables integration with System Security Services Daemon (SSSD). SSSD can be used to provide user ID mapping from a number of sources. See https://sssd.io for more information about the SSSD project.",
								Attributes: map[string]schema.Attribute{
									"sidecar": schema.SingleNestedAttribute{
										Description:         "Sidecar tells Rook to run SSSD in a sidecar alongside the NFS-Ganesha server in each NFS pod.",
										MarkdownDescription: "Sidecar tells Rook to run SSSD in a sidecar alongside the NFS-Ganesha server in each NFS pod.",
										Attributes: map[string]schema.Attribute{
											"additional_files": schema.ListNestedAttribute{
												Description:         "AdditionalFiles defines any number of additional files that should be mounted into the SSSD sidecar. These files may be referenced by the sssd.conf config file.",
												MarkdownDescription: "AdditionalFiles defines any number of additional files that should be mounted into the SSSD sidecar. These files may be referenced by the sssd.conf config file.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"sub_path": schema.StringAttribute{
															Description:         "SubPath defines the sub-path in '/etc/sssd/rook-additional/' where the additional file(s) will be placed. Each subPath definition must be unique and must not contain ':'.",
															MarkdownDescription: "SubPath defines the sub-path in '/etc/sssd/rook-additional/' where the additional file(s) will be placed. Each subPath definition must be unique and must not contain ':'.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.RegexMatches(regexp.MustCompile(`^[^:]+$`), ""),
															},
														},

														"volume_source": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"config_map": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"default_mode": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"items": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"mode": schema.Int64Attribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"path": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																		"name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"optional": schema.BoolAttribute{
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

																"empty_dir": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"medium": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"size_limit": schema.StringAttribute{
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

																"host_path": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"path": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"type": schema.StringAttribute{
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

																"persistent_volume_claim": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"claim_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
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

																"projected": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"default_mode": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"sources": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"cluster_trust_bundle": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"label_selector": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"match_expressions": schema.ListNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            true,
																													Optional:            false,
																													Computed:            false,
																												},

																												"values": schema.ListAttribute{
																													Description:         "",
																													MarkdownDescription: "",
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

																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"optional": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"path": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"signer_name": schema.StringAttribute{
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

																					"config_map": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"items": schema.ListNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"key": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            true,
																											Optional:            false,
																											Computed:            false,
																										},

																										"mode": schema.Int64Attribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"path": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
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

																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"optional": schema.BoolAttribute{
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

																					"downward_api": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"items": schema.ListNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"field_ref": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"api_version": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            true,
																													Computed:            false,
																												},

																												"field_path": schema.StringAttribute{
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

																										"mode": schema.Int64Attribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"path": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            true,
																											Optional:            false,
																											Computed:            false,
																										},

																										"resource_field_ref": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"container_name": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            true,
																													Computed:            false,
																												},

																												"divisor": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            true,
																													Computed:            false,
																												},

																												"resource": schema.StringAttribute{
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
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"items": schema.ListNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"key": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            true,
																											Optional:            false,
																											Computed:            false,
																										},

																										"mode": schema.Int64Attribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"path": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
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

																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"optional": schema.BoolAttribute{
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

																					"service_account_token": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"audience": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"expiration_seconds": schema.Int64Attribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"path": schema.StringAttribute{
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
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"default_mode": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"items": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"mode": schema.Int64Attribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"path": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																		"optional": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_name": schema.StringAttribute{
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

											"debug_level": schema.Int64Attribute{
												Description:         "DebugLevel sets the debug level for SSSD. If unset or set to 0, Rook does nothing. Otherwise, this may be a value between 1 and 10. See SSSD docs for more info: https://sssd.io/troubleshooting/basics.html#sssd-debug-logs",
												MarkdownDescription: "DebugLevel sets the debug level for SSSD. If unset or set to 0, Rook does nothing. Otherwise, this may be a value between 1 and 10. See SSSD docs for more info: https://sssd.io/troubleshooting/basics.html#sssd-debug-logs",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(10),
												},
											},

											"image": schema.StringAttribute{
												Description:         "Image defines the container image that should be used for the SSSD sidecar.",
												MarkdownDescription: "Image defines the container image that should be used for the SSSD sidecar.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},

											"resources": schema.SingleNestedAttribute{
												Description:         "Resources allow specifying resource requests/limits on the SSSD sidecar container.",
												MarkdownDescription: "Resources allow specifying resource requests/limits on the SSSD sidecar container.",
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

											"sssd_config_file": schema.SingleNestedAttribute{
												Description:         "SSSDConfigFile defines where the SSSD configuration should be sourced from. The config file will be placed into '/etc/sssd/sssd.conf'. If this is left empty, Rook will not add the file. This allows you to manage the 'sssd.conf' file yourself however you wish. For example, you may build it into your custom Ceph container image or use the Vault agent injector to securely add the file via annotations on the CephNFS spec (passed to the NFS server pods).",
												MarkdownDescription: "SSSDConfigFile defines where the SSSD configuration should be sourced from. The config file will be placed into '/etc/sssd/sssd.conf'. If this is left empty, Rook will not add the file. This allows you to manage the 'sssd.conf' file yourself however you wish. For example, you may build it into your custom Ceph container image or use the Vault agent injector to securely add the file via annotations on the CephNFS spec (passed to the NFS server pods).",
												Attributes: map[string]schema.Attribute{
													"volume_source": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"config_map": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"default_mode": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"items": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"key": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"mode": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"path": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"optional": schema.BoolAttribute{
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

															"empty_dir": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"medium": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"size_limit": schema.StringAttribute{
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

															"host_path": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"path": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"type": schema.StringAttribute{
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

															"persistent_volume_claim": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"claim_name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"read_only": schema.BoolAttribute{
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

															"projected": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"default_mode": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"sources": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"cluster_trust_bundle": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"label_selector": schema.SingleNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Attributes: map[string]schema.Attribute{
																								"match_expressions": schema.ListNestedAttribute{
																									Description:         "",
																									MarkdownDescription: "",
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
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            true,
																												Optional:            false,
																												Computed:            false,
																											},

																											"values": schema.ListAttribute{
																												Description:         "",
																												MarkdownDescription: "",
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

																						"name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"optional": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"path": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"signer_name": schema.StringAttribute{
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

																				"config_map": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"items": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"key": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"mode": schema.Int64Attribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"path": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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

																						"name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"optional": schema.BoolAttribute{
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

																				"downward_api": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"items": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"field_ref": schema.SingleNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Attributes: map[string]schema.Attribute{
																											"api_version": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            true,
																												Computed:            false,
																											},

																											"field_path": schema.StringAttribute{
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

																									"mode": schema.Int64Attribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"path": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"resource_field_ref": schema.SingleNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Attributes: map[string]schema.Attribute{
																											"container_name": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            true,
																												Computed:            false,
																											},

																											"divisor": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            true,
																												Computed:            false,
																											},

																											"resource": schema.StringAttribute{
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
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"items": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"key": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"mode": schema.Int64Attribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"path": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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

																						"name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"optional": schema.BoolAttribute{
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

																				"service_account_token": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"audience": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"expiration_seconds": schema.Int64Attribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"path": schema.StringAttribute{
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
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"default_mode": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"items": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"key": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"mode": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"path": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

																	"optional": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"secret_name": schema.StringAttribute{
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
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"server": schema.SingleNestedAttribute{
						Description:         "Server is the Ganesha Server specification",
						MarkdownDescription: "Server is the Ganesha Server specification",
						Attributes: map[string]schema.Attribute{
							"active": schema.Int64Attribute{
								Description:         "The number of active Ganesha servers",
								MarkdownDescription: "The number of active Ganesha servers",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"annotations": schema.MapAttribute{
								Description:         "The annotations-related configuration to add/set on each Pod related object.",
								MarkdownDescription: "The annotations-related configuration to add/set on each Pod related object.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host_network": schema.BoolAttribute{
								Description:         "Whether host networking is enabled for the Ganesha server. If not set, the network settings from the cluster CR will be applied.",
								MarkdownDescription: "Whether host networking is enabled for the Ganesha server. If not set, the network settings from the cluster CR will be applied.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "The labels-related configuration to add/set on each Pod related object.",
								MarkdownDescription: "The labels-related configuration to add/set on each Pod related object.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"liveness_probe": schema.SingleNestedAttribute{
								Description:         "A liveness-probe to verify that Ganesha server has valid run-time state. If LivenessProbe.Disabled is false and LivenessProbe.Probe is nil uses default probe.",
								MarkdownDescription: "A liveness-probe to verify that Ganesha server has valid run-time state. If LivenessProbe.Disabled is false and LivenessProbe.Probe is nil uses default probe.",
								Attributes: map[string]schema.Attribute{
									"disabled": schema.BoolAttribute{
										Description:         "Disabled determines whether probe is disable or not",
										MarkdownDescription: "Disabled determines whether probe is disable or not",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"probe": schema.SingleNestedAttribute{
										Description:         "Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.",
										MarkdownDescription: "Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.",
										Attributes: map[string]schema.Attribute{
											"exec": schema.SingleNestedAttribute{
												Description:         "Exec specifies the action to take.",
												MarkdownDescription: "Exec specifies the action to take.",
												Attributes: map[string]schema.Attribute{
													"command": schema.ListAttribute{
														Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
														MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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
												Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"grpc": schema.SingleNestedAttribute{
												Description:         "GRPC specifies an action involving a GRPC port.",
												MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
												Attributes: map[string]schema.Attribute{
													"port": schema.Int64Attribute{
														Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
														MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"service": schema.StringAttribute{
														Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
														MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"http_get": schema.SingleNestedAttribute{
												Description:         "HTTPGet specifies the http request to perform.",
												MarkdownDescription: "HTTPGet specifies the http request to perform.",
												Attributes: map[string]schema.Attribute{
													"host": schema.StringAttribute{
														Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
														MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
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
																	Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																	MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
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
														Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
														MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"scheme": schema.StringAttribute{
														Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
														MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
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
												Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"period_seconds": schema.Int64Attribute{
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"success_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tcp_socket": schema.SingleNestedAttribute{
												Description:         "TCPSocket specifies an action involving a TCP port.",
												MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
												Attributes: map[string]schema.Attribute{
													"host": schema.StringAttribute{
														Description:         "Optional: Host name to connect to, defaults to the pod IP.",
														MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"port": schema.StringAttribute{
														Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
														MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"termination_grace_period_seconds": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"timeout_seconds": schema.Int64Attribute{
												Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
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

							"log_level": schema.StringAttribute{
								Description:         "LogLevel set logging level",
								MarkdownDescription: "LogLevel set logging level",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"placement": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"node_affinity": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"preference": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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
																	Description:         "",
																	MarkdownDescription: "",
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
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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
															Description:         "",
															MarkdownDescription: "",
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
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"node_selector_terms": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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
																	Description:         "",
																	MarkdownDescription: "",
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
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"pod_affinity_term": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"label_selector": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"match_expressions": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																"match_label_keys": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"mismatch_label_keys": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"namespace_selector": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"match_expressions": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																"namespaces": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"topology_key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
															Description:         "",
															MarkdownDescription: "",
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
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"label_selector": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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

														"match_label_keys": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"mismatch_label_keys": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace_selector": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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

														"namespaces": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"topology_key": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"pod_affinity_term": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"label_selector": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"match_expressions": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																"match_label_keys": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"mismatch_label_keys": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"namespace_selector": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"match_expressions": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																"namespaces": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"topology_key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
															Description:         "",
															MarkdownDescription: "",
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
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"label_selector": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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

														"match_label_keys": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"mismatch_label_keys": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace_selector": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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

														"namespaces": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"topology_key": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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

									"tolerations": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"effect": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"toleration_seconds": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
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

									"topology_spread_constraints": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"label_selector": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
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
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

												"match_label_keys": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_skew": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"min_domains": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_affinity_policy": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_taints_policy": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"topology_key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"when_unsatisfiable": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

							"priority_class_name": schema.StringAttribute{
								Description:         "PriorityClassName sets the priority class on the pods",
								MarkdownDescription: "PriorityClassName sets the priority class on the pods",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Resources set resource requests and limits",
								MarkdownDescription: "Resources set resource requests and limits",
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
	}
}

func (r *CephRookIoCephNfsV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ceph_rook_io_ceph_nfs_v1_manifest")

	var model CephRookIoCephNfsV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("ceph.rook.io/v1")
	model.Kind = pointer.String("CephNFS")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
