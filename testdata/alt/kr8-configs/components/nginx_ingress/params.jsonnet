{
kr8_spec: {
    includes: [
//      'rbac.jsonnet',
    ],
  },
    local config = self,
  namespace: 'kube-system',
  release_name: 'nginx-ingress',
  kubecfg_gc_enable: true,
  annotations: {
    'nginx.ingress.kubernetes.io/force-ssl-redirect': 'true',
  },
//  local home = {
//    test: config.wat
//
//  }
}

