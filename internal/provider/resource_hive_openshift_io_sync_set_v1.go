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

type HiveOpenshiftIoSyncSetV1Resource struct{}

var (
	_ resource.Resource = (*HiveOpenshiftIoSyncSetV1Resource)(nil)
)

type HiveOpenshiftIoSyncSetV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type HiveOpenshiftIoSyncSetV1GoModel struct {
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
		ApplyBehavior *string `tfsdk:"apply_behavior" yaml:"applyBehavior,omitempty"`

		ClusterDeploymentRefs *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"cluster_deployment_refs" yaml:"clusterDeploymentRefs,omitempty"`

		Patches *[]struct {
			ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

			Patch *string `tfsdk:"patch" yaml:"patch,omitempty"`

			PatchType *string `tfsdk:"patch_type" yaml:"patchType,omitempty"`
		} `tfsdk:"patches" yaml:"patches,omitempty"`

		ResourceApplyMode *string `tfsdk:"resource_apply_mode" yaml:"resourceApplyMode,omitempty"`

		Resources *[]map[string]string `tfsdk:"resources" yaml:"resources,omitempty"`

		SecretMappings *[]struct {
			SourceRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"source_ref" yaml:"sourceRef,omitempty"`

			TargetRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"target_ref" yaml:"targetRef,omitempty"`
		} `tfsdk:"secret_mappings" yaml:"secretMappings,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewHiveOpenshiftIoSyncSetV1Resource() resource.Resource {
	return &HiveOpenshiftIoSyncSetV1Resource{}
}

func (r *HiveOpenshiftIoSyncSetV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hive_openshift_io_sync_set_v1"
}

