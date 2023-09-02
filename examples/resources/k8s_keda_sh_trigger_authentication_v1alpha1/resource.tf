resource "k8s_keda_sh_trigger_authentication_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
