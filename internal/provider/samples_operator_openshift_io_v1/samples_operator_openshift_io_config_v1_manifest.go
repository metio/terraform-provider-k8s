/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package samples_operator_openshift_io_v1

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
	_ datasource.DataSource = &SamplesOperatorOpenshiftIoConfigV1Manifest{}
)

func NewSamplesOperatorOpenshiftIoConfigV1Manifest() datasource.DataSource {
	return &SamplesOperatorOpenshiftIoConfigV1Manifest{}
}

type SamplesOperatorOpenshiftIoConfigV1Manifest struct{}

type SamplesOperatorOpenshiftIoConfigV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Architectures       *[]string `tfsdk:"architectures" json:"architectures,omitempty"`
		ManagementState     *string   `tfsdk:"management_state" json:"managementState,omitempty"`
		SamplesRegistry     *string   `tfsdk:"samples_registry" json:"samplesRegistry,omitempty"`
		SkippedHelmCharts   *[]string `tfsdk:"skipped_helm_charts" json:"skippedHelmCharts,omitempty"`
		SkippedImagestreams *[]string `tfsdk:"skipped_imagestreams" json:"skippedImagestreams,omitempty"`
		SkippedTemplates    *[]string `tfsdk:"skipped_templates" json:"skippedTemplates,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SamplesOperatorOpenshiftIoConfigV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_samples_operator_openshift_io_config_v1_manifest"
}

func (r *SamplesOperatorOpenshiftIoConfigV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Config contains the configuration and detailed condition status for the Samples Operator.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "Config contains the configuration and detailed condition status for the Samples Operator.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
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
				Description:         "ConfigSpec contains the desired configuration and state for the Samples Operator, controlling various behavior around the imagestreams and templates it creates/updates in the openshift namespace.",
				MarkdownDescription: "ConfigSpec contains the desired configuration and state for the Samples Operator, controlling various behavior around the imagestreams and templates it creates/updates in the openshift namespace.",
				Attributes: map[string]schema.Attribute{
					"architectures": schema.ListAttribute{
						Description:         "architectures determine which hardware architecture(s) to install, where x86_64, ppc64le, and s390x are the only supported choices currently.",
						MarkdownDescription: "architectures determine which hardware architecture(s) to install, where x86_64, ppc64le, and s390x are the only supported choices currently.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"management_state": schema.StringAttribute{
						Description:         "managementState is top level on/off type of switch for all operators. When 'Managed', this operator processes config and manipulates the samples accordingly. When 'Unmanaged', this operator ignores any updates to the resources it watches. When 'Removed', it reacts that same wasy as it does if the Config object is deleted, meaning any ImageStreams or Templates it manages (i.e. it honors the skipped lists) and the registry secret are deleted, along with the ConfigMap in the operator's namespace that represents the last config used to manipulate the samples,",
						MarkdownDescription: "managementState is top level on/off type of switch for all operators. When 'Managed', this operator processes config and manipulates the samples accordingly. When 'Unmanaged', this operator ignores any updates to the resources it watches. When 'Removed', it reacts that same wasy as it does if the Config object is deleted, meaning any ImageStreams or Templates it manages (i.e. it honors the skipped lists) and the registry secret are deleted, along with the ConfigMap in the operator's namespace that represents the last config used to manipulate the samples,",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(Managed|Unmanaged|Force|Removed)$`), ""),
						},
					},

					"samples_registry": schema.StringAttribute{
						Description:         "samplesRegistry allows for the specification of which registry is accessed by the ImageStreams for their image content.  Defaults on the content in https://github.com/openshift/library that are pulled into this github repository, but based on our pulling only ocp content it typically defaults to registry.redhat.io.",
						MarkdownDescription: "samplesRegistry allows for the specification of which registry is accessed by the ImageStreams for their image content.  Defaults on the content in https://github.com/openshift/library that are pulled into this github repository, but based on our pulling only ocp content it typically defaults to registry.redhat.io.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"skipped_helm_charts": schema.ListAttribute{
						Description:         "skippedHelmCharts specifies names of helm charts that should NOT be managed. Admins can use this to allow them to delete content they don’t want. They will still have to MANUALLY DELETE the content but the operator will not recreate(or update) anything listed here. Few examples of the name of helmcharts which can be skipped are 'redhat-redhat-perl-imagestreams','redhat-redhat-nodejs-imagestreams','redhat-nginx-imagestreams', 'redhat-redhat-ruby-imagestreams','redhat-redhat-python-imagestreams','redhat-redhat-php-imagestreams', 'redhat-httpd-imagestreams','redhat-redhat-dotnet-imagestreams'. Rest of the names can be obtained from openshift console --> helmcharts -->installed helmcharts. This will display the list of all the 12 helmcharts(of imagestreams)being installed by Samples Operator. The skippedHelmCharts must be a valid Kubernetes resource name. May contain only lowercase alphanumeric characters, hyphens and periods, and each period separated segment must begin and end with an alphanumeric character. It must be non-empty and at most 253 characters in length",
						MarkdownDescription: "skippedHelmCharts specifies names of helm charts that should NOT be managed. Admins can use this to allow them to delete content they don’t want. They will still have to MANUALLY DELETE the content but the operator will not recreate(or update) anything listed here. Few examples of the name of helmcharts which can be skipped are 'redhat-redhat-perl-imagestreams','redhat-redhat-nodejs-imagestreams','redhat-nginx-imagestreams', 'redhat-redhat-ruby-imagestreams','redhat-redhat-python-imagestreams','redhat-redhat-php-imagestreams', 'redhat-httpd-imagestreams','redhat-redhat-dotnet-imagestreams'. Rest of the names can be obtained from openshift console --> helmcharts -->installed helmcharts. This will display the list of all the 12 helmcharts(of imagestreams)being installed by Samples Operator. The skippedHelmCharts must be a valid Kubernetes resource name. May contain only lowercase alphanumeric characters, hyphens and periods, and each period separated segment must begin and end with an alphanumeric character. It must be non-empty and at most 253 characters in length",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"skipped_imagestreams": schema.ListAttribute{
						Description:         "skippedImagestreams specifies names of image streams that should NOT be created/updated.  Admins can use this to allow them to delete content they don’t want.  They will still have to manually delete the content but the operator will not recreate(or update) anything listed here.",
						MarkdownDescription: "skippedImagestreams specifies names of image streams that should NOT be created/updated.  Admins can use this to allow them to delete content they don’t want.  They will still have to manually delete the content but the operator will not recreate(or update) anything listed here.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"skipped_templates": schema.ListAttribute{
						Description:         "skippedTemplates specifies names of templates that should NOT be created/updated.  Admins can use this to allow them to delete content they don’t want.  They will still have to manually delete the content but the operator will not recreate(or update) anything listed here.",
						MarkdownDescription: "skippedTemplates specifies names of templates that should NOT be created/updated.  Admins can use this to allow them to delete content they don’t want.  They will still have to manually delete the content but the operator will not recreate(or update) anything listed here.",
						ElementType:         types.StringType,
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
	}
}

func (r *SamplesOperatorOpenshiftIoConfigV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_samples_operator_openshift_io_config_v1_manifest")

	var model SamplesOperatorOpenshiftIoConfigV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("samples.operator.openshift.io/v1")
	model.Kind = pointer.String("Config")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
