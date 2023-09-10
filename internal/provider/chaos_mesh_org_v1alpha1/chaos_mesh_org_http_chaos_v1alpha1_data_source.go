/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package chaos_mesh_org_v1alpha1

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
	_ datasource.DataSource              = &ChaosMeshOrgHttpchaosV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &ChaosMeshOrgHttpchaosV1Alpha1DataSource{}
)

func NewChaosMeshOrgHttpchaosV1Alpha1DataSource() datasource.DataSource {
	return &ChaosMeshOrgHttpchaosV1Alpha1DataSource{}
}

type ChaosMeshOrgHttpchaosV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ChaosMeshOrgHttpchaosV1Alpha1DataSourceData struct {
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
		Abort    *bool   `tfsdk:"abort" json:"abort,omitempty"`
		Code     *int64  `tfsdk:"code" json:"code,omitempty"`
		Delay    *string `tfsdk:"delay" json:"delay,omitempty"`
		Duration *string `tfsdk:"duration" json:"duration,omitempty"`
		Method   *string `tfsdk:"method" json:"method,omitempty"`
		Mode     *string `tfsdk:"mode" json:"mode,omitempty"`
		Patch    *struct {
			Body *struct {
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"body" json:"body,omitempty"`
			Headers *[]string `tfsdk:"headers" json:"headers,omitempty"`
			Queries *[]string `tfsdk:"queries" json:"queries,omitempty"`
		} `tfsdk:"patch" json:"patch,omitempty"`
		Path          *string `tfsdk:"path" json:"path,omitempty"`
		Port          *int64  `tfsdk:"port" json:"port,omitempty"`
		RemoteCluster *string `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
		Replace       *struct {
			Body    *string            `tfsdk:"body" json:"body,omitempty"`
			Code    *int64             `tfsdk:"code" json:"code,omitempty"`
			Headers *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
			Method  *string            `tfsdk:"method" json:"method,omitempty"`
			Path    *string            `tfsdk:"path" json:"path,omitempty"`
			Queries *map[string]string `tfsdk:"queries" json:"queries,omitempty"`
		} `tfsdk:"replace" json:"replace,omitempty"`
		Request_headers  *map[string]string `tfsdk:"request_headers" json:"request_headers,omitempty"`
		Response_headers *map[string]string `tfsdk:"response_headers" json:"response_headers,omitempty"`
		Selector         *struct {
			AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
			ExpressionSelectors *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
			FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
			LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
			Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
			NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
			Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
			PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
			Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
		Target *string `tfsdk:"target" json:"target,omitempty"`
		Tls    *struct {
			CaName          *string `tfsdk:"ca_name" json:"caName,omitempty"`
			CertName        *string `tfsdk:"cert_name" json:"certName,omitempty"`
			KeyName         *string `tfsdk:"key_name" json:"keyName,omitempty"`
			SecretName      *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			SecretNamespace *string `tfsdk:"secret_namespace" json:"secretNamespace,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
		Value *string `tfsdk:"value" json:"value,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ChaosMeshOrgHttpchaosV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_chaos_mesh_org_http_chaos_v1alpha1"
}

func (r *ChaosMeshOrgHttpchaosV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "HTTPChaos is the Schema for the HTTPchaos API",
		MarkdownDescription: "HTTPChaos is the Schema for the HTTPchaos API",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"abort": schema.BoolAttribute{
						Description:         "Abort is a rule to abort a http session.",
						MarkdownDescription: "Abort is a rule to abort a http session.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"code": schema.Int64Attribute{
						Description:         "Code is a rule to select target by http status code in response.",
						MarkdownDescription: "Code is a rule to select target by http status code in response.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"delay": schema.StringAttribute{
						Description:         "Delay represents the delay of the target request/response. A duration string is a possibly unsigned sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
						MarkdownDescription: "Delay represents the delay of the target request/response. A duration string is a possibly unsigned sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"duration": schema.StringAttribute{
						Description:         "Duration represents the duration of the chaos action.",
						MarkdownDescription: "Duration represents the duration of the chaos action.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"method": schema.StringAttribute{
						Description:         "Method is a rule to select target by http method in request.",
						MarkdownDescription: "Method is a rule to select target by http method in request.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"mode": schema.StringAttribute{
						Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
						MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"patch": schema.SingleNestedAttribute{
						Description:         "Patch is a rule to patch some contents in target.",
						MarkdownDescription: "Patch is a rule to patch some contents in target.",
						Attributes: map[string]schema.Attribute{
							"body": schema.SingleNestedAttribute{
								Description:         "Body is a rule to patch message body of target.",
								MarkdownDescription: "Body is a rule to patch message body of target.",
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										Description:         "Type represents the patch type, only support 'JSON' as [merge patch json](https://tools.ietf.org/html/rfc7396) currently.",
										MarkdownDescription: "Type represents the patch type, only support 'JSON' as [merge patch json](https://tools.ietf.org/html/rfc7396) currently.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"value": schema.StringAttribute{
										Description:         "Value is the patch contents.",
										MarkdownDescription: "Value is the patch contents.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"headers": schema.ListAttribute{
								Description:         "Headers is a rule to append http headers of target. For example: '[['Set-Cookie', '<one cookie>'], ['Set-Cookie', '<another cookie>']]'.",
								MarkdownDescription: "Headers is a rule to append http headers of target. For example: '[['Set-Cookie', '<one cookie>'], ['Set-Cookie', '<another cookie>']]'.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"queries": schema.ListAttribute{
								Description:         "Queries is a rule to append uri queries of target(Request only). For example: '[['foo', 'bar'], ['foo', 'unknown']]'.",
								MarkdownDescription: "Queries is a rule to append uri queries of target(Request only). For example: '[['foo', 'bar'], ['foo', 'unknown']]'.",
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

					"path": schema.StringAttribute{
						Description:         "Path is a rule to select target by uri path in http request.",
						MarkdownDescription: "Path is a rule to select target by uri path in http request.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"port": schema.Int64Attribute{
						Description:         "Port represents the target port to be proxy of.",
						MarkdownDescription: "Port represents the target port to be proxy of.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"remote_cluster": schema.StringAttribute{
						Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
						MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"replace": schema.SingleNestedAttribute{
						Description:         "Replace is a rule to replace some contents in target.",
						MarkdownDescription: "Replace is a rule to replace some contents in target.",
						Attributes: map[string]schema.Attribute{
							"body": schema.StringAttribute{
								Description:         "Body is a rule to replace http message body in target.",
								MarkdownDescription: "Body is a rule to replace http message body in target.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"code": schema.Int64Attribute{
								Description:         "Code is a rule to replace http status code in response.",
								MarkdownDescription: "Code is a rule to replace http status code in response.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"headers": schema.MapAttribute{
								Description:         "Headers is a rule to replace http headers of target. The key-value pairs represent header name and header value pairs.",
								MarkdownDescription: "Headers is a rule to replace http headers of target. The key-value pairs represent header name and header value pairs.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"method": schema.StringAttribute{
								Description:         "Method is a rule to replace http method in request.",
								MarkdownDescription: "Method is a rule to replace http method in request.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"path": schema.StringAttribute{
								Description:         "Path is rule to to replace uri path in http request.",
								MarkdownDescription: "Path is rule to to replace uri path in http request.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"queries": schema.MapAttribute{
								Description:         "Queries is a rule to replace uri queries in http request. For example, with value '{ 'foo': 'unknown' }', the '/?foo=bar' will be altered to '/?foo=unknown',",
								MarkdownDescription: "Queries is a rule to replace uri queries in http request. For example, with value '{ 'foo': 'unknown' }', the '/?foo=bar' will be altered to '/?foo=unknown',",
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

					"request_headers": schema.MapAttribute{
						Description:         "RequestHeaders is a rule to select target by http headers in request. The key-value pairs represent header name and header value pairs.",
						MarkdownDescription: "RequestHeaders is a rule to select target by http headers in request. The key-value pairs represent header name and header value pairs.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"response_headers": schema.MapAttribute{
						Description:         "ResponseHeaders is a rule to select target by http headers in response. The key-value pairs represent header name and header value pairs.",
						MarkdownDescription: "ResponseHeaders is a rule to select target by http headers in response. The key-value pairs represent header name and header value pairs.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"selector": schema.SingleNestedAttribute{
						Description:         "Selector is used to select pods that are used to inject chaos action.",
						MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
						Attributes: map[string]schema.Attribute{
							"annotation_selectors": schema.MapAttribute{
								Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"expression_selectors": schema.ListNestedAttribute{
								Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
								MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
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

							"field_selectors": schema.MapAttribute{
								Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"label_selectors": schema.MapAttribute{
								Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespaces": schema.ListAttribute{
								Description:         "Namespaces is a set of namespace to which objects belong.",
								MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"node_selectors": schema.MapAttribute{
								Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
								MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"nodes": schema.ListAttribute{
								Description:         "Nodes is a set of node name and objects must belong to these nodes.",
								MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"pod_phase_selectors": schema.ListAttribute{
								Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
								MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"pods": schema.MapAttribute{
								Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
								MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
								ElementType:         types.ListType{ElemType: types.StringType},
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"target": schema.StringAttribute{
						Description:         "Target is the object to be selected and injected.",
						MarkdownDescription: "Target is the object to be selected and injected.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"tls": schema.SingleNestedAttribute{
						Description:         "TLS is the tls config, will override PodHttpChaos if there are multiple HTTPChaos experiments are applied",
						MarkdownDescription: "TLS is the tls config, will override PodHttpChaos if there are multiple HTTPChaos experiments are applied",
						Attributes: map[string]schema.Attribute{
							"ca_name": schema.StringAttribute{
								Description:         "CAName represents the data name of ca file in secret, 'ca.crt' for example",
								MarkdownDescription: "CAName represents the data name of ca file in secret, 'ca.crt' for example",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"cert_name": schema.StringAttribute{
								Description:         "CertName represents the data name of cert file in secret, 'tls.crt' for example",
								MarkdownDescription: "CertName represents the data name of cert file in secret, 'tls.crt' for example",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"key_name": schema.StringAttribute{
								Description:         "KeyName represents the data name of key file in secret, 'tls.key' for example",
								MarkdownDescription: "KeyName represents the data name of key file in secret, 'tls.key' for example",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"secret_name": schema.StringAttribute{
								Description:         "SecretName represents the name of required secret resource",
								MarkdownDescription: "SecretName represents the name of required secret resource",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"secret_namespace": schema.StringAttribute{
								Description:         "SecretNamespace represents the namespace of required secret resource",
								MarkdownDescription: "SecretNamespace represents the namespace of required secret resource",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"value": schema.StringAttribute{
						Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
						MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
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

func (r *ChaosMeshOrgHttpchaosV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ChaosMeshOrgHttpchaosV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_chaos_mesh_org_http_chaos_v1alpha1")

	var data ChaosMeshOrgHttpchaosV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "chaos-mesh.org", Version: "v1alpha1", Resource: "HTTPChaos"}).
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

	var readResponse ChaosMeshOrgHttpchaosV1Alpha1DataSourceData
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
	data.ApiVersion = pointer.String("chaos-mesh.org/v1alpha1")
	data.Kind = pointer.String("HTTPChaos")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
