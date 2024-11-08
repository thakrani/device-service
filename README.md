
**Objective**
Design a RESTful service to manage a device database, allowing operations such as adding, retrieving, updating, and deleting devices.

**Design Patterns**
Repository Pattern: To abstract the data layer, making it easier to test and maintain.
Service Layer: A separate service layer to handle business logic, which will interact with the repository.

**Endpoints**:

- POST /devices: Add a new device.
- GET /devices/{id}: Retrieve a device by its identifier.
- GET /devices: List all devices.
- PUT /devices/{id}: Update a device (both full and partial updates).
- DELETE /devices/{id}: Delete a device.
- GET /devices/search/{brand}: Search for devices by brand.


**Curls**:

- Add Device:
curl -X POST -H "Content-Type: application/json" -d '{"device_name": "Phone X", "device_brand": "BrandA"}' http://localhost:8080/devices

- List Devices:
curl -X GET http://localhost:8080/devices

- Get Device by ID:
curl -X GET http://localhost:8080/devices/{id}

- Update Device:
curl -X PUT -H "Content-Type: application/json" -d '{"device_name": "Phone Y"}' http://localhost:8080/devices/{id}

- Delete Device:
curl -X DELETE http://localhost:8080/devices/{id}

- Search Device by Brand:
curl -X GET "http://localhost:8080/devices/search/{brand}"

**Mocks**
cd repository
mockery --name=IDeviceRepository --case=snake --output=../mocks --outpkg=mocks
cd ..
cd services
mockery --name=IDeviceService --case=snake --output=../mocks --outpkg=mocks
cd ..
cd db
mockery --name=IDB --case=snake --output=../mocks --outpkg=mocks
cd ..

**Install Postgres**
brew services start postgresql@15
brew install postgresql@15
brew services start postgresql@15
sudo -u postgres psql
psql postgres
CREATE ROLE postgres WITH LOGIN SUPERUSER CREATEDB CREATEROLE REPLICATION PASSWORD 'yourpassword';
Create Table DB script in db folder

Update username and password in config.json to connect to db.

FINAL COMMIT 18:19:00 Nov 7