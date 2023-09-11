/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_open_cluster_management_io_v1

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
	_ datasource.DataSource              = &OperatorOpenClusterManagementIoKlusterletV1DataSource{}
	_ datasource.DataSourceWithConfigure = &OperatorOpenClusterManagementIoKlusterletV1DataSource{}
)

func NewOperatorOpenClusterManagementIoKlusterletV1DataSource() datasource.DataSource {
	return &OperatorOpenClusterManagementIoKlusterletV1DataSource{}
}

type OperatorOpenClusterManagementIoKlusterletV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type OperatorOpenClusterManagementIoKlusterletV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ClusterName  *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
		DeployOption *struct {
			Mode *string `tfsdk:"mode" json:"mode,omitempty"`
		} `tfsdk:"deploy_option" json:"deployOption,omitempty"`
		ExternalServerURLs *[]struct {
			CaBundle *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
			Url      *string `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"external_server_urls" json:"externalServerURLs,omitempty"`
		HubApiServerHostAlias *struct {
			Hostname *string `tfsdk:"hostname" json:"hostname,omitempty"`
			Ip       *string `tfsdk:"ip" json:"ip,omitempty"`
		} `tfsdk:"hub_api_server_host_alias" json:"hubApiServerHostAlias,omitempty"`
		Namespace     *string `tfsdk:"namespace" json:"namespace,omitempty"`
		NodePlacement *struct {
			NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			Tolerations  *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
		} `tfsdk:"node_placement" json:"nodePlacement,omitempty"`
		RegistrationConfiguration *struct {
			ClientCertExpirationSeconds *int64 `tfsdk:"client_cert_expiration_seconds" json:"clientCertExpirationSeconds,omitempty"`
			FeatureGates                *[]struct {
				Feature *string `tfsdk:"feature" json:"feature,omitempty"`
				Mode    *string `tfsdk:"mode" json:"mode,omitempty"`
			} `tfsdk:"feature_gates" json:"featureGates,omitempty"`
		} `tfsdk:"registration_configuration" json:"registrationConfiguration,omitempty"`
		RegistrationImagePullSpec *string `tfsdk:"registration_image_pull_spec" json:"registrationImagePullSpec,omitempty"`
		WorkConfiguration         *struct {
			FeatureGates *[]struct {
				Feature *string `tfsdk:"feature" json:"feature,omitempty"`
				Mode    *string `tfsdk:"mode" json:"mode,omitempty"`
			} `tfsdk:"feature_gates" json:"featureGates,omitempty"`
		} `tfsdk:"work_configuration" json:"workConfiguration,omitempty"`
		WorkImagePullSpec *string `tfsdk:"work_image_pull_spec" json:"workImagePullSpec,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorOpenClusterManagementIoKlusterletV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_open_cluster_management_io_klusterlet_v1"
}

func (r *OperatorOpenClusterManagementIoKlusterletV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Klusterlet represents controllers to install the resources for a managed cluster. When configured, the Klusterlet requires a secret named bootstrap-hub-kubeconfig in the agent namespace to allow API requests to the hub for the registration protocol. In Hosted mode, the Klusterlet requires an additional secret named external-managed-kubeconfig in the agent namespace to allow API requests to the managed cluster for resources installation.",
		MarkdownDescription: "Klusterlet represents controllers to install the resources for a managed cluster. When configured, the Klusterlet requires a secret named bootstrap-hub-kubeconfig in the agent namespace to allow API requests to the hub for the registration protocol. In Hosted mode, the Klusterlet requires an additional secret named external-managed-kubeconfig in the agent namespace to allow API requests to the managed cluster for resources installation.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "Spec represents the desired deployment configuration of Klusterlet agent.",
				MarkdownDescription: "Spec represents the desired deployment configuration of Klusterlet agent.",
				Attributes: map[string]schema.Attribute{
					"cluster_name": schema.StringAttribute{
						Description:         "ClusterName is the name of the managed cluster to be created on hub. The Klusterlet agent generates a random name if it is not set, or discovers the appropriate cluster name on OpenShift.",
						MarkdownDescription: "ClusterName is the name of the managed cluster to be created on hub. The Klusterlet agent generates a random name if it is not set, or discovers the appropriate cluster name on OpenShift.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"deploy_option": schema.SingleNestedAttribute{
						Description:         "DeployOption contains the options of deploying a klusterlet",
						MarkdownDescription: "DeployOption contains the options of deploying a klusterlet",
						Attributes: map[string]schema.Attribute{
							"mode": schema.StringAttribute{
								Description:         "Mode can be Default or Hosted. It is Default mode if not specified In Default mode, all klusterlet related resources are deployed on the managed cluster. In Hosted mode, only crd and configurations are installed on the spoke/managed cluster. Controllers run in another cluster (defined as management-cluster) and connect to the mangaged cluster with the kubeconfig in secret of 'external-managed-kubeconfig'(a kubeconfig of managed-cluster with cluster-admin permission). Note: Do not modify the Mode field once it's applied.",
								MarkdownDescription: "Mode can be Default or Hosted. It is Default mode if not specified In Default mode, all klusterlet related resources are deployed on the managed cluster. In Hosted mode, only crd and configurations are installed on the spoke/managed cluster. Controllers run in another cluster (defined as management-cluster) and connect to the mangaged cluster with the kubeconfig in secret of 'external-managed-kubeconfig'(a kubeconfig of managed-cluster with cluster-admin permission). Note: Do not modify the Mode field once it's applied.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"external_server_urls": schema.ListNestedAttribute{
						Description:         "ExternalServerURLs represents the a list of apiserver urls and ca bundles that is accessible externally If it is set empty, managed cluster has no externally accessible url that hub cluster can visit.",
						MarkdownDescription: "ExternalServerURLs represents the a list of apiserver urls and ca bundles that is accessible externally If it is set empty, managed cluster has no externally accessible url that hub cluster can visit.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"ca_bundle": schema.StringAttribute{
									Description:         "CABundle is the ca bundle to connect to apiserver of the managed cluster. System certs are used if it is not set.",
									MarkdownDescription: "CABundle is the ca bundle to connect to apiserver of the managed cluster. System certs are used if it is not set.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"url": schema.StringAttribute{
									Description:         "URL is the url of apiserver endpoint of the managed cluster.",
									MarkdownDescription: "URL is the url of apiserver endpoint of the managed cluster.",
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

					"hub_api_server_host_alias": schema.SingleNestedAttribute{
						Description:         "HubApiServerHostAlias contains the host alias for hub api server. registration-agent and work-agent will use it to communicate with hub api server.",
						MarkdownDescription: "HubApiServerHostAlias contains the host alias for hub api server. registration-agent and work-agent will use it to communicate with hub api server.",
						Attributes: map[string]schema.Attribute{
							"hostname": schema.StringAttribute{
								Description:         "Hostname for the above IP address.",
								MarkdownDescription: "Hostname for the above IP address.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"ip": schema.StringAttribute{
								Description:         "IP address of the host file entry.",
								MarkdownDescription: "IP address of the host file entry.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"namespace": schema.StringAttribute{
						Description:         "Namespace is the namespace to deploy the agent on the managed cluster. The namespace must have a prefix of 'open-cluster-management-', and if it is not set, the namespace of 'open-cluster-management-agent' is used to deploy agent. In addition, the add-ons are deployed to the namespace of '{Namespace}-addon'. In the Hosted mode, this namespace still exists on the managed cluster to contain necessary resources, like service accounts, roles and rolebindings, while the agent is deployed to the namespace with the same name as klusterlet on the management cluster.",
						MarkdownDescription: "Namespace is the namespace to deploy the agent on the managed cluster. The namespace must have a prefix of 'open-cluster-management-', and if it is not set, the namespace of 'open-cluster-management-agent' is used to deploy agent. In addition, the add-ons are deployed to the namespace of '{Namespace}-addon'. In the Hosted mode, this namespace still exists on the managed cluster to contain necessary resources, like service accounts, roles and rolebindings, while the agent is deployed to the namespace with the same name as klusterlet on the management cluster.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"node_placement": schema.SingleNestedAttribute{
						Description:         "NodePlacement enables explicit control over the scheduling of the deployed pods.",
						MarkdownDescription: "NodePlacement enables explicit control over the scheduling of the deployed pods.",
						Attributes: map[string]schema.Attribute{
							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector defines which Nodes the Pods are scheduled on. The default is an empty list.",
								MarkdownDescription: "NodeSelector defines which Nodes the Pods are scheduled on. The default is an empty list.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "Tolerations is attached by pods to tolerate any taint that matches the triple <key,value,effect> using the matching operator <operator>. The default is an empty list.",
								MarkdownDescription: "Tolerations is attached by pods to tolerate any taint that matches the triple <key,value,effect> using the matching operator <operator>. The default is an empty list.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"effect": schema.StringAttribute{
											Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"key": schema.StringAttribute{
											Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"operator": schema.StringAttribute{
											Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"toleration_seconds": schema.Int64Attribute{
											Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
											MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"registration_configuration": schema.SingleNestedAttribute{
						Description:         "RegistrationConfiguration contains the configuration of registration",
						MarkdownDescription: "RegistrationConfiguration contains the configuration of registration",
						Attributes: map[string]schema.Attribute{
							"client_cert_expiration_seconds": schema.Int64Attribute{
								Description:         "clientCertExpirationSeconds represents the seconds of a client certificate to expire. If it is not set or 0, the default duration seconds will be set by the hub cluster. If the value is larger than the max signing duration seconds set on the hub cluster, the max signing duration seconds will be set.",
								MarkdownDescription: "clientCertExpirationSeconds represents the seconds of a client certificate to expire. If it is not set or 0, the default duration seconds will be set by the hub cluster. If the value is larger than the max signing duration seconds set on the hub cluster, the max signing duration seconds will be set.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"feature_gates": schema.ListNestedAttribute{
								Description:         "FeatureGates represents the list of feature gates for registration If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates: 1. If featuregate/Foo does not exist, registration-operator will discard it 2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true] 3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false, he can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
								MarkdownDescription: "FeatureGates represents the list of feature gates for registration If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates: 1. If featuregate/Foo does not exist, registration-operator will discard it 2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true] 3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false, he can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"feature": schema.StringAttribute{
											Description:         "Feature is the key of feature gate. e.g. featuregate/Foo.",
											MarkdownDescription: "Feature is the key of feature gate. e.g. featuregate/Foo.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mode": schema.StringAttribute{
											Description:         "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
											MarkdownDescription: "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"registration_image_pull_spec": schema.StringAttribute{
						Description:         "RegistrationImagePullSpec represents the desired image configuration of registration agent. quay.io/open-cluster-management.io/registration:latest will be used if unspecified.",
						MarkdownDescription: "RegistrationImagePullSpec represents the desired image configuration of registration agent. quay.io/open-cluster-management.io/registration:latest will be used if unspecified.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"work_configuration": schema.SingleNestedAttribute{
						Description:         "WorkConfiguration contains the configuration of work",
						MarkdownDescription: "WorkConfiguration contains the configuration of work",
						Attributes: map[string]schema.Attribute{
							"feature_gates": schema.ListNestedAttribute{
								Description:         "FeatureGates represents the list of feature gates for work If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates: 1. If featuregate/Foo does not exist, registration-operator will discard it 2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true] 3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false, he can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
								MarkdownDescription: "FeatureGates represents the list of feature gates for work If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates: 1. If featuregate/Foo does not exist, registration-operator will discard it 2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true] 3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false, he can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"feature": schema.StringAttribute{
											Description:         "Feature is the key of feature gate. e.g. featuregate/Foo.",
											MarkdownDescription: "Feature is the key of feature gate. e.g. featuregate/Foo.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mode": schema.StringAttribute{
											Description:         "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
											MarkdownDescription: "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"work_image_pull_spec": schema.StringAttribute{
						Description:         "WorkImagePullSpec represents the desired image configuration of work agent. quay.io/open-cluster-management.io/work:latest will be used if unspecified.",
						MarkdownDescription: "WorkImagePullSpec represents the desired image configuration of work agent. quay.io/open-cluster-management.io/work:latest will be used if unspecified.",
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

func (r *OperatorOpenClusterManagementIoKlusterletV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *OperatorOpenClusterManagementIoKlusterletV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_operator_open_cluster_management_io_klusterlet_v1")

	var data OperatorOpenClusterManagementIoKlusterletV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "operator.open-cluster-management.io", Version: "v1", Resource: "klusterlets"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetResourceError(err, data.Metadata.Name))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse OperatorOpenClusterManagementIoKlusterletV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("operator.open-cluster-management.io/v1")
	data.Kind = pointer.String("Klusterlet")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
