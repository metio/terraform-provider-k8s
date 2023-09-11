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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"regexp"
)

var (
	_ resource.Resource                = &OperatorOpenClusterManagementIoKlusterletV1Resource{}
	_ resource.ResourceWithConfigure   = &OperatorOpenClusterManagementIoKlusterletV1Resource{}
	_ resource.ResourceWithImportState = &OperatorOpenClusterManagementIoKlusterletV1Resource{}
)

func NewOperatorOpenClusterManagementIoKlusterletV1Resource() resource.Resource {
	return &OperatorOpenClusterManagementIoKlusterletV1Resource{}
}

type OperatorOpenClusterManagementIoKlusterletV1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type OperatorOpenClusterManagementIoKlusterletV1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

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

func (r *OperatorOpenClusterManagementIoKlusterletV1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_open_cluster_management_io_klusterlet_v1"
}

func (r *OperatorOpenClusterManagementIoKlusterletV1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
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
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(63),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
						},
					},

					"deploy_option": schema.SingleNestedAttribute{
						Description:         "DeployOption contains the options of deploying a klusterlet",
						MarkdownDescription: "DeployOption contains the options of deploying a klusterlet",
						Attributes: map[string]schema.Attribute{
							"mode": schema.StringAttribute{
								Description:         "Mode can be Default or Hosted. It is Default mode if not specified In Default mode, all klusterlet related resources are deployed on the managed cluster. In Hosted mode, only crd and configurations are installed on the spoke/managed cluster. Controllers run in another cluster (defined as management-cluster) and connect to the mangaged cluster with the kubeconfig in secret of 'external-managed-kubeconfig'(a kubeconfig of managed-cluster with cluster-admin permission). Note: Do not modify the Mode field once it's applied.",
								MarkdownDescription: "Mode can be Default or Hosted. It is Default mode if not specified In Default mode, all klusterlet related resources are deployed on the managed cluster. In Hosted mode, only crd and configurations are installed on the spoke/managed cluster. Controllers run in another cluster (defined as management-cluster) and connect to the mangaged cluster with the kubeconfig in secret of 'external-managed-kubeconfig'(a kubeconfig of managed-cluster with cluster-admin permission). Note: Do not modify the Mode field once it's applied.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
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
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										validators.Base64Validator(),
									},
								},

								"url": schema.StringAttribute{
									Description:         "URL is the url of apiserver endpoint of the managed cluster.",
									MarkdownDescription: "URL is the url of apiserver endpoint of the managed cluster.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"hub_api_server_host_alias": schema.SingleNestedAttribute{
						Description:         "HubApiServerHostAlias contains the host alias for hub api server. registration-agent and work-agent will use it to communicate with hub api server.",
						MarkdownDescription: "HubApiServerHostAlias contains the host alias for hub api server. registration-agent and work-agent will use it to communicate with hub api server.",
						Attributes: map[string]schema.Attribute{
							"hostname": schema.StringAttribute{
								Description:         "Hostname for the above IP address.",
								MarkdownDescription: "Hostname for the above IP address.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`), ""),
								},
							},

							"ip": schema.StringAttribute{
								Description:         "IP address of the host file entry.",
								MarkdownDescription: "IP address of the host file entry.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"namespace": schema.StringAttribute{
						Description:         "Namespace is the namespace to deploy the agent on the managed cluster. The namespace must have a prefix of 'open-cluster-management-', and if it is not set, the namespace of 'open-cluster-management-agent' is used to deploy agent. In addition, the add-ons are deployed to the namespace of '{Namespace}-addon'. In the Hosted mode, this namespace still exists on the managed cluster to contain necessary resources, like service accounts, roles and rolebindings, while the agent is deployed to the namespace with the same name as klusterlet on the management cluster.",
						MarkdownDescription: "Namespace is the namespace to deploy the agent on the managed cluster. The namespace must have a prefix of 'open-cluster-management-', and if it is not set, the namespace of 'open-cluster-management-agent' is used to deploy agent. In addition, the add-ons are deployed to the namespace of '{Namespace}-addon'. In the Hosted mode, this namespace still exists on the managed cluster to contain necessary resources, like service accounts, roles and rolebindings, while the agent is deployed to the namespace with the same name as klusterlet on the management cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(63),
							stringvalidator.RegexMatches(regexp.MustCompile(`^open-cluster-management-[-a-z0-9]*[a-z0-9]$`), ""),
						},
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
								Optional:            true,
								Computed:            false,
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
											Optional:            true,
											Computed:            false,
										},

										"key": schema.StringAttribute{
											Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"toleration_seconds": schema.Int64Attribute{
											Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
											MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"registration_configuration": schema.SingleNestedAttribute{
						Description:         "RegistrationConfiguration contains the configuration of registration",
						MarkdownDescription: "RegistrationConfiguration contains the configuration of registration",
						Attributes: map[string]schema.Attribute{
							"client_cert_expiration_seconds": schema.Int64Attribute{
								Description:         "clientCertExpirationSeconds represents the seconds of a client certificate to expire. If it is not set or 0, the default duration seconds will be set by the hub cluster. If the value is larger than the max signing duration seconds set on the hub cluster, the max signing duration seconds will be set.",
								MarkdownDescription: "clientCertExpirationSeconds represents the seconds of a client certificate to expire. If it is not set or 0, the default duration seconds will be set by the hub cluster. If the value is larger than the max signing duration seconds set on the hub cluster, the max signing duration seconds will be set.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"feature_gates": schema.ListNestedAttribute{
								Description:         "FeatureGates represents the list of feature gates for registration If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates: 1. If featuregate/Foo does not exist, registration-operator will discard it 2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true] 3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false, he can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
								MarkdownDescription: "FeatureGates represents the list of feature gates for registration If it is set empty, default feature gates will be used. If it is set, featuregate/Foo is an example of one item in FeatureGates: 1. If featuregate/Foo does not exist, registration-operator will discard it 2. If featuregate/Foo exists and is false by default. It is now possible to set featuregate/Foo=[false|true] 3. If featuregate/Foo exists and is true by default. If a cluster-admin upgrading from 1 to 2 wants to continue having featuregate/Foo=false, he can set featuregate/Foo=false before upgrading. Let's say the cluster-admin wants featuregate/Foo=false.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"feature": schema.StringAttribute{
											Description:         "Feature is the key of feature gate. e.g. featuregate/Foo.",
											MarkdownDescription: "Feature is the key of feature gate. e.g. featuregate/Foo.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"mode": schema.StringAttribute{
											Description:         "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
											MarkdownDescription: "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Enable", "Disable"),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"registration_image_pull_spec": schema.StringAttribute{
						Description:         "RegistrationImagePullSpec represents the desired image configuration of registration agent. quay.io/open-cluster-management.io/registration:latest will be used if unspecified.",
						MarkdownDescription: "RegistrationImagePullSpec represents the desired image configuration of registration agent. quay.io/open-cluster-management.io/registration:latest will be used if unspecified.",
						Required:            false,
						Optional:            true,
						Computed:            false,
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
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"mode": schema.StringAttribute{
											Description:         "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
											MarkdownDescription: "Mode is either Enable, Disable, '' where '' is Disable by default. In Enable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=true'. In Disable mode, a valid feature gate 'featuregate/Foo' will be set to '--featuregate/Foo=false'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Enable", "Disable"),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"work_image_pull_spec": schema.StringAttribute{
						Description:         "WorkImagePullSpec represents the desired image configuration of work agent. quay.io/open-cluster-management.io/work:latest will be used if unspecified.",
						MarkdownDescription: "WorkImagePullSpec represents the desired image configuration of work agent. quay.io/open-cluster-management.io/work:latest will be used if unspecified.",
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

func (r *OperatorOpenClusterManagementIoKlusterletV1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *OperatorOpenClusterManagementIoKlusterletV1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_operator_open_cluster_management_io_klusterlet_v1")

	var model OperatorOpenClusterManagementIoKlusterletV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("operator.open-cluster-management.io/v1")
	model.Kind = pointer.String("Klusterlet")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "operator.open-cluster-management.io", Version: "v1", Resource: "klusterlets"}).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse OperatorOpenClusterManagementIoKlusterletV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *OperatorOpenClusterManagementIoKlusterletV1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_open_cluster_management_io_klusterlet_v1")

	var data OperatorOpenClusterManagementIoKlusterletV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "operator.open-cluster-management.io", Version: "v1", Resource: "klusterlets"}).
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

	var readResponse OperatorOpenClusterManagementIoKlusterletV1ResourceData
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

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *OperatorOpenClusterManagementIoKlusterletV1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_operator_open_cluster_management_io_klusterlet_v1")

	var model OperatorOpenClusterManagementIoKlusterletV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.open-cluster-management.io/v1")
	model.Kind = pointer.String("Klusterlet")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "operator.open-cluster-management.io", Version: "v1", Resource: "klusterlets"}).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse OperatorOpenClusterManagementIoKlusterletV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *OperatorOpenClusterManagementIoKlusterletV1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_operator_open_cluster_management_io_klusterlet_v1")

	var data OperatorOpenClusterManagementIoKlusterletV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "operator.open-cluster-management.io", Version: "v1", Resource: "klusterlets"}).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *OperatorOpenClusterManagementIoKlusterletV1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	if request.ID == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'name' Got: '%q'", request.ID),
		)
		return
	}
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	resource.ImportStatePassthroughID(ctx, path.Root("metadata").AtName("name"), request, response)
}
