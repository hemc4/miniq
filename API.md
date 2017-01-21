#API Enpoints

Note: This is minimal documentation for the understanding of the functioanlity 

* Add consumer 
    Registering the consumer to the service 
    
    Request -  ConsumerName(string)
    Response - ConsumerId(string)

* Create queue
    Create a new queue 
    
    Request - ConsumerId(string), QueueName(string)
    Response - Status (bool)

* Add message 
    Publish a message to service 
    
    Request- ConsumerId(string), MsgPayload(string)
    Response- MessageId(string)

* Get message 
    Consume the message for processing 
    
    Request- ConsumerId(string)
    Response- MessageData(Msgid, payload)

* Delete queue 
    To remove the list from the service 
    
    Request- ConsumerId(string), QueueName(string)
    Response- Status(bool)

* Ack Message 
    Once the processing is done consumer can notify teh service for the same.
     
    Request- ConsumerId(string), MessageId(string), Ack(bool)
    Response - Status(bool)
    
     


