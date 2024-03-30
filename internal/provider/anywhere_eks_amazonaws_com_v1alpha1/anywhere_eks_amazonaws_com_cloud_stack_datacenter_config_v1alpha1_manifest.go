/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package anywhere_eks_amazonaws_com_v1alpha1

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
	_ datasource.DataSource = &AnywhereEksAmazonawsComCloudStackDatacenterConfigV1Alpha1Manifest{}
)

func NewAnywhereEksAmazonawsComCloudStackDatacenterConfigV1Alpha1Manifest() datasource.DataSource {
	return &AnywhereEksAmazonawsComCloudStackDatacenterConfigV1Alpha1Manifest{}
}

type AnywhereEksAmazonawsComCloudStackDatacenterConfigV1Alpha1Manifest struct{}

type AnywhereEksAmazonawsComCloudStackDatacenterConfigV1Alpha1ManifestData struct {
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
		Account           *string `tfsdk:"account" json:"account,omitempty"`
		AvailabilityZones *[]struct {
			Account               *string `tfsdk:"account" json:"account,omitempty"`
			CredentialsRef        *string `tfsdk:"credentials_ref" json:"credentialsRef,omitempty"`
			Domain                *string `tfsdk:"domain" json:"domain,omitempty"`
			ManagementApiEndpoint *string `tfsdk:"management_api_endpoint" json:"managementApiEndpoint,omitempty"`
			Name                  *string `tfsdk:"name" json:"name,omitempty"`
			Zone                  *struct {
				Id      *string `tfsdk:"id" json:"id,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Network *struct {
					Id   *string `tfsdk:"id" json:"id,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"network" json:"network,omitempty"`
			} `tfsdk:"zone" json:"zone,omitempty"`
		} `tfsdk:"availability_zones" json:"availabilityZones,omitempty"`
		Domain                *string `tfsdk:"domain" json:"domain,omitempty"`
		ManagementApiEndpoint *string `tfsdk:"management_api_endpoint" json:"managementApiEndpoint,omitempty"`
		Zones                 *[]struct {
			Id      *string `tfsdk:"id" json:"id,omitempty"`
			Name    *string `tfsdk:"name" json:"name,omitempty"`
			Network *struct {
				Id   *string `tfsdk:"id" json:"id,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"network" json:"network,omitempty"`
		} `tfsdk:"zones" json:"zones,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AnywhereEksAmazonawsComCloudStackDatacenterConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_anywhere_eks_amazonaws_com_cloud_stack_datacenter_config_v1alpha1_manifest"
}

func (r *AnywhereEksAmazonawsComCloudStackDatacenterConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CloudStackDatacenterConfig is the Schema for the cloudstackdatacenterconfigs API.",
		MarkdownDescription: "CloudStackDatacenterConfig is the Schema for the cloudstackdatacenterconfigs API.",
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
				Description:         "CloudStackDatacenterConfigSpec defines the desired state of CloudStackDatacenterConfig.",
				MarkdownDescription: "CloudStackDatacenterConfigSpec defines the desired state of CloudStackDatacenterConfig.",
				Attributes: map[string]schema.Attribute{
					"account": schema.StringAttribute{
						Description:         "Account typically represents a customer of the service provider or a department in a large organization. Multiple users can exist in an account, and all CloudStack resources belong to an account. Accounts have users and users have credentials to operate on resources within that account. If an account name is provided, a domain must also be provided. Deprecated: Please use AvailabilityZones instead",
						MarkdownDescription: "Account typically represents a customer of the service provider or a department in a large organization. Multiple users can exist in an account, and all CloudStack resources belong to an account. Accounts have users and users have credentials to operate on resources within that account. If an account name is provided, a domain must also be provided. Deprecated: Please use AvailabilityZones instead",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"availability_zones": schema.ListNestedAttribute{
						Description:         "AvailabilityZones list of different partitions to distribute VMs across - corresponds to a list of CAPI failure domains",
						MarkdownDescription: "AvailabilityZones list of different partitions to distribute VMs across - corresponds to a list of CAPI failure domains",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"account": schema.StringAttribute{
									Description:         "Account typically represents a customer of the service provider or a department in a large organization. Multiple users can exist in an account, and all CloudStack resources belong to an account. Accounts have users and users have credentials to operate on resources within that account. If an account name is provided, a domain must also be provided.",
									MarkdownDescription: "Account typically represents a customer of the service provider or a department in a large organization. Multiple users can exist in an account, and all CloudStack resources belong to an account. Accounts have users and users have credentials to operate on resources within that account. If an account name is provided, a domain must also be provided.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"credentials_ref": schema.StringAttribute{
									Description:         "CredentialRef is used to reference a secret in the eksa-system namespace",
									MarkdownDescription: "CredentialRef is used to reference a secret in the eksa-system namespace",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"domain": schema.StringAttribute{
									Description:         "Domain contains a grouping of accounts. Domains usually contain multiple accounts that have some logical relationship to each other and a set of delegated administrators with some authority over the domain and its subdomains This field is considered as a fully qualified domain name which is the same as the domain path without 'ROOT/' prefix. For example, if 'foo' is specified then a domain with 'ROOT/foo' domain path is picked. The value 'ROOT' is a special case that points to 'the' ROOT domain of the CloudStack. That is, a domain with a path 'ROOT/ROOT' is not allowed.",
									MarkdownDescription: "Domain contains a grouping of accounts. Domains usually contain multiple accounts that have some logical relationship to each other and a set of delegated administrators with some authority over the domain and its subdomains This field is considered as a fully qualified domain name which is the same as the domain path without 'ROOT/' prefix. For example, if 'foo' is specified then a domain with 'ROOT/foo' domain path is picked. The value 'ROOT' is a special case that points to 'the' ROOT domain of the CloudStack. That is, a domain with a path 'ROOT/ROOT' is not allowed.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"management_api_endpoint": schema.StringAttribute{
									Description:         "CloudStack Management API endpoint's IP. It is added to VM's noproxy list",
									MarkdownDescription: "CloudStack Management API endpoint's IP. It is added to VM's noproxy list",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is used as a unique identifier for each availability zone",
									MarkdownDescription: "Name is used as a unique identifier for each availability zone",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"zone": schema.SingleNestedAttribute{
									Description:         "Zone represents the properties of the CloudStack zone in which clusters should be created, like the network.",
									MarkdownDescription: "Zone represents the properties of the CloudStack zone in which clusters should be created, like the network.",
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Description:         "Zone is the name or UUID of the CloudStack zone in which clusters should be created. Zones should be managed by a single CloudStack Management endpoint.",
											MarkdownDescription: "Zone is the name or UUID of the CloudStack zone in which clusters should be created. Zones should be managed by a single CloudStack Management endpoint.",
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

										"network": schema.SingleNestedAttribute{
											Description:         "Network is the name or UUID of the CloudStack network in which clusters should be created. It can either be an isolated or shared network. If it doesn’t already exist in CloudStack, it’ll automatically be created by CAPC as an isolated network. It can either be specified as a UUID or name In multiple-zones situation, only 'Shared' network is supported.",
											MarkdownDescription: "Network is the name or UUID of the CloudStack network in which clusters should be created. It can either be an isolated or shared network. If it doesn’t already exist in CloudStack, it’ll automatically be created by CAPC as an isolated network. It can either be specified as a UUID or name In multiple-zones situation, only 'Shared' network is supported.",
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Description:         "Id of a resource in the CloudStack environment. Mutually exclusive with Name",
													MarkdownDescription: "Id of a resource in the CloudStack environment. Mutually exclusive with Name",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of a resource in the CloudStack environment. Mutually exclusive with Id",
													MarkdownDescription: "Name of a resource in the CloudStack environment. Mutually exclusive with Id",
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
									Required: true,
									Optional: false,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"domain": schema.StringAttribute{
						Description:         "Domain contains a grouping of accounts. Domains usually contain multiple accounts that have some logical relationship to each other and a set of delegated administrators with some authority over the domain and its subdomains This field is considered as a fully qualified domain name which is the same as the domain path without 'ROOT/' prefix. For example, if 'foo' is specified then a domain with 'ROOT/foo' domain path is picked. The value 'ROOT' is a special case that points to 'the' ROOT domain of the CloudStack. That is, a domain with a path 'ROOT/ROOT' is not allowed. Deprecated: Please use AvailabilityZones instead",
						MarkdownDescription: "Domain contains a grouping of accounts. Domains usually contain multiple accounts that have some logical relationship to each other and a set of delegated administrators with some authority over the domain and its subdomains This field is considered as a fully qualified domain name which is the same as the domain path without 'ROOT/' prefix. For example, if 'foo' is specified then a domain with 'ROOT/foo' domain path is picked. The value 'ROOT' is a special case that points to 'the' ROOT domain of the CloudStack. That is, a domain with a path 'ROOT/ROOT' is not allowed. Deprecated: Please use AvailabilityZones instead",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"management_api_endpoint": schema.StringAttribute{
						Description:         "CloudStack Management API endpoint's IP. It is added to VM's noproxy list Deprecated: Please use AvailabilityZones instead",
						MarkdownDescription: "CloudStack Management API endpoint's IP. It is added to VM's noproxy list Deprecated: Please use AvailabilityZones instead",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"zones": schema.ListNestedAttribute{
						Description:         "Zones is a list of one or more zones that are managed by a single CloudStack management endpoint. Deprecated: Please use AvailabilityZones instead",
						MarkdownDescription: "Zones is a list of one or more zones that are managed by a single CloudStack management endpoint. Deprecated: Please use AvailabilityZones instead",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description:         "Zone is the name or UUID of the CloudStack zone in which clusters should be created. Zones should be managed by a single CloudStack Management endpoint.",
									MarkdownDescription: "Zone is the name or UUID of the CloudStack zone in which clusters should be created. Zones should be managed by a single CloudStack Management endpoint.",
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

								"network": schema.SingleNestedAttribute{
									Description:         "Network is the name or UUID of the CloudStack network in which clusters should be created. It can either be an isolated or shared network. If it doesn’t already exist in CloudStack, it’ll automatically be created by CAPC as an isolated network. It can either be specified as a UUID or name In multiple-zones situation, only 'Shared' network is supported.",
									MarkdownDescription: "Network is the name or UUID of the CloudStack network in which clusters should be created. It can either be an isolated or shared network. If it doesn’t already exist in CloudStack, it’ll automatically be created by CAPC as an isolated network. It can either be specified as a UUID or name In multiple-zones situation, only 'Shared' network is supported.",
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Description:         "Id of a resource in the CloudStack environment. Mutually exclusive with Name",
											MarkdownDescription: "Id of a resource in the CloudStack environment. Mutually exclusive with Name",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of a resource in the CloudStack environment. Mutually exclusive with Id",
											MarkdownDescription: "Name of a resource in the CloudStack environment. Mutually exclusive with Id",
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

func (r *AnywhereEksAmazonawsComCloudStackDatacenterConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_anywhere_eks_amazonaws_com_cloud_stack_datacenter_config_v1alpha1_manifest")

	var model AnywhereEksAmazonawsComCloudStackDatacenterConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("anywhere.eks.amazonaws.com/v1alpha1")
	model.Kind = pointer.String("CloudStackDatacenterConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
