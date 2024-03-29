data "k8s_ibmcloud_ibm_com_composable_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    template = {}
  }
}
