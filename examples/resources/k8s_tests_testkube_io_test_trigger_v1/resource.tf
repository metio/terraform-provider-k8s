resource "k8s_tests_testkube_io_test_trigger_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
