# Go CRUD API with Box.com

This is a simple CRUD service with Box.com API. We can upload file which also insert the data to database, and delete a file. Pattern used is Hexagonal Pattern with a bootstrap to wrap things up. Still not tested yet, and unit test still need to be added.

Next Goal
- Test All Endpoints
- Add redis/sessions to store jwt token for 30 minutes
- Adding Login Authentication
- Build unit test

### Dependency
Web Framework : Gin (https://github.com/gin-gonic/gin)
ORM           : GORM (https://gorm.io/docs)
Redis         : TBD
Config        : Viper (https://github.com/spf13/viper)


## Contributing
Please note this code still not fully tested yet. If want to ask something or request for pull request please open for issue.

## License

Proprietary


