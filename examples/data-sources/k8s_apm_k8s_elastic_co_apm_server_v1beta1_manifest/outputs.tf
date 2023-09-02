output "manifests" {
  value = {
    "example" = data.k8s_apm_k8s_elastic_co_apm_server_v1beta1_manifest.example.yaml
  }
}
