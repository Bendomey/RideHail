# RideHail

Ride Hailing Service tryouts with golang, gorm, postgres, redis, gin, uber h3 🔥

# Services Development

Below is the flow for the development of the ridehailing service. We have 5 micro services to develop. Accounts service is to handle the management of users on the system, Listener Service is to listen to "online" drivers location every time it is logged to the service (Say 3 secs), Dispatch service, to help customers make requests and let drivers accept the request, and Miscellaneous is to handle some low level work. Or lets say assistive works for the other microservices

## Accounts

- [x] Create Super Admin
- [x] Create Admin
- [x] Login Admin
- [ ] Update Admin
- [ ] Update Admin Password
- [ ] Delete Admin
- [ ] Forgot Admin Password Functionality

- [ ] Create Customer
- [ ] Login Customer
- [ ] Update Customer
- [ ] Delete Customer

- [ ] Create Driver
- [ ] Approve Driver
- [ ] Login Driver
- [ ] Update Driver
- [ ] Update Driver Password
- [ ] Delete Driver
- [ ] Forgot Driver Password Functionality
- [ ] Update Driver Status from online to offline and vice versa

## Listener

- [ ] Get Long and Lat from online Drivers
- [ ] Get H3 geo-index at resolution 8
- [ ] Check if its the same thing as the old h3 in db and then add latlng to h3 index group in redis, else delete it from initial h3 index group and save in the new h3 index group. Afterwards, you save the new h3 in db.
- Get Drivers location to users so we can show on map [{driverId, lat, lng}]

## Dispatch

- [ ] Create Trip (request) with latlng plus id of customer, we convert to h3 and get that group.
- [ ] Populate redis data with googles duration and distance. Now we sort the data by distance, (Getting ETAs).
- [ ] We send requests to drivers. If one accepts, we create update Trip Request to waitingForPickup and update driver to engaged.
- [ ] Cancel A Trip (Request).
- [ ] Pay For A Trip

## Miscellaneous

- [ ] Create Ratings
- [ ] Create Favorites
- [ ] Create Faqs by Admin
- [ ] Create Enquiries
- [ ] Answer Enquiries by Admin
- [ ] Create Coupon Codes by Admin
- [ ] Use Coupon Codes
