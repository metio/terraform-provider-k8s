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

type SourceToolkitFluxcdIoBucketV1Beta1Resource struct{}

var (
	_ resource.Resource = (*SourceToolkitFluxcdIoBucketV1Beta1Resource)(nil)
)

type SourceToolkitFluxcdIoBucketV1Beta1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type SourceToolkitFluxcdIoBucketV1Beta1GoModel struct {
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
		AccessFrom *struct {
			NamespaceSelectors *[]struct {
				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"namespace_selectors" yaml:"namespaceSelectors,omitempty"`
		} `tfsdk:"access_from" yaml:"accessFrom,omitempty"`

		BucketName *string `tfsdk:"bucket_name" yaml:"bucketName,omitempty"`

		Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

		Ignore *string `tfsdk:"ignore" yaml:"ignore,omitempty"`

		Insecure *bool `tfsdk:"insecure" yaml:"insecure,omitempty"`

		Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

		Provider *string `tfsdk:"provider" yaml:"provider,omitempty"`

		Region *string `tfsdk:"region" yaml:"region,omitempty"`

		SecretRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

		Suspend *bool `tfsdk:"suspend" yaml:"suspend,omitempty"`

		Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewSourceToolkitFluxcdIoBucketV1Beta1Resource() resource.Resource {
	return &SourceToolkitFluxcdIoBucketV1Beta1Resource{}
}

func (r *SourceToolkitFluxcdIoBucketV1Beta1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_toolkit_fluxcd_io_bucket_v1beta1"
}

func (r *SourceToolkitFluxcdIoBucketV1Beta1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Bucket is the Schema for the buckets API",
		MarkdownDescription: "Bucket is the Schema for the buckets API",
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
				Description:         "BucketSpec defines the desired state of an S3 compatible bucket",
				MarkdownDescription: "BucketSpec defines the desired state of an S3 compatible bucket",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"access_from": {
						Description:         "AccessFrom defines an Access Control List for allowing cross-namespace references to this object.",
						MarkdownDescription: "AccessFrom defines an Access Control List for allowing cross-namespace references to this object.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"namespace_selectors": {
								Description:         "NamespaceSelectors is the list of namespace selectors to which this ACL applies. Items in this list are evaluated using a logical OR operation.",
								MarkdownDescription: "NamespaceSelectors is the list of namespace selectors to which this ACL applies. Items in this list are evaluated using a logical OR operation.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_labels": {
										Description:         "MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

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

					"bucket_name": {
						Description:         "The bucket name.",
						MarkdownDescription: "The bucket name.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"endpoint": {
						Description:         "The bucket endpoint address.",
						MarkdownDescription: "The bucket endpoint address.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"ignore": {
						Description:         "Ignore overrides the set of excluded patterns in the .sourceignore format (which is the same as .gitignore). If not provided, a default will be used, consult the documentation for your version to find out what those are.",
						MarkdownDescription: "Ignore overrides the set of excluded patterns in the .sourceignore format (which is the same as .gitignore). If not provided, a default will be used, consult the documentation for your version to find out what those are.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"insecure": {
						Description:         "Insecure allows connecting to a non-TLS S3 HTTP endpoint.",
						MarkdownDescription: "Insecure allows connecting to a non-TLS S3 HTTP endpoint.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"interval": {
						Description:         "The interval at which to check for bucket updates.",
						MarkdownDescription: "The interval at which to check for bucket updates.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"provider": {
						Description:         "The S3 compatible storage provider name, default ('generic').",
						MarkdownDescription: "The S3 compatible storage provider name, default ('generic').",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("generic", "aws", "gcp"),
						},
					},

					"region": {
						Description:         "The bucket region.",
						MarkdownDescription: "The bucket region.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"secret_ref": {
						Description:         "The name of the secret containing authentication credentials for the Bucket.",
						MarkdownDescription: "The name of the secret containing authentication credentials for the Bucket.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",

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

					"suspend": {
						Description:         "This flag tells the controller to suspend the reconciliation of this source.",
						MarkdownDescription: "This flag tells the controller to suspend the reconciliation of this source.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"timeout": {
						Description:         "The timeout for download operations, defaults to 60s.",
						MarkdownDescription: "The timeout for download operations, defaults to 60s.",

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

func (r *SourceToolkitFluxcdIoBucketV1Beta1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_source_toolkit_fluxcd_io_bucket_v1beta1")

	var state SourceToolkitFluxcdIoBucketV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SourceToolkitFluxcdIoBucketV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("source.toolkit.fluxcd.io/v1beta1")
	goModel.Kind = utilities.Ptr("Bucket")

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

func (r *SourceToolkitFluxcdIoBucketV1Beta1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_source_toolkit_fluxcd_io_bucket_v1beta1")
	// NO-OP: All data is already in Terraform state
}

func (r *SourceToolkitFluxcdIoBucketV1Beta1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_source_toolkit_fluxcd_io_bucket_v1beta1")

	var state SourceToolkitFluxcdIoBucketV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SourceToolkitFluxcdIoBucketV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("source.toolkit.fluxcd.io/v1beta1")
	goModel.Kind = utilities.Ptr("Bucket")

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

func (r *SourceToolkitFluxcdIoBucketV1Beta1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_source_toolkit_fluxcd_io_bucket_v1beta1")
	// NO-OP: Terraform removes the state automatically for us
}
