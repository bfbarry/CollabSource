import pymongo as pm
import json
import os

"""
TODO: use chatgpt to create projects and users
"""
config_path = os.path.join(os.path.dirname(__file__), 'db_config.json')

with open(config_path, 'r') as f:
    config = json.load(f)
    CONN_STR = config['CONN_STR']
    N_PROJECTS = config['N_PROJECTS']
    N_USERS = config['N_USERS']

if __name__ == '__main__':

    client = pm.MongoClient(CONN_STR)
    print('connected to mongodb')
    db = client['local']

    project_docs = [
        {
            "name"        : f"DevProj{i}",
            "description" : f"DevDesc{i}",
            "category" 	  : "sample_category",
            "tags"        : ["tag1", "tag2"]
        }
        for i in range(N_PROJECTS)
    ]

    user_docs = [
        {
            "name": f"DevUser{i}",
            "email": f"DevUser{i}@z.com",
            "description": "hey :)",
            "skills": ["sorcery", "tiktok dance"]
        }
        for i in range(N_USERS)
    ]

    db['projects'].insert_many(project_docs)
    print(f'inserted {N_PROJECTS} projects')

    db['users'].insert_many(user_docs)
    print(f'inserted {N_USERS} users')
