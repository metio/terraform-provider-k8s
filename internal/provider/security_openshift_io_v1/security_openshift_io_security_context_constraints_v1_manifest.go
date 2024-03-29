/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package security_openshift_io_v1

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
	_ datasource.DataSource = &SecurityOpenshiftIoSecurityContextConstraintsV1Manifest{}
)

func NewSecurityOpenshiftIoSecurityContextConstraintsV1Manifest() datasource.DataSource {
	return &SecurityOpenshiftIoSecurityContextConstraintsV1Manifest{}
}

type SecurityOpenshiftIoSecurityContextConstraintsV1Manifest struct{}

type SecurityOpenshiftIoSecurityContextConstraintsV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	AllowHostDirVolumePlugin *bool     `tfsdk:"allow_host_dir_volume_plugin" json:"allowHostDirVolumePlugin,omitempty"`
	AllowHostIPC             *bool     `tfsdk:"allow_host_ipc" json:"allowHostIPC,omitempty"`
	AllowHostNetwork         *bool     `tfsdk:"allow_host_network" json:"allowHostNetwork,omitempty"`
	AllowHostPID             *bool     `tfsdk:"allow_host_pid" json:"allowHostPID,omitempty"`
	AllowHostPorts           *bool     `tfsdk:"allow_host_ports" json:"allowHostPorts,omitempty"`
	AllowPrivilegeEscalation *bool     `tfsdk:"allow_privilege_escalation" json:"allowPrivilegeEscalation,omitempty"`
	AllowPrivilegedContainer *bool     `tfsdk:"allow_privileged_container" json:"allowPrivilegedContainer,omitempty"`
	AllowedCapabilities      *[]string `tfsdk:"allowed_capabilities" json:"allowedCapabilities,omitempty"`
	AllowedFlexVolumes       *[]struct {
		Driver *string `tfsdk:"driver" json:"driver,omitempty"`
	} `tfsdk:"allowed_flex_volumes" json:"allowedFlexVolumes,omitempty"`
	AllowedUnsafeSysctls            *[]string `tfsdk:"allowed_unsafe_sysctls" json:"allowedUnsafeSysctls,omitempty"`
	DefaultAddCapabilities          *[]string `tfsdk:"default_add_capabilities" json:"defaultAddCapabilities,omitempty"`
	DefaultAllowPrivilegeEscalation *bool     `tfsdk:"default_allow_privilege_escalation" json:"defaultAllowPrivilegeEscalation,omitempty"`
	ForbiddenSysctls                *[]string `tfsdk:"forbidden_sysctls" json:"forbiddenSysctls,omitempty"`
	FsGroup                         *struct {
		Ranges *[]struct {
			Max *int64 `tfsdk:"max" json:"max,omitempty"`
			Min *int64 `tfsdk:"min" json:"min,omitempty"`
		} `tfsdk:"ranges" json:"ranges,omitempty"`
		Type *string `tfsdk:"type" json:"type,omitempty"`
	} `tfsdk:"fs_group" json:"fsGroup,omitempty"`
	Groups                   *[]string `tfsdk:"groups" json:"groups,omitempty"`
	Priority                 *int64    `tfsdk:"priority" json:"priority,omitempty"`
	ReadOnlyRootFilesystem   *bool     `tfsdk:"read_only_root_filesystem" json:"readOnlyRootFilesystem,omitempty"`
	RequiredDropCapabilities *[]string `tfsdk:"required_drop_capabilities" json:"requiredDropCapabilities,omitempty"`
	RunAsUser                *struct {
		Type        *string `tfsdk:"type" json:"type,omitempty"`
		Uid         *int64  `tfsdk:"uid" json:"uid,omitempty"`
		UidRangeMax *int64  `tfsdk:"uid_range_max" json:"uidRangeMax,omitempty"`
		UidRangeMin *int64  `tfsdk:"uid_range_min" json:"uidRangeMin,omitempty"`
	} `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
	SeLinuxContext *struct {
		SeLinuxOptions *struct {
			Level *string `tfsdk:"level" json:"level,omitempty"`
			Role  *string `tfsdk:"role" json:"role,omitempty"`
			Type  *string `tfsdk:"type" json:"type,omitempty"`
			User  *string `tfsdk:"user" json:"user,omitempty"`
		} `tfsdk:"se_linux_options" json:"seLinuxOptions,omitempty"`
		Type *string `tfsdk:"type" json:"type,omitempty"`
	} `tfsdk:"se_linux_context" json:"seLinuxContext,omitempty"`
	SeccompProfiles    *[]string `tfsdk:"seccomp_profiles" json:"seccompProfiles,omitempty"`
	SupplementalGroups *struct {
		Ranges *[]struct {
			Max *int64 `tfsdk:"max" json:"max,omitempty"`
			Min *int64 `tfsdk:"min" json:"min,omitempty"`
		} `tfsdk:"ranges" json:"ranges,omitempty"`
		Type *string `tfsdk:"type" json:"type,omitempty"`
	} `tfsdk:"supplemental_groups" json:"supplementalGroups,omitempty"`
	Users   *[]string `tfsdk:"users" json:"users,omitempty"`
	Volumes *[]string `tfsdk:"volumes" json:"volumes,omitempty"`
}

func (r *SecurityOpenshiftIoSecurityContextConstraintsV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_security_openshift_io_security_context_constraints_v1_manifest"
}

func (r *SecurityOpenshiftIoSecurityContextConstraintsV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SecurityContextConstraints governs the ability to make requests that affect the SecurityContext that will be applied to a container. For historical reasons SCC was exposed under the core Kubernetes API group. That exposure is deprecated and will be removed in a future release - users should instead use the security.openshift.io group to manage SecurityContextConstraints.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "SecurityContextConstraints governs the ability to make requests that affect the SecurityContext that will be applied to a container. For historical reasons SCC was exposed under the core Kubernetes API group. That exposure is deprecated and will be removed in a future release - users should instead use the security.openshift.io group to manage SecurityContextConstraints.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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

			"allow_host_dir_volume_plugin": schema.BoolAttribute{
				Description:         "AllowHostDirVolumePlugin determines if the policy allow containers to use the HostDir volume plugin",
				MarkdownDescription: "AllowHostDirVolumePlugin determines if the policy allow containers to use the HostDir volume plugin",
				Required:            true,
				Optional:            false,
				Computed:            false,
			},

			"allow_host_ipc": schema.BoolAttribute{
				Description:         "AllowHostIPC determines if the policy allows host ipc in the containers.",
				MarkdownDescription: "AllowHostIPC determines if the policy allows host ipc in the containers.",
				Required:            true,
				Optional:            false,
				Computed:            false,
			},

			"allow_host_network": schema.BoolAttribute{
				Description:         "AllowHostNetwork determines if the policy allows the use of HostNetwork in the pod spec.",
				MarkdownDescription: "AllowHostNetwork determines if the policy allows the use of HostNetwork in the pod spec.",
				Required:            true,
				Optional:            false,
				Computed:            false,
			},

			"allow_host_pid": schema.BoolAttribute{
				Description:         "AllowHostPID determines if the policy allows host pid in the containers.",
				MarkdownDescription: "AllowHostPID determines if the policy allows host pid in the containers.",
				Required:            true,
				Optional:            false,
				Computed:            false,
			},

			"allow_host_ports": schema.BoolAttribute{
				Description:         "AllowHostPorts determines if the policy allows host ports in the containers.",
				MarkdownDescription: "AllowHostPorts determines if the policy allows host ports in the containers.",
				Required:            true,
				Optional:            false,
				Computed:            false,
			},

			"allow_privilege_escalation": schema.BoolAttribute{
				Description:         "AllowPrivilegeEscalation determines if a pod can request to allow privilege escalation. If unspecified, defaults to true.",
				MarkdownDescription: "AllowPrivilegeEscalation determines if a pod can request to allow privilege escalation. If unspecified, defaults to true.",
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"allow_privileged_container": schema.BoolAttribute{
				Description:         "AllowPrivilegedContainer determines if a container can request to be run as privileged.",
				MarkdownDescription: "AllowPrivilegedContainer determines if a container can request to be run as privileged.",
				Required:            true,
				Optional:            false,
				Computed:            false,
			},

			"allowed_capabilities": schema.ListAttribute{
				Description:         "AllowedCapabilities is a list of capabilities that can be requested to add to the container. Capabilities in this field maybe added at the pod author's discretion. You must not list a capability in both AllowedCapabilities and RequiredDropCapabilities. To allow all capabilities you may use '*'.",
				MarkdownDescription: "AllowedCapabilities is a list of capabilities that can be requested to add to the container. Capabilities in this field maybe added at the pod author's discretion. You must not list a capability in both AllowedCapabilities and RequiredDropCapabilities. To allow all capabilities you may use '*'.",
				ElementType:         types.StringType,
				Required:            true,
				Optional:            false,
				Computed:            false,
			},

			"allowed_flex_volumes": schema.ListNestedAttribute{
				Description:         "AllowedFlexVolumes is a whitelist of allowed Flexvolumes.  Empty or nil indicates that all Flexvolumes may be used.  This parameter is effective only when the usage of the Flexvolumes is allowed in the 'Volumes' field.",
				MarkdownDescription: "AllowedFlexVolumes is a whitelist of allowed Flexvolumes.  Empty or nil indicates that all Flexvolumes may be used.  This parameter is effective only when the usage of the Flexvolumes is allowed in the 'Volumes' field.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"driver": schema.StringAttribute{
							Description:         "Driver is the name of the Flexvolume driver.",
							MarkdownDescription: "Driver is the name of the Flexvolume driver.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},

			"allowed_unsafe_sysctls": schema.ListAttribute{
				Description:         "AllowedUnsafeSysctls is a list of explicitly allowed unsafe sysctls, defaults to none. Each entry is either a plain sysctl name or ends in '*' in which case it is considered as a prefix of allowed sysctls. Single * means all unsafe sysctls are allowed. Kubelet has to whitelist all allowed unsafe sysctls explicitly to avoid rejection.  Examples: e.g. 'foo/*' allows 'foo/bar', 'foo/baz', etc. e.g. 'foo.*' allows 'foo.bar', 'foo.baz', etc.",
				MarkdownDescription: "AllowedUnsafeSysctls is a list of explicitly allowed unsafe sysctls, defaults to none. Each entry is either a plain sysctl name or ends in '*' in which case it is considered as a prefix of allowed sysctls. Single * means all unsafe sysctls are allowed. Kubelet has to whitelist all allowed unsafe sysctls explicitly to avoid rejection.  Examples: e.g. 'foo/*' allows 'foo/bar', 'foo/baz', etc. e.g. 'foo.*' allows 'foo.bar', 'foo.baz', etc.",
				ElementType:         types.StringType,
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"default_add_capabilities": schema.ListAttribute{
				Description:         "DefaultAddCapabilities is the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability.  You may not list a capabiility in both DefaultAddCapabilities and RequiredDropCapabilities.",
				MarkdownDescription: "DefaultAddCapabilities is the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability.  You may not list a capabiility in both DefaultAddCapabilities and RequiredDropCapabilities.",
				ElementType:         types.StringType,
				Required:            true,
				Optional:            false,
				Computed:            false,
			},

			"default_allow_privilege_escalation": schema.BoolAttribute{
				Description:         "DefaultAllowPrivilegeEscalation controls the default setting for whether a process can gain more privileges than its parent process.",
				MarkdownDescription: "DefaultAllowPrivilegeEscalation controls the default setting for whether a process can gain more privileges than its parent process.",
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"forbidden_sysctls": schema.ListAttribute{
				Description:         "ForbiddenSysctls is a list of explicitly forbidden sysctls, defaults to none. Each entry is either a plain sysctl name or ends in '*' in which case it is considered as a prefix of forbidden sysctls. Single * means all sysctls are forbidden.  Examples: e.g. 'foo/*' forbids 'foo/bar', 'foo/baz', etc. e.g. 'foo.*' forbids 'foo.bar', 'foo.baz', etc.",
				MarkdownDescription: "ForbiddenSysctls is a list of explicitly forbidden sysctls, defaults to none. Each entry is either a plain sysctl name or ends in '*' in which case it is considered as a prefix of forbidden sysctls. Single * means all sysctls are forbidden.  Examples: e.g. 'foo/*' forbids 'foo/bar', 'foo/baz', etc. e.g. 'foo.*' forbids 'foo.bar', 'foo.baz', etc.",
				ElementType:         types.StringType,
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"fs_group": schema.SingleNestedAttribute{
				Description:         "FSGroup is the strategy that will dictate what fs group is used by the SecurityContext.",
				MarkdownDescription: "FSGroup is the strategy that will dictate what fs group is used by the SecurityContext.",
				Attributes: map[string]schema.Attribute{
					"ranges": schema.ListNestedAttribute{
						Description:         "Ranges are the allowed ranges of fs groups.  If you would like to force a single fs group then supply a single range with the same start and end.",
						MarkdownDescription: "Ranges are the allowed ranges of fs groups.  If you would like to force a single fs group then supply a single range with the same start and end.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"max": schema.Int64Attribute{
									Description:         "Max is the end of the range, inclusive.",
									MarkdownDescription: "Max is the end of the range, inclusive.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"min": schema.Int64Attribute{
									Description:         "Min is the start of the range, inclusive.",
									MarkdownDescription: "Min is the start of the range, inclusive.",
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

					"type": schema.StringAttribute{
						Description:         "Type is the strategy that will dictate what FSGroup is used in the SecurityContext.",
						MarkdownDescription: "Type is the strategy that will dictate what FSGroup is used in the SecurityContext.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},

			"groups": schema.ListAttribute{
				Description:         "The groups that have permission to use this security context constraints",
				MarkdownDescription: "The groups that have permission to use this security context constraints",
				ElementType:         types.StringType,
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"priority": schema.Int64Attribute{
				Description:         "Priority influences the sort order of SCCs when evaluating which SCCs to try first for a given pod request based on access in the Users and Groups fields.  The higher the int, the higher priority. An unset value is considered a 0 priority. If scores for multiple SCCs are equal they will be sorted from most restrictive to least restrictive. If both priorities and restrictions are equal the SCCs will be sorted by name.",
				MarkdownDescription: "Priority influences the sort order of SCCs when evaluating which SCCs to try first for a given pod request based on access in the Users and Groups fields.  The higher the int, the higher priority. An unset value is considered a 0 priority. If scores for multiple SCCs are equal they will be sorted from most restrictive to least restrictive. If both priorities and restrictions are equal the SCCs will be sorted by name.",
				Required:            true,
				Optional:            false,
				Computed:            false,
			},

			"read_only_root_filesystem": schema.BoolAttribute{
				Description:         "ReadOnlyRootFilesystem when set to true will force containers to run with a read only root file system.  If the container specifically requests to run with a non-read only root file system the SCC should deny the pod. If set to false the container may run with a read only root file system if it wishes but it will not be forced to.",
				MarkdownDescription: "ReadOnlyRootFilesystem when set to true will force containers to run with a read only root file system.  If the container specifically requests to run with a non-read only root file system the SCC should deny the pod. If set to false the container may run with a read only root file system if it wishes but it will not be forced to.",
				Required:            true,
				Optional:            false,
				Computed:            false,
			},

			"required_drop_capabilities": schema.ListAttribute{
				Description:         "RequiredDropCapabilities are the capabilities that will be dropped from the container.  These are required to be dropped and cannot be added.",
				MarkdownDescription: "RequiredDropCapabilities are the capabilities that will be dropped from the container.  These are required to be dropped and cannot be added.",
				ElementType:         types.StringType,
				Required:            true,
				Optional:            false,
				Computed:            false,
			},

			"run_as_user": schema.SingleNestedAttribute{
				Description:         "RunAsUser is the strategy that will dictate what RunAsUser is used in the SecurityContext.",
				MarkdownDescription: "RunAsUser is the strategy that will dictate what RunAsUser is used in the SecurityContext.",
				Attributes: map[string]schema.Attribute{
					"type": schema.StringAttribute{
						Description:         "Type is the strategy that will dictate what RunAsUser is used in the SecurityContext.",
						MarkdownDescription: "Type is the strategy that will dictate what RunAsUser is used in the SecurityContext.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"uid": schema.Int64Attribute{
						Description:         "UID is the user id that containers must run as.  Required for the MustRunAs strategy if not using namespace/service account allocated uids.",
						MarkdownDescription: "UID is the user id that containers must run as.  Required for the MustRunAs strategy if not using namespace/service account allocated uids.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"uid_range_max": schema.Int64Attribute{
						Description:         "UIDRangeMax defines the max value for a strategy that allocates by range.",
						MarkdownDescription: "UIDRangeMax defines the max value for a strategy that allocates by range.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"uid_range_min": schema.Int64Attribute{
						Description:         "UIDRangeMin defines the min value for a strategy that allocates by range.",
						MarkdownDescription: "UIDRangeMin defines the min value for a strategy that allocates by range.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},

			"se_linux_context": schema.SingleNestedAttribute{
				Description:         "SELinuxContext is the strategy that will dictate what labels will be set in the SecurityContext.",
				MarkdownDescription: "SELinuxContext is the strategy that will dictate what labels will be set in the SecurityContext.",
				Attributes: map[string]schema.Attribute{
					"se_linux_options": schema.SingleNestedAttribute{
						Description:         "seLinuxOptions required to run as; required for MustRunAs",
						MarkdownDescription: "seLinuxOptions required to run as; required for MustRunAs",
						Attributes: map[string]schema.Attribute{
							"level": schema.StringAttribute{
								Description:         "Level is SELinux level label that applies to the container.",
								MarkdownDescription: "Level is SELinux level label that applies to the container.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"role": schema.StringAttribute{
								Description:         "Role is a SELinux role label that applies to the container.",
								MarkdownDescription: "Role is a SELinux role label that applies to the container.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type is a SELinux type label that applies to the container.",
								MarkdownDescription: "Type is a SELinux type label that applies to the container.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"user": schema.StringAttribute{
								Description:         "User is a SELinux user label that applies to the container.",
								MarkdownDescription: "User is a SELinux user label that applies to the container.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"type": schema.StringAttribute{
						Description:         "Type is the strategy that will dictate what SELinux context is used in the SecurityContext.",
						MarkdownDescription: "Type is the strategy that will dictate what SELinux context is used in the SecurityContext.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},

			"seccomp_profiles": schema.ListAttribute{
				Description:         "SeccompProfiles lists the allowed profiles that may be set for the pod or container's seccomp annotations.  An unset (nil) or empty value means that no profiles may be specifid by the pod or container.	The wildcard '*' may be used to allow all profiles.  When used to generate a value for a pod the first non-wildcard profile will be used as the default.",
				MarkdownDescription: "SeccompProfiles lists the allowed profiles that may be set for the pod or container's seccomp annotations.  An unset (nil) or empty value means that no profiles may be specifid by the pod or container.	The wildcard '*' may be used to allow all profiles.  When used to generate a value for a pod the first non-wildcard profile will be used as the default.",
				ElementType:         types.StringType,
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"supplemental_groups": schema.SingleNestedAttribute{
				Description:         "SupplementalGroups is the strategy that will dictate what supplemental groups are used by the SecurityContext.",
				MarkdownDescription: "SupplementalGroups is the strategy that will dictate what supplemental groups are used by the SecurityContext.",
				Attributes: map[string]schema.Attribute{
					"ranges": schema.ListNestedAttribute{
						Description:         "Ranges are the allowed ranges of supplemental groups.  If you would like to force a single supplemental group then supply a single range with the same start and end.",
						MarkdownDescription: "Ranges are the allowed ranges of supplemental groups.  If you would like to force a single supplemental group then supply a single range with the same start and end.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"max": schema.Int64Attribute{
									Description:         "Max is the end of the range, inclusive.",
									MarkdownDescription: "Max is the end of the range, inclusive.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"min": schema.Int64Attribute{
									Description:         "Min is the start of the range, inclusive.",
									MarkdownDescription: "Min is the start of the range, inclusive.",
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

					"type": schema.StringAttribute{
						Description:         "Type is the strategy that will dictate what supplemental groups is used in the SecurityContext.",
						MarkdownDescription: "Type is the strategy that will dictate what supplemental groups is used in the SecurityContext.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},

			"users": schema.ListAttribute{
				Description:         "The users who have permissions to use this security context constraints",
				MarkdownDescription: "The users who have permissions to use this security context constraints",
				ElementType:         types.StringType,
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"volumes": schema.ListAttribute{
				Description:         "Volumes is a white list of allowed volume plugins.  FSType corresponds directly with the field names of a VolumeSource (azureFile, configMap, emptyDir).  To allow all volumes you may use '*'. To allow no volumes, set to ['none'].",
				MarkdownDescription: "Volumes is a white list of allowed volume plugins.  FSType corresponds directly with the field names of a VolumeSource (azureFile, configMap, emptyDir).  To allow all volumes you may use '*'. To allow no volumes, set to ['none'].",
				ElementType:         types.StringType,
				Required:            true,
				Optional:            false,
				Computed:            false,
			},
		},
	}
}

func (r *SecurityOpenshiftIoSecurityContextConstraintsV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_security_openshift_io_security_context_constraints_v1_manifest")

	var model SecurityOpenshiftIoSecurityContextConstraintsV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("security.openshift.io/v1")
	model.Kind = pointer.String("SecurityContextConstraints")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
