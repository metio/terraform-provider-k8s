/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package edp_epam_com_v1alpha1

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
	_ datasource.DataSource = &EdpEpamComNexusUserV1Alpha1Manifest{}
)

func NewEdpEpamComNexusUserV1Alpha1Manifest() datasource.DataSource {
	return &EdpEpamComNexusUserV1Alpha1Manifest{}
}

type EdpEpamComNexusUserV1Alpha1Manifest struct{}

type EdpEpamComNexusUserV1Alpha1ManifestData struct {
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
		Email     *string `tfsdk:"email" json:"email,omitempty"`
		FirstName *string `tfsdk:"first_name" json:"firstName,omitempty"`
		Id        *string `tfsdk:"id" json:"id,omitempty"`
		LastName  *string `tfsdk:"last_name" json:"lastName,omitempty"`
		NexusRef  *struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"nexus_ref" json:"nexusRef,omitempty"`
		Roles  *[]string `tfsdk:"roles" json:"roles,omitempty"`
		Secret *string   `tfsdk:"secret" json:"secret,omitempty"`
		Status *string   `tfsdk:"status" json:"status,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *EdpEpamComNexusUserV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_edp_epam_com_nexus_user_v1alpha1_manifest"
}

func (r *EdpEpamComNexusUserV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "NexusUser is the Schema for the nexususers API.",
		MarkdownDescription: "NexusUser is the Schema for the nexususers API.",
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
				Description:         "NexusUserSpec defines the desired state of NexusUser.",
				MarkdownDescription: "NexusUserSpec defines the desired state of NexusUser.",
				Attributes: map[string]schema.Attribute{
					"email": schema.StringAttribute{
						Description:         "Email is the email address of the user.",
						MarkdownDescription: "Email is the email address of the user.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(254),
						},
					},

					"first_name": schema.StringAttribute{
						Description:         "FirstName of the user.",
						MarkdownDescription: "FirstName of the user.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"id": schema.StringAttribute{
						Description:         "ID is the username of the user. ID should be unique across all users.",
						MarkdownDescription: "ID is the username of the user. ID should be unique across all users.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},

					"last_name": schema.StringAttribute{
						Description:         "LastName of the user.",
						MarkdownDescription: "LastName of the user.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"nexus_ref": schema.SingleNestedAttribute{
						Description:         "NexusRef is a reference to Nexus custom resource.",
						MarkdownDescription: "NexusRef is a reference to Nexus custom resource.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind specifies the kind of the Nexus resource.",
								MarkdownDescription: "Kind specifies the kind of the Nexus resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name specifies the name of the Nexus resource.",
								MarkdownDescription: "Name specifies the name of the Nexus resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"roles": schema.ListAttribute{
						Description:         "Roles is a list of roles assigned to user.",
						MarkdownDescription: "Roles is a list of roles assigned to user.",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"secret": schema.StringAttribute{
						Description:         "Secret is the reference of the k8s object Secret for the user password. Format: $secret-name:secret-key. After updating Secret user password will be updated.",
						MarkdownDescription: "Secret is the reference of the k8s object Secret for the user password. Format: $secret-name:secret-key. After updating Secret user password will be updated.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"status": schema.StringAttribute{
						Description:         "Status is a status of the user.",
						MarkdownDescription: "Status is a status of the user.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("active", "disabled"),
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

func (r *EdpEpamComNexusUserV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_edp_epam_com_nexus_user_v1alpha1_manifest")

	var model EdpEpamComNexusUserV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("edp.epam.com/v1alpha1")
	model.Kind = pointer.String("NexusUser")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
