/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package organizations_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &OrganizationsServicesK8SAwsOrganizationalUnitV1Alpha1Manifest{}
)

func NewOrganizationsServicesK8SAwsOrganizationalUnitV1Alpha1Manifest() datasource.DataSource {
	return &OrganizationsServicesK8SAwsOrganizationalUnitV1Alpha1Manifest{}
}

type OrganizationsServicesK8SAwsOrganizationalUnitV1Alpha1Manifest struct{}

type OrganizationsServicesK8SAwsOrganizationalUnitV1Alpha1ManifestData struct {
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
		Name     *string `tfsdk:"name" json:"name,omitempty"`
		ParentID *string `tfsdk:"parent_id" json:"parentID,omitempty"`
		Tags     *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OrganizationsServicesK8SAwsOrganizationalUnitV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_organizations_services_k8s_aws_organizational_unit_v1alpha1_manifest"
}

func (r *OrganizationsServicesK8SAwsOrganizationalUnitV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "OrganizationalUnit is the Schema for the OrganizationalUnits API",
		MarkdownDescription: "OrganizationalUnit is the Schema for the OrganizationalUnits API",
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
				Description:         "OrganizationalUnitSpec defines the desired state of OrganizationalUnit.Contains details about an organizational unit (OU). An OU is a containerof Amazon Web Services accounts within a root of an organization. Policiesthat are attached to an OU apply to all accounts contained in that OU andin any child OUs.",
				MarkdownDescription: "OrganizationalUnitSpec defines the desired state of OrganizationalUnit.Contains details about an organizational unit (OU). An OU is a containerof Amazon Web Services accounts within a root of an organization. Policiesthat are attached to an OU apply to all accounts contained in that OU andin any child OUs.",
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description:         "The friendly name to assign to the new OU.",
						MarkdownDescription: "The friendly name to assign to the new OU.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"parent_id": schema.StringAttribute{
						Description:         "The unique identifier (ID) of the parent root or OU that you want to createthe new OU in.The regex pattern (http://wikipedia.org/wiki/regex) for a parent ID stringrequires one of the following:   * Root - A string that begins with 'r-' followed by from 4 to 32 lowercase   letters or digits.   * Organizational unit (OU) - A string that begins with 'ou-' followed   by from 4 to 32 lowercase letters or digits (the ID of the root that the   OU is in). This string is followed by a second '-' dash and from 8 to   32 additional lowercase letters or digits.",
						MarkdownDescription: "The unique identifier (ID) of the parent root or OU that you want to createthe new OU in.The regex pattern (http://wikipedia.org/wiki/regex) for a parent ID stringrequires one of the following:   * Root - A string that begins with 'r-' followed by from 4 to 32 lowercase   letters or digits.   * Organizational unit (OU) - A string that begins with 'ou-' followed   by from 4 to 32 lowercase letters or digits (the ID of the root that the   OU is in). This string is followed by a second '-' dash and from 8 to   32 additional lowercase letters or digits.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "A list of tags that you want to attach to the newly created OU. For eachtag in the list, you must specify both a tag key and a value. You can setthe value to an empty string, but you can't set it to null. For more informationabout tagging, see Tagging Organizations resources (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_tagging.html)in the Organizations User Guide.If any one of the tags is invalid or if you exceed the allowed number oftags for an OU, then the entire request fails and the OU is not created.",
						MarkdownDescription: "A list of tags that you want to attach to the newly created OU. For eachtag in the list, you must specify both a tag key and a value. You can setthe value to an empty string, but you can't set it to null. For more informationabout tagging, see Tagging Organizations resources (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_tagging.html)in the Organizations User Guide.If any one of the tags is invalid or if you exceed the allowed number oftags for an OU, then the entire request fails and the OU is not created.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *OrganizationsServicesK8SAwsOrganizationalUnitV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_organizations_services_k8s_aws_organizational_unit_v1alpha1_manifest")

	var model OrganizationsServicesK8SAwsOrganizationalUnitV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("organizations.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("OrganizationalUnit")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
