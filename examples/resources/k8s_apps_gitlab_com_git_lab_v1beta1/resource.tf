resource "k8s_apps_gitlab_com_git_lab_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_apps_gitlab_com_git_lab_v1beta1" "example" {
  metadata = {
    name      = "gitlab"
    namespace = "gitlab-system"
  }
  spec = {
    chart = {
      version = "6.4.0"
      values = {
        certmanager = {
          install = false
        }
        global = {
          hosts = {
            domain      = "example.com"
            host_suffix = ""
          }
          ingress = {
            configure_certmanager = false
            tls = {
              secret_name = "gitlab-tls"
            }
          }
        }
      }
    }
  }
}
