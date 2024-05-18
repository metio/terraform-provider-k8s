data "k8s_gateway_nginx_org_nginx_proxy_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {}
}
