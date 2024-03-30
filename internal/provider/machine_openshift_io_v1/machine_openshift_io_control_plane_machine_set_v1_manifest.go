/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package machine_openshift_io_v1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &MachineOpenshiftIoControlPlaneMachineSetV1Manifest{}
)

func NewMachineOpenshiftIoControlPlaneMachineSetV1Manifest() datasource.DataSource {
	return &MachineOpenshiftIoControlPlaneMachineSetV1Manifest{}
}

type MachineOpenshiftIoControlPlaneMachineSetV1Manifest struct{}

type MachineOpenshiftIoControlPlaneMachineSetV1ManifestData struct {
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
		Replicas *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
		Selector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
		State    *string `tfsdk:"state" json:"state,omitempty"`
		Strategy *struct {
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"strategy" json:"strategy,omitempty"`
		Template *struct {
			MachineType                           *string `tfsdk:"machine_type" json:"machineType,omitempty"`
			Machines_v1beta1_machine_openshift_io *struct {
				FailureDomains *struct {
					Aws *[]struct {
						Placement *struct {
							AvailabilityZone *string `tfsdk:"availability_zone" json:"availabilityZone,omitempty"`
						} `tfsdk:"placement" json:"placement,omitempty"`
						Subnet *struct {
							Arn     *string `tfsdk:"arn" json:"arn,omitempty"`
							Filters *[]struct {
								Name   *string   `tfsdk:"name" json:"name,omitempty"`
								Values *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"filters" json:"filters,omitempty"`
							Id   *string `tfsdk:"id" json:"id,omitempty"`
							Type *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"subnet" json:"subnet,omitempty"`
					} `tfsdk:"aws" json:"aws,omitempty"`
					Azure *[]struct {
						Subnet *string `tfsdk:"subnet" json:"subnet,omitempty"`
						Zone   *string `tfsdk:"zone" json:"zone,omitempty"`
					} `tfsdk:"azure" json:"azure,omitempty"`
					Gcp *[]struct {
						Zone *string `tfsdk:"zone" json:"zone,omitempty"`
					} `tfsdk:"gcp" json:"gcp,omitempty"`
					Nutanix *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"nutanix" json:"nutanix,omitempty"`
					Openstack *[]struct {
						AvailabilityZone *string `tfsdk:"availability_zone" json:"availabilityZone,omitempty"`
						RootVolume       *struct {
							AvailabilityZone *string `tfsdk:"availability_zone" json:"availabilityZone,omitempty"`
							VolumeType       *string `tfsdk:"volume_type" json:"volumeType,omitempty"`
						} `tfsdk:"root_volume" json:"rootVolume,omitempty"`
					} `tfsdk:"openstack" json:"openstack,omitempty"`
					Platform *string `tfsdk:"platform" json:"platform,omitempty"`
					Vsphere  *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"vsphere" json:"vsphere,omitempty"`
				} `tfsdk:"failure_domains" json:"failureDomains,omitempty"`
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec *struct {
					LifecycleHooks *struct {
						PreDrain *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Owner *string `tfsdk:"owner" json:"owner,omitempty"`
						} `tfsdk:"pre_drain" json:"preDrain,omitempty"`
						PreTerminate *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Owner *string `tfsdk:"owner" json:"owner,omitempty"`
						} `tfsdk:"pre_terminate" json:"preTerminate,omitempty"`
					} `tfsdk:"lifecycle_hooks" json:"lifecycleHooks,omitempty"`
					Metadata *struct {
						Annotations     *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						GenerateName    *string            `tfsdk:"generate_name" json:"generateName,omitempty"`
						Labels          *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						Name            *string            `tfsdk:"name" json:"name,omitempty"`
						Namespace       *string            `tfsdk:"namespace" json:"namespace,omitempty"`
						OwnerReferences *[]struct {
							ApiVersion         *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
							BlockOwnerDeletion *bool   `tfsdk:"block_owner_deletion" json:"blockOwnerDeletion,omitempty"`
							Controller         *bool   `tfsdk:"controller" json:"controller,omitempty"`
							Kind               *string `tfsdk:"kind" json:"kind,omitempty"`
							Name               *string `tfsdk:"name" json:"name,omitempty"`
							Uid                *string `tfsdk:"uid" json:"uid,omitempty"`
						} `tfsdk:"owner_references" json:"ownerReferences,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
					ProviderID   *string `tfsdk:"provider_id" json:"providerID,omitempty"`
					ProviderSpec *struct {
						Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"provider_spec" json:"providerSpec,omitempty"`
					Taints *[]struct {
						Effect    *string `tfsdk:"effect" json:"effect,omitempty"`
						Key       *string `tfsdk:"key" json:"key,omitempty"`
						TimeAdded *string `tfsdk:"time_added" json:"timeAdded,omitempty"`
						Value     *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"taints" json:"taints,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"machines_v1beta1_machine_openshift_io" json:"machines_v1beta1_machine_openshift_io,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MachineOpenshiftIoControlPlaneMachineSetV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_machine_openshift_io_control_plane_machine_set_v1_manifest"
}

func (r *MachineOpenshiftIoControlPlaneMachineSetV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ControlPlaneMachineSet ensures that a specified number of control plane machine replicas are running at any given time. Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "ControlPlaneMachineSet ensures that a specified number of control plane machine replicas are running at any given time. Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
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
				Description:         "ControlPlaneMachineSet represents the configuration of the ControlPlaneMachineSet.",
				MarkdownDescription: "ControlPlaneMachineSet represents the configuration of the ControlPlaneMachineSet.",
				Attributes: map[string]schema.Attribute{
					"replicas": schema.Int64Attribute{
						Description:         "Replicas defines how many Control Plane Machines should be created by this ControlPlaneMachineSet. This field is immutable and cannot be changed after cluster installation. The ControlPlaneMachineSet only operates with 3 or 5 node control planes, 3 and 5 are the only valid values for this field.",
						MarkdownDescription: "Replicas defines how many Control Plane Machines should be created by this ControlPlaneMachineSet. This field is immutable and cannot be changed after cluster installation. The ControlPlaneMachineSet only operates with 3 or 5 node control planes, 3 and 5 are the only valid values for this field.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.OneOf(3, 5),
						},
					},

					"selector": schema.SingleNestedAttribute{
						Description:         "Label selector for Machines. Existing Machines selected by this selector will be the ones affected by this ControlPlaneMachineSet. It must match the template's labels. This field is considered immutable after creation of the resource.",
						MarkdownDescription: "Label selector for Machines. Existing Machines selected by this selector will be the ones affected by this ControlPlaneMachineSet. It must match the template's labels. This field is considered immutable after creation of the resource.",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"state": schema.StringAttribute{
						Description:         "State defines whether the ControlPlaneMachineSet is Active or Inactive. When Inactive, the ControlPlaneMachineSet will not take any action on the state of the Machines within the cluster. When Active, the ControlPlaneMachineSet will reconcile the Machines and will update the Machines as necessary. Once Active, a ControlPlaneMachineSet cannot be made Inactive. To prevent further action please remove the ControlPlaneMachineSet.",
						MarkdownDescription: "State defines whether the ControlPlaneMachineSet is Active or Inactive. When Inactive, the ControlPlaneMachineSet will not take any action on the state of the Machines within the cluster. When Active, the ControlPlaneMachineSet will reconcile the Machines and will update the Machines as necessary. Once Active, a ControlPlaneMachineSet cannot be made Inactive. To prevent further action please remove the ControlPlaneMachineSet.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Active", "Inactive"),
						},
					},

					"strategy": schema.SingleNestedAttribute{
						Description:         "Strategy defines how the ControlPlaneMachineSet will update Machines when it detects a change to the ProviderSpec.",
						MarkdownDescription: "Strategy defines how the ControlPlaneMachineSet will update Machines when it detects a change to the ProviderSpec.",
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Description:         "Type defines the type of update strategy that should be used when updating Machines owned by the ControlPlaneMachineSet. Valid values are 'RollingUpdate' and 'OnDelete'. The current default value is 'RollingUpdate'.",
								MarkdownDescription: "Type defines the type of update strategy that should be used when updating Machines owned by the ControlPlaneMachineSet. Valid values are 'RollingUpdate' and 'OnDelete'. The current default value is 'RollingUpdate'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("RollingUpdate", "OnDelete"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"template": schema.SingleNestedAttribute{
						Description:         "Template describes the Control Plane Machines that will be created by this ControlPlaneMachineSet.",
						MarkdownDescription: "Template describes the Control Plane Machines that will be created by this ControlPlaneMachineSet.",
						Attributes: map[string]schema.Attribute{
							"machine_type": schema.StringAttribute{
								Description:         "MachineType determines the type of Machines that should be managed by the ControlPlaneMachineSet. Currently, the only valid value is machines_v1beta1_machine_openshift_io.",
								MarkdownDescription: "MachineType determines the type of Machines that should be managed by the ControlPlaneMachineSet. Currently, the only valid value is machines_v1beta1_machine_openshift_io.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("machines_v1beta1_machine_openshift_io"),
								},
							},

							"machines_v1beta1_machine_openshift_io": schema.SingleNestedAttribute{
								Description:         "OpenShiftMachineV1Beta1Machine defines the template for creating Machines from the v1beta1.machine.openshift.io API group.",
								MarkdownDescription: "OpenShiftMachineV1Beta1Machine defines the template for creating Machines from the v1beta1.machine.openshift.io API group.",
								Attributes: map[string]schema.Attribute{
									"failure_domains": schema.SingleNestedAttribute{
										Description:         "FailureDomains is the list of failure domains (sometimes called availability zones) in which the ControlPlaneMachineSet should balance the Control Plane Machines. This will be merged into the ProviderSpec given in the template. This field is optional on platforms that do not require placement information.",
										MarkdownDescription: "FailureDomains is the list of failure domains (sometimes called availability zones) in which the ControlPlaneMachineSet should balance the Control Plane Machines. This will be merged into the ProviderSpec given in the template. This field is optional on platforms that do not require placement information.",
										Attributes: map[string]schema.Attribute{
											"aws": schema.ListNestedAttribute{
												Description:         "AWS configures failure domain information for the AWS platform.",
												MarkdownDescription: "AWS configures failure domain information for the AWS platform.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"placement": schema.SingleNestedAttribute{
															Description:         "Placement configures the placement information for this instance.",
															MarkdownDescription: "Placement configures the placement information for this instance.",
															Attributes: map[string]schema.Attribute{
																"availability_zone": schema.StringAttribute{
																	Description:         "AvailabilityZone is the availability zone of the instance.",
																	MarkdownDescription: "AvailabilityZone is the availability zone of the instance.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
															Validators: []validator.Object{
																objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("subnet")),
															},
														},

														"subnet": schema.SingleNestedAttribute{
															Description:         "Subnet is a reference to the subnet to use for this instance.",
															MarkdownDescription: "Subnet is a reference to the subnet to use for this instance.",
															Attributes: map[string]schema.Attribute{
																"arn": schema.StringAttribute{
																	Description:         "ARN of resource.",
																	MarkdownDescription: "ARN of resource.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"filters": schema.ListNestedAttribute{
																	Description:         "Filters is a set of filters used to identify a resource.",
																	MarkdownDescription: "Filters is a set of filters used to identify a resource.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "Name of the filter. Filter names are case-sensitive.",
																				MarkdownDescription: "Name of the filter. Filter names are case-sensitive.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "Values includes one or more filter values. Filter values are case-sensitive.",
																				MarkdownDescription: "Values includes one or more filter values. Filter values are case-sensitive.",
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

																"id": schema.StringAttribute{
																	Description:         "ID of resource.",
																	MarkdownDescription: "ID of resource.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "Type determines how the reference will fetch the AWS resource.",
																	MarkdownDescription: "Type determines how the reference will fetch the AWS resource.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("ID", "ARN", "Filters"),
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
															Validators: []validator.Object{
																objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("placement")),
															},
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"azure": schema.ListNestedAttribute{
												Description:         "Azure configures failure domain information for the Azure platform.",
												MarkdownDescription: "Azure configures failure domain information for the Azure platform.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"subnet": schema.StringAttribute{
															Description:         "subnet is the name of the network subnet in which the VM will be created. When omitted, the subnet value from the machine providerSpec template will be used.",
															MarkdownDescription: "subnet is the name of the network subnet in which the VM will be created. When omitted, the subnet value from the machine providerSpec template will be used.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtMost(80),
																stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9](?:[a-zA-Z0-9._-]*[a-zA-Z0-9_])?$`), ""),
															},
														},

														"zone": schema.StringAttribute{
															Description:         "Availability Zone for the virtual machine. If nil, the virtual machine should be deployed to no zone.",
															MarkdownDescription: "Availability Zone for the virtual machine. If nil, the virtual machine should be deployed to no zone.",
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

											"gcp": schema.ListNestedAttribute{
												Description:         "GCP configures failure domain information for the GCP platform.",
												MarkdownDescription: "GCP configures failure domain information for the GCP platform.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"zone": schema.StringAttribute{
															Description:         "Zone is the zone in which the GCP machine provider will create the VM.",
															MarkdownDescription: "Zone is the zone in which the GCP machine provider will create the VM.",
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

											"nutanix": schema.ListNestedAttribute{
												Description:         "nutanix configures failure domain information for the Nutanix platform.",
												MarkdownDescription: "nutanix configures failure domain information for the Nutanix platform.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "name of the failure domain in which the nutanix machine provider will create the VM. Failure domains are defined in a cluster's config.openshift.io/Infrastructure resource.",
															MarkdownDescription: "name of the failure domain in which the nutanix machine provider will create the VM. Failure domains are defined in a cluster's config.openshift.io/Infrastructure resource.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.LengthAtMost(64),
																stringvalidator.RegexMatches(regexp.MustCompile(`[a-z0-9]([-a-z0-9]*[a-z0-9])?`), ""),
															},
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"openstack": schema.ListNestedAttribute{
												Description:         "OpenStack configures failure domain information for the OpenStack platform.",
												MarkdownDescription: "OpenStack configures failure domain information for the OpenStack platform.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"availability_zone": schema.StringAttribute{
															Description:         "availabilityZone is the nova availability zone in which the OpenStack machine provider will create the VM. If not specified, the VM will be created in the default availability zone specified in the nova configuration. Availability zone names must NOT contain : since it is used by admin users to specify hosts where instances are launched in server creation. Also, it must not contain spaces otherwise it will lead to node that belongs to this availability zone register failure, see kubernetes/cloud-provider-openstack#1379 for further information. The maximum length of availability zone name is 63 as per labels limits.",
															MarkdownDescription: "availabilityZone is the nova availability zone in which the OpenStack machine provider will create the VM. If not specified, the VM will be created in the default availability zone specified in the nova configuration. Availability zone names must NOT contain : since it is used by admin users to specify hosts where instances are launched in server creation. Also, it must not contain spaces otherwise it will lead to node that belongs to this availability zone register failure, see kubernetes/cloud-provider-openstack#1379 for further information. The maximum length of availability zone name is 63 as per labels limits.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.LengthAtMost(63),
																stringvalidator.RegexMatches(regexp.MustCompile(`^[^: ]*$`), ""),
																stringvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("root_volume")),
															},
														},

														"root_volume": schema.SingleNestedAttribute{
															Description:         "rootVolume contains settings that will be used by the OpenStack machine provider to create the root volume attached to the VM. If not specified, no root volume will be created.",
															MarkdownDescription: "rootVolume contains settings that will be used by the OpenStack machine provider to create the root volume attached to the VM. If not specified, no root volume will be created.",
															Attributes: map[string]schema.Attribute{
																"availability_zone": schema.StringAttribute{
																	Description:         "availabilityZone specifies the Cinder availability zone where the root volume will be created. If not specifified, the root volume will be created in the availability zone specified by the volume type in the cinder configuration. If the volume type (configured in the OpenStack cluster) does not specify an availability zone, the root volume will be created in the default availability zone specified in the cinder configuration. See https://docs.openstack.org/cinder/latest/admin/availability-zone-type.html for more details. If the OpenStack cluster is deployed with the cross_az_attach configuration option set to false, the root volume will have to be in the same availability zone as the VM (defined by OpenStackFailureDomain.AvailabilityZone). Availability zone names must NOT contain spaces otherwise it will lead to volume that belongs to this availability zone register failure, see kubernetes/cloud-provider-openstack#1379 for further information. The maximum length of availability zone name is 63 as per labels limits.",
																	MarkdownDescription: "availabilityZone specifies the Cinder availability zone where the root volume will be created. If not specifified, the root volume will be created in the availability zone specified by the volume type in the cinder configuration. If the volume type (configured in the OpenStack cluster) does not specify an availability zone, the root volume will be created in the default availability zone specified in the cinder configuration. See https://docs.openstack.org/cinder/latest/admin/availability-zone-type.html for more details. If the OpenStack cluster is deployed with the cross_az_attach configuration option set to false, the root volume will have to be in the same availability zone as the VM (defined by OpenStackFailureDomain.AvailabilityZone). Availability zone names must NOT contain spaces otherwise it will lead to volume that belongs to this availability zone register failure, see kubernetes/cloud-provider-openstack#1379 for further information. The maximum length of availability zone name is 63 as per labels limits.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(63),
																		stringvalidator.RegexMatches(regexp.MustCompile(`^[^ ]*$`), ""),
																	},
																},

																"volume_type": schema.StringAttribute{
																	Description:         "volumeType specifies the type of the root volume that will be provisioned. The maximum length of a volume type name is 255 characters, as per the OpenStack limit.",
																	MarkdownDescription: "volumeType specifies the type of the root volume that will be provisioned. The maximum length of a volume type name is 255 characters, as per the OpenStack limit.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(255),
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
															Validators: []validator.Object{
																objectvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("availability_zone")),
															},
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"platform": schema.StringAttribute{
												Description:         "Platform identifies the platform for which the FailureDomain represents. Currently supported values are AWS, Azure, GCP, OpenStack, VSphere and Nutanix.",
												MarkdownDescription: "Platform identifies the platform for which the FailureDomain represents. Currently supported values are AWS, Azure, GCP, OpenStack, VSphere and Nutanix.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("", "AWS", "Azure", "BareMetal", "GCP", "Libvirt", "OpenStack", "None", "VSphere", "oVirt", "IBMCloud", "KubeVirt", "EquinixMetal", "PowerVS", "AlibabaCloud", "Nutanix", "External"),
												},
											},

											"vsphere": schema.ListNestedAttribute{
												Description:         "vsphere configures failure domain information for the VSphere platform.",
												MarkdownDescription: "vsphere configures failure domain information for the VSphere platform.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "name of the failure domain in which the vSphere machine provider will create the VM. Failure domains are defined in a cluster's config.openshift.io/Infrastructure resource. When balancing machines across failure domains, the control plane machine set will inject configuration from the Infrastructure resource into the machine providerSpec to allocate the machine to a failure domain.",
															MarkdownDescription: "name of the failure domain in which the vSphere machine provider will create the VM. Failure domains are defined in a cluster's config.openshift.io/Infrastructure resource. When balancing machines across failure domains, the control plane machine set will inject configuration from the Infrastructure resource into the machine providerSpec to allocate the machine to a failure domain.",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"metadata": schema.SingleNestedAttribute{
										Description:         "ObjectMeta is the standard object metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata Labels are required to match the ControlPlaneMachineSet selector.",
										MarkdownDescription: "ObjectMeta is the standard object metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata Labels are required to match the ControlPlaneMachineSet selector.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
												MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels. This field must contain both the 'machine.openshift.io/cluster-api-machine-role' and 'machine.openshift.io/cluster-api-machine-type' labels, both with a value of 'master'. It must also contain a label with the key 'machine.openshift.io/cluster-api-cluster'.",
												MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels. This field must contain both the 'machine.openshift.io/cluster-api-machine-role' and 'machine.openshift.io/cluster-api-machine-type' labels, both with a value of 'master'. It must also contain a label with the key 'machine.openshift.io/cluster-api-cluster'.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"spec": schema.SingleNestedAttribute{
										Description:         "Spec contains the desired configuration of the Control Plane Machines. The ProviderSpec within contains platform specific details for creating the Control Plane Machines. The ProviderSe should be complete apart from the platform specific failure domain field. This will be overriden when the Machines are created based on the FailureDomains field.",
										MarkdownDescription: "Spec contains the desired configuration of the Control Plane Machines. The ProviderSpec within contains platform specific details for creating the Control Plane Machines. The ProviderSe should be complete apart from the platform specific failure domain field. This will be overriden when the Machines are created based on the FailureDomains field.",
										Attributes: map[string]schema.Attribute{
											"lifecycle_hooks": schema.SingleNestedAttribute{
												Description:         "LifecycleHooks allow users to pause operations on the machine at certain predefined points within the machine lifecycle.",
												MarkdownDescription: "LifecycleHooks allow users to pause operations on the machine at certain predefined points within the machine lifecycle.",
												Attributes: map[string]schema.Attribute{
													"pre_drain": schema.ListNestedAttribute{
														Description:         "PreDrain hooks prevent the machine from being drained. This also blocks further lifecycle events, such as termination.",
														MarkdownDescription: "PreDrain hooks prevent the machine from being drained. This also blocks further lifecycle events, such as termination.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name defines a unique name for the lifcycle hook. The name should be unique and descriptive, ideally 1-3 words, in CamelCase or it may be namespaced, eg. foo.example.com/CamelCase. Names must be unique and should only be managed by a single entity.",
																	MarkdownDescription: "Name defines a unique name for the lifcycle hook. The name should be unique and descriptive, ideally 1-3 words, in CamelCase or it may be namespaced, eg. foo.example.com/CamelCase. Names must be unique and should only be managed by a single entity.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(3),
																		stringvalidator.LengthAtMost(256),
																		stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$`), ""),
																	},
																},

																"owner": schema.StringAttribute{
																	Description:         "Owner defines the owner of the lifecycle hook. This should be descriptive enough so that users can identify who/what is responsible for blocking the lifecycle. This could be the name of a controller (e.g. clusteroperator/etcd) or an administrator managing the hook.",
																	MarkdownDescription: "Owner defines the owner of the lifecycle hook. This should be descriptive enough so that users can identify who/what is responsible for blocking the lifecycle. This could be the name of a controller (e.g. clusteroperator/etcd) or an administrator managing the hook.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(3),
																		stringvalidator.LengthAtMost(512),
																	},
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pre_terminate": schema.ListNestedAttribute{
														Description:         "PreTerminate hooks prevent the machine from being terminated. PreTerminate hooks be actioned after the Machine has been drained.",
														MarkdownDescription: "PreTerminate hooks prevent the machine from being terminated. PreTerminate hooks be actioned after the Machine has been drained.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name defines a unique name for the lifcycle hook. The name should be unique and descriptive, ideally 1-3 words, in CamelCase or it may be namespaced, eg. foo.example.com/CamelCase. Names must be unique and should only be managed by a single entity.",
																	MarkdownDescription: "Name defines a unique name for the lifcycle hook. The name should be unique and descriptive, ideally 1-3 words, in CamelCase or it may be namespaced, eg. foo.example.com/CamelCase. Names must be unique and should only be managed by a single entity.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(3),
																		stringvalidator.LengthAtMost(256),
																		stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$`), ""),
																	},
																},

																"owner": schema.StringAttribute{
																	Description:         "Owner defines the owner of the lifecycle hook. This should be descriptive enough so that users can identify who/what is responsible for blocking the lifecycle. This could be the name of a controller (e.g. clusteroperator/etcd) or an administrator managing the hook.",
																	MarkdownDescription: "Owner defines the owner of the lifecycle hook. This should be descriptive enough so that users can identify who/what is responsible for blocking the lifecycle. This could be the name of a controller (e.g. clusteroperator/etcd) or an administrator managing the hook.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(3),
																		stringvalidator.LengthAtMost(512),
																	},
																},
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

											"metadata": schema.SingleNestedAttribute{
												Description:         "ObjectMeta will autopopulate the Node created. Use this to indicate what labels, annotations, name prefix, etc., should be used when creating the Node.",
												MarkdownDescription: "ObjectMeta will autopopulate the Node created. Use this to indicate what labels, annotations, name prefix, etc., should be used when creating the Node.",
												Attributes: map[string]schema.Attribute{
													"annotations": schema.MapAttribute{
														Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
														MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"generate_name": schema.StringAttribute{
														Description:         "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.  If this field is specified and the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry (optionally after the time indicated in the Retry-After header).  Applied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#idempotency",
														MarkdownDescription: "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.  If this field is specified and the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry (optionally after the time indicated in the Retry-After header).  Applied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#idempotency",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"labels": schema.MapAttribute{
														Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
														MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
														MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace defines the space within each name must be unique. An empty namespace is equivalent to the 'default' namespace, but 'default' is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.  Must be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces",
														MarkdownDescription: "Namespace defines the space within each name must be unique. An empty namespace is equivalent to the 'default' namespace, but 'default' is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.  Must be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"owner_references": schema.ListNestedAttribute{
														Description:         "List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller.",
														MarkdownDescription: "List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"api_version": schema.StringAttribute{
																	Description:         "API version of the referent.",
																	MarkdownDescription: "API version of the referent.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"block_owner_deletion": schema.BoolAttribute{
																	Description:         "If true, AND if the owner has the 'foregroundDeletion' finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. See https://kubernetes.io/docs/concepts/architecture/garbage-collection/#foreground-deletion for how the garbage collector interacts with this field and enforces the foreground deletion. Defaults to false. To set this field, a user needs 'delete' permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.",
																	MarkdownDescription: "If true, AND if the owner has the 'foregroundDeletion' finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. See https://kubernetes.io/docs/concepts/architecture/garbage-collection/#foreground-deletion for how the garbage collector interacts with this field and enforces the foreground deletion. Defaults to false. To set this field, a user needs 'delete' permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"controller": schema.BoolAttribute{
																	Description:         "If true, this reference points to the managing controller.",
																	MarkdownDescription: "If true, this reference points to the managing controller.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"kind": schema.StringAttribute{
																	Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																	MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
																	MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"uid": schema.StringAttribute{
																	Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids",
																	MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids",
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"provider_id": schema.StringAttribute{
												Description:         "ProviderID is the identification ID of the machine provided by the provider. This field must match the provider ID as seen on the node object corresponding to this machine. This field is required by higher level consumers of cluster-api. Example use case is cluster autoscaler with cluster-api as provider. Clean-up logic in the autoscaler compares machines to nodes to find out machines at provider which could not get registered as Kubernetes nodes. With cluster-api as a generic out-of-tree provider for autoscaler, this field is required by autoscaler to be able to have a provider view of the list of machines. Another list of nodes is queried from the k8s apiserver and then a comparison is done to find out unregistered machines and are marked for delete. This field will be set by the actuators and consumed by higher level entities like autoscaler that will be interfacing with cluster-api as generic provider.",
												MarkdownDescription: "ProviderID is the identification ID of the machine provided by the provider. This field must match the provider ID as seen on the node object corresponding to this machine. This field is required by higher level consumers of cluster-api. Example use case is cluster autoscaler with cluster-api as provider. Clean-up logic in the autoscaler compares machines to nodes to find out machines at provider which could not get registered as Kubernetes nodes. With cluster-api as a generic out-of-tree provider for autoscaler, this field is required by autoscaler to be able to have a provider view of the list of machines. Another list of nodes is queried from the k8s apiserver and then a comparison is done to find out unregistered machines and are marked for delete. This field will be set by the actuators and consumed by higher level entities like autoscaler that will be interfacing with cluster-api as generic provider.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"provider_spec": schema.SingleNestedAttribute{
												Description:         "ProviderSpec details Provider-specific configuration to use during node creation.",
												MarkdownDescription: "ProviderSpec details Provider-specific configuration to use during node creation.",
												Attributes: map[string]schema.Attribute{
													"value": schema.MapAttribute{
														Description:         "Value is an inlined, serialized representation of the resource configuration. It is recommended that providers maintain their own versioned API types that should be serialized/deserialized from this field, akin to component config.",
														MarkdownDescription: "Value is an inlined, serialized representation of the resource configuration. It is recommended that providers maintain their own versioned API types that should be serialized/deserialized from this field, akin to component config.",
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

											"taints": schema.ListNestedAttribute{
												Description:         "The list of the taints to be applied to the corresponding Node in additive manner. This list will not overwrite any other taints added to the Node on an ongoing basis by other entities. These taints should be actively reconciled e.g. if you ask the machine controller to apply a taint and then manually remove the taint the machine controller will put it back) but not have the machine controller remove any taints",
												MarkdownDescription: "The list of the taints to be applied to the corresponding Node in additive manner. This list will not overwrite any other taints added to the Node on an ongoing basis by other entities. These taints should be actively reconciled e.g. if you ask the machine controller to apply a taint and then manually remove the taint the machine controller will put it back) but not have the machine controller remove any taints",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"effect": schema.StringAttribute{
															Description:         "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
															MarkdownDescription: "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"key": schema.StringAttribute{
															Description:         "Required. The taint key to be applied to a node.",
															MarkdownDescription: "Required. The taint key to be applied to a node.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"time_added": schema.StringAttribute{
															Description:         "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
															MarkdownDescription: "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																validators.DateTime64Validator(),
															},
														},

														"value": schema.StringAttribute{
															Description:         "The taint value corresponding to the taint key.",
															MarkdownDescription: "The taint value corresponding to the taint key.",
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
										},
										Required: true,
										Optional: false,
										Computed: false,
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *MachineOpenshiftIoControlPlaneMachineSetV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_machine_openshift_io_control_plane_machine_set_v1_manifest")

	var model MachineOpenshiftIoControlPlaneMachineSetV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("machine.openshift.io/v1")
	model.Kind = pointer.String("ControlPlaneMachineSet")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
