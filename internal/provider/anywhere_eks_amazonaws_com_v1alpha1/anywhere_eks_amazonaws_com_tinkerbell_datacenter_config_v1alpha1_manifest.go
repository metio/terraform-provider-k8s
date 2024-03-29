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
	_ datasource.DataSource = &AnywhereEksAmazonawsComTinkerbellDatacenterConfigV1Alpha1Manifest{}
)

func NewAnywhereEksAmazonawsComTinkerbellDatacenterConfigV1Alpha1Manifest() datasource.DataSource {
	return &AnywhereEksAmazonawsComTinkerbellDatacenterConfigV1Alpha1Manifest{}
}

type AnywhereEksAmazonawsComTinkerbellDatacenterConfigV1Alpha1Manifest struct{}

type AnywhereEksAmazonawsComTinkerbellDatacenterConfigV1Alpha1ManifestData struct {
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
		HookImagesURLPath          *string `tfsdk:"hook_images_url_path" json:"hookImagesURLPath,omitempty"`
		OsImageURL                 *string `tfsdk:"os_image_url" json:"osImageURL,omitempty"`
		SkipLoadBalancerDeployment *bool   `tfsdk:"skip_load_balancer_deployment" json:"skipLoadBalancerDeployment,omitempty"`
		TinkerbellIP               *string `tfsdk:"tinkerbell_ip" json:"tinkerbellIP,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AnywhereEksAmazonawsComTinkerbellDatacenterConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_anywhere_eks_amazonaws_com_tinkerbell_datacenter_config_v1alpha1_manifest"
}

func (r *AnywhereEksAmazonawsComTinkerbellDatacenterConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TinkerbellDatacenterConfig is the Schema for the TinkerbellDatacenterConfigs API.",
		MarkdownDescription: "TinkerbellDatacenterConfig is the Schema for the TinkerbellDatacenterConfigs API.",
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
				Description:         "TinkerbellDatacenterConfigSpec defines the desired state of TinkerbellDatacenterConfig.",
				MarkdownDescription: "TinkerbellDatacenterConfigSpec defines the desired state of TinkerbellDatacenterConfig.",
				Attributes: map[string]schema.Attribute{
					"hook_images_url_path": schema.StringAttribute{
						Description:         "HookImagesURLPath can be used to override the default Hook images path to pull from a local server.",
						MarkdownDescription: "HookImagesURLPath can be used to override the default Hook images path to pull from a local server.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"os_image_url": schema.StringAttribute{
						Description:         "OSImageURL can be used to override the default OS image path to pull from a local server. OSImageURL is a URL to the OS image used during provisioning. To perform modular upgrades the OSImageURL must be specified on the TinkerbellMachineConfig objects. You cannot specify an OSImageURL on the TinkerbellDatacenterConfig and TinkerbellMachineConfigs simultaneously. It must include the Kubernetes version(s). For example, a URL used for Kubernetes 1.27 could be http://localhost:8080/ubuntu-2204-1.27.tgz",
						MarkdownDescription: "OSImageURL can be used to override the default OS image path to pull from a local server. OSImageURL is a URL to the OS image used during provisioning. To perform modular upgrades the OSImageURL must be specified on the TinkerbellMachineConfig objects. You cannot specify an OSImageURL on the TinkerbellDatacenterConfig and TinkerbellMachineConfigs simultaneously. It must include the Kubernetes version(s). For example, a URL used for Kubernetes 1.27 could be http://localhost:8080/ubuntu-2204-1.27.tgz",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"skip_load_balancer_deployment": schema.BoolAttribute{
						Description:         "SkipLoadBalancerDeployment when set to 'true' can be used to skip deploying a load balancer to expose Tinkerbell stack. Users will need to deploy and configure a load balancer manually after the cluster is created.",
						MarkdownDescription: "SkipLoadBalancerDeployment when set to 'true' can be used to skip deploying a load balancer to expose Tinkerbell stack. Users will need to deploy and configure a load balancer manually after the cluster is created.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tinkerbell_ip": schema.StringAttribute{
						Description:         "TinkerbellIP is used to configure a VIP for hosting the Tinkerbell services.",
						MarkdownDescription: "TinkerbellIP is used to configure a VIP for hosting the Tinkerbell services.",
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
	}
}

func (r *AnywhereEksAmazonawsComTinkerbellDatacenterConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_anywhere_eks_amazonaws_com_tinkerbell_datacenter_config_v1alpha1_manifest")

	var model AnywhereEksAmazonawsComTinkerbellDatacenterConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("anywhere.eks.amazonaws.com/v1alpha1")
	model.Kind = pointer.String("TinkerbellDatacenterConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
