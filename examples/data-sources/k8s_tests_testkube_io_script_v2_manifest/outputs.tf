output "manifests" {
  value = {
    "example" = data.k8s_tests_testkube_io_script_v2_manifest.example.yaml
  }
}
