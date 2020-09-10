create table tbluser (
uid serial primary key,
fname varchar(100) not null,
lname varchar(100) not null,
username varchar(100) not null,
password varchar(100) not null,
branch varchar(100) not null,
insti varchar(100) not null
);

create table tblclient (
	flag varchar(10) not null,
	sn varchar(20) not null,	
	area varchar(100) not null,
	unit varchar(100) not null,
	cid varchar(20) primary key,
	centername varchar(100) not null,
	membername varchar(100) not null,
	rdate text not null,
	bday text not null,
	mlength numeric(15, 2) not null,
	bcode varchar(20) not null,
	newcid varchar(20) not null,
	contact varchar(20) not null
);

create table tblloan (
lid serial primary key not null,
amount numeric(15, 2) not null,
paymentno numeric(15) not null,
purpose varchar(100) not null,
pn varchar(100) not null,
dateapplied date not null,
encoded varchar(100) not null,
status varchar(100) not null,
category varchar(100) not null,
pgroup varchar(100) not null,
product varchar(100) not null,
frequency varchar(100) not null,
guarantor varchar(100) not null,
cid varchar(100) not null
);

select * from tbluser;
select * from tblloan;
select * from tblclient;
truncate table tblclient;

insert into tblclient (flag, sn, area, unit, cid, centername, membername, rdate, bday, mlength, bcode, newcid, contact)
values ('1', '43', 'CARD Inc. Agusan Del Sur 1', 'BAROBO 1', '400610',	'QUEENS - Barobo', 'Cario, Leonie L.', '05/06/10',	'11/12/81', '10.16', '1193', '1193-0040-06104',	'09102756782');
insert into tblclient (flag, sn, area, unit, cid, centername, membername, rdate, bday, mlength, bcode, newcid, contact)
values( '1', '62', 'CARD Inc. Agusan Del Sur 1', 'BAROBO 1', '400724',	 'SAMPAGUITA- Barobo', 'Malbas, Felecidad', '05/24/10', '03/06/54', '10.11', '1193', '1193-0040-07243', '09309198745');
insert into tblclient (flag, sn, area, unit, cid, centername, membername, rdate, bday, mlength, bcode, newcid, contact)
values ('1', '68', 'CARD Inc. Agusan Del Sur 1', 'BAROBO 1', '400878',	 'TALISAY 2- Barobo', 'Frasco, Salome M.', '04/29/10', '10/30/78',	'10.18', '1193', '1193-0040-08787', '09516840489');
insert into tblclient (flag, sn, area, unit, cid, centername, membername, rdate, bday, mlength, bcode, newcid, contact)
values ('1', '11', 'CARD Inc. Agusan Del Sur 1','BAROBO 1',	'400124',	'BAYBAY - Barobo',	'Paninsoro, Snooky P.', '04/27/10', '07/24/81', '10.18', '1193', '1193-0040-01246', '09104008811');
insert into tblclient (flag, sn, area, unit, cid, centername, membername, rdate, bday, mlength, bcode, newcid, contact)
values ('1', '12', 'CARD Inc. Agusan Del Sur 1', 'BAROBO 1', '400171',	 'TERMINAL - Barobo', 'Henobiagon, Mary Ann P.',  '06/10/09', '01/04/83', '11.06', '1193', '1193-0040-01717', '09072864633');
insert into tblclient (flag, sn, area, unit, cid, centername, membername, rdate, bday, mlength, bcode, newcid, contact)
values ('1', '29', 'CARD Inc. Agusan Del Sur 1', 'BAROBO 1', '400370', 'RIZAL - Barobo', 'Diaz, Ma. Victoria A.', '10/05/09', '04/08/62', '10.74',	'1193', '1193-0040-03705', '09124581610');
insert into tblclient (flag, sn, area, unit, cid, centername, membername, rdate, bday, mlength, bcode, newcid, contact)
values ('1', '30', 'CARD Inc. Agusan Del Sur 1', 'BAROBO 1', '400376',	'RIZAL - Barobo', 'Abrao, Josephine D.', '10/05/09', '11/05/67', '10.74', '1193', '1193-0040-03762', '09514106994');
insert into tblclient (flag, sn, area, unit, cid, centername, membername, rdate, bday, mlength, bcode, newcid, contact)
values ('1', '32', 'CARD Inc. Agusan Del Sur 1', 'BAROBO 1', '400386', 'RIZAL - Barobo', 'Camporedondo, Delia T.', '12/14/09', '07/06/64', '10.55', '1193', '1193-0040-03861', '09488490624');
insert into tblclient (flag, sn, area, unit, cid, centername, membername, rdate, bday, mlength, bcode, newcid, contact)
values ('1', '34', 'CARD Inc. Agusan Del Sur 1', 'BAROBO 1', '400389',	'RIZAL - Barobo', 'Payapaya, Marcelina C.', '01/11/10', '02/05/65', '10.47', '1193', '1193-0040-03895', '09101996903');
insert into tblclient (flag, sn, area, unit, cid, centername, membername, rdate, bday, mlength, bcode, newcid, contact)
values ('1', '43', 'CARD Inc. Agusan Del Sur 1', 'BAROBO 1', '400610',	'QUEENS - Barobo', 'Cario, Leonie L.', '05/06/10',	'11/12/81', '10.16', '1193', '1193-0040-06104',	'09102756782');


