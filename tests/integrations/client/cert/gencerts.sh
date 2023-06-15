#!/bin/bash
if ! [[ "$0" =~ "./gencerts.sh"  ]]; then
      echo "must be run from 'cert'"
        exit 255
fi

if ! which cfssl; then
      echo "cfssl is not installed"
        exit 255
fi

cfssl gencert -initca ca-csr.json | cfssljson -bare ca -

# tm-server
echo '{"CN":"tm-server","hosts":[""],"key":{"algo":"rsa","size":2048}}' | cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=server -hostname="localhost,127.0.0.1" - | cfssljson -bare tm-server

# client
echo '{"CN":"client","hosts":[""],"key":{"algo":"rsa","size":2048}}' | cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=client -hostname="" - | cfssljson -bare client
