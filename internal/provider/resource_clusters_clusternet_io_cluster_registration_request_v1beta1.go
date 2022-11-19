/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"regexp"

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

type ClustersClusternetIoClusterRegistrationRequestV1Beta1Resource struct{}

var (
	_ resource.Resource = (*ClustersClusternetIoClusterRegistrationRequestV1Beta1Resource)(nil)
)

type ClustersClusternetIoClusterRegistrationRequestV1Beta1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ClustersClusternetIoClusterRegistrationRequestV1Beta1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		ClusterId *string `tfsdk:"cluster_id" yaml:"clusterId,omitempty"`

		ClusterLabels *map[string]string `tfsdk:"cluster_labels" yaml:"clusterLabels,omitempty"`

		ClusterName *string `tfsdk:"cluster_name" yaml:"clusterName,omitempty"`

		ClusterNamespace *string `tfsdk:"cluster_namespace" yaml:"clusterNamespace,omitempty"`

		ClusterType *string `tfsdk:"cluster_type" yaml:"clusterType,omitempty"`

		SyncMode *string `tfsdk:"sync_mode" yaml:"syncMode,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewClustersClusternetIoClusterRegistrationRequestV1Beta1Resource() resource.Resource {
	return &ClustersClusternetIoClusterRegistrationRequestV1Beta1Resource{}
}

func (r *ClustersClusternetIoClusterRegistrationRequestV1Beta1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusters_clusternet_io_cluster_registration_request_v1beta1"
}

func (r *ClustersClusternetIoClusterRegistrationRequestV1Beta1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ClusterRegistrationRequest is the Schema for the clusterregistrationrequests API",
		MarkdownDescription: "ClusterRegistrationRequest is the Schema for the clusterregistrationrequests API",
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
				Description:         "ClusterRegistrationRequestSpec defines the desired state of ClusterRegistrationRequest",
				MarkdownDescription: "ClusterRegistrationRequestSpec defines the desired state of ClusterRegistrationRequest",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"cluster_id": {
						Description:         "ClusterID, a Random (Version 4) UUID, is a unique value in time and space value representing for child cluster. It is typically generated by the clusternet agent on the successful creation of a 'self-cluster' Lease in the child cluster. Also it is not allowed to change on PUT operations.",
						MarkdownDescription: "ClusterID, a Random (Version 4) UUID, is a unique value in time and space value representing for child cluster. It is typically generated by the clusternet agent on the successful creation of a 'self-cluster' Lease in the child cluster. Also it is not allowed to change on PUT operations.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
						},
					},

					"cluster_labels": {
						Description:         "ClusterLabels is the labels of the child cluster.",
						MarkdownDescription: "ClusterLabels is the labels of the child cluster.",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_name": {
						Description:         "ClusterName is the cluster name. a lower case alphanumeric characters or '-', and must start and end with an alphanumeric character",
						MarkdownDescription: "ClusterName is the cluster name. a lower case alphanumeric characters or '-', and must start and end with an alphanumeric character",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.LengthAtMost(30),

							stringvalidator.RegexMatches(regexp.MustCompile(`[a-z0-9]([-a-z0-9]*[a-z0-9])?([a-z0-9]([-a-z0-9]*[a-z0-9]))*`), ""),
						},
					},

					"cluster_namespace": {
						Description:         "ClusterNamespace is the dedicated namespace of the cluster.",
						MarkdownDescription: "ClusterNamespace is the dedicated namespace of the cluster.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.LengthAtMost(63),

							stringvalidator.RegexMatches(regexp.MustCompile(`[a-z0-9]([-a-z0-9]*[a-z0-9])?`), ""),
						},
					},

					"cluster_type": {
						Description:         "ClusterType denotes the type of the child cluster.",
						MarkdownDescription: "ClusterType denotes the type of the child cluster.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"sync_mode": {
						Description:         "SyncMode decides how to sync resources from parent cluster to child cluster.",
						MarkdownDescription: "SyncMode decides how to sync resources from parent cluster to child cluster.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("Push", "Pull", "Dual"),
						},
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *ClustersClusternetIoClusterRegistrationRequestV1Beta1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_clusters_clusternet_io_cluster_registration_request_v1beta1")

	var state ClustersClusternetIoClusterRegistrationRequestV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ClustersClusternetIoClusterRegistrationRequestV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("clusters.clusternet.io/v1beta1")
	goModel.Kind = utilities.Ptr("ClusterRegistrationRequest")

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

func (r *ClustersClusternetIoClusterRegistrationRequestV1Beta1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_clusters_clusternet_io_cluster_registration_request_v1beta1")
	// NO-OP: All data is already in Terraform state
}

func (r *ClustersClusternetIoClusterRegistrationRequestV1Beta1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_clusters_clusternet_io_cluster_registration_request_v1beta1")

	var state ClustersClusternetIoClusterRegistrationRequestV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ClustersClusternetIoClusterRegistrationRequestV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("clusters.clusternet.io/v1beta1")
	goModel.Kind = utilities.Ptr("ClusterRegistrationRequest")

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

func (r *ClustersClusternetIoClusterRegistrationRequestV1Beta1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_clusters_clusternet_io_cluster_registration_request_v1beta1")
	// NO-OP: Terraform removes the state automatically for us
}
