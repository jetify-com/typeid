#! env bash

echo "Convert to JSON"
yq eval typeid/typeid/spec/valid.yml --tojson > typeid/typeid/spec/valid.json
yq eval typeid/typeid/spec/invalid.yml --tojson > typeid/typeid/spec/invalid.json

echo "Update typeid-go"
cp typeid/typeid/spec/valid.yml typeid/typeid-go/testdata/valid.yml
cp typeid/typeid/spec/invalid.yml typeid/typeid-go/testdata/invalid.yml

echo "Update typeid-js"
cat <<-TS > typeid/typeid-js/test/valid.ts
// Data copied from the valid.yml spec file
export default $(cat typeid/typeid/spec/valid.json)
TS
cat <<-TS > typeid/typeid-js/test/invalid.ts
// Data copied from the invalid.yml spec file
export default $(cat typeid/typeid/spec/invalid.json)
TS
