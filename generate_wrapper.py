#!/usr/bin/env python3

import os, glob, re

funcre = re.compile(r'func.*Api\) ')
apire  = re.compile(r'type .*Api struct')

TEMPLATE = """package packswgo

import "os"

// APITokenVar is the name of teh envvar which holds auth token to Packet API
const APITokenVar = "PACKET_AUTH_TOKEN"

type Client struct {{
{clientStructBody}
}}

func getCfg() *Configuration {{
	cfg := NewConfiguration()
	cfg.APIKey["X-Auth-Token"] = os.Getenv(APITokenVar)
	return cfg
}}

// NewClient returns API client
func NewClient() Client {{
	cfg := getCfg()
	return Client{{
{clientStructCreation}
	}}
}}

{generatedMethods}

"""

def funcStart(l):
    return funcre.sub("", l)

def api(l):
    return (l[8:].split(")")[0])[:-3]

def funcCall(l):
    fs = funcStart(l)
    name, rest = fs.split("(",1)
    fargs = (rest.split(")", 1)[0]).split(",")
    fargnames = [i.strip().split(" ")[0] for i in fargs]
    return name + "(" + ", ".join(fargnames) + ")"


def generateFunc(l):
    return "func (c *Client) " + funcStart(l) + "\n\treturn c." + api(l) + "." + funcCall(l) + "\n}"

def generateApiMethods():
    for f in glob.glob("*_api.go"):
        for l in open(f).readlines():
            if funcre.match(l):
                yield(generateFunc(l[:-1]))

def findApis():
    for f in glob.glob("*_api.go"):
        for l in open(f).readlines():
            if apire.match(l):
                yield(l.split(" ")[1][:-3])
d = {}
d['clientStructBody'] = "\n".join(
    ["\t{0} {0}Api".format(i) for i in findApis()]
)
d['clientStructCreation'] = "\n".join(
    ["\t\t{0}: {0}Api{{Configuration: cfg}},".format(i) for i in findApis()]
)
d['generatedMethods'] = "\n\n".join(generateApiMethods())

print(TEMPLATE.format(**d))


