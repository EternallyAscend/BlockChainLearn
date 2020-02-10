# Software Requirements Specification

Version 1.0

</br></br>

<font size="8" weight="2"><b><center>Medical Data Blockchain Services Platform with Hospital-ledger</center></b></font>

</br></br>

## Content

[toc]

Version 1.0 last edited by H.L. Yu on 2020-02-10 Mon.

</br><center>Nankai University 2020 &copy; All Rights Reserved.</center></br>

## 1. Introduction

### 1.1 Purpose

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;This software project is an open-source system and it has two parts. One is "Hospital-ledger", another is "Medical Data Blockchain Services Platform". "Hospital-ledger" is provided for us to record the medical data among hospital workers and normal user securely. "Medical Data Blockchain Services Platform" is provided for us to manage the blockchain network and to exchange the  information and data. The "Hospital-ledger" is the basic of the project and "Medical Data Blockchain Services Platform".

### 1.2 Intended Audience and Reading Suggestions

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;This document is intended for developers, project managers, testers, users and document writers. All parts are provided for develops, project managers, testers and document writers. The part 2.4 Operating Environment and part 2.6 User Documentation are provided for users. These two part are provided for user to learn how to use the software.

### 1.3 Product Scope

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;This software project is designed for us to record the medical data of patients securely and share the data among different hospitals in order to give patient a better treatment, avoiding missing the medical history, reduce the risk of treatment and improve the relationship between the doctors and patients. As well as it will act as an open platform for data transparency without personal identity information in details.

### 1.4 References

1. Apache CouchDB. http://couchdb.apache.org.
2. Apache Kafka. http://kafka.apache.org.
3. Hyperledger. http://www.hyperledger.org.
4. Hyperledger Fabric. http://github.com/hyperledger/fabric.
5. The Linux Foundation. http://www.linuxfoundation.org.
6. Raft https://raft.github.io.
7. E. Androulaki, C. Cachin, C. Ferris, S. Muralidharan, C. Stathakopoulou: Hyperledger Fabric: A Distributed Operating System for Permissioned Blockchains. In EuroSys '18 April 23-26, 2018, Porto, Portugal.
8. H.Y. Li, L.H. Zhu, M. Shen,  F. Gao, X.L. Tao. S. Liu: Blockchain-Based Data Preservation System for Medical Data. In Springer Science+Business Media, LLC, part of Springer Nature 2018, (2018)42:141.

## 2. Overall Description

### 2.1 Product Perspective

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The software project includes two subsystems. One part is "Hospital-ledger", it provides a distributed network to store the medical data by blockchain, make the network lightweight and share the data securely between hospitals and patients by using the distributed alliance blockchain system Hyperledger Fabric which designed by IBM. Another part is "Medical Data Blockchain Services Platform", it provides a platform to create, manage, edit and destroy the blockchain network, give the authentications to doctors, common users and managers, edit the smart-contact, verify the data and transport the data.

### 2.2 Product Functions

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The first subsystem "Hospital-ledger" has basic functions for storing the medical data. It provides data writing services for doctors, data querying services for both doctors and common users. 

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;On the other hand, the second subsystem "Medical Data Blockchain Services Platform" has basic blockchain services management functions, identify verifying and register services, information pushing services, data transporting services and verifying services for blockchain network. 

### 2.3 User Classes and Characteristics

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;There are two types of users in subsystem "Hospital-ledger" and four types of users in subsystem "Medical Data Blockchain Services Platform". There only show the users of them without other services peers which will show in details in part 3 External Interface Requirements.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;In "Hospital-ledger", there are two types of users: doctors and patients, actually they log in with their account of "Medical Data Blockchain Services Platform" and then the system will connect to the relevant peer in blockchain network. The doctors belong to certain hospital and they can commit the request to write and query the medical data of patients, these operations will record on the "Medical Data Blockchain Services Platform", they also can pull requests of original authentications for themselves and reset the key of their account. By contrast, the patients only have the authority to query the data of themselves, pull requests of original authentications for themselves and reset the key of their accounts.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;In "Medical Data Blockchain Services Platform", there are four types of users: doctors, common users, organization managers and platform managers. There are still some other peers in this subsystem, such as the transfer peer and verify peer for the security of the medical data of common users. The doctors and common users (patients in "Hospital-ledger") in this subsystem can log in with their account key and digital certificate, then system will connect to relevant peer in blockchain network, and push the medical data return to their clients which they request, especially, the common users can download the multimedia medical data of themselves from the HDFS services on cloud. The managers, their authorities are higher than doctors and common users, they can log in with their account key and digital certificate, then the system will connect to the relevant peer of them. The managers of organizations can verify the identity of doctors and common users, send the digital certificate to them via the client, query medical data without personal information of common users, create and delete the organizations of doctors or common users,  create accounts and edit the accounts status of doctors and common users. The platform managers can verify the identity of the managers of organizations, monitor the status of the network system and the status of services platform, create, edit, update and stop the smart contract which named "chain code" in Hyperledger Fabric.

