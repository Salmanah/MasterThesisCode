Usage() {
 echo ""
 echo "Usage: ./createPeerAdminCard.sh [-h host] [-n]"
 echo ""
 echo "Options:"
 echo -e "\t-h or --host:\t\t(Optional) name of the host to specify in the connection profile"
 echo -e "\t-n or --noimport:\t(Optional) don't import into card store"
 echo ""
 echo "Example: ./createPeerAdminCard.sh"
 echo ""
 exit 1
}

Parse_Arguments() {
 while [ $# -gt 0 ]; do
 case $1 in
 --help)
 HELPINFO=true
 ;;
 --host | -h)
 shift
 HOST="$1"
 ;;
 --noimport | -n)
 NOIMPORT=true
 ;;
 esac
 shift
 done
}

HOST=localhost
Parse_Arguments $@

if [ "${HELPINFO}" == "true" ]; then
 Usage
fi

# Grab the current directory
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

if [ -z "${HL_COMPOSER_CLI}" ]; then
 HL_COMPOSER_CLI=$(which composer)
fi

echo
# check that the composer command exists at a version >v0.16
COMPOSER_VERSION=$("${HL_COMPOSER_CLI}" --version 2>/dev/null)
COMPOSER_RC=$?

if [ $COMPOSER_RC -eq 0 ]; then
 AWKRET=$(echo $COMPOSER_VERSION | awk -F. '{if ($2<19) print "1"; else print "0";}')
 if [ $AWKRET -eq 1 ]; then
 echo Cannot use $COMPOSER_VERSION version of composer with fabric 1.1, v0.19 or higher is required
 exit 1
 else
 echo Using composer-cli at $COMPOSER_VERSION
 fi
else
 echo 'No version of composer-cli has been detected, you need to install composer-cli at v0.19 or higher'
 exit 1
fi

