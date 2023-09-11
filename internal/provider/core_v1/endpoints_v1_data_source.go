/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package core_v1

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
	_ datasource.DataSource              = &EndpointsV1DataSource{}
	_ datasource.DataSourceWithConfigure = &EndpointsV1DataSource{}
)

func NewEndpointsV1DataSource() datasource.DataSource {
	return &EndpointsV1DataSource{}
}

type EndpointsV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type EndpointsV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Subsets *[]struct {
		Addresses *[]struct {
			Hostname  *string `tfsdk:"hostname" json:"hostname,omitempty"`
			Ip        *string `tfsdk:"ip" json:"ip,omitempty"`
			NodeName  *string `tfsdk:"node_name" json:"nodeName,omitempty"`
			TargetRef *struct {
				ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
				Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
				Name            *string `tfsdk:"name" json:"name,omitempty"`
				Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
				ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
				Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
			} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		} `tfsdk:"addresses" json:"addresses,omitempty"`
		NotReadyAddresses *[]struct {
			Hostname  *string `tfsdk:"hostname" json:"hostname,omitempty"`
			Ip        *string `tfsdk:"ip" json:"ip,omitempty"`
			NodeName  *string `tfsdk:"node_name" json:"nodeName,omitempty"`
			TargetRef *struct {
				ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
				Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
				Name            *string `tfsdk:"name" json:"name,omitempty"`
				Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
				ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
				Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
			} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		} `tfsdk:"not_ready_addresses" json:"notReadyAddresses,omitempty"`
		Ports *[]struct {
			AppProtocol *string `tfsdk:"app_protocol" json:"appProtocol,omitempty"`
			Name        *string `tfsdk:"name" json:"name,omitempty"`
			Port        *int64  `tfsdk:"port" json:"port,omitempty"`
			Protocol    *string `tfsdk:"protocol" json:"protocol,omitempty"`
		} `tfsdk:"ports" json:"ports,omitempty"`
	} `tfsdk:"subsets" json:"subsets,omitempty"`
}

func (r *EndpointsV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_endpoints_v1"
}

