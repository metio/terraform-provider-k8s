output "manifests" {
  value = {
    "example" = data.k8s_kyverno_io_policy_exception_v2beta1_manifest.example.yaml
  }
}
