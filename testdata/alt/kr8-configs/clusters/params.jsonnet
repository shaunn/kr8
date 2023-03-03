{
  _kr8_spec: {
//    postprocessor: "(import 'postprocess-v2.jsonnet').postprocess",  // must return a function that processes a list of objects
    generate_dir: 'generated',
    generate_short_names: true,
  },
  _cluster: {

    // Can/should debug level be linked to env?
    // should errors here be noted and moved on?
    // Global Defaults and Mandatory Parameters
    cluster_name: error '"cluster_name" must be set to the cluster name',
    cluster_type: error '"cluster_type" must be set to the cluster type',
    region_name: error '"region_name" must be set to the region name',
    tier: error '"tier" must be set to the tier name',

    aws_region: if $._cluster.cluster_type == 'aws' then error 'aws_region must be set for aws clusters' else null,
  },
} +
{
  // Components
  _components: {
    sealed_secrets: { path: 'components/sealed_secrets' },
    clamav: { path: 'components/clamav' },

  },
    // Components for AWS clusters
}
