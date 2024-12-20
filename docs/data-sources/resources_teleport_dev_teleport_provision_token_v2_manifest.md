---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_resources_teleport_dev_teleport_provision_token_v2_manifest Data Source - terraform-provider-k8s"
subcategory: "resources.teleport.dev"
description: |-
  ProvisionToken is the Schema for the provisiontokens API
---

# k8s_resources_teleport_dev_teleport_provision_token_v2_manifest (Data Source)

ProvisionToken is the Schema for the provisiontokens API

## Example Usage

```terraform
data "k8s_resources_teleport_dev_teleport_provision_token_v2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) ProvisionToken resource definition v2 from Teleport (see [below for nested schema](#nestedatt--spec))

### Read-Only

- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.
- `namespace` (String) Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Optional:

- `allow` (Attributes List) Allow is a list of TokenRules, nodes using this token must match one allow rule to use this token. (see [below for nested schema](#nestedatt--spec--allow))
- `aws_iid_ttl` (String) AWSIIDTTL is the TTL to use for AWS EC2 Instance Identity Documents used to join the cluster with this token.
- `azure` (Attributes) Azure allows the configuration of options specific to the 'azure' join method. (see [below for nested schema](#nestedatt--spec--azure))
- `bot_name` (String) BotName is the name of the bot this token grants access to, if any
- `circleci` (Attributes) CircleCI allows the configuration of options specific to the 'circleci' join method. (see [below for nested schema](#nestedatt--spec--circleci))
- `gcp` (Attributes) GCP allows the configuration of options specific to the 'gcp' join method. (see [below for nested schema](#nestedatt--spec--gcp))
- `github` (Attributes) GitHub allows the configuration of options specific to the 'github' join method. (see [below for nested schema](#nestedatt--spec--github))
- `gitlab` (Attributes) GitLab allows the configuration of options specific to the 'gitlab' join method. (see [below for nested schema](#nestedatt--spec--gitlab))
- `join_method` (String) JoinMethod is the joining method required in order to use this token. Supported joining methods include: azure, circleci, ec2, gcp, github, gitlab, iam, kubernetes, spacelift, token, tpm
- `kubernetes` (Attributes) Kubernetes allows the configuration of options specific to the 'kubernetes' join method. (see [below for nested schema](#nestedatt--spec--kubernetes))
- `roles` (List of String) Roles is a list of roles associated with the token, that will be converted to metadata in the SSH and X509 certificates issued to the user of the token
- `spacelift` (Attributes) Spacelift allows the configuration of options specific to the 'spacelift' join method. (see [below for nested schema](#nestedatt--spec--spacelift))
- `suggested_agent_matcher_labels` (Map of String) SuggestedAgentMatcherLabels is a set of labels to be used by agents to match on resources. When an agent uses this token, the agent should monitor resources that match those labels. For databases, this means adding the labels to 'db_service.resources.labels'. Currently, only node-join scripts create a configuration according to the suggestion.
- `suggested_labels` (Map of String) SuggestedLabels is a set of labels that resources should set when using this token to enroll themselves in the cluster. Currently, only node-join scripts create a configuration according to the suggestion.
- `terraform_cloud` (Attributes) TerraformCloud allows the configuration of options specific to the 'terraform_cloud' join method. (see [below for nested schema](#nestedatt--spec--terraform_cloud))
- `tpm` (Attributes) TPM allows the configuration of options specific to the 'tpm' join method. (see [below for nested schema](#nestedatt--spec--tpm))

<a id="nestedatt--spec--allow"></a>
### Nested Schema for `spec.allow`

Optional:

- `aws_account` (String) AWSAccount is the AWS account ID.
- `aws_arn` (String) AWSARN is used for the IAM join method, the AWS identity of joining nodes must match this ARN. Supports wildcards '*' and '?'.
- `aws_regions` (List of String) AWSRegions is used for the EC2 join method and is a list of AWS regions a node is allowed to join from.
- `aws_role` (String) AWSRole is used for the EC2 join method and is the ARN of the AWS role that the Auth Service will assume in order to call the ec2 API.


<a id="nestedatt--spec--azure"></a>
### Nested Schema for `spec.azure`

Optional:

- `allow` (Attributes List) Allow is a list of Rules, nodes using this token must match one allow rule to use this token. (see [below for nested schema](#nestedatt--spec--azure--allow))

<a id="nestedatt--spec--azure--allow"></a>
### Nested Schema for `spec.azure.allow`

Optional:

- `resource_groups` (List of String)
- `subscription` (String)



<a id="nestedatt--spec--circleci"></a>
### Nested Schema for `spec.circleci`

Optional:

- `allow` (Attributes List) Allow is a list of TokenRules, nodes using this token must match one allow rule to use this token. (see [below for nested schema](#nestedatt--spec--circleci--allow))
- `organization_id` (String)

<a id="nestedatt--spec--circleci--allow"></a>
### Nested Schema for `spec.circleci.allow`

Optional:

- `context_id` (String)
- `project_id` (String)



<a id="nestedatt--spec--gcp"></a>
### Nested Schema for `spec.gcp`

Optional:

- `allow` (Attributes List) Allow is a list of Rules, nodes using this token must match one allow rule to use this token. (see [below for nested schema](#nestedatt--spec--gcp--allow))

<a id="nestedatt--spec--gcp--allow"></a>
### Nested Schema for `spec.gcp.allow`

Optional:

- `locations` (List of String)
- `project_ids` (List of String)
- `service_accounts` (List of String)



<a id="nestedatt--spec--github"></a>
### Nested Schema for `spec.github`

Optional:

- `allow` (Attributes List) Allow is a list of TokenRules, nodes using this token must match one allow rule to use this token. (see [below for nested schema](#nestedatt--spec--github--allow))
- `enterprise_server_host` (String) EnterpriseServerHost allows joining from runners associated with a GitHub Enterprise Server instance. When unconfigured, tokens will be validated against github.com, but when configured to the host of a GHES instance, then the tokens will be validated against host. This value should be the hostname of the GHES instance, and should not include the scheme or a path. The instance must be accessible over HTTPS at this hostname and the certificate must be trusted by the Auth Service.
- `enterprise_slug` (String) EnterpriseSlug allows the slug of a GitHub Enterprise organisation to be included in the expected issuer of the OIDC tokens. This is for compatibility with the 'include_enterprise_slug' option in GHE. This field should be set to the slug of your enterprise if this is enabled. If this is not enabled, then this field must be left empty. This field cannot be specified if 'enterprise_server_host' is specified. See https://docs.github.com/en/enterprise-cloud@latest/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect#customizing-the-issuer-value-for-an-enterprise for more information about customized issuer values.

<a id="nestedatt--spec--github--allow"></a>
### Nested Schema for `spec.github.allow`

Optional:

- `actor` (String)
- `environment` (String)
- `ref` (String)
- `ref_type` (String)
- `repository` (String)
- `repository_owner` (String)
- `sub` (String)
- `workflow` (String)



<a id="nestedatt--spec--gitlab"></a>
### Nested Schema for `spec.gitlab`

Optional:

- `allow` (Attributes List) Allow is a list of TokenRules, nodes using this token must match one allow rule to use this token. (see [below for nested schema](#nestedatt--spec--gitlab--allow))
- `domain` (String) Domain is the domain of your GitLab instance. This will default to 'gitlab.com' - but can be set to the domain of your self-hosted GitLab e.g 'gitlab.example.com'.

<a id="nestedatt--spec--gitlab--allow"></a>
### Nested Schema for `spec.gitlab.allow`

Optional:

- `ci_config_ref_uri` (String)
- `ci_config_sha` (String)
- `deployment_tier` (String)
- `environment` (String)
- `environment_protected` (Boolean)
- `namespace_path` (String)
- `pipeline_source` (String)
- `project_path` (String)
- `project_visibility` (String)
- `ref` (String)
- `ref_protected` (Boolean)
- `ref_type` (String)
- `sub` (String)
- `user_email` (String)
- `user_id` (String)
- `user_login` (String)



<a id="nestedatt--spec--kubernetes"></a>
### Nested Schema for `spec.kubernetes`

Optional:

- `allow` (Attributes List) Allow is a list of Rules, nodes using this token must match one allow rule to use this token. (see [below for nested schema](#nestedatt--spec--kubernetes--allow))
- `static_jwks` (Attributes) StaticJWKS is the configuration specific to the 'static_jwks' type. (see [below for nested schema](#nestedatt--spec--kubernetes--static_jwks))
- `type` (String) Type controls which behavior should be used for validating the Kubernetes Service Account token. Support values: - 'in_cluster' - 'static_jwks' If unset, this defaults to 'in_cluster'.

<a id="nestedatt--spec--kubernetes--allow"></a>
### Nested Schema for `spec.kubernetes.allow`

Optional:

- `service_account` (String)


<a id="nestedatt--spec--kubernetes--static_jwks"></a>
### Nested Schema for `spec.kubernetes.static_jwks`

Optional:

- `jwks` (String)



<a id="nestedatt--spec--spacelift"></a>
### Nested Schema for `spec.spacelift`

Optional:

- `allow` (Attributes List) Allow is a list of Rules, nodes using this token must match one allow rule to use this token. (see [below for nested schema](#nestedatt--spec--spacelift--allow))
- `hostname` (String) Hostname is the hostname of the Spacelift tenant that tokens will originate from. E.g 'example.app.spacelift.io'

<a id="nestedatt--spec--spacelift--allow"></a>
### Nested Schema for `spec.spacelift.allow`

Optional:

- `caller_id` (String)
- `caller_type` (String)
- `scope` (String)
- `space_id` (String)



<a id="nestedatt--spec--terraform_cloud"></a>
### Nested Schema for `spec.terraform_cloud`

Optional:

- `allow` (Attributes List) Allow is a list of Rules, nodes using this token must match one allow rule to use this token. (see [below for nested schema](#nestedatt--spec--terraform_cloud--allow))
- `audience` (String) Audience is the JWT audience as configured in the TFC_WORKLOAD_IDENTITY_AUDIENCE(_$TAG) variable in Terraform Cloud. If unset, defaults to the Teleport cluster name. For example, if 'TFC_WORKLOAD_IDENTITY_AUDIENCE_TELEPORT=foo' is set in Terraform Cloud, this value should be 'foo'. If the variable is set to match the cluster name, it does not need to be set here.
- `hostname` (String) Hostname is the hostname of the Terraform Enterprise instance expected to issue JWTs allowed by this token. This may be unset for regular Terraform Cloud use, in which case it will be assumed to be 'app.terraform.io'. Otherwise, it must both match the 'iss' (issuer) field included in JWTs, and provide standard JWKS endpoints.

<a id="nestedatt--spec--terraform_cloud--allow"></a>
### Nested Schema for `spec.terraform_cloud.allow`

Optional:

- `organization_id` (String)
- `organization_name` (String)
- `project_id` (String)
- `project_name` (String)
- `run_phase` (String)
- `workspace_id` (String)
- `workspace_name` (String)



<a id="nestedatt--spec--tpm"></a>
### Nested Schema for `spec.tpm`

Optional:

- `allow` (Attributes List) Allow is a list of Rules, the presented delegated identity must match one allow rule to permit joining. (see [below for nested schema](#nestedatt--spec--tpm--allow))
- `ekcert_allowed_cas` (List of String) EKCertAllowedCAs is a list of CA certificates that will be used to validate TPM EKCerts. When specified, joining TPMs must present an EKCert signed by one of the specified CAs. TPMs that do not present an EKCert will be not permitted to join. When unspecified, TPMs will be allowed to join with either an EKCert or an EKPubHash.

<a id="nestedatt--spec--tpm--allow"></a>
### Nested Schema for `spec.tpm.allow`

Optional:

- `description` (String)
- `ek_certificate_serial` (String)
- `ek_public_hash` (String)
