resource "k8s_chaos_mesh_org_stress_chaos_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
