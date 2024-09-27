/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kueue_x_k8s_io_v1alpha1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &KueueXK8SIoCohortV1Alpha1Manifest{}
)

func NewKueueXK8SIoCohortV1Alpha1Manifest() datasource.DataSource {
	return &KueueXK8SIoCohortV1Alpha1Manifest{}
}

type KueueXK8SIoCohortV1Alpha1Manifest struct{}

type KueueXK8SIoCohortV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Parent         *string `tfsdk:"parent" json:"parent,omitempty"`
		ResourceGroups *[]struct {
			CoveredResources *[]string `tfsdk:"covered_resources" json:"coveredResources,omitempty"`
			Flavors          *[]struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Resources *[]struct {
					BorrowingLimit *string `tfsdk:"borrowing_limit" json:"borrowingLimit,omitempty"`
					LendingLimit   *string `tfsdk:"lending_limit" json:"lendingLimit,omitempty"`
					Name           *string `tfsdk:"name" json:"name,omitempty"`
					NominalQuota   *string `tfsdk:"nominal_quota" json:"nominalQuota,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"flavors" json:"flavors,omitempty"`
		} `tfsdk:"resource_groups" json:"resourceGroups,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KueueXK8SIoCohortV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kueue_x_k8s_io_cohort_v1alpha1_manifest"
}

func (r *KueueXK8SIoCohortV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Cohort is the Schema for the cohorts API. Using Hierarchical Cohorts (any Cohort which has a parent) with Fair Sharing results in undefined behavior in 0.9",
		MarkdownDescription: "Cohort is the Schema for the cohorts API. Using Hierarchical Cohorts (any Cohort which has a parent) with Fair Sharing results in undefined behavior in 0.9",
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
				Description:         "CohortSpec defines the desired state of Cohort",
				MarkdownDescription: "CohortSpec defines the desired state of Cohort",
				Attributes: map[string]schema.Attribute{
					"parent": schema.StringAttribute{
						Description:         "Parent references the name of the Cohort's parent, if any. It satisfies one of three cases: 1) Unset. This Cohort is the root of its Cohort tree. 2) References a non-existent Cohort. We use default Cohort (no borrowing/lending limits). 3) References an existent Cohort. If a cycle is created, we disable all members of the Cohort, including ClusterQueues, until the cycle is removed. We prevent further admission while the cycle exists.",
						MarkdownDescription: "Parent references the name of the Cohort's parent, if any. It satisfies one of three cases: 1) Unset. This Cohort is the root of its Cohort tree. 2) References a non-existent Cohort. We use default Cohort (no borrowing/lending limits). 3) References an existent Cohort. If a cycle is created, we disable all members of the Cohort, including ClusterQueues, until the cycle is removed. We prevent further admission while the cycle exists.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(253),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
						},
					},

					"resource_groups": schema.ListNestedAttribute{
						Description:         "ResourceGroups describes groupings of Resources and Flavors. Each ResourceGroup defines a list of Resources and a list of Flavors which provide quotas for these Resources. Each Resource and each Flavor may only form part of one ResourceGroup. There may be up to 16 ResourceGroups within a Cohort. BorrowingLimit limits how much members of this Cohort subtree can borrow from the parent subtree. LendingLimit limits how much members of this Cohort subtree can lend to the parent subtree. Borrowing and Lending limits must only be set when the Cohort has a parent. Otherwise, the Cohort create/update will be rejected by the webhook.",
						MarkdownDescription: "ResourceGroups describes groupings of Resources and Flavors. Each ResourceGroup defines a list of Resources and a list of Flavors which provide quotas for these Resources. Each Resource and each Flavor may only form part of one ResourceGroup. There may be up to 16 ResourceGroups within a Cohort. BorrowingLimit limits how much members of this Cohort subtree can borrow from the parent subtree. LendingLimit limits how much members of this Cohort subtree can lend to the parent subtree. Borrowing and Lending limits must only be set when the Cohort has a parent. Otherwise, the Cohort create/update will be rejected by the webhook.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"covered_resources": schema.ListAttribute{
									Description:         "coveredResources is the list of resources covered by the flavors in this group. Examples: cpu, memory, vendor.com/gpu. The list cannot be empty and it can contain up to 16 resources.",
									MarkdownDescription: "coveredResources is the list of resources covered by the flavors in this group. Examples: cpu, memory, vendor.com/gpu. The list cannot be empty and it can contain up to 16 resources.",
									ElementType:         types.StringType,
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"flavors": schema.ListNestedAttribute{
									Description:         "flavors is the list of flavors that provide the resources of this group. Typically, different flavors represent different hardware models (e.g., gpu models, cpu architectures) or pricing models (on-demand vs spot cpus). Each flavor MUST list all the resources listed for this group in the same order as the .resources field. The list cannot be empty and it can contain up to 16 flavors.",
									MarkdownDescription: "flavors is the list of flavors that provide the resources of this group. Typically, different flavors represent different hardware models (e.g., gpu models, cpu architectures) or pricing models (on-demand vs spot cpus). Each flavor MUST list all the resources listed for this group in the same order as the .resources field. The list cannot be empty and it can contain up to 16 flavors.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "name of this flavor. The name should match the .metadata.name of a ResourceFlavor. If a matching ResourceFlavor does not exist, the ClusterQueue will have an Active condition set to False.",
												MarkdownDescription: "name of this flavor. The name should match the .metadata.name of a ResourceFlavor. If a matching ResourceFlavor does not exist, the ClusterQueue will have an Active condition set to False.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(253),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
												},
											},

											"resources": schema.ListNestedAttribute{
												Description:         "resources is the list of quotas for this flavor per resource. There could be up to 16 resources.",
												MarkdownDescription: "resources is the list of quotas for this flavor per resource. There could be up to 16 resources.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"borrowing_limit": schema.StringAttribute{
															Description:         "borrowingLimit is the maximum amount of quota for the [flavor, resource] combination that this ClusterQueue is allowed to borrow from the unused quota of other ClusterQueues in the same cohort. In total, at a given time, Workloads in a ClusterQueue can consume a quantity of quota equal to nominalQuota+borrowingLimit, assuming the other ClusterQueues in the cohort have enough unused quota. If null, it means that there is no borrowing limit. If not null, it must be non-negative. borrowingLimit must be null if spec.cohort is empty.",
															MarkdownDescription: "borrowingLimit is the maximum amount of quota for the [flavor, resource] combination that this ClusterQueue is allowed to borrow from the unused quota of other ClusterQueues in the same cohort. In total, at a given time, Workloads in a ClusterQueue can consume a quantity of quota equal to nominalQuota+borrowingLimit, assuming the other ClusterQueues in the cohort have enough unused quota. If null, it means that there is no borrowing limit. If not null, it must be non-negative. borrowingLimit must be null if spec.cohort is empty.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"lending_limit": schema.StringAttribute{
															Description:         "lendingLimit is the maximum amount of unused quota for the [flavor, resource] combination that this ClusterQueue can lend to other ClusterQueues in the same cohort. In total, at a given time, ClusterQueue reserves for its exclusive use a quantity of quota equals to nominalQuota - lendingLimit. If null, it means that there is no lending limit, meaning that all the nominalQuota can be borrowed by other clusterQueues in the cohort. If not null, it must be non-negative. lendingLimit must be null if spec.cohort is empty. This field is in beta stage and is enabled by default.",
															MarkdownDescription: "lendingLimit is the maximum amount of unused quota for the [flavor, resource] combination that this ClusterQueue can lend to other ClusterQueues in the same cohort. In total, at a given time, ClusterQueue reserves for its exclusive use a quantity of quota equals to nominalQuota - lendingLimit. If null, it means that there is no lending limit, meaning that all the nominalQuota can be borrowed by other clusterQueues in the cohort. If not null, it must be non-negative. lendingLimit must be null if spec.cohort is empty. This field is in beta stage and is enabled by default.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "name of this resource.",
															MarkdownDescription: "name of this resource.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"nominal_quota": schema.StringAttribute{
															Description:         "nominalQuota is the quantity of this resource that is available for Workloads admitted by this ClusterQueue at a point in time. The nominalQuota must be non-negative. nominalQuota should represent the resources in the cluster available for running jobs (after discounting resources consumed by system components and pods not managed by kueue). In an autoscaled cluster, nominalQuota should account for resources that can be provided by a component such as Kubernetes cluster-autoscaler. If the ClusterQueue belongs to a cohort, the sum of the quotas for each (flavor, resource) combination defines the maximum quantity that can be allocated by a ClusterQueue in the cohort.",
															MarkdownDescription: "nominalQuota is the quantity of this resource that is available for Workloads admitted by this ClusterQueue at a point in time. The nominalQuota must be non-negative. nominalQuota should represent the resources in the cluster available for running jobs (after discounting resources consumed by system components and pods not managed by kueue). In an autoscaled cluster, nominalQuota should account for resources that can be provided by a component such as Kubernetes cluster-autoscaler. If the ClusterQueue belongs to a cohort, the sum of the quotas for each (flavor, resource) combination defines the maximum quantity that can be allocated by a ClusterQueue in the cohort.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
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

func (r *KueueXK8SIoCohortV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kueue_x_k8s_io_cohort_v1alpha1_manifest")

	var model KueueXK8SIoCohortV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kueue.x-k8s.io/v1alpha1")
	model.Kind = pointer.String("Cohort")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
