/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package lb_lbconfig_carlosedp_com_v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &LbLbconfigCarlosedpComExternalLoadBalancerV1DataSource{}
	_ datasource.DataSourceWithConfigure = &LbLbconfigCarlosedpComExternalLoadBalancerV1DataSource{}
)

func NewLbLbconfigCarlosedpComExternalLoadBalancerV1DataSource() datasource.DataSource {
	return &LbLbconfigCarlosedpComExternalLoadBalancerV1DataSource{}
}

type LbLbconfigCarlosedpComExternalLoadBalancerV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type LbLbconfigCarlosedpComExternalLoadBalancerV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *LbLbconfigCarlosedpComExternalLoadBalancerV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_lb_lbconfig_carlosedp_com_external_load_balancer_v1"
}

func (r *LbLbconfigCarlosedpComExternalLoadBalancerV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ExternalLoadBalancer is the Schema for the externalloadbalancers API",
		MarkdownDescription: "ExternalLoadBalancer is the Schema for the externalloadbalancers API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
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
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name is the monitor name, it is set by the controller",
								MarkdownDescription: "Name is the monitor name, it is set by the controller",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"path": schema.StringAttribute{
								Description:         "Path is the path URL to check for the pool members in the format '/healthz'",
								MarkdownDescription: "Path is the path URL to check for the pool members in the format '/healthz'",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"port": schema.Int64Attribute{
								Description:         "Port is the port this monitor should check the pool members",
								MarkdownDescription: "Port is the port this monitor should check the pool members",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"nodelabels": schema.MapAttribute{
						Description:         "NodeLabels are the node labels used for router sharding as an alternative to 'type'. Optional.",
						MarkdownDescription: "NodeLabels are the node labels used for router sharding as an alternative to 'type'. Optional.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ports": schema.ListAttribute{
						Description:         "Ports is the ports exposed by this LoadBalancer instance",
						MarkdownDescription: "Ports is the ports exposed by this LoadBalancer instance",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"provider": schema.SingleNestedAttribute{
						Description:         "Provider is the LoadBalancer backend provider",
						MarkdownDescription: "Provider is the LoadBalancer backend provider",
						Attributes: map[string]schema.Attribute{
							"creds": schema.StringAttribute{
								Description:         "Creds is the credentials secret holding the 'username' and 'password' keys. Generate with: 'kubectl create secret generic <secret-name> --from-literal=username=<username> --from-literal=password=<password>'",
								MarkdownDescription: "Creds is the credentials secret holding the 'username' and 'password' keys. Generate with: 'kubectl create secret generic <secret-name> --from-literal=username=<username> --from-literal=password=<password>'",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"debug": schema.BoolAttribute{
								Description:         "Debug is a flag to enable debug on the backend log output. Defaults to false.",
								MarkdownDescription: "Debug is a flag to enable debug on the backend log output. Defaults to false.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"host": schema.StringAttribute{
								Description:         "Host is the Load Balancer API IP or Hostname in URL format. Eg. 'http://10.25.10.10'.",
								MarkdownDescription: "Host is the Load Balancer API IP or Hostname in URL format. Eg. 'http://10.25.10.10'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"lbmethod": schema.StringAttribute{
								Description:         "Type is the Load-Balancing method. Defaults to 'round-robin'. Options are: ROUNDROBIN, LEASTCONNECTION, LEASTRESPONSETIME",
								MarkdownDescription: "Type is the Load-Balancing method. Defaults to 'round-robin'. Options are: ROUNDROBIN, LEASTCONNECTION, LEASTRESPONSETIME",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"partition": schema.StringAttribute{
								Description:         "Partition is the F5 partition to create the Load Balancer instances. Defaults to 'Common'. (F5 BigIP only)",
								MarkdownDescription: "Partition is the F5 partition to create the Load Balancer instances. Defaults to 'Common'. (F5 BigIP only)",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"port": schema.Int64Attribute{
								Description:         "Port is the Load Balancer API Port.",
								MarkdownDescription: "Port is the Load Balancer API Port.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"validatecerts": schema.BoolAttribute{
								Description:         "ValidateCerts is a flag to validate or not the Load Balancer API certificate. Defaults to false.",
								MarkdownDescription: "ValidateCerts is a flag to validate or not the Load Balancer API certificate. Defaults to false.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"vendor": schema.StringAttribute{
								Description:         "Vendor is the backend provider vendor",
								MarkdownDescription: "Vendor is the backend provider vendor",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"type": schema.StringAttribute{
						Description:         "Type is the node role type (master or infra) for the LoadBalancer instance",
						MarkdownDescription: "Type is the node role type (master or infra) for the LoadBalancer instance",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"vip": schema.StringAttribute{
						Description:         "Vip is the Virtual IP configured in  this LoadBalancer instance",
						MarkdownDescription: "Vip is the Virtual IP configured in  this LoadBalancer instance",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *LbLbconfigCarlosedpComExternalLoadBalancerV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *LbLbconfigCarlosedpComExternalLoadBalancerV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_lb_lbconfig_carlosedp_com_external_load_balancer_v1")

	var data LbLbconfigCarlosedpComExternalLoadBalancerV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "lb.lbconfig.carlosedp.com", Version: "v1", Resource: "externalloadbalancers"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse LbLbconfigCarlosedpComExternalLoadBalancerV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("lb.lbconfig.carlosedp.com/v1")
	data.Kind = pointer.String("ExternalLoadBalancer")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
