---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_hive_openshift_io_cluster_pool_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "hive.openshift.io"
description: |-
  ClusterPool represents a pool of clusters that should be kept ready to be given out to users. Clusters are removed from the pool once claimed and then automatically replaced with a new one.
---

# k8s_hive_openshift_io_cluster_pool_v1_manifest (Data Source)

ClusterPool represents a pool of clusters that should be kept ready to be given out to users. Clusters are removed from the pool once claimed and then automatically replaced with a new one.

## Example Usage

```terraform
data "k8s_hive_openshift_io_cluster_pool_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    base_domain = "example.com"
    image_set_ref = {
      name = "some-image-set"
    }
    size     = 123
    platform = {}
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) ClusterPoolSpec defines the desired state of the ClusterPool. (see [below for nested schema](#nestedatt--spec))

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

Required:

- `base_domain` (String) BaseDomain is the base domain to use for all clusters created in this pool.
- `image_set_ref` (Attributes) ImageSetRef is a reference to a ClusterImageSet. The release image specified in the ClusterImageSet will be used by clusters created for this cluster pool. (see [below for nested schema](#nestedatt--spec--image_set_ref))
- `platform` (Attributes) Platform encompasses the desired platform for the cluster. (see [below for nested schema](#nestedatt--spec--platform))
- `size` (Number) Size is the default number of clusters that we should keep provisioned and waiting for use.

Optional:

- `annotations` (Map of String) Annotations to be applied to new ClusterDeployments created for the pool. ClusterDeployments that have already been claimed will not be affected when this value is modified.
- `claim_lifetime` (Attributes) ClaimLifetime defines the lifetimes for claims for the cluster pool. (see [below for nested schema](#nestedatt--spec--claim_lifetime))
- `hibernate_after` (String) HibernateAfter will be applied to new ClusterDeployments created for the pool. HibernateAfter will transition clusters in the clusterpool to hibernating power state after it has been running for the given duration. The time that a cluster has been running is the time since the cluster was installed or the time since the cluster last came out of hibernation. This is a Duration value; see https://pkg.go.dev/time#ParseDuration for accepted formats. Note: due to discrepancies in validation vs parsing, we use a Pattern instead of 'Format=duration'. See https://bugzilla.redhat.com/show_bug.cgi?id=2050332 https://github.com/kubernetes/apimachinery/issues/131 https://github.com/kubernetes/apiextensions-apiserver/issues/56
- `hibernation_config` (Attributes) HibernationConfig configures the hibernation/resume behavior of ClusterDeployments owned by the ClusterPool. (see [below for nested schema](#nestedatt--spec--hibernation_config))
- `install_attempts_limit` (Number) InstallAttemptsLimit is the maximum number of times Hive will attempt to install the cluster.
- `install_config_secret_template_ref` (Attributes) InstallConfigSecretTemplateRef is a secret with the key install-config.yaml consisting of the content of the install-config.yaml to be used as a template for all clusters in this pool. Cluster specific settings (name, basedomain) will be injected dynamically when the ClusterDeployment install-config Secret is generated. (see [below for nested schema](#nestedatt--spec--install_config_secret_template_ref))
- `inventory` (Attributes List) Inventory maintains a list of entries consumed by the ClusterPool to customize the default ClusterDeployment. (see [below for nested schema](#nestedatt--spec--inventory))
- `labels` (Map of String) Labels to be applied to new ClusterDeployments created for the pool. ClusterDeployments that have already been claimed will not be affected when this value is modified.
- `max_concurrent` (Number) MaxConcurrent is the maximum number of clusters that will be provisioned or deprovisioned at an time. This includes the claimed clusters being deprovisioned. By default there is no limit.
- `max_size` (Number) MaxSize is the maximum number of clusters that will be provisioned including clusters that have been claimed and ones waiting to be used. By default there is no limit.
- `pull_secret_ref` (Attributes) PullSecretRef is the reference to the secret to use when pulling images. (see [below for nested schema](#nestedatt--spec--pull_secret_ref))
- `running_count` (Number) RunningCount is the number of clusters we should keep running. The remainder will be kept hibernated until claimed. By default no clusters will be kept running (all will be hibernated).
- `skip_machine_pools` (Boolean) SkipMachinePools allows creating clusterpools where the machinepools are not managed by hive after cluster creation

<a id="nestedatt--spec--image_set_ref"></a>
### Nested Schema for `spec.image_set_ref`

Required:

- `name` (String) Name is the name of the ClusterImageSet that this refers to


<a id="nestedatt--spec--platform"></a>
### Nested Schema for `spec.platform`

Optional:

- `agent_bare_metal` (Attributes) AgentBareMetal is the configuration used when performing an Assisted Agent based installation to bare metal. (see [below for nested schema](#nestedatt--spec--platform--agent_bare_metal))
- `aws` (Attributes) AWS is the configuration used when installing on AWS. (see [below for nested schema](#nestedatt--spec--platform--aws))
- `azure` (Attributes) Azure is the configuration used when installing on Azure. (see [below for nested schema](#nestedatt--spec--platform--azure))
- `baremetal` (Attributes) BareMetal is the configuration used when installing on bare metal. (see [below for nested schema](#nestedatt--spec--platform--baremetal))
- `gcp` (Attributes) GCP is the configuration used when installing on Google Cloud Platform. (see [below for nested schema](#nestedatt--spec--platform--gcp))
- `ibmcloud` (Attributes) IBMCloud is the configuration used when installing on IBM Cloud (see [below for nested schema](#nestedatt--spec--platform--ibmcloud))
- `none` (Map of String) None indicates platform-agnostic install. https://docs.openshift.com/container-platform/4.7/installing/installing_platform_agnostic/installing-platform-agnostic.html
- `openstack` (Attributes) OpenStack is the configuration used when installing on OpenStack (see [below for nested schema](#nestedatt--spec--platform--openstack))
- `ovirt` (Attributes) Ovirt is the configuration used when installing on oVirt (see [below for nested schema](#nestedatt--spec--platform--ovirt))
- `vsphere` (Attributes) VSphere is the configuration used when installing on vSphere (see [below for nested schema](#nestedatt--spec--platform--vsphere))

<a id="nestedatt--spec--platform--agent_bare_metal"></a>
### Nested Schema for `spec.platform.agent_bare_metal`

Required:

- `agent_selector` (Attributes) AgentSelector is a label selector used for associating relevant custom resources with this cluster. (Agent, BareMetalHost, etc) (see [below for nested schema](#nestedatt--spec--platform--agent_bare_metal--agent_selector))

<a id="nestedatt--spec--platform--agent_bare_metal--agent_selector"></a>
### Nested Schema for `spec.platform.agent_bare_metal.agent_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--platform--agent_bare_metal--agent_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--platform--agent_bare_metal--agent_selector--match_expressions"></a>
### Nested Schema for `spec.platform.agent_bare_metal.agent_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.




