data "k8s_route53_services_k8s_aws_hosted_zone_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
