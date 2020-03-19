local parseYaml = std.native('parseYaml');
local config = std.extVar('kr8');  // imports the config from params.jsonnet
local cluster = std.extVar('kr8_cluster');  // imports the config from params.jsonnet

// Stream format needs a list...
[
  {
    cluster: cluster,
  },
  {
    config: config,
  },
]
