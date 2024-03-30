/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package oracle_db_anthosapis_com_v1alpha1

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
	_ datasource.DataSource = &OracleDbAnthosapisComExportV1Alpha1Manifest{}
)

func NewOracleDbAnthosapisComExportV1Alpha1Manifest() datasource.DataSource {
	return &OracleDbAnthosapisComExportV1Alpha1Manifest{}
}

type OracleDbAnthosapisComExportV1Alpha1Manifest struct{}

type OracleDbAnthosapisComExportV1Alpha1ManifestData struct {
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
		DatabaseName     *string   `tfsdk:"database_name" json:"databaseName,omitempty"`
		ExportObjectType *string   `tfsdk:"export_object_type" json:"exportObjectType,omitempty"`
		ExportObjects    *[]string `tfsdk:"export_objects" json:"exportObjects,omitempty"`
		FlashbackTime    *string   `tfsdk:"flashback_time" json:"flashbackTime,omitempty"`
		GcsLogPath       *string   `tfsdk:"gcs_log_path" json:"gcsLogPath,omitempty"`
		GcsPath          *string   `tfsdk:"gcs_path" json:"gcsPath,omitempty"`
		Instance         *string   `tfsdk:"instance" json:"instance,omitempty"`
		Type             *string   `tfsdk:"type" json:"type,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OracleDbAnthosapisComExportV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_oracle_db_anthosapis_com_export_v1alpha1_manifest"
}

func (r *OracleDbAnthosapisComExportV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Export is the Schema for the exports API.",
		MarkdownDescription: "Export is the Schema for the exports API.",
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
				Description:         "ExportSpec defines the desired state of Export",
				MarkdownDescription: "ExportSpec defines the desired state of Export",
				Attributes: map[string]schema.Attribute{
					"database_name": schema.StringAttribute{
						Description:         "DatabaseName is the database resource name within Instance to export from.",
						MarkdownDescription: "DatabaseName is the database resource name within Instance to export from.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"export_object_type": schema.StringAttribute{
						Description:         "ExportObjectType is the type of objects to export. If omitted, the default of Schemas is assumed. Supported options at this point are: Schemas or Tables.",
						MarkdownDescription: "ExportObjectType is the type of objects to export. If omitted, the default of Schemas is assumed. Supported options at this point are: Schemas or Tables.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Schemas", "Tables"),
						},
					},

					"export_objects": schema.ListAttribute{
						Description:         "ExportObjects are objects, schemas or tables, exported by DataPump.",
						MarkdownDescription: "ExportObjects are objects, schemas or tables, exported by DataPump.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"flashback_time": schema.StringAttribute{
						Description:         "FlashbackTime is an optional time. If this time is set, the SCN that most closely matches the time is found, and this SCN is used to enable the Flashback utility. The export operation is performed with data that is consistent up to this SCN.",
						MarkdownDescription: "FlashbackTime is an optional time. If this time is set, the SCN that most closely matches the time is found, and this SCN is used to enable the Flashback utility. The export operation is performed with data that is consistent up to this SCN.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							validators.DateTime64Validator(),
						},
					},

					"gcs_log_path": schema.StringAttribute{
						Description:         "GcsLogPath is an optional full path in GCS. If set up ahead of time, export logs can be optionally transferred to set GCS bucket. A user is to ensure proper write access to the bucket from within the Oracle Operator.",
						MarkdownDescription: "GcsLogPath is an optional full path in GCS. If set up ahead of time, export logs can be optionally transferred to set GCS bucket. A user is to ensure proper write access to the bucket from within the Oracle Operator.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"gcs_path": schema.StringAttribute{
						Description:         "GcsPath is a full path in GCS bucket to transfer exported files to. A user is to ensure proper write access to the bucket from within the Oracle Operator.",
						MarkdownDescription: "GcsPath is a full path in GCS bucket to transfer exported files to. A user is to ensure proper write access to the bucket from within the Oracle Operator.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"instance": schema.StringAttribute{
						Description:         "Instance is the resource name within namespace to export from.",
						MarkdownDescription: "Instance is the resource name within namespace to export from.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"type": schema.StringAttribute{
						Description:         "Type of the Export. If omitted, the default of DataPump is assumed.",
						MarkdownDescription: "Type of the Export. If omitted, the default of DataPump is assumed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("DataPump"),
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

func (r *OracleDbAnthosapisComExportV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_oracle_db_anthosapis_com_export_v1alpha1_manifest")

	var model OracleDbAnthosapisComExportV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("oracle.db.anthosapis.com/v1alpha1")
	model.Kind = pointer.String("Export")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
