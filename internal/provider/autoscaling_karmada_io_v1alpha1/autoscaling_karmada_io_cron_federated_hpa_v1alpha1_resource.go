/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package autoscaling_karmada_io_v1alpha1

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
	_ resource.Resource                = &AutoscalingKarmadaIoCronFederatedHpaV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &AutoscalingKarmadaIoCronFederatedHpaV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &AutoscalingKarmadaIoCronFederatedHpaV1Alpha1Resource{}
)

func NewAutoscalingKarmadaIoCronFederatedHpaV1Alpha1Resource() resource.Resource {
	return &AutoscalingKarmadaIoCronFederatedHpaV1Alpha1Resource{}
}

type AutoscalingKarmadaIoCronFederatedHpaV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type AutoscalingKarmadaIoCronFederatedHpaV1Alpha1ResourceData struct {
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
		Rules *[]struct {
			FailedHistoryLimit     *int64  `tfsdk:"failed_history_limit" json:"failedHistoryLimit,omitempty"`
			Name                   *string `tfsdk:"name" json:"name,omitempty"`
			Schedule               *string `tfsdk:"schedule" json:"schedule,omitempty"`
			SuccessfulHistoryLimit *int64  `tfsdk:"successful_history_limit" json:"successfulHistoryLimit,omitempty"`
			Suspend                *bool   `tfsdk:"suspend" json:"suspend,omitempty"`
			TargetMaxReplicas      *int64  `tfsdk:"target_max_replicas" json:"targetMaxReplicas,omitempty"`
			TargetMinReplicas      *int64  `tfsdk:"target_min_replicas" json:"targetMinReplicas,omitempty"`
			TargetReplicas         *int64  `tfsdk:"target_replicas" json:"targetReplicas,omitempty"`
			TimeZone               *string `tfsdk:"time_zone" json:"timeZone,omitempty"`
		} `tfsdk:"rules" json:"rules,omitempty"`
		ScaleTargetRef *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"scale_target_ref" json:"scaleTargetRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AutoscalingKarmadaIoCronFederatedHpaV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_autoscaling_karmada_io_cron_federated_hpa_v1alpha1"
}

