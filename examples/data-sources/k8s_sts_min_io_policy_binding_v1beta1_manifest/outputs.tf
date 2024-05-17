output "manifests" {
  value = {
    "example" = data.k8s_sts_min_io_policy_binding_v1beta1_manifest.example.yaml
  }
}
