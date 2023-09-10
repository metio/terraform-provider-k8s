data "k8s_gateway_nginx_org_nginx_gateway_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
