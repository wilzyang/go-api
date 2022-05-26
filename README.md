# Go CRUD API with Google Storage

This is a simple CRUD service with Google Storage. At first, the service used box.com, but after some consideration of box.com API usage is only for premium users only (which is not free), so I change the storage to Google Cloud Storage which is can be accessed by a free account. You could get a Google Cloud Platform (GCP) account free for six months and get some credits for testing.

The purpose of this API is to upload files to a bucket that also inserts the data to the database, there will be more additional functionality like deleting files and so on. The Pattern used is Hexagonal Pattern with a bootstrap to wrap things up. Still not been tested yet, and the unit test still needs to be added.

### Requirement
This Project Requirement is
- Go Version 1.6 above (its build at 1.8 but 1.6 is quite good enough)
- Mysql or Postgres database (right now I only add postgres connection, mysql will be added later)
- Google Cloud Platform account (you can register free account for 6 month which have some balance too, but please be careful and add some billing alert)


### Next Goal
- Refaktor to Google Storage from Box.com [Done]
- Test All Endpoints and debugging
- Create Makefile to easily install service to docker and set ENV
- Add redis/sessions to store jwt token for 30 minutes
- Add middleware for verify data request
- Adding Login Authentication
- Build unit test

### Change Log
#### [22-05-2022]
- Refaktor box.com API to Google Storage
- Box.com JWT auth and api still there, but other configuration at main already changed to Google Storage
- Modify database upload query and functions, this way we can manipulate data inside domain and doesn't need to touch repository file if there any changes on table column
- Pass success upload message to respond, so we can easily debug if there any error

#### [23-05-2022]
- Refaktor pattern and change some folder naming
- Debugging and fix errors
- Handling file name to url encode before insert to database
- Pass error message to response

#### [26-05-2022]
- Update handler not found
- Add Handlers to Delete Uploaded Files, this action also delete data from database [not tested yet]

### Dependency
Dependency for this project,
- Web Framework : Gin
- ORM           : GORM
- Redis         : TBD
- Config        : Viper
- File Storage  : Google Cloud Storage

### Reference
- Google Cloud Libraries: https://cloud.google.com/storage/docs/reference/libraries
- Google Cloud Storage API: https://cloud.google.com/storage/docs
- Gorm (database): https://gorm.io/docs
- Gin Web Framwork: https://github.com/gin-gonic/gin
- Viper Config: https://github.com/spf13/viper

## Contributing
Please note this code is not fully tested yet. If you want to ask or discuss anything please open an issue.

## License

Proprietary


