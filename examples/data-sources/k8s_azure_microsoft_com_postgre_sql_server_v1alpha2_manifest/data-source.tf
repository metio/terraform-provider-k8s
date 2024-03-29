data "k8s_azure_microsoft_com_postgre_sql_server_v1alpha2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
