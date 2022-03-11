// START: Queries.
// Show all nodes with relationships.
match (n) return n;
// Show Items that relate to specific Order (by order id).
match (od:Order {id: 2})-[:CONTAINS]->(items) return od,items;
// Count price of specific order.
match (od:Order {id: 2})-[:CONTAINS]->(items) return sum(items.price);
// Find all orders of specific customer.
match (customer:Customer {id: 0})-[:BOUGHT]->(orders) return customer,orders;
// Find all items bought by specific customer.
match (customer:Customer {id: 0})-[:BOUGHT]->(orders)-[:CONTAINS]->(items) return customer,orders,items;
// Find number of items bought by specific customer.
match (customer:Customer {id: 0})-[:BOUGHT]->(orders)-[:CONTAINS]->(items) return customer,count(items);
// Count the amount payed by specific customer.
match (customer:Customer {id: 0})-[:BOUGHT]->(orders)-[:CONTAINS]->(items) return customer,sum(items.price);
// Count how many times every item was bought. Sort by this value.
match (it:Item)<-[:CONTAINS]-(orders) return it,count(orders) as num order by num desc;
// Show all items viewed by specific customer.
match (customer:Customer {id: 0})-[:VIEW]->(items) return customer,items;
// Show all other items that were bought together with specific item.
match (it:Item {id:1})<-[:CONTAINS]-(orders)-[:CONTAINS]->(items) return DISTINCT it, items;
// Find all customer that bought specific item.
match (it:Item {id:0})<-[:CONTAINS]-(orders)<-[:BOUGHT]-(customers) return it,customers;
// Find for specific customer items he viewed, but has not bought.
match (customer:Customer {id:0})-[:VIEW]->(itemsViewed)
where not (customer)-[:BOUGHT]->()-[:CONTAINS]->(itemsViewed)
return itemsViewed;
// END: Queries.
