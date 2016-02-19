module.exports = function(doc) {
  doc._id = doc._id['$oid']; 
  doc["fullName"] = doc["firstName"] + " " + doc["lastName"];
  return doc;
}