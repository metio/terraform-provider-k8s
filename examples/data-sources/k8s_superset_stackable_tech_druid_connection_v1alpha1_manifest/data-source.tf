data "k8s_superset_stackable_tech_druid_connection_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    druid = {
      name = "some-name"
    }
    superset = {
      name = "some-name"
    }
  }
}
