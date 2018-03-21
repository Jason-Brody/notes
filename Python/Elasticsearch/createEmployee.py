import json
import subprocess

index = 'employees'
tp = 'employee'

base_url = 'http://localhost:9200'


data = {
    "first_name" :  "Jane",
    "last_name" :   "Smith",
    "age" :         32,
    "about" :       "I like to collect rock albums",
    "interests":  [ "music" ]
}

id = 2
url = '/'.join([base_url,index,tp,str(id)])
cmd = "curl -XPUT '%s' -d '%s' "  % (url,json.dumps(data)) 
subprocess.call(cmd,shell=True, close_fds=True, timeout=30)
