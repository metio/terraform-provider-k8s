data "k8s_multicluster_crd_antrea_io_cluster_set_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
