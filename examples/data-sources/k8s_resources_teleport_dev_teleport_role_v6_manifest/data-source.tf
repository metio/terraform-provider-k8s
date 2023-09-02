data "k8s_resources_teleport_dev_teleport_role_v6_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
