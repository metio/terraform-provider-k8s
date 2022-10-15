resource "k8s_persistent_volume_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_persistent_volume_v1" "example" {
  metadata = {
    name = "terraform-example"
  }
  spec = {
    capacity = {
      storage = "2Gi"
    }
    access_modes = ["ReadWriteMany"]

    vsphere_volume = {
      volume_path = "/absolute/path"
    }
  }
}
