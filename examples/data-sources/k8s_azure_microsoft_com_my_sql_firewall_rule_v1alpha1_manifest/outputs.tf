output "manifests" {
  value = {
    "example" = data.k8s_azure_microsoft_com_my_sql_firewall_rule_v1alpha1_manifest.example.yaml
  }
}
