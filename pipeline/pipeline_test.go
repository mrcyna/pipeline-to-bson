package pipeline

import (
	"io/ioutil"
	"testing"
)

func TestSingleLinePipelineWillBeTransformedCorrectly(t *testing.T) {
	input := `[{"$count": "language"}]`
	expected := `bson.A{bson.M{"$count": "language"}}
`
	if !Validate(input) {
		t.Errorf("unable to validate correctly")
		return
	}

	result := Transform(input)
	if result != expected {
		t.Errorf("unable to transformed correctly, expected:\n%v\nbut got:\n%v", expected, result)
		return
	}
}

func TestMultiLinePipelineWillBeTransformedCorrectly(t *testing.T) {
	input := `[{$project: {"_id": 0, "email": 1, "members.access": 1}},
{$set: {"pending": false}}]`

	expected := `bson.A{bson.M{"$project": bson.M{"_id": 0, "email": 1, "members.access": 1}},
bson.M{"$set": bson.M{"pending": false}}}
`
	if !Validate(input) {
		t.Errorf("unable to validate correctly")
		return
	}

	result := Transform(input)
	if result != expected {
		t.Errorf("unable to transformed correctly, expected:\n%s\nbut got:\n%s", expected, result)
		return
	}
}

func TestBigQueryPipelineWillBeTransformedCorrectly(t *testing.T) {
	content, err := ioutil.ReadFile("../example.txt")
	if err != nil {
		t.Errorf("unable to open example file")
	}
	input := string(content)

	expected := `bson.A{
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
},
`
	if !Validate(input) {
		t.Errorf("unable to validate correctly")
		return
	}

	result := Transform(input)
	if result != expected {
		t.Errorf("unable to transformed correctly, expected:\n%s\nbut got:\n%s", expected, result)
	}
}

func TestWrongQueryWillFailedWithIncorrectBrackets(t *testing.T) {
	input := `[{{{{{{{{{{{{{$project: {"_id": 0, "email": 1, "members.access": 1}},
{$set: {"pending": false}}]`

	if Validate(input) {
		t.Errorf("unable to validate correctly")
		return
	}
}

func TestWrongQueryWillFailedWithIncorrectOpening(t *testing.T) {
	input := `{"$count": "language"}`

	if Validate(input) {
		t.Errorf("unable to validate correctly")
		return
	}
}