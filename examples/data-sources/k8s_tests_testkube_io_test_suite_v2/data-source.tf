data "k8s_tests_testkube_io_test_suite_v2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
