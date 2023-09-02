/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package lb_lbconfig_carlosedp_com_v1

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
	_ resource.Resource                = &LbLbconfigCarlosedpComExternalLoadBalancerV1Resource{}
	_ resource.ResourceWithConfigure   = &LbLbconfigCarlosedpComExternalLoadBalancerV1Resource{}
	_ resource.ResourceWithImportState = &LbLbconfigCarlosedpComExternalLoadBalancerV1Resource{}
)

func NewLbLbconfigCarlosedpComExternalLoadBalancerV1Resource() resource.Resource {
	return &LbLbconfigCarlosedpComExternalLoadBalancerV1Resource{}
}

type LbLbconfigCarlosedpComExternalLoadBalancerV1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type LbLbconfigCarlosedpComExternalLoadBalancerV1ResourceData struct {
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
		Monitor *struct {
			Monitortype *string `tfsdk:"monitortype" json:"monitortype,omitempty"`
			Name        *string `tfsdk:"name" json:"name,omitempty"`
			Path        *string `tfsdk:"path" json:"path,omitempty"`
			Port        *int64  `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"monitor" json:"monitor,omitempty"`
		Nodelabels *map[string]string `tfsdk:"nodelabels" json:"nodelabels,omitempty"`
		Ports      *[]string          `tfsdk:"ports" json:"ports,omitempty"`
		Provider   *struct {
			Creds         *string `tfsdk:"creds" json:"creds,omitempty"`
			Debug         *bool   `tfsdk:"debug" json:"debug,omitempty"`
			Host          *string `tfsdk:"host" json:"host,omitempty"`
			Lbmethod      *string `tfsdk:"lbmethod" json:"lbmethod,omitempty"`
			Partition     *string `tfsdk:"partition" json:"partition,omitempty"`
			Port          *int64  `tfsdk:"port" json:"port,omitempty"`
			Validatecerts *bool   `tfsdk:"validatecerts" json:"validatecerts,omitempty"`
			Vendor        *string `tfsdk:"vendor" json:"vendor,omitempty"`
		} `tfsdk:"provider" json:"provider,omitempty"`
		Type *string `tfsdk:"type" json:"type,omitempty"`
		Vip  *string `tfsdk:"vip" json:"vip,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *LbLbconfigCarlosedpComExternalLoadBalancerV1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_lb_lbconfig_carlosedp_com_external_load_balancer_v1"
}

func (r *LbLbconfigCarlosedpComExternalLoadBalancerV1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ExternalLoadBalancer is the Schema for the externalloadbalancers API",
		MarkdownDescription: "ExternalLoadBalancer is the Schema for the externalloadbalancers API",
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
				Description:         "ExternalLoadBalancerSpec is the spec of a LoadBalancer instance.",
				MarkdownDescription: "ExternalLoadBalancerSpec is the spec of a LoadBalancer instance.",
				Attributes: map[string]schema.Attribute{
					"monitor": schema.SingleNestedAttribute{
						Description:         "Monitor is the path and port to monitor the LoadBalancer members",
						MarkdownDescription: "Monitor is the path and port to monitor the LoadBalancer members",
						Attributes: map[string]schema.Attribute{
							"monitortype": schema.StringAttribute{
								Description:         "MonitorType is the monitor parent type. <monitorType> must be one of 'http', 'https', 'icmp'.",
								MarkdownDescription: "MonitorType is the monitor parent type. <monitorType> must be one of 'http', 'https', 'icmp'.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("http", "https", "icmp"),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name is the monitor name, it is set by the controller",
								MarkdownDescription: "Name is the monitor name, it is set by the controller",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"path": schema.StringAttribute{
								Description:         "Path is the path URL to check for the pool members in the format '/healthz'",
								MarkdownDescription: "Path is the path URL to check for the pool members in the format '/healthz'",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"port": schema.Int64Attribute{
								Description:         "Port is the port this monitor should check the pool members",
								MarkdownDescription: "Port is the port this monitor should check the pool members",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"nodelabels": schema.MapAttribute{
						Description:         "NodeLabels are the node labels used for router sharding as an alternative to 'type'. Optional.",
						MarkdownDescription: "NodeLabels are the node labels used for router sharding as an alternative to 'type'. Optional.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ports": schema.ListAttribute{
						Description:         "Ports is the ports exposed by this LoadBalancer instance",
						MarkdownDescription: "Ports is the ports exposed by this LoadBalancer instance",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"provider": schema.SingleNestedAttribute{
						Description:         "Provider is the LoadBalancer backend provider",
						MarkdownDescription: "Provider is the LoadBalancer backend provider",
						Attributes: map[string]schema.Attribute{
							"creds": schema.StringAttribute{
								Description:         "Creds is the credentials secret holding the 'username' and 'password' keys. Generate with: 'kubectl create secret generic <secret-name> --from-literal=username=<username> --from-literal=password=<password>'",
								MarkdownDescription: "Creds is the credentials secret holding the 'username' and 'password' keys. Generate with: 'kubectl create secret generic <secret-name> --from-literal=username=<username> --from-literal=password=<password>'",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"debug": schema.BoolAttribute{
								Description:         "Debug is a flag to enable debug on the backend log output. Defaults to false.",
								MarkdownDescription: "Debug is a flag to enable debug on the backend log output. Defaults to false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host": schema.StringAttribute{
								Description:         "Host is the Load Balancer API IP or Hostname in URL format. Eg. 'http://10.25.10.10'.",
								MarkdownDescription: "Host is the Load Balancer API IP or Hostname in URL format. Eg. 'http://10.25.10.10'.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(255),
								},
							},

							"lbmethod": schema.StringAttribute{
								Description:         "Type is the Load-Balancing method. Defaults to 'round-robin'. Options are: ROUNDROBIN, LEASTCONNECTION, LEASTRESPONSETIME",
								MarkdownDescription: "Type is the Load-Balancing method. Defaults to 'round-robin'. Options are: ROUNDROBIN, LEASTCONNECTION, LEASTRESPONSETIME",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("ROUNDROBIN", "LEASTCONNECTION", "LEASTRESPONSETIME"),
								},
							},

							"partition": schema.StringAttribute{
								Description:         "Partition is the F5 partition to create the Load Balancer instances. Defaults to 'Common'. (F5 BigIP only)",
								MarkdownDescription: "Partition is the F5 partition to create the Load Balancer instances. Defaults to 'Common'. (F5 BigIP only)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "Port is the Load Balancer API Port.",
								MarkdownDescription: "Port is the Load Balancer API Port.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(65535),
								},
							},

							"validatecerts": schema.BoolAttribute{
								Description:         "ValidateCerts is a flag to validate or not the Load Balancer API certificate. Defaults to false.",
								MarkdownDescription: "ValidateCerts is a flag to validate or not the Load Balancer API certificate. Defaults to false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vendor": schema.StringAttribute{
								Description:         "Vendor is the backend provider vendor",
								MarkdownDescription: "Vendor is the backend provider vendor",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Dummy", "F5_BigIP", "Citrix_ADC", "HAProxy"),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"type": schema.StringAttribute{
						Description:         "Type is the node role type (master or infra) for the LoadBalancer instance",
						MarkdownDescription: "Type is the node role type (master or infra) for the LoadBalancer instance",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("master", "infra"),
						},
					},

					"vip": schema.StringAttribute{
						Description:         "Vip is the Virtual IP configured in  this LoadBalancer instance",
						MarkdownDescription: "Vip is the Virtual IP configured in  this LoadBalancer instance",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.LengthAtMost(15),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *LbLbconfigCarlosedpComExternalLoadBalancerV1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *LbLbconfigCarlosedpComExternalLoadBalancerV1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_lb_lbconfig_carlosedp_com_external_load_balancer_v1")

	var model LbLbconfigCarlosedpComExternalLoadBalancerV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("lb.lbconfig.carlosedp.com/v1")
	model.Kind = pointer.String("ExternalLoadBalancer")

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

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "lb.lbconfig.carlosedp.com", Version: "v1", Resource: "ExternalLoadBalancer"}).
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

	var readResponse LbLbconfigCarlosedpComExternalLoadBalancerV1ResourceData
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

func (r *LbLbconfigCarlosedpComExternalLoadBalancerV1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_lb_lbconfig_carlosedp_com_external_load_balancer_v1")

	var data LbLbconfigCarlosedpComExternalLoadBalancerV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "lb.lbconfig.carlosedp.com", Version: "v1", Resource: "ExternalLoadBalancer"}).
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

	var readResponse LbLbconfigCarlosedpComExternalLoadBalancerV1ResourceData
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

func (r *LbLbconfigCarlosedpComExternalLoadBalancerV1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_lb_lbconfig_carlosedp_com_external_load_balancer_v1")

	var model LbLbconfigCarlosedpComExternalLoadBalancerV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("lb.lbconfig.carlosedp.com/v1")
	model.Kind = pointer.String("ExternalLoadBalancer")

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

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "lb.lbconfig.carlosedp.com", Version: "v1", Resource: "ExternalLoadBalancer"}).
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

	var readResponse LbLbconfigCarlosedpComExternalLoadBalancerV1ResourceData
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

func (r *LbLbconfigCarlosedpComExternalLoadBalancerV1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_lb_lbconfig_carlosedp_com_external_load_balancer_v1")

	var data LbLbconfigCarlosedpComExternalLoadBalancerV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "lb.lbconfig.carlosedp.com", Version: "v1", Resource: "ExternalLoadBalancer"}).
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

func (r *LbLbconfigCarlosedpComExternalLoadBalancerV1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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