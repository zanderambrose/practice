db = db.getSiblingDB('admin');

db.auth("root", "examplepassword");

db = db.getSiblingDB('whoshittin');

db.createUser({
    'user': "whoshittinuser",
    'pwd': "whoshittinpassword",
    'roles': [{
        'role': 'dbOwner',
        'db': 'whoshittin'
    }]
});

db.createCollection('client');

db.client.insertOne({
    clientId: "admin",
    apiKey: "admin",
    permissions: { isAdmin: true }
});
