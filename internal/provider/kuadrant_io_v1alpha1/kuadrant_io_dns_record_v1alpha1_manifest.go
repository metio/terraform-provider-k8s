/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuadrant_io_v1alpha1

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
	_ datasource.DataSource = &KuadrantIoDnsrecordV1Alpha1Manifest{}
)

func NewKuadrantIoDnsrecordV1Alpha1Manifest() datasource.DataSource {
	return &KuadrantIoDnsrecordV1Alpha1Manifest{}
}

type KuadrantIoDnsrecordV1Alpha1Manifest struct{}

type KuadrantIoDnsrecordV1Alpha1ManifestData struct {
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
		Endpoints *[]struct {
			DnsName          *string            `tfsdk:"dns_name" json:"dnsName,omitempty"`
			Labels           *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			ProviderSpecific *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"provider_specific" json:"providerSpecific,omitempty"`
			RecordTTL     *int64    `tfsdk:"record_ttl" json:"recordTTL,omitempty"`
			RecordType    *string   `tfsdk:"record_type" json:"recordType,omitempty"`
			SetIdentifier *string   `tfsdk:"set_identifier" json:"setIdentifier,omitempty"`
			Targets       *[]string `tfsdk:"targets" json:"targets,omitempty"`
		} `tfsdk:"endpoints" json:"endpoints,omitempty"`
		HealthCheck *struct {
			Endpoint         *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
			FailureThreshold *int64  `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
			Port             *int64  `tfsdk:"port" json:"port,omitempty"`
			Protocol         *string `tfsdk:"protocol" json:"protocol,omitempty"`
		} `tfsdk:"health_check" json:"healthCheck,omitempty"`
		ManagedZone *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"managed_zone" json:"managedZone,omitempty"`
		OwnerID  *string `tfsdk:"owner_id" json:"ownerID,omitempty"`
		RootHost *string `tfsdk:"root_host" json:"rootHost,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KuadrantIoDnsrecordV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuadrant_io_dns_record_v1alpha1_manifest"
}

func (r *KuadrantIoDnsrecordV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DNSRecord is the Schema for the dnsrecords API",
		MarkdownDescription: "DNSRecord is the Schema for the dnsrecords API",
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
				Description:         "DNSRecordSpec defines the desired state of DNSRecord",
				MarkdownDescription: "DNSRecordSpec defines the desired state of DNSRecord",
				Attributes: map[string]schema.Attribute{
					"endpoints": schema.ListNestedAttribute{
						Description:         "endpoints is a list of endpoints that will be published into the dns provider.",
						MarkdownDescription: "endpoints is a list of endpoints that will be published into the dns provider.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"dns_name": schema.StringAttribute{
									Description:         "The hostname of the DNS record",
									MarkdownDescription: "The hostname of the DNS record",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels stores labels defined for the Endpoint",
									MarkdownDescription: "Labels stores labels defined for the Endpoint",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"provider_specific": schema.ListNestedAttribute{
									Description:         "ProviderSpecific stores provider specific config",
									MarkdownDescription: "ProviderSpecific stores provider specific config",
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
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"record_ttl": schema.Int64Attribute{
									Description:         "TTL for the record",
									MarkdownDescription: "TTL for the record",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"record_type": schema.StringAttribute{
									Description:         "RecordType type of record, e.g. CNAME, A, AAAA, SRV, TXT etc",
									MarkdownDescription: "RecordType type of record, e.g. CNAME, A, AAAA, SRV, TXT etc",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"set_identifier": schema.StringAttribute{
									Description:         "Identifier to distinguish multiple records with the same name and type (e.g. Route53 records with routing policies other than 'simple')",
									MarkdownDescription: "Identifier to distinguish multiple records with the same name and type (e.g. Route53 records with routing policies other than 'simple')",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"targets": schema.ListAttribute{
									Description:         "The targets the DNS record points to",
									MarkdownDescription: "The targets the DNS record points to",
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

					"health_check": schema.SingleNestedAttribute{
						Description:         "HealthCheckSpec configures health checks in the DNS provider.By default this health check will be applied to each unique DNS A Record forthe listeners assigned to the target gateway",
						MarkdownDescription: "HealthCheckSpec configures health checks in the DNS provider.By default this health check will be applied to each unique DNS A Record forthe listeners assigned to the target gateway",
						Attributes: map[string]schema.Attribute{
							"endpoint": schema.StringAttribute{
								Description:         "Endpoint is the path to append to the host to reach the expected health check.Must start with '?' or '/', contain only valid URL characters and end with alphanumeric char or '/'. For example '/' or '/healthz' are common",
								MarkdownDescription: "Endpoint is the path to append to the host to reach the expected health check.Must start with '?' or '/', contain only valid URL characters and end with alphanumeric char or '/'. For example '/' or '/healthz' are common",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(?:\?|\/)[\w\-.~:\/?#\[\]@!$&'()*+,;=]+(?:[a-zA-Z0-9]|\/){1}$`), ""),
								},
							},

							"failure_threshold": schema.Int64Attribute{
								Description:         "FailureThreshold is a limit of consecutive failures that must occur for a host to be considered unhealthy",
								MarkdownDescription: "FailureThreshold is a limit of consecutive failures that must occur for a host to be considered unhealthy",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "Port to connect to the host on. Must be either 80, 443 or 1024-49151",
								MarkdownDescription: "Port to connect to the host on. Must be either 80, 443 or 1024-49151",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"protocol": schema.StringAttribute{
								Description:         "Protocol to use when connecting to the host, valid values are 'HTTP' or 'HTTPS'",
								MarkdownDescription: "Protocol to use when connecting to the host, valid values are 'HTTP' or 'HTTPS'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"managed_zone": schema.SingleNestedAttribute{
						Description:         "managedZone is a reference to a ManagedZone instance to which this record will publish its endpoints.",
						MarkdownDescription: "managedZone is a reference to a ManagedZone instance to which this record will publish its endpoints.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "'name' is the name of the managed zone.Required",
								MarkdownDescription: "'name' is the name of the managed zone.Required",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"owner_id": schema.StringAttribute{
						Description:         "ownerID is a unique string used to identify the owner of this record.If unset or set to an empty string the record UID will be used.",
						MarkdownDescription: "ownerID is a unique string used to identify the owner of this record.If unset or set to an empty string the record UID will be used.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(6),
							stringvalidator.LengthAtMost(36),
						},
					},

					"root_host": schema.StringAttribute{
						Description:         "rootHost is the single root for all endpoints in a DNSRecord.it is expected all defined endpoints are children of or equal to this rootHostMust contain at least two groups of valid URL characters separated by a '.'",
						MarkdownDescription: "rootHost is the single root for all endpoints in a DNSRecord.it is expected all defined endpoints are children of or equal to this rootHostMust contain at least two groups of valid URL characters separated by a '.'",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.LengthAtMost(255),
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?:[\w\-.~:\/?#[\]@!$&'()*+,;=]+)\.(?:[\w\-.~:\/?#[\]@!$&'()*+,;=]+)$`), ""),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *KuadrantIoDnsrecordV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kuadrant_io_dns_record_v1alpha1_manifest")

	var model KuadrantIoDnsrecordV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kuadrant.io/v1alpha1")
	model.Kind = pointer.String("DNSRecord")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
