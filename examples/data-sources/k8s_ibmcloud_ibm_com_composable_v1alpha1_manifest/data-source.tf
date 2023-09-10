data "k8s_ibmcloud_ibm_com_composable_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    template = {
      apiVersion = "ibmcloud.ibm.com/v1alpha1"
      kind       = "Service"
    }
  }
}
