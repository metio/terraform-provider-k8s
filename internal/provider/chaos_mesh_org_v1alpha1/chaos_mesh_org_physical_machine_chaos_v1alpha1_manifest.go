/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package chaos_mesh_org_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &ChaosMeshOrgPhysicalMachineChaosV1Alpha1Manifest{}
)

func NewChaosMeshOrgPhysicalMachineChaosV1Alpha1Manifest() datasource.DataSource {
	return &ChaosMeshOrgPhysicalMachineChaosV1Alpha1Manifest{}
}

type ChaosMeshOrgPhysicalMachineChaosV1Alpha1Manifest struct{}

type ChaosMeshOrgPhysicalMachineChaosV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Action  *string   `tfsdk:"action" json:"action,omitempty"`
		Address *[]string `tfsdk:"address" json:"address,omitempty"`
		Clock   *struct {
			Clock_ids_slice *string `tfsdk:"clock_ids_slice" json:"clock-ids-slice,omitempty"`
			Pid             *int64  `tfsdk:"pid" json:"pid,omitempty"`
			Time_offset     *string `tfsdk:"time_offset" json:"time-offset,omitempty"`
		} `tfsdk:"clock" json:"clock,omitempty"`
		Disk_fill *struct {
			Fill_by_fallocate *bool   `tfsdk:"fill_by_fallocate" json:"fill-by-fallocate,omitempty"`
			Path              *string `tfsdk:"path" json:"path,omitempty"`
			Size              *string `tfsdk:"size" json:"size,omitempty"`
		} `tfsdk:"disk_fill" json:"disk-fill,omitempty"`
		Disk_read_payload *struct {
			Path                *string `tfsdk:"path" json:"path,omitempty"`
			Payload_process_num *int64  `tfsdk:"payload_process_num" json:"payload-process-num,omitempty"`
			Size                *string `tfsdk:"size" json:"size,omitempty"`
		} `tfsdk:"disk_read_payload" json:"disk-read-payload,omitempty"`
		Disk_write_payload *struct {
			Path                *string `tfsdk:"path" json:"path,omitempty"`
			Payload_process_num *int64  `tfsdk:"payload_process_num" json:"payload-process-num,omitempty"`
			Size                *string `tfsdk:"size" json:"size,omitempty"`
		} `tfsdk:"disk_write_payload" json:"disk-write-payload,omitempty"`
		Duration    *string `tfsdk:"duration" json:"duration,omitempty"`
		File_append *struct {
			Count     *int64  `tfsdk:"count" json:"count,omitempty"`
			Data      *string `tfsdk:"data" json:"data,omitempty"`
			File_name *string `tfsdk:"file_name" json:"file-name,omitempty"`
		} `tfsdk:"file_append" json:"file-append,omitempty"`
		File_create *struct {
			Dir_name  *string `tfsdk:"dir_name" json:"dir-name,omitempty"`
			File_name *string `tfsdk:"file_name" json:"file-name,omitempty"`
		} `tfsdk:"file_create" json:"file-create,omitempty"`
		File_delete *struct {
			Dir_name  *string `tfsdk:"dir_name" json:"dir-name,omitempty"`
			File_name *string `tfsdk:"file_name" json:"file-name,omitempty"`
		} `tfsdk:"file_delete" json:"file-delete,omitempty"`
		File_modify *struct {
			File_name *string `tfsdk:"file_name" json:"file-name,omitempty"`
			Privilege *int64  `tfsdk:"privilege" json:"privilege,omitempty"`
		} `tfsdk:"file_modify" json:"file-modify,omitempty"`
		File_rename *struct {
			Dest_file   *string `tfsdk:"dest_file" json:"dest-file,omitempty"`
			Source_file *string `tfsdk:"source_file" json:"source-file,omitempty"`
		} `tfsdk:"file_rename" json:"file-rename,omitempty"`
		File_replace *struct {
			Dest_string   *string `tfsdk:"dest_string" json:"dest-string,omitempty"`
			File_name     *string `tfsdk:"file_name" json:"file-name,omitempty"`
			Line          *int64  `tfsdk:"line" json:"line,omitempty"`
			Origin_string *string `tfsdk:"origin_string" json:"origin-string,omitempty"`
		} `tfsdk:"file_replace" json:"file-replace,omitempty"`
		Http_abort *struct {
			Code        *string   `tfsdk:"code" json:"code,omitempty"`
			Method      *string   `tfsdk:"method" json:"method,omitempty"`
			Path        *string   `tfsdk:"path" json:"path,omitempty"`
			Port        *int64    `tfsdk:"port" json:"port,omitempty"`
			Proxy_ports *[]string `tfsdk:"proxy_ports" json:"proxy_ports,omitempty"`
			Target      *string   `tfsdk:"target" json:"target,omitempty"`
		} `tfsdk:"http_abort" json:"http-abort,omitempty"`
		Http_config *struct {
			File_path *string `tfsdk:"file_path" json:"file_path,omitempty"`
		} `tfsdk:"http_config" json:"http-config,omitempty"`
		Http_delay *struct {
			Code        *string   `tfsdk:"code" json:"code,omitempty"`
			Delay       *string   `tfsdk:"delay" json:"delay,omitempty"`
			Method      *string   `tfsdk:"method" json:"method,omitempty"`
			Path        *string   `tfsdk:"path" json:"path,omitempty"`
			Port        *int64    `tfsdk:"port" json:"port,omitempty"`
			Proxy_ports *[]string `tfsdk:"proxy_ports" json:"proxy_ports,omitempty"`
			Target      *string   `tfsdk:"target" json:"target,omitempty"`
		} `tfsdk:"http_delay" json:"http-delay,omitempty"`
		Http_request *struct {
			Count            *int64  `tfsdk:"count" json:"count,omitempty"`
			Enable_conn_pool *bool   `tfsdk:"enable_conn_pool" json:"enable-conn-pool,omitempty"`
			Url              *string `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"http_request" json:"http-request,omitempty"`
		Jvm_exception *struct {
			Class     *string `tfsdk:"class" json:"class,omitempty"`
			Exception *string `tfsdk:"exception" json:"exception,omitempty"`
			Method    *string `tfsdk:"method" json:"method,omitempty"`
			Pid       *int64  `tfsdk:"pid" json:"pid,omitempty"`
			Port      *int64  `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"jvm_exception" json:"jvm-exception,omitempty"`
		Jvm_gc *struct {
			Pid  *int64 `tfsdk:"pid" json:"pid,omitempty"`
			Port *int64 `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"jvm_gc" json:"jvm-gc,omitempty"`
		Jvm_latency *struct {
			Class   *string `tfsdk:"class" json:"class,omitempty"`
			Latency *int64  `tfsdk:"latency" json:"latency,omitempty"`
			Method  *string `tfsdk:"method" json:"method,omitempty"`
			Pid     *int64  `tfsdk:"pid" json:"pid,omitempty"`
			Port    *int64  `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"jvm_latency" json:"jvm-latency,omitempty"`
		Jvm_mysql *struct {
			Database              *string `tfsdk:"database" json:"database,omitempty"`
			Exception             *string `tfsdk:"exception" json:"exception,omitempty"`
			Latency               *int64  `tfsdk:"latency" json:"latency,omitempty"`
			MysqlConnectorVersion *string `tfsdk:"mysql_connector_version" json:"mysqlConnectorVersion,omitempty"`
			Pid                   *int64  `tfsdk:"pid" json:"pid,omitempty"`
			Port                  *int64  `tfsdk:"port" json:"port,omitempty"`
			SqlType               *string `tfsdk:"sql_type" json:"sqlType,omitempty"`
			Table                 *string `tfsdk:"table" json:"table,omitempty"`
		} `tfsdk:"jvm_mysql" json:"jvm-mysql,omitempty"`
		Jvm_return *struct {
			Class  *string `tfsdk:"class" json:"class,omitempty"`
			Method *string `tfsdk:"method" json:"method,omitempty"`
			Pid    *int64  `tfsdk:"pid" json:"pid,omitempty"`
			Port   *int64  `tfsdk:"port" json:"port,omitempty"`
			Value  *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"jvm_return" json:"jvm-return,omitempty"`
		Jvm_rule_data *struct {
			Pid       *int64  `tfsdk:"pid" json:"pid,omitempty"`
			Port      *int64  `tfsdk:"port" json:"port,omitempty"`
			Rule_data *string `tfsdk:"rule_data" json:"rule-data,omitempty"`
		} `tfsdk:"jvm_rule_data" json:"jvm-rule-data,omitempty"`
		Jvm_stress *struct {
			Cpu_count *int64  `tfsdk:"cpu_count" json:"cpu-count,omitempty"`
			Mem_type  *string `tfsdk:"mem_type" json:"mem-type,omitempty"`
			Pid       *int64  `tfsdk:"pid" json:"pid,omitempty"`
			Port      *int64  `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"jvm_stress" json:"jvm-stress,omitempty"`
		Kafka_fill *struct {
			Host          *string `tfsdk:"host" json:"host,omitempty"`
			MaxBytes      *int64  `tfsdk:"max_bytes" json:"maxBytes,omitempty"`
			MessageSize   *int64  `tfsdk:"message_size" json:"messageSize,omitempty"`
			Password      *string `tfsdk:"password" json:"password,omitempty"`
			Port          *int64  `tfsdk:"port" json:"port,omitempty"`
			ReloadCommand *string `tfsdk:"reload_command" json:"reloadCommand,omitempty"`
			Topic         *string `tfsdk:"topic" json:"topic,omitempty"`
			Username      *string `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"kafka_fill" json:"kafka-fill,omitempty"`
		Kafka_flood *struct {
			Host        *string `tfsdk:"host" json:"host,omitempty"`
			MessageSize *int64  `tfsdk:"message_size" json:"messageSize,omitempty"`
			Password    *string `tfsdk:"password" json:"password,omitempty"`
			Port        *int64  `tfsdk:"port" json:"port,omitempty"`
			Threads     *int64  `tfsdk:"threads" json:"threads,omitempty"`
			Topic       *string `tfsdk:"topic" json:"topic,omitempty"`
			Username    *string `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"kafka_flood" json:"kafka-flood,omitempty"`
		Kafka_io *struct {
			ConfigFile  *string `tfsdk:"config_file" json:"configFile,omitempty"`
			NonReadable *bool   `tfsdk:"non_readable" json:"nonReadable,omitempty"`
			NonWritable *bool   `tfsdk:"non_writable" json:"nonWritable,omitempty"`
			Topic       *string `tfsdk:"topic" json:"topic,omitempty"`
		} `tfsdk:"kafka_io" json:"kafka-io,omitempty"`
		Mode              *string `tfsdk:"mode" json:"mode,omitempty"`
		Network_bandwidth *struct {
			Buffer     *int64  `tfsdk:"buffer" json:"buffer,omitempty"`
			Device     *string `tfsdk:"device" json:"device,omitempty"`
			Hostname   *string `tfsdk:"hostname" json:"hostname,omitempty"`
			Ip_address *string `tfsdk:"ip_address" json:"ip-address,omitempty"`
			Limit      *int64  `tfsdk:"limit" json:"limit,omitempty"`
			Minburst   *int64  `tfsdk:"minburst" json:"minburst,omitempty"`
			Peakrate   *int64  `tfsdk:"peakrate" json:"peakrate,omitempty"`
			Rate       *string `tfsdk:"rate" json:"rate,omitempty"`
		} `tfsdk:"network_bandwidth" json:"network-bandwidth,omitempty"`
		Network_corrupt *struct {
			Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
			Device      *string `tfsdk:"device" json:"device,omitempty"`
			Egress_port *string `tfsdk:"egress_port" json:"egress-port,omitempty"`
			Hostname    *string `tfsdk:"hostname" json:"hostname,omitempty"`
			Ip_address  *string `tfsdk:"ip_address" json:"ip-address,omitempty"`
			Ip_protocol *string `tfsdk:"ip_protocol" json:"ip-protocol,omitempty"`
			Percent     *string `tfsdk:"percent" json:"percent,omitempty"`
			Source_port *string `tfsdk:"source_port" json:"source-port,omitempty"`
		} `tfsdk:"network_corrupt" json:"network-corrupt,omitempty"`
		Network_delay *struct {
			Accept_tcp_flags *string `tfsdk:"accept_tcp_flags" json:"accept-tcp-flags,omitempty"`
			Correlation      *string `tfsdk:"correlation" json:"correlation,omitempty"`
			Device           *string `tfsdk:"device" json:"device,omitempty"`
			Egress_port      *string `tfsdk:"egress_port" json:"egress-port,omitempty"`
			Hostname         *string `tfsdk:"hostname" json:"hostname,omitempty"`
			Ip_address       *string `tfsdk:"ip_address" json:"ip-address,omitempty"`
			Ip_protocol      *string `tfsdk:"ip_protocol" json:"ip-protocol,omitempty"`
			Jitter           *string `tfsdk:"jitter" json:"jitter,omitempty"`
			Latency          *string `tfsdk:"latency" json:"latency,omitempty"`
			Source_port      *string `tfsdk:"source_port" json:"source-port,omitempty"`
		} `tfsdk:"network_delay" json:"network-delay,omitempty"`
		Network_dns *struct {
			Dns_domain_name *string `tfsdk:"dns_domain_name" json:"dns-domain-name,omitempty"`
			Dns_ip          *string `tfsdk:"dns_ip" json:"dns-ip,omitempty"`
			Dns_server      *string `tfsdk:"dns_server" json:"dns-server,omitempty"`
		} `tfsdk:"network_dns" json:"network-dns,omitempty"`
		Network_down *struct {
			Device   *string `tfsdk:"device" json:"device,omitempty"`
			Duration *string `tfsdk:"duration" json:"duration,omitempty"`
		} `tfsdk:"network_down" json:"network-down,omitempty"`
		Network_duplicate *struct {
			Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
			Device      *string `tfsdk:"device" json:"device,omitempty"`
			Egress_port *string `tfsdk:"egress_port" json:"egress-port,omitempty"`
			Hostname    *string `tfsdk:"hostname" json:"hostname,omitempty"`
			Ip_address  *string `tfsdk:"ip_address" json:"ip-address,omitempty"`
			Ip_protocol *string `tfsdk:"ip_protocol" json:"ip-protocol,omitempty"`
			Percent     *string `tfsdk:"percent" json:"percent,omitempty"`
			Source_port *string `tfsdk:"source_port" json:"source-port,omitempty"`
		} `tfsdk:"network_duplicate" json:"network-duplicate,omitempty"`
		Network_flood *struct {
			Duration   *string `tfsdk:"duration" json:"duration,omitempty"`
			Ip_address *string `tfsdk:"ip_address" json:"ip-address,omitempty"`
			Parallel   *int64  `tfsdk:"parallel" json:"parallel,omitempty"`
			Port       *string `tfsdk:"port" json:"port,omitempty"`
			Rate       *string `tfsdk:"rate" json:"rate,omitempty"`
		} `tfsdk:"network_flood" json:"network-flood,omitempty"`
		Network_loss *struct {
			Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
			Device      *string `tfsdk:"device" json:"device,omitempty"`
			Egress_port *string `tfsdk:"egress_port" json:"egress-port,omitempty"`
			Hostname    *string `tfsdk:"hostname" json:"hostname,omitempty"`
			Ip_address  *string `tfsdk:"ip_address" json:"ip-address,omitempty"`
			Ip_protocol *string `tfsdk:"ip_protocol" json:"ip-protocol,omitempty"`
			Percent     *string `tfsdk:"percent" json:"percent,omitempty"`
			Source_port *string `tfsdk:"source_port" json:"source-port,omitempty"`
		} `tfsdk:"network_loss" json:"network-loss,omitempty"`
		Network_partition *struct {
			Accept_tcp_flags *string `tfsdk:"accept_tcp_flags" json:"accept-tcp-flags,omitempty"`
			Device           *string `tfsdk:"device" json:"device,omitempty"`
			Direction        *string `tfsdk:"direction" json:"direction,omitempty"`
			Hostname         *string `tfsdk:"hostname" json:"hostname,omitempty"`
			Ip_address       *string `tfsdk:"ip_address" json:"ip-address,omitempty"`
			Ip_protocol      *string `tfsdk:"ip_protocol" json:"ip-protocol,omitempty"`
		} `tfsdk:"network_partition" json:"network-partition,omitempty"`
		Process *struct {
			Process    *string `tfsdk:"process" json:"process,omitempty"`
			RecoverCmd *string `tfsdk:"recover_cmd" json:"recoverCmd,omitempty"`
			Signal     *int64  `tfsdk:"signal" json:"signal,omitempty"`
		} `tfsdk:"process" json:"process,omitempty"`
		Redis_cacheLimit *struct {
			Addr      *string `tfsdk:"addr" json:"addr,omitempty"`
			CacheSize *string `tfsdk:"cache_size" json:"cacheSize,omitempty"`
			Password  *string `tfsdk:"password" json:"password,omitempty"`
			Percent   *string `tfsdk:"percent" json:"percent,omitempty"`
		} `tfsdk:"redis_cache_limit" json:"redis-cacheLimit,omitempty"`
		Redis_expiration *struct {
			Addr       *string `tfsdk:"addr" json:"addr,omitempty"`
			Expiration *string `tfsdk:"expiration" json:"expiration,omitempty"`
			Key        *string `tfsdk:"key" json:"key,omitempty"`
			Option     *string `tfsdk:"option" json:"option,omitempty"`
			Password   *string `tfsdk:"password" json:"password,omitempty"`
		} `tfsdk:"redis_expiration" json:"redis-expiration,omitempty"`
		Redis_penetration *struct {
			Addr       *string `tfsdk:"addr" json:"addr,omitempty"`
			Password   *string `tfsdk:"password" json:"password,omitempty"`
			RequestNum *int64  `tfsdk:"request_num" json:"requestNum,omitempty"`
		} `tfsdk:"redis_penetration" json:"redis-penetration,omitempty"`
		Redis_restart *struct {
			Addr        *string `tfsdk:"addr" json:"addr,omitempty"`
			Conf        *string `tfsdk:"conf" json:"conf,omitempty"`
			FlushConfig *bool   `tfsdk:"flush_config" json:"flushConfig,omitempty"`
			Password    *string `tfsdk:"password" json:"password,omitempty"`
			RedisPath   *bool   `tfsdk:"redis_path" json:"redisPath,omitempty"`
		} `tfsdk:"redis_restart" json:"redis-restart,omitempty"`
		Redis_stop *struct {
			Addr        *string `tfsdk:"addr" json:"addr,omitempty"`
			Conf        *string `tfsdk:"conf" json:"conf,omitempty"`
			FlushConfig *bool   `tfsdk:"flush_config" json:"flushConfig,omitempty"`
			Password    *string `tfsdk:"password" json:"password,omitempty"`
			RedisPath   *bool   `tfsdk:"redis_path" json:"redisPath,omitempty"`
		} `tfsdk:"redis_stop" json:"redis-stop,omitempty"`
		RemoteCluster *string `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
		Selector      *struct {
			AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
			ExpressionSelectors *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
			FieldSelectors   *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
			LabelSelectors   *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
			Namespaces       *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
			PhysicalMachines *map[string][]string `tfsdk:"physical_machines" json:"physicalMachines,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
		Stress_cpu *struct {
			Load    *int64    `tfsdk:"load" json:"load,omitempty"`
			Options *[]string `tfsdk:"options" json:"options,omitempty"`
			Workers *int64    `tfsdk:"workers" json:"workers,omitempty"`
		} `tfsdk:"stress_cpu" json:"stress-cpu,omitempty"`
		Stress_mem *struct {
			Options *[]string `tfsdk:"options" json:"options,omitempty"`
			Size    *string   `tfsdk:"size" json:"size,omitempty"`
		} `tfsdk:"stress_mem" json:"stress-mem,omitempty"`
		Uid          *string `tfsdk:"uid" json:"uid,omitempty"`
		User_defined *struct {
			AttackCmd  *string `tfsdk:"attack_cmd" json:"attackCmd,omitempty"`
			RecoverCmd *string `tfsdk:"recover_cmd" json:"recoverCmd,omitempty"`
		} `tfsdk:"user_defined" json:"user_defined,omitempty"`
		Value *string `tfsdk:"value" json:"value,omitempty"`
		Vm    *struct {
			Vm_name *string `tfsdk:"vm_name" json:"vm-name,omitempty"`
		} `tfsdk:"vm" json:"vm,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ChaosMeshOrgPhysicalMachineChaosV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_chaos_mesh_org_physical_machine_chaos_v1alpha1_manifest"
}

func (r *ChaosMeshOrgPhysicalMachineChaosV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PhysicalMachineChaos is the Schema for the physical machine chaos API",
		MarkdownDescription: "PhysicalMachineChaos is the Schema for the physical machine chaos API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"yaml": schema.StringAttribute{
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
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
						Optional:            true,
						Computed:            false,
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
						Computed:            false,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "Spec defines the behavior of a physical machine chaos experiment",
				MarkdownDescription: "Spec defines the behavior of a physical machine chaos experiment",
				Attributes: map[string]schema.Attribute{
					"action": schema.StringAttribute{
						Description:         "the subAction, generate automatically",
						MarkdownDescription: "the subAction, generate automatically",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("stress-cpu", "stress-mem", "disk-read-payload", "disk-write-payload", "disk-fill", "network-corrupt", "network-duplicate", "network-loss", "network-delay", "network-partition", "network-dns", "network-bandwidth", "network-flood", "network-down", "process", "jvm-exception", "jvm-gc", "jvm-latency", "jvm-return", "jvm-stress", "jvm-rule-data", "jvm-mysql", "clock", "redis-expiration", "redis-penetration", "redis-cacheLimit", "redis-restart", "redis-stop", "kafka-fill", "kafka-flood", "kafka-io", "file-create", "file-modify", "file-delete", "file-rename", "file-append", "file-replace", "vm", "user_defined"),
						},
					},

					"address": schema.ListAttribute{
						Description:         "DEPRECATED: Use Selector instead. Only one of Address and Selector could be specified.",
						MarkdownDescription: "DEPRECATED: Use Selector instead. Only one of Address and Selector could be specified.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"clock": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"clock_ids_slice": schema.StringAttribute{
								Description:         "the identifier of the particular clock on which to act. More clock description in linux kernel can be found in man page of clock_getres, clock_gettime, clock_settime. Muti clock ids should be split with ','",
								MarkdownDescription: "the identifier of the particular clock on which to act. More clock description in linux kernel can be found in man page of clock_getres, clock_gettime, clock_settime. Muti clock ids should be split with ','",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pid": schema.Int64Attribute{
								Description:         "the pid of target program.",
								MarkdownDescription: "the pid of target program.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_offset": schema.StringAttribute{
								Description:         "specifies the length of time offset.",
								MarkdownDescription: "specifies the length of time offset.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"disk_fill": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"fill_by_fallocate": schema.BoolAttribute{
								Description:         "fill disk by fallocate",
								MarkdownDescription: "fill disk by fallocate",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"path": schema.StringAttribute{
								Description:         "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
								MarkdownDescription: "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"size": schema.StringAttribute{
								Description:         "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
								MarkdownDescription: "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"disk_read_payload": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"path": schema.StringAttribute{
								Description:         "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
								MarkdownDescription: "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"payload_process_num": schema.Int64Attribute{
								Description:         "specifies the number of process work on writing, default 1, only 1-255 is valid value",
								MarkdownDescription: "specifies the number of process work on writing, default 1, only 1-255 is valid value",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"size": schema.StringAttribute{
								Description:         "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
								MarkdownDescription: "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"disk_write_payload": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"path": schema.StringAttribute{
								Description:         "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
								MarkdownDescription: "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"payload_process_num": schema.Int64Attribute{
								Description:         "specifies the number of process work on writing, default 1, only 1-255 is valid value",
								MarkdownDescription: "specifies the number of process work on writing, default 1, only 1-255 is valid value",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"size": schema.StringAttribute{
								Description:         "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
								MarkdownDescription: "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"duration": schema.StringAttribute{
						Description:         "Duration represents the duration of the chaos action",
						MarkdownDescription: "Duration represents the duration of the chaos action",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"file_append": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"count": schema.Int64Attribute{
								Description:         "Count is the number of times to append the data.",
								MarkdownDescription: "Count is the number of times to append the data.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"data": schema.StringAttribute{
								Description:         "Data is the data for append.",
								MarkdownDescription: "Data is the data for append.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"file_name": schema.StringAttribute{
								Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
								MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"file_create": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"dir_name": schema.StringAttribute{
								Description:         "DirName is the directory name to create or delete.",
								MarkdownDescription: "DirName is the directory name to create or delete.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"file_name": schema.StringAttribute{
								Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
								MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"file_delete": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"dir_name": schema.StringAttribute{
								Description:         "DirName is the directory name to create or delete.",
								MarkdownDescription: "DirName is the directory name to create or delete.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"file_name": schema.StringAttribute{
								Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
								MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"file_modify": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"file_name": schema.StringAttribute{
								Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
								MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"privilege": schema.Int64Attribute{
								Description:         "Privilege is the file privilege to be set.",
								MarkdownDescription: "Privilege is the file privilege to be set.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"file_rename": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"dest_file": schema.StringAttribute{
								Description:         "DestFile is the name to be renamed.",
								MarkdownDescription: "DestFile is the name to be renamed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"source_file": schema.StringAttribute{
								Description:         "SourceFile is the name need to be renamed.",
								MarkdownDescription: "SourceFile is the name need to be renamed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"file_replace": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"dest_string": schema.StringAttribute{
								Description:         "DestStr is the destination string of the file.",
								MarkdownDescription: "DestStr is the destination string of the file.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"file_name": schema.StringAttribute{
								Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
								MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"line": schema.Int64Attribute{
								Description:         "Line is the line number of the file to be replaced.",
								MarkdownDescription: "Line is the line number of the file to be replaced.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"origin_string": schema.StringAttribute{
								Description:         "OriginStr is the origin string of the file.",
								MarkdownDescription: "OriginStr is the origin string of the file.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"http_abort": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"code": schema.StringAttribute{
								Description:         "Code is a rule to select target by http status code in response",
								MarkdownDescription: "Code is a rule to select target by http status code in response",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"method": schema.StringAttribute{
								Description:         "HTTP method",
								MarkdownDescription: "HTTP method",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"path": schema.StringAttribute{
								Description:         "Match path of Uri with wildcard matches",
								MarkdownDescription: "Match path of Uri with wildcard matches",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "The TCP port that the target service listens on",
								MarkdownDescription: "The TCP port that the target service listens on",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_ports": schema.ListAttribute{
								Description:         "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",
								MarkdownDescription: "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",
								ElementType:         types.StringType,
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"target": schema.StringAttribute{
								Description:         "HTTP target: Request or Response",
								MarkdownDescription: "HTTP target: Request or Response",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"http_config": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"file_path": schema.StringAttribute{
								Description:         "The config file path",
								MarkdownDescription: "The config file path",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"http_delay": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"code": schema.StringAttribute{
								Description:         "Code is a rule to select target by http status code in response",
								MarkdownDescription: "Code is a rule to select target by http status code in response",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"delay": schema.StringAttribute{
								Description:         "Delay represents the delay of the target request/response",
								MarkdownDescription: "Delay represents the delay of the target request/response",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"method": schema.StringAttribute{
								Description:         "HTTP method",
								MarkdownDescription: "HTTP method",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"path": schema.StringAttribute{
								Description:         "Match path of Uri with wildcard matches",
								MarkdownDescription: "Match path of Uri with wildcard matches",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "The TCP port that the target service listens on",
								MarkdownDescription: "The TCP port that the target service listens on",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_ports": schema.ListAttribute{
								Description:         "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",
								MarkdownDescription: "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",
								ElementType:         types.StringType,
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"target": schema.StringAttribute{
								Description:         "HTTP target: Request or Response",
								MarkdownDescription: "HTTP target: Request or Response",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"http_request": schema.SingleNestedAttribute{
						Description:         "used for HTTP request, now only support GET",
						MarkdownDescription: "used for HTTP request, now only support GET",
						Attributes: map[string]schema.Attribute{
							"count": schema.Int64Attribute{
								Description:         "The number of requests to send",
								MarkdownDescription: "The number of requests to send",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_conn_pool": schema.BoolAttribute{
								Description:         "Enable connection pool",
								MarkdownDescription: "Enable connection pool",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"url": schema.StringAttribute{
								Description:         "Request to send'",
								MarkdownDescription: "Request to send'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"jvm_exception": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"class": schema.StringAttribute{
								Description:         "Java class",
								MarkdownDescription: "Java class",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"exception": schema.StringAttribute{
								Description:         "the exception which needs to throw for action 'exception'",
								MarkdownDescription: "the exception which needs to throw for action 'exception'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"method": schema.StringAttribute{
								Description:         "the method in Java class",
								MarkdownDescription: "the method in Java class",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pid": schema.Int64Attribute{
								Description:         "the pid of Java process which needs to attach",
								MarkdownDescription: "the pid of Java process which needs to attach",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "the port of agent server, default 9277",
								MarkdownDescription: "the port of agent server, default 9277",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"jvm_gc": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"pid": schema.Int64Attribute{
								Description:         "the pid of Java process which needs to attach",
								MarkdownDescription: "the pid of Java process which needs to attach",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "the port of agent server, default 9277",
								MarkdownDescription: "the port of agent server, default 9277",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"jvm_latency": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"class": schema.StringAttribute{
								Description:         "Java class",
								MarkdownDescription: "Java class",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"latency": schema.Int64Attribute{
								Description:         "the latency duration for action 'latency', unit ms",
								MarkdownDescription: "the latency duration for action 'latency', unit ms",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"method": schema.StringAttribute{
								Description:         "the method in Java class",
								MarkdownDescription: "the method in Java class",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pid": schema.Int64Attribute{
								Description:         "the pid of Java process which needs to attach",
								MarkdownDescription: "the pid of Java process which needs to attach",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "the port of agent server, default 9277",
								MarkdownDescription: "the port of agent server, default 9277",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"jvm_mysql": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"database": schema.StringAttribute{
								Description:         "the match database default value is '', means match all database",
								MarkdownDescription: "the match database default value is '', means match all database",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"exception": schema.StringAttribute{
								Description:         "The exception which needs to throw for action 'exception' or the exception message needs to throw in action 'mysql'",
								MarkdownDescription: "The exception which needs to throw for action 'exception' or the exception message needs to throw in action 'mysql'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"latency": schema.Int64Attribute{
								Description:         "The latency duration for action 'latency' or the latency duration in action 'mysql'",
								MarkdownDescription: "The latency duration for action 'latency' or the latency duration in action 'mysql'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mysql_connector_version": schema.StringAttribute{
								Description:         "the version of mysql-connector-java, only support 5.X.X(set to '5') and 8.X.X(set to '8') now",
								MarkdownDescription: "the version of mysql-connector-java, only support 5.X.X(set to '5') and 8.X.X(set to '8') now",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pid": schema.Int64Attribute{
								Description:         "the pid of Java process which needs to attach",
								MarkdownDescription: "the pid of Java process which needs to attach",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "the port of agent server, default 9277",
								MarkdownDescription: "the port of agent server, default 9277",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sql_type": schema.StringAttribute{
								Description:         "the match sql type default value is '', means match all SQL type. The value can be 'select', 'insert', 'update', 'delete', 'replace'.",
								MarkdownDescription: "the match sql type default value is '', means match all SQL type. The value can be 'select', 'insert', 'update', 'delete', 'replace'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"table": schema.StringAttribute{
								Description:         "the match table default value is '', means match all table",
								MarkdownDescription: "the match table default value is '', means match all table",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"jvm_return": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"class": schema.StringAttribute{
								Description:         "Java class",
								MarkdownDescription: "Java class",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"method": schema.StringAttribute{
								Description:         "the method in Java class",
								MarkdownDescription: "the method in Java class",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pid": schema.Int64Attribute{
								Description:         "the pid of Java process which needs to attach",
								MarkdownDescription: "the pid of Java process which needs to attach",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "the port of agent server, default 9277",
								MarkdownDescription: "the port of agent server, default 9277",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"value": schema.StringAttribute{
								Description:         "the return value for action 'return'",
								MarkdownDescription: "the return value for action 'return'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"jvm_rule_data": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"pid": schema.Int64Attribute{
								Description:         "the pid of Java process which needs to attach",
								MarkdownDescription: "the pid of Java process which needs to attach",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "the port of agent server, default 9277",
								MarkdownDescription: "the port of agent server, default 9277",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rule_data": schema.StringAttribute{
								Description:         "RuleData used to save the rule file's data, will use it when recover",
								MarkdownDescription: "RuleData used to save the rule file's data, will use it when recover",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"jvm_stress": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"cpu_count": schema.Int64Attribute{
								Description:         "the CPU core number need to use, only set it when action is stress",
								MarkdownDescription: "the CPU core number need to use, only set it when action is stress",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mem_type": schema.StringAttribute{
								Description:         "the memory type need to locate, only set it when action is stress, the value can be 'stack' or 'heap'",
								MarkdownDescription: "the memory type need to locate, only set it when action is stress, the value can be 'stack' or 'heap'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pid": schema.Int64Attribute{
								Description:         "the pid of Java process which needs to attach",
								MarkdownDescription: "the pid of Java process which needs to attach",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "the port of agent server, default 9277",
								MarkdownDescription: "the port of agent server, default 9277",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"kafka_fill": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"host": schema.StringAttribute{
								Description:         "The host of kafka server",
								MarkdownDescription: "The host of kafka server",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_bytes": schema.Int64Attribute{
								Description:         "The max bytes to fill",
								MarkdownDescription: "The max bytes to fill",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"message_size": schema.Int64Attribute{
								Description:         "The size of each message",
								MarkdownDescription: "The size of each message",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"password": schema.StringAttribute{
								Description:         "The password of kafka client",
								MarkdownDescription: "The password of kafka client",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "The port of kafka server",
								MarkdownDescription: "The port of kafka server",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"reload_command": schema.StringAttribute{
								Description:         "The command to reload kafka config",
								MarkdownDescription: "The command to reload kafka config",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"topic": schema.StringAttribute{
								Description:         "The topic to attack",
								MarkdownDescription: "The topic to attack",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"username": schema.StringAttribute{
								Description:         "The username of kafka client",
								MarkdownDescription: "The username of kafka client",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"kafka_flood": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"host": schema.StringAttribute{
								Description:         "The host of kafka server",
								MarkdownDescription: "The host of kafka server",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"message_size": schema.Int64Attribute{
								Description:         "The size of each message",
								MarkdownDescription: "The size of each message",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"password": schema.StringAttribute{
								Description:         "The password of kafka client",
								MarkdownDescription: "The password of kafka client",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "The port of kafka server",
								MarkdownDescription: "The port of kafka server",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"threads": schema.Int64Attribute{
								Description:         "The number of worker threads",
								MarkdownDescription: "The number of worker threads",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"topic": schema.StringAttribute{
								Description:         "The topic to attack",
								MarkdownDescription: "The topic to attack",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"username": schema.StringAttribute{
								Description:         "The username of kafka client",
								MarkdownDescription: "The username of kafka client",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"kafka_io": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"config_file": schema.StringAttribute{
								Description:         "The path of server config",
								MarkdownDescription: "The path of server config",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"non_readable": schema.BoolAttribute{
								Description:         "Make kafka cluster non-readable",
								MarkdownDescription: "Make kafka cluster non-readable",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"non_writable": schema.BoolAttribute{
								Description:         "Make kafka cluster non-writable",
								MarkdownDescription: "Make kafka cluster non-writable",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"topic": schema.StringAttribute{
								Description:         "The topic to attack",
								MarkdownDescription: "The topic to attack",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"mode": schema.StringAttribute{
						Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
						MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
						},
					},

					"network_bandwidth": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"buffer": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"device": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"hostname": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ip_address": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"limit": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"minburst": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"peakrate": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rate": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"network_corrupt": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"correlation": schema.StringAttribute{
								Description:         "correlation is percentage (10 is 10%)",
								MarkdownDescription: "correlation is percentage (10 is 10%)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"device": schema.StringAttribute{
								Description:         "the network interface to impact",
								MarkdownDescription: "the network interface to impact",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"egress_port": schema.StringAttribute{
								Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"hostname": schema.StringAttribute{
								Description:         "only impact traffic to these hostnames",
								MarkdownDescription: "only impact traffic to these hostnames",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ip_address": schema.StringAttribute{
								Description:         "only impact egress traffic to these IP addresses",
								MarkdownDescription: "only impact egress traffic to these IP addresses",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ip_protocol": schema.StringAttribute{
								Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
								MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"percent": schema.StringAttribute{
								Description:         "percentage of packets to corrupt (10 is 10%)",
								MarkdownDescription: "percentage of packets to corrupt (10 is 10%)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"source_port": schema.StringAttribute{
								Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"network_delay": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"accept_tcp_flags": schema.StringAttribute{
								Description:         "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",
								MarkdownDescription: "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"correlation": schema.StringAttribute{
								Description:         "correlation is percentage (10 is 10%)",
								MarkdownDescription: "correlation is percentage (10 is 10%)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"device": schema.StringAttribute{
								Description:         "the network interface to impact",
								MarkdownDescription: "the network interface to impact",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"egress_port": schema.StringAttribute{
								Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"hostname": schema.StringAttribute{
								Description:         "only impact traffic to these hostnames",
								MarkdownDescription: "only impact traffic to these hostnames",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ip_address": schema.StringAttribute{
								Description:         "only impact egress traffic to these IP addresses",
								MarkdownDescription: "only impact egress traffic to these IP addresses",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ip_protocol": schema.StringAttribute{
								Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
								MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"jitter": schema.StringAttribute{
								Description:         "jitter time, time units: ns, us (or µs), ms, s, m, h.",
								MarkdownDescription: "jitter time, time units: ns, us (or µs), ms, s, m, h.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"latency": schema.StringAttribute{
								Description:         "delay egress time, time units: ns, us (or µs), ms, s, m, h.",
								MarkdownDescription: "delay egress time, time units: ns, us (or µs), ms, s, m, h.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"source_port": schema.StringAttribute{
								Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"network_dns": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"dns_domain_name": schema.StringAttribute{
								Description:         "map this host to specified IP",
								MarkdownDescription: "map this host to specified IP",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dns_ip": schema.StringAttribute{
								Description:         "map specified host to this IP address",
								MarkdownDescription: "map specified host to this IP address",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dns_server": schema.StringAttribute{
								Description:         "update the DNS server in /etc/resolv.conf with this value",
								MarkdownDescription: "update the DNS server in /etc/resolv.conf with this value",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"network_down": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"device": schema.StringAttribute{
								Description:         "The network interface to impact",
								MarkdownDescription: "The network interface to impact",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"duration": schema.StringAttribute{
								Description:         "NIC down time, time units: ns, us (or µs), ms, s, m, h.",
								MarkdownDescription: "NIC down time, time units: ns, us (or µs), ms, s, m, h.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"network_duplicate": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"correlation": schema.StringAttribute{
								Description:         "correlation is percentage (10 is 10%)",
								MarkdownDescription: "correlation is percentage (10 is 10%)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"device": schema.StringAttribute{
								Description:         "the network interface to impact",
								MarkdownDescription: "the network interface to impact",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"egress_port": schema.StringAttribute{
								Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"hostname": schema.StringAttribute{
								Description:         "only impact traffic to these hostnames",
								MarkdownDescription: "only impact traffic to these hostnames",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ip_address": schema.StringAttribute{
								Description:         "only impact egress traffic to these IP addresses",
								MarkdownDescription: "only impact egress traffic to these IP addresses",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ip_protocol": schema.StringAttribute{
								Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
								MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"percent": schema.StringAttribute{
								Description:         "percentage of packets to duplicate (10 is 10%)",
								MarkdownDescription: "percentage of packets to duplicate (10 is 10%)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"source_port": schema.StringAttribute{
								Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"network_flood": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"duration": schema.StringAttribute{
								Description:         "The number of seconds to run the iperf test",
								MarkdownDescription: "The number of seconds to run the iperf test",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"ip_address": schema.StringAttribute{
								Description:         "Generate traffic to this IP address",
								MarkdownDescription: "Generate traffic to this IP address",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"parallel": schema.Int64Attribute{
								Description:         "The number of iperf parallel client threads to run",
								MarkdownDescription: "The number of iperf parallel client threads to run",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.StringAttribute{
								Description:         "Generate traffic to this port on the IP address",
								MarkdownDescription: "Generate traffic to this port on the IP address",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rate": schema.StringAttribute{
								Description:         "The speed of network traffic, allows bps, kbps, mbps, gbps, tbps unit. bps means bytes per second",
								MarkdownDescription: "The speed of network traffic, allows bps, kbps, mbps, gbps, tbps unit. bps means bytes per second",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"network_loss": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"correlation": schema.StringAttribute{
								Description:         "correlation is percentage (10 is 10%)",
								MarkdownDescription: "correlation is percentage (10 is 10%)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"device": schema.StringAttribute{
								Description:         "the network interface to impact",
								MarkdownDescription: "the network interface to impact",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"egress_port": schema.StringAttribute{
								Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"hostname": schema.StringAttribute{
								Description:         "only impact traffic to these hostnames",
								MarkdownDescription: "only impact traffic to these hostnames",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ip_address": schema.StringAttribute{
								Description:         "only impact egress traffic to these IP addresses",
								MarkdownDescription: "only impact egress traffic to these IP addresses",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ip_protocol": schema.StringAttribute{
								Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
								MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"percent": schema.StringAttribute{
								Description:         "percentage of packets to loss (10 is 10%)",
								MarkdownDescription: "percentage of packets to loss (10 is 10%)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"source_port": schema.StringAttribute{
								Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"network_partition": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"accept_tcp_flags": schema.StringAttribute{
								Description:         "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",
								MarkdownDescription: "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"device": schema.StringAttribute{
								Description:         "the network interface to impact",
								MarkdownDescription: "the network interface to impact",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"direction": schema.StringAttribute{
								Description:         "specifies the partition direction, values can be 'from', 'to'. 'from' means packets coming from the 'IPAddress' or 'Hostname' and going to your server, 'to' means packets originating from your server and going to the 'IPAddress' or 'Hostname'.",
								MarkdownDescription: "specifies the partition direction, values can be 'from', 'to'. 'from' means packets coming from the 'IPAddress' or 'Hostname' and going to your server, 'to' means packets originating from your server and going to the 'IPAddress' or 'Hostname'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"hostname": schema.StringAttribute{
								Description:         "only impact traffic to these hostnames",
								MarkdownDescription: "only impact traffic to these hostnames",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ip_address": schema.StringAttribute{
								Description:         "only impact egress traffic to these IP addresses",
								MarkdownDescription: "only impact egress traffic to these IP addresses",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ip_protocol": schema.StringAttribute{
								Description:         "only impact egress traffic to these IP addresses",
								MarkdownDescription: "only impact egress traffic to these IP addresses",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"process": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"process": schema.StringAttribute{
								Description:         "the process name or the process ID",
								MarkdownDescription: "the process name or the process ID",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"recover_cmd": schema.StringAttribute{
								Description:         "the command to be run when recovering experiment",
								MarkdownDescription: "the command to be run when recovering experiment",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"signal": schema.Int64Attribute{
								Description:         "the signal number to send",
								MarkdownDescription: "the signal number to send",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"redis_cache_limit": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"addr": schema.StringAttribute{
								Description:         "The adress of Redis server",
								MarkdownDescription: "The adress of Redis server",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cache_size": schema.StringAttribute{
								Description:         "The size of 'maxmemory'",
								MarkdownDescription: "The size of 'maxmemory'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"password": schema.StringAttribute{
								Description:         "The password of Redis server",
								MarkdownDescription: "The password of Redis server",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"percent": schema.StringAttribute{
								Description:         "Specifies maxmemory as a percentage of the original value",
								MarkdownDescription: "Specifies maxmemory as a percentage of the original value",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"redis_expiration": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"addr": schema.StringAttribute{
								Description:         "The adress of Redis server",
								MarkdownDescription: "The adress of Redis server",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"expiration": schema.StringAttribute{
								Description:         "The expiration of the keys",
								MarkdownDescription: "The expiration of the keys",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"key": schema.StringAttribute{
								Description:         "The keys to be expired",
								MarkdownDescription: "The keys to be expired",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"option": schema.StringAttribute{
								Description:         "Additional options for 'expiration'",
								MarkdownDescription: "Additional options for 'expiration'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"password": schema.StringAttribute{
								Description:         "The password of Redis server",
								MarkdownDescription: "The password of Redis server",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"redis_penetration": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"addr": schema.StringAttribute{
								Description:         "The adress of Redis server",
								MarkdownDescription: "The adress of Redis server",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"password": schema.StringAttribute{
								Description:         "The password of Redis server",
								MarkdownDescription: "The password of Redis server",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"request_num": schema.Int64Attribute{
								Description:         "The number of requests to be sent",
								MarkdownDescription: "The number of requests to be sent",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"redis_restart": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"addr": schema.StringAttribute{
								Description:         "The adress of Redis server",
								MarkdownDescription: "The adress of Redis server",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"conf": schema.StringAttribute{
								Description:         "The path of Sentinel conf",
								MarkdownDescription: "The path of Sentinel conf",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"flush_config": schema.BoolAttribute{
								Description:         "The control flag determines whether to flush config",
								MarkdownDescription: "The control flag determines whether to flush config",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"password": schema.StringAttribute{
								Description:         "The password of Redis server",
								MarkdownDescription: "The password of Redis server",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"redis_path": schema.BoolAttribute{
								Description:         "The path of 'redis-server' command-line tool",
								MarkdownDescription: "The path of 'redis-server' command-line tool",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"redis_stop": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"addr": schema.StringAttribute{
								Description:         "The adress of Redis server",
								MarkdownDescription: "The adress of Redis server",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"conf": schema.StringAttribute{
								Description:         "The path of Sentinel conf",
								MarkdownDescription: "The path of Sentinel conf",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"flush_config": schema.BoolAttribute{
								Description:         "The control flag determines whether to flush config",
								MarkdownDescription: "The control flag determines whether to flush config",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"password": schema.StringAttribute{
								Description:         "The password of Redis server",
								MarkdownDescription: "The password of Redis server",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"redis_path": schema.BoolAttribute{
								Description:         "The path of 'redis-server' command-line tool",
								MarkdownDescription: "The path of 'redis-server' command-line tool",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"remote_cluster": schema.StringAttribute{
						Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
						MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"selector": schema.SingleNestedAttribute{
						Description:         "Selector is used to select physical machines that are used to inject chaos action.",
						MarkdownDescription: "Selector is used to select physical machines that are used to inject chaos action.",
						Attributes: map[string]schema.Attribute{
							"annotation_selectors": schema.MapAttribute{
								Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"expression_selectors": schema.ListNestedAttribute{
								Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
								MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											ElementType:         types.StringType,
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

							"field_selectors": schema.MapAttribute{
								Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"label_selectors": schema.MapAttribute{
								Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespaces": schema.ListAttribute{
								Description:         "Namespaces is a set of namespace to which objects belong.",
								MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"physical_machines": schema.MapAttribute{
								Description:         "PhysicalMachines is a map of string keys and a set values that used to select physical machines. The key defines the namespace which physical machine belong, and each value is a set of physical machine names.",
								MarkdownDescription: "PhysicalMachines is a map of string keys and a set values that used to select physical machines. The key defines the namespace which physical machine belong, and each value is a set of physical machine names.",
								ElementType:         types.ListType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"stress_cpu": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"load": schema.Int64Attribute{
								Description:         "specifies P percent loading per CPU worker. 0 is effectively a sleep (no load) and 100 is full loading.",
								MarkdownDescription: "specifies P percent loading per CPU worker. 0 is effectively a sleep (no load) and 100 is full loading.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"options": schema.ListAttribute{
								Description:         "extend stress-ng options",
								MarkdownDescription: "extend stress-ng options",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"workers": schema.Int64Attribute{
								Description:         "specifies N workers to apply the stressor.",
								MarkdownDescription: "specifies N workers to apply the stressor.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"stress_mem": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"options": schema.ListAttribute{
								Description:         "extend stress-ng options",
								MarkdownDescription: "extend stress-ng options",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"size": schema.StringAttribute{
								Description:         "specifies N bytes consumed per vm worker, default is the total available memory. One can specify the size as % of total available memory or in units of B, KB/KiB, MB/MiB, GB/GiB, TB/TiB..",
								MarkdownDescription: "specifies N bytes consumed per vm worker, default is the total available memory. One can specify the size as % of total available memory or in units of B, KB/KiB, MB/MiB, GB/GiB, TB/TiB..",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"uid": schema.StringAttribute{
						Description:         "the experiment ID",
						MarkdownDescription: "the experiment ID",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"user_defined": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"attack_cmd": schema.StringAttribute{
								Description:         "The command to be executed when attack",
								MarkdownDescription: "The command to be executed when attack",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"recover_cmd": schema.StringAttribute{
								Description:         "The command to be executed when recover",
								MarkdownDescription: "The command to be executed when recover",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"value": schema.StringAttribute{
						Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of physical machines to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of physical machines the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
						MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of physical machines to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of physical machines the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vm": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"vm_name": schema.StringAttribute{
								Description:         "The name of the VM to be injected",
								MarkdownDescription: "The name of the VM to be injected",
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ChaosMeshOrgPhysicalMachineChaosV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_chaos_mesh_org_physical_machine_chaos_v1alpha1_manifest")

	var model ChaosMeshOrgPhysicalMachineChaosV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("chaos-mesh.org/v1alpha1")
	model.Kind = pointer.String("PhysicalMachineChaos")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
