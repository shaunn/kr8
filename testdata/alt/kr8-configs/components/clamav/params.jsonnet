{
  local config = self,
  kr8_spec: {
    includes: [
      'psp.jsonnet',
      'rbac.jsonnet',
      'clamav.jsonnet',
    ],
  },

  registry: '493203180918.dkr.ecr.%s.amazonaws.com' % [$._cluster.aws_region],
  repo: 'akp/releases/clamav-akp',
  release_name: 'clamav-akp',
  release_ver: '0.104.1-1',
  image: '%s/%s:%s' % [config.registry, config.repo, config.release_ver],
  comp_name: 'clamav',
  namespace: 'kube-system',
  tolerations: [{ operator: 'Exists' }],
  serviceAccount: 'clamav-akp-sa',
  secrets: {},
}