func (r *EndpointsV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Endpoints is a collection of endpoints that implement the actual service. Example:	 Name: 'mysvc',	 Subsets: [	   {	     Addresses: [{'ip': '10.10.1.1'}, {'ip': '10.10.2.2'}],	     Ports: [{'name': 'a', 'port': 8675}, {'name': 'b', 'port': 309}]	   },	   {	     Addresses: [{'ip': '10.10.3.3'}],	     Ports: [{'name': 'a', 'port': 93}, {'name': 'b', 'port': 76}]	   },	]",
		MarkdownDescription: "Endpoints is a collection of endpoints that implement the actual service. Example:	 Name: 'mysvc',	 Subsets: [	   {	     Addresses: [{'ip': '10.10.1.1'}, {'ip': '10.10.2.2'}],	     Ports: [{'name': 'a', 'port': 8675}, {'name': 'b', 'port': 309}]	   },	   {	     Addresses: [{'ip': '10.10.3.3'}],	     Ports: [{'name': 'a', 'port': 93}, {'name': 'b', 'port': 76}]	   },	]",
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

			"subsets": schema.ListNestedAttribute{
				Description:         "The set of all endpoints is the union of all subsets. Addresses are placed into subsets according to the IPs they share. A single address with multiple ports, some of which are ready and some of which are not (because they come from different containers) will result in the address being displayed in different subsets for the different ports. No address will appear in both Addresses and NotReadyAddresses in the same subset. Sets of addresses and ports that comprise a service.",
				MarkdownDescription: "The set of all endpoints is the union of all subsets. Addresses are placed into subsets according to the IPs they share. A single address with multiple ports, some of which are ready and some of which are not (because they come from different containers) will result in the address being displayed in different subsets for the different ports. No address will appear in both Addresses and NotReadyAddresses in the same subset. Sets of addresses and ports that comprise a service.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"addresses": schema.ListNestedAttribute{
							Description:         "IP addresses which offer the related ports that are marked as ready. These endpoints should be considered safe for load balancers and clients to utilize.",
							MarkdownDescription: "IP addresses which offer the related ports that are marked as ready. These endpoints should be considered safe for load balancers and clients to utilize.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"hostname": schema.StringAttribute{
										Description:         "The Hostname of this endpoint",
										MarkdownDescription: "The Hostname of this endpoint",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"ip": schema.StringAttribute{
										Description:         "The IP of this endpoint. May not be loopback (127.0.0.0/8 or ::1), link-local (169.254.0.0/16 or fe80::/10), or link-local multicast (224.0.0.0/24 or ff02::/16).",
										MarkdownDescription: "The IP of this endpoint. May not be loopback (127.0.0.0/8 or ::1), link-local (169.254.0.0/16 or fe80::/10), or link-local multicast (224.0.0.0/24 or ff02::/16).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"node_name": schema.StringAttribute{
										Description:         "Optional: Node hosting this endpoint. This can be used to determine endpoints local to a node.",
										MarkdownDescription: "Optional: Node hosting this endpoint. This can be used to determine endpoints local to a node.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"target_ref": schema.SingleNestedAttribute{
										Description:         "ObjectReference contains enough information to let you inspect or modify the referred object.",
										MarkdownDescription: "ObjectReference contains enough information to let you inspect or modify the referred object.",
										Attributes: map[string]schema.Attribute{
											"api_version": schema.StringAttribute{
												Description:         "API version of the referent.",
												MarkdownDescription: "API version of the referent.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"field_path": schema.StringAttribute{
												Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
												MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"kind": schema.StringAttribute{
												Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
												MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
												MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"resource_version": schema.StringAttribute{
												Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
												MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"uid": schema.StringAttribute{
												Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
												MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
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
							},
							Required: false,
							Optional: false,
							Computed: true,
						},

						"not_ready_addresses": schema.ListNestedAttribute{
							Description:         "IP addresses which offer the related ports but are not currently marked as ready because they have not yet finished starting, have recently failed a readiness check, or have recently failed a liveness check.",
							MarkdownDescription: "IP addresses which offer the related ports but are not currently marked as ready because they have not yet finished starting, have recently failed a readiness check, or have recently failed a liveness check.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"hostname": schema.StringAttribute{
										Description:         "The Hostname of this endpoint",
										MarkdownDescription: "The Hostname of this endpoint",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"ip": schema.StringAttribute{
										Description:         "The IP of this endpoint. May not be loopback (127.0.0.0/8 or ::1), link-local (169.254.0.0/16 or fe80::/10), or link-local multicast (224.0.0.0/24 or ff02::/16).",
										MarkdownDescription: "The IP of this endpoint. May not be loopback (127.0.0.0/8 or ::1), link-local (169.254.0.0/16 or fe80::/10), or link-local multicast (224.0.0.0/24 or ff02::/16).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"node_name": schema.StringAttribute{
										Description:         "Optional: Node hosting this endpoint. This can be used to determine endpoints local to a node.",
										MarkdownDescription: "Optional: Node hosting this endpoint. This can be used to determine endpoints local to a node.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"target_ref": schema.SingleNestedAttribute{
										Description:         "ObjectReference contains enough information to let you inspect or modify the referred object.",
										MarkdownDescription: "ObjectReference contains enough information to let you inspect or modify the referred object.",
										Attributes: map[string]schema.Attribute{
											"api_version": schema.StringAttribute{
												Description:         "API version of the referent.",
												MarkdownDescription: "API version of the referent.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"field_path": schema.StringAttribute{
												Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
												MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"kind": schema.StringAttribute{
												Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
												MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
												MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"resource_version": schema.StringAttribute{
												Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
												MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"uid": schema.StringAttribute{
												Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
												MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
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
							},
							Required: false,
							Optional: false,
							Computed: true,
						},

						"ports": schema.ListNestedAttribute{
							Description:         "Port numbers available on the related IP addresses.",
							MarkdownDescription: "Port numbers available on the related IP addresses.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"app_protocol": schema.StringAttribute{
										Description:         "The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either:* Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names).* Kubernetes-defined prefixed names:  * 'kubernetes.io/h2c' - HTTP/2 over cleartext as described in https://www.rfc-editor.org/rfc/rfc7540  * 'kubernetes.io/ws'  - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455  * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455* Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.",
										MarkdownDescription: "The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either:* Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names).* Kubernetes-defined prefixed names:  * 'kubernetes.io/h2c' - HTTP/2 over cleartext as described in https://www.rfc-editor.org/rfc/rfc7540  * 'kubernetes.io/ws'  - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455  * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455* Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "The name of this port.  This must match the 'name' field in the corresponding ServicePort. Must be a DNS_LABEL. Optional only if one port is defined.",
										MarkdownDescription: "The name of this port.  This must match the 'name' field in the corresponding ServicePort. Must be a DNS_LABEL. Optional only if one port is defined.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"port": schema.Int64Attribute{
										Description:         "The port number of the endpoint.",
										MarkdownDescription: "The port number of the endpoint.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"protocol": schema.StringAttribute{
										Description:         "The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.",
										MarkdownDescription: "The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.",
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
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *EndpointsV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *EndpointsV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_endpoints_v1")

	var data EndpointsV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "", Version: "v1", Resource: "endpoints"}).
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

	var readResponse EndpointsV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("v1")
	data.Kind = pointer.String("Endpoints")
	data.Metadata = readResponse.Metadata
	data.Subsets = readResponse.Subsets

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
