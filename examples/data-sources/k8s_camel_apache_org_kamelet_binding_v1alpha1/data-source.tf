data "k8s_camel_apache_org_kamelet_binding_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
