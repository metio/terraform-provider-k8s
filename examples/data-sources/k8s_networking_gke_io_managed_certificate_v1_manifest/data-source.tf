data "k8s_networking_gke_io_managed_certificate_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