### 2.4 Operating Environment

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Now, the subsystem "Hospital-ledger" can run one-click setup on the Linux operating system CentOS 7 or Red Hat which based on CPU whose architecture is X86_64(amd64). The Ubuntu 18.04 LTS and other system which based on Debian with the same CPU can also run the subsystem correctly theoretically, but we don't provide the quick install scripts now, these scripts will appear on the GitHub repository later. Under the fast install scripts, the subsystem need the environment with docker latest version and docker compose latest version, Go 1.13.4 or higher, Node.js v8.9.4, GCC 7.3.0 and Hyperledger Fabric 2.0.0, the scripts on GitHub can help you to complete the installation procedure of the environment. 

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Another subsystem "Medical Data Blockchain Services Platform" can install with packages provided, the clients can run on Android, iOS, Mac OS, Windows and Linux.

### 2.5 Design and Implementation Constraints

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Firstly, the subsystem "Hospital-ledger", hardware requirements of the software project need a group of computers with disk space more than 40G free per computer with memory more than 2G per computer. The operating system requirement is Linux operating system CentOS 7 or Red Hat based on X86_64(amd64) CPU. Theoretically it can run on Ubuntu 18.04 LTS and other operating system based on Debian. Another subsystem run on cloud, the client need hardware based on Android, iOS, Mac OS, Windows or Linux.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Secondly, the software requirements of subsystem "Hospital-ledger", the one-click installations have wrote by Shell and it will consolidate into the managers' clients of "Medical Data Blockchain Services Platform" and this part will be controlled by the managers. The blockchain network employ the Hyperledger Fabric designed by IBM and Linux foundation, the main language is Golang. When it comes to the data of multimedia, the "Hospital-ledger" using the HDFS technology in Hadoop 2.0 to realize the peer-to-peer data storage of common users.

### 2.6 User Documentation

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The subsystem "Hospital-ledger" do not need any operations of users, it will work automatically with "Medical Data Blockchain Services Platform". There are four types of users in subsystem "Medical Data Blockchain Services Platform", there are introductions in details behind.

#### 2.6.1 Basic operations: Register, log in and query. 

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Register: After install the client on your devices, click the register button and submit the necessary information for verify your identity and wait the organization managers to verify, and then store the digital certificate send to you on your devices which is useful when log in.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Log in: After register on the client, the users can log in with the account, key and the digital certificate.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Query: After log in, there will show some important information of the users, the users can choose the query type and filter the data.

#### 2.6.2 Special Operations

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Doctor users, common users, organizations managers and platform managers special operations.

#### 2.6.3 Issues

### 2.7 Assumptions and Dependencies

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The subsystem "Hospital-ledger" depends on Hyperledger Fabric an open-source system designed by IBM, this tool is an open-source system, fabric project will maintain by Linux Foundation for a long time, it also depends on Hadoop Distributed Files System an open-source system designed by Apache.

## 3. External Interface Requirements

### 3.1 User Interfaces

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The "Hospital-ledger" do not have any GUI for users and the GUI of "Medical Data Blockchain Services Platform" looks like a search engine.

### 3.2 Hardware Interfaces

### 3.3 Software Interfaces

### 3.4 Communications Interfaces

## 4. System Features

## 5. Other Nonfunctional Requirements

### 5.1 Performance Requirements

### 5.2 Safety Requirements

### 5.3 Security Requirements

### 5.4 Software Quality Attributes

### 5.5 Business Rules

## 6. Other Requirements



```Python
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; [toc] Content &copy;
```
