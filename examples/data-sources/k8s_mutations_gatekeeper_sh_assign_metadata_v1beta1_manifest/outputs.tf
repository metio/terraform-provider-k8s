output "manifests" {
  value = {
    "example" = data.k8s_mutations_gatekeeper_sh_assign_metadata_v1beta1_manifest.example.yaml
  }
}
