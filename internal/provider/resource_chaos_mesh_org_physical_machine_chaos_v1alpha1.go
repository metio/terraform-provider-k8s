/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type ChaosMeshOrgPhysicalMachineChaosV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ChaosMeshOrgPhysicalMachineChaosV1Alpha1Resource)(nil)
)

type ChaosMeshOrgPhysicalMachineChaosV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ChaosMeshOrgPhysicalMachineChaosV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Action *string `tfsdk:"action" yaml:"action,omitempty"`

		Address *[]string `tfsdk:"address" yaml:"address,omitempty"`

		Clock *struct {
			Clock_ids_slice *string `tfsdk:"clock_ids_slice" yaml:"clock-ids-slice,omitempty"`

			Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

			Time_offset *string `tfsdk:"time_offset" yaml:"time-offset,omitempty"`
		} `tfsdk:"clock" yaml:"clock,omitempty"`

		Disk_fill *struct {
			Fill_by_fallocate *bool `tfsdk:"fill_by_fallocate" yaml:"fill-by-fallocate,omitempty"`

			Path *string `tfsdk:"path" yaml:"path,omitempty"`

			Size *string `tfsdk:"size" yaml:"size,omitempty"`
		} `tfsdk:"disk_fill" yaml:"disk-fill,omitempty"`

		Disk_read_payload *struct {
			Path *string `tfsdk:"path" yaml:"path,omitempty"`

			Payload_process_num *int64 `tfsdk:"payload_process_num" yaml:"payload-process-num,omitempty"`

			Size *string `tfsdk:"size" yaml:"size,omitempty"`
		} `tfsdk:"disk_read_payload" yaml:"disk-read-payload,omitempty"`

		Disk_write_payload *struct {
			Path *string `tfsdk:"path" yaml:"path,omitempty"`

			Payload_process_num *int64 `tfsdk:"payload_process_num" yaml:"payload-process-num,omitempty"`

			Size *string `tfsdk:"size" yaml:"size,omitempty"`
		} `tfsdk:"disk_write_payload" yaml:"disk-write-payload,omitempty"`

		Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

		File_append *struct {
			Count *int64 `tfsdk:"count" yaml:"count,omitempty"`

			Data *string `tfsdk:"data" yaml:"data,omitempty"`

			File_name *string `tfsdk:"file_name" yaml:"file-name,omitempty"`
		} `tfsdk:"file_append" yaml:"file-append,omitempty"`

		File_create *struct {
			Dir_name *string `tfsdk:"dir_name" yaml:"dir-name,omitempty"`

			File_name *string `tfsdk:"file_name" yaml:"file-name,omitempty"`
		} `tfsdk:"file_create" yaml:"file-create,omitempty"`

		File_delete *struct {
			Dir_name *string `tfsdk:"dir_name" yaml:"dir-name,omitempty"`

			File_name *string `tfsdk:"file_name" yaml:"file-name,omitempty"`
		} `tfsdk:"file_delete" yaml:"file-delete,omitempty"`

		File_modify *struct {
			File_name *string `tfsdk:"file_name" yaml:"file-name,omitempty"`

			Privilege *int64 `tfsdk:"privilege" yaml:"privilege,omitempty"`
		} `tfsdk:"file_modify" yaml:"file-modify,omitempty"`

		File_rename *struct {
			Dest_file *string `tfsdk:"dest_file" yaml:"dest-file,omitempty"`

			Source_file *string `tfsdk:"source_file" yaml:"source-file,omitempty"`
		} `tfsdk:"file_rename" yaml:"file-rename,omitempty"`

		File_replace *struct {
			Dest_string *string `tfsdk:"dest_string" yaml:"dest-string,omitempty"`

			File_name *string `tfsdk:"file_name" yaml:"file-name,omitempty"`

			Line *int64 `tfsdk:"line" yaml:"line,omitempty"`

			Origin_string *string `tfsdk:"origin_string" yaml:"origin-string,omitempty"`
		} `tfsdk:"file_replace" yaml:"file-replace,omitempty"`

		Http_abort *struct {
			Code *string `tfsdk:"code" yaml:"code,omitempty"`

			Method *string `tfsdk:"method" yaml:"method,omitempty"`

			Path *string `tfsdk:"path" yaml:"path,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

			Proxy_ports *[]string `tfsdk:"proxy_ports" yaml:"proxy_ports,omitempty"`

			Target *string `tfsdk:"target" yaml:"target,omitempty"`
		} `tfsdk:"http_abort" yaml:"http-abort,omitempty"`

		Http_config *struct {
			File_path *string `tfsdk:"file_path" yaml:"file_path,omitempty"`
		} `tfsdk:"http_config" yaml:"http-config,omitempty"`

		Http_delay *struct {
			Code *string `tfsdk:"code" yaml:"code,omitempty"`

			Delay *string `tfsdk:"delay" yaml:"delay,omitempty"`

			Method *string `tfsdk:"method" yaml:"method,omitempty"`

			Path *string `tfsdk:"path" yaml:"path,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

			Proxy_ports *[]string `tfsdk:"proxy_ports" yaml:"proxy_ports,omitempty"`

			Target *string `tfsdk:"target" yaml:"target,omitempty"`
		} `tfsdk:"http_delay" yaml:"http-delay,omitempty"`

		Http_request *struct {
			Count *int64 `tfsdk:"count" yaml:"count,omitempty"`

			Enable_conn_pool *bool `tfsdk:"enable_conn_pool" yaml:"enable-conn-pool,omitempty"`

			Url *string `tfsdk:"url" yaml:"url,omitempty"`
		} `tfsdk:"http_request" yaml:"http-request,omitempty"`

		Jvm_exception *struct {
			Class *string `tfsdk:"class" yaml:"class,omitempty"`

			Exception *string `tfsdk:"exception" yaml:"exception,omitempty"`

			Method *string `tfsdk:"method" yaml:"method,omitempty"`

			Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`
		} `tfsdk:"jvm_exception" yaml:"jvm-exception,omitempty"`

		Jvm_gc *struct {
			Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`
		} `tfsdk:"jvm_gc" yaml:"jvm-gc,omitempty"`

		Jvm_latency *struct {
			Class *string `tfsdk:"class" yaml:"class,omitempty"`

			Latency *int64 `tfsdk:"latency" yaml:"latency,omitempty"`

			Method *string `tfsdk:"method" yaml:"method,omitempty"`

			Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`
		} `tfsdk:"jvm_latency" yaml:"jvm-latency,omitempty"`

		Jvm_mysql *struct {
			Database *string `tfsdk:"database" yaml:"database,omitempty"`

			Exception *string `tfsdk:"exception" yaml:"exception,omitempty"`

			Latency *int64 `tfsdk:"latency" yaml:"latency,omitempty"`

			MysqlConnectorVersion *string `tfsdk:"mysql_connector_version" yaml:"mysqlConnectorVersion,omitempty"`

			Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

			SqlType *string `tfsdk:"sql_type" yaml:"sqlType,omitempty"`

			Table *string `tfsdk:"table" yaml:"table,omitempty"`
		} `tfsdk:"jvm_mysql" yaml:"jvm-mysql,omitempty"`

		Jvm_return *struct {
			Class *string `tfsdk:"class" yaml:"class,omitempty"`

			Method *string `tfsdk:"method" yaml:"method,omitempty"`

			Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"jvm_return" yaml:"jvm-return,omitempty"`

		Jvm_rule_data *struct {
			Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

			Rule_data *string `tfsdk:"rule_data" yaml:"rule-data,omitempty"`
		} `tfsdk:"jvm_rule_data" yaml:"jvm-rule-data,omitempty"`

		Jvm_stress *struct {
			Cpu_count *int64 `tfsdk:"cpu_count" yaml:"cpu-count,omitempty"`

			Mem_type *string `tfsdk:"mem_type" yaml:"mem-type,omitempty"`

			Pid *int64 `tfsdk:"pid" yaml:"pid,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`
		} `tfsdk:"jvm_stress" yaml:"jvm-stress,omitempty"`

		Kafka_fill *struct {
			Host *string `tfsdk:"host" yaml:"host,omitempty"`

			MaxBytes *int64 `tfsdk:"max_bytes" yaml:"maxBytes,omitempty"`

			MessageSize *int64 `tfsdk:"message_size" yaml:"messageSize,omitempty"`

			Password *string `tfsdk:"password" yaml:"password,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

			ReloadCommand *string `tfsdk:"reload_command" yaml:"reloadCommand,omitempty"`

			Topic *string `tfsdk:"topic" yaml:"topic,omitempty"`

			Username *string `tfsdk:"username" yaml:"username,omitempty"`
		} `tfsdk:"kafka_fill" yaml:"kafka-fill,omitempty"`

		Kafka_flood *struct {
			Host *string `tfsdk:"host" yaml:"host,omitempty"`

			MessageSize *int64 `tfsdk:"message_size" yaml:"messageSize,omitempty"`

			Password *string `tfsdk:"password" yaml:"password,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

			Threads *int64 `tfsdk:"threads" yaml:"threads,omitempty"`

			Topic *string `tfsdk:"topic" yaml:"topic,omitempty"`

			Username *string `tfsdk:"username" yaml:"username,omitempty"`
		} `tfsdk:"kafka_flood" yaml:"kafka-flood,omitempty"`

		Kafka_io *struct {
			ConfigFile *string `tfsdk:"config_file" yaml:"configFile,omitempty"`

			NonReadable *bool `tfsdk:"non_readable" yaml:"nonReadable,omitempty"`

			NonWritable *bool `tfsdk:"non_writable" yaml:"nonWritable,omitempty"`

			Topic *string `tfsdk:"topic" yaml:"topic,omitempty"`
		} `tfsdk:"kafka_io" yaml:"kafka-io,omitempty"`

		Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

		Network_bandwidth *struct {
			Buffer *int64 `tfsdk:"buffer" yaml:"buffer,omitempty"`

			Device *string `tfsdk:"device" yaml:"device,omitempty"`

			Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

			Ip_address *string `tfsdk:"ip_address" yaml:"ip-address,omitempty"`

			Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`

			Minburst *int64 `tfsdk:"minburst" yaml:"minburst,omitempty"`

			Peakrate *int64 `tfsdk:"peakrate" yaml:"peakrate,omitempty"`

			Rate *string `tfsdk:"rate" yaml:"rate,omitempty"`
		} `tfsdk:"network_bandwidth" yaml:"network-bandwidth,omitempty"`

		Network_corrupt *struct {
			Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

			Device *string `tfsdk:"device" yaml:"device,omitempty"`

			Egress_port *string `tfsdk:"egress_port" yaml:"egress-port,omitempty"`

			Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

			Ip_address *string `tfsdk:"ip_address" yaml:"ip-address,omitempty"`

			Ip_protocol *string `tfsdk:"ip_protocol" yaml:"ip-protocol,omitempty"`

			Percent *string `tfsdk:"percent" yaml:"percent,omitempty"`

			Source_port *string `tfsdk:"source_port" yaml:"source-port,omitempty"`
		} `tfsdk:"network_corrupt" yaml:"network-corrupt,omitempty"`

		Network_delay *struct {
			Accept_tcp_flags *string `tfsdk:"accept_tcp_flags" yaml:"accept-tcp-flags,omitempty"`

			Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

			Device *string `tfsdk:"device" yaml:"device,omitempty"`

			Egress_port *string `tfsdk:"egress_port" yaml:"egress-port,omitempty"`

			Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

			Ip_address *string `tfsdk:"ip_address" yaml:"ip-address,omitempty"`

			Ip_protocol *string `tfsdk:"ip_protocol" yaml:"ip-protocol,omitempty"`

			Jitter *string `tfsdk:"jitter" yaml:"jitter,omitempty"`

			Latency *string `tfsdk:"latency" yaml:"latency,omitempty"`

			Source_port *string `tfsdk:"source_port" yaml:"source-port,omitempty"`
		} `tfsdk:"network_delay" yaml:"network-delay,omitempty"`

		Network_dns *struct {
			Dns_domain_name *string `tfsdk:"dns_domain_name" yaml:"dns-domain-name,omitempty"`

			Dns_ip *string `tfsdk:"dns_ip" yaml:"dns-ip,omitempty"`

			Dns_server *string `tfsdk:"dns_server" yaml:"dns-server,omitempty"`
		} `tfsdk:"network_dns" yaml:"network-dns,omitempty"`

		Network_down *struct {
			Device *string `tfsdk:"device" yaml:"device,omitempty"`

			Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`
		} `tfsdk:"network_down" yaml:"network-down,omitempty"`

		Network_duplicate *struct {
			Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

			Device *string `tfsdk:"device" yaml:"device,omitempty"`

			Egress_port *string `tfsdk:"egress_port" yaml:"egress-port,omitempty"`

			Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

			Ip_address *string `tfsdk:"ip_address" yaml:"ip-address,omitempty"`

			Ip_protocol *string `tfsdk:"ip_protocol" yaml:"ip-protocol,omitempty"`

			Percent *string `tfsdk:"percent" yaml:"percent,omitempty"`

			Source_port *string `tfsdk:"source_port" yaml:"source-port,omitempty"`
		} `tfsdk:"network_duplicate" yaml:"network-duplicate,omitempty"`

		Network_flood *struct {
			Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

			Ip_address *string `tfsdk:"ip_address" yaml:"ip-address,omitempty"`

			Parallel *int64 `tfsdk:"parallel" yaml:"parallel,omitempty"`

			Port *string `tfsdk:"port" yaml:"port,omitempty"`

			Rate *string `tfsdk:"rate" yaml:"rate,omitempty"`
		} `tfsdk:"network_flood" yaml:"network-flood,omitempty"`

		Network_loss *struct {
			Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

			Device *string `tfsdk:"device" yaml:"device,omitempty"`

			Egress_port *string `tfsdk:"egress_port" yaml:"egress-port,omitempty"`

			Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

			Ip_address *string `tfsdk:"ip_address" yaml:"ip-address,omitempty"`

			Ip_protocol *string `tfsdk:"ip_protocol" yaml:"ip-protocol,omitempty"`

			Percent *string `tfsdk:"percent" yaml:"percent,omitempty"`

			Source_port *string `tfsdk:"source_port" yaml:"source-port,omitempty"`
		} `tfsdk:"network_loss" yaml:"network-loss,omitempty"`

		Network_partition *struct {
			Accept_tcp_flags *string `tfsdk:"accept_tcp_flags" yaml:"accept-tcp-flags,omitempty"`

			Device *string `tfsdk:"device" yaml:"device,omitempty"`

			Direction *string `tfsdk:"direction" yaml:"direction,omitempty"`

			Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

			Ip_address *string `tfsdk:"ip_address" yaml:"ip-address,omitempty"`

			Ip_protocol *string `tfsdk:"ip_protocol" yaml:"ip-protocol,omitempty"`
		} `tfsdk:"network_partition" yaml:"network-partition,omitempty"`

		Process *struct {
			Process *string `tfsdk:"process" yaml:"process,omitempty"`

			RecoverCmd *string `tfsdk:"recover_cmd" yaml:"recoverCmd,omitempty"`

			Signal *int64 `tfsdk:"signal" yaml:"signal,omitempty"`
		} `tfsdk:"process" yaml:"process,omitempty"`

		Redis_cacheLimit *struct {
			Addr *string `tfsdk:"addr" yaml:"addr,omitempty"`

			CacheSize *string `tfsdk:"cache_size" yaml:"cacheSize,omitempty"`

			Password *string `tfsdk:"password" yaml:"password,omitempty"`

			Percent *string `tfsdk:"percent" yaml:"percent,omitempty"`
		} `tfsdk:"redis_cache_limit" yaml:"redis-cacheLimit,omitempty"`

		Redis_expiration *struct {
			Addr *string `tfsdk:"addr" yaml:"addr,omitempty"`

			Expiration *string `tfsdk:"expiration" yaml:"expiration,omitempty"`

			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Option *string `tfsdk:"option" yaml:"option,omitempty"`

			Password *string `tfsdk:"password" yaml:"password,omitempty"`
		} `tfsdk:"redis_expiration" yaml:"redis-expiration,omitempty"`

		Redis_penetration *struct {
			Addr *string `tfsdk:"addr" yaml:"addr,omitempty"`

			Password *string `tfsdk:"password" yaml:"password,omitempty"`

			RequestNum *int64 `tfsdk:"request_num" yaml:"requestNum,omitempty"`
		} `tfsdk:"redis_penetration" yaml:"redis-penetration,omitempty"`

		Redis_restart *struct {
			Addr *string `tfsdk:"addr" yaml:"addr,omitempty"`

			Conf *string `tfsdk:"conf" yaml:"conf,omitempty"`

			FlushConfig *bool `tfsdk:"flush_config" yaml:"flushConfig,omitempty"`

			Password *string `tfsdk:"password" yaml:"password,omitempty"`

			RedisPath *bool `tfsdk:"redis_path" yaml:"redisPath,omitempty"`
		} `tfsdk:"redis_restart" yaml:"redis-restart,omitempty"`

		Redis_stop *struct {
			Addr *string `tfsdk:"addr" yaml:"addr,omitempty"`

			Conf *string `tfsdk:"conf" yaml:"conf,omitempty"`

			FlushConfig *bool `tfsdk:"flush_config" yaml:"flushConfig,omitempty"`

			Password *string `tfsdk:"password" yaml:"password,omitempty"`

			RedisPath *bool `tfsdk:"redis_path" yaml:"redisPath,omitempty"`
		} `tfsdk:"redis_stop" yaml:"redis-stop,omitempty"`

		Selector *struct {
			AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

			ExpressionSelectors *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

			FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

			LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

			Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

			PhysicalMachines *map[string][]string `tfsdk:"physical_machines" yaml:"physicalMachines,omitempty"`
		} `tfsdk:"selector" yaml:"selector,omitempty"`

		Stress_cpu *struct {
			Load *int64 `tfsdk:"load" yaml:"load,omitempty"`

			Options *[]string `tfsdk:"options" yaml:"options,omitempty"`

			Workers *int64 `tfsdk:"workers" yaml:"workers,omitempty"`
		} `tfsdk:"stress_cpu" yaml:"stress-cpu,omitempty"`

		Stress_mem *struct {
			Options *[]string `tfsdk:"options" yaml:"options,omitempty"`

			Size *string `tfsdk:"size" yaml:"size,omitempty"`
		} `tfsdk:"stress_mem" yaml:"stress-mem,omitempty"`

		Uid *string `tfsdk:"uid" yaml:"uid,omitempty"`

		User_defined *struct {
			AttackCmd *string `tfsdk:"attack_cmd" yaml:"attackCmd,omitempty"`

			RecoverCmd *string `tfsdk:"recover_cmd" yaml:"recoverCmd,omitempty"`
		} `tfsdk:"user_defined" yaml:"user_defined,omitempty"`

		Value *string `tfsdk:"value" yaml:"value,omitempty"`

		Vm *struct {
			Vm_name *string `tfsdk:"vm_name" yaml:"vm-name,omitempty"`
		} `tfsdk:"vm" yaml:"vm,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewChaosMeshOrgPhysicalMachineChaosV1Alpha1Resource() resource.Resource {
	return &ChaosMeshOrgPhysicalMachineChaosV1Alpha1Resource{}
}

func (r *ChaosMeshOrgPhysicalMachineChaosV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_chaos_mesh_org_physical_machine_chaos_v1alpha1"
}

func (r *ChaosMeshOrgPhysicalMachineChaosV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "PhysicalMachineChaos is the Schema for the physical machine chaos API",
		MarkdownDescription: "PhysicalMachineChaos is the Schema for the physical machine chaos API",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "Spec defines the behavior of a physical machine chaos experiment",
				MarkdownDescription: "Spec defines the behavior of a physical machine chaos experiment",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"action": {
						Description:         "the subAction, generate automatically",
						MarkdownDescription: "the subAction, generate automatically",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("stress-cpu", "stress-mem", "disk-read-payload", "disk-write-payload", "disk-fill", "network-corrupt", "network-duplicate", "network-loss", "network-delay", "network-partition", "network-dns", "network-bandwidth", "network-flood", "network-down", "process", "jvm-exception", "jvm-gc", "jvm-latency", "jvm-return", "jvm-stress", "jvm-rule-data", "jvm-mysql", "clock", "redis-expiration", "redis-penetration", "redis-cacheLimit", "redis-restart", "redis-stop", "kafka-fill", "kafka-flood", "kafka-io", "file-create", "file-modify", "file-delete", "file-rename", "file-append", "file-replace", "vm", "user_defined"),
						},
					},

					"address": {
						Description:         "DEPRECATED: Use Selector instead. Only one of Address and Selector could be specified.",
						MarkdownDescription: "DEPRECATED: Use Selector instead. Only one of Address and Selector could be specified.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"clock": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"clock_ids_slice": {
								Description:         "the identifier of the particular clock on which to act. More clock description in linux kernel can be found in man page of clock_getres, clock_gettime, clock_settime. Muti clock ids should be split with ','",
								MarkdownDescription: "the identifier of the particular clock on which to act. More clock description in linux kernel can be found in man page of clock_getres, clock_gettime, clock_settime. Muti clock ids should be split with ','",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pid": {
								Description:         "the pid of target program.",
								MarkdownDescription: "the pid of target program.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"time_offset": {
								Description:         "specifies the length of time offset.",
								MarkdownDescription: "specifies the length of time offset.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"disk_fill": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"fill_by_fallocate": {
								Description:         "fill disk by fallocate",
								MarkdownDescription: "fill disk by fallocate",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"path": {
								Description:         "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
								MarkdownDescription: "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"size": {
								Description:         "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
								MarkdownDescription: "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"disk_read_payload": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"path": {
								Description:         "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
								MarkdownDescription: "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"payload_process_num": {
								Description:         "specifies the number of process work on writing, default 1, only 1-255 is valid value",
								MarkdownDescription: "specifies the number of process work on writing, default 1, only 1-255 is valid value",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"size": {
								Description:         "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
								MarkdownDescription: "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"disk_write_payload": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"path": {
								Description:         "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",
								MarkdownDescription: "specifies the location to fill data in. if path not provided, payload will read/write from/into a temp file, temp file will be deleted after writing",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"payload_process_num": {
								Description:         "specifies the number of process work on writing, default 1, only 1-255 is valid value",
								MarkdownDescription: "specifies the number of process work on writing, default 1, only 1-255 is valid value",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"size": {
								Description:         "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",
								MarkdownDescription: "specifies how many units of data will write into the file path. support unit: c=1, w=2, b=512, kB=1000, K=1024, MB=1000*1000, M=1024*1024, GB=1000*1000*1000, G=1024*1024*1024 BYTES. example : 1M | 512kB",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"duration": {
						Description:         "Duration represents the duration of the chaos action",
						MarkdownDescription: "Duration represents the duration of the chaos action",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"file_append": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"count": {
								Description:         "Count is the number of times to append the data.",
								MarkdownDescription: "Count is the number of times to append the data.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"data": {
								Description:         "Data is the data for append.",
								MarkdownDescription: "Data is the data for append.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"file_name": {
								Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
								MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"file_create": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"dir_name": {
								Description:         "DirName is the directory name to create or delete.",
								MarkdownDescription: "DirName is the directory name to create or delete.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"file_name": {
								Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
								MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"file_delete": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"dir_name": {
								Description:         "DirName is the directory name to create or delete.",
								MarkdownDescription: "DirName is the directory name to create or delete.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"file_name": {
								Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
								MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"file_modify": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"file_name": {
								Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
								MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"privilege": {
								Description:         "Privilege is the file privilege to be set.",
								MarkdownDescription: "Privilege is the file privilege to be set.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"file_rename": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"dest_file": {
								Description:         "DestFile is the name to be renamed.",
								MarkdownDescription: "DestFile is the name to be renamed.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"source_file": {
								Description:         "SourceFile is the name need to be renamed.",
								MarkdownDescription: "SourceFile is the name need to be renamed.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"file_replace": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"dest_string": {
								Description:         "DestStr is the destination string of the file.",
								MarkdownDescription: "DestStr is the destination string of the file.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"file_name": {
								Description:         "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",
								MarkdownDescription: "FileName is the name of the file to be created, modified, deleted, renamed, or appended.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"line": {
								Description:         "Line is the line number of the file to be replaced.",
								MarkdownDescription: "Line is the line number of the file to be replaced.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"origin_string": {
								Description:         "OriginStr is the origin string of the file.",
								MarkdownDescription: "OriginStr is the origin string of the file.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"http_abort": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"code": {
								Description:         "Code is a rule to select target by http status code in response",
								MarkdownDescription: "Code is a rule to select target by http status code in response",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"method": {
								Description:         "HTTP method",
								MarkdownDescription: "HTTP method",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"path": {
								Description:         "Match path of Uri with wildcard matches",
								MarkdownDescription: "Match path of Uri with wildcard matches",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "The TCP port that the target service listens on",
								MarkdownDescription: "The TCP port that the target service listens on",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"proxy_ports": {
								Description:         "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",
								MarkdownDescription: "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",

								Type: types.ListType{ElemType: types.StringType},

								Required: true,
								Optional: false,
								Computed: false,
							},

							"target": {
								Description:         "HTTP target: Request or Response",
								MarkdownDescription: "HTTP target: Request or Response",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"http_config": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"file_path": {
								Description:         "The config file path",
								MarkdownDescription: "The config file path",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"http_delay": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"code": {
								Description:         "Code is a rule to select target by http status code in response",
								MarkdownDescription: "Code is a rule to select target by http status code in response",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"delay": {
								Description:         "Delay represents the delay of the target request/response",
								MarkdownDescription: "Delay represents the delay of the target request/response",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"method": {
								Description:         "HTTP method",
								MarkdownDescription: "HTTP method",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"path": {
								Description:         "Match path of Uri with wildcard matches",
								MarkdownDescription: "Match path of Uri with wildcard matches",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "The TCP port that the target service listens on",
								MarkdownDescription: "The TCP port that the target service listens on",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"proxy_ports": {
								Description:         "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",
								MarkdownDescription: "Composed with one of the port of HTTP connection, we will only attack HTTP connection with port inside proxy_ports",

								Type: types.ListType{ElemType: types.StringType},

								Required: true,
								Optional: false,
								Computed: false,
							},

							"target": {
								Description:         "HTTP target: Request or Response",
								MarkdownDescription: "HTTP target: Request or Response",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"http_request": {
						Description:         "used for HTTP request, now only support GET",
						MarkdownDescription: "used for HTTP request, now only support GET",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"count": {
								Description:         "The number of requests to send",
								MarkdownDescription: "The number of requests to send",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_conn_pool": {
								Description:         "Enable connection pool",
								MarkdownDescription: "Enable connection pool",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"url": {
								Description:         "Request to send'",
								MarkdownDescription: "Request to send'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"jvm_exception": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"class": {
								Description:         "Java class",
								MarkdownDescription: "Java class",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"exception": {
								Description:         "the exception which needs to throw for action 'exception'",
								MarkdownDescription: "the exception which needs to throw for action 'exception'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"method": {
								Description:         "the method in Java class",
								MarkdownDescription: "the method in Java class",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pid": {
								Description:         "the pid of Java process which needs to attach",
								MarkdownDescription: "the pid of Java process which needs to attach",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "the port of agent server, default 9277",
								MarkdownDescription: "the port of agent server, default 9277",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"jvm_gc": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"pid": {
								Description:         "the pid of Java process which needs to attach",
								MarkdownDescription: "the pid of Java process which needs to attach",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "the port of agent server, default 9277",
								MarkdownDescription: "the port of agent server, default 9277",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"jvm_latency": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"class": {
								Description:         "Java class",
								MarkdownDescription: "Java class",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"latency": {
								Description:         "the latency duration for action 'latency', unit ms",
								MarkdownDescription: "the latency duration for action 'latency', unit ms",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"method": {
								Description:         "the method in Java class",
								MarkdownDescription: "the method in Java class",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pid": {
								Description:         "the pid of Java process which needs to attach",
								MarkdownDescription: "the pid of Java process which needs to attach",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "the port of agent server, default 9277",
								MarkdownDescription: "the port of agent server, default 9277",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"jvm_mysql": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"database": {
								Description:         "the match database default value is '', means match all database",
								MarkdownDescription: "the match database default value is '', means match all database",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"exception": {
								Description:         "The exception which needs to throw for action 'exception' or the exception message needs to throw in action 'mysql'",
								MarkdownDescription: "The exception which needs to throw for action 'exception' or the exception message needs to throw in action 'mysql'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"latency": {
								Description:         "The latency duration for action 'latency' or the latency duration in action 'mysql'",
								MarkdownDescription: "The latency duration for action 'latency' or the latency duration in action 'mysql'",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mysql_connector_version": {
								Description:         "the version of mysql-connector-java, only support 5.X.X(set to '5') and 8.X.X(set to '8') now",
								MarkdownDescription: "the version of mysql-connector-java, only support 5.X.X(set to '5') and 8.X.X(set to '8') now",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pid": {
								Description:         "the pid of Java process which needs to attach",
								MarkdownDescription: "the pid of Java process which needs to attach",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "the port of agent server, default 9277",
								MarkdownDescription: "the port of agent server, default 9277",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"sql_type": {
								Description:         "the match sql type default value is '', means match all SQL type. The value can be 'select', 'insert', 'update', 'delete', 'replace'.",
								MarkdownDescription: "the match sql type default value is '', means match all SQL type. The value can be 'select', 'insert', 'update', 'delete', 'replace'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"table": {
								Description:         "the match table default value is '', means match all table",
								MarkdownDescription: "the match table default value is '', means match all table",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"jvm_return": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"class": {
								Description:         "Java class",
								MarkdownDescription: "Java class",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"method": {
								Description:         "the method in Java class",
								MarkdownDescription: "the method in Java class",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pid": {
								Description:         "the pid of Java process which needs to attach",
								MarkdownDescription: "the pid of Java process which needs to attach",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "the port of agent server, default 9277",
								MarkdownDescription: "the port of agent server, default 9277",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"value": {
								Description:         "the return value for action 'return'",
								MarkdownDescription: "the return value for action 'return'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"jvm_rule_data": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"pid": {
								Description:         "the pid of Java process which needs to attach",
								MarkdownDescription: "the pid of Java process which needs to attach",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "the port of agent server, default 9277",
								MarkdownDescription: "the port of agent server, default 9277",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"rule_data": {
								Description:         "RuleData used to save the rule file's data, will use it when recover",
								MarkdownDescription: "RuleData used to save the rule file's data, will use it when recover",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"jvm_stress": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cpu_count": {
								Description:         "the CPU core number need to use, only set it when action is stress",
								MarkdownDescription: "the CPU core number need to use, only set it when action is stress",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mem_type": {
								Description:         "the memory type need to locate, only set it when action is stress, the value can be 'stack' or 'heap'",
								MarkdownDescription: "the memory type need to locate, only set it when action is stress, the value can be 'stack' or 'heap'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pid": {
								Description:         "the pid of Java process which needs to attach",
								MarkdownDescription: "the pid of Java process which needs to attach",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "the port of agent server, default 9277",
								MarkdownDescription: "the port of agent server, default 9277",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"kafka_fill": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"host": {
								Description:         "The host of kafka server",
								MarkdownDescription: "The host of kafka server",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_bytes": {
								Description:         "The max bytes to fill",
								MarkdownDescription: "The max bytes to fill",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"message_size": {
								Description:         "The size of each message",
								MarkdownDescription: "The size of each message",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"password": {
								Description:         "The password of kafka client",
								MarkdownDescription: "The password of kafka client",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "The port of kafka server",
								MarkdownDescription: "The port of kafka server",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"reload_command": {
								Description:         "The command to reload kafka config",
								MarkdownDescription: "The command to reload kafka config",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"topic": {
								Description:         "The topic to attack",
								MarkdownDescription: "The topic to attack",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"username": {
								Description:         "The username of kafka client",
								MarkdownDescription: "The username of kafka client",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"kafka_flood": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"host": {
								Description:         "The host of kafka server",
								MarkdownDescription: "The host of kafka server",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"message_size": {
								Description:         "The size of each message",
								MarkdownDescription: "The size of each message",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"password": {
								Description:         "The password of kafka client",
								MarkdownDescription: "The password of kafka client",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "The port of kafka server",
								MarkdownDescription: "The port of kafka server",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"threads": {
								Description:         "The number of worker threads",
								MarkdownDescription: "The number of worker threads",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"topic": {
								Description:         "The topic to attack",
								MarkdownDescription: "The topic to attack",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"username": {
								Description:         "The username of kafka client",
								MarkdownDescription: "The username of kafka client",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"kafka_io": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"config_file": {
								Description:         "The path of server config",
								MarkdownDescription: "The path of server config",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"non_readable": {
								Description:         "Make kafka cluster non-readable",
								MarkdownDescription: "Make kafka cluster non-readable",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"non_writable": {
								Description:         "Make kafka cluster non-writable",
								MarkdownDescription: "Make kafka cluster non-writable",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"topic": {
								Description:         "The topic to attack",
								MarkdownDescription: "The topic to attack",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"mode": {
						Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
						MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
						},
					},

					"network_bandwidth": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"buffer": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(1),
								},
							},

							"device": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"hostname": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ip_address": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"limit": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(1),
								},
							},

							"minburst": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"peakrate": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"rate": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"network_corrupt": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"correlation": {
								Description:         "correlation is percentage (10 is 10%)",
								MarkdownDescription: "correlation is percentage (10 is 10%)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"device": {
								Description:         "the network interface to impact",
								MarkdownDescription: "the network interface to impact",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"egress_port": {
								Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"hostname": {
								Description:         "only impact traffic to these hostnames",
								MarkdownDescription: "only impact traffic to these hostnames",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ip_address": {
								Description:         "only impact egress traffic to these IP addresses",
								MarkdownDescription: "only impact egress traffic to these IP addresses",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ip_protocol": {
								Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
								MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"percent": {
								Description:         "percentage of packets to corrupt (10 is 10%)",
								MarkdownDescription: "percentage of packets to corrupt (10 is 10%)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"source_port": {
								Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"network_delay": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"accept_tcp_flags": {
								Description:         "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",
								MarkdownDescription: "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"correlation": {
								Description:         "correlation is percentage (10 is 10%)",
								MarkdownDescription: "correlation is percentage (10 is 10%)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"device": {
								Description:         "the network interface to impact",
								MarkdownDescription: "the network interface to impact",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"egress_port": {
								Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"hostname": {
								Description:         "only impact traffic to these hostnames",
								MarkdownDescription: "only impact traffic to these hostnames",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ip_address": {
								Description:         "only impact egress traffic to these IP addresses",
								MarkdownDescription: "only impact egress traffic to these IP addresses",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ip_protocol": {
								Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
								MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"jitter": {
								Description:         "jitter time, time units: ns, us (or µs), ms, s, m, h.",
								MarkdownDescription: "jitter time, time units: ns, us (or µs), ms, s, m, h.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"latency": {
								Description:         "delay egress time, time units: ns, us (or µs), ms, s, m, h.",
								MarkdownDescription: "delay egress time, time units: ns, us (or µs), ms, s, m, h.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"source_port": {
								Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"network_dns": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"dns_domain_name": {
								Description:         "map this host to specified IP",
								MarkdownDescription: "map this host to specified IP",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dns_ip": {
								Description:         "map specified host to this IP address",
								MarkdownDescription: "map specified host to this IP address",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dns_server": {
								Description:         "update the DNS server in /etc/resolv.conf with this value",
								MarkdownDescription: "update the DNS server in /etc/resolv.conf with this value",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"network_down": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"device": {
								Description:         "The network interface to impact",
								MarkdownDescription: "The network interface to impact",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"duration": {
								Description:         "NIC down time, time units: ns, us (or µs), ms, s, m, h.",
								MarkdownDescription: "NIC down time, time units: ns, us (or µs), ms, s, m, h.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"network_duplicate": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"correlation": {
								Description:         "correlation is percentage (10 is 10%)",
								MarkdownDescription: "correlation is percentage (10 is 10%)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"device": {
								Description:         "the network interface to impact",
								MarkdownDescription: "the network interface to impact",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"egress_port": {
								Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"hostname": {
								Description:         "only impact traffic to these hostnames",
								MarkdownDescription: "only impact traffic to these hostnames",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ip_address": {
								Description:         "only impact egress traffic to these IP addresses",
								MarkdownDescription: "only impact egress traffic to these IP addresses",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ip_protocol": {
								Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
								MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"percent": {
								Description:         "percentage of packets to duplicate (10 is 10%)",
								MarkdownDescription: "percentage of packets to duplicate (10 is 10%)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"source_port": {
								Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"network_flood": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"duration": {
								Description:         "The number of seconds to run the iperf test",
								MarkdownDescription: "The number of seconds to run the iperf test",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"ip_address": {
								Description:         "Generate traffic to this IP address",
								MarkdownDescription: "Generate traffic to this IP address",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"parallel": {
								Description:         "The number of iperf parallel client threads to run",
								MarkdownDescription: "The number of iperf parallel client threads to run",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "Generate traffic to this port on the IP address",
								MarkdownDescription: "Generate traffic to this port on the IP address",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"rate": {
								Description:         "The speed of network traffic, allows bps, kbps, mbps, gbps, tbps unit. bps means bytes per second",
								MarkdownDescription: "The speed of network traffic, allows bps, kbps, mbps, gbps, tbps unit. bps means bytes per second",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"network_loss": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"correlation": {
								Description:         "correlation is percentage (10 is 10%)",
								MarkdownDescription: "correlation is percentage (10 is 10%)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"device": {
								Description:         "the network interface to impact",
								MarkdownDescription: "the network interface to impact",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"egress_port": {
								Description:         "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								MarkdownDescription: "only impact egress traffic to these destination ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"hostname": {
								Description:         "only impact traffic to these hostnames",
								MarkdownDescription: "only impact traffic to these hostnames",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ip_address": {
								Description:         "only impact egress traffic to these IP addresses",
								MarkdownDescription: "only impact egress traffic to these IP addresses",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ip_protocol": {
								Description:         "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",
								MarkdownDescription: "only impact traffic using this IP protocol, supported: tcp, udp, icmp, all",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"percent": {
								Description:         "percentage of packets to loss (10 is 10%)",
								MarkdownDescription: "percentage of packets to loss (10 is 10%)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"source_port": {
								Description:         "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",
								MarkdownDescription: "only impact egress traffic from these source ports, use a ',' to separate or to indicate the range, such as 80, 8001:8010. it can only be used in conjunction with -p tcp or -p udp",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"network_partition": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"accept_tcp_flags": {
								Description:         "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",
								MarkdownDescription: "only the packet which match the tcp flag can be accepted, others will be dropped. only set when the IPProtocol is tcp, used for partition.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"device": {
								Description:         "the network interface to impact",
								MarkdownDescription: "the network interface to impact",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"direction": {
								Description:         "specifies the partition direction, values can be 'from', 'to'. 'from' means packets coming from the 'IPAddress' or 'Hostname' and going to your server, 'to' means packets originating from your server and going to the 'IPAddress' or 'Hostname'.",
								MarkdownDescription: "specifies the partition direction, values can be 'from', 'to'. 'from' means packets coming from the 'IPAddress' or 'Hostname' and going to your server, 'to' means packets originating from your server and going to the 'IPAddress' or 'Hostname'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"hostname": {
								Description:         "only impact traffic to these hostnames",
								MarkdownDescription: "only impact traffic to these hostnames",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ip_address": {
								Description:         "only impact egress traffic to these IP addresses",
								MarkdownDescription: "only impact egress traffic to these IP addresses",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ip_protocol": {
								Description:         "only impact egress traffic to these IP addresses",
								MarkdownDescription: "only impact egress traffic to these IP addresses",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"process": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"process": {
								Description:         "the process name or the process ID",
								MarkdownDescription: "the process name or the process ID",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"recover_cmd": {
								Description:         "the command to be run when recovering experiment",
								MarkdownDescription: "the command to be run when recovering experiment",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"signal": {
								Description:         "the signal number to send",
								MarkdownDescription: "the signal number to send",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"redis_cache_limit": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"addr": {
								Description:         "The adress of Redis server",
								MarkdownDescription: "The adress of Redis server",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cache_size": {
								Description:         "The size of 'maxmemory'",
								MarkdownDescription: "The size of 'maxmemory'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"password": {
								Description:         "The password of Redis server",
								MarkdownDescription: "The password of Redis server",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"percent": {
								Description:         "Specifies maxmemory as a percentage of the original value",
								MarkdownDescription: "Specifies maxmemory as a percentage of the original value",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"redis_expiration": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"addr": {
								Description:         "The adress of Redis server",
								MarkdownDescription: "The adress of Redis server",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"expiration": {
								Description:         "The expiration of the keys",
								MarkdownDescription: "The expiration of the keys",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"key": {
								Description:         "The keys to be expired",
								MarkdownDescription: "The keys to be expired",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"option": {
								Description:         "Additional options for 'expiration'",
								MarkdownDescription: "Additional options for 'expiration'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"password": {
								Description:         "The password of Redis server",
								MarkdownDescription: "The password of Redis server",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"redis_penetration": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"addr": {
								Description:         "The adress of Redis server",
								MarkdownDescription: "The adress of Redis server",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"password": {
								Description:         "The password of Redis server",
								MarkdownDescription: "The password of Redis server",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"request_num": {
								Description:         "The number of requests to be sent",
								MarkdownDescription: "The number of requests to be sent",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"redis_restart": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"addr": {
								Description:         "The adress of Redis server",
								MarkdownDescription: "The adress of Redis server",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"conf": {
								Description:         "The path of Sentinel conf",
								MarkdownDescription: "The path of Sentinel conf",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"flush_config": {
								Description:         "The control flag determines whether to flush config",
								MarkdownDescription: "The control flag determines whether to flush config",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"password": {
								Description:         "The password of Redis server",
								MarkdownDescription: "The password of Redis server",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"redis_path": {
								Description:         "The path of 'redis-server' command-line tool",
								MarkdownDescription: "The path of 'redis-server' command-line tool",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"redis_stop": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"addr": {
								Description:         "The adress of Redis server",
								MarkdownDescription: "The adress of Redis server",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"conf": {
								Description:         "The path of Sentinel conf",
								MarkdownDescription: "The path of Sentinel conf",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"flush_config": {
								Description:         "The control flag determines whether to flush config",
								MarkdownDescription: "The control flag determines whether to flush config",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"password": {
								Description:         "The password of Redis server",
								MarkdownDescription: "The password of Redis server",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"redis_path": {
								Description:         "The path of 'redis-server' command-line tool",
								MarkdownDescription: "The path of 'redis-server' command-line tool",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"selector": {
						Description:         "Selector is used to select physical machines that are used to inject chaos action.",
						MarkdownDescription: "Selector is used to select physical machines that are used to inject chaos action.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"annotation_selectors": {
								Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"expression_selectors": {
								Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
								MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "key is the label key that the selector applies to.",
										MarkdownDescription: "key is the label key that the selector applies to.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"operator": {
										Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
										MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"values": {
										Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
										MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"field_selectors": {
								Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"label_selectors": {
								Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespaces": {
								Description:         "Namespaces is a set of namespace to which objects belong.",
								MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"physical_machines": {
								Description:         "PhysicalMachines is a map of string keys and a set values that used to select physical machines. The key defines the namespace which physical machine belong, and each value is a set of physical machine names.",
								MarkdownDescription: "PhysicalMachines is a map of string keys and a set values that used to select physical machines. The key defines the namespace which physical machine belong, and each value is a set of physical machine names.",

								Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"stress_cpu": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"load": {
								Description:         "specifies P percent loading per CPU worker. 0 is effectively a sleep (no load) and 100 is full loading.",
								MarkdownDescription: "specifies P percent loading per CPU worker. 0 is effectively a sleep (no load) and 100 is full loading.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"options": {
								Description:         "extend stress-ng options",
								MarkdownDescription: "extend stress-ng options",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"workers": {
								Description:         "specifies N workers to apply the stressor.",
								MarkdownDescription: "specifies N workers to apply the stressor.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"stress_mem": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"options": {
								Description:         "extend stress-ng options",
								MarkdownDescription: "extend stress-ng options",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"size": {
								Description:         "specifies N bytes consumed per vm worker, default is the total available memory. One can specify the size as % of total available memory or in units of B, KB/KiB, MB/MiB, GB/GiB, TB/TiB..",
								MarkdownDescription: "specifies N bytes consumed per vm worker, default is the total available memory. One can specify the size as % of total available memory or in units of B, KB/KiB, MB/MiB, GB/GiB, TB/TiB..",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"uid": {
						Description:         "the experiment ID",
						MarkdownDescription: "the experiment ID",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"user_defined": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"attack_cmd": {
								Description:         "The command to be executed when attack",
								MarkdownDescription: "The command to be executed when attack",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"recover_cmd": {
								Description:         "The command to be executed when recover",
								MarkdownDescription: "The command to be executed when recover",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"value": {
						Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of physical machines to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of physical machines the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
						MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of physical machines to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of physical machines the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"vm": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"vm_name": {
								Description:         "The name of the VM to be injected",
								MarkdownDescription: "The name of the VM to be injected",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *ChaosMeshOrgPhysicalMachineChaosV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_chaos_mesh_org_physical_machine_chaos_v1alpha1")

	var state ChaosMeshOrgPhysicalMachineChaosV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ChaosMeshOrgPhysicalMachineChaosV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("chaos-mesh.org/v1alpha1")
	goModel.Kind = utilities.Ptr("PhysicalMachineChaos")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *ChaosMeshOrgPhysicalMachineChaosV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_chaos_mesh_org_physical_machine_chaos_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ChaosMeshOrgPhysicalMachineChaosV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_chaos_mesh_org_physical_machine_chaos_v1alpha1")

	var state ChaosMeshOrgPhysicalMachineChaosV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ChaosMeshOrgPhysicalMachineChaosV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("chaos-mesh.org/v1alpha1")
	goModel.Kind = utilities.Ptr("PhysicalMachineChaos")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *ChaosMeshOrgPhysicalMachineChaosV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_chaos_mesh_org_physical_machine_chaos_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
