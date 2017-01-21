#MiniQ
This is a simple message queue service using Redis as storage and Go language 

Service is designed keeping below main poitns in mind

* Persistent - in case of the crash, There should be no data loss 
* Atomic - One message should be processed by one consumer 


#Basic Design 
Redis provide RPOPPUSH functionality as atomic operation so we are going to use for reliability purpose.
Each queue will be a list and new message will be LPUSH to the list. 

Each consumer will have it's own working queue and and process message from this queue and mark them ack once 
done with processing 

When a consumer get a message it goes to "processing queue " and the consumers "working queue". Once processes it 
goes POP'ed  from the processing queue otherwise goes back to Main queue. This way we can handle the re-
processing of the message in case of consumer failure.



#Scaling 

##Redis 
To maintain high load on the storage, we require to implement clustering on the redis. On the top
of clustering data sharding need to be implemented

Redis provide tuning of the params based on desired persistence and performance 


##Monitoring 

To support high availbility on scale, we need some kind of monitoring on the redis level and machine level.
A simple json api to fetch following iinforamtion form the redis and display on a simple web interface 
will be a nice starting point 

* All consumers 
* All queues
* Queue lengths 
* Failed messages 

For machine level checks we can use Sensu or Nagios 
