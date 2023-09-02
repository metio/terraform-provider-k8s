output "manifests" {
  value = {
    "example" = data.k8s_ripsaw_cloudbulldozer_io_benchmark_v1alpha1_manifest.example.yaml
  }
}
