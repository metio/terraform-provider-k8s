data "k8s_secscan_quay_redhat_com_image_manifest_vuln_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
