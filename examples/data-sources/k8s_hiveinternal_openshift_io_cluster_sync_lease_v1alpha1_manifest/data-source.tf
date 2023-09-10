data "k8s_hiveinternal_openshift_io_cluster_sync_lease_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
