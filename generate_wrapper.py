#!/usr/bin/env python3

import os, glob, re

funcre = re.compile(r'func.*Api\) ')
apire  = re.compile(r'type .*Api struct')

TEMPLATE = """package packswgo

import (
    "os"
    "log"
)

// APITokenVar is the name of teh envvar which holds auth token to Packet API
const (
    APITokenVar = "PACKET_AUTH_TOKEN"
    APIURLVar = "PACKET_API_URL"
)

type Client struct {{
{clientStructBody}
}}

func getCfg(token, apiURL string) *Configuration {{
	cfg := NewConfiguration()
        cfg.APIKey["X-Auth-Token"] = token
        cfg.BasePath = apiURL
	return cfg
}}

// NewClient returns Client
func NewClient() Client {{
        return NewClientWithToken(os.Getenv(APITokenVar))
}}

// NewClientWithToken returns API client
func NewClientWithToken(token string) Client {{
        if token == "" {{
            log.Fatal("You must provide Packet API Auth token")
        }}
        apiURL := "https://api.packet.net"
        if os.Getenv(APIURLVar) != "" {{
            apiURL = os.Getenv(APIURLVar)
        }}

	cfg := getCfg(token, apiURL)
	return Client{{
{clientStructCreation}
	}}
}}

"""


def findApis():
    for f in glob.glob("*_api.go"):
        for l in open(f).readlines():
            if apire.match(l):
                yield(l.split(" ")[1][:-3])
d = {}
d['clientStructBody'] = "\n".join(
    ["\t{0}Api".format(i) for i in findApis()]
)
d['clientStructCreation'] = "\n".join(
    ["\t\t{0}Api{{Configuration: cfg}},".format(i) for i in findApis()]
)

print(TEMPLATE.format(**d))


