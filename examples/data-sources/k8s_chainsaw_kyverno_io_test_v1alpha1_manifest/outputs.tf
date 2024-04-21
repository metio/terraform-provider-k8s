output "manifests" {
  value = {
    "example" = data.k8s_chainsaw_kyverno_io_test_v1alpha1_manifest.example.yaml
  }
}
