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

type ElasticacheServicesK8SAwsUserV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ElasticacheServicesK8SAwsUserV1Alpha1Resource)(nil)
)

type ElasticacheServicesK8SAwsUserV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ElasticacheServicesK8SAwsUserV1Alpha1GoModel struct {
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
		AccessString *string `tfsdk:"access_string" yaml:"accessString,omitempty"`

		Engine *string `tfsdk:"engine" yaml:"engine,omitempty"`

		NoPasswordRequired *bool `tfsdk:"no_password_required" yaml:"noPasswordRequired,omitempty"`

		Passwords *[]struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
		} `tfsdk:"passwords" yaml:"passwords,omitempty"`

		Tags *[]struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"tags" yaml:"tags,omitempty"`

		UserID *string `tfsdk:"user_id" yaml:"userID,omitempty"`

		UserName *string `tfsdk:"user_name" yaml:"userName,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewElasticacheServicesK8SAwsUserV1Alpha1Resource() resource.Resource {
	return &ElasticacheServicesK8SAwsUserV1Alpha1Resource{}
}

func (r *ElasticacheServicesK8SAwsUserV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_elasticache_services_k8s_aws_user_v1alpha1"
}

func (r *ElasticacheServicesK8SAwsUserV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "User is the Schema for the Users API",
		MarkdownDescription: "User is the Schema for the Users API",
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
				Description:         "",
				MarkdownDescription: "",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"access_string": {
						Description:         "Access permissions string used for this user.",
						MarkdownDescription: "Access permissions string used for this user.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"engine": {
						Description:         "The current supported value is Redis.",
						MarkdownDescription: "The current supported value is Redis.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"no_password_required": {
						Description:         "Indicates a password is not required for this user.",
						MarkdownDescription: "Indicates a password is not required for this user.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"passwords": {
						Description:         "Passwords used for this user. You can create up to two passwords for each user.",
						MarkdownDescription: "Passwords used for this user. You can create up to two passwords for each user.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"key": {
								Description:         "Key is the key within the secret",
								MarkdownDescription: "Key is the key within the secret",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"name": {
								Description:         "Name is unique within a namespace to reference a secret resource.",
								MarkdownDescription: "Name is unique within a namespace to reference a secret resource.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace defines the space within which the secret name must be unique.",
								MarkdownDescription: "Namespace defines the space within which the secret name must be unique.",

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

					"tags": {
						Description:         "A list of tags to be added to this resource. A tag is a key-value pair. A tag key must be accompanied by a tag value, although null is accepted.",
						MarkdownDescription: "A list of tags to be added to this resource. A tag is a key-value pair. A tag key must be accompanied by a tag value, although null is accepted.",

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

					"user_id": {
						Description:         "The ID of the user.",
						MarkdownDescription: "The ID of the user.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"user_name": {
						Description:         "The username of the user.",
						MarkdownDescription: "The username of the user.",

						Type: types.StringType,

						Required: true,
						Optional: false,
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

func (r *ElasticacheServicesK8SAwsUserV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_elasticache_services_k8s_aws_user_v1alpha1")

	var state ElasticacheServicesK8SAwsUserV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ElasticacheServicesK8SAwsUserV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("elasticache.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("User")

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

func (r *ElasticacheServicesK8SAwsUserV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_elasticache_services_k8s_aws_user_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ElasticacheServicesK8SAwsUserV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_elasticache_services_k8s_aws_user_v1alpha1")

	var state ElasticacheServicesK8SAwsUserV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ElasticacheServicesK8SAwsUserV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("elasticache.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("User")

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

func (r *ElasticacheServicesK8SAwsUserV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_elasticache_services_k8s_aws_user_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
