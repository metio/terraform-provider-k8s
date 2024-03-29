data "k8s_azure_microsoft_com_my_sql_server_administrator_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
