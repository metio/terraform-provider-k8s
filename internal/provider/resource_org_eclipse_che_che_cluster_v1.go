/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type OrgEclipseCheCheClusterV1Resource struct{}

var (
	_ resource.Resource = (*OrgEclipseCheCheClusterV1Resource)(nil)
)

type OrgEclipseCheCheClusterV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type OrgEclipseCheCheClusterV1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Auth *struct {
			Debug *bool `tfsdk:"debug" yaml:"debug,omitempty"`

			ExternalIdentityProvider *bool `tfsdk:"external_identity_provider" yaml:"externalIdentityProvider,omitempty"`

			GatewayAuthenticationSidecarImage *string `tfsdk:"gateway_authentication_sidecar_image" yaml:"gatewayAuthenticationSidecarImage,omitempty"`

			GatewayAuthorizationSidecarImage *string `tfsdk:"gateway_authorization_sidecar_image" yaml:"gatewayAuthorizationSidecarImage,omitempty"`

			GatewayConfigBumpEnv *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`

				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

					FieldRef *struct {
						ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

						FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
					} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

					ResourceFieldRef *struct {
						ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

						Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

						Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
					} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

					SecretKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"gateway_config_bump_env" yaml:"gatewayConfigBumpEnv,omitempty"`

			GatewayEnv *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`

				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

					FieldRef *struct {
						ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

						FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
					} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

					ResourceFieldRef *struct {
						ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

						Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

						Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
					} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

					SecretKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"gateway_env" yaml:"gatewayEnv,omitempty"`

			GatewayHeaderRewriteSidecarImage *string `tfsdk:"gateway_header_rewrite_sidecar_image" yaml:"gatewayHeaderRewriteSidecarImage,omitempty"`

			GatewayKubeRbacProxyEnv *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`

				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

					FieldRef *struct {
						ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

						FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
					} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

					ResourceFieldRef *struct {
						ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

						Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

						Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
					} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

					SecretKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"gateway_kube_rbac_proxy_env" yaml:"gatewayKubeRbacProxyEnv,omitempty"`

			GatewayOAuthProxyEnv *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`

				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

					FieldRef *struct {
						ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

						FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
					} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

					ResourceFieldRef *struct {
						ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

						Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

						Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
					} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

					SecretKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"gateway_o_auth_proxy_env" yaml:"gatewayOAuthProxyEnv,omitempty"`

			IdentityProviderAdminUserName *string `tfsdk:"identity_provider_admin_user_name" yaml:"identityProviderAdminUserName,omitempty"`

			IdentityProviderClientId *string `tfsdk:"identity_provider_client_id" yaml:"identityProviderClientId,omitempty"`

			IdentityProviderContainerResources *struct {
				Limits *struct {
					Cpu *string `tfsdk:"cpu" yaml:"cpu,omitempty"`

					Memory *string `tfsdk:"memory" yaml:"memory,omitempty"`
				} `tfsdk:"limits" yaml:"limits,omitempty"`

				Request *struct {
					Cpu *string `tfsdk:"cpu" yaml:"cpu,omitempty"`

					Memory *string `tfsdk:"memory" yaml:"memory,omitempty"`
				} `tfsdk:"request" yaml:"request,omitempty"`
			} `tfsdk:"identity_provider_container_resources" yaml:"identityProviderContainerResources,omitempty"`

			IdentityProviderImage *string `tfsdk:"identity_provider_image" yaml:"identityProviderImage,omitempty"`

			IdentityProviderImagePullPolicy *string `tfsdk:"identity_provider_image_pull_policy" yaml:"identityProviderImagePullPolicy,omitempty"`

			IdentityProviderIngress *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Labels *string `tfsdk:"labels" yaml:"labels,omitempty"`
			} `tfsdk:"identity_provider_ingress" yaml:"identityProviderIngress,omitempty"`

			IdentityProviderPassword *string `tfsdk:"identity_provider_password" yaml:"identityProviderPassword,omitempty"`

			IdentityProviderPostgresPassword *string `tfsdk:"identity_provider_postgres_password" yaml:"identityProviderPostgresPassword,omitempty"`

			IdentityProviderPostgresSecret *string `tfsdk:"identity_provider_postgres_secret" yaml:"identityProviderPostgresSecret,omitempty"`

			IdentityProviderRealm *string `tfsdk:"identity_provider_realm" yaml:"identityProviderRealm,omitempty"`

			IdentityProviderRoute *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Domain *string `tfsdk:"domain" yaml:"domain,omitempty"`

				Labels *string `tfsdk:"labels" yaml:"labels,omitempty"`
			} `tfsdk:"identity_provider_route" yaml:"identityProviderRoute,omitempty"`

			IdentityProviderSecret *string `tfsdk:"identity_provider_secret" yaml:"identityProviderSecret,omitempty"`

			IdentityProviderURL *string `tfsdk:"identity_provider_url" yaml:"identityProviderURL,omitempty"`

			IdentityToken *string `tfsdk:"identity_token" yaml:"identityToken,omitempty"`

			InitialOpenShiftOAuthUser *bool `tfsdk:"initial_open_shift_o_auth_user" yaml:"initialOpenShiftOAuthUser,omitempty"`

			NativeUserMode *bool `tfsdk:"native_user_mode" yaml:"nativeUserMode,omitempty"`

			OAuthClientName *string `tfsdk:"o_auth_client_name" yaml:"oAuthClientName,omitempty"`

			OAuthScope *string `tfsdk:"o_auth_scope" yaml:"oAuthScope,omitempty"`

			OAuthSecret *string `tfsdk:"o_auth_secret" yaml:"oAuthSecret,omitempty"`

			OpenShiftoAuth *bool `tfsdk:"open_shifto_auth" yaml:"openShiftoAuth,omitempty"`

			UpdateAdminPassword *bool `tfsdk:"update_admin_password" yaml:"updateAdminPassword,omitempty"`
		} `tfsdk:"auth" yaml:"auth,omitempty"`

		Dashboard *struct {
			Warning *string `tfsdk:"warning" yaml:"warning,omitempty"`
		} `tfsdk:"dashboard" yaml:"dashboard,omitempty"`

		Database *struct {
			ChePostgresContainerResources *struct {
				Limits *struct {
					Cpu *string `tfsdk:"cpu" yaml:"cpu,omitempty"`

					Memory *string `tfsdk:"memory" yaml:"memory,omitempty"`
				} `tfsdk:"limits" yaml:"limits,omitempty"`

				Request *struct {
					Cpu *string `tfsdk:"cpu" yaml:"cpu,omitempty"`

					Memory *string `tfsdk:"memory" yaml:"memory,omitempty"`
				} `tfsdk:"request" yaml:"request,omitempty"`
			} `tfsdk:"che_postgres_container_resources" yaml:"chePostgresContainerResources,omitempty"`

			ChePostgresDb *string `tfsdk:"che_postgres_db" yaml:"chePostgresDb,omitempty"`

			ChePostgresHostName *string `tfsdk:"che_postgres_host_name" yaml:"chePostgresHostName,omitempty"`

			ChePostgresPassword *string `tfsdk:"che_postgres_password" yaml:"chePostgresPassword,omitempty"`

			ChePostgresPort *string `tfsdk:"che_postgres_port" yaml:"chePostgresPort,omitempty"`

			ChePostgresSecret *string `tfsdk:"che_postgres_secret" yaml:"chePostgresSecret,omitempty"`

			ChePostgresUser *string `tfsdk:"che_postgres_user" yaml:"chePostgresUser,omitempty"`

			ExternalDb *bool `tfsdk:"external_db" yaml:"externalDb,omitempty"`

			PostgresEnv *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`

				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

					FieldRef *struct {
						ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

						FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
					} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

					ResourceFieldRef *struct {
						ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

						Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

						Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
					} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

					SecretKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"postgres_env" yaml:"postgresEnv,omitempty"`

			PostgresImage *string `tfsdk:"postgres_image" yaml:"postgresImage,omitempty"`

			PostgresImagePullPolicy *string `tfsdk:"postgres_image_pull_policy" yaml:"postgresImagePullPolicy,omitempty"`

			PostgresVersion *string `tfsdk:"postgres_version" yaml:"postgresVersion,omitempty"`

			PvcClaimSize *string `tfsdk:"pvc_claim_size" yaml:"pvcClaimSize,omitempty"`
		} `tfsdk:"database" yaml:"database,omitempty"`

		DevWorkspace *struct {
			ControllerImage *string `tfsdk:"controller_image" yaml:"controllerImage,omitempty"`

			Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

			Env *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`

				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

					FieldRef *struct {
						ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

						FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
					} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

					ResourceFieldRef *struct {
						ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

						Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

						Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
					} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

					SecretKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"env" yaml:"env,omitempty"`

			RunningLimit *string `tfsdk:"running_limit" yaml:"runningLimit,omitempty"`

			SecondsOfInactivityBeforeIdling *int64 `tfsdk:"seconds_of_inactivity_before_idling" yaml:"secondsOfInactivityBeforeIdling,omitempty"`

			SecondsOfRunBeforeIdling *int64 `tfsdk:"seconds_of_run_before_idling" yaml:"secondsOfRunBeforeIdling,omitempty"`
		} `tfsdk:"dev_workspace" yaml:"devWorkspace,omitempty"`

		GitServices *struct {
			Bitbucket *[]struct {
				Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

				SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
			} `tfsdk:"bitbucket" yaml:"bitbucket,omitempty"`

			Github *[]struct {
				Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

				SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
			} `tfsdk:"github" yaml:"github,omitempty"`

			Gitlab *[]struct {
				Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

				SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
			} `tfsdk:"gitlab" yaml:"gitlab,omitempty"`
		} `tfsdk:"git_services" yaml:"gitServices,omitempty"`

		ImagePuller *struct {
			Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

			Spec *struct {
				Affinity *string `tfsdk:"affinity" yaml:"affinity,omitempty"`

				CachingCPULimit *string `tfsdk:"caching_cpu_limit" yaml:"cachingCPULimit,omitempty"`

				CachingCPURequest *string `tfsdk:"caching_cpu_request" yaml:"cachingCPURequest,omitempty"`

				CachingIntervalHours *string `tfsdk:"caching_interval_hours" yaml:"cachingIntervalHours,omitempty"`

				CachingMemoryLimit *string `tfsdk:"caching_memory_limit" yaml:"cachingMemoryLimit,omitempty"`

				CachingMemoryRequest *string `tfsdk:"caching_memory_request" yaml:"cachingMemoryRequest,omitempty"`

				ConfigMapName *string `tfsdk:"config_map_name" yaml:"configMapName,omitempty"`

				DaemonsetName *string `tfsdk:"daemonset_name" yaml:"daemonsetName,omitempty"`

				DeploymentName *string `tfsdk:"deployment_name" yaml:"deploymentName,omitempty"`

				ImagePullSecrets *string `tfsdk:"image_pull_secrets" yaml:"imagePullSecrets,omitempty"`

				ImagePullerImage *string `tfsdk:"image_puller_image" yaml:"imagePullerImage,omitempty"`

				Images *string `tfsdk:"images" yaml:"images,omitempty"`

				NodeSelector *string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`
			} `tfsdk:"spec" yaml:"spec,omitempty"`
		} `tfsdk:"image_puller" yaml:"imagePuller,omitempty"`

		K8s *struct {
			IngressClass *string `tfsdk:"ingress_class" yaml:"ingressClass,omitempty"`

			IngressDomain *string `tfsdk:"ingress_domain" yaml:"ingressDomain,omitempty"`

			IngressStrategy *string `tfsdk:"ingress_strategy" yaml:"ingressStrategy,omitempty"`

			SecurityContextFsGroup *string `tfsdk:"security_context_fs_group" yaml:"securityContextFsGroup,omitempty"`

			SecurityContextRunAsUser *string `tfsdk:"security_context_run_as_user" yaml:"securityContextRunAsUser,omitempty"`

			SingleHostExposureType *string `tfsdk:"single_host_exposure_type" yaml:"singleHostExposureType,omitempty"`

			TlsSecretName *string `tfsdk:"tls_secret_name" yaml:"tlsSecretName,omitempty"`
		} `tfsdk:"k8s" yaml:"k8s,omitempty"`

		Metrics *struct {
			Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`
		} `tfsdk:"metrics" yaml:"metrics,omitempty"`

		Server *struct {
			AirGapContainerRegistryHostname *string `tfsdk:"air_gap_container_registry_hostname" yaml:"airGapContainerRegistryHostname,omitempty"`

			AirGapContainerRegistryOrganization *string `tfsdk:"air_gap_container_registry_organization" yaml:"airGapContainerRegistryOrganization,omitempty"`

			AllowAutoProvisionUserNamespace *bool `tfsdk:"allow_auto_provision_user_namespace" yaml:"allowAutoProvisionUserNamespace,omitempty"`

			AllowUserDefinedWorkspaceNamespaces *bool `tfsdk:"allow_user_defined_workspace_namespaces" yaml:"allowUserDefinedWorkspaceNamespaces,omitempty"`

			CheClusterRoles *string `tfsdk:"che_cluster_roles" yaml:"cheClusterRoles,omitempty"`

			CheDebug *string `tfsdk:"che_debug" yaml:"cheDebug,omitempty"`

			CheFlavor *string `tfsdk:"che_flavor" yaml:"cheFlavor,omitempty"`

			CheHost *string `tfsdk:"che_host" yaml:"cheHost,omitempty"`

			CheHostTLSSecret *string `tfsdk:"che_host_tls_secret" yaml:"cheHostTLSSecret,omitempty"`

			CheImage *string `tfsdk:"che_image" yaml:"cheImage,omitempty"`

			CheImagePullPolicy *string `tfsdk:"che_image_pull_policy" yaml:"cheImagePullPolicy,omitempty"`

			CheImageTag *string `tfsdk:"che_image_tag" yaml:"cheImageTag,omitempty"`

			CheLogLevel *string `tfsdk:"che_log_level" yaml:"cheLogLevel,omitempty"`

			CheServerEnv *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`

				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

					FieldRef *struct {
						ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

						FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
					} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

					ResourceFieldRef *struct {
						ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

						Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

						Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
					} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

					SecretKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"che_server_env" yaml:"cheServerEnv,omitempty"`

			CheServerIngress *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Labels *string `tfsdk:"labels" yaml:"labels,omitempty"`
			} `tfsdk:"che_server_ingress" yaml:"cheServerIngress,omitempty"`

			CheServerRoute *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Domain *string `tfsdk:"domain" yaml:"domain,omitempty"`

				Labels *string `tfsdk:"labels" yaml:"labels,omitempty"`
			} `tfsdk:"che_server_route" yaml:"cheServerRoute,omitempty"`

			CheWorkspaceClusterRole *string `tfsdk:"che_workspace_cluster_role" yaml:"cheWorkspaceClusterRole,omitempty"`

			CustomCheProperties *map[string]string `tfsdk:"custom_che_properties" yaml:"customCheProperties,omitempty"`

			DashboardCpuLimit *string `tfsdk:"dashboard_cpu_limit" yaml:"dashboardCpuLimit,omitempty"`

			DashboardCpuRequest *string `tfsdk:"dashboard_cpu_request" yaml:"dashboardCpuRequest,omitempty"`

			DashboardEnv *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`

				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

					FieldRef *struct {
						ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

						FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
					} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

					ResourceFieldRef *struct {
						ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

						Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

						Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
					} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

					SecretKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"dashboard_env" yaml:"dashboardEnv,omitempty"`

			DashboardImage *string `tfsdk:"dashboard_image" yaml:"dashboardImage,omitempty"`

			DashboardImagePullPolicy *string `tfsdk:"dashboard_image_pull_policy" yaml:"dashboardImagePullPolicy,omitempty"`

			DashboardIngress *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Labels *string `tfsdk:"labels" yaml:"labels,omitempty"`
			} `tfsdk:"dashboard_ingress" yaml:"dashboardIngress,omitempty"`

			DashboardMemoryLimit *string `tfsdk:"dashboard_memory_limit" yaml:"dashboardMemoryLimit,omitempty"`

			DashboardMemoryRequest *string `tfsdk:"dashboard_memory_request" yaml:"dashboardMemoryRequest,omitempty"`

			DashboardRoute *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Domain *string `tfsdk:"domain" yaml:"domain,omitempty"`

				Labels *string `tfsdk:"labels" yaml:"labels,omitempty"`
			} `tfsdk:"dashboard_route" yaml:"dashboardRoute,omitempty"`

			DevfileRegistryCpuLimit *string `tfsdk:"devfile_registry_cpu_limit" yaml:"devfileRegistryCpuLimit,omitempty"`

			DevfileRegistryCpuRequest *string `tfsdk:"devfile_registry_cpu_request" yaml:"devfileRegistryCpuRequest,omitempty"`

			DevfileRegistryEnv *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`

				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

					FieldRef *struct {
						ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

						FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
					} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

					ResourceFieldRef *struct {
						ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

						Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

						Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
					} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

					SecretKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"devfile_registry_env" yaml:"devfileRegistryEnv,omitempty"`

			DevfileRegistryImage *string `tfsdk:"devfile_registry_image" yaml:"devfileRegistryImage,omitempty"`

			DevfileRegistryIngress *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Labels *string `tfsdk:"labels" yaml:"labels,omitempty"`
			} `tfsdk:"devfile_registry_ingress" yaml:"devfileRegistryIngress,omitempty"`

			DevfileRegistryMemoryLimit *string `tfsdk:"devfile_registry_memory_limit" yaml:"devfileRegistryMemoryLimit,omitempty"`

			DevfileRegistryMemoryRequest *string `tfsdk:"devfile_registry_memory_request" yaml:"devfileRegistryMemoryRequest,omitempty"`

			DevfileRegistryPullPolicy *string `tfsdk:"devfile_registry_pull_policy" yaml:"devfileRegistryPullPolicy,omitempty"`

			DevfileRegistryRoute *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Domain *string `tfsdk:"domain" yaml:"domain,omitempty"`

				Labels *string `tfsdk:"labels" yaml:"labels,omitempty"`
			} `tfsdk:"devfile_registry_route" yaml:"devfileRegistryRoute,omitempty"`

			DevfileRegistryUrl *string `tfsdk:"devfile_registry_url" yaml:"devfileRegistryUrl,omitempty"`

			DisableInternalClusterSVCNames *bool `tfsdk:"disable_internal_cluster_svc_names" yaml:"disableInternalClusterSVCNames,omitempty"`

			ExternalDevfileRegistries *[]struct {
				Url *string `tfsdk:"url" yaml:"url,omitempty"`
			} `tfsdk:"external_devfile_registries" yaml:"externalDevfileRegistries,omitempty"`

			ExternalDevfileRegistry *bool `tfsdk:"external_devfile_registry" yaml:"externalDevfileRegistry,omitempty"`

			ExternalPluginRegistry *bool `tfsdk:"external_plugin_registry" yaml:"externalPluginRegistry,omitempty"`

			GitSelfSignedCert *bool `tfsdk:"git_self_signed_cert" yaml:"gitSelfSignedCert,omitempty"`

			NonProxyHosts *string `tfsdk:"non_proxy_hosts" yaml:"nonProxyHosts,omitempty"`

			OpenVSXRegistryURL *string `tfsdk:"open_vsx_registry_url" yaml:"openVSXRegistryURL,omitempty"`

			PluginRegistryCpuLimit *string `tfsdk:"plugin_registry_cpu_limit" yaml:"pluginRegistryCpuLimit,omitempty"`

			PluginRegistryCpuRequest *string `tfsdk:"plugin_registry_cpu_request" yaml:"pluginRegistryCpuRequest,omitempty"`

			PluginRegistryEnv *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`

				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

					FieldRef *struct {
						ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

						FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
					} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

					ResourceFieldRef *struct {
						ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

						Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

						Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
					} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

					SecretKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"plugin_registry_env" yaml:"pluginRegistryEnv,omitempty"`

			PluginRegistryImage *string `tfsdk:"plugin_registry_image" yaml:"pluginRegistryImage,omitempty"`

			PluginRegistryIngress *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Labels *string `tfsdk:"labels" yaml:"labels,omitempty"`
			} `tfsdk:"plugin_registry_ingress" yaml:"pluginRegistryIngress,omitempty"`

			PluginRegistryMemoryLimit *string `tfsdk:"plugin_registry_memory_limit" yaml:"pluginRegistryMemoryLimit,omitempty"`

			PluginRegistryMemoryRequest *string `tfsdk:"plugin_registry_memory_request" yaml:"pluginRegistryMemoryRequest,omitempty"`

			PluginRegistryPullPolicy *string `tfsdk:"plugin_registry_pull_policy" yaml:"pluginRegistryPullPolicy,omitempty"`

			PluginRegistryRoute *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Domain *string `tfsdk:"domain" yaml:"domain,omitempty"`

				Labels *string `tfsdk:"labels" yaml:"labels,omitempty"`
			} `tfsdk:"plugin_registry_route" yaml:"pluginRegistryRoute,omitempty"`

			PluginRegistryUrl *string `tfsdk:"plugin_registry_url" yaml:"pluginRegistryUrl,omitempty"`

			ProxyPassword *string `tfsdk:"proxy_password" yaml:"proxyPassword,omitempty"`

			ProxyPort *string `tfsdk:"proxy_port" yaml:"proxyPort,omitempty"`

			ProxySecret *string `tfsdk:"proxy_secret" yaml:"proxySecret,omitempty"`

			ProxyURL *string `tfsdk:"proxy_url" yaml:"proxyURL,omitempty"`

			ProxyUser *string `tfsdk:"proxy_user" yaml:"proxyUser,omitempty"`

			SelfSignedCert *bool `tfsdk:"self_signed_cert" yaml:"selfSignedCert,omitempty"`

			ServerCpuLimit *string `tfsdk:"server_cpu_limit" yaml:"serverCpuLimit,omitempty"`

			ServerCpuRequest *string `tfsdk:"server_cpu_request" yaml:"serverCpuRequest,omitempty"`

			ServerExposureStrategy *string `tfsdk:"server_exposure_strategy" yaml:"serverExposureStrategy,omitempty"`

			ServerMemoryLimit *string `tfsdk:"server_memory_limit" yaml:"serverMemoryLimit,omitempty"`

			ServerMemoryRequest *string `tfsdk:"server_memory_request" yaml:"serverMemoryRequest,omitempty"`

			ServerTrustStoreConfigMapName *string `tfsdk:"server_trust_store_config_map_name" yaml:"serverTrustStoreConfigMapName,omitempty"`

			SingleHostGatewayConfigMapLabels *map[string]string `tfsdk:"single_host_gateway_config_map_labels" yaml:"singleHostGatewayConfigMapLabels,omitempty"`

			SingleHostGatewayConfigSidecarImage *string `tfsdk:"single_host_gateway_config_sidecar_image" yaml:"singleHostGatewayConfigSidecarImage,omitempty"`

			SingleHostGatewayImage *string `tfsdk:"single_host_gateway_image" yaml:"singleHostGatewayImage,omitempty"`

			TlsSupport *bool `tfsdk:"tls_support" yaml:"tlsSupport,omitempty"`

			UseInternalClusterSVCNames *bool `tfsdk:"use_internal_cluster_svc_names" yaml:"useInternalClusterSVCNames,omitempty"`

			WorkspaceDefaultComponents *[]struct {
				Attributes utilities.Dynamic `tfsdk:"attributes" yaml:"attributes,omitempty"`

				ComponentType *string `tfsdk:"component_type" yaml:"componentType,omitempty"`

				Container *struct {
					Annotation *struct {
						Deployment *map[string]string `tfsdk:"deployment" yaml:"deployment,omitempty"`

						Service *map[string]string `tfsdk:"service" yaml:"service,omitempty"`
					} `tfsdk:"annotation" yaml:"annotation,omitempty"`

					Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

					Command *[]string `tfsdk:"command" yaml:"command,omitempty"`

					CpuLimit *string `tfsdk:"cpu_limit" yaml:"cpuLimit,omitempty"`

					CpuRequest *string `tfsdk:"cpu_request" yaml:"cpuRequest,omitempty"`

					DedicatedPod *bool `tfsdk:"dedicated_pod" yaml:"dedicatedPod,omitempty"`

					Endpoints *[]struct {
						Annotation *map[string]string `tfsdk:"annotation" yaml:"annotation,omitempty"`

						Attributes utilities.Dynamic `tfsdk:"attributes" yaml:"attributes,omitempty"`

						Exposure *string `tfsdk:"exposure" yaml:"exposure,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

						Secure *bool `tfsdk:"secure" yaml:"secure,omitempty"`

						TargetPort *int64 `tfsdk:"target_port" yaml:"targetPort,omitempty"`
					} `tfsdk:"endpoints" yaml:"endpoints,omitempty"`

					Env *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"env" yaml:"env,omitempty"`

					Image *string `tfsdk:"image" yaml:"image,omitempty"`

					MemoryLimit *string `tfsdk:"memory_limit" yaml:"memoryLimit,omitempty"`

					MemoryRequest *string `tfsdk:"memory_request" yaml:"memoryRequest,omitempty"`

					MountSources *bool `tfsdk:"mount_sources" yaml:"mountSources,omitempty"`

					SourceMapping *string `tfsdk:"source_mapping" yaml:"sourceMapping,omitempty"`

					VolumeMounts *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`
					} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`
				} `tfsdk:"container" yaml:"container,omitempty"`

				Custom *struct {
					ComponentClass *string `tfsdk:"component_class" yaml:"componentClass,omitempty"`

					EmbeddedResource utilities.Dynamic `tfsdk:"embedded_resource" yaml:"embeddedResource,omitempty"`
				} `tfsdk:"custom" yaml:"custom,omitempty"`

				Image *struct {
					AutoBuild *bool `tfsdk:"auto_build" yaml:"autoBuild,omitempty"`

					Dockerfile *struct {
						Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

						BuildContext *string `tfsdk:"build_context" yaml:"buildContext,omitempty"`

						DevfileRegistry *struct {
							Id *string `tfsdk:"id" yaml:"id,omitempty"`

							RegistryUrl *string `tfsdk:"registry_url" yaml:"registryUrl,omitempty"`
						} `tfsdk:"devfile_registry" yaml:"devfileRegistry,omitempty"`

						Git *struct {
							CheckoutFrom *struct {
								Remote *string `tfsdk:"remote" yaml:"remote,omitempty"`

								Revision *string `tfsdk:"revision" yaml:"revision,omitempty"`
							} `tfsdk:"checkout_from" yaml:"checkoutFrom,omitempty"`

							FileLocation *string `tfsdk:"file_location" yaml:"fileLocation,omitempty"`

							Remotes *map[string]string `tfsdk:"remotes" yaml:"remotes,omitempty"`
						} `tfsdk:"git" yaml:"git,omitempty"`

						RootRequired *bool `tfsdk:"root_required" yaml:"rootRequired,omitempty"`

						SrcType *string `tfsdk:"src_type" yaml:"srcType,omitempty"`

						Uri *string `tfsdk:"uri" yaml:"uri,omitempty"`
					} `tfsdk:"dockerfile" yaml:"dockerfile,omitempty"`

					ImageName *string `tfsdk:"image_name" yaml:"imageName,omitempty"`

					ImageType *string `tfsdk:"image_type" yaml:"imageType,omitempty"`
				} `tfsdk:"image" yaml:"image,omitempty"`

				Kubernetes *struct {
					DeployByDefault *bool `tfsdk:"deploy_by_default" yaml:"deployByDefault,omitempty"`

					Endpoints *[]struct {
						Annotation *map[string]string `tfsdk:"annotation" yaml:"annotation,omitempty"`

						Attributes utilities.Dynamic `tfsdk:"attributes" yaml:"attributes,omitempty"`

						Exposure *string `tfsdk:"exposure" yaml:"exposure,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

						Secure *bool `tfsdk:"secure" yaml:"secure,omitempty"`

						TargetPort *int64 `tfsdk:"target_port" yaml:"targetPort,omitempty"`
					} `tfsdk:"endpoints" yaml:"endpoints,omitempty"`

					Inlined *string `tfsdk:"inlined" yaml:"inlined,omitempty"`

					LocationType *string `tfsdk:"location_type" yaml:"locationType,omitempty"`

					Uri *string `tfsdk:"uri" yaml:"uri,omitempty"`
				} `tfsdk:"kubernetes" yaml:"kubernetes,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Openshift *struct {
					DeployByDefault *bool `tfsdk:"deploy_by_default" yaml:"deployByDefault,omitempty"`

					Endpoints *[]struct {
						Annotation *map[string]string `tfsdk:"annotation" yaml:"annotation,omitempty"`

						Attributes utilities.Dynamic `tfsdk:"attributes" yaml:"attributes,omitempty"`

						Exposure *string `tfsdk:"exposure" yaml:"exposure,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

						Secure *bool `tfsdk:"secure" yaml:"secure,omitempty"`

						TargetPort *int64 `tfsdk:"target_port" yaml:"targetPort,omitempty"`
					} `tfsdk:"endpoints" yaml:"endpoints,omitempty"`

					Inlined *string `tfsdk:"inlined" yaml:"inlined,omitempty"`

					LocationType *string `tfsdk:"location_type" yaml:"locationType,omitempty"`

					Uri *string `tfsdk:"uri" yaml:"uri,omitempty"`
				} `tfsdk:"openshift" yaml:"openshift,omitempty"`

				Plugin *struct {
					Commands *[]struct {
						Apply *struct {
							Component *string `tfsdk:"component" yaml:"component,omitempty"`

							Group *struct {
								IsDefault *bool `tfsdk:"is_default" yaml:"isDefault,omitempty"`

								Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`
							} `tfsdk:"group" yaml:"group,omitempty"`

							Label *string `tfsdk:"label" yaml:"label,omitempty"`
						} `tfsdk:"apply" yaml:"apply,omitempty"`

						Attributes utilities.Dynamic `tfsdk:"attributes" yaml:"attributes,omitempty"`

						CommandType *string `tfsdk:"command_type" yaml:"commandType,omitempty"`

						Composite *struct {
							Commands *[]string `tfsdk:"commands" yaml:"commands,omitempty"`

							Group *struct {
								IsDefault *bool `tfsdk:"is_default" yaml:"isDefault,omitempty"`

								Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`
							} `tfsdk:"group" yaml:"group,omitempty"`

							Label *string `tfsdk:"label" yaml:"label,omitempty"`

							Parallel *bool `tfsdk:"parallel" yaml:"parallel,omitempty"`
						} `tfsdk:"composite" yaml:"composite,omitempty"`

						Exec *struct {
							CommandLine *string `tfsdk:"command_line" yaml:"commandLine,omitempty"`

							Component *string `tfsdk:"component" yaml:"component,omitempty"`

							Env *[]struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"env" yaml:"env,omitempty"`

							Group *struct {
								IsDefault *bool `tfsdk:"is_default" yaml:"isDefault,omitempty"`

								Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`
							} `tfsdk:"group" yaml:"group,omitempty"`

							HotReloadCapable *bool `tfsdk:"hot_reload_capable" yaml:"hotReloadCapable,omitempty"`

							Label *string `tfsdk:"label" yaml:"label,omitempty"`

							WorkingDir *string `tfsdk:"working_dir" yaml:"workingDir,omitempty"`
						} `tfsdk:"exec" yaml:"exec,omitempty"`

						Id *string `tfsdk:"id" yaml:"id,omitempty"`
					} `tfsdk:"commands" yaml:"commands,omitempty"`

					Components *[]struct {
						Attributes utilities.Dynamic `tfsdk:"attributes" yaml:"attributes,omitempty"`

						ComponentType *string `tfsdk:"component_type" yaml:"componentType,omitempty"`

						Container *struct {
							Annotation *struct {
								Deployment *map[string]string `tfsdk:"deployment" yaml:"deployment,omitempty"`

								Service *map[string]string `tfsdk:"service" yaml:"service,omitempty"`
							} `tfsdk:"annotation" yaml:"annotation,omitempty"`

							Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

							Command *[]string `tfsdk:"command" yaml:"command,omitempty"`

							CpuLimit *string `tfsdk:"cpu_limit" yaml:"cpuLimit,omitempty"`

							CpuRequest *string `tfsdk:"cpu_request" yaml:"cpuRequest,omitempty"`

							DedicatedPod *bool `tfsdk:"dedicated_pod" yaml:"dedicatedPod,omitempty"`

							Endpoints *[]struct {
								Annotation *map[string]string `tfsdk:"annotation" yaml:"annotation,omitempty"`

								Attributes utilities.Dynamic `tfsdk:"attributes" yaml:"attributes,omitempty"`

								Exposure *string `tfsdk:"exposure" yaml:"exposure,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`

								Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

								Secure *bool `tfsdk:"secure" yaml:"secure,omitempty"`

								TargetPort *int64 `tfsdk:"target_port" yaml:"targetPort,omitempty"`
							} `tfsdk:"endpoints" yaml:"endpoints,omitempty"`

							Env *[]struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"env" yaml:"env,omitempty"`

							Image *string `tfsdk:"image" yaml:"image,omitempty"`

							MemoryLimit *string `tfsdk:"memory_limit" yaml:"memoryLimit,omitempty"`

							MemoryRequest *string `tfsdk:"memory_request" yaml:"memoryRequest,omitempty"`

							MountSources *bool `tfsdk:"mount_sources" yaml:"mountSources,omitempty"`

							SourceMapping *string `tfsdk:"source_mapping" yaml:"sourceMapping,omitempty"`

							VolumeMounts *[]struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`
							} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`
						} `tfsdk:"container" yaml:"container,omitempty"`

						Image *struct {
							AutoBuild *bool `tfsdk:"auto_build" yaml:"autoBuild,omitempty"`

							Dockerfile *struct {
								Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

								BuildContext *string `tfsdk:"build_context" yaml:"buildContext,omitempty"`

								DevfileRegistry *struct {
									Id *string `tfsdk:"id" yaml:"id,omitempty"`

									RegistryUrl *string `tfsdk:"registry_url" yaml:"registryUrl,omitempty"`
								} `tfsdk:"devfile_registry" yaml:"devfileRegistry,omitempty"`

								Git *struct {
									CheckoutFrom *struct {
										Remote *string `tfsdk:"remote" yaml:"remote,omitempty"`

										Revision *string `tfsdk:"revision" yaml:"revision,omitempty"`
									} `tfsdk:"checkout_from" yaml:"checkoutFrom,omitempty"`

									FileLocation *string `tfsdk:"file_location" yaml:"fileLocation,omitempty"`

									Remotes *map[string]string `tfsdk:"remotes" yaml:"remotes,omitempty"`
								} `tfsdk:"git" yaml:"git,omitempty"`

								RootRequired *bool `tfsdk:"root_required" yaml:"rootRequired,omitempty"`

								SrcType *string `tfsdk:"src_type" yaml:"srcType,omitempty"`

								Uri *string `tfsdk:"uri" yaml:"uri,omitempty"`
							} `tfsdk:"dockerfile" yaml:"dockerfile,omitempty"`

							ImageName *string `tfsdk:"image_name" yaml:"imageName,omitempty"`

							ImageType *string `tfsdk:"image_type" yaml:"imageType,omitempty"`
						} `tfsdk:"image" yaml:"image,omitempty"`

						Kubernetes *struct {
							DeployByDefault *bool `tfsdk:"deploy_by_default" yaml:"deployByDefault,omitempty"`

							Endpoints *[]struct {
								Annotation *map[string]string `tfsdk:"annotation" yaml:"annotation,omitempty"`

								Attributes utilities.Dynamic `tfsdk:"attributes" yaml:"attributes,omitempty"`

								Exposure *string `tfsdk:"exposure" yaml:"exposure,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`

								Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

								Secure *bool `tfsdk:"secure" yaml:"secure,omitempty"`

								TargetPort *int64 `tfsdk:"target_port" yaml:"targetPort,omitempty"`
							} `tfsdk:"endpoints" yaml:"endpoints,omitempty"`

							Inlined *string `tfsdk:"inlined" yaml:"inlined,omitempty"`

							LocationType *string `tfsdk:"location_type" yaml:"locationType,omitempty"`

							Uri *string `tfsdk:"uri" yaml:"uri,omitempty"`
						} `tfsdk:"kubernetes" yaml:"kubernetes,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Openshift *struct {
							DeployByDefault *bool `tfsdk:"deploy_by_default" yaml:"deployByDefault,omitempty"`

							Endpoints *[]struct {
								Annotation *map[string]string `tfsdk:"annotation" yaml:"annotation,omitempty"`

								Attributes utilities.Dynamic `tfsdk:"attributes" yaml:"attributes,omitempty"`

								Exposure *string `tfsdk:"exposure" yaml:"exposure,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`

								Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

								Secure *bool `tfsdk:"secure" yaml:"secure,omitempty"`

								TargetPort *int64 `tfsdk:"target_port" yaml:"targetPort,omitempty"`
							} `tfsdk:"endpoints" yaml:"endpoints,omitempty"`

							Inlined *string `tfsdk:"inlined" yaml:"inlined,omitempty"`

							LocationType *string `tfsdk:"location_type" yaml:"locationType,omitempty"`

							Uri *string `tfsdk:"uri" yaml:"uri,omitempty"`
						} `tfsdk:"openshift" yaml:"openshift,omitempty"`

						Volume *struct {
							Ephemeral *bool `tfsdk:"ephemeral" yaml:"ephemeral,omitempty"`

							Size *string `tfsdk:"size" yaml:"size,omitempty"`
						} `tfsdk:"volume" yaml:"volume,omitempty"`
					} `tfsdk:"components" yaml:"components,omitempty"`

					Id *string `tfsdk:"id" yaml:"id,omitempty"`

					ImportReferenceType *string `tfsdk:"import_reference_type" yaml:"importReferenceType,omitempty"`

					Kubernetes *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"kubernetes" yaml:"kubernetes,omitempty"`

					RegistryUrl *string `tfsdk:"registry_url" yaml:"registryUrl,omitempty"`

					Uri *string `tfsdk:"uri" yaml:"uri,omitempty"`

					Version *string `tfsdk:"version" yaml:"version,omitempty"`
				} `tfsdk:"plugin" yaml:"plugin,omitempty"`

				Volume *struct {
					Ephemeral *bool `tfsdk:"ephemeral" yaml:"ephemeral,omitempty"`

					Size *string `tfsdk:"size" yaml:"size,omitempty"`
				} `tfsdk:"volume" yaml:"volume,omitempty"`
			} `tfsdk:"workspace_default_components" yaml:"workspaceDefaultComponents,omitempty"`

			WorkspaceDefaultEditor *string `tfsdk:"workspace_default_editor" yaml:"workspaceDefaultEditor,omitempty"`

			WorkspaceNamespaceDefault *string `tfsdk:"workspace_namespace_default" yaml:"workspaceNamespaceDefault,omitempty"`

			WorkspacePodNodeSelector *map[string]string `tfsdk:"workspace_pod_node_selector" yaml:"workspacePodNodeSelector,omitempty"`

			WorkspacePodTolerations *[]struct {
				Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"workspace_pod_tolerations" yaml:"workspacePodTolerations,omitempty"`

			WorkspacesDefaultPlugins *[]struct {
				Editor *string `tfsdk:"editor" yaml:"editor,omitempty"`

				Plugins *[]string `tfsdk:"plugins" yaml:"plugins,omitempty"`
			} `tfsdk:"workspaces_default_plugins" yaml:"workspacesDefaultPlugins,omitempty"`
		} `tfsdk:"server" yaml:"server,omitempty"`

		Storage *struct {
			PerWorkspaceStrategyPVCStorageClassName *string `tfsdk:"per_workspace_strategy_pvc_storage_class_name" yaml:"perWorkspaceStrategyPVCStorageClassName,omitempty"`

			PerWorkspaceStrategyPvcClaimSize *string `tfsdk:"per_workspace_strategy_pvc_claim_size" yaml:"perWorkspaceStrategyPvcClaimSize,omitempty"`

			PostgresPVCStorageClassName *string `tfsdk:"postgres_pvc_storage_class_name" yaml:"postgresPVCStorageClassName,omitempty"`

			PreCreateSubPaths *bool `tfsdk:"pre_create_sub_paths" yaml:"preCreateSubPaths,omitempty"`

			PvcClaimSize *string `tfsdk:"pvc_claim_size" yaml:"pvcClaimSize,omitempty"`

			PvcJobsImage *string `tfsdk:"pvc_jobs_image" yaml:"pvcJobsImage,omitempty"`

			PvcStrategy *string `tfsdk:"pvc_strategy" yaml:"pvcStrategy,omitempty"`

			WorkspacePVCStorageClassName *string `tfsdk:"workspace_pvc_storage_class_name" yaml:"workspacePVCStorageClassName,omitempty"`
		} `tfsdk:"storage" yaml:"storage,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewOrgEclipseCheCheClusterV1Resource() resource.Resource {
	return &OrgEclipseCheCheClusterV1Resource{}
}

func (r *OrgEclipseCheCheClusterV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_eclipse_che_che_cluster_v1"
}

func (r *OrgEclipseCheCheClusterV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "The 'CheCluster' custom resource allows defining and managing a Che server installation",
		MarkdownDescription: "The 'CheCluster' custom resource allows defining and managing a Che server installation",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "Desired configuration of the Che installation. Based on these settings, the  Operator automatically creates and maintains several ConfigMaps that will contain the appropriate environment variables the various components of the Che installation. These generated ConfigMaps must NOT be updated manually.",
				MarkdownDescription: "Desired configuration of the Che installation. Based on these settings, the  Operator automatically creates and maintains several ConfigMaps that will contain the appropriate environment variables the various components of the Che installation. These generated ConfigMaps must NOT be updated manually.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"auth": {
						Description:         "Configuration settings related to the Authentication used by the Che installation.",
						MarkdownDescription: "Configuration settings related to the Authentication used by the Che installation.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"debug": {
								Description:         "Deprecated. The value of this flag is ignored. Debug internal identity provider.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Debug internal identity provider.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"external_identity_provider": {
								Description:         "Deprecated. The value of this flag is ignored. Instructs the Operator on whether or not to deploy a dedicated Identity Provider (Keycloak or RH SSO instance). Instructs the Operator on whether to deploy a dedicated Identity Provider (Keycloak or RH-SSO instance). By default, a dedicated Identity Provider server is deployed as part of the Che installation. When 'externalIdentityProvider' is 'true', no dedicated identity provider will be deployed by the Operator and you will need to provide details about the external identity provider you are about to use. See also all the other fields starting with: 'identityProvider'.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Instructs the Operator on whether or not to deploy a dedicated Identity Provider (Keycloak or RH SSO instance). Instructs the Operator on whether to deploy a dedicated Identity Provider (Keycloak or RH-SSO instance). By default, a dedicated Identity Provider server is deployed as part of the Che installation. When 'externalIdentityProvider' is 'true', no dedicated identity provider will be deployed by the Operator and you will need to provide details about the external identity provider you are about to use. See also all the other fields starting with: 'identityProvider'.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"gateway_authentication_sidecar_image": {
								Description:         "Gateway sidecar responsible for authentication when NativeUserMode is enabled. See link:https://github.com/oauth2-proxy/oauth2-proxy[oauth2-proxy] or link:https://github.com/openshift/oauth-proxy[openshift/oauth-proxy].",
								MarkdownDescription: "Gateway sidecar responsible for authentication when NativeUserMode is enabled. See link:https://github.com/oauth2-proxy/oauth2-proxy[oauth2-proxy] or link:https://github.com/openshift/oauth-proxy[openshift/oauth-proxy].",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"gateway_authorization_sidecar_image": {
								Description:         "Gateway sidecar responsible for authorization when NativeUserMode is enabled. See link:https://github.com/brancz/kube-rbac-proxy[kube-rbac-proxy] or link:https://github.com/openshift/kube-rbac-proxy[openshift/kube-rbac-proxy]",
								MarkdownDescription: "Gateway sidecar responsible for authorization when NativeUserMode is enabled. See link:https://github.com/brancz/kube-rbac-proxy[kube-rbac-proxy] or link:https://github.com/openshift/kube-rbac-proxy[openshift/kube-rbac-proxy]",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"gateway_config_bump_env": {
								Description:         "List of environment variables to set in the Configbump container.",
								MarkdownDescription: "List of environment variables to set in the Configbump container.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
										MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"value": {
										Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
										MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value_from": {
										Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
										MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_key_ref": {
												Description:         "Selects a key of a ConfigMap.",
												MarkdownDescription: "Selects a key of a ConfigMap.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the ConfigMap or its key must be defined",
														MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_ref": {
												Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
												MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"api_version": {
														Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
														MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_path": {
														Description:         "Path of the field to select in the specified API version.",
														MarkdownDescription: "Path of the field to select in the specified API version.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resource_field_ref": {
												Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
												MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"container_name": {
														Description:         "Container name: required for volumes, optional for env vars",
														MarkdownDescription: "Container name: required for volumes, optional for env vars",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"divisor": {
														Description:         "Specifies the output format of the exposed resources, defaults to '1'",
														MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resource": {
														Description:         "Required: resource to select",
														MarkdownDescription: "Required: resource to select",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_key_ref": {
												Description:         "Selects a key of a secret in the pod's namespace",
												MarkdownDescription: "Selects a key of a secret in the pod's namespace",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the Secret or its key must be defined",
														MarkdownDescription: "Specify whether the Secret or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"gateway_env": {
								Description:         "List of environment variables to set in the Gateway container.",
								MarkdownDescription: "List of environment variables to set in the Gateway container.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
										MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"value": {
										Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
										MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value_from": {
										Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
										MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_key_ref": {
												Description:         "Selects a key of a ConfigMap.",
												MarkdownDescription: "Selects a key of a ConfigMap.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the ConfigMap or its key must be defined",
														MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_ref": {
												Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
												MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"api_version": {
														Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
														MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_path": {
														Description:         "Path of the field to select in the specified API version.",
														MarkdownDescription: "Path of the field to select in the specified API version.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resource_field_ref": {
												Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
												MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"container_name": {
														Description:         "Container name: required for volumes, optional for env vars",
														MarkdownDescription: "Container name: required for volumes, optional for env vars",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"divisor": {
														Description:         "Specifies the output format of the exposed resources, defaults to '1'",
														MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resource": {
														Description:         "Required: resource to select",
														MarkdownDescription: "Required: resource to select",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_key_ref": {
												Description:         "Selects a key of a secret in the pod's namespace",
												MarkdownDescription: "Selects a key of a secret in the pod's namespace",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the Secret or its key must be defined",
														MarkdownDescription: "Specify whether the Secret or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"gateway_header_rewrite_sidecar_image": {
								Description:         "Deprecated. The value of this flag is ignored. Sidecar functionality is now implemented in Traefik plugin.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Sidecar functionality is now implemented in Traefik plugin.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"gateway_kube_rbac_proxy_env": {
								Description:         "List of environment variables to set in the Kube rbac proxy container.",
								MarkdownDescription: "List of environment variables to set in the Kube rbac proxy container.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
										MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"value": {
										Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
										MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value_from": {
										Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
										MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_key_ref": {
												Description:         "Selects a key of a ConfigMap.",
												MarkdownDescription: "Selects a key of a ConfigMap.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the ConfigMap or its key must be defined",
														MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_ref": {
												Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
												MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"api_version": {
														Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
														MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_path": {
														Description:         "Path of the field to select in the specified API version.",
														MarkdownDescription: "Path of the field to select in the specified API version.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resource_field_ref": {
												Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
												MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"container_name": {
														Description:         "Container name: required for volumes, optional for env vars",
														MarkdownDescription: "Container name: required for volumes, optional for env vars",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"divisor": {
														Description:         "Specifies the output format of the exposed resources, defaults to '1'",
														MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resource": {
														Description:         "Required: resource to select",
														MarkdownDescription: "Required: resource to select",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_key_ref": {
												Description:         "Selects a key of a secret in the pod's namespace",
												MarkdownDescription: "Selects a key of a secret in the pod's namespace",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the Secret or its key must be defined",
														MarkdownDescription: "Specify whether the Secret or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"gateway_o_auth_proxy_env": {
								Description:         "List of environment variables to set in the OAuth proxy container.",
								MarkdownDescription: "List of environment variables to set in the OAuth proxy container.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
										MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"value": {
										Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
										MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value_from": {
										Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
										MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_key_ref": {
												Description:         "Selects a key of a ConfigMap.",
												MarkdownDescription: "Selects a key of a ConfigMap.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the ConfigMap or its key must be defined",
														MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_ref": {
												Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
												MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"api_version": {
														Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
														MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_path": {
														Description:         "Path of the field to select in the specified API version.",
														MarkdownDescription: "Path of the field to select in the specified API version.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resource_field_ref": {
												Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
												MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"container_name": {
														Description:         "Container name: required for volumes, optional for env vars",
														MarkdownDescription: "Container name: required for volumes, optional for env vars",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"divisor": {
														Description:         "Specifies the output format of the exposed resources, defaults to '1'",
														MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resource": {
														Description:         "Required: resource to select",
														MarkdownDescription: "Required: resource to select",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_key_ref": {
												Description:         "Selects a key of a secret in the pod's namespace",
												MarkdownDescription: "Selects a key of a secret in the pod's namespace",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the Secret or its key must be defined",
														MarkdownDescription: "Specify whether the Secret or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"identity_provider_admin_user_name": {
								Description:         "Deprecated. The value of this flag is ignored. Overrides the name of the Identity Provider administrator user. Defaults to 'admin'.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Overrides the name of the Identity Provider administrator user. Defaults to 'admin'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"identity_provider_client_id": {
								Description:         "Deprecated. The value of this flag is ignored. Name of a Identity provider, Keycloak or RH-SSO, 'client-id' that is used for Che. Override this when an external Identity Provider is in use. See the 'externalIdentityProvider' field. When omitted or left blank, it is set to the value of the 'flavour' field suffixed with '-public'.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Name of a Identity provider, Keycloak or RH-SSO, 'client-id' that is used for Che. Override this when an external Identity Provider is in use. See the 'externalIdentityProvider' field. When omitted or left blank, it is set to the value of the 'flavour' field suffixed with '-public'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"identity_provider_container_resources": {
								Description:         "Deprecated. The value of this flag is ignored. Identity provider container custom settings.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Identity provider container custom settings.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"limits": {
										Description:         "Limits describes the maximum amount of compute resources allowed.",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cpu": {
												Description:         "CPU, in cores. (500m = .5 cores)",
												MarkdownDescription: "CPU, in cores. (500m = .5 cores)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"memory": {
												Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
												MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"request": {
										Description:         "Requests describes the minimum amount of compute resources required.",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cpu": {
												Description:         "CPU, in cores. (500m = .5 cores)",
												MarkdownDescription: "CPU, in cores. (500m = .5 cores)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"memory": {
												Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
												MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"identity_provider_image": {
								Description:         "Deprecated. The value of this flag is ignored. Overrides the container image used in the Identity Provider, Keycloak or RH-SSO, deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Overrides the container image used in the Identity Provider, Keycloak or RH-SSO, deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"identity_provider_image_pull_policy": {
								Description:         "Deprecated. The value of this flag is ignored. Overrides the image pull policy used in the Identity Provider, Keycloak or RH-SSO, deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Overrides the image pull policy used in the Identity Provider, Keycloak or RH-SSO, deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"identity_provider_ingress": {
								Description:         "Deprecated. The value of this flag is ignored. Ingress custom settings.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Ingress custom settings.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										MarkdownDescription: "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"labels": {
										Description:         "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										MarkdownDescription: "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"identity_provider_password": {
								Description:         "Deprecated. The value of this flag is ignored. Overrides the password of Keycloak administrator user. Override this when an external Identity Provider is in use. See the 'externalIdentityProvider' field. When omitted or left blank, it is set to an auto-generated password.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Overrides the password of Keycloak administrator user. Override this when an external Identity Provider is in use. See the 'externalIdentityProvider' field. When omitted or left blank, it is set to an auto-generated password.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"identity_provider_postgres_password": {
								Description:         "Deprecated. The value of this flag is ignored. Password for a Identity Provider, Keycloak or RH-SSO, to connect to the database. Override this when an external Identity Provider is in use. See the 'externalIdentityProvider' field. When omitted or left blank, it is set to an auto-generated password.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Password for a Identity Provider, Keycloak or RH-SSO, to connect to the database. Override this when an external Identity Provider is in use. See the 'externalIdentityProvider' field. When omitted or left blank, it is set to an auto-generated password.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"identity_provider_postgres_secret": {
								Description:         "Deprecated. The value of this flag is ignored. The secret that contains 'password' for the Identity Provider, Keycloak or RH-SSO, to connect to the database. When the secret is defined, the 'identityProviderPostgresPassword' is ignored. When the value is omitted or left blank, the one of following scenarios applies: 1. 'identityProviderPostgresPassword' is defined, then it will be used to connect to the database. 2. 'identityProviderPostgresPassword' is not defined, then a new secret with the name 'che-identity-postgres-secret' will be created with an auto-generated value for 'password'. The secret must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. The secret that contains 'password' for the Identity Provider, Keycloak or RH-SSO, to connect to the database. When the secret is defined, the 'identityProviderPostgresPassword' is ignored. When the value is omitted or left blank, the one of following scenarios applies: 1. 'identityProviderPostgresPassword' is defined, then it will be used to connect to the database. 2. 'identityProviderPostgresPassword' is not defined, then a new secret with the name 'che-identity-postgres-secret' will be created with an auto-generated value for 'password'. The secret must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"identity_provider_realm": {
								Description:         "Deprecated. The value of this flag is ignored. Name of a Identity provider, Keycloak or RH-SSO, realm that is used for Che. Override this when an external Identity Provider is in use. See the 'externalIdentityProvider' field. When omitted or left blank, it is set to the value of the 'flavour' field.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Name of a Identity provider, Keycloak or RH-SSO, realm that is used for Che. Override this when an external Identity Provider is in use. See the 'externalIdentityProvider' field. When omitted or left blank, it is set to the value of the 'flavour' field.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"identity_provider_route": {
								Description:         "Deprecated. The value of this flag is ignored. Route custom settings.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Route custom settings.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										MarkdownDescription: "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"domain": {
										Description:         "Operator uses the domain to generate a hostname for a route. In a conjunction with labels it creates a route, which is served by a non-default Ingress controller. The generated host name will follow this pattern: '<route-name>-<route-namespace>.<domain>'.",
										MarkdownDescription: "Operator uses the domain to generate a hostname for a route. In a conjunction with labels it creates a route, which is served by a non-default Ingress controller. The generated host name will follow this pattern: '<route-name>-<route-namespace>.<domain>'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"labels": {
										Description:         "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										MarkdownDescription: "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"identity_provider_secret": {
								Description:         "Deprecated. The value of this flag is ignored. The secret that contains 'user' and 'password' for Identity Provider. When the secret is defined, the 'identityProviderAdminUserName' and 'identityProviderPassword' are ignored. When the value is omitted or left blank, the one of following scenarios applies: 1. 'identityProviderAdminUserName' and 'identityProviderPassword' are defined, then they will be used. 2. 'identityProviderAdminUserName' or 'identityProviderPassword' are not defined, then a new secret with the name 'che-identity-secret' will be created with default value 'admin' for 'user' and with an auto-generated value for 'password'. The secret must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. The secret that contains 'user' and 'password' for Identity Provider. When the secret is defined, the 'identityProviderAdminUserName' and 'identityProviderPassword' are ignored. When the value is omitted or left blank, the one of following scenarios applies: 1. 'identityProviderAdminUserName' and 'identityProviderPassword' are defined, then they will be used. 2. 'identityProviderAdminUserName' or 'identityProviderPassword' are not defined, then a new secret with the name 'che-identity-secret' will be created with default value 'admin' for 'user' and with an auto-generated value for 'password'. The secret must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"identity_provider_url": {
								Description:         "Public URL of the Identity Provider server (Keycloak / RH-SSO server). Set this ONLY when a use of an external Identity Provider is needed. See the 'externalIdentityProvider' field. By default, this will be automatically calculated and set by the Operator.",
								MarkdownDescription: "Public URL of the Identity Provider server (Keycloak / RH-SSO server). Set this ONLY when a use of an external Identity Provider is needed. See the 'externalIdentityProvider' field. By default, this will be automatically calculated and set by the Operator.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"identity_token": {
								Description:         "Identity token to be passed to upstream. There are two types of tokens supported: 'id_token' and 'access_token'. Default value is 'id_token'. This field is specific to Che installations made for Kubernetes only and ignored for OpenShift.",
								MarkdownDescription: "Identity token to be passed to upstream. There are two types of tokens supported: 'id_token' and 'access_token'. Default value is 'id_token'. This field is specific to Che installations made for Kubernetes only and ignored for OpenShift.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"initial_open_shift_o_auth_user": {
								Description:         "Deprecated. The value of this flag is ignored. For operating with the OpenShift OAuth authentication, create a new user account since the kubeadmin can not be used. If the value is true, then a new OpenShift OAuth user will be created for the HTPasswd identity provider. If the value is false and the user has already been created, then it will be removed. If value is an empty, then do nothing. The user's credentials are stored in the 'openshift-oauth-user-credentials' secret in 'openshift-config' namespace by Operator. Note that this solution is Openshift 4 platform-specific.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. For operating with the OpenShift OAuth authentication, create a new user account since the kubeadmin can not be used. If the value is true, then a new OpenShift OAuth user will be created for the HTPasswd identity provider. If the value is false and the user has already been created, then it will be removed. If value is an empty, then do nothing. The user's credentials are stored in the 'openshift-oauth-user-credentials' secret in 'openshift-config' namespace by Operator. Note that this solution is Openshift 4 platform-specific.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"native_user_mode": {
								Description:         "Deprecated. The value of this flag is ignored. Enables native user mode. Currently works only on OpenShift and DevWorkspace engine. Native User mode uses OpenShift OAuth directly as identity provider, without Keycloak.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Enables native user mode. Currently works only on OpenShift and DevWorkspace engine. Native User mode uses OpenShift OAuth directly as identity provider, without Keycloak.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"o_auth_client_name": {
								Description:         "Name of the OpenShift 'OAuthClient' resource used to setup identity federation on the OpenShift side. Auto-generated when left blank. See also the 'OpenShiftoAuth' field.",
								MarkdownDescription: "Name of the OpenShift 'OAuthClient' resource used to setup identity federation on the OpenShift side. Auto-generated when left blank. See also the 'OpenShiftoAuth' field.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"o_auth_scope": {
								Description:         "Access Token Scope. This field is specific to Che installations made for Kubernetes only and ignored for OpenShift.",
								MarkdownDescription: "Access Token Scope. This field is specific to Che installations made for Kubernetes only and ignored for OpenShift.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"o_auth_secret": {
								Description:         "Name of the secret set in the OpenShift 'OAuthClient' resource used to setup identity federation on the OpenShift side. Auto-generated when left blank. See also the 'OAuthClientName' field.",
								MarkdownDescription: "Name of the secret set in the OpenShift 'OAuthClient' resource used to setup identity federation on the OpenShift side. Auto-generated when left blank. See also the 'OAuthClientName' field.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"open_shifto_auth": {
								Description:         "Deprecated. The value of this flag is ignored. Enables the integration of the identity provider (Keycloak / RHSSO) with OpenShift OAuth. Empty value on OpenShift by default. This will allow users to directly login with their OpenShift user through the OpenShift login, and have their workspaces created under personal OpenShift namespaces. WARNING: the 'kubeadmin' user is NOT supported, and logging through it will NOT allow accessing the Che Dashboard.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Enables the integration of the identity provider (Keycloak / RHSSO) with OpenShift OAuth. Empty value on OpenShift by default. This will allow users to directly login with their OpenShift user through the OpenShift login, and have their workspaces created under personal OpenShift namespaces. WARNING: the 'kubeadmin' user is NOT supported, and logging through it will NOT allow accessing the Che Dashboard.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"update_admin_password": {
								Description:         "Deprecated. The value of this flag is ignored. Forces the default 'admin' Che user to update password on first login. Defaults to 'false'.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Forces the default 'admin' Che user to update password on first login. Defaults to 'false'.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"dashboard": {
						Description:         "Configuration settings related to the User Dashboard used by the Che installation.",
						MarkdownDescription: "Configuration settings related to the User Dashboard used by the Che installation.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"warning": {
								Description:         "Warning message that will be displayed on the User Dashboard",
								MarkdownDescription: "Warning message that will be displayed on the User Dashboard",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"database": {
						Description:         "Configuration settings related to the database used by the Che installation.",
						MarkdownDescription: "Configuration settings related to the database used by the Che installation.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"che_postgres_container_resources": {
								Description:         "PostgreSQL container custom settings",
								MarkdownDescription: "PostgreSQL container custom settings",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"limits": {
										Description:         "Limits describes the maximum amount of compute resources allowed.",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cpu": {
												Description:         "CPU, in cores. (500m = .5 cores)",
												MarkdownDescription: "CPU, in cores. (500m = .5 cores)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"memory": {
												Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
												MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"request": {
										Description:         "Requests describes the minimum amount of compute resources required.",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cpu": {
												Description:         "CPU, in cores. (500m = .5 cores)",
												MarkdownDescription: "CPU, in cores. (500m = .5 cores)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"memory": {
												Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
												MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"che_postgres_db": {
								Description:         "PostgreSQL database name that the Che server uses to connect to the DB. Defaults to 'dbche'.",
								MarkdownDescription: "PostgreSQL database name that the Che server uses to connect to the DB. Defaults to 'dbche'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"che_postgres_host_name": {
								Description:         "PostgreSQL Database host name that the Che server uses to connect to. Defaults is 'postgres'. Override this value ONLY when using an external database. See field 'externalDb'. In the default case it will be automatically set by the Operator.",
								MarkdownDescription: "PostgreSQL Database host name that the Che server uses to connect to. Defaults is 'postgres'. Override this value ONLY when using an external database. See field 'externalDb'. In the default case it will be automatically set by the Operator.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"che_postgres_password": {
								Description:         "PostgreSQL password that the Che server uses to connect to the DB. When omitted or left blank, it will be set to an automatically generated value.",
								MarkdownDescription: "PostgreSQL password that the Che server uses to connect to the DB. When omitted or left blank, it will be set to an automatically generated value.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"che_postgres_port": {
								Description:         "PostgreSQL Database port that the Che server uses to connect to. Defaults to 5432. Override this value ONLY when using an external database. See field 'externalDb'. In the default case it will be automatically set by the Operator.",
								MarkdownDescription: "PostgreSQL Database port that the Che server uses to connect to. Defaults to 5432. Override this value ONLY when using an external database. See field 'externalDb'. In the default case it will be automatically set by the Operator.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"che_postgres_secret": {
								Description:         "The secret that contains PostgreSQL'user' and 'password' that the Che server uses to connect to the DB. When the secret is defined, the 'chePostgresUser' and 'chePostgresPassword' are ignored. When the value is omitted or left blank, the one of following scenarios applies: 1. 'chePostgresUser' and 'chePostgresPassword' are defined, then they will be used to connect to the DB. 2. 'chePostgresUser' or 'chePostgresPassword' are not defined, then a new secret with the name 'postgres-credentials' will be created with default value of 'pgche' for 'user' and with an auto-generated value for 'password'. The secret must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",
								MarkdownDescription: "The secret that contains PostgreSQL'user' and 'password' that the Che server uses to connect to the DB. When the secret is defined, the 'chePostgresUser' and 'chePostgresPassword' are ignored. When the value is omitted or left blank, the one of following scenarios applies: 1. 'chePostgresUser' and 'chePostgresPassword' are defined, then they will be used to connect to the DB. 2. 'chePostgresUser' or 'chePostgresPassword' are not defined, then a new secret with the name 'postgres-credentials' will be created with default value of 'pgche' for 'user' and with an auto-generated value for 'password'. The secret must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"che_postgres_user": {
								Description:         "PostgreSQL user that the Che server uses to connect to the DB. Defaults to 'pgche'.",
								MarkdownDescription: "PostgreSQL user that the Che server uses to connect to the DB. Defaults to 'pgche'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"external_db": {
								Description:         "Instructs the Operator on whether to deploy a dedicated database. By default, a dedicated PostgreSQL database is deployed as part of the Che installation. When 'externalDb' is 'true', no dedicated database will be deployed by the Operator and you will need to provide connection details to the external DB you are about to use. See also all the fields starting with: 'chePostgres'.",
								MarkdownDescription: "Instructs the Operator on whether to deploy a dedicated database. By default, a dedicated PostgreSQL database is deployed as part of the Che installation. When 'externalDb' is 'true', no dedicated database will be deployed by the Operator and you will need to provide connection details to the external DB you are about to use. See also all the fields starting with: 'chePostgres'.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"postgres_env": {
								Description:         "List of environment variables to set in the PostgreSQL container.",
								MarkdownDescription: "List of environment variables to set in the PostgreSQL container.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
										MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"value": {
										Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
										MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value_from": {
										Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
										MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_key_ref": {
												Description:         "Selects a key of a ConfigMap.",
												MarkdownDescription: "Selects a key of a ConfigMap.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the ConfigMap or its key must be defined",
														MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_ref": {
												Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
												MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"api_version": {
														Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
														MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_path": {
														Description:         "Path of the field to select in the specified API version.",
														MarkdownDescription: "Path of the field to select in the specified API version.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resource_field_ref": {
												Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
												MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"container_name": {
														Description:         "Container name: required for volumes, optional for env vars",
														MarkdownDescription: "Container name: required for volumes, optional for env vars",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"divisor": {
														Description:         "Specifies the output format of the exposed resources, defaults to '1'",
														MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resource": {
														Description:         "Required: resource to select",
														MarkdownDescription: "Required: resource to select",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_key_ref": {
												Description:         "Selects a key of a secret in the pod's namespace",
												MarkdownDescription: "Selects a key of a secret in the pod's namespace",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the Secret or its key must be defined",
														MarkdownDescription: "Specify whether the Secret or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"postgres_image": {
								Description:         "Overrides the container image used in the PostgreSQL database deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",
								MarkdownDescription: "Overrides the container image used in the PostgreSQL database deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"postgres_image_pull_policy": {
								Description:         "Overrides the image pull policy used in the PostgreSQL database deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
								MarkdownDescription: "Overrides the image pull policy used in the PostgreSQL database deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"postgres_version": {
								Description:         "Indicates a PostgreSQL version image to use. Allowed values are: '9.6' and '13.3'. Migrate your PostgreSQL database to switch from one version to another.",
								MarkdownDescription: "Indicates a PostgreSQL version image to use. Allowed values are: '9.6' and '13.3'. Migrate your PostgreSQL database to switch from one version to another.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pvc_claim_size": {
								Description:         "Size of the persistent volume claim for database. Defaults to '1Gi'. To update pvc storageclass that provisions it must support resize when Eclipse Che has been already deployed.",
								MarkdownDescription: "Size of the persistent volume claim for database. Defaults to '1Gi'. To update pvc storageclass that provisions it must support resize when Eclipse Che has been already deployed.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"dev_workspace": {
						Description:         "DevWorkspace operator configuration",
						MarkdownDescription: "DevWorkspace operator configuration",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"controller_image": {
								Description:         "Overrides the container image used in the DevWorkspace controller deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",
								MarkdownDescription: "Overrides the container image used in the DevWorkspace controller deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable": {
								Description:         "Deploys the DevWorkspace Operator in the cluster. Does nothing when a matching version of the Operator is already installed. Fails when a non-matching version of the Operator is already installed.",
								MarkdownDescription: "Deploys the DevWorkspace Operator in the cluster. Does nothing when a matching version of the Operator is already installed. Fails when a non-matching version of the Operator is already installed.",

								Type: types.BoolType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"env": {
								Description:         "List of environment variables to set in the DevWorkspace container.",
								MarkdownDescription: "List of environment variables to set in the DevWorkspace container.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
										MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"value": {
										Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
										MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value_from": {
										Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
										MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_key_ref": {
												Description:         "Selects a key of a ConfigMap.",
												MarkdownDescription: "Selects a key of a ConfigMap.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the ConfigMap or its key must be defined",
														MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_ref": {
												Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
												MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"api_version": {
														Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
														MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_path": {
														Description:         "Path of the field to select in the specified API version.",
														MarkdownDescription: "Path of the field to select in the specified API version.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resource_field_ref": {
												Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
												MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"container_name": {
														Description:         "Container name: required for volumes, optional for env vars",
														MarkdownDescription: "Container name: required for volumes, optional for env vars",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"divisor": {
														Description:         "Specifies the output format of the exposed resources, defaults to '1'",
														MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resource": {
														Description:         "Required: resource to select",
														MarkdownDescription: "Required: resource to select",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_key_ref": {
												Description:         "Selects a key of a secret in the pod's namespace",
												MarkdownDescription: "Selects a key of a secret in the pod's namespace",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the Secret or its key must be defined",
														MarkdownDescription: "Specify whether the Secret or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"running_limit": {
								Description:         "Maximum number of the running workspaces per user.",
								MarkdownDescription: "Maximum number of the running workspaces per user.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"seconds_of_inactivity_before_idling": {
								Description:         "Idle timeout for workspaces in seconds. This timeout is the duration after which a workspace will be idled if there is no activity. To disable workspace idling due to inactivity, set this value to -1.",
								MarkdownDescription: "Idle timeout for workspaces in seconds. This timeout is the duration after which a workspace will be idled if there is no activity. To disable workspace idling due to inactivity, set this value to -1.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"seconds_of_run_before_idling": {
								Description:         "Run timeout for workspaces in seconds. This timeout is the maximum duration a workspace runs. To disable workspace run timeout, set this value to -1.",
								MarkdownDescription: "Run timeout for workspaces in seconds. This timeout is the maximum duration a workspace runs. To disable workspace run timeout, set this value to -1.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"git_services": {
						Description:         "A configuration that allows users to work with remote Git repositories.",
						MarkdownDescription: "A configuration that allows users to work with remote Git repositories.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"bitbucket": {
								Description:         "Enables users to work with repositories hosted on Bitbucket (bitbucket.org or self-hosted).",
								MarkdownDescription: "Enables users to work with repositories hosted on Bitbucket (bitbucket.org or self-hosted).",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"endpoint": {
										Description:         "Bitbucket server endpoint URL.",
										MarkdownDescription: "Bitbucket server endpoint URL.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_name": {
										Description:         "Kubernetes secret, that contains Base64-encoded Bitbucket OAuth 1.0 or OAuth 2.0 data. For OAuth 1.0: private key, Bitbucket Application link consumer key and Bitbucket Application link shared secret must be stored in 'private.key', 'consumer.key' and 'shared_secret' keys respectively. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-1-for-a-bitbucket-server/. For OAuth 2.0: Bitbucket OAuth consumer key and Bitbucket OAuth consumer secret must be stored in 'id' and 'secret' keys respectively. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-the-bitbucket-cloud/.",
										MarkdownDescription: "Kubernetes secret, that contains Base64-encoded Bitbucket OAuth 1.0 or OAuth 2.0 data. For OAuth 1.0: private key, Bitbucket Application link consumer key and Bitbucket Application link shared secret must be stored in 'private.key', 'consumer.key' and 'shared_secret' keys respectively. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-1-for-a-bitbucket-server/. For OAuth 2.0: Bitbucket OAuth consumer key and Bitbucket OAuth consumer secret must be stored in 'id' and 'secret' keys respectively. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-the-bitbucket-cloud/.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"github": {
								Description:         "Enables users to work with repositories hosted on GitHub (github.com or GitHub Enterprise).",
								MarkdownDescription: "Enables users to work with repositories hosted on GitHub (github.com or GitHub Enterprise).",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"endpoint": {
										Description:         "GitHub server endpoint URL.",
										MarkdownDescription: "GitHub server endpoint URL.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"secret_name": {
										Description:         "Kubernetes secret, that contains Base64-encoded GitHub OAuth Client id and GitHub OAuth Client secret, that stored in 'id' and 'secret' keys respectively. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-github/.",
										MarkdownDescription: "Kubernetes secret, that contains Base64-encoded GitHub OAuth Client id and GitHub OAuth Client secret, that stored in 'id' and 'secret' keys respectively. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-github/.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"gitlab": {
								Description:         "Enables users to work with repositories hosted on GitLab (gitlab.com or self-hosted).",
								MarkdownDescription: "Enables users to work with repositories hosted on GitLab (gitlab.com or self-hosted).",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"endpoint": {
										Description:         "GitLab server endpoint URL.",
										MarkdownDescription: "GitLab server endpoint URL.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"secret_name": {
										Description:         "Kubernetes secret, that contains Base64-encoded GitHub Application id and GitLab Application Client secret, that stored in 'id' and 'secret' keys respectively. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-gitlab/.",
										MarkdownDescription: "Kubernetes secret, that contains Base64-encoded GitHub Application id and GitLab Application Client secret, that stored in 'id' and 'secret' keys respectively. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-gitlab/.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"image_puller": {
						Description:         "Kubernetes Image Puller configuration",
						MarkdownDescription: "Kubernetes Image Puller configuration",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enable": {
								Description:         "Install and configure the Community Supported Kubernetes Image Puller Operator. When set to 'true' and no spec is provided, it will create a default KubernetesImagePuller object to be managed by the Operator. When set to 'false', the KubernetesImagePuller object will be deleted, and the Operator will be uninstalled, regardless of whether a spec is provided. If the 'spec.images' field is empty, a set of recommended workspace-related images will be automatically detected and pre-pulled after installation. Note that while this Operator and its behavior is community-supported, its payload may be commercially-supported for pulling commercially-supported images.",
								MarkdownDescription: "Install and configure the Community Supported Kubernetes Image Puller Operator. When set to 'true' and no spec is provided, it will create a default KubernetesImagePuller object to be managed by the Operator. When set to 'false', the KubernetesImagePuller object will be deleted, and the Operator will be uninstalled, regardless of whether a spec is provided. If the 'spec.images' field is empty, a set of recommended workspace-related images will be automatically detected and pre-pulled after installation. Note that while this Operator and its behavior is community-supported, its payload may be commercially-supported for pulling commercially-supported images.",

								Type: types.BoolType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"spec": {
								Description:         "A KubernetesImagePullerSpec to configure the image puller in the CheCluster",
								MarkdownDescription: "A KubernetesImagePullerSpec to configure the image puller in the CheCluster",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"affinity": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"caching_cpu_limit": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"caching_cpu_request": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"caching_interval_hours": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"caching_memory_limit": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"caching_memory_request": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"config_map_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"daemonset_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"deployment_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image_pull_secrets": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image_puller_image": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"images": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"node_selector": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"k8s": {
						Description:         "Configuration settings specific to Che installations made on upstream Kubernetes.",
						MarkdownDescription: "Configuration settings specific to Che installations made on upstream Kubernetes.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"ingress_class": {
								Description:         "Ingress class that will define the which controller will manage ingresses. Defaults to 'nginx'. NB: This drives the 'kubernetes.io/ingress.class' annotation on Che-related ingresses.",
								MarkdownDescription: "Ingress class that will define the which controller will manage ingresses. Defaults to 'nginx'. NB: This drives the 'kubernetes.io/ingress.class' annotation on Che-related ingresses.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ingress_domain": {
								Description:         "Global ingress domain for a Kubernetes cluster. This MUST be explicitly specified: there are no defaults.",
								MarkdownDescription: "Global ingress domain for a Kubernetes cluster. This MUST be explicitly specified: there are no defaults.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ingress_strategy": {
								Description:         "Deprecated. The value of this flag is ignored. Strategy for ingress creation. Options are: 'multi-host' (host is explicitly provided in ingress), 'single-host' (host is provided, path-based rules) and 'default-host' (no host is provided, path-based rules). Defaults to 'multi-host' Deprecated in favor of 'serverExposureStrategy' in the 'server' section, which defines this regardless of the cluster type. When both are defined, the 'serverExposureStrategy' option takes precedence.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Strategy for ingress creation. Options are: 'multi-host' (host is explicitly provided in ingress), 'single-host' (host is provided, path-based rules) and 'default-host' (no host is provided, path-based rules). Defaults to 'multi-host' Deprecated in favor of 'serverExposureStrategy' in the 'server' section, which defines this regardless of the cluster type. When both are defined, the 'serverExposureStrategy' option takes precedence.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"security_context_fs_group": {
								Description:         "The FSGroup in which the Che Pod and workspace Pods containers runs in. Default value is '1724'.",
								MarkdownDescription: "The FSGroup in which the Che Pod and workspace Pods containers runs in. Default value is '1724'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"security_context_run_as_user": {
								Description:         "ID of the user the Che Pod and workspace Pods containers run as. Default value is '1724'.",
								MarkdownDescription: "ID of the user the Che Pod and workspace Pods containers run as. Default value is '1724'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"single_host_exposure_type": {
								Description:         "Deprecated. The value of this flag is ignored. When the serverExposureStrategy is set to 'single-host', the way the server, registries and workspaces are exposed is further configured by this property. The possible values are 'native', which means that the server and workspaces are exposed using ingresses on K8s or 'gateway' where the server and workspaces are exposed using a custom gateway based on link:https://doc.traefik.io/traefik/[Traefik]. All the endpoints whether backed by the ingress or gateway 'route' always point to the subpaths on the same domain. Defaults to 'native'.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. When the serverExposureStrategy is set to 'single-host', the way the server, registries and workspaces are exposed is further configured by this property. The possible values are 'native', which means that the server and workspaces are exposed using ingresses on K8s or 'gateway' where the server and workspaces are exposed using a custom gateway based on link:https://doc.traefik.io/traefik/[Traefik]. All the endpoints whether backed by the ingress or gateway 'route' always point to the subpaths on the same domain. Defaults to 'native'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls_secret_name": {
								Description:         "Name of a secret that will be used to setup ingress TLS termination when TLS is enabled. When the field is empty string, the default cluster certificate will be used. See also the 'tlsSupport' field.",
								MarkdownDescription: "Name of a secret that will be used to setup ingress TLS termination when TLS is enabled. When the field is empty string, the default cluster certificate will be used. See also the 'tlsSupport' field.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"metrics": {
						Description:         "Configuration settings related to the metrics collection used by the Che installation.",
						MarkdownDescription: "Configuration settings related to the metrics collection used by the Che installation.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enable": {
								Description:         "Enables 'metrics' the Che server endpoint. Default to 'true'.",
								MarkdownDescription: "Enables 'metrics' the Che server endpoint. Default to 'true'.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"server": {
						Description:         "General configuration settings related to the Che server, the plugin and devfile registries",
						MarkdownDescription: "General configuration settings related to the Che server, the plugin and devfile registries",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"air_gap_container_registry_hostname": {
								Description:         "Optional host name, or URL, to an alternate container registry to pull images from. This value overrides the container registry host name defined in all the default container images involved in a Che deployment. This is particularly useful to install Che in a restricted environment.",
								MarkdownDescription: "Optional host name, or URL, to an alternate container registry to pull images from. This value overrides the container registry host name defined in all the default container images involved in a Che deployment. This is particularly useful to install Che in a restricted environment.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"air_gap_container_registry_organization": {
								Description:         "Optional repository name of an alternate container registry to pull images from. This value overrides the container registry organization defined in all the default container images involved in a Che deployment. This is particularly useful to install Eclipse Che in a restricted environment.",
								MarkdownDescription: "Optional repository name of an alternate container registry to pull images from. This value overrides the container registry organization defined in all the default container images involved in a Che deployment. This is particularly useful to install Eclipse Che in a restricted environment.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"allow_auto_provision_user_namespace": {
								Description:         "Indicates if is allowed to automatically create a user namespace. If it set to false, then user namespace must be pre-created by a cluster administrator.",
								MarkdownDescription: "Indicates if is allowed to automatically create a user namespace. If it set to false, then user namespace must be pre-created by a cluster administrator.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"allow_user_defined_workspace_namespaces": {
								Description:         "Deprecated. The value of this flag is ignored. Defines that a user is allowed to specify a Kubernetes namespace, or an OpenShift project, which differs from the default. It's NOT RECOMMENDED to set to 'true' without OpenShift OAuth configured. The OpenShift infrastructure also uses this property.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Defines that a user is allowed to specify a Kubernetes namespace, or an OpenShift project, which differs from the default. It's NOT RECOMMENDED to set to 'true' without OpenShift OAuth configured. The OpenShift infrastructure also uses this property.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"che_cluster_roles": {
								Description:         "A comma-separated list of ClusterRoles that will be assigned to Che ServiceAccount. Each role must have 'app.kubernetes.io/part-of=che.eclipse.org' label. Be aware that the Che Operator has to already have all permissions in these ClusterRoles to grant them.",
								MarkdownDescription: "A comma-separated list of ClusterRoles that will be assigned to Che ServiceAccount. Each role must have 'app.kubernetes.io/part-of=che.eclipse.org' label. Be aware that the Che Operator has to already have all permissions in these ClusterRoles to grant them.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"che_debug": {
								Description:         "Enables the debug mode for Che server. Defaults to 'false'.",
								MarkdownDescription: "Enables the debug mode for Che server. Defaults to 'false'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"che_flavor": {
								Description:         "Deprecated. The value of this flag is ignored. Specifies a variation of the installation. The options are  'che' for upstream Che installations or 'devspaces' for Red Hat OpenShift Dev Spaces (formerly Red Hat CodeReady Workspaces) installation",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Specifies a variation of the installation. The options are  'che' for upstream Che installations or 'devspaces' for Red Hat OpenShift Dev Spaces (formerly Red Hat CodeReady Workspaces) installation",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"che_host": {
								Description:         "Public host name of the installed Che server. When value is omitted, the value it will be automatically set by the Operator. See the 'cheHostTLSSecret' field.",
								MarkdownDescription: "Public host name of the installed Che server. When value is omitted, the value it will be automatically set by the Operator. See the 'cheHostTLSSecret' field.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"che_host_tls_secret": {
								Description:         "Name of a secret containing certificates to secure ingress or route for the custom host name of the installed Che server. The secret must have 'app.kubernetes.io/part-of=che.eclipse.org' label. See the 'cheHost' field.",
								MarkdownDescription: "Name of a secret containing certificates to secure ingress or route for the custom host name of the installed Che server. The secret must have 'app.kubernetes.io/part-of=che.eclipse.org' label. See the 'cheHost' field.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"che_image": {
								Description:         "Overrides the container image used in Che deployment. This does NOT include the container image tag. Omit it or leave it empty to use the default container image provided by the Operator.",
								MarkdownDescription: "Overrides the container image used in Che deployment. This does NOT include the container image tag. Omit it or leave it empty to use the default container image provided by the Operator.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"che_image_pull_policy": {
								Description:         "Overrides the image pull policy used in Che deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
								MarkdownDescription: "Overrides the image pull policy used in Che deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"che_image_tag": {
								Description:         "Overrides the tag of the container image used in Che deployment. Omit it or leave it empty to use the default image tag provided by the Operator.",
								MarkdownDescription: "Overrides the tag of the container image used in Che deployment. Omit it or leave it empty to use the default image tag provided by the Operator.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"che_log_level": {
								Description:         "Log level for the Che server: 'INFO' or 'DEBUG'. Defaults to 'INFO'.",
								MarkdownDescription: "Log level for the Che server: 'INFO' or 'DEBUG'. Defaults to 'INFO'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"che_server_env": {
								Description:         "List of environment variables to set in the Che server container.",
								MarkdownDescription: "List of environment variables to set in the Che server container.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
										MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"value": {
										Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
										MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value_from": {
										Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
										MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_key_ref": {
												Description:         "Selects a key of a ConfigMap.",
												MarkdownDescription: "Selects a key of a ConfigMap.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the ConfigMap or its key must be defined",
														MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_ref": {
												Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
												MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"api_version": {
														Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
														MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_path": {
														Description:         "Path of the field to select in the specified API version.",
														MarkdownDescription: "Path of the field to select in the specified API version.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resource_field_ref": {
												Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
												MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"container_name": {
														Description:         "Container name: required for volumes, optional for env vars",
														MarkdownDescription: "Container name: required for volumes, optional for env vars",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"divisor": {
														Description:         "Specifies the output format of the exposed resources, defaults to '1'",
														MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resource": {
														Description:         "Required: resource to select",
														MarkdownDescription: "Required: resource to select",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_key_ref": {
												Description:         "Selects a key of a secret in the pod's namespace",
												MarkdownDescription: "Selects a key of a secret in the pod's namespace",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the Secret or its key must be defined",
														MarkdownDescription: "Specify whether the Secret or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"che_server_ingress": {
								Description:         "The Che server ingress custom settings.",
								MarkdownDescription: "The Che server ingress custom settings.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										MarkdownDescription: "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"labels": {
										Description:         "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										MarkdownDescription: "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"che_server_route": {
								Description:         "The Che server route custom settings.",
								MarkdownDescription: "The Che server route custom settings.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										MarkdownDescription: "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"domain": {
										Description:         "Operator uses the domain to generate a hostname for a route. In a conjunction with labels it creates a route, which is served by a non-default Ingress controller. The generated host name will follow this pattern: '<route-name>-<route-namespace>.<domain>'.",
										MarkdownDescription: "Operator uses the domain to generate a hostname for a route. In a conjunction with labels it creates a route, which is served by a non-default Ingress controller. The generated host name will follow this pattern: '<route-name>-<route-namespace>.<domain>'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"labels": {
										Description:         "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										MarkdownDescription: "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"che_workspace_cluster_role": {
								Description:         "Custom cluster role bound to the user for the Che workspaces. The role must have 'app.kubernetes.io/part-of=che.eclipse.org' label. The default roles are used when omitted or left blank.",
								MarkdownDescription: "Custom cluster role bound to the user for the Che workspaces. The role must have 'app.kubernetes.io/part-of=che.eclipse.org' label. The default roles are used when omitted or left blank.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"custom_che_properties": {
								Description:         "Map of additional environment variables that will be applied in the generated 'che' ConfigMap to be used by the Che server, in addition to the values already generated from other fields of the 'CheCluster' custom resource (CR). When 'customCheProperties' contains a property that would be normally generated in 'che' ConfigMap from other CR fields, the value defined in the 'customCheProperties' is used instead.",
								MarkdownDescription: "Map of additional environment variables that will be applied in the generated 'che' ConfigMap to be used by the Che server, in addition to the values already generated from other fields of the 'CheCluster' custom resource (CR). When 'customCheProperties' contains a property that would be normally generated in 'che' ConfigMap from other CR fields, the value defined in the 'customCheProperties' is used instead.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dashboard_cpu_limit": {
								Description:         "Overrides the CPU limit used in the dashboard deployment. In cores. (500m = .5 cores). Default to 500m.",
								MarkdownDescription: "Overrides the CPU limit used in the dashboard deployment. In cores. (500m = .5 cores). Default to 500m.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dashboard_cpu_request": {
								Description:         "Overrides the CPU request used in the dashboard deployment. In cores. (500m = .5 cores). Default to 100m.",
								MarkdownDescription: "Overrides the CPU request used in the dashboard deployment. In cores. (500m = .5 cores). Default to 100m.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dashboard_env": {
								Description:         "List of environment variables to set in the dashboard container.",
								MarkdownDescription: "List of environment variables to set in the dashboard container.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
										MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"value": {
										Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
										MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value_from": {
										Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
										MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_key_ref": {
												Description:         "Selects a key of a ConfigMap.",
												MarkdownDescription: "Selects a key of a ConfigMap.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the ConfigMap or its key must be defined",
														MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_ref": {
												Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
												MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"api_version": {
														Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
														MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_path": {
														Description:         "Path of the field to select in the specified API version.",
														MarkdownDescription: "Path of the field to select in the specified API version.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resource_field_ref": {
												Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
												MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"container_name": {
														Description:         "Container name: required for volumes, optional for env vars",
														MarkdownDescription: "Container name: required for volumes, optional for env vars",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"divisor": {
														Description:         "Specifies the output format of the exposed resources, defaults to '1'",
														MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resource": {
														Description:         "Required: resource to select",
														MarkdownDescription: "Required: resource to select",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_key_ref": {
												Description:         "Selects a key of a secret in the pod's namespace",
												MarkdownDescription: "Selects a key of a secret in the pod's namespace",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the Secret or its key must be defined",
														MarkdownDescription: "Specify whether the Secret or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dashboard_image": {
								Description:         "Overrides the container image used in the dashboard deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",
								MarkdownDescription: "Overrides the container image used in the dashboard deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dashboard_image_pull_policy": {
								Description:         "Overrides the image pull policy used in the dashboard deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
								MarkdownDescription: "Overrides the image pull policy used in the dashboard deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dashboard_ingress": {
								Description:         "Deprecated. The value of this flag is ignored. Dashboard ingress custom settings.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Dashboard ingress custom settings.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										MarkdownDescription: "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"labels": {
										Description:         "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										MarkdownDescription: "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dashboard_memory_limit": {
								Description:         "Overrides the memory limit used in the dashboard deployment. Defaults to 256Mi.",
								MarkdownDescription: "Overrides the memory limit used in the dashboard deployment. Defaults to 256Mi.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dashboard_memory_request": {
								Description:         "Overrides the memory request used in the dashboard deployment. Defaults to 16Mi.",
								MarkdownDescription: "Overrides the memory request used in the dashboard deployment. Defaults to 16Mi.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dashboard_route": {
								Description:         "Deprecated. The value of this flag is ignored. Dashboard route custom settings.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Dashboard route custom settings.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										MarkdownDescription: "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"domain": {
										Description:         "Operator uses the domain to generate a hostname for a route. In a conjunction with labels it creates a route, which is served by a non-default Ingress controller. The generated host name will follow this pattern: '<route-name>-<route-namespace>.<domain>'.",
										MarkdownDescription: "Operator uses the domain to generate a hostname for a route. In a conjunction with labels it creates a route, which is served by a non-default Ingress controller. The generated host name will follow this pattern: '<route-name>-<route-namespace>.<domain>'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"labels": {
										Description:         "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										MarkdownDescription: "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"devfile_registry_cpu_limit": {
								Description:         "Overrides the CPU limit used in the devfile registry deployment. In cores. (500m = .5 cores). Default to 500m.",
								MarkdownDescription: "Overrides the CPU limit used in the devfile registry deployment. In cores. (500m = .5 cores). Default to 500m.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"devfile_registry_cpu_request": {
								Description:         "Overrides the CPU request used in the devfile registry deployment. In cores. (500m = .5 cores). Default to 100m.",
								MarkdownDescription: "Overrides the CPU request used in the devfile registry deployment. In cores. (500m = .5 cores). Default to 100m.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"devfile_registry_env": {
								Description:         "List of environment variables to set in the plugin registry container.",
								MarkdownDescription: "List of environment variables to set in the plugin registry container.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
										MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"value": {
										Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
										MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value_from": {
										Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
										MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_key_ref": {
												Description:         "Selects a key of a ConfigMap.",
												MarkdownDescription: "Selects a key of a ConfigMap.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the ConfigMap or its key must be defined",
														MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_ref": {
												Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
												MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"api_version": {
														Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
														MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_path": {
														Description:         "Path of the field to select in the specified API version.",
														MarkdownDescription: "Path of the field to select in the specified API version.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resource_field_ref": {
												Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
												MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"container_name": {
														Description:         "Container name: required for volumes, optional for env vars",
														MarkdownDescription: "Container name: required for volumes, optional for env vars",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"divisor": {
														Description:         "Specifies the output format of the exposed resources, defaults to '1'",
														MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resource": {
														Description:         "Required: resource to select",
														MarkdownDescription: "Required: resource to select",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_key_ref": {
												Description:         "Selects a key of a secret in the pod's namespace",
												MarkdownDescription: "Selects a key of a secret in the pod's namespace",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the Secret or its key must be defined",
														MarkdownDescription: "Specify whether the Secret or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"devfile_registry_image": {
								Description:         "Overrides the container image used in the devfile registry deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",
								MarkdownDescription: "Overrides the container image used in the devfile registry deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"devfile_registry_ingress": {
								Description:         "Deprecated. The value of this flag is ignored. The devfile registry ingress custom settings.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. The devfile registry ingress custom settings.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										MarkdownDescription: "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"labels": {
										Description:         "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										MarkdownDescription: "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"devfile_registry_memory_limit": {
								Description:         "Overrides the memory limit used in the devfile registry deployment. Defaults to 256Mi.",
								MarkdownDescription: "Overrides the memory limit used in the devfile registry deployment. Defaults to 256Mi.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"devfile_registry_memory_request": {
								Description:         "Overrides the memory request used in the devfile registry deployment. Defaults to 16Mi.",
								MarkdownDescription: "Overrides the memory request used in the devfile registry deployment. Defaults to 16Mi.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"devfile_registry_pull_policy": {
								Description:         "Overrides the image pull policy used in the devfile registry deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
								MarkdownDescription: "Overrides the image pull policy used in the devfile registry deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"devfile_registry_route": {
								Description:         "Deprecated. The value of this flag is ignored. The devfile registry route custom settings.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. The devfile registry route custom settings.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										MarkdownDescription: "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"domain": {
										Description:         "Operator uses the domain to generate a hostname for a route. In a conjunction with labels it creates a route, which is served by a non-default Ingress controller. The generated host name will follow this pattern: '<route-name>-<route-namespace>.<domain>'.",
										MarkdownDescription: "Operator uses the domain to generate a hostname for a route. In a conjunction with labels it creates a route, which is served by a non-default Ingress controller. The generated host name will follow this pattern: '<route-name>-<route-namespace>.<domain>'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"labels": {
										Description:         "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										MarkdownDescription: "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"devfile_registry_url": {
								Description:         "Deprecated in favor of 'externalDevfileRegistries' fields.",
								MarkdownDescription: "Deprecated in favor of 'externalDevfileRegistries' fields.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disable_internal_cluster_svc_names": {
								Description:         "Deprecated. The value of this flag is ignored. Disable internal cluster SVC names usage to communicate between components to speed up the traffic and avoid proxy issues.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Disable internal cluster SVC names usage to communicate between components to speed up the traffic and avoid proxy issues.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"external_devfile_registries": {
								Description:         "External devfile registries, that serves sample, ready-to-use devfiles. Configure this in addition to a dedicated devfile registry (when 'externalDevfileRegistry' is 'false') or instead of it (when 'externalDevfileRegistry' is 'true')",
								MarkdownDescription: "External devfile registries, that serves sample, ready-to-use devfiles. Configure this in addition to a dedicated devfile registry (when 'externalDevfileRegistry' is 'false') or instead of it (when 'externalDevfileRegistry' is 'true')",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"url": {
										Description:         "Public URL of the devfile registry.",
										MarkdownDescription: "Public URL of the devfile registry.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"external_devfile_registry": {
								Description:         "Instructs the Operator on whether to deploy a dedicated devfile registry server. By default, a dedicated devfile registry server is started. When 'externalDevfileRegistry' is 'true', no such dedicated server will be started by the Operator and configure at least one devfile registry with 'externalDevfileRegistries' field.",
								MarkdownDescription: "Instructs the Operator on whether to deploy a dedicated devfile registry server. By default, a dedicated devfile registry server is started. When 'externalDevfileRegistry' is 'true', no such dedicated server will be started by the Operator and configure at least one devfile registry with 'externalDevfileRegistries' field.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"external_plugin_registry": {
								Description:         "Instructs the Operator on whether to deploy a dedicated plugin registry server. By default, a dedicated plugin registry server is started. When 'externalPluginRegistry' is 'true', no such dedicated server will be started by the Operator and you will have to manually set the 'pluginRegistryUrl' field.",
								MarkdownDescription: "Instructs the Operator on whether to deploy a dedicated plugin registry server. By default, a dedicated plugin registry server is started. When 'externalPluginRegistry' is 'true', no such dedicated server will be started by the Operator and you will have to manually set the 'pluginRegistryUrl' field.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"git_self_signed_cert": {
								Description:         "When enabled, the certificate from 'che-git-self-signed-cert' ConfigMap will be propagated to the Che components and provide particular configuration for Git. Note, the 'che-git-self-signed-cert' ConfigMap must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",
								MarkdownDescription: "When enabled, the certificate from 'che-git-self-signed-cert' ConfigMap will be propagated to the Che components and provide particular configuration for Git. Note, the 'che-git-self-signed-cert' ConfigMap must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"non_proxy_hosts": {
								Description:         "List of hosts that will be reached directly, bypassing the proxy. Specify wild card domain use the following form '.<DOMAIN>' and '|' as delimiter, for example: 'localhost|.my.host.com|123.42.12.32' Only use when configuring a proxy is required. Operator respects OpenShift cluster wide proxy configuration and no additional configuration is required, but defining 'nonProxyHosts' in a custom resource leads to merging non proxy hosts lists from the cluster proxy configuration and ones defined in the custom resources. See the doc https://docs.openshift.com/container-platform/4.4/networking/enable-cluster-wide-proxy.html. See also the 'proxyURL' fields.",
								MarkdownDescription: "List of hosts that will be reached directly, bypassing the proxy. Specify wild card domain use the following form '.<DOMAIN>' and '|' as delimiter, for example: 'localhost|.my.host.com|123.42.12.32' Only use when configuring a proxy is required. Operator respects OpenShift cluster wide proxy configuration and no additional configuration is required, but defining 'nonProxyHosts' in a custom resource leads to merging non proxy hosts lists from the cluster proxy configuration and ones defined in the custom resources. See the doc https://docs.openshift.com/container-platform/4.4/networking/enable-cluster-wide-proxy.html. See also the 'proxyURL' fields.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"open_vsx_registry_url": {
								Description:         "Open VSX registry URL. If omitted an embedded instance will be used.",
								MarkdownDescription: "Open VSX registry URL. If omitted an embedded instance will be used.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"plugin_registry_cpu_limit": {
								Description:         "Overrides the CPU limit used in the plugin registry deployment. In cores. (500m = .5 cores). Default to 500m.",
								MarkdownDescription: "Overrides the CPU limit used in the plugin registry deployment. In cores. (500m = .5 cores). Default to 500m.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"plugin_registry_cpu_request": {
								Description:         "Overrides the CPU request used in the plugin registry deployment. In cores. (500m = .5 cores). Default to 100m.",
								MarkdownDescription: "Overrides the CPU request used in the plugin registry deployment. In cores. (500m = .5 cores). Default to 100m.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"plugin_registry_env": {
								Description:         "List of environment variables to set in the devfile registry container.",
								MarkdownDescription: "List of environment variables to set in the devfile registry container.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
										MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"value": {
										Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
										MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value_from": {
										Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
										MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_key_ref": {
												Description:         "Selects a key of a ConfigMap.",
												MarkdownDescription: "Selects a key of a ConfigMap.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key to select.",
														MarkdownDescription: "The key to select.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the ConfigMap or its key must be defined",
														MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"field_ref": {
												Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
												MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"api_version": {
														Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
														MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"field_path": {
														Description:         "Path of the field to select in the specified API version.",
														MarkdownDescription: "Path of the field to select in the specified API version.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resource_field_ref": {
												Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
												MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"container_name": {
														Description:         "Container name: required for volumes, optional for env vars",
														MarkdownDescription: "Container name: required for volumes, optional for env vars",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"divisor": {
														Description:         "Specifies the output format of the exposed resources, defaults to '1'",
														MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resource": {
														Description:         "Required: resource to select",
														MarkdownDescription: "Required: resource to select",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_key_ref": {
												Description:         "Selects a key of a secret in the pod's namespace",
												MarkdownDescription: "Selects a key of a secret in the pod's namespace",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the Secret or its key must be defined",
														MarkdownDescription: "Specify whether the Secret or its key must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"plugin_registry_image": {
								Description:         "Overrides the container image used in the plugin registry deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",
								MarkdownDescription: "Overrides the container image used in the plugin registry deployment. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"plugin_registry_ingress": {
								Description:         "Deprecated. The value of this flag is ignored. Plugin registry ingress custom settings.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Plugin registry ingress custom settings.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										MarkdownDescription: "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"labels": {
										Description:         "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										MarkdownDescription: "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"plugin_registry_memory_limit": {
								Description:         "Overrides the memory limit used in the plugin registry deployment. Defaults to 1536Mi.",
								MarkdownDescription: "Overrides the memory limit used in the plugin registry deployment. Defaults to 1536Mi.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"plugin_registry_memory_request": {
								Description:         "Overrides the memory request used in the plugin registry deployment. Defaults to 16Mi.",
								MarkdownDescription: "Overrides the memory request used in the plugin registry deployment. Defaults to 16Mi.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"plugin_registry_pull_policy": {
								Description:         "Overrides the image pull policy used in the plugin registry deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
								MarkdownDescription: "Overrides the image pull policy used in the plugin registry deployment. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"plugin_registry_route": {
								Description:         "Deprecated. The value of this flag is ignored. Plugin registry route custom settings.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Plugin registry route custom settings.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",
										MarkdownDescription: "Unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"domain": {
										Description:         "Operator uses the domain to generate a hostname for a route. In a conjunction with labels it creates a route, which is served by a non-default Ingress controller. The generated host name will follow this pattern: '<route-name>-<route-namespace>.<domain>'.",
										MarkdownDescription: "Operator uses the domain to generate a hostname for a route. In a conjunction with labels it creates a route, which is served by a non-default Ingress controller. The generated host name will follow this pattern: '<route-name>-<route-namespace>.<domain>'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"labels": {
										Description:         "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",
										MarkdownDescription: "Comma separated list of labels that can be used to organize and categorize objects by scoping and selecting.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"plugin_registry_url": {
								Description:         "Public URL of the plugin registry that serves sample ready-to-use devfiles. Set this ONLY when a use of an external devfile registry is needed. See the 'externalPluginRegistry' field. By default, this will be automatically calculated by the Operator.",
								MarkdownDescription: "Public URL of the plugin registry that serves sample ready-to-use devfiles. Set this ONLY when a use of an external devfile registry is needed. See the 'externalPluginRegistry' field. By default, this will be automatically calculated by the Operator.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"proxy_password": {
								Description:         "Password of the proxy server. Only use when proxy configuration is required. See the 'proxyURL', 'proxyUser' and 'proxySecret' fields.",
								MarkdownDescription: "Password of the proxy server. Only use when proxy configuration is required. See the 'proxyURL', 'proxyUser' and 'proxySecret' fields.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"proxy_port": {
								Description:         "Port of the proxy server. Only use when configuring a proxy is required. See also the 'proxyURL' and 'nonProxyHosts' fields.",
								MarkdownDescription: "Port of the proxy server. Only use when configuring a proxy is required. See also the 'proxyURL' and 'nonProxyHosts' fields.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"proxy_secret": {
								Description:         "The secret that contains 'user' and 'password' for a proxy server. When the secret is defined, the 'proxyUser' and 'proxyPassword' are ignored. The secret must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",
								MarkdownDescription: "The secret that contains 'user' and 'password' for a proxy server. When the secret is defined, the 'proxyUser' and 'proxyPassword' are ignored. The secret must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"proxy_url": {
								Description:         "URL (protocol+host name) of the proxy server. This drives the appropriate changes in the 'JAVA_OPTS' and 'https(s)_proxy' variables in the Che server and workspaces containers. Only use when configuring a proxy is required. Operator respects OpenShift cluster wide proxy configuration and no additional configuration is required, but defining 'proxyUrl' in a custom resource leads to overrides the cluster proxy configuration with fields 'proxyUrl', 'proxyPort', 'proxyUser' and 'proxyPassword' from the custom resource. See the doc https://docs.openshift.com/container-platform/4.4/networking/enable-cluster-wide-proxy.html. See also the 'proxyPort' and 'nonProxyHosts' fields.",
								MarkdownDescription: "URL (protocol+host name) of the proxy server. This drives the appropriate changes in the 'JAVA_OPTS' and 'https(s)_proxy' variables in the Che server and workspaces containers. Only use when configuring a proxy is required. Operator respects OpenShift cluster wide proxy configuration and no additional configuration is required, but defining 'proxyUrl' in a custom resource leads to overrides the cluster proxy configuration with fields 'proxyUrl', 'proxyPort', 'proxyUser' and 'proxyPassword' from the custom resource. See the doc https://docs.openshift.com/container-platform/4.4/networking/enable-cluster-wide-proxy.html. See also the 'proxyPort' and 'nonProxyHosts' fields.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"proxy_user": {
								Description:         "User name of the proxy server. Only use when configuring a proxy is required. See also the 'proxyURL', 'proxyPassword' and 'proxySecret' fields.",
								MarkdownDescription: "User name of the proxy server. Only use when configuring a proxy is required. See also the 'proxyURL', 'proxyPassword' and 'proxySecret' fields.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"self_signed_cert": {
								Description:         "Deprecated. The value of this flag is ignored. The Che Operator will automatically detect whether the router certificate is self-signed and propagate it to other components, such as the Che server.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. The Che Operator will automatically detect whether the router certificate is self-signed and propagate it to other components, such as the Che server.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"server_cpu_limit": {
								Description:         "Overrides the CPU limit used in the Che server deployment In cores. (500m = .5 cores). Default to 1.",
								MarkdownDescription: "Overrides the CPU limit used in the Che server deployment In cores. (500m = .5 cores). Default to 1.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"server_cpu_request": {
								Description:         "Overrides the CPU request used in the Che server deployment In cores. (500m = .5 cores). Default to 100m.",
								MarkdownDescription: "Overrides the CPU request used in the Che server deployment In cores. (500m = .5 cores). Default to 100m.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"server_exposure_strategy": {
								Description:         "Deprecated. The value of this flag is ignored. Sets the server and workspaces exposure type. Possible values are 'multi-host', 'single-host', 'default-host'. Defaults to 'multi-host', which creates a separate ingress, or OpenShift routes, for every required endpoint. 'single-host' makes Che exposed on a single host name with workspaces exposed on subpaths. Read the docs to learn about the limitations of this approach. Also consult the 'singleHostExposureType' property to further configure how the Operator and the Che server make that happen on Kubernetes. 'default-host' exposes the Che server on the host of the cluster. Read the docs to learn about the limitations of this approach.",
								MarkdownDescription: "Deprecated. The value of this flag is ignored. Sets the server and workspaces exposure type. Possible values are 'multi-host', 'single-host', 'default-host'. Defaults to 'multi-host', which creates a separate ingress, or OpenShift routes, for every required endpoint. 'single-host' makes Che exposed on a single host name with workspaces exposed on subpaths. Read the docs to learn about the limitations of this approach. Also consult the 'singleHostExposureType' property to further configure how the Operator and the Che server make that happen on Kubernetes. 'default-host' exposes the Che server on the host of the cluster. Read the docs to learn about the limitations of this approach.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"server_memory_limit": {
								Description:         "Overrides the memory limit used in the Che server deployment. Defaults to 1Gi.",
								MarkdownDescription: "Overrides the memory limit used in the Che server deployment. Defaults to 1Gi.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"server_memory_request": {
								Description:         "Overrides the memory request used in the Che server deployment. Defaults to 512Mi.",
								MarkdownDescription: "Overrides the memory request used in the Che server deployment. Defaults to 512Mi.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"server_trust_store_config_map_name": {
								Description:         "Name of the ConfigMap with public certificates to add to Java trust store of the Che server. This is often required when adding the OpenShift OAuth provider, which has HTTPS endpoint signed with self-signed cert. The Che server must be aware of its CA cert to be able to request it. This is disabled by default. The Config Map must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",
								MarkdownDescription: "Name of the ConfigMap with public certificates to add to Java trust store of the Che server. This is often required when adding the OpenShift OAuth provider, which has HTTPS endpoint signed with self-signed cert. The Che server must be aware of its CA cert to be able to request it. This is disabled by default. The Config Map must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"single_host_gateway_config_map_labels": {
								Description:         "The labels that need to be present in the ConfigMaps representing the gateway configuration.",
								MarkdownDescription: "The labels that need to be present in the ConfigMaps representing the gateway configuration.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"single_host_gateway_config_sidecar_image": {
								Description:         "The image used for the gateway sidecar that provides configuration to the gateway. Omit it or leave it empty to use the default container image provided by the Operator.",
								MarkdownDescription: "The image used for the gateway sidecar that provides configuration to the gateway. Omit it or leave it empty to use the default container image provided by the Operator.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"single_host_gateway_image": {
								Description:         "The image used for the gateway in the single host mode. Omit it or leave it empty to use the default container image provided by the Operator.",
								MarkdownDescription: "The image used for the gateway in the single host mode. Omit it or leave it empty to use the default container image provided by the Operator.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls_support": {
								Description:         "Deprecated. Instructs the Operator to deploy Che in TLS mode. This is enabled by default. Disabling TLS sometimes cause malfunction of some Che components.",
								MarkdownDescription: "Deprecated. Instructs the Operator to deploy Che in TLS mode. This is enabled by default. Disabling TLS sometimes cause malfunction of some Che components.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"use_internal_cluster_svc_names": {
								Description:         "Deprecated in favor of 'disableInternalClusterSVCNames'.",
								MarkdownDescription: "Deprecated in favor of 'disableInternalClusterSVCNames'.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"workspace_default_components": {
								Description:         "Default components applied to DevWorkspaces. These default components are meant to be used when a Devfile does not contain any components.",
								MarkdownDescription: "Default components applied to DevWorkspaces. These default components are meant to be used when a Devfile does not contain any components.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"attributes": {
										Description:         "Map of implementation-dependant free-form YAML attributes.",
										MarkdownDescription: "Map of implementation-dependant free-form YAML attributes.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"component_type": {
										Description:         "Type of component",
										MarkdownDescription: "Type of component",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("Container", "Kubernetes", "Openshift", "Volume", "Image", "Plugin", "Custom"),
										},
									},

									"container": {
										Description:         "Allows adding and configuring devworkspace-related containers",
										MarkdownDescription: "Allows adding and configuring devworkspace-related containers",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotation": {
												Description:         "Annotations that should be added to specific resources for this container",
												MarkdownDescription: "Annotations that should be added to specific resources for this container",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"deployment": {
														Description:         "Annotations to be added to deployment",
														MarkdownDescription: "Annotations to be added to deployment",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"service": {
														Description:         "Annotations to be added to service",
														MarkdownDescription: "Annotations to be added to service",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"args": {
												Description:         "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command.  Defaults to an empty array, meaning use whatever is defined in the image.",
												MarkdownDescription: "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command.  Defaults to an empty array, meaning use whatever is defined in the image.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"command": {
												Description:         "The command to run in the dockerimage component instead of the default one provided in the image.  Defaults to an empty array, meaning use whatever is defined in the image.",
												MarkdownDescription: "The command to run in the dockerimage component instead of the default one provided in the image.  Defaults to an empty array, meaning use whatever is defined in the image.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"cpu_limit": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"cpu_request": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"dedicated_pod": {
												Description:         "Specify if a container should run in its own separated pod, instead of running as part of the main development environment pod.  Default value is 'false'",
												MarkdownDescription: "Specify if a container should run in its own separated pod, instead of running as part of the main development environment pod.  Default value is 'false'",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"endpoints": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"annotation": {
														Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
														MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"attributes": {
														Description:         "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
														MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"exposure": {
														Description:         "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
														MarkdownDescription: "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("public", "internal", "none"),
														},
													},

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtMost(63),

															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
													},

													"path": {
														Description:         "Path of the endpoint URL",
														MarkdownDescription: "Path of the endpoint URL",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"protocol": {
														Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
														MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
														},
													},

													"secure": {
														Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
														MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"target_port": {
														Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
														MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"env": {
												Description:         "Environment variables used in this container.  The following variables are reserved and cannot be overridden via env:   - '$PROJECTS_ROOT'   - '$PROJECT_SOURCE'",
												MarkdownDescription: "Environment variables used in this container.  The following variables are reserved and cannot be overridden via env:   - '$PROJECTS_ROOT'   - '$PROJECT_SOURCE'",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"value": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"image": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"memory_limit": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"memory_request": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mount_sources": {
												Description:         "Toggles whether or not the project source code should be mounted in the component.  Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",
												MarkdownDescription: "Toggles whether or not the project source code should be mounted in the component.  Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"source_mapping": {
												Description:         "Optional specification of the path in the container where project sources should be transferred/mounted when 'mountSources' is 'true'. When omitted, the default value of /projects is used.",
												MarkdownDescription: "Optional specification of the path in the container where project sources should be transferred/mounted when 'mountSources' is 'true'. When omitted, the default value of /projects is used.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"volume_mounts": {
												Description:         "List of volumes mounts that should be mounted is this container.",
												MarkdownDescription: "List of volumes mounts that should be mounted is this container.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "The volume mount name is the name of an existing 'Volume' component. If several containers mount the same volume name then they will reuse the same volume and will be able to access to the same files.",
														MarkdownDescription: "The volume mount name is the name of an existing 'Volume' component. If several containers mount the same volume name then they will reuse the same volume and will be able to access to the same files.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtMost(63),

															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
													},

													"path": {
														Description:         "The path in the component container where the volume should be mounted. If not path is mentioned, default path is the is '/<name>'.",
														MarkdownDescription: "The path in the component container where the volume should be mounted. If not path is mentioned, default path is the is '/<name>'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"custom": {
										Description:         "Custom component whose logic is implementation-dependant and should be provided by the user possibly through some dedicated controller",
										MarkdownDescription: "Custom component whose logic is implementation-dependant and should be provided by the user possibly through some dedicated controller",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"component_class": {
												Description:         "Class of component that the associated implementation controller should use to process this command with the appropriate logic",
												MarkdownDescription: "Class of component that the associated implementation controller should use to process this command with the appropriate logic",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"embedded_resource": {
												Description:         "Additional free-form configuration for this custom component that the implementation controller will know how to use",
												MarkdownDescription: "Additional free-form configuration for this custom component that the implementation controller will know how to use",

												Type: utilities.DynamicType{},

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image": {
										Description:         "Allows specifying the definition of an image for outer loop builds",
										MarkdownDescription: "Allows specifying the definition of an image for outer loop builds",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"auto_build": {
												Description:         "Defines if the image should be built during startup.  Default value is 'false'",
												MarkdownDescription: "Defines if the image should be built during startup.  Default value is 'false'",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"dockerfile": {
												Description:         "Allows specifying dockerfile type build",
												MarkdownDescription: "Allows specifying dockerfile type build",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"args": {
														Description:         "The arguments to supply to the dockerfile build.",
														MarkdownDescription: "The arguments to supply to the dockerfile build.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"build_context": {
														Description:         "Path of source directory to establish build context. Defaults to ${PROJECT_SOURCE} in the container",
														MarkdownDescription: "Path of source directory to establish build context. Defaults to ${PROJECT_SOURCE} in the container",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"devfile_registry": {
														Description:         "Dockerfile's Devfile Registry source",
														MarkdownDescription: "Dockerfile's Devfile Registry source",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"id": {
																Description:         "Id in a devfile registry that contains a Dockerfile. The src in the OCI registry required for the Dockerfile build will be downloaded for building the image.",
																MarkdownDescription: "Id in a devfile registry that contains a Dockerfile. The src in the OCI registry required for the Dockerfile build will be downloaded for building the image.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"registry_url": {
																Description:         "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src. To ensure the Dockerfile gets resolved consistently in different environments, it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",
																MarkdownDescription: "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src. To ensure the Dockerfile gets resolved consistently in different environments, it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"git": {
														Description:         "Dockerfile's Git source",
														MarkdownDescription: "Dockerfile's Git source",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"checkout_from": {
																Description:         "Defines from what the project should be checked out. Required if there are more than one remote configured",
																MarkdownDescription: "Defines from what the project should be checked out. Required if there are more than one remote configured",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"remote": {
																		Description:         "The remote name should be used as init. Required if there are more than one remote configured",
																		MarkdownDescription: "The remote name should be used as init. Required if there are more than one remote configured",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"revision": {
																		Description:         "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
																		MarkdownDescription: "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"file_location": {
																Description:         "Location of the Dockerfile in the Git repository when using git as Dockerfile src. Defaults to Dockerfile.",
																MarkdownDescription: "Location of the Dockerfile in the Git repository when using git as Dockerfile src. Defaults to Dockerfile.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"remotes": {
																Description:         "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
																MarkdownDescription: "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",

																Type: types.MapType{ElemType: types.StringType},

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"root_required": {
														Description:         "Specify if a privileged builder pod is required.  Default value is 'false'",
														MarkdownDescription: "Specify if a privileged builder pod is required.  Default value is 'false'",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"src_type": {
														Description:         "Type of Dockerfile src",
														MarkdownDescription: "Type of Dockerfile src",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("Uri", "DevfileRegistry", "Git"),
														},
													},

													"uri": {
														Description:         "URI Reference of a Dockerfile. It can be a full URL or a relative URI from the current devfile as the base URI.",
														MarkdownDescription: "URI Reference of a Dockerfile. It can be a full URL or a relative URI from the current devfile as the base URI.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"image_name": {
												Description:         "Name of the image for the resulting outerloop build",
												MarkdownDescription: "Name of the image for the resulting outerloop build",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"image_type": {
												Description:         "Type of image",
												MarkdownDescription: "Type of image",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("Dockerfile"),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kubernetes": {
										Description:         "Allows importing into the devworkspace the Kubernetes resources defined in a given manifest. For example this allows reusing the Kubernetes definitions used to deploy some runtime components in production.",
										MarkdownDescription: "Allows importing into the devworkspace the Kubernetes resources defined in a given manifest. For example this allows reusing the Kubernetes definitions used to deploy some runtime components in production.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"deploy_by_default": {
												Description:         "Defines if the component should be deployed during startup.  Default value is 'false'",
												MarkdownDescription: "Defines if the component should be deployed during startup.  Default value is 'false'",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"endpoints": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"annotation": {
														Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
														MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"attributes": {
														Description:         "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
														MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"exposure": {
														Description:         "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
														MarkdownDescription: "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("public", "internal", "none"),
														},
													},

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtMost(63),

															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
													},

													"path": {
														Description:         "Path of the endpoint URL",
														MarkdownDescription: "Path of the endpoint URL",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"protocol": {
														Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
														MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
														},
													},

													"secure": {
														Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
														MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"target_port": {
														Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
														MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"inlined": {
												Description:         "Inlined manifest",
												MarkdownDescription: "Inlined manifest",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"location_type": {
												Description:         "Type of Kubernetes-like location",
												MarkdownDescription: "Type of Kubernetes-like location",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("Uri", "Inlined"),
												},
											},

											"uri": {
												Description:         "Location in a file fetched from a uri.",
												MarkdownDescription: "Location in a file fetched from a uri.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Mandatory name that allows referencing the component from other elements (such as commands) or from an external devfile that may reference this component through a parent or a plugin.",
										MarkdownDescription: "Mandatory name that allows referencing the component from other elements (such as commands) or from an external devfile that may reference this component through a parent or a plugin.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.LengthAtMost(63),

											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
										},
									},

									"openshift": {
										Description:         "Allows importing into the devworkspace the OpenShift resources defined in a given manifest. For example this allows reusing the OpenShift definitions used to deploy some runtime components in production.",
										MarkdownDescription: "Allows importing into the devworkspace the OpenShift resources defined in a given manifest. For example this allows reusing the OpenShift definitions used to deploy some runtime components in production.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"deploy_by_default": {
												Description:         "Defines if the component should be deployed during startup.  Default value is 'false'",
												MarkdownDescription: "Defines if the component should be deployed during startup.  Default value is 'false'",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"endpoints": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"annotation": {
														Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
														MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"attributes": {
														Description:         "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
														MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"exposure": {
														Description:         "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
														MarkdownDescription: "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("public", "internal", "none"),
														},
													},

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtMost(63),

															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
													},

													"path": {
														Description:         "Path of the endpoint URL",
														MarkdownDescription: "Path of the endpoint URL",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"protocol": {
														Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
														MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
														},
													},

													"secure": {
														Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
														MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"target_port": {
														Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
														MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"inlined": {
												Description:         "Inlined manifest",
												MarkdownDescription: "Inlined manifest",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"location_type": {
												Description:         "Type of Kubernetes-like location",
												MarkdownDescription: "Type of Kubernetes-like location",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("Uri", "Inlined"),
												},
											},

											"uri": {
												Description:         "Location in a file fetched from a uri.",
												MarkdownDescription: "Location in a file fetched from a uri.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"plugin": {
										Description:         "Allows importing a plugin.  Plugins are mainly imported devfiles that contribute components, commands and events as a consistent single unit. They are defined in either YAML files following the devfile syntax, or as 'DevWorkspaceTemplate' Kubernetes Custom Resources",
										MarkdownDescription: "Allows importing a plugin.  Plugins are mainly imported devfiles that contribute components, commands and events as a consistent single unit. They are defined in either YAML files following the devfile syntax, or as 'DevWorkspaceTemplate' Kubernetes Custom Resources",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"commands": {
												Description:         "Overrides of commands encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",
												MarkdownDescription: "Overrides of commands encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"apply": {
														Description:         "Command that consists in applying a given component definition, typically bound to a devworkspace event.  For example, when an 'apply' command is bound to a 'preStart' event, and references a 'container' component, it will start the container as a K8S initContainer in the devworkspace POD, unless the component has its 'dedicatedPod' field set to 'true'.  When no 'apply' command exist for a given component, it is assumed the component will be applied at devworkspace start by default, unless 'deployByDefault' for that component is set to false.",
														MarkdownDescription: "Command that consists in applying a given component definition, typically bound to a devworkspace event.  For example, when an 'apply' command is bound to a 'preStart' event, and references a 'container' component, it will start the container as a K8S initContainer in the devworkspace POD, unless the component has its 'dedicatedPod' field set to 'true'.  When no 'apply' command exist for a given component, it is assumed the component will be applied at devworkspace start by default, unless 'deployByDefault' for that component is set to false.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"component": {
																Description:         "Describes component that will be applied",
																MarkdownDescription: "Describes component that will be applied",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"group": {
																Description:         "Defines the group this command is part of",
																MarkdownDescription: "Defines the group this command is part of",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"is_default": {
																		Description:         "Identifies the default command for a given group kind",
																		MarkdownDescription: "Identifies the default command for a given group kind",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"kind": {
																		Description:         "Kind of group the command is part of",
																		MarkdownDescription: "Kind of group the command is part of",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.OneOf("build", "run", "test", "debug", "deploy"),
																		},
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"label": {
																Description:         "Optional label that provides a label for this command to be used in Editor UI menus for example",
																MarkdownDescription: "Optional label that provides a label for this command to be used in Editor UI menus for example",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"attributes": {
														Description:         "Map of implementation-dependant free-form YAML attributes.",
														MarkdownDescription: "Map of implementation-dependant free-form YAML attributes.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"command_type": {
														Description:         "Type of devworkspace command",
														MarkdownDescription: "Type of devworkspace command",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("Exec", "Apply", "Composite"),
														},
													},

													"composite": {
														Description:         "Composite command that allows executing several sub-commands either sequentially or concurrently",
														MarkdownDescription: "Composite command that allows executing several sub-commands either sequentially or concurrently",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"commands": {
																Description:         "The commands that comprise this composite command",
																MarkdownDescription: "The commands that comprise this composite command",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"group": {
																Description:         "Defines the group this command is part of",
																MarkdownDescription: "Defines the group this command is part of",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"is_default": {
																		Description:         "Identifies the default command for a given group kind",
																		MarkdownDescription: "Identifies the default command for a given group kind",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"kind": {
																		Description:         "Kind of group the command is part of",
																		MarkdownDescription: "Kind of group the command is part of",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.OneOf("build", "run", "test", "debug", "deploy"),
																		},
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"label": {
																Description:         "Optional label that provides a label for this command to be used in Editor UI menus for example",
																MarkdownDescription: "Optional label that provides a label for this command to be used in Editor UI menus for example",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"parallel": {
																Description:         "Indicates if the sub-commands should be executed concurrently",
																MarkdownDescription: "Indicates if the sub-commands should be executed concurrently",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"exec": {
														Description:         "CLI Command executed in an existing component container",
														MarkdownDescription: "CLI Command executed in an existing component container",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"command_line": {
																Description:         "The actual command-line string  Special variables that can be used:   - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping.   - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
																MarkdownDescription: "The actual command-line string  Special variables that can be used:   - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping.   - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"component": {
																Description:         "Describes component to which given action relates",
																MarkdownDescription: "Describes component to which given action relates",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"env": {
																Description:         "Optional list of environment variables that have to be set before running the command",
																MarkdownDescription: "Optional list of environment variables that have to be set before running the command",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"group": {
																Description:         "Defines the group this command is part of",
																MarkdownDescription: "Defines the group this command is part of",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"is_default": {
																		Description:         "Identifies the default command for a given group kind",
																		MarkdownDescription: "Identifies the default command for a given group kind",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"kind": {
																		Description:         "Kind of group the command is part of",
																		MarkdownDescription: "Kind of group the command is part of",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.OneOf("build", "run", "test", "debug", "deploy"),
																		},
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"hot_reload_capable": {
																Description:         "Whether the command is capable to reload itself when source code changes. If set to 'true' the command won't be restarted and it is expected to handle file changes on its own.  Default value is 'false'",
																MarkdownDescription: "Whether the command is capable to reload itself when source code changes. If set to 'true' the command won't be restarted and it is expected to handle file changes on its own.  Default value is 'false'",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"label": {
																Description:         "Optional label that provides a label for this command to be used in Editor UI menus for example",
																MarkdownDescription: "Optional label that provides a label for this command to be used in Editor UI menus for example",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"working_dir": {
																Description:         "Working directory where the command should be executed  Special variables that can be used:   - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping.   - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
																MarkdownDescription: "Working directory where the command should be executed  Special variables that can be used:   - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping.   - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"id": {
														Description:         "Mandatory identifier that allows referencing this command in composite commands, from a parent, or in events.",
														MarkdownDescription: "Mandatory identifier that allows referencing this command in composite commands, from a parent, or in events.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtMost(63),

															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"components": {
												Description:         "Overrides of components encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",
												MarkdownDescription: "Overrides of components encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"attributes": {
														Description:         "Map of implementation-dependant free-form YAML attributes.",
														MarkdownDescription: "Map of implementation-dependant free-form YAML attributes.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"component_type": {
														Description:         "Type of component",
														MarkdownDescription: "Type of component",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("Container", "Kubernetes", "Openshift", "Volume", "Image"),
														},
													},

													"container": {
														Description:         "Allows adding and configuring devworkspace-related containers",
														MarkdownDescription: "Allows adding and configuring devworkspace-related containers",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"annotation": {
																Description:         "Annotations that should be added to specific resources for this container",
																MarkdownDescription: "Annotations that should be added to specific resources for this container",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"deployment": {
																		Description:         "Annotations to be added to deployment",
																		MarkdownDescription: "Annotations to be added to deployment",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"service": {
																		Description:         "Annotations to be added to service",
																		MarkdownDescription: "Annotations to be added to service",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"args": {
																Description:         "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command.  Defaults to an empty array, meaning use whatever is defined in the image.",
																MarkdownDescription: "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command.  Defaults to an empty array, meaning use whatever is defined in the image.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"command": {
																Description:         "The command to run in the dockerimage component instead of the default one provided in the image.  Defaults to an empty array, meaning use whatever is defined in the image.",
																MarkdownDescription: "The command to run in the dockerimage component instead of the default one provided in the image.  Defaults to an empty array, meaning use whatever is defined in the image.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"cpu_limit": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"cpu_request": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"dedicated_pod": {
																Description:         "Specify if a container should run in its own separated pod, instead of running as part of the main development environment pod.  Default value is 'false'",
																MarkdownDescription: "Specify if a container should run in its own separated pod, instead of running as part of the main development environment pod.  Default value is 'false'",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoints": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"annotation": {
																		Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
																		MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"attributes": {
																		Description:         "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
																		MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",

																		Type: utilities.DynamicType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"exposure": {
																		Description:         "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
																		MarkdownDescription: "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.OneOf("public", "internal", "none"),
																		},
																	},

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.LengthAtMost(63),

																			stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																		},
																	},

																	"path": {
																		Description:         "Path of the endpoint URL",
																		MarkdownDescription: "Path of the endpoint URL",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"protocol": {
																		Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
																		MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																		},
																	},

																	"secure": {
																		Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																		MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"target_port": {
																		Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
																		MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"env": {
																Description:         "Environment variables used in this container.  The following variables are reserved and cannot be overridden via env:   - '$PROJECTS_ROOT'   - '$PROJECT_SOURCE'",
																MarkdownDescription: "Environment variables used in this container.  The following variables are reserved and cannot be overridden via env:   - '$PROJECTS_ROOT'   - '$PROJECT_SOURCE'",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"image": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"memory_limit": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"memory_request": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"mount_sources": {
																Description:         "Toggles whether or not the project source code should be mounted in the component.  Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",
																MarkdownDescription: "Toggles whether or not the project source code should be mounted in the component.  Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"source_mapping": {
																Description:         "Optional specification of the path in the container where project sources should be transferred/mounted when 'mountSources' is 'true'. When omitted, the default value of /projects is used.",
																MarkdownDescription: "Optional specification of the path in the container where project sources should be transferred/mounted when 'mountSources' is 'true'. When omitted, the default value of /projects is used.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"volume_mounts": {
																Description:         "List of volumes mounts that should be mounted is this container.",
																MarkdownDescription: "List of volumes mounts that should be mounted is this container.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "The volume mount name is the name of an existing 'Volume' component. If several containers mount the same volume name then they will reuse the same volume and will be able to access to the same files.",
																		MarkdownDescription: "The volume mount name is the name of an existing 'Volume' component. If several containers mount the same volume name then they will reuse the same volume and will be able to access to the same files.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.LengthAtMost(63),

																			stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																		},
																	},

																	"path": {
																		Description:         "The path in the component container where the volume should be mounted. If not path is mentioned, default path is the is '/<name>'.",
																		MarkdownDescription: "The path in the component container where the volume should be mounted. If not path is mentioned, default path is the is '/<name>'.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"image": {
														Description:         "Allows specifying the definition of an image for outer loop builds",
														MarkdownDescription: "Allows specifying the definition of an image for outer loop builds",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"auto_build": {
																Description:         "Defines if the image should be built during startup.  Default value is 'false'",
																MarkdownDescription: "Defines if the image should be built during startup.  Default value is 'false'",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"dockerfile": {
																Description:         "Allows specifying dockerfile type build",
																MarkdownDescription: "Allows specifying dockerfile type build",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"args": {
																		Description:         "The arguments to supply to the dockerfile build.",
																		MarkdownDescription: "The arguments to supply to the dockerfile build.",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"build_context": {
																		Description:         "Path of source directory to establish build context. Defaults to ${PROJECT_SOURCE} in the container",
																		MarkdownDescription: "Path of source directory to establish build context. Defaults to ${PROJECT_SOURCE} in the container",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"devfile_registry": {
																		Description:         "Dockerfile's Devfile Registry source",
																		MarkdownDescription: "Dockerfile's Devfile Registry source",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"id": {
																				Description:         "Id in a devfile registry that contains a Dockerfile. The src in the OCI registry required for the Dockerfile build will be downloaded for building the image.",
																				MarkdownDescription: "Id in a devfile registry that contains a Dockerfile. The src in the OCI registry required for the Dockerfile build will be downloaded for building the image.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"registry_url": {
																				Description:         "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src. To ensure the Dockerfile gets resolved consistently in different environments, it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",
																				MarkdownDescription: "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src. To ensure the Dockerfile gets resolved consistently in different environments, it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"git": {
																		Description:         "Dockerfile's Git source",
																		MarkdownDescription: "Dockerfile's Git source",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"checkout_from": {
																				Description:         "Defines from what the project should be checked out. Required if there are more than one remote configured",
																				MarkdownDescription: "Defines from what the project should be checked out. Required if there are more than one remote configured",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"remote": {
																						Description:         "The remote name should be used as init. Required if there are more than one remote configured",
																						MarkdownDescription: "The remote name should be used as init. Required if there are more than one remote configured",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"revision": {
																						Description:         "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
																						MarkdownDescription: "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"file_location": {
																				Description:         "Location of the Dockerfile in the Git repository when using git as Dockerfile src. Defaults to Dockerfile.",
																				MarkdownDescription: "Location of the Dockerfile in the Git repository when using git as Dockerfile src. Defaults to Dockerfile.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"remotes": {
																				Description:         "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
																				MarkdownDescription: "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"root_required": {
																		Description:         "Specify if a privileged builder pod is required.  Default value is 'false'",
																		MarkdownDescription: "Specify if a privileged builder pod is required.  Default value is 'false'",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"src_type": {
																		Description:         "Type of Dockerfile src",
																		MarkdownDescription: "Type of Dockerfile src",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.OneOf("Uri", "DevfileRegistry", "Git"),
																		},
																	},

																	"uri": {
																		Description:         "URI Reference of a Dockerfile. It can be a full URL or a relative URI from the current devfile as the base URI.",
																		MarkdownDescription: "URI Reference of a Dockerfile. It can be a full URL or a relative URI from the current devfile as the base URI.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"image_name": {
																Description:         "Name of the image for the resulting outerloop build",
																MarkdownDescription: "Name of the image for the resulting outerloop build",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"image_type": {
																Description:         "Type of image",
																MarkdownDescription: "Type of image",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.OneOf("Dockerfile", "AutoBuild"),
																},
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"kubernetes": {
														Description:         "Allows importing into the devworkspace the Kubernetes resources defined in a given manifest. For example this allows reusing the Kubernetes definitions used to deploy some runtime components in production.",
														MarkdownDescription: "Allows importing into the devworkspace the Kubernetes resources defined in a given manifest. For example this allows reusing the Kubernetes definitions used to deploy some runtime components in production.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"deploy_by_default": {
																Description:         "Defines if the component should be deployed during startup.  Default value is 'false'",
																MarkdownDescription: "Defines if the component should be deployed during startup.  Default value is 'false'",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoints": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"annotation": {
																		Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
																		MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"attributes": {
																		Description:         "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
																		MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",

																		Type: utilities.DynamicType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"exposure": {
																		Description:         "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
																		MarkdownDescription: "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.OneOf("public", "internal", "none"),
																		},
																	},

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.LengthAtMost(63),

																			stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																		},
																	},

																	"path": {
																		Description:         "Path of the endpoint URL",
																		MarkdownDescription: "Path of the endpoint URL",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"protocol": {
																		Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
																		MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																		},
																	},

																	"secure": {
																		Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																		MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"target_port": {
																		Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
																		MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"inlined": {
																Description:         "Inlined manifest",
																MarkdownDescription: "Inlined manifest",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"location_type": {
																Description:         "Type of Kubernetes-like location",
																MarkdownDescription: "Type of Kubernetes-like location",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.OneOf("Uri", "Inlined"),
																},
															},

															"uri": {
																Description:         "Location in a file fetched from a uri.",
																MarkdownDescription: "Location in a file fetched from a uri.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "Mandatory name that allows referencing the component from other elements (such as commands) or from an external devfile that may reference this component through a parent or a plugin.",
														MarkdownDescription: "Mandatory name that allows referencing the component from other elements (such as commands) or from an external devfile that may reference this component through a parent or a plugin.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtMost(63),

															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
													},

													"openshift": {
														Description:         "Allows importing into the devworkspace the OpenShift resources defined in a given manifest. For example this allows reusing the OpenShift definitions used to deploy some runtime components in production.",
														MarkdownDescription: "Allows importing into the devworkspace the OpenShift resources defined in a given manifest. For example this allows reusing the OpenShift definitions used to deploy some runtime components in production.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"deploy_by_default": {
																Description:         "Defines if the component should be deployed during startup.  Default value is 'false'",
																MarkdownDescription: "Defines if the component should be deployed during startup.  Default value is 'false'",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoints": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"annotation": {
																		Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
																		MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"attributes": {
																		Description:         "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
																		MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",

																		Type: utilities.DynamicType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"exposure": {
																		Description:         "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
																		MarkdownDescription: "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.OneOf("public", "internal", "none"),
																		},
																	},

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.LengthAtMost(63),

																			stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																		},
																	},

																	"path": {
																		Description:         "Path of the endpoint URL",
																		MarkdownDescription: "Path of the endpoint URL",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"protocol": {
																		Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
																		MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																		},
																	},

																	"secure": {
																		Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																		MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"target_port": {
																		Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
																		MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"inlined": {
																Description:         "Inlined manifest",
																MarkdownDescription: "Inlined manifest",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"location_type": {
																Description:         "Type of Kubernetes-like location",
																MarkdownDescription: "Type of Kubernetes-like location",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.OneOf("Uri", "Inlined"),
																},
															},

															"uri": {
																Description:         "Location in a file fetched from a uri.",
																MarkdownDescription: "Location in a file fetched from a uri.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"volume": {
														Description:         "Allows specifying the definition of a volume shared by several other components",
														MarkdownDescription: "Allows specifying the definition of a volume shared by several other components",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"ephemeral": {
																Description:         "Ephemeral volumes are not stored persistently across restarts. Defaults to false",
																MarkdownDescription: "Ephemeral volumes are not stored persistently across restarts. Defaults to false",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"size": {
																Description:         "Size of the volume",
																MarkdownDescription: "Size of the volume",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"id": {
												Description:         "Id in a registry that contains a Devfile yaml file",
												MarkdownDescription: "Id in a registry that contains a Devfile yaml file",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"import_reference_type": {
												Description:         "type of location from where the referenced template structure should be retrieved",
												MarkdownDescription: "type of location from where the referenced template structure should be retrieved",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("Uri", "Id", "Kubernetes"),
												},
											},

											"kubernetes": {
												Description:         "Reference to a Kubernetes CRD of type DevWorkspaceTemplate",
												MarkdownDescription: "Reference to a Kubernetes CRD of type DevWorkspaceTemplate",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"namespace": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"registry_url": {
												Description:         "Registry URL to pull the parent devfile from when using id in the parent reference. To ensure the parent devfile gets resolved consistently in different environments, it is recommended to always specify the 'registryUrl' when 'id' is used.",
												MarkdownDescription: "Registry URL to pull the parent devfile from when using id in the parent reference. To ensure the parent devfile gets resolved consistently in different environments, it is recommended to always specify the 'registryUrl' when 'id' is used.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"uri": {
												Description:         "URI Reference of a parent devfile YAML file. It can be a full URL or a relative URI with the current devfile as the base URI.",
												MarkdownDescription: "URI Reference of a parent devfile YAML file. It can be a full URL or a relative URI with the current devfile as the base URI.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"version": {
												Description:         "Specific stack/sample version to pull the parent devfile from, when using id in the parent reference. To specify 'version', 'id' must be defined and used as the import reference source. 'version' can be either a specific stack version, or 'latest'. If no 'version' specified, default version will be used.",
												MarkdownDescription: "Specific stack/sample version to pull the parent devfile from, when using id in the parent reference. To specify 'version', 'id' must be defined and used as the import reference source. 'version' can be either a specific stack version, or 'latest'. If no 'version' specified, default version will be used.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^(latest)|(([1-9])\.([0-9]+)\.([0-9]+)(\-[0-9a-z-]+(\.[0-9a-z-]+)*)?(\+[0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*)?)$`), ""),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"volume": {
										Description:         "Allows specifying the definition of a volume shared by several other components",
										MarkdownDescription: "Allows specifying the definition of a volume shared by several other components",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ephemeral": {
												Description:         "Ephemeral volumes are not stored persistently across restarts. Defaults to false",
												MarkdownDescription: "Ephemeral volumes are not stored persistently across restarts. Defaults to false",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"size": {
												Description:         "Size of the volume",
												MarkdownDescription: "Size of the volume",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"workspace_default_editor": {
								Description:         "The default editor to workspace create with. It could be a plugin ID or a URI. The plugin ID must have 'publisher/plugin/version'. The URI must start from 'http'.",
								MarkdownDescription: "The default editor to workspace create with. It could be a plugin ID or a URI. The plugin ID must have 'publisher/plugin/version'. The URI must start from 'http'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"workspace_namespace_default": {
								Description:         "Defines Kubernetes default namespace in which user's workspaces are created for a case when a user does not override it. It's possible to use '<username>', '<userid>' and '<workspaceid>' placeholders, such as che-workspace-<username>. In that case, a new namespace will be created for each user or workspace.",
								MarkdownDescription: "Defines Kubernetes default namespace in which user's workspaces are created for a case when a user does not override it. It's possible to use '<username>', '<userid>' and '<workspaceid>' placeholders, such as che-workspace-<username>. In that case, a new namespace will be created for each user or workspace.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"workspace_pod_node_selector": {
								Description:         "The node selector that limits the nodes that can run the workspace pods.",
								MarkdownDescription: "The node selector that limits the nodes that can run the workspace pods.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"workspace_pod_tolerations": {
								Description:         "The pod tolerations put on the workspace pods to limit where the workspace pods can run.",
								MarkdownDescription: "The pod tolerations put on the workspace pods to limit where the workspace pods can run.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"effect": {
										Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
										MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"key": {
										Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
										MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"operator": {
										Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
										MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"toleration_seconds": {
										Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
										MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": {
										Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
										MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"workspaces_default_plugins": {
								Description:         "Default plug-ins applied to Devworkspaces.",
								MarkdownDescription: "Default plug-ins applied to Devworkspaces.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"editor": {
										Description:         "The editor id to specify default plug-ins for.",
										MarkdownDescription: "The editor id to specify default plug-ins for.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"plugins": {
										Description:         "Default plug-in uris for the specified editor.",
										MarkdownDescription: "Default plug-in uris for the specified editor.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"storage": {
						Description:         "Configuration settings related to the persistent storage used by the Che installation.",
						MarkdownDescription: "Configuration settings related to the persistent storage used by the Che installation.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"per_workspace_strategy_pvc_storage_class_name": {
								Description:         "Storage class for the Persistent Volume Claims dedicated to the Che workspaces. When omitted or left blank, a default storage class is used.",
								MarkdownDescription: "Storage class for the Persistent Volume Claims dedicated to the Che workspaces. When omitted or left blank, a default storage class is used.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"per_workspace_strategy_pvc_claim_size": {
								Description:         "Size of the persistent volume claim for workspaces.",
								MarkdownDescription: "Size of the persistent volume claim for workspaces.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"postgres_pvc_storage_class_name": {
								Description:         "Storage class for the Persistent Volume Claim dedicated to the PostgreSQL database. When omitted or left blank, a default storage class is used.",
								MarkdownDescription: "Storage class for the Persistent Volume Claim dedicated to the PostgreSQL database. When omitted or left blank, a default storage class is used.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pre_create_sub_paths": {
								Description:         "Instructs the Che server to start a special Pod to pre-create a sub-path in the Persistent Volumes. Defaults to 'false', however it will need to enable it according to the configuration of your Kubernetes cluster.",
								MarkdownDescription: "Instructs the Che server to start a special Pod to pre-create a sub-path in the Persistent Volumes. Defaults to 'false', however it will need to enable it according to the configuration of your Kubernetes cluster.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pvc_claim_size": {
								Description:         "Size of the persistent volume claim for workspaces. Defaults to '10Gi'.",
								MarkdownDescription: "Size of the persistent volume claim for workspaces. Defaults to '10Gi'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pvc_jobs_image": {
								Description:         "Overrides the container image used to create sub-paths in the Persistent Volumes. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator. See also the 'preCreateSubPaths' field.",
								MarkdownDescription: "Overrides the container image used to create sub-paths in the Persistent Volumes. This includes the image tag. Omit it or leave it empty to use the default container image provided by the Operator. See also the 'preCreateSubPaths' field.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pvc_strategy": {
								Description:         "Persistent volume claim strategy for the Che server. This Can be:'common' (all workspaces PVCs in one volume), 'per-workspace' (one PVC per workspace for all declared volumes) and 'unique' (one PVC per declared volume). Defaults to 'common'.",
								MarkdownDescription: "Persistent volume claim strategy for the Che server. This Can be:'common' (all workspaces PVCs in one volume), 'per-workspace' (one PVC per workspace for all declared volumes) and 'unique' (one PVC per declared volume). Defaults to 'common'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"workspace_pvc_storage_class_name": {
								Description:         "Storage class for the Persistent Volume Claims dedicated to the Che workspaces. When omitted or left blank, a default storage class is used.",
								MarkdownDescription: "Storage class for the Persistent Volume Claims dedicated to the Che workspaces. When omitted or left blank, a default storage class is used.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *OrgEclipseCheCheClusterV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_org_eclipse_che_che_cluster_v1")

	var state OrgEclipseCheCheClusterV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel OrgEclipseCheCheClusterV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("org.eclipse.che/v1")
	goModel.Kind = utilities.Ptr("CheCluster")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *OrgEclipseCheCheClusterV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_org_eclipse_che_che_cluster_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *OrgEclipseCheCheClusterV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_org_eclipse_che_che_cluster_v1")

	var state OrgEclipseCheCheClusterV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel OrgEclipseCheCheClusterV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("org.eclipse.che/v1")
	goModel.Kind = utilities.Ptr("CheCluster")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *OrgEclipseCheCheClusterV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_org_eclipse_che_che_cluster_v1")
	// NO-OP: Terraform removes the state automatically for us
}
