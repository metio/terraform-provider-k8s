data "k8s_chaos_mesh_org_physical_machine_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    address = "some-address"
  }
}
