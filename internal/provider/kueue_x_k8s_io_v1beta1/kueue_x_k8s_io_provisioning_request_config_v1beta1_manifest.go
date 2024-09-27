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
	_ datasource.DataSource = &KueueXK8SIoProvisioningRequestConfigV1Beta1Manifest{}
)

func NewKueueXK8SIoProvisioningRequestConfigV1Beta1Manifest() datasource.DataSource {
	return &KueueXK8SIoProvisioningRequestConfigV1Beta1Manifest{}
}

type KueueXK8SIoProvisioningRequestConfigV1Beta1Manifest struct{}

type KueueXK8SIoProvisioningRequestConfigV1Beta1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ManagedResources      *[]string          `tfsdk:"managed_resources" json:"managedResources,omitempty"`
		Parameters            *map[string]string `tfsdk:"parameters" json:"parameters,omitempty"`
		ProvisioningClassName *string            `tfsdk:"provisioning_class_name" json:"provisioningClassName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KueueXK8SIoProvisioningRequestConfigV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kueue_x_k8s_io_provisioning_request_config_v1beta1_manifest"
}

func (r *KueueXK8SIoProvisioningRequestConfigV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ProvisioningRequestConfig is the Schema for the provisioningrequestconfig API",
		MarkdownDescription: "ProvisioningRequestConfig is the Schema for the provisioningrequestconfig API",
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
				Description:         "ProvisioningRequestConfigSpec defines the desired state of ProvisioningRequestConfig",
				MarkdownDescription: "ProvisioningRequestConfigSpec defines the desired state of ProvisioningRequestConfig",
				Attributes: map[string]schema.Attribute{
					"managed_resources": schema.ListAttribute{
						Description:         "managedResources contains the list of resources managed by the autoscaling. If empty, all resources are considered managed. If not empty, the ProvisioningRequest will contain only the podsets that are requesting at least one of them. If none of the workloads podsets is requesting at least a managed resource, the workload is considered ready.",
						MarkdownDescription: "managedResources contains the list of resources managed by the autoscaling. If empty, all resources are considered managed. If not empty, the ProvisioningRequest will contain only the podsets that are requesting at least one of them. If none of the workloads podsets is requesting at least a managed resource, the workload is considered ready.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"parameters": schema.MapAttribute{
						Description:         "Parameters contains all other parameters classes may require.",
						MarkdownDescription: "Parameters contains all other parameters classes may require.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"provisioning_class_name": schema.StringAttribute{
						Description:         "ProvisioningClassName describes the different modes of provisioning the resources. Check autoscaling.x-k8s.io ProvisioningRequestSpec.ProvisioningClassName for details.",
						MarkdownDescription: "ProvisioningClassName describes the different modes of provisioning the resources. Check autoscaling.x-k8s.io ProvisioningRequestSpec.ProvisioningClassName for details.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(253),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
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

func (r *KueueXK8SIoProvisioningRequestConfigV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kueue_x_k8s_io_provisioning_request_config_v1beta1_manifest")

	var model KueueXK8SIoProvisioningRequestConfigV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kueue.x-k8s.io/v1beta1")
	model.Kind = pointer.String("ProvisioningRequestConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
