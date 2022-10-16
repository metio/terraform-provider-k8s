output "resources" {
  value = {
    "minimal" = k8s_k8gb_absa_oss_gslb_v1beta1.minimal.yaml
    "example" = k8s_k8gb_absa_oss_gslb_v1beta1.example.yaml
  }
}