func (r *HiveOpenshiftIoSyncSetV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "SyncSet is the Schema for the SyncSet API",
		MarkdownDescription: "SyncSet is the Schema for the SyncSet API",
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
				Description:         "SyncSetSpec defines the SyncSetCommonSpec resources and patches to sync along with ClusterDeploymentRefs indicating which clusters the SyncSet applies to in the SyncSet's namespace.",
				MarkdownDescription: "SyncSetSpec defines the SyncSetCommonSpec resources and patches to sync along with ClusterDeploymentRefs indicating which clusters the SyncSet applies to in the SyncSet's namespace.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"apply_behavior": {
						Description:         "ApplyBehavior indicates how resources in this syncset will be applied to the target cluster. The default value of 'Apply' indicates that resources should be applied using the 'oc apply' command. If no value is set, 'Apply' is assumed. A value of 'CreateOnly' indicates that the resource will only be created if it does not already exist in the target cluster. Otherwise, it will be left alone. A value of 'CreateOrUpdate' indicates that the resource will be created/updated without the use of the 'oc apply' command, allowing larger resources to be synced, but losing some functionality of the 'oc apply' command such as the ability to remove annotations, labels, and other map entries in general.",
						MarkdownDescription: "ApplyBehavior indicates how resources in this syncset will be applied to the target cluster. The default value of 'Apply' indicates that resources should be applied using the 'oc apply' command. If no value is set, 'Apply' is assumed. A value of 'CreateOnly' indicates that the resource will only be created if it does not already exist in the target cluster. Otherwise, it will be left alone. A value of 'CreateOrUpdate' indicates that the resource will be created/updated without the use of the 'oc apply' command, allowing larger resources to be synced, but losing some functionality of the 'oc apply' command such as the ability to remove annotations, labels, and other map entries in general.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("", "Apply", "CreateOnly", "CreateOrUpdate"),
						},
					},

					"cluster_deployment_refs": {
						Description:         "ClusterDeploymentRefs is the list of LocalObjectReference indicating which clusters the SyncSet applies to in the SyncSet's namespace.",
						MarkdownDescription: "ClusterDeploymentRefs is the list of LocalObjectReference indicating which clusters the SyncSet applies to in the SyncSet's namespace.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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

					"patches": {
						Description:         "Patches is the list of patches to apply.",
						MarkdownDescription: "Patches is the list of patches to apply.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"api_version": {
								Description:         "APIVersion is the Group and Version of the object to be patched.",
								MarkdownDescription: "APIVersion is the Group and Version of the object to be patched.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"kind": {
								Description:         "Kind is the Kind of the object to be patched.",
								MarkdownDescription: "Kind is the Kind of the object to be patched.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"name": {
								Description:         "Name is the name of the object to be patched.",
								MarkdownDescription: "Name is the name of the object to be patched.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace is the Namespace in which the object to patch exists. Defaults to the SyncSet's Namespace.",
								MarkdownDescription: "Namespace is the Namespace in which the object to patch exists. Defaults to the SyncSet's Namespace.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"patch": {
								Description:         "Patch is the patch to apply.",
								MarkdownDescription: "Patch is the patch to apply.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"patch_type": {
								Description:         "PatchType indicates the PatchType as 'strategic' (default), 'json', or 'merge'.",
								MarkdownDescription: "PatchType indicates the PatchType as 'strategic' (default), 'json', or 'merge'.",

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

					"resource_apply_mode": {
						Description:         "ResourceApplyMode indicates if the Resource apply mode is 'Upsert' (default) or 'Sync'. ApplyMode 'Upsert' indicates create and update. ApplyMode 'Sync' indicates create, update and delete.",
						MarkdownDescription: "ResourceApplyMode indicates if the Resource apply mode is 'Upsert' (default) or 'Sync'. ApplyMode 'Upsert' indicates create and update. ApplyMode 'Sync' indicates create, update and delete.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resources": {
						Description:         "Resources is the list of objects to sync from RawExtension definitions.",
						MarkdownDescription: "Resources is the list of objects to sync from RawExtension definitions.",

						Type: types.ListType{ElemType: types.MapType{ElemType: types.StringType}},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"secret_mappings": {
						Description:         "Secrets is the list of secrets to sync along with their respective destinations.",
						MarkdownDescription: "Secrets is the list of secrets to sync along with their respective destinations.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"source_ref": {
								Description:         "SourceRef specifies the name and namespace of a secret on the management cluster",
								MarkdownDescription: "SourceRef specifies the name and namespace of a secret on the management cluster",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name is the name of the secret",
										MarkdownDescription: "Name is the name of the secret",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"namespace": {
										Description:         "Namespace is the namespace where the secret lives. If not present for the source secret reference, it is assumed to be the same namespace as the syncset with the reference.",
										MarkdownDescription: "Namespace is the namespace where the secret lives. If not present for the source secret reference, it is assumed to be the same namespace as the syncset with the reference.",

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

							"target_ref": {
								Description:         "TargetRef specifies the target name and namespace of the secret on the target cluster",
								MarkdownDescription: "TargetRef specifies the target name and namespace of the secret on the target cluster",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name is the name of the secret",
										MarkdownDescription: "Name is the name of the secret",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"namespace": {
										Description:         "Namespace is the namespace where the secret lives. If not present for the source secret reference, it is assumed to be the same namespace as the syncset with the reference.",
										MarkdownDescription: "Namespace is the namespace where the secret lives. If not present for the source secret reference, it is assumed to be the same namespace as the syncset with the reference.",

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
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *HiveOpenshiftIoSyncSetV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_hive_openshift_io_sync_set_v1")

	var state HiveOpenshiftIoSyncSetV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HiveOpenshiftIoSyncSetV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hive.openshift.io/v1")
	goModel.Kind = utilities.Ptr("SyncSet")

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

func (r *HiveOpenshiftIoSyncSetV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hive_openshift_io_sync_set_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *HiveOpenshiftIoSyncSetV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_hive_openshift_io_sync_set_v1")

	var state HiveOpenshiftIoSyncSetV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HiveOpenshiftIoSyncSetV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hive.openshift.io/v1")
	goModel.Kind = utilities.Ptr("SyncSet")

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

func (r *HiveOpenshiftIoSyncSetV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_hive_openshift_io_sync_set_v1")
	// NO-OP: Terraform removes the state automatically for us
}
