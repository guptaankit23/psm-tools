#!/usr/bin/python3
import json
import requests
import re
import os
import sys
import logging
from login import get_psm_config
import warnings
import collections
warnings.simplefilter("ignore")

HOME = os.environ["HOME"]
psm_config_path = HOME+"/.psm/config.json"
psm_config = {}

def downloadSwaggerFiles(target):
    host = psm_config["psm-ip"]
    dirname = "swagger_" + target    
    if not os.path.exists(dirname):
        os.mkdir(dirname)

    response = requests.get("https://"+host+"/docs/generated/swaggeruri.html", verify=False)
    matches = re.findall("href=\"([^\"]*)", response.text)

    for match in matches:
        resp = requests.get("https://"+host+match, verify=False)
        filename = match.replace("/swagger/", "")
        jsondata = processSwagger(filename, resp.json())
        with open(dirname + "/" + filename + ".json", "w") as f:
            json.dump(jsondata, f, indent=4)

def removeRequired(filename, jsondata):
    for key in jsondata["definitions"]:
        if "required" in jsondata["definitions"][key] and len(jsondata["definitions"][key]["required"]):
            print(filename, key, jsondata["definitions"][key]["required"])
            jsondata["definitions"][key]["required"] = []
    return jsondata

def removeDuplicates(filename, jsondata):

    if type(jsondata) != dict:
        return

    for key in jsondata.keys():

        value = jsondata[key]
        if type(value) is dict:
            removeDuplicates(filename, jsondata[key])
        elif type(value) is list:
            duplicate_map = set()
            newlist = []

            for i,v in enumerate(value):
                removeDuplicates(filename, jsondata[key][i])

            value = jsondata[key]
            for v in value:
                hv = None
                if isinstance(v, collections.Hashable):
                    hv = v
                else:
                    hv = frozenset(v)
                if hv not in duplicate_map:
                    duplicate_map.add(hv)
                    newlist.append(v)
            jsondata[key] = newlist
    return jsondata
            
def processSwagger(filename, jsondata):

    jsondata = removeRequired(filename, jsondata)
    removeDuplicates(filename, jsondata)

    '''
    fwlog apigroup:
        fwloglist.meta.name is always empty, serialization fails because of the checks on apiObjectMeta
    since all apigroups are merged into the same folder now, checks on apiObjectMeta and apiListWatchOptions are removed from all swagger files
    '''
    del jsondata["definitions"]["apiListWatchOptions"]["properties"]["name"]["minLength"]
    del jsondata["definitions"]["apiListWatchOptions"]["properties"]["tenant"]["minLength"]
    del jsondata["definitions"]["apiListWatchOptions"]["properties"]["namespace"]["minLength"]
    del jsondata["definitions"]["apiListWatchOptions"]["properties"]["name"]["pattern"]
    del jsondata["definitions"]["apiListWatchOptions"]["properties"]["tenant"]["pattern"]
    del jsondata["definitions"]["apiListWatchOptions"]["properties"]["namespace"]["pattern"]

    del jsondata["definitions"]["apiObjectMeta"]["properties"]["name"]["minLength"]
    del jsondata["definitions"]["apiObjectMeta"]["properties"]["tenant"]["minLength"]
    del jsondata["definitions"]["apiObjectMeta"]["properties"]["namespace"]["minLength"]
    del jsondata["definitions"]["apiObjectMeta"]["properties"]["name"]["pattern"]
    del jsondata["definitions"]["apiObjectMeta"]["properties"]["tenant"]["pattern"]
    del jsondata["definitions"]["apiObjectMeta"]["properties"]["namespace"]["pattern"]
    
    if filename == "objstore":
        del jsondata["paths"]["/objstore/v1/uploads/snapshots"]
        del jsondata["paths"]["/objstore/v1/uploads/images"]
    if filename == "cluster":
        params = jsondata["paths"]["/configs/cluster/v1/distributedservicecards/{O.Name}"]["get"]["parameters"]
        newparams = []
        param_set = set()

        for param in params:
            if param["name"] not in param_set:
                param_set.add(param["name"])
                newparams.append(param)
        jsondata["paths"]["/configs/cluster/v1/distributedservicecards/{O.Name}"]["get"]["parameters"] = newparams

    return jsondata

def usage():
    print ("Usage: getswagger.py [cloud|dss|ent]")
    sys.exit(1)

if __name__ == "__main__":
    if ((len(sys.argv) != 2) or (sys.argv[1] not in "cloud dss ent")):
        usage()
    psm_config = get_psm_config()
    downloadSwaggerFiles(sys.argv[1])
