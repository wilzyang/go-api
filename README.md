# Go CRUD API with Box.com

This is a simple CRUD service with Box.com API. We can upload file which also insert the data to database, and delete a file. Pattern used is Hexagonal Pattern with a bootstrap to wrap things up. Still not tested yet, and unit still need to be added.

Next Goal
- Test Endpoints
- Add redis/sessions to store jwt token for 30 minutes
- Build unit test

### Dependency
Web Framework : Gin (https://github.com/gin-gonic/gin)
ORM           : GORM (https://gorm.io/docs)
Redis         : TBD
Config        : Viper (https://github.com/spf13/viper)


## Contributing
Pull request all welcome. For major changes, please open an issue first to discuss what would you want to change. Please note this code still not fully tested yet.

## License

Proprietary


