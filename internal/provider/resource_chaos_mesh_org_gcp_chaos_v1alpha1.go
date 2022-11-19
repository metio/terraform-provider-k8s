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

type ChaosMeshOrgGCPChaosV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ChaosMeshOrgGCPChaosV1Alpha1Resource)(nil)
)

type ChaosMeshOrgGCPChaosV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ChaosMeshOrgGCPChaosV1Alpha1GoModel struct {
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
		Action *string `tfsdk:"action" yaml:"action,omitempty"`

		DeviceNames *[]string `tfsdk:"device_names" yaml:"deviceNames,omitempty"`

		Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

		Instance *string `tfsdk:"instance" yaml:"instance,omitempty"`

		Project *string `tfsdk:"project" yaml:"project,omitempty"`

		SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`

		Zone *string `tfsdk:"zone" yaml:"zone,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewChaosMeshOrgGCPChaosV1Alpha1Resource() resource.Resource {
	return &ChaosMeshOrgGCPChaosV1Alpha1Resource{}
}

func (r *ChaosMeshOrgGCPChaosV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_chaos_mesh_org_gcp_chaos_v1alpha1"
}

func (r *ChaosMeshOrgGCPChaosV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "GCPChaos is the Schema for the gcpchaos API",
		MarkdownDescription: "GCPChaos is the Schema for the gcpchaos API",
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
				Description:         "GCPChaosSpec is the content of the specification for a GCPChaos",
				MarkdownDescription: "GCPChaosSpec is the content of the specification for a GCPChaos",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"action": {
						Description:         "Action defines the specific gcp chaos action. Supported action: node-stop / node-reset / disk-loss Default action: node-stop",
						MarkdownDescription: "Action defines the specific gcp chaos action. Supported action: node-stop / node-reset / disk-loss Default action: node-stop",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("node-stop", "node-reset", "disk-loss"),
						},
					},

					"device_names": {
						Description:         "The device name of disks to detach. Needed in disk-loss.",
						MarkdownDescription: "The device name of disks to detach. Needed in disk-loss.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"duration": {
						Description:         "Duration represents the duration of the chaos action.",
						MarkdownDescription: "Duration represents the duration of the chaos action.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"instance": {
						Description:         "Instance defines the name of the instance",
						MarkdownDescription: "Instance defines the name of the instance",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"project": {
						Description:         "Project defines the ID of gcp project.",
						MarkdownDescription: "Project defines the ID of gcp project.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"secret_name": {
						Description:         "SecretName defines the name of kubernetes secret. It is used for GCP credentials.",
						MarkdownDescription: "SecretName defines the name of kubernetes secret. It is used for GCP credentials.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"zone": {
						Description:         "Zone defines the zone of gcp project.",
						MarkdownDescription: "Zone defines the zone of gcp project.",

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
		},
	}, nil
}

func (r *ChaosMeshOrgGCPChaosV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_chaos_mesh_org_gcp_chaos_v1alpha1")

	var state ChaosMeshOrgGCPChaosV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ChaosMeshOrgGCPChaosV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("chaos-mesh.org/v1alpha1")
	goModel.Kind = utilities.Ptr("GCPChaos")

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

func (r *ChaosMeshOrgGCPChaosV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_chaos_mesh_org_gcp_chaos_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ChaosMeshOrgGCPChaosV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_chaos_mesh_org_gcp_chaos_v1alpha1")

	var state ChaosMeshOrgGCPChaosV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ChaosMeshOrgGCPChaosV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("chaos-mesh.org/v1alpha1")
	goModel.Kind = utilities.Ptr("GCPChaos")

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

func (r *ChaosMeshOrgGCPChaosV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_chaos_mesh_org_gcp_chaos_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
