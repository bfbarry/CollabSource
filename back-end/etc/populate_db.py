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



    user_docs = [
        {
            "name": f"Dev User {i}",
            "email": f"DevUser{i}@z.com",
            "description": "hey :) I'm a sample user I like certain hobbies, and I am skilled at certain things.",
            "skills": ["sorcery", "tiktok dance"],
            "password": "123"
        }
        for i in range(N_USERS)
    ]

    db['users'].insert_many(user_docs)
    print(f'inserted {N_USERS} users')

    get_random_id = lambda : db['users'].aggregate([{'$sample': {'size': 1}}]).next()['_id']
    project_docs = [
        {
            "name"        : f"Dev Project {i}",
            "description" : f"This project is about certain things. It will accomplish a cool goal!",
            "category" 	  : "sample_category",
            "tags"        : ["tag1", "tag2"],
            "seeking"     : ["wizard", "engineer", "chef"],
            "ownerId"     : get_random_id()
        }
        for i in range(N_PROJECTS)
    ]

    db['projects'].insert_many(project_docs)
    print(f'inserted {N_PROJECTS} projects')


