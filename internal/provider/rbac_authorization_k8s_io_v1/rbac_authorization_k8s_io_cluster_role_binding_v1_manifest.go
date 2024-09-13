/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package rbac_authorization_k8s_io_v1

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
	_ datasource.DataSource = &RbacAuthorizationK8SIoClusterRoleBindingV1Manifest{}
)

func NewRbacAuthorizationK8SIoClusterRoleBindingV1Manifest() datasource.DataSource {
	return &RbacAuthorizationK8SIoClusterRoleBindingV1Manifest{}
}

type RbacAuthorizationK8SIoClusterRoleBindingV1Manifest struct{}

type RbacAuthorizationK8SIoClusterRoleBindingV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	RoleRef *struct {
		ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
		Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
		Name     *string `tfsdk:"name" json:"name,omitempty"`
	} `tfsdk:"role_ref" json:"roleRef,omitempty"`
	Subjects *[]struct {
		ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
		Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
		Name      *string `tfsdk:"name" json:"name,omitempty"`
		Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
	} `tfsdk:"subjects" json:"subjects,omitempty"`
}

func (r *RbacAuthorizationK8SIoClusterRoleBindingV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_rbac_authorization_k8s_io_cluster_role_binding_v1_manifest"
}

func (r *RbacAuthorizationK8SIoClusterRoleBindingV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterRoleBinding references a ClusterRole, but not contain it. It can reference a ClusterRole in the global namespace, and adds who information via Subject.",
		MarkdownDescription: "ClusterRoleBinding references a ClusterRole, but not contain it. It can reference a ClusterRole in the global namespace, and adds who information via Subject.",
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

			"role_ref": schema.SingleNestedAttribute{
				Description:         "RoleRef contains information that points to the role being used",
				MarkdownDescription: "RoleRef contains information that points to the role being used",
				Attributes: map[string]schema.Attribute{
					"api_group": schema.StringAttribute{
						Description:         "APIGroup is the group for the resource being referenced",
						MarkdownDescription: "APIGroup is the group for the resource being referenced",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"kind": schema.StringAttribute{
						Description:         "Kind is the type of resource being referenced",
						MarkdownDescription: "Kind is the type of resource being referenced",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "Name is the name of resource being referenced",
						MarkdownDescription: "Name is the name of resource being referenced",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},

			"subjects": schema.ListNestedAttribute{
				Description:         "Subjects holds references to the objects the role applies to.",
				MarkdownDescription: "Subjects holds references to the objects the role applies to.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"api_group": schema.StringAttribute{
							Description:         "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
							MarkdownDescription: "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
							Required:            false,
							Optional:            true,
							Computed:            false,
						},

						"kind": schema.StringAttribute{
							Description:         "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
							MarkdownDescription: "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},

						"name": schema.StringAttribute{
							Description:         "Name of the object being referenced.",
							MarkdownDescription: "Name of the object being referenced.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},

						"namespace": schema.StringAttribute{
							Description:         "Namespace of the referenced object. If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",
							MarkdownDescription: "Namespace of the referenced object. If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",
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
	}
}

func (r *RbacAuthorizationK8SIoClusterRoleBindingV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_rbac_authorization_k8s_io_cluster_role_binding_v1_manifest")

	var model RbacAuthorizationK8SIoClusterRoleBindingV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("rbac.authorization.k8s.io/v1")
	model.Kind = pointer.String("ClusterRoleBinding")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
