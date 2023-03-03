local config = std.extVar('kr8');
local kube = import 'kube.libsonnet';

local clamavSA = kube.ServiceAccount(config.serviceAccount) + {
  metadata+: {
    namespace: config.namespace,
  },
};

local clamavRole = kube.Role(config.release_name) + {
  rules+: [
    {
      apiGroups: [
        '',
      ],
      resources: [
        'pods',
        'daemonsets',
      ],
      verbs: [
        '*',
      ],
    },
  ],
};

local subjectClamav = {
  kind: 'ServiceAccount',
  name: config.serviceAccount,
  namespace: config.namespace,
};

local clamavRoleBinding = kube.RoleBinding(config.release_name) + {
  roleRef_: clamavRole,
  subjects: [subjectClamav],
};

[
  clamavSA,
  clamavRole,
  clamavRoleBinding,
]
