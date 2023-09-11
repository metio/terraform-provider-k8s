/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ec2_services_k8s_aws_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
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
	_ resource.Resource                = &Ec2ServicesK8SAwsVpcV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &Ec2ServicesK8SAwsVpcV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &Ec2ServicesK8SAwsVpcV1Alpha1Resource{}
)

func NewEc2ServicesK8SAwsVpcV1Alpha1Resource() resource.Resource {
	return &Ec2ServicesK8SAwsVpcV1Alpha1Resource{}
}

type Ec2ServicesK8SAwsVpcV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type Ec2ServicesK8SAwsVpcV1Alpha1ResourceData struct {
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
		AmazonProvidedIPv6CIDRBlock     *bool     `tfsdk:"amazon_provided_i_pv6_cidr_block" json:"amazonProvidedIPv6CIDRBlock,omitempty"`
		CidrBlocks                      *[]string `tfsdk:"cidr_blocks" json:"cidrBlocks,omitempty"`
		EnableDNSHostnames              *bool     `tfsdk:"enable_dns_hostnames" json:"enableDNSHostnames,omitempty"`
		EnableDNSSupport                *bool     `tfsdk:"enable_dns_support" json:"enableDNSSupport,omitempty"`
		InstanceTenancy                 *string   `tfsdk:"instance_tenancy" json:"instanceTenancy,omitempty"`
		Ipv4IPAMPoolID                  *string   `tfsdk:"ipv4_ipam_pool_id" json:"ipv4IPAMPoolID,omitempty"`
		Ipv4NetmaskLength               *int64    `tfsdk:"ipv4_netmask_length" json:"ipv4NetmaskLength,omitempty"`
		Ipv6CIDRBlock                   *string   `tfsdk:"ipv6_cidr_block" json:"ipv6CIDRBlock,omitempty"`
		Ipv6CIDRBlockNetworkBorderGroup *string   `tfsdk:"ipv6_cidr_block_network_border_group" json:"ipv6CIDRBlockNetworkBorderGroup,omitempty"`
		Ipv6IPAMPoolID                  *string   `tfsdk:"ipv6_ipam_pool_id" json:"ipv6IPAMPoolID,omitempty"`
		Ipv6NetmaskLength               *int64    `tfsdk:"ipv6_netmask_length" json:"ipv6NetmaskLength,omitempty"`
		Ipv6Pool                        *string   `tfsdk:"ipv6_pool" json:"ipv6Pool,omitempty"`
		Tags                            *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Ec2ServicesK8SAwsVpcV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ec2_services_k8s_aws_vpc_v1alpha1"
}

