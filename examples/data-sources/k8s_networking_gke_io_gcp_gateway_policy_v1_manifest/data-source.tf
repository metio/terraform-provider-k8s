data "k8s_networking_gke_io_gcp_gateway_policy_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
