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

type HiveOpenshiftIoClusterProvisionV1Resource struct{}

var (
	_ resource.Resource = (*HiveOpenshiftIoClusterProvisionV1Resource)(nil)
)

type HiveOpenshiftIoClusterProvisionV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type HiveOpenshiftIoClusterProvisionV1GoModel struct {
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
		AdminKubeconfigSecretRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"admin_kubeconfig_secret_ref" yaml:"adminKubeconfigSecretRef,omitempty"`

		AdminPasswordSecretRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"admin_password_secret_ref" yaml:"adminPasswordSecretRef,omitempty"`

		Attempt *int64 `tfsdk:"attempt" yaml:"attempt,omitempty"`

		ClusterDeploymentRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"cluster_deployment_ref" yaml:"clusterDeploymentRef,omitempty"`

		ClusterID *string `tfsdk:"cluster_id" yaml:"clusterID,omitempty"`

		InfraID *string `tfsdk:"infra_id" yaml:"infraID,omitempty"`

		InstallLog *string `tfsdk:"install_log" yaml:"installLog,omitempty"`

		Metadata *map[string]string `tfsdk:"metadata" yaml:"metadata,omitempty"`

		PrevClusterID *string `tfsdk:"prev_cluster_id" yaml:"prevClusterID,omitempty"`

		PrevInfraID *string `tfsdk:"prev_infra_id" yaml:"prevInfraID,omitempty"`

		PrevProvisionName *string `tfsdk:"prev_provision_name" yaml:"prevProvisionName,omitempty"`

		Stage *string `tfsdk:"stage" yaml:"stage,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewHiveOpenshiftIoClusterProvisionV1Resource() resource.Resource {
	return &HiveOpenshiftIoClusterProvisionV1Resource{}
}

func (r *HiveOpenshiftIoClusterProvisionV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hive_openshift_io_cluster_provision_v1"
}

func (r *HiveOpenshiftIoClusterProvisionV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ClusterProvision is the Schema for the clusterprovisions API",
		MarkdownDescription: "ClusterProvision is the Schema for the clusterprovisions API",
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
				Description:         "ClusterProvisionSpec defines the results of provisioning a cluster.",
				MarkdownDescription: "ClusterProvisionSpec defines the results of provisioning a cluster.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"admin_kubeconfig_secret_ref": {
						Description:         "AdminKubeconfigSecretRef references the secret containing the admin kubeconfig for this cluster.",
						MarkdownDescription: "AdminKubeconfigSecretRef references the secret containing the admin kubeconfig for this cluster.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

					"admin_password_secret_ref": {
						Description:         "AdminPasswordSecretRef references the secret containing the admin username/password which can be used to login to this cluster.",
						MarkdownDescription: "AdminPasswordSecretRef references the secret containing the admin username/password which can be used to login to this cluster.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

					"attempt": {
						Description:         "Attempt is which attempt number of the cluster deployment that this ClusterProvision is",
						MarkdownDescription: "Attempt is which attempt number of the cluster deployment that this ClusterProvision is",

						Type: types.Int64Type,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"cluster_deployment_ref": {
						Description:         "ClusterDeploymentRef references the cluster deployment provisioned.",
						MarkdownDescription: "ClusterDeploymentRef references the cluster deployment provisioned.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

					"cluster_id": {
						Description:         "ClusterID is a globally unique identifier for this cluster generated during installation. Used for reporting metrics among other places.",
						MarkdownDescription: "ClusterID is a globally unique identifier for this cluster generated during installation. Used for reporting metrics among other places.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"infra_id": {
						Description:         "InfraID is an identifier for this cluster generated during installation and used for tagging/naming resources in cloud providers.",
						MarkdownDescription: "InfraID is an identifier for this cluster generated during installation and used for tagging/naming resources in cloud providers.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"install_log": {
						Description:         "InstallLog is the log from the installer.",
						MarkdownDescription: "InstallLog is the log from the installer.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"metadata": {
						Description:         "Metadata is the metadata.json generated by the installer, providing metadata information about the cluster created.",
						MarkdownDescription: "Metadata is the metadata.json generated by the installer, providing metadata information about the cluster created.",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"prev_cluster_id": {
						Description:         "PrevClusterID is the cluster ID of the previous failed provision attempt.",
						MarkdownDescription: "PrevClusterID is the cluster ID of the previous failed provision attempt.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"prev_infra_id": {
						Description:         "PrevInfraID is the infra ID of the previous failed provision attempt.",
						MarkdownDescription: "PrevInfraID is the infra ID of the previous failed provision attempt.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"prev_provision_name": {
						Description:         "PrevProvisionName is the name of the previous failed provision attempt.",
						MarkdownDescription: "PrevProvisionName is the name of the previous failed provision attempt.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"stage": {
						Description:         "Stage is the stage of provisioning that the cluster deployment has reached.",
						MarkdownDescription: "Stage is the stage of provisioning that the cluster deployment has reached.",

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

func (r *HiveOpenshiftIoClusterProvisionV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_hive_openshift_io_cluster_provision_v1")

	var state HiveOpenshiftIoClusterProvisionV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HiveOpenshiftIoClusterProvisionV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hive.openshift.io/v1")
	goModel.Kind = utilities.Ptr("ClusterProvision")

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

func (r *HiveOpenshiftIoClusterProvisionV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hive_openshift_io_cluster_provision_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *HiveOpenshiftIoClusterProvisionV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_hive_openshift_io_cluster_provision_v1")

	var state HiveOpenshiftIoClusterProvisionV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HiveOpenshiftIoClusterProvisionV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hive.openshift.io/v1")
	goModel.Kind = utilities.Ptr("ClusterProvision")

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

func (r *HiveOpenshiftIoClusterProvisionV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_hive_openshift_io_cluster_provision_v1")
	// NO-OP: Terraform removes the state automatically for us
}
