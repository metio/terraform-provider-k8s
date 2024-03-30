/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package azure_microsoft_com_v1alpha2

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
	_ datasource.DataSource = &AzureMicrosoftComMySqlaaduserV1Alpha2Manifest{}
)

func NewAzureMicrosoftComMySqlaaduserV1Alpha2Manifest() datasource.DataSource {
	return &AzureMicrosoftComMySqlaaduserV1Alpha2Manifest{}
}

type AzureMicrosoftComMySqlaaduserV1Alpha2Manifest struct{}

type AzureMicrosoftComMySqlaaduserV1Alpha2ManifestData struct {
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
		AadId         *string              `tfsdk:"aad_id" json:"aadId,omitempty"`
		DatabaseRoles *map[string][]string `tfsdk:"database_roles" json:"databaseRoles,omitempty"`
		ResourceGroup *string              `tfsdk:"resource_group" json:"resourceGroup,omitempty"`
		Roles         *[]string            `tfsdk:"roles" json:"roles,omitempty"`
		Server        *string              `tfsdk:"server" json:"server,omitempty"`
		Username      *string              `tfsdk:"username" json:"username,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AzureMicrosoftComMySqlaaduserV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_azure_microsoft_com_my_sqlaad_user_v1alpha2_manifest"
}

func (r *AzureMicrosoftComMySqlaaduserV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "MySQLAADUser is the Schema for an AAD user for MySQL",
		MarkdownDescription: "MySQLAADUser is the Schema for an AAD user for MySQL",
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
				Description:         "MySQLAADUserSpec defines the desired state of MySQLAADUser",
				MarkdownDescription: "MySQLAADUserSpec defines the desired state of MySQLAADUser",
				Attributes: map[string]schema.Attribute{
					"aad_id": schema.StringAttribute{
						Description:         "AAD ID is the ID of the user in Azure Active Directory. When creating a user for a managed identity this must be the client id (sometimes called app id) of the managed identity. When creating a user for a 'normal' (non-managed identity) user or group, this is the OID of the user or group.",
						MarkdownDescription: "AAD ID is the ID of the user in Azure Active Directory. When creating a user for a managed identity this must be the client id (sometimes called app id) of the managed identity. When creating a user for a 'normal' (non-managed identity) user or group, this is the OID of the user or group.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"database_roles": schema.MapAttribute{
						Description:         "The database-level roles assigned to the user (keyed by database name).",
						MarkdownDescription: "The database-level roles assigned to the user (keyed by database name).",
						ElementType:         types.ListType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
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

					"roles": schema.ListAttribute{
						Description:         "The server-level roles assigned to the user.",
						MarkdownDescription: "The server-level roles assigned to the user.",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"server": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"username": schema.StringAttribute{
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
		},
	}
}

func (r *AzureMicrosoftComMySqlaaduserV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_azure_microsoft_com_my_sqlaad_user_v1alpha2_manifest")

	var model AzureMicrosoftComMySqlaaduserV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("azure.microsoft.com/v1alpha2")
	model.Kind = pointer.String("MySQLAADUser")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
