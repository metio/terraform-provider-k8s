output "manifests" {
  value = {
    "example" = data.k8s_tests_testkube_io_test_suite_v1_manifest.example.yaml
  }
}
