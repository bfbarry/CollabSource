import pymongo as pm
import json
import os

config_path = os.path.join(os.path.dirname(__file__), 'db_config.json')
with open(config_path, 'r') as f:
    config = json.load(f)
    CONN_STR = config['CONN_STR']

TABLES_TO_DROP = ['projects', 'users']

if __name__ == '__main__':
    client = pm.MongoClient(CONN_STR)
    print('connected to mongodb')
    db = client['local']

    for t in TABLES_TO_DROP:
        db[t].drop()
        print(f'dropped collection {t}')