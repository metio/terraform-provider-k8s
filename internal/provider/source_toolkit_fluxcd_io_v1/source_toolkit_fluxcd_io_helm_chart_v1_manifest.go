/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package source_toolkit_fluxcd_io_v1

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
	_ datasource.DataSource = &SourceToolkitFluxcdIoHelmChartV1Manifest{}
)

func NewSourceToolkitFluxcdIoHelmChartV1Manifest() datasource.DataSource {
	return &SourceToolkitFluxcdIoHelmChartV1Manifest{}
}

type SourceToolkitFluxcdIoHelmChartV1Manifest struct{}

type SourceToolkitFluxcdIoHelmChartV1ManifestData struct {
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
		Chart                    *string `tfsdk:"chart" json:"chart,omitempty"`
		IgnoreMissingValuesFiles *bool   `tfsdk:"ignore_missing_values_files" json:"ignoreMissingValuesFiles,omitempty"`
		Interval                 *string `tfsdk:"interval" json:"interval,omitempty"`
		ReconcileStrategy        *string `tfsdk:"reconcile_strategy" json:"reconcileStrategy,omitempty"`
		SourceRef                *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"source_ref" json:"sourceRef,omitempty"`
		Suspend     *bool     `tfsdk:"suspend" json:"suspend,omitempty"`
		ValuesFiles *[]string `tfsdk:"values_files" json:"valuesFiles,omitempty"`
		Verify      *struct {
			MatchOIDCIdentity *[]struct {
				Issuer  *string `tfsdk:"issuer" json:"issuer,omitempty"`
				Subject *string `tfsdk:"subject" json:"subject,omitempty"`
			} `tfsdk:"match_oidc_identity" json:"matchOIDCIdentity,omitempty"`
			Provider  *string `tfsdk:"provider" json:"provider,omitempty"`
			SecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"verify" json:"verify,omitempty"`
		Version *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SourceToolkitFluxcdIoHelmChartV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_source_toolkit_fluxcd_io_helm_chart_v1_manifest"
}

