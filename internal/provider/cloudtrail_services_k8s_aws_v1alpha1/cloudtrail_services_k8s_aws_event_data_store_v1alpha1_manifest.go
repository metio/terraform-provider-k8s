/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cloudtrail_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &CloudtrailServicesK8SAwsEventDataStoreV1Alpha1Manifest{}
)

func NewCloudtrailServicesK8SAwsEventDataStoreV1Alpha1Manifest() datasource.DataSource {
	return &CloudtrailServicesK8SAwsEventDataStoreV1Alpha1Manifest{}
}

type CloudtrailServicesK8SAwsEventDataStoreV1Alpha1Manifest struct{}

type CloudtrailServicesK8SAwsEventDataStoreV1Alpha1ManifestData struct {
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
		AdvancedEventSelectors *[]struct {
			FieldSelectors *[]struct {
				EndsWith      *[]string `tfsdk:"ends_with" json:"endsWith,omitempty"`
				Equals        *[]string `tfsdk:"equals" json:"equals,omitempty"`
				Field         *string   `tfsdk:"field" json:"field,omitempty"`
				NotEndsWith   *[]string `tfsdk:"not_ends_with" json:"notEndsWith,omitempty"`
				NotEquals     *[]string `tfsdk:"not_equals" json:"notEquals,omitempty"`
				NotStartsWith *[]string `tfsdk:"not_starts_with" json:"notStartsWith,omitempty"`
				StartsWith    *[]string `tfsdk:"starts_with" json:"startsWith,omitempty"`
			} `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"advanced_event_selectors" json:"advancedEventSelectors,omitempty"`
		MultiRegionEnabled  *bool   `tfsdk:"multi_region_enabled" json:"multiRegionEnabled,omitempty"`
		Name                *string `tfsdk:"name" json:"name,omitempty"`
		OrganizationEnabled *bool   `tfsdk:"organization_enabled" json:"organizationEnabled,omitempty"`
		RetentionPeriod     *int64  `tfsdk:"retention_period" json:"retentionPeriod,omitempty"`
		Tags                *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		TerminationProtectionEnabled *bool `tfsdk:"termination_protection_enabled" json:"terminationProtectionEnabled,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CloudtrailServicesK8SAwsEventDataStoreV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cloudtrail_services_k8s_aws_event_data_store_v1alpha1_manifest"
}

func (r *CloudtrailServicesK8SAwsEventDataStoreV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "EventDataStore is the Schema for the EventDataStores API",
		MarkdownDescription: "EventDataStore is the Schema for the EventDataStores API",
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
				Description:         "EventDataStoreSpec defines the desired state of EventDataStore.A storage lake of event data against which you can run complex SQL-basedqueries. An event data store can include events that you have logged on youraccount from the last 90 to 2555 days (about three months to up to sevenyears). To select events for an event data store, use advanced event selectors(https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced).",
				MarkdownDescription: "EventDataStoreSpec defines the desired state of EventDataStore.A storage lake of event data against which you can run complex SQL-basedqueries. An event data store can include events that you have logged on youraccount from the last 90 to 2555 days (about three months to up to sevenyears). To select events for an event data store, use advanced event selectors(https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced).",
				Attributes: map[string]schema.Attribute{
					"advanced_event_selectors": schema.ListNestedAttribute{
						Description:         "The advanced event selectors to use to select the events for the data store.For more information about how to use advanced event selectors, see Log eventsby using advanced event selectors (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced)in the CloudTrail User Guide.",
						MarkdownDescription: "The advanced event selectors to use to select the events for the data store.For more information about how to use advanced event selectors, see Log eventsby using advanced event selectors (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced)in the CloudTrail User Guide.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"field_selectors": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"ends_with": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"equals": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"field": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"not_ends_with": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"not_equals": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"not_starts_with": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"starts_with": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
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

								"name": schema.StringAttribute{
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

					"multi_region_enabled": schema.BoolAttribute{
						Description:         "Specifies whether the event data store includes events from all regions,or only from the region in which the event data store is created.",
						MarkdownDescription: "Specifies whether the event data store includes events from all regions,or only from the region in which the event data store is created.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "The name of the event data store.",
						MarkdownDescription: "The name of the event data store.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"organization_enabled": schema.BoolAttribute{
						Description:         "Specifies whether an event data store collects events logged for an organizationin Organizations.",
						MarkdownDescription: "Specifies whether an event data store collects events logged for an organizationin Organizations.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"retention_period": schema.Int64Attribute{
						Description:         "The retention period of the event data store, in days. You can set a retentionperiod of up to 2555 days, the equivalent of seven years.",
						MarkdownDescription: "The retention period of the event data store, in days. You can set a retentionperiod of up to 2555 days, the equivalent of seven years.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
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

					"termination_protection_enabled": schema.BoolAttribute{
						Description:         "Specifies whether termination protection is enabled for the event data store.If termination protection is enabled, you cannot delete the event data storeuntil termination protection is disabled.",
						MarkdownDescription: "Specifies whether termination protection is enabled for the event data store.If termination protection is enabled, you cannot delete the event data storeuntil termination protection is disabled.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *CloudtrailServicesK8SAwsEventDataStoreV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cloudtrail_services_k8s_aws_event_data_store_v1alpha1_manifest")

	var model CloudtrailServicesK8SAwsEventDataStoreV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cloudtrail.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("EventDataStore")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
