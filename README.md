# titan-meta
titan search meta service


## Build

1. Make sure [​*Go*​](https://golang.org/) (version 1.20) is installed.
2. Use `make` to install TM. TM is installed in the `bin` directory.

## Usage

### Single Node with default ports

You can run `tm-server` directly on your local machine, if you want to connect to TM from outside,
you can let TM listen on the host IP.

```bash
# Set correct HostIP here.
export HostIP="192.168.199.105"

tm-server --name="tm" \
          --data-dir="tm" \
          --client-urls="http://${HostIP}:2379" \
          --peer-urls="http://${HostIP}:2380" \
          --log-file=tm.log
```

Using `curl` to see TM member:

```bash
curl http://${HostIP}:2379/tm/api/v1/members

{
    "members": [
        {
            "name":"tm",
            "member_id":"f62e88a6e81c149",
            "peer_urls": [
                "http://192.168.199.105:2380"
            ],
            "client_urls": [
                "http://192.168.199.105:2379"
            ]
        }
    ]
}
```

A better tool [httpie](https://github.com/jkbrzt/httpie) is recommended:

```bash
http http://${HostIP}:2379/tm/api/v1/members
Access-Control-Allow-Headers: accept, content-type, authorization
Access-Control-Allow-Methods: POST, GET, OPTIONS, PUT, DELETE
Access-Control-Allow-Origin: *
Content-Length: 673
Content-Type: application/json; charset=UTF-8
Date: Thu, 20 Feb 2020 09:49:42 GMT

{
    "members": [
        {
            "client_urls": [
                "http://192.168.199.105:2379"
            ],
            "member_id": "f62e88a6e81c149",
            "name": "tm",
            "peer_urls": [
                "http://192.168.199.105:2380"
            ]
        }
    ]
}
```

### Cluster

As a component of TiKV project, TM needs to run with TiKV to work. The cluster can also include TiDB to provide SQL services. You can refer [Deploy a TiDB Cluster Using TiUP](https://docs.pingcap.com/tidb/stable/production-deployment-using-tiup) or [TiDB in Kubernetes Documentation](https://docs.pingcap.com/tidb-in-kubernetes/stable) for detailed instructions to deploy a cluster.