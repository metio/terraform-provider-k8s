/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package anywhere_eks_amazonaws_com_v1alpha1

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
	_ datasource.DataSource = &AnywhereEksAmazonawsComVsphereDatacenterConfigV1Alpha1Manifest{}
)

func NewAnywhereEksAmazonawsComVsphereDatacenterConfigV1Alpha1Manifest() datasource.DataSource {
	return &AnywhereEksAmazonawsComVsphereDatacenterConfigV1Alpha1Manifest{}
}

type AnywhereEksAmazonawsComVsphereDatacenterConfigV1Alpha1Manifest struct{}

type AnywhereEksAmazonawsComVsphereDatacenterConfigV1Alpha1ManifestData struct {
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
		Datacenter     *string `tfsdk:"datacenter" json:"datacenter,omitempty"`
		FailureDomains *[]struct {
			ComputeCluster *string `tfsdk:"compute_cluster" json:"computeCluster,omitempty"`
			Datastore      *string `tfsdk:"datastore" json:"datastore,omitempty"`
			Folder         *string `tfsdk:"folder" json:"folder,omitempty"`
			Name           *string `tfsdk:"name" json:"name,omitempty"`
			Network        *string `tfsdk:"network" json:"network,omitempty"`
			ResourcePool   *string `tfsdk:"resource_pool" json:"resourcePool,omitempty"`
		} `tfsdk:"failure_domains" json:"failureDomains,omitempty"`
		Insecure   *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
		Network    *string `tfsdk:"network" json:"network,omitempty"`
		Server     *string `tfsdk:"server" json:"server,omitempty"`
		Thumbprint *string `tfsdk:"thumbprint" json:"thumbprint,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AnywhereEksAmazonawsComVsphereDatacenterConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_anywhere_eks_amazonaws_com_v_sphere_datacenter_config_v1alpha1_manifest"
}

func (r *AnywhereEksAmazonawsComVsphereDatacenterConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VSphereDatacenterConfig is the Schema for the VSphereDatacenterConfigs API.",
		MarkdownDescription: "VSphereDatacenterConfig is the Schema for the VSphereDatacenterConfigs API.",
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
				Description:         "VSphereDatacenterConfigSpec defines the desired state of VSphereDatacenterConfig.",
				MarkdownDescription: "VSphereDatacenterConfigSpec defines the desired state of VSphereDatacenterConfig.",
				Attributes: map[string]schema.Attribute{
					"datacenter": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"failure_domains": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"compute_cluster": schema.StringAttribute{
									Description:         "ComputeCluster is the name or inventory path of the computecluster in which the VM is created/located",
									MarkdownDescription: "ComputeCluster is the name or inventory path of the computecluster in which the VM is created/located",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"datastore": schema.StringAttribute{
									Description:         "Datastore is the name or inventory path of the datastore in which the VM is created/located",
									MarkdownDescription: "Datastore is the name or inventory path of the datastore in which the VM is created/located",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"folder": schema.StringAttribute{
									Description:         "Folder is the name or inventory path of the folder in which the the VM is created/located",
									MarkdownDescription: "Folder is the name or inventory path of the folder in which the the VM is created/located",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is used as a unique identifier for each failure domain.",
									MarkdownDescription: "Name is used as a unique identifier for each failure domain.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"network": schema.StringAttribute{
									Description:         "Network is the name or inventory path of the network which will be added to the VM",
									MarkdownDescription: "Network is the name or inventory path of the network which will be added to the VM",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"resource_pool": schema.StringAttribute{
									Description:         "ResourcePool is the name or inventory path of the resource pool in which the VM is created/located",
									MarkdownDescription: "ResourcePool is the name or inventory path of the resource pool in which the VM is created/located",
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

					"insecure": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"network": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"server": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"thumbprint": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
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

func (r *AnywhereEksAmazonawsComVsphereDatacenterConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_anywhere_eks_amazonaws_com_v_sphere_datacenter_config_v1alpha1_manifest")

	var model AnywhereEksAmazonawsComVsphereDatacenterConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("anywhere.eks.amazonaws.com/v1alpha1")
	model.Kind = pointer.String("VSphereDatacenterConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
