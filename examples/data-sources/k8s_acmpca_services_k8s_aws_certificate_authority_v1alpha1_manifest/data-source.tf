data "k8s_acmpca_services_k8s_aws_certificate_authority_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
