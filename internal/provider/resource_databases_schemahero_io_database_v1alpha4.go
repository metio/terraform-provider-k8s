/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

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

type DatabasesSchemaheroIoDatabaseV1Alpha4Resource struct{}

var (
	_ resource.Resource = (*DatabasesSchemaheroIoDatabaseV1Alpha4Resource)(nil)
)

type DatabasesSchemaheroIoDatabaseV1Alpha4TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type DatabasesSchemaheroIoDatabaseV1Alpha4GoModel struct {
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
		Connection *struct {
			Cassandra *struct {
				Hosts *[]string `tfsdk:"hosts" yaml:"hosts,omitempty"`

				Keyspace *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"keyspace" yaml:"keyspace,omitempty"`

				Password *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"password" yaml:"password,omitempty"`

				Username *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"username" yaml:"username,omitempty"`
			} `tfsdk:"cassandra" yaml:"cassandra,omitempty"`

			Cockroachdb *struct {
				Dbname *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"dbname" yaml:"dbname,omitempty"`

				Host *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"host" yaml:"host,omitempty"`

				Password *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"password" yaml:"password,omitempty"`

				Port *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"port" yaml:"port,omitempty"`

				Schema *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"schema" yaml:"schema,omitempty"`

				Sslmode *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"sslmode" yaml:"sslmode,omitempty"`

				Uri *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"uri" yaml:"uri,omitempty"`

				User *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"user" yaml:"user,omitempty"`
			} `tfsdk:"cockroachdb" yaml:"cockroachdb,omitempty"`

			Mysql *struct {
				Collation *string `tfsdk:"collation" yaml:"collation,omitempty"`

				Dbname *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"dbname" yaml:"dbname,omitempty"`

				DefaultCharset *string `tfsdk:"default_charset" yaml:"defaultCharset,omitempty"`

				DisableTLS *bool `tfsdk:"disable_tls" yaml:"disableTLS,omitempty"`

				Host *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"host" yaml:"host,omitempty"`

				Password *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"password" yaml:"password,omitempty"`

				Port *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"port" yaml:"port,omitempty"`

				Uri *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"uri" yaml:"uri,omitempty"`

				User *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"user" yaml:"user,omitempty"`
			} `tfsdk:"mysql" yaml:"mysql,omitempty"`

			Postgres *struct {
				Dbname *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"dbname" yaml:"dbname,omitempty"`

				Host *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"host" yaml:"host,omitempty"`

				Password *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"password" yaml:"password,omitempty"`

				Port *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"port" yaml:"port,omitempty"`

				Schema *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"schema" yaml:"schema,omitempty"`

				Sslmode *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"sslmode" yaml:"sslmode,omitempty"`

				Uri *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"uri" yaml:"uri,omitempty"`

				User *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"user" yaml:"user,omitempty"`
			} `tfsdk:"postgres" yaml:"postgres,omitempty"`

			Rqlite *struct {
				DisableTLS *bool `tfsdk:"disable_tls" yaml:"disableTLS,omitempty"`

				Host *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"host" yaml:"host,omitempty"`

				Password *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"password" yaml:"password,omitempty"`

				Port *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"port" yaml:"port,omitempty"`

				Uri *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"uri" yaml:"uri,omitempty"`

				User *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						Ssm *struct {
							AccessKeyId *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"access_key_id" yaml:"accessKeyId,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Region *string `tfsdk:"region" yaml:"region,omitempty"`

							SecretAccessKey *struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"secret_access_key" yaml:"secretAccessKey,omitempty"`

							WithDecryption *bool `tfsdk:"with_decryption" yaml:"withDecryption,omitempty"`
						} `tfsdk:"ssm" yaml:"ssm,omitempty"`

						Vault *struct {
							AgentInject *bool `tfsdk:"agent_inject" yaml:"agentInject,omitempty"`

							ConnectionTemplate *string `tfsdk:"connection_template" yaml:"connectionTemplate,omitempty"`

							Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

							KubernetesAuthEndpoint *string `tfsdk:"kubernetes_auth_endpoint" yaml:"kubernetesAuthEndpoint,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`

							ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

							ServiceAccountNamespace *string `tfsdk:"service_account_namespace" yaml:"serviceAccountNamespace,omitempty"`
						} `tfsdk:"vault" yaml:"vault,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"user" yaml:"user,omitempty"`
			} `tfsdk:"rqlite" yaml:"rqlite,omitempty"`

			Sqlite *struct {
				Dsn *string `tfsdk:"dsn" yaml:"dsn,omitempty"`
			} `tfsdk:"sqlite" yaml:"sqlite,omitempty"`
		} `tfsdk:"connection" yaml:"connection,omitempty"`

		DeploySeedData *bool `tfsdk:"deploy_seed_data" yaml:"deploySeedData,omitempty"`

		EnableShellCommand *bool `tfsdk:"enable_shell_command" yaml:"enableShellCommand,omitempty"`

		ImmediateDeploy *bool `tfsdk:"immediate_deploy" yaml:"immediateDeploy,omitempty"`

		Schemahero *struct {
			Image *string `tfsdk:"image" yaml:"image,omitempty"`

			NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`
		} `tfsdk:"schemahero" yaml:"schemahero,omitempty"`

		Template *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

				Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"metadata" yaml:"metadata,omitempty"`
		} `tfsdk:"template" yaml:"template,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewDatabasesSchemaheroIoDatabaseV1Alpha4Resource() resource.Resource {
	return &DatabasesSchemaheroIoDatabaseV1Alpha4Resource{}
}

func (r *DatabasesSchemaheroIoDatabaseV1Alpha4Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_databases_schemahero_io_database_v1alpha4"
}

func (r *DatabasesSchemaheroIoDatabaseV1Alpha4Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Database is the Schema for the databases API",
		MarkdownDescription: "Database is the Schema for the databases API",
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
				Description:         "",
				MarkdownDescription: "",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"connection": {
						Description:         "DatabaseConnection defines connection parameters for the database driver",
						MarkdownDescription: "DatabaseConnection defines connection parameters for the database driver",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cassandra": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"hosts": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: true,
										Optional: false,
										Computed: false,
									},

									"keyspace": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"password": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"username": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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

							"cockroachdb": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"dbname": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"host": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"password": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"schema": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sslmode": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"uri": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"user": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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

							"mysql": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"collation": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"dbname": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"default_charset": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"disable_tls": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"host": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"password": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"uri": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"user": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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

							"postgres": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"dbname": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"host": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"password": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"schema": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sslmode": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"uri": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"user": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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

							"rqlite": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"disable_tls": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"host": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"password": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"uri": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"user": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"secret_key_ref": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"ssm": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"access_key_id": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"region": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_access_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"value": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"secret_key_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

															"with_decryption": {
																Description:         "",
																MarkdownDescription: "",

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

													"vault": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"agent_inject": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"connection_template": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes_auth_endpoint": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service_account": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"service_account_namespace": {
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

							"sqlite": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"dsn": {
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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"deploy_seed_data": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_shell_command": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"immediate_deploy": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"schemahero": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"image": {
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

					"template": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"metadata": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"finalizers": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"labels": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *DatabasesSchemaheroIoDatabaseV1Alpha4Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_databases_schemahero_io_database_v1alpha4")

	var state DatabasesSchemaheroIoDatabaseV1Alpha4TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel DatabasesSchemaheroIoDatabaseV1Alpha4GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("databases.schemahero.io/v1alpha4")
	goModel.Kind = utilities.Ptr("Database")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *DatabasesSchemaheroIoDatabaseV1Alpha4Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_databases_schemahero_io_database_v1alpha4")
	// NO-OP: All data is already in Terraform state
}

func (r *DatabasesSchemaheroIoDatabaseV1Alpha4Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_databases_schemahero_io_database_v1alpha4")

	var state DatabasesSchemaheroIoDatabaseV1Alpha4TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel DatabasesSchemaheroIoDatabaseV1Alpha4GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("databases.schemahero.io/v1alpha4")
	goModel.Kind = utilities.Ptr("Database")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *DatabasesSchemaheroIoDatabaseV1Alpha4Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_databases_schemahero_io_database_v1alpha4")
	// NO-OP: Terraform removes the state automatically for us
}
