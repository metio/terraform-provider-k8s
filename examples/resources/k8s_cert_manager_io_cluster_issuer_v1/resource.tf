resource "k8s_cert_manager_io_cluster_issuer_v1" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
  spec = {}
}
