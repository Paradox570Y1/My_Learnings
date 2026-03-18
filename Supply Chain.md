supply - order place , land - how it flows

example
need 100 bouquet flower how it reaches user , that's our responsibility
earlier our vertical doesn't exist, it was scattered earlier
Now we have managed to put in more than 70 % of scattered into current system

CEO
Saurav (CTO)
Himanshu (AVP)
Divyansh 
Navdeep
Amrit (SDE3)
Rohit (SDE2)


# Functions

Micro Services
- fnp-identity (RBA- role based access) all users comes through this
- fnp-so (sale order)
- fnp-fullfillment (manufacturing order, last delivery boy logic)
- fnp-ops (purchase order creations, most crucial, manages components, catalog to select quantitites, supplier sends bills GRN (Goods reciept Notes), match if items are correct or not)
- fnp-stats (kafka oriented reporting, wastage, purchase events)
- fnp-go-kit  (powerhouse powering all the cervice, libraries with prebuilt services)
- fnp-locations (facility location, coco, fofo, cfc)
- fnp-auditlogs (tags audit)

components - raw materials -> makes product
(flowers, paper, ribbion)  ->  bouquet
our inventory manages components


Databases - Mongo, MySQL, Redis, mongoose -> atlas, Cloud Bucket