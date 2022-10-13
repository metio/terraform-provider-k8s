resource "k8s_operator_aquasec_com_aqua_kube_enforcer_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    config = {}
  }
}

resource "k8s_operator_aquasec_com_aqua_kube_enforcer_v1alpha1" "example" {
  metadata = {
    name = "test"
  }
  spec = {
    infra = {
      requirements    = true
      version         = "2022.4"
      service_account = "aqua-kube-enforcer-sa"
    }
    config = {
      gateway_address   = "aqua-gateway:8443"
      cluster_name      = "aqua-secure"
      image_pull_secret = "aqua-registry"
    }
    deploy = {
      replicas = 3
      service  = "ClusterIP"
      image = {
        registry    = "registry.aquasec.com"
        repository  = "kube-enforcer"
        tag         = "<<IMAGE TAG>>"
        pull_policy = "IfNotPresent"
      }
    }
    token = "<<KUBE_ENFORCER_GROUP_TOKEN>>"
    starboard = {
      infra = {
        requirements    = true
        service_account = "starboard-operator"
      }
      config = {
        image_pull_secret = "starboard-registry"
      }
      deploy = {
        replicas = 1
      }
    }
  }
}
