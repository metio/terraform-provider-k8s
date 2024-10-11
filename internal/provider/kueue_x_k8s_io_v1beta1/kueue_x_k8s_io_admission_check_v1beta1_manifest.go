/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kueue_x_k8s_io_v1beta1

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
	_ datasource.DataSource = &KueueXK8SIoAdmissionCheckV1Beta1Manifest{}
)

func NewKueueXK8SIoAdmissionCheckV1Beta1Manifest() datasource.DataSource {
	return &KueueXK8SIoAdmissionCheckV1Beta1Manifest{}
}

type KueueXK8SIoAdmissionCheckV1Beta1Manifest struct{}

type KueueXK8SIoAdmissionCheckV1Beta1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ControllerName *string `tfsdk:"controller_name" json:"controllerName,omitempty"`
		Parameters     *struct {
			ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
			Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"parameters" json:"parameters,omitempty"`
		RetryDelayMinutes *int64 `tfsdk:"retry_delay_minutes" json:"retryDelayMinutes,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KueueXK8SIoAdmissionCheckV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kueue_x_k8s_io_admission_check_v1beta1_manifest"
}

func (r *KueueXK8SIoAdmissionCheckV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AdmissionCheck is the Schema for the admissionchecks API",
		MarkdownDescription: "AdmissionCheck is the Schema for the admissionchecks API",
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
				Description:         "AdmissionCheckSpec defines the desired state of AdmissionCheck",
				MarkdownDescription: "AdmissionCheckSpec defines the desired state of AdmissionCheck",
				Attributes: map[string]schema.Attribute{
					"controller_name": schema.StringAttribute{
						Description:         "controllerName identifies the controller that processes the AdmissionCheck, not necessarily a Kubernetes Pod or Deployment name. Cannot be empty.",
						MarkdownDescription: "controllerName identifies the controller that processes the AdmissionCheck, not necessarily a Kubernetes Pod or Deployment name. Cannot be empty.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"parameters": schema.SingleNestedAttribute{
						Description:         "Parameters identifies a configuration with additional parameters for the check.",
						MarkdownDescription: "Parameters identifies a configuration with additional parameters for the check.",
						Attributes: map[string]schema.Attribute{
							"api_group": schema.StringAttribute{
								Description:         "ApiGroup is the group for the resource being referenced.",
								MarkdownDescription: "ApiGroup is the group for the resource being referenced.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
								},
							},

							"kind": schema.StringAttribute{
								Description:         "Kind is the type of the resource being referenced.",
								MarkdownDescription: "Kind is the type of the resource being referenced.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)[a-z]([-a-z0-9]*[a-z0-9])?$`), ""),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of the resource being referenced.",
								MarkdownDescription: "Name is the name of the resource being referenced.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"retry_delay_minutes": schema.Int64Attribute{
						Description:         "RetryDelayMinutes specifies how long to keep the workload suspended after a failed check (after it transitioned to False). When the delay period has passed, the check state goes to 'Unknown'. The default is 15 min. Deprecated: retryDelayMinutes has already been deprecated since v0.8 and will be removed in v1beta2.",
						MarkdownDescription: "RetryDelayMinutes specifies how long to keep the workload suspended after a failed check (after it transitioned to False). When the delay period has passed, the check state goes to 'Unknown'. The default is 15 min. Deprecated: retryDelayMinutes has already been deprecated since v0.8 and will be removed in v1beta2.",
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

func (r *KueueXK8SIoAdmissionCheckV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kueue_x_k8s_io_admission_check_v1beta1_manifest")

	var model KueueXK8SIoAdmissionCheckV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kueue.x-k8s.io/v1beta1")
	model.Kind = pointer.String("AdmissionCheck")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
