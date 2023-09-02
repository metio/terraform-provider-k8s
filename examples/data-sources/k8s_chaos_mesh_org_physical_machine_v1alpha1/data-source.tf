data "k8s_chaos_mesh_org_physical_machine_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
