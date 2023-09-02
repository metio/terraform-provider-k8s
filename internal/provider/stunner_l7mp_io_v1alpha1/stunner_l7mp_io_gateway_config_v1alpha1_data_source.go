/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package stunner_l7mp_io_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &StunnerL7MpIoGatewayConfigV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &StunnerL7MpIoGatewayConfigV1Alpha1DataSource{}
)

func NewStunnerL7MpIoGatewayConfigV1Alpha1DataSource() datasource.DataSource {
	return &StunnerL7MpIoGatewayConfigV1Alpha1DataSource{}
}

type StunnerL7MpIoGatewayConfigV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type StunnerL7MpIoGatewayConfigV1Alpha1DataSourceData struct {
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
		AuthLifetime *int64 `tfsdk:"auth_lifetime" json:"authLifetime,omitempty"`
		AuthRef      *struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"auth_ref" json:"authRef,omitempty"`
		AuthType                       *string            `tfsdk:"auth_type" json:"authType,omitempty"`
		HealthCheckEndpoint            *string            `tfsdk:"health_check_endpoint" json:"healthCheckEndpoint,omitempty"`
		LoadBalancerServiceAnnotations *map[string]string `tfsdk:"load_balancer_service_annotations" json:"loadBalancerServiceAnnotations,omitempty"`
		LogLevel                       *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
		MaxPort                        *int64             `tfsdk:"max_port" json:"maxPort,omitempty"`
		MetricsEndpoint                *string            `tfsdk:"metrics_endpoint" json:"metricsEndpoint,omitempty"`
		MinPort                        *int64             `tfsdk:"min_port" json:"minPort,omitempty"`
		Password                       *string            `tfsdk:"password" json:"password,omitempty"`
		Realm                          *string            `tfsdk:"realm" json:"realm,omitempty"`
		SharedSecret                   *string            `tfsdk:"shared_secret" json:"sharedSecret,omitempty"`
		StunnerConfig                  *string            `tfsdk:"stunner_config" json:"stunnerConfig,omitempty"`
		UserName                       *string            `tfsdk:"user_name" json:"userName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *StunnerL7MpIoGatewayConfigV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_stunner_l7mp_io_gateway_config_v1alpha1"
}

func (r *StunnerL7MpIoGatewayConfigV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "GatewayConfig is the Schema for the gatewayconfigs API",
		MarkdownDescription: "GatewayConfig is the Schema for the gatewayconfigs API",
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
				Description:         "GatewayConfigSpec defines the desired state of GatewayConfig",
				MarkdownDescription: "GatewayConfigSpec defines the desired state of GatewayConfig",
				Attributes: map[string]schema.Attribute{
					"auth_lifetime": schema.Int64Attribute{
						Description:         "AuthLifetime defines the lifetime of 'longterm' authentication credentials in seconds.",
						MarkdownDescription: "AuthLifetime defines the lifetime of 'longterm' authentication credentials in seconds.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"auth_ref": schema.SingleNestedAttribute{
						Description:         "Note that externally set credentials override any inline auth credentials (AuthType, AuthUsername, etc.): if AuthRef is nonempty then it is expected that the referenced Secret exists and *all* authentication credentials are correctly set in the referenced Secret (username/password or shared secret). Mixing of credential sources (inline/external) is not supported.",
						MarkdownDescription: "Note that externally set credentials override any inline auth credentials (AuthType, AuthUsername, etc.): if AuthRef is nonempty then it is expected that the referenced Secret exists and *all* authentication credentials are correctly set in the referenced Secret (username/password or shared secret). Mixing of credential sources (inline/external) is not supported.",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "Group is the group of the referent. For example, 'gateway.networking.k8s.io'. When unspecified or empty string, core API group is inferred.",
								MarkdownDescription: "Group is the group of the referent. For example, 'gateway.networking.k8s.io'. When unspecified or empty string, core API group is inferred.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",
								MarkdownDescription: "Kind is kind of the referent. For example 'HTTPRoute' or 'Service'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of the referent.",
								MarkdownDescription: "Name is the name of the referent.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace is the namespace of the backend. When unspecified, the local namespace is inferred.  Note that when a namespace is specified, a ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.  Support: Core",
								MarkdownDescription: "Namespace is the namespace of the backend. When unspecified, the local namespace is inferred.  Note that when a namespace is specified, a ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.  Support: Core",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"auth_type": schema.StringAttribute{
						Description:         "AuthType is the type of the STUN/TURN authentication mechanism.",
						MarkdownDescription: "AuthType is the type of the STUN/TURN authentication mechanism.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"health_check_endpoint": schema.StringAttribute{
						Description:         "HealthCheckEndpoint is the URI of the form 'http://address:port' exposed for external HTTP health-checking. A liveness probe responder will be exposed on path '/live' and readiness probe on path '/ready'. The scheme ('http://') is mandatory, default is to enable health-checking at 'http://0.0.0.0:8086'.",
						MarkdownDescription: "HealthCheckEndpoint is the URI of the form 'http://address:port' exposed for external HTTP health-checking. A liveness probe responder will be exposed on path '/live' and readiness probe on path '/ready'. The scheme ('http://') is mandatory, default is to enable health-checking at 'http://0.0.0.0:8086'.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"load_balancer_service_annotations": schema.MapAttribute{
						Description:         "LoadBalancerServiceAnnotations is a list of annotations that will go into the LoadBalancer services created automatically by the operator to wrap Gateways.  NOTE: removing annotations from a GatewayConfig will not result in the removal of the corresponding annotations from the LoadBalancer service, in order to prevent the accidental removal of an annotation installed there by Kubernetes or the cloud provider. If you really want to remove an annotation, do this manually or simply remove all Gateways (which will remove the corresponding LoadBalancer services), update the GatewayConfig and then recreate the Gateways, so that the newly created LoadBalancer services will contain the required annotations.",
						MarkdownDescription: "LoadBalancerServiceAnnotations is a list of annotations that will go into the LoadBalancer services created automatically by the operator to wrap Gateways.  NOTE: removing annotations from a GatewayConfig will not result in the removal of the corresponding annotations from the LoadBalancer service, in order to prevent the accidental removal of an annotation installed there by Kubernetes or the cloud provider. If you really want to remove an annotation, do this manually or simply remove all Gateways (which will remove the corresponding LoadBalancer services), update the GatewayConfig and then recreate the Gateways, so that the newly created LoadBalancer services will contain the required annotations.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"log_level": schema.StringAttribute{
						Description:         "LogLevel specifies the default loglevel for the STUNner daemon.",
						MarkdownDescription: "LogLevel specifies the default loglevel for the STUNner daemon.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"max_port": schema.Int64Attribute{
						Description:         "MaxRelayPort is the smallest relay port assigned for STUNner relay connections.",
						MarkdownDescription: "MaxRelayPort is the smallest relay port assigned for STUNner relay connections.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"metrics_endpoint": schema.StringAttribute{
						Description:         "MetricsEndpoint is the URI in the form 'http://address:port/path' exposed for metric scraping (Prometheus). The scheme ('http://') is mandatory. Default is to expose no metric endpoint.",
						MarkdownDescription: "MetricsEndpoint is the URI in the form 'http://address:port/path' exposed for metric scraping (Prometheus). The scheme ('http://') is mandatory. Default is to expose no metric endpoint.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"min_port": schema.Int64Attribute{
						Description:         "MinRelayPort is the smallest relay port assigned for STUNner relay connections.",
						MarkdownDescription: "MinRelayPort is the smallest relay port assigned for STUNner relay connections.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"password": schema.StringAttribute{
						Description:         "Password defines the 'password' credential for 'plaintext' authentication.",
						MarkdownDescription: "Password defines the 'password' credential for 'plaintext' authentication.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"realm": schema.StringAttribute{
						Description:         "Realm defines the STUN/TURN authentication realm to be used for clients toauthenticate with STUNner.  The realm must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character. No other punctuation is allowed.",
						MarkdownDescription: "Realm defines the STUN/TURN authentication realm to be used for clients toauthenticate with STUNner.  The realm must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character. No other punctuation is allowed.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"shared_secret": schema.StringAttribute{
						Description:         "SharedSecret defines the shared secret to be used for 'longterm' authentication.",
						MarkdownDescription: "SharedSecret defines the shared secret to be used for 'longterm' authentication.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"stunner_config": schema.StringAttribute{
						Description:         "StunnerConfig specifies the name of the ConfigMap into which the operator renders the stunnerd configfile.",
						MarkdownDescription: "StunnerConfig specifies the name of the ConfigMap into which the operator renders the stunnerd configfile.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"user_name": schema.StringAttribute{
						Description:         "Username defines the 'username' credential for 'plaintext' authentication.",
						MarkdownDescription: "Username defines the 'username' credential for 'plaintext' authentication.",
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

func (r *StunnerL7MpIoGatewayConfigV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *StunnerL7MpIoGatewayConfigV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_stunner_l7mp_io_gateway_config_v1alpha1")

	var data StunnerL7MpIoGatewayConfigV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "stunner.l7mp.io", Version: "v1alpha1", Resource: "GatewayConfig"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
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

	var readResponse StunnerL7MpIoGatewayConfigV1Alpha1DataSourceData
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
	data.ApiVersion = pointer.String("stunner.l7mp.io/v1alpha1")
	data.Kind = pointer.String("GatewayConfig")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}