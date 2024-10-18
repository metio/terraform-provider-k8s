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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &AnywhereEksAmazonawsComNutanixDatacenterConfigV1Alpha1Manifest{}
)

func NewAnywhereEksAmazonawsComNutanixDatacenterConfigV1Alpha1Manifest() datasource.DataSource {
	return &AnywhereEksAmazonawsComNutanixDatacenterConfigV1Alpha1Manifest{}
}

type AnywhereEksAmazonawsComNutanixDatacenterConfigV1Alpha1Manifest struct{}

type AnywhereEksAmazonawsComNutanixDatacenterConfigV1Alpha1ManifestData struct {
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
		AdditionalTrustBundle *string `tfsdk:"additional_trust_bundle" json:"additionalTrustBundle,omitempty"`
		CredentialRef         *struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"credential_ref" json:"credentialRef,omitempty"`
		Endpoint       *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
		FailureDomains *[]struct {
			Cluster *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
				Uuid *string `tfsdk:"uuid" json:"uuid,omitempty"`
			} `tfsdk:"cluster" json:"cluster,omitempty"`
			Name    *string `tfsdk:"name" json:"name,omitempty"`
			Subnets *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
				Uuid *string `tfsdk:"uuid" json:"uuid,omitempty"`
			} `tfsdk:"subnets" json:"subnets,omitempty"`
			WorkerMachineGroups *[]string `tfsdk:"worker_machine_groups" json:"workerMachineGroups,omitempty"`
		} `tfsdk:"failure_domains" json:"failureDomains,omitempty"`
		Insecure *bool  `tfsdk:"insecure" json:"insecure,omitempty"`
		Port     *int64 `tfsdk:"port" json:"port,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AnywhereEksAmazonawsComNutanixDatacenterConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_anywhere_eks_amazonaws_com_nutanix_datacenter_config_v1alpha1_manifest"
}

func (r *AnywhereEksAmazonawsComNutanixDatacenterConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "NutanixDatacenterConfig is the Schema for the NutanixDatacenterConfigs API",
		MarkdownDescription: "NutanixDatacenterConfig is the Schema for the NutanixDatacenterConfigs API",
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
				Description:         "NutanixDatacenterConfigSpec defines the desired state of NutanixDatacenterConfig.",
				MarkdownDescription: "NutanixDatacenterConfigSpec defines the desired state of NutanixDatacenterConfig.",
				Attributes: map[string]schema.Attribute{
					"additional_trust_bundle": schema.StringAttribute{
						Description:         "AdditionalTrustBundle is the optional PEM-encoded certificate bundle for users that configured their Prism Central with certificates from non-publicly trusted CAs",
						MarkdownDescription: "AdditionalTrustBundle is the optional PEM-encoded certificate bundle for users that configured their Prism Central with certificates from non-publicly trusted CAs",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"credential_ref": schema.SingleNestedAttribute{
						Description:         "CredentialRef is the reference to the secret name that contains the credentials for the Nutanix Prism Central. The namespace for the secret is assumed to be a constant i.e. eksa-system.",
						MarkdownDescription: "CredentialRef is the reference to the secret name that contains the credentials for the Nutanix Prism Central. The namespace for the secret is assumed to be a constant i.e. eksa-system.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"endpoint": schema.StringAttribute{
						Description:         "Endpoint is the Endpoint of Nutanix Prism Central",
						MarkdownDescription: "Endpoint is the Endpoint of Nutanix Prism Central",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"failure_domains": schema.ListNestedAttribute{
						Description:         "FailureDomains is the optional list of failure domains for the Nutanix Datacenter.",
						MarkdownDescription: "FailureDomains is the optional list of failure domains for the Nutanix Datacenter.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cluster": schema.SingleNestedAttribute{
									Description:         "Cluster is the Prism Element cluster name or uuid that is connected to the Prism Central.",
									MarkdownDescription: "Cluster is the Prism Element cluster name or uuid that is connected to the Prism Central.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "name is the resource name in the PC",
											MarkdownDescription: "name is the resource name in the PC",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"type": schema.StringAttribute{
											Description:         "Type is the identifier type to use for this resource.",
											MarkdownDescription: "Type is the identifier type to use for this resource.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("uuid", "name"),
											},
										},

										"uuid": schema.StringAttribute{
											Description:         "uuid is the UUID of the resource in the PC.",
											MarkdownDescription: "uuid is the UUID of the resource in the PC.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is the unique name of the failure domain. Name must be between 1 and 64 characters long. It must consist of only lower case alphanumeric characters and hyphens (-). It must start and end with an alphanumeric character.",
									MarkdownDescription: "Name is the unique name of the failure domain. Name must be between 1 and 64 characters long. It must consist of only lower case alphanumeric characters and hyphens (-). It must start and end with an alphanumeric character.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(64),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
									},
								},

								"subnets": schema.ListNestedAttribute{
									Description:         "Subnets holds the list of subnets identifiers cluster's network subnets.",
									MarkdownDescription: "Subnets holds the list of subnets identifiers cluster's network subnets.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "name is the resource name in the PC",
												MarkdownDescription: "name is the resource name in the PC",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Type is the identifier type to use for this resource.",
												MarkdownDescription: "Type is the identifier type to use for this resource.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("uuid", "name"),
												},
											},

											"uuid": schema.StringAttribute{
												Description:         "uuid is the UUID of the resource in the PC.",
												MarkdownDescription: "uuid is the UUID of the resource in the PC.",
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

								"worker_machine_groups": schema.ListAttribute{
									Description:         "Worker Machine Groups holds the list of worker machine group names that will use this failure domain.",
									MarkdownDescription: "Worker Machine Groups holds the list of worker machine group names that will use this failure domain.",
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

					"insecure": schema.BoolAttribute{
						Description:         "Insecure is the optional flag to skip TLS verification. Nutanix Prism Central installation by default ships with a self-signed certificate that will fail TLS verification because the certificate is not issued by a public CA and does not have the IP SANs with the Prism Central endpoint. To accommodate the scenario where the user has not changed the default Certificate that ships with Prism Central, we allow the user to skip TLS verification. This is not recommended for production use.",
						MarkdownDescription: "Insecure is the optional flag to skip TLS verification. Nutanix Prism Central installation by default ships with a self-signed certificate that will fail TLS verification because the certificate is not issued by a public CA and does not have the IP SANs with the Prism Central endpoint. To accommodate the scenario where the user has not changed the default Certificate that ships with Prism Central, we allow the user to skip TLS verification. This is not recommended for production use.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"port": schema.Int64Attribute{
						Description:         "Port is the Port of Nutanix Prism Central",
						MarkdownDescription: "Port is the Port of Nutanix Prism Central",
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

func (r *AnywhereEksAmazonawsComNutanixDatacenterConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_anywhere_eks_amazonaws_com_nutanix_datacenter_config_v1alpha1_manifest")

	var model AnywhereEksAmazonawsComNutanixDatacenterConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("anywhere.eks.amazonaws.com/v1alpha1")
	model.Kind = pointer.String("NutanixDatacenterConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
