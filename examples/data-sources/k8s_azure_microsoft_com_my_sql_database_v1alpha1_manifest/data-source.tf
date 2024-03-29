data "k8s_azure_microsoft_com_my_sql_database_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
