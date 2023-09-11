/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hive_openshift_io_v1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &HiveOpenshiftIoDnszoneV1Manifest{}
)

func NewHiveOpenshiftIoDnszoneV1Manifest() datasource.DataSource {
	return &HiveOpenshiftIoDnszoneV1Manifest{}
}

type HiveOpenshiftIoDnszoneV1Manifest struct{}

type HiveOpenshiftIoDnszoneV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		Aws *struct {
			AdditionalTags *[]struct {
				Key   *string `tfsdk:"key" json:"key,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"additional_tags" json:"additionalTags,omitempty"`
			CredentialsAssumeRole *struct {
				ExternalID *string `tfsdk:"external_id" json:"externalID,omitempty"`
				RoleARN    *string `tfsdk:"role_arn" json:"roleARN,omitempty"`
			} `tfsdk:"credentials_assume_role" json:"credentialsAssumeRole,omitempty"`
			CredentialsSecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
			Region *string `tfsdk:"region" json:"region,omitempty"`
		} `tfsdk:"aws" json:"aws,omitempty"`
		Azure *struct {
			CloudName            *string `tfsdk:"cloud_name" json:"cloudName,omitempty"`
			CredentialsSecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
			ResourceGroupName *string `tfsdk:"resource_group_name" json:"resourceGroupName,omitempty"`
		} `tfsdk:"azure" json:"azure,omitempty"`
		Gcp *struct {
			CredentialsSecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"credentials_secret_ref" json:"credentialsSecretRef,omitempty"`
		} `tfsdk:"gcp" json:"gcp,omitempty"`
		LinkToParentDomain *bool   `tfsdk:"link_to_parent_domain" json:"linkToParentDomain,omitempty"`
		PreserveOnDelete   *bool   `tfsdk:"preserve_on_delete" json:"preserveOnDelete,omitempty"`
		Zone               *string `tfsdk:"zone" json:"zone,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *HiveOpenshiftIoDnszoneV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hive_openshift_io_dns_zone_v1_manifest"
}

func (r *HiveOpenshiftIoDnszoneV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DNSZone is the Schema for the dnszones API",
		MarkdownDescription: "DNSZone is the Schema for the dnszones API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
				Description:         "DNSZoneSpec defines the desired state of DNSZone",
				MarkdownDescription: "DNSZoneSpec defines the desired state of DNSZone",
				Attributes: map[string]schema.Attribute{
					"aws": schema.SingleNestedAttribute{
						Description:         "AWS specifies AWS-specific cloud configuration",
						MarkdownDescription: "AWS specifies AWS-specific cloud configuration",
						Attributes: map[string]schema.Attribute{
							"additional_tags": schema.ListNestedAttribute{
								Description:         "AdditionalTags is a set of additional tags to set on the DNS hosted zone. In addition to these tags,the DNS Zone controller will set a hive.openhsift.io/hostedzone tag identifying the HostedZone record that it belongs to.",
								MarkdownDescription: "AdditionalTags is a set of additional tags to set on the DNS hosted zone. In addition to these tags,the DNS Zone controller will set a hive.openhsift.io/hostedzone tag identifying the HostedZone record that it belongs to.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "Key is the key for the tag",
											MarkdownDescription: "Key is the key for the tag",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the value for the tag",
											MarkdownDescription: "Value is the value for the tag",
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

							"credentials_assume_role": schema.SingleNestedAttribute{
								Description:         "CredentialsAssumeRole refers to the IAM role that must be assumed to obtain AWS account access for the DNS CRUD operations.",
								MarkdownDescription: "CredentialsAssumeRole refers to the IAM role that must be assumed to obtain AWS account access for the DNS CRUD operations.",
								Attributes: map[string]schema.Attribute{
									"external_id": schema.StringAttribute{
										Description:         "ExternalID is random string generated by platform so that assume role is protected from confused deputy problem. more info: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html",
										MarkdownDescription: "ExternalID is random string generated by platform so that assume role is protected from confused deputy problem. more info: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"role_arn": schema.StringAttribute{
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

							"credentials_secret_ref": schema.SingleNestedAttribute{
								Description:         "CredentialsSecretRef contains a reference to a secret that contains AWS credentials for CRUD operations",
								MarkdownDescription: "CredentialsSecretRef contains a reference to a secret that contains AWS credentials for CRUD operations",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"region": schema.StringAttribute{
								Description:         "Region is the AWS region to use for route53 operations. This defaults to us-east-1. For AWS China, use cn-northwest-1.",
								MarkdownDescription: "Region is the AWS region to use for route53 operations. This defaults to us-east-1. For AWS China, use cn-northwest-1.",
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
						Description:         "Azure specifes Azure-specific cloud configuration",
						MarkdownDescription: "Azure specifes Azure-specific cloud configuration",
						Attributes: map[string]schema.Attribute{
							"cloud_name": schema.StringAttribute{
								Description:         "CloudName is the name of the Azure cloud environment which can be used to configure the Azure SDK with the appropriate Azure API endpoints. If empty, the value is equal to 'AzurePublicCloud'.",
								MarkdownDescription: "CloudName is the name of the Azure cloud environment which can be used to configure the Azure SDK with the appropriate Azure API endpoints. If empty, the value is equal to 'AzurePublicCloud'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("", "AzurePublicCloud", "AzureUSGovernmentCloud", "AzureChinaCloud", "AzureGermanCloud"),
								},
							},

							"credentials_secret_ref": schema.SingleNestedAttribute{
								Description:         "CredentialsSecretRef references a secret that will be used to authenticate with Azure CloudDNS. It will need permission to create and manage CloudDNS Hosted Zones. Secret should have a key named 'osServicePrincipal.json'. The credentials must specify the project to use.",
								MarkdownDescription: "CredentialsSecretRef references a secret that will be used to authenticate with Azure CloudDNS. It will need permission to create and manage CloudDNS Hosted Zones. Secret should have a key named 'osServicePrincipal.json'. The credentials must specify the project to use.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"resource_group_name": schema.StringAttribute{
								Description:         "ResourceGroupName specifies the Azure resource group in which the Hosted Zone should be created.",
								MarkdownDescription: "ResourceGroupName specifies the Azure resource group in which the Hosted Zone should be created.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"gcp": schema.SingleNestedAttribute{
						Description:         "GCP specifies GCP-specific cloud configuration",
						MarkdownDescription: "GCP specifies GCP-specific cloud configuration",
						Attributes: map[string]schema.Attribute{
							"credentials_secret_ref": schema.SingleNestedAttribute{
								Description:         "CredentialsSecretRef references a secret that will be used to authenticate with GCP CloudDNS. It will need permission to create and manage CloudDNS Hosted Zones. Secret should have a key named 'osServiceAccount.json'. The credentials must specify the project to use.",
								MarkdownDescription: "CredentialsSecretRef references a secret that will be used to authenticate with GCP CloudDNS. It will need permission to create and manage CloudDNS Hosted Zones. Secret should have a key named 'osServiceAccount.json'. The credentials must specify the project to use.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"link_to_parent_domain": schema.BoolAttribute{
						Description:         "LinkToParentDomain specifies whether DNS records should be automatically created to link this DNSZone with a parent domain.",
						MarkdownDescription: "LinkToParentDomain specifies whether DNS records should be automatically created to link this DNSZone with a parent domain.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"preserve_on_delete": schema.BoolAttribute{
						Description:         "PreserveOnDelete allows the user to disconnect a DNSZone from Hive without deprovisioning it. This can also be used to abandon ongoing DNSZone deprovision. Typically set automatically due to PreserveOnDelete being set on a ClusterDeployment.",
						MarkdownDescription: "PreserveOnDelete allows the user to disconnect a DNSZone from Hive without deprovisioning it. This can also be used to abandon ongoing DNSZone deprovision. Typically set automatically due to PreserveOnDelete being set on a ClusterDeployment.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"zone": schema.StringAttribute{
						Description:         "Zone is the DNS zone to host",
						MarkdownDescription: "Zone is the DNS zone to host",
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

func (r *HiveOpenshiftIoDnszoneV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hive_openshift_io_dns_zone_v1_manifest")

	var model HiveOpenshiftIoDnszoneV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("hive.openshift.io/v1")
	model.Kind = pointer.String("DNSZone")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
