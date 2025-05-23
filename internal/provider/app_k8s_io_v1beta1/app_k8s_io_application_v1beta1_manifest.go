/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package app_k8s_io_v1beta1

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
	_ datasource.DataSource = &AppK8SIoApplicationV1Beta1Manifest{}
)

func NewAppK8SIoApplicationV1Beta1Manifest() datasource.DataSource {
	return &AppK8SIoApplicationV1Beta1Manifest{}
}

type AppK8SIoApplicationV1Beta1Manifest struct{}

type AppK8SIoApplicationV1Beta1ManifestData struct {
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
		AddOwnerRef    *bool   `tfsdk:"add_owner_ref" json:"addOwnerRef,omitempty"`
		AssemblyPhase  *string `tfsdk:"assembly_phase" json:"assemblyPhase,omitempty"`
		ComponentKinds *[]struct {
			Group *string `tfsdk:"group" json:"group,omitempty"`
			Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
		} `tfsdk:"component_kinds" json:"componentKinds,omitempty"`
		Descriptor *struct {
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Icons       *[]struct {
				Size *string `tfsdk:"size" json:"size,omitempty"`
				Src  *string `tfsdk:"src" json:"src,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"icons" json:"icons,omitempty"`
			Keywords *[]string `tfsdk:"keywords" json:"keywords,omitempty"`
			Links    *[]struct {
				Description *string `tfsdk:"description" json:"description,omitempty"`
				Url         *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"links" json:"links,omitempty"`
			Maintainers *[]struct {
				Email *string `tfsdk:"email" json:"email,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Url   *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"maintainers" json:"maintainers,omitempty"`
			Notes  *string `tfsdk:"notes" json:"notes,omitempty"`
			Owners *[]struct {
				Email *string `tfsdk:"email" json:"email,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Url   *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"owners" json:"owners,omitempty"`
			Type    *string `tfsdk:"type" json:"type,omitempty"`
			Version *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"descriptor" json:"descriptor,omitempty"`
		Info *[]struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Type      *string `tfsdk:"type" json:"type,omitempty"`
			Value     *string `tfsdk:"value" json:"value,omitempty"`
			ValueFrom *struct {
				ConfigMapKeyRef *struct {
					ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
					Key             *string `tfsdk:"key" json:"key,omitempty"`
					Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
					Name            *string `tfsdk:"name" json:"name,omitempty"`
					Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
					ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
					Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
				} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
				IngressRef *struct {
					ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
					Host            *string `tfsdk:"host" json:"host,omitempty"`
					Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
					Name            *string `tfsdk:"name" json:"name,omitempty"`
					Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Path            *string `tfsdk:"path" json:"path,omitempty"`
					Protocol        *string `tfsdk:"protocol" json:"protocol,omitempty"`
					ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
					Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
				} `tfsdk:"ingress_ref" json:"ingressRef,omitempty"`
				SecretKeyRef *struct {
					ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
					Key             *string `tfsdk:"key" json:"key,omitempty"`
					Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
					Name            *string `tfsdk:"name" json:"name,omitempty"`
					Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
					ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
					Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
				} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				ServiceRef *struct {
					ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
					Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
					Name            *string `tfsdk:"name" json:"name,omitempty"`
					Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Path            *string `tfsdk:"path" json:"path,omitempty"`
					Port            *int64  `tfsdk:"port" json:"port,omitempty"`
					Protocol        *string `tfsdk:"protocol" json:"protocol,omitempty"`
					ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
					Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
				} `tfsdk:"service_ref" json:"serviceRef,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"value_from" json:"valueFrom,omitempty"`
		} `tfsdk:"info" json:"info,omitempty"`
		Selector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppK8SIoApplicationV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_app_k8s_io_application_v1beta1_manifest"
}

func (r *AppK8SIoApplicationV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Application is the Schema for the applications API",
		MarkdownDescription: "Application is the Schema for the applications API",
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
				Description:         "ApplicationSpec defines the specification for an Application.",
				MarkdownDescription: "ApplicationSpec defines the specification for an Application.",
				Attributes: map[string]schema.Attribute{
					"add_owner_ref": schema.BoolAttribute{
						Description:         "AddOwnerRef objects - flag to indicate if we need to add OwnerRefs to matching objects Matching is done by using Selector to query all ComponentGroupKinds",
						MarkdownDescription: "AddOwnerRef objects - flag to indicate if we need to add OwnerRefs to matching objects Matching is done by using Selector to query all ComponentGroupKinds",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"assembly_phase": schema.StringAttribute{
						Description:         "AssemblyPhase represents the current phase of the application's assembly. An empty value is equivalent to 'Succeeded'.",
						MarkdownDescription: "AssemblyPhase represents the current phase of the application's assembly. An empty value is equivalent to 'Succeeded'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"component_kinds": schema.ListNestedAttribute{
						Description:         "ComponentGroupKinds is a list of Kinds for Application's components (e.g. Deployments, Pods, Services, CRDs). It can be used in conjunction with the Application's Selector to list or watch the Applications components.",
						MarkdownDescription: "ComponentGroupKinds is a list of Kinds for Application's components (e.g. Deployments, Pods, Services, CRDs). It can be used in conjunction with the Application's Selector to list or watch the Applications components.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
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

					"descriptor": schema.SingleNestedAttribute{
						Description:         "Descriptor regroups information and metadata about an application.",
						MarkdownDescription: "Descriptor regroups information and metadata about an application.",
						Attributes: map[string]schema.Attribute{
							"description": schema.StringAttribute{
								Description:         "Description is a brief string description of the Application.",
								MarkdownDescription: "Description is a brief string description of the Application.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"icons": schema.ListNestedAttribute{
								Description:         "Icons is an optional list of icons for an application. Icon information includes the source, size, and mime type.",
								MarkdownDescription: "Icons is an optional list of icons for an application. Icon information includes the source, size, and mime type.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"size": schema.StringAttribute{
											Description:         "(optional) The size of the image in pixels (e.g., 25x25).",
											MarkdownDescription: "(optional) The size of the image in pixels (e.g., 25x25).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"src": schema.StringAttribute{
											Description:         "The source for image represented as either an absolute URL to the image or a Data URL containing the image. Data URLs are defined in RFC 2397.",
											MarkdownDescription: "The source for image represented as either an absolute URL to the image or a Data URL containing the image. Data URLs are defined in RFC 2397.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"type": schema.StringAttribute{
											Description:         "(optional) The mine type of the image (e.g., 'image/png').",
											MarkdownDescription: "(optional) The mine type of the image (e.g., 'image/png').",
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

							"keywords": schema.ListAttribute{
								Description:         "Keywords is an optional list of key words associated with the application (e.g. MySQL, RDBMS, database).",
								MarkdownDescription: "Keywords is an optional list of key words associated with the application (e.g. MySQL, RDBMS, database).",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"links": schema.ListNestedAttribute{
								Description:         "Links are a list of descriptive URLs intended to be used to surface additional documentation, dashboards, etc.",
								MarkdownDescription: "Links are a list of descriptive URLs intended to be used to surface additional documentation, dashboards, etc.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"description": schema.StringAttribute{
											Description:         "Description is human readable content explaining the purpose of the link.",
											MarkdownDescription: "Description is human readable content explaining the purpose of the link.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"url": schema.StringAttribute{
											Description:         "Url typically points at a website address.",
											MarkdownDescription: "Url typically points at a website address.",
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

							"maintainers": schema.ListNestedAttribute{
								Description:         "Maintainers is an optional list of maintainers of the application. The maintainers in this list maintain the the source code, images, and package for the application.",
								MarkdownDescription: "Maintainers is an optional list of maintainers of the application. The maintainers in this list maintain the the source code, images, and package for the application.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"email": schema.StringAttribute{
											Description:         "Email is the email address.",
											MarkdownDescription: "Email is the email address.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name is the descriptive name.",
											MarkdownDescription: "Name is the descriptive name.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"url": schema.StringAttribute{
											Description:         "Url could typically be a website address.",
											MarkdownDescription: "Url could typically be a website address.",
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

							"notes": schema.StringAttribute{
								Description:         "Notes contain a human readable snippets intended as a quick start for the users of the Application. CommonMark markdown syntax may be used for rich text representation.",
								MarkdownDescription: "Notes contain a human readable snippets intended as a quick start for the users of the Application. CommonMark markdown syntax may be used for rich text representation.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"owners": schema.ListNestedAttribute{
								Description:         "Owners is an optional list of the owners of the installed application. The owners of the application should be contacted in the event of a planned or unplanned disruption affecting the application.",
								MarkdownDescription: "Owners is an optional list of the owners of the installed application. The owners of the application should be contacted in the event of a planned or unplanned disruption affecting the application.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"email": schema.StringAttribute{
											Description:         "Email is the email address.",
											MarkdownDescription: "Email is the email address.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name is the descriptive name.",
											MarkdownDescription: "Name is the descriptive name.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"url": schema.StringAttribute{
											Description:         "Url could typically be a website address.",
											MarkdownDescription: "Url could typically be a website address.",
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

							"type": schema.StringAttribute{
								Description:         "Type is the type of the application (e.g. WordPress, MySQL, Cassandra).",
								MarkdownDescription: "Type is the type of the application (e.g. WordPress, MySQL, Cassandra).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"version": schema.StringAttribute{
								Description:         "Version is an optional version indicator for the Application.",
								MarkdownDescription: "Version is an optional version indicator for the Application.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"info": schema.ListNestedAttribute{
						Description:         "Info contains human readable key,value pairs for the Application.",
						MarkdownDescription: "Info contains human readable key,value pairs for the Application.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name is a human readable title for this piece of information.",
									MarkdownDescription: "Name is a human readable title for this piece of information.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"type": schema.StringAttribute{
									Description:         "Type of the value for this InfoItem.",
									MarkdownDescription: "Type of the value for this InfoItem.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Value is human readable content.",
									MarkdownDescription: "Value is human readable content.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value_from": schema.SingleNestedAttribute{
									Description:         "ValueFrom defines a reference to derive the value from another source.",
									MarkdownDescription: "ValueFrom defines a reference to derive the value from another source.",
									Attributes: map[string]schema.Attribute{
										"config_map_key_ref": schema.SingleNestedAttribute{
											Description:         "Selects a key of a ConfigMap.",
											MarkdownDescription: "Selects a key of a ConfigMap.",
											Attributes: map[string]schema.Attribute{
												"api_version": schema.StringAttribute{
													Description:         "API version of the referent.",
													MarkdownDescription: "API version of the referent.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"field_path": schema.StringAttribute{
													Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
													MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"key": schema.StringAttribute{
													Description:         "The key to select.",
													MarkdownDescription: "The key to select.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kind": schema.StringAttribute{
													Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
													MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
													MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"resource_version": schema.StringAttribute{
													Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
													MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uid": schema.StringAttribute{
													Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
													MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"ingress_ref": schema.SingleNestedAttribute{
											Description:         "Select an Ingress.",
											MarkdownDescription: "Select an Ingress.",
											Attributes: map[string]schema.Attribute{
												"api_version": schema.StringAttribute{
													Description:         "API version of the referent.",
													MarkdownDescription: "API version of the referent.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"field_path": schema.StringAttribute{
													Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
													MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"host": schema.StringAttribute{
													Description:         "The optional host to select.",
													MarkdownDescription: "The optional host to select.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kind": schema.StringAttribute{
													Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
													MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
													MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"path": schema.StringAttribute{
													Description:         "The optional HTTP path.",
													MarkdownDescription: "The optional HTTP path.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"protocol": schema.StringAttribute{
													Description:         "Protocol for the ingress",
													MarkdownDescription: "Protocol for the ingress",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"resource_version": schema.StringAttribute{
													Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
													MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uid": schema.StringAttribute{
													Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
													MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"secret_key_ref": schema.SingleNestedAttribute{
											Description:         "Selects a key of a Secret.",
											MarkdownDescription: "Selects a key of a Secret.",
											Attributes: map[string]schema.Attribute{
												"api_version": schema.StringAttribute{
													Description:         "API version of the referent.",
													MarkdownDescription: "API version of the referent.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"field_path": schema.StringAttribute{
													Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
													MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"key": schema.StringAttribute{
													Description:         "The key to select.",
													MarkdownDescription: "The key to select.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kind": schema.StringAttribute{
													Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
													MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
													MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"resource_version": schema.StringAttribute{
													Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
													MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uid": schema.StringAttribute{
													Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
													MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"service_ref": schema.SingleNestedAttribute{
											Description:         "Select a Service.",
											MarkdownDescription: "Select a Service.",
											Attributes: map[string]schema.Attribute{
												"api_version": schema.StringAttribute{
													Description:         "API version of the referent.",
													MarkdownDescription: "API version of the referent.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"field_path": schema.StringAttribute{
													Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
													MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kind": schema.StringAttribute{
													Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
													MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
													MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"path": schema.StringAttribute{
													Description:         "The optional HTTP path.",
													MarkdownDescription: "The optional HTTP path.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "The optional port to select.",
													MarkdownDescription: "The optional port to select.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"protocol": schema.StringAttribute{
													Description:         "Protocol for the service",
													MarkdownDescription: "Protocol for the service",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"resource_version": schema.StringAttribute{
													Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
													MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uid": schema.StringAttribute{
													Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
													MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"type": schema.StringAttribute{
											Description:         "Type of source.",
											MarkdownDescription: "Type of source.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"selector": schema.SingleNestedAttribute{
						Description:         "Selector is a label query over kinds that created by the application. It must match the component objects' labels. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors",
						MarkdownDescription: "Selector is a label query over kinds that created by the application. It must match the component objects' labels. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *AppK8SIoApplicationV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_app_k8s_io_application_v1beta1_manifest")

	var model AppK8SIoApplicationV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("app.k8s.io/v1beta1")
	model.Kind = pointer.String("Application")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