cat << EOF > DevServer_connection.json
{
    "name": "hlfv1",
    "x-type": "hlfv1",
    "version": "1.0.0",
    "client": {
        "organization": "Org1",
        "connection": {
            "timeout": {
                "peer": {
                    "endorser": "300",
                    "eventHub": "300",
                    "eventReg": "300"
                },
                "orderer": "300"
            }
        }
    },
    "channels": {
        "mychannel": {
            "orderers": [
                "orderer0.example.com",
                "orderer1.example.com",
                "orderer2.example.com"
            ],
            "peers": {
                "peer0.org1.example.com": {
                    "endorsingPeer": true,
                    "chaincodeQuery": true,
                    "eventSource": true
                },
                "peer1.org1.example.com": {
                    "endorsingPeer": true,
                    "chaincodeQuery": true,
                    "eventSource": true
                },
                "peer2.org1.example.com": {
                    "endorsingPeer": true,
                    "chaincodeQuery": true,
                    "eventSource": true
                }
            }
        }
    },
    "organizations": {
        "Org1": {
            "mspid": "Org1MSP",
            "peers": [
                "peer0.org1.example.com",
                "peer1.org1.example.com",
                "peer2.org1.example.com"
            ],
            "certificateAuthorities": [
                "ca.example.com"
            ]
        }
    },
    "orderers": {
        "orderer0.example.com": {
            "url": "grpc://158.39.77.205:7050",
            "grpcOptions": {
                "ssl-target-name-override": "orderer0.example.com"
            },
            "tlsCACerts": {
                "pem": "-----BEGIN CERTIFICATE-----MIICNTCCAdygAwIBAgIRAPnz3O5jyXbCAKuhCWEVYSMwCgYIKoZIzj0EAwIwbDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBGcmFuY2lzY28xFDASBgNVBAoTC2V4YW1wbGUuY29tMRowGAYDVQQDExF0bHNjYS5leGFtcGxlLmNvbTAeFw0xOTAzMTAxNzU3NDRaFw0yOTAzMDcxNzU3NDRaMGwxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMRQwEgYDVQQKEwtleGFtcGxlLmNvbTEaMBgGA1UEAxMRdGxzY2EuZXhhbXBsZS5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASAWJ9neYId6SKPrxxkcAPQG+wqdd5KVcV6cUb+koKbsHVEgotDulooCpeb40pxr6O3pQ46oJeNzY1EbLJRc3hNo18wXTAOBgNVHQ8BAf8EBAMCAaYwDwYDVR0lBAgwBgYEVR0lADAPBgNVHRMBAf8EBTADAQH/MCkGA1UdDgQiBCBGmExpfyeFHDKoF/QcJT2Hf2iFjDAXDxCzZ2+1w96o+jAKBggqhkjOPQQDAgNHADBEAiA2GVftE/gSIEKVVTDUo35z/JEElOSTl9fYwtYEX85mwgIgah3bj+g178VeTpPVAmQ8wj/rFZA3McMpfwjShMwHSLQ=-----END CERTIFICATE-----"
            }
        },
         "orderer1.example.com": {
            "url": "grpc://158.39.77.205:7050",
            "grpcOptions": {
                "ssl-target-name-override": "orderer1.example.com"
            },
            "tlsCACerts": {
                "pem": "-----BEGIN CERTIFICATE-----MIICNTCCAdygAwIBAgIRAPnz3O5jyXbCAKuhCWEVYSMwCgYIKoZIzj0EAwIwbDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBGcmFuY2lzY28xFDASBgNVBAoTC2V4YW1wbGUuY29tMRowGAYDVQQDExF0bHNjYS5leGFtcGxlLmNvbTAeFw0xOTAzMTAxNzU3NDRaFw0yOTAzMDcxNzU3NDRaMGwxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMRQwEgYDVQQKEwtleGFtcGxlLmNvbTEaMBgGA1UEAxMRdGxzY2EuZXhhbXBsZS5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASAWJ9neYId6SKPrxxkcAPQG+wqdd5KVcV6cUb+koKbsHVEgotDulooCpeb40pxr6O3pQ46oJeNzY1EbLJRc3hNo18wXTAOBgNVHQ8BAf8EBAMCAaYwDwYDVR0lBAgwBgYEVR0lADAPBgNVHRMBAf8EBTADAQH/MCkGA1UdDgQiBCBGmExpfyeFHDKoF/QcJT2Hf2iFjDAXDxCzZ2+1w96o+jAKBggqhkjOPQQDAgNHADBEAiA2GVftE/gSIEKVVTDUo35z/JEElOSTl9fYwtYEX85mwgIgah3bj+g178VeTpPVAmQ8wj/rFZA3McMpfwjShMwHSLQ=-----END CERTIFICATE-----"
            }
        },
         "orderer2.example.com": {
            "url": "grpc://158.39.77.205:7050",
            "grpcOptions": {
                "ssl-target-name-override": "orderer2.example.com"
            },
            "tlsCACerts": {
                "pem": "-----BEGIN CERTIFICATE-----MIICNTCCAdygAwIBAgIRAPnz3O5jyXbCAKuhCWEVYSMwCgYIKoZIzj0EAwIwbDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBGcmFuY2lzY28xFDASBgNVBAoTC2V4YW1wbGUuY29tMRowGAYDVQQDExF0bHNjYS5leGFtcGxlLmNvbTAeFw0xOTAzMTAxNzU3NDRaFw0yOTAzMDcxNzU3NDRaMGwxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMRQwEgYDVQQKEwtleGFtcGxlLmNvbTEaMBgGA1UEAxMRdGxzY2EuZXhhbXBsZS5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASAWJ9neYId6SKPrxxkcAPQG+wqdd5KVcV6cUb+koKbsHVEgotDulooCpeb40pxr6O3pQ46oJeNzY1EbLJRc3hNo18wXTAOBgNVHQ8BAf8EBAMCAaYwDwYDVR0lBAgwBgYEVR0lADAPBgNVHRMBAf8EBTADAQH/MCkGA1UdDgQiBCBGmExpfyeFHDKoF/QcJT2Hf2iFjDAXDxCzZ2+1w96o+jAKBggqhkjOPQQDAgNHADBEAiA2GVftE/gSIEKVVTDUo35z/JEElOSTl9fYwtYEX85mwgIgah3bj+g178VeTpPVAmQ8wj/rFZA3McMpfwjShMwHSLQ=-----END CERTIFICATE-----"
            }
        }
    },
    "peers": {
        "peer0.org1.example.com": {
            "url": "grpc://158.39.77.41:7051",
            "eventUrl": "grpc://158.39.77.41:7053",
            "grpcOptions": {
                "ssl-target-name-override": "peer0.org1.example.com"
            },
            "tlsCACerts": {
                "pem": "-----BEGIN CERTIFICATE-----MIICSTCCAe+gAwIBAgIQFhCDcahCMC490GGh44ivSjAKBggqhkjOPQQDAjB2MQswCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEZMBcGA1UEChMQb3JnMS5leGFtcGxlLmNvbTEfMB0GA1UEAxMWdGxzY2Eub3JnMS5leGFtcGxlLmNvbTAeFw0xOTAzMTAxNzU3NDRaFw0yOTAzMDcxNzU3NDRaMHYxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMRkwFwYDVQQKExBvcmcxLmV4YW1wbGUuY29tMR8wHQYDVQQDExZ0bHNjYS5vcmcxLmV4YW1wbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEXNU8GTdbDCvXeXW70PLRN8tuzNd/c5wVKZM8t2TOzq/j2ty26UY6RMd9vINsJS6yE+nKG5ao1bpzQdRCVGyDgqNfMF0wDgYDVR0PAQH/BAQDAgGmMA8GA1UdJQQIMAYGBFUdJQAwDwYDVR0TAQH/BAUwAwEB/zApBgNVHQ4EIgQgDeSGYgpXz89UB4wlOkMiTe0zuxm8ae/AWMfcfW4tyTgwCgYIKoZIzj0EAwIDSAAwRQIhAMdGnYIpPlBhK+vVVVo2iF4/vlet0KPmn1qphLnemOy/AiBM/c8TwHCqG58GEwl3T9wR4OWBRS0yLG3sYWNAt0EzMg==-----END CERTIFICATE-----"
            }
        },
        "peer1.org1.example.com": {
            "url": "grpc://158.39.77.152:8051",
            "eventUrl": "grpc://158.39.77.152:8053",
            "grpcOptions": {
                "ssl-target-name-override": "peer1.org1.example.com"
            },
            "tlsCACerts": {
                "pem": "-----BEGIN CERTIFICATE-----MIICSTCCAe+gAwIBAgIQFhCDcahCMC490GGh44ivSjAKBggqhkjOPQQDAjB2MQswCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEZMBcGA1UEChMQb3JnMS5leGFtcGxlLmNvbTEfMB0GA1UEAxMWdGxzY2Eub3JnMS5leGFtcGxlLmNvbTAeFw0xOTAzMTAxNzU3NDRaFw0yOTAzMDcxNzU3NDRaMHYxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMRkwFwYDVQQKExBvcmcxLmV4YW1wbGUuY29tMR8wHQYDVQQDExZ0bHNjYS5vcmcxLmV4YW1wbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEXNU8GTdbDCvXeXW70PLRN8tuzNd/c5wVKZM8t2TOzq/j2ty26UY6RMd9vINsJS6yE+nKG5ao1bpzQdRCVGyDgqNfMF0wDgYDVR0PAQH/BAQDAgGmMA8GA1UdJQQIMAYGBFUdJQAwDwYDVR0TAQH/BAUwAwEB/zApBgNVHQ4EIgQgDeSGYgpXz89UB4wlOkMiTe0zuxm8ae/AWMfcfW4tyTgwCgYIKoZIzj0EAwIDSAAwRQIhAMdGnYIpPlBhK+vVVVo2iF4/vlet0KPmn1qphLnemOy/AiBM/c8TwHCqG58GEwl3T9wR4OWBRS0yLG3sYWNAt0EzMg==-----END CERTIFICATE-----"
            }
        },
        "peer2.org1.example.com": {
            "url": "grpc://158.39.77.131:9051",
            "eventUrl": "grpc://158.39.77.205:9053",
            "grpcOptions": {
                "ssl-target-name-override": "peer2.org1.example.com"
            },
            "tlsCACerts": {
                "pem": "-----BEGIN CERTIFICATE-----MIICSTCCAe+gAwIBAgIQFhCDcahCMC490GGh44ivSjAKBggqhkjOPQQDAjB2MQswCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEZMBcGA1UEChMQb3JnMS5leGFtcGxlLmNvbTEfMB0GA1UEAxMWdGxzY2Eub3JnMS5leGFtcGxlLmNvbTAeFw0xOTAzMTAxNzU3NDRaFw0yOTAzMDcxNzU3NDRaMHYxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMRkwFwYDVQQKExBvcmcxLmV4YW1wbGUuY29tMR8wHQYDVQQDExZ0bHNjYS5vcmcxLmV4YW1wbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEXNU8GTdbDCvXeXW70PLRN8tuzNd/c5wVKZM8t2TOzq/j2ty26UY6RMd9vINsJS6yE+nKG5ao1bpzQdRCVGyDgqNfMF0wDgYDVR0PAQH/BAQDAgGmMA8GA1UdJQQIMAYGBFUdJQAwDwYDVR0TAQH/BAUwAwEB/zApBgNVHQ4EIgQgDeSGYgpXz89UB4wlOkMiTe0zuxm8ae/AWMfcfW4tyTgwCgYIKoZIzj0EAwIDSAAwRQIhAMdGnYIpPlBhK+vVVVo2iF4/vlet0KPmn1qphLnemOy/AiBM/c8TwHCqG58GEwl3T9wR4OWBRS0yLG3sYWNAt0EzMg==-----END CERTIFICATE-----"
            }
        }
    },
    "certificateAuthorities": {
        "ca.example.com": {
            "url": "http://158.39.77.205:7054",
            "caName": "ca.example.com",
            "httpOptions": {
                "verify": false
            }
        }
    }
}



