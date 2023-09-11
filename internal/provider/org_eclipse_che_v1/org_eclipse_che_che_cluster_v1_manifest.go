/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package org_eclipse_che_v1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &OrgEclipseCheCheClusterV1Manifest{}
)

func NewOrgEclipseCheCheClusterV1Manifest() datasource.DataSource {
	return &OrgEclipseCheCheClusterV1Manifest{}
}

type OrgEclipseCheCheClusterV1Manifest struct{}

type OrgEclipseCheCheClusterV1ManifestData struct {
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
		Auth *struct {
			Debug                             *bool   `tfsdk:"debug" json:"debug,omitempty"`
			ExternalIdentityProvider          *bool   `tfsdk:"external_identity_provider" json:"externalIdentityProvider,omitempty"`
			GatewayAuthenticationSidecarImage *string `tfsdk:"gateway_authentication_sidecar_image" json:"gatewayAuthenticationSidecarImage,omitempty"`
			GatewayAuthorizationSidecarImage  *string `tfsdk:"gateway_authorization_sidecar_image" json:"gatewayAuthorizationSidecarImage,omitempty"`
			GatewayConfigBumpEnv              *[]struct {
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
			} `tfsdk:"gateway_config_bump_env" json:"gatewayConfigBumpEnv,omitempty"`
			GatewayEnv *[]struct {
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
			} `tfsdk:"gateway_env" json:"gatewayEnv,omitempty"`
			GatewayHeaderRewriteSidecarImage *string `tfsdk:"gateway_header_rewrite_sidecar_image" json:"gatewayHeaderRewriteSidecarImage,omitempty"`
			GatewayKubeRbacProxyEnv          *[]struct {
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
			} `tfsdk:"gateway_kube_rbac_proxy_env" json:"gatewayKubeRbacProxyEnv,omitempty"`
			GatewayOAuthProxyEnv *[]struct {
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
			} `tfsdk:"gateway_o_auth_proxy_env" json:"gatewayOAuthProxyEnv,omitempty"`
			IdentityProviderAdminUserName      *string `tfsdk:"identity_provider_admin_user_name" json:"identityProviderAdminUserName,omitempty"`
			IdentityProviderClientId           *string `tfsdk:"identity_provider_client_id" json:"identityProviderClientId,omitempty"`
			IdentityProviderContainerResources *struct {
				Limits *struct {
					Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
					Memory *string `tfsdk:"memory" json:"memory,omitempty"`
				} `tfsdk:"limits" json:"limits,omitempty"`
				Request *struct {
					Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
					Memory *string `tfsdk:"memory" json:"memory,omitempty"`
				} `tfsdk:"request" json:"request,omitempty"`
			} `tfsdk:"identity_provider_container_resources" json:"identityProviderContainerResources,omitempty"`
			IdentityProviderImage           *string `tfsdk:"identity_provider_image" json:"identityProviderImage,omitempty"`
			IdentityProviderImagePullPolicy *string `tfsdk:"identity_provider_image_pull_policy" json:"identityProviderImagePullPolicy,omitempty"`
			IdentityProviderIngress         *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *string            `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"identity_provider_ingress" json:"identityProviderIngress,omitempty"`
			IdentityProviderPassword         *string `tfsdk:"identity_provider_password" json:"identityProviderPassword,omitempty"`
			IdentityProviderPostgresPassword *string `tfsdk:"identity_provider_postgres_password" json:"identityProviderPostgresPassword,omitempty"`
			IdentityProviderPostgresSecret   *string `tfsdk:"identity_provider_postgres_secret" json:"identityProviderPostgresSecret,omitempty"`
			IdentityProviderRealm            *string `tfsdk:"identity_provider_realm" json:"identityProviderRealm,omitempty"`
			IdentityProviderRoute            *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Domain      *string            `tfsdk:"domain" json:"domain,omitempty"`
				Labels      *string            `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"identity_provider_route" json:"identityProviderRoute,omitempty"`
			IdentityProviderSecret    *string `tfsdk:"identity_provider_secret" json:"identityProviderSecret,omitempty"`
			IdentityProviderURL       *string `tfsdk:"identity_provider_url" json:"identityProviderURL,omitempty"`
			IdentityToken             *string `tfsdk:"identity_token" json:"identityToken,omitempty"`
			InitialOpenShiftOAuthUser *bool   `tfsdk:"initial_open_shift_o_auth_user" json:"initialOpenShiftOAuthUser,omitempty"`
			NativeUserMode            *bool   `tfsdk:"native_user_mode" json:"nativeUserMode,omitempty"`
			OAuthClientName           *string `tfsdk:"o_auth_client_name" json:"oAuthClientName,omitempty"`
			OAuthScope                *string `tfsdk:"o_auth_scope" json:"oAuthScope,omitempty"`
			OAuthSecret               *string `tfsdk:"o_auth_secret" json:"oAuthSecret,omitempty"`
			OpenShiftoAuth            *bool   `tfsdk:"open_shifto_auth" json:"openShiftoAuth,omitempty"`
			UpdateAdminPassword       *bool   `tfsdk:"update_admin_password" json:"updateAdminPassword,omitempty"`
		} `tfsdk:"auth" json:"auth,omitempty"`
		Dashboard *struct {
			Warning *string `tfsdk:"warning" json:"warning,omitempty"`
		} `tfsdk:"dashboard" json:"dashboard,omitempty"`
		Database *struct {
			ChePostgresContainerResources *struct {
				Limits *struct {
					Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
					Memory *string `tfsdk:"memory" json:"memory,omitempty"`
				} `tfsdk:"limits" json:"limits,omitempty"`
				Request *struct {
					Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
					Memory *string `tfsdk:"memory" json:"memory,omitempty"`
				} `tfsdk:"request" json:"request,omitempty"`
			} `tfsdk:"che_postgres_container_resources" json:"chePostgresContainerResources,omitempty"`
			ChePostgresDb       *string `tfsdk:"che_postgres_db" json:"chePostgresDb,omitempty"`
			ChePostgresHostName *string `tfsdk:"che_postgres_host_name" json:"chePostgresHostName,omitempty"`
			ChePostgresPassword *string `tfsdk:"che_postgres_password" json:"chePostgresPassword,omitempty"`
			ChePostgresPort     *string `tfsdk:"che_postgres_port" json:"chePostgresPort,omitempty"`
			ChePostgresSecret   *string `tfsdk:"che_postgres_secret" json:"chePostgresSecret,omitempty"`
			ChePostgresUser     *string `tfsdk:"che_postgres_user" json:"chePostgresUser,omitempty"`
			ExternalDb          *bool   `tfsdk:"external_db" json:"externalDb,omitempty"`
			PostgresEnv         *[]struct {
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
			} `tfsdk:"postgres_env" json:"postgresEnv,omitempty"`
			PostgresImage           *string `tfsdk:"postgres_image" json:"postgresImage,omitempty"`
			PostgresImagePullPolicy *string `tfsdk:"postgres_image_pull_policy" json:"postgresImagePullPolicy,omitempty"`
			PostgresVersion         *string `tfsdk:"postgres_version" json:"postgresVersion,omitempty"`
			PvcClaimSize            *string `tfsdk:"pvc_claim_size" json:"pvcClaimSize,omitempty"`
		} `tfsdk:"database" json:"database,omitempty"`
		DevWorkspace *struct {
			ControllerImage *string `tfsdk:"controller_image" json:"controllerImage,omitempty"`
			Enable          *bool   `tfsdk:"enable" json:"enable,omitempty"`
			Env             *[]struct {
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
			} `tfsdk:"env" json:"env,omitempty"`
			RunningLimit                    *string `tfsdk:"running_limit" json:"runningLimit,omitempty"`
			SecondsOfInactivityBeforeIdling *int64  `tfsdk:"seconds_of_inactivity_before_idling" json:"secondsOfInactivityBeforeIdling,omitempty"`
			SecondsOfRunBeforeIdling        *int64  `tfsdk:"seconds_of_run_before_idling" json:"secondsOfRunBeforeIdling,omitempty"`
		} `tfsdk:"dev_workspace" json:"devWorkspace,omitempty"`
		GitServices *struct {
			Bitbucket *[]struct {
				Endpoint   *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
				SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"bitbucket" json:"bitbucket,omitempty"`
			Github *[]struct {
				Endpoint   *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
				SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"github" json:"github,omitempty"`
			Gitlab *[]struct {
				Endpoint   *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
				SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"gitlab" json:"gitlab,omitempty"`
		} `tfsdk:"git_services" json:"gitServices,omitempty"`
		ImagePuller *struct {
			Enable *bool `tfsdk:"enable" json:"enable,omitempty"`
			Spec   *struct {
				Affinity             *string `tfsdk:"affinity" json:"affinity,omitempty"`
				CachingCPULimit      *string `tfsdk:"caching_cpu_limit" json:"cachingCPULimit,omitempty"`
				CachingCPURequest    *string `tfsdk:"caching_cpu_request" json:"cachingCPURequest,omitempty"`
				CachingIntervalHours *string `tfsdk:"caching_interval_hours" json:"cachingIntervalHours,omitempty"`
				CachingMemoryLimit   *string `tfsdk:"caching_memory_limit" json:"cachingMemoryLimit,omitempty"`
				CachingMemoryRequest *string `tfsdk:"caching_memory_request" json:"cachingMemoryRequest,omitempty"`
				ConfigMapName        *string `tfsdk:"config_map_name" json:"configMapName,omitempty"`
				DaemonsetName        *string `tfsdk:"daemonset_name" json:"daemonsetName,omitempty"`
				DeploymentName       *string `tfsdk:"deployment_name" json:"deploymentName,omitempty"`
				ImagePullSecrets     *string `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
				ImagePullerImage     *string `tfsdk:"image_puller_image" json:"imagePullerImage,omitempty"`
				Images               *string `tfsdk:"images" json:"images,omitempty"`
				NodeSelector         *string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"image_puller" json:"imagePuller,omitempty"`
		K8s *struct {
			IngressClass             *string `tfsdk:"ingress_class" json:"ingressClass,omitempty"`
			IngressDomain            *string `tfsdk:"ingress_domain" json:"ingressDomain,omitempty"`
			IngressStrategy          *string `tfsdk:"ingress_strategy" json:"ingressStrategy,omitempty"`
			SecurityContextFsGroup   *string `tfsdk:"security_context_fs_group" json:"securityContextFsGroup,omitempty"`
			SecurityContextRunAsUser *string `tfsdk:"security_context_run_as_user" json:"securityContextRunAsUser,omitempty"`
			SingleHostExposureType   *string `tfsdk:"single_host_exposure_type" json:"singleHostExposureType,omitempty"`
			TlsSecretName            *string `tfsdk:"tls_secret_name" json:"tlsSecretName,omitempty"`
		} `tfsdk:"k8s" json:"k8s,omitempty"`
		Metrics *struct {
			Enable *bool `tfsdk:"enable" json:"enable,omitempty"`
		} `tfsdk:"metrics" json:"metrics,omitempty"`
		Server *struct {
			AirGapContainerRegistryHostname     *string `tfsdk:"air_gap_container_registry_hostname" json:"airGapContainerRegistryHostname,omitempty"`
			AirGapContainerRegistryOrganization *string `tfsdk:"air_gap_container_registry_organization" json:"airGapContainerRegistryOrganization,omitempty"`
			AllowAutoProvisionUserNamespace     *bool   `tfsdk:"allow_auto_provision_user_namespace" json:"allowAutoProvisionUserNamespace,omitempty"`
			AllowUserDefinedWorkspaceNamespaces *bool   `tfsdk:"allow_user_defined_workspace_namespaces" json:"allowUserDefinedWorkspaceNamespaces,omitempty"`
			CheClusterRoles                     *string `tfsdk:"che_cluster_roles" json:"cheClusterRoles,omitempty"`
			CheDebug                            *string `tfsdk:"che_debug" json:"cheDebug,omitempty"`
			CheFlavor                           *string `tfsdk:"che_flavor" json:"cheFlavor,omitempty"`
			CheHost                             *string `tfsdk:"che_host" json:"cheHost,omitempty"`
			CheHostTLSSecret                    *string `tfsdk:"che_host_tls_secret" json:"cheHostTLSSecret,omitempty"`
			CheImage                            *string `tfsdk:"che_image" json:"cheImage,omitempty"`
			CheImagePullPolicy                  *string `tfsdk:"che_image_pull_policy" json:"cheImagePullPolicy,omitempty"`
			CheImageTag                         *string `tfsdk:"che_image_tag" json:"cheImageTag,omitempty"`
			CheLogLevel                         *string `tfsdk:"che_log_level" json:"cheLogLevel,omitempty"`
			CheServerEnv                        *[]struct {
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
			} `tfsdk:"che_server_env" json:"cheServerEnv,omitempty"`
			CheServerIngress *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *string            `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"che_server_ingress" json:"cheServerIngress,omitempty"`
			CheServerRoute *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Domain      *string            `tfsdk:"domain" json:"domain,omitempty"`
				Labels      *string            `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"che_server_route" json:"cheServerRoute,omitempty"`
			CheWorkspaceClusterRole *string            `tfsdk:"che_workspace_cluster_role" json:"cheWorkspaceClusterRole,omitempty"`
			CustomCheProperties     *map[string]string `tfsdk:"custom_che_properties" json:"customCheProperties,omitempty"`
			DashboardCpuLimit       *string            `tfsdk:"dashboard_cpu_limit" json:"dashboardCpuLimit,omitempty"`
			DashboardCpuRequest     *string            `tfsdk:"dashboard_cpu_request" json:"dashboardCpuRequest,omitempty"`
			DashboardEnv            *[]struct {
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
			} `tfsdk:"dashboard_env" json:"dashboardEnv,omitempty"`
			DashboardImage           *string `tfsdk:"dashboard_image" json:"dashboardImage,omitempty"`
			DashboardImagePullPolicy *string `tfsdk:"dashboard_image_pull_policy" json:"dashboardImagePullPolicy,omitempty"`
			DashboardIngress         *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *string            `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"dashboard_ingress" json:"dashboardIngress,omitempty"`
			DashboardMemoryLimit   *string `tfsdk:"dashboard_memory_limit" json:"dashboardMemoryLimit,omitempty"`
			DashboardMemoryRequest *string `tfsdk:"dashboard_memory_request" json:"dashboardMemoryRequest,omitempty"`
			DashboardRoute         *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Domain      *string            `tfsdk:"domain" json:"domain,omitempty"`
				Labels      *string            `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"dashboard_route" json:"dashboardRoute,omitempty"`
			DevfileRegistryCpuLimit   *string `tfsdk:"devfile_registry_cpu_limit" json:"devfileRegistryCpuLimit,omitempty"`
			DevfileRegistryCpuRequest *string `tfsdk:"devfile_registry_cpu_request" json:"devfileRegistryCpuRequest,omitempty"`
			DevfileRegistryEnv        *[]struct {
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
			} `tfsdk:"devfile_registry_env" json:"devfileRegistryEnv,omitempty"`
			DevfileRegistryImage   *string `tfsdk:"devfile_registry_image" json:"devfileRegistryImage,omitempty"`
			DevfileRegistryIngress *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *string            `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"devfile_registry_ingress" json:"devfileRegistryIngress,omitempty"`
			DevfileRegistryMemoryLimit   *string `tfsdk:"devfile_registry_memory_limit" json:"devfileRegistryMemoryLimit,omitempty"`
			DevfileRegistryMemoryRequest *string `tfsdk:"devfile_registry_memory_request" json:"devfileRegistryMemoryRequest,omitempty"`
			DevfileRegistryPullPolicy    *string `tfsdk:"devfile_registry_pull_policy" json:"devfileRegistryPullPolicy,omitempty"`
			DevfileRegistryRoute         *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Domain      *string            `tfsdk:"domain" json:"domain,omitempty"`
				Labels      *string            `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"devfile_registry_route" json:"devfileRegistryRoute,omitempty"`
			DevfileRegistryUrl             *string `tfsdk:"devfile_registry_url" json:"devfileRegistryUrl,omitempty"`
			DisableInternalClusterSVCNames *bool   `tfsdk:"disable_internal_cluster_svc_names" json:"disableInternalClusterSVCNames,omitempty"`
			ExternalDevfileRegistries      *[]struct {
				Url *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"external_devfile_registries" json:"externalDevfileRegistries,omitempty"`
			ExternalDevfileRegistry  *bool   `tfsdk:"external_devfile_registry" json:"externalDevfileRegistry,omitempty"`
			ExternalPluginRegistry   *bool   `tfsdk:"external_plugin_registry" json:"externalPluginRegistry,omitempty"`
			GitSelfSignedCert        *bool   `tfsdk:"git_self_signed_cert" json:"gitSelfSignedCert,omitempty"`
			NonProxyHosts            *string `tfsdk:"non_proxy_hosts" json:"nonProxyHosts,omitempty"`
			OpenVSXRegistryURL       *string `tfsdk:"open_vsx_registry_url" json:"openVSXRegistryURL,omitempty"`
			PluginRegistryCpuLimit   *string `tfsdk:"plugin_registry_cpu_limit" json:"pluginRegistryCpuLimit,omitempty"`
			PluginRegistryCpuRequest *string `tfsdk:"plugin_registry_cpu_request" json:"pluginRegistryCpuRequest,omitempty"`
			PluginRegistryEnv        *[]struct {
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
			} `tfsdk:"plugin_registry_env" json:"pluginRegistryEnv,omitempty"`
			PluginRegistryImage   *string `tfsdk:"plugin_registry_image" json:"pluginRegistryImage,omitempty"`
			PluginRegistryIngress *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *string            `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"plugin_registry_ingress" json:"pluginRegistryIngress,omitempty"`
			PluginRegistryMemoryLimit   *string `tfsdk:"plugin_registry_memory_limit" json:"pluginRegistryMemoryLimit,omitempty"`
			PluginRegistryMemoryRequest *string `tfsdk:"plugin_registry_memory_request" json:"pluginRegistryMemoryRequest,omitempty"`
			PluginRegistryPullPolicy    *string `tfsdk:"plugin_registry_pull_policy" json:"pluginRegistryPullPolicy,omitempty"`
			PluginRegistryRoute         *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Domain      *string            `tfsdk:"domain" json:"domain,omitempty"`
				Labels      *string            `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"plugin_registry_route" json:"pluginRegistryRoute,omitempty"`
			PluginRegistryUrl                   *string            `tfsdk:"plugin_registry_url" json:"pluginRegistryUrl,omitempty"`
			ProxyPassword                       *string            `tfsdk:"proxy_password" json:"proxyPassword,omitempty"`
			ProxyPort                           *string            `tfsdk:"proxy_port" json:"proxyPort,omitempty"`
			ProxySecret                         *string            `tfsdk:"proxy_secret" json:"proxySecret,omitempty"`
			ProxyURL                            *string            `tfsdk:"proxy_url" json:"proxyURL,omitempty"`
			ProxyUser                           *string            `tfsdk:"proxy_user" json:"proxyUser,omitempty"`
			SelfSignedCert                      *bool              `tfsdk:"self_signed_cert" json:"selfSignedCert,omitempty"`
			ServerCpuLimit                      *string            `tfsdk:"server_cpu_limit" json:"serverCpuLimit,omitempty"`
			ServerCpuRequest                    *string            `tfsdk:"server_cpu_request" json:"serverCpuRequest,omitempty"`
			ServerExposureStrategy              *string            `tfsdk:"server_exposure_strategy" json:"serverExposureStrategy,omitempty"`
			ServerMemoryLimit                   *string            `tfsdk:"server_memory_limit" json:"serverMemoryLimit,omitempty"`
			ServerMemoryRequest                 *string            `tfsdk:"server_memory_request" json:"serverMemoryRequest,omitempty"`
			ServerTrustStoreConfigMapName       *string            `tfsdk:"server_trust_store_config_map_name" json:"serverTrustStoreConfigMapName,omitempty"`
			SingleHostGatewayConfigMapLabels    *map[string]string `tfsdk:"single_host_gateway_config_map_labels" json:"singleHostGatewayConfigMapLabels,omitempty"`
			SingleHostGatewayConfigSidecarImage *string            `tfsdk:"single_host_gateway_config_sidecar_image" json:"singleHostGatewayConfigSidecarImage,omitempty"`
			SingleHostGatewayImage              *string            `tfsdk:"single_host_gateway_image" json:"singleHostGatewayImage,omitempty"`
			TlsSupport                          *bool              `tfsdk:"tls_support" json:"tlsSupport,omitempty"`
			UseInternalClusterSVCNames          *bool              `tfsdk:"use_internal_cluster_svc_names" json:"useInternalClusterSVCNames,omitempty"`
			WorkspaceDefaultComponents          *[]struct {
				Attributes    *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
				ComponentType *string            `tfsdk:"component_type" json:"componentType,omitempty"`
				Container     *struct {
					Annotation *struct {
						Deployment *map[string]string `tfsdk:"deployment" json:"deployment,omitempty"`
						Service    *map[string]string `tfsdk:"service" json:"service,omitempty"`
					} `tfsdk:"annotation" json:"annotation,omitempty"`
					Args         *[]string `tfsdk:"args" json:"args,omitempty"`
					Command      *[]string `tfsdk:"command" json:"command,omitempty"`
					CpuLimit     *string   `tfsdk:"cpu_limit" json:"cpuLimit,omitempty"`
					CpuRequest   *string   `tfsdk:"cpu_request" json:"cpuRequest,omitempty"`
					DedicatedPod *bool     `tfsdk:"dedicated_pod" json:"dedicatedPod,omitempty"`
					Endpoints    *[]struct {
						Annotation *map[string]string `tfsdk:"annotation" json:"annotation,omitempty"`
						Attributes *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
						Exposure   *string            `tfsdk:"exposure" json:"exposure,omitempty"`
						Name       *string            `tfsdk:"name" json:"name,omitempty"`
						Path       *string            `tfsdk:"path" json:"path,omitempty"`
						Protocol   *string            `tfsdk:"protocol" json:"protocol,omitempty"`
						Secure     *bool              `tfsdk:"secure" json:"secure,omitempty"`
						TargetPort *int64             `tfsdk:"target_port" json:"targetPort,omitempty"`
					} `tfsdk:"endpoints" json:"endpoints,omitempty"`
					Env *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
					Image         *string `tfsdk:"image" json:"image,omitempty"`
					MemoryLimit   *string `tfsdk:"memory_limit" json:"memoryLimit,omitempty"`
					MemoryRequest *string `tfsdk:"memory_request" json:"memoryRequest,omitempty"`
					MountSources  *bool   `tfsdk:"mount_sources" json:"mountSources,omitempty"`
					SourceMapping *string `tfsdk:"source_mapping" json:"sourceMapping,omitempty"`
					VolumeMounts  *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
						Path *string `tfsdk:"path" json:"path,omitempty"`
					} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
				} `tfsdk:"container" json:"container,omitempty"`
				Custom *struct {
					ComponentClass   *string            `tfsdk:"component_class" json:"componentClass,omitempty"`
					EmbeddedResource *map[string]string `tfsdk:"embedded_resource" json:"embeddedResource,omitempty"`
				} `tfsdk:"custom" json:"custom,omitempty"`
				Image *struct {
					AutoBuild  *bool `tfsdk:"auto_build" json:"autoBuild,omitempty"`
					Dockerfile *struct {
						Args            *[]string `tfsdk:"args" json:"args,omitempty"`
						BuildContext    *string   `tfsdk:"build_context" json:"buildContext,omitempty"`
						DevfileRegistry *struct {
							Id          *string `tfsdk:"id" json:"id,omitempty"`
							RegistryUrl *string `tfsdk:"registry_url" json:"registryUrl,omitempty"`
						} `tfsdk:"devfile_registry" json:"devfileRegistry,omitempty"`
						Git *struct {
							CheckoutFrom *struct {
								Remote   *string `tfsdk:"remote" json:"remote,omitempty"`
								Revision *string `tfsdk:"revision" json:"revision,omitempty"`
							} `tfsdk:"checkout_from" json:"checkoutFrom,omitempty"`
							FileLocation *string            `tfsdk:"file_location" json:"fileLocation,omitempty"`
							Remotes      *map[string]string `tfsdk:"remotes" json:"remotes,omitempty"`
						} `tfsdk:"git" json:"git,omitempty"`
						RootRequired *bool   `tfsdk:"root_required" json:"rootRequired,omitempty"`
						SrcType      *string `tfsdk:"src_type" json:"srcType,omitempty"`
						Uri          *string `tfsdk:"uri" json:"uri,omitempty"`
					} `tfsdk:"dockerfile" json:"dockerfile,omitempty"`
					ImageName *string `tfsdk:"image_name" json:"imageName,omitempty"`
					ImageType *string `tfsdk:"image_type" json:"imageType,omitempty"`
				} `tfsdk:"image" json:"image,omitempty"`
				Kubernetes *struct {
					DeployByDefault *bool `tfsdk:"deploy_by_default" json:"deployByDefault,omitempty"`
					Endpoints       *[]struct {
						Annotation *map[string]string `tfsdk:"annotation" json:"annotation,omitempty"`
						Attributes *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
						Exposure   *string            `tfsdk:"exposure" json:"exposure,omitempty"`
						Name       *string            `tfsdk:"name" json:"name,omitempty"`
						Path       *string            `tfsdk:"path" json:"path,omitempty"`
						Protocol   *string            `tfsdk:"protocol" json:"protocol,omitempty"`
						Secure     *bool              `tfsdk:"secure" json:"secure,omitempty"`
						TargetPort *int64             `tfsdk:"target_port" json:"targetPort,omitempty"`
					} `tfsdk:"endpoints" json:"endpoints,omitempty"`
					Inlined      *string `tfsdk:"inlined" json:"inlined,omitempty"`
					LocationType *string `tfsdk:"location_type" json:"locationType,omitempty"`
					Uri          *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Openshift *struct {
					DeployByDefault *bool `tfsdk:"deploy_by_default" json:"deployByDefault,omitempty"`
					Endpoints       *[]struct {
						Annotation *map[string]string `tfsdk:"annotation" json:"annotation,omitempty"`
						Attributes *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
						Exposure   *string            `tfsdk:"exposure" json:"exposure,omitempty"`
						Name       *string            `tfsdk:"name" json:"name,omitempty"`
						Path       *string            `tfsdk:"path" json:"path,omitempty"`
						Protocol   *string            `tfsdk:"protocol" json:"protocol,omitempty"`
						Secure     *bool              `tfsdk:"secure" json:"secure,omitempty"`
						TargetPort *int64             `tfsdk:"target_port" json:"targetPort,omitempty"`
					} `tfsdk:"endpoints" json:"endpoints,omitempty"`
					Inlined      *string `tfsdk:"inlined" json:"inlined,omitempty"`
					LocationType *string `tfsdk:"location_type" json:"locationType,omitempty"`
					Uri          *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"openshift" json:"openshift,omitempty"`
				Plugin *struct {
					Commands *[]struct {
						Apply *struct {
							Component *string `tfsdk:"component" json:"component,omitempty"`
							Group     *struct {
								IsDefault *bool   `tfsdk:"is_default" json:"isDefault,omitempty"`
								Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
							} `tfsdk:"group" json:"group,omitempty"`
							Label *string `tfsdk:"label" json:"label,omitempty"`
						} `tfsdk:"apply" json:"apply,omitempty"`
						Attributes  *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
						CommandType *string            `tfsdk:"command_type" json:"commandType,omitempty"`
						Composite   *struct {
							Commands *[]string `tfsdk:"commands" json:"commands,omitempty"`
							Group    *struct {
								IsDefault *bool   `tfsdk:"is_default" json:"isDefault,omitempty"`
								Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
							} `tfsdk:"group" json:"group,omitempty"`
							Label    *string `tfsdk:"label" json:"label,omitempty"`
							Parallel *bool   `tfsdk:"parallel" json:"parallel,omitempty"`
						} `tfsdk:"composite" json:"composite,omitempty"`
						Exec *struct {
							CommandLine *string `tfsdk:"command_line" json:"commandLine,omitempty"`
							Component   *string `tfsdk:"component" json:"component,omitempty"`
							Env         *[]struct {
								Name  *string `tfsdk:"name" json:"name,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"env" json:"env,omitempty"`
							Group *struct {
								IsDefault *bool   `tfsdk:"is_default" json:"isDefault,omitempty"`
								Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
							} `tfsdk:"group" json:"group,omitempty"`
							HotReloadCapable *bool   `tfsdk:"hot_reload_capable" json:"hotReloadCapable,omitempty"`
							Label            *string `tfsdk:"label" json:"label,omitempty"`
							WorkingDir       *string `tfsdk:"working_dir" json:"workingDir,omitempty"`
						} `tfsdk:"exec" json:"exec,omitempty"`
						Id *string `tfsdk:"id" json:"id,omitempty"`
					} `tfsdk:"commands" json:"commands,omitempty"`
					Components *[]struct {
						Attributes    *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
						ComponentType *string            `tfsdk:"component_type" json:"componentType,omitempty"`
						Container     *struct {
							Annotation *struct {
								Deployment *map[string]string `tfsdk:"deployment" json:"deployment,omitempty"`
								Service    *map[string]string `tfsdk:"service" json:"service,omitempty"`
							} `tfsdk:"annotation" json:"annotation,omitempty"`
							Args         *[]string `tfsdk:"args" json:"args,omitempty"`
							Command      *[]string `tfsdk:"command" json:"command,omitempty"`
							CpuLimit     *string   `tfsdk:"cpu_limit" json:"cpuLimit,omitempty"`
							CpuRequest   *string   `tfsdk:"cpu_request" json:"cpuRequest,omitempty"`
							DedicatedPod *bool     `tfsdk:"dedicated_pod" json:"dedicatedPod,omitempty"`
							Endpoints    *[]struct {
								Annotation *map[string]string `tfsdk:"annotation" json:"annotation,omitempty"`
								Attributes *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
								Exposure   *string            `tfsdk:"exposure" json:"exposure,omitempty"`
								Name       *string            `tfsdk:"name" json:"name,omitempty"`
								Path       *string            `tfsdk:"path" json:"path,omitempty"`
								Protocol   *string            `tfsdk:"protocol" json:"protocol,omitempty"`
								Secure     *bool              `tfsdk:"secure" json:"secure,omitempty"`
								TargetPort *int64             `tfsdk:"target_port" json:"targetPort,omitempty"`
							} `tfsdk:"endpoints" json:"endpoints,omitempty"`
							Env *[]struct {
								Name  *string `tfsdk:"name" json:"name,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"env" json:"env,omitempty"`
							Image         *string `tfsdk:"image" json:"image,omitempty"`
							MemoryLimit   *string `tfsdk:"memory_limit" json:"memoryLimit,omitempty"`
							MemoryRequest *string `tfsdk:"memory_request" json:"memoryRequest,omitempty"`
							MountSources  *bool   `tfsdk:"mount_sources" json:"mountSources,omitempty"`
							SourceMapping *string `tfsdk:"source_mapping" json:"sourceMapping,omitempty"`
							VolumeMounts  *[]struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
								Path *string `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
						} `tfsdk:"container" json:"container,omitempty"`
						Image *struct {
							AutoBuild  *bool `tfsdk:"auto_build" json:"autoBuild,omitempty"`
							Dockerfile *struct {
								Args            *[]string `tfsdk:"args" json:"args,omitempty"`
								BuildContext    *string   `tfsdk:"build_context" json:"buildContext,omitempty"`
								DevfileRegistry *struct {
									Id          *string `tfsdk:"id" json:"id,omitempty"`
									RegistryUrl *string `tfsdk:"registry_url" json:"registryUrl,omitempty"`
								} `tfsdk:"devfile_registry" json:"devfileRegistry,omitempty"`
								Git *struct {
									CheckoutFrom *struct {
										Remote   *string `tfsdk:"remote" json:"remote,omitempty"`
										Revision *string `tfsdk:"revision" json:"revision,omitempty"`
									} `tfsdk:"checkout_from" json:"checkoutFrom,omitempty"`
									FileLocation *string            `tfsdk:"file_location" json:"fileLocation,omitempty"`
									Remotes      *map[string]string `tfsdk:"remotes" json:"remotes,omitempty"`
								} `tfsdk:"git" json:"git,omitempty"`
								RootRequired *bool   `tfsdk:"root_required" json:"rootRequired,omitempty"`
								SrcType      *string `tfsdk:"src_type" json:"srcType,omitempty"`
								Uri          *string `tfsdk:"uri" json:"uri,omitempty"`
							} `tfsdk:"dockerfile" json:"dockerfile,omitempty"`
							ImageName *string `tfsdk:"image_name" json:"imageName,omitempty"`
							ImageType *string `tfsdk:"image_type" json:"imageType,omitempty"`
						} `tfsdk:"image" json:"image,omitempty"`
						Kubernetes *struct {
							DeployByDefault *bool `tfsdk:"deploy_by_default" json:"deployByDefault,omitempty"`
							Endpoints       *[]struct {
								Annotation *map[string]string `tfsdk:"annotation" json:"annotation,omitempty"`
								Attributes *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
								Exposure   *string            `tfsdk:"exposure" json:"exposure,omitempty"`
								Name       *string            `tfsdk:"name" json:"name,omitempty"`
								Path       *string            `tfsdk:"path" json:"path,omitempty"`
								Protocol   *string            `tfsdk:"protocol" json:"protocol,omitempty"`
								Secure     *bool              `tfsdk:"secure" json:"secure,omitempty"`
								TargetPort *int64             `tfsdk:"target_port" json:"targetPort,omitempty"`
							} `tfsdk:"endpoints" json:"endpoints,omitempty"`
							Inlined      *string `tfsdk:"inlined" json:"inlined,omitempty"`
							LocationType *string `tfsdk:"location_type" json:"locationType,omitempty"`
							Uri          *string `tfsdk:"uri" json:"uri,omitempty"`
						} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Openshift *struct {
							DeployByDefault *bool `tfsdk:"deploy_by_default" json:"deployByDefault,omitempty"`
							Endpoints       *[]struct {
								Annotation *map[string]string `tfsdk:"annotation" json:"annotation,omitempty"`
								Attributes *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
								Exposure   *string            `tfsdk:"exposure" json:"exposure,omitempty"`
								Name       *string            `tfsdk:"name" json:"name,omitempty"`
								Path       *string            `tfsdk:"path" json:"path,omitempty"`
								Protocol   *string            `tfsdk:"protocol" json:"protocol,omitempty"`
								Secure     *bool              `tfsdk:"secure" json:"secure,omitempty"`
								TargetPort *int64             `tfsdk:"target_port" json:"targetPort,omitempty"`
							} `tfsdk:"endpoints" json:"endpoints,omitempty"`
							Inlined      *string `tfsdk:"inlined" json:"inlined,omitempty"`
							LocationType *string `tfsdk:"location_type" json:"locationType,omitempty"`
							Uri          *string `tfsdk:"uri" json:"uri,omitempty"`
						} `tfsdk:"openshift" json:"openshift,omitempty"`
						Volume *struct {
							Ephemeral *bool   `tfsdk:"ephemeral" json:"ephemeral,omitempty"`
							Size      *string `tfsdk:"size" json:"size,omitempty"`
						} `tfsdk:"volume" json:"volume,omitempty"`
					} `tfsdk:"components" json:"components,omitempty"`
					Id                  *string `tfsdk:"id" json:"id,omitempty"`
					ImportReferenceType *string `tfsdk:"import_reference_type" json:"importReferenceType,omitempty"`
					Kubernetes          *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
					RegistryUrl *string `tfsdk:"registry_url" json:"registryUrl,omitempty"`
					Uri         *string `tfsdk:"uri" json:"uri,omitempty"`
					Version     *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"plugin" json:"plugin,omitempty"`
				Volume *struct {
					Ephemeral *bool   `tfsdk:"ephemeral" json:"ephemeral,omitempty"`
					Size      *string `tfsdk:"size" json:"size,omitempty"`
				} `tfsdk:"volume" json:"volume,omitempty"`
			} `tfsdk:"workspace_default_components" json:"workspaceDefaultComponents,omitempty"`
			WorkspaceDefaultEditor    *string            `tfsdk:"workspace_default_editor" json:"workspaceDefaultEditor,omitempty"`
			WorkspaceNamespaceDefault *string            `tfsdk:"workspace_namespace_default" json:"workspaceNamespaceDefault,omitempty"`
			WorkspacePodNodeSelector  *map[string]string `tfsdk:"workspace_pod_node_selector" json:"workspacePodNodeSelector,omitempty"`
			WorkspacePodTolerations   *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"workspace_pod_tolerations" json:"workspacePodTolerations,omitempty"`
			WorkspacesDefaultPlugins *[]struct {
				Editor  *string   `tfsdk:"editor" json:"editor,omitempty"`
				Plugins *[]string `tfsdk:"plugins" json:"plugins,omitempty"`
			} `tfsdk:"workspaces_default_plugins" json:"workspacesDefaultPlugins,omitempty"`
		} `tfsdk:"server" json:"server,omitempty"`
		Storage *struct {
			PerWorkspaceStrategyPVCStorageClassName *string `tfsdk:"per_workspace_strategy_pvc_storage_class_name" json:"perWorkspaceStrategyPVCStorageClassName,omitempty"`
			PerWorkspaceStrategyPvcClaimSize        *string `tfsdk:"per_workspace_strategy_pvc_claim_size" json:"perWorkspaceStrategyPvcClaimSize,omitempty"`
			PostgresPVCStorageClassName             *string `tfsdk:"postgres_pvc_storage_class_name" json:"postgresPVCStorageClassName,omitempty"`
			PreCreateSubPaths                       *bool   `tfsdk:"pre_create_sub_paths" json:"preCreateSubPaths,omitempty"`
			PvcClaimSize                            *string `tfsdk:"pvc_claim_size" json:"pvcClaimSize,omitempty"`
			PvcJobsImage                            *string `tfsdk:"pvc_jobs_image" json:"pvcJobsImage,omitempty"`
			PvcStrategy                             *string `tfsdk:"pvc_strategy" json:"pvcStrategy,omitempty"`
			WorkspacePVCStorageClassName            *string `tfsdk:"workspace_pvc_storage_class_name" json:"workspacePVCStorageClassName,omitempty"`
		} `tfsdk:"storage" json:"storage,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OrgEclipseCheCheClusterV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_org_eclipse_che_che_cluster_v1_manifest"
}

func (r *OrgEclipseCheCheClusterV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "The 'CheCluster' custom resource allows defining and managing a Che server installation",
		MarkdownDescription: "The 'CheCluster' custom resource allows defining and managing a Che server installation",
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
				Description:         "Desired configuration of the Che installation. Based on these settings, the  Operator automatically creates and maintains several ConfigMaps that will contain the appropriate environment variables the various components of the Che installation. These generated ConfigMaps must NOT be updated manually.",
				MarkdownDescription: "Desired configuration of the Che installation. Based on these settings, the  Operator automatically creates and maintains several ConfigMaps that will contain the appropriate environment variables the various components of the Che installation. These generated ConfigMaps must NOT be updated manually.",
				Attributes: map[string]schema.Attribute{
					"auth": schema.SingleNestedAttribute{
						Description:         "Configuration settings related to the Authentication used by the Che installation.",
						MarkdownDescription: "Configuration settings related to the Authentication used by the Che installation.",
						Attributes: map[string]schema.Attribute{
							"debug": schema.BoolAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Debug internal identity provider.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Debug internal identity provider.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"external_identity_provider": schema.BoolAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Instructs the Operator on whether or not to deploy a dedicated Identity Provider (Keycloak or RH SSO instance). Instructs the Operator on whether to deploy a dedicated Identity Provider (Keycloak or RH-SSO instance). By default, a dedicated Identity Provider server is deployed as part of the Che installation. When 'externalIdentityProvider' is 'true', no dedicated identity provider will be deployed by the Operator and you will need to provide details about the external identity provider you are about to use. See also all the other fields starting with: 'identityProvider'.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Instructs the Operator on whether or not to deploy a dedicated Identity Provider (Keycloak or RH SSO instance). Instructs the Operator on whether to deploy a dedicated Identity Provider (Keycloak or RH-SSO instance). By default, a dedicated Identity Provider server is deployed as part of the Che installation. When 'externalIdentityProvider' is 'true', no dedicated identity provider will be deployed by the Operator and you will need to provide details about the external identity provider you are about to use. See also all the other fields starting with: 'identityProvider'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"gateway_authentication_sidecar_image": schema.StringAttribute{
								Description:         "Gateway sidecar responsible for authentication when NativeUserMode is enabled. See link:https://github.com/oauth2-proxy/oauth2-proxy[oauth2-proxy] or link:https://github.com/openshift/oauth-proxy[openshift/oauth-proxy].",
								MarkdownDescription: "Gateway sidecar responsible for authentication when NativeUserMode is enabled. See link:https://github.com/oauth2-proxy/oauth2-proxy[oauth2-proxy] or link:https://github.com/openshift/oauth-proxy[openshift/oauth-proxy].",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"gateway_authorization_sidecar_image": schema.StringAttribute{
								Description:         "Gateway sidecar responsible for authorization when NativeUserMode is enabled. See link:https://github.com/brancz/kube-rbac-proxy[kube-rbac-proxy] or link:https://github.com/openshift/kube-rbac-proxy[openshift/kube-rbac-proxy]",
								MarkdownDescription: "Gateway sidecar responsible for authorization when NativeUserMode is enabled. See link:https://github.com/brancz/kube-rbac-proxy[kube-rbac-proxy] or link:https://github.com/openshift/kube-rbac-proxy[openshift/kube-rbac-proxy]",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"gateway_config_bump_env": schema.ListNestedAttribute{
								Description:         "List of environment variables to set in the Configbump container.",
								MarkdownDescription: "List of environment variables to set in the Configbump container.",
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

							"gateway_env": schema.ListNestedAttribute{
								Description:         "List of environment variables to set in the Gateway container.",
								MarkdownDescription: "List of environment variables to set in the Gateway container.",
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

							"gateway_header_rewrite_sidecar_image": schema.StringAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Sidecar functionality is now implemented in Traefik plugin.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Sidecar functionality is now implemented in Traefik plugin.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"gateway_kube_rbac_proxy_env": schema.ListNestedAttribute{
								Description:         "List of environment variables to set in the Kube rbac proxy container.",
								MarkdownDescription: "List of environment variables to set in the Kube rbac proxy container.",
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

							"gateway_o_auth_proxy_env": schema.ListNestedAttribute{
								Description:         "List of environment variables to set in the OAuth proxy container.",
								MarkdownDescription: "List of environment variables to set in the OAuth proxy container.",
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

							"identity_provider_admin_user_name": schema.StringAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Overrides the name of the Identity Provider administrator user. Defaults to 'admin'.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Overrides the name of the Identity Provider administrator user. Defaults to 'admin'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"identity_provider_client_id": schema.StringAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Name of a Identity provider, Keycloak or RH-SSO, 'client-id' that is used for Che. Override this when an external Identity Provider is in use. See the 'externalIdentityProvider' field. When omitted or left blank, it is set to the value of the 'flavour' field suffixed with '-public'.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Name of a Identity provider, Keycloak or RH-SSO, 'client-id' that is used for Che. Override this when an external Identity Provider is in use. See the 'externalIdentityProvider' field. When omitted or left blank, it is set to the value of the 'flavour' field suffixed with '-public'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"identity_provider_container_resources": schema.SingleNestedAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Identity provider container custom settings.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Identity provider container custom settings.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.SingleNestedAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed.",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.",
										Attributes: map[string]schema.Attribute{
											"cpu": schema.StringAttribute{
												Description:         "CPU, in cores. (500m = .5 cores)",
												MarkdownDescription: "CPU, in cores. (500m = .5 cores)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"memory": schema.StringAttribute{
												Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
												MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"request": schema.SingleNestedAttribute{
										Description:         "Requests describes the minimum amount of compute resources required.",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required.",
										Attributes: map[string]schema.Attribute{
											"cpu": schema.StringAttribute{
												Description:         "CPU, in cores. (500m = .5 cores)",
												MarkdownDescription: "CPU, in cores. (500m = .5 cores)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"memory": schema.StringAttribute{
												Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
												MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
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

							"identity_provider_image": schema.StringAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Overrides the container image used in the Identity Provider, Keycloak or RH-SSO, deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Overrides the container image used in the Identity Provider, Keycloak or RH-SSO, deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"identity_provider_image_pull_policy": schema.StringAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Overrides the image pull policy used in the Identity Provider, Keycloak or RH-SSO, deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Overrides the image pull policy used in the Identity Provider, Keycloak or RH-SSO, deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"identity_provider_ingress": schema.SingleNestedAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Ingress custom settings.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Ingress custom settings.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										MarkdownDescription: "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.StringAttribute{
										Description:         "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										MarkdownDescription: "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"identity_provider_password": schema.StringAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Overrides the password of Keycloak administrator user. Override this when an external Identity Provider is in use. See the 'externalIdentityProvider' field. When omitted or left blank, it is set to an auto-generated password.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Overrides the password of Keycloak administrator user. Override this when an external Identity Provider is in use. See the 'externalIdentityProvider' field. When omitted or left blank, it is set to an auto-generated password.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"identity_provider_postgres_password": schema.StringAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Password for a Identity Provider, Keycloak or RH-SSO, to connect to the database. Override this when an external Identity Provider is in use. See the 'externalIdentityProvider' field. When omitted or left blank, it is set to an auto-generated password.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Password for a Identity Provider, Keycloak or RH-SSO, to connect to the database. Override this when an external Identity Provider is in use. See the 'externalIdentityProvider' field. When omitted or left blank, it is set to an auto-generated password.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"identity_provider_postgres_secret": schema.StringAttribute{
								Description:         "Deprecated. The value of this flag is ignored. The secret that contains 'password' for the Identity Provider, Keycloak or RH-SSO, to connect to the database. When the secret is defined, the 'identityProviderPostgresPassword' is ignored. When the value is omitted or left blank, the one of following scenarios applies: 1. 'identityProviderPostgresPassword' is defined, then it will be used to connect to the database. 2. 'identityProviderPostgresPassword' is not defined, then a new secret with the name 'che-identity-postgres-secret' will be created with an auto-generated value for 'password'. The secret must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. The secret that contains 'password' for the Identity Provider, Keycloak or RH-SSO, to connect to the database. When the secret is defined, the 'identityProviderPostgresPassword' is ignored. When the value is omitted or left blank, the one of following scenarios applies: 1. 'identityProviderPostgresPassword' is defined, then it will be used to connect to the database. 2. 'identityProviderPostgresPassword' is not defined, then a new secret with the name 'che-identity-postgres-secret' will be created with an auto-generated value for 'password'. The secret must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"identity_provider_realm": schema.StringAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Name of a Identity provider, Keycloak or RH-SSO, realm that is used for Che. Override this when an external Identity Provider is in use. See the 'externalIdentityProvider' field. When omitted or left blank, it is set to the value of the 'flavour' field.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Name of a Identity provider, Keycloak or RH-SSO, realm that is used for Che. Override this when an external Identity Provider is in use. See the 'externalIdentityProvider' field. When omitted or left blank, it is set to the value of the 'flavour' field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"identity_provider_route": schema.SingleNestedAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Route custom settings.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Route custom settings.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										MarkdownDescription: "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"domain": schema.StringAttribute{
										Description:         "Operator uses the domain to generate a hostname for a route. In a conjunction with labels it creates a route, which is served by a non-default Ingress controller. The generated host name will follow this pattern: '<route-name>-<route-namespace>.<domain>'.",
										MarkdownDescription: "Operator uses the domain to generate a hostname for a route. In a conjunction with labels it creates a route, which is served by a non-default Ingress controller. The generated host name will follow this pattern: '<route-name>-<route-namespace>.<domain>'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.StringAttribute{
										Description:         "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										MarkdownDescription: "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"identity_provider_secret": schema.StringAttribute{
								Description:         "Deprecated. The value of this flag is ignored. The secret that contains 'user' and 'password' for Identity Provider. When the secret is defined, the 'identityProviderAdminUserName' and 'identityProviderPassword' are ignored. When the value is omitted or left blank, the one of following scenarios applies: 1. 'identityProviderAdminUserName' and 'identityProviderPassword' are defined, then they will be used. 2. 'identityProviderAdminUserName' or 'identityProviderPassword' are not defined, then a new secret with the name 'che-identity-secret' will be created with default value 'admin' for 'user' and with an auto-generated value for 'password'. The secret must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. The secret that contains 'user' and 'password' for Identity Provider. When the secret is defined, the 'identityProviderAdminUserName' and 'identityProviderPassword' are ignored. When the value is omitted or left blank, the one of following scenarios applies: 1. 'identityProviderAdminUserName' and 'identityProviderPassword' are defined, then they will be used. 2. 'identityProviderAdminUserName' or 'identityProviderPassword' are not defined, then a new secret with the name 'che-identity-secret' will be created with default value 'admin' for 'user' and with an auto-generated value for 'password'. The secret must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"identity_provider_url": schema.StringAttribute{
								Description:         "Public URL of the Identity Provider server (Keycloak / RH-SSO server). Set this ONLY when a use of an external Identity Provider is needed. See the 'externalIdentityProvider' field. By default, this will be automatically calculated and set by the Operator.",
								MarkdownDescription: "Public URL of the Identity Provider server (Keycloak / RH-SSO server). Set this ONLY when a use of an external Identity Provider is needed. See the 'externalIdentityProvider' field. By default, this will be automatically calculated and set by the Operator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"identity_token": schema.StringAttribute{
								Description:         "Identity token to be passed to upstream. There are two types of tokens supported: 'id_token' and 'access_token'. Default value is 'id_token'. This field is specific to Che installations made for Kubernetes only and ignored for OpenShift.",
								MarkdownDescription: "Identity token to be passed to upstream. There are two types of tokens supported: 'id_token' and 'access_token'. Default value is 'id_token'. This field is specific to Che installations made for Kubernetes only and ignored for OpenShift.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"initial_open_shift_o_auth_user": schema.BoolAttribute{
								Description:         "Deprecated. The value of this flag is ignored. For operating with the OpenShift OAuth authentication, create a new user account since the kubeadmin can not be used. If the value is true, then a new OpenShift OAuth user will be created for the HTPasswd identity provider. If the value is false and the user has already been created, then it will be removed. If value is an empty, then do nothing. The user's credentials are stored in the 'openshift-oauth-user-credentials' secret in 'openshift-config' namespace by Operator. Note that this solution is Openshift 4 platform-specific.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. For operating with the OpenShift OAuth authentication, create a new user account since the kubeadmin can not be used. If the value is true, then a new OpenShift OAuth user will be created for the HTPasswd identity provider. If the value is false and the user has already been created, then it will be removed. If value is an empty, then do nothing. The user's credentials are stored in the 'openshift-oauth-user-credentials' secret in 'openshift-config' namespace by Operator. Note that this solution is Openshift 4 platform-specific.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"native_user_mode": schema.BoolAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Enables native user mode. Currently works only on OpenShift and DevWorkspace engine. Native User mode uses OpenShift OAuth directly as identity provider, without Keycloak.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Enables native user mode. Currently works only on OpenShift and DevWorkspace engine. Native User mode uses OpenShift OAuth directly as identity provider, without Keycloak.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"o_auth_client_name": schema.StringAttribute{
								Description:         "Name of the OpenShift 'OAuthClient' resource used to setup identity federation on the OpenShift side. Auto-generated when left blank. See also the 'OpenShiftoAuth' field.",
								MarkdownDescription: "Name of the OpenShift 'OAuthClient' resource used to setup identity federation on the OpenShift side. Auto-generated when left blank. See also the 'OpenShiftoAuth' field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"o_auth_scope": schema.StringAttribute{
								Description:         "Access Token Scope. This field is specific to Che installations made for Kubernetes only and ignored for OpenShift.",
								MarkdownDescription: "Access Token Scope. This field is specific to Che installations made for Kubernetes only and ignored for OpenShift.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"o_auth_secret": schema.StringAttribute{
								Description:         "Name of the secret set in the OpenShift 'OAuthClient' resource used to setup identity federation on the OpenShift side. Auto-generated when left blank. See also the 'OAuthClientName' field.",
								MarkdownDescription: "Name of the secret set in the OpenShift 'OAuthClient' resource used to setup identity federation on the OpenShift side. Auto-generated when left blank. See also the 'OAuthClientName' field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"open_shifto_auth": schema.BoolAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Enables the integration of the identity provider (Keycloak / RHSSO) with OpenShift OAuth. Empty value on OpenShift by default. This will allow users to directly login with their OpenShift user through the OpenShift login, and have their workspaces created under personal OpenShift namespaces. WARNING: the 'kubeadmin' user is NOT supported, and logging through it will NOT allow accessing the Che Dashboard.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Enables the integration of the identity provider (Keycloak / RHSSO) with OpenShift OAuth. Empty value on OpenShift by default. This will allow users to directly login with their OpenShift user through the OpenShift login, and have their workspaces created under personal OpenShift namespaces. WARNING: the 'kubeadmin' user is NOT supported, and logging through it will NOT allow accessing the Che Dashboard.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"update_admin_password": schema.BoolAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Forces the default 'admin' Che user to update password on first login. Defaults to 'false'.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Forces the default 'admin' Che user to update password on first login. Defaults to 'false'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"dashboard": schema.SingleNestedAttribute{
						Description:         "Configuration settings related to the User Dashboard used by the Che installation.",
						MarkdownDescription: "Configuration settings related to the User Dashboard used by the Che installation.",
						Attributes: map[string]schema.Attribute{
							"warning": schema.StringAttribute{
								Description:         "Warning message that will be displayed on the User Dashboard",
								MarkdownDescription: "Warning message that will be displayed on the User Dashboard",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"database": schema.SingleNestedAttribute{
						Description:         "Configuration settings related to the database used by the Che installation.",
						MarkdownDescription: "Configuration settings related to the database used by the Che installation.",
						Attributes: map[string]schema.Attribute{
							"che_postgres_container_resources": schema.SingleNestedAttribute{
								Description:         "PostgreSQL container custom settings",
								MarkdownDescription: "PostgreSQL container custom settings",
								Attributes: map[string]schema.Attribute{
									"limits": schema.SingleNestedAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed.",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.",
										Attributes: map[string]schema.Attribute{
											"cpu": schema.StringAttribute{
												Description:         "CPU, in cores. (500m = .5 cores)",
												MarkdownDescription: "CPU, in cores. (500m = .5 cores)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"memory": schema.StringAttribute{
												Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
												MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"request": schema.SingleNestedAttribute{
										Description:         "Requests describes the minimum amount of compute resources required.",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required.",
										Attributes: map[string]schema.Attribute{
											"cpu": schema.StringAttribute{
												Description:         "CPU, in cores. (500m = .5 cores)",
												MarkdownDescription: "CPU, in cores. (500m = .5 cores)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"memory": schema.StringAttribute{
												Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
												MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
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

							"che_postgres_db": schema.StringAttribute{
								Description:         "PostgreSQL database name that the Che server uses to connect to the DB. Defaults to 'dbche'.",
								MarkdownDescription: "PostgreSQL database name that the Che server uses to connect to the DB. Defaults to 'dbche'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"che_postgres_host_name": schema.StringAttribute{
								Description:         "PostgreSQL Database host name that the Che server uses to connect to. Defaults is 'postgres'. Override this value ONLY when using an external database. See field 'externalDb'. In the default case it will be automatically set by the Operator.",
								MarkdownDescription: "PostgreSQL Database host name that the Che server uses to connect to. Defaults is 'postgres'. Override this value ONLY when using an external database. See field 'externalDb'. In the default case it will be automatically set by the Operator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"che_postgres_password": schema.StringAttribute{
								Description:         "PostgreSQL password that the Che server uses to connect to the DB. When omitted or left blank, it will be set to an automatically generated value.",
								MarkdownDescription: "PostgreSQL password that the Che server uses to connect to the DB. When omitted or left blank, it will be set to an automatically generated value.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"che_postgres_port": schema.StringAttribute{
								Description:         "PostgreSQL Database port that the Che server uses to connect to. Defaults to 5432. Override this value ONLY when using an external database. See field 'externalDb'. In the default case it will be automatically set by the Operator.",
								MarkdownDescription: "PostgreSQL Database port that the Che server uses to connect to. Defaults to 5432. Override this value ONLY when using an external database. See field 'externalDb'. In the default case it will be automatically set by the Operator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"che_postgres_secret": schema.StringAttribute{
								Description:         "The secret that contains PostgreSQL'user' and 'password' that the Che server uses to connect to the DB. When the secret is defined, the 'chePostgresUser' and 'chePostgresPassword' are ignored. When the value is omitted or left blank, the one of following scenarios applies: 1. 'chePostgresUser' and 'chePostgresPassword' are defined, then they will be used to connect to the DB. 2. 'chePostgresUser' or 'chePostgresPassword' are not defined, then a new secret with the name 'postgres-credentials' will be created with default value of 'pgche' for 'user' and with an auto-generated value for 'password'. The secret must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",
								MarkdownDescription: "The secret that contains PostgreSQL'user' and 'password' that the Che server uses to connect to the DB. When the secret is defined, the 'chePostgresUser' and 'chePostgresPassword' are ignored. When the value is omitted or left blank, the one of following scenarios applies: 1. 'chePostgresUser' and 'chePostgresPassword' are defined, then they will be used to connect to the DB. 2. 'chePostgresUser' or 'chePostgresPassword' are not defined, then a new secret with the name 'postgres-credentials' will be created with default value of 'pgche' for 'user' and with an auto-generated value for 'password'. The secret must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"che_postgres_user": schema.StringAttribute{
								Description:         "PostgreSQL user that the Che server uses to connect to the DB. Defaults to 'pgche'.",
								MarkdownDescription: "PostgreSQL user that the Che server uses to connect to the DB. Defaults to 'pgche'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"external_db": schema.BoolAttribute{
								Description:         "Instructs the Operator on whether to deploy a dedicated database. By default, a dedicated PostgreSQL database is deployed as part of the Che installation. When 'externalDb' is 'true', no dedicated database will be deployed by the Operator and you will need to provide connection details to the external DB you are about to use. See also all the fields starting with: 'chePostgres'.",
								MarkdownDescription: "Instructs the Operator on whether to deploy a dedicated database. By default, a dedicated PostgreSQL database is deployed as part of the Che installation. When 'externalDb' is 'true', no dedicated database will be deployed by the Operator and you will need to provide connection details to the external DB you are about to use. See also all the fields starting with: 'chePostgres'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"postgres_env": schema.ListNestedAttribute{
								Description:         "List of environment variables to set in the PostgreSQL container.",
								MarkdownDescription: "List of environment variables to set in the PostgreSQL container.",
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

							"postgres_image": schema.StringAttribute{
								Description:         "Overrides the container image used in the PostgreSQL database deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",
								MarkdownDescription: "Overrides the container image used in the PostgreSQL database deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"postgres_image_pull_policy": schema.StringAttribute{
								Description:         "Overrides the image pull policy used in the PostgreSQL database deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
								MarkdownDescription: "Overrides the image pull policy used in the PostgreSQL database deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"postgres_version": schema.StringAttribute{
								Description:         "Indicates a PostgreSQL version image to use. Allowed values are: '9.6' and '13.3'. Migrate your PostgreSQL database to switch from one version to another.",
								MarkdownDescription: "Indicates a PostgreSQL version image to use. Allowed values are: '9.6' and '13.3'. Migrate your PostgreSQL database to switch from one version to another.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pvc_claim_size": schema.StringAttribute{
								Description:         "Size of the persistent volume claim for database. Defaults to '1Gi'. To update pvc storageclass that provisions it must support resize when Eclipse Che has been already deployed.",
								MarkdownDescription: "Size of the persistent volume claim for database. Defaults to '1Gi'. To update pvc storageclass that provisions it must support resize when Eclipse Che has been already deployed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"dev_workspace": schema.SingleNestedAttribute{
						Description:         "DevWorkspace operator configuration",
						MarkdownDescription: "DevWorkspace operator configuration",
						Attributes: map[string]schema.Attribute{
							"controller_image": schema.StringAttribute{
								Description:         "Overrides the container image used in the DevWorkspace controller deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",
								MarkdownDescription: "Overrides the container image used in the DevWorkspace controller deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable": schema.BoolAttribute{
								Description:         "Deploys the DevWorkspace Operator in the cluster. Does nothing when a matching version of the Operator is already installed. Fails when a non-matching version of the Operator is already installed.",
								MarkdownDescription: "Deploys the DevWorkspace Operator in the cluster. Does nothing when a matching version of the Operator is already installed. Fails when a non-matching version of the Operator is already installed.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"env": schema.ListNestedAttribute{
								Description:         "List of environment variables to set in the DevWorkspace container.",
								MarkdownDescription: "List of environment variables to set in the DevWorkspace container.",
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

							"running_limit": schema.StringAttribute{
								Description:         "Maximum number of the running workspaces per user.",
								MarkdownDescription: "Maximum number of the running workspaces per user.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"seconds_of_inactivity_before_idling": schema.Int64Attribute{
								Description:         "Idle timeout for workspaces in seconds. This timeout is the duration after which a workspace will be idled if there is no activity. To disable workspace idling due to inactivity, set this value to -1.",
								MarkdownDescription: "Idle timeout for workspaces in seconds. This timeout is the duration after which a workspace will be idled if there is no activity. To disable workspace idling due to inactivity, set this value to -1.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"seconds_of_run_before_idling": schema.Int64Attribute{
								Description:         "Run timeout for workspaces in seconds. This timeout is the maximum duration a workspace runs. To disable workspace run timeout, set this value to -1.",
								MarkdownDescription: "Run timeout for workspaces in seconds. This timeout is the maximum duration a workspace runs. To disable workspace run timeout, set this value to -1.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"git_services": schema.SingleNestedAttribute{
						Description:         "A configuration that allows users to work with remote Git repositories.",
						MarkdownDescription: "A configuration that allows users to work with remote Git repositories.",
						Attributes: map[string]schema.Attribute{
							"bitbucket": schema.ListNestedAttribute{
								Description:         "Enables users to work with repositories hosted on Bitbucket (bitbucket.org or self-hosted).",
								MarkdownDescription: "Enables users to work with repositories hosted on Bitbucket (bitbucket.org or self-hosted).",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"endpoint": schema.StringAttribute{
											Description:         "Bitbucket server endpoint URL. Deprecated in favor of 'che.eclipse.org/scm-server-endpoint' annotation. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-1-for-a-bitbucket-server/.",
											MarkdownDescription: "Bitbucket server endpoint URL. Deprecated in favor of 'che.eclipse.org/scm-server-endpoint' annotation. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-1-for-a-bitbucket-server/.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_name": schema.StringAttribute{
											Description:         "Kubernetes secret, that contains Base64-encoded Bitbucket OAuth 1.0 or OAuth 2.0 data. See the following pages for details: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-1-for-a-bitbucket-server/ and https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-the-bitbucket-cloud/.",
											MarkdownDescription: "Kubernetes secret, that contains Base64-encoded Bitbucket OAuth 1.0 or OAuth 2.0 data. See the following pages for details: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-1-for-a-bitbucket-server/ and https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-the-bitbucket-cloud/.",
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

							"github": schema.ListNestedAttribute{
								Description:         "Enables users to work with repositories hosted on GitHub (github.com or GitHub Enterprise).",
								MarkdownDescription: "Enables users to work with repositories hosted on GitHub (github.com or GitHub Enterprise).",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"endpoint": schema.StringAttribute{
											Description:         "GitHub server endpoint URL. Deprecated in favor of 'che.eclipse.org/scm-server-endpoint' annotation. See the following page for details: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-github/.",
											MarkdownDescription: "GitHub server endpoint URL. Deprecated in favor of 'che.eclipse.org/scm-server-endpoint' annotation. See the following page for details: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-github/.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_name": schema.StringAttribute{
											Description:         "Kubernetes secret, that contains Base64-encoded GitHub OAuth Client id and GitHub OAuth Client secret. See the following page for details: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-github/.",
											MarkdownDescription: "Kubernetes secret, that contains Base64-encoded GitHub OAuth Client id and GitHub OAuth Client secret. See the following page for details: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-github/.",
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

							"gitlab": schema.ListNestedAttribute{
								Description:         "Enables users to work with repositories hosted on GitLab (gitlab.com or self-hosted).",
								MarkdownDescription: "Enables users to work with repositories hosted on GitLab (gitlab.com or self-hosted).",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"endpoint": schema.StringAttribute{
											Description:         "GitLab server endpoint URL. Deprecated in favor of 'che.eclipse.org/scm-server-endpoint' annotation. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-gitlab/.",
											MarkdownDescription: "GitLab server endpoint URL. Deprecated in favor of 'che.eclipse.org/scm-server-endpoint' annotation. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-gitlab/.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_name": schema.StringAttribute{
											Description:         "Kubernetes secret, that contains Base64-encoded GitHub Application id and GitLab Application Client secret. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-gitlab/.",
											MarkdownDescription: "Kubernetes secret, that contains Base64-encoded GitHub Application id and GitLab Application Client secret. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-gitlab/.",
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

					"image_puller": schema.SingleNestedAttribute{
						Description:         "Kubernetes Image Puller configuration",
						MarkdownDescription: "Kubernetes Image Puller configuration",
						Attributes: map[string]schema.Attribute{
							"enable": schema.BoolAttribute{
								Description:         "Install and configure the Community Supported Kubernetes Image Puller Operator. When set to 'true' and no spec is provided, it will create a default KubernetesImagePuller object to be managed by the Operator. When set to 'false', the KubernetesImagePuller object will be deleted, and the Operator will be uninstalled, regardless of whether a spec is provided. If the 'spec.images' field is empty, a set of recommended workspace-related images will be automatically detected and pre-pulled after installation. Note that while this Operator and its behavior is community-supported, its payload may be commercially-supported for pulling commercially-supported images.",
								MarkdownDescription: "Install and configure the Community Supported Kubernetes Image Puller Operator. When set to 'true' and no spec is provided, it will create a default KubernetesImagePuller object to be managed by the Operator. When set to 'false', the KubernetesImagePuller object will be deleted, and the Operator will be uninstalled, regardless of whether a spec is provided. If the 'spec.images' field is empty, a set of recommended workspace-related images will be automatically detected and pre-pulled after installation. Note that while this Operator and its behavior is community-supported, its payload may be commercially-supported for pulling commercially-supported images.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"spec": schema.SingleNestedAttribute{
								Description:         "A KubernetesImagePullerSpec to configure the image puller in the CheCluster",
								MarkdownDescription: "A KubernetesImagePullerSpec to configure the image puller in the CheCluster",
								Attributes: map[string]schema.Attribute{
									"affinity": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"caching_cpu_limit": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"caching_cpu_request": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"caching_interval_hours": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"caching_memory_limit": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"caching_memory_request": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"config_map_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"daemonset_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"deployment_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"image_pull_secrets": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"image_puller_image": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"images": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"node_selector": schema.StringAttribute{
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

					"k8s": schema.SingleNestedAttribute{
						Description:         "Configuration settings specific to Che installations made on upstream Kubernetes.",
						MarkdownDescription: "Configuration settings specific to Che installations made on upstream Kubernetes.",
						Attributes: map[string]schema.Attribute{
							"ingress_class": schema.StringAttribute{
								Description:         "Ingress class that will define the which controller will manage ingresses. Defaults to 'nginx'. NB: This drives the 'kubernetes.io/ingress.class' annotation on Che-related ingresses.",
								MarkdownDescription: "Ingress class that will define the which controller will manage ingresses. Defaults to 'nginx'. NB: This drives the 'kubernetes.io/ingress.class' annotation on Che-related ingresses.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ingress_domain": schema.StringAttribute{
								Description:         "Global ingress domain for a Kubernetes cluster. This MUST be explicitly specified: there are no defaults.",
								MarkdownDescription: "Global ingress domain for a Kubernetes cluster. This MUST be explicitly specified: there are no defaults.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ingress_strategy": schema.StringAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Strategy for ingress creation. Options are: 'multi-host' (host is explicitly provided in ingress), 'single-host' (host is provided, path-based rules) and 'default-host' (no host is provided, path-based rules). Defaults to 'multi-host' Deprecated in favor of 'serverExposureStrategy' in the 'server' section, which defines this regardless of the cluster type. When both are defined, the 'serverExposureStrategy' option takes precedence.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Strategy for ingress creation. Options are: 'multi-host' (host is explicitly provided in ingress), 'single-host' (host is provided, path-based rules) and 'default-host' (no host is provided, path-based rules). Defaults to 'multi-host' Deprecated in favor of 'serverExposureStrategy' in the 'server' section, which defines this regardless of the cluster type. When both are defined, the 'serverExposureStrategy' option takes precedence.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"security_context_fs_group": schema.StringAttribute{
								Description:         "The FSGroup in which the Che Pod and workspace Pods containers runs in. Default value is '1724'.",
								MarkdownDescription: "The FSGroup in which the Che Pod and workspace Pods containers runs in. Default value is '1724'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"security_context_run_as_user": schema.StringAttribute{
								Description:         "ID of the user the Che Pod and workspace Pods containers run as. Default value is '1724'.",
								MarkdownDescription: "ID of the user the Che Pod and workspace Pods containers run as. Default value is '1724'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"single_host_exposure_type": schema.StringAttribute{
								Description:         "Deprecated. The value of this flag is ignored. When the serverExposureStrategy is set to 'single-host', the way the server, registries and workspaces are exposed is further configured by this property. The possible values are 'native', which means that the server and workspaces are exposed using ingresses on K8s or 'gateway' where the server and workspaces are exposed using a custom gateway based on link:https://doc.traefik.io/traefik/[Traefik]. All the endpoints whether backed by the ingress or gateway 'route' always point to the subpaths on the same domain. Defaults to 'native'.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. When the serverExposureStrategy is set to 'single-host', the way the server, registries and workspaces are exposed is further configured by this property. The possible values are 'native', which means that the server and workspaces are exposed using ingresses on K8s or 'gateway' where the server and workspaces are exposed using a custom gateway based on link:https://doc.traefik.io/traefik/[Traefik]. All the endpoints whether backed by the ingress or gateway 'route' always point to the subpaths on the same domain. Defaults to 'native'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls_secret_name": schema.StringAttribute{
								Description:         "Name of a secret that will be used to setup ingress TLS termination when TLS is enabled. When the field is empty string, the default cluster certificate will be used. See also the 'tlsSupport' field.",
								MarkdownDescription: "Name of a secret that will be used to setup ingress TLS termination when TLS is enabled. When the field is empty string, the default cluster certificate will be used. See also the 'tlsSupport' field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"metrics": schema.SingleNestedAttribute{
						Description:         "Configuration settings related to the metrics collection used by the Che installation.",
						MarkdownDescription: "Configuration settings related to the metrics collection used by the Che installation.",
						Attributes: map[string]schema.Attribute{
							"enable": schema.BoolAttribute{
								Description:         "Enables 'metrics' the Che server endpoint. Default to 'true'.",
								MarkdownDescription: "Enables 'metrics' the Che server endpoint. Default to 'true'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"server": schema.SingleNestedAttribute{
						Description:         "General configuration settings related to the Che server, the plugin and devfile registries",
						MarkdownDescription: "General configuration settings related to the Che server, the plugin and devfile registries",
						Attributes: map[string]schema.Attribute{
							"air_gap_container_registry_hostname": schema.StringAttribute{
								Description:         "Optional host name, or URL, to an alternate container registry to pull images from. This value overrides the container registry host name defined in all the default container images involved in a Che deployment. This is particularly useful to install Che in a restricted environment.",
								MarkdownDescription: "Optional host name, or URL, to an alternate container registry to pull images from. This value overrides the container registry host name defined in all the default container images involved in a Che deployment. This is particularly useful to install Che in a restricted environment.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"air_gap_container_registry_organization": schema.StringAttribute{
								Description:         "Optional repository name of an alternate container registry to pull images from. This value overrides the container registry organization defined in all the default container images involved in a Che deployment. This is particularly useful to install Eclipse Che in a restricted environment.",
								MarkdownDescription: "Optional repository name of an alternate container registry to pull images from. This value overrides the container registry organization defined in all the default container images involved in a Che deployment. This is particularly useful to install Eclipse Che in a restricted environment.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"allow_auto_provision_user_namespace": schema.BoolAttribute{
								Description:         "Indicates if is allowed to automatically create a user namespace. If it set to false, then user namespace must be pre-created by a cluster administrator.",
								MarkdownDescription: "Indicates if is allowed to automatically create a user namespace. If it set to false, then user namespace must be pre-created by a cluster administrator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"allow_user_defined_workspace_namespaces": schema.BoolAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Defines that a user is allowed to specify a Kubernetes namespace, or an OpenShift project, which differs from the default. It's NOT RECOMMENDED to set to 'true' without OpenShift OAuth configured. The OpenShift infrastructure also uses this property.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Defines that a user is allowed to specify a Kubernetes namespace, or an OpenShift project, which differs from the default. It's NOT RECOMMENDED to set to 'true' without OpenShift OAuth configured. The OpenShift infrastructure also uses this property.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"che_cluster_roles": schema.StringAttribute{
								Description:         "A comma-separated list of ClusterRoles that will be assigned to Che ServiceAccount. Each role must have 'app.kubernetes.io/part-of=che.eclipse.org' label. Be aware that the Che Operator has to already have all permissions in these ClusterRoles to grant them.",
								MarkdownDescription: "A comma-separated list of ClusterRoles that will be assigned to Che ServiceAccount. Each role must have 'app.kubernetes.io/part-of=che.eclipse.org' label. Be aware that the Che Operator has to already have all permissions in these ClusterRoles to grant them.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"che_debug": schema.StringAttribute{
								Description:         "Enables the debug mode for Che server. Defaults to 'false'.",
								MarkdownDescription: "Enables the debug mode for Che server. Defaults to 'false'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"che_flavor": schema.StringAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Specifies a variation of the installation. The options are  'che' for upstream Che installations or 'devspaces' for Red Hat OpenShift Dev Spaces (formerly Red Hat CodeReady Workspaces) installation",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Specifies a variation of the installation. The options are  'che' for upstream Che installations or 'devspaces' for Red Hat OpenShift Dev Spaces (formerly Red Hat CodeReady Workspaces) installation",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"che_host": schema.StringAttribute{
								Description:         "Public host name of the installed Che server. When value is omitted, the value it will be automatically set by the Operator. See the 'cheHostTLSSecret' field.",
								MarkdownDescription: "Public host name of the installed Che server. When value is omitted, the value it will be automatically set by the Operator. See the 'cheHostTLSSecret' field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"che_host_tls_secret": schema.StringAttribute{
								Description:         "Name of a secret containing certificates to secure ingress or route for the custom host name of the installed Che server. The secret must have 'app.kubernetes.io/part-of=che.eclipse.org' label. See the 'cheHost' field.",
								MarkdownDescription: "Name of a secret containing certificates to secure ingress or route for the custom host name of the installed Che server. The secret must have 'app.kubernetes.io/part-of=che.eclipse.org' label. See the 'cheHost' field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"che_image": schema.StringAttribute{
								Description:         "Overrides the container image used in Che deployment. This does NOT include the container image tag. Omit it or leave it empty to use the default container image provided by the Operator.",
								MarkdownDescription: "Overrides the container image used in Che deployment. This does NOT include the container image tag. Omit it or leave it empty to use the default container image provided by the Operator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"che_image_pull_policy": schema.StringAttribute{
								Description:         "Overrides the image pull policy used in Che deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
								MarkdownDescription: "Overrides the image pull policy used in Che deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"che_image_tag": schema.StringAttribute{
								Description:         "Overrides the tag of the container image used in Che deployment. Omit it or leave it empty to use the default image tag provided by the Operator.",
								MarkdownDescription: "Overrides the tag of the container image used in Che deployment. Omit it or leave it empty to use the default image tag provided by the Operator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"che_log_level": schema.StringAttribute{
								Description:         "Log level for the Che server: 'INFO' or 'DEBUG'. Defaults to 'INFO'.",
								MarkdownDescription: "Log level for the Che server: 'INFO' or 'DEBUG'. Defaults to 'INFO'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"che_server_env": schema.ListNestedAttribute{
								Description:         "List of environment variables to set in the Che server container.",
								MarkdownDescription: "List of environment variables to set in the Che server container.",
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

							"che_server_ingress": schema.SingleNestedAttribute{
								Description:         "The Che server ingress custom settings.",
								MarkdownDescription: "The Che server ingress custom settings.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										MarkdownDescription: "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.StringAttribute{
										Description:         "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										MarkdownDescription: "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"che_server_route": schema.SingleNestedAttribute{
								Description:         "The Che server route custom settings.",
								MarkdownDescription: "The Che server route custom settings.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										MarkdownDescription: "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"domain": schema.StringAttribute{
										Description:         "Operator uses the domain to generate a hostname for a route. In a conjunction with labels it creates a route, which is served by a non-default Ingress controller. The generated host name will follow this pattern: '<route-name>-<route-namespace>.<domain>'.",
										MarkdownDescription: "Operator uses the domain to generate a hostname for a route. In a conjunction with labels it creates a route, which is served by a non-default Ingress controller. The generated host name will follow this pattern: '<route-name>-<route-namespace>.<domain>'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.StringAttribute{
										Description:         "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										MarkdownDescription: "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"che_workspace_cluster_role": schema.StringAttribute{
								Description:         "Custom cluster role bound to the user for the Che workspaces. The role must have 'app.kubernetes.io/part-of=che.eclipse.org' label. The default roles are used when omitted or left blank.",
								MarkdownDescription: "Custom cluster role bound to the user for the Che workspaces. The role must have 'app.kubernetes.io/part-of=che.eclipse.org' label. The default roles are used when omitted or left blank.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"custom_che_properties": schema.MapAttribute{
								Description:         "Map of additional environment variables that will be applied in the generated 'che' ConfigMap to be used by the Che server, in addition to the values already generated from other fields of the 'CheCluster' custom resource (CR). When 'customCheProperties' contains a property that would be normally generated in 'che' ConfigMap from other CR fields, the value defined in the 'customCheProperties' is used instead.",
								MarkdownDescription: "Map of additional environment variables that will be applied in the generated 'che' ConfigMap to be used by the Che server, in addition to the values already generated from other fields of the 'CheCluster' custom resource (CR). When 'customCheProperties' contains a property that would be normally generated in 'che' ConfigMap from other CR fields, the value defined in the 'customCheProperties' is used instead.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dashboard_cpu_limit": schema.StringAttribute{
								Description:         "Overrides the CPU limit used in the dashboard deployment. In cores. (500m = .5 cores). Default to 500m.",
								MarkdownDescription: "Overrides the CPU limit used in the dashboard deployment. In cores. (500m = .5 cores). Default to 500m.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dashboard_cpu_request": schema.StringAttribute{
								Description:         "Overrides the CPU request used in the dashboard deployment. In cores. (500m = .5 cores). Default to 100m.",
								MarkdownDescription: "Overrides the CPU request used in the dashboard deployment. In cores. (500m = .5 cores). Default to 100m.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dashboard_env": schema.ListNestedAttribute{
								Description:         "List of environment variables to set in the dashboard container.",
								MarkdownDescription: "List of environment variables to set in the dashboard container.",
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

							"dashboard_image": schema.StringAttribute{
								Description:         "Overrides the container image used in the dashboard deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",
								MarkdownDescription: "Overrides the container image used in the dashboard deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dashboard_image_pull_policy": schema.StringAttribute{
								Description:         "Overrides the image pull policy used in the dashboard deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
								MarkdownDescription: "Overrides the image pull policy used in the dashboard deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dashboard_ingress": schema.SingleNestedAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Dashboard ingress custom settings.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Dashboard ingress custom settings.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										MarkdownDescription: "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.StringAttribute{
										Description:         "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										MarkdownDescription: "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"dashboard_memory_limit": schema.StringAttribute{
								Description:         "Overrides the memory limit used in the dashboard deployment. Defaults to 256Mi.",
								MarkdownDescription: "Overrides the memory limit used in the dashboard deployment. Defaults to 256Mi.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dashboard_memory_request": schema.StringAttribute{
								Description:         "Overrides the memory request used in the dashboard deployment. Defaults to 16Mi.",
								MarkdownDescription: "Overrides the memory request used in the dashboard deployment. Defaults to 16Mi.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dashboard_route": schema.SingleNestedAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Dashboard route custom settings.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Dashboard route custom settings.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										MarkdownDescription: "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"domain": schema.StringAttribute{
										Description:         "Operator uses the domain to generate a hostname for a route. In a conjunction with labels it creates a route, which is served by a non-default Ingress controller. The generated host name will follow this pattern: '<route-name>-<route-namespace>.<domain>'.",
										MarkdownDescription: "Operator uses the domain to generate a hostname for a route. In a conjunction with labels it creates a route, which is served by a non-default Ingress controller. The generated host name will follow this pattern: '<route-name>-<route-namespace>.<domain>'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.StringAttribute{
										Description:         "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										MarkdownDescription: "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"devfile_registry_cpu_limit": schema.StringAttribute{
								Description:         "Overrides the CPU limit used in the devfile registry deployment. In cores. (500m = .5 cores). Default to 500m.",
								MarkdownDescription: "Overrides the CPU limit used in the devfile registry deployment. In cores. (500m = .5 cores). Default to 500m.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"devfile_registry_cpu_request": schema.StringAttribute{
								Description:         "Overrides the CPU request used in the devfile registry deployment. In cores. (500m = .5 cores). Default to 100m.",
								MarkdownDescription: "Overrides the CPU request used in the devfile registry deployment. In cores. (500m = .5 cores). Default to 100m.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"devfile_registry_env": schema.ListNestedAttribute{
								Description:         "List of environment variables to set in the plugin registry container.",
								MarkdownDescription: "List of environment variables to set in the plugin registry container.",
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

							"devfile_registry_image": schema.StringAttribute{
								Description:         "Overrides the container image used in the devfile registry deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",
								MarkdownDescription: "Overrides the container image used in the devfile registry deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"devfile_registry_ingress": schema.SingleNestedAttribute{
								Description:         "Deprecated. The value of this flag is ignored. The devfile registry ingress custom settings.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. The devfile registry ingress custom settings.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										MarkdownDescription: "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.StringAttribute{
										Description:         "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										MarkdownDescription: "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"devfile_registry_memory_limit": schema.StringAttribute{
								Description:         "Overrides the memory limit used in the devfile registry deployment. Defaults to 256Mi.",
								MarkdownDescription: "Overrides the memory limit used in the devfile registry deployment. Defaults to 256Mi.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"devfile_registry_memory_request": schema.StringAttribute{
								Description:         "Overrides the memory request used in the devfile registry deployment. Defaults to 16Mi.",
								MarkdownDescription: "Overrides the memory request used in the devfile registry deployment. Defaults to 16Mi.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"devfile_registry_pull_policy": schema.StringAttribute{
								Description:         "Overrides the image pull policy used in the devfile registry deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
								MarkdownDescription: "Overrides the image pull policy used in the devfile registry deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"devfile_registry_route": schema.SingleNestedAttribute{
								Description:         "Deprecated. The value of this flag is ignored. The devfile registry route custom settings.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. The devfile registry route custom settings.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										MarkdownDescription: "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"domain": schema.StringAttribute{
										Description:         "Operator uses the domain to generate a hostname for a route. In a conjunction with labels it creates a route, which is served by a non-default Ingress controller. The generated host name will follow this pattern: '<route-name>-<route-namespace>.<domain>'.",
										MarkdownDescription: "Operator uses the domain to generate a hostname for a route. In a conjunction with labels it creates a route, which is served by a non-default Ingress controller. The generated host name will follow this pattern: '<route-name>-<route-namespace>.<domain>'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.StringAttribute{
										Description:         "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										MarkdownDescription: "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"devfile_registry_url": schema.StringAttribute{
								Description:         "Deprecated in favor of 'externalDevfileRegistries' fields.",
								MarkdownDescription: "Deprecated in favor of 'externalDevfileRegistries' fields.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_internal_cluster_svc_names": schema.BoolAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Disable internal cluster SVC names usage to communicate between components to speed up the traffic and avoid proxy issues.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Disable internal cluster SVC names usage to communicate between components to speed up the traffic and avoid proxy issues.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"external_devfile_registries": schema.ListNestedAttribute{
								Description:         "External devfile registries, that serves sample, ready-to-use devfiles. Configure this in addition to a dedicated devfile registry (when 'externalDevfileRegistry' is 'false') or instead of it (when 'externalDevfileRegistry' is 'true')",
								MarkdownDescription: "External devfile registries, that serves sample, ready-to-use devfiles. Configure this in addition to a dedicated devfile registry (when 'externalDevfileRegistry' is 'false') or instead of it (when 'externalDevfileRegistry' is 'true')",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"url": schema.StringAttribute{
											Description:         "Public URL of the devfile registry.",
											MarkdownDescription: "Public URL of the devfile registry.",
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

							"external_devfile_registry": schema.BoolAttribute{
								Description:         "Instructs the Operator on whether to deploy a dedicated devfile registry server. By default, a dedicated devfile registry server is started. When 'externalDevfileRegistry' is 'true', no such dedicated server will be started by the Operator and configure at least one devfile registry with 'externalDevfileRegistries' field.",
								MarkdownDescription: "Instructs the Operator on whether to deploy a dedicated devfile registry server. By default, a dedicated devfile registry server is started. When 'externalDevfileRegistry' is 'true', no such dedicated server will be started by the Operator and configure at least one devfile registry with 'externalDevfileRegistries' field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"external_plugin_registry": schema.BoolAttribute{
								Description:         "Instructs the Operator on whether to deploy a dedicated plugin registry server. By default, a dedicated plugin registry server is started. When 'externalPluginRegistry' is 'true', no such dedicated server will be started by the Operator and you will have to manually set the 'pluginRegistryUrl' field.",
								MarkdownDescription: "Instructs the Operator on whether to deploy a dedicated plugin registry server. By default, a dedicated plugin registry server is started. When 'externalPluginRegistry' is 'true', no such dedicated server will be started by the Operator and you will have to manually set the 'pluginRegistryUrl' field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"git_self_signed_cert": schema.BoolAttribute{
								Description:         "When enabled, the certificate from 'che-git-self-signed-cert' ConfigMap will be propagated to the Che components and provide particular configuration for Git. Note, the 'che-git-self-signed-cert' ConfigMap must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",
								MarkdownDescription: "When enabled, the certificate from 'che-git-self-signed-cert' ConfigMap will be propagated to the Che components and provide particular configuration for Git. Note, the 'che-git-self-signed-cert' ConfigMap must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"non_proxy_hosts": schema.StringAttribute{
								Description:         "List of hosts that will be reached directly, bypassing the proxy. Specify wild card domain use the following form '.<DOMAIN>' and '|' as delimiter, for example: 'localhost|.my.host.com|123.42.12.32' Only use when configuring a proxy is required. Operator respects OpenShift cluster wide proxy configuration and no additional configuration is required, but defining 'nonProxyHosts' in a custom resource leads to merging non proxy hosts lists from the cluster proxy configuration and ones defined in the custom resources. See the doc https://docs.openshift.com/container-platform/4.4/networking/enable-cluster-wide-proxy.html. See also the 'proxyURL' fields.",
								MarkdownDescription: "List of hosts that will be reached directly, bypassing the proxy. Specify wild card domain use the following form '.<DOMAIN>' and '|' as delimiter, for example: 'localhost|.my.host.com|123.42.12.32' Only use when configuring a proxy is required. Operator respects OpenShift cluster wide proxy configuration and no additional configuration is required, but defining 'nonProxyHosts' in a custom resource leads to merging non proxy hosts lists from the cluster proxy configuration and ones defined in the custom resources. See the doc https://docs.openshift.com/container-platform/4.4/networking/enable-cluster-wide-proxy.html. See also the 'proxyURL' fields.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"open_vsx_registry_url": schema.StringAttribute{
								Description:         "Open VSX registry URL. If omitted an embedded instance will be used.",
								MarkdownDescription: "Open VSX registry URL. If omitted an embedded instance will be used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"plugin_registry_cpu_limit": schema.StringAttribute{
								Description:         "Overrides the CPU limit used in the plugin registry deployment. In cores. (500m = .5 cores). Default to 500m.",
								MarkdownDescription: "Overrides the CPU limit used in the plugin registry deployment. In cores. (500m = .5 cores). Default to 500m.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"plugin_registry_cpu_request": schema.StringAttribute{
								Description:         "Overrides the CPU request used in the plugin registry deployment. In cores. (500m = .5 cores). Default to 100m.",
								MarkdownDescription: "Overrides the CPU request used in the plugin registry deployment. In cores. (500m = .5 cores). Default to 100m.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"plugin_registry_env": schema.ListNestedAttribute{
								Description:         "List of environment variables to set in the devfile registry container.",
								MarkdownDescription: "List of environment variables to set in the devfile registry container.",
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

							"plugin_registry_image": schema.StringAttribute{
								Description:         "Overrides the container image used in the plugin registry deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",
								MarkdownDescription: "Overrides the container image used in the plugin registry deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"plugin_registry_ingress": schema.SingleNestedAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Plugin registry ingress custom settings.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Plugin registry ingress custom settings.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										MarkdownDescription: "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.StringAttribute{
										Description:         "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										MarkdownDescription: "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"plugin_registry_memory_limit": schema.StringAttribute{
								Description:         "Overrides the memory limit used in the plugin registry deployment. Defaults to 1536Mi.",
								MarkdownDescription: "Overrides the memory limit used in the plugin registry deployment. Defaults to 1536Mi.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"plugin_registry_memory_request": schema.StringAttribute{
								Description:         "Overrides the memory request used in the plugin registry deployment. Defaults to 16Mi.",
								MarkdownDescription: "Overrides the memory request used in the plugin registry deployment. Defaults to 16Mi.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"plugin_registry_pull_policy": schema.StringAttribute{
								Description:         "Overrides the image pull policy used in the plugin registry deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
								MarkdownDescription: "Overrides the image pull policy used in the plugin registry deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"plugin_registry_route": schema.SingleNestedAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Plugin registry route custom settings.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Plugin registry route custom settings.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										MarkdownDescription: "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"domain": schema.StringAttribute{
										Description:         "Operator uses the domain to generate a hostname for a route. In a conjunction with labels it creates a route, which is served by a non-default Ingress controller. The generated host name will follow this pattern: '<route-name>-<route-namespace>.<domain>'.",
										MarkdownDescription: "Operator uses the domain to generate a hostname for a route. In a conjunction with labels it creates a route, which is served by a non-default Ingress controller. The generated host name will follow this pattern: '<route-name>-<route-namespace>.<domain>'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.StringAttribute{
										Description:         "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										MarkdownDescription: "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"plugin_registry_url": schema.StringAttribute{
								Description:         "Public URL of the plugin registry that serves sample ready-to-use devfiles. Set this ONLY when a use of an external devfile registry is needed. See the 'externalPluginRegistry' field. By default, this will be automatically calculated by the Operator.",
								MarkdownDescription: "Public URL of the plugin registry that serves sample ready-to-use devfiles. Set this ONLY when a use of an external devfile registry is needed. See the 'externalPluginRegistry' field. By default, this will be automatically calculated by the Operator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_password": schema.StringAttribute{
								Description:         "Password of the proxy server. Only use when proxy configuration is required. See the 'proxyURL', 'proxyUser' and 'proxySecret' fields.",
								MarkdownDescription: "Password of the proxy server. Only use when proxy configuration is required. See the 'proxyURL', 'proxyUser' and 'proxySecret' fields.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_port": schema.StringAttribute{
								Description:         "Port of the proxy server. Only use when configuring a proxy is required. See also the 'proxyURL' and 'nonProxyHosts' fields.",
								MarkdownDescription: "Port of the proxy server. Only use when configuring a proxy is required. See also the 'proxyURL' and 'nonProxyHosts' fields.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_secret": schema.StringAttribute{
								Description:         "The secret that contains 'user' and 'password' for a proxy server. When the secret is defined, the 'proxyUser' and 'proxyPassword' are ignored. The secret must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",
								MarkdownDescription: "The secret that contains 'user' and 'password' for a proxy server. When the secret is defined, the 'proxyUser' and 'proxyPassword' are ignored. The secret must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_url": schema.StringAttribute{
								Description:         "URL (protocol+host name) of the proxy server. This drives the appropriate changes in the 'JAVA_OPTS' and 'https(s)_proxy' variables in the Che server and workspaces containers. Only use when configuring a proxy is required. Operator respects OpenShift cluster wide proxy configuration and no additional configuration is required, but defining 'proxyUrl' in a custom resource leads to overrides the cluster proxy configuration with fields 'proxyUrl', 'proxyPort', 'proxyUser' and 'proxyPassword' from the custom resource. See the doc https://docs.openshift.com/container-platform/4.4/networking/enable-cluster-wide-proxy.html. See also the 'proxyPort' and 'nonProxyHosts' fields.",
								MarkdownDescription: "URL (protocol+host name) of the proxy server. This drives the appropriate changes in the 'JAVA_OPTS' and 'https(s)_proxy' variables in the Che server and workspaces containers. Only use when configuring a proxy is required. Operator respects OpenShift cluster wide proxy configuration and no additional configuration is required, but defining 'proxyUrl' in a custom resource leads to overrides the cluster proxy configuration with fields 'proxyUrl', 'proxyPort', 'proxyUser' and 'proxyPassword' from the custom resource. See the doc https://docs.openshift.com/container-platform/4.4/networking/enable-cluster-wide-proxy.html. See also the 'proxyPort' and 'nonProxyHosts' fields.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_user": schema.StringAttribute{
								Description:         "User name of the proxy server. Only use when configuring a proxy is required. See also the 'proxyURL', 'proxyPassword' and 'proxySecret' fields.",
								MarkdownDescription: "User name of the proxy server. Only use when configuring a proxy is required. See also the 'proxyURL', 'proxyPassword' and 'proxySecret' fields.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"self_signed_cert": schema.BoolAttribute{
								Description:         "Deprecated. The value of this flag is ignored. The Che Operator will automatically detect whether the router certificate is self-signed and propagate it to other components, such as the Che server.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. The Che Operator will automatically detect whether the router certificate is self-signed and propagate it to other components, such as the Che server.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"server_cpu_limit": schema.StringAttribute{
								Description:         "Overrides the CPU limit used in the Che server deployment In cores. (500m = .5 cores). Default to 1.",
								MarkdownDescription: "Overrides the CPU limit used in the Che server deployment In cores. (500m = .5 cores). Default to 1.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"server_cpu_request": schema.StringAttribute{
								Description:         "Overrides the CPU request used in the Che server deployment In cores. (500m = .5 cores). Default to 100m.",
								MarkdownDescription: "Overrides the CPU request used in the Che server deployment In cores. (500m = .5 cores). Default to 100m.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"server_exposure_strategy": schema.StringAttribute{
								Description:         "Deprecated. The value of this flag is ignored. Sets the server and workspaces exposure type. Possible values are 'multi-host', 'single-host', 'default-host'. Defaults to 'multi-host', which creates a separate ingress, or OpenShift routes, for every required endpoint. 'single-host' makes Che exposed on a single host name with workspaces exposed on subpaths. Read the docs to learn about the limitations of this approach. Also consult the 'singleHostExposureType' property to further configure how the Operator and the Che server make that happen on Kubernetes. 'default-host' exposes the Che server on the host of the cluster. Read the docs to learn about the limitations of this approach.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Sets the server and workspaces exposure type. Possible values are 'multi-host', 'single-host', 'default-host'. Defaults to 'multi-host', which creates a separate ingress, or OpenShift routes, for every required endpoint. 'single-host' makes Che exposed on a single host name with workspaces exposed on subpaths. Read the docs to learn about the limitations of this approach. Also consult the 'singleHostExposureType' property to further configure how the Operator and the Che server make that happen on Kubernetes. 'default-host' exposes the Che server on the host of the cluster. Read the docs to learn about the limitations of this approach.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"server_memory_limit": schema.StringAttribute{
								Description:         "Overrides the memory limit used in the Che server deployment. Defaults to 1Gi.",
								MarkdownDescription: "Overrides the memory limit used in the Che server deployment. Defaults to 1Gi.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"server_memory_request": schema.StringAttribute{
								Description:         "Overrides the memory request used in the Che server deployment. Defaults to 512Mi.",
								MarkdownDescription: "Overrides the memory request used in the Che server deployment. Defaults to 512Mi.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"server_trust_store_config_map_name": schema.StringAttribute{
								Description:         "Name of the ConfigMap with public certificates to add to Java trust store of the Che server. This is often required when adding the OpenShift OAuth provider, which has HTTPS endpoint signed with self-signed cert. The Che server must be aware of its CA cert to be able to request it. This is disabled by default. The Config Map must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",
								MarkdownDescription: "Name of the ConfigMap with public certificates to add to Java trust store of the Che server. This is often required when adding the OpenShift OAuth provider, which has HTTPS endpoint signed with self-signed cert. The Che server must be aware of its CA cert to be able to request it. This is disabled by default. The Config Map must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"single_host_gateway_config_map_labels": schema.MapAttribute{
								Description:         "The labels that need to be present in the ConfigMaps representing the gateway configuration.",
								MarkdownDescription: "The labels that need to be present in the ConfigMaps representing the gateway configuration.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"single_host_gateway_config_sidecar_image": schema.StringAttribute{
								Description:         "The image used for the gateway sidecar that provides configuration to the gateway. Omit it or leave it empty to use the default container image provided by the Operator.",
								MarkdownDescription: "The image used for the gateway sidecar that provides configuration to the gateway. Omit it or leave it empty to use the default container image provided by the Operator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"single_host_gateway_image": schema.StringAttribute{
								Description:         "The image used for the gateway in the single host mode. Omit it or leave it empty to use the default container image provided by the Operator.",
								MarkdownDescription: "The image used for the gateway in the single host mode. Omit it or leave it empty to use the default container image provided by the Operator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls_support": schema.BoolAttribute{
								Description:         "Deprecated. Instructs the Operator to deploy Che in TLS mode. This is enabled by default. Disabling TLS sometimes cause malfunction of some Che components.",
								MarkdownDescription: "Deprecated. Instructs the Operator to deploy Che in TLS mode. This is enabled by default. Disabling TLS sometimes cause malfunction of some Che components.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"use_internal_cluster_svc_names": schema.BoolAttribute{
								Description:         "Deprecated in favor of 'disableInternalClusterSVCNames'.",
								MarkdownDescription: "Deprecated in favor of 'disableInternalClusterSVCNames'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"workspace_default_components": schema.ListNestedAttribute{
								Description:         "Default components applied to DevWorkspaces. These default components are meant to be used when a Devfile does not contain any components.",
								MarkdownDescription: "Default components applied to DevWorkspaces. These default components are meant to be used when a Devfile does not contain any components.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"attributes": schema.MapAttribute{
											Description:         "Map of implementation-dependant free-form YAML attributes.",
											MarkdownDescription: "Map of implementation-dependant free-form YAML attributes.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"component_type": schema.StringAttribute{
											Description:         "Type of component",
											MarkdownDescription: "Type of component",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Container", "Kubernetes", "Openshift", "Volume", "Image", "Plugin", "Custom"),
											},
										},

										"container": schema.SingleNestedAttribute{
											Description:         "Allows adding and configuring devworkspace-related containers",
											MarkdownDescription: "Allows adding and configuring devworkspace-related containers",
											Attributes: map[string]schema.Attribute{
												"annotation": schema.SingleNestedAttribute{
													Description:         "Annotations that should be added to specific resources for this container",
													MarkdownDescription: "Annotations that should be added to specific resources for this container",
													Attributes: map[string]schema.Attribute{
														"deployment": schema.MapAttribute{
															Description:         "Annotations to be added to deployment",
															MarkdownDescription: "Annotations to be added to deployment",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"service": schema.MapAttribute{
															Description:         "Annotations to be added to service",
															MarkdownDescription: "Annotations to be added to service",
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

												"args": schema.ListAttribute{
													Description:         "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command.  Defaults to an empty array, meaning use whatever is defined in the image.",
													MarkdownDescription: "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command.  Defaults to an empty array, meaning use whatever is defined in the image.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"command": schema.ListAttribute{
													Description:         "The command to run in the dockerimage component instead of the default one provided in the image.  Defaults to an empty array, meaning use whatever is defined in the image.",
													MarkdownDescription: "The command to run in the dockerimage component instead of the default one provided in the image.  Defaults to an empty array, meaning use whatever is defined in the image.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"cpu_limit": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"cpu_request": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"dedicated_pod": schema.BoolAttribute{
													Description:         "Specify if a container should run in its own separated pod, instead of running as part of the main development environment pod.  Default value is 'false'",
													MarkdownDescription: "Specify if a container should run in its own separated pod, instead of running as part of the main development environment pod.  Default value is 'false'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"endpoints": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"annotation": schema.MapAttribute{
																Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
																MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"attributes": schema.MapAttribute{
																Description:         "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
																MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"exposure": schema.StringAttribute{
																Description:         "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
																MarkdownDescription: "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("public", "internal", "none"),
																},
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																},
															},

															"path": schema.StringAttribute{
																Description:         "Path of the endpoint URL",
																MarkdownDescription: "Path of the endpoint URL",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"protocol": schema.StringAttribute{
																Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
																MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																},
															},

															"secure": schema.BoolAttribute{
																Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"target_port": schema.Int64Attribute{
																Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
																MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",
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

												"env": schema.ListNestedAttribute{
													Description:         "Environment variables used in this container.  The following variables are reserved and cannot be overridden via env:   - '$PROJECTS_ROOT'   - '$PROJECT_SOURCE'",
													MarkdownDescription: "Environment variables used in this container.  The following variables are reserved and cannot be overridden via env:   - '$PROJECTS_ROOT'   - '$PROJECT_SOURCE'",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"value": schema.StringAttribute{
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

												"image": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"memory_limit": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"memory_request": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mount_sources": schema.BoolAttribute{
													Description:         "Toggles whether or not the project source code should be mounted in the component.  Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",
													MarkdownDescription: "Toggles whether or not the project source code should be mounted in the component.  Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"source_mapping": schema.StringAttribute{
													Description:         "Optional specification of the path in the container where project sources should be transferred/mounted when 'mountSources' is 'true'. When omitted, the default value of /projects is used.",
													MarkdownDescription: "Optional specification of the path in the container where project sources should be transferred/mounted when 'mountSources' is 'true'. When omitted, the default value of /projects is used.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volume_mounts": schema.ListNestedAttribute{
													Description:         "List of volumes mounts that should be mounted is this container.",
													MarkdownDescription: "List of volumes mounts that should be mounted is this container.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "The volume mount name is the name of an existing 'Volume' component. If several containers mount the same volume name then they will reuse the same volume and will be able to access to the same files.",
																MarkdownDescription: "The volume mount name is the name of an existing 'Volume' component. If several containers mount the same volume name then they will reuse the same volume and will be able to access to the same files.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																},
															},

															"path": schema.StringAttribute{
																Description:         "The path in the component container where the volume should be mounted. If not path is mentioned, default path is the is '/<name>'.",
																MarkdownDescription: "The path in the component container where the volume should be mounted. If not path is mentioned, default path is the is '/<name>'.",
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

										"custom": schema.SingleNestedAttribute{
											Description:         "Custom component whose logic is implementation-dependant and should be provided by the user possibly through some dedicated controller",
											MarkdownDescription: "Custom component whose logic is implementation-dependant and should be provided by the user possibly through some dedicated controller",
											Attributes: map[string]schema.Attribute{
												"component_class": schema.StringAttribute{
													Description:         "Class of component that the associated implementation controller should use to process this command with the appropriate logic",
													MarkdownDescription: "Class of component that the associated implementation controller should use to process this command with the appropriate logic",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"embedded_resource": schema.MapAttribute{
													Description:         "Additional free-form configuration for this custom component that the implementation controller will know how to use",
													MarkdownDescription: "Additional free-form configuration for this custom component that the implementation controller will know how to use",
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

										"image": schema.SingleNestedAttribute{
											Description:         "Allows specifying the definition of an image for outer loop builds",
											MarkdownDescription: "Allows specifying the definition of an image for outer loop builds",
											Attributes: map[string]schema.Attribute{
												"auto_build": schema.BoolAttribute{
													Description:         "Defines if the image should be built during startup.  Default value is 'false'",
													MarkdownDescription: "Defines if the image should be built during startup.  Default value is 'false'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"dockerfile": schema.SingleNestedAttribute{
													Description:         "Allows specifying dockerfile type build",
													MarkdownDescription: "Allows specifying dockerfile type build",
													Attributes: map[string]schema.Attribute{
														"args": schema.ListAttribute{
															Description:         "The arguments to supply to the dockerfile build.",
															MarkdownDescription: "The arguments to supply to the dockerfile build.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"build_context": schema.StringAttribute{
															Description:         "Path of source directory to establish build context. Defaults to ${PROJECT_SOURCE} in the container",
															MarkdownDescription: "Path of source directory to establish build context. Defaults to ${PROJECT_SOURCE} in the container",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"devfile_registry": schema.SingleNestedAttribute{
															Description:         "Dockerfile's Devfile Registry source",
															MarkdownDescription: "Dockerfile's Devfile Registry source",
															Attributes: map[string]schema.Attribute{
																"id": schema.StringAttribute{
																	Description:         "Id in a devfile registry that contains a Dockerfile. The src in the OCI registry required for the Dockerfile build will be downloaded for building the image.",
																	MarkdownDescription: "Id in a devfile registry that contains a Dockerfile. The src in the OCI registry required for the Dockerfile build will be downloaded for building the image.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"registry_url": schema.StringAttribute{
																	Description:         "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src. To ensure the Dockerfile gets resolved consistently in different environments, it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",
																	MarkdownDescription: "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src. To ensure the Dockerfile gets resolved consistently in different environments, it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"git": schema.SingleNestedAttribute{
															Description:         "Dockerfile's Git source",
															MarkdownDescription: "Dockerfile's Git source",
															Attributes: map[string]schema.Attribute{
																"checkout_from": schema.SingleNestedAttribute{
																	Description:         "Defines from what the project should be checked out. Required if there are more than one remote configured",
																	MarkdownDescription: "Defines from what the project should be checked out. Required if there are more than one remote configured",
																	Attributes: map[string]schema.Attribute{
																		"remote": schema.StringAttribute{
																			Description:         "The remote name should be used as init. Required if there are more than one remote configured",
																			MarkdownDescription: "The remote name should be used as init. Required if there are more than one remote configured",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"revision": schema.StringAttribute{
																			Description:         "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
																			MarkdownDescription: "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"file_location": schema.StringAttribute{
																	Description:         "Location of the Dockerfile in the Git repository when using git as Dockerfile src. Defaults to Dockerfile.",
																	MarkdownDescription: "Location of the Dockerfile in the Git repository when using git as Dockerfile src. Defaults to Dockerfile.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"remotes": schema.MapAttribute{
																	Description:         "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
																	MarkdownDescription: "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
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

														"root_required": schema.BoolAttribute{
															Description:         "Specify if a privileged builder pod is required.  Default value is 'false'",
															MarkdownDescription: "Specify if a privileged builder pod is required.  Default value is 'false'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"src_type": schema.StringAttribute{
															Description:         "Type of Dockerfile src",
															MarkdownDescription: "Type of Dockerfile src",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("Uri", "DevfileRegistry", "Git"),
															},
														},

														"uri": schema.StringAttribute{
															Description:         "URI Reference of a Dockerfile. It can be a full URL or a relative URI from the current devfile as the base URI.",
															MarkdownDescription: "URI Reference of a Dockerfile. It can be a full URL or a relative URI from the current devfile as the base URI.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"image_name": schema.StringAttribute{
													Description:         "Name of the image for the resulting outerloop build",
													MarkdownDescription: "Name of the image for the resulting outerloop build",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"image_type": schema.StringAttribute{
													Description:         "Type of image",
													MarkdownDescription: "Type of image",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Dockerfile"),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"kubernetes": schema.SingleNestedAttribute{
											Description:         "Allows importing into the devworkspace the Kubernetes resources defined in a given manifest. For example this allows reusing the Kubernetes definitions used to deploy some runtime components in production.",
											MarkdownDescription: "Allows importing into the devworkspace the Kubernetes resources defined in a given manifest. For example this allows reusing the Kubernetes definitions used to deploy some runtime components in production.",
											Attributes: map[string]schema.Attribute{
												"deploy_by_default": schema.BoolAttribute{
													Description:         "Defines if the component should be deployed during startup.  Default value is 'false'",
													MarkdownDescription: "Defines if the component should be deployed during startup.  Default value is 'false'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"endpoints": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"annotation": schema.MapAttribute{
																Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
																MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"attributes": schema.MapAttribute{
																Description:         "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
																MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"exposure": schema.StringAttribute{
																Description:         "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
																MarkdownDescription: "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("public", "internal", "none"),
																},
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																},
															},

															"path": schema.StringAttribute{
																Description:         "Path of the endpoint URL",
																MarkdownDescription: "Path of the endpoint URL",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"protocol": schema.StringAttribute{
																Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
																MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																},
															},

															"secure": schema.BoolAttribute{
																Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"target_port": schema.Int64Attribute{
																Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
																MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",
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

												"inlined": schema.StringAttribute{
													Description:         "Inlined manifest",
													MarkdownDescription: "Inlined manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"location_type": schema.StringAttribute{
													Description:         "Type of Kubernetes-like location",
													MarkdownDescription: "Type of Kubernetes-like location",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Uri", "Inlined"),
													},
												},

												"uri": schema.StringAttribute{
													Description:         "Location in a file fetched from a uri.",
													MarkdownDescription: "Location in a file fetched from a uri.",
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
											Description:         "Mandatory name that allows referencing the component from other elements (such as commands) or from an external devfile that may reference this component through a parent or a plugin.",
											MarkdownDescription: "Mandatory name that allows referencing the component from other elements (such as commands) or from an external devfile that may reference this component through a parent or a plugin.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
											},
										},

										"openshift": schema.SingleNestedAttribute{
											Description:         "Allows importing into the devworkspace the OpenShift resources defined in a given manifest. For example this allows reusing the OpenShift definitions used to deploy some runtime components in production.",
											MarkdownDescription: "Allows importing into the devworkspace the OpenShift resources defined in a given manifest. For example this allows reusing the OpenShift definitions used to deploy some runtime components in production.",
											Attributes: map[string]schema.Attribute{
												"deploy_by_default": schema.BoolAttribute{
													Description:         "Defines if the component should be deployed during startup.  Default value is 'false'",
													MarkdownDescription: "Defines if the component should be deployed during startup.  Default value is 'false'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"endpoints": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"annotation": schema.MapAttribute{
																Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
																MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"attributes": schema.MapAttribute{
																Description:         "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
																MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"exposure": schema.StringAttribute{
																Description:         "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
																MarkdownDescription: "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("public", "internal", "none"),
																},
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																},
															},

															"path": schema.StringAttribute{
																Description:         "Path of the endpoint URL",
																MarkdownDescription: "Path of the endpoint URL",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"protocol": schema.StringAttribute{
																Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
																MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																},
															},

															"secure": schema.BoolAttribute{
																Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"target_port": schema.Int64Attribute{
																Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
																MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",
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

												"inlined": schema.StringAttribute{
													Description:         "Inlined manifest",
													MarkdownDescription: "Inlined manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"location_type": schema.StringAttribute{
													Description:         "Type of Kubernetes-like location",
													MarkdownDescription: "Type of Kubernetes-like location",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Uri", "Inlined"),
													},
												},

												"uri": schema.StringAttribute{
													Description:         "Location in a file fetched from a uri.",
													MarkdownDescription: "Location in a file fetched from a uri.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"plugin": schema.SingleNestedAttribute{
											Description:         "Allows importing a plugin.  Plugins are mainly imported devfiles that contribute components, commands and events as a consistent single unit. They are defined in either YAML files following the devfile syntax, or as 'DevWorkspaceTemplate' Kubernetes Custom Resources",
											MarkdownDescription: "Allows importing a plugin.  Plugins are mainly imported devfiles that contribute components, commands and events as a consistent single unit. They are defined in either YAML files following the devfile syntax, or as 'DevWorkspaceTemplate' Kubernetes Custom Resources",
											Attributes: map[string]schema.Attribute{
												"commands": schema.ListNestedAttribute{
													Description:         "Overrides of commands encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",
													MarkdownDescription: "Overrides of commands encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"apply": schema.SingleNestedAttribute{
																Description:         "Command that consists in applying a given component definition, typically bound to a devworkspace event.  For example, when an 'apply' command is bound to a 'preStart' event, and references a 'container' component, it will start the container as a K8S initContainer in the devworkspace POD, unless the component has its 'dedicatedPod' field set to 'true'.  When no 'apply' command exist for a given component, it is assumed the component will be applied at devworkspace start by default, unless 'deployByDefault' for that component is set to false.",
																MarkdownDescription: "Command that consists in applying a given component definition, typically bound to a devworkspace event.  For example, when an 'apply' command is bound to a 'preStart' event, and references a 'container' component, it will start the container as a K8S initContainer in the devworkspace POD, unless the component has its 'dedicatedPod' field set to 'true'.  When no 'apply' command exist for a given component, it is assumed the component will be applied at devworkspace start by default, unless 'deployByDefault' for that component is set to false.",
																Attributes: map[string]schema.Attribute{
																	"component": schema.StringAttribute{
																		Description:         "Describes component that will be applied",
																		MarkdownDescription: "Describes component that will be applied",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"group": schema.SingleNestedAttribute{
																		Description:         "Defines the group this command is part of",
																		MarkdownDescription: "Defines the group this command is part of",
																		Attributes: map[string]schema.Attribute{
																			"is_default": schema.BoolAttribute{
																				Description:         "Identifies the default command for a given group kind",
																				MarkdownDescription: "Identifies the default command for a given group kind",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"kind": schema.StringAttribute{
																				Description:         "Kind of group the command is part of",
																				MarkdownDescription: "Kind of group the command is part of",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.OneOf("build", "run", "test", "debug", "deploy"),
																				},
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"label": schema.StringAttribute{
																		Description:         "Optional label that provides a label for this command to be used in Editor UI menus for example",
																		MarkdownDescription: "Optional label that provides a label for this command to be used in Editor UI menus for example",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"attributes": schema.MapAttribute{
																Description:         "Map of implementation-dependant free-form YAML attributes.",
																MarkdownDescription: "Map of implementation-dependant free-form YAML attributes.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"command_type": schema.StringAttribute{
																Description:         "Type of devworkspace command",
																MarkdownDescription: "Type of devworkspace command",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("Exec", "Apply", "Composite"),
																},
															},

															"composite": schema.SingleNestedAttribute{
																Description:         "Composite command that allows executing several sub-commands either sequentially or concurrently",
																MarkdownDescription: "Composite command that allows executing several sub-commands either sequentially or concurrently",
																Attributes: map[string]schema.Attribute{
																	"commands": schema.ListAttribute{
																		Description:         "The commands that comprise this composite command",
																		MarkdownDescription: "The commands that comprise this composite command",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"group": schema.SingleNestedAttribute{
																		Description:         "Defines the group this command is part of",
																		MarkdownDescription: "Defines the group this command is part of",
																		Attributes: map[string]schema.Attribute{
																			"is_default": schema.BoolAttribute{
																				Description:         "Identifies the default command for a given group kind",
																				MarkdownDescription: "Identifies the default command for a given group kind",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"kind": schema.StringAttribute{
																				Description:         "Kind of group the command is part of",
																				MarkdownDescription: "Kind of group the command is part of",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.OneOf("build", "run", "test", "debug", "deploy"),
																				},
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"label": schema.StringAttribute{
																		Description:         "Optional label that provides a label for this command to be used in Editor UI menus for example",
																		MarkdownDescription: "Optional label that provides a label for this command to be used in Editor UI menus for example",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"parallel": schema.BoolAttribute{
																		Description:         "Indicates if the sub-commands should be executed concurrently",
																		MarkdownDescription: "Indicates if the sub-commands should be executed concurrently",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"exec": schema.SingleNestedAttribute{
																Description:         "CLI Command executed in an existing component container",
																MarkdownDescription: "CLI Command executed in an existing component container",
																Attributes: map[string]schema.Attribute{
																	"command_line": schema.StringAttribute{
																		Description:         "The actual command-line string  Special variables that can be used:   - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping.   - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
																		MarkdownDescription: "The actual command-line string  Special variables that can be used:   - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping.   - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"component": schema.StringAttribute{
																		Description:         "Describes component to which given action relates",
																		MarkdownDescription: "Describes component to which given action relates",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"env": schema.ListNestedAttribute{
																		Description:         "Optional list of environment variables that have to be set before running the command",
																		MarkdownDescription: "Optional list of environment variables that have to be set before running the command",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
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

																	"group": schema.SingleNestedAttribute{
																		Description:         "Defines the group this command is part of",
																		MarkdownDescription: "Defines the group this command is part of",
																		Attributes: map[string]schema.Attribute{
																			"is_default": schema.BoolAttribute{
																				Description:         "Identifies the default command for a given group kind",
																				MarkdownDescription: "Identifies the default command for a given group kind",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"kind": schema.StringAttribute{
																				Description:         "Kind of group the command is part of",
																				MarkdownDescription: "Kind of group the command is part of",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.OneOf("build", "run", "test", "debug", "deploy"),
																				},
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"hot_reload_capable": schema.BoolAttribute{
																		Description:         "Specify whether the command is restarted or not when the source code changes. If set to 'true' the command won't be restarted. A *hotReloadCapable* 'run' or 'debug' command is expected to handle file changes on its own and won't be restarted. A *hotReloadCapable* 'build' command is expected to be executed only once and won't be executed again. This field is taken into account only for commands 'build', 'run' and 'debug' with 'isDefault' set to 'true'.  Default value is 'false'",
																		MarkdownDescription: "Specify whether the command is restarted or not when the source code changes. If set to 'true' the command won't be restarted. A *hotReloadCapable* 'run' or 'debug' command is expected to handle file changes on its own and won't be restarted. A *hotReloadCapable* 'build' command is expected to be executed only once and won't be executed again. This field is taken into account only for commands 'build', 'run' and 'debug' with 'isDefault' set to 'true'.  Default value is 'false'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"label": schema.StringAttribute{
																		Description:         "Optional label that provides a label for this command to be used in Editor UI menus for example",
																		MarkdownDescription: "Optional label that provides a label for this command to be used in Editor UI menus for example",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"working_dir": schema.StringAttribute{
																		Description:         "Working directory where the command should be executed  Special variables that can be used:   - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping.   - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
																		MarkdownDescription: "Working directory where the command should be executed  Special variables that can be used:   - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping.   - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"id": schema.StringAttribute{
																Description:         "Mandatory identifier that allows referencing this command in composite commands, from a parent, or in events.",
																MarkdownDescription: "Mandatory identifier that allows referencing this command in composite commands, from a parent, or in events.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"components": schema.ListNestedAttribute{
													Description:         "Overrides of components encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",
													MarkdownDescription: "Overrides of components encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"attributes": schema.MapAttribute{
																Description:         "Map of implementation-dependant free-form YAML attributes.",
																MarkdownDescription: "Map of implementation-dependant free-form YAML attributes.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"component_type": schema.StringAttribute{
																Description:         "Type of component",
																MarkdownDescription: "Type of component",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("Container", "Kubernetes", "Openshift", "Volume", "Image"),
																},
															},

															"container": schema.SingleNestedAttribute{
																Description:         "Allows adding and configuring devworkspace-related containers",
																MarkdownDescription: "Allows adding and configuring devworkspace-related containers",
																Attributes: map[string]schema.Attribute{
																	"annotation": schema.SingleNestedAttribute{
																		Description:         "Annotations that should be added to specific resources for this container",
																		MarkdownDescription: "Annotations that should be added to specific resources for this container",
																		Attributes: map[string]schema.Attribute{
																			"deployment": schema.MapAttribute{
																				Description:         "Annotations to be added to deployment",
																				MarkdownDescription: "Annotations to be added to deployment",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"service": schema.MapAttribute{
																				Description:         "Annotations to be added to service",
																				MarkdownDescription: "Annotations to be added to service",
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

																	"args": schema.ListAttribute{
																		Description:         "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command.  Defaults to an empty array, meaning use whatever is defined in the image.",
																		MarkdownDescription: "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command.  Defaults to an empty array, meaning use whatever is defined in the image.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"command": schema.ListAttribute{
																		Description:         "The command to run in the dockerimage component instead of the default one provided in the image.  Defaults to an empty array, meaning use whatever is defined in the image.",
																		MarkdownDescription: "The command to run in the dockerimage component instead of the default one provided in the image.  Defaults to an empty array, meaning use whatever is defined in the image.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"cpu_limit": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"cpu_request": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"dedicated_pod": schema.BoolAttribute{
																		Description:         "Specify if a container should run in its own separated pod, instead of running as part of the main development environment pod.  Default value is 'false'",
																		MarkdownDescription: "Specify if a container should run in its own separated pod, instead of running as part of the main development environment pod.  Default value is 'false'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"endpoints": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"annotation": schema.MapAttribute{
																					Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
																					MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"attributes": schema.MapAttribute{
																					Description:         "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
																					MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"exposure": schema.StringAttribute{
																					Description:         "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
																					MarkdownDescription: "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.OneOf("public", "internal", "none"),
																					},
																				},

																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.LengthAtMost(63),
																						stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																					},
																				},

																				"path": schema.StringAttribute{
																					Description:         "Path of the endpoint URL",
																					MarkdownDescription: "Path of the endpoint URL",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"protocol": schema.StringAttribute{
																					Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
																					MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																					},
																				},

																				"secure": schema.BoolAttribute{
																					Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																					MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"target_port": schema.Int64Attribute{
																					Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
																					MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",
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

																	"env": schema.ListNestedAttribute{
																		Description:         "Environment variables used in this container.  The following variables are reserved and cannot be overridden via env:   - '$PROJECTS_ROOT'   - '$PROJECT_SOURCE'",
																		MarkdownDescription: "Environment variables used in this container.  The following variables are reserved and cannot be overridden via env:   - '$PROJECTS_ROOT'   - '$PROJECT_SOURCE'",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
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

																	"image": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"memory_limit": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"memory_request": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"mount_sources": schema.BoolAttribute{
																		Description:         "Toggles whether or not the project source code should be mounted in the component.  Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",
																		MarkdownDescription: "Toggles whether or not the project source code should be mounted in the component.  Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"source_mapping": schema.StringAttribute{
																		Description:         "Optional specification of the path in the container where project sources should be transferred/mounted when 'mountSources' is 'true'. When omitted, the default value of /projects is used.",
																		MarkdownDescription: "Optional specification of the path in the container where project sources should be transferred/mounted when 'mountSources' is 'true'. When omitted, the default value of /projects is used.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"volume_mounts": schema.ListNestedAttribute{
																		Description:         "List of volumes mounts that should be mounted is this container.",
																		MarkdownDescription: "List of volumes mounts that should be mounted is this container.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "The volume mount name is the name of an existing 'Volume' component. If several containers mount the same volume name then they will reuse the same volume and will be able to access to the same files.",
																					MarkdownDescription: "The volume mount name is the name of an existing 'Volume' component. If several containers mount the same volume name then they will reuse the same volume and will be able to access to the same files.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.LengthAtMost(63),
																						stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																					},
																				},

																				"path": schema.StringAttribute{
																					Description:         "The path in the component container where the volume should be mounted. If not path is mentioned, default path is the is '/<name>'.",
																					MarkdownDescription: "The path in the component container where the volume should be mounted. If not path is mentioned, default path is the is '/<name>'.",
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

															"image": schema.SingleNestedAttribute{
																Description:         "Allows specifying the definition of an image for outer loop builds",
																MarkdownDescription: "Allows specifying the definition of an image for outer loop builds",
																Attributes: map[string]schema.Attribute{
																	"auto_build": schema.BoolAttribute{
																		Description:         "Defines if the image should be built during startup.  Default value is 'false'",
																		MarkdownDescription: "Defines if the image should be built during startup.  Default value is 'false'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"dockerfile": schema.SingleNestedAttribute{
																		Description:         "Allows specifying dockerfile type build",
																		MarkdownDescription: "Allows specifying dockerfile type build",
																		Attributes: map[string]schema.Attribute{
																			"args": schema.ListAttribute{
																				Description:         "The arguments to supply to the dockerfile build.",
																				MarkdownDescription: "The arguments to supply to the dockerfile build.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"build_context": schema.StringAttribute{
																				Description:         "Path of source directory to establish build context. Defaults to ${PROJECT_SOURCE} in the container",
																				MarkdownDescription: "Path of source directory to establish build context. Defaults to ${PROJECT_SOURCE} in the container",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"devfile_registry": schema.SingleNestedAttribute{
																				Description:         "Dockerfile's Devfile Registry source",
																				MarkdownDescription: "Dockerfile's Devfile Registry source",
																				Attributes: map[string]schema.Attribute{
																					"id": schema.StringAttribute{
																						Description:         "Id in a devfile registry that contains a Dockerfile. The src in the OCI registry required for the Dockerfile build will be downloaded for building the image.",
																						MarkdownDescription: "Id in a devfile registry that contains a Dockerfile. The src in the OCI registry required for the Dockerfile build will be downloaded for building the image.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"registry_url": schema.StringAttribute{
																						Description:         "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src. To ensure the Dockerfile gets resolved consistently in different environments, it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",
																						MarkdownDescription: "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src. To ensure the Dockerfile gets resolved consistently in different environments, it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"git": schema.SingleNestedAttribute{
																				Description:         "Dockerfile's Git source",
																				MarkdownDescription: "Dockerfile's Git source",
																				Attributes: map[string]schema.Attribute{
																					"checkout_from": schema.SingleNestedAttribute{
																						Description:         "Defines from what the project should be checked out. Required if there are more than one remote configured",
																						MarkdownDescription: "Defines from what the project should be checked out. Required if there are more than one remote configured",
																						Attributes: map[string]schema.Attribute{
																							"remote": schema.StringAttribute{
																								Description:         "The remote name should be used as init. Required if there are more than one remote configured",
																								MarkdownDescription: "The remote name should be used as init. Required if there are more than one remote configured",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"revision": schema.StringAttribute{
																								Description:         "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
																								MarkdownDescription: "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},
																						},
																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"file_location": schema.StringAttribute{
																						Description:         "Location of the Dockerfile in the Git repository when using git as Dockerfile src. Defaults to Dockerfile.",
																						MarkdownDescription: "Location of the Dockerfile in the Git repository when using git as Dockerfile src. Defaults to Dockerfile.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"remotes": schema.MapAttribute{
																						Description:         "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
																						MarkdownDescription: "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
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

																			"root_required": schema.BoolAttribute{
																				Description:         "Specify if a privileged builder pod is required.  Default value is 'false'",
																				MarkdownDescription: "Specify if a privileged builder pod is required.  Default value is 'false'",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"src_type": schema.StringAttribute{
																				Description:         "Type of Dockerfile src",
																				MarkdownDescription: "Type of Dockerfile src",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.OneOf("Uri", "DevfileRegistry", "Git"),
																				},
																			},

																			"uri": schema.StringAttribute{
																				Description:         "URI Reference of a Dockerfile. It can be a full URL or a relative URI from the current devfile as the base URI.",
																				MarkdownDescription: "URI Reference of a Dockerfile. It can be a full URL or a relative URI from the current devfile as the base URI.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"image_name": schema.StringAttribute{
																		Description:         "Name of the image for the resulting outerloop build",
																		MarkdownDescription: "Name of the image for the resulting outerloop build",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"image_type": schema.StringAttribute{
																		Description:         "Type of image",
																		MarkdownDescription: "Type of image",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("Dockerfile", "AutoBuild"),
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes": schema.SingleNestedAttribute{
																Description:         "Allows importing into the devworkspace the Kubernetes resources defined in a given manifest. For example this allows reusing the Kubernetes definitions used to deploy some runtime components in production.",
																MarkdownDescription: "Allows importing into the devworkspace the Kubernetes resources defined in a given manifest. For example this allows reusing the Kubernetes definitions used to deploy some runtime components in production.",
																Attributes: map[string]schema.Attribute{
																	"deploy_by_default": schema.BoolAttribute{
																		Description:         "Defines if the component should be deployed during startup.  Default value is 'false'",
																		MarkdownDescription: "Defines if the component should be deployed during startup.  Default value is 'false'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"endpoints": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"annotation": schema.MapAttribute{
																					Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
																					MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"attributes": schema.MapAttribute{
																					Description:         "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
																					MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"exposure": schema.StringAttribute{
																					Description:         "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
																					MarkdownDescription: "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.OneOf("public", "internal", "none"),
																					},
																				},

																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.LengthAtMost(63),
																						stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																					},
																				},

																				"path": schema.StringAttribute{
																					Description:         "Path of the endpoint URL",
																					MarkdownDescription: "Path of the endpoint URL",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"protocol": schema.StringAttribute{
																					Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
																					MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																					},
																				},

																				"secure": schema.BoolAttribute{
																					Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																					MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"target_port": schema.Int64Attribute{
																					Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
																					MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",
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

																	"inlined": schema.StringAttribute{
																		Description:         "Inlined manifest",
																		MarkdownDescription: "Inlined manifest",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"location_type": schema.StringAttribute{
																		Description:         "Type of Kubernetes-like location",
																		MarkdownDescription: "Type of Kubernetes-like location",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("Uri", "Inlined"),
																		},
																	},

																	"uri": schema.StringAttribute{
																		Description:         "Location in a file fetched from a uri.",
																		MarkdownDescription: "Location in a file fetched from a uri.",
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
																Description:         "Mandatory name that allows referencing the component from other elements (such as commands) or from an external devfile that may reference this component through a parent or a plugin.",
																MarkdownDescription: "Mandatory name that allows referencing the component from other elements (such as commands) or from an external devfile that may reference this component through a parent or a plugin.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																},
															},

															"openshift": schema.SingleNestedAttribute{
																Description:         "Allows importing into the devworkspace the OpenShift resources defined in a given manifest. For example this allows reusing the OpenShift definitions used to deploy some runtime components in production.",
																MarkdownDescription: "Allows importing into the devworkspace the OpenShift resources defined in a given manifest. For example this allows reusing the OpenShift definitions used to deploy some runtime components in production.",
																Attributes: map[string]schema.Attribute{
																	"deploy_by_default": schema.BoolAttribute{
																		Description:         "Defines if the component should be deployed during startup.  Default value is 'false'",
																		MarkdownDescription: "Defines if the component should be deployed during startup.  Default value is 'false'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"endpoints": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"annotation": schema.MapAttribute{
																					Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
																					MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"attributes": schema.MapAttribute{
																					Description:         "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
																					MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"exposure": schema.StringAttribute{
																					Description:         "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
																					MarkdownDescription: "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.OneOf("public", "internal", "none"),
																					},
																				},

																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.LengthAtMost(63),
																						stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																					},
																				},

																				"path": schema.StringAttribute{
																					Description:         "Path of the endpoint URL",
																					MarkdownDescription: "Path of the endpoint URL",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"protocol": schema.StringAttribute{
																					Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
																					MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																					},
																				},

																				"secure": schema.BoolAttribute{
																					Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																					MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"target_port": schema.Int64Attribute{
																					Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
																					MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",
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

																	"inlined": schema.StringAttribute{
																		Description:         "Inlined manifest",
																		MarkdownDescription: "Inlined manifest",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"location_type": schema.StringAttribute{
																		Description:         "Type of Kubernetes-like location",
																		MarkdownDescription: "Type of Kubernetes-like location",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("Uri", "Inlined"),
																		},
																	},

																	"uri": schema.StringAttribute{
																		Description:         "Location in a file fetched from a uri.",
																		MarkdownDescription: "Location in a file fetched from a uri.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"volume": schema.SingleNestedAttribute{
																Description:         "Allows specifying the definition of a volume shared by several other components",
																MarkdownDescription: "Allows specifying the definition of a volume shared by several other components",
																Attributes: map[string]schema.Attribute{
																	"ephemeral": schema.BoolAttribute{
																		Description:         "Ephemeral volumes are not stored persistently across restarts. Defaults to false",
																		MarkdownDescription: "Ephemeral volumes are not stored persistently across restarts. Defaults to false",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"size": schema.StringAttribute{
																		Description:         "Size of the volume",
																		MarkdownDescription: "Size of the volume",
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

												"id": schema.StringAttribute{
													Description:         "Id in a registry that contains a Devfile yaml file",
													MarkdownDescription: "Id in a registry that contains a Devfile yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"import_reference_type": schema.StringAttribute{
													Description:         "type of location from where the referenced template structure should be retrieved",
													MarkdownDescription: "type of location from where the referenced template structure should be retrieved",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Uri", "Id", "Kubernetes"),
													},
												},

												"kubernetes": schema.SingleNestedAttribute{
													Description:         "Reference to a Kubernetes CRD of type DevWorkspaceTemplate",
													MarkdownDescription: "Reference to a Kubernetes CRD of type DevWorkspaceTemplate",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
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

												"registry_url": schema.StringAttribute{
													Description:         "Registry URL to pull the parent devfile from when using id in the parent reference. To ensure the parent devfile gets resolved consistently in different environments, it is recommended to always specify the 'registryUrl' when 'id' is used.",
													MarkdownDescription: "Registry URL to pull the parent devfile from when using id in the parent reference. To ensure the parent devfile gets resolved consistently in different environments, it is recommended to always specify the 'registryUrl' when 'id' is used.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "URI Reference of a parent devfile YAML file. It can be a full URL or a relative URI with the current devfile as the base URI.",
													MarkdownDescription: "URI Reference of a parent devfile YAML file. It can be a full URL or a relative URI with the current devfile as the base URI.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"version": schema.StringAttribute{
													Description:         "Specific stack/sample version to pull the parent devfile from, when using id in the parent reference. To specify 'version', 'id' must be defined and used as the import reference source. 'version' can be either a specific stack version, or 'latest'. If no 'version' specified, default version will be used.",
													MarkdownDescription: "Specific stack/sample version to pull the parent devfile from, when using id in the parent reference. To specify 'version', 'id' must be defined and used as the import reference source. 'version' can be either a specific stack version, or 'latest'. If no 'version' specified, default version will be used.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^(latest)|(([1-9])\.([0-9]+)\.([0-9]+)(\-[0-9a-z-]+(\.[0-9a-z-]+)*)?(\+[0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*)?)$`), ""),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"volume": schema.SingleNestedAttribute{
											Description:         "Allows specifying the definition of a volume shared by several other components",
											MarkdownDescription: "Allows specifying the definition of a volume shared by several other components",
											Attributes: map[string]schema.Attribute{
												"ephemeral": schema.BoolAttribute{
													Description:         "Ephemeral volumes are not stored persistently across restarts. Defaults to false",
													MarkdownDescription: "Ephemeral volumes are not stored persistently across restarts. Defaults to false",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"size": schema.StringAttribute{
													Description:         "Size of the volume",
													MarkdownDescription: "Size of the volume",
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

							"workspace_default_editor": schema.StringAttribute{
								Description:         "The default editor to workspace create with. It could be a plugin ID or a URI. The plugin ID must have 'publisher/plugin/version'. The URI must start from 'http'.",
								MarkdownDescription: "The default editor to workspace create with. It could be a plugin ID or a URI. The plugin ID must have 'publisher/plugin/version'. The URI must start from 'http'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"workspace_namespace_default": schema.StringAttribute{
								Description:         "Defines Kubernetes default namespace in which user's workspaces are created for a case when a user does not override it. It's possible to use '<username>', '<userid>' and '<workspaceid>' placeholders, such as che-workspace-<username>. In that case, a new namespace will be created for each user or workspace.",
								MarkdownDescription: "Defines Kubernetes default namespace in which user's workspaces are created for a case when a user does not override it. It's possible to use '<username>', '<userid>' and '<workspaceid>' placeholders, such as che-workspace-<username>. In that case, a new namespace will be created for each user or workspace.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"workspace_pod_node_selector": schema.MapAttribute{
								Description:         "The node selector that limits the nodes that can run the workspace pods.",
								MarkdownDescription: "The node selector that limits the nodes that can run the workspace pods.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"workspace_pod_tolerations": schema.ListNestedAttribute{
								Description:         "The pod tolerations put on the workspace pods to limit where the workspace pods can run.",
								MarkdownDescription: "The pod tolerations put on the workspace pods to limit where the workspace pods can run.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"effect": schema.StringAttribute{
											Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key": schema.StringAttribute{
											Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"toleration_seconds": schema.Int64Attribute{
											Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
											MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

							"workspaces_default_plugins": schema.ListNestedAttribute{
								Description:         "Default plug-ins applied to Devworkspaces.",
								MarkdownDescription: "Default plug-ins applied to Devworkspaces.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"editor": schema.StringAttribute{
											Description:         "The editor id to specify default plug-ins for.",
											MarkdownDescription: "The editor id to specify default plug-ins for.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"plugins": schema.ListAttribute{
											Description:         "Default plug-in uris for the specified editor.",
											MarkdownDescription: "Default plug-in uris for the specified editor.",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"storage": schema.SingleNestedAttribute{
						Description:         "Configuration settings related to the persistent storage used by the Che installation.",
						MarkdownDescription: "Configuration settings related to the persistent storage used by the Che installation.",
						Attributes: map[string]schema.Attribute{
							"per_workspace_strategy_pvc_storage_class_name": schema.StringAttribute{
								Description:         "Storage class for the Persistent Volume Claims dedicated to the Che workspaces. When omitted or left blank, a default storage class is used.",
								MarkdownDescription: "Storage class for the Persistent Volume Claims dedicated to the Che workspaces. When omitted or left blank, a default storage class is used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"per_workspace_strategy_pvc_claim_size": schema.StringAttribute{
								Description:         "Size of the persistent volume claim for workspaces.",
								MarkdownDescription: "Size of the persistent volume claim for workspaces.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"postgres_pvc_storage_class_name": schema.StringAttribute{
								Description:         "Storage class for the Persistent Volume Claim dedicated to the PostgreSQL database. When omitted or left blank, a default storage class is used.",
								MarkdownDescription: "Storage class for the Persistent Volume Claim dedicated to the PostgreSQL database. When omitted or left blank, a default storage class is used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pre_create_sub_paths": schema.BoolAttribute{
								Description:         "Instructs the Che server to start a special Pod to pre-create a sub-path in the Persistent Volumes. Defaults to 'false', however it will need to enable it according to the configuration of your Kubernetes cluster.",
								MarkdownDescription: "Instructs the Che server to start a special Pod to pre-create a sub-path in the Persistent Volumes. Defaults to 'false', however it will need to enable it according to the configuration of your Kubernetes cluster.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pvc_claim_size": schema.StringAttribute{
								Description:         "Size of the persistent volume claim for workspaces. Defaults to '10Gi'.",
								MarkdownDescription: "Size of the persistent volume claim for workspaces. Defaults to '10Gi'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pvc_jobs_image": schema.StringAttribute{
								Description:         "Overrides the container image used to create sub-paths in the Persistent Volumes. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator. See also the 'preCreateSubPaths' field.",
								MarkdownDescription: "Overrides the container image used to create sub-paths in the Persistent Volumes. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator. See also the 'preCreateSubPaths' field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pvc_strategy": schema.StringAttribute{
								Description:         "Persistent volume claim strategy for the Che server. This Can be:'common' (all workspaces PVCs in one volume), 'per-workspace' (one PVC per workspace for all declared volumes) and 'unique' (one PVC per declared volume). Defaults to 'common'.",
								MarkdownDescription: "Persistent volume claim strategy for the Che server. This Can be:'common' (all workspaces PVCs in one volume), 'per-workspace' (one PVC per workspace for all declared volumes) and 'unique' (one PVC per declared volume). Defaults to 'common'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"workspace_pvc_storage_class_name": schema.StringAttribute{
								Description:         "Storage class for the Persistent Volume Claims dedicated to the Che workspaces. When omitted or left blank, a default storage class is used.",
								MarkdownDescription: "Storage class for the Persistent Volume Claims dedicated to the Che workspaces. When omitted or left blank, a default storage class is used.",
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

func (r *OrgEclipseCheCheClusterV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_org_eclipse_che_che_cluster_v1_manifest")

	var model OrgEclipseCheCheClusterV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("org.eclipse.che/v1")
	model.Kind = pointer.String("CheCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
