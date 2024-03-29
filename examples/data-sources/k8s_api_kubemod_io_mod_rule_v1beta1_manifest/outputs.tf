output "manifests" {
  value = {
    "example" = data.k8s_api_kubemod_io_mod_rule_v1beta1_manifest.example.yaml
  }
}
