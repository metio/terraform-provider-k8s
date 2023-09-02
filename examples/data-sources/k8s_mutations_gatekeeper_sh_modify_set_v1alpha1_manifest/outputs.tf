output "manifests" {
  value = {
    "example" = data.k8s_mutations_gatekeeper_sh_modify_set_v1alpha1_manifest.example.yaml
  }
}
