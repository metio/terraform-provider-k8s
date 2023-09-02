/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hazelcast_com_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	"strings"
)

var (
	_ resource.Resource                = &HazelcastComMapV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &HazelcastComMapV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &HazelcastComMapV1Alpha1Resource{}
)

func NewHazelcastComMapV1Alpha1Resource() resource.Resource {
	return &HazelcastComMapV1Alpha1Resource{}
}

type HazelcastComMapV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type HazelcastComMapV1Alpha1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

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

func (r *HazelcastComMapV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hazelcast_com_map_v1alpha1"
}

func (r *HazelcastComMapV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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
				Description:         "MapSpec defines the desired state of Hazelcast Map Config",
				MarkdownDescription: "MapSpec defines the desired state of Hazelcast Map Config",
				Attributes: map[string]schema.Attribute{
					"async_backup_count": schema.Int64Attribute{
						Description:         "Number of asynchronous backups.",
						MarkdownDescription: "Number of asynchronous backups.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(6),
						},
					},

					"backup_count": schema.Int64Attribute{
						Description:         "Number of synchronous backups.",
						MarkdownDescription: "Number of synchronous backups.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(6),
						},
					},

					"entry_listeners": schema.ListNestedAttribute{
						Description:         "EntryListeners contains the configuration for the map-level or entry-based events listeners provided by the Hazelcast’s eventing framework. You can learn more at https://docs.hazelcast.com/hazelcast/latest/events/object-events.",
						MarkdownDescription: "EntryListeners contains the configuration for the map-level or entry-based events listeners provided by the Hazelcast’s eventing framework. You can learn more at https://docs.hazelcast.com/hazelcast/latest/events/object-events.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"class_name": schema.StringAttribute{
									Description:         "ClassName is the fully qualified name of the class that implements any of the Listener interface.",
									MarkdownDescription: "ClassName is the fully qualified name of the class that implements any of the Listener interface.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},

								"include_values": schema.BoolAttribute{
									Description:         "IncludeValues is an optional attribute that indicates whether the event will contain the map value. Defaults to true.",
									MarkdownDescription: "IncludeValues is an optional attribute that indicates whether the event will contain the map value. Defaults to true.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"local": schema.BoolAttribute{
									Description:         "Local is an optional attribute that indicates whether the map on the local member can be listened to. Defaults to false.",
									MarkdownDescription: "Local is an optional attribute that indicates whether the map on the local member can be listened to. Defaults to false.",
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

					"event_journal": schema.SingleNestedAttribute{
						Description:         "EventJournal specifies event journal configuration of the Map",
						MarkdownDescription: "EventJournal specifies event journal configuration of the Map",
						Attributes: map[string]schema.Attribute{
							"capacity": schema.Int64Attribute{
								Description:         "Capacity sets the capacity of the ringbuffer underlying the event journal.",
								MarkdownDescription: "Capacity sets the capacity of the ringbuffer underlying the event journal.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_to_live_seconds": schema.Int64Attribute{
								Description:         "TimeToLiveSeconds indicates how long the items remain in the event journal before they are expired.",
								MarkdownDescription: "TimeToLiveSeconds indicates how long the items remain in the event journal before they are expired.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"eviction": schema.SingleNestedAttribute{
						Description:         "Configuration for removing data from the map when it reaches its max size. It can be updated.",
						MarkdownDescription: "Configuration for removing data from the map when it reaches its max size. It can be updated.",
						Attributes: map[string]schema.Attribute{
							"eviction_policy": schema.StringAttribute{
								Description:         "Eviction policy to be applied when map reaches its max size according to the max size policy.",
								MarkdownDescription: "Eviction policy to be applied when map reaches its max size according to the max size policy.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("NONE", "LRU", "LFU", "RANDOM"),
								},
							},

							"max_size": schema.Int64Attribute{
								Description:         "Max size of the map.",
								MarkdownDescription: "Max size of the map.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_size_policy": schema.StringAttribute{
								Description:         "Policy for deciding if the maxSize is reached.",
								MarkdownDescription: "Policy for deciding if the maxSize is reached.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("PER_NODE", "PER_PARTITION", "USED_HEAP_SIZE", "USED_HEAP_PERCENTAGE", "FREE_HEAP_SIZE", "FREE_HEAP_PERCENTAGE", "USED_NATIVE_MEMORY_SIZE", "USED_NATIVE_MEMORY_PERCENTAGE", "FREE_NATIVE_MEMORY_SIZE", "FREE_NATIVE_MEMORY_PERCENTAGE", "ENTRY_COUNT"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"hazelcast_resource_name": schema.StringAttribute{
						Description:         "HazelcastResourceName defines the name of the Hazelcast resource that this resource is created for.",
						MarkdownDescription: "HazelcastResourceName defines the name of the Hazelcast resource that this resource is created for.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"in_memory_format": schema.StringAttribute{
						Description:         "InMemoryFormat specifies in which format data will be stored in your map",
						MarkdownDescription: "InMemoryFormat specifies in which format data will be stored in your map",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("BINARY", "OBJECT", "NATIVE"),
						},
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
									Optional:            true,
									Computed:            false,
								},

								"bit_map_index_options": schema.SingleNestedAttribute{
									Description:         "Options for 'BITMAP' index type. See https://docs.hazelcast.com/hazelcast/latest/query/indexing-maps#configuring-bitmap-indexes",
									MarkdownDescription: "Options for 'BITMAP' index type. See https://docs.hazelcast.com/hazelcast/latest/query/indexing-maps#configuring-bitmap-indexes",
									Attributes: map[string]schema.Attribute{
										"unique_key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"unique_key_transition": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("OBJECT", "LONG", "RAW"),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the index config.",
									MarkdownDescription: "Name of the index config.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"type": schema.StringAttribute{
									Description:         "Type of the index. See https://docs.hazelcast.com/hazelcast/latest/query/indexing-maps#index-types",
									MarkdownDescription: "Type of the index. See https://docs.hazelcast.com/hazelcast/latest/query/indexing-maps#index-types",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("SORTED", "HASH", "BITMAP"),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"map_store": schema.SingleNestedAttribute{
						Description:         "Configuration options when you want to load/store the map entries from/to a persistent data store such as a relational database You can learn more at https://docs.hazelcast.com/hazelcast/latest/data-structures/working-with-external-data",
						MarkdownDescription: "Configuration options when you want to load/store the map entries from/to a persistent data store such as a relational database You can learn more at https://docs.hazelcast.com/hazelcast/latest/data-structures/working-with-external-data",
						Attributes: map[string]schema.Attribute{
							"class_name": schema.StringAttribute{
								Description:         "Name of your class implementing MapLoader and/or MapStore interface.",
								MarkdownDescription: "Name of your class implementing MapLoader and/or MapStore interface.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"initial_mode": schema.StringAttribute{
								Description:         "Sets the initial entry loading mode.",
								MarkdownDescription: "Sets the initial entry loading mode.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("LAZY", "EAGER"),
								},
							},

							"properties_secret_name": schema.StringAttribute{
								Description:         "Properties can be used for giving information to the MapStore implementation",
								MarkdownDescription: "Properties can be used for giving information to the MapStore implementation",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"write_batch_size": schema.Int64Attribute{
								Description:         "Used to create batches when writing to map store.",
								MarkdownDescription: "Used to create batches when writing to map store.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"write_coalescing": schema.BoolAttribute{
								Description:         "It is meaningful if you are using write behind in MapStore. When it is set to true, only the latest store operation on a key during the write-delay-seconds will be reflected to MapStore.",
								MarkdownDescription: "It is meaningful if you are using write behind in MapStore. When it is set to true, only the latest store operation on a key during the write-delay-seconds will be reflected to MapStore.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"write_delay_seconds": schema.Int64Attribute{
								Description:         "Number of seconds to delay the storing of entries.",
								MarkdownDescription: "Number of seconds to delay the storing of entries.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"max_idle_seconds": schema.Int64Attribute{
						Description:         "Maximum time in seconds for each entry to stay idle in the map. Entries that are idle for more than this time are evicted automatically. It can be updated.",
						MarkdownDescription: "Maximum time in seconds for each entry to stay idle in the map. Entries that are idle for more than this time are evicted automatically. It can be updated.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "Name of the data structure config to be created. If empty, CR name will be used. It cannot be updated after the config is created successfully.",
						MarkdownDescription: "Name of the data structure config to be created. If empty, CR name will be used. It cannot be updated after the config is created successfully.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"near_cache": schema.SingleNestedAttribute{
						Description:         "InMemoryFormat specifies near cache configuration for map",
						MarkdownDescription: "InMemoryFormat specifies near cache configuration for map",
						Attributes: map[string]schema.Attribute{
							"cache_local_entries": schema.BoolAttribute{
								Description:         "CacheLocalEntries specifies whether the local entries are cached",
								MarkdownDescription: "CacheLocalEntries specifies whether the local entries are cached",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"eviction": schema.SingleNestedAttribute{
								Description:         "NearCacheEviction specifies the eviction behavior in Near Cache",
								MarkdownDescription: "NearCacheEviction specifies the eviction behavior in Near Cache",
								Attributes: map[string]schema.Attribute{
									"eviction_policy": schema.StringAttribute{
										Description:         "EvictionPolicy to be applied when near cache reaches its max size according to the max size policy.",
										MarkdownDescription: "EvictionPolicy to be applied when near cache reaches its max size according to the max size policy.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("NONE", "LRU", "LFU", "RANDOM"),
										},
									},

									"max_size_policy": schema.StringAttribute{
										Description:         "MaxSizePolicy for deciding if the maxSize is reached.",
										MarkdownDescription: "MaxSizePolicy for deciding if the maxSize is reached.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("PER_NODE", "PER_PARTITION", "USED_HEAP_SIZE", "USED_HEAP_PERCENTAGE", "FREE_HEAP_SIZE", "FREE_HEAP_PERCENTAGE", "USED_NATIVE_MEMORY_SIZE", "USED_NATIVE_MEMORY_PERCENTAGE", "FREE_NATIVE_MEMORY_SIZE", "FREE_NATIVE_MEMORY_PERCENTAGE", "ENTRY_COUNT"),
										},
									},

									"size": schema.Int64Attribute{
										Description:         "Size is maximum size of the Near Cache used for max-size-policy",
										MarkdownDescription: "Size is maximum size of the Near Cache used for max-size-policy",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"in_memory_format": schema.StringAttribute{
								Description:         "InMemoryFormat specifies in which format data will be stored in your near cache",
								MarkdownDescription: "InMemoryFormat specifies in which format data will be stored in your near cache",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("BINARY", "OBJECT", "NATIVE"),
								},
							},

							"invalidate_on_change": schema.BoolAttribute{
								Description:         "InvalidateOnChange specifies whether the cached entries are evicted when the entries are updated or removed",
								MarkdownDescription: "InvalidateOnChange specifies whether the cached entries are evicted when the entries are updated or removed",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_idle_seconds": schema.Int64Attribute{
								Description:         "MaxIdleSeconds Maximum number of seconds each entry can stay in the Near Cache as untouched (not read)",
								MarkdownDescription: "MaxIdleSeconds Maximum number of seconds each entry can stay in the Near Cache as untouched (not read)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name is name of the near cache",
								MarkdownDescription: "Name is name of the near cache",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_to_live_seconds": schema.Int64Attribute{
								Description:         "TimeToLiveSeconds maximum number of seconds for each entry to stay in the Near Cache",
								MarkdownDescription: "TimeToLiveSeconds maximum number of seconds for each entry to stay in the Near Cache",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"persistence_enabled": schema.BoolAttribute{
						Description:         "When enabled, map data will be persisted. It cannot be updated after map config is created successfully.",
						MarkdownDescription: "When enabled, map data will be persisted. It cannot be updated after map config is created successfully.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"time_to_live_seconds": schema.Int64Attribute{
						Description:         "Maximum time in seconds for each entry to stay in the map. If it is not 0, entries that are older than this time and not updated for this time are evicted automatically. It can be updated.",
						MarkdownDescription: "Maximum time in seconds for each entry to stay in the map. If it is not 0, entries that are older than this time and not updated for this time are evicted automatically. It can be updated.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *HazelcastComMapV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *HazelcastComMapV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_hazelcast_com_map_v1alpha1")

	var model HazelcastComMapV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("hazelcast.com/v1alpha1")
	model.Kind = pointer.String("Map")

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

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "hazelcast.com", Version: "v1alpha1", Resource: "Map"}).
		Namespace(model.Metadata.Namespace).
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

	var readResponse HazelcastComMapV1Alpha1ResourceData
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

func (r *HazelcastComMapV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hazelcast_com_map_v1alpha1")

	var data HazelcastComMapV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
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

	var readResponse HazelcastComMapV1Alpha1ResourceData
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

func (r *HazelcastComMapV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_hazelcast_com_map_v1alpha1")

	var model HazelcastComMapV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("hazelcast.com/v1alpha1")
	model.Kind = pointer.String("Map")

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

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "hazelcast.com", Version: "v1alpha1", Resource: "Map"}).
		Namespace(model.Metadata.Namespace).
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

	var readResponse HazelcastComMapV1Alpha1ResourceData
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

func (r *HazelcastComMapV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_hazelcast_com_map_v1alpha1")

	var data HazelcastComMapV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hazelcast.com", Version: "v1alpha1", Resource: "Map"}).
		Namespace(data.Metadata.Namespace).
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

func (r *HazelcastComMapV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