EOF

PRIVATE_KEY="${DIR}"/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/keystore/b0bcbb24d6e090ad699f65b37d2b932883a2f069c25473a07e7151d20efe685b_sk
CERT="${DIR}"/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/signcerts/Admin@org1.example.com-cert.pem

if [ "${NOIMPORT}" != "true" ]; then
 CARDOUTPUT=/tmp/composer/PeerAdmin@hlfv1.card
else
 CARDOUTPUT=PeerAdmin@hlfv1.card
fi

"${HL_COMPOSER_CLI}" card create -p DevServer_connection.json -u PeerAdmin -c "${CERT}" -k "${PRIVATE_KEY}" -r PeerAdmin -r ChannelAdmin --file $CARDOUTPUT

if [ "${NOIMPORT}" != "true" ]; then
 if "${HL_COMPOSER_CLI}" card list -c PeerAdmin@hlfv1 > /dev/null; then
 "${HL_COMPOSER_CLI}" card delete -c PeerAdmin@hlfv1
 fi

"${HL_COMPOSER_CLI}" card import --file /tmp/composer/PeerAdmin@hlfv1.card 
 "${HL_COMPOSER_CLI}" card list
 echo "Hyperledger Composer PeerAdmin card has been imported, host of fabric specified as '${HOST}'"
 rm /tmp/composer/PeerAdmin@hlfv1.card
else
 echo "Hyperledger Composer PeerAdmin card has been created, host of fabric specified as '${HOST}'"
fi