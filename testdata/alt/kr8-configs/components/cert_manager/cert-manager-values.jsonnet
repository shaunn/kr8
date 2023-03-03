local config = std.extVar('kr8');

{
//kr8_spec: {},
  rbac: {
    create: true,
  },

  ingressShim: {
    defaultIssuerName: 'letsencrypt',
    defaultIssuerKind: 'ClusterIssuer',
  },

}
