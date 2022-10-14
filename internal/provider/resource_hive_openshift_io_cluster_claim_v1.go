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

type HiveOpenshiftIoClusterClaimV1Resource struct{}

var (
	_ resource.Resource = (*HiveOpenshiftIoClusterClaimV1Resource)(nil)
)

type HiveOpenshiftIoClusterClaimV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type HiveOpenshiftIoClusterClaimV1GoModel struct {
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
		ClusterPoolName *string `tfsdk:"cluster_pool_name" yaml:"clusterPoolName,omitempty"`

		Lifetime *string `tfsdk:"lifetime" yaml:"lifetime,omitempty"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Subjects *[]struct {
			ApiGroup *string `tfsdk:"api_group" yaml:"apiGroup,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
		} `tfsdk:"subjects" yaml:"subjects,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewHiveOpenshiftIoClusterClaimV1Resource() resource.Resource {
	return &HiveOpenshiftIoClusterClaimV1Resource{}
}

func (r *HiveOpenshiftIoClusterClaimV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hive_openshift_io_cluster_claim_v1"
}

func (r *HiveOpenshiftIoClusterClaimV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ClusterClaim represents a claim to a cluster from a cluster pool.",
		MarkdownDescription: "ClusterClaim represents a claim to a cluster from a cluster pool.",
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
				Description:         "ClusterClaimSpec defines the desired state of the ClusterClaim.",
				MarkdownDescription: "ClusterClaimSpec defines the desired state of the ClusterClaim.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"cluster_pool_name": {
						Description:         "ClusterPoolName is the name of the cluster pool from which to claim a cluster.",
						MarkdownDescription: "ClusterPoolName is the name of the cluster pool from which to claim a cluster.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"lifetime": {
						Description:         "Lifetime is the maximum lifetime of the claim after it is assigned a cluster. If the claim still exists when the lifetime has elapsed, the claim will be deleted by Hive. This is a Duration value; see https://pkg.go.dev/time#ParseDuration for accepted formats. Note: due to discrepancies in validation vs parsing, we use a Pattern instead of 'Format=duration'. See https://bugzilla.redhat.com/show_bug.cgi?id=2050332 https://github.com/kubernetes/apimachinery/issues/131 https://github.com/kubernetes/apiextensions-apiserver/issues/56",
						MarkdownDescription: "Lifetime is the maximum lifetime of the claim after it is assigned a cluster. If the claim still exists when the lifetime has elapsed, the claim will be deleted by Hive. This is a Duration value; see https://pkg.go.dev/time#ParseDuration for accepted formats. Note: due to discrepancies in validation vs parsing, we use a Pattern instead of 'Format=duration'. See https://bugzilla.redhat.com/show_bug.cgi?id=2050332 https://github.com/kubernetes/apimachinery/issues/131 https://github.com/kubernetes/apiextensions-apiserver/issues/56",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ns|us|µs|ms|s|m|h))+$`), ""),
						},
					},

					"namespace": {
						Description:         "Namespace is the namespace containing the ClusterDeployment (name will match the namespace) of the claimed cluster. This field will be set as soon as a suitable cluster can be found, however that cluster may still be resuming and not yet ready for use. Wait for the ClusterRunning condition to be true to avoid this issue.",
						MarkdownDescription: "Namespace is the namespace containing the ClusterDeployment (name will match the namespace) of the claimed cluster. This field will be set as soon as a suitable cluster can be found, however that cluster may still be resuming and not yet ready for use. Wait for the ClusterRunning condition to be true to avoid this issue.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"subjects": {
						Description:         "Subjects hold references to which to authorize access to the claimed cluster.",
						MarkdownDescription: "Subjects hold references to which to authorize access to the claimed cluster.",

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
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *HiveOpenshiftIoClusterClaimV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_hive_openshift_io_cluster_claim_v1")

	var state HiveOpenshiftIoClusterClaimV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HiveOpenshiftIoClusterClaimV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hive.openshift.io/v1")
	goModel.Kind = utilities.Ptr("ClusterClaim")

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

func (r *HiveOpenshiftIoClusterClaimV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hive_openshift_io_cluster_claim_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *HiveOpenshiftIoClusterClaimV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_hive_openshift_io_cluster_claim_v1")

	var state HiveOpenshiftIoClusterClaimV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HiveOpenshiftIoClusterClaimV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hive.openshift.io/v1")
	goModel.Kind = utilities.Ptr("ClusterClaim")

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

func (r *HiveOpenshiftIoClusterClaimV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_hive_openshift_io_cluster_claim_v1")
	// NO-OP: Terraform removes the state automatically for us
}
