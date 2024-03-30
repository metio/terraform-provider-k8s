/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package azure_microsoft_com_v1alpha1

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
	_ datasource.DataSource = &AzureMicrosoftComApimgmtApiV1Alpha1Manifest{}
)

func NewAzureMicrosoftComApimgmtApiV1Alpha1Manifest() datasource.DataSource {
	return &AzureMicrosoftComApimgmtApiV1Alpha1Manifest{}
}

type AzureMicrosoftComApimgmtApiV1Alpha1Manifest struct{}

type AzureMicrosoftComApimgmtApiV1Alpha1ManifestData struct {
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
		ApiId      *string `tfsdk:"api_id" json:"apiId,omitempty"`
		ApiService *string `tfsdk:"api_service" json:"apiService,omitempty"`
		Location   *string `tfsdk:"location" json:"location,omitempty"`
		Properties *struct {
			ApiRevision            *string `tfsdk:"api_revision" json:"apiRevision,omitempty"`
			ApiRevisionDescription *string `tfsdk:"api_revision_description" json:"apiRevisionDescription,omitempty"`
			ApiVersion             *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			ApiVersionDescription  *string `tfsdk:"api_version_description" json:"apiVersionDescription,omitempty"`
			ApiVersionSetId        *string `tfsdk:"api_version_set_id" json:"apiVersionSetId,omitempty"`
			ApiVersionSets         *struct {
				Description *string `tfsdk:"description" json:"description,omitempty"`
				Id          *string `tfsdk:"id" json:"id,omitempty"`
				Name        *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"api_version_sets" json:"apiVersionSets,omitempty"`
			Description          *string   `tfsdk:"description" json:"description,omitempty"`
			DisplayName          *string   `tfsdk:"display_name" json:"displayName,omitempty"`
			Format               *string   `tfsdk:"format" json:"format,omitempty"`
			IsCurrent            *bool     `tfsdk:"is_current" json:"isCurrent,omitempty"`
			IsOnline             *bool     `tfsdk:"is_online" json:"isOnline,omitempty"`
			Path                 *string   `tfsdk:"path" json:"path,omitempty"`
			Protocols            *[]string `tfsdk:"protocols" json:"protocols,omitempty"`
			ServiceUrl           *string   `tfsdk:"service_url" json:"serviceUrl,omitempty"`
			SourceApiId          *string   `tfsdk:"source_api_id" json:"sourceApiId,omitempty"`
			SubscriptionRequired *bool     `tfsdk:"subscription_required" json:"subscriptionRequired,omitempty"`
		} `tfsdk:"properties" json:"properties,omitempty"`
		ResourceGroup *string `tfsdk:"resource_group" json:"resourceGroup,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AzureMicrosoftComApimgmtApiV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_azure_microsoft_com_api_mgmt_api_v1alpha1_manifest"
}

func (r *AzureMicrosoftComApimgmtApiV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "APIMgmtSpec defines the desired state of APIMgmt",
				MarkdownDescription: "APIMgmtSpec defines the desired state of APIMgmt",
				Attributes: map[string]schema.Attribute{
					"api_id": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"api_service": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"location": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"properties": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"api_revision": schema.StringAttribute{
								Description:         "APIRevision - Describes the Revision of the Api. If no value is provided, default revision 1 is created",
								MarkdownDescription: "APIRevision - Describes the Revision of the Api. If no value is provided, default revision 1 is created",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"api_revision_description": schema.StringAttribute{
								Description:         "APIRevisionDescription - Description of the Api Revision.",
								MarkdownDescription: "APIRevisionDescription - Description of the Api Revision.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"api_version": schema.StringAttribute{
								Description:         "APIVersion - Indicates the Version identifier of the API if the API is versioned",
								MarkdownDescription: "APIVersion - Indicates the Version identifier of the API if the API is versioned",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"api_version_description": schema.StringAttribute{
								Description:         "APIVersionDescription - Description of the Api Version.",
								MarkdownDescription: "APIVersionDescription - Description of the Api Version.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"api_version_set_id": schema.StringAttribute{
								Description:         "APIVersionSetID - A resource identifier for the related ApiVersionSet.",
								MarkdownDescription: "APIVersionSetID - A resource identifier for the related ApiVersionSet.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"api_version_sets": schema.SingleNestedAttribute{
								Description:         "APIVersionSet - APIVersionSetContractDetails an API Version Set contains the common configuration for a set of API versions.",
								MarkdownDescription: "APIVersionSet - APIVersionSetContractDetails an API Version Set contains the common configuration for a set of API versions.",
								Attributes: map[string]schema.Attribute{
									"description": schema.StringAttribute{
										Description:         "Description - Description of API Version Set.",
										MarkdownDescription: "Description - Description of API Version Set.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"id": schema.StringAttribute{
										Description:         "ID - Identifier for existing API Version Set. Omit this value to create a new Version Set.",
										MarkdownDescription: "ID - Identifier for existing API Version Set. Omit this value to create a new Version Set.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name - The display Name of the API Version Set.",
										MarkdownDescription: "Name - The display Name of the API Version Set.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"description": schema.StringAttribute{
								Description:         "Description - Description of the API. May include HTML formatting tags.",
								MarkdownDescription: "Description - Description of the API. May include HTML formatting tags.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"display_name": schema.StringAttribute{
								Description:         "DisplayName - API name. Must be 1 to 300 characters long.",
								MarkdownDescription: "DisplayName - API name. Must be 1 to 300 characters long.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"format": schema.StringAttribute{
								Description:         "Format - Format of the Content in which the API is getting imported. Possible values include: 'WadlXML', 'WadlLinkJSON', 'SwaggerJSON', 'SwaggerLinkJSON', 'Wsdl', 'WsdlLink', 'Openapi', 'Openapijson', 'OpenapiLink'",
								MarkdownDescription: "Format - Format of the Content in which the API is getting imported. Possible values include: 'WadlXML', 'WadlLinkJSON', 'SwaggerJSON', 'SwaggerLinkJSON', 'Wsdl', 'WsdlLink', 'Openapi', 'Openapijson', 'OpenapiLink'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"is_current": schema.BoolAttribute{
								Description:         "IsCurrent - Indicates if API revision is current api revision.",
								MarkdownDescription: "IsCurrent - Indicates if API revision is current api revision.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"is_online": schema.BoolAttribute{
								Description:         "IsOnline - READ-ONLY; Indicates if API revision is accessible via the gateway.",
								MarkdownDescription: "IsOnline - READ-ONLY; Indicates if API revision is accessible via the gateway.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"path": schema.StringAttribute{
								Description:         "Path - Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.",
								MarkdownDescription: "Path - Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"protocols": schema.ListAttribute{
								Description:         "Protocols - Describes on which protocols the operations in this API can be invoked.",
								MarkdownDescription: "Protocols - Describes on which protocols the operations in this API can be invoked.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_url": schema.StringAttribute{
								Description:         "ServiceURL - Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.",
								MarkdownDescription: "ServiceURL - Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"source_api_id": schema.StringAttribute{
								Description:         "SourceAPIID - API identifier of the source API.",
								MarkdownDescription: "SourceAPIID - API identifier of the source API.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"subscription_required": schema.BoolAttribute{
								Description:         "SubscriptionRequired - Specifies whether an API or Product subscription is required for accessing the API.",
								MarkdownDescription: "SubscriptionRequired - Specifies whether an API or Product subscription is required for accessing the API.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"resource_group": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[-\w\._\(\)]+$`), ""),
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

func (r *AzureMicrosoftComApimgmtApiV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_azure_microsoft_com_api_mgmt_api_v1alpha1_manifest")

	var model AzureMicrosoftComApimgmtApiV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("azure.microsoft.com/v1alpha1")
	model.Kind = pointer.String("APIMgmtAPI")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
