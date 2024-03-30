/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package lb_lbconfig_carlosedp_com_v1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &LbLbconfigCarlosedpComExternalLoadBalancerV1Manifest{}
)

func NewLbLbconfigCarlosedpComExternalLoadBalancerV1Manifest() datasource.DataSource {
	return &LbLbconfigCarlosedpComExternalLoadBalancerV1Manifest{}
}

type LbLbconfigCarlosedpComExternalLoadBalancerV1Manifest struct{}

type LbLbconfigCarlosedpComExternalLoadBalancerV1ManifestData struct {
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
		Monitor *struct {
			Monitortype *string `tfsdk:"monitortype" json:"monitortype,omitempty"`
			Name        *string `tfsdk:"name" json:"name,omitempty"`
			Path        *string `tfsdk:"path" json:"path,omitempty"`
			Port        *int64  `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"monitor" json:"monitor,omitempty"`
		Nodelabels *map[string]string `tfsdk:"nodelabels" json:"nodelabels,omitempty"`
		Ports      *[]string          `tfsdk:"ports" json:"ports,omitempty"`
		Provider   *struct {
			Creds         *string `tfsdk:"creds" json:"creds,omitempty"`
			Debug         *bool   `tfsdk:"debug" json:"debug,omitempty"`
			Host          *string `tfsdk:"host" json:"host,omitempty"`
			Lbmethod      *string `tfsdk:"lbmethod" json:"lbmethod,omitempty"`
			Partition     *string `tfsdk:"partition" json:"partition,omitempty"`
			Port          *int64  `tfsdk:"port" json:"port,omitempty"`
			Validatecerts *bool   `tfsdk:"validatecerts" json:"validatecerts,omitempty"`
			Vendor        *string `tfsdk:"vendor" json:"vendor,omitempty"`
		} `tfsdk:"provider" json:"provider,omitempty"`
		Type *string `tfsdk:"type" json:"type,omitempty"`
		Vip  *string `tfsdk:"vip" json:"vip,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *LbLbconfigCarlosedpComExternalLoadBalancerV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_lb_lbconfig_carlosedp_com_external_load_balancer_v1_manifest"
}

func (r *LbLbconfigCarlosedpComExternalLoadBalancerV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ExternalLoadBalancer is the Schema for the externalloadbalancers API",
		MarkdownDescription: "ExternalLoadBalancer is the Schema for the externalloadbalancers API",
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
				Description:         "ExternalLoadBalancerSpec is the spec of a LoadBalancer instance.",
				MarkdownDescription: "ExternalLoadBalancerSpec is the spec of a LoadBalancer instance.",
				Attributes: map[string]schema.Attribute{
					"monitor": schema.SingleNestedAttribute{
						Description:         "Monitor is the path and port to monitor the LoadBalancer members",
						MarkdownDescription: "Monitor is the path and port to monitor the LoadBalancer members",
						Attributes: map[string]schema.Attribute{
							"monitortype": schema.StringAttribute{
								Description:         "MonitorType is the monitor parent type. <monitorType> must be one of 'http', 'https', 'icmp'.",
								MarkdownDescription: "MonitorType is the monitor parent type. <monitorType> must be one of 'http', 'https', 'icmp'.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("http", "https", "icmp"),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name is the monitor name, it is set by the controller",
								MarkdownDescription: "Name is the monitor name, it is set by the controller",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"path": schema.StringAttribute{
								Description:         "Path is the path URL to check for the pool members in the format '/healthz'",
								MarkdownDescription: "Path is the path URL to check for the pool members in the format '/healthz'",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"port": schema.Int64Attribute{
								Description:         "Port is the port this monitor should check the pool members",
								MarkdownDescription: "Port is the port this monitor should check the pool members",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"nodelabels": schema.MapAttribute{
						Description:         "NodeLabels are the node labels used for router sharding as an alternative to 'type'. Optional.",
						MarkdownDescription: "NodeLabels are the node labels used for router sharding as an alternative to 'type'. Optional.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ports": schema.ListAttribute{
						Description:         "Ports is the ports exposed by this LoadBalancer instance",
						MarkdownDescription: "Ports is the ports exposed by this LoadBalancer instance",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"provider": schema.SingleNestedAttribute{
						Description:         "Provider is the LoadBalancer backend provider",
						MarkdownDescription: "Provider is the LoadBalancer backend provider",
						Attributes: map[string]schema.Attribute{
							"creds": schema.StringAttribute{
								Description:         "Creds is the credentials secret holding the 'username' and 'password' keys. Generate with: 'kubectl create secret generic <secret-name> --from-literal=username=<username> --from-literal=password=<password>'",
								MarkdownDescription: "Creds is the credentials secret holding the 'username' and 'password' keys. Generate with: 'kubectl create secret generic <secret-name> --from-literal=username=<username> --from-literal=password=<password>'",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"debug": schema.BoolAttribute{
								Description:         "Debug is a flag to enable debug on the backend log output. Defaults to false.",
								MarkdownDescription: "Debug is a flag to enable debug on the backend log output. Defaults to false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host": schema.StringAttribute{
								Description:         "Host is the Load Balancer API IP or Hostname in URL format. Eg. 'http://10.25.10.10'.",
								MarkdownDescription: "Host is the Load Balancer API IP or Hostname in URL format. Eg. 'http://10.25.10.10'.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(255),
								},
							},

							"lbmethod": schema.StringAttribute{
								Description:         "Type is the Load-Balancing method. Defaults to 'round-robin'. Options are: ROUNDROBIN, LEASTCONNECTION, LEASTRESPONSETIME",
								MarkdownDescription: "Type is the Load-Balancing method. Defaults to 'round-robin'. Options are: ROUNDROBIN, LEASTCONNECTION, LEASTRESPONSETIME",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("ROUNDROBIN", "LEASTCONNECTION", "LEASTRESPONSETIME"),
								},
							},

							"partition": schema.StringAttribute{
								Description:         "Partition is the F5 partition to create the Load Balancer instances. Defaults to 'Common'. (F5 BigIP only)",
								MarkdownDescription: "Partition is the F5 partition to create the Load Balancer instances. Defaults to 'Common'. (F5 BigIP only)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "Port is the Load Balancer API Port.",
								MarkdownDescription: "Port is the Load Balancer API Port.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"validatecerts": schema.BoolAttribute{
								Description:         "ValidateCerts is a flag to validate or not the Load Balancer API certificate. Defaults to false.",
								MarkdownDescription: "ValidateCerts is a flag to validate or not the Load Balancer API certificate. Defaults to false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vendor": schema.StringAttribute{
								Description:         "Vendor is the backend provider vendor",
								MarkdownDescription: "Vendor is the backend provider vendor",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Dummy", "F5_BigIP", "Citrix_ADC", "HAProxy"),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"type": schema.StringAttribute{
						Description:         "Type is the node role type (master or infra) for the LoadBalancer instance",
						MarkdownDescription: "Type is the node role type (master or infra) for the LoadBalancer instance",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("master", "infra"),
						},
					},

					"vip": schema.StringAttribute{
						Description:         "Vip is the Virtual IP configured in  this LoadBalancer instance",
						MarkdownDescription: "Vip is the Virtual IP configured in  this LoadBalancer instance",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.LengthAtMost(15),
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

func (r *LbLbconfigCarlosedpComExternalLoadBalancerV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_lb_lbconfig_carlosedp_com_external_load_balancer_v1_manifest")

	var model LbLbconfigCarlosedpComExternalLoadBalancerV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("lb.lbconfig.carlosedp.com/v1")
	model.Kind = pointer.String("ExternalLoadBalancer")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
