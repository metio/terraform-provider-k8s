/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hazelcast_com_v1alpha1

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
	_ datasource.DataSource              = &HazelcastComWanReplicationV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &HazelcastComWanReplicationV1Alpha1DataSource{}
)

func NewHazelcastComWanReplicationV1Alpha1DataSource() datasource.DataSource {
	return &HazelcastComWanReplicationV1Alpha1DataSource{}
}

type HazelcastComWanReplicationV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type HazelcastComWanReplicationV1Alpha1DataSourceData struct {
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
		Acknowledgement *struct {
			Timeout *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
			Type    *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"acknowledgement" json:"acknowledgement,omitempty"`
		Batch *struct {
			MaximumDelay *int64 `tfsdk:"maximum_delay" json:"maximumDelay,omitempty"`
			Size         *int64 `tfsdk:"size" json:"size,omitempty"`
		} `tfsdk:"batch" json:"batch,omitempty"`
		Endpoints *string `tfsdk:"endpoints" json:"endpoints,omitempty"`
		Queue     *struct {
			Capacity     *int64  `tfsdk:"capacity" json:"capacity,omitempty"`
			FullBehavior *string `tfsdk:"full_behavior" json:"fullBehavior,omitempty"`
		} `tfsdk:"queue" json:"queue,omitempty"`
		Resources *[]struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		TargetClusterName *string `tfsdk:"target_cluster_name" json:"targetClusterName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *HazelcastComWanReplicationV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hazelcast_com_wan_replication_v1alpha1"
}

func (r *HazelcastComWanReplicationV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "WanReplication is the Schema for the wanreplications API",
		MarkdownDescription: "WanReplication is the Schema for the wanreplications API",
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
				Description:         "WanReplicationSpec defines the desired state of WanReplication",
				MarkdownDescription: "WanReplicationSpec defines the desired state of WanReplication",
				Attributes: map[string]schema.Attribute{
					"acknowledgement": schema.SingleNestedAttribute{
						Description:         "Acknowledgement is the configuration for the condition when the next batch of WAN events are sent.",
						MarkdownDescription: "Acknowledgement is the configuration for the condition when the next batch of WAN events are sent.",
						Attributes: map[string]schema.Attribute{
							"timeout": schema.Int64Attribute{
								Description:         "Timeout represents the time in milliseconds the source cluster waits for the acknowledgement. After timeout, the events will be resent.",
								MarkdownDescription: "Timeout represents the time in milliseconds the source cluster waits for the acknowledgement. After timeout, the events will be resent.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"type": schema.StringAttribute{
								Description:         "Type represents how a batch of replication events is considered successfully replicated.",
								MarkdownDescription: "Type represents how a batch of replication events is considered successfully replicated.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"batch": schema.SingleNestedAttribute{
						Description:         "Batch is the configuration for WAN events batch.",
						MarkdownDescription: "Batch is the configuration for WAN events batch.",
						Attributes: map[string]schema.Attribute{
							"maximum_delay": schema.Int64Attribute{
								Description:         "MaximumDelay represents the maximum delay in milliseconds. If the batch size is not reached, the events will be sent after the maximum delay.",
								MarkdownDescription: "MaximumDelay represents the maximum delay in milliseconds. If the batch size is not reached, the events will be sent after the maximum delay.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"size": schema.Int64Attribute{
								Description:         "Size represents the maximum batch size.",
								MarkdownDescription: "Size represents the maximum batch size.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"endpoints": schema.StringAttribute{
						Description:         "Endpoints is the target cluster comma separated endpoint list .",
						MarkdownDescription: "Endpoints is the target cluster comma separated endpoint list .",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"queue": schema.SingleNestedAttribute{
						Description:         "Queue is the configuration for WAN events queue.",
						MarkdownDescription: "Queue is the configuration for WAN events queue.",
						Attributes: map[string]schema.Attribute{
							"capacity": schema.Int64Attribute{
								Description:         "Capacity is the total capacity of WAN queue.",
								MarkdownDescription: "Capacity is the total capacity of WAN queue.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"full_behavior": schema.StringAttribute{
								Description:         "FullBehavior represents the behavior of the new arrival when the queue is full.",
								MarkdownDescription: "FullBehavior represents the behavior of the new arrival when the queue is full.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"resources": schema.ListNestedAttribute{
						Description:         "Resources is the list of custom resources to which WAN replication applies.",
						MarkdownDescription: "Resources is the list of custom resources to which WAN replication applies.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"kind": schema.StringAttribute{
									Description:         "Kind is the kind of custom resource to which WAN replication applies.",
									MarkdownDescription: "Kind is the kind of custom resource to which WAN replication applies.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "Name is the name of custom resource to which WAN replication applies.",
									MarkdownDescription: "Name is the name of custom resource to which WAN replication applies.",
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

					"target_cluster_name": schema.StringAttribute{
						Description:         "ClusterName is the clusterName field of the target Hazelcast resource.",
						MarkdownDescription: "ClusterName is the clusterName field of the target Hazelcast resource.",
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

func (r *HazelcastComWanReplicationV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *HazelcastComWanReplicationV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_hazelcast_com_wan_replication_v1alpha1")

	var data HazelcastComWanReplicationV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hazelcast.com", Version: "v1alpha1", Resource: "wanreplications"}).
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

	var readResponse HazelcastComWanReplicationV1Alpha1DataSourceData
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
	data.ApiVersion = pointer.String("hazelcast.com/v1alpha1")
	data.Kind = pointer.String("WanReplication")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
