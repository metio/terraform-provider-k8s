data "k8s_tests_testkube_io_test_suite_execution_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
