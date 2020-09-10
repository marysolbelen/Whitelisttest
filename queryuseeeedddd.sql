drop table tblclient;
create table tblclient (
area varchar(100) ,
bday text,
cid varchar(100),
centername varchar(100),
contact varchar(100) ,
flag varchar(100) ,
mlength varchar(100) ,
membername varchar(100),
bcode varchar(100) ,
newcid varchar(100) ,
rdate text,
sn varchar(100) ,
unit varchar(100)
);


create table tblclient (
Area varchar(100) not null,
Birthday text not null,
CID varchar(100) not null,
CenterName varchar(100) not null,
Contact varchar(100) not null,
Flag varchar(100) not null,
LengthOfMembership varchar(100) not null,
MemberName varchar(100) not null,
NewBranchCode varchar(100) not null,
NewCID varchar(100) not null,
RecognizedDate text not null,
SN varchar(100) not null,
Unit varchar(100) not null
);
/*Used table*/
create table tblclient (
Area varchar(100) not null,
Birthday date not null,
CID varchar(100) not null,
CenterName varchar(100) not null,
Contact varchar(100) not null,
Flag varchar(100) not null,
LengthOfMembership numeric(15, 2) not null,
MemberName varchar(100) not null,
NewBranchCode int not null,
NewCID varchar(100) not null,
RecognizedDate date not null,
SN int not null,
Unit varchar(100) not null
);