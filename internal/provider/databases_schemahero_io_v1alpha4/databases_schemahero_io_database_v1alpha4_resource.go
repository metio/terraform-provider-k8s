/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package databases_schemahero_io_v1alpha4

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"strings"
)

var (
	_ resource.Resource                = &DatabasesSchemaheroIoDatabaseV1Alpha4Resource{}
	_ resource.ResourceWithConfigure   = &DatabasesSchemaheroIoDatabaseV1Alpha4Resource{}
	_ resource.ResourceWithImportState = &DatabasesSchemaheroIoDatabaseV1Alpha4Resource{}
)

func NewDatabasesSchemaheroIoDatabaseV1Alpha4Resource() resource.Resource {
	return &DatabasesSchemaheroIoDatabaseV1Alpha4Resource{}
}

type DatabasesSchemaheroIoDatabaseV1Alpha4Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type DatabasesSchemaheroIoDatabaseV1Alpha4ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

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

func (r *DatabasesSchemaheroIoDatabaseV1Alpha4Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_databases_schemahero_io_database_v1alpha4"
}

func (r *DatabasesSchemaheroIoDatabaseV1Alpha4Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
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

func (r *DatabasesSchemaheroIoDatabaseV1Alpha4Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *DatabasesSchemaheroIoDatabaseV1Alpha4Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_databases_schemahero_io_database_v1alpha4")

	var model DatabasesSchemaheroIoDatabaseV1Alpha4ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("databases.schemahero.io/v1alpha4")
	model.Kind = pointer.String("Database")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "databases.schemahero.io", Version: "v1alpha4", Resource: "databases"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse DatabasesSchemaheroIoDatabaseV1Alpha4ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *DatabasesSchemaheroIoDatabaseV1Alpha4Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_databases_schemahero_io_database_v1alpha4")

	var data DatabasesSchemaheroIoDatabaseV1Alpha4ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "databases.schemahero.io", Version: "v1alpha4", Resource: "databases"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
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

	var readResponse DatabasesSchemaheroIoDatabaseV1Alpha4ResourceData
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

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *DatabasesSchemaheroIoDatabaseV1Alpha4Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_databases_schemahero_io_database_v1alpha4")

	var model DatabasesSchemaheroIoDatabaseV1Alpha4ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("databases.schemahero.io/v1alpha4")
	model.Kind = pointer.String("Database")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "databases.schemahero.io", Version: "v1alpha4", Resource: "databases"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse DatabasesSchemaheroIoDatabaseV1Alpha4ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *DatabasesSchemaheroIoDatabaseV1Alpha4Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_databases_schemahero_io_database_v1alpha4")

	var data DatabasesSchemaheroIoDatabaseV1Alpha4ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "databases.schemahero.io", Version: "v1alpha4", Resource: "databases"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *DatabasesSchemaheroIoDatabaseV1Alpha4Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
