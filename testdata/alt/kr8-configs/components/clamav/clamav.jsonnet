local config = std.extVar('kr8');

local daemonSet = {
  kind: 'DaemonSet',
  metadata: {
    name: config.release_name,
    namespace: config.namespace,
  },
  apiVersion: 'apps/v1',
  spec: {
    selector: { matchLabels: { name: config.release_name } },
    updateStrategy: { type: 'RollingUpdate' },
    template: {
      metadata: { labels: { name: config.release_name } },
      spec: {
        serviceAccount: config.serviceAccount,
        tolerations: config.tolerations,
        nodeSelector: {
          'node-role.akp/node': '',
        },
        containers: [{
          name: config.release_name,
          image: config.image,
          imagePullPolicy: 'IfNotPresent',
          resources: {
            requests: { cpu: '0', memory: '0' },
            limits: { cpu: '500m', memory: '1Gi' },
          },
          env: [
            { name: 'AKP_CLAMAV_CRONJOB', value: "0 */4 * * * sh -c 'sleep $((RANDOM%3600))'; clamscan --stdout --infected /host-fs/tmp /host-fs/var/tmp /host-fs/var/spool/mail >> /var/log/clamav/clamscan.log 2>&1" },
            { name: 'CLAMAV_NO_CLAMD', value: 'true' },
            { name: 'CLAMAV_NO_MILTERD', value: 'true' },
          ],
          securityContext: {
            privileged: true,
          },
          volumeMounts: [{
            name: 'host-vol-var',
            mountPath: '/host-fs/var',
            readOnly: true,
          }] + [{
            name: 'host-vol-tmp',
            mountPath: '/host-fs/tmp',
            readOnly: true,
          }] + [{
            name: 'logs',
            mountPath: '/var/log/clamav/',
          }],
        }],
        terminationGracePeriodSeconds: 30,
        volumes: [{
          name: 'host-vol-var',
          hostPath: { path: '/var' },
        }] + [{
          name: 'host-vol-tmp',
          hostPath: { path: '/tmp' },
        }] + [{
          name: 'logs',
          hostPath: { path: '/var/log/clamav', type: 'DirectoryOrCreate' },
        }],
      },
    },
  },
};

[daemonSet]
