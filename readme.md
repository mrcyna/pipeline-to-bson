# Pipeline to BSON
A simple CLI app to convert MongoDB pipeline command in JSON format to Golang BSON structs.

## How it works?
simply by running the app:
```
$ ./pipeline-to-bson
Enter the filepath:
```
And by validating the path of the file, and it's syntax transformation will be written to stdout.

## Example
Input:
```
  [
	{$match: {"members": {$elemMatch: {"account_id": "0db28ed0-7d67-42ef-bd2d-f1cb9a814198"}}}},
 	{$project: {"_id": 0, "email": 1, "members.access": 1}},
 	{$set: {"pending": false}},
        {$unwind: "$members"},
        { "$replaceRoot": {
            "newRoot": { "$mergeObjects": ["$$ROOT", "$members"] }
        }},
        {$unionWith: {"coll": "admins", "pipeline": [
     			{$match: {"account_id": "0db28ed0-7d67-42ef-bd2d-f1cb9a814198", "status": "pending"}},
 			{$project: {"_id": 0,  "email": 1, "access": 1}},
 			{$set: {"pending": true}}
	     	]}
	}
]
```

Desired Output:
```
bson.A{
bson.M{"$match": bson.M{"members": bson.M{"$elemMatch": bson.M{"account_id": "0db28ed0-7d67-42ef-bd2d-f1cb9a814198"}}}},
bson.M{"$project": bson.M{"_id": 0, "email": 1, "members.access": 1}},
bson.M{"$set": bson.M{"pending": false}},
bson.M{"$unwind": "$members"},
bson.M{ "$replaceRoot": bson.M{
"newRoot": bson.M{ "$mergeObjects": bson.A{"$$ROOT", "$members"} },
}},
bson.M{"$unionWith": bson.M{"coll": "admins", "pipeline": bson.A{
bson.M{"$match": bson.M{"account_id": "0db28ed0-7d67-42ef-bd2d-f1cb9a814198", "status": "pending"}},
bson.M{"$project": bson.M{"_id": 0,  "email": 1, "access": 1}},
bson.M{"$set": bson.M{"pending": true}},
}},
},
```