/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package console_openshift_io_v1

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
	_ datasource.DataSource = &ConsoleOpenshiftIoConsoleExternalLogLinkV1Manifest{}
)

func NewConsoleOpenshiftIoConsoleExternalLogLinkV1Manifest() datasource.DataSource {
	return &ConsoleOpenshiftIoConsoleExternalLogLinkV1Manifest{}
}

type ConsoleOpenshiftIoConsoleExternalLogLinkV1Manifest struct{}

type ConsoleOpenshiftIoConsoleExternalLogLinkV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		HrefTemplate    *string `tfsdk:"href_template" json:"hrefTemplate,omitempty"`
		NamespaceFilter *string `tfsdk:"namespace_filter" json:"namespaceFilter,omitempty"`
		Text            *string `tfsdk:"text" json:"text,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ConsoleOpenshiftIoConsoleExternalLogLinkV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_console_openshift_io_console_external_log_link_v1_manifest"
}

func (r *ConsoleOpenshiftIoConsoleExternalLogLinkV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ConsoleExternalLogLink is an extension for customizing OpenShift web console log links.  Compatibility level 2: Stable within a major release for a minimum of 9 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "ConsoleExternalLogLink is an extension for customizing OpenShift web console log links.  Compatibility level 2: Stable within a major release for a minimum of 9 months or 3 minor releases (whichever is longer).",
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
				Description:         "ConsoleExternalLogLinkSpec is the desired log link configuration. The log link will appear on the logs tab of the pod details page.",
				MarkdownDescription: "ConsoleExternalLogLinkSpec is the desired log link configuration. The log link will appear on the logs tab of the pod details page.",
				Attributes: map[string]schema.Attribute{
					"href_template": schema.StringAttribute{
						Description:         "hrefTemplate is an absolute secure URL (must use https) for the log link including variables to be replaced. Variables are specified in the URL with the format ${variableName}, for instance, ${containerName} and will be replaced with the corresponding values from the resource. Resource is a pod. Supported variables are: - ${resourceName} - name of the resource which containes the logs - ${resourceUID} - UID of the resource which contains the logs - e.g. '11111111-2222-3333-4444-555555555555' - ${containerName} - name of the resource's container that contains the logs - ${resourceNamespace} - namespace of the resource that contains the logs - ${resourceNamespaceUID} - namespace UID of the resource that contains the logs - ${podLabels} - JSON representation of labels matching the pod with the logs - e.g. '{'key1':'value1','key2':'value2'}'  e.g., https://example.com/logs?resourceName=${resourceName}&containerName=${containerName}&resourceNamespace=${resourceNamespace}&podLabels=${podLabels}",
						MarkdownDescription: "hrefTemplate is an absolute secure URL (must use https) for the log link including variables to be replaced. Variables are specified in the URL with the format ${variableName}, for instance, ${containerName} and will be replaced with the corresponding values from the resource. Resource is a pod. Supported variables are: - ${resourceName} - name of the resource which containes the logs - ${resourceUID} - UID of the resource which contains the logs - e.g. '11111111-2222-3333-4444-555555555555' - ${containerName} - name of the resource's container that contains the logs - ${resourceNamespace} - namespace of the resource that contains the logs - ${resourceNamespaceUID} - namespace UID of the resource that contains the logs - ${podLabels} - JSON representation of labels matching the pod with the logs - e.g. '{'key1':'value1','key2':'value2'}'  e.g., https://example.com/logs?resourceName=${resourceName}&containerName=${containerName}&resourceNamespace=${resourceNamespace}&podLabels=${podLabels}",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^https://`), ""),
						},
					},

					"namespace_filter": schema.StringAttribute{
						Description:         "namespaceFilter is a regular expression used to restrict a log link to a matching set of namespaces (e.g., '^openshift-'). The string is converted into a regular expression using the JavaScript RegExp constructor. If not specified, links will be displayed for all the namespaces.",
						MarkdownDescription: "namespaceFilter is a regular expression used to restrict a log link to a matching set of namespaces (e.g., '^openshift-'). The string is converted into a regular expression using the JavaScript RegExp constructor. If not specified, links will be displayed for all the namespaces.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"text": schema.StringAttribute{
						Description:         "text is the display text for the link",
						MarkdownDescription: "text is the display text for the link",
						Required:            true,
						Optional:            false,
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

func (r *ConsoleOpenshiftIoConsoleExternalLogLinkV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_console_openshift_io_console_external_log_link_v1_manifest")

	var model ConsoleOpenshiftIoConsoleExternalLogLinkV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("console.openshift.io/v1")
	model.Kind = pointer.String("ConsoleExternalLogLink")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
