/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"
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

type NotificationToolkitFluxcdIoReceiverV1Beta1Resource struct{}

var (
	_ resource.Resource = (*NotificationToolkitFluxcdIoReceiverV1Beta1Resource)(nil)
)

type NotificationToolkitFluxcdIoReceiverV1Beta1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type NotificationToolkitFluxcdIoReceiverV1Beta1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Suspend *bool `tfsdk:"suspend" yaml:"suspend,omitempty"`

		Type *string `tfsdk:"type" yaml:"type,omitempty"`

		Events *[]string `tfsdk:"events" yaml:"events,omitempty"`

		Resources *[]struct {
			ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
		} `tfsdk:"resources" yaml:"resources,omitempty"`

		SecretRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewNotificationToolkitFluxcdIoReceiverV1Beta1Resource() resource.Resource {
	return &NotificationToolkitFluxcdIoReceiverV1Beta1Resource{}
}

func (r *NotificationToolkitFluxcdIoReceiverV1Beta1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_notification_toolkit_fluxcd_io_receiver_v1beta1"
}

func (r *NotificationToolkitFluxcdIoReceiverV1Beta1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Receiver is the Schema for the receivers API",
		MarkdownDescription: "Receiver is the Schema for the receivers API",
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
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
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
				Description:         "ReceiverSpec defines the desired state of Receiver",
				MarkdownDescription: "ReceiverSpec defines the desired state of Receiver",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"suspend": {
						Description:         "This flag tells the controller to suspend subsequent events handling. Defaults to false.",
						MarkdownDescription: "This flag tells the controller to suspend subsequent events handling. Defaults to false.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"type": {
						Description:         "Type of webhook sender, used to determine the validation procedure and payload deserialization.",
						MarkdownDescription: "Type of webhook sender, used to determine the validation procedure and payload deserialization.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"events": {
						Description:         "A list of events to handle, e.g. 'push' for GitHub or 'Push Hook' for GitLab.",
						MarkdownDescription: "A list of events to handle, e.g. 'push' for GitHub or 'Push Hook' for GitLab.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resources": {
						Description:         "A list of resources to be notified about changes.",
						MarkdownDescription: "A list of resources to be notified about changes.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"api_version": {
								Description:         "API version of the referent",
								MarkdownDescription: "API version of the referent",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kind": {
								Description:         "Kind of the referent",
								MarkdownDescription: "Kind of the referent",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"match_labels": {
								Description:         "MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name of the referent",
								MarkdownDescription: "Name of the referent",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace of the referent",
								MarkdownDescription: "Namespace of the referent",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"secret_ref": {
						Description:         "Secret reference containing the token used to validate the payload authenticity",
						MarkdownDescription: "Secret reference containing the token used to validate the payload authenticity",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",

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
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *NotificationToolkitFluxcdIoReceiverV1Beta1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_notification_toolkit_fluxcd_io_receiver_v1beta1")

	var state NotificationToolkitFluxcdIoReceiverV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel NotificationToolkitFluxcdIoReceiverV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("notification.toolkit.fluxcd.io/v1beta1")
	goModel.Kind = utilities.Ptr("Receiver")

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

func (r *NotificationToolkitFluxcdIoReceiverV1Beta1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_notification_toolkit_fluxcd_io_receiver_v1beta1")
	// NO-OP: All data is already in Terraform state
}

func (r *NotificationToolkitFluxcdIoReceiverV1Beta1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_notification_toolkit_fluxcd_io_receiver_v1beta1")

	var state NotificationToolkitFluxcdIoReceiverV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel NotificationToolkitFluxcdIoReceiverV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("notification.toolkit.fluxcd.io/v1beta1")
	goModel.Kind = utilities.Ptr("Receiver")

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

func (r *NotificationToolkitFluxcdIoReceiverV1Beta1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_notification_toolkit_fluxcd_io_receiver_v1beta1")
	// NO-OP: Terraform removes the state automatically for us
}
