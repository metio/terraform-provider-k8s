data "k8s_oracle_db_anthosapis_com_pitr_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
