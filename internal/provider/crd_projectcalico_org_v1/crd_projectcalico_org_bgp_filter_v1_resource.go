/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package crd_projectcalico_org_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
	"time"
)

var (
	_ resource.Resource                = &CrdProjectcalicoOrgBgpfilterV1Resource{}
	_ resource.ResourceWithConfigure   = &CrdProjectcalicoOrgBgpfilterV1Resource{}
	_ resource.ResourceWithImportState = &CrdProjectcalicoOrgBgpfilterV1Resource{}
)

func NewCrdProjectcalicoOrgBgpfilterV1Resource() resource.Resource {
	return &CrdProjectcalicoOrgBgpfilterV1Resource{}
}

type CrdProjectcalicoOrgBgpfilterV1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type CrdProjectcalicoOrgBgpfilterV1ResourceData struct {
	ID                  types.String `tfsdk:"id" json:"-"`
	ForceConflicts      types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager        types.String `tfsdk:"field_manager" json:"-"`
	DeletionPropagation types.String `tfsdk:"deletion_propagation" json:"-"`
	WaitForUpsert       types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete       types.Object `tfsdk:"wait_for_delete" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ExportV4 *[]struct {
			Action        *string `tfsdk:"action" json:"action,omitempty"`
			Cidr          *string `tfsdk:"cidr" json:"cidr,omitempty"`
			MatchOperator *string `tfsdk:"match_operator" json:"matchOperator,omitempty"`
		} `tfsdk:"export_v4" json:"exportV4,omitempty"`
		ExportV6 *[]struct {
			Action        *string `tfsdk:"action" json:"action,omitempty"`
			Cidr          *string `tfsdk:"cidr" json:"cidr,omitempty"`
			MatchOperator *string `tfsdk:"match_operator" json:"matchOperator,omitempty"`
		} `tfsdk:"export_v6" json:"exportV6,omitempty"`
		ImportV4 *[]struct {
			Action        *string `tfsdk:"action" json:"action,omitempty"`
			Cidr          *string `tfsdk:"cidr" json:"cidr,omitempty"`
			MatchOperator *string `tfsdk:"match_operator" json:"matchOperator,omitempty"`
		} `tfsdk:"import_v4" json:"importV4,omitempty"`
		ImportV6 *[]struct {
			Action        *string `tfsdk:"action" json:"action,omitempty"`
			Cidr          *string `tfsdk:"cidr" json:"cidr,omitempty"`
			MatchOperator *string `tfsdk:"match_operator" json:"matchOperator,omitempty"`
		} `tfsdk:"import_v6" json:"importV6,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CrdProjectcalicoOrgBgpfilterV1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_crd_projectcalico_org_bgp_filter_v1"
}

func (r *CrdProjectcalicoOrgBgpfilterV1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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

			"field_manager": schema.StringAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"deletion_propagation": schema.StringAttribute{
				Description:         "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				MarkdownDescription: "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("Orphan", "Background", "Foreground"),
				},
			},

			"wait_for_upsert": schema.ListNestedAttribute{
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
						"timeout": schema.Int64Attribute{
							Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(30),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"poll_interval": schema.Int64Attribute{
							Description:         "The number of seconds to wait before checking again.",
							MarkdownDescription: "The number of seconds to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(5),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
					},
				},
			},

			"wait_for_delete": schema.SingleNestedAttribute{
				Description:         "Wait for deletion of resources.",
				MarkdownDescription: "Wait for deletion of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"timeout": schema.Int64Attribute{
						Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(30),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"poll_interval": schema.Int64Attribute{
						Description:         "The number of seconds to wait before checking again.",
						MarkdownDescription: "The number of seconds to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(5),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
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
				Description:         "BGPFilterSpec contains the IPv4 and IPv6 filter rules of the BGP Filter.",
				MarkdownDescription: "BGPFilterSpec contains the IPv4 and IPv6 filter rules of the BGP Filter.",
				Attributes: map[string]schema.Attribute{
					"export_v4": schema.ListNestedAttribute{
						Description:         "The ordered set of IPv4 BGPFilter rules acting on exporting routes to a peer.",
						MarkdownDescription: "The ordered set of IPv4 BGPFilter rules acting on exporting routes to a peer.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"cidr": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"match_operator": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"export_v6": schema.ListNestedAttribute{
						Description:         "The ordered set of IPv6 BGPFilter rules acting on exporting routes to a peer.",
						MarkdownDescription: "The ordered set of IPv6 BGPFilter rules acting on exporting routes to a peer.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"cidr": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"match_operator": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"import_v4": schema.ListNestedAttribute{
						Description:         "The ordered set of IPv4 BGPFilter rules acting on importing routes from a peer.",
						MarkdownDescription: "The ordered set of IPv4 BGPFilter rules acting on importing routes from a peer.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"cidr": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"match_operator": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"import_v6": schema.ListNestedAttribute{
						Description:         "The ordered set of IPv6 BGPFilter rules acting on importing routes from a peer.",
						MarkdownDescription: "The ordered set of IPv6 BGPFilter rules acting on importing routes from a peer.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"cidr": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"match_operator": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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
		},
	}
}

func (r *CrdProjectcalicoOrgBgpfilterV1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedResourceDataError(request.ProviderData))
	}
}

