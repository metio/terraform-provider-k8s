resource "k8s_caching_ibm_com_varnish_cluster_v1alpha1" "minimal" {
  metadata = {
    name = "test"
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

resource "k8s_caching_ibm_com_varnish_cluster_v1alpha1" "example" {
  metadata = {
    name = "varnishcluster-sample"
  }
  spec = {
    backend = {
      port = "web"
      selector = {
        app = "nginx"
      }
    }
    replicas = 1
    service = {
      port = 80
    }
    varnish = {
      args = [
        "-p",
        "default_ttl=3600",
        "-p",
        "default_grace=3600",
      ]
    }
    vcl = {
      config_map_name      = "vcl-files"
      entrypoint_file_name = "entrypoint.vcl"
    }
  }
}
