/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

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

type CapsuleClastixIoTenantV1Beta1Resource struct{}

var (
	_ resource.Resource = (*CapsuleClastixIoTenantV1Beta1Resource)(nil)
)

type CapsuleClastixIoTenantV1Beta1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CapsuleClastixIoTenantV1Beta1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		AdditionalRoleBindings *[]struct {
			ClusterRoleName *string `tfsdk:"cluster_role_name" yaml:"clusterRoleName,omitempty"`

			Subjects *[]struct {
				ApiGroup *string `tfsdk:"api_group" yaml:"apiGroup,omitempty"`

				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"subjects" yaml:"subjects,omitempty"`
		} `tfsdk:"additional_role_bindings" yaml:"additionalRoleBindings,omitempty"`

		ContainerRegistries *struct {
			Allowed *[]string `tfsdk:"allowed" yaml:"allowed,omitempty"`

			AllowedRegex *string `tfsdk:"allowed_regex" yaml:"allowedRegex,omitempty"`
		} `tfsdk:"container_registries" yaml:"containerRegistries,omitempty"`

		ImagePullPolicies *[]string `tfsdk:"image_pull_policies" yaml:"imagePullPolicies,omitempty"`

		IngressOptions *struct {
			AllowedClasses *struct {
				Allowed *[]string `tfsdk:"allowed" yaml:"allowed,omitempty"`

				AllowedRegex *string `tfsdk:"allowed_regex" yaml:"allowedRegex,omitempty"`
			} `tfsdk:"allowed_classes" yaml:"allowedClasses,omitempty"`

			AllowedHostnames *struct {
				Allowed *[]string `tfsdk:"allowed" yaml:"allowed,omitempty"`

				AllowedRegex *string `tfsdk:"allowed_regex" yaml:"allowedRegex,omitempty"`
			} `tfsdk:"allowed_hostnames" yaml:"allowedHostnames,omitempty"`

			HostnameCollisionScope *string `tfsdk:"hostname_collision_scope" yaml:"hostnameCollisionScope,omitempty"`
		} `tfsdk:"ingress_options" yaml:"ingressOptions,omitempty"`

		LimitRanges *struct {
			Items *[]struct {
				Limits *[]struct {
					Default *map[string]string `tfsdk:"default" yaml:"default,omitempty"`

					DefaultRequest *map[string]string `tfsdk:"default_request" yaml:"defaultRequest,omitempty"`

					Max *map[string]string `tfsdk:"max" yaml:"max,omitempty"`

					MaxLimitRequestRatio *map[string]string `tfsdk:"max_limit_request_ratio" yaml:"maxLimitRequestRatio,omitempty"`

					Min *map[string]string `tfsdk:"min" yaml:"min,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"limits" yaml:"limits,omitempty"`
			} `tfsdk:"items" yaml:"items,omitempty"`
		} `tfsdk:"limit_ranges" yaml:"limitRanges,omitempty"`

		NamespaceOptions *struct {
			AdditionalMetadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
			} `tfsdk:"additional_metadata" yaml:"additionalMetadata,omitempty"`

			Quota *int64 `tfsdk:"quota" yaml:"quota,omitempty"`
		} `tfsdk:"namespace_options" yaml:"namespaceOptions,omitempty"`

		NetworkPolicies *struct {
			Items *[]struct {
				Egress *[]struct {
					Ports *[]struct {
						EndPort *int64 `tfsdk:"end_port" yaml:"endPort,omitempty"`

						Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

						Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
					} `tfsdk:"ports" yaml:"ports,omitempty"`

					To *[]struct {
						IpBlock *struct {
							Cidr *string `tfsdk:"cidr" yaml:"cidr,omitempty"`

							Except *[]string `tfsdk:"except" yaml:"except,omitempty"`
						} `tfsdk:"ip_block" yaml:"ipBlock,omitempty"`

						NamespaceSelector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

						PodSelector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"pod_selector" yaml:"podSelector,omitempty"`
					} `tfsdk:"to" yaml:"to,omitempty"`
				} `tfsdk:"egress" yaml:"egress,omitempty"`

				Ingress *[]struct {
					From *[]struct {
						IpBlock *struct {
							Cidr *string `tfsdk:"cidr" yaml:"cidr,omitempty"`

							Except *[]string `tfsdk:"except" yaml:"except,omitempty"`
						} `tfsdk:"ip_block" yaml:"ipBlock,omitempty"`

						NamespaceSelector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

						PodSelector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"pod_selector" yaml:"podSelector,omitempty"`
					} `tfsdk:"from" yaml:"from,omitempty"`

					Ports *[]struct {
						EndPort *int64 `tfsdk:"end_port" yaml:"endPort,omitempty"`

						Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

						Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
					} `tfsdk:"ports" yaml:"ports,omitempty"`
				} `tfsdk:"ingress" yaml:"ingress,omitempty"`

				PodSelector *struct {
					MatchExpressions *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

					MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
				} `tfsdk:"pod_selector" yaml:"podSelector,omitempty"`

				PolicyTypes *[]string `tfsdk:"policy_types" yaml:"policyTypes,omitempty"`
			} `tfsdk:"items" yaml:"items,omitempty"`
		} `tfsdk:"network_policies" yaml:"networkPolicies,omitempty"`

		NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

		Owners *[]struct {
			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			ProxySettings *[]struct {
				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				Operations *[]string `tfsdk:"operations" yaml:"operations,omitempty"`
			} `tfsdk:"proxy_settings" yaml:"proxySettings,omitempty"`
		} `tfsdk:"owners" yaml:"owners,omitempty"`

		PriorityClasses *struct {
			Allowed *[]string `tfsdk:"allowed" yaml:"allowed,omitempty"`

			AllowedRegex *string `tfsdk:"allowed_regex" yaml:"allowedRegex,omitempty"`
		} `tfsdk:"priority_classes" yaml:"priorityClasses,omitempty"`

		ResourceQuotas *struct {
			Items *[]struct {
				Hard *map[string]string `tfsdk:"hard" yaml:"hard,omitempty"`

				ScopeSelector *struct {
					MatchExpressions *[]struct {
						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						ScopeName *string `tfsdk:"scope_name" yaml:"scopeName,omitempty"`

						Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`
				} `tfsdk:"scope_selector" yaml:"scopeSelector,omitempty"`

				Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`
			} `tfsdk:"items" yaml:"items,omitempty"`

			Scope *string `tfsdk:"scope" yaml:"scope,omitempty"`
		} `tfsdk:"resource_quotas" yaml:"resourceQuotas,omitempty"`

		ServiceOptions *struct {
			AdditionalMetadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
			} `tfsdk:"additional_metadata" yaml:"additionalMetadata,omitempty"`

			AllowedServices *struct {
				ExternalName *bool `tfsdk:"external_name" yaml:"externalName,omitempty"`

				LoadBalancer *bool `tfsdk:"load_balancer" yaml:"loadBalancer,omitempty"`

				NodePort *bool `tfsdk:"node_port" yaml:"nodePort,omitempty"`
			} `tfsdk:"allowed_services" yaml:"allowedServices,omitempty"`

			ExternalIPs *struct {
				Allowed *[]string `tfsdk:"allowed" yaml:"allowed,omitempty"`
			} `tfsdk:"external_i_ps" yaml:"externalIPs,omitempty"`
		} `tfsdk:"service_options" yaml:"serviceOptions,omitempty"`

		StorageClasses *struct {
			Allowed *[]string `tfsdk:"allowed" yaml:"allowed,omitempty"`

			AllowedRegex *string `tfsdk:"allowed_regex" yaml:"allowedRegex,omitempty"`
		} `tfsdk:"storage_classes" yaml:"storageClasses,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCapsuleClastixIoTenantV1Beta1Resource() resource.Resource {
	return &CapsuleClastixIoTenantV1Beta1Resource{}
}

func (r *CapsuleClastixIoTenantV1Beta1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_capsule_clastix_io_tenant_v1beta1"
}

func (r *CapsuleClastixIoTenantV1Beta1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Tenant is the Schema for the tenants API.",
		MarkdownDescription: "Tenant is the Schema for the tenants API.",
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
				Description:         "TenantSpec defines the desired state of Tenant.",
				MarkdownDescription: "TenantSpec defines the desired state of Tenant.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"additional_role_bindings": {
						Description:         "Specifies additional RoleBindings assigned to the Tenant. Capsule will ensure that all namespaces in the Tenant always contain the RoleBinding for the given ClusterRole. Optional.",
						MarkdownDescription: "Specifies additional RoleBindings assigned to the Tenant. Capsule will ensure that all namespaces in the Tenant always contain the RoleBinding for the given ClusterRole. Optional.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"cluster_role_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"subjects": {
								Description:         "kubebuilder:validation:Minimum=1",
								MarkdownDescription: "kubebuilder:validation:Minimum=1",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"api_group": {
										Description:         "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
										MarkdownDescription: "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kind": {
										Description:         "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
										MarkdownDescription: "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": {
										Description:         "Name of the object being referenced.",
										MarkdownDescription: "Name of the object being referenced.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"namespace": {
										Description:         "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",
										MarkdownDescription: "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"container_registries": {
						Description:         "Specifies the trusted Image Registries assigned to the Tenant. Capsule assures that all Pods resources created in the Tenant can use only one of the allowed trusted registries. Optional.",
						MarkdownDescription: "Specifies the trusted Image Registries assigned to the Tenant. Capsule assures that all Pods resources created in the Tenant can use only one of the allowed trusted registries. Optional.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"allowed": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"allowed_regex": {
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

					"image_pull_policies": {
						Description:         "Specify the allowed values for the imagePullPolicies option in Pod resources. Capsule assures that all Pod resources created in the Tenant can use only one of the allowed policy. Optional.",
						MarkdownDescription: "Specify the allowed values for the imagePullPolicies option in Pod resources. Capsule assures that all Pod resources created in the Tenant can use only one of the allowed policy. Optional.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress_options": {
						Description:         "Specifies options for the Ingress resources, such as allowed hostnames and IngressClass. Optional.",
						MarkdownDescription: "Specifies options for the Ingress resources, such as allowed hostnames and IngressClass. Optional.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"allowed_classes": {
								Description:         "Specifies the allowed IngressClasses assigned to the Tenant. Capsule assures that all Ingress resources created in the Tenant can use only one of the allowed IngressClasses. Optional.",
								MarkdownDescription: "Specifies the allowed IngressClasses assigned to the Tenant. Capsule assures that all Ingress resources created in the Tenant can use only one of the allowed IngressClasses. Optional.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"allowed": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"allowed_regex": {
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

							"allowed_hostnames": {
								Description:         "Specifies the allowed hostnames in Ingresses for the given Tenant. Capsule assures that all Ingress resources created in the Tenant can use only one of the allowed hostnames. Optional.",
								MarkdownDescription: "Specifies the allowed hostnames in Ingresses for the given Tenant. Capsule assures that all Ingress resources created in the Tenant can use only one of the allowed hostnames. Optional.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"allowed": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"allowed_regex": {
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

							"hostname_collision_scope": {
								Description:         "Defines the scope of hostname collision check performed when Tenant Owners create Ingress with allowed hostnames.  - Cluster: disallow the creation of an Ingress if the pair hostname and path is already used across the Namespaces managed by Capsule.  - Tenant: disallow the creation of an Ingress if the pair hostname and path is already used across the Namespaces of the Tenant.  - Namespace: disallow the creation of an Ingress if the pair hostname and path is already used in the Ingress Namespace.  Optional.",
								MarkdownDescription: "Defines the scope of hostname collision check performed when Tenant Owners create Ingress with allowed hostnames.  - Cluster: disallow the creation of an Ingress if the pair hostname and path is already used across the Namespaces managed by Capsule.  - Tenant: disallow the creation of an Ingress if the pair hostname and path is already used across the Namespaces of the Tenant.  - Namespace: disallow the creation of an Ingress if the pair hostname and path is already used in the Ingress Namespace.  Optional.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("Cluster", "Tenant", "Namespace", "Disabled"),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"limit_ranges": {
						Description:         "Specifies the resource min/max usage restrictions to the Tenant. The assigned values are inherited by any namespace created in the Tenant. Optional.",
						MarkdownDescription: "Specifies the resource min/max usage restrictions to the Tenant. The assigned values are inherited by any namespace created in the Tenant. Optional.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"items": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"limits": {
										Description:         "Limits is the list of LimitRangeItem objects that are enforced.",
										MarkdownDescription: "Limits is the list of LimitRangeItem objects that are enforced.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"default": {
												Description:         "Default resource requirement limit value by resource name if resource limit is omitted.",
												MarkdownDescription: "Default resource requirement limit value by resource name if resource limit is omitted.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"default_request": {
												Description:         "DefaultRequest is the default resource requirement request value by resource name if resource request is omitted.",
												MarkdownDescription: "DefaultRequest is the default resource requirement request value by resource name if resource request is omitted.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max": {
												Description:         "Max usage constraints on this kind by resource name.",
												MarkdownDescription: "Max usage constraints on this kind by resource name.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_limit_request_ratio": {
												Description:         "MaxLimitRequestRatio if specified, the named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource.",
												MarkdownDescription: "MaxLimitRequestRatio if specified, the named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"min": {
												Description:         "Min usage constraints on this kind by resource name.",
												MarkdownDescription: "Min usage constraints on this kind by resource name.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
												Description:         "Type of resource that this limit applies to.",
												MarkdownDescription: "Type of resource that this limit applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

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

					"namespace_options": {
						Description:         "Specifies options for the Namespaces, such as additional metadata or maximum number of namespaces allowed for that Tenant. Once the namespace quota assigned to the Tenant has been reached, the Tenant owner cannot create further namespaces. Optional.",
						MarkdownDescription: "Specifies options for the Namespaces, such as additional metadata or maximum number of namespaces allowed for that Tenant. Once the namespace quota assigned to the Tenant has been reached, the Tenant owner cannot create further namespaces. Optional.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"additional_metadata": {
								Description:         "Specifies additional labels and annotations the Capsule operator places on any Namespace resource in the Tenant. Optional.",
								MarkdownDescription: "Specifies additional labels and annotations the Capsule operator places on any Namespace resource in the Tenant. Optional.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"quota": {
								Description:         "Specifies the maximum number of namespaces allowed for that Tenant. Once the namespace quota assigned to the Tenant has been reached, the Tenant owner cannot create further namespaces. Optional.",
								MarkdownDescription: "Specifies the maximum number of namespaces allowed for that Tenant. Once the namespace quota assigned to the Tenant has been reached, the Tenant owner cannot create further namespaces. Optional.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(1),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"network_policies": {
						Description:         "Specifies the NetworkPolicies assigned to the Tenant. The assigned NetworkPolicies are inherited by any namespace created in the Tenant. Optional.",
						MarkdownDescription: "Specifies the NetworkPolicies assigned to the Tenant. The assigned NetworkPolicies are inherited by any namespace created in the Tenant. Optional.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"items": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"egress": {
										Description:         "List of egress rules to be applied to the selected pods. Outgoing traffic is allowed if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic matches at least one egress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy limits all outgoing traffic (and serves solely to ensure that the pods it selects are isolated by default). This field is beta-level in 1.8",
										MarkdownDescription: "List of egress rules to be applied to the selected pods. Outgoing traffic is allowed if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic matches at least one egress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy limits all outgoing traffic (and serves solely to ensure that the pods it selects are isolated by default). This field is beta-level in 1.8",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"ports": {
												Description:         "List of destination ports for outgoing traffic. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.",
												MarkdownDescription: "List of destination ports for outgoing traffic. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"end_port": {
														Description:         "If set, indicates that the range of ports from port to endPort, inclusive, should be allowed by the policy. This field cannot be defined if the port field is not defined or if the port field is defined as a named (string) port. The endPort must be equal or greater than port. This feature is in Beta state and is enabled by default. It can be disabled using the Feature Gate 'NetworkPolicyEndPort'.",
														MarkdownDescription: "If set, indicates that the range of ports from port to endPort, inclusive, should be allowed by the policy. This field cannot be defined if the port field is not defined or if the port field is defined as a named (string) port. The endPort must be equal or greater than port. This feature is in Beta state and is enabled by default. It can be disabled using the Feature Gate 'NetworkPolicyEndPort'.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"port": {
														Description:         "The port on the given protocol. This can either be a numerical or named port on a pod. If this field is not provided, this matches all port names and numbers. If present, only traffic on the specified protocol AND port will be matched.",
														MarkdownDescription: "The port on the given protocol. This can either be a numerical or named port on a pod. If this field is not provided, this matches all port names and numbers. If present, only traffic on the specified protocol AND port will be matched.",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"protocol": {
														Description:         "The protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
														MarkdownDescription: "The protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",

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

											"to": {
												Description:         "List of destinations for outgoing traffic of pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all destinations (traffic not restricted by destination). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the to list.",
												MarkdownDescription: "List of destinations for outgoing traffic of pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all destinations (traffic not restricted by destination). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the to list.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"ip_block": {
														Description:         "IPBlock defines policy on a particular IPBlock. If this field is set then neither of the other fields can be.",
														MarkdownDescription: "IPBlock defines policy on a particular IPBlock. If this field is set then neither of the other fields can be.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"cidr": {
																Description:         "CIDR is a string representing the IP Block Valid examples are '192.168.1.1/24' or '2001:db9::/64'",
																MarkdownDescription: "CIDR is a string representing the IP Block Valid examples are '192.168.1.1/24' or '2001:db9::/64'",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"except": {
																Description:         "Except is a slice of CIDRs that should not be included within an IP Block Valid examples are '192.168.1.1/24' or '2001:db9::/64' Except values will be rejected if they are outside the CIDR range",
																MarkdownDescription: "Except is a slice of CIDRs that should not be included within an IP Block Valid examples are '192.168.1.1/24' or '2001:db9::/64' Except values will be rejected if they are outside the CIDR range",

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

													"namespace_selector": {
														Description:         "Selects Namespaces using cluster-scoped labels. This field follows standard label selector semantics; if present but empty, it selects all namespaces.  If PodSelector is also set, then the NetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects all Pods in the Namespaces selected by NamespaceSelector.",
														MarkdownDescription: "Selects Namespaces using cluster-scoped labels. This field follows standard label selector semantics; if present but empty, it selects all namespaces.  If PodSelector is also set, then the NetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects all Pods in the Namespaces selected by NamespaceSelector.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"match_expressions": {
																Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"operator": {
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

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

															"match_labels": {
																Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

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

													"pod_selector": {
														Description:         "This is a label selector which selects Pods. This field follows standard label selector semantics; if present but empty, it selects all pods.  If NamespaceSelector is also set, then the NetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects the Pods matching PodSelector in the policy's own Namespace.",
														MarkdownDescription: "This is a label selector which selects Pods. This field follows standard label selector semantics; if present but empty, it selects all pods.  If NamespaceSelector is also set, then the NetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects the Pods matching PodSelector in the policy's own Namespace.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"match_expressions": {
																Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"operator": {
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

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

															"match_labels": {
																Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

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

									"ingress": {
										Description:         "List of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic source is the pod's local node, OR if the traffic matches at least one ingress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy does not allow any traffic (and serves solely to ensure that the pods it selects are isolated by default)",
										MarkdownDescription: "List of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic source is the pod's local node, OR if the traffic matches at least one ingress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy does not allow any traffic (and serves solely to ensure that the pods it selects are isolated by default)",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"from": {
												Description:         "List of sources which should be able to access the pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all sources (traffic not restricted by source). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the from list.",
												MarkdownDescription: "List of sources which should be able to access the pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all sources (traffic not restricted by source). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the from list.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"ip_block": {
														Description:         "IPBlock defines policy on a particular IPBlock. If this field is set then neither of the other fields can be.",
														MarkdownDescription: "IPBlock defines policy on a particular IPBlock. If this field is set then neither of the other fields can be.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"cidr": {
																Description:         "CIDR is a string representing the IP Block Valid examples are '192.168.1.1/24' or '2001:db9::/64'",
																MarkdownDescription: "CIDR is a string representing the IP Block Valid examples are '192.168.1.1/24' or '2001:db9::/64'",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"except": {
																Description:         "Except is a slice of CIDRs that should not be included within an IP Block Valid examples are '192.168.1.1/24' or '2001:db9::/64' Except values will be rejected if they are outside the CIDR range",
																MarkdownDescription: "Except is a slice of CIDRs that should not be included within an IP Block Valid examples are '192.168.1.1/24' or '2001:db9::/64' Except values will be rejected if they are outside the CIDR range",

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

													"namespace_selector": {
														Description:         "Selects Namespaces using cluster-scoped labels. This field follows standard label selector semantics; if present but empty, it selects all namespaces.  If PodSelector is also set, then the NetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects all Pods in the Namespaces selected by NamespaceSelector.",
														MarkdownDescription: "Selects Namespaces using cluster-scoped labels. This field follows standard label selector semantics; if present but empty, it selects all namespaces.  If PodSelector is also set, then the NetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects all Pods in the Namespaces selected by NamespaceSelector.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"match_expressions": {
																Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"operator": {
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

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

															"match_labels": {
																Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

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

													"pod_selector": {
														Description:         "This is a label selector which selects Pods. This field follows standard label selector semantics; if present but empty, it selects all pods.  If NamespaceSelector is also set, then the NetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects the Pods matching PodSelector in the policy's own Namespace.",
														MarkdownDescription: "This is a label selector which selects Pods. This field follows standard label selector semantics; if present but empty, it selects all pods.  If NamespaceSelector is also set, then the NetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects the Pods matching PodSelector in the policy's own Namespace.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"match_expressions": {
																Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"operator": {
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

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

															"match_labels": {
																Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ports": {
												Description:         "List of ports which should be made accessible on the pods selected for this rule. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.",
												MarkdownDescription: "List of ports which should be made accessible on the pods selected for this rule. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"end_port": {
														Description:         "If set, indicates that the range of ports from port to endPort, inclusive, should be allowed by the policy. This field cannot be defined if the port field is not defined or if the port field is defined as a named (string) port. The endPort must be equal or greater than port. This feature is in Beta state and is enabled by default. It can be disabled using the Feature Gate 'NetworkPolicyEndPort'.",
														MarkdownDescription: "If set, indicates that the range of ports from port to endPort, inclusive, should be allowed by the policy. This field cannot be defined if the port field is not defined or if the port field is defined as a named (string) port. The endPort must be equal or greater than port. This feature is in Beta state and is enabled by default. It can be disabled using the Feature Gate 'NetworkPolicyEndPort'.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"port": {
														Description:         "The port on the given protocol. This can either be a numerical or named port on a pod. If this field is not provided, this matches all port names and numbers. If present, only traffic on the specified protocol AND port will be matched.",
														MarkdownDescription: "The port on the given protocol. This can either be a numerical or named port on a pod. If this field is not provided, this matches all port names and numbers. If present, only traffic on the specified protocol AND port will be matched.",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"protocol": {
														Description:         "The protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
														MarkdownDescription: "The protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",

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

									"pod_selector": {
										Description:         "Selects the pods to which this NetworkPolicy object applies. The array of ingress rules is applied to any pods selected by this field. Multiple network policies can select the same set of pods. In this case, the ingress rules for each are combined additively. This field is NOT optional and follows standard label selector semantics. An empty podSelector matches all pods in this namespace.",
										MarkdownDescription: "Selects the pods to which this NetworkPolicy object applies. The array of ingress rules is applied to any pods selected by this field. Multiple network policies can select the same set of pods. In this case, the ingress rules for each are combined additively. This field is NOT optional and follows standard label selector semantics. An empty podSelector matches all pods in this namespace.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"match_expressions": {
												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"operator": {
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"values": {
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

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

											"match_labels": {
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"policy_types": {
										Description:         "List of rule types that the NetworkPolicy relates to. Valid options are ['Ingress'], ['Egress'], or ['Ingress', 'Egress']. If this field is not specified, it will default based on the existence of Ingress or Egress rules; policies that contain an Egress section are assumed to affect Egress, and all policies (whether or not they contain an Ingress section) are assumed to affect Ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [ 'Egress' ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a policyTypes value that include 'Egress' (since such a policy would not include an Egress section and would otherwise default to just [ 'Ingress' ]). This field is beta-level in 1.8",
										MarkdownDescription: "List of rule types that the NetworkPolicy relates to. Valid options are ['Ingress'], ['Egress'], or ['Ingress', 'Egress']. If this field is not specified, it will default based on the existence of Ingress or Egress rules; policies that contain an Egress section are assumed to affect Egress, and all policies (whether or not they contain an Ingress section) are assumed to affect Ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [ 'Egress' ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a policyTypes value that include 'Egress' (since such a policy would not include an Egress section and would otherwise default to just [ 'Ingress' ]). This field is beta-level in 1.8",

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

					"node_selector": {
						Description:         "Specifies the label to control the placement of pods on a given pool of worker nodes. All namespaces created within the Tenant will have the node selector annotation. This annotation tells the Kubernetes scheduler to place pods on the nodes having the selector label. Optional.",
						MarkdownDescription: "Specifies the label to control the placement of pods on a given pool of worker nodes. All namespaces created within the Tenant will have the node selector annotation. This annotation tells the Kubernetes scheduler to place pods on the nodes having the selector label. Optional.",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"owners": {
						Description:         "Specifies the owners of the Tenant. Mandatory.",
						MarkdownDescription: "Specifies the owners of the Tenant. Mandatory.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"kind": {
								Description:         "Kind of tenant owner. Possible values are 'User', 'Group', and 'ServiceAccount'",
								MarkdownDescription: "Kind of tenant owner. Possible values are 'User', 'Group', and 'ServiceAccount'",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("User", "Group", "ServiceAccount"),
								},
							},

							"name": {
								Description:         "Name of tenant owner.",
								MarkdownDescription: "Name of tenant owner.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"proxy_settings": {
								Description:         "Proxy settings for tenant owner.",
								MarkdownDescription: "Proxy settings for tenant owner.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"kind": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("Nodes", "StorageClasses", "IngressClasses", "PriorityClasses"),
										},
									},

									"operations": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

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

						Required: true,
						Optional: false,
						Computed: false,
					},

					"priority_classes": {
						Description:         "Specifies the allowed priorityClasses assigned to the Tenant. Capsule assures that all Pods resources created in the Tenant can use only one of the allowed PriorityClasses. Optional.",
						MarkdownDescription: "Specifies the allowed priorityClasses assigned to the Tenant. Capsule assures that all Pods resources created in the Tenant can use only one of the allowed PriorityClasses. Optional.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"allowed": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"allowed_regex": {
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

					"resource_quotas": {
						Description:         "Specifies a list of ResourceQuota resources assigned to the Tenant. The assigned values are inherited by any namespace created in the Tenant. The Capsule operator aggregates ResourceQuota at Tenant level, so that the hard quota is never crossed for the given Tenant. This permits the Tenant owner to consume resources in the Tenant regardless of the namespace. Optional.",
						MarkdownDescription: "Specifies a list of ResourceQuota resources assigned to the Tenant. The assigned values are inherited by any namespace created in the Tenant. The Capsule operator aggregates ResourceQuota at Tenant level, so that the hard quota is never crossed for the given Tenant. This permits the Tenant owner to consume resources in the Tenant regardless of the namespace. Optional.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"items": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"hard": {
										Description:         "hard is the set of desired hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/",
										MarkdownDescription: "hard is the set of desired hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"scope_selector": {
										Description:         "scopeSelector is also a collection of filters like scopes that must match each object tracked by a quota but expressed using ScopeSelectorOperator in combination with possible values. For a resource to match, both scopes AND scopeSelector (if specified in spec), must be matched.",
										MarkdownDescription: "scopeSelector is also a collection of filters like scopes that must match each object tracked by a quota but expressed using ScopeSelectorOperator in combination with possible values. For a resource to match, both scopes AND scopeSelector (if specified in spec), must be matched.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"match_expressions": {
												Description:         "A list of scope selector requirements by scope of the resources.",
												MarkdownDescription: "A list of scope selector requirements by scope of the resources.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"operator": {
														Description:         "Represents a scope's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist.",
														MarkdownDescription: "Represents a scope's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"scope_name": {
														Description:         "The name of the scope that the selector applies to.",
														MarkdownDescription: "The name of the scope that the selector applies to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"values": {
														Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

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

									"scopes": {
										Description:         "A collection of filters that must match each object tracked by a quota. If not specified, the quota matches all objects.",
										MarkdownDescription: "A collection of filters that must match each object tracked by a quota. If not specified, the quota matches all objects.",

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

							"scope": {
								Description:         "Define if the Resource Budget should compute resource across all Namespaces in the Tenant or individually per cluster. Default is Tenant",
								MarkdownDescription: "Define if the Resource Budget should compute resource across all Namespaces in the Tenant or individually per cluster. Default is Tenant",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("Tenant", "Namespace"),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_options": {
						Description:         "Specifies options for the Service, such as additional metadata or block of certain type of Services. Optional.",
						MarkdownDescription: "Specifies options for the Service, such as additional metadata or block of certain type of Services. Optional.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"additional_metadata": {
								Description:         "Specifies additional labels and annotations the Capsule operator places on any Service resource in the Tenant. Optional.",
								MarkdownDescription: "Specifies additional labels and annotations the Capsule operator places on any Service resource in the Tenant. Optional.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"allowed_services": {
								Description:         "Block or deny certain type of Services. Optional.",
								MarkdownDescription: "Block or deny certain type of Services. Optional.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"external_name": {
										Description:         "Specifies if ExternalName service type resources are allowed for the Tenant. Default is true. Optional.",
										MarkdownDescription: "Specifies if ExternalName service type resources are allowed for the Tenant. Default is true. Optional.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"load_balancer": {
										Description:         "Specifies if LoadBalancer service type resources are allowed for the Tenant. Default is true. Optional.",
										MarkdownDescription: "Specifies if LoadBalancer service type resources are allowed for the Tenant. Default is true. Optional.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"node_port": {
										Description:         "Specifies if NodePort service type resources are allowed for the Tenant. Default is true. Optional.",
										MarkdownDescription: "Specifies if NodePort service type resources are allowed for the Tenant. Default is true. Optional.",

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

							"external_i_ps": {
								Description:         "Specifies the external IPs that can be used in Services with type ClusterIP. An empty list means no IPs are allowed. Optional.",
								MarkdownDescription: "Specifies the external IPs that can be used in Services with type ClusterIP. An empty list means no IPs are allowed. Optional.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"allowed": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

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

					"storage_classes": {
						Description:         "Specifies the allowed StorageClasses assigned to the Tenant. Capsule assures that all PersistentVolumeClaim resources created in the Tenant can use only one of the allowed StorageClasses. Optional.",
						MarkdownDescription: "Specifies the allowed StorageClasses assigned to the Tenant. Capsule assures that all PersistentVolumeClaim resources created in the Tenant can use only one of the allowed StorageClasses. Optional.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"allowed": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"allowed_regex": {
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
		},
	}, nil
}

func (r *CapsuleClastixIoTenantV1Beta1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_capsule_clastix_io_tenant_v1beta1")

	var state CapsuleClastixIoTenantV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CapsuleClastixIoTenantV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("capsule.clastix.io/v1beta1")
	goModel.Kind = utilities.Ptr("Tenant")

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

func (r *CapsuleClastixIoTenantV1Beta1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_capsule_clastix_io_tenant_v1beta1")
	// NO-OP: All data is already in Terraform state
}

func (r *CapsuleClastixIoTenantV1Beta1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_capsule_clastix_io_tenant_v1beta1")

	var state CapsuleClastixIoTenantV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CapsuleClastixIoTenantV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("capsule.clastix.io/v1beta1")
	goModel.Kind = utilities.Ptr("Tenant")

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

func (r *CapsuleClastixIoTenantV1Beta1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_capsule_clastix_io_tenant_v1beta1")
	// NO-OP: Terraform removes the state automatically for us
}
