data "k8s_ec2_services_k8s_aws_dhcp_options_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
