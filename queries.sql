create table tblclient (
cid serial primary key not null,
firstname varchar(100) not null,
lastname varchar(100) not null,
contact varchar(100) not null,
username varchar(100) not null,
password varchar(100) not null,
date date not null
);

select * from tblclient;

insert into tblclient (firstname, lastname, contact, username, password, date)
values ('Cherrie', 'Arrogancia', '09104487623', 'aiai', 'arrogancia', '2020-07-24');
insert into tblclient (firstname, lastname, contact, username, password, date)
values ('Kimwen', 'Del Valle', '09104487623', 'kim', 'delvalle', '2020-07-24');
insert into tblclient (firstname, lastname, contact, username, password, date)
values ('Jamilla', 'Punzalan', '09104487623', 'jamilla', 'punzalan', '2020-07-24');


create table tblloan (
lid serial primary key not null,
amount numeric(15, 2) not null,
paymentno numeric(15) not null,
purpose varchar(100) not null,
pn varchar(100) not null,
dateapplied date not null,
encoded varchar(100) not null,
cid numeric(100) not null,
status varchar(100) not null
);
select * from tblloan;

insert into tblloan (amount, paymentno, purpose, pn, dateapplied, encoded, cid, status)
values ('5000', '5', 'Calamity', '12345', '2020-07-24', 'N/A', '1', 'Unassigned');
insert into tblloan (amount, paymentno, purpose, pn, dateapplied, encoded, cid, status)
values ('4000', '4', 'Sick', '12345', '2020-07-24', 'N/A', '2', 'Assigned');
insert into tblloan (amount, paymentno, purpose, pn, dateapplied, encoded, cid, status)
values ('3000', '3', 'Emergency', '12345', '2020-07-24', 'N/A', '3', 'Assigned');

create table tblinfo (
Iid serial primary key,
category varchar(100) not null,
pgroup varchar(100) not null,
product varchar(100) not null,
frequency varchar(100) not null,
guarantor varchar(100) not null,
lid numeric(15) not null
);

select * from tblinfo;

insert into tblinfo (category, pgroup, product, frequency, guarantor, lid)
values ('Lending', 'CB.OL.LOAN.GRP', 'CB.OL.LOAN.18', 'WEEKLY', 'NO', '5');

select * from tblloan inner join tblclient on tblloan.cid = tblclient.cid where tblclient.cid='1';


insert into tblloan (amount, paymentno, purpose, pn, dateapplied, encoded, status, category, pgroup, product, frequency, guarantor, cid )
values ('3000', '3', 'Emergency', '12345', '2020-07-24', 'N/A', 'Unassigned', 'Lending', 'CB.OL.LOAN.GRP', 'CB.OL.LOAN.18','Weekly', 'No', '3');

create table tbluser (
uid serial primary key,
fname varchar(100) not null,
lname varchar(100) not null,
username varchar(100) not null,
password varchar(100) not null,
branch varchar(100) not null,
insti varchar(100) not null
);