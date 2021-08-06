db = db.getSiblingDB("admin");

if( !db.getUser(username)){
    db.createUser({
        user: username,
        pwd: password,
        mechanisms: ["SCRAM-SHA-1",],
        roles: [
            {
                role: 'dbOwner',
                db: database
            }
        ]
    });
}