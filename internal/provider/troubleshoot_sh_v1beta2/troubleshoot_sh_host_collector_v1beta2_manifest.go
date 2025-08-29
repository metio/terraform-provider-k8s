/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package troubleshoot_sh_v1beta2

import (
	"context"
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
	_ datasource.DataSource = &TroubleshootShHostCollectorV1Beta2Manifest{}
)

func NewTroubleshootShHostCollectorV1Beta2Manifest() datasource.DataSource {
	return &TroubleshootShHostCollectorV1Beta2Manifest{}
}

type TroubleshootShHostCollectorV1Beta2Manifest struct{}

type TroubleshootShHostCollectorV1Beta2ManifestData struct {
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
		Analyzers *[]struct {
			BlockDevices *struct {
				Annotations                *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName                  *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName              *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude                    UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				IncludeUnmountedPartitions *bool              `tfsdk:"include_unmounted_partitions" json:"includeUnmountedPartitions,omitempty"`
				MinimumAcceptableSize      *int64             `tfsdk:"minimum_acceptable_size" json:"minimumAcceptableSize,omitempty"`
				Outcomes                   *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"block_devices" json:"blockDevices,omitempty"`
			Certificate *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"certificate" json:"certificate,omitempty"`
			CertificatesCollection *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"certificates_collection" json:"certificatesCollection,omitempty"`
			Cpu *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"cpu" json:"cpu,omitempty"`
			DiskUsage *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"disk_usage" json:"diskUsage,omitempty"`
			FilesystemPerformance *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"filesystem_performance" json:"filesystemPerformance,omitempty"`
			HostOS *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"host_os" json:"hostOS,omitempty"`
			HostServices *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"host_services" json:"hostServices,omitempty"`
			Http *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"http" json:"http,omitempty"`
			HttpLoadBalancer *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"http_load_balancer" json:"httpLoadBalancer,omitempty"`
			Ipv4Interfaces *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"ipv4_interfaces" json:"ipv4Interfaces,omitempty"`
			JsonCompare *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				FileName      *string            `tfsdk:"file_name" json:"fileName,omitempty"`
				JsonPath      *string            `tfsdk:"json_path" json:"jsonPath,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Path   *string `tfsdk:"path" json:"path,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
				Value  *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"json_compare" json:"jsonCompare,omitempty"`
			KernelConfigs *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				SelectedConfigs *[]string `tfsdk:"selected_configs" json:"selectedConfigs,omitempty"`
				Strict          UNKNOWN   `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"kernel_configs" json:"kernelConfigs,omitempty"`
			KernelModules *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"kernel_modules" json:"kernelModules,omitempty"`
			Memory *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"memory" json:"memory,omitempty"`
			NetworkNamespaceConnectivity *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"network_namespace_connectivity" json:"networkNamespaceConnectivity,omitempty"`
			SubnetAvailable *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"subnet_available" json:"subnetAvailable,omitempty"`
			SubnetContainsIP *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				Cidr          *string            `tfsdk:"cidr" json:"cidr,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Ip            *string            `tfsdk:"ip" json:"ip,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"subnet_contains_ip" json:"subnetContainsIP,omitempty"`
			Sysctl *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"sysctl" json:"sysctl,omitempty"`
			SystemPackages *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"system_packages" json:"systemPackages,omitempty"`
			TcpConnect *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"tcp_connect" json:"tcpConnect,omitempty"`
			TcpLoadBalancer *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"tcp_load_balancer" json:"tcpLoadBalancer,omitempty"`
			TcpPortStatus *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"tcp_port_status" json:"tcpPortStatus,omitempty"`
			TextAnalyze *struct {
				Annotations     *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName       *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName   *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude         UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				ExcludeFiles    *[]string          `tfsdk:"exclude_files" json:"excludeFiles,omitempty"`
				FileName        *string            `tfsdk:"file_name" json:"fileName,omitempty"`
				IgnoreIfNoFiles *bool              `tfsdk:"ignore_if_no_files" json:"ignoreIfNoFiles,omitempty"`
				Outcomes        *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Regex       *string `tfsdk:"regex" json:"regex,omitempty"`
				RegexGroups *string `tfsdk:"regex_groups" json:"regexGroups,omitempty"`
				Strict      UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"text_analyze" json:"textAnalyze,omitempty"`
			Time *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"time" json:"time,omitempty"`
			UdpPortStatus *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				CheckName     *string            `tfsdk:"check_name" json:"checkName,omitempty"`
				CollectorName *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				Outcomes      *[]struct {
					Fail *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"fail" json:"fail,omitempty"`
					Pass *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"pass" json:"pass,omitempty"`
					Warn *struct {
						Message *string `tfsdk:"message" json:"message,omitempty"`
						Uri     *string `tfsdk:"uri" json:"uri,omitempty"`
						When    *string `tfsdk:"when" json:"when,omitempty"`
					} `tfsdk:"warn" json:"warn,omitempty"`
				} `tfsdk:"outcomes" json:"outcomes,omitempty"`
				Strict UNKNOWN `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"udp_port_status" json:"udpPortStatus,omitempty"`
		} `tfsdk:"analyzers" json:"analyzers,omitempty"`
		Collectors *[]struct {
			BlockDevices *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"block_devices" json:"blockDevices,omitempty"`
			Certificate *struct {
				CertificatePath *string `tfsdk:"certificate_path" json:"certificatePath,omitempty"`
				CollectorName   *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude         UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
				KeyPath         *string `tfsdk:"key_path" json:"keyPath,omitempty"`
			} `tfsdk:"certificate" json:"certificate,omitempty"`
			CertificatesCollection *struct {
				CollectorName *string   `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN   `tfsdk:"exclude" json:"exclude,omitempty"`
				Paths         *[]string `tfsdk:"paths" json:"paths,omitempty"`
			} `tfsdk:"certificates_collection" json:"certificatesCollection,omitempty"`
			Cgroups *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
				MountPoint    *string `tfsdk:"mount_point" json:"mountPoint,omitempty"`
			} `tfsdk:"cgroups" json:"cgroups,omitempty"`
			Copy *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
				Path          *string `tfsdk:"path" json:"path,omitempty"`
			} `tfsdk:"copy" json:"copy,omitempty"`
			Cpu *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"cpu" json:"cpu,omitempty"`
			DiskUsage *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
				Path          *string `tfsdk:"path" json:"path,omitempty"`
			} `tfsdk:"disk_usage" json:"diskUsage,omitempty"`
			Dns *struct {
				CollectorName *string   `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN   `tfsdk:"exclude" json:"exclude,omitempty"`
				Hostnames     *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
			} `tfsdk:"dns" json:"dns,omitempty"`
			FilesystemPerformance *struct {
				BackgroundIOPSWarmupSeconds *int64  `tfsdk:"background_iops_warmup_seconds" json:"backgroundIOPSWarmupSeconds,omitempty"`
				BackgroundReadIOPS          *int64  `tfsdk:"background_read_iops" json:"backgroundReadIOPS,omitempty"`
				BackgroundReadIOPSJobs      *int64  `tfsdk:"background_read_iops_jobs" json:"backgroundReadIOPSJobs,omitempty"`
				BackgroundWriteIOPS         *int64  `tfsdk:"background_write_iops" json:"backgroundWriteIOPS,omitempty"`
				BackgroundWriteIOPSJobs     *int64  `tfsdk:"background_write_iops_jobs" json:"backgroundWriteIOPSJobs,omitempty"`
				CollectorName               *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Datasync                    *bool   `tfsdk:"datasync" json:"datasync,omitempty"`
				Directory                   *string `tfsdk:"directory" json:"directory,omitempty"`
				EnableBackgroundIOPS        *bool   `tfsdk:"enable_background_iops" json:"enableBackgroundIOPS,omitempty"`
				Exclude                     UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
				FileSize                    *string `tfsdk:"file_size" json:"fileSize,omitempty"`
				OperationSize               *int64  `tfsdk:"operation_size" json:"operationSize,omitempty"`
				RunTime                     *string `tfsdk:"run_time" json:"runTime,omitempty"`
				Sync                        *bool   `tfsdk:"sync" json:"sync,omitempty"`
				Timeout                     *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"filesystem_performance" json:"filesystemPerformance,omitempty"`
			HostOS *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"host_os" json:"hostOS,omitempty"`
			HostServices *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"host_services" json:"hostServices,omitempty"`
			Http *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
				Get           *struct {
					Headers            *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
					InsecureSkipVerify *bool              `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
					Proxy              *string            `tfsdk:"proxy" json:"proxy,omitempty"`
					Timeout            *string            `tfsdk:"timeout" json:"timeout,omitempty"`
					Tls                *struct {
						Cacert     *string `tfsdk:"cacert" json:"cacert,omitempty"`
						ClientCert *string `tfsdk:"client_cert" json:"clientCert,omitempty"`
						ClientKey  *string `tfsdk:"client_key" json:"clientKey,omitempty"`
						Secret     *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"secret" json:"secret,omitempty"`
						SkipVerify *bool `tfsdk:"skip_verify" json:"skipVerify,omitempty"`
					} `tfsdk:"tls" json:"tls,omitempty"`
					Url *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"get" json:"get,omitempty"`
				Post *struct {
					Body               *string            `tfsdk:"body" json:"body,omitempty"`
					Headers            *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
					InsecureSkipVerify *bool              `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
					Proxy              *string            `tfsdk:"proxy" json:"proxy,omitempty"`
					Timeout            *string            `tfsdk:"timeout" json:"timeout,omitempty"`
					Tls                *struct {
						Cacert     *string `tfsdk:"cacert" json:"cacert,omitempty"`
						ClientCert *string `tfsdk:"client_cert" json:"clientCert,omitempty"`
						ClientKey  *string `tfsdk:"client_key" json:"clientKey,omitempty"`
						Secret     *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"secret" json:"secret,omitempty"`
						SkipVerify *bool `tfsdk:"skip_verify" json:"skipVerify,omitempty"`
					} `tfsdk:"tls" json:"tls,omitempty"`
					Url *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"post" json:"post,omitempty"`
				Put *struct {
					Body               *string            `tfsdk:"body" json:"body,omitempty"`
					Headers            *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
					InsecureSkipVerify *bool              `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
					Proxy              *string            `tfsdk:"proxy" json:"proxy,omitempty"`
					Timeout            *string            `tfsdk:"timeout" json:"timeout,omitempty"`
					Tls                *struct {
						Cacert     *string `tfsdk:"cacert" json:"cacert,omitempty"`
						ClientCert *string `tfsdk:"client_cert" json:"clientCert,omitempty"`
						ClientKey  *string `tfsdk:"client_key" json:"clientKey,omitempty"`
						Secret     *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"secret" json:"secret,omitempty"`
						SkipVerify *bool `tfsdk:"skip_verify" json:"skipVerify,omitempty"`
					} `tfsdk:"tls" json:"tls,omitempty"`
					Url *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"put" json:"put,omitempty"`
			} `tfsdk:"http" json:"http,omitempty"`
			HttpLoadBalancer *struct {
				Address       *string `tfsdk:"address" json:"address,omitempty"`
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
				Path          *string `tfsdk:"path" json:"path,omitempty"`
				Port          *int64  `tfsdk:"port" json:"port,omitempty"`
				Timeout       *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"http_load_balancer" json:"httpLoadBalancer,omitempty"`
			Ipv4Interfaces *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"ipv4_interfaces" json:"ipv4Interfaces,omitempty"`
			Journald *struct {
				CollectorName *string   `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Dmesg         *bool     `tfsdk:"dmesg" json:"dmesg,omitempty"`
				Exclude       UNKNOWN   `tfsdk:"exclude" json:"exclude,omitempty"`
				Lines         *int64    `tfsdk:"lines" json:"lines,omitempty"`
				Output        *string   `tfsdk:"output" json:"output,omitempty"`
				Reverse       *bool     `tfsdk:"reverse" json:"reverse,omitempty"`
				Since         *string   `tfsdk:"since" json:"since,omitempty"`
				System        *bool     `tfsdk:"system" json:"system,omitempty"`
				Timeout       *string   `tfsdk:"timeout" json:"timeout,omitempty"`
				Units         *[]string `tfsdk:"units" json:"units,omitempty"`
				Until         *string   `tfsdk:"until" json:"until,omitempty"`
				Utc           *bool     `tfsdk:"utc" json:"utc,omitempty"`
			} `tfsdk:"journald" json:"journald,omitempty"`
			KernelConfigs *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"kernel_configs" json:"kernelConfigs,omitempty"`
			KernelModules *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"kernel_modules" json:"kernelModules,omitempty"`
			Kubernetes *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
			Memory *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"memory" json:"memory,omitempty"`
			NetworkNamespaceConnectivity *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
				FromCIDR      *string `tfsdk:"from_cidr" json:"fromCIDR,omitempty"`
				Port          *int64  `tfsdk:"port" json:"port,omitempty"`
				Timeout       *string `tfsdk:"timeout" json:"timeout,omitempty"`
				ToCIDR        *string `tfsdk:"to_cidr" json:"toCIDR,omitempty"`
			} `tfsdk:"network_namespace_connectivity" json:"networkNamespaceConnectivity,omitempty"`
			Run *struct {
				Args             *[]string          `tfsdk:"args" json:"args,omitempty"`
				CollectorName    *string            `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Command          *string            `tfsdk:"command" json:"command,omitempty"`
				Env              *[]string          `tfsdk:"env" json:"env,omitempty"`
				Exclude          UNKNOWN            `tfsdk:"exclude" json:"exclude,omitempty"`
				IgnoreParentEnvs *bool              `tfsdk:"ignore_parent_envs" json:"ignoreParentEnvs,omitempty"`
				InheritEnvs      *[]string          `tfsdk:"inherit_envs" json:"inheritEnvs,omitempty"`
				Input            *map[string]string `tfsdk:"input" json:"input,omitempty"`
				OutputDir        *string            `tfsdk:"output_dir" json:"outputDir,omitempty"`
				Timeout          *string            `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"run" json:"run,omitempty"`
			SubnetAvailable *struct {
				CIDRRangeAlloc *string `tfsdk:"cidr_range_alloc" json:"CIDRRangeAlloc,omitempty"`
				CollectorName  *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				DesiredCIDR    *int64  `tfsdk:"desired_cidr" json:"desiredCIDR,omitempty"`
				Exclude        UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"subnet_available" json:"subnetAvailable,omitempty"`
			Sysctl *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"sysctl" json:"sysctl,omitempty"`
			SystemPackages *struct {
				Amzn          *[]string `tfsdk:"amzn" json:"amzn,omitempty"`
				Amzn2         *[]string `tfsdk:"amzn2" json:"amzn2,omitempty"`
				Centos        *[]string `tfsdk:"centos" json:"centos,omitempty"`
				Centos7       *[]string `tfsdk:"centos7" json:"centos7,omitempty"`
				Centos8       *[]string `tfsdk:"centos8" json:"centos8,omitempty"`
				Centos9       *[]string `tfsdk:"centos9" json:"centos9,omitempty"`
				CollectorName *string   `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN   `tfsdk:"exclude" json:"exclude,omitempty"`
				Ol            *[]string `tfsdk:"ol" json:"ol,omitempty"`
				Ol7           *[]string `tfsdk:"ol7" json:"ol7,omitempty"`
				Ol8           *[]string `tfsdk:"ol8" json:"ol8,omitempty"`
				Ol9           *[]string `tfsdk:"ol9" json:"ol9,omitempty"`
				Rhel          *[]string `tfsdk:"rhel" json:"rhel,omitempty"`
				Rhel7         *[]string `tfsdk:"rhel7" json:"rhel7,omitempty"`
				Rhel8         *[]string `tfsdk:"rhel8" json:"rhel8,omitempty"`
				Rhel9         *[]string `tfsdk:"rhel9" json:"rhel9,omitempty"`
				Rocky         *[]string `tfsdk:"rocky" json:"rocky,omitempty"`
				Rocky8        *[]string `tfsdk:"rocky8" json:"rocky8,omitempty"`
				Rocky9        *[]string `tfsdk:"rocky9" json:"rocky9,omitempty"`
				Ubuntu        *[]string `tfsdk:"ubuntu" json:"ubuntu,omitempty"`
				Ubuntu16      *[]string `tfsdk:"ubuntu16" json:"ubuntu16,omitempty"`
				Ubuntu18      *[]string `tfsdk:"ubuntu18" json:"ubuntu18,omitempty"`
				Ubuntu20      *[]string `tfsdk:"ubuntu20" json:"ubuntu20,omitempty"`
			} `tfsdk:"system_packages" json:"systemPackages,omitempty"`
			TcpConnect *struct {
				Address       *string `tfsdk:"address" json:"address,omitempty"`
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
				Timeout       *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"tcp_connect" json:"tcpConnect,omitempty"`
			TcpLoadBalancer *struct {
				Address       *string `tfsdk:"address" json:"address,omitempty"`
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
				Port          *int64  `tfsdk:"port" json:"port,omitempty"`
				Timeout       *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"tcp_load_balancer" json:"tcpLoadBalancer,omitempty"`
			TcpPortStatus *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
				Interface     *string `tfsdk:"interface" json:"interface,omitempty"`
				Port          *int64  `tfsdk:"port" json:"port,omitempty"`
			} `tfsdk:"tcp_port_status" json:"tcpPortStatus,omitempty"`
			Time *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
			} `tfsdk:"time" json:"time,omitempty"`
			UdpPortStatus *struct {
				CollectorName *string `tfsdk:"collector_name" json:"collectorName,omitempty"`
				Exclude       UNKNOWN `tfsdk:"exclude" json:"exclude,omitempty"`
				Interface     *string `tfsdk:"interface" json:"interface,omitempty"`
				Port          *int64  `tfsdk:"port" json:"port,omitempty"`
			} `tfsdk:"udp_port_status" json:"udpPortStatus,omitempty"`
		} `tfsdk:"collectors" json:"collectors,omitempty"`
		Uri *string `tfsdk:"uri" json:"uri,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TroubleshootShHostCollectorV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_troubleshoot_sh_host_collector_v1beta2_manifest"
}

func (r *TroubleshootShHostCollectorV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "HostCollector is the Schema for the collectors API",
		MarkdownDescription: "HostCollector is the Schema for the collectors API",
		Attributes: map[string]schema.Attribute{
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
				Description:         "HostCollectorSpec defines the desired state of HostCollector",
				MarkdownDescription: "HostCollectorSpec defines the desired state of HostCollector",
				Attributes: map[string]schema.Attribute{
					"analyzers": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"block_devices": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"include_unmounted_partitions": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"minimum_acceptable_size": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"certificate": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"certificates_collection": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"cpu": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"disk_usage": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"filesystem_performance": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"host_os": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"host_services": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"http": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"http_load_balancer": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"ipv4_interfaces": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"json_compare": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"file_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"json_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"kernel_configs": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"selected_configs": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"kernel_modules": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"memory": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"network_namespace_connectivity": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"subnet_available": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"subnet_contains_ip": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"cidr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ip": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"sysctl": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"system_packages": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"tcp_connect": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"tcp_load_balancer": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"tcp_port_status": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"text_analyze": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude_files": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"file_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ignore_if_no_files": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"regex": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"regex_groups": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"time": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"udp_port_status": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"check_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outcomes": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"fail": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pass": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"warn": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"message": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uri": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"when": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"strict": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"collectors": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"block_devices": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"certificate": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"certificate_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key_path": schema.StringAttribute{
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

								"certificates_collection": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"paths": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"cgroups": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mount_point": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"copy": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
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

								"cpu": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"disk_usage": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
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

								"dns": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"hostnames": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"filesystem_performance": schema.SingleNestedAttribute{
									Description:         "FilesystemPerformance benchmarks sequential write latency on a single file. The optional background IOPS feature attempts to mimic real-world conditions by running read and write workloads prior to and during benchmark execution.",
									MarkdownDescription: "FilesystemPerformance benchmarks sequential write latency on a single file. The optional background IOPS feature attempts to mimic real-world conditions by running read and write workloads prior to and during benchmark execution.",
									Attributes: map[string]schema.Attribute{
										"background_iops_warmup_seconds": schema.Int64Attribute{
											Description:         "How long to run the background IOPS read and write workloads prior to starting the benchmarks.",
											MarkdownDescription: "How long to run the background IOPS read and write workloads prior to starting the benchmarks.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"background_read_iops": schema.Int64Attribute{
											Description:         "The target read IOPS to run while benchmarking. This is a limit and there is no guarantee it will be reached. This is the total IOPS for all background read jobs.",
											MarkdownDescription: "The target read IOPS to run while benchmarking. This is a limit and there is no guarantee it will be reached. This is the total IOPS for all background read jobs.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"background_read_iops_jobs": schema.Int64Attribute{
											Description:         "Number of threads to use for background read IOPS. This should be set high enough to reach the target specified in BackgrounReadIOPS.",
											MarkdownDescription: "Number of threads to use for background read IOPS. This should be set high enough to reach the target specified in BackgrounReadIOPS.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"background_write_iops": schema.Int64Attribute{
											Description:         "The target write IOPS to run while benchmarking. This is a limit and there is no guarantee it will be reached. This is the total IOPS for all background write jobs.",
											MarkdownDescription: "The target write IOPS to run while benchmarking. This is a limit and there is no guarantee it will be reached. This is the total IOPS for all background write jobs.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"background_write_iops_jobs": schema.Int64Attribute{
											Description:         "Number of threads to use for background write IOPS. This should be set high enough to reach the target specified in BackgroundWriteIOPS. Example: If BackgroundWriteIOPS is 100 and write latency is 10ms then a single job would barely be able to reach 100 IOPS so this should be at least 2.",
											MarkdownDescription: "Number of threads to use for background write IOPS. This should be set high enough to reach the target specified in BackgroundWriteIOPS. Example: If BackgroundWriteIOPS is 100 and write latency is 10ms then a single job would barely be able to reach 100 IOPS so this should be at least 2.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"datasync": schema.BoolAttribute{
											Description:         "Whether to call datasync on the file after each write. Skipped if Sync is also true. Does not apply to background IOPS task.",
											MarkdownDescription: "Whether to call datasync on the file after each write. Skipped if Sync is also true. Does not apply to background IOPS task.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"directory": schema.StringAttribute{
											Description:         "The directory where the benchmark will create files.",
											MarkdownDescription: "The directory where the benchmark will create files.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enable_background_iops": schema.BoolAttribute{
											Description:         "Enable the background IOPS feature.",
											MarkdownDescription: "Enable the background IOPS feature.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"file_size": schema.StringAttribute{
											Description:         "The size of the file used in the benchmark. The number of IO operations for the benchmark will be FileSize / OperationSizeBytes. Accepts valid Kubernetes resource units such as Mi.",
											MarkdownDescription: "The size of the file used in the benchmark. The number of IO operations for the benchmark will be FileSize / OperationSizeBytes. Accepts valid Kubernetes resource units such as Mi.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"operation_size": schema.Int64Attribute{
											Description:         "The size of each write operation performed while benchmarking. This does not apply to the background IOPS feature if enabled, since those must be fixed at 4096.",
											MarkdownDescription: "The size of each write operation performed while benchmarking. This does not apply to the background IOPS feature if enabled, since those must be fixed at 4096.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"run_time": schema.StringAttribute{
											Description:         "Limit runtime. The test will run until it completes the configured I/O workload or until it has run for this specified amount of time, whichever occurs first. When the unit is omitted, the value is interpreted in seconds. Defaults to 120 seconds. Set to '0' to disable.",
											MarkdownDescription: "Limit runtime. The test will run until it completes the configured I/O workload or until it has run for this specified amount of time, whichever occurs first. When the unit is omitted, the value is interpreted in seconds. Defaults to 120 seconds. Set to '0' to disable.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sync": schema.BoolAttribute{
											Description:         "Whether to call sync on the file after each write. Does not apply to background IOPS task.",
											MarkdownDescription: "Whether to call sync on the file after each write. Does not apply to background IOPS task.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"timeout": schema.StringAttribute{
											Description:         "Total timeout, including background IOPS setup and warmup if enabled.",
											MarkdownDescription: "Total timeout, including background IOPS setup and warmup if enabled.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"host_os": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"host_services": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"http": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"get": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"headers": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"insecure_skip_verify": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"proxy": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"timeout": schema.StringAttribute{
													Description:         "Timeout is the time to wait for a server's response. Its a duration e.g 15s, 2h30m. Missing value or empty string or means no timeout.",
													MarkdownDescription: "Timeout is the time to wait for a server's response. Its a duration e.g 15s, 2h30m. Missing value or empty string or means no timeout.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tls": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"cacert": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"client_cert": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"client_key": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"secret": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"namespace": schema.StringAttribute{
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

														"skip_verify": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"url": schema.StringAttribute{
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

										"post": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"body": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"headers": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"insecure_skip_verify": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"proxy": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"timeout": schema.StringAttribute{
													Description:         "Timeout is the time to wait for a server's response. Its a duration e.g 15s, 2h30m. Missing value or empty string or means no timeout.",
													MarkdownDescription: "Timeout is the time to wait for a server's response. Its a duration e.g 15s, 2h30m. Missing value or empty string or means no timeout.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tls": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"cacert": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"client_cert": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"client_key": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"secret": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"namespace": schema.StringAttribute{
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

														"skip_verify": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"url": schema.StringAttribute{
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

										"put": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"body": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"headers": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"insecure_skip_verify": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"proxy": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"timeout": schema.StringAttribute{
													Description:         "Timeout is the time to wait for a server's response. Its a duration e.g 15s, 2h30m. Missing value or empty string or means no timeout.",
													MarkdownDescription: "Timeout is the time to wait for a server's response. Its a duration e.g 15s, 2h30m. Missing value or empty string or means no timeout.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tls": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"cacert": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"client_cert": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"client_key": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"secret": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"namespace": schema.StringAttribute{
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

														"skip_verify": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"url": schema.StringAttribute{
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"http_load_balancer": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"address": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"timeout": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"ipv4_interfaces": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"journald": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"dmesg": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"lines": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"output": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"reverse": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"since": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"system": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"timeout": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"units": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"until": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"utc": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"kernel_configs": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"kernel_modules": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"kubernetes": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"memory": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"network_namespace_connectivity": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"from_cidr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"timeout": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"to_cidr": schema.StringAttribute{
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

								"run": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"args": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"command": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"env": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ignore_parent_envs": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"inherit_envs": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"input": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"output_dir": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"timeout": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"subnet_available": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"cidr_range_alloc": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"desired_cidr": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"sysctl": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"system_packages": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"amzn": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"amzn2": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"centos": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"centos7": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"centos8": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"centos9": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ol": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ol7": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ol8": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ol9": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"rhel": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"rhel7": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"rhel8": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"rhel9": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"rocky": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"rocky8": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"rocky9": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ubuntu": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ubuntu16": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ubuntu18": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ubuntu20": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"tcp_connect": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"address": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"timeout": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"tcp_load_balancer": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"address": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"timeout": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"tcp_port_status": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"interface": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
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

								"time": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"udp_port_status": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"collector_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"exclude": UNKNOWN{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         UNKNOWN,
											CustomType:          UNKNOWN,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"interface": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"uri": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
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

func (r *TroubleshootShHostCollectorV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_troubleshoot_sh_host_collector_v1beta2_manifest")

	var model TroubleshootShHostCollectorV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("troubleshoot.sh/v1beta2")
	model.Kind = pointer.String("HostCollector")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
