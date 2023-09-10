data "k8s_acid_zalan_do_operator_configuration_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  configuration = {
    enable_crd_registration               = true
    max_instances                         = 7
    min_instances                         = 3
    enable_spilo_wal_path_compat          = true
    workers                               = 5
    docker_image                          = "internal.registry.com/some/image:yes"
    enable_team_id_clustername_prefix     = true
    kubernetes_use_configmaps             = true
    set_memory_request_to_limit           = true
    etcd_host                             = "localhost"
    ignore_instance_limits_annotation_key = "some-key/limit"
    enable_pgversion_env_var              = true
    enable_lazy_spilo_upgrade             = true
    enable_shm_volume                     = true
  }
}
