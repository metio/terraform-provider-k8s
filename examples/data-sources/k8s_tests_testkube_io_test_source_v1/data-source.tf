data "k8s_tests_testkube_io_test_source_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
