{
  _cluster+: {
  // comment out any of these for general breakage
    tier: 'dev',
    region_name: 'sfo2',
    cluster_name: 'cursedone',
    cluster_type: 'cursedone',
//    cluster_type: error 'cursedone!!!!!!!',
    dns_domain: 'example.com',

    required_variable: "given", //meaningless

  },
  _components+: {
    sealed_secrets: { path: 'components/sealed_secrets' },
    nginx_ingress: { path: 'components/nginx_ingress' },
    external_dns: { path: 'components/external_dns' },
    cert_manager: { path: 'components/cert_manager' },
    component_fail: { path: 'components/eternally_cursed_dns' },
  },

  external_dns+: {
    extraEnv: {
      CF_API_KEY: 'some_key',
      CF_API_EMAIL: 'admin@example.com',
    },

    provider: 'cloudflare',
//    txtPrefix: 'do',
    domainFilters: [
      'example.com',
      'example.io',
      'bad',
      'example.work',
    ],
  },


    component_fail+:{
//        required_variable: null   // Fails when commented out
//        required_variable: 0   // Fails when commented out/
    },
}
