{
  namespace: 'kube-system',
  release_name: 'sealed_secrets',
  kubecfg_gc_enable: true,

  kr8_spec: {
      includes: [
  //      'rbac.jsonnet',
      ],
    },

}
