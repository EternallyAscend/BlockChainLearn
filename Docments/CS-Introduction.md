# Introduction of Client / Server API Designing

[toc]

## 1. Client

### 1.1 Communication Specification

#### 1.1.1 Users' Client

##### 1.1.1.1 To Information Exchange Server

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;When the users' client send messages to server, server need a formatted message to deal with. The rules are as follows. "Doc." means doctors, "Use." means common users and "Dev." means medical equipment and private health devices.

| IP : Port (IPv4)      | Command         | Authority ID    | Parameter (Table in details) |
| --------------------- | --------------- | --------------- | ---------------------------- |
| String 0 to 24        | String 25 to 31 | String 32 to 48 | String 48 to end             |
| xxx.xxx.xxx.xxx:xxxx; | Query;          | Doc. / Use.     | Account, Date and Type List; |
| xxx.xxx.xxx.xxx:xxxx; | Write;          | Doc. / Dev.     | Account, Date and Data List; |
| xxx.xxx.xxx.xxx:xxxx; | Verify;         | Doc. / Use.     | Hash String;                 |

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"IP : Port (IPv4)" is using to verify the user, return data (Query and Write) and record the operations, "Command" is using to decide to function and method of returning data to clients,  "Authority ID" is using to verify the hash id of the account and judge whether he has the authority to operate, "Parameter" is using to operate the network.



##### 1.1.1.2 To Account Server

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;When the users log in, register or verify authentication information.

| IP : Port (IPv4)      | Command         | Authority ID    | Parameter                           |
| --------------------- | --------------- | --------------- | ----------------------------------- |
| String 0 to 24        | String 25 to 31 | String 32 to 48 | String 48 to end                    |
| xxx.xxx.xxx.xxx:xxxx; | Login;          | Doc. / Use.     | Account, Digital Certificate, Date. |
| xxx.xxx.xxx.xxx:xxxx; | Logout;         | Doc. / Use.     | Account, Date.                      |
| xxx.xxx.xxx.xxx:xxxx; | Verify;         | Doc. / Use.     | ID or others                        |



#### 1.1.2 Functions Specification

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Users can do 



### 1.2 Managers' Clients

#### 1.2.1

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Managers can allow and verify the user and manage the chaincode on blockchain network. The managers can connect to the anchor node.

| IP : Port (IPv4)      | Command         | Authority ID    | Parameter        |
| --------------------- | --------------- | --------------- | ---------------- |
| String 0 to 24        | String 25 to 26 | String 32 to 48 | String 48 to end |
| xxx.xxx.xxx.xxx:xxxx; | Verify; (00)    | ID              | User ID, Status. |
| xxx.xxx.xxx.xxx:xxxx; | New Org(01)     | ID              | Org Info, Date.  |
| xxx.xxx.xxx.xxx:xxxx; | Add Node(10)    | ID              | Node Info, Date. |
| xxx.xxx.xxx.xxx:xxxx; | Chaincode(11)   | ID              | Chaincode, Date. |





## 2. Server

### 2.1 Information Stored

#### 2.1.1 Users' and devices' Information

| User Type       | ID        | Hash ID | Date        | Hash Code        | Status        |
| --------------- | --------- | ------- | ----------- | ---------------- | ------------- |
| String 0 to 1   | String 18 | String  | String 18   | String           | Boolean       |
| Common Users    | ID        | SM2(ID) | Record Date | Accumulated Data | Alive / Other |
| Doctor Users    | ID        | SM2(ID) | Record Date | Operations       | Legal         |
| Medical Devices | ID Code   | ID Code | Related Org | Accumulated Data | Legal         |
| Users' Devices  | ID Code   | ID Code | Related ID  | Accumulated Data | Legal         |



#### 2.1.2 User Types

| User Type       | Code | ID        | Hash ID |
| --------------- | ---- | --------- | ------- |
| Common Users    | 00   | 18 Digits | SM2(ID) |
| Doctor Users    | 01   | 18 Digits | SM2(ID) |
| Medical Devices | 11   | 32 Digits | ID      |
| Users' Devices  | 10   | 32 Digits | ID      |



#### 2.1.3 Hash Code Composition

| User Type    | Left         | Right             | Lock          |
| ------------ | ------------ | ----------------- | ------------- |
| Doctors      | Query Record | Write Record      | Null          |
| Common Users | Text Record  | Multimedia Record | enc(Hash ID); |
| Devices      | Query Record | Write Record      | enc(ID Code); |

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The :enc()" is a encryption function to get lock type for data encryption.



#### 2.1.4 Encryption function



## 3. Chain Code

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Block Data:

| Hash ID   | Date             | Code     | Data      |
| --------- | ---------------- | -------- | --------- |
| String 32 | String 16        | String 8 | String 32 |
| SM2(ID)   | YYYYMMDDHHMMSSXX | 00000000 | Data      |
| ID Code   | YYYYMMDDHHMMSSXX | 01000000 | Data      |

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Code List

| Code     | Name            | Data   | Remark    |
| -------- | --------------- | ------ | --------- |
| 00Length | Case            | String | length    |
| 01XXXXXX | Medical Devices | String | Data type |
| 10XXXXXX | Multimedia Data | String | Data type |
| 11XXXXXX | Users' Devices  | String | Data type |

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The block has only 50 lines in one record.  The case of patients will divided into pieces with length as other data types.

## 4. Other



```
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
```

