/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package databases_schemahero_io_v1alpha4

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &DatabasesSchemaheroIoDatabaseV1Alpha4DataSource{}
	_ datasource.DataSourceWithConfigure = &DatabasesSchemaheroIoDatabaseV1Alpha4DataSource{}
)

func NewDatabasesSchemaheroIoDatabaseV1Alpha4DataSource() datasource.DataSource {
	return &DatabasesSchemaheroIoDatabaseV1Alpha4DataSource{}
}

type DatabasesSchemaheroIoDatabaseV1Alpha4DataSource struct {
	kubernetesClient dynamic.Interface
}

type DatabasesSchemaheroIoDatabaseV1Alpha4DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *DatabasesSchemaheroIoDatabaseV1Alpha4DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_databases_schemahero_io_database_v1alpha4"
}

func (r *DatabasesSchemaheroIoDatabaseV1Alpha4DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Database is the Schema for the databases API",
		MarkdownDescription: "Database is the Schema for the databases API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
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
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"keyspace": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"password": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"username": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"host": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"password": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"port": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"schema": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"sslmode": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"uri": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"user": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"mysql": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"collation": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"dbname": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"default_charset": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"disable_tls": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"host": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"password": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"port": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"uri": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"user": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"host": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"password": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"port": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"schema": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"sslmode": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"uri": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"user": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"rqlite": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"disable_tls": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"host": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"password": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"port": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"uri": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"user": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"sqlite": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"dsn": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"host": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"password": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"port": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"schema": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"sslmode": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"uri": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"user": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"value": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"region": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_access_key": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"with_decryption": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"vault": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"agent_inject": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"connection_template": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"kubernetes_auth_endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"role": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"service_account_namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"deploy_seed_data": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_shell_command": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"immediate_deploy": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"schemahero": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"image": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"node_selector": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
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
										Optional:            false,
										Computed:            true,
									},

									"finalizers": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"labels": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"namespace": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *DatabasesSchemaheroIoDatabaseV1Alpha4DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *DatabasesSchemaheroIoDatabaseV1Alpha4DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_databases_schemahero_io_database_v1alpha4")

	var data DatabasesSchemaheroIoDatabaseV1Alpha4DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "databases.schemahero.io", Version: "v1alpha4", Resource: "databases"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse DatabasesSchemaheroIoDatabaseV1Alpha4DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("databases.schemahero.io/v1alpha4")
	data.Kind = pointer.String("Database")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
