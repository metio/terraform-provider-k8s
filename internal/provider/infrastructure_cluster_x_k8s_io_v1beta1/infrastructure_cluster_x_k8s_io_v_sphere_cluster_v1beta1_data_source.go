/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1beta1

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
	_ datasource.DataSource              = &InfrastructureClusterXK8SIoVsphereClusterV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &InfrastructureClusterXK8SIoVsphereClusterV1Beta1DataSource{}
)

func NewInfrastructureClusterXK8SIoVsphereClusterV1Beta1DataSource() datasource.DataSource {
	return &InfrastructureClusterXK8SIoVsphereClusterV1Beta1DataSource{}
}

type InfrastructureClusterXK8SIoVsphereClusterV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type InfrastructureClusterXK8SIoVsphereClusterV1Beta1DataSourceData struct {
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
		ClusterModules *[]struct {
			ControlPlane     *bool   `tfsdk:"control_plane" json:"controlPlane,omitempty"`
			ModuleUUID       *string `tfsdk:"module_uuid" json:"moduleUUID,omitempty"`
			TargetObjectName *string `tfsdk:"target_object_name" json:"targetObjectName,omitempty"`
		} `tfsdk:"cluster_modules" json:"clusterModules,omitempty"`
		ControlPlaneEndpoint *struct {
			Host *string `tfsdk:"host" json:"host,omitempty"`
			Port *int64  `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"control_plane_endpoint" json:"controlPlaneEndpoint,omitempty"`
		FailureDomainSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"failure_domain_selector" json:"failureDomainSelector,omitempty"`
		IdentityRef *struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"identity_ref" json:"identityRef,omitempty"`
		Server     *string `tfsdk:"server" json:"server,omitempty"`
		Thumbprint *string `tfsdk:"thumbprint" json:"thumbprint,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoVsphereClusterV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1beta1"
}

func (r *InfrastructureClusterXK8SIoVsphereClusterV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VSphereCluster is the Schema for the vsphereclusters API",
		MarkdownDescription: "VSphereCluster is the Schema for the vsphereclusters API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "VSphereClusterSpec defines the desired state of VSphereCluster",
				MarkdownDescription: "VSphereClusterSpec defines the desired state of VSphereCluster",
				Attributes: map[string]schema.Attribute{
					"cluster_modules": schema.ListNestedAttribute{
						Description:         "ClusterModules hosts information regarding the anti-affinity vSphere constructs for each of the objects responsible for creation of VM objects belonging to the cluster.",
						MarkdownDescription: "ClusterModules hosts information regarding the anti-affinity vSphere constructs for each of the objects responsible for creation of VM objects belonging to the cluster.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"control_plane": schema.BoolAttribute{
									Description:         "ControlPlane indicates whether the referred object is responsible for control plane nodes. Currently, only the KubeadmControlPlane objects have this flag set to true. Only a single object in the slice can have this value set to true.",
									MarkdownDescription: "ControlPlane indicates whether the referred object is responsible for control plane nodes. Currently, only the KubeadmControlPlane objects have this flag set to true. Only a single object in the slice can have this value set to true.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"module_uuid": schema.StringAttribute{
									Description:         "ModuleUUID is the unique identifier of the 'ClusterModule' used by the object.",
									MarkdownDescription: "ModuleUUID is the unique identifier of the 'ClusterModule' used by the object.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"target_object_name": schema.StringAttribute{
									Description:         "TargetObjectName points to the object that uses the Cluster Module information to enforce anti-affinity amongst its descendant VM objects.",
									MarkdownDescription: "TargetObjectName points to the object that uses the Cluster Module information to enforce anti-affinity amongst its descendant VM objects.",
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

					"control_plane_endpoint": schema.SingleNestedAttribute{
						Description:         "ControlPlaneEndpoint represents the endpoint used to communicate with the control plane.",
						MarkdownDescription: "ControlPlaneEndpoint represents the endpoint used to communicate with the control plane.",
						Attributes: map[string]schema.Attribute{
							"host": schema.StringAttribute{
								Description:         "The hostname on which the API server is serving.",
								MarkdownDescription: "The hostname on which the API server is serving.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"port": schema.Int64Attribute{
								Description:         "The port on which the API server is serving.",
								MarkdownDescription: "The port on which the API server is serving.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"failure_domain_selector": schema.SingleNestedAttribute{
						Description:         "FailureDomainSelector is the label selector to use for failure domain selection for the control plane nodes of the cluster. If not set ('nil'), selecting failure domains will be disabled. An empty value ('{}') selects all existing failure domains. A valid selector will select all failure domains which match the selector.",
						MarkdownDescription: "FailureDomainSelector is the label selector to use for failure domain selection for the control plane nodes of the cluster. If not set ('nil'), selecting failure domains will be disabled. An empty value ('{}') selects all existing failure domains. A valid selector will select all failure domains which match the selector.",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											ElementType:         types.StringType,
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

					"identity_ref": schema.SingleNestedAttribute{
						Description:         "IdentityRef is a reference to either a Secret or VSphereClusterIdentity that contains the identity to use when reconciling the cluster.",
						MarkdownDescription: "IdentityRef is a reference to either a Secret or VSphereClusterIdentity that contains the identity to use when reconciling the cluster.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind of the identity. Can either be VSphereClusterIdentity or Secret",
								MarkdownDescription: "Kind of the identity. Can either be VSphereClusterIdentity or Secret",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the identity.",
								MarkdownDescription: "Name of the identity.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"server": schema.StringAttribute{
						Description:         "Server is the address of the vSphere endpoint.",
						MarkdownDescription: "Server is the address of the vSphere endpoint.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"thumbprint": schema.StringAttribute{
						Description:         "Thumbprint is the colon-separated SHA-1 checksum of the given vCenter server's host certificate",
						MarkdownDescription: "Thumbprint is the colon-separated SHA-1 checksum of the given vCenter server's host certificate",
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

func (r *InfrastructureClusterXK8SIoVsphereClusterV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *InfrastructureClusterXK8SIoVsphereClusterV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1beta1")

	var data InfrastructureClusterXK8SIoVsphereClusterV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "infrastructure.cluster.x-k8s.io", Version: "v1beta1", Resource: "vsphereclusters"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse InfrastructureClusterXK8SIoVsphereClusterV1Beta1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1beta1")
	data.Kind = pointer.String("VSphereCluster")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
