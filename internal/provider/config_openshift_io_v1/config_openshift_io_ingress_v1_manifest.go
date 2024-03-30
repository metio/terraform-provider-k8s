/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package config_openshift_io_v1

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
	_ datasource.DataSource = &ConfigOpenshiftIoIngressV1Manifest{}
)

func NewConfigOpenshiftIoIngressV1Manifest() datasource.DataSource {
	return &ConfigOpenshiftIoIngressV1Manifest{}
}

type ConfigOpenshiftIoIngressV1Manifest struct{}

type ConfigOpenshiftIoIngressV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AppsDomain      *string `tfsdk:"apps_domain" json:"appsDomain,omitempty"`
		ComponentRoutes *[]struct {
			Hostname                 *string `tfsdk:"hostname" json:"hostname,omitempty"`
			Name                     *string `tfsdk:"name" json:"name,omitempty"`
			Namespace                *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ServingCertKeyPairSecret *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"serving_cert_key_pair_secret" json:"servingCertKeyPairSecret,omitempty"`
		} `tfsdk:"component_routes" json:"componentRoutes,omitempty"`
		Domain       *string `tfsdk:"domain" json:"domain,omitempty"`
		LoadBalancer *struct {
			Platform *struct {
				Aws *struct {
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"aws" json:"aws,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"platform" json:"platform,omitempty"`
		} `tfsdk:"load_balancer" json:"loadBalancer,omitempty"`
		RequiredHSTSPolicies *[]struct {
			DomainPatterns          *[]string `tfsdk:"domain_patterns" json:"domainPatterns,omitempty"`
			IncludeSubDomainsPolicy *string   `tfsdk:"include_sub_domains_policy" json:"includeSubDomainsPolicy,omitempty"`
			MaxAge                  *struct {
				LargestMaxAge  *int64 `tfsdk:"largest_max_age" json:"largestMaxAge,omitempty"`
				SmallestMaxAge *int64 `tfsdk:"smallest_max_age" json:"smallestMaxAge,omitempty"`
			} `tfsdk:"max_age" json:"maxAge,omitempty"`
			NamespaceSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
			PreloadPolicy *string `tfsdk:"preload_policy" json:"preloadPolicy,omitempty"`
		} `tfsdk:"required_hsts_policies" json:"requiredHSTSPolicies,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ConfigOpenshiftIoIngressV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_config_openshift_io_ingress_v1_manifest"
}

func (r *ConfigOpenshiftIoIngressV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Ingress holds cluster-wide information about ingress, including the default ingress domain used for routes. The canonical name is 'cluster'.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "Ingress holds cluster-wide information about ingress, including the default ingress domain used for routes. The canonical name is 'cluster'.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
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
				Description:         "spec holds user settable values for configuration",
				MarkdownDescription: "spec holds user settable values for configuration",
				Attributes: map[string]schema.Attribute{
					"apps_domain": schema.StringAttribute{
						Description:         "appsDomain is an optional domain to use instead of the one specified in the domain field when a Route is created without specifying an explicit host. If appsDomain is nonempty, this value is used to generate default host values for Route. Unlike domain, appsDomain may be modified after installation. This assumes a new ingresscontroller has been setup with a wildcard certificate.",
						MarkdownDescription: "appsDomain is an optional domain to use instead of the one specified in the domain field when a Route is created without specifying an explicit host. If appsDomain is nonempty, this value is used to generate default host values for Route. Unlike domain, appsDomain may be modified after installation. This assumes a new ingresscontroller has been setup with a wildcard certificate.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"component_routes": schema.ListNestedAttribute{
						Description:         "componentRoutes is an optional list of routes that are managed by OpenShift components that a cluster-admin is able to configure the hostname and serving certificate for. The namespace and name of each route in this list should match an existing entry in the status.componentRoutes list.  To determine the set of configurable Routes, look at namespace and name of entries in the .status.componentRoutes list, where participating operators write the status of configurable routes.",
						MarkdownDescription: "componentRoutes is an optional list of routes that are managed by OpenShift components that a cluster-admin is able to configure the hostname and serving certificate for. The namespace and name of each route in this list should match an existing entry in the status.componentRoutes list.  To determine the set of configurable Routes, look at namespace and name of entries in the .status.componentRoutes list, where participating operators write the status of configurable routes.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"hostname": schema.StringAttribute{
									Description:         "hostname is the hostname that should be used by the route.",
									MarkdownDescription: "hostname is the hostname that should be used by the route.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^([a-zA-Z0-9\p{S}\p{L}]((-?[a-zA-Z0-9\p{S}\p{L}]{0,62})?)|([a-zA-Z0-9\p{S}\p{L}](([a-zA-Z0-9-\p{S}\p{L}]{0,61}[a-zA-Z0-9\p{S}\p{L}])?)(\.)){1,}([a-zA-Z\p{L}]){2,63})$|^(([a-z0-9][-a-z0-9]{0,61}[a-z0-9]|[a-z0-9]{1,63})[\.]){0,}([a-z0-9][-a-z0-9]{0,61}[a-z0-9]|[a-z0-9]{1,63})$`), ""),
									},
								},

								"name": schema.StringAttribute{
									Description:         "name is the logical name of the route to customize.  The namespace and name of this componentRoute must match a corresponding entry in the list of status.componentRoutes if the route is to be customized.",
									MarkdownDescription: "name is the logical name of the route to customize.  The namespace and name of this componentRoute must match a corresponding entry in the list of status.componentRoutes if the route is to be customized.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(256),
									},
								},

								"namespace": schema.StringAttribute{
									Description:         "namespace is the namespace of the route to customize.  The namespace and name of this componentRoute must match a corresponding entry in the list of status.componentRoutes if the route is to be customized.",
									MarkdownDescription: "namespace is the namespace of the route to customize.  The namespace and name of this componentRoute must match a corresponding entry in the list of status.componentRoutes if the route is to be customized.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(63),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
									},
								},

								"serving_cert_key_pair_secret": schema.SingleNestedAttribute{
									Description:         "servingCertKeyPairSecret is a reference to a secret of type 'kubernetes.io/tls' in the openshift-config namespace. The serving cert/key pair must match and will be used by the operator to fulfill the intent of serving with this name. If the custom hostname uses the default routing suffix of the cluster, the Secret specification for a serving certificate will not be needed.",
									MarkdownDescription: "servingCertKeyPairSecret is a reference to a secret of type 'kubernetes.io/tls' in the openshift-config namespace. The serving cert/key pair must match and will be used by the operator to fulfill the intent of serving with this name. If the custom hostname uses the default routing suffix of the cluster, the Secret specification for a serving certificate will not be needed.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "name is the metadata.name of the referenced secret",
											MarkdownDescription: "name is the metadata.name of the referenced secret",
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

					"domain": schema.StringAttribute{
						Description:         "domain is used to generate a default host name for a route when the route's host name is empty. The generated host name will follow this pattern: '<route-name>.<route-namespace>.<domain>'.  It is also used as the default wildcard domain suffix for ingress. The default ingresscontroller domain will follow this pattern: '*.<domain>'.  Once set, changing domain is not currently supported.",
						MarkdownDescription: "domain is used to generate a default host name for a route when the route's host name is empty. The generated host name will follow this pattern: '<route-name>.<route-namespace>.<domain>'.  It is also used as the default wildcard domain suffix for ingress. The default ingresscontroller domain will follow this pattern: '*.<domain>'.  Once set, changing domain is not currently supported.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"load_balancer": schema.SingleNestedAttribute{
						Description:         "loadBalancer contains the load balancer details in general which are not only specific to the underlying infrastructure provider of the current cluster and are required for Ingress Controller to work on OpenShift.",
						MarkdownDescription: "loadBalancer contains the load balancer details in general which are not only specific to the underlying infrastructure provider of the current cluster and are required for Ingress Controller to work on OpenShift.",
						Attributes: map[string]schema.Attribute{
							"platform": schema.SingleNestedAttribute{
								Description:         "platform holds configuration specific to the underlying infrastructure provider for the ingress load balancers. When omitted, this means the user has no opinion and the platform is left to choose reasonable defaults. These defaults are subject to change over time.",
								MarkdownDescription: "platform holds configuration specific to the underlying infrastructure provider for the ingress load balancers. When omitted, this means the user has no opinion and the platform is left to choose reasonable defaults. These defaults are subject to change over time.",
								Attributes: map[string]schema.Attribute{
									"aws": schema.SingleNestedAttribute{
										Description:         "aws contains settings specific to the Amazon Web Services infrastructure provider.",
										MarkdownDescription: "aws contains settings specific to the Amazon Web Services infrastructure provider.",
										Attributes: map[string]schema.Attribute{
											"type": schema.StringAttribute{
												Description:         "type allows user to set a load balancer type. When this field is set the default ingresscontroller will get created using the specified LBType. If this field is not set then the default ingress controller of LBType Classic will be created. Valid values are:  * 'Classic': A Classic Load Balancer that makes routing decisions at either the transport layer (TCP/SSL) or the application layer (HTTP/HTTPS). See the following for additional details:  https://docs.aws.amazon.com/AmazonECS/latest/developerguide/load-balancer-types.html#clb  * 'NLB': A Network Load Balancer that makes routing decisions at the transport layer (TCP/SSL). See the following for additional details:  https://docs.aws.amazon.com/AmazonECS/latest/developerguide/load-balancer-types.html#nlb",
												MarkdownDescription: "type allows user to set a load balancer type. When this field is set the default ingresscontroller will get created using the specified LBType. If this field is not set then the default ingress controller of LBType Classic will be created. Valid values are:  * 'Classic': A Classic Load Balancer that makes routing decisions at either the transport layer (TCP/SSL) or the application layer (HTTP/HTTPS). See the following for additional details:  https://docs.aws.amazon.com/AmazonECS/latest/developerguide/load-balancer-types.html#clb  * 'NLB': A Network Load Balancer that makes routing decisions at the transport layer (TCP/SSL). See the following for additional details:  https://docs.aws.amazon.com/AmazonECS/latest/developerguide/load-balancer-types.html#nlb",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("NLB", "Classic"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": schema.StringAttribute{
										Description:         "type is the underlying infrastructure provider for the cluster. Allowed values are 'AWS', 'Azure', 'BareMetal', 'GCP', 'Libvirt', 'OpenStack', 'VSphere', 'oVirt', 'KubeVirt', 'EquinixMetal', 'PowerVS', 'AlibabaCloud', 'Nutanix' and 'None'. Individual components may not support all platforms, and must handle unrecognized platforms as None if they do not support that platform.",
										MarkdownDescription: "type is the underlying infrastructure provider for the cluster. Allowed values are 'AWS', 'Azure', 'BareMetal', 'GCP', 'Libvirt', 'OpenStack', 'VSphere', 'oVirt', 'KubeVirt', 'EquinixMetal', 'PowerVS', 'AlibabaCloud', 'Nutanix' and 'None'. Individual components may not support all platforms, and must handle unrecognized platforms as None if they do not support that platform.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("", "AWS", "Azure", "BareMetal", "GCP", "Libvirt", "OpenStack", "None", "VSphere", "oVirt", "IBMCloud", "KubeVirt", "EquinixMetal", "PowerVS", "AlibabaCloud", "Nutanix", "External"),
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

					"required_hsts_policies": schema.ListNestedAttribute{
						Description:         "requiredHSTSPolicies specifies HSTS policies that are required to be set on newly created  or updated routes matching the domainPattern/s and namespaceSelector/s that are specified in the policy. Each requiredHSTSPolicy must have at least a domainPattern and a maxAge to validate a route HSTS Policy route annotation, and affect route admission.  A candidate route is checked for HSTS Policies if it has the HSTS Policy route annotation: 'haproxy.router.openshift.io/hsts_header' E.g. haproxy.router.openshift.io/hsts_header: max-age=31536000;preload;includeSubDomains  - For each candidate route, if it matches a requiredHSTSPolicy domainPattern and optional namespaceSelector, then the maxAge, preloadPolicy, and includeSubdomainsPolicy must be valid to be admitted.  Otherwise, the route is rejected. - The first match, by domainPattern and optional namespaceSelector, in the ordering of the RequiredHSTSPolicies determines the route's admission status. - If the candidate route doesn't match any requiredHSTSPolicy domainPattern and optional namespaceSelector, then it may use any HSTS Policy annotation.  The HSTS policy configuration may be changed after routes have already been created. An update to a previously admitted route may then fail if the updated route does not conform to the updated HSTS policy configuration. However, changing the HSTS policy configuration will not cause a route that is already admitted to stop working.  Note that if there are no RequiredHSTSPolicies, any HSTS Policy annotation on the route is valid.",
						MarkdownDescription: "requiredHSTSPolicies specifies HSTS policies that are required to be set on newly created  or updated routes matching the domainPattern/s and namespaceSelector/s that are specified in the policy. Each requiredHSTSPolicy must have at least a domainPattern and a maxAge to validate a route HSTS Policy route annotation, and affect route admission.  A candidate route is checked for HSTS Policies if it has the HSTS Policy route annotation: 'haproxy.router.openshift.io/hsts_header' E.g. haproxy.router.openshift.io/hsts_header: max-age=31536000;preload;includeSubDomains  - For each candidate route, if it matches a requiredHSTSPolicy domainPattern and optional namespaceSelector, then the maxAge, preloadPolicy, and includeSubdomainsPolicy must be valid to be admitted.  Otherwise, the route is rejected. - The first match, by domainPattern and optional namespaceSelector, in the ordering of the RequiredHSTSPolicies determines the route's admission status. - If the candidate route doesn't match any requiredHSTSPolicy domainPattern and optional namespaceSelector, then it may use any HSTS Policy annotation.  The HSTS policy configuration may be changed after routes have already been created. An update to a previously admitted route may then fail if the updated route does not conform to the updated HSTS policy configuration. However, changing the HSTS policy configuration will not cause a route that is already admitted to stop working.  Note that if there are no RequiredHSTSPolicies, any HSTS Policy annotation on the route is valid.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"domain_patterns": schema.ListAttribute{
									Description:         "domainPatterns is a list of domains for which the desired HSTS annotations are required. If domainPatterns is specified and a route is created with a spec.host matching one of the domains, the route must specify the HSTS Policy components described in the matching RequiredHSTSPolicy.  The use of wildcards is allowed like this: *.foo.com matches everything under foo.com. foo.com only matches foo.com, so to cover foo.com and everything under it, you must specify *both*.",
									MarkdownDescription: "domainPatterns is a list of domains for which the desired HSTS annotations are required. If domainPatterns is specified and a route is created with a spec.host matching one of the domains, the route must specify the HSTS Policy components described in the matching RequiredHSTSPolicy.  The use of wildcards is allowed like this: *.foo.com matches everything under foo.com. foo.com only matches foo.com, so to cover foo.com and everything under it, you must specify *both*.",
									ElementType:         types.StringType,
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"include_sub_domains_policy": schema.StringAttribute{
									Description:         "includeSubDomainsPolicy means the HSTS Policy should apply to any subdomains of the host's domain name.  Thus, for the host bar.foo.com, if includeSubDomainsPolicy was set to RequireIncludeSubDomains: - the host app.bar.foo.com would inherit the HSTS Policy of bar.foo.com - the host bar.foo.com would inherit the HSTS Policy of bar.foo.com - the host foo.com would NOT inherit the HSTS Policy of bar.foo.com - the host def.foo.com would NOT inherit the HSTS Policy of bar.foo.com",
									MarkdownDescription: "includeSubDomainsPolicy means the HSTS Policy should apply to any subdomains of the host's domain name.  Thus, for the host bar.foo.com, if includeSubDomainsPolicy was set to RequireIncludeSubDomains: - the host app.bar.foo.com would inherit the HSTS Policy of bar.foo.com - the host bar.foo.com would inherit the HSTS Policy of bar.foo.com - the host foo.com would NOT inherit the HSTS Policy of bar.foo.com - the host def.foo.com would NOT inherit the HSTS Policy of bar.foo.com",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("RequireIncludeSubDomains", "RequireNoIncludeSubDomains", "NoOpinion"),
									},
								},

								"max_age": schema.SingleNestedAttribute{
									Description:         "maxAge is the delta time range in seconds during which hosts are regarded as HSTS hosts. If set to 0, it negates the effect, and hosts are removed as HSTS hosts. If set to 0 and includeSubdomains is specified, all subdomains of the host are also removed as HSTS hosts. maxAge is a time-to-live value, and if this policy is not refreshed on a client, the HSTS policy will eventually expire on that client.",
									MarkdownDescription: "maxAge is the delta time range in seconds during which hosts are regarded as HSTS hosts. If set to 0, it negates the effect, and hosts are removed as HSTS hosts. If set to 0 and includeSubdomains is specified, all subdomains of the host are also removed as HSTS hosts. maxAge is a time-to-live value, and if this policy is not refreshed on a client, the HSTS policy will eventually expire on that client.",
									Attributes: map[string]schema.Attribute{
										"largest_max_age": schema.Int64Attribute{
											Description:         "The largest allowed value (in seconds) of the RequiredHSTSPolicy max-age This value can be left unspecified, in which case no upper limit is enforced.",
											MarkdownDescription: "The largest allowed value (in seconds) of the RequiredHSTSPolicy max-age This value can be left unspecified, in which case no upper limit is enforced.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
												int64validator.AtMost(2.147483647e+09),
											},
										},

										"smallest_max_age": schema.Int64Attribute{
											Description:         "The smallest allowed value (in seconds) of the RequiredHSTSPolicy max-age Setting max-age=0 allows the deletion of an existing HSTS header from a host.  This is a necessary tool for administrators to quickly correct mistakes. This value can be left unspecified, in which case no lower limit is enforced.",
											MarkdownDescription: "The smallest allowed value (in seconds) of the RequiredHSTSPolicy max-age Setting max-age=0 allows the deletion of an existing HSTS header from a host.  This is a necessary tool for administrators to quickly correct mistakes. This value can be left unspecified, in which case no lower limit is enforced.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
												int64validator.AtMost(2.147483647e+09),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"namespace_selector": schema.SingleNestedAttribute{
									Description:         "namespaceSelector specifies a label selector such that the policy applies only to those routes that are in namespaces with labels that match the selector, and are in one of the DomainPatterns. Defaults to the empty LabelSelector, which matches everything.",
									MarkdownDescription: "namespaceSelector specifies a label selector such that the policy applies only to those routes that are in namespaces with labels that match the selector, and are in one of the DomainPatterns. Defaults to the empty LabelSelector, which matches everything.",
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
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"values": schema.ListAttribute{
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
											Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

								"preload_policy": schema.StringAttribute{
									Description:         "preloadPolicy directs the client to include hosts in its host preload list so that it never needs to do an initial load to get the HSTS header (note that this is not defined in RFC 6797 and is therefore client implementation-dependent).",
									MarkdownDescription: "preloadPolicy directs the client to include hosts in its host preload list so that it never needs to do an initial load to get the HSTS header (note that this is not defined in RFC 6797 and is therefore client implementation-dependent).",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("RequirePreload", "RequireNoPreload", "NoOpinion"),
									},
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
		},
	}
}

func (r *ConfigOpenshiftIoIngressV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_config_openshift_io_ingress_v1_manifest")

	var model ConfigOpenshiftIoIngressV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("config.openshift.io/v1")
	model.Kind = pointer.String("Ingress")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
