import random
from pymongo import MongoClient
from bson import ObjectId
from datetime import datetime, timedelta
import random
import json, os

make_random_date = lambda days : datetime.now() - timedelta(days=days) + timedelta(days=random.randint(0, days))
config_path = os.path.join(os.path.dirname(__file__), 'db_config.json')

with open(config_path, 'r') as f:
    config = json.load(f)
    CONN_STR = config['PROD']

"""
note: training db is all art&music for some reason

"""

# USE ENV TODO TODO
FromClient = MongoClient(CONN_STR)
ToClient = MongoClient('mongodb://root:example@localhost:27017')

db_train = FromClient['training_data']
db_end = ToClient['local']


# Collections
old_projects_collection = db_train['projects']
new_projects_collection = db_end['projects']
users_collection = db_end['users']

# Fetch all documents from ProjectOld
old_projects = old_projects_collection.aggregate([{"$sample": {"size": old_projects_collection.count_documents({})}}])

# Loop through each document in ProjectOld
for old_project in old_projects:
    # Get a random user from the Users collection
    user_count = users_collection.count_documents({})
    random_index = random.randint(0, user_count - 1)
    random_user = users_collection.find().skip(random_index).limit(1)[0]

    # Create a new Project document
    new_project = {
        "_id": old_project.get("_id", ObjectId()),
        "ownerId": random_user["_id"],
        "ownerEmail": random_user["email"],
        "name": old_project.get("name", ""),
        "description": old_project.get("description", ""),
        "category": old_project.get("category", ""),
        "tags": old_project.get("tags", []),
        "seeking": old_project.get("seeking", []),
        "members": [],
        "dateCreated": make_random_date(90),
        "links": [],
        "memberRequests": []
    }

    # Insert the new project into the Project collection
    new_projects_collection.insert_one(new_project)

print("Data transfer complete.")
