/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package eks_services_k8s_aws_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &EksServicesK8SAwsFargateProfileV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &EksServicesK8SAwsFargateProfileV1Alpha1DataSource{}
)

func NewEksServicesK8SAwsFargateProfileV1Alpha1DataSource() datasource.DataSource {
	return &EksServicesK8SAwsFargateProfileV1Alpha1DataSource{}
}

type EksServicesK8SAwsFargateProfileV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type EksServicesK8SAwsFargateProfileV1Alpha1DataSourceData struct {
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
		ClientRequestToken *string `tfsdk:"client_request_token" json:"clientRequestToken,omitempty"`
		ClusterName        *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
		ClusterRef         *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"cluster_ref" json:"clusterRef,omitempty"`
		Name                *string `tfsdk:"name" json:"name,omitempty"`
		PodExecutionRoleARN *string `tfsdk:"pod_execution_role_arn" json:"podExecutionRoleARN,omitempty"`
		PodExecutionRoleRef *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"pod_execution_role_ref" json:"podExecutionRoleRef,omitempty"`
		Selectors *[]struct {
			Labels    *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Namespace *string            `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"selectors" json:"selectors,omitempty"`
		SubnetRefs *[]struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"subnet_refs" json:"subnetRefs,omitempty"`
		Subnets *[]string          `tfsdk:"subnets" json:"subnets,omitempty"`
		Tags    *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *EksServicesK8SAwsFargateProfileV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_eks_services_k8s_aws_fargate_profile_v1alpha1"
}

func (r *EksServicesK8SAwsFargateProfileV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "FargateProfile is the Schema for the FargateProfiles API",
		MarkdownDescription: "FargateProfile is the Schema for the FargateProfiles API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "FargateProfileSpec defines the desired state of FargateProfile.  An object representing an Fargate profile.",
				MarkdownDescription: "FargateProfileSpec defines the desired state of FargateProfile.  An object representing an Fargate profile.",
				Attributes: map[string]schema.Attribute{
					"client_request_token": schema.StringAttribute{
						Description:         "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.",
						MarkdownDescription: "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cluster_name": schema.StringAttribute{
						Description:         "The name of the Amazon EKS cluster to apply the Fargate profile to.",
						MarkdownDescription: "The name of the Amazon EKS cluster to apply the Fargate profile to.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cluster_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef:  from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef:  from: name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
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

					"name": schema.StringAttribute{
						Description:         "The name of the Fargate profile.",
						MarkdownDescription: "The name of the Fargate profile.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"pod_execution_role_arn": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) of the pod execution role to use for pods that match the selectors in the Fargate profile. The pod execution role allows Fargate infrastructure to register with your cluster as a node, and it provides read access to Amazon ECR image repositories. For more information, see Pod Execution Role (https://docs.aws.amazon.com/eks/latest/userguide/pod-execution-role.html) in the Amazon EKS User Guide.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of the pod execution role to use for pods that match the selectors in the Fargate profile. The pod execution role allows Fargate infrastructure to register with your cluster as a node, and it provides read access to Amazon ECR image repositories. For more information, see Pod Execution Role (https://docs.aws.amazon.com/eks/latest/userguide/pod-execution-role.html) in the Amazon EKS User Guide.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"pod_execution_role_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef:  from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef:  from: name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
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

					"selectors": schema.ListNestedAttribute{
						Description:         "The selectors to match for pods to use this Fargate profile. Each selector must have an associated namespace. Optionally, you can also specify labels for a namespace. You may specify up to five selectors in a Fargate profile.",
						MarkdownDescription: "The selectors to match for pods to use this Fargate profile. Each selector must have an associated namespace. Optionally, you can also specify labels for a namespace. You may specify up to five selectors in a Fargate profile.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"labels": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"namespace": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"subnet_refs": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"from": schema.SingleNestedAttribute{
									Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
									MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"subnets": schema.ListAttribute{
						Description:         "The IDs of subnets to launch your pods into. At this time, pods running on Fargate are not assigned public IP addresses, so only private subnets (with no direct route to an Internet Gateway) are accepted for this parameter.",
						MarkdownDescription: "The IDs of subnets to launch your pods into. At this time, pods running on Fargate are not assigned public IP addresses, so only private subnets (with no direct route to an Internet Gateway) are accepted for this parameter.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"tags": schema.MapAttribute{
						Description:         "The metadata to apply to the Fargate profile to assist with categorization and organization. Each tag consists of a key and an optional value. You define both. Fargate profile tags do not propagate to any other resources associated with the Fargate profile, such as the pods that are scheduled with it.",
						MarkdownDescription: "The metadata to apply to the Fargate profile to assist with categorization and organization. Each tag consists of a key and an optional value. You define both. Fargate profile tags do not propagate to any other resources associated with the Fargate profile, such as the pods that are scheduled with it.",
						ElementType:         types.StringType,
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
	}
}

func (r *EksServicesK8SAwsFargateProfileV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *EksServicesK8SAwsFargateProfileV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_eks_services_k8s_aws_fargate_profile_v1alpha1")

	var data EksServicesK8SAwsFargateProfileV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "eks.services.k8s.aws", Version: "v1alpha1", Resource: "fargateprofiles"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse EksServicesK8SAwsFargateProfileV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("eks.services.k8s.aws/v1alpha1")
	data.Kind = pointer.String("FargateProfile")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
