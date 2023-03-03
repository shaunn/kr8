{
local config = self,
kr8_spec: {
    includes: [
//      'rbac.jsonnet',
    ],

  # Evil starts here...
  array_test: ['a','b'],

  },
  release_name: 'eternally_cursed_dns',  // equivalent of name used for helm install --name ...
  namespace: 'eternally_cursed_dns',
  kubecfg_gc_enable: true,

  provider: 'aws',
  // must be a list
  domainFilters: [],
  txtOwnerId: std.format('%s', $._cluster.cluster_name),
//    local home = {
//      test: config.wat
//
//    }


# ... and here
local expecting_array = std.all(config.array_test),
//    result:  " d " + std.toString(error "ddd")
//    result:  eval("d")

//  expression: 3 < 1,
  expression: 3 > 1,
//  expression: error "'{[[false]]}'",
  required_variable: error "This is a required variable",

  local test = if config.required_variable == 0 then 1/0,
}
