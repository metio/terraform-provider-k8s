/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package config_openshift_io_v1

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
	_ datasource.DataSource = &ConfigOpenshiftIoDnsV1Manifest{}
)

func NewConfigOpenshiftIoDnsV1Manifest() datasource.DataSource {
	return &ConfigOpenshiftIoDnsV1Manifest{}
}

type ConfigOpenshiftIoDnsV1Manifest struct{}

type ConfigOpenshiftIoDnsV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		BaseDomain *string `tfsdk:"base_domain" json:"baseDomain,omitempty"`
		Platform   *struct {
			Aws *struct {
				PrivateZoneIAMRole *string `tfsdk:"private_zone_iam_role" json:"privateZoneIAMRole,omitempty"`
			} `tfsdk:"aws" json:"aws,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"platform" json:"platform,omitempty"`
		PrivateZone *struct {
			Id   *string            `tfsdk:"id" json:"id,omitempty"`
			Tags *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		} `tfsdk:"private_zone" json:"privateZone,omitempty"`
		PublicZone *struct {
			Id   *string            `tfsdk:"id" json:"id,omitempty"`
			Tags *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		} `tfsdk:"public_zone" json:"publicZone,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ConfigOpenshiftIoDnsV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_config_openshift_io_dns_v1_manifest"
}

func (r *ConfigOpenshiftIoDnsV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DNS holds cluster-wide information about DNS. The canonical name is 'cluster'  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "DNS holds cluster-wide information about DNS. The canonical name is 'cluster'  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
					"base_domain": schema.StringAttribute{
						Description:         "baseDomain is the base domain of the cluster. All managed DNS records will be sub-domains of this base.  For example, given the base domain 'openshift.example.com', an API server DNS record may be created for 'cluster-api.openshift.example.com'.  Once set, this field cannot be changed.",
						MarkdownDescription: "baseDomain is the base domain of the cluster. All managed DNS records will be sub-domains of this base.  For example, given the base domain 'openshift.example.com', an API server DNS record may be created for 'cluster-api.openshift.example.com'.  Once set, this field cannot be changed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"platform": schema.SingleNestedAttribute{
						Description:         "platform holds configuration specific to the underlying infrastructure provider for DNS. When omitted, this means the user has no opinion and the platform is left to choose reasonable defaults. These defaults are subject to change over time.",
						MarkdownDescription: "platform holds configuration specific to the underlying infrastructure provider for DNS. When omitted, this means the user has no opinion and the platform is left to choose reasonable defaults. These defaults are subject to change over time.",
						Attributes: map[string]schema.Attribute{
							"aws": schema.SingleNestedAttribute{
								Description:         "aws contains DNS configuration specific to the Amazon Web Services cloud provider.",
								MarkdownDescription: "aws contains DNS configuration specific to the Amazon Web Services cloud provider.",
								Attributes: map[string]schema.Attribute{
									"private_zone_iam_role": schema.StringAttribute{
										Description:         "privateZoneIAMRole contains the ARN of an IAM role that should be assumed when performing operations on the cluster's private hosted zone specified in the cluster DNS config. When left empty, no role should be assumed.",
										MarkdownDescription: "privateZoneIAMRole contains the ARN of an IAM role that should be assumed when performing operations on the cluster's private hosted zone specified in the cluster DNS config. When left empty, no role should be assumed.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^arn:(aws|aws-cn|aws-us-gov):iam::[0-9]{12}:role\/.*$`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": schema.StringAttribute{
								Description:         "type is the underlying infrastructure provider for the cluster. Allowed values: '', 'AWS'.  Individual components may not support all platforms, and must handle unrecognized platforms with best-effort defaults.",
								MarkdownDescription: "type is the underlying infrastructure provider for the cluster. Allowed values: '', 'AWS'.  Individual components may not support all platforms, and must handle unrecognized platforms with best-effort defaults.",
								Required:            true,
								Optional:            false,
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

					"private_zone": schema.SingleNestedAttribute{
						Description:         "privateZone is the location where all the DNS records that are only available internally to the cluster exist.  If this field is nil, no private records should be created.  Once set, this field cannot be changed.",
						MarkdownDescription: "privateZone is the location where all the DNS records that are only available internally to the cluster exist.  If this field is nil, no private records should be created.  Once set, this field cannot be changed.",
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description:         "id is the identifier that can be used to find the DNS hosted zone.  on AWS zone can be fetched using 'ID' as id in [1] on Azure zone can be fetched using 'ID' as a pre-determined name in [2], on GCP zone can be fetched using 'ID' as a pre-determined name in [3].  [1]: https://docs.aws.amazon.com/cli/latest/reference/route53/get-hosted-zone.html#options [2]: https://docs.microsoft.com/en-us/cli/azure/network/dns/zone?view=azure-cli-latest#az-network-dns-zone-show [3]: https://cloud.google.com/dns/docs/reference/v1/managedZones/get",
								MarkdownDescription: "id is the identifier that can be used to find the DNS hosted zone.  on AWS zone can be fetched using 'ID' as id in [1] on Azure zone can be fetched using 'ID' as a pre-determined name in [2], on GCP zone can be fetched using 'ID' as a pre-determined name in [3].  [1]: https://docs.aws.amazon.com/cli/latest/reference/route53/get-hosted-zone.html#options [2]: https://docs.microsoft.com/en-us/cli/azure/network/dns/zone?view=azure-cli-latest#az-network-dns-zone-show [3]: https://cloud.google.com/dns/docs/reference/v1/managedZones/get",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tags": schema.MapAttribute{
								Description:         "tags can be used to query the DNS hosted zone.  on AWS, resourcegroupstaggingapi [1] can be used to fetch a zone using 'Tags' as tag-filters,  [1]: https://docs.aws.amazon.com/cli/latest/reference/resourcegroupstaggingapi/get-resources.html#options",
								MarkdownDescription: "tags can be used to query the DNS hosted zone.  on AWS, resourcegroupstaggingapi [1] can be used to fetch a zone using 'Tags' as tag-filters,  [1]: https://docs.aws.amazon.com/cli/latest/reference/resourcegroupstaggingapi/get-resources.html#options",
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

					"public_zone": schema.SingleNestedAttribute{
						Description:         "publicZone is the location where all the DNS records that are publicly accessible to the internet exist.  If this field is nil, no public records should be created.  Once set, this field cannot be changed.",
						MarkdownDescription: "publicZone is the location where all the DNS records that are publicly accessible to the internet exist.  If this field is nil, no public records should be created.  Once set, this field cannot be changed.",
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description:         "id is the identifier that can be used to find the DNS hosted zone.  on AWS zone can be fetched using 'ID' as id in [1] on Azure zone can be fetched using 'ID' as a pre-determined name in [2], on GCP zone can be fetched using 'ID' as a pre-determined name in [3].  [1]: https://docs.aws.amazon.com/cli/latest/reference/route53/get-hosted-zone.html#options [2]: https://docs.microsoft.com/en-us/cli/azure/network/dns/zone?view=azure-cli-latest#az-network-dns-zone-show [3]: https://cloud.google.com/dns/docs/reference/v1/managedZones/get",
								MarkdownDescription: "id is the identifier that can be used to find the DNS hosted zone.  on AWS zone can be fetched using 'ID' as id in [1] on Azure zone can be fetched using 'ID' as a pre-determined name in [2], on GCP zone can be fetched using 'ID' as a pre-determined name in [3].  [1]: https://docs.aws.amazon.com/cli/latest/reference/route53/get-hosted-zone.html#options [2]: https://docs.microsoft.com/en-us/cli/azure/network/dns/zone?view=azure-cli-latest#az-network-dns-zone-show [3]: https://cloud.google.com/dns/docs/reference/v1/managedZones/get",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tags": schema.MapAttribute{
								Description:         "tags can be used to query the DNS hosted zone.  on AWS, resourcegroupstaggingapi [1] can be used to fetch a zone using 'Tags' as tag-filters,  [1]: https://docs.aws.amazon.com/cli/latest/reference/resourcegroupstaggingapi/get-resources.html#options",
								MarkdownDescription: "tags can be used to query the DNS hosted zone.  on AWS, resourcegroupstaggingapi [1] can be used to fetch a zone using 'Tags' as tag-filters,  [1]: https://docs.aws.amazon.com/cli/latest/reference/resourcegroupstaggingapi/get-resources.html#options",
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
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ConfigOpenshiftIoDnsV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_config_openshift_io_dns_v1_manifest")

	var model ConfigOpenshiftIoDnsV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("config.openshift.io/v1")
	model.Kind = pointer.String("DNS")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
