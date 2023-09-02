/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hive_openshift_io_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &HiveOpenshiftIoDNSZoneV1DataSource{}
	_ datasource.DataSourceWithConfigure = &HiveOpenshiftIoDNSZoneV1DataSource{}
)

func NewHiveOpenshiftIoDNSZoneV1DataSource() datasource.DataSource {
	return &HiveOpenshiftIoDNSZoneV1DataSource{}
}

type HiveOpenshiftIoDNSZoneV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type HiveOpenshiftIoDNSZoneV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *HiveOpenshiftIoDNSZoneV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hive_openshift_io_dns_zone_v1"
}

func (r *HiveOpenshiftIoDNSZoneV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
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
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the value for the tag",
											MarkdownDescription: "Value is the value for the tag",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"credentials_assume_role": schema.SingleNestedAttribute{
								Description:         "CredentialsAssumeRole refers to the IAM role that must be assumed to obtain AWS account access for the DNS CRUD operations.",
								MarkdownDescription: "CredentialsAssumeRole refers to the IAM role that must be assumed to obtain AWS account access for the DNS CRUD operations.",
								Attributes: map[string]schema.Attribute{
									"external_id": schema.StringAttribute{
										Description:         "ExternalID is random string generated by platform so that assume role is protected from confused deputy problem. more info: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html",
										MarkdownDescription: "ExternalID is random string generated by platform so that assume role is protected from confused deputy problem. more info: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"role_arn": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"credentials_secret_ref": schema.SingleNestedAttribute{
								Description:         "CredentialsSecretRef contains a reference to a secret that contains AWS credentials for CRUD operations",
								MarkdownDescription: "CredentialsSecretRef contains a reference to a secret that contains AWS credentials for CRUD operations",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"region": schema.StringAttribute{
								Description:         "Region is the AWS region to use for route53 operations. This defaults to us-east-1. For AWS China, use cn-northwest-1.",
								MarkdownDescription: "Region is the AWS region to use for route53 operations. This defaults to us-east-1. For AWS China, use cn-northwest-1.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"azure": schema.SingleNestedAttribute{
						Description:         "Azure specifes Azure-specific cloud configuration",
						MarkdownDescription: "Azure specifes Azure-specific cloud configuration",
						Attributes: map[string]schema.Attribute{
							"cloud_name": schema.StringAttribute{
								Description:         "CloudName is the name of the Azure cloud environment which can be used to configure the Azure SDK with the appropriate Azure API endpoints. If empty, the value is equal to 'AzurePublicCloud'.",
								MarkdownDescription: "CloudName is the name of the Azure cloud environment which can be used to configure the Azure SDK with the appropriate Azure API endpoints. If empty, the value is equal to 'AzurePublicCloud'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"credentials_secret_ref": schema.SingleNestedAttribute{
								Description:         "CredentialsSecretRef references a secret that will be used to authenticate with Azure CloudDNS. It will need permission to create and manage CloudDNS Hosted Zones. Secret should have a key named 'osServicePrincipal.json'. The credentials must specify the project to use.",
								MarkdownDescription: "CredentialsSecretRef references a secret that will be used to authenticate with Azure CloudDNS. It will need permission to create and manage CloudDNS Hosted Zones. Secret should have a key named 'osServicePrincipal.json'. The credentials must specify the project to use.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"resource_group_name": schema.StringAttribute{
								Description:         "ResourceGroupName specifies the Azure resource group in which the Hosted Zone should be created.",
								MarkdownDescription: "ResourceGroupName specifies the Azure resource group in which the Hosted Zone should be created.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
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
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"link_to_parent_domain": schema.BoolAttribute{
						Description:         "LinkToParentDomain specifies whether DNS records should be automatically created to link this DNSZone with a parent domain.",
						MarkdownDescription: "LinkToParentDomain specifies whether DNS records should be automatically created to link this DNSZone with a parent domain.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"preserve_on_delete": schema.BoolAttribute{
						Description:         "PreserveOnDelete allows the user to disconnect a DNSZone from Hive without deprovisioning it. This can also be used to abandon ongoing DNSZone deprovision. Typically set automatically due to PreserveOnDelete being set on a ClusterDeployment.",
						MarkdownDescription: "PreserveOnDelete allows the user to disconnect a DNSZone from Hive without deprovisioning it. This can also be used to abandon ongoing DNSZone deprovision. Typically set automatically due to PreserveOnDelete being set on a ClusterDeployment.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"zone": schema.StringAttribute{
						Description:         "Zone is the DNS zone to host",
						MarkdownDescription: "Zone is the DNS zone to host",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *HiveOpenshiftIoDNSZoneV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *HiveOpenshiftIoDNSZoneV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_hive_openshift_io_dns_zone_v1")

	var data HiveOpenshiftIoDNSZoneV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hive.openshift.io", Version: "v1", Resource: "DNSZone"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse HiveOpenshiftIoDNSZoneV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("hive.openshift.io/v1")
	data.Kind = pointer.String("DNSZone")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
