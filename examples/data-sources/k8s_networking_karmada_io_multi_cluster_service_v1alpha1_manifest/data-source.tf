data "k8s_networking_karmada_io_multi_cluster_service_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    types = ["Service"]
  }
}
