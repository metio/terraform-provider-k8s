resource "k8s_chaos_mesh_org_physical_machine_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    address = "some-address"
  }
}
