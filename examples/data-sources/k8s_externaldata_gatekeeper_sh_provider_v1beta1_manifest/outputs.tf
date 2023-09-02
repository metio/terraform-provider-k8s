output "manifests" {
  value = {
    "example" = data.k8s_externaldata_gatekeeper_sh_provider_v1beta1_manifest.example.yaml
  }
}
