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

type ImagingIngestionAlvearieOrgDimseProxyV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ImagingIngestionAlvearieOrgDimseProxyV1Alpha1Resource)(nil)
)

type ImagingIngestionAlvearieOrgDimseProxyV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ImagingIngestionAlvearieOrgDimseProxyV1Alpha1GoModel struct {
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
		ApplicationEntityTitle *string `tfsdk:"application_entity_title" yaml:"applicationEntityTitle,omitempty"`

		ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

		ImagePullSecrets *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"image_pull_secrets" yaml:"imagePullSecrets,omitempty"`

		NatsSecure *bool `tfsdk:"nats_secure" yaml:"natsSecure,omitempty"`

		NatsSubjectChannel *string `tfsdk:"nats_subject_channel" yaml:"natsSubjectChannel,omitempty"`

		NatsSubjectRoot *string `tfsdk:"nats_subject_root" yaml:"natsSubjectRoot,omitempty"`

		NatsTokenSecret *string `tfsdk:"nats_token_secret" yaml:"natsTokenSecret,omitempty"`

		NatsUrl *string `tfsdk:"nats_url" yaml:"natsUrl,omitempty"`

		Proxy *struct {
			Image *string `tfsdk:"image" yaml:"image,omitempty"`
		} `tfsdk:"proxy" yaml:"proxy,omitempty"`

		TargetDimseHost *string `tfsdk:"target_dimse_host" yaml:"targetDimseHost,omitempty"`

		TargetDimsePort *int64 `tfsdk:"target_dimse_port" yaml:"targetDimsePort,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewImagingIngestionAlvearieOrgDimseProxyV1Alpha1Resource() resource.Resource {
	return &ImagingIngestionAlvearieOrgDimseProxyV1Alpha1Resource{}
}

func (r *ImagingIngestionAlvearieOrgDimseProxyV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_imaging_ingestion_alvearie_org_dimse_proxy_v1alpha1"
}

func (r *ImagingIngestionAlvearieOrgDimseProxyV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Provides a bidirectional proxied DIMSE Application Entity (AE) in the cluster",
		MarkdownDescription: "Provides a bidirectional proxied DIMSE Application Entity (AE) in the cluster",
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
				Description:         "DimseProxySpec defines the desired state of DimseProxy",
				MarkdownDescription: "DimseProxySpec defines the desired state of DimseProxy",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"application_entity_title": {
						Description:         "Application Entity Title",
						MarkdownDescription: "Application Entity Title",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"image_pull_policy": {
						Description:         "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
						MarkdownDescription: "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"image_pull_secrets": {
						Description:         "Image Pull Secrets",
						MarkdownDescription: "Image Pull Secrets",

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

						Required: false,
						Optional: true,
						Computed: false,
					},

					"nats_secure": {
						Description:         "Make NATS Connection Secure",
						MarkdownDescription: "Make NATS Connection Secure",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"nats_subject_channel": {
						Description:         "NATS Subject Channel to use",
						MarkdownDescription: "NATS Subject Channel to use",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"nats_subject_root": {
						Description:         "NATS Subject Root",
						MarkdownDescription: "NATS Subject Root",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"nats_token_secret": {
						Description:         "NATS Token Secret Name",
						MarkdownDescription: "NATS Token Secret Name",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"nats_url": {
						Description:         "NATS URL",
						MarkdownDescription: "NATS URL",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"proxy": {
						Description:         "DIMSE Proxy Spec",
						MarkdownDescription: "DIMSE Proxy Spec",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"image": {
								Description:         "Image",
								MarkdownDescription: "Image",

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

					"target_dimse_host": {
						Description:         "Target Dimse Host",
						MarkdownDescription: "Target Dimse Host",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"target_dimse_port": {
						Description:         "Target Dimse Port",
						MarkdownDescription: "Target Dimse Port",

						Type: types.Int64Type,

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

func (r *ImagingIngestionAlvearieOrgDimseProxyV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_imaging_ingestion_alvearie_org_dimse_proxy_v1alpha1")

	var state ImagingIngestionAlvearieOrgDimseProxyV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ImagingIngestionAlvearieOrgDimseProxyV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("imaging-ingestion.alvearie.org/v1alpha1")
	goModel.Kind = utilities.Ptr("DimseProxy")

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

func (r *ImagingIngestionAlvearieOrgDimseProxyV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_imaging_ingestion_alvearie_org_dimse_proxy_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ImagingIngestionAlvearieOrgDimseProxyV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_imaging_ingestion_alvearie_org_dimse_proxy_v1alpha1")

	var state ImagingIngestionAlvearieOrgDimseProxyV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ImagingIngestionAlvearieOrgDimseProxyV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("imaging-ingestion.alvearie.org/v1alpha1")
	goModel.Kind = utilities.Ptr("DimseProxy")

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

func (r *ImagingIngestionAlvearieOrgDimseProxyV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_imaging_ingestion_alvearie_org_dimse_proxy_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
