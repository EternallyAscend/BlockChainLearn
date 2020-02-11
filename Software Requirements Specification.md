# Software Requirements Specification

Version 1.0

</br></br>

<font size="4.5" weight="2"><b><center>Medical Data Blockchain Services Platform with Hospital-ledger</center></b></font>

</br></br>

<center><font size="4">Content</font></center>
[toc]

</br>Version 1.0 last edited by H.L. Yu on 2020-02-10 Mon.

</br><center>Nankai University 2020 &copy; All Rights Reserved.</center></br>

</br>

## 1. Introduction

### 1.1 Purpose

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Nowadays, the issues of treatments and the process of share and transfer medical data are still waiting to be solved. Fortunately, the security and the privacy of the bitcoin display well in the digital currency area,  it follows that the blockchain technology can also provide a trustworthy platform to store the medical data for common users, such as doctors and patients. This software project devote to develop a secure, shared and distributed system for storing medical data for doctors to learn the medical data history of the patients quickly without omission and provide an open medical data statistics platform for necessary medical data analyses or emergencies such as flus.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;This software project is an open-source system and it has two parts. One is "Hospital-ledger", another is "Medical Data Blockchain Services Platform". "Hospital-ledger" is provided for us to record the medical data among hospital workers and normal user securely. "Medical Data Blockchain Services Platform" is provided for us to manage the blockchain network and to exchange the  information and data. The "Hospital-ledger" is the basic of the project and "Medical Data Blockchain Services Platform".

### 1.2 Intended Audience and Reading Suggestions

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;This document is intended for developers, project managers, testers, users and document writers. All parts are provided for develops, project managers, testers and document writers. The part 2.4 Operating Environment and part 2.6 User Documentation are provided for users. These two part are provided for user to learn how to use the software.

### 1.3 Product Scope

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;This software project is designed for us to record the medical data of patients securely and share the data among different hospitals in order to give patient a better treatment, avoiding missing the medical history, reduce the risk of treatments and improve the relationship between the doctors and patients. As well as it will act as an open platform for data transparency without personal identity information in details for necessary medical data analyses or other uses such as the emergencies of flus.

### 1.4 References

