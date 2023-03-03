```azure
/usr/local/Cellar/go/1.19.6/libexec/bin/go test -json ./... -v ./... generate | jq -r '.Output'
    

/usr/local/Cellar/go/1.19.6/libexec/bin/go test ./... -v ./... generate
```


```
testdata/
├── .kr8.env.modified
├── .kr8.env_disable
├── .kr8.test.env
├── Testing.md
├── alt
│   └── kr8-configs
│     ├── .gitignore
│     ├── .kr8.env          # Provides .kr8.env file
│     ├── LICENSE
│     ├── README.md
│     ├── Taskfile.yaml
│     ├── Taskfile.yml
│     ├── bin
│     ├── clusters
│     ├── components
│     ├── config.env
│     ├── generated
│     ├── kr8-configs
│     ├── kr8.cfg
│     ├── lib
│     ├── metadata
│     ├── templates
│     └── tmp
└── default                     # No default .kr8.env file, just an alternative
    └── kr8-configs
        ├── .gitignore
        ├── .kr8.env.test
        ├── LICENSE
        ├── README.md
        ├── Taskfile.yaml
        ├── Taskfile.yml
        ├── bin
        ├── clusters
        ├── components
        ├── config.env
        ├── generated
        ├── kr8-configs
        ├── kr8.cfg
        ├── lib
        ├── metadata
        ├── templates
        └── tmp
```