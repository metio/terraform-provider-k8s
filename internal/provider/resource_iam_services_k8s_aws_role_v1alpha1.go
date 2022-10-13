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

type IamServicesK8SAwsRoleV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*IamServicesK8SAwsRoleV1Alpha1Resource)(nil)
)

type IamServicesK8SAwsRoleV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type IamServicesK8SAwsRoleV1Alpha1GoModel struct {
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

	Spec *struct {
		AssumeRolePolicyDocument *string `tfsdk:"assume_role_policy_document" yaml:"assumeRolePolicyDocument,omitempty"`

		Description *string `tfsdk:"description" yaml:"description,omitempty"`

		MaxSessionDuration *int64 `tfsdk:"max_session_duration" yaml:"maxSessionDuration,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`

		Path *string `tfsdk:"path" yaml:"path,omitempty"`

		PermissionsBoundary *string `tfsdk:"permissions_boundary" yaml:"permissionsBoundary,omitempty"`

		Policies *[]string `tfsdk:"policies" yaml:"policies,omitempty"`

		Tags *[]struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"tags" yaml:"tags,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewIamServicesK8SAwsRoleV1Alpha1Resource() resource.Resource {
	return &IamServicesK8SAwsRoleV1Alpha1Resource{}
}

func (r *IamServicesK8SAwsRoleV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iam_services_k8s_aws_role_v1alpha1"
}

func (r *IamServicesK8SAwsRoleV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Role is the Schema for the Roles API",
		MarkdownDescription: "Role is the Schema for the Roles API",
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

			"spec": {
				Description:         "RoleSpec defines the desired state of Role.  Contains information about an IAM role. This structure is returned as a response element in several API operations that interact with roles.",
				MarkdownDescription: "RoleSpec defines the desired state of Role.  Contains information about an IAM role. This structure is returned as a response element in several API operations that interact with roles.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"assume_role_policy_document": {
						Description:         "The trust relationship policy document that grants an entity permission to assume the role.  In IAM, you must provide a JSON policy that has been converted to a string. However, for CloudFormation templates formatted in YAML, you can provide the policy in JSON or YAML format. CloudFormation always converts a YAML policy to JSON format before submitting it to IAM.  The regex pattern (http://wikipedia.org/wiki/regex) used to validate this parameter is a string of characters consisting of the following:  * Any printable ASCII character ranging from the space character (u0020) through the end of the ASCII character range  * The printable characters in the Basic Latin and Latin-1 Supplement character set (through u00FF)  * The special characters tab (u0009), line feed (u000A), and carriage return (u000D)  Upon success, the response includes the same trust policy in JSON format.",
						MarkdownDescription: "The trust relationship policy document that grants an entity permission to assume the role.  In IAM, you must provide a JSON policy that has been converted to a string. However, for CloudFormation templates formatted in YAML, you can provide the policy in JSON or YAML format. CloudFormation always converts a YAML policy to JSON format before submitting it to IAM.  The regex pattern (http://wikipedia.org/wiki/regex) used to validate this parameter is a string of characters consisting of the following:  * Any printable ASCII character ranging from the space character (u0020) through the end of the ASCII character range  * The printable characters in the Basic Latin and Latin-1 Supplement character set (through u00FF)  * The special characters tab (u0009), line feed (u000A), and carriage return (u000D)  Upon success, the response includes the same trust policy in JSON format.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"description": {
						Description:         "A description of the role.",
						MarkdownDescription: "A description of the role.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"max_session_duration": {
						Description:         "The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default value of one hour is applied. This setting can have a value from 1 hour to 12 hours.  Anyone who assumes the role from the or API can use the DurationSeconds API parameter or the duration-seconds CLI parameter to request a longer session. The MaxSessionDuration setting determines the maximum duration that can be requested using the DurationSeconds parameter. If users don't specify a value for the DurationSeconds parameter, their security credentials are valid for one hour by default. This applies when you use the AssumeRole* API operations or the assume-role* CLI operations but does not apply when you use those operations to create a console URL. For more information, see Using IAM roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html) in the IAM User Guide.",
						MarkdownDescription: "The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default value of one hour is applied. This setting can have a value from 1 hour to 12 hours.  Anyone who assumes the role from the or API can use the DurationSeconds API parameter or the duration-seconds CLI parameter to request a longer session. The MaxSessionDuration setting determines the maximum duration that can be requested using the DurationSeconds parameter. If users don't specify a value for the DurationSeconds parameter, their security credentials are valid for one hour by default. This applies when you use the AssumeRole* API operations or the assume-role* CLI operations but does not apply when you use those operations to create a console URL. For more information, see Using IAM roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html) in the IAM User Guide.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"name": {
						Description:         "The name of the role to create.  IAM user, group, role, and policy names must be unique within the account. Names are not distinguished by case. For example, you cannot create resources named both 'MyResource' and 'myresource'.",
						MarkdownDescription: "The name of the role to create.  IAM user, group, role, and policy names must be unique within the account. Names are not distinguished by case. For example, you cannot create resources named both 'MyResource' and 'myresource'.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"path": {
						Description:         "The path to the role. For more information about paths, see IAM Identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide.  This parameter is optional. If it is not included, it defaults to a slash (/).  This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. In addition, it can contain any ASCII character from the ! (u0021) through the DEL character (u007F), including most punctuation characters, digits, and upper and lowercased letters.",
						MarkdownDescription: "The path to the role. For more information about paths, see IAM Identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide.  This parameter is optional. If it is not included, it defaults to a slash (/).  This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. In addition, it can contain any ASCII character from the ! (u0021) through the DEL character (u007F), including most punctuation characters, digits, and upper and lowercased letters.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"permissions_boundary": {
						Description:         "The ARN of the policy that is used to set the permissions boundary for the role.",
						MarkdownDescription: "The ARN of the policy that is used to set the permissions boundary for the role.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"policies": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tags": {
						Description:         "A list of tags that you want to attach to the new role. Each tag consists of a key name and an associated value. For more information about tagging, see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the IAM User Guide.  If any one of the tags is invalid or if you exceed the allowed maximum number of tags, then the entire request fails and the resource is not created.",
						MarkdownDescription: "A list of tags that you want to attach to the new role. Each tag consists of a key name and an associated value. For more information about tagging, see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the IAM User Guide.  If any one of the tags is invalid or if you exceed the allowed maximum number of tags, then the entire request fails and the resource is not created.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"key": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"value": {
								Description:         "",
								MarkdownDescription: "",

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
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *IamServicesK8SAwsRoleV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_iam_services_k8s_aws_role_v1alpha1")

	var state IamServicesK8SAwsRoleV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel IamServicesK8SAwsRoleV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("iam.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("Role")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *IamServicesK8SAwsRoleV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_iam_services_k8s_aws_role_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *IamServicesK8SAwsRoleV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_iam_services_k8s_aws_role_v1alpha1")

	var state IamServicesK8SAwsRoleV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel IamServicesK8SAwsRoleV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("iam.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("Role")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *IamServicesK8SAwsRoleV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_iam_services_k8s_aws_role_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
