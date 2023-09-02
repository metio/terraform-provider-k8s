output "manifests" {
  value = {
    "example" = data.k8s_expansion_gatekeeper_sh_expansion_template_v1beta1_manifest.example.yaml
  }
}
