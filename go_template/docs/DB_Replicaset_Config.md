# Mongo Db ReplicaSet Configure 
steps : 
* Connect to one of the mongodb servers. You can run  ```make mongo_ui_run```
* Execute the following code 
```js 
mongodb> config = { "_id" : "activegeo_set", "members" : [ { "_id" : 0, "host" : "mongodb1:27017" }, { "_id" : 1, "host" : "mongodb2:27017" }, { "_id" : 2, "host" : "mongodb3:27017" } ] }
```
* and Initioate the databaes with folowing code 
```js 
mongodb> rs.initiate(config)
```

* If everything goes well you will see the `{ "ok" : 1 }` on terminal results.

* to add new database to rs please run: 
```js
mongodb:activegeo_set> rs.add( { "_id":4, "host": "new_mongo:27017" } )
```

* to remove new database to re please run: 
```js
mongodb:activegeo_set> rs.add("new_mongo:27017")
```