<a id="nestedatt--spec--platform--aws"></a>
### Nested Schema for `spec.platform.aws`

Required:

- `region` (String) Region specifies the AWS region where the cluster will be created.

Optional:

- `credentials_assume_role` (Attributes) CredentialsAssumeRole refers to the IAM role that must be assumed to obtain AWS account access for the cluster operations. (see [below for nested schema](#nestedatt--spec--platform--aws--credentials_assume_role))
- `credentials_secret_ref` (Attributes) CredentialsSecretRef refers to a secret that contains the AWS account access credentials. (see [below for nested schema](#nestedatt--spec--platform--aws--credentials_secret_ref))
- `private_link` (Attributes) PrivateLink allows uses to enable access to the cluster's API server using AWS PrivateLink. AWS PrivateLink includes a pair of VPC Endpoint Service and VPC Endpoint accross AWS accounts and allows clients to connect to services using AWS's internal networking instead of the Internet. (see [below for nested schema](#nestedatt--spec--platform--aws--private_link))
- `user_tags` (Map of String) UserTags specifies additional tags for AWS resources created for the cluster.

<a id="nestedatt--spec--platform--aws--credentials_assume_role"></a>
### Nested Schema for `spec.platform.aws.credentials_assume_role`

Required:

- `role_arn` (String)

Optional:

- `external_id` (String) ExternalID is random string generated by platform so that assume role is protected from confused deputy problem. more info: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html


<a id="nestedatt--spec--platform--aws--credentials_secret_ref"></a>
### Nested Schema for `spec.platform.aws.credentials_secret_ref`

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?


<a id="nestedatt--spec--platform--aws--private_link"></a>
### Nested Schema for `spec.platform.aws.private_link`

Required:

- `enabled` (Boolean)

Optional:

- `additional_allowed_principals` (List of String) AdditionalAllowedPrincipals is a list of additional allowed principal ARNs to be configured for the Private Link cluster's VPC Endpoint Service. ARNs provided as AdditionalAllowedPrincipals will be configured for the cluster's VPC Endpoint Service in addition to the IAM entity used by Hive.



<a id="nestedatt--spec--platform--azure"></a>
### Nested Schema for `spec.platform.azure`

Required:

- `credentials_secret_ref` (Attributes) CredentialsSecretRef refers to a secret that contains the Azure account access credentials. (see [below for nested schema](#nestedatt--spec--platform--azure--credentials_secret_ref))
- `region` (String) Region specifies the Azure region where the cluster will be created.

Optional:

- `base_domain_resource_group_name` (String) BaseDomainResourceGroupName specifies the resource group where the azure DNS zone for the base domain is found
- `cloud_name` (String) cloudName is the name of the Azure cloud environment which can be used to configure the Azure SDK with the appropriate Azure API endpoints. If empty, the value is equal to 'AzurePublicCloud'.

<a id="nestedatt--spec--platform--azure--credentials_secret_ref"></a>
### Nested Schema for `spec.platform.azure.credentials_secret_ref`

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?



<a id="nestedatt--spec--platform--baremetal"></a>
### Nested Schema for `spec.platform.baremetal`

Required:

- `libvirt_ssh_private_key_secret_ref` (Attributes) LibvirtSSHPrivateKeySecretRef is the reference to the secret that contains the private SSH key to use for access to the libvirt provisioning host. The SSH private key is expected to be in the secret data under the 'ssh-privatekey' key. (see [below for nested schema](#nestedatt--spec--platform--baremetal--libvirt_ssh_private_key_secret_ref))

<a id="nestedatt--spec--platform--baremetal--libvirt_ssh_private_key_secret_ref"></a>
### Nested Schema for `spec.platform.baremetal.libvirt_ssh_private_key_secret_ref`

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?



<a id="nestedatt--spec--platform--gcp"></a>
### Nested Schema for `spec.platform.gcp`

Required:

- `credentials_secret_ref` (Attributes) CredentialsSecretRef refers to a secret that contains the GCP account access credentials. (see [below for nested schema](#nestedatt--spec--platform--gcp--credentials_secret_ref))
- `region` (String) Region specifies the GCP region where the cluster will be created.

<a id="nestedatt--spec--platform--gcp--credentials_secret_ref"></a>
### Nested Schema for `spec.platform.gcp.credentials_secret_ref`

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?



<a id="nestedatt--spec--platform--ibmcloud"></a>
### Nested Schema for `spec.platform.ibmcloud`

Required:

- `credentials_secret_ref` (Attributes) CredentialsSecretRef refers to a secret that contains IBM Cloud account access credentials. (see [below for nested schema](#nestedatt--spec--platform--ibmcloud--credentials_secret_ref))
- `region` (String) Region specifies the IBM Cloud region where the cluster will be created.

Optional:

- `account_id` (String) AccountID is the IBM Cloud Account ID. AccountID is DEPRECATED and is gathered via the IBM Cloud API for the provided credentials. This field will be ignored.
- `cis_instance_crn` (String) CISInstanceCRN is the IBM Cloud Internet Services Instance CRN CISInstanceCRN is DEPRECATED and gathered via the IBM Cloud API for the provided credentials and cluster deployment base domain. This field will be ignored.

<a id="nestedatt--spec--platform--ibmcloud--credentials_secret_ref"></a>
### Nested Schema for `spec.platform.ibmcloud.credentials_secret_ref`

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?



<a id="nestedatt--spec--platform--openstack"></a>
### Nested Schema for `spec.platform.openstack`

Required:

- `cloud` (String) Cloud will be used to indicate the OS_CLOUD value to use the right section from the clouds.yaml in the CredentialsSecretRef.
- `credentials_secret_ref` (Attributes) CredentialsSecretRef refers to a secret that contains the OpenStack account access credentials. (see [below for nested schema](#nestedatt--spec--platform--openstack--credentials_secret_ref))

Optional:

- `certificates_secret_ref` (Attributes) CertificatesSecretRef refers to a secret that contains CA certificates necessary for communicating with the OpenStack. There is additional configuration required for the OpenShift cluster to trust the certificates provided in this secret. The 'clouds.yaml' file included in the credentialsSecretRef Secret must also include a reference to the certificate bundle file for the OpenShift cluster being created to trust the OpenStack endpoints. The 'clouds.yaml' file must set the 'cacert' field to either '/etc/openstack-ca/<key name containing the trust bundle in credentialsSecretRef Secret>' or '/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem'.  For example, '''clouds.yaml clouds: shiftstack: auth: ... cacert: '/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem' ''' (see [below for nested schema](#nestedatt--spec--platform--openstack--certificates_secret_ref))
- `trunk_support` (Boolean) TrunkSupport indicates whether or not to use trunk ports in your OpenShift cluster.

<a id="nestedatt--spec--platform--openstack--credentials_secret_ref"></a>
### Nested Schema for `spec.platform.openstack.credentials_secret_ref`

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?


<a id="nestedatt--spec--platform--openstack--certificates_secret_ref"></a>
### Nested Schema for `spec.platform.openstack.certificates_secret_ref`

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?



<a id="nestedatt--spec--platform--ovirt"></a>
### Nested Schema for `spec.platform.ovirt`

Required:

- `certificates_secret_ref` (Attributes) CertificatesSecretRef refers to a secret that contains the oVirt CA certificates necessary for communicating with oVirt. (see [below for nested schema](#nestedatt--spec--platform--ovirt--certificates_secret_ref))
- `credentials_secret_ref` (Attributes) CredentialsSecretRef refers to a secret that contains the oVirt account access credentials with fields: ovirt_url, ovirt_username, ovirt_password, ovirt_ca_bundle (see [below for nested schema](#nestedatt--spec--platform--ovirt--credentials_secret_ref))
- `ovirt_cluster_id` (String) The target cluster under which all VMs will run
- `storage_domain_id` (String) The target storage domain under which all VM disk would be created.

Optional:

- `ovirt_network_name` (String) The target network of all the network interfaces of the nodes. Omitting defaults to ovirtmgmt network which is a default network for evert ovirt cluster.

<a id="nestedatt--spec--platform--ovirt--certificates_secret_ref"></a>
### Nested Schema for `spec.platform.ovirt.certificates_secret_ref`

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?


<a id="nestedatt--spec--platform--ovirt--credentials_secret_ref"></a>
### Nested Schema for `spec.platform.ovirt.credentials_secret_ref`

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?



<a id="nestedatt--spec--platform--vsphere"></a>
### Nested Schema for `spec.platform.vsphere`

Required:

- `certificates_secret_ref` (Attributes) CertificatesSecretRef refers to a secret that contains the vSphere CA certificates necessary for communicating with the VCenter. (see [below for nested schema](#nestedatt--spec--platform--vsphere--certificates_secret_ref))
- `credentials_secret_ref` (Attributes) CredentialsSecretRef refers to a secret that contains the vSphere account access credentials: GOVC_USERNAME, GOVC_PASSWORD fields. (see [below for nested schema](#nestedatt--spec--platform--vsphere--credentials_secret_ref))
- `datacenter` (String) Datacenter is the name of the datacenter to use in the vCenter.
- `default_datastore` (String) DefaultDatastore is the default datastore to use for provisioning volumes.
- `v_center` (String) VCenter is the domain name or IP address of the vCenter.

Optional:

- `cluster` (String) Cluster is the name of the cluster virtual machines will be cloned into.
- `folder` (String) Folder is the name of the folder that will be used and/or created for virtual machines.
- `network` (String) Network specifies the name of the network to be used by the cluster.

<a id="nestedatt--spec--platform--vsphere--certificates_secret_ref"></a>
### Nested Schema for `spec.platform.vsphere.certificates_secret_ref`

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?


<a id="nestedatt--spec--platform--vsphere--credentials_secret_ref"></a>
### Nested Schema for `spec.platform.vsphere.credentials_secret_ref`

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?




<a id="nestedatt--spec--claim_lifetime"></a>
### Nested Schema for `spec.claim_lifetime`

Optional:

- `default` (String) Default is the default lifetime of the claim when no lifetime is set on the claim itself. This is a Duration value; see https://pkg.go.dev/time#ParseDuration for accepted formats. Note: due to discrepancies in validation vs parsing, we use a Pattern instead of 'Format=duration'. See https://bugzilla.redhat.com/show_bug.cgi?id=2050332 https://github.com/kubernetes/apimachinery/issues/131 https://github.com/kubernetes/apiextensions-apiserver/issues/56
- `maximum` (String) Maximum is the maximum lifetime of the claim after it is assigned a cluster. If the claim still exists when the lifetime has elapsed, the claim will be deleted by Hive. The lifetime of a claim is the mimimum of the lifetimes set by the cluster pool and the claim itself. This is a Duration value; see https://pkg.go.dev/time#ParseDuration for accepted formats. Note: due to discrepancies in validation vs parsing, we use a Pattern instead of 'Format=duration'. See https://bugzilla.redhat.com/show_bug.cgi?id=2050332 https://github.com/kubernetes/apimachinery/issues/131 https://github.com/kubernetes/apiextensions-apiserver/issues/56


<a id="nestedatt--spec--hibernation_config"></a>
### Nested Schema for `spec.hibernation_config`

Optional:

- `resume_timeout` (String) ResumeTimeout is the maximum amount of time we will wait for an unclaimed ClusterDeployment to resume from hibernation (e.g. at the behest of runningCount, or in preparation for being claimed). If this time is exceeded, the ClusterDeployment will be considered Broken and we will replace it. The default (unspecified or zero) means no timeout -- we will allow the ClusterDeployment to continue trying to resume 'forever'. This is a Duration value; see https://pkg.go.dev/time#ParseDuration for accepted formats. Note: due to discrepancies in validation vs parsing, we use a Pattern instead of 'Format=duration'. See https://bugzilla.redhat.com/show_bug.cgi?id=2050332 https://github.com/kubernetes/apimachinery/issues/131 https://github.com/kubernetes/apiextensions-apiserver/issues/56


<a id="nestedatt--spec--install_config_secret_template_ref"></a>
### Nested Schema for `spec.install_config_secret_template_ref`

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?


<a id="nestedatt--spec--inventory"></a>
### Nested Schema for `spec.inventory`

Optional:

- `kind` (String) Kind denotes the kind of the referenced resource. The default is ClusterDeploymentCustomization, which is also currently the only supported value.
- `name` (String) Name is the name of the referenced resource.


<a id="nestedatt--spec--pull_secret_ref"></a>
### Nested Schema for `spec.pull_secret_ref`

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?