data "k8s_organizations_services_k8s_aws_organizational_unit_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
