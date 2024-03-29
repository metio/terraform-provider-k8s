data "k8s_azure_microsoft_com_postgre_sql_firewall_rule_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
