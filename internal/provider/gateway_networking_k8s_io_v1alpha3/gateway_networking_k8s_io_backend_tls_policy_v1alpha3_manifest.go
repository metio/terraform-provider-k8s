/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package gateway_networking_k8s_io_v1alpha3

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &GatewayNetworkingK8SIoBackendTlspolicyV1Alpha3Manifest{}
)

func NewGatewayNetworkingK8SIoBackendTlspolicyV1Alpha3Manifest() datasource.DataSource {
	return &GatewayNetworkingK8SIoBackendTlspolicyV1Alpha3Manifest{}
}

type GatewayNetworkingK8SIoBackendTlspolicyV1Alpha3Manifest struct{}

type GatewayNetworkingK8SIoBackendTlspolicyV1Alpha3ManifestData struct {
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
		Options    *map[string]string `tfsdk:"options" json:"options,omitempty"`
		TargetRefs *[]struct {
			Group       *string `tfsdk:"group" json:"group,omitempty"`
			Kind        *string `tfsdk:"kind" json:"kind,omitempty"`
			Name        *string `tfsdk:"name" json:"name,omitempty"`
			SectionName *string `tfsdk:"section_name" json:"sectionName,omitempty"`
		} `tfsdk:"target_refs" json:"targetRefs,omitempty"`
		Validation *struct {
			CaCertificateRefs *[]struct {
				Group *string `tfsdk:"group" json:"group,omitempty"`
				Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"ca_certificate_refs" json:"caCertificateRefs,omitempty"`
			Hostname        *string `tfsdk:"hostname" json:"hostname,omitempty"`
			SubjectAltNames *[]struct {
				Hostname *string `tfsdk:"hostname" json:"hostname,omitempty"`
				Type     *string `tfsdk:"type" json:"type,omitempty"`
				Uri      *string `tfsdk:"uri" json:"uri,omitempty"`
			} `tfsdk:"subject_alt_names" json:"subjectAltNames,omitempty"`
			WellKnownCACertificates *string `tfsdk:"well_known_ca_certificates" json:"wellKnownCACertificates,omitempty"`
		} `tfsdk:"validation" json:"validation,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GatewayNetworkingK8SIoBackendTlspolicyV1Alpha3Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_gateway_networking_k8s_io_backend_tls_policy_v1alpha3_manifest"
}

func (r *GatewayNetworkingK8SIoBackendTlspolicyV1Alpha3Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "BackendTLSPolicy provides a way to configure how a Gateway connects to a Backend via TLS.",
		MarkdownDescription: "BackendTLSPolicy provides a way to configure how a Gateway connects to a Backend via TLS.",
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
				Description:         "Spec defines the desired state of BackendTLSPolicy.",
				MarkdownDescription: "Spec defines the desired state of BackendTLSPolicy.",
				Attributes: map[string]schema.Attribute{
					"options": schema.MapAttribute{
						Description:         "Options are a list of key/value pairs to enable extended TLS configuration for each implementation. For example, configuring the minimum TLS version or supported cipher suites. A set of common keys MAY be defined by the API in the future. To avoid any ambiguity, implementation-specific definitions MUST use domain-prefixed names, such as 'example.com/my-custom-option'. Un-prefixed names are reserved for key names defined by Gateway API. Support: Implementation-specific",
						MarkdownDescription: "Options are a list of key/value pairs to enable extended TLS configuration for each implementation. For example, configuring the minimum TLS version or supported cipher suites. A set of common keys MAY be defined by the API in the future. To avoid any ambiguity, implementation-specific definitions MUST use domain-prefixed names, such as 'example.com/my-custom-option'. Un-prefixed names are reserved for key names defined by Gateway API. Support: Implementation-specific",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"target_refs": schema.ListNestedAttribute{
						Description:         "TargetRefs identifies an API object to apply the policy to. Only Services have Extended support. Implementations MAY support additional objects, with Implementation Specific support. Note that this config applies to the entire referenced resource by default, but this default may change in the future to provide a more granular application of the policy. Support: Extended for Kubernetes Service Support: Implementation-specific for any other resource",
						MarkdownDescription: "TargetRefs identifies an API object to apply the policy to. Only Services have Extended support. Implementations MAY support additional objects, with Implementation Specific support. Note that this config applies to the entire referenced resource by default, but this default may change in the future to provide a more granular application of the policy. Support: Extended for Kubernetes Service Support: Implementation-specific for any other resource",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group": schema.StringAttribute{
									Description:         "Group is the group of the target resource.",
									MarkdownDescription: "Group is the group of the target resource.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(253),
										stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
									},
								},

								"kind": schema.StringAttribute{
									Description:         "Kind is kind of the target resource.",
									MarkdownDescription: "Kind is kind of the target resource.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(63),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
									},
								},

								"name": schema.StringAttribute{
									Description:         "Name is the name of the target resource.",
									MarkdownDescription: "Name is the name of the target resource.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(253),
									},
								},

								"section_name": schema.StringAttribute{
									Description:         "SectionName is the name of a section within the target resource. When unspecified, this targetRef targets the entire resource. In the following resources, SectionName is interpreted as the following: * Gateway: Listener name * HTTPRoute: HTTPRouteRule name * Service: Port name If a SectionName is specified, but does not exist on the targeted object, the Policy must fail to attach, and the policy implementation should record a 'ResolvedRefs' or similar Condition in the Policy's status.",
									MarkdownDescription: "SectionName is the name of a section within the target resource. When unspecified, this targetRef targets the entire resource. In the following resources, SectionName is interpreted as the following: * Gateway: Listener name * HTTPRoute: HTTPRouteRule name * Service: Port name If a SectionName is specified, but does not exist on the targeted object, the Policy must fail to attach, and the policy implementation should record a 'ResolvedRefs' or similar Condition in the Policy's status.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(253),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
									},
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"validation": schema.SingleNestedAttribute{
						Description:         "Validation contains backend TLS validation configuration.",
						MarkdownDescription: "Validation contains backend TLS validation configuration.",
						Attributes: map[string]schema.Attribute{
							"ca_certificate_refs": schema.ListNestedAttribute{
								Description:         "CACertificateRefs contains one or more references to Kubernetes objects that contain a PEM-encoded TLS CA certificate bundle, which is used to validate a TLS handshake between the Gateway and backend Pod. If CACertificateRefs is empty or unspecified, then WellKnownCACertificates must be specified. Only one of CACertificateRefs or WellKnownCACertificates may be specified, not both. If CACertifcateRefs is empty or unspecified, the configuration for WellKnownCACertificates MUST be honored instead if supported by the implementation. References to a resource in a different namespace are invalid for the moment, although we will revisit this in the future. A single CACertificateRef to a Kubernetes ConfigMap kind has 'Core' support. Implementations MAY choose to support attaching multiple certificates to a backend, but this behavior is implementation-specific. Support: Core - An optional single reference to a Kubernetes ConfigMap, with the CA certificate in a key named 'ca.crt'. Support: Implementation-specific (More than one reference, or other kinds of resources).",
								MarkdownDescription: "CACertificateRefs contains one or more references to Kubernetes objects that contain a PEM-encoded TLS CA certificate bundle, which is used to validate a TLS handshake between the Gateway and backend Pod. If CACertificateRefs is empty or unspecified, then WellKnownCACertificates must be specified. Only one of CACertificateRefs or WellKnownCACertificates may be specified, not both. If CACertifcateRefs is empty or unspecified, the configuration for WellKnownCACertificates MUST be honored instead if supported by the implementation. References to a resource in a different namespace are invalid for the moment, although we will revisit this in the future. A single CACertificateRef to a Kubernetes ConfigMap kind has 'Core' support. Implementations MAY choose to support attaching multiple certificates to a backend, but this behavior is implementation-specific. Support: Core - An optional single reference to a Kubernetes ConfigMap, with the CA certificate in a key named 'ca.crt'. Support: Implementation-specific (More than one reference, or other kinds of resources).",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"group": schema.StringAttribute{
											Description:         "Group is the group of the referent. For example, 'gateway.networking.k8s.io'. When unspecified or empty string, core API group is inferred.",
											MarkdownDescription: "Group is the group of the referent. For example, 'gateway.networking.k8s.io'. When unspecified or empty string, core API group is inferred.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(253),
												stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
											},
										},

										"kind": schema.StringAttribute{
											Description:         "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",
											MarkdownDescription: "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
											},
										},

										"name": schema.StringAttribute{
											Description:         "Name is the name of the referent.",
											MarkdownDescription: "Name is the name of the referent.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(253),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"hostname": schema.StringAttribute{
								Description:         "Hostname is used for two purposes in the connection between Gateways and backends: 1. Hostname MUST be used as the SNI to connect to the backend (RFC 6066). 2. If SubjectAltNames is not specified, Hostname MUST be used for authentication and MUST match the certificate served by the matching backend. Support: Core",
								MarkdownDescription: "Hostname is used for two purposes in the connection between Gateways and backends: 1. Hostname MUST be used as the SNI to connect to the backend (RFC 6066). 2. If SubjectAltNames is not specified, Hostname MUST be used for authentication and MUST match the certificate served by the matching backend. Support: Core",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
								},
							},

							"subject_alt_names": schema.ListNestedAttribute{
								Description:         "SubjectAltNames contains one or more Subject Alternative Names. When specified, the certificate served from the backend MUST have at least one Subject Alternate Name matching one of the specified SubjectAltNames. Support: Core",
								MarkdownDescription: "SubjectAltNames contains one or more Subject Alternative Names. When specified, the certificate served from the backend MUST have at least one Subject Alternate Name matching one of the specified SubjectAltNames. Support: Core",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"hostname": schema.StringAttribute{
											Description:         "Hostname contains Subject Alternative Name specified in DNS name format. Required when Type is set to Hostname, ignored otherwise. Support: Core",
											MarkdownDescription: "Hostname contains Subject Alternative Name specified in DNS name format. Required when Type is set to Hostname, ignored otherwise. Support: Core",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(253),
												stringvalidator.RegexMatches(regexp.MustCompile(`^(\*\.)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
											},
										},

										"type": schema.StringAttribute{
											Description:         "Type determines the format of the Subject Alternative Name. Always required. Support: Core",
											MarkdownDescription: "Type determines the format of the Subject Alternative Name. Always required. Support: Core",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Hostname", "URI"),
											},
										},

										"uri": schema.StringAttribute{
											Description:         "URI contains Subject Alternative Name specified in a full URI format. It MUST include both a scheme (e.g., 'http' or 'ftp') and a scheme-specific-part. Common values include SPIFFE IDs like 'spiffe://mycluster.example.com/ns/myns/sa/svc1sa'. Required when Type is set to URI, ignored otherwise. Support: Core",
											MarkdownDescription: "URI contains Subject Alternative Name specified in a full URI format. It MUST include both a scheme (e.g., 'http' or 'ftp') and a scheme-specific-part. Common values include SPIFFE IDs like 'spiffe://mycluster.example.com/ns/myns/sa/svc1sa'. Required when Type is set to URI, ignored otherwise. Support: Core",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(253),
												stringvalidator.RegexMatches(regexp.MustCompile(`^(([^:/?#]+):)(//([^/?#]*))([^?#]*)(\?([^#]*))?(#(.*))?`), ""),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"well_known_ca_certificates": schema.StringAttribute{
								Description:         "WellKnownCACertificates specifies whether system CA certificates may be used in the TLS handshake between the gateway and backend pod. If WellKnownCACertificates is unspecified or empty (''), then CACertificateRefs must be specified with at least one entry for a valid configuration. Only one of CACertificateRefs or WellKnownCACertificates may be specified, not both. If an implementation does not support the WellKnownCACertificates field or the value supplied is not supported, the Status Conditions on the Policy MUST be updated to include an Accepted: False Condition with Reason: Invalid. Support: Implementation-specific",
								MarkdownDescription: "WellKnownCACertificates specifies whether system CA certificates may be used in the TLS handshake between the gateway and backend pod. If WellKnownCACertificates is unspecified or empty (''), then CACertificateRefs must be specified with at least one entry for a valid configuration. Only one of CACertificateRefs or WellKnownCACertificates may be specified, not both. If an implementation does not support the WellKnownCACertificates field or the value supplied is not supported, the Status Conditions on the Policy MUST be updated to include an Accepted: False Condition with Reason: Invalid. Support: Implementation-specific",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("System"),
								},
							},
						},
						Required: true,
						Optional: false,
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

func (r *GatewayNetworkingK8SIoBackendTlspolicyV1Alpha3Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gateway_networking_k8s_io_backend_tls_policy_v1alpha3_manifest")

	var model GatewayNetworkingK8SIoBackendTlspolicyV1Alpha3ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("gateway.networking.k8s.io/v1alpha3")
	model.Kind = pointer.String("BackendTLSPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
