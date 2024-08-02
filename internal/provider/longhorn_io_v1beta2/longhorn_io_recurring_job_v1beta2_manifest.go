/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package longhorn_io_v1beta2

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
	_ datasource.DataSource = &LonghornIoRecurringJobV1Beta2Manifest{}
)

func NewLonghornIoRecurringJobV1Beta2Manifest() datasource.DataSource {
	return &LonghornIoRecurringJobV1Beta2Manifest{}
}

type LonghornIoRecurringJobV1Beta2Manifest struct{}

type LonghornIoRecurringJobV1Beta2ManifestData struct {
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
		Concurrency *int64             `tfsdk:"concurrency" json:"concurrency,omitempty"`
		Cron        *string            `tfsdk:"cron" json:"cron,omitempty"`
		Groups      *[]string          `tfsdk:"groups" json:"groups,omitempty"`
		Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Name        *string            `tfsdk:"name" json:"name,omitempty"`
		Parameters  *map[string]string `tfsdk:"parameters" json:"parameters,omitempty"`
		Retain      *int64             `tfsdk:"retain" json:"retain,omitempty"`
		Task        *string            `tfsdk:"task" json:"task,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *LonghornIoRecurringJobV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_longhorn_io_recurring_job_v1beta2_manifest"
}

func (r *LonghornIoRecurringJobV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "RecurringJob is where Longhorn stores recurring job object.",
		MarkdownDescription: "RecurringJob is where Longhorn stores recurring job object.",
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
				Description:         "RecurringJobSpec defines the desired state of the Longhorn recurring job",
				MarkdownDescription: "RecurringJobSpec defines the desired state of the Longhorn recurring job",
				Attributes: map[string]schema.Attribute{
					"concurrency": schema.Int64Attribute{
						Description:         "The concurrency of taking the snapshot/backup.",
						MarkdownDescription: "The concurrency of taking the snapshot/backup.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cron": schema.StringAttribute{
						Description:         "The cron setting.",
						MarkdownDescription: "The cron setting.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"groups": schema.ListAttribute{
						Description:         "The recurring job group.",
						MarkdownDescription: "The recurring job group.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"labels": schema.MapAttribute{
						Description:         "The label of the snapshot/backup.",
						MarkdownDescription: "The label of the snapshot/backup.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "The recurring job name.",
						MarkdownDescription: "The recurring job name.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"parameters": schema.MapAttribute{
						Description:         "The parameters of the snapshot/backup.Support parameters: 'full-backup-interval'.",
						MarkdownDescription: "The parameters of the snapshot/backup.Support parameters: 'full-backup-interval'.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"retain": schema.Int64Attribute{
						Description:         "The retain count of the snapshot/backup.",
						MarkdownDescription: "The retain count of the snapshot/backup.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"task": schema.StringAttribute{
						Description:         "The recurring job task.Can be 'snapshot', 'snapshot-force-create', 'snapshot-cleanup', 'snapshot-delete', 'backup', 'backup-force-create' or 'filesystem-trim'",
						MarkdownDescription: "The recurring job task.Can be 'snapshot', 'snapshot-force-create', 'snapshot-cleanup', 'snapshot-delete', 'backup', 'backup-force-create' or 'filesystem-trim'",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("snapshot", "snapshot-force-create", "snapshot-cleanup", "snapshot-delete", "backup", "backup-force-create", "filesystem-trim"),
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

func (r *LonghornIoRecurringJobV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_longhorn_io_recurring_job_v1beta2_manifest")

	var model LonghornIoRecurringJobV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("longhorn.io/v1beta2")
	model.Kind = pointer.String("RecurringJob")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
