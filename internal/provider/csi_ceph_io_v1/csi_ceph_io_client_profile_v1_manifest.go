/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package csi_ceph_io_v1

import (
	"context"
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
	_ datasource.DataSource = &CsiCephIoClientProfileV1Manifest{}
)

func NewCsiCephIoClientProfileV1Manifest() datasource.DataSource {
	return &CsiCephIoClientProfileV1Manifest{}
}

type CsiCephIoClientProfileV1Manifest struct{}

type CsiCephIoClientProfileV1ManifestData struct {
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
		CephConnectionRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"ceph_connection_ref" json:"cephConnectionRef,omitempty"`
		CephFs *struct {
			CephCsiSecrets *struct {
				ControllerPublishSecret *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"controller_publish_secret" json:"controllerPublishSecret,omitempty"`
			} `tfsdk:"ceph_csi_secrets" json:"cephCsiSecrets,omitempty"`
			FuseMountOptions   *map[string]string `tfsdk:"fuse_mount_options" json:"fuseMountOptions,omitempty"`
			KernelMountOptions *map[string]string `tfsdk:"kernel_mount_options" json:"kernelMountOptions,omitempty"`
			RadosNamespace     *string            `tfsdk:"rados_namespace" json:"radosNamespace,omitempty"`
			SubVolumeGroup     *string            `tfsdk:"sub_volume_group" json:"subVolumeGroup,omitempty"`
		} `tfsdk:"ceph_fs" json:"cephFs,omitempty"`
		Nfs *map[string]string `tfsdk:"nfs" json:"nfs,omitempty"`
		Rbd *struct {
			CephCsiSecrets *struct {
				ControllerPublishSecret *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"controller_publish_secret" json:"controllerPublishSecret,omitempty"`
			} `tfsdk:"ceph_csi_secrets" json:"cephCsiSecrets,omitempty"`
			RadosNamespace *string `tfsdk:"rados_namespace" json:"radosNamespace,omitempty"`
		} `tfsdk:"rbd" json:"rbd,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CsiCephIoClientProfileV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_csi_ceph_io_client_profile_v1_manifest"
}

func (r *CsiCephIoClientProfileV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClientProfile is the Schema for the clientprofiles API",
		MarkdownDescription: "ClientProfile is the Schema for the clientprofiles API",
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
				Description:         "ClientProfileSpec defines the desired state of Ceph CSI configuration for volumes and snapshots configured to use this profile",
				MarkdownDescription: "ClientProfileSpec defines the desired state of Ceph CSI configuration for volumes and snapshots configured to use this profile",
				Attributes: map[string]schema.Attribute{
					"ceph_connection_ref": schema.SingleNestedAttribute{
						Description:         "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
						MarkdownDescription: "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"ceph_fs": schema.SingleNestedAttribute{
						Description:         "CephFsConfigSpec defines the desired CephFs configuration",
						MarkdownDescription: "CephFsConfigSpec defines the desired CephFs configuration",
						Attributes: map[string]schema.Attribute{
							"ceph_csi_secrets": schema.SingleNestedAttribute{
								Description:         "CephCsiSecretsSpec defines the secrets used by the client profile to access the Ceph cluster and perform operations on volumes.",
								MarkdownDescription: "CephCsiSecretsSpec defines the secrets used by the client profile to access the Ceph cluster and perform operations on volumes.",
								Attributes: map[string]schema.Attribute{
									"controller_publish_secret": schema.SingleNestedAttribute{
										Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
										MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "name is unique within a namespace to reference a secret resource.",
												MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "namespace defines the space within which the secret name must be unique.",
												MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
												Required:            false,
												Optional:            true,
												Computed:            false,
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

							"fuse_mount_options": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kernel_mount_options": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rados_namespace": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sub_volume_group": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"nfs": schema.MapAttribute{
						Description:         "NfsConfigSpec cdefines the desired NFS configuration",
						MarkdownDescription: "NfsConfigSpec cdefines the desired NFS configuration",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"rbd": schema.SingleNestedAttribute{
						Description:         "RbdConfigSpec defines the desired RBD configuration",
						MarkdownDescription: "RbdConfigSpec defines the desired RBD configuration",
						Attributes: map[string]schema.Attribute{
							"ceph_csi_secrets": schema.SingleNestedAttribute{
								Description:         "CephCsiSecretsSpec defines the secrets used by the client profile to access the Ceph cluster and perform operations on volumes.",
								MarkdownDescription: "CephCsiSecretsSpec defines the secrets used by the client profile to access the Ceph cluster and perform operations on volumes.",
								Attributes: map[string]schema.Attribute{
									"controller_publish_secret": schema.SingleNestedAttribute{
										Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
										MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "name is unique within a namespace to reference a secret resource.",
												MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "namespace defines the space within which the secret name must be unique.",
												MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
												Required:            false,
												Optional:            true,
												Computed:            false,
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

							"rados_namespace": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
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
	}
}

func (r *CsiCephIoClientProfileV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_csi_ceph_io_client_profile_v1_manifest")

	var model CsiCephIoClientProfileV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("csi.ceph.io/v1")
	model.Kind = pointer.String("ClientProfile")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
