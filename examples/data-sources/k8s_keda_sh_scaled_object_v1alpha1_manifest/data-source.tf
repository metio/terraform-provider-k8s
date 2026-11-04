data "k8s_keda_sh_scaled_object_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }

  spec = {
    scale_target_ref = {
      name = "some-value"
    }
    triggers = []
  }
}
