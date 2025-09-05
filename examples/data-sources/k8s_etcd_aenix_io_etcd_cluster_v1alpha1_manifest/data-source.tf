data "k8s_etcd_aenix_io_etcd_cluster_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
