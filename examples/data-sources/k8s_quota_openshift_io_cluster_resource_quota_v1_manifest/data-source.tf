data "k8s_quota_openshift_io_cluster_resource_quota_v1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
