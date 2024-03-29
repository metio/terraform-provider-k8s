output "manifests" {
  value = {
    "example" = data.k8s_k6_io_test_run_v1alpha1_manifest.example.yaml
  }
}
