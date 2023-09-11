/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package metal3_io_v1alpha1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &Metal3IoFirmwareSchemaV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &Metal3IoFirmwareSchemaV1Alpha1DataSource{}
)

func NewMetal3IoFirmwareSchemaV1Alpha1DataSource() datasource.DataSource {
	return &Metal3IoFirmwareSchemaV1Alpha1DataSource{}
}

type Metal3IoFirmwareSchemaV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type Metal3IoFirmwareSchemaV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		HardwareModel  *string `tfsdk:"hardware_model" json:"hardwareModel,omitempty"`
		HardwareVendor *string `tfsdk:"hardware_vendor" json:"hardwareVendor,omitempty"`
		Schema         *struct {
			Allowable_values *[]string `tfsdk:"allowable_values" json:"allowable_values,omitempty"`
			Attribute_type   *string   `tfsdk:"attribute_type" json:"attribute_type,omitempty"`
			Lower_bound      *int64    `tfsdk:"lower_bound" json:"lower_bound,omitempty"`
			Max_length       *int64    `tfsdk:"max_length" json:"max_length,omitempty"`
			Min_length       *int64    `tfsdk:"min_length" json:"min_length,omitempty"`
			Read_only        *bool     `tfsdk:"read_only" json:"read_only,omitempty"`
			Unique           *bool     `tfsdk:"unique" json:"unique,omitempty"`
			Upper_bound      *int64    `tfsdk:"upper_bound" json:"upper_bound,omitempty"`
		} `tfsdk:"schema" json:"schema,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Metal3IoFirmwareSchemaV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_metal3_io_firmware_schema_v1alpha1"
}

func (r *Metal3IoFirmwareSchemaV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "FirmwareSchema is the Schema for the firmwareschemas API",
		MarkdownDescription: "FirmwareSchema is the Schema for the firmwareschemas API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "FirmwareSchemaSpec defines the desired state of FirmwareSchema",
				MarkdownDescription: "FirmwareSchemaSpec defines the desired state of FirmwareSchema",
				Attributes: map[string]schema.Attribute{
					"hardware_model": schema.StringAttribute{
						Description:         "The hardware model associated with this schema",
						MarkdownDescription: "The hardware model associated with this schema",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"hardware_vendor": schema.StringAttribute{
						Description:         "The hardware vendor associated with this schema",
						MarkdownDescription: "The hardware vendor associated with this schema",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"schema": schema.SingleNestedAttribute{
						Description:         "Map of firmware name to schema",
						MarkdownDescription: "Map of firmware name to schema",
						Attributes: map[string]schema.Attribute{
							"allowable_values": schema.ListAttribute{
								Description:         "The allowable value for an Enumeration type setting.",
								MarkdownDescription: "The allowable value for an Enumeration type setting.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"attribute_type": schema.StringAttribute{
								Description:         "The type of setting.",
								MarkdownDescription: "The type of setting.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"lower_bound": schema.Int64Attribute{
								Description:         "The lowest value for an Integer type setting.",
								MarkdownDescription: "The lowest value for an Integer type setting.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"max_length": schema.Int64Attribute{
								Description:         "Maximum length for a String type setting.",
								MarkdownDescription: "Maximum length for a String type setting.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"min_length": schema.Int64Attribute{
								Description:         "Minimum length for a String type setting.",
								MarkdownDescription: "Minimum length for a String type setting.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"read_only": schema.BoolAttribute{
								Description:         "Whether or not this setting is read only.",
								MarkdownDescription: "Whether or not this setting is read only.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"unique": schema.BoolAttribute{
								Description:         "Whether or not this setting's value is unique to this node, e.g. a serial number.",
								MarkdownDescription: "Whether or not this setting's value is unique to this node, e.g. a serial number.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"upper_bound": schema.Int64Attribute{
								Description:         "The highest value for an Integer type setting.",
								MarkdownDescription: "The highest value for an Integer type setting.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *Metal3IoFirmwareSchemaV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *Metal3IoFirmwareSchemaV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_metal3_io_firmware_schema_v1alpha1")

	var data Metal3IoFirmwareSchemaV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "metal3.io", Version: "v1alpha1", Resource: "firmwareschemas"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
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

	var readResponse Metal3IoFirmwareSchemaV1Alpha1DataSourceData
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

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("metal3.io/v1alpha1")
	data.Kind = pointer.String("FirmwareSchema")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
