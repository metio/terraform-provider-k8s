resource "k8s_apps_3scale_net_ap_icast_v1alpha1" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
  spec = {

  }
}

resource "k8s_apps_3scale_net_ap_icast_v1alpha1" "example" {
  metadata = {
    name = "test"
  }
  spec = {
    admin_portal_credentials_ref = {
      name = "mysecretname"
    }
  }
}
