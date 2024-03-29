data "k8s_caching_ibm_com_varnish_cluster_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    backend = {
      port     = "web"
      selector = {}
    }
    service = {
      port = 8080
    }
    vcl = {
      config_map_name      = "vcl-files"
      entrypoint_file_name = "entrypoint.vcl"
    }
  }
}
