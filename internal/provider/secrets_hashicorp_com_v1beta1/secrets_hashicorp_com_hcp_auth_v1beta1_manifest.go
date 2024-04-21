/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package secrets_hashicorp_com_v1beta1

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
	_ datasource.DataSource = &SecretsHashicorpComHcpauthV1Beta1Manifest{}
)

func NewSecretsHashicorpComHcpauthV1Beta1Manifest() datasource.DataSource {
	return &SecretsHashicorpComHcpauthV1Beta1Manifest{}
}

type SecretsHashicorpComHcpauthV1Beta1Manifest struct{}

type SecretsHashicorpComHcpauthV1Beta1ManifestData struct {
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
		AllowedNamespaces *[]string `tfsdk:"allowed_namespaces" json:"allowedNamespaces,omitempty"`
		Method            *string   `tfsdk:"method" json:"method,omitempty"`
		OrganizationID    *string   `tfsdk:"organization_id" json:"organizationID,omitempty"`
		ProjectID         *string   `tfsdk:"project_id" json:"projectID,omitempty"`
		ServicePrincipal  *struct {
			SecretRef *string `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"service_principal" json:"servicePrincipal,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SecretsHashicorpComHcpauthV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_secrets_hashicorp_com_hcp_auth_v1beta1_manifest"
}

func (r *SecretsHashicorpComHcpauthV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "HCPAuth is the Schema for the hcpauths API",
		MarkdownDescription: "HCPAuth is the Schema for the hcpauths API",
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
				Description:         "HCPAuthSpec defines the desired state of HCPAuth",
				MarkdownDescription: "HCPAuthSpec defines the desired state of HCPAuth",
				Attributes: map[string]schema.Attribute{
					"allowed_namespaces": schema.ListAttribute{
						Description:         "AllowedNamespaces Kubernetes Namespaces which are allow-listed for use with this AuthMethod.This field allows administrators to customize which Kubernetes namespaces are authorized touse with this AuthMethod. While Vault will still enforce its own rules, this has the addedconfigurability of restricting which HCPAuthMethods can be used by which namespaces.Accepted values:[]{'*'} - wildcard, all namespaces.[]{'a', 'b'} - list of namespaces.unset - disallow all namespaces except the Operator's the HCPAuthMethod's namespace, thisis the default behavior.",
						MarkdownDescription: "AllowedNamespaces Kubernetes Namespaces which are allow-listed for use with this AuthMethod.This field allows administrators to customize which Kubernetes namespaces are authorized touse with this AuthMethod. While Vault will still enforce its own rules, this has the addedconfigurability of restricting which HCPAuthMethods can be used by which namespaces.Accepted values:[]{'*'} - wildcard, all namespaces.[]{'a', 'b'} - list of namespaces.unset - disallow all namespaces except the Operator's the HCPAuthMethod's namespace, thisis the default behavior.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"method": schema.StringAttribute{
						Description:         "Method to use when authenticating to Vault.",
						MarkdownDescription: "Method to use when authenticating to Vault.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("servicePrincipal"),
						},
					},

					"organization_id": schema.StringAttribute{
						Description:         "OrganizationID of the HCP organization.",
						MarkdownDescription: "OrganizationID of the HCP organization.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"project_id": schema.StringAttribute{
						Description:         "ProjectID of the HCP project.",
						MarkdownDescription: "ProjectID of the HCP project.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"service_principal": schema.SingleNestedAttribute{
						Description:         "ServicePrincipal provides the necessary configuration for authenticating toHCP using a service principal. For security reasons, only project-levelservice principals should ever be used.",
						MarkdownDescription: "ServicePrincipal provides the necessary configuration for authenticating toHCP using a service principal. For security reasons, only project-levelservice principals should ever be used.",
						Attributes: map[string]schema.Attribute{
							"secret_ref": schema.StringAttribute{
								Description:         "SecretRef is the name of a Kubernetes secret in the consumer's(VDS/VSS/PKI/HCP) namespace which provides the HCP ServicePrincipal clientID,and clientSecret.The secret data must have the following structure {  'clientID': 'clientID',  'clientSecret': 'clientSecret',}",
								MarkdownDescription: "SecretRef is the name of a Kubernetes secret in the consumer's(VDS/VSS/PKI/HCP) namespace which provides the HCP ServicePrincipal clientID,and clientSecret.The secret data must have the following structure {  'clientID': 'clientID',  'clientSecret': 'clientSecret',}",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *SecretsHashicorpComHcpauthV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_secrets_hashicorp_com_hcp_auth_v1beta1_manifest")

	var model SecretsHashicorpComHcpauthV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("secrets.hashicorp.com/v1beta1")
	model.Kind = pointer.String("HCPAuth")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
