/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hazelcast_com_v1alpha1

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
	_ datasource.DataSource              = &HazelcastComMapV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &HazelcastComMapV1Alpha1DataSource{}
)

func NewHazelcastComMapV1Alpha1DataSource() datasource.DataSource {
	return &HazelcastComMapV1Alpha1DataSource{}
}

type HazelcastComMapV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type HazelcastComMapV1Alpha1DataSourceData struct {
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
		AsyncBackupCount *int64 `tfsdk:"async_backup_count" json:"asyncBackupCount,omitempty"`
		BackupCount      *int64 `tfsdk:"backup_count" json:"backupCount,omitempty"`
		EntryListeners   *[]struct {
			ClassName     *string `tfsdk:"class_name" json:"className,omitempty"`
			IncludeValues *bool   `tfsdk:"include_values" json:"includeValues,omitempty"`
			Local         *bool   `tfsdk:"local" json:"local,omitempty"`
		} `tfsdk:"entry_listeners" json:"entryListeners,omitempty"`
		EventJournal *struct {
			Capacity          *int64 `tfsdk:"capacity" json:"capacity,omitempty"`
			TimeToLiveSeconds *int64 `tfsdk:"time_to_live_seconds" json:"timeToLiveSeconds,omitempty"`
		} `tfsdk:"event_journal" json:"eventJournal,omitempty"`
		Eviction *struct {
			EvictionPolicy *string `tfsdk:"eviction_policy" json:"evictionPolicy,omitempty"`
			MaxSize        *int64  `tfsdk:"max_size" json:"maxSize,omitempty"`
			MaxSizePolicy  *string `tfsdk:"max_size_policy" json:"maxSizePolicy,omitempty"`
		} `tfsdk:"eviction" json:"eviction,omitempty"`
		HazelcastResourceName *string `tfsdk:"hazelcast_resource_name" json:"hazelcastResourceName,omitempty"`
		InMemoryFormat        *string `tfsdk:"in_memory_format" json:"inMemoryFormat,omitempty"`
		Indexes               *[]struct {
			Attributes         *[]string `tfsdk:"attributes" json:"attributes,omitempty"`
			BitMapIndexOptions *struct {
				UniqueKey           *string `tfsdk:"unique_key" json:"uniqueKey,omitempty"`
				UniqueKeyTransition *string `tfsdk:"unique_key_transition" json:"uniqueKeyTransition,omitempty"`
			} `tfsdk:"bit_map_index_options" json:"bitMapIndexOptions,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"indexes" json:"indexes,omitempty"`
		MapStore *struct {
			ClassName            *string `tfsdk:"class_name" json:"className,omitempty"`
			InitialMode          *string `tfsdk:"initial_mode" json:"initialMode,omitempty"`
			PropertiesSecretName *string `tfsdk:"properties_secret_name" json:"propertiesSecretName,omitempty"`
			WriteBatchSize       *int64  `tfsdk:"write_batch_size" json:"writeBatchSize,omitempty"`
			WriteCoalescing      *bool   `tfsdk:"write_coalescing" json:"writeCoalescing,omitempty"`
			WriteDelaySeconds    *int64  `tfsdk:"write_delay_seconds" json:"writeDelaySeconds,omitempty"`
		} `tfsdk:"map_store" json:"mapStore,omitempty"`
		MaxIdleSeconds *int64  `tfsdk:"max_idle_seconds" json:"maxIdleSeconds,omitempty"`
		Name           *string `tfsdk:"name" json:"name,omitempty"`
		NearCache      *struct {
			CacheLocalEntries *bool `tfsdk:"cache_local_entries" json:"cacheLocalEntries,omitempty"`
			Eviction          *struct {
				EvictionPolicy *string `tfsdk:"eviction_policy" json:"evictionPolicy,omitempty"`
				MaxSizePolicy  *string `tfsdk:"max_size_policy" json:"maxSizePolicy,omitempty"`
				Size           *int64  `tfsdk:"size" json:"size,omitempty"`
			} `tfsdk:"eviction" json:"eviction,omitempty"`
			InMemoryFormat     *string `tfsdk:"in_memory_format" json:"inMemoryFormat,omitempty"`
			InvalidateOnChange *bool   `tfsdk:"invalidate_on_change" json:"invalidateOnChange,omitempty"`
			MaxIdleSeconds     *int64  `tfsdk:"max_idle_seconds" json:"maxIdleSeconds,omitempty"`
			Name               *string `tfsdk:"name" json:"name,omitempty"`
			TimeToLiveSeconds  *int64  `tfsdk:"time_to_live_seconds" json:"timeToLiveSeconds,omitempty"`
		} `tfsdk:"near_cache" json:"nearCache,omitempty"`
		PersistenceEnabled *bool  `tfsdk:"persistence_enabled" json:"persistenceEnabled,omitempty"`
		TimeToLiveSeconds  *int64 `tfsdk:"time_to_live_seconds" json:"timeToLiveSeconds,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *HazelcastComMapV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hazelcast_com_map_v1alpha1"
}

func (r *HazelcastComMapV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Map is the Schema for the maps API",
		MarkdownDescription: "Map is the Schema for the maps API",
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
				Description:         "MapSpec defines the desired state of Hazelcast Map Config",
				MarkdownDescription: "MapSpec defines the desired state of Hazelcast Map Config",
				Attributes: map[string]schema.Attribute{
					"async_backup_count": schema.Int64Attribute{
						Description:         "Number of asynchronous backups.",
						MarkdownDescription: "Number of asynchronous backups.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"backup_count": schema.Int64Attribute{
						Description:         "Number of synchronous backups.",
						MarkdownDescription: "Number of synchronous backups.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"entry_listeners": schema.ListNestedAttribute{
						Description:         "EntryListeners contains the configuration for the map-level or entry-based events listeners provided by the Hazelcast’s eventing framework. You can learn more at https://docs.hazelcast.com/hazelcast/latest/events/object-events.",
						MarkdownDescription: "EntryListeners contains the configuration for the map-level or entry-based events listeners provided by the Hazelcast’s eventing framework. You can learn more at https://docs.hazelcast.com/hazelcast/latest/events/object-events.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"class_name": schema.StringAttribute{
									Description:         "ClassName is the fully qualified name of the class that implements any of the Listener interface.",
									MarkdownDescription: "ClassName is the fully qualified name of the class that implements any of the Listener interface.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"include_values": schema.BoolAttribute{
									Description:         "IncludeValues is an optional attribute that indicates whether the event will contain the map value. Defaults to true.",
									MarkdownDescription: "IncludeValues is an optional attribute that indicates whether the event will contain the map value. Defaults to true.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"local": schema.BoolAttribute{
									Description:         "Local is an optional attribute that indicates whether the map on the local member can be listened to. Defaults to false.",
									MarkdownDescription: "Local is an optional attribute that indicates whether the map on the local member can be listened to. Defaults to false.",
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

					"event_journal": schema.SingleNestedAttribute{
						Description:         "EventJournal specifies event journal configuration of the Map",
						MarkdownDescription: "EventJournal specifies event journal configuration of the Map",
						Attributes: map[string]schema.Attribute{
							"capacity": schema.Int64Attribute{
								Description:         "Capacity sets the capacity of the ringbuffer underlying the event journal.",
								MarkdownDescription: "Capacity sets the capacity of the ringbuffer underlying the event journal.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"time_to_live_seconds": schema.Int64Attribute{
								Description:         "TimeToLiveSeconds indicates how long the items remain in the event journal before they are expired.",
								MarkdownDescription: "TimeToLiveSeconds indicates how long the items remain in the event journal before they are expired.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"eviction": schema.SingleNestedAttribute{
						Description:         "Configuration for removing data from the map when it reaches its max size. It can be updated.",
						MarkdownDescription: "Configuration for removing data from the map when it reaches its max size. It can be updated.",
						Attributes: map[string]schema.Attribute{
							"eviction_policy": schema.StringAttribute{
								Description:         "Eviction policy to be applied when map reaches its max size according to the max size policy.",
								MarkdownDescription: "Eviction policy to be applied when map reaches its max size according to the max size policy.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"max_size": schema.Int64Attribute{
								Description:         "Max size of the map.",
								MarkdownDescription: "Max size of the map.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"max_size_policy": schema.StringAttribute{
								Description:         "Policy for deciding if the maxSize is reached.",
								MarkdownDescription: "Policy for deciding if the maxSize is reached.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"hazelcast_resource_name": schema.StringAttribute{
						Description:         "HazelcastResourceName defines the name of the Hazelcast resource that this resource is created for.",
						MarkdownDescription: "HazelcastResourceName defines the name of the Hazelcast resource that this resource is created for.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"in_memory_format": schema.StringAttribute{
						Description:         "InMemoryFormat specifies in which format data will be stored in your map",
						MarkdownDescription: "InMemoryFormat specifies in which format data will be stored in your map",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"indexes": schema.ListNestedAttribute{
						Description:         "Indexes to be created for the map data. You can learn more at https://docs.hazelcast.com/hazelcast/latest/query/indexing-maps. It cannot be updated after map config is created successfully.",
						MarkdownDescription: "Indexes to be created for the map data. You can learn more at https://docs.hazelcast.com/hazelcast/latest/query/indexing-maps. It cannot be updated after map config is created successfully.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"attributes": schema.ListAttribute{
									Description:         "Attributes of the index.",
									MarkdownDescription: "Attributes of the index.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"bit_map_index_options": schema.SingleNestedAttribute{
									Description:         "Options for 'BITMAP' index type. See https://docs.hazelcast.com/hazelcast/latest/query/indexing-maps#configuring-bitmap-indexes",
									MarkdownDescription: "Options for 'BITMAP' index type. See https://docs.hazelcast.com/hazelcast/latest/query/indexing-maps#configuring-bitmap-indexes",
									Attributes: map[string]schema.Attribute{
										"unique_key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"unique_key_transition": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the index config.",
									MarkdownDescription: "Name of the index config.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"type": schema.StringAttribute{
									Description:         "Type of the index. See https://docs.hazelcast.com/hazelcast/latest/query/indexing-maps#index-types",
									MarkdownDescription: "Type of the index. See https://docs.hazelcast.com/hazelcast/latest/query/indexing-maps#index-types",
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

					"map_store": schema.SingleNestedAttribute{
						Description:         "Configuration options when you want to load/store the map entries from/to a persistent data store such as a relational database You can learn more at https://docs.hazelcast.com/hazelcast/latest/data-structures/working-with-external-data",
						MarkdownDescription: "Configuration options when you want to load/store the map entries from/to a persistent data store such as a relational database You can learn more at https://docs.hazelcast.com/hazelcast/latest/data-structures/working-with-external-data",
						Attributes: map[string]schema.Attribute{
							"class_name": schema.StringAttribute{
								Description:         "Name of your class implementing MapLoader and/or MapStore interface.",
								MarkdownDescription: "Name of your class implementing MapLoader and/or MapStore interface.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"initial_mode": schema.StringAttribute{
								Description:         "Sets the initial entry loading mode.",
								MarkdownDescription: "Sets the initial entry loading mode.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"properties_secret_name": schema.StringAttribute{
								Description:         "Properties can be used for giving information to the MapStore implementation",
								MarkdownDescription: "Properties can be used for giving information to the MapStore implementation",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"write_batch_size": schema.Int64Attribute{
								Description:         "Used to create batches when writing to map store.",
								MarkdownDescription: "Used to create batches when writing to map store.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"write_coalescing": schema.BoolAttribute{
								Description:         "It is meaningful if you are using write behind in MapStore. When it is set to true, only the latest store operation on a key during the write-delay-seconds will be reflected to MapStore.",
								MarkdownDescription: "It is meaningful if you are using write behind in MapStore. When it is set to true, only the latest store operation on a key during the write-delay-seconds will be reflected to MapStore.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"write_delay_seconds": schema.Int64Attribute{
								Description:         "Number of seconds to delay the storing of entries.",
								MarkdownDescription: "Number of seconds to delay the storing of entries.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"max_idle_seconds": schema.Int64Attribute{
						Description:         "Maximum time in seconds for each entry to stay idle in the map. Entries that are idle for more than this time are evicted automatically. It can be updated.",
						MarkdownDescription: "Maximum time in seconds for each entry to stay idle in the map. Entries that are idle for more than this time are evicted automatically. It can be updated.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"name": schema.StringAttribute{
						Description:         "Name of the data structure config to be created. If empty, CR name will be used. It cannot be updated after the config is created successfully.",
						MarkdownDescription: "Name of the data structure config to be created. If empty, CR name will be used. It cannot be updated after the config is created successfully.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"near_cache": schema.SingleNestedAttribute{
						Description:         "InMemoryFormat specifies near cache configuration for map",
						MarkdownDescription: "InMemoryFormat specifies near cache configuration for map",
						Attributes: map[string]schema.Attribute{
							"cache_local_entries": schema.BoolAttribute{
								Description:         "CacheLocalEntries specifies whether the local entries are cached",
								MarkdownDescription: "CacheLocalEntries specifies whether the local entries are cached",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"eviction": schema.SingleNestedAttribute{
								Description:         "NearCacheEviction specifies the eviction behavior in Near Cache",
								MarkdownDescription: "NearCacheEviction specifies the eviction behavior in Near Cache",
								Attributes: map[string]schema.Attribute{
									"eviction_policy": schema.StringAttribute{
										Description:         "EvictionPolicy to be applied when near cache reaches its max size according to the max size policy.",
										MarkdownDescription: "EvictionPolicy to be applied when near cache reaches its max size according to the max size policy.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"max_size_policy": schema.StringAttribute{
										Description:         "MaxSizePolicy for deciding if the maxSize is reached.",
										MarkdownDescription: "MaxSizePolicy for deciding if the maxSize is reached.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"size": schema.Int64Attribute{
										Description:         "Size is maximum size of the Near Cache used for max-size-policy",
										MarkdownDescription: "Size is maximum size of the Near Cache used for max-size-policy",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"in_memory_format": schema.StringAttribute{
								Description:         "InMemoryFormat specifies in which format data will be stored in your near cache",
								MarkdownDescription: "InMemoryFormat specifies in which format data will be stored in your near cache",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"invalidate_on_change": schema.BoolAttribute{
								Description:         "InvalidateOnChange specifies whether the cached entries are evicted when the entries are updated or removed",
								MarkdownDescription: "InvalidateOnChange specifies whether the cached entries are evicted when the entries are updated or removed",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"max_idle_seconds": schema.Int64Attribute{
								Description:         "MaxIdleSeconds Maximum number of seconds each entry can stay in the Near Cache as untouched (not read)",
								MarkdownDescription: "MaxIdleSeconds Maximum number of seconds each entry can stay in the Near Cache as untouched (not read)",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name is name of the near cache",
								MarkdownDescription: "Name is name of the near cache",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"time_to_live_seconds": schema.Int64Attribute{
								Description:         "TimeToLiveSeconds maximum number of seconds for each entry to stay in the Near Cache",
								MarkdownDescription: "TimeToLiveSeconds maximum number of seconds for each entry to stay in the Near Cache",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"persistence_enabled": schema.BoolAttribute{
						Description:         "When enabled, map data will be persisted. It cannot be updated after map config is created successfully.",
						MarkdownDescription: "When enabled, map data will be persisted. It cannot be updated after map config is created successfully.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"time_to_live_seconds": schema.Int64Attribute{
						Description:         "Maximum time in seconds for each entry to stay in the map. If it is not 0, entries that are older than this time and not updated for this time are evicted automatically. It can be updated.",
						MarkdownDescription: "Maximum time in seconds for each entry to stay in the map. If it is not 0, entries that are older than this time and not updated for this time are evicted automatically. It can be updated.",
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

func (r *HazelcastComMapV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *HazelcastComMapV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_hazelcast_com_map_v1alpha1")

	var data HazelcastComMapV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hazelcast.com", Version: "v1alpha1", Resource: "Map"}).
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

	var readResponse HazelcastComMapV1Alpha1DataSourceData
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
	data.Kind = pointer.String("Map")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
