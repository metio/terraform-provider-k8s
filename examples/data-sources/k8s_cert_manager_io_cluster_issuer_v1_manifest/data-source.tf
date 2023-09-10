data "k8s_cert_manager_io_cluster_issuer_v1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {}
}
