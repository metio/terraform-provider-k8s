/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package canaries_flanksource_com_v1

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
	_ datasource.DataSource = &CanariesFlanksourceComCanaryV1Manifest{}
)

func NewCanariesFlanksourceComCanaryV1Manifest() datasource.DataSource {
	return &CanariesFlanksourceComCanaryV1Manifest{}
}

type CanariesFlanksourceComCanaryV1Manifest struct{}

type CanariesFlanksourceComCanaryV1ManifestData struct {
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
		Alertmanager *[]struct {
			Alerts      *[]string `tfsdk:"alerts" json:"alerts,omitempty"`
			Connection  *string   `tfsdk:"connection" json:"connection,omitempty"`
			Description *string   `tfsdk:"description" json:"description,omitempty"`
			Display     *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Exclude_filters *map[string]string `tfsdk:"exclude_filters" json:"exclude_filters,omitempty"`
			Filters         *map[string]string `tfsdk:"filters" json:"filters,omitempty"`
			Icon            *string            `tfsdk:"icon" json:"icon,omitempty"`
			Ignore          *[]string          `tfsdk:"ignore" json:"ignore,omitempty"`
			Labels          *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics         *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Password  *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"password" json:"password,omitempty"`
			Relationships *struct {
				Components *[]struct {
					Name *struct {
						Expr  *string `tfsdk:"expr" json:"expr,omitempty"`
						Label *string `tfsdk:"label" json:"label,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"name" json:"name,omitempty"`
					Namespace *struct {
						Expr  *string `tfsdk:"expr" json:"expr,omitempty"`
						Label *string `tfsdk:"label" json:"label,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"namespace" json:"namespace,omitempty"`
					Type *struct {
						Expr  *string `tfsdk:"expr" json:"expr,omitempty"`
						Label *string `tfsdk:"label" json:"label,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"components" json:"components,omitempty"`
				Configs *[]struct {
					Name *struct {
						Expr  *string `tfsdk:"expr" json:"expr,omitempty"`
						Label *string `tfsdk:"label" json:"label,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"name" json:"name,omitempty"`
					Namespace *struct {
						Expr  *string `tfsdk:"expr" json:"expr,omitempty"`
						Label *string `tfsdk:"label" json:"label,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"namespace" json:"namespace,omitempty"`
					Type *struct {
						Expr  *string `tfsdk:"expr" json:"expr,omitempty"`
						Label *string `tfsdk:"label" json:"label,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"configs" json:"configs,omitempty"`
			} `tfsdk:"relationships" json:"relationships,omitempty"`
			Test *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			Transform *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
			Url                     *string `tfsdk:"url" json:"url,omitempty"`
			Username                *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"alertmanager" json:"alertmanager,omitempty"`
		AwsConfig *[]struct {
			AccessKey *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"access_key" json:"accessKey,omitempty"`
			AggregatorName *string `tfsdk:"aggregator_name" json:"aggregatorName,omitempty"`
			Connection     *string `tfsdk:"connection" json:"connection,omitempty"`
			Description    *string `tfsdk:"description" json:"description,omitempty"`
			Display        *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Endpoint *string            `tfsdk:"endpoint" json:"endpoint,omitempty"`
			Icon     *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels   *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics  *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Query     *string `tfsdk:"query" json:"query,omitempty"`
			Region    *string `tfsdk:"region" json:"region,omitempty"`
			SecretKey *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"secret_key" json:"secretKey,omitempty"`
			SessionToken *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"session_token" json:"sessionToken,omitempty"`
			SkipTLSVerify *bool `tfsdk:"skip_tls_verify" json:"skipTLSVerify,omitempty"`
			Test          *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			Transform *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"aws_config" json:"awsConfig,omitempty"`
		AwsConfigRule *[]struct {
			AccessKey *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"access_key" json:"accessKey,omitempty"`
			ComplianceTypes *[]string `tfsdk:"compliance_types" json:"complianceTypes,omitempty"`
			Connection      *string   `tfsdk:"connection" json:"connection,omitempty"`
			Description     *string   `tfsdk:"description" json:"description,omitempty"`
			Display         *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Endpoint    *string            `tfsdk:"endpoint" json:"endpoint,omitempty"`
			Icon        *string            `tfsdk:"icon" json:"icon,omitempty"`
			IgnoreRules *[]string          `tfsdk:"ignore_rules" json:"ignoreRules,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics     *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string   `tfsdk:"name" json:"name,omitempty"`
			Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
			Region    *string   `tfsdk:"region" json:"region,omitempty"`
			Rules     *[]string `tfsdk:"rules" json:"rules,omitempty"`
			SecretKey *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"secret_key" json:"secretKey,omitempty"`
			SessionToken *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"session_token" json:"sessionToken,omitempty"`
			SkipTLSVerify *bool `tfsdk:"skip_tls_verify" json:"skipTLSVerify,omitempty"`
			Test          *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			Transform *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"aws_config_rule" json:"awsConfigRule,omitempty"`
		AzureDevops *[]struct {
			Branch      *[]string `tfsdk:"branch" json:"branch,omitempty"`
			Connection  *string   `tfsdk:"connection" json:"connection,omitempty"`
			Description *string   `tfsdk:"description" json:"description,omitempty"`
			Display     *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Icon    *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels  *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name                *string `tfsdk:"name" json:"name,omitempty"`
			Namespace           *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Organization        *string `tfsdk:"organization" json:"organization,omitempty"`
			PersonalAccessToken *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"personal_access_token" json:"personalAccessToken,omitempty"`
			Pipeline *string `tfsdk:"pipeline" json:"pipeline,omitempty"`
			Project  *string `tfsdk:"project" json:"project,omitempty"`
			Test     *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			ThresholdMillis *int64 `tfsdk:"threshold_millis" json:"thresholdMillis,omitempty"`
			Transform       *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string            `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
			Variables               *map[string]string `tfsdk:"variables" json:"variables,omitempty"`
		} `tfsdk:"azure_devops" json:"azureDevops,omitempty"`
		Catalog *[]struct {
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Display     *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Icon    *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels  *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Selector  *[]struct {
				Agent         *string   `tfsdk:"agent" json:"agent,omitempty"`
				Cache         *string   `tfsdk:"cache" json:"cache,omitempty"`
				FieldSelector *string   `tfsdk:"field_selector" json:"fieldSelector,omitempty"`
				Id            *string   `tfsdk:"id" json:"id,omitempty"`
				LabelSelector *string   `tfsdk:"label_selector" json:"labelSelector,omitempty"`
				Name          *string   `tfsdk:"name" json:"name,omitempty"`
				Namespace     *string   `tfsdk:"namespace" json:"namespace,omitempty"`
				Statuses      *[]string `tfsdk:"statuses" json:"statuses,omitempty"`
				Types         *[]string `tfsdk:"types" json:"types,omitempty"`
			} `tfsdk:"selector" json:"selector,omitempty"`
			Test *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			Transform *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"catalog" json:"catalog,omitempty"`
		Cloudwatch *[]struct {
			AccessKey *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"access_key" json:"accessKey,omitempty"`
			ActionPrefix *string   `tfsdk:"action_prefix" json:"actionPrefix,omitempty"`
			AlarmPrefix  *string   `tfsdk:"alarm_prefix" json:"alarmPrefix,omitempty"`
			Alarms       *[]string `tfsdk:"alarms" json:"alarms,omitempty"`
			Connection   *string   `tfsdk:"connection" json:"connection,omitempty"`
			Description  *string   `tfsdk:"description" json:"description,omitempty"`
			Display      *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Endpoint *string            `tfsdk:"endpoint" json:"endpoint,omitempty"`
			Icon     *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels   *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics  *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Region    *string `tfsdk:"region" json:"region,omitempty"`
			SecretKey *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"secret_key" json:"secretKey,omitempty"`
			SessionToken *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"session_token" json:"sessionToken,omitempty"`
			SkipTLSVerify *bool   `tfsdk:"skip_tls_verify" json:"skipTLSVerify,omitempty"`
			State         *string `tfsdk:"state" json:"state,omitempty"`
			Test          *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			Transform *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"cloudwatch" json:"cloudwatch,omitempty"`
		ConfigDB *[]struct {
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Display     *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Icon    *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels  *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Query     *string `tfsdk:"query" json:"query,omitempty"`
			Test      *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			Transform *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"config_db" json:"configDB,omitempty"`
		Containerd *[]struct {
			Auth *struct {
				Password *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						HelmRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Username *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						HelmRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"auth" json:"auth,omitempty"`
			Description    *string            `tfsdk:"description" json:"description,omitempty"`
			ExpectedDigest *string            `tfsdk:"expected_digest" json:"expectedDigest,omitempty"`
			ExpectedSize   *int64             `tfsdk:"expected_size" json:"expectedSize,omitempty"`
			Icon           *string            `tfsdk:"icon" json:"icon,omitempty"`
			Image          *string            `tfsdk:"image" json:"image,omitempty"`
			Labels         *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics        *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name                    *string `tfsdk:"name" json:"name,omitempty"`
			Namespace               *string `tfsdk:"namespace" json:"namespace,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"containerd" json:"containerd,omitempty"`
		ContainerdPush *[]struct {
			Description *string            `tfsdk:"description" json:"description,omitempty"`
			Icon        *string            `tfsdk:"icon" json:"icon,omitempty"`
			Image       *string            `tfsdk:"image" json:"image,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics     *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name                    *string `tfsdk:"name" json:"name,omitempty"`
			Namespace               *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Password                *string `tfsdk:"password" json:"password,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
			Username                *string `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"containerd_push" json:"containerdPush,omitempty"`
		DatabaseBackup *[]struct {
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Display     *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Gcp *struct {
				GcpConnection *struct {
					Connection  *string `tfsdk:"connection" json:"connection,omitempty"`
					Credentials *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Value     *string `tfsdk:"value" json:"value,omitempty"`
						ValueFrom *struct {
							ConfigMapKeyRef *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
							HelmRef *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
							SecretKeyRef *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
							ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
						} `tfsdk:"value_from" json:"valueFrom,omitempty"`
					} `tfsdk:"credentials" json:"credentials,omitempty"`
					Endpoint *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
				} `tfsdk:"gcp_connection" json:"gcpConnection,omitempty"`
				Instance *string `tfsdk:"instance" json:"instance,omitempty"`
				Project  *string `tfsdk:"project" json:"project,omitempty"`
			} `tfsdk:"gcp" json:"gcp,omitempty"`
			Icon    *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels  *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			MaxAge  *string            `tfsdk:"max_age" json:"maxAge,omitempty"`
			Metrics *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Test      *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			Transform *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"database_backup" json:"databaseBackup,omitempty"`
		Dns *[]struct {
			Description *string            `tfsdk:"description" json:"description,omitempty"`
			Exactreply  *[]string          `tfsdk:"exactreply" json:"exactreply,omitempty"`
			Icon        *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics     *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Minrecords              *int64  `tfsdk:"minrecords" json:"minrecords,omitempty"`
			Name                    *string `tfsdk:"name" json:"name,omitempty"`
			Namespace               *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Port                    *int64  `tfsdk:"port" json:"port,omitempty"`
			Query                   *string `tfsdk:"query" json:"query,omitempty"`
			Querytype               *string `tfsdk:"querytype" json:"querytype,omitempty"`
			Server                  *string `tfsdk:"server" json:"server,omitempty"`
			ThresholdMillis         *int64  `tfsdk:"threshold_millis" json:"thresholdMillis,omitempty"`
			Timeout                 *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"dns" json:"dns,omitempty"`
		Docker *[]struct {
			Auth *struct {
				Password *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						HelmRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Username *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						HelmRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"auth" json:"auth,omitempty"`
			Description    *string            `tfsdk:"description" json:"description,omitempty"`
			ExpectedDigest *string            `tfsdk:"expected_digest" json:"expectedDigest,omitempty"`
			ExpectedSize   *int64             `tfsdk:"expected_size" json:"expectedSize,omitempty"`
			Icon           *string            `tfsdk:"icon" json:"icon,omitempty"`
			Image          *string            `tfsdk:"image" json:"image,omitempty"`
			Labels         *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics        *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name                    *string `tfsdk:"name" json:"name,omitempty"`
			Namespace               *string `tfsdk:"namespace" json:"namespace,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"docker" json:"docker,omitempty"`
		DockerPush *[]struct {
			Auth *struct {
				Password *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						HelmRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Username *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						HelmRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"auth" json:"auth,omitempty"`
			Description *string            `tfsdk:"description" json:"description,omitempty"`
			Icon        *string            `tfsdk:"icon" json:"icon,omitempty"`
			Image       *string            `tfsdk:"image" json:"image,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics     *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name                    *string `tfsdk:"name" json:"name,omitempty"`
			Namespace               *string `tfsdk:"namespace" json:"namespace,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"docker_push" json:"dockerPush,omitempty"`
		Dynatrace *[]struct {
			ApiKey *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"api_key" json:"apiKey,omitempty"`
			Connection  *string `tfsdk:"connection" json:"connection,omitempty"`
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Display     *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Host    *string            `tfsdk:"host" json:"host,omitempty"`
			Icon    *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels  *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Scheme    *string `tfsdk:"scheme" json:"scheme,omitempty"`
			Test      *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			Transform *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"dynatrace" json:"dynatrace,omitempty"`
		Ec2 *[]struct {
			AccessKey *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"access_key" json:"accessKey,omitempty"`
			Ami       *string `tfsdk:"ami" json:"ami,omitempty"`
			CanaryRef *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"canary_ref" json:"canaryRef,omitempty"`
			Connection  *string            `tfsdk:"connection" json:"connection,omitempty"`
			Description *string            `tfsdk:"description" json:"description,omitempty"`
			Endpoint    *string            `tfsdk:"endpoint" json:"endpoint,omitempty"`
			Icon        *string            `tfsdk:"icon" json:"icon,omitempty"`
			KeepAlive   *bool              `tfsdk:"keep_alive" json:"keepAlive,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics     *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Region    *string `tfsdk:"region" json:"region,omitempty"`
			SecretKey *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"secret_key" json:"secretKey,omitempty"`
			SecurityGroup *string `tfsdk:"security_group" json:"securityGroup,omitempty"`
			SessionToken  *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"session_token" json:"sessionToken,omitempty"`
			SkipTLSVerify           *bool   `tfsdk:"skip_tls_verify" json:"skipTLSVerify,omitempty"`
			TimeOut                 *int64  `tfsdk:"time_out" json:"timeOut,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
			UserData                *string `tfsdk:"user_data" json:"userData,omitempty"`
			WaitTime                *int64  `tfsdk:"wait_time" json:"waitTime,omitempty"`
		} `tfsdk:"ec2" json:"ec2,omitempty"`
		Elasticsearch *[]struct {
			Connection  *string `tfsdk:"connection" json:"connection,omitempty"`
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Display     *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Icon    *string            `tfsdk:"icon" json:"icon,omitempty"`
			Index   *string            `tfsdk:"index" json:"index,omitempty"`
			Labels  *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Password  *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"password" json:"password,omitempty"`
			Query   *string `tfsdk:"query" json:"query,omitempty"`
			Results *int64  `tfsdk:"results" json:"results,omitempty"`
			Test    *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			Transform *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
			Url                     *string `tfsdk:"url" json:"url,omitempty"`
			Username                *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"elasticsearch" json:"elasticsearch,omitempty"`
		Env *struct {
			ConfigMapKeyRef *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
			FieldRef *struct {
				ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
			} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
			SecretKeyRef *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"env" json:"env,omitempty"`
		Exec *[]struct {
			Artifacts *[]struct {
				Path *string `tfsdk:"path" json:"path,omitempty"`
			} `tfsdk:"artifacts" json:"artifacts,omitempty"`
			Checkout *struct {
				Certificate *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						HelmRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"certificate" json:"certificate,omitempty"`
				Connection  *string `tfsdk:"connection" json:"connection,omitempty"`
				Destination *string `tfsdk:"destination" json:"destination,omitempty"`
				Password    *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						HelmRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Url      *string `tfsdk:"url" json:"url,omitempty"`
				Username *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						HelmRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"checkout" json:"checkout,omitempty"`
			Connections *struct {
				Aws *struct {
					AccessKey *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Value     *string `tfsdk:"value" json:"value,omitempty"`
						ValueFrom *struct {
							ConfigMapKeyRef *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
							HelmRef *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
							SecretKeyRef *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
							ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
						} `tfsdk:"value_from" json:"valueFrom,omitempty"`
					} `tfsdk:"access_key" json:"accessKey,omitempty"`
					Connection *string `tfsdk:"connection" json:"connection,omitempty"`
					Endpoint   *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
					Region     *string `tfsdk:"region" json:"region,omitempty"`
					SecretKey  *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Value     *string `tfsdk:"value" json:"value,omitempty"`
						ValueFrom *struct {
							ConfigMapKeyRef *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
							HelmRef *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
							SecretKeyRef *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
							ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
						} `tfsdk:"value_from" json:"valueFrom,omitempty"`
					} `tfsdk:"secret_key" json:"secretKey,omitempty"`
					SessionToken *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Value     *string `tfsdk:"value" json:"value,omitempty"`
						ValueFrom *struct {
							ConfigMapKeyRef *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
							HelmRef *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
							SecretKeyRef *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
							ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
						} `tfsdk:"value_from" json:"valueFrom,omitempty"`
					} `tfsdk:"session_token" json:"sessionToken,omitempty"`
					SkipTLSVerify *bool `tfsdk:"skip_tls_verify" json:"skipTLSVerify,omitempty"`
				} `tfsdk:"aws" json:"aws,omitempty"`
				Azure *struct {
					ClientID *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Value     *string `tfsdk:"value" json:"value,omitempty"`
						ValueFrom *struct {
							ConfigMapKeyRef *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
							HelmRef *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
							SecretKeyRef *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
							ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
						} `tfsdk:"value_from" json:"valueFrom,omitempty"`
					} `tfsdk:"client_id" json:"clientID,omitempty"`
					ClientSecret *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Value     *string `tfsdk:"value" json:"value,omitempty"`
						ValueFrom *struct {
							ConfigMapKeyRef *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
							HelmRef *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
							SecretKeyRef *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
							ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
						} `tfsdk:"value_from" json:"valueFrom,omitempty"`
					} `tfsdk:"client_secret" json:"clientSecret,omitempty"`
					Connection *string `tfsdk:"connection" json:"connection,omitempty"`
					TenantID   *string `tfsdk:"tenant_id" json:"tenantID,omitempty"`
				} `tfsdk:"azure" json:"azure,omitempty"`
				Gcp *struct {
					Connection  *string `tfsdk:"connection" json:"connection,omitempty"`
					Credentials *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Value     *string `tfsdk:"value" json:"value,omitempty"`
						ValueFrom *struct {
							ConfigMapKeyRef *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
							HelmRef *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
							SecretKeyRef *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
							ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
						} `tfsdk:"value_from" json:"valueFrom,omitempty"`
					} `tfsdk:"credentials" json:"credentials,omitempty"`
					Endpoint *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
				} `tfsdk:"gcp" json:"gcp,omitempty"`
			} `tfsdk:"connections" json:"connections,omitempty"`
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Display     *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Env *[]struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"env" json:"env,omitempty"`
			Icon    *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels  *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Script    *string `tfsdk:"script" json:"script,omitempty"`
			Test      *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			Transform *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"exec" json:"exec,omitempty"`
		Folder *[]struct {
			AvailableSize *string `tfsdk:"available_size" json:"availableSize,omitempty"`
			AwsConnection *struct {
				AccessKey *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						HelmRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"access_key" json:"accessKey,omitempty"`
				Bucket     *string `tfsdk:"bucket" json:"bucket,omitempty"`
				Connection *string `tfsdk:"connection" json:"connection,omitempty"`
				Endpoint   *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
				ObjectPath *string `tfsdk:"object_path" json:"objectPath,omitempty"`
				Region     *string `tfsdk:"region" json:"region,omitempty"`
				SecretKey  *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						HelmRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"secret_key" json:"secretKey,omitempty"`
				SessionToken *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						HelmRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"session_token" json:"sessionToken,omitempty"`
				SkipTLSVerify *bool `tfsdk:"skip_tls_verify" json:"skipTLSVerify,omitempty"`
				UsePathStyle  *bool `tfsdk:"use_path_style" json:"usePathStyle,omitempty"`
			} `tfsdk:"aws_connection" json:"awsConnection,omitempty"`
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Display     *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Filter *struct {
				MaxAge  *string `tfsdk:"max_age" json:"maxAge,omitempty"`
				MaxSize *string `tfsdk:"max_size" json:"maxSize,omitempty"`
				MinAge  *string `tfsdk:"min_age" json:"minAge,omitempty"`
				MinSize *string `tfsdk:"min_size" json:"minSize,omitempty"`
				Regex   *string `tfsdk:"regex" json:"regex,omitempty"`
				Since   *string `tfsdk:"since" json:"since,omitempty"`
			} `tfsdk:"filter" json:"filter,omitempty"`
			GcpConnection *struct {
				Bucket      *string `tfsdk:"bucket" json:"bucket,omitempty"`
				Connection  *string `tfsdk:"connection" json:"connection,omitempty"`
				Credentials *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						HelmRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"credentials" json:"credentials,omitempty"`
				Endpoint *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
			} `tfsdk:"gcp_connection" json:"gcpConnection,omitempty"`
			Icon     *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels   *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			MaxAge   *string            `tfsdk:"max_age" json:"maxAge,omitempty"`
			MaxCount *int64             `tfsdk:"max_count" json:"maxCount,omitempty"`
			MaxSize  *string            `tfsdk:"max_size" json:"maxSize,omitempty"`
			Metrics  *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			MinAge         *string `tfsdk:"min_age" json:"minAge,omitempty"`
			MinCount       *int64  `tfsdk:"min_count" json:"minCount,omitempty"`
			MinSize        *string `tfsdk:"min_size" json:"minSize,omitempty"`
			Name           *string `tfsdk:"name" json:"name,omitempty"`
			Namespace      *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Path           *string `tfsdk:"path" json:"path,omitempty"`
			Recursive      *bool   `tfsdk:"recursive" json:"recursive,omitempty"`
			SftpConnection *struct {
				Connection *string `tfsdk:"connection" json:"connection,omitempty"`
				Host       *string `tfsdk:"host" json:"host,omitempty"`
				Password   *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						HelmRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Port     *int64 `tfsdk:"port" json:"port,omitempty"`
				Username *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						HelmRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"sftp_connection" json:"sftpConnection,omitempty"`
			SmbConnection *struct {
				Connection *string `tfsdk:"connection" json:"connection,omitempty"`
				Domain     *string `tfsdk:"domain" json:"domain,omitempty"`
				Password   *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						HelmRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Port     *int64 `tfsdk:"port" json:"port,omitempty"`
				Username *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						HelmRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"smb_connection" json:"smbConnection,omitempty"`
			Test *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			TotalSize *string `tfsdk:"total_size" json:"totalSize,omitempty"`
			Transform *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"folder" json:"folder,omitempty"`
		GitProtocol *[]struct {
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Display     *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Filename *string            `tfsdk:"filename" json:"filename,omitempty"`
			Icon     *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels   *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics  *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Password  *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"password" json:"password,omitempty"`
			Repository *string `tfsdk:"repository" json:"repository,omitempty"`
			Test       *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			Transform *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
			Username                *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"git_protocol" json:"gitProtocol,omitempty"`
		Github *[]struct {
			Connection  *string `tfsdk:"connection" json:"connection,omitempty"`
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Display     *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			GithubToken *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"github_token" json:"githubToken,omitempty"`
			Icon    *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels  *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Query     *string `tfsdk:"query" json:"query,omitempty"`
			Test      *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			Transform *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"github" json:"github,omitempty"`
		Helm *[]struct {
			Auth *struct {
				Password *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						HelmRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Username *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						HelmRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
						ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"auth" json:"auth,omitempty"`
			Cafile      *string            `tfsdk:"cafile" json:"cafile,omitempty"`
			Chartmuseum *string            `tfsdk:"chartmuseum" json:"chartmuseum,omitempty"`
			Description *string            `tfsdk:"description" json:"description,omitempty"`
			Icon        *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics     *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name                    *string `tfsdk:"name" json:"name,omitempty"`
			Namespace               *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Project                 *string `tfsdk:"project" json:"project,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"helm" json:"helm,omitempty"`
		Http *[]struct {
			Body        *string `tfsdk:"body" json:"body,omitempty"`
			Connection  *string `tfsdk:"connection" json:"connection,omitempty"`
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Display     *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Endpoint *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
			Env      *[]struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"env" json:"env,omitempty"`
			Headers *[]struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"headers" json:"headers,omitempty"`
			Icon         *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels       *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			MaxSSLExpiry *int64             `tfsdk:"max_ssl_expiry" json:"maxSSLExpiry,omitempty"`
			Method       *string            `tfsdk:"method" json:"method,omitempty"`
			Metrics      *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Ntlm      *bool   `tfsdk:"ntlm" json:"ntlm,omitempty"`
			Ntlmv2    *bool   `tfsdk:"ntlmv2" json:"ntlmv2,omitempty"`
			Oauth2    *struct {
				Params   *map[string]string `tfsdk:"params" json:"params,omitempty"`
				Scope    *[]string          `tfsdk:"scope" json:"scope,omitempty"`
				TokenURL *string            `tfsdk:"token_url" json:"tokenURL,omitempty"`
			} `tfsdk:"oauth2" json:"oauth2,omitempty"`
			Password *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"password" json:"password,omitempty"`
			ResponseCodes       *[]string `tfsdk:"response_codes" json:"responseCodes,omitempty"`
			ResponseContent     *string   `tfsdk:"response_content" json:"responseContent,omitempty"`
			ResponseJSONContent *struct {
				Path  *string `tfsdk:"path" json:"path,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"response_json_content" json:"responseJSONContent,omitempty"`
			TemplateBody *bool `tfsdk:"template_body" json:"templateBody,omitempty"`
			Test         *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			ThresholdMillis *int64 `tfsdk:"threshold_millis" json:"thresholdMillis,omitempty"`
			Transform       *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
			Url                     *string `tfsdk:"url" json:"url,omitempty"`
			Username                *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"http" json:"http,omitempty"`
		Icmp *[]struct {
			Description *string            `tfsdk:"description" json:"description,omitempty"`
			Endpoint    *string            `tfsdk:"endpoint" json:"endpoint,omitempty"`
			Icon        *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics     *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name                    *string `tfsdk:"name" json:"name,omitempty"`
			Namespace               *string `tfsdk:"namespace" json:"namespace,omitempty"`
			PacketCount             *int64  `tfsdk:"packet_count" json:"packetCount,omitempty"`
			PacketLossThreshold     *int64  `tfsdk:"packet_loss_threshold" json:"packetLossThreshold,omitempty"`
			ThresholdMillis         *int64  `tfsdk:"threshold_millis" json:"thresholdMillis,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"icmp" json:"icmp,omitempty"`
		Icon     *string `tfsdk:"icon" json:"icon,omitempty"`
		Interval *int64  `tfsdk:"interval" json:"interval,omitempty"`
		Jmeter   *[]struct {
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Host        *string `tfsdk:"host" json:"host,omitempty"`
			Icon        *string `tfsdk:"icon" json:"icon,omitempty"`
			Jmx         *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"jmx" json:"jmx,omitempty"`
			Labels  *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name                    *string   `tfsdk:"name" json:"name,omitempty"`
			Namespace               *string   `tfsdk:"namespace" json:"namespace,omitempty"`
			Port                    *int64    `tfsdk:"port" json:"port,omitempty"`
			Properties              *[]string `tfsdk:"properties" json:"properties,omitempty"`
			ResponseDuration        *string   `tfsdk:"response_duration" json:"responseDuration,omitempty"`
			SystemProperties        *[]string `tfsdk:"system_properties" json:"systemProperties,omitempty"`
			TransformDeleteStrategy *string   `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"jmeter" json:"jmeter,omitempty"`
		Junit *[]struct {
			Artifacts *[]struct {
				Path *string `tfsdk:"path" json:"path,omitempty"`
			} `tfsdk:"artifacts" json:"artifacts,omitempty"`
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Display     *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Icon    *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels  *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string            `tfsdk:"name" json:"name,omitempty"`
			Namespace *string            `tfsdk:"namespace" json:"namespace,omitempty"`
			Spec      *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
			Test      *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			TestResults *string `tfsdk:"test_results" json:"testResults,omitempty"`
			Timeout     *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
			Transform   *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"junit" json:"junit,omitempty"`
		Kubernetes *[]struct {
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Display     *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Icon       *string   `tfsdk:"icon" json:"icon,omitempty"`
			Ignore     *[]string `tfsdk:"ignore" json:"ignore,omitempty"`
			Kind       *string   `tfsdk:"kind" json:"kind,omitempty"`
			Kubeconfig *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"kubeconfig" json:"kubeconfig,omitempty"`
			Labels  *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name              *string `tfsdk:"name" json:"name,omitempty"`
			Namespace         *string `tfsdk:"namespace" json:"namespace,omitempty"`
			NamespaceSelector *struct {
				FieldSelector *string `tfsdk:"field_selector" json:"fieldSelector,omitempty"`
				LabelSelector *string `tfsdk:"label_selector" json:"labelSelector,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
			Ready    *bool `tfsdk:"ready" json:"ready,omitempty"`
			Resource *struct {
				FieldSelector *string `tfsdk:"field_selector" json:"fieldSelector,omitempty"`
				LabelSelector *string `tfsdk:"label_selector" json:"labelSelector,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"resource" json:"resource,omitempty"`
			Test *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			Transform *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
		Ldap *[]struct {
			BindDN      *string            `tfsdk:"bind_dn" json:"bindDN,omitempty"`
			Connection  *string            `tfsdk:"connection" json:"connection,omitempty"`
			Description *string            `tfsdk:"description" json:"description,omitempty"`
			Icon        *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics     *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Password  *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"password" json:"password,omitempty"`
			SkipTLSVerify           *bool   `tfsdk:"skip_tls_verify" json:"skipTLSVerify,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
			Url                     *string `tfsdk:"url" json:"url,omitempty"`
			UserSearch              *string `tfsdk:"user_search" json:"userSearch,omitempty"`
			Username                *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"ldap" json:"ldap,omitempty"`
		Mongodb *[]struct {
			Connection  *string            `tfsdk:"connection" json:"connection,omitempty"`
			Description *string            `tfsdk:"description" json:"description,omitempty"`
			Icon        *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics     *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Password  *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"password" json:"password,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
			Url                     *string `tfsdk:"url" json:"url,omitempty"`
			Username                *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"mongodb" json:"mongodb,omitempty"`
		Mssql *[]struct {
			Connection  *string `tfsdk:"connection" json:"connection,omitempty"`
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Display     *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Icon    *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels  *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Password  *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"password" json:"password,omitempty"`
			Query   *string `tfsdk:"query" json:"query,omitempty"`
			Results *int64  `tfsdk:"results" json:"results,omitempty"`
			Test    *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			Transform *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
			Url                     *string `tfsdk:"url" json:"url,omitempty"`
			Username                *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"mssql" json:"mssql,omitempty"`
		Mysql *[]struct {
			Connection  *string `tfsdk:"connection" json:"connection,omitempty"`
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Display     *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Icon    *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels  *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Password  *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"password" json:"password,omitempty"`
			Query   *string `tfsdk:"query" json:"query,omitempty"`
			Results *int64  `tfsdk:"results" json:"results,omitempty"`
			Test    *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			Transform *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
			Url                     *string `tfsdk:"url" json:"url,omitempty"`
			Username                *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"mysql" json:"mysql,omitempty"`
		Namespace *[]struct {
			Deadline             *int64             `tfsdk:"deadline" json:"deadline,omitempty"`
			DeleteTimeout        *int64             `tfsdk:"delete_timeout" json:"deleteTimeout,omitempty"`
			Description          *string            `tfsdk:"description" json:"description,omitempty"`
			ExpectedContent      *string            `tfsdk:"expected_content" json:"expectedContent,omitempty"`
			ExpectedHttpStatuses *[]string          `tfsdk:"expected_http_statuses" json:"expectedHttpStatuses,omitempty"`
			HttpRetryInterval    *int64             `tfsdk:"http_retry_interval" json:"httpRetryInterval,omitempty"`
			HttpTimeout          *int64             `tfsdk:"http_timeout" json:"httpTimeout,omitempty"`
			Icon                 *string            `tfsdk:"icon" json:"icon,omitempty"`
			IngressHost          *string            `tfsdk:"ingress_host" json:"ingressHost,omitempty"`
			IngressName          *string            `tfsdk:"ingress_name" json:"ingressName,omitempty"`
			IngressTimeout       *int64             `tfsdk:"ingress_timeout" json:"ingressTimeout,omitempty"`
			Labels               *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics              *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name                    *string            `tfsdk:"name" json:"name,omitempty"`
			Namespace               *string            `tfsdk:"namespace" json:"namespace,omitempty"`
			NamespaceAnnotations    *map[string]string `tfsdk:"namespace_annotations" json:"namespaceAnnotations,omitempty"`
			NamespaceLabels         *map[string]string `tfsdk:"namespace_labels" json:"namespaceLabels,omitempty"`
			NamespaceNamePrefix     *string            `tfsdk:"namespace_name_prefix" json:"namespaceNamePrefix,omitempty"`
			Path                    *string            `tfsdk:"path" json:"path,omitempty"`
			PodSpec                 *string            `tfsdk:"pod_spec" json:"podSpec,omitempty"`
			Port                    *int64             `tfsdk:"port" json:"port,omitempty"`
			PriorityClass           *string            `tfsdk:"priority_class" json:"priorityClass,omitempty"`
			ReadyTimeout            *int64             `tfsdk:"ready_timeout" json:"readyTimeout,omitempty"`
			Schedule_timeout        *int64             `tfsdk:"schedule_timeout" json:"schedule_timeout,omitempty"`
			TransformDeleteStrategy *string            `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"namespace" json:"namespace,omitempty"`
		Opensearch *[]struct {
			Connection  *string `tfsdk:"connection" json:"connection,omitempty"`
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Display     *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Icon    *string            `tfsdk:"icon" json:"icon,omitempty"`
			Index   *string            `tfsdk:"index" json:"index,omitempty"`
			Labels  *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Password  *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"password" json:"password,omitempty"`
			Query   *string `tfsdk:"query" json:"query,omitempty"`
			Results *int64  `tfsdk:"results" json:"results,omitempty"`
			Test    *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			Transform *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
			Url                     *string `tfsdk:"url" json:"url,omitempty"`
			Username                *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"opensearch" json:"opensearch,omitempty"`
		Owner *string `tfsdk:"owner" json:"owner,omitempty"`
		Pod   *[]struct {
			Deadline             *int64             `tfsdk:"deadline" json:"deadline,omitempty"`
			DeleteTimeout        *int64             `tfsdk:"delete_timeout" json:"deleteTimeout,omitempty"`
			Description          *string            `tfsdk:"description" json:"description,omitempty"`
			ExpectedContent      *string            `tfsdk:"expected_content" json:"expectedContent,omitempty"`
			ExpectedHttpStatuses *[]string          `tfsdk:"expected_http_statuses" json:"expectedHttpStatuses,omitempty"`
			HttpRetryInterval    *int64             `tfsdk:"http_retry_interval" json:"httpRetryInterval,omitempty"`
			HttpTimeout          *int64             `tfsdk:"http_timeout" json:"httpTimeout,omitempty"`
			Icon                 *string            `tfsdk:"icon" json:"icon,omitempty"`
			IngressClass         *string            `tfsdk:"ingress_class" json:"ingressClass,omitempty"`
			IngressHost          *string            `tfsdk:"ingress_host" json:"ingressHost,omitempty"`
			IngressName          *string            `tfsdk:"ingress_name" json:"ingressName,omitempty"`
			IngressTimeout       *int64             `tfsdk:"ingress_timeout" json:"ingressTimeout,omitempty"`
			Labels               *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics              *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name                    *string `tfsdk:"name" json:"name,omitempty"`
			Namespace               *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Path                    *string `tfsdk:"path" json:"path,omitempty"`
			Port                    *int64  `tfsdk:"port" json:"port,omitempty"`
			PriorityClass           *string `tfsdk:"priority_class" json:"priorityClass,omitempty"`
			ReadyTimeout            *int64  `tfsdk:"ready_timeout" json:"readyTimeout,omitempty"`
			RoundRobinNodes         *bool   `tfsdk:"round_robin_nodes" json:"roundRobinNodes,omitempty"`
			ScheduleTimeout         *int64  `tfsdk:"schedule_timeout" json:"scheduleTimeout,omitempty"`
			Spec                    *string `tfsdk:"spec" json:"spec,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"pod" json:"pod,omitempty"`
		Postgres *[]struct {
			Connection  *string `tfsdk:"connection" json:"connection,omitempty"`
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Display     *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Icon    *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels  *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Password  *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"password" json:"password,omitempty"`
			Query   *string `tfsdk:"query" json:"query,omitempty"`
			Results *int64  `tfsdk:"results" json:"results,omitempty"`
			Test    *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			Transform *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
			Url                     *string `tfsdk:"url" json:"url,omitempty"`
			Username                *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"postgres" json:"postgres,omitempty"`
		Prometheus *[]struct {
			Connection  *string `tfsdk:"connection" json:"connection,omitempty"`
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Display     *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Host    *string            `tfsdk:"host" json:"host,omitempty"`
			Icon    *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels  *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Password  *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"password" json:"password,omitempty"`
			Query *string `tfsdk:"query" json:"query,omitempty"`
			Test  *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			Transform *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
			Url                     *string `tfsdk:"url" json:"url,omitempty"`
			Username                *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"prometheus" json:"prometheus,omitempty"`
		Redis *[]struct {
			Addr        *string            `tfsdk:"addr" json:"addr,omitempty"`
			Connection  *string            `tfsdk:"connection" json:"connection,omitempty"`
			Db          *int64             `tfsdk:"db" json:"db,omitempty"`
			Description *string            `tfsdk:"description" json:"description,omitempty"`
			Icon        *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics     *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Password  *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"password" json:"password,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
			Url                     *string `tfsdk:"url" json:"url,omitempty"`
			Username                *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"redis" json:"redis,omitempty"`
		Replicas *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
		Restic   *[]struct {
			AccessKey *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"access_key" json:"accessKey,omitempty"`
			AwsConnectionName *string            `tfsdk:"aws_connection_name" json:"awsConnectionName,omitempty"`
			CaCert            *string            `tfsdk:"ca_cert" json:"caCert,omitempty"`
			CheckIntegrity    *bool              `tfsdk:"check_integrity" json:"checkIntegrity,omitempty"`
			Connection        *string            `tfsdk:"connection" json:"connection,omitempty"`
			Description       *string            `tfsdk:"description" json:"description,omitempty"`
			Icon              *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels            *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			MaxAge            *string            `tfsdk:"max_age" json:"maxAge,omitempty"`
			Metrics           *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Password  *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"password" json:"password,omitempty"`
			Repository *string `tfsdk:"repository" json:"repository,omitempty"`
			SecretKey  *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"secret_key" json:"secretKey,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"restic" json:"restic,omitempty"`
		ResultMode *string `tfsdk:"result_mode" json:"resultMode,omitempty"`
		S3         *[]struct {
			AccessKey *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"access_key" json:"accessKey,omitempty"`
			Bucket      *string            `tfsdk:"bucket" json:"bucket,omitempty"`
			BucketName  *string            `tfsdk:"bucket_name" json:"bucketName,omitempty"`
			Connection  *string            `tfsdk:"connection" json:"connection,omitempty"`
			Description *string            `tfsdk:"description" json:"description,omitempty"`
			Endpoint    *string            `tfsdk:"endpoint" json:"endpoint,omitempty"`
			Icon        *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics     *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			Namespace  *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ObjectPath *string `tfsdk:"object_path" json:"objectPath,omitempty"`
			Region     *string `tfsdk:"region" json:"region,omitempty"`
			SecretKey  *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"secret_key" json:"secretKey,omitempty"`
			SessionToken *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"session_token" json:"sessionToken,omitempty"`
			SkipTLSVerify           *bool   `tfsdk:"skip_tls_verify" json:"skipTLSVerify,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
			UsePathStyle            *bool   `tfsdk:"use_path_style" json:"usePathStyle,omitempty"`
		} `tfsdk:"s3" json:"s3,omitempty"`
		Schedule *string `tfsdk:"schedule" json:"schedule,omitempty"`
		Severity *string `tfsdk:"severity" json:"severity,omitempty"`
		Tcp      *[]struct {
			Description *string            `tfsdk:"description" json:"description,omitempty"`
			Endpoint    *string            `tfsdk:"endpoint" json:"endpoint,omitempty"`
			Icon        *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics     *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name                    *string `tfsdk:"name" json:"name,omitempty"`
			Namespace               *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ThresholdMillis         *int64  `tfsdk:"threshold_millis" json:"thresholdMillis,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"tcp" json:"tcp,omitempty"`
		Webhook *struct {
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Display     *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Icon    *string            `tfsdk:"icon" json:"icon,omitempty"`
			Labels  *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Metrics *[]struct {
				Labels *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueExpr *string `tfsdk:"value_expr" json:"valueExpr,omitempty"`
				} `tfsdk:"labels" json:"labels,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Test      *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"test" json:"test,omitempty"`
			Token *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					HelmRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"helm_ref" json:"helmRef,omitempty"`
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"token" json:"token,omitempty"`
			Transform *struct {
				Expr       *string `tfsdk:"expr" json:"expr,omitempty"`
				Javascript *string `tfsdk:"javascript" json:"javascript,omitempty"`
				JsonPath   *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Template   *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"transform" json:"transform,omitempty"`
			TransformDeleteStrategy *string `tfsdk:"transform_delete_strategy" json:"transformDeleteStrategy,omitempty"`
		} `tfsdk:"webhook" json:"webhook,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CanariesFlanksourceComCanaryV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_canaries_flanksource_com_canary_v1_manifest"
}

func (r *CanariesFlanksourceComCanaryV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Canary is the Schema for the canaries API",
		MarkdownDescription: "Canary is the Schema for the canaries API",
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
				Description:         "CanarySpec defines the desired state of Canary",
				MarkdownDescription: "CanarySpec defines the desired state of Canary",
				Attributes: map[string]schema.Attribute{
					"alertmanager": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"alerts": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"connection": schema.StringAttribute{
									Description:         "Connection name e.g. connection://http/google",
									MarkdownDescription: "Connection name e.g. connection://http/google",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"display": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"exclude_filters": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"filters": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ignore": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"password": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"relationships": schema.SingleNestedAttribute{
									Description:         "Relationships defines a way to link the check results to components and configsusing lookup expressions.",
									MarkdownDescription: "Relationships defines a way to link the check results to components and configsusing lookup expressions.",
									Attributes: map[string]schema.Attribute{
										"components": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.SingleNestedAttribute{
														Description:         "Lookup specifies the type of lookup to perform.",
														MarkdownDescription: "Lookup specifies the type of lookup to perform.",
														Attributes: map[string]schema.Attribute{
															"expr": schema.StringAttribute{
																Description:         "Expr is a cel-expression.",
																MarkdownDescription: "Expr is a cel-expression.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"label": schema.StringAttribute{
																Description:         "Label specifies the key to lookup on the label.",
																MarkdownDescription: "Label specifies the key to lookup on the label.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.StringAttribute{
																Description:         "Value is the static value to use.",
																MarkdownDescription: "Value is the static value to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": schema.SingleNestedAttribute{
														Description:         "Lookup specifies the type of lookup to perform.",
														MarkdownDescription: "Lookup specifies the type of lookup to perform.",
														Attributes: map[string]schema.Attribute{
															"expr": schema.StringAttribute{
																Description:         "Expr is a cel-expression.",
																MarkdownDescription: "Expr is a cel-expression.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"label": schema.StringAttribute{
																Description:         "Label specifies the key to lookup on the label.",
																MarkdownDescription: "Label specifies the key to lookup on the label.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.StringAttribute{
																Description:         "Value is the static value to use.",
																MarkdownDescription: "Value is the static value to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"type": schema.SingleNestedAttribute{
														Description:         "Lookup specifies the type of lookup to perform.",
														MarkdownDescription: "Lookup specifies the type of lookup to perform.",
														Attributes: map[string]schema.Attribute{
															"expr": schema.StringAttribute{
																Description:         "Expr is a cel-expression.",
																MarkdownDescription: "Expr is a cel-expression.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"label": schema.StringAttribute{
																Description:         "Label specifies the key to lookup on the label.",
																MarkdownDescription: "Label specifies the key to lookup on the label.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.StringAttribute{
																Description:         "Value is the static value to use.",
																MarkdownDescription: "Value is the static value to use.",
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

										"configs": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.SingleNestedAttribute{
														Description:         "Lookup specifies the type of lookup to perform.",
														MarkdownDescription: "Lookup specifies the type of lookup to perform.",
														Attributes: map[string]schema.Attribute{
															"expr": schema.StringAttribute{
																Description:         "Expr is a cel-expression.",
																MarkdownDescription: "Expr is a cel-expression.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"label": schema.StringAttribute{
																Description:         "Label specifies the key to lookup on the label.",
																MarkdownDescription: "Label specifies the key to lookup on the label.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.StringAttribute{
																Description:         "Value is the static value to use.",
																MarkdownDescription: "Value is the static value to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": schema.SingleNestedAttribute{
														Description:         "Lookup specifies the type of lookup to perform.",
														MarkdownDescription: "Lookup specifies the type of lookup to perform.",
														Attributes: map[string]schema.Attribute{
															"expr": schema.StringAttribute{
																Description:         "Expr is a cel-expression.",
																MarkdownDescription: "Expr is a cel-expression.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"label": schema.StringAttribute{
																Description:         "Label specifies the key to lookup on the label.",
																MarkdownDescription: "Label specifies the key to lookup on the label.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.StringAttribute{
																Description:         "Value is the static value to use.",
																MarkdownDescription: "Value is the static value to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"type": schema.SingleNestedAttribute{
														Description:         "Lookup specifies the type of lookup to perform.",
														MarkdownDescription: "Lookup specifies the type of lookup to perform.",
														Attributes: map[string]schema.Attribute{
															"expr": schema.StringAttribute{
																Description:         "Expr is a cel-expression.",
																MarkdownDescription: "Expr is a cel-expression.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"label": schema.StringAttribute{
																Description:         "Label specifies the key to lookup on the label.",
																MarkdownDescription: "Label specifies the key to lookup on the label.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.StringAttribute{
																Description:         "Value is the static value to use.",
																MarkdownDescription: "Value is the static value to use.",
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"test": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"url": schema.StringAttribute{
									Description:         "Connection url, interpolated with username,password",
									MarkdownDescription: "Connection url, interpolated with username,password",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"username": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

					"aws_config": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"access_key": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"aggregator_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"connection": schema.StringAttribute{
									Description:         "ConnectionName of the connection. It'll be used to populate the endpoint, accessKey and secretKey.",
									MarkdownDescription: "ConnectionName of the connection. It'll be used to populate the endpoint, accessKey and secretKey.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"display": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"endpoint": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"query": schema.StringAttribute{
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

								"secret_key": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"session_token": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"skip_tls_verify": schema.BoolAttribute{
									Description:         "Skip TLS verify when connecting to aws",
									MarkdownDescription: "Skip TLS verify when connecting to aws",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"test": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"aws_config_rule": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"access_key": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"compliance_types": schema.ListAttribute{
									Description:         "Filters the results by compliance. The allowed values are INSUFFICIENT_DATA, NON_COMPLIANT, NOT_APPLICABLE, COMPLIANT",
									MarkdownDescription: "Filters the results by compliance. The allowed values are INSUFFICIENT_DATA, NON_COMPLIANT, NOT_APPLICABLE, COMPLIANT",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"connection": schema.StringAttribute{
									Description:         "ConnectionName of the connection. It'll be used to populate the endpoint, accessKey and secretKey.",
									MarkdownDescription: "ConnectionName of the connection. It'll be used to populate the endpoint, accessKey and secretKey.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"display": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"endpoint": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ignore_rules": schema.ListAttribute{
									Description:         "List of rules which would be omitted from the fetch result",
									MarkdownDescription: "List of rules which would be omitted from the fetch result",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"region": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"rules": schema.ListAttribute{
									Description:         "Specify one or more Config rule names to filter the results by rule.",
									MarkdownDescription: "Specify one or more Config rule names to filter the results by rule.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"secret_key": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"session_token": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"skip_tls_verify": schema.BoolAttribute{
									Description:         "Skip TLS verify when connecting to aws",
									MarkdownDescription: "Skip TLS verify when connecting to aws",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"test": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"azure_devops": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"branch": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"connection": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"display": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"organization": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"personal_access_token": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"pipeline": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"project": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"test": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"threshold_millis": schema.Int64Attribute{
									Description:         "ThresholdMillis the maximum duration of a Run. (Optional)",
									MarkdownDescription: "ThresholdMillis the maximum duration of a Run. (Optional)",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"transform": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"variables": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"catalog": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"display": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"selector": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"agent": schema.StringAttribute{
												Description:         "Agent can be the agent id or the name of the agent. Additionally, the special 'self' value can be used to select resources without an agent.",
												MarkdownDescription: "Agent can be the agent id or the name of the agent. Additionally, the special 'self' value can be used to select resources without an agent.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cache": schema.StringAttribute{
												Description:         "Cache directives 'no-cache' (should not fetch from cache but can be cached) 'no-store' (should not cache) 'max-age=X' (cache for X duration)",
												MarkdownDescription: "Cache directives 'no-cache' (should not fetch from cache but can be cached) 'no-store' (should not cache) 'max-age=X' (cache for X duration)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"field_selector": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"id": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"label_selector": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
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

											"statuses": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"types": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
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

								"test": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"cloudwatch": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"access_key": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"action_prefix": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"alarm_prefix": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"alarms": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"connection": schema.StringAttribute{
									Description:         "ConnectionName of the connection. It'll be used to populate the endpoint, accessKey and secretKey.",
									MarkdownDescription: "ConnectionName of the connection. It'll be used to populate the endpoint, accessKey and secretKey.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"display": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"endpoint": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"region": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"secret_key": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"session_token": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"skip_tls_verify": schema.BoolAttribute{
									Description:         "Skip TLS verify when connecting to aws",
									MarkdownDescription: "Skip TLS verify when connecting to aws",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"state": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"test": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"config_db": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"display": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"query": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"test": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"containerd": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"auth": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"password": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
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

												"value_from": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"helm_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"service_account": schema.StringAttribute{
															Description:         "ServiceAccount specifies the service account whose token should be fetched",
															MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

										"username": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
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

												"value_from": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"helm_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"service_account": schema.StringAttribute{
															Description:         "ServiceAccount specifies the service account whose token should be fetched",
															MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"expected_digest": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"expected_size": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"image": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"containerd_push": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"image": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"password": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"username": schema.StringAttribute{
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

					"database_backup": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"display": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"gcp": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"gcp_connection": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"connection": schema.StringAttribute{
													Description:         "ConnectionName of the connection. It'll be used to populate the endpoint and credentials.",
													MarkdownDescription: "ConnectionName of the connection. It'll be used to populate the endpoint and credentials.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"credentials": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
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

														"value_from": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"config_map_key_ref": schema.SingleNestedAttribute{
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
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"helm_ref": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																			MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
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
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"service_account": schema.StringAttribute{
																	Description:         "ServiceAccount specifies the service account whose token should be fetched",
																	MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

												"endpoint": schema.StringAttribute{
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

										"instance": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"project": schema.StringAttribute{
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

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_age": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"test": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"dns": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"exactreply": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"minrecords": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"port": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"query": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"querytype": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"server": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"threshold_millis": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"timeout": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"docker": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"auth": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"password": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
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

												"value_from": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"helm_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"service_account": schema.StringAttribute{
															Description:         "ServiceAccount specifies the service account whose token should be fetched",
															MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

										"username": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
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

												"value_from": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"helm_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"service_account": schema.StringAttribute{
															Description:         "ServiceAccount specifies the service account whose token should be fetched",
															MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"expected_digest": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"expected_size": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"image": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"docker_push": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"auth": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"password": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
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

												"value_from": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"helm_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"service_account": schema.StringAttribute{
															Description:         "ServiceAccount specifies the service account whose token should be fetched",
															MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

										"username": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
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

												"value_from": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"helm_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"service_account": schema.StringAttribute{
															Description:         "ServiceAccount specifies the service account whose token should be fetched",
															MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"image": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"dynatrace": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"api_key": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"connection": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"display": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"host": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"scheme": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"test": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"ec2": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"access_key": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"ami": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"canary_ref": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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

								"connection": schema.StringAttribute{
									Description:         "ConnectionName of the connection. It'll be used to populate the endpoint, accessKey and secretKey.",
									MarkdownDescription: "ConnectionName of the connection. It'll be used to populate the endpoint, accessKey and secretKey.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
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

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"keep_alive": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"region": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"secret_key": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"security_group": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"session_token": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"skip_tls_verify": schema.BoolAttribute{
									Description:         "Skip TLS verify when connecting to aws",
									MarkdownDescription: "Skip TLS verify when connecting to aws",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"time_out": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"user_data": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"wait_time": schema.Int64Attribute{
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

					"elasticsearch": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"connection": schema.StringAttribute{
									Description:         "Connection name e.g. connection://http/google",
									MarkdownDescription: "Connection name e.g. connection://http/google",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"display": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"index": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"password": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"query": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"results": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"test": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"url": schema.StringAttribute{
									Description:         "Connection url, interpolated with username,password",
									MarkdownDescription: "Connection url, interpolated with username,password",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"username": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

					"env": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
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
										Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
								Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations,spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
								MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations,spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
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
										Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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

							"value": schema.StringAttribute{
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

					"exec": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"artifacts": schema.ListNestedAttribute{
									Description:         "Artifacts configure the artifacts generated by the check",
									MarkdownDescription: "Artifacts configure the artifacts generated by the check",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"path": schema.StringAttribute{
												Description:         "Path to the artifact on the check runner.Special paths: /dev/stdout & /dev/stdin",
												MarkdownDescription: "Path to the artifact on the check runner.Special paths: /dev/stdout & /dev/stdin",
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

								"checkout": schema.SingleNestedAttribute{
									Description:         "Checkout details the git repository that should be mounted to the process",
									MarkdownDescription: "Checkout details the git repository that should be mounted to the process",
									Attributes: map[string]schema.Attribute{
										"certificate": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
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

												"value_from": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"helm_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"service_account": schema.StringAttribute{
															Description:         "ServiceAccount specifies the service account whose token should be fetched",
															MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

										"connection": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"destination": schema.StringAttribute{
											Description:         "Destination is the full path to where the contents of the URL should be downloaded to.If left empty, the sha256 hash of the URL will be used as the dir name.",
											MarkdownDescription: "Destination is the full path to where the contents of the URL should be downloaded to.If left empty, the sha256 hash of the URL will be used as the dir name.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"password": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
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

												"value_from": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"helm_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"service_account": schema.StringAttribute{
															Description:         "ServiceAccount specifies the service account whose token should be fetched",
															MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

										"url": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"username": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
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

												"value_from": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"helm_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"service_account": schema.StringAttribute{
															Description:         "ServiceAccount specifies the service account whose token should be fetched",
															MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"connections": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"aws": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"access_key": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
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

														"value_from": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"config_map_key_ref": schema.SingleNestedAttribute{
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
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"helm_ref": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																			MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
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
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"service_account": schema.StringAttribute{
																	Description:         "ServiceAccount specifies the service account whose token should be fetched",
																	MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

												"connection": schema.StringAttribute{
													Description:         "ConnectionName of the connection. It'll be used to populate the endpoint, accessKey and secretKey.",
													MarkdownDescription: "ConnectionName of the connection. It'll be used to populate the endpoint, accessKey and secretKey.",
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

												"region": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret_key": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
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

														"value_from": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"config_map_key_ref": schema.SingleNestedAttribute{
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
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"helm_ref": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																			MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
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
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"service_account": schema.StringAttribute{
																	Description:         "ServiceAccount specifies the service account whose token should be fetched",
																	MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

												"session_token": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
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

														"value_from": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"config_map_key_ref": schema.SingleNestedAttribute{
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
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"helm_ref": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																			MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
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
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"service_account": schema.StringAttribute{
																	Description:         "ServiceAccount specifies the service account whose token should be fetched",
																	MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

												"skip_tls_verify": schema.BoolAttribute{
													Description:         "Skip TLS verify when connecting to aws",
													MarkdownDescription: "Skip TLS verify when connecting to aws",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"azure": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"client_id": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
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

														"value_from": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"config_map_key_ref": schema.SingleNestedAttribute{
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
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"helm_ref": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																			MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
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
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"service_account": schema.StringAttribute{
																	Description:         "ServiceAccount specifies the service account whose token should be fetched",
																	MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

												"client_secret": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
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

														"value_from": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"config_map_key_ref": schema.SingleNestedAttribute{
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
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"helm_ref": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																			MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
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
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"service_account": schema.StringAttribute{
																	Description:         "ServiceAccount specifies the service account whose token should be fetched",
																	MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

												"connection": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tenant_id": schema.StringAttribute{
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

										"gcp": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"connection": schema.StringAttribute{
													Description:         "ConnectionName of the connection. It'll be used to populate the endpoint and credentials.",
													MarkdownDescription: "ConnectionName of the connection. It'll be used to populate the endpoint and credentials.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"credentials": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
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

														"value_from": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"config_map_key_ref": schema.SingleNestedAttribute{
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
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"helm_ref": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																			MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
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
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"service_account": schema.StringAttribute{
																	Description:         "ServiceAccount specifies the service account whose token should be fetched",
																	MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

												"endpoint": schema.StringAttribute{
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

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"display": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"env": schema.ListNestedAttribute{
									Description:         "EnvVars are the environment variables that are accessible to exec processes",
									MarkdownDescription: "EnvVars are the environment variables that are accessible to exec processes",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
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

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"config_map_key_ref": schema.SingleNestedAttribute{
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
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"helm_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
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
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"service_account": schema.StringAttribute{
														Description:         "ServiceAccount specifies the service account whose token should be fetched",
														MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"script": schema.StringAttribute{
									Description:         "Script can be a inline script or a path to a script that needs to be executedOn windows executed via powershell and in darwin and linux executed using bash",
									MarkdownDescription: "Script can be a inline script or a path to a script that needs to be executedOn windows executed via powershell and in darwin and linux executed using bash",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"test": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"folder": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"available_size": schema.StringAttribute{
									Description:         "AvailableSize present on the filesystem",
									MarkdownDescription: "AvailableSize present on the filesystem",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"aws_connection": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"access_key": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
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

												"value_from": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"helm_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"service_account": schema.StringAttribute{
															Description:         "ServiceAccount specifies the service account whose token should be fetched",
															MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

										"bucket": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"connection": schema.StringAttribute{
											Description:         "ConnectionName of the connection. It'll be used to populate the endpoint, accessKey and secretKey.",
											MarkdownDescription: "ConnectionName of the connection. It'll be used to populate the endpoint, accessKey and secretKey.",
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

										"object_path": schema.StringAttribute{
											Description:         "glob path to restrict matches to a subset",
											MarkdownDescription: "glob path to restrict matches to a subset",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"region": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_key": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
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

												"value_from": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"helm_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"service_account": schema.StringAttribute{
															Description:         "ServiceAccount specifies the service account whose token should be fetched",
															MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

										"session_token": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
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

												"value_from": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"helm_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"service_account": schema.StringAttribute{
															Description:         "ServiceAccount specifies the service account whose token should be fetched",
															MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

										"skip_tls_verify": schema.BoolAttribute{
											Description:         "Skip TLS verify when connecting to aws",
											MarkdownDescription: "Skip TLS verify when connecting to aws",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"use_path_style": schema.BoolAttribute{
											Description:         "Use path style path: http://s3.amazonaws.com/BUCKET/KEY instead of http://BUCKET.s3.amazonaws.com/KEY",
											MarkdownDescription: "Use path style path: http://s3.amazonaws.com/BUCKET/KEY instead of http://BUCKET.s3.amazonaws.com/KEY",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"display": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"filter": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"max_age": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"max_size": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"min_age": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"min_size": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"regex": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"since": schema.StringAttribute{
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

								"gcp_connection": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"bucket": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"connection": schema.StringAttribute{
											Description:         "ConnectionName of the connection. It'll be used to populate the endpoint and credentials.",
											MarkdownDescription: "ConnectionName of the connection. It'll be used to populate the endpoint and credentials.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"credentials": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
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

												"value_from": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"helm_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"service_account": schema.StringAttribute{
															Description:         "ServiceAccount specifies the service account whose token should be fetched",
															MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

										"endpoint": schema.StringAttribute{
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

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_age": schema.StringAttribute{
									Description:         "MaxAge the latest object should be younger than defined age",
									MarkdownDescription: "MaxAge the latest object should be younger than defined age",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_count": schema.Int64Attribute{
									Description:         "MinCount the minimum number of files inside the searchPath",
									MarkdownDescription: "MinCount the minimum number of files inside the searchPath",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_size": schema.StringAttribute{
									Description:         "MaxSize of the files inside the searchPath",
									MarkdownDescription: "MaxSize of the files inside the searchPath",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"min_age": schema.StringAttribute{
									Description:         "MinAge the latest object should be older than defined age",
									MarkdownDescription: "MinAge the latest object should be older than defined age",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"min_count": schema.Int64Attribute{
									Description:         "MinCount the minimum number of files inside the searchPath",
									MarkdownDescription: "MinCount the minimum number of files inside the searchPath",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"min_size": schema.StringAttribute{
									Description:         "MinSize of the files inside the searchPath",
									MarkdownDescription: "MinSize of the files inside the searchPath",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"path": schema.StringAttribute{
									Description:         "Path  to folder or object storage, e.g. 's3://<bucket-name>',  'gcs://<bucket-name>', '/path/tp/folder'",
									MarkdownDescription: "Path  to folder or object storage, e.g. 's3://<bucket-name>',  'gcs://<bucket-name>', '/path/tp/folder'",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"recursive": schema.BoolAttribute{
									Description:         "Recursive when set to true will recursively scan the folder to list the files in it.However, symlinks are simply listed but not traversed.",
									MarkdownDescription: "Recursive when set to true will recursively scan the folder to list the files in it.However, symlinks are simply listed but not traversed.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"sftp_connection": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"connection": schema.StringAttribute{
											Description:         "ConnectionName of the connection. It'll be used to populate the connection fields.",
											MarkdownDescription: "ConnectionName of the connection. It'll be used to populate the connection fields.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"host": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"password": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
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

												"value_from": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"helm_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"service_account": schema.StringAttribute{
															Description:         "ServiceAccount specifies the service account whose token should be fetched",
															MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

										"port": schema.Int64Attribute{
											Description:         "Port for the SSH server. Defaults to 22",
											MarkdownDescription: "Port for the SSH server. Defaults to 22",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"username": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
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

												"value_from": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"helm_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"service_account": schema.StringAttribute{
															Description:         "ServiceAccount specifies the service account whose token should be fetched",
															MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"smb_connection": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"connection": schema.StringAttribute{
											Description:         "ConnectionName of the connection. It'll be used to populate the connection fields.",
											MarkdownDescription: "ConnectionName of the connection. It'll be used to populate the connection fields.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"domain": schema.StringAttribute{
											Description:         "Domain...",
											MarkdownDescription: "Domain...",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"password": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
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

												"value_from": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"helm_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"service_account": schema.StringAttribute{
															Description:         "ServiceAccount specifies the service account whose token should be fetched",
															MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

										"port": schema.Int64Attribute{
											Description:         "Port on which smb server is running. Defaults to 445",
											MarkdownDescription: "Port on which smb server is running. Defaults to 445",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"username": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
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

												"value_from": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"helm_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"service_account": schema.StringAttribute{
															Description:         "ServiceAccount specifies the service account whose token should be fetched",
															MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"test": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"total_size": schema.StringAttribute{
									Description:         "TotalSize present on the filesystem",
									MarkdownDescription: "TotalSize present on the filesystem",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"transform": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"git_protocol": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"display": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"filename": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"password": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"repository": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"test": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"username": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

					"github": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"connection": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"display": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"github_token": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"query": schema.StringAttribute{
									Description:         "Query to be executed. Please see https://github.com/askgitdev/askgit for more details regarding syntax",
									MarkdownDescription: "Query to be executed. Please see https://github.com/askgitdev/askgit for more details regarding syntax",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"test": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"helm": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"auth": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"password": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
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

												"value_from": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"helm_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"service_account": schema.StringAttribute{
															Description:         "ServiceAccount specifies the service account whose token should be fetched",
															MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

										"username": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
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

												"value_from": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"helm_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"service_account": schema.StringAttribute{
															Description:         "ServiceAccount specifies the service account whose token should be fetched",
															MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"cafile": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"chartmuseum": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"project": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"http": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"body": schema.StringAttribute{
									Description:         "Request Body Contents",
									MarkdownDescription: "Request Body Contents",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"connection": schema.StringAttribute{
									Description:         "Connection name e.g. connection://http/google",
									MarkdownDescription: "Connection name e.g. connection://http/google",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"display": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"endpoint": schema.StringAttribute{
									Description:         "Deprecated: Use url instead",
									MarkdownDescription: "Deprecated: Use url instead",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"env": schema.ListNestedAttribute{
									Description:         "EnvVars are the environment variables that are accesible to templated body",
									MarkdownDescription: "EnvVars are the environment variables that are accesible to templated body",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
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

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"config_map_key_ref": schema.SingleNestedAttribute{
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
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"helm_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
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
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"service_account": schema.StringAttribute{
														Description:         "ServiceAccount specifies the service account whose token should be fetched",
														MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"headers": schema.ListNestedAttribute{
									Description:         "Header fields to be used in the query",
									MarkdownDescription: "Header fields to be used in the query",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
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

											"value_from": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"config_map_key_ref": schema.SingleNestedAttribute{
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
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"helm_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
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
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"service_account": schema.StringAttribute{
														Description:         "ServiceAccount specifies the service account whose token should be fetched",
														MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_ssl_expiry": schema.Int64Attribute{
									Description:         "Maximum number of days until the SSL Certificate expires.",
									MarkdownDescription: "Maximum number of days until the SSL Certificate expires.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"method": schema.StringAttribute{
									Description:         "Method to use - defaults to GET",
									MarkdownDescription: "Method to use - defaults to GET",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ntlm": schema.BoolAttribute{
									Description:         "NTLM when set to true will do authentication using NTLM v1 protocol",
									MarkdownDescription: "NTLM when set to true will do authentication using NTLM v1 protocol",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ntlmv2": schema.BoolAttribute{
									Description:         "NTLM when set to true will do authentication using NTLM v2 protocol",
									MarkdownDescription: "NTLM when set to true will do authentication using NTLM v2 protocol",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"oauth2": schema.SingleNestedAttribute{
									Description:         "Oauth2 Configuration. The client ID & Client secret should go to username & password respectively.",
									MarkdownDescription: "Oauth2 Configuration. The client ID & Client secret should go to username & password respectively.",
									Attributes: map[string]schema.Attribute{
										"params": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"scope": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"token_url": schema.StringAttribute{
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

								"password": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"response_codes": schema.ListAttribute{
									Description:         "Expected response codes for the HTTP Request.",
									MarkdownDescription: "Expected response codes for the HTTP Request.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"response_content": schema.StringAttribute{
									Description:         "Exact response content expected to be returned by the endpoint.",
									MarkdownDescription: "Exact response content expected to be returned by the endpoint.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"response_json_content": schema.SingleNestedAttribute{
									Description:         "Deprecated, use expr and jsonpath function",
									MarkdownDescription: "Deprecated, use expr and jsonpath function",
									Attributes: map[string]schema.Attribute{
										"path": schema.StringAttribute{
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"template_body": schema.BoolAttribute{
									Description:         "Template the request body",
									MarkdownDescription: "Template the request body",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"test": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"threshold_millis": schema.Int64Attribute{
									Description:         "Maximum duration in milliseconds for the HTTP request. It will fail the check if it takes longer.",
									MarkdownDescription: "Maximum duration in milliseconds for the HTTP request. It will fail the check if it takes longer.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"transform": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"url": schema.StringAttribute{
									Description:         "Connection url, interpolated with username,password",
									MarkdownDescription: "Connection url, interpolated with username,password",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"username": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

					"icmp": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
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

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"packet_count": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"packet_loss_threshold": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"threshold_millis": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"icon": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"interval": schema.Int64Attribute{
						Description:         "interval (in seconds) to run checks on Deprecated in favor of Schedule",
						MarkdownDescription: "interval (in seconds) to run checks on Deprecated in favor of Schedule",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"jmeter": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"host": schema.StringAttribute{
									Description:         "Host is the server against which test plan needs to be executed",
									MarkdownDescription: "Host is the server against which test plan needs to be executed",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"jmx": schema.SingleNestedAttribute{
									Description:         "Jmx defines the ConfigMap or Secret reference to get the JMX test plan",
									MarkdownDescription: "Jmx defines the ConfigMap or Secret reference to get the JMX test plan",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"port": schema.Int64Attribute{
									Description:         "Port on which the server is running",
									MarkdownDescription: "Port on which the server is running",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"properties": schema.ListAttribute{
									Description:         "Properties defines the local Jmeter properties",
									MarkdownDescription: "Properties defines the local Jmeter properties",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"response_duration": schema.StringAttribute{
									Description:         "ResponseDuration under which the all the test should pass",
									MarkdownDescription: "ResponseDuration under which the all the test should pass",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"system_properties": schema.ListAttribute{
									Description:         "SystemProperties defines the java system property",
									MarkdownDescription: "SystemProperties defines the java system property",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"junit": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"artifacts": schema.ListNestedAttribute{
									Description:         "Artifacts configure the artifacts generated by the check",
									MarkdownDescription: "Artifacts configure the artifacts generated by the check",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"path": schema.StringAttribute{
												Description:         "Path to the artifact on the check runner.Special paths: /dev/stdout & /dev/stdin",
												MarkdownDescription: "Path to the artifact on the check runner.Special paths: /dev/stdout & /dev/stdin",
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

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"display": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"spec": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"test": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"test_results": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"timeout": schema.Int64Attribute{
									Description:         "Timeout in minutes to wait for specified container to finish its job. Defaults to 5 minutes",
									MarkdownDescription: "Timeout in minutes to wait for specified container to finish its job. Defaults to 5 minutes",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"transform": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"kubernetes": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"display": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ignore": schema.ListAttribute{
									Description:         "Ignore the specified resources from the fetched resources. Can be a glob pattern.",
									MarkdownDescription: "Ignore the specified resources from the fetched resources. Can be a glob pattern.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"kubeconfig": schema.SingleNestedAttribute{
									Description:         "KubeConfig is the kubeconfig or the path to the kubeconfig file.",
									MarkdownDescription: "KubeConfig is the kubeconfig or the path to the kubeconfig file.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"namespace_selector": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"field_selector": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"label_selector": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"ready": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"resource": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"field_selector": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"label_selector": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"test": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"ldap": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"bind_dn": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"connection": schema.StringAttribute{
									Description:         "Connection name e.g. connection://http/google",
									MarkdownDescription: "Connection name e.g. connection://http/google",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"password": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"skip_tls_verify": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"url": schema.StringAttribute{
									Description:         "Connection url, interpolated with username,password",
									MarkdownDescription: "Connection url, interpolated with username,password",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"user_search": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"username": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

					"mongodb": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"connection": schema.StringAttribute{
									Description:         "Connection name e.g. connection://http/google",
									MarkdownDescription: "Connection name e.g. connection://http/google",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"password": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"url": schema.StringAttribute{
									Description:         "Connection url, interpolated with username,password",
									MarkdownDescription: "Connection url, interpolated with username,password",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"username": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

					"mssql": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"connection": schema.StringAttribute{
									Description:         "Connection name e.g. connection://http/google",
									MarkdownDescription: "Connection name e.g. connection://http/google",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"display": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"password": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"query": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"results": schema.Int64Attribute{
									Description:         "Number rows to check for",
									MarkdownDescription: "Number rows to check for",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"test": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"url": schema.StringAttribute{
									Description:         "Connection url, interpolated with username,password",
									MarkdownDescription: "Connection url, interpolated with username,password",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"username": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

					"mysql": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"connection": schema.StringAttribute{
									Description:         "Connection name e.g. connection://http/google",
									MarkdownDescription: "Connection name e.g. connection://http/google",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"display": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"password": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"query": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"results": schema.Int64Attribute{
									Description:         "Number rows to check for",
									MarkdownDescription: "Number rows to check for",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"test": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"url": schema.StringAttribute{
									Description:         "Connection url, interpolated with username,password",
									MarkdownDescription: "Connection url, interpolated with username,password",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"username": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

					"namespace": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"deadline": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"delete_timeout": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"expected_content": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"expected_http_statuses": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"http_retry_interval": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"http_timeout": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ingress_host": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ingress_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ingress_timeout": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"namespace_annotations": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"namespace_labels": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"namespace_name_prefix": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"path": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"pod_spec": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"port": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"priority_class": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ready_timeout": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"schedule_timeout": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"opensearch": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"connection": schema.StringAttribute{
									Description:         "Connection name e.g. connection://http/google",
									MarkdownDescription: "Connection name e.g. connection://http/google",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"display": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"index": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"password": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"query": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"results": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"test": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"url": schema.StringAttribute{
									Description:         "Connection url, interpolated with username,password",
									MarkdownDescription: "Connection url, interpolated with username,password",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"username": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

					"owner": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pod": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"deadline": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"delete_timeout": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"expected_content": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"expected_http_statuses": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"http_retry_interval": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"http_timeout": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ingress_class": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ingress_host": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ingress_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ingress_timeout": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"path": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"port": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"priority_class": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ready_timeout": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"round_robin_nodes": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"schedule_timeout": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"spec": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"postgres": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"connection": schema.StringAttribute{
									Description:         "Connection name e.g. connection://http/google",
									MarkdownDescription: "Connection name e.g. connection://http/google",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"display": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"password": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"query": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"results": schema.Int64Attribute{
									Description:         "Number rows to check for",
									MarkdownDescription: "Number rows to check for",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"test": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"url": schema.StringAttribute{
									Description:         "Connection url, interpolated with username,password",
									MarkdownDescription: "Connection url, interpolated with username,password",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"username": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

					"prometheus": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"connection": schema.StringAttribute{
									Description:         "Connection name e.g. connection://http/google",
									MarkdownDescription: "Connection name e.g. connection://http/google",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"display": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"host": schema.StringAttribute{
									Description:         "Deprecated: use 'url' instead",
									MarkdownDescription: "Deprecated: use 'url' instead",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"password": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"query": schema.StringAttribute{
									Description:         "PromQL query",
									MarkdownDescription: "PromQL query",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"test": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"expr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"javascript": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"url": schema.StringAttribute{
									Description:         "Connection url, interpolated with username,password",
									MarkdownDescription: "Connection url, interpolated with username,password",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"username": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

					"redis": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"addr": schema.StringAttribute{
									Description:         "Deprecated: Use url instead",
									MarkdownDescription: "Deprecated: Use url instead",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"connection": schema.StringAttribute{
									Description:         "Connection name e.g. connection://http/google",
									MarkdownDescription: "Connection name e.g. connection://http/google",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"db": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"password": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"url": schema.StringAttribute{
									Description:         "Connection url, interpolated with username,password",
									MarkdownDescription: "Connection url, interpolated with username,password",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"username": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

					"replicas": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"restic": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"access_key": schema.SingleNestedAttribute{
									Description:         "AccessKey access key id for connection with aws s3, minio, wasabi, alibaba oss",
									MarkdownDescription: "AccessKey access key id for connection with aws s3, minio, wasabi, alibaba oss",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"aws_connection_name": schema.StringAttribute{
									Description:         "Name of the AWS connection used to derive the access key and secret key.",
									MarkdownDescription: "Name of the AWS connection used to derive the access key and secret key.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ca_cert": schema.StringAttribute{
									Description:         "CaCert path to the root cert. In case of self-signed certificates",
									MarkdownDescription: "CaCert path to the root cert. In case of self-signed certificates",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"check_integrity": schema.BoolAttribute{
									Description:         "CheckIntegrity when enabled will check the Integrity and consistency of the restic reposiotry",
									MarkdownDescription: "CheckIntegrity when enabled will check the Integrity and consistency of the restic reposiotry",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"connection": schema.StringAttribute{
									Description:         "Name of the connection used to derive restic password.",
									MarkdownDescription: "Name of the connection used to derive restic password.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_age": schema.StringAttribute{
									Description:         "MaxAge for backup freshness",
									MarkdownDescription: "MaxAge for backup freshness",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"password": schema.SingleNestedAttribute{
									Description:         "Password for the restic repository",
									MarkdownDescription: "Password for the restic repository",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"repository": schema.StringAttribute{
									Description:         "Repository The restic repository path eg: rest:https://user:pass@host:8000/ or rest:https://host:8000/ or s3:s3.amazonaws.com/bucket_name",
									MarkdownDescription: "Repository The restic repository path eg: rest:https://user:pass@host:8000/ or rest:https://host:8000/ or s3:s3.amazonaws.com/bucket_name",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"secret_key": schema.SingleNestedAttribute{
									Description:         "SecretKey secret access key for connection with aws s3, minio, wasabi, alibaba oss",
									MarkdownDescription: "SecretKey secret access key for connection with aws s3, minio, wasabi, alibaba oss",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"result_mode": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"s3": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"access_key": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"bucket": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"bucket_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"connection": schema.StringAttribute{
									Description:         "ConnectionName of the connection. It'll be used to populate the endpoint, accessKey and secretKey.",
									MarkdownDescription: "ConnectionName of the connection. It'll be used to populate the endpoint, accessKey and secretKey.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
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

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"object_path": schema.StringAttribute{
									Description:         "glob path to restrict matches to a subset",
									MarkdownDescription: "glob path to restrict matches to a subset",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"region": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"secret_key": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"session_token": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

										"value_from": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"helm_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_account": schema.StringAttribute{
													Description:         "ServiceAccount specifies the service account whose token should be fetched",
													MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

								"skip_tls_verify": schema.BoolAttribute{
									Description:         "Skip TLS verify when connecting to aws",
									MarkdownDescription: "Skip TLS verify when connecting to aws",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"use_path_style": schema.BoolAttribute{
									Description:         "Use path style path: http://s3.amazonaws.com/BUCKET/KEY instead of http://BUCKET.s3.amazonaws.com/KEY",
									MarkdownDescription: "Use path style path: http://s3.amazonaws.com/BUCKET/KEY instead of http://BUCKET.s3.amazonaws.com/KEY",
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

					"schedule": schema.StringAttribute{
						Description:         "Schedule to run checks on. Supports all cron expression, example: '30 3-6,20-23 * * *'. For more info about cron expression syntax see https://en.wikipedia.org/wiki/Cron Also supports golang duration, can be set as '@every 1m30s' which runs the check every 1 minute and 30 seconds.",
						MarkdownDescription: "Schedule to run checks on. Supports all cron expression, example: '30 3-6,20-23 * * *'. For more info about cron expression syntax see https://en.wikipedia.org/wiki/Cron Also supports golang duration, can be set as '@every 1m30s' which runs the check every 1 minute and 30 seconds.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"severity": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tcp": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"description": schema.StringAttribute{
									Description:         "Description for the check",
									MarkdownDescription: "Description for the check",
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

								"icon": schema.StringAttribute{
									Description:         "Icon for overwriting default icon on the dashboard",
									MarkdownDescription: "Icon for overwriting default icon on the dashboard",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels for the check",
									MarkdownDescription: "Labels for the check",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metrics": schema.ListNestedAttribute{
									Description:         "Metrics to expose from check results",
									MarkdownDescription: "Metrics to expose from check results",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
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

														"value_expr": schema.StringAttribute{
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

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "Name of the check",
									MarkdownDescription: "Name of the check",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"threshold_millis": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"transform_delete_strategy": schema.StringAttribute{
									Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
									MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

					"webhook": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"description": schema.StringAttribute{
								Description:         "Description for the check",
								MarkdownDescription: "Description for the check",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"display": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"expr": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"javascript": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"json_path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"template": schema.StringAttribute{
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

							"icon": schema.StringAttribute{
								Description:         "Icon for overwriting default icon on the dashboard",
								MarkdownDescription: "Icon for overwriting default icon on the dashboard",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "Labels for the check",
								MarkdownDescription: "Labels for the check",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metrics": schema.ListNestedAttribute{
								Description:         "Metrics to expose from check results",
								MarkdownDescription: "Metrics to expose from check results",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"labels": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
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

													"value_expr": schema.StringAttribute{
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

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"type": schema.StringAttribute{
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

							"name": schema.StringAttribute{
								Description:         "Name of the check",
								MarkdownDescription: "Name of the check",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
								MarkdownDescription: "Namespace to insert the check into, if different to the namespace the canary is defined, e.g.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"test": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"expr": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"javascript": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"json_path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"template": schema.StringAttribute{
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

							"token": schema.SingleNestedAttribute{
								Description:         "Token is an optional authorization token to run this check",
								MarkdownDescription: "Token is an optional authorization token to run this check",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
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

									"value_from": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"config_map_key_ref": schema.SingleNestedAttribute{
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
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"helm_ref": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "Key is a JSONPath expression used to fetch the key from the merged JSON.",
														MarkdownDescription: "Key is a JSONPath expression used to fetch the key from the merged JSON.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
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
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"service_account": schema.StringAttribute{
												Description:         "ServiceAccount specifies the service account whose token should be fetched",
												MarkdownDescription: "ServiceAccount specifies the service account whose token should be fetched",
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

							"transform": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"expr": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"javascript": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"json_path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"template": schema.StringAttribute{
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

							"transform_delete_strategy": schema.StringAttribute{
								Description:         "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
								MarkdownDescription: "Transformed checks have a delete strategy on deletion they can either be marked healthy, unhealthy or left as is",
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

func (r *CanariesFlanksourceComCanaryV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_canaries_flanksource_com_canary_v1_manifest")

	var model CanariesFlanksourceComCanaryV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("canaries.flanksource.com/v1")
	model.Kind = pointer.String("Canary")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}