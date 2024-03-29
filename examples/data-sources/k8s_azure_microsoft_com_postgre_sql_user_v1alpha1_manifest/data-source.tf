data "k8s_azure_microsoft_com_postgre_sql_user_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