func (r *SourceToolkitFluxcdIoHelmChartV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "HelmChart is the Schema for the helmcharts API.",
		MarkdownDescription: "HelmChart is the Schema for the helmcharts API.",
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
				Description:         "HelmChartSpec specifies the desired state of a Helm chart.",
				MarkdownDescription: "HelmChartSpec specifies the desired state of a Helm chart.",
				Attributes: map[string]schema.Attribute{
					"chart": schema.StringAttribute{
						Description:         "Chart is the name or path the Helm chart is available at in theSourceRef.",
						MarkdownDescription: "Chart is the name or path the Helm chart is available at in theSourceRef.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"ignore_missing_values_files": schema.BoolAttribute{
						Description:         "IgnoreMissingValuesFiles controls whether to silently ignore missing valuesfiles rather than failing.",
						MarkdownDescription: "IgnoreMissingValuesFiles controls whether to silently ignore missing valuesfiles rather than failing.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"interval": schema.StringAttribute{
						Description:         "Interval at which the HelmChart SourceRef is checked for updates.This interval is approximate and may be subject to jitter to ensureefficient use of resources.",
						MarkdownDescription: "Interval at which the HelmChart SourceRef is checked for updates.This interval is approximate and may be subject to jitter to ensureefficient use of resources.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m|h))+$`), ""),
						},
					},

					"reconcile_strategy": schema.StringAttribute{
						Description:         "ReconcileStrategy determines what enables the creation of a new artifact.Valid values are ('ChartVersion', 'Revision').See the documentation of the values for an explanation on their behavior.Defaults to ChartVersion when omitted.",
						MarkdownDescription: "ReconcileStrategy determines what enables the creation of a new artifact.Valid values are ('ChartVersion', 'Revision').See the documentation of the values for an explanation on their behavior.Defaults to ChartVersion when omitted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("ChartVersion", "Revision"),
						},
					},

					"source_ref": schema.SingleNestedAttribute{
						Description:         "SourceRef is the reference to the Source the chart is available at.",
						MarkdownDescription: "SourceRef is the reference to the Source the chart is available at.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "APIVersion of the referent.",
								MarkdownDescription: "APIVersion of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent, valid values are ('HelmRepository', 'GitRepository','Bucket').",
								MarkdownDescription: "Kind of the referent, valid values are ('HelmRepository', 'GitRepository','Bucket').",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("HelmRepository", "GitRepository", "Bucket"),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"suspend": schema.BoolAttribute{
						Description:         "Suspend tells the controller to suspend the reconciliation of thissource.",
						MarkdownDescription: "Suspend tells the controller to suspend the reconciliation of thissource.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"values_files": schema.ListAttribute{
						Description:         "ValuesFiles is an alternative list of values files to use as the chartvalues (values.yaml is not included by default), expected to be arelative path in the SourceRef.Values files are merged in the order of this list with the last fileoverriding the first. Ignored when omitted.",
						MarkdownDescription: "ValuesFiles is an alternative list of values files to use as the chartvalues (values.yaml is not included by default), expected to be arelative path in the SourceRef.Values files are merged in the order of this list with the last fileoverriding the first. Ignored when omitted.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"verify": schema.SingleNestedAttribute{
						Description:         "Verify contains the secret name containing the trusted public keysused to verify the signature and specifies which provider to use to checkwhether OCI image is authentic.This field is only supported when using HelmRepository source with spec.type 'oci'.Chart dependencies, which are not bundled in the umbrella chart artifact, are not verified.",
						MarkdownDescription: "Verify contains the secret name containing the trusted public keysused to verify the signature and specifies which provider to use to checkwhether OCI image is authentic.This field is only supported when using HelmRepository source with spec.type 'oci'.Chart dependencies, which are not bundled in the umbrella chart artifact, are not verified.",
						Attributes: map[string]schema.Attribute{
							"match_oidc_identity": schema.ListNestedAttribute{
								Description:         "MatchOIDCIdentity specifies the identity matching criteria to usewhile verifying an OCI artifact which was signed using Cosign keylesssigning. The artifact's identity is deemed to be verified if any of thespecified matchers match against the identity.",
								MarkdownDescription: "MatchOIDCIdentity specifies the identity matching criteria to usewhile verifying an OCI artifact which was signed using Cosign keylesssigning. The artifact's identity is deemed to be verified if any of thespecified matchers match against the identity.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"issuer": schema.StringAttribute{
											Description:         "Issuer specifies the regex pattern to match against to verifythe OIDC issuer in the Fulcio certificate. The pattern must be avalid Go regular expression.",
											MarkdownDescription: "Issuer specifies the regex pattern to match against to verifythe OIDC issuer in the Fulcio certificate. The pattern must be avalid Go regular expression.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"subject": schema.StringAttribute{
											Description:         "Subject specifies the regex pattern to match against to verifythe identity subject in the Fulcio certificate. The pattern mustbe a valid Go regular expression.",
											MarkdownDescription: "Subject specifies the regex pattern to match against to verifythe identity subject in the Fulcio certificate. The pattern mustbe a valid Go regular expression.",
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

							"provider": schema.StringAttribute{
								Description:         "Provider specifies the technology used to sign the OCI Artifact.",
								MarkdownDescription: "Provider specifies the technology used to sign the OCI Artifact.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("cosign", "notation"),
								},
							},

							"secret_ref": schema.SingleNestedAttribute{
								Description:         "SecretRef specifies the Kubernetes Secret containing thetrusted public keys.",
								MarkdownDescription: "SecretRef specifies the Kubernetes Secret containing thetrusted public keys.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"version": schema.StringAttribute{
						Description:         "Version is the chart version semver expression, ignored for charts fromGitRepository and Bucket sources. Defaults to latest when omitted.",
						MarkdownDescription: "Version is the chart version semver expression, ignored for charts fromGitRepository and Bucket sources. Defaults to latest when omitted.",
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

func (r *SourceToolkitFluxcdIoHelmChartV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_source_toolkit_fluxcd_io_helm_chart_v1_manifest")

	var model SourceToolkitFluxcdIoHelmChartV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("source.toolkit.fluxcd.io/v1")
	model.Kind = pointer.String("HelmChart")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
