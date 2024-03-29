data "k8s_azure_microsoft_com_azure_sql_server_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
