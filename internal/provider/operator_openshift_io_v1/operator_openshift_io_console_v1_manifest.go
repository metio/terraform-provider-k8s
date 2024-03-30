/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_openshift_io_v1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
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
	_ datasource.DataSource = &OperatorOpenshiftIoConsoleV1Manifest{}
)

func NewOperatorOpenshiftIoConsoleV1Manifest() datasource.DataSource {
	return &OperatorOpenshiftIoConsoleV1Manifest{}
}

type OperatorOpenshiftIoConsoleV1Manifest struct{}

type OperatorOpenshiftIoConsoleV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Customization *struct {
			AddPage *struct {
				DisabledActions *[]string `tfsdk:"disabled_actions" json:"disabledActions,omitempty"`
			} `tfsdk:"add_page" json:"addPage,omitempty"`
			Brand          *string `tfsdk:"brand" json:"brand,omitempty"`
			CustomLogoFile *struct {
				Key  *string `tfsdk:"key" json:"key,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"custom_logo_file" json:"customLogoFile,omitempty"`
			CustomProductName *string `tfsdk:"custom_product_name" json:"customProductName,omitempty"`
			DeveloperCatalog  *struct {
				Categories *[]struct {
					Id            *string `tfsdk:"id" json:"id,omitempty"`
					Label         *string `tfsdk:"label" json:"label,omitempty"`
					Subcategories *[]struct {
						Id    *string   `tfsdk:"id" json:"id,omitempty"`
						Label *string   `tfsdk:"label" json:"label,omitempty"`
						Tags  *[]string `tfsdk:"tags" json:"tags,omitempty"`
					} `tfsdk:"subcategories" json:"subcategories,omitempty"`
					Tags *[]string `tfsdk:"tags" json:"tags,omitempty"`
				} `tfsdk:"categories" json:"categories,omitempty"`
				Types *struct {
					Disabled *[]string `tfsdk:"disabled" json:"disabled,omitempty"`
					Enabled  *[]string `tfsdk:"enabled" json:"enabled,omitempty"`
					State    *string   `tfsdk:"state" json:"state,omitempty"`
				} `tfsdk:"types" json:"types,omitempty"`
			} `tfsdk:"developer_catalog" json:"developerCatalog,omitempty"`
			DocumentationBaseURL *string `tfsdk:"documentation_base_url" json:"documentationBaseURL,omitempty"`
			Perspectives         *[]struct {
				Id              *string `tfsdk:"id" json:"id,omitempty"`
				PinnedResources *[]struct {
					Group    *string `tfsdk:"group" json:"group,omitempty"`
					Resource *string `tfsdk:"resource" json:"resource,omitempty"`
					Version  *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"pinned_resources" json:"pinnedResources,omitempty"`
				Visibility *struct {
					AccessReview *struct {
						Missing *[]struct {
							Group       *string `tfsdk:"group" json:"group,omitempty"`
							Name        *string `tfsdk:"name" json:"name,omitempty"`
							Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
							Resource    *string `tfsdk:"resource" json:"resource,omitempty"`
							Subresource *string `tfsdk:"subresource" json:"subresource,omitempty"`
							Verb        *string `tfsdk:"verb" json:"verb,omitempty"`
							Version     *string `tfsdk:"version" json:"version,omitempty"`
						} `tfsdk:"missing" json:"missing,omitempty"`
						Required *[]struct {
							Group       *string `tfsdk:"group" json:"group,omitempty"`
							Name        *string `tfsdk:"name" json:"name,omitempty"`
							Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
							Resource    *string `tfsdk:"resource" json:"resource,omitempty"`
							Subresource *string `tfsdk:"subresource" json:"subresource,omitempty"`
							Verb        *string `tfsdk:"verb" json:"verb,omitempty"`
							Version     *string `tfsdk:"version" json:"version,omitempty"`
						} `tfsdk:"required" json:"required,omitempty"`
					} `tfsdk:"access_review" json:"accessReview,omitempty"`
					State *string `tfsdk:"state" json:"state,omitempty"`
				} `tfsdk:"visibility" json:"visibility,omitempty"`
			} `tfsdk:"perspectives" json:"perspectives,omitempty"`
			ProjectAccess *struct {
				AvailableClusterRoles *[]string `tfsdk:"available_cluster_roles" json:"availableClusterRoles,omitempty"`
			} `tfsdk:"project_access" json:"projectAccess,omitempty"`
			QuickStarts *struct {
				Disabled *[]string `tfsdk:"disabled" json:"disabled,omitempty"`
			} `tfsdk:"quick_starts" json:"quickStarts,omitempty"`
		} `tfsdk:"customization" json:"customization,omitempty"`
		LogLevel         *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
		ManagementState  *string            `tfsdk:"management_state" json:"managementState,omitempty"`
		ObservedConfig   *map[string]string `tfsdk:"observed_config" json:"observedConfig,omitempty"`
		OperatorLogLevel *string            `tfsdk:"operator_log_level" json:"operatorLogLevel,omitempty"`
		Plugins          *[]string          `tfsdk:"plugins" json:"plugins,omitempty"`
		Providers        *struct {
			Statuspage *struct {
				PageID *string `tfsdk:"page_id" json:"pageID,omitempty"`
			} `tfsdk:"statuspage" json:"statuspage,omitempty"`
		} `tfsdk:"providers" json:"providers,omitempty"`
		Route *struct {
			Hostname *string `tfsdk:"hostname" json:"hostname,omitempty"`
			Secret   *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"secret" json:"secret,omitempty"`
		} `tfsdk:"route" json:"route,omitempty"`
		UnsupportedConfigOverrides *map[string]string `tfsdk:"unsupported_config_overrides" json:"unsupportedConfigOverrides,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorOpenshiftIoConsoleV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_openshift_io_console_v1_manifest"
}

func (r *OperatorOpenshiftIoConsoleV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Console provides a means to configure an operator to manage the console.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "Console provides a means to configure an operator to manage the console.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
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
				Description:         "ConsoleSpec is the specification of the desired behavior of the Console.",
				MarkdownDescription: "ConsoleSpec is the specification of the desired behavior of the Console.",
				Attributes: map[string]schema.Attribute{
					"customization": schema.SingleNestedAttribute{
						Description:         "customization is used to optionally provide a small set of customization options to the web console.",
						MarkdownDescription: "customization is used to optionally provide a small set of customization options to the web console.",
						Attributes: map[string]schema.Attribute{
							"add_page": schema.SingleNestedAttribute{
								Description:         "addPage allows customizing actions on the Add page in developer perspective.",
								MarkdownDescription: "addPage allows customizing actions on the Add page in developer perspective.",
								Attributes: map[string]schema.Attribute{
									"disabled_actions": schema.ListAttribute{
										Description:         "disabledActions is a list of actions that are not shown to users. Each action in the list is represented by its ID.",
										MarkdownDescription: "disabledActions is a list of actions that are not shown to users. Each action in the list is represented by its ID.",
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

							"brand": schema.StringAttribute{
								Description:         "brand is the default branding of the web console which can be overridden by providing the brand field.  There is a limited set of specific brand options. This field controls elements of the console such as the logo. Invalid value will prevent a console rollout.",
								MarkdownDescription: "brand is the default branding of the web console which can be overridden by providing the brand field.  There is a limited set of specific brand options. This field controls elements of the console such as the logo. Invalid value will prevent a console rollout.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("openshift", "okd", "online", "ocp", "dedicated", "azure", "OpenShift", "OKD", "Online", "OCP", "Dedicated", "Azure", "ROSA"),
								},
							},

							"custom_logo_file": schema.SingleNestedAttribute{
								Description:         "customLogoFile replaces the default OpenShift logo in the masthead and about dialog. It is a reference to a ConfigMap in the openshift-config namespace. This can be created with a command like 'oc create configmap custom-logo --from-file=/path/to/file -n openshift-config'. Image size must be less than 1 MB due to constraints on the ConfigMap size. The ConfigMap key should include a file extension so that the console serves the file with the correct MIME type. Recommended logo specifications: Dimensions: Max height of 68px and max width of 200px SVG format preferred",
								MarkdownDescription: "customLogoFile replaces the default OpenShift logo in the masthead and about dialog. It is a reference to a ConfigMap in the openshift-config namespace. This can be created with a command like 'oc create configmap custom-logo --from-file=/path/to/file -n openshift-config'. Image size must be less than 1 MB due to constraints on the ConfigMap size. The ConfigMap key should include a file extension so that the console serves the file with the correct MIME type. Recommended logo specifications: Dimensions: Max height of 68px and max width of 200px SVG format preferred",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "Key allows pointing to a specific key/value inside of the configmap.  This is useful for logical file references.",
										MarkdownDescription: "Key allows pointing to a specific key/value inside of the configmap.  This is useful for logical file references.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"custom_product_name": schema.StringAttribute{
								Description:         "customProductName is the name that will be displayed in page titles, logo alt text, and the about dialog instead of the normal OpenShift product name.",
								MarkdownDescription: "customProductName is the name that will be displayed in page titles, logo alt text, and the about dialog instead of the normal OpenShift product name.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"developer_catalog": schema.SingleNestedAttribute{
								Description:         "developerCatalog allows to configure the shown developer catalog categories (filters) and types (sub-catalogs).",
								MarkdownDescription: "developerCatalog allows to configure the shown developer catalog categories (filters) and types (sub-catalogs).",
								Attributes: map[string]schema.Attribute{
									"categories": schema.ListNestedAttribute{
										Description:         "categories which are shown in the developer catalog.",
										MarkdownDescription: "categories which are shown in the developer catalog.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Description:         "ID is an identifier used in the URL to enable deep linking in console. ID is required and must have 1-32 URL safe (A-Z, a-z, 0-9, - and _) characters.",
													MarkdownDescription: "ID is an identifier used in the URL to enable deep linking in console. ID is required and must have 1-32 URL safe (A-Z, a-z, 0-9, - and _) characters.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.LengthAtMost(32),
														stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9-_]+$`), ""),
													},
												},

												"label": schema.StringAttribute{
													Description:         "label defines a category display label. It is required and must have 1-64 characters.",
													MarkdownDescription: "label defines a category display label. It is required and must have 1-64 characters.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.LengthAtMost(64),
													},
												},

												"subcategories": schema.ListNestedAttribute{
													Description:         "subcategories defines a list of child categories.",
													MarkdownDescription: "subcategories defines a list of child categories.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"id": schema.StringAttribute{
																Description:         "ID is an identifier used in the URL to enable deep linking in console. ID is required and must have 1-32 URL safe (A-Z, a-z, 0-9, - and _) characters.",
																MarkdownDescription: "ID is an identifier used in the URL to enable deep linking in console. ID is required and must have 1-32 URL safe (A-Z, a-z, 0-9, - and _) characters.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(32),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9-_]+$`), ""),
																},
															},

															"label": schema.StringAttribute{
																Description:         "label defines a category display label. It is required and must have 1-64 characters.",
																MarkdownDescription: "label defines a category display label. It is required and must have 1-64 characters.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(64),
																},
															},

															"tags": schema.ListAttribute{
																Description:         "tags is a list of strings that will match the category. A selected category show all items which has at least one overlapping tag between category and item.",
																MarkdownDescription: "tags is a list of strings that will match the category. A selected category show all items which has at least one overlapping tag between category and item.",
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

												"tags": schema.ListAttribute{
													Description:         "tags is a list of strings that will match the category. A selected category show all items which has at least one overlapping tag between category and item.",
													MarkdownDescription: "tags is a list of strings that will match the category. A selected category show all items which has at least one overlapping tag between category and item.",
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

									"types": schema.SingleNestedAttribute{
										Description:         "types allows enabling or disabling of sub-catalog types that user can see in the Developer catalog. When omitted, all the sub-catalog types will be shown.",
										MarkdownDescription: "types allows enabling or disabling of sub-catalog types that user can see in the Developer catalog. When omitted, all the sub-catalog types will be shown.",
										Attributes: map[string]schema.Attribute{
											"disabled": schema.ListAttribute{
												Description:         "disabled is a list of developer catalog types (sub-catalogs IDs) that are not shown to users. Types (sub-catalogs) are added via console plugins, the available types (sub-catalog IDs) are available in the console on the cluster configuration page, or when editing the YAML in the console. Example: 'Devfile', 'HelmChart', 'BuilderImage' If the list is empty or all the available sub-catalog types are added, then the complete developer catalog should be hidden.",
												MarkdownDescription: "disabled is a list of developer catalog types (sub-catalogs IDs) that are not shown to users. Types (sub-catalogs) are added via console plugins, the available types (sub-catalog IDs) are available in the console on the cluster configuration page, or when editing the YAML in the console. Example: 'Devfile', 'HelmChart', 'BuilderImage' If the list is empty or all the available sub-catalog types are added, then the complete developer catalog should be hidden.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.ListAttribute{
												Description:         "enabled is a list of developer catalog types (sub-catalogs IDs) that will be shown to users. Types (sub-catalogs) are added via console plugins, the available types (sub-catalog IDs) are available in the console on the cluster configuration page, or when editing the YAML in the console. Example: 'Devfile', 'HelmChart', 'BuilderImage' If the list is non-empty, a new type will not be shown to the user until it is added to list. If the list is empty the complete developer catalog will be shown.",
												MarkdownDescription: "enabled is a list of developer catalog types (sub-catalogs IDs) that will be shown to users. Types (sub-catalogs) are added via console plugins, the available types (sub-catalog IDs) are available in the console on the cluster configuration page, or when editing the YAML in the console. Example: 'Devfile', 'HelmChart', 'BuilderImage' If the list is non-empty, a new type will not be shown to the user until it is added to list. If the list is empty the complete developer catalog will be shown.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"state": schema.StringAttribute{
												Description:         "state defines if a list of catalog types should be enabled or disabled.",
												MarkdownDescription: "state defines if a list of catalog types should be enabled or disabled.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Enabled", "Disabled"),
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

							"documentation_base_url": schema.StringAttribute{
								Description:         "documentationBaseURL links to external documentation are shown in various sections of the web console.  Providing documentationBaseURL will override the default documentation URL. Invalid value will prevent a console rollout.",
								MarkdownDescription: "documentationBaseURL links to external documentation are shown in various sections of the web console.  Providing documentationBaseURL will override the default documentation URL. Invalid value will prevent a console rollout.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^$|^((https):\/\/?)[^\s()<>]+(?:\([\w\d]+\)|([^[:punct:]\s]|\/?))\/$`), ""),
								},
							},

							"perspectives": schema.ListNestedAttribute{
								Description:         "perspectives allows enabling/disabling of perspective(s) that user can see in the Perspective switcher dropdown.",
								MarkdownDescription: "perspectives allows enabling/disabling of perspective(s) that user can see in the Perspective switcher dropdown.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Description:         "id defines the id of the perspective. Example: 'dev', 'admin'. The available perspective ids can be found in the code snippet section next to the yaml editor. Incorrect or unknown ids will be ignored.",
											MarkdownDescription: "id defines the id of the perspective. Example: 'dev', 'admin'. The available perspective ids can be found in the code snippet section next to the yaml editor. Incorrect or unknown ids will be ignored.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"pinned_resources": schema.ListNestedAttribute{
											Description:         "pinnedResources defines the list of default pinned resources that users will see on the perspective navigation if they have not customized these pinned resources themselves. The list of available Kubernetes resources could be read via 'kubectl api-resources'. The console will also provide a configuration UI and a YAML snippet that will list the available resources that can be pinned to the navigation. Incorrect or unknown resources will be ignored.",
											MarkdownDescription: "pinnedResources defines the list of default pinned resources that users will see on the perspective navigation if they have not customized these pinned resources themselves. The list of available Kubernetes resources could be read via 'kubectl api-resources'. The console will also provide a configuration UI and a YAML snippet that will list the available resources that can be pinned to the navigation. Incorrect or unknown resources will be ignored.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"group": schema.StringAttribute{
														Description:         "group is the API Group of the Resource. Enter empty string for the core group. This value should consist of only lowercase alphanumeric characters, hyphens and periods. Example: '', 'apps', 'build.openshift.io', etc.",
														MarkdownDescription: "group is the API Group of the Resource. Enter empty string for the core group. This value should consist of only lowercase alphanumeric characters, hyphens and periods. Example: '', 'apps', 'build.openshift.io', etc.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
														},
													},

													"resource": schema.StringAttribute{
														Description:         "resource is the type that is being referenced. It is normally the plural form of the resource kind in lowercase. This value should consist of only lowercase alphanumeric characters and hyphens. Example: 'deployments', 'deploymentconfigs', 'pods', etc.",
														MarkdownDescription: "resource is the type that is being referenced. It is normally the plural form of the resource kind in lowercase. This value should consist of only lowercase alphanumeric characters and hyphens. Example: 'deployments', 'deploymentconfigs', 'pods', etc.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
													},

													"version": schema.StringAttribute{
														Description:         "version is the API Version of the Resource. This value should consist of only lowercase alphanumeric characters. Example: 'v1', 'v1beta1', etc.",
														MarkdownDescription: "version is the API Version of the Resource. This value should consist of only lowercase alphanumeric characters. Example: 'v1', 'v1beta1', etc.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]+$`), ""),
														},
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"visibility": schema.SingleNestedAttribute{
											Description:         "visibility defines the state of perspective along with access review checks if needed for that perspective.",
											MarkdownDescription: "visibility defines the state of perspective along with access review checks if needed for that perspective.",
											Attributes: map[string]schema.Attribute{
												"access_review": schema.SingleNestedAttribute{
													Description:         "accessReview defines required and missing access review checks.",
													MarkdownDescription: "accessReview defines required and missing access review checks.",
													Attributes: map[string]schema.Attribute{
														"missing": schema.ListNestedAttribute{
															Description:         "missing defines a list of permission checks. The perspective will only be shown when at least one check fails. When omitted, the access review is skipped and the perspective will not be shown unless it is required to do so based on the configuration of the required access review list.",
															MarkdownDescription: "missing defines a list of permission checks. The perspective will only be shown when at least one check fails. When omitted, the access review is skipped and the perspective will not be shown unless it is required to do so based on the configuration of the required access review list.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"group": schema.StringAttribute{
																		Description:         "Group is the API Group of the Resource.  '*' means all.",
																		MarkdownDescription: "Group is the API Group of the Resource.  '*' means all.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name is the name of the resource being requested for a 'get' or deleted for a 'delete'. '' (empty) means all.",
																		MarkdownDescription: "Name is the name of the resource being requested for a 'get' or deleted for a 'delete'. '' (empty) means all.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "Namespace is the namespace of the action being requested.  Currently, there is no distinction between no namespace and all namespaces '' (empty) is defaulted for LocalSubjectAccessReviews '' (empty) is empty for cluster-scoped resources '' (empty) means 'all' for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview",
																		MarkdownDescription: "Namespace is the namespace of the action being requested.  Currently, there is no distinction between no namespace and all namespaces '' (empty) is defaulted for LocalSubjectAccessReviews '' (empty) is empty for cluster-scoped resources '' (empty) means 'all' for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"resource": schema.StringAttribute{
																		Description:         "Resource is one of the existing resource types.  '*' means all.",
																		MarkdownDescription: "Resource is one of the existing resource types.  '*' means all.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"subresource": schema.StringAttribute{
																		Description:         "Subresource is one of the existing resource types.  '' means none.",
																		MarkdownDescription: "Subresource is one of the existing resource types.  '' means none.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"verb": schema.StringAttribute{
																		Description:         "Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy.  '*' means all.",
																		MarkdownDescription: "Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy.  '*' means all.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"version": schema.StringAttribute{
																		Description:         "Version is the API Version of the Resource.  '*' means all.",
																		MarkdownDescription: "Version is the API Version of the Resource.  '*' means all.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
															Validators: []validator.List{
																listvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("required")),
															},
														},

														"required": schema.ListNestedAttribute{
															Description:         "required defines a list of permission checks. The perspective will only be shown when all checks are successful. When omitted, the access review is skipped and the perspective will not be shown unless it is required to do so based on the configuration of the missing access review list.",
															MarkdownDescription: "required defines a list of permission checks. The perspective will only be shown when all checks are successful. When omitted, the access review is skipped and the perspective will not be shown unless it is required to do so based on the configuration of the missing access review list.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"group": schema.StringAttribute{
																		Description:         "Group is the API Group of the Resource.  '*' means all.",
																		MarkdownDescription: "Group is the API Group of the Resource.  '*' means all.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name is the name of the resource being requested for a 'get' or deleted for a 'delete'. '' (empty) means all.",
																		MarkdownDescription: "Name is the name of the resource being requested for a 'get' or deleted for a 'delete'. '' (empty) means all.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "Namespace is the namespace of the action being requested.  Currently, there is no distinction between no namespace and all namespaces '' (empty) is defaulted for LocalSubjectAccessReviews '' (empty) is empty for cluster-scoped resources '' (empty) means 'all' for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview",
																		MarkdownDescription: "Namespace is the namespace of the action being requested.  Currently, there is no distinction between no namespace and all namespaces '' (empty) is defaulted for LocalSubjectAccessReviews '' (empty) is empty for cluster-scoped resources '' (empty) means 'all' for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"resource": schema.StringAttribute{
																		Description:         "Resource is one of the existing resource types.  '*' means all.",
																		MarkdownDescription: "Resource is one of the existing resource types.  '*' means all.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"subresource": schema.StringAttribute{
																		Description:         "Subresource is one of the existing resource types.  '' means none.",
																		MarkdownDescription: "Subresource is one of the existing resource types.  '' means none.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"verb": schema.StringAttribute{
																		Description:         "Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy.  '*' means all.",
																		MarkdownDescription: "Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy.  '*' means all.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"version": schema.StringAttribute{
																		Description:         "Version is the API Version of the Resource.  '*' means all.",
																		MarkdownDescription: "Version is the API Version of the Resource.  '*' means all.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
															Validators: []validator.List{
																listvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("missing")),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"state": schema.StringAttribute{
													Description:         "state defines the perspective is enabled or disabled or access review check is required.",
													MarkdownDescription: "state defines the perspective is enabled or disabled or access review check is required.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Enabled", "Disabled", "AccessReview"),
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

							"project_access": schema.SingleNestedAttribute{
								Description:         "projectAccess allows customizing the available list of ClusterRoles in the Developer perspective Project access page which can be used by a project admin to specify roles to other users and restrict access within the project. If set, the list will replace the default ClusterRole options.",
								MarkdownDescription: "projectAccess allows customizing the available list of ClusterRoles in the Developer perspective Project access page which can be used by a project admin to specify roles to other users and restrict access within the project. If set, the list will replace the default ClusterRole options.",
								Attributes: map[string]schema.Attribute{
									"available_cluster_roles": schema.ListAttribute{
										Description:         "availableClusterRoles is the list of ClusterRole names that are assignable to users through the project access tab.",
										MarkdownDescription: "availableClusterRoles is the list of ClusterRole names that are assignable to users through the project access tab.",
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

							"quick_starts": schema.SingleNestedAttribute{
								Description:         "quickStarts allows customization of available ConsoleQuickStart resources in console.",
								MarkdownDescription: "quickStarts allows customization of available ConsoleQuickStart resources in console.",
								Attributes: map[string]schema.Attribute{
									"disabled": schema.ListAttribute{
										Description:         "disabled is a list of ConsoleQuickStart resource names that are not shown to users.",
										MarkdownDescription: "disabled is a list of ConsoleQuickStart resource names that are not shown to users.",
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

					"log_level": schema.StringAttribute{
						Description:         "logLevel is an intent based logging for an overall component.  It does not give fine grained control, but it is a simple way to manage coarse grained logging choices that operators have to interpret for their operands.  Valid values are: 'Normal', 'Debug', 'Trace', 'TraceAll'. Defaults to 'Normal'.",
						MarkdownDescription: "logLevel is an intent based logging for an overall component.  It does not give fine grained control, but it is a simple way to manage coarse grained logging choices that operators have to interpret for their operands.  Valid values are: 'Normal', 'Debug', 'Trace', 'TraceAll'. Defaults to 'Normal'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("", "Normal", "Debug", "Trace", "TraceAll"),
						},
					},

					"management_state": schema.StringAttribute{
						Description:         "managementState indicates whether and how the operator should manage the component",
						MarkdownDescription: "managementState indicates whether and how the operator should manage the component",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(Managed|Unmanaged|Force|Removed)$`), ""),
						},
					},

					"observed_config": schema.MapAttribute{
						Description:         "observedConfig holds a sparse config that controller has observed from the cluster state.  It exists in spec because it is an input to the level for the operator",
						MarkdownDescription: "observedConfig holds a sparse config that controller has observed from the cluster state.  It exists in spec because it is an input to the level for the operator",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"operator_log_level": schema.StringAttribute{
						Description:         "operatorLogLevel is an intent based logging for the operator itself.  It does not give fine grained control, but it is a simple way to manage coarse grained logging choices that operators have to interpret for themselves.  Valid values are: 'Normal', 'Debug', 'Trace', 'TraceAll'. Defaults to 'Normal'.",
						MarkdownDescription: "operatorLogLevel is an intent based logging for the operator itself.  It does not give fine grained control, but it is a simple way to manage coarse grained logging choices that operators have to interpret for themselves.  Valid values are: 'Normal', 'Debug', 'Trace', 'TraceAll'. Defaults to 'Normal'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("", "Normal", "Debug", "Trace", "TraceAll"),
						},
					},

					"plugins": schema.ListAttribute{
						Description:         "plugins defines a list of enabled console plugin names.",
						MarkdownDescription: "plugins defines a list of enabled console plugin names.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"providers": schema.SingleNestedAttribute{
						Description:         "providers contains configuration for using specific service providers.",
						MarkdownDescription: "providers contains configuration for using specific service providers.",
						Attributes: map[string]schema.Attribute{
							"statuspage": schema.SingleNestedAttribute{
								Description:         "statuspage contains ID for statuspage.io page that provides status info about.",
								MarkdownDescription: "statuspage contains ID for statuspage.io page that provides status info about.",
								Attributes: map[string]schema.Attribute{
									"page_id": schema.StringAttribute{
										Description:         "pageID is the unique ID assigned by Statuspage for your page. This must be a public page.",
										MarkdownDescription: "pageID is the unique ID assigned by Statuspage for your page. This must be a public page.",
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

					"route": schema.SingleNestedAttribute{
						Description:         "route contains hostname and secret reference that contains the serving certificate. If a custom route is specified, a new route will be created with the provided hostname, under which console will be available. In case of custom hostname uses the default routing suffix of the cluster, the Secret specification for a serving certificate will not be needed. In case of custom hostname points to an arbitrary domain, manual DNS configurations steps are necessary. The default console route will be maintained to reserve the default hostname for console if the custom route is removed. If not specified, default route will be used. DEPRECATED",
						MarkdownDescription: "route contains hostname and secret reference that contains the serving certificate. If a custom route is specified, a new route will be created with the provided hostname, under which console will be available. In case of custom hostname uses the default routing suffix of the cluster, the Secret specification for a serving certificate will not be needed. In case of custom hostname points to an arbitrary domain, manual DNS configurations steps are necessary. The default console route will be maintained to reserve the default hostname for console if the custom route is removed. If not specified, default route will be used. DEPRECATED",
						Attributes: map[string]schema.Attribute{
							"hostname": schema.StringAttribute{
								Description:         "hostname is the desired custom domain under which console will be available.",
								MarkdownDescription: "hostname is the desired custom domain under which console will be available.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret": schema.SingleNestedAttribute{
								Description:         "secret points to secret in the openshift-config namespace that contains custom certificate and key and needs to be created manually by the cluster admin. Referenced Secret is required to contain following key value pairs: - 'tls.crt' - to specifies custom certificate - 'tls.key' - to specifies private key of the custom certificate If the custom hostname uses the default routing suffix of the cluster, the Secret specification for a serving certificate will not be needed.",
								MarkdownDescription: "secret points to secret in the openshift-config namespace that contains custom certificate and key and needs to be created manually by the cluster admin. Referenced Secret is required to contain following key value pairs: - 'tls.crt' - to specifies custom certificate - 'tls.key' - to specifies private key of the custom certificate If the custom hostname uses the default routing suffix of the cluster, the Secret specification for a serving certificate will not be needed.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "name is the metadata.name of the referenced secret",
										MarkdownDescription: "name is the metadata.name of the referenced secret",
										Required:            true,
										Optional:            false,
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

					"unsupported_config_overrides": schema.MapAttribute{
						Description:         "unsupportedConfigOverrides overrides the final configuration that was computed by the operator. Red Hat does not support the use of this field. Misuse of this field could lead to unexpected behavior or conflict with other configuration options. Seek guidance from the Red Hat support before using this field. Use of this property blocks cluster upgrades, it must be removed before upgrading your cluster.",
						MarkdownDescription: "unsupportedConfigOverrides overrides the final configuration that was computed by the operator. Red Hat does not support the use of this field. Misuse of this field could lead to unexpected behavior or conflict with other configuration options. Seek guidance from the Red Hat support before using this field. Use of this property blocks cluster upgrades, it must be removed before upgrading your cluster.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *OperatorOpenshiftIoConsoleV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_openshift_io_console_v1_manifest")

	var model OperatorOpenshiftIoConsoleV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.openshift.io/v1")
	model.Kind = pointer.String("Console")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
