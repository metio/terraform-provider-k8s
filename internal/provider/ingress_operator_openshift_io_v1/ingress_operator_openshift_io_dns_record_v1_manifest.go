/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ingress_operator_openshift_io_v1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &IngressOperatorOpenshiftIoDnsrecordV1Manifest{}
)

func NewIngressOperatorOpenshiftIoDnsrecordV1Manifest() datasource.DataSource {
	return &IngressOperatorOpenshiftIoDnsrecordV1Manifest{}
}

type IngressOperatorOpenshiftIoDnsrecordV1Manifest struct{}

type IngressOperatorOpenshiftIoDnsrecordV1ManifestData struct {
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
		DnsManagementPolicy *string   `tfsdk:"dns_management_policy" json:"dnsManagementPolicy,omitempty"`
		DnsName             *string   `tfsdk:"dns_name" json:"dnsName,omitempty"`
		RecordTTL           *int64    `tfsdk:"record_ttl" json:"recordTTL,omitempty"`
		RecordType          *string   `tfsdk:"record_type" json:"recordType,omitempty"`
		Targets             *[]string `tfsdk:"targets" json:"targets,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *IngressOperatorOpenshiftIoDnsrecordV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ingress_operator_openshift_io_dns_record_v1_manifest"
}

func (r *IngressOperatorOpenshiftIoDnsrecordV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DNSRecord is a DNS record managed in the zones defined by dns.config.openshift.io/cluster .spec.publicZone and .spec.privateZone.  Cluster admin manipulation of this resource is not supported. This resource is only for internal communication of OpenShift operators.  If DNSManagementPolicy is 'Unmanaged', the operator will not be responsible for managing the DNS records on the cloud provider.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "DNSRecord is a DNS record managed in the zones defined by dns.config.openshift.io/cluster .spec.publicZone and .spec.privateZone.  Cluster admin manipulation of this resource is not supported. This resource is only for internal communication of OpenShift operators.  If DNSManagementPolicy is 'Unmanaged', the operator will not be responsible for managing the DNS records on the cloud provider.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
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
				Description:         "spec is the specification of the desired behavior of the dnsRecord.",
				MarkdownDescription: "spec is the specification of the desired behavior of the dnsRecord.",
				Attributes: map[string]schema.Attribute{
					"dns_management_policy": schema.StringAttribute{
						Description:         "dnsManagementPolicy denotes the current policy applied on the DNS record. Records that have policy set as 'Unmanaged' are ignored by the ingress operator.  This means that the DNS record on the cloud provider is not managed by the operator, and the 'Published' status condition will be updated to 'Unknown' status, since it is externally managed. Any existing record on the cloud provider can be deleted at the discretion of the cluster admin.  This field defaults to Managed. Valid values are 'Managed' and 'Unmanaged'.",
						MarkdownDescription: "dnsManagementPolicy denotes the current policy applied on the DNS record. Records that have policy set as 'Unmanaged' are ignored by the ingress operator.  This means that the DNS record on the cloud provider is not managed by the operator, and the 'Published' status condition will be updated to 'Unknown' status, since it is externally managed. Any existing record on the cloud provider can be deleted at the discretion of the cluster admin.  This field defaults to Managed. Valid values are 'Managed' and 'Unmanaged'.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Managed", "Unmanaged"),
						},
					},

					"dns_name": schema.StringAttribute{
						Description:         "dnsName is the hostname of the DNS record",
						MarkdownDescription: "dnsName is the hostname of the DNS record",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"record_ttl": schema.Int64Attribute{
						Description:         "recordTTL is the record TTL in seconds. If zero, the default is 30. RecordTTL will not be used in AWS regions Alias targets, but will be used in CNAME targets, per AWS API contract.",
						MarkdownDescription: "recordTTL is the record TTL in seconds. If zero, the default is 30. RecordTTL will not be used in AWS regions Alias targets, but will be used in CNAME targets, per AWS API contract.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"record_type": schema.StringAttribute{
						Description:         "recordType is the DNS record type. For example, 'A' or 'CNAME'.",
						MarkdownDescription: "recordType is the DNS record type. For example, 'A' or 'CNAME'.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("CNAME", "A"),
						},
					},

					"targets": schema.ListAttribute{
						Description:         "targets are record targets.",
						MarkdownDescription: "targets are record targets.",
						ElementType:         types.StringType,
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
	}
}

func (r *IngressOperatorOpenshiftIoDnsrecordV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ingress_operator_openshift_io_dns_record_v1_manifest")

	var model IngressOperatorOpenshiftIoDnsrecordV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("ingress.operator.openshift.io/v1")
	model.Kind = pointer.String("DNSRecord")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
