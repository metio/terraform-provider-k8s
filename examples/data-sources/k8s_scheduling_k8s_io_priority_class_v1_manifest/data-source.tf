data "k8s_scheduling_k8s_io_priority_class_v1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  value = 100
}
