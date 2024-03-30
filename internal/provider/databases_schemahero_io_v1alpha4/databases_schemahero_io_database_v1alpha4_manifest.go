/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package databases_schemahero_io_v1alpha4

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
	_ datasource.DataSource = &DatabasesSchemaheroIoDatabaseV1Alpha4Manifest{}
)

func NewDatabasesSchemaheroIoDatabaseV1Alpha4Manifest() datasource.DataSource {
	return &DatabasesSchemaheroIoDatabaseV1Alpha4Manifest{}
}

type DatabasesSchemaheroIoDatabaseV1Alpha4Manifest struct{}

type DatabasesSchemaheroIoDatabaseV1Alpha4ManifestData struct {
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
		Connection *struct {
			Cassandra *struct {
				Hosts    *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
				Keyspace *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"keyspace" json:"keyspace,omitempty"`
				Password *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Username *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"cassandra" json:"cassandra,omitempty"`
			Cockroachdb *struct {
				Dbname *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"dbname" json:"dbname,omitempty"`
				Host *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"host" json:"host,omitempty"`
				Password *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Port *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"port" json:"port,omitempty"`
				Schema *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"schema" json:"schema,omitempty"`
				Sslmode *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"sslmode" json:"sslmode,omitempty"`
				Uri *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"uri" json:"uri,omitempty"`
				User *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"user" json:"user,omitempty"`
			} `tfsdk:"cockroachdb" json:"cockroachdb,omitempty"`
			Mysql *struct {
				Collation *string `tfsdk:"collation" json:"collation,omitempty"`
				Dbname    *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"dbname" json:"dbname,omitempty"`
				DefaultCharset *string `tfsdk:"default_charset" json:"defaultCharset,omitempty"`
				DisableTLS     *bool   `tfsdk:"disable_tls" json:"disableTLS,omitempty"`
				Host           *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"host" json:"host,omitempty"`
				Password *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Port *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"port" json:"port,omitempty"`
				Uri *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"uri" json:"uri,omitempty"`
				User *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"user" json:"user,omitempty"`
			} `tfsdk:"mysql" json:"mysql,omitempty"`
			Postgres *struct {
				Dbname *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"dbname" json:"dbname,omitempty"`
				Host *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"host" json:"host,omitempty"`
				Password *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Port *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"port" json:"port,omitempty"`
				Schema *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"schema" json:"schema,omitempty"`
				Sslmode *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"sslmode" json:"sslmode,omitempty"`
				Uri *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"uri" json:"uri,omitempty"`
				User *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"user" json:"user,omitempty"`
			} `tfsdk:"postgres" json:"postgres,omitempty"`
			Rqlite *struct {
				DisableTLS *bool `tfsdk:"disable_tls" json:"disableTLS,omitempty"`
				Host       *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"host" json:"host,omitempty"`
				Password *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Port *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"port" json:"port,omitempty"`
				Uri *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"uri" json:"uri,omitempty"`
				User *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"user" json:"user,omitempty"`
			} `tfsdk:"rqlite" json:"rqlite,omitempty"`
			Sqlite *struct {
				Dsn *string `tfsdk:"dsn" json:"dsn,omitempty"`
			} `tfsdk:"sqlite" json:"sqlite,omitempty"`
			Timescaledb *struct {
				Dbname *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"dbname" json:"dbname,omitempty"`
				Host *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"host" json:"host,omitempty"`
				Password *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Port *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"port" json:"port,omitempty"`
				Schema *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"schema" json:"schema,omitempty"`
				Sslmode *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"sslmode" json:"sslmode,omitempty"`
				Uri *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"uri" json:"uri,omitempty"`
				User *struct {
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						Ssm *struct {
							AccessKeyId *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Region          *string `tfsdk:"region" json:"region,omitempty"`
							SecretAccessKey *struct {
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									SecretKeyRef *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
							WithDecryption *bool `tfsdk:"with_decryption" json:"withDecryption,omitempty"`
						} `tfsdk:"ssm" json:"ssm,omitempty"`
						Vault *struct {
							AgentInject             *bool   `tfsdk:"agent_inject" json:"agentInject,omitempty"`
							ConnectionTemplate      *string `tfsdk:"connection_template" json:"connectionTemplate,omitempty"`
							Endpoint                *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
							KubernetesAuthEndpoint  *string `tfsdk:"kubernetes_auth_endpoint" json:"kubernetesAuthEndpoint,omitempty"`
							Role                    *string `tfsdk:"role" json:"role,omitempty"`
							Secret                  *string `tfsdk:"secret" json:"secret,omitempty"`
							ServiceAccount          *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" json:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" json:"vault,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"user" json:"user,omitempty"`
			} `tfsdk:"timescaledb" json:"timescaledb,omitempty"`
		} `tfsdk:"connection" json:"connection,omitempty"`
		DeploySeedData     *bool `tfsdk:"deploy_seed_data" json:"deploySeedData,omitempty"`
		EnableShellCommand *bool `tfsdk:"enable_shell_command" json:"enableShellCommand,omitempty"`
		ImmediateDeploy    *bool `tfsdk:"immediate_deploy" json:"immediateDeploy,omitempty"`
		Schemahero         *struct {
			Image        *string            `tfsdk:"image" json:"image,omitempty"`
			NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			Tolerations  *[]struct {
				Effect   *string `tfsdk:"effect" json:"effect,omitempty"`
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Operator *string `tfsdk:"operator" json:"operator,omitempty"`
				Value    *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
		} `tfsdk:"schemahero" json:"schemahero,omitempty"`
		Template *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Finalizers  *[]string          `tfsdk:"finalizers" json:"finalizers,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
				Namespace   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DatabasesSchemaheroIoDatabaseV1Alpha4Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_databases_schemahero_io_database_v1alpha4_manifest"
}

func (r *DatabasesSchemaheroIoDatabaseV1Alpha4Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Database is the Schema for the databases API",
		MarkdownDescription: "Database is the Schema for the databases API",
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
					"connection": schema.SingleNestedAttribute{
						Description:         "DatabaseConnection defines connection parameters for the database driver",
						MarkdownDescription: "DatabaseConnection defines connection parameters for the database driver",
						Attributes: map[string]schema.Attribute{
							"cassandra": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"hosts": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"keyspace": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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
										Required: true,
										Optional: false,
										Computed: false,
									},

									"password": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"username": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

							"cockroachdb": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"dbname": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"host": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"password": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"port": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"schema": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"sslmode": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"uri": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"user": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

							"mysql": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"collation": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dbname": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"default_charset": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"disable_tls": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"host": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"password": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"port": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"uri": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"user": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

							"postgres": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"dbname": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"host": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"password": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"port": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"schema": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"sslmode": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"uri": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"user": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

							"rqlite": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"disable_tls": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"host": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"password": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"port": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"uri": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"user": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

							"sqlite": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"dsn": schema.StringAttribute{
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

							"timescaledb": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"dbname": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"host": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"password": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"port": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"schema": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"sslmode": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"uri": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

									"user": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ssm": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"access_key_id": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

															"with_decryption": schema.BoolAttribute{
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

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service_account_namespace": schema.StringAttribute{
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

					"deploy_seed_data": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_shell_command": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"immediate_deploy": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"schemahero": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"image": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_selector": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"template": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"finalizers": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
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

func (r *DatabasesSchemaheroIoDatabaseV1Alpha4Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_databases_schemahero_io_database_v1alpha4_manifest")

	var model DatabasesSchemaheroIoDatabaseV1Alpha4ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("databases.schemahero.io/v1alpha4")
	model.Kind = pointer.String("Database")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
