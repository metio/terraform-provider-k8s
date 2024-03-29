output "manifests" {
  value = {
    "example" = data.k8s_azure_microsoft_com_my_sqlv_net_rule_v1alpha1_manifest.example.yaml
  }
}
