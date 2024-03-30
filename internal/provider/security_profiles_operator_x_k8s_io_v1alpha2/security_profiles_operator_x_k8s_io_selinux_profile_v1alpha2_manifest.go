/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package security_profiles_operator_x_k8s_io_v1alpha2

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
	_ datasource.DataSource = &SecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2Manifest{}
)

func NewSecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2Manifest() datasource.DataSource {
	return &SecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2Manifest{}
}

type SecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2Manifest struct{}

type SecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2ManifestData struct {
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
		Allow    *map[string]map[string][]string `tfsdk:"allow" json:"allow,omitempty"`
		Disabled *bool                           `tfsdk:"disabled" json:"disabled,omitempty"`
		Inherit  *[]struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"inherit" json:"inherit,omitempty"`
		Permissive *bool `tfsdk:"permissive" json:"permissive,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_security_profiles_operator_x_k8s_io_selinux_profile_v1alpha2_manifest"
}

func (r *SecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SelinuxProfile is the Schema for the selinuxprofiles API.",
		MarkdownDescription: "SelinuxProfile is the Schema for the selinuxprofiles API.",
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
				Description:         "SelinuxProfileSpec defines the desired state of SelinuxProfile.",
				MarkdownDescription: "SelinuxProfileSpec defines the desired state of SelinuxProfile.",
				Attributes: map[string]schema.Attribute{
					"allow": schema.MapAttribute{
						Description:         "Defines the allow policy for the profile",
						MarkdownDescription: "Defines the allow policy for the profile",
						ElementType:         types.MapType{ElemType: types.ListType{ElemType: types.StringType}},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disabled": schema.BoolAttribute{
						Description:         "Whether the profile is disabled and should be skipped during reconciliation.",
						MarkdownDescription: "Whether the profile is disabled and should be skipped during reconciliation.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"inherit": schema.ListNestedAttribute{
						Description:         "A SELinuxProfile or set of profiles that this inherits from.Note that they need to be in the same namespace.",
						MarkdownDescription: "A SELinuxProfile or set of profiles that this inherits from.Note that they need to be in the same namespace.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"kind": schema.StringAttribute{
									Description:         "The Kind of the policy that this inherits from.Can be a SelinuxProfile object Or 'System' if an alreadyinstalled policy will be used.The allowed 'System' policies are available in theSecurityProfilesOperatorDaemon instance.",
									MarkdownDescription: "The Kind of the policy that this inherits from.Can be a SelinuxProfile object Or 'System' if an alreadyinstalled policy will be used.The allowed 'System' policies are available in theSecurityProfilesOperatorDaemon instance.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("System", "SelinuxProfile"),
									},
								},

								"name": schema.StringAttribute{
									Description:         "The name of the policy that this inherits from.",
									MarkdownDescription: "The name of the policy that this inherits from.",
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

					"permissive": schema.BoolAttribute{
						Description:         "Permissive, when true will cause the SELinux profile to onlylog violations instead of enforcing them.",
						MarkdownDescription: "Permissive, when true will cause the SELinux profile to onlylog violations instead of enforcing them.",
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

func (r *SecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_security_profiles_operator_x_k8s_io_selinux_profile_v1alpha2_manifest")

	var model SecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("security-profiles-operator.x-k8s.io/v1alpha2")
	model.Kind = pointer.String("SelinuxProfile")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
