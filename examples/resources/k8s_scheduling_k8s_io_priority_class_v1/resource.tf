resource "k8s_scheduling_k8s_io_priority_class_v1" "minimal" {
  metadata = {
    name = "test"
  }
  value = -100
}

resource "k8s_scheduling_k8s_io_priority_class_v1" "example" {
  metadata = {
    name = "test"
  }
  value = 100
}