func (r *AutoscalingKarmadaIoCronFederatedHpaV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CronFederatedHPA represents a collection of repeating schedule to scale replica number of a specific workload. It can scale any resource implementing the scale subresource as well as FederatedHPA.",
		MarkdownDescription: "CronFederatedHPA represents a collection of repeating schedule to scale replica number of a specific workload. It can scale any resource implementing the scale subresource as well as FederatedHPA.",
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
				Description:         "Spec is the specification of the CronFederatedHPA.",
				MarkdownDescription: "Spec is the specification of the CronFederatedHPA.",
				Attributes: map[string]schema.Attribute{
					"rules": schema.ListNestedAttribute{
						Description:         "Rules contains a collection of schedules that declares when and how the referencing target resource should be scaled.",
						MarkdownDescription: "Rules contains a collection of schedules that declares when and how the referencing target resource should be scaled.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"failed_history_limit": schema.Int64Attribute{
									Description:         "FailedHistoryLimit represents the count of failed execution items for each rule. The value must be a positive integer. It defaults to 3.",
									MarkdownDescription: "FailedHistoryLimit represents the count of failed execution items for each rule. The value must be a positive integer. It defaults to 3.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
										int64validator.AtMost(32),
									},
								},

								"name": schema.StringAttribute{
									Description:         "Name of the rule. Each rule in a CronFederatedHPA must have a unique name.  Note: the name will be used as an identifier to record its execution history. Changing the name will be considered as deleting the old rule and adding a new rule, that means the original execution history will be discarded.",
									MarkdownDescription: "Name of the rule. Each rule in a CronFederatedHPA must have a unique name.  Note: the name will be used as an identifier to record its execution history. Changing the name will be considered as deleting the old rule and adding a new rule, that means the original execution history will be discarded.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(32),
									},
								},

								"schedule": schema.StringAttribute{
									Description:         "Schedule is the cron expression that represents a periodical time. The syntax follows https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#schedule-syntax.",
									MarkdownDescription: "Schedule is the cron expression that represents a periodical time. The syntax follows https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#schedule-syntax.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"successful_history_limit": schema.Int64Attribute{
									Description:         "SuccessfulHistoryLimit represents the count of successful execution items for each rule. The value must be a positive integer. It defaults to 3.",
									MarkdownDescription: "SuccessfulHistoryLimit represents the count of successful execution items for each rule. The value must be a positive integer. It defaults to 3.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(1),
										int64validator.AtMost(32),
									},
								},

								"suspend": schema.BoolAttribute{
									Description:         "Suspend tells the controller to suspend subsequent executions. Defaults to false.",
									MarkdownDescription: "Suspend tells the controller to suspend subsequent executions. Defaults to false.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target_max_replicas": schema.Int64Attribute{
									Description:         "TargetMaxReplicas is the target MaxReplicas to be set for FederatedHPA. Only needed when referencing resource is FederatedHPA. TargetMinReplicas and TargetMaxReplicas can be specified together or either one can be specified alone. nil means the MaxReplicas(.spec.maxReplicas) of the referencing FederatedHPA will not be updated.",
									MarkdownDescription: "TargetMaxReplicas is the target MaxReplicas to be set for FederatedHPA. Only needed when referencing resource is FederatedHPA. TargetMinReplicas and TargetMaxReplicas can be specified together or either one can be specified alone. nil means the MaxReplicas(.spec.maxReplicas) of the referencing FederatedHPA will not be updated.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target_min_replicas": schema.Int64Attribute{
									Description:         "TargetMinReplicas is the target MinReplicas to be set for FederatedHPA. Only needed when referencing resource is FederatedHPA. TargetMinReplicas and TargetMaxReplicas can be specified together or either one can be specified alone. nil means the MinReplicas(.spec.minReplicas) of the referencing FederatedHPA will not be updated.",
									MarkdownDescription: "TargetMinReplicas is the target MinReplicas to be set for FederatedHPA. Only needed when referencing resource is FederatedHPA. TargetMinReplicas and TargetMaxReplicas can be specified together or either one can be specified alone. nil means the MinReplicas(.spec.minReplicas) of the referencing FederatedHPA will not be updated.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target_replicas": schema.Int64Attribute{
									Description:         "TargetReplicas is the target replicas to be scaled for resources referencing by ScaleTargetRef of this CronFederatedHPA. Only needed when referencing resource is not FederatedHPA.",
									MarkdownDescription: "TargetReplicas is the target replicas to be scaled for resources referencing by ScaleTargetRef of this CronFederatedHPA. Only needed when referencing resource is not FederatedHPA.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"time_zone": schema.StringAttribute{
									Description:         "TimeZone for the giving schedule. If not specified, this will default to the time zone of the karmada-controller-manager process. Invalid TimeZone will be rejected when applying by karmada-webhook. see https://en.wikipedia.org/wiki/List_of_tz_database_time_zones for the all timezones.",
									MarkdownDescription: "TimeZone for the giving schedule. If not specified, this will default to the time zone of the karmada-controller-manager process. Invalid TimeZone will be rejected when applying by karmada-webhook. see https://en.wikipedia.org/wiki/List_of_tz_database_time_zones for the all timezones.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"scale_target_ref": schema.SingleNestedAttribute{
						Description:         "ScaleTargetRef points to the target resource to scale. Target resource could be any resource that implementing the scale subresource like Deployment, or FederatedHPA.",
						MarkdownDescription: "ScaleTargetRef points to the target resource to scale. Target resource could be any resource that implementing the scale subresource like Deployment, or FederatedHPA.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "apiVersion is the API version of the referent",
								MarkdownDescription: "apiVersion is the API version of the referent",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
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

func (r *AutoscalingKarmadaIoCronFederatedHpaV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *AutoscalingKarmadaIoCronFederatedHpaV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_autoscaling_karmada_io_cron_federated_hpa_v1alpha1")

	var model AutoscalingKarmadaIoCronFederatedHpaV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("autoscaling.karmada.io/v1alpha1")
	model.Kind = pointer.String("CronFederatedHPA")

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

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "autoscaling.karmada.io", Version: "v1alpha1", Resource: "CronFederatedHPA"}).
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

	var readResponse AutoscalingKarmadaIoCronFederatedHpaV1Alpha1ResourceData
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

func (r *AutoscalingKarmadaIoCronFederatedHpaV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_autoscaling_karmada_io_cron_federated_hpa_v1alpha1")

	var data AutoscalingKarmadaIoCronFederatedHpaV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "autoscaling.karmada.io", Version: "v1alpha1", Resource: "CronFederatedHPA"}).
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

	var readResponse AutoscalingKarmadaIoCronFederatedHpaV1Alpha1ResourceData
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

func (r *AutoscalingKarmadaIoCronFederatedHpaV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_autoscaling_karmada_io_cron_federated_hpa_v1alpha1")

	var model AutoscalingKarmadaIoCronFederatedHpaV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("autoscaling.karmada.io/v1alpha1")
	model.Kind = pointer.String("CronFederatedHPA")

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

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "autoscaling.karmada.io", Version: "v1alpha1", Resource: "CronFederatedHPA"}).
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

	var readResponse AutoscalingKarmadaIoCronFederatedHpaV1Alpha1ResourceData
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

func (r *AutoscalingKarmadaIoCronFederatedHpaV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_autoscaling_karmada_io_cron_federated_hpa_v1alpha1")

	var data AutoscalingKarmadaIoCronFederatedHpaV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "autoscaling.karmada.io", Version: "v1alpha1", Resource: "CronFederatedHPA"}).
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

func (r *AutoscalingKarmadaIoCronFederatedHpaV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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
