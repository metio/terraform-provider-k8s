/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package authzed_com_v1alpha1

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
	_ datasource.DataSource = &AuthzedComSpiceDbclusterV1Alpha1Manifest{}
)

func NewAuthzedComSpiceDbclusterV1Alpha1Manifest() datasource.DataSource {
	return &AuthzedComSpiceDbclusterV1Alpha1Manifest{}
}

type AuthzedComSpiceDbclusterV1Alpha1Manifest struct{}

type AuthzedComSpiceDbclusterV1Alpha1ManifestData struct {
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
		Channel *string            `tfsdk:"channel" json:"channel,omitempty"`
		Config  *map[string]string `tfsdk:"config" json:"config,omitempty"`
		Patches *[]struct {
			Kind  *string            `tfsdk:"kind" json:"kind,omitempty"`
			Patch *map[string]string `tfsdk:"patch" json:"patch,omitempty"`
		} `tfsdk:"patches" json:"patches,omitempty"`
		SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
		Version    *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AuthzedComSpiceDbclusterV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_authzed_com_spice_db_cluster_v1alpha1_manifest"
}

func (r *AuthzedComSpiceDbclusterV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SpiceDBCluster defines all options for a full SpiceDB cluster",
		MarkdownDescription: "SpiceDBCluster defines all options for a full SpiceDB cluster",
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
				Description:         "ClusterSpec holds the desired state of the cluster.",
				MarkdownDescription: "ClusterSpec holds the desired state of the cluster.",
				Attributes: map[string]schema.Attribute{
					"channel": schema.StringAttribute{
						Description:         "Channel is a defined series of updates that operator should follow. The operator is configured with a datasource that configures available channels and update paths. If 'version' is not specified, then the operator will keep SpiceDB up-to-date with the current head of the channel. If 'version' is specified, then the operator will write available updates in the status.",
						MarkdownDescription: "Channel is a defined series of updates that operator should follow. The operator is configured with a datasource that configures available channels and update paths. If 'version' is not specified, then the operator will keep SpiceDB up-to-date with the current head of the channel. If 'version' is specified, then the operator will write available updates in the status.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"config": schema.MapAttribute{
						Description:         "Config values to be passed to the cluster",
						MarkdownDescription: "Config values to be passed to the cluster",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"patches": schema.ListNestedAttribute{
						Description:         "Patches is a list of patches to apply to generated resources. If multiple patches apply to the same object and field, later patches in the list take precedence over earlier ones.",
						MarkdownDescription: "Patches is a list of patches to apply to generated resources. If multiple patches apply to the same object and field, later patches in the list take precedence over earlier ones.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"kind": schema.StringAttribute{
									Description:         "Kind targets an object by its kubernetes Kind name.",
									MarkdownDescription: "Kind targets an object by its kubernetes Kind name.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"patch": schema.MapAttribute{
									Description:         "Patch is an inlined representation of a structured merge patch (one that just specifies the structure and fields to be modified) or a an explicit JSON6902 patch operation.",
									MarkdownDescription: "Patch is an inlined representation of a structured merge patch (one that just specifies the structure and fields to be modified) or a an explicit JSON6902 patch operation.",
									ElementType:         types.StringType,
									Required:            true,
									Optional:            false,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"secret_name": schema.StringAttribute{
						Description:         "SecretName points to a secret (in the same namespace) that holds secret config for the cluster like passwords, credentials, etc. If the secret is omitted, one will be generated",
						MarkdownDescription: "SecretName points to a secret (in the same namespace) that holds secret config for the cluster like passwords, credentials, etc. If the secret is omitted, one will be generated",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"version": schema.StringAttribute{
						Description:         "Version is the name of the version of SpiceDB that will be run. The version is usually a simple version string like 'v1.13.0', but the operator is configured with a data source that tells it what versions are allowed, and they may have other names. If omitted, the newest version in the head of the channel will be used. Note that the 'config.image' field will take precedence over version/channel, if it is specified",
						MarkdownDescription: "Version is the name of the version of SpiceDB that will be run. The version is usually a simple version string like 'v1.13.0', but the operator is configured with a data source that tells it what versions are allowed, and they may have other names. If omitted, the newest version in the head of the channel will be used. Note that the 'config.image' field will take precedence over version/channel, if it is specified",
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

func (r *AuthzedComSpiceDbclusterV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_authzed_com_spice_db_cluster_v1alpha1_manifest")

	var model AuthzedComSpiceDbclusterV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("authzed.com/v1alpha1")
	model.Kind = pointer.String("SpiceDBCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
