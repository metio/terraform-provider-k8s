/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package anywhere_eks_amazonaws_com_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &AnywhereEksAmazonawsComFluxConfigV1Alpha1Manifest{}
)

func NewAnywhereEksAmazonawsComFluxConfigV1Alpha1Manifest() datasource.DataSource {
	return &AnywhereEksAmazonawsComFluxConfigV1Alpha1Manifest{}
}

type AnywhereEksAmazonawsComFluxConfigV1Alpha1Manifest struct{}

type AnywhereEksAmazonawsComFluxConfigV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		Branch            *string `tfsdk:"branch" json:"branch,omitempty"`
		ClusterConfigPath *string `tfsdk:"cluster_config_path" json:"clusterConfigPath,omitempty"`
		Git               *struct {
			RepositoryUrl   *string `tfsdk:"repository_url" json:"repositoryUrl,omitempty"`
			SshKeyAlgorithm *string `tfsdk:"ssh_key_algorithm" json:"sshKeyAlgorithm,omitempty"`
		} `tfsdk:"git" json:"git,omitempty"`
		Github *struct {
			Owner      *string `tfsdk:"owner" json:"owner,omitempty"`
			Personal   *bool   `tfsdk:"personal" json:"personal,omitempty"`
			Repository *string `tfsdk:"repository" json:"repository,omitempty"`
		} `tfsdk:"github" json:"github,omitempty"`
		SystemNamespace *string `tfsdk:"system_namespace" json:"systemNamespace,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AnywhereEksAmazonawsComFluxConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_anywhere_eks_amazonaws_com_flux_config_v1alpha1_manifest"
}

func (r *AnywhereEksAmazonawsComFluxConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "FluxConfig is the Schema for the fluxconfigs API and defines the configurations of the Flux GitOps Toolkit and Git repository it links to.",
		MarkdownDescription: "FluxConfig is the Schema for the fluxconfigs API and defines the configurations of the Flux GitOps Toolkit and Git repository it links to.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "FluxConfigSpec defines the desired state of FluxConfig.",
				MarkdownDescription: "FluxConfigSpec defines the desired state of FluxConfig.",
				Attributes: map[string]schema.Attribute{
					"branch": schema.StringAttribute{
						Description:         "Git branch. Defaults to main.",
						MarkdownDescription: "Git branch. Defaults to main.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_config_path": schema.StringAttribute{
						Description:         "ClusterConfigPath relative to the repository root, when specified the cluster sync will be scoped to this path.",
						MarkdownDescription: "ClusterConfigPath relative to the repository root, when specified the cluster sync will be scoped to this path.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"git": schema.SingleNestedAttribute{
						Description:         "Used to specify Git provider that will be used to host the git files",
						MarkdownDescription: "Used to specify Git provider that will be used to host the git files",
						Attributes: map[string]schema.Attribute{
							"repository_url": schema.StringAttribute{
								Description:         "Repository URL for the repository to be used with flux. Can be either an SSH or HTTPS url.",
								MarkdownDescription: "Repository URL for the repository to be used with flux. Can be either an SSH or HTTPS url.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"ssh_key_algorithm": schema.StringAttribute{
								Description:         "SSH public key algorithm for the private key specified (rsa, ecdsa, ed25519) (default ecdsa)",
								MarkdownDescription: "SSH public key algorithm for the private key specified (rsa, ecdsa, ed25519) (default ecdsa)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"github": schema.SingleNestedAttribute{
						Description:         "Used to specify Github provider to host the Git repo and host the git files",
						MarkdownDescription: "Used to specify Github provider to host the Git repo and host the git files",
						Attributes: map[string]schema.Attribute{
							"owner": schema.StringAttribute{
								Description:         "Owner is the user or organization name of the Git provider.",
								MarkdownDescription: "Owner is the user or organization name of the Git provider.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"personal": schema.BoolAttribute{
								Description:         "if true, the owner is assumed to be a Git user; otherwise an org.",
								MarkdownDescription: "if true, the owner is assumed to be a Git user; otherwise an org.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"repository": schema.StringAttribute{
								Description:         "Repository name.",
								MarkdownDescription: "Repository name.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"system_namespace": schema.StringAttribute{
						Description:         "SystemNamespace scope for this operation. Defaults to flux-system",
						MarkdownDescription: "SystemNamespace scope for this operation. Defaults to flux-system",
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

func (r *AnywhereEksAmazonawsComFluxConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_anywhere_eks_amazonaws_com_flux_config_v1alpha1_manifest")

	var model AnywhereEksAmazonawsComFluxConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("anywhere.eks.amazonaws.com/v1alpha1")
	model.Kind = pointer.String("FluxConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
