resource "k8s_ibmcloud_ibm_com_composable_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    template = {}
  }
}

resource "k8s_ibmcloud_ibm_com_composable_v1alpha1" "example" {
  metadata = {
    name = "comp"
  }
  spec = {
    template = {
      apiVersion = "ibmcloud.ibm.com/v1alpha1"
      kind       = "Service"
    }
  }
}
