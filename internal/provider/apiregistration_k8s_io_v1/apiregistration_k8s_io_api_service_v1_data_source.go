/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apiregistration_k8s_io_v1

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
	_ datasource.DataSource              = &ApiregistrationK8SIoApiserviceV1DataSource{}
	_ datasource.DataSourceWithConfigure = &ApiregistrationK8SIoApiserviceV1DataSource{}
)

func NewApiregistrationK8SIoApiserviceV1DataSource() datasource.DataSource {
	return &ApiregistrationK8SIoApiserviceV1DataSource{}
}

type ApiregistrationK8SIoApiserviceV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ApiregistrationK8SIoApiserviceV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		CaBundle              *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
		Group                 *string `tfsdk:"group" json:"group,omitempty"`
		GroupPriorityMinimum  *int64  `tfsdk:"group_priority_minimum" json:"groupPriorityMinimum,omitempty"`
		InsecureSkipTLSVerify *bool   `tfsdk:"insecure_skip_tls_verify" json:"insecureSkipTLSVerify,omitempty"`
		Service               *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Port      *int64  `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"service" json:"service,omitempty"`
		Version         *string `tfsdk:"version" json:"version,omitempty"`
		VersionPriority *int64  `tfsdk:"version_priority" json:"versionPriority,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ApiregistrationK8SIoApiserviceV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apiregistration_k8s_io_api_service_v1"
}

func (r *ApiregistrationK8SIoApiserviceV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "APIService represents a server for a particular GroupVersion. Name must be 'version.group'.",
		MarkdownDescription: "APIService represents a server for a particular GroupVersion. Name must be 'version.group'.",
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
				Description:         "APIServiceSpec contains information for locating and communicating with a server. Only https is supported, though you are able to disable certificate verification.",
				MarkdownDescription: "APIServiceSpec contains information for locating and communicating with a server. Only https is supported, though you are able to disable certificate verification.",
				Attributes: map[string]schema.Attribute{
					"ca_bundle": schema.StringAttribute{
						Description:         "CABundle is a PEM encoded CA bundle which will be used to validate an API server's serving certificate. If unspecified, system trust roots on the apiserver are used.",
						MarkdownDescription: "CABundle is a PEM encoded CA bundle which will be used to validate an API server's serving certificate. If unspecified, system trust roots on the apiserver are used.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"group": schema.StringAttribute{
						Description:         "Group is the API group name this server hosts",
						MarkdownDescription: "Group is the API group name this server hosts",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"group_priority_minimum": schema.Int64Attribute{
						Description:         "GroupPriorityMininum is the priority this group should have at least. Higher priority means that the group is preferred by clients over lower priority ones. Note that other versions of this group might specify even higher GroupPriorityMininum values such that the whole group gets a higher priority. The primary sort is based on GroupPriorityMinimum, ordered highest number to lowest (20 before 10). The secondary sort is based on the alphabetical comparison of the name of the object.  (v1.bar before v1.foo) We'd recommend something like: *.k8s.io (except extensions) at 18000 and PaaSes (OpenShift, Deis) are recommended to be in the 2000s",
						MarkdownDescription: "GroupPriorityMininum is the priority this group should have at least. Higher priority means that the group is preferred by clients over lower priority ones. Note that other versions of this group might specify even higher GroupPriorityMininum values such that the whole group gets a higher priority. The primary sort is based on GroupPriorityMinimum, ordered highest number to lowest (20 before 10). The secondary sort is based on the alphabetical comparison of the name of the object.  (v1.bar before v1.foo) We'd recommend something like: *.k8s.io (except extensions) at 18000 and PaaSes (OpenShift, Deis) are recommended to be in the 2000s",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"insecure_skip_tls_verify": schema.BoolAttribute{
						Description:         "InsecureSkipTLSVerify disables TLS certificate verification when communicating with this server. This is strongly discouraged.  You should use the CABundle instead.",
						MarkdownDescription: "InsecureSkipTLSVerify disables TLS certificate verification when communicating with this server. This is strongly discouraged.  You should use the CABundle instead.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"service": schema.SingleNestedAttribute{
						Description:         "ServiceReference holds a reference to Service.legacy.k8s.io",
						MarkdownDescription: "ServiceReference holds a reference to Service.legacy.k8s.io",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name is the name of the service",
								MarkdownDescription: "Name is the name of the service",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace is the namespace of the service",
								MarkdownDescription: "Namespace is the namespace of the service",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"port": schema.Int64Attribute{
								Description:         "If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. 'port' should be a valid port number (1-65535, inclusive).",
								MarkdownDescription: "If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. 'port' should be a valid port number (1-65535, inclusive).",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"version": schema.StringAttribute{
						Description:         "Version is the API version this server hosts.  For example, 'v1'",
						MarkdownDescription: "Version is the API version this server hosts.  For example, 'v1'",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"version_priority": schema.Int64Attribute{
						Description:         "VersionPriority controls the ordering of this API version inside of its group.  Must be greater than zero. The primary sort is based on VersionPriority, ordered highest to lowest (20 before 10). Since it's inside of a group, the number can be small, probably in the 10s. In case of equal version priorities, the version string will be used to compute the order inside a group. If the version string is 'kube-like', it will sort above non 'kube-like' version strings, which are ordered lexicographically. 'Kube-like' versions start with a 'v', then are followed by a number (the major version), then optionally the string 'alpha' or 'beta' and another number (the minor version). These are sorted first by GA > beta > alpha (where GA is a version with no suffix such as beta or alpha), and then by comparing major version, then minor version. An example sorted list of versions: v10, v2, v1, v11beta2, v10beta3, v3beta1, v12alpha1, v11alpha2, foo1, foo10.",
						MarkdownDescription: "VersionPriority controls the ordering of this API version inside of its group.  Must be greater than zero. The primary sort is based on VersionPriority, ordered highest to lowest (20 before 10). Since it's inside of a group, the number can be small, probably in the 10s. In case of equal version priorities, the version string will be used to compute the order inside a group. If the version string is 'kube-like', it will sort above non 'kube-like' version strings, which are ordered lexicographically. 'Kube-like' versions start with a 'v', then are followed by a number (the major version), then optionally the string 'alpha' or 'beta' and another number (the minor version). These are sorted first by GA > beta > alpha (where GA is a version with no suffix such as beta or alpha), and then by comparing major version, then minor version. An example sorted list of versions: v10, v2, v1, v11beta2, v10beta3, v3beta1, v12alpha1, v11alpha2, foo1, foo10.",
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

func (r *ApiregistrationK8SIoApiserviceV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ApiregistrationK8SIoApiserviceV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_apiregistration_k8s_io_api_service_v1")

	var data ApiregistrationK8SIoApiserviceV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "apiregistration.k8s.io", Version: "v1", Resource: "APIService"}).
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

	var readResponse ApiregistrationK8SIoApiserviceV1DataSourceData
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
	data.ApiVersion = pointer.String("apiregistration.k8s.io/v1")
	data.Kind = pointer.String("APIService")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
