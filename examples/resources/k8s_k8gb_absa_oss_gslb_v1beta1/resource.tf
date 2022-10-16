resource "k8s_k8gb_absa_oss_gslb_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_k8gb_absa_oss_gslb_v1beta1" "example" {
  metadata = {
    name = "test-gslb-failover"
    namespace = "test-gslb"
  }
  spec = {
    ingress = {
      rules = [
        {
          host = "failover.test.k8gb.io"
          http = {
            paths = [
              {
                path = "/"
                path_type = "Prefix"
                backend = {
                  service = {
                    name = "frontend-podinfo"
                    port = {
                      name = "http"
                    }
                  }
                }
              }
            ]
          }
        }
      ]
    }
    strategy = {
    primary_geo_tag = "eu-west-1"
    type = "failover"
    }
  }
}
