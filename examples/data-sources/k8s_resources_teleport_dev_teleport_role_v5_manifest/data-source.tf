data "k8s_resources_teleport_dev_teleport_role_v5_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
