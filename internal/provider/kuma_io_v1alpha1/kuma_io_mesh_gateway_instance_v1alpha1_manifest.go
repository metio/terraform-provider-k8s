/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuma_io_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &KumaIoMeshGatewayInstanceV1Alpha1Manifest{}
)

func NewKumaIoMeshGatewayInstanceV1Alpha1Manifest() datasource.DataSource {
	return &KumaIoMeshGatewayInstanceV1Alpha1Manifest{}
}

type KumaIoMeshGatewayInstanceV1Alpha1Manifest struct{}

type KumaIoMeshGatewayInstanceV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		PodTemplate *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec *struct {
				Container *struct {
					SecurityContext *struct {
						ReadOnlyRootFilesystem *bool `tfsdk:"read_only_root_filesystem" json:"readOnlyRootFilesystem,omitempty"`
					} `tfsdk:"security_context" json:"securityContext,omitempty"`
				} `tfsdk:"container" json:"container,omitempty"`
				SecurityContext *struct {
					FsGroup *int64 `tfsdk:"fs_group" json:"fsGroup,omitempty"`
				} `tfsdk:"security_context" json:"securityContext,omitempty"`
				ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"pod_template" json:"podTemplate,omitempty"`
		Replicas  *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
		Resources *struct {
			Claims *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"claims" json:"claims,omitempty"`
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		ServiceTemplate *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec *struct {
				LoadBalancerIP *string `tfsdk:"load_balancer_ip" json:"loadBalancerIP,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"service_template" json:"serviceTemplate,omitempty"`
		ServiceType *string            `tfsdk:"service_type" json:"serviceType,omitempty"`
		Tags        *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KumaIoMeshGatewayInstanceV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuma_io_mesh_gateway_instance_v1alpha1_manifest"
}

func (r *KumaIoMeshGatewayInstanceV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "MeshGatewayInstance represents a managed instance of a dataplane proxy for a Kuma Gateway.",
		MarkdownDescription: "MeshGatewayInstance represents a managed instance of a dataplane proxy for a Kuma Gateway.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "MeshGatewayInstanceSpec specifies the options available for a GatewayDataplane.",
				MarkdownDescription: "MeshGatewayInstanceSpec specifies the options available for a GatewayDataplane.",
				Attributes: map[string]schema.Attribute{
					"pod_template": schema.SingleNestedAttribute{
						Description:         "PodTemplate configures the Pod owned by this config.",
						MarkdownDescription: "PodTemplate configures the Pod owned by this config.",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description:         "Metadata holds metadata configuration for a Service.",
								MarkdownDescription: "Metadata holds metadata configuration for a Service.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations holds annotations to be set on an object.",
										MarkdownDescription: "Annotations holds annotations to be set on an object.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels holds labels to be set on an objects.",
										MarkdownDescription: "Labels holds labels to be set on an objects.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"spec": schema.SingleNestedAttribute{
								Description:         "Spec holds some customizable fields of a Pod.",
								MarkdownDescription: "Spec holds some customizable fields of a Pod.",
								Attributes: map[string]schema.Attribute{
									"container": schema.SingleNestedAttribute{
										Description:         "Container corresponds to PodSpec.Container",
										MarkdownDescription: "Container corresponds to PodSpec.Container",
										Attributes: map[string]schema.Attribute{
											"security_context": schema.SingleNestedAttribute{
												Description:         "ContainerSecurityContext corresponds to PodSpec.Container.SecurityContext",
												MarkdownDescription: "ContainerSecurityContext corresponds to PodSpec.Container.SecurityContext",
												Attributes: map[string]schema.Attribute{
													"read_only_root_filesystem": schema.BoolAttribute{
														Description:         "ReadOnlyRootFilesystem corresponds to PodSpec.Container.SecurityContext.ReadOnlyRootFilesystem",
														MarkdownDescription: "ReadOnlyRootFilesystem corresponds to PodSpec.Container.SecurityContext.ReadOnlyRootFilesystem",
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"security_context": schema.SingleNestedAttribute{
										Description:         "PodSecurityContext corresponds to PodSpec.SecurityContext",
										MarkdownDescription: "PodSecurityContext corresponds to PodSpec.SecurityContext",
										Attributes: map[string]schema.Attribute{
											"fs_group": schema.Int64Attribute{
												Description:         "FSGroup corresponds to PodSpec.SecurityContext.FSGroup",
												MarkdownDescription: "FSGroup corresponds to PodSpec.SecurityContext.FSGroup",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"service_account_name": schema.StringAttribute{
										Description:         "ServiceAccountName corresponds to PodSpec.ServiceAccountName.",
										MarkdownDescription: "ServiceAccountName corresponds to PodSpec.ServiceAccountName.",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"replicas": schema.Int64Attribute{
						Description:         "Replicas is the number of dataplane proxy replicas to create. For now this is a fixed number, but in the future it could be automatically scaled based on metrics.",
						MarkdownDescription: "Replicas is the number of dataplane proxy replicas to create. For now this is a fixed number, but in the future it could be automatically scaled based on metrics.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "Resources specifies the compute resources for the proxy container. The default can be set in the control plane config.",
						MarkdownDescription: "Resources specifies the compute resources for the proxy container. The default can be set in the control plane config.",
						Attributes: map[string]schema.Attribute{
							"claims": schema.ListNestedAttribute{
								Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
								MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
											MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

							"limits": schema.MapAttribute{
								Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"requests": schema.MapAttribute{
								Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_template": schema.SingleNestedAttribute{
						Description:         "ServiceTemplate configures the Service owned by this config.",
						MarkdownDescription: "ServiceTemplate configures the Service owned by this config.",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description:         "Metadata holds metadata configuration for a Service.",
								MarkdownDescription: "Metadata holds metadata configuration for a Service.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations holds annotations to be set on an object.",
										MarkdownDescription: "Annotations holds annotations to be set on an object.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels holds labels to be set on an objects.",
										MarkdownDescription: "Labels holds labels to be set on an objects.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"spec": schema.SingleNestedAttribute{
								Description:         "Spec holds some customizable fields of a Service.",
								MarkdownDescription: "Spec holds some customizable fields of a Service.",
								Attributes: map[string]schema.Attribute{
									"load_balancer_ip": schema.StringAttribute{
										Description:         "LoadBalancerIP corresponds to ServiceSpec.LoadBalancerIP.",
										MarkdownDescription: "LoadBalancerIP corresponds to ServiceSpec.LoadBalancerIP.",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_type": schema.StringAttribute{
						Description:         "ServiceType specifies the type of managed Service that will be created to expose the dataplane proxies to traffic from outside the cluster. The ports to expose will be taken from the matching Gateway resource. If there is no matching Gateway, the managed Service will be deleted.",
						MarkdownDescription: "ServiceType specifies the type of managed Service that will be created to expose the dataplane proxies to traffic from outside the cluster. The ports to expose will be taken from the matching Gateway resource. If there is no matching Gateway, the managed Service will be deleted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("LoadBalancer", "ClusterIP", "NodePort"),
						},
					},

					"tags": schema.MapAttribute{
						Description:         "Tags specifies the Kuma tags that are propagated to the managed dataplane proxies. These tags should include exactly one 'kuma.io/service' tag, and should match exactly one Gateway resource.",
						MarkdownDescription: "Tags specifies the Kuma tags that are propagated to the managed dataplane proxies. These tags should include exactly one 'kuma.io/service' tag, and should match exactly one Gateway resource.",
						ElementType:         types.StringType,
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

func (r *KumaIoMeshGatewayInstanceV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kuma_io_mesh_gateway_instance_v1alpha1_manifest")

	var model KumaIoMeshGatewayInstanceV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("kuma.io/v1alpha1")
	model.Kind = pointer.String("MeshGatewayInstance")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
