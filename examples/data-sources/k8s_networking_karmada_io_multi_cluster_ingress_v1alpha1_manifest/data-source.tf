data "k8s_networking_karmada_io_multi_cluster_ingress_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
