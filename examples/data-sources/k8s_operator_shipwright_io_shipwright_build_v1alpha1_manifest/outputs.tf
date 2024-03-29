output "manifests" {
  value = {
    "example" = data.k8s_operator_shipwright_io_shipwright_build_v1alpha1_manifest.example.yaml
  }
}
