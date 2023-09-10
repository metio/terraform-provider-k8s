data "k8s_resources_teleport_dev_teleport_login_rule_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
