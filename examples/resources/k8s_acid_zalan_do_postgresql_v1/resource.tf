resource "k8s_acid_zalan_do_postgresql_v1" "big" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
    labels = {
      "test" = "abc"
    }
    annotations = {
      "try" = "this"
    }
  }
  spec = {
    allowed_source_ranges = ["10.0.0.0/8"]
    clone = {
      cluster              = "first-cluster"
      s3_access_key_id     = "123456"
      s3_endpoint          = "s3://eu-central.example.com"
      s3_force_path_style  = true
      s3_secret_access_key = "hunter2 huntering"
      s3_wal_path          = "wals/abc"
      timestamp            = "123"
      uid                  = "abc"
    }
    databases = {
      "yes" = "no"
    }
    docker_image                     = "hub.docker.com/example/image:version"
    enable_replica_connection_pooler = true
    pod_priority_class_name          = "critical"
    service_annotations = {
      "sidecar.inject" = "true"
    }
    tls = {
      ca_secret_name   = "mystery"
      certificate_file = "cert.pem"
      private_key_file = "private.pem"
      secret_name      = "secret"
      ca_file          = "ca.pem"
    }
    additional_volumes  = []
    number_of_instances = 3
    postgresql = {
      version = "13.0"
    }
    team_id = "abc"
    volume = {
      storage_class = "gp3"
      selector = {
        match_labels = {
          "app.kubernetes.io/name" = "some-example"
        }
      }
      size = "17G"
    }
  }
}

resource "k8s_acid_zalan_do_postgresql_v1" "small" {
  metadata = {
    name = "test"
  }
  spec = {
    number_of_instances = 3
    postgresql = {
      version = "13.0"
    }
    team_id = "abc"
    volume = {
      storage_class = "gp3"
      selector = {
        match_labels = {
          "app.kubernetes.io/name" = "some-example"
        }
      }
      size = "17G"
    }
  }
}
