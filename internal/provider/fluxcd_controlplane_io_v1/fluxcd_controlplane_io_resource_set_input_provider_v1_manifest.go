/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package fluxcd_controlplane_io_v1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &FluxcdControlplaneIoResourceSetInputProviderV1Manifest{}
)

func NewFluxcdControlplaneIoResourceSetInputProviderV1Manifest() datasource.DataSource {
	return &FluxcdControlplaneIoResourceSetInputProviderV1Manifest{}
}

type FluxcdControlplaneIoResourceSetInputProviderV1Manifest struct{}

type FluxcdControlplaneIoResourceSetInputProviderV1ManifestData struct {
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
		CertSecretRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"cert_secret_ref" json:"certSecretRef,omitempty"`
		DefaultValues *map[string]string `tfsdk:"default_values" json:"defaultValues,omitempty"`
		Filter        *struct {
			ExcludeBranch *string   `tfsdk:"exclude_branch" json:"excludeBranch,omitempty"`
			IncludeBranch *string   `tfsdk:"include_branch" json:"includeBranch,omitempty"`
			Labels        *[]string `tfsdk:"labels" json:"labels,omitempty"`
			Limit         *int64    `tfsdk:"limit" json:"limit,omitempty"`
		} `tfsdk:"filter" json:"filter,omitempty"`
		SecretRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		Skip *struct {
			Labels *[]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"skip" json:"skip,omitempty"`
		Type *string `tfsdk:"type" json:"type,omitempty"`
		Url  *string `tfsdk:"url" json:"url,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *FluxcdControlplaneIoResourceSetInputProviderV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_fluxcd_controlplane_io_resource_set_input_provider_v1_manifest"
}

func (r *FluxcdControlplaneIoResourceSetInputProviderV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ResourceSetInputProvider is the Schema for the ResourceSetInputProviders API.",
		MarkdownDescription: "ResourceSetInputProvider is the Schema for the ResourceSetInputProviders API.",
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
				Description:         "ResourceSetInputProviderSpec defines the desired state of ResourceSetInputProvider",
				MarkdownDescription: "ResourceSetInputProviderSpec defines the desired state of ResourceSetInputProvider",
				Attributes: map[string]schema.Attribute{
					"cert_secret_ref": schema.SingleNestedAttribute{
						Description:         "CertSecretRef specifies the Kubernetes Secret containing either or both of - a PEM-encoded CA certificate ('ca.crt') - a PEM-encoded client certificate ('tls.crt') and private key ('tls.key') When connecting to a Git provider that uses self-signed certificates, the CA certificate must be set in the Secret under the 'ca.crt' key to establish the trust relationship.",
						MarkdownDescription: "CertSecretRef specifies the Kubernetes Secret containing either or both of - a PEM-encoded CA certificate ('ca.crt') - a PEM-encoded client certificate ('tls.crt') and private key ('tls.key') When connecting to a Git provider that uses self-signed certificates, the CA certificate must be set in the Secret under the 'ca.crt' key to establish the trust relationship.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"default_values": schema.MapAttribute{
						Description:         "DefaultValues contains the default values for the inputs. These values are used to populate the inputs when the provider response does not contain them.",
						MarkdownDescription: "DefaultValues contains the default values for the inputs. These values are used to populate the inputs when the provider response does not contain them.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"filter": schema.SingleNestedAttribute{
						Description:         "Filter defines the filter to apply to the input provider response.",
						MarkdownDescription: "Filter defines the filter to apply to the input provider response.",
						Attributes: map[string]schema.Attribute{
							"exclude_branch": schema.StringAttribute{
								Description:         "ExcludeBranch specifies the regular expression to filter the branches that the input provider should exclude.",
								MarkdownDescription: "ExcludeBranch specifies the regular expression to filter the branches that the input provider should exclude.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"include_branch": schema.StringAttribute{
								Description:         "IncludeBranch specifies the regular expression to filter the branches that the input provider should include.",
								MarkdownDescription: "IncludeBranch specifies the regular expression to filter the branches that the input provider should include.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.ListAttribute{
								Description:         "Labels specifies the list of labels to filter the input provider response.",
								MarkdownDescription: "Labels specifies the list of labels to filter the input provider response.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"limit": schema.Int64Attribute{
								Description:         "Limit specifies the maximum number of input sets to return. When not set, the default limit is 100.",
								MarkdownDescription: "Limit specifies the maximum number of input sets to return. When not set, the default limit is 100.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"secret_ref": schema.SingleNestedAttribute{
						Description:         "SecretRef specifies the Kubernetes Secret containing the basic-auth credentials to access the input provider. The secret must contain the keys 'username' and 'password'. When connecting to a Git provider, the password should be a personal access token that grants read-only access to the repository.",
						MarkdownDescription: "SecretRef specifies the Kubernetes Secret containing the basic-auth credentials to access the input provider. The secret must contain the keys 'username' and 'password'. When connecting to a Git provider, the password should be a personal access token that grants read-only access to the repository.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"skip": schema.SingleNestedAttribute{
						Description:         "Skip defines whether we need to skip input provider response updates.",
						MarkdownDescription: "Skip defines whether we need to skip input provider response updates.",
						Attributes: map[string]schema.Attribute{
							"labels": schema.ListAttribute{
								Description:         "Labels specifies list of labels to skip input provider response when any of the label conditions matched. When prefixed with !, input provider response will be skipped if it does not have this label.",
								MarkdownDescription: "Labels specifies list of labels to skip input provider response when any of the label conditions matched. When prefixed with !, input provider response will be skipped if it does not have this label.",
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

					"type": schema.StringAttribute{
						Description:         "Type specifies the type of the input provider.",
						MarkdownDescription: "Type specifies the type of the input provider.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Static", "GitHubBranch", "GitHubPullRequest", "GitLabBranch", "GitLabMergeRequest"),
						},
					},

					"url": schema.StringAttribute{
						Description:         "URL specifies the HTTP/S address of the input provider API. When connecting to a Git provider, the URL should point to the repository address.",
						MarkdownDescription: "URL specifies the HTTP/S address of the input provider API. When connecting to a Git provider, the URL should point to the repository address.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^((http|https)://.*){0,1}$`), ""),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *FluxcdControlplaneIoResourceSetInputProviderV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_fluxcd_controlplane_io_resource_set_input_provider_v1_manifest")

	var model FluxcdControlplaneIoResourceSetInputProviderV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("fluxcd.controlplane.io/v1")
	model.Kind = pointer.String("ResourceSetInputProvider")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
