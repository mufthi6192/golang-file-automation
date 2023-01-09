
# Golang File Automation

This app is the tools for handle and simpilify your file export using MySQL.
This app This app running on CLI which build with Go, Gorm, MySQL, and YAML configuration. 
I hope you helped with this app 🙌🙌🙌

## Authors

- [Mufthi Ryanda](https://www.instagram.com/mufthi_ryanda)


## Basic Configuration

To use this app you need to setup configuration file. Here the step for use this app.


- Setup database on ```config.yaml```

```yaml
db_host : "YOUR MYSQL HOST" {Optional}
db_port : "YOUR MYSQL PORT" {Usually 3306}
db_name : "YOUR DATABASE NAME"
db_username : "YOUR DATABASE USERNAME"
db_password : "YOUR DATABASE PASSWORD" {Optional}
```
- Setup file path on ```config.yaml```
```yaml
current_path : "/public/assets/images"
destination_path : "YOUR DESTINATION PATH"
```
You can change the current path without paste the root path. You just
need to change the current path into your path that you make on the app.

## Custom Configuration
We know that your table structure not exactly same with me.
So here the things, how to custom the mysql table structure

- Declaring model on ```app/entity/model/product.go```
```go
type TableName struct{
ColumnName     data_type  `json:"column_name"`
}
```

Or you can make your own model.

- Change model on ```app/services/file/file.go```
```go
var testingModel = model.ProductModels{
ColumnName = Value			
}
```

## Support

If you have a quetion with this app, feel free to contact me on 
[Instagram](https://www.instagram.com/mufthi_ryanda).

