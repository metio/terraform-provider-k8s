data "k8s_tests_testkube_io_test_suite_v3" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
