module.exports = function(doc) {
  doc._id = doc._id['$oid']; 
//  doc["fullName"] = doc["firstName"] + " " + doc["lastName"];
  doc["fullName"] = "full name";
  return doc;
}