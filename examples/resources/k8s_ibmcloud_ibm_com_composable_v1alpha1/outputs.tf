output "resources" {
  value = {
    "minimal" = k8s_ibmcloud_ibm_com_composable_v1alpha1.minimal.yaml
    "example" = k8s_ibmcloud_ibm_com_composable_v1alpha1.example.yaml
  }
}
