Source({name:"localmongo"})
.transform({filename: "transformers/addFullName.js", "namespace": "transformer./.*/"})
.save({name:"es"});