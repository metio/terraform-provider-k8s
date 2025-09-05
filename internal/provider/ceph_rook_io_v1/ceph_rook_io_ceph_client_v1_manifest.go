/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ceph_rook_io_v1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &CephRookIoCephClientV1Manifest{}
)

func NewCephRookIoCephClientV1Manifest() datasource.DataSource {
	return &CephRookIoCephClientV1Manifest{}
}

type CephRookIoCephClientV1Manifest struct{}

type CephRookIoCephClientV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Caps         *map[string]string `tfsdk:"caps" json:"caps,omitempty"`
		Name         *string            `tfsdk:"name" json:"name,omitempty"`
		RemoveSecret *bool              `tfsdk:"remove_secret" json:"removeSecret,omitempty"`
		SecretName   *string            `tfsdk:"secret_name" json:"secretName,omitempty"`
		Security     *struct {
			Cephx *struct {
				KeyGeneration     *int64  `tfsdk:"key_generation" json:"keyGeneration,omitempty"`
				KeyRotationPolicy *string `tfsdk:"key_rotation_policy" json:"keyRotationPolicy,omitempty"`
			} `tfsdk:"cephx" json:"cephx,omitempty"`
		} `tfsdk:"security" json:"security,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CephRookIoCephClientV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ceph_rook_io_ceph_client_v1_manifest"
}

func (r *CephRookIoCephClientV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CephClient represents a Ceph Client",
		MarkdownDescription: "CephClient represents a Ceph Client",
		Attributes: map[string]schema.Attribute{
			"yaml": schema.StringAttribute{
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"metadata": schema.SingleNestedAttribute{
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Optional:            false,
				Computed:            false,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.NameValidator(),
							stringvalidator.LengthAtLeast(1),
						},
					},

					"namespace": schema.StringAttribute{
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.NameValidator(),
							stringvalidator.LengthAtLeast(1),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Map{
							validators.LabelValidator(),
						},
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "Spec represents the specification of a Ceph Client",
				MarkdownDescription: "Spec represents the specification of a Ceph Client",
				Attributes: map[string]schema.Attribute{
					"caps": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"remove_secret": schema.BoolAttribute{
						Description:         "RemoveSecret indicates whether the current secret for this ceph client should be removed or not. If true, the K8s secret will be deleted, but the cephx keyring will remain until the CR is deleted.",
						MarkdownDescription: "RemoveSecret indicates whether the current secret for this ceph client should be removed or not. If true, the K8s secret will be deleted, but the cephx keyring will remain until the CR is deleted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"secret_name": schema.StringAttribute{
						Description:         "SecretName is the name of the secret created for this ceph client. If not specified, the default name is 'rook-ceph-client-' as a prefix to the CR name.",
						MarkdownDescription: "SecretName is the name of the secret created for this ceph client. If not specified, the default name is 'rook-ceph-client-' as a prefix to the CR name.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"security": schema.SingleNestedAttribute{
						Description:         "Security represents security settings",
						MarkdownDescription: "Security represents security settings",
						Attributes: map[string]schema.Attribute{
							"cephx": schema.SingleNestedAttribute{
								Description:         "CephX configures CephX key settings. More: https://docs.ceph.com/en/latest/dev/cephx/",
								MarkdownDescription: "CephX configures CephX key settings. More: https://docs.ceph.com/en/latest/dev/cephx/",
								Attributes: map[string]schema.Attribute{
									"key_generation": schema.Int64Attribute{
										Description:         "KeyGeneration specifies the desired CephX key generation. This is used when KeyRotationPolicy is KeyGeneration and ignored for other policies. If this is set to greater than the current key generation, relevant keys will be rotated, and the generation value will be updated to this new value (generation values are not necessarily incremental, though that is the intended use case). If this is set to less than or equal to the current key generation, keys are not rotated.",
										MarkdownDescription: "KeyGeneration specifies the desired CephX key generation. This is used when KeyRotationPolicy is KeyGeneration and ignored for other policies. If this is set to greater than the current key generation, relevant keys will be rotated, and the generation value will be updated to this new value (generation values are not necessarily incremental, though that is the intended use case). If this is set to less than or equal to the current key generation, keys are not rotated.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
											int64validator.AtMost(4.294967295e+09),
										},
									},

									"key_rotation_policy": schema.StringAttribute{
										Description:         "KeyRotationPolicy controls if and when CephX keys are rotated after initial creation. One of Disabled, or KeyGeneration. Default Disabled.",
										MarkdownDescription: "KeyRotationPolicy controls if and when CephX keys are rotated after initial creation. One of Disabled, or KeyGeneration. Default Disabled.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("", "Disabled", "KeyGeneration"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *CephRookIoCephClientV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ceph_rook_io_ceph_client_v1_manifest")

	var model CephRookIoCephClientV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("ceph.rook.io/v1")
	model.Kind = pointer.String("CephClient")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
