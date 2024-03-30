/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package redhatcop_redhat_io_v1alpha1

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
	_ datasource.DataSource = &RedhatcopRedhatIoKeepalivedGroupV1Alpha1Manifest{}
)

func NewRedhatcopRedhatIoKeepalivedGroupV1Alpha1Manifest() datasource.DataSource {
	return &RedhatcopRedhatIoKeepalivedGroupV1Alpha1Manifest{}
}

type RedhatcopRedhatIoKeepalivedGroupV1Alpha1Manifest struct{}

type RedhatcopRedhatIoKeepalivedGroupV1Alpha1ManifestData struct {
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
		BlacklistRouterIDs            *[]string          `tfsdk:"blacklist_router_i_ds" json:"blacklistRouterIDs,omitempty"`
		DaemonsetPodAnnotations       *map[string]string `tfsdk:"daemonset_pod_annotations" json:"daemonsetPodAnnotations,omitempty"`
		DaemonsetPodPriorityClassName *string            `tfsdk:"daemonset_pod_priority_class_name" json:"daemonsetPodPriorityClassName,omitempty"`
		Image                         *string            `tfsdk:"image" json:"image,omitempty"`
		Interface                     *string            `tfsdk:"interface" json:"interface,omitempty"`
		InterfaceFromIP               *string            `tfsdk:"interface_from_ip" json:"interfaceFromIP,omitempty"`
		NodeSelector                  *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		PasswordAuth                  *struct {
			SecretKey *string `tfsdk:"secret_key" json:"secretKey,omitempty"`
			SecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"password_auth" json:"passwordAuth,omitempty"`
		UnicastEnabled *bool              `tfsdk:"unicast_enabled" json:"unicastEnabled,omitempty"`
		VerbatimConfig *map[string]string `tfsdk:"verbatim_config" json:"verbatimConfig,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RedhatcopRedhatIoKeepalivedGroupV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_redhatcop_redhat_io_keepalived_group_v1alpha1_manifest"
}

func (r *RedhatcopRedhatIoKeepalivedGroupV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "KeepalivedGroup is the Schema for the keepalivedgroups API",
		MarkdownDescription: "KeepalivedGroup is the Schema for the keepalivedgroups API",
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
				Description:         "KeepalivedGroupSpec defines the desired state of KeepalivedGroup",
				MarkdownDescription: "KeepalivedGroupSpec defines the desired state of KeepalivedGroup",
				Attributes: map[string]schema.Attribute{
					"blacklist_router_i_ds": schema.ListAttribute{
						Description:         "// +kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "// +kubebuilder:validation:UniqueItems=true",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"daemonset_pod_annotations": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"daemonset_pod_priority_class_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image": schema.StringAttribute{
						Description:         "//+kubebuilder:validation:Optional",
						MarkdownDescription: "//+kubebuilder:validation:Optional",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"interface": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"interface_from_ip": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"node_selector": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"password_auth": schema.SingleNestedAttribute{
						Description:         "PasswordAuth references a Kubernetes secret to extract the password for VRRP authentication",
						MarkdownDescription: "PasswordAuth references a Kubernetes secret to extract the password for VRRP authentication",
						Attributes: map[string]schema.Attribute{
							"secret_key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_ref": schema.SingleNestedAttribute{
								Description:         "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
								MarkdownDescription: "LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"unicast_enabled": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"verbatim_config": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
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
	}
}

func (r *RedhatcopRedhatIoKeepalivedGroupV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_redhatcop_redhat_io_keepalived_group_v1alpha1_manifest")

	var model RedhatcopRedhatIoKeepalivedGroupV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("redhatcop.redhat.io/v1alpha1")
	model.Kind = pointer.String("KeepalivedGroup")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
