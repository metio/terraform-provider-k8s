output "manifests" {
  value = {
    "example" = data.k8s_ibmcloud_ibm_com_composable_v1alpha1_manifest.example.yaml
  }
}
