data "k8s_hiveinternal_openshift_io_cluster_sync_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
