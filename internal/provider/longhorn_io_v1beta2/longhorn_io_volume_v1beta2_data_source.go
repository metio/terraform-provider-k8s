/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package longhorn_io_v1beta2

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
	_ datasource.DataSource              = &LonghornIoVolumeV1Beta2DataSource{}
	_ datasource.DataSourceWithConfigure = &LonghornIoVolumeV1Beta2DataSource{}
)

func NewLonghornIoVolumeV1Beta2DataSource() datasource.DataSource {
	return &LonghornIoVolumeV1Beta2DataSource{}
}

type LonghornIoVolumeV1Beta2DataSource struct {
	kubernetesClient dynamic.Interface
}

type LonghornIoVolumeV1Beta2DataSourceData struct {
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
		Standby                     *bool     `tfsdk:"standby" json:"Standby,omitempty"`
		AccessMode                  *string   `tfsdk:"access_mode" json:"accessMode,omitempty"`
		BackendStoreDriver          *string   `tfsdk:"backend_store_driver" json:"backendStoreDriver,omitempty"`
		BackingImage                *string   `tfsdk:"backing_image" json:"backingImage,omitempty"`
		BackupCompressionMethod     *string   `tfsdk:"backup_compression_method" json:"backupCompressionMethod,omitempty"`
		DataLocality                *string   `tfsdk:"data_locality" json:"dataLocality,omitempty"`
		DataSource                  *string   `tfsdk:"data_source" json:"dataSource,omitempty"`
		DisableFrontend             *bool     `tfsdk:"disable_frontend" json:"disableFrontend,omitempty"`
		DiskSelector                *[]string `tfsdk:"disk_selector" json:"diskSelector,omitempty"`
		Encrypted                   *bool     `tfsdk:"encrypted" json:"encrypted,omitempty"`
		EngineImage                 *string   `tfsdk:"engine_image" json:"engineImage,omitempty"`
		FromBackup                  *string   `tfsdk:"from_backup" json:"fromBackup,omitempty"`
		Frontend                    *string   `tfsdk:"frontend" json:"frontend,omitempty"`
		LastAttachedBy              *string   `tfsdk:"last_attached_by" json:"lastAttachedBy,omitempty"`
		Migratable                  *bool     `tfsdk:"migratable" json:"migratable,omitempty"`
		MigrationNodeID             *string   `tfsdk:"migration_node_id" json:"migrationNodeID,omitempty"`
		NodeID                      *string   `tfsdk:"node_id" json:"nodeID,omitempty"`
		NodeSelector                *[]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		NumberOfReplicas            *int64    `tfsdk:"number_of_replicas" json:"numberOfReplicas,omitempty"`
		OfflineReplicaRebuilding    *string   `tfsdk:"offline_replica_rebuilding" json:"offlineReplicaRebuilding,omitempty"`
		ReplicaAutoBalance          *string   `tfsdk:"replica_auto_balance" json:"replicaAutoBalance,omitempty"`
		ReplicaDiskSoftAntiAffinity *string   `tfsdk:"replica_disk_soft_anti_affinity" json:"replicaDiskSoftAntiAffinity,omitempty"`
		ReplicaSoftAntiAffinity     *string   `tfsdk:"replica_soft_anti_affinity" json:"replicaSoftAntiAffinity,omitempty"`
		ReplicaZoneSoftAntiAffinity *string   `tfsdk:"replica_zone_soft_anti_affinity" json:"replicaZoneSoftAntiAffinity,omitempty"`
		RestoreVolumeRecurringJob   *string   `tfsdk:"restore_volume_recurring_job" json:"restoreVolumeRecurringJob,omitempty"`
		RevisionCounterDisabled     *bool     `tfsdk:"revision_counter_disabled" json:"revisionCounterDisabled,omitempty"`
		Size                        *string   `tfsdk:"size" json:"size,omitempty"`
		SnapshotDataIntegrity       *string   `tfsdk:"snapshot_data_integrity" json:"snapshotDataIntegrity,omitempty"`
		StaleReplicaTimeout         *int64    `tfsdk:"stale_replica_timeout" json:"staleReplicaTimeout,omitempty"`
		UnmapMarkSnapChainRemoved   *string   `tfsdk:"unmap_mark_snap_chain_removed" json:"unmapMarkSnapChainRemoved,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *LonghornIoVolumeV1Beta2DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_longhorn_io_volume_v1beta2"
}

func (r *LonghornIoVolumeV1Beta2DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Volume is where Longhorn stores volume object.",
		MarkdownDescription: "Volume is where Longhorn stores volume object.",
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
				Description:         "VolumeSpec defines the desired state of the Longhorn volume",
				MarkdownDescription: "VolumeSpec defines the desired state of the Longhorn volume",
				Attributes: map[string]schema.Attribute{
					"standby": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"access_mode": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"backend_store_driver": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"backing_image": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"backup_compression_method": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"data_locality": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"data_source": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"disable_frontend": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"disk_selector": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"encrypted": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"engine_image": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"from_backup": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"frontend": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"last_attached_by": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"migratable": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"migration_node_id": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"node_id": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"node_selector": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"number_of_replicas": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"offline_replica_rebuilding": schema.StringAttribute{
						Description:         "OfflineReplicaRebuilding is used to determine if the offline replica rebuilding feature is enabled or not",
						MarkdownDescription: "OfflineReplicaRebuilding is used to determine if the offline replica rebuilding feature is enabled or not",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"replica_auto_balance": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"replica_disk_soft_anti_affinity": schema.StringAttribute{
						Description:         "Replica disk soft anti affinity of the volume. Set enabled to allow replicas to be scheduled in the same disk.",
						MarkdownDescription: "Replica disk soft anti affinity of the volume. Set enabled to allow replicas to be scheduled in the same disk.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"replica_soft_anti_affinity": schema.StringAttribute{
						Description:         "Replica soft anti affinity of the volume. Set enabled to allow replicas to be scheduled on the same node.",
						MarkdownDescription: "Replica soft anti affinity of the volume. Set enabled to allow replicas to be scheduled on the same node.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"replica_zone_soft_anti_affinity": schema.StringAttribute{
						Description:         "Replica zone soft anti affinity of the volume. Set enabled to allow replicas to be scheduled in the same zone.",
						MarkdownDescription: "Replica zone soft anti affinity of the volume. Set enabled to allow replicas to be scheduled in the same zone.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"restore_volume_recurring_job": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"revision_counter_disabled": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"size": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"snapshot_data_integrity": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"stale_replica_timeout": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"unmap_mark_snap_chain_removed": schema.StringAttribute{
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
		},
	}
}

func (r *LonghornIoVolumeV1Beta2DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *LonghornIoVolumeV1Beta2DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_longhorn_io_volume_v1beta2")

	var data LonghornIoVolumeV1Beta2DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "longhorn.io", Version: "v1beta2", Resource: "volumes"}).
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

	var readResponse LonghornIoVolumeV1Beta2DataSourceData
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
	data.ApiVersion = pointer.String("longhorn.io/v1beta2")
	data.Kind = pointer.String("Volume")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
