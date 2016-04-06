//Source({name:"localmongo", namespace:"foo.bar|user|order"})
//.save({name:"es", namespace:"foo.bar|user|order"});
Source({name:"localmongo", namespace:"foo.user"})
.save({name:"es", namespace:"foo.user"});
