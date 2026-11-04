data "k8s_tinkerbell_org_stack_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }

  spec = {
    services = {
      tink_controller = {}
      tink_server     = {}
    }
    version = "some-value"
  }
}
