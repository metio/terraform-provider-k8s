data "k8s_persistent_volume_claim_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    access_modes = ["ReadWriteMany"]
    resources = {
      requests = {
        storage = "5Gi"
      }
    }
    volume_name = "some-volume"
  }
}
