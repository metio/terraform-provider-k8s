/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type RbacAuthorizationK8SIoClusterRoleBindingV1Resource struct{}

var (
	_ resource.Resource = (*RbacAuthorizationK8SIoClusterRoleBindingV1Resource)(nil)
)

type RbacAuthorizationK8SIoClusterRoleBindingV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	RoleRef    types.Object `tfsdk:"role_ref"`
	Subjects   types.List   `tfsdk:"subjects"`
}

type RbacAuthorizationK8SIoClusterRoleBindingV1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	RoleRef *struct {
		ApiGroup *string `tfsdk:"api_group" yaml:"apiGroup,omitempty"`

		Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`
	} `tfsdk:"role_ref" yaml:"roleRef,omitempty"`

	Subjects *[]struct {
		ApiGroup *string `tfsdk:"api_group" yaml:"apiGroup,omitempty"`

		Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
	} `tfsdk:"subjects" yaml:"subjects,omitempty"`
}

func NewRbacAuthorizationK8SIoClusterRoleBindingV1Resource() resource.Resource {
	return &RbacAuthorizationK8SIoClusterRoleBindingV1Resource{}
}

func (r *RbacAuthorizationK8SIoClusterRoleBindingV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rbac_authorization_k8s_io_cluster_role_binding_v1"
}

func (r *RbacAuthorizationK8SIoClusterRoleBindingV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ClusterRoleBinding references a ClusterRole, but not contain it.  It can reference a ClusterRole in the global namespace, and adds who information via Subject.",
		MarkdownDescription: "ClusterRoleBinding references a ClusterRole, but not contain it.  It can reference a ClusterRole in the global namespace, and adds who information via Subject.",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"role_ref": {
				Description:         "RoleRef contains information that points to the role being used",
				MarkdownDescription: "RoleRef contains information that points to the role being used",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"api_group": {
						Description:         "APIGroup is the group for the resource being referenced",
						MarkdownDescription: "APIGroup is the group for the resource being referenced",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"kind": {
						Description:         "Kind is the type of resource being referenced",
						MarkdownDescription: "Kind is the type of resource being referenced",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"name": {
						Description:         "Name is the name of resource being referenced",
						MarkdownDescription: "Name is the name of resource being referenced",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},

			"subjects": {
				Description:         "Subjects holds references to the objects the role applies to.",
				MarkdownDescription: "Subjects holds references to the objects the role applies to.",

				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

					"api_group": {
						Description:         "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
						MarkdownDescription: "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"kind": {
						Description:         "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
						MarkdownDescription: "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"name": {
						Description:         "Name of the object being referenced.",
						MarkdownDescription: "Name of the object being referenced.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"namespace": {
						Description:         "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",
						MarkdownDescription: "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *RbacAuthorizationK8SIoClusterRoleBindingV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_rbac_authorization_k8s_io_cluster_role_binding_v1")

	var state RbacAuthorizationK8SIoClusterRoleBindingV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel RbacAuthorizationK8SIoClusterRoleBindingV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("rbac.authorization.k8s.io/v1")
	goModel.Kind = utilities.Ptr("ClusterRoleBinding")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *RbacAuthorizationK8SIoClusterRoleBindingV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_rbac_authorization_k8s_io_cluster_role_binding_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *RbacAuthorizationK8SIoClusterRoleBindingV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_rbac_authorization_k8s_io_cluster_role_binding_v1")

	var state RbacAuthorizationK8SIoClusterRoleBindingV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel RbacAuthorizationK8SIoClusterRoleBindingV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("rbac.authorization.k8s.io/v1")
	goModel.Kind = utilities.Ptr("ClusterRoleBinding")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *RbacAuthorizationK8SIoClusterRoleBindingV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_rbac_authorization_k8s_io_cluster_role_binding_v1")
	// NO-OP: Terraform removes the state automatically for us
}