1. Apache CouchDB. [http://couchdb.apache.org](http://couchdb.apache.org).
2. Apache Kafka. [http://kafka.apache.org](http://kafka.apache.org).
3. Hyperledger. [http://www.hyperledger.org](http://www.hyperledger.org).
4. Hyperledger Fabric. [http://github.com/hyperledger/fabric](http://github.com/hyperledger/fabric).
5. The Linux Foundation. [http://www.linuxfoundation.org](http://www.linuxfoundation.org).
6. Raft [https://raft.github.io](https://raft.github.io).
7. E. Androulaki, C. Cachin, C. Ferris, S. Muralidharan, C. Stathakopoulou: Hyperledger Fabric: A Distributed Operating System for Permissioned Blockchains. In EuroSys '18 April 23-26, 2018, Porto, Portugal.
8. H.Y. Li, L.H. Zhu, M. Shen,  F. Gao, X.L. Tao. S. Liu: Blockchain-Based Data Preservation System for Medical Data. In Springer Science+Business Media, LLC, part of Springer Nature 2018, (2018)42:141.

</br>

## 2. Overall Description

### 2.1 Product Perspective

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The software project includes two subsystems. One part is "Hospital-ledger", it provides a distributed network to store the medical data by blockchain, make the network lightweight and share the data securely between hospitals and patients by using the distributed alliance blockchain system Hyperledger Fabric which designed by IBM. Another part is "Medical Data Blockchain Services Platform", it provides a platform to create, manage, edit and destroy the blockchain network, give the authentications to doctors, common users and managers, edit the smart-contact, verify the data and transport the data.

### 2.2 Product Functions

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The first subsystem "Hospital-ledger" has basic functions for storing the medical data. It provides data blockchain services and data writing services for doctors, data querying services for both doctors and common users. 

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;On the other hand, the Client/Server(C/S) architecture software, the second subsystem "Medical Data Blockchain Services Platform" has basic blockchain services management functions such as create,edit and delete the organizations, create, edit, update and stop smart-contact(chain code in Hyperledger Fabric), create,edit, verify and delete the users, identify verifying and register services, information pushing services, data transporting services and digital certificates verifying services for blockchain network. 

### 2.3 User Classes and Characteristics

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;There are two types of users in subsystem "Hospital-ledger" and four types of users in subsystem "Medical Data Blockchain Services Platform". There only show the users of them without other services peers which will show in details in part 3 External Interface Requirements.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;In "Hospital-ledger", there are two types of users: doctors and patients, actually they log in with their account of "Medical Data Blockchain Services Platform" and then the system will connect to the relevant peer in blockchain network. The doctors belong to certain hospital and they can commit the request to write and query the medical data of patients, these operations will record on the "Medical Data Blockchain Services Platform", they also can pull requests of original authentications for themselves and reset the key of their account. By contrast, the patients only have the authority to query the data of themselves, pull requests of original authentications for themselves and reset the key of their accounts.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;In "Medical Data Blockchain Services Platform", there are four types of users: doctors, common users, organization managers and platform managers. There are still some other peers in this subsystem, such as the transfer peer and verify peer for the security of the medical data of common users. The doctors and common users (patients in "Hospital-ledger") in this subsystem can log in with their account key and digital certificate, then system will connect to relevant peer in blockchain network, and push the medical data return to their clients which they request, especially, the common users can download the multimedia medical data of themselves from the HDFS services on cloud. The managers, their authorities are higher than doctors and common users, they can log in with their account key and digital certificate, then the system will connect to the relevant peer of them. The managers of organizations can verify the identity of doctors and common users, send the digital certificate to them via the client, query medical data without personal information of common users, create and delete the organizations of doctors or common users,  create accounts and edit the accounts status of doctors and common users. The platform managers can verify the identity of the managers of organizations, monitor the status of the network system and the status of services platform, create, edit, update and stop the smart contract which named "chain code" in Hyperledger Fabric.

### 2.4 Operating Environment

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Now, the subsystem "Hospital-ledger" can run one-click setup on the Linux operating system CentOS 7 or Red Hat which based on CPU whose architecture is X86_64(amd64). The Ubuntu 18.04 LTS and other system which based on Debian with the same CPU can also run the subsystem correctly theoretically, but we don't provide the quick install scripts now, these scripts will appear on the GitHub repository later. Under the fast install scripts, the subsystem need the environment with docker latest version and docker compose latest version, Go 1.13.4 or higher, Node.js v8.9.4, GCC 7.3.0 and Hyperledger Fabric 2.0.0, the scripts on GitHub can help you to complete the installation procedure of the environment. 

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Another subsystem "Medical Data Blockchain Services Platform" can install with packages provided, the clients can run on Android, iOS, Mac OS, Windows and Linux.

### 2.5 Design and Implementation Constraints

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Firstly, the  hardware requirements of the subsystem "Hospital-ledger", the blockchain network need a group of computers with disk space more than 40G free per computer with memory more than 2G per computer to store the medical data of common users. The operating system requirement is Linux operating system CentOS 7 or Red Hat based on X86_64(amd64) CPU. Theoretically it can run on Ubuntu 18.04 LTS and other operating system based on Debian. Another subsystem run on cloud, the client need hardware based on Android, iOS, Mac OS, Windows or Linux. There still need some devices for running Hadoop Distributed Files System (HDFS).

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Secondly, the software requirements of the subsystem "Hospital-ledger", the one-click installations have wrote by Shell and it will consolidate into the managers' clients of "Medical Data Blockchain Services Platform" and this part will be controlled by the managers. The blockchain network employ the Hyperledger Fabric designed by IBM and Linux foundation, the main language is Golang. When it comes to the data of multimedia, the "Hospital-ledger" using the HDFS technology in Hadoop 2.0 to realize the peer-to-peer data storage of common users.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Thirdly, the hardware requirements of the subsystem "Medical Data Blockchain Services Platform", this system need some devices to verify the Hash value of the users and the data and to transport the data between "Hospital-ledger" network and users' clients, these system will isolate the blockchain network of data and the users to improve the security level.  

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Fourth, the software requirements of the subsystem "Medical Data Blockchain Services Platform", the client of desktop will designed by Qt and the mobile client will designed by Flutter.

### 2.6 User Documentation

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The subsystem "Hospital-ledger" do not need any operations of users, it will work automatically with "Medical Data Blockchain Services Platform", common users do not have the authority to operator the network as well. 

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;There are four types of users in subsystem "Medical Data Blockchain Services Platform", doctors, common users, organizations managers and platform managers, there are introductions in details behind. The client of users can be installed on most of major systems, including Windows, Mac OS, Linux, Android and iOS.

#### 2.6.1 Basic operations: Register, log in and query. 

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Register: After install the client on your devices, click the register button and submit the necessary information for verify your identity and wait the organization managers to verify, and then store the digital certificate send to you on your devices which is useful when log in.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Log in: After register on the client, the users can log in with the account, key and the digital certificate.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Query: After log in, there will show some important information of the users, the users can choose the query type and filter the data.

#### 2.6.2 Special Operations

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Doctor users can query the patients' medical data with the information of them, but these query operations will record on the network to reduce the risk of malicious information query, doctor users can also request the multimedia data of patients, these data will send to client via the network query by the peer of patient. 

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Common users can query the multimedia data of themselves and store these multimedia data on their devices. Common users have the authority to end the treatments of the specific doctor, avoid the medical data queried by doctors maliciously.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Organizations managers can manage the doctors and common users organizations, including verifying the identify of the doctors and common users, send the digital certificates of users which calculated by the blockchain network service, create, edit or stop the smart contact of every organization, create, edit or close the organizations of them.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Platform managers must be a trustworthy organization, for they have a lot of special operations to manage the whole network and the system. Platform managers can verify the identifies of organizations managers, check the status of organizations managers, check the status of the blockchain network, manage the services of the computer peers cluster and monitor the log status of the distributed network system. 

#### 2.6.3 Issues

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The probable issues may occurred when the user lose the digital certificate or account key. The users need to verify the identify again via the client, they can log in until the organizations mangers confirm the request, calculate the new Hash value and then resend digital certificate to the user through client.

### 2.7 Assumptions and Dependencies

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The subsystem "Hospital-ledger" depends on Hyperledger Fabric an open-source system designed by IBM, it also depends on Hadoop Distributed Files System an open-source system designed by Apache, these tools are both open-source system, these projects will maintain by Linux Foundation for a long time. The Hyperledger Fabric is mature platform of alliance chains of blockchain with high level security for the data. Another subsystem "Medical Data Blockchain Services Platform" depends on Nginx of Apache, Qt of Digia and Flutter frame of Google, these tools and frames are all open-source system.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The dependencies are open-source system, there is not much risk of using them. As well as that these dependencies are mainstream software developed by powerful company and organizations, the developers of them can provide support for a long time. The development prospect of this software project guaranteed by these basic tools and frames.

</br>

## 3. External Interface Requirements

### 3.1 User Interfaces

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Firstly, the GUI of the client. The "Hospital-ledger" do not have any GUI for users and the GUI of "Medical Data Blockchain Services Platform" looks like a search engine, provide a method to search the medical data of themselves with log in page for common users. The client provides other special operators interfaces under the search bar.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Secondly, the client interfaces for the diverse operators of different users: register, log in, log out, query,  display, write, request, verify, organizations, peers, users account, smart contact and blockchain network platform manage interfaces.  These interfaces will show in details behind.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Register interface: It provides functions for users to upload the information of their identity to relevant organization manager peer.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Log in interface: It provides functions for users to connect to the relevant peer of users, get data form the isolated computers at the instance of the security of medical data. 

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Log out interface: It provides functions for users to disconnect from the relevant peer.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Query interface: It provides functions for users to query the data from the relevant peer and pull request to Hadoop Files Distributed System to query the multimedia medical data of the users.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Display interface: It provides functions for the client to show most important abstract data or all data in details for users. There still some functions to show what the users search from summary data, text data or multimedia medical data.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Write interface: It provides functions for the doctors to write medical data to the hospital blockchain network and users organizations blockchain network, these operations will be monitored and recorded the Hash value of the operations by the "Hospital-ledger" and the public chain in the "Medical Data Blockchain Services Platform" cloud server. The medical data will endorsed in doctor's hospital organization and user's organization respectively to keep the security of the medical data. It will calculate the medical data into  a Hash value and endorse the record on public chain on "Medical Data Blockchain Services Platform" cloud servers.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Identify verify interface: It provides functions for the organizations managers to manage the identifies of the doctors and the common users. It will calculate the identify information and a random number assign to the account as a Hash value and make it into the digital certificates for the user to log in. When the user reset the key or request for recertification, it will recalculate the Hash value for new digital certificates.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Data verify interface: It provides functions for the users to verify the data which return form the relevant peers when query the data. When query the medical data for users, to verify the primitiveness of the data.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Organizations manage interface: It provides functions for the organizations managers to manage the peers and the size of the organizations. The organizations managers can create new organization when the users register and the latest organization is full, edit the state of the peers of the organizations when the doctors leave hospital or exchange work places and the state of common users change, delete the organization when the organization is empty and confirmed by most organizations managers.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Peers manage interface: It provides functions for the organizations managers to manage the state of the users. When the doctors exchange the work place or leave the medical area or the common users leave the world, the organizations managers can edit the state of them to make sure the normal operation of these blockchain system.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Users account manage interface: It provides functions for the organizations managers to manage the account information and authority of the users.

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Smart contact manage interface:

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Blockchain Platform manage interface:

### 3.2 Hardware Interfaces

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The subsystem "Medical Data Blockchain Services Platform" provides the hardware interfaces to manage the peers of "Hospital-ledger" blockchain network. These hardware interfaces provide functions to start, reboot and shut down the computers on cloud and these functions will be provided for organizations managers to set up the blockchain network.

### 3.3 Software Interfaces

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;There are a lot of interfaces to make sure the security of the medical data.

### 3.4 Communications Interfaces

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The client pull request to the devices on the cloud, there are the interfaces for transport the medical data and  between them.

</br>

## 4. System Features

</br>

## 5. Other Nonfunctional Requirements

### 5.1 Performance Requirements

### 5.2 Safety Requirements

### 5.3 Security Requirements

### 5.4 Software Quality Attributes

### 5.5 Business Rules

</br>

## 6. Other Requirements

### 6.1 Verify Requirements

### 6.2 Data Structure Requirements



```Python
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; [toc] Content &copy;
```