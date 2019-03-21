Feature: IoT smarthome network
 
    Background:
        Given I have deployed the business network definition ..
        And I have added the following participants
        """
        [
        {"$class":"org.blockchain.models.Home", "homeId":"Salmans_home", "homeOwner":{"$class":"org.blockchain.model.Owner", "firstname":"Salman","lastname:Ahmed"},"address":{"$class":"org.blockchain.model.Address","street":"Problemveien 21","city":"Oslo"}},
        {"$class":"org.blockchain.models.Home", "homeId":"Franks_home", "homeOwner":{"$class":"org.blockchain.model.Owner", "firstname":"Frank","lastname:Johnson"},"address":{"$class":"org.blockchain.model.Address","street":"Problemveien 22","city":"Oslo"}},
        {"$class":"org.blockchain.models.Home", "homeId":"Johns_home", "homeOwner":{"$class":"org.blockchain.model.Owner", "firstname":"John","lastname:Wick"},"address":{"$class":"org.blockchain.model.Address","street":"Problemveien 23","city":"Oslo"}}
        ]
        """
        And I have added the following asset of type org.blockchain.model.Device
            | deviceId          | home          | deviceLocation        | temperature  | rights        | transmissionInterval | nickname       |   
            | Salmans_Fridge    | Salmans_home  | second_floor          | 0            | READ,WRITE    | 200                  | Fridge_midge   |
            | Johns_Iphone      | John_home     | third_floor           | 0            | READ          | 3000                 | my_iphone      |
            | Franks_hub        | Frank_home    | first_floor           | 0            | WRITE         | 5000                 | the_hub        |    
        When I submit the following transactions of type org.blockchain.model.DeviceReading
            | device         | temperature  |   readingTime     |
            | Salmans_Fridge | 30           |    15:30          |
            | Johns_Iphone   | 5            |    12:00          |   
            | Franks_hub     | 10           |    2:00           |  

Scenario: When the device has sent device readings
    When I submit the following transaction of type org.acme.shipping.perishable.DeviceReading
        | deviceId       |
        | Salmans_Fridge |
     
    Then I should have the following participants
    """
    [
     {"$class":"org.blockchain.models.Home", "homeId":"Salmans_home", "homeOwner":{"$class":"org.blockchain.model.Owner", "firstname":"Salman","lastname:Ahmed"},"address":{"$class":"org.blockchain.model.Address","street":"Problemveien 21","city":"Oslo"}}
    ]
    """