func (r *Ec2ServicesK8SAwsVpcV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VPC is the Schema for the VPCS API",
		MarkdownDescription: "VPC is the Schema for the VPCS API",
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
				Description:         "VpcSpec defines the desired state of Vpc.  Describes a VPC.",
				MarkdownDescription: "VpcSpec defines the desired state of Vpc.  Describes a VPC.",
				Attributes: map[string]schema.Attribute{
					"amazon_provided_i_pv6_cidr_block": schema.BoolAttribute{
						Description:         "Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or the size of the CIDR block.",
						MarkdownDescription: "Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or the size of the CIDR block.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cidr_blocks": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"enable_dns_hostnames": schema.BoolAttribute{
						Description:         "The attribute value. The valid values are true or false.",
						MarkdownDescription: "The attribute value. The valid values are true or false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_dns_support": schema.BoolAttribute{
						Description:         "The attribute value. The valid values are true or false.",
						MarkdownDescription: "The attribute value. The valid values are true or false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"instance_tenancy": schema.StringAttribute{
						Description:         "The tenancy options for instances launched into the VPC. For default, instances are launched with shared tenancy by default. You can launch instances with any tenancy into a shared tenancy VPC. For dedicated, instances are launched as dedicated tenancy instances by default. You can only launch instances with a tenancy of dedicated or host into a dedicated tenancy VPC.  Important: The host value cannot be used with this parameter. Use the default or dedicated values only.  Default: default",
						MarkdownDescription: "The tenancy options for instances launched into the VPC. For default, instances are launched with shared tenancy by default. You can launch instances with any tenancy into a shared tenancy VPC. For dedicated, instances are launched as dedicated tenancy instances by default. You can only launch instances with a tenancy of dedicated or host into a dedicated tenancy VPC.  Important: The host value cannot be used with this parameter. Use the default or dedicated values only.  Default: default",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipv4_ipam_pool_id": schema.StringAttribute{
						Description:         "The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. For more information, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						MarkdownDescription: "The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. For more information, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipv4_netmask_length": schema.Int64Attribute{
						Description:         "The netmask length of the IPv4 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool. For more information about IPAM, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						MarkdownDescription: "The netmask length of the IPv4 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool. For more information about IPAM, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipv6_cidr_block": schema.StringAttribute{
						Description:         "The IPv6 CIDR block from the IPv6 address pool. You must also specify Ipv6Pool in the request.  To let Amazon choose the IPv6 CIDR block for you, omit this parameter.",
						MarkdownDescription: "The IPv6 CIDR block from the IPv6 address pool. You must also specify Ipv6Pool in the request.  To let Amazon choose the IPv6 CIDR block for you, omit this parameter.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipv6_cidr_block_network_border_group": schema.StringAttribute{
						Description:         "The name of the location from which we advertise the IPV6 CIDR block. Use this parameter to limit the address to this location.  You must set AmazonProvidedIpv6CidrBlock to true to use this parameter.",
						MarkdownDescription: "The name of the location from which we advertise the IPV6 CIDR block. Use this parameter to limit the address to this location.  You must set AmazonProvidedIpv6CidrBlock to true to use this parameter.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipv6_ipam_pool_id": schema.StringAttribute{
						Description:         "The ID of an IPv6 IPAM pool which will be used to allocate this VPC an IPv6 CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across Amazon Web Services Regions and accounts throughout your Amazon Web Services Organization. For more information, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						MarkdownDescription: "The ID of an IPv6 IPAM pool which will be used to allocate this VPC an IPv6 CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across Amazon Web Services Regions and accounts throughout your Amazon Web Services Organization. For more information, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipv6_netmask_length": schema.Int64Attribute{
						Description:         "The netmask length of the IPv6 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool. For more information about IPAM, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						MarkdownDescription: "The netmask length of the IPv6 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool. For more information about IPAM, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipv6_pool": schema.StringAttribute{
						Description:         "The ID of an IPv6 address pool from which to allocate the IPv6 CIDR block.",
						MarkdownDescription: "The ID of an IPv6 address pool from which to allocate the IPv6 CIDR block.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "The tags. The value parameter is required, but if you don't want the tag to have a value, specify the parameter with no value, and we set the value to an empty string.",
						MarkdownDescription: "The tags. The value parameter is required, but if you don't want the tag to have a value, specify the parameter with no value, and we set the value to an empty string.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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
		},
	}
}

func (r *Ec2ServicesK8SAwsVpcV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *Ec2ServicesK8SAwsVpcV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_ec2_services_k8s_aws_vpc_v1alpha1")

	var model Ec2ServicesK8SAwsVpcV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("ec2.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("VPC")

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

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "ec2.services.k8s.aws", Version: "v1alpha1", Resource: "vpcs"}).
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

	var readResponse Ec2ServicesK8SAwsVpcV1Alpha1ResourceData
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

func (r *Ec2ServicesK8SAwsVpcV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ec2_services_k8s_aws_vpc_v1alpha1")

	var data Ec2ServicesK8SAwsVpcV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "ec2.services.k8s.aws", Version: "v1alpha1", Resource: "vpcs"}).
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

	var readResponse Ec2ServicesK8SAwsVpcV1Alpha1ResourceData
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

func (r *Ec2ServicesK8SAwsVpcV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_ec2_services_k8s_aws_vpc_v1alpha1")

	var model Ec2ServicesK8SAwsVpcV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("ec2.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("VPC")

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

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "ec2.services.k8s.aws", Version: "v1alpha1", Resource: "vpcs"}).
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

	var readResponse Ec2ServicesK8SAwsVpcV1Alpha1ResourceData
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

func (r *Ec2ServicesK8SAwsVpcV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_ec2_services_k8s_aws_vpc_v1alpha1")

	var data Ec2ServicesK8SAwsVpcV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "ec2.services.k8s.aws", Version: "v1alpha1", Resource: "vpcs"}).
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

func (r *Ec2ServicesK8SAwsVpcV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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
