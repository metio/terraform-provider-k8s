/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kibana_k8s_elastic_co_v1beta1

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
	_ datasource.DataSource = &KibanaK8SElasticCoKibanaV1Beta1Manifest{}
)

func NewKibanaK8SElasticCoKibanaV1Beta1Manifest() datasource.DataSource {
	return &KibanaK8SElasticCoKibanaV1Beta1Manifest{}
}

type KibanaK8SElasticCoKibanaV1Beta1Manifest struct{}

type KibanaK8SElasticCoKibanaV1Beta1ManifestData struct {
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
		Config           *map[string]string `tfsdk:"config" json:"config,omitempty"`
		Count            *int64             `tfsdk:"count" json:"count,omitempty"`
		ElasticsearchRef *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"elasticsearch_ref" json:"elasticsearchRef,omitempty"`
		Http *struct {
			Service *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Finalizers  *[]string          `tfsdk:"finalizers" json:"finalizers,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name        *string            `tfsdk:"name" json:"name,omitempty"`
					Namespace   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec *struct {
					AllocateLoadBalancerNodePorts *bool     `tfsdk:"allocate_load_balancer_node_ports" json:"allocateLoadBalancerNodePorts,omitempty"`
					ClusterIP                     *string   `tfsdk:"cluster_ip" json:"clusterIP,omitempty"`
					ClusterIPs                    *[]string `tfsdk:"cluster_i_ps" json:"clusterIPs,omitempty"`
					ExternalIPs                   *[]string `tfsdk:"external_i_ps" json:"externalIPs,omitempty"`
					ExternalName                  *string   `tfsdk:"external_name" json:"externalName,omitempty"`
					ExternalTrafficPolicy         *string   `tfsdk:"external_traffic_policy" json:"externalTrafficPolicy,omitempty"`
					HealthCheckNodePort           *int64    `tfsdk:"health_check_node_port" json:"healthCheckNodePort,omitempty"`
					InternalTrafficPolicy         *string   `tfsdk:"internal_traffic_policy" json:"internalTrafficPolicy,omitempty"`
					IpFamilies                    *[]string `tfsdk:"ip_families" json:"ipFamilies,omitempty"`
					IpFamilyPolicy                *string   `tfsdk:"ip_family_policy" json:"ipFamilyPolicy,omitempty"`
					LoadBalancerClass             *string   `tfsdk:"load_balancer_class" json:"loadBalancerClass,omitempty"`
					LoadBalancerIP                *string   `tfsdk:"load_balancer_ip" json:"loadBalancerIP,omitempty"`
					LoadBalancerSourceRanges      *[]string `tfsdk:"load_balancer_source_ranges" json:"loadBalancerSourceRanges,omitempty"`
					Ports                         *[]struct {
						AppProtocol *string `tfsdk:"app_protocol" json:"appProtocol,omitempty"`
						Name        *string `tfsdk:"name" json:"name,omitempty"`
						NodePort    *int64  `tfsdk:"node_port" json:"nodePort,omitempty"`
						Port        *int64  `tfsdk:"port" json:"port,omitempty"`
						Protocol    *string `tfsdk:"protocol" json:"protocol,omitempty"`
						TargetPort  *string `tfsdk:"target_port" json:"targetPort,omitempty"`
					} `tfsdk:"ports" json:"ports,omitempty"`
					PublishNotReadyAddresses *bool              `tfsdk:"publish_not_ready_addresses" json:"publishNotReadyAddresses,omitempty"`
					Selector                 *map[string]string `tfsdk:"selector" json:"selector,omitempty"`
					SessionAffinity          *string            `tfsdk:"session_affinity" json:"sessionAffinity,omitempty"`
					SessionAffinityConfig    *struct {
						ClientIP *struct {
							TimeoutSeconds *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
						} `tfsdk:"client_ip" json:"clientIP,omitempty"`
					} `tfsdk:"session_affinity_config" json:"sessionAffinityConfig,omitempty"`
					TrafficDistribution *string `tfsdk:"traffic_distribution" json:"trafficDistribution,omitempty"`
					Type                *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"service" json:"service,omitempty"`
			Tls *struct {
				Certificate *struct {
					SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
				} `tfsdk:"certificate" json:"certificate,omitempty"`
				SelfSignedCertificate *struct {
					Disabled        *bool `tfsdk:"disabled" json:"disabled,omitempty"`
					SubjectAltNames *[]struct {
						Dns *string `tfsdk:"dns" json:"dns,omitempty"`
						Ip  *string `tfsdk:"ip" json:"ip,omitempty"`
					} `tfsdk:"subject_alt_names" json:"subjectAltNames,omitempty"`
				} `tfsdk:"self_signed_certificate" json:"selfSignedCertificate,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"http" json:"http,omitempty"`
		Image          *string            `tfsdk:"image" json:"image,omitempty"`
		PodTemplate    *map[string]string `tfsdk:"pod_template" json:"podTemplate,omitempty"`
		SecureSettings *[]struct {
			Entries *[]struct {
				Key  *string `tfsdk:"key" json:"key,omitempty"`
				Path *string `tfsdk:"path" json:"path,omitempty"`
			} `tfsdk:"entries" json:"entries,omitempty"`
			SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
		} `tfsdk:"secure_settings" json:"secureSettings,omitempty"`
		Version *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KibanaK8SElasticCoKibanaV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kibana_k8s_elastic_co_kibana_v1beta1_manifest"
}

func (r *KibanaK8SElasticCoKibanaV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Kibana represents a Kibana resource in a Kubernetes cluster.",
		MarkdownDescription: "Kibana represents a Kibana resource in a Kubernetes cluster.",
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
				Description:         "KibanaSpec holds the specification of a Kibana instance.",
				MarkdownDescription: "KibanaSpec holds the specification of a Kibana instance.",
				Attributes: map[string]schema.Attribute{
					"config": schema.MapAttribute{
						Description:         "Config holds the Kibana configuration. See: https://www.elastic.co/guide/en/kibana/current/settings.html",
						MarkdownDescription: "Config holds the Kibana configuration. See: https://www.elastic.co/guide/en/kibana/current/settings.html",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"count": schema.Int64Attribute{
						Description:         "Count of Kibana instances to deploy.",
						MarkdownDescription: "Count of Kibana instances to deploy.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"elasticsearch_ref": schema.SingleNestedAttribute{
						Description:         "ElasticsearchRef is a reference to an Elasticsearch cluster running in the same Kubernetes cluster.",
						MarkdownDescription: "ElasticsearchRef is a reference to an Elasticsearch cluster running in the same Kubernetes cluster.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the Kubernetes object.",
								MarkdownDescription: "Name of the Kubernetes object.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the Kubernetes object. If empty, defaults to the current namespace.",
								MarkdownDescription: "Namespace of the Kubernetes object. If empty, defaults to the current namespace.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"http": schema.SingleNestedAttribute{
						Description:         "HTTP holds the HTTP layer configuration for Kibana.",
						MarkdownDescription: "HTTP holds the HTTP layer configuration for Kibana.",
						Attributes: map[string]schema.Attribute{
							"service": schema.SingleNestedAttribute{
								Description:         "Service defines the template for the associated Kubernetes Service object.",
								MarkdownDescription: "Service defines the template for the associated Kubernetes Service object.",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "ObjectMeta is the metadata of the service. The name and namespace provided here are managed by ECK and will be ignored.",
										MarkdownDescription: "ObjectMeta is the metadata of the service. The name and namespace provided here are managed by ECK and will be ignored.",
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

									"spec": schema.SingleNestedAttribute{
										Description:         "Spec is the specification of the service.",
										MarkdownDescription: "Spec is the specification of the service.",
										Attributes: map[string]schema.Attribute{
											"allocate_load_balancer_node_ports": schema.BoolAttribute{
												Description:         "allocateLoadBalancerNodePorts defines if NodePorts will be automatically allocated for services with type LoadBalancer. Default is 'true'. It may be set to 'false' if the cluster load-balancer does not rely on NodePorts. If the caller requests specific NodePorts (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type LoadBalancer and will be cleared if the type is changed to any other type.",
												MarkdownDescription: "allocateLoadBalancerNodePorts defines if NodePorts will be automatically allocated for services with type LoadBalancer. Default is 'true'. It may be set to 'false' if the cluster load-balancer does not rely on NodePorts. If the caller requests specific NodePorts (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type LoadBalancer and will be cleared if the type is changed to any other type.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cluster_ip": schema.StringAttribute{
												Description:         "clusterIP is the IP address of the service and is usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be blank) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above). Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												MarkdownDescription: "clusterIP is the IP address of the service and is usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be blank) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above). Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cluster_i_ps": schema.ListAttribute{
												Description:         "ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be empty) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above). Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName. If this field is not specified, it will be initialized from the clusterIP field. If this field is specified, clients must ensure that clusterIPs[0] and clusterIP have the same value. This field may hold a maximum of two entries (dual-stack IPs, in either order). These IPs must correspond to the values of the ipFamilies field. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												MarkdownDescription: "ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be empty) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above). Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName. If this field is not specified, it will be initialized from the clusterIP field. If this field is specified, clients must ensure that clusterIPs[0] and clusterIP have the same value. This field may hold a maximum of two entries (dual-stack IPs, in either order). These IPs must correspond to the values of the ipFamilies field. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"external_i_ps": schema.ListAttribute{
												Description:         "externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service. These IPs are not managed by Kubernetes. The user is responsible for ensuring that traffic arrives at a node with this IP. A common example is external load-balancers that are not part of the Kubernetes system.",
												MarkdownDescription: "externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service. These IPs are not managed by Kubernetes. The user is responsible for ensuring that traffic arrives at a node with this IP. A common example is external load-balancers that are not part of the Kubernetes system.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"external_name": schema.StringAttribute{
												Description:         "externalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record). No proxying will be involved. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires 'type' to be 'ExternalName'.",
												MarkdownDescription: "externalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record). No proxying will be involved. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires 'type' to be 'ExternalName'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"external_traffic_policy": schema.StringAttribute{
												Description:         "externalTrafficPolicy describes how nodes distribute service traffic they receive on one of the Service's 'externally-facing' addresses (NodePorts, ExternalIPs, and LoadBalancer IPs). If set to 'Local', the proxy will configure the service in a way that assumes that external load balancers will take care of balancing the service traffic between nodes, and so each node will deliver traffic only to the node-local endpoints of the service, without masquerading the client source IP. (Traffic mistakenly sent to a node with no endpoints will be dropped.) The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features). Note that traffic sent to an External IP or LoadBalancer IP from within the cluster will always get 'Cluster' semantics, but clients sending to a NodePort from within the cluster may need to take traffic policy into account when picking a node.",
												MarkdownDescription: "externalTrafficPolicy describes how nodes distribute service traffic they receive on one of the Service's 'externally-facing' addresses (NodePorts, ExternalIPs, and LoadBalancer IPs). If set to 'Local', the proxy will configure the service in a way that assumes that external load balancers will take care of balancing the service traffic between nodes, and so each node will deliver traffic only to the node-local endpoints of the service, without masquerading the client source IP. (Traffic mistakenly sent to a node with no endpoints will be dropped.) The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features). Note that traffic sent to an External IP or LoadBalancer IP from within the cluster will always get 'Cluster' semantics, but clients sending to a NodePort from within the cluster may need to take traffic policy into account when picking a node.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"health_check_node_port": schema.Int64Attribute{
												Description:         "healthCheckNodePort specifies the healthcheck nodePort for the service. This only applies when type is set to LoadBalancer and externalTrafficPolicy is set to Local. If a value is specified, is in-range, and is not in use, it will be used. If not specified, a value will be automatically allocated. External systems (e.g. load-balancers) can use this port to determine if a given node holds endpoints for this service or not. If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type). This field cannot be updated once set.",
												MarkdownDescription: "healthCheckNodePort specifies the healthcheck nodePort for the service. This only applies when type is set to LoadBalancer and externalTrafficPolicy is set to Local. If a value is specified, is in-range, and is not in use, it will be used. If not specified, a value will be automatically allocated. External systems (e.g. load-balancers) can use this port to determine if a given node holds endpoints for this service or not. If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type). This field cannot be updated once set.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"internal_traffic_policy": schema.StringAttribute{
												Description:         "InternalTrafficPolicy describes how nodes distribute service traffic they receive on the ClusterIP. If set to 'Local', the proxy will assume that pods only want to talk to endpoints of the service on the same node as the pod, dropping the traffic if there are no local endpoints. The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features).",
												MarkdownDescription: "InternalTrafficPolicy describes how nodes distribute service traffic they receive on the ClusterIP. If set to 'Local', the proxy will assume that pods only want to talk to endpoints of the service on the same node as the pod, dropping the traffic if there are no local endpoints. The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ip_families": schema.ListAttribute{
												Description:         "IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the Service. Valid values are 'IPv4' and 'IPv6'. This field only applies to Services of types ClusterIP, NodePort, and LoadBalancer, and does apply to 'headless' services. This field will be wiped when updating a Service to type ExternalName. This field may hold a maximum of two entries (dual-stack families, in either order). These families must correspond to the values of the clusterIPs field, if specified. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field.",
												MarkdownDescription: "IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the Service. Valid values are 'IPv4' and 'IPv6'. This field only applies to Services of types ClusterIP, NodePort, and LoadBalancer, and does apply to 'headless' services. This field will be wiped when updating a Service to type ExternalName. This field may hold a maximum of two entries (dual-stack families, in either order). These families must correspond to the values of the clusterIPs field, if specified. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ip_family_policy": schema.StringAttribute{
												Description:         "IPFamilyPolicy represents the dual-stack-ness requested or required by this Service. If there is no value provided, then this field will be set to SingleStack. Services can be 'SingleStack' (a single IP family), 'PreferDualStack' (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or 'RequireDualStack' (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field. This field will be wiped when updating a service to type ExternalName.",
												MarkdownDescription: "IPFamilyPolicy represents the dual-stack-ness requested or required by this Service. If there is no value provided, then this field will be set to SingleStack. Services can be 'SingleStack' (a single IP family), 'PreferDualStack' (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or 'RequireDualStack' (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field. This field will be wiped when updating a service to type ExternalName.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"load_balancer_class": schema.StringAttribute{
												Description:         "loadBalancerClass is the class of the load balancer implementation this Service belongs to. If specified, the value of this field must be a label-style identifier, with an optional prefix, e.g. 'internal-vip' or 'example.com/internal-vip'. Unprefixed names are reserved for end-users. This field can only be set when the Service type is 'LoadBalancer'. If not set, the default load balancer implementation is used, today this is typically done through the cloud provider integration, but should apply for any default implementation. If set, it is assumed that a load balancer implementation is watching for Services with a matching class. Any default load balancer implementation (e.g. cloud providers) should ignore Services that set this field. This field can only be set when creating or updating a Service to type 'LoadBalancer'. Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type.",
												MarkdownDescription: "loadBalancerClass is the class of the load balancer implementation this Service belongs to. If specified, the value of this field must be a label-style identifier, with an optional prefix, e.g. 'internal-vip' or 'example.com/internal-vip'. Unprefixed names are reserved for end-users. This field can only be set when the Service type is 'LoadBalancer'. If not set, the default load balancer implementation is used, today this is typically done through the cloud provider integration, but should apply for any default implementation. If set, it is assumed that a load balancer implementation is watching for Services with a matching class. Any default load balancer implementation (e.g. cloud providers) should ignore Services that set this field. This field can only be set when creating or updating a Service to type 'LoadBalancer'. Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"load_balancer_ip": schema.StringAttribute{
												Description:         "Only applies to Service Type: LoadBalancer. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature. Deprecated: This field was under-specified and its meaning varies across implementations. Using it is non-portable and it may not support dual-stack. Users are encouraged to use implementation-specific annotations when available.",
												MarkdownDescription: "Only applies to Service Type: LoadBalancer. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature. Deprecated: This field was under-specified and its meaning varies across implementations. Using it is non-portable and it may not support dual-stack. Users are encouraged to use implementation-specific annotations when available.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"load_balancer_source_ranges": schema.ListAttribute{
												Description:         "If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature.' More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/",
												MarkdownDescription: "If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature.' More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ports": schema.ListNestedAttribute{
												Description:         "The list of ports that are exposed by this service. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												MarkdownDescription: "The list of ports that are exposed by this service. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"app_protocol": schema.StringAttribute{
															Description:         "The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either: * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). * Kubernetes-defined prefixed names: * 'kubernetes.io/h2c' - HTTP/2 prior knowledge over cleartext as described in https://www.rfc-editor.org/rfc/rfc9113.html#name-starting-http-2-with-prior- * 'kubernetes.io/ws' - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455 * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455 * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.",
															MarkdownDescription: "The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either: * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). * Kubernetes-defined prefixed names: * 'kubernetes.io/h2c' - HTTP/2 prior knowledge over cleartext as described in https://www.rfc-editor.org/rfc/rfc9113.html#name-starting-http-2-with-prior- * 'kubernetes.io/ws' - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455 * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455 * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the 'name' field in the EndpointPort. Optional if only one ServicePort is defined on this service.",
															MarkdownDescription: "The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the 'name' field in the EndpointPort. Optional if only one ServicePort is defined on this service.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"node_port": schema.Int64Attribute{
															Description:         "The port on each node on which this service is exposed when type is NodePort or LoadBalancer. Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail. If not specified, a port will be allocated if this Service requires one. If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type from NodePort to ClusterIP). More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport",
															MarkdownDescription: "The port on each node on which this service is exposed when type is NodePort or LoadBalancer. Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail. If not specified, a port will be allocated if this Service requires one. If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type from NodePort to ClusterIP). More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.Int64Attribute{
															Description:         "The port that will be exposed by this service.",
															MarkdownDescription: "The port that will be exposed by this service.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"protocol": schema.StringAttribute{
															Description:         "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'. Default is TCP.",
															MarkdownDescription: "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'. Default is TCP.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"target_port": schema.StringAttribute{
															Description:         "Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If this is a string, it will be looked up as a named port in the target Pod's container ports. If this is not specified, the value of the 'port' field is used (an identity map). This field is ignored for services with clusterIP=None, and should be omitted or set equal to the 'port' field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service",
															MarkdownDescription: "Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If this is a string, it will be looked up as a named port in the target Pod's container ports. If this is not specified, the value of the 'port' field is used (an identity map). This field is ignored for services with clusterIP=None, and should be omitted or set equal to the 'port' field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service",
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

											"publish_not_ready_addresses": schema.BoolAttribute{
												Description:         "publishNotReadyAddresses indicates that any agent which deals with endpoints for this Service should disregard any indications of ready/not-ready. The primary use case for setting this field is for a StatefulSet's Headless Service to propagate SRV DNS records for its Pods for the purpose of peer discovery. The Kubernetes controllers that generate Endpoints and EndpointSlice resources for Services interpret this to mean that all endpoints are considered 'ready' even if the Pods themselves are not. Agents which consume only Kubernetes generated endpoints through the Endpoints or EndpointSlice resources can safely assume this behavior.",
												MarkdownDescription: "publishNotReadyAddresses indicates that any agent which deals with endpoints for this Service should disregard any indications of ready/not-ready. The primary use case for setting this field is for a StatefulSet's Headless Service to propagate SRV DNS records for its Pods for the purpose of peer discovery. The Kubernetes controllers that generate Endpoints and EndpointSlice resources for Services interpret this to mean that all endpoints are considered 'ready' even if the Pods themselves are not. Agents which consume only Kubernetes generated endpoints through the Endpoints or EndpointSlice resources can safely assume this behavior.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"selector": schema.MapAttribute{
												Description:         "Route service traffic to pods with label keys and values matching this selector. If empty or not present, the service is assumed to have an external process managing its endpoints, which Kubernetes will not modify. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/",
												MarkdownDescription: "Route service traffic to pods with label keys and values matching this selector. If empty or not present, the service is assumed to have an external process managing its endpoints, which Kubernetes will not modify. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"session_affinity": schema.StringAttribute{
												Description:         "Supports 'ClientIP' and 'None'. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												MarkdownDescription: "Supports 'ClientIP' and 'None'. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"session_affinity_config": schema.SingleNestedAttribute{
												Description:         "sessionAffinityConfig contains the configurations of session affinity.",
												MarkdownDescription: "sessionAffinityConfig contains the configurations of session affinity.",
												Attributes: map[string]schema.Attribute{
													"client_ip": schema.SingleNestedAttribute{
														Description:         "clientIP contains the configurations of Client IP based session affinity.",
														MarkdownDescription: "clientIP contains the configurations of Client IP based session affinity.",
														Attributes: map[string]schema.Attribute{
															"timeout_seconds": schema.Int64Attribute{
																Description:         "timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'. Default value is 10800(for 3 hours).",
																MarkdownDescription: "timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'. Default value is 10800(for 3 hours).",
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

											"traffic_distribution": schema.StringAttribute{
												Description:         "TrafficDistribution offers a way to express preferences for how traffic is distributed to Service endpoints. Implementations can use this field as a hint, but are not required to guarantee strict adherence. If the field is not set, the implementation will apply its default routing strategy. If set to 'PreferClose', implementations should prioritize endpoints that are in the same zone.",
												MarkdownDescription: "TrafficDistribution offers a way to express preferences for how traffic is distributed to Service endpoints. Implementations can use this field as a hint, but are not required to guarantee strict adherence. If the field is not set, the implementation will apply its default routing strategy. If set to 'PreferClose', implementations should prioritize endpoints that are in the same zone.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is 'None', no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. 'ExternalName' aliases this service to the specified externalName. Several other fields do not apply to ExternalName services. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types",
												MarkdownDescription: "type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is 'None', no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. 'ExternalName' aliases this service to the specified externalName. Several other fields do not apply to ExternalName services. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types",
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

							"tls": schema.SingleNestedAttribute{
								Description:         "TLS defines options for configuring TLS for HTTP.",
								MarkdownDescription: "TLS defines options for configuring TLS for HTTP.",
								Attributes: map[string]schema.Attribute{
									"certificate": schema.SingleNestedAttribute{
										Description:         "Certificate is a reference to a Kubernetes secret that contains the certificate and private key for enabling TLS. The referenced secret should contain the following: - 'ca.crt': The certificate authority (optional). - 'tls.crt': The certificate (or a chain). - 'tls.key': The private key to the first certificate in the certificate chain.",
										MarkdownDescription: "Certificate is a reference to a Kubernetes secret that contains the certificate and private key for enabling TLS. The referenced secret should contain the following: - 'ca.crt': The certificate authority (optional). - 'tls.crt': The certificate (or a chain). - 'tls.key': The private key to the first certificate in the certificate chain.",
										Attributes: map[string]schema.Attribute{
											"secret_name": schema.StringAttribute{
												Description:         "SecretName is the name of the secret.",
												MarkdownDescription: "SecretName is the name of the secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"self_signed_certificate": schema.SingleNestedAttribute{
										Description:         "SelfSignedCertificate allows configuring the self-signed certificate generated by the operator.",
										MarkdownDescription: "SelfSignedCertificate allows configuring the self-signed certificate generated by the operator.",
										Attributes: map[string]schema.Attribute{
											"disabled": schema.BoolAttribute{
												Description:         "Disabled indicates that the provisioning of the self-signed certifcate should be disabled.",
												MarkdownDescription: "Disabled indicates that the provisioning of the self-signed certifcate should be disabled.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"subject_alt_names": schema.ListNestedAttribute{
												Description:         "SubjectAlternativeNames is a list of SANs to include in the generated HTTP TLS certificate.",
												MarkdownDescription: "SubjectAlternativeNames is a list of SANs to include in the generated HTTP TLS certificate.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"dns": schema.StringAttribute{
															Description:         "DNS is the DNS name of the subject.",
															MarkdownDescription: "DNS is the DNS name of the subject.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ip": schema.StringAttribute{
															Description:         "IP is the IP address of the subject.",
															MarkdownDescription: "IP is the IP address of the subject.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"image": schema.StringAttribute{
						Description:         "Image is the Kibana Docker image to deploy.",
						MarkdownDescription: "Image is the Kibana Docker image to deploy.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pod_template": schema.MapAttribute{
						Description:         "PodTemplate provides customisation options (labels, annotations, affinity rules, resource requests, and so on) for the Kibana pods",
						MarkdownDescription: "PodTemplate provides customisation options (labels, annotations, affinity rules, resource requests, and so on) for the Kibana pods",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"secure_settings": schema.ListNestedAttribute{
						Description:         "SecureSettings is a list of references to Kubernetes secrets containing sensitive configuration options for Kibana.",
						MarkdownDescription: "SecureSettings is a list of references to Kubernetes secrets containing sensitive configuration options for Kibana.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"entries": schema.ListNestedAttribute{
									Description:         "Entries define how to project each key-value pair in the secret to filesystem paths. If not defined, all keys will be projected to similarly named paths in the filesystem. If defined, only the specified keys will be projected to the corresponding paths.",
									MarkdownDescription: "Entries define how to project each key-value pair in the secret to filesystem paths. If not defined, all keys will be projected to similarly named paths in the filesystem. If defined, only the specified keys will be projected to the corresponding paths.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "Key is the key contained in the secret.",
												MarkdownDescription: "Key is the key contained in the secret.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"path": schema.StringAttribute{
												Description:         "Path is the relative file path to map the key to. Path must not be an absolute file path and must not contain any '..' components.",
												MarkdownDescription: "Path is the relative file path to map the key to. Path must not be an absolute file path and must not contain any '..' components.",
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

								"secret_name": schema.StringAttribute{
									Description:         "SecretName is the name of the secret.",
									MarkdownDescription: "SecretName is the name of the secret.",
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

					"version": schema.StringAttribute{
						Description:         "Version of Kibana.",
						MarkdownDescription: "Version of Kibana.",
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
	}
}

func (r *KibanaK8SElasticCoKibanaV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kibana_k8s_elastic_co_kibana_v1beta1_manifest")

	var model KibanaK8SElasticCoKibanaV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kibana.k8s.elastic.co/v1beta1")
	model.Kind = pointer.String("Kibana")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