func (r *CrdProjectcalicoOrgBgpfilterV1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_crd_projectcalico_org_bgp_filter_v1")

	var model CrdProjectcalicoOrgBgpfilterV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("crd.projectcalico.org/v1")
	model.Kind = pointer.String("BGPFilter")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
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
		Resource(k8sSchema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "v1", Resource: "bgpfilters"}).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse CrdProjectcalicoOrgBgpfilterV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec
	if model.ForceConflicts.IsUnknown() {
		model.ForceConflicts = types.BoolNull()
	}
	if model.FieldManager.IsUnknown() {
		model.FieldManager = types.StringNull()
	}
	if model.DeletionPropagation.IsUnknown() {
		model.DeletionPropagation = types.StringNull()
	}
	if model.WaitForUpsert.IsUnknown() {
		model.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if model.WaitForDelete.IsUnknown() {
		model.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CrdProjectcalicoOrgBgpfilterV1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_crd_projectcalico_org_bgp_filter_v1")

	var data CrdProjectcalicoOrgBgpfilterV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "v1", Resource: "bgpfilters"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetResourceError(err, data.Metadata.Name))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse CrdProjectcalicoOrgBgpfilterV1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec
	if data.ForceConflicts.IsUnknown() {
		data.ForceConflicts = types.BoolNull()
	}
	if data.FieldManager.IsUnknown() {
		data.FieldManager = types.StringNull()
	}
	if data.DeletionPropagation.IsUnknown() {
		data.DeletionPropagation = types.StringNull()
	}
	if data.WaitForUpsert.IsUnknown() {
		data.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if data.WaitForDelete.IsUnknown() {
		data.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *CrdProjectcalicoOrgBgpfilterV1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_crd_projectcalico_org_bgp_filter_v1")

	var model CrdProjectcalicoOrgBgpfilterV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("crd.projectcalico.org/v1")
	model.Kind = pointer.String("BGPFilter")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
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
		Resource(k8sSchema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "v1", Resource: "bgpfilters"}).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse CrdProjectcalicoOrgBgpfilterV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CrdProjectcalicoOrgBgpfilterV1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_crd_projectcalico_org_bgp_filter_v1")

	var data CrdProjectcalicoOrgBgpfilterV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	deleteOptions := meta.DeleteOptions{}
	if !data.DeletionPropagation.IsNull() && !data.DeletionPropagation.IsUnknown() {
		deleteOptions.PropagationPolicy = utilities.MapDeletionPropagation(data.DeletionPropagation.ValueString())
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "v1", Resource: "bgpfilters"}).
		Delete(ctx, data.Metadata.Name, deleteOptions)
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() && !data.WaitForDelete.IsUnknown() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "v1", Resource: "bgpfilters"}).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout.Milliseconds() == 0 {
				break
			}
			if time.Now().After(startTime.Add(timeout)) {
				response.Diagnostics.Append(utilities.WaitTimeoutExceeded())
				return
			}
			time.Sleep(pollInterval)
		}
	}
}

func (r *CrdProjectcalicoOrgBgpfilterV1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	if request.ID == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'name' Got: '%q'", request.ID),
		)
		return
	}
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	resource.ImportStatePassthroughID(ctx, path.Root("metadata").AtName("name"), request, response)
}
