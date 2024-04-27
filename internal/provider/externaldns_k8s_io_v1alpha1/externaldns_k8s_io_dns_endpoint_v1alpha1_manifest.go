/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package externaldns_k8s_io_v1alpha1

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
	_ datasource.DataSource = &ExternaldnsK8SIoDnsendpointV1Alpha1Manifest{}
)

func NewExternaldnsK8SIoDnsendpointV1Alpha1Manifest() datasource.DataSource {
	return &ExternaldnsK8SIoDnsendpointV1Alpha1Manifest{}
}

type ExternaldnsK8SIoDnsendpointV1Alpha1Manifest struct{}

type ExternaldnsK8SIoDnsendpointV1Alpha1ManifestData struct {
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
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ExternaldnsK8SIoDnsendpointV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_externaldns_k8s_io_dns_endpoint_v1alpha1_manifest"
}

func (r *ExternaldnsK8SIoDnsendpointV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "DNSEndpointSpec defines the desired state of DNSEndpoint",
				MarkdownDescription: "DNSEndpointSpec defines the desired state of DNSEndpoint",
				Attributes: map[string]schema.Attribute{
					"endpoints": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *ExternaldnsK8SIoDnsendpointV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_externaldns_k8s_io_dns_endpoint_v1alpha1_manifest")

	var model ExternaldnsK8SIoDnsendpointV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("externaldns.k8s.io/v1alpha1")
	model.Kind = pointer.String("DNSEndpoint")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
