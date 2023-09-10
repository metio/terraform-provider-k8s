data "k8s_chaos_mesh_org_azure_chaos_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    action              = "vm-stop"
    resource_group_name = "some-name"
    subscription_id     = "abc"
    vm_name             = "some-vm"
  }
}
