/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

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

type ServicesK8SAwsFieldExportV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ServicesK8SAwsFieldExportV1Alpha1Resource)(nil)
)

type ServicesK8SAwsFieldExportV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ServicesK8SAwsFieldExportV1Alpha1GoModel struct {
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
		From *struct {
			Path *string `tfsdk:"path" yaml:"path,omitempty"`

			Resource *struct {
				Group *string `tfsdk:"group" yaml:"group,omitempty"`

				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"resource" yaml:"resource,omitempty"`
		} `tfsdk:"from" yaml:"from,omitempty"`

		To *struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
		} `tfsdk:"to" yaml:"to,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewServicesK8SAwsFieldExportV1Alpha1Resource() resource.Resource {
	return &ServicesK8SAwsFieldExportV1Alpha1Resource{}
}

func (r *ServicesK8SAwsFieldExportV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_services_k8s_aws_field_export_v1alpha1"
}

func (r *ServicesK8SAwsFieldExportV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "FieldExport is the schema for the FieldExport API.",
		MarkdownDescription: "FieldExport is the schema for the FieldExport API.",
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
				Description:         "FieldExportSpec defines the desired state of the FieldExport.",
				MarkdownDescription: "FieldExportSpec defines the desired state of the FieldExport.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"from": {
						Description:         "ResourceFieldSelector provides the values necessary to identify an individual field on an individual K8s resource.",
						MarkdownDescription: "ResourceFieldSelector provides the values necessary to identify an individual field on an individual K8s resource.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"path": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"resource": {
								Description:         "NamespacedResource provides all the values necessary to identify an ACK resource of a given type (within the same namespace as the custom resource containing this type).",
								MarkdownDescription: "NamespacedResource provides all the values necessary to identify an ACK resource of a given type (within the same namespace as the custom resource containing this type).",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"group": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"kind": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": {
										Description:         "",
										MarkdownDescription: "",

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
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"to": {
						Description:         "FieldExportTarget provides the values necessary to identify the output path for a field export.",
						MarkdownDescription: "FieldExportTarget provides the values necessary to identify the output path for a field export.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"key": {
								Description:         "Key overrides the default value ('<namespace>.<FieldExport-resource-name>') for the FieldExport target",
								MarkdownDescription: "Key overrides the default value ('<namespace>.<FieldExport-resource-name>') for the FieldExport target",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kind": {
								Description:         "FieldExportOutputType represents all types that can be produced by a field export operation",
								MarkdownDescription: "FieldExportOutputType represents all types that can be produced by a field export operation",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("configmap", "secret"),
								},
							},

							"name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace is marked as optional, so we cannot compose 'NamespacedName'",
								MarkdownDescription: "Namespace is marked as optional, so we cannot compose 'NamespacedName'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

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

func (r *ServicesK8SAwsFieldExportV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_services_k8s_aws_field_export_v1alpha1")

	var state ServicesK8SAwsFieldExportV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ServicesK8SAwsFieldExportV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("FieldExport")

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

func (r *ServicesK8SAwsFieldExportV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_services_k8s_aws_field_export_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ServicesK8SAwsFieldExportV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_services_k8s_aws_field_export_v1alpha1")

	var state ServicesK8SAwsFieldExportV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ServicesK8SAwsFieldExportV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("FieldExport")

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

func (r *ServicesK8SAwsFieldExportV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_services_k8s_aws_field_export_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
