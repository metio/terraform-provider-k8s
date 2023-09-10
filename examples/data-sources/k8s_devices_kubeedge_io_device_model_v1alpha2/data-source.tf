data "k8s_devices_kubeedge_io_device_model_v1alpha2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
