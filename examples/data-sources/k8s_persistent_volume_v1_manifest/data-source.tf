data "k8s_persistent_volume_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
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
