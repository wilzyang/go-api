# Go CRUD API with Google Storage

This is a simple CRUD service with Google Storage. At first the project is using box.com, but after some consideration of box.com API usage is only for premium users only (which is not free), so I change the storage to Google Cloud Storage which is can be access by free account. You could get Google Cloud Platform (GCP) account free for 6 months and get some credits to for testing.

The purpose of this API is to upload file to bucket which also insert the data to database, there will be more additional functionality like delete files and so on. Pattern used is Hexagonal Pattern with a bootstrap to wrap things up. Still not tested yet, and unit test still need to be added.

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
Date 22-05-2022 
- Refaktor box.com API to Google Storage
- Box.com JWT auth and api still there, but other configuration at main already changed to Google Storage
- Modify database upload query and functions, this way we can manipulate data on domain side

### Dependency
Dependency for this project,
- Web Framework : Gin
- ORM           : GORM
- Redis         : TBD
- Config        : Viper
- File Storage  : Google Cloud Storage

### Reference
- Google Cloud OAuth Authentication: https://cloud.google.com/storage/docs/authentication
- Google Cloud Storage API: https://cloud.google.com/storage/docs/reference/libraries
- Gorm (database): https://gorm.io/docs
- Gin Web Framwork: https://github.com/gin-gonic/gin
- Viper Config: https://github.com/spf13/viper

## Contributing
Please note this code still not fully tested yet. If want to ask something or request for pull request please open for issue.

## License

Proprietary


