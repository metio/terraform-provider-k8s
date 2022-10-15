resource "k8s_persistent_volume_claim_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_persistent_volume_claim_v1" "example" {
  metadata = {
    name = "terraform-example"
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
