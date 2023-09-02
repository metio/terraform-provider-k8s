output "manifests" {
  value = {
    "example" = data.k8s_secscan_quay_redhat_com_image_manifest_vuln_v1alpha1_manifest.example.yaml
  }
}