insert into tblloan (amount, paymentno, purpose, pn, dateapplied, encoded, status, category, pgroup, product, frequency, guarantor, cid)
values ('5000', '5', 'Personal Consumption', '0012345', '2020-07-01', 'N/A', 'Unassigned', 'Lending','CB.OL.LOAN.GRP', 'CB.OL.LOAN.18', 'Weekly', 'No', '400124');

insert into tblloan (amount, paymentno, purpose, pn, dateapplied, encoded, status, category, pgroup, product, frequency, guarantor, cid)
values ('5000', '5', 'Personal Consumption', '0012345', '2020-07-01', 'N/A', 'Unassigned', 'Lending','CB.OL.LOAN.GRP', 'CB.OL.LOAN.18', 'Weekly', 'No', '400171');

insert into tblloan (amount, paymentno, purpose, pn, dateapplied, encoded, status, category, pgroup, product, frequency, guarantor, cid)
values ('5000', '5', 'Personal Consumption', '0012345', '2020-07-01', 'N/A', 'Unassigned', 'Lending','CB.OL.LOAN.GRP', 'CB.OL.LOAN.18', 'Weekly', 'No', '400370');

insert into tblloan (amount, paymentno, purpose, pn, dateapplied, encoded, status, category, pgroup, product, frequency, guarantor, cid)
values ('5000', '5', 'Personal Consumption', '0012345', '2020-07-01', 'N/A', 'Unassigned', 'Lending','CB.OL.LOAN.GRP', 'CB.OL.LOAN.18', 'Weekly', 'No', '400376');

insert into tblloan (amount, paymentno, purpose, pn, dateapplied, encoded, status, category, pgroup, product, frequency, guarantor, cid)
values ('5000', '5', 'Personal Consumption', '0012345', '2020-07-01', 'N/A', 'Unassigned', 'Lending','CB.OL.LOAN.GRP', 'CB.OL.LOAN.18', 'Weekly', 'No', '400386');

insert into tblloan (amount, paymentno, purpose, pn, dateapplied, encoded, status, category, pgroup, product, frequency, guarantor, cid)
values ('5000', '5', 'Personal Consumption', '0012345', '2020-07-01', 'N/A', 'Unassigned', 'Lending','CB.OL.LOAN.GRP', 'CB.OL.LOAN.18', 'Weekly', 'No', '400389');

insert into tblloan (amount, paymentno, purpose, pn, dateapplied, encoded, status, category, pgroup, product, frequency, guarantor, cid)
values ('5000', '5', 'Personal Consumption', '0012345', '2020-07-01', 'N/A', 'Unassigned', 'Lending','CB.OL.LOAN.GRP', 'CB.OL.LOAN.18', 'Weekly', 'No', '400610');

insert into tblloan (amount, paymentno, purpose, pn, dateapplied, encoded, status, category, pgroup, product, frequency, guarantor, cid)
values ('5000', '5', 'Personal Consumption', '0012345', '2020-07-01', 'N/A', 'Unassigned', 'Lending','CB.OL.LOAN.GRP', 'CB.OL.LOAN.18', 'Weekly', 'No', '400724');


insert into tbluser (fname, lname, username, password, branch, insti) values
('Marysol', 'Belen', 'Marysol', 'marysol4', 'San Pablo City', 'FDSAP');

insert into tblloan (amount, paymentno, purpose, pn, dateapplied, encoded, status, category, pgroup, product, frequency, guarantor, cid)
values ('5000', '5', 'Personal Consumption', '0012345', '2020-07-01', 'N/A', 'Unassigned', 'Lending','CB.OL.LOAN.GRP', 'CB.OL.LOAN.18', 'Weekly', 'No', '400878');
