/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuma_io_v1alpha1

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
	_ datasource.DataSource              = &KumaIoMeshGatewayConfigV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &KumaIoMeshGatewayConfigV1Alpha1DataSource{}
)

func NewKumaIoMeshGatewayConfigV1Alpha1DataSource() datasource.DataSource {
	return &KumaIoMeshGatewayConfigV1Alpha1DataSource{}
}

type KumaIoMeshGatewayConfigV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type KumaIoMeshGatewayConfigV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		CrossMesh   *bool `tfsdk:"cross_mesh" json:"crossMesh,omitempty"`
		PodTemplate *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec *struct {
				Container *struct {
					SecurityContext *struct {
						ReadOnlyRootFilesystem *bool `tfsdk:"read_only_root_filesystem" json:"readOnlyRootFilesystem,omitempty"`
					} `tfsdk:"security_context" json:"securityContext,omitempty"`
				} `tfsdk:"container" json:"container,omitempty"`
				SecurityContext *struct {
					FsGroup *int64 `tfsdk:"fs_group" json:"fsGroup,omitempty"`
				} `tfsdk:"security_context" json:"securityContext,omitempty"`
				ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"pod_template" json:"podTemplate,omitempty"`
		Replicas  *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
		Resources *struct {
			Claims *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"claims" json:"claims,omitempty"`
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		ServiceTemplate *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec *struct {
				LoadBalancerIP *string `tfsdk:"load_balancer_ip" json:"loadBalancerIP,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"service_template" json:"serviceTemplate,omitempty"`
		ServiceType *string            `tfsdk:"service_type" json:"serviceType,omitempty"`
		Tags        *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KumaIoMeshGatewayConfigV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuma_io_mesh_gateway_config_v1alpha1"
}

func (r *KumaIoMeshGatewayConfigV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "MeshGatewayConfig holds the configuration of a MeshGateway. A GatewayClass can refer to a MeshGatewayConfig via parametersRef.",
		MarkdownDescription: "MeshGatewayConfig holds the configuration of a MeshGateway. A GatewayClass can refer to a MeshGatewayConfig via parametersRef.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "MeshGatewayConfigSpec specifies the options available for a Kuma MeshGateway.",
				MarkdownDescription: "MeshGatewayConfigSpec specifies the options available for a Kuma MeshGateway.",
				Attributes: map[string]schema.Attribute{
					"cross_mesh": schema.BoolAttribute{
						Description:         "CrossMesh specifies whether listeners configured by this gateway are cross mesh listeners.",
						MarkdownDescription: "CrossMesh specifies whether listeners configured by this gateway are cross mesh listeners.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"pod_template": schema.SingleNestedAttribute{
						Description:         "PodTemplate configures the Pod owned by this config.",
						MarkdownDescription: "PodTemplate configures the Pod owned by this config.",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description:         "Metadata holds metadata configuration for a Service.",
								MarkdownDescription: "Metadata holds metadata configuration for a Service.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations holds annotations to be set on an object.",
										MarkdownDescription: "Annotations holds annotations to be set on an object.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels holds labels to be set on an objects.",
										MarkdownDescription: "Labels holds labels to be set on an objects.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"spec": schema.SingleNestedAttribute{
								Description:         "Spec holds some customizable fields of a Pod.",
								MarkdownDescription: "Spec holds some customizable fields of a Pod.",
								Attributes: map[string]schema.Attribute{
									"container": schema.SingleNestedAttribute{
										Description:         "Container corresponds to PodSpec.Container",
										MarkdownDescription: "Container corresponds to PodSpec.Container",
										Attributes: map[string]schema.Attribute{
											"security_context": schema.SingleNestedAttribute{
												Description:         "ContainerSecurityContext corresponds to PodSpec.Container.SecurityContext",
												MarkdownDescription: "ContainerSecurityContext corresponds to PodSpec.Container.SecurityContext",
												Attributes: map[string]schema.Attribute{
													"read_only_root_filesystem": schema.BoolAttribute{
														Description:         "ReadOnlyRootFilesystem corresponds to PodSpec.Container.SecurityContext.ReadOnlyRootFilesystem",
														MarkdownDescription: "ReadOnlyRootFilesystem corresponds to PodSpec.Container.SecurityContext.ReadOnlyRootFilesystem",
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
										Required: false,
										Optional: false,
										Computed: true,
									},

									"security_context": schema.SingleNestedAttribute{
										Description:         "PodSecurityContext corresponds to PodSpec.SecurityContext",
										MarkdownDescription: "PodSecurityContext corresponds to PodSpec.SecurityContext",
										Attributes: map[string]schema.Attribute{
											"fs_group": schema.Int64Attribute{
												Description:         "FSGroup corresponds to PodSpec.SecurityContext.FSGroup",
												MarkdownDescription: "FSGroup corresponds to PodSpec.SecurityContext.FSGroup",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"service_account_name": schema.StringAttribute{
										Description:         "ServiceAccountName corresponds to PodSpec.ServiceAccountName.",
										MarkdownDescription: "ServiceAccountName corresponds to PodSpec.ServiceAccountName.",
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
						Required: false,
						Optional: false,
						Computed: true,
					},

					"replicas": schema.Int64Attribute{
						Description:         "Replicas is the number of dataplane proxy replicas to create. For now this is a fixed number, but in the future it could be automatically scaled based on metrics.",
						MarkdownDescription: "Replicas is the number of dataplane proxy replicas to create. For now this is a fixed number, but in the future it could be automatically scaled based on metrics.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "Resources specifies the compute resources for the proxy container. The default can be set in the control plane config.",
						MarkdownDescription: "Resources specifies the compute resources for the proxy container. The default can be set in the control plane config.",
						Attributes: map[string]schema.Attribute{
							"claims": schema.ListNestedAttribute{
								Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
								MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
											MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"limits": schema.MapAttribute{
								Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"requests": schema.MapAttribute{
								Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"service_template": schema.SingleNestedAttribute{
						Description:         "ServiceTemplate configures the Service owned by this config.",
						MarkdownDescription: "ServiceTemplate configures the Service owned by this config.",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description:         "Metadata holds metadata configuration for a Service.",
								MarkdownDescription: "Metadata holds metadata configuration for a Service.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations holds annotations to be set on an object.",
										MarkdownDescription: "Annotations holds annotations to be set on an object.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels holds labels to be set on an objects.",
										MarkdownDescription: "Labels holds labels to be set on an objects.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"spec": schema.SingleNestedAttribute{
								Description:         "Spec holds some customizable fields of a Service.",
								MarkdownDescription: "Spec holds some customizable fields of a Service.",
								Attributes: map[string]schema.Attribute{
									"load_balancer_ip": schema.StringAttribute{
										Description:         "LoadBalancerIP corresponds to ServiceSpec.LoadBalancerIP.",
										MarkdownDescription: "LoadBalancerIP corresponds to ServiceSpec.LoadBalancerIP.",
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
						Required: false,
						Optional: false,
						Computed: true,
					},

					"service_type": schema.StringAttribute{
						Description:         "ServiceType specifies the type of managed Service that will be created to expose the dataplane proxies to traffic from outside the cluster. The ports to expose will be taken from the matching Gateway resource. If there is no matching Gateway, the managed Service will be deleted.",
						MarkdownDescription: "ServiceType specifies the type of managed Service that will be created to expose the dataplane proxies to traffic from outside the cluster. The ports to expose will be taken from the matching Gateway resource. If there is no matching Gateway, the managed Service will be deleted.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"tags": schema.MapAttribute{
						Description:         "Tags specifies a set of Kuma tags that are included in the MeshGatewayInstance and thus propagated to every Dataplane generated to serve the MeshGateway. These tags should include a maximum of one 'kuma.io/service' tag.",
						MarkdownDescription: "Tags specifies a set of Kuma tags that are included in the MeshGatewayInstance and thus propagated to every Dataplane generated to serve the MeshGateway. These tags should include a maximum of one 'kuma.io/service' tag.",
						ElementType:         types.StringType,
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

func (r *KumaIoMeshGatewayConfigV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *KumaIoMeshGatewayConfigV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_kuma_io_mesh_gateway_config_v1alpha1")

	var data KumaIoMeshGatewayConfigV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kuma.io", Version: "v1alpha1", Resource: "MeshGatewayConfig"}).
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

	var readResponse KumaIoMeshGatewayConfigV1Alpha1DataSourceData
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

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("kuma.io/v1alpha1")
	data.Kind = pointer.String("MeshGatewayConfig")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
