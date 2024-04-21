output "manifests" {
  value = {
    "example" = data.k8s_chainsaw_kyverno_io_configuration_v1alpha2_manifest.example.yaml
  }
}
