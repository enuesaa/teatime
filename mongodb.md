# MongoDB
- https://nodejs.keicode.com/mongodb/backup-and-restore.php

## Backup
```console
$ mongodump
2024-09-16T04:47:28.316+0000	writing admin.system.version to dump/admin/system.version.bson
2024-09-16T04:47:28.317+0000	done dumping admin.system.version (1 document)
2024-09-16T04:47:28.318+0000	writing a.delete_me to dump/a/delete_me.bson
2024-09-16T04:47:28.319+0000	done dumping a.delete_me (1 document)
```

## Restore
データが入っていたからエラーが発生したけど実用上問題はないだろう
```console
$ mongorestore
2024-09-16T04:49:01.897+0000	using default 'dump' directory
2024-09-16T04:49:01.897+0000	preparing collections to restore from
2024-09-16T04:49:01.897+0000	reading metadata for a.delete_me from dump/a/delete_me.metadata.json
2024-09-16T04:49:01.898+0000	restoring to existing collection a.delete_me without dropping
2024-09-16T04:49:01.898+0000	restoring a.delete_me from dump/a/delete_me.bson
2024-09-16T04:49:01.902+0000	continuing through error: E11000 duplicate key error collection: a.delete_me index: _id_ dup key: { _id: ObjectId('66e7b7cc550efad752a8587f') }
2024-09-16T04:49:01.909+0000	finished restoring a.delete_me (0 documents, 1 failure)
2024-09-16T04:49:01.910+0000	no indexes to restore for collection a.delete_me
2024-09-16T04:49:01.910+0000	0 document(s) restored successfully. 1 document(s) failed to restore.
```
