Structure of the Bolt Database: 

Every separate label has a bucket. 
There is a bucket for things with the "important" tag
There is a bucket for anything unlabeled. 

Really, once you have these labels, it's better to do a SQL database.

for a (key, value) pair, the key will be the name of the task to be completed. 

Outside of the "imp" bucket, the value, if it exists, will be the timestamp of when that task was completed in RFC3339 format. 

Inside of the "imp" bucket, the value will be the label of the task. 

