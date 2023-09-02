output "manifests" {
  value = {
    "example" = data.k8s_k8gb_absa_oss_gslb_v1beta1_manifest.example.yaml
  }
}
