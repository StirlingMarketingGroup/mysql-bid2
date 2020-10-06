# MySQL bid2

A small MySQL UDF library for creating 8-byte time-oriented IDs, called "bid2"s, written in Golang.

---

### `bid2`

Returns a new bid2. The first 7 bytes of the ID are the unix timestamp in nano seconds, and the last byte is a crypto-random byte.

```sql
`bid2` (  )
```

## Examples

```sql
select`bid2`();
-- 0x163B688ADC795B68

select`bid2`();
-- 0x163B688C81BB80D9

```
---

## Dependencies

You will need Golang, which you can get from here https://golang.org/doc/install. You will also need the MySQL dev library.

Debian / Ubuntu
```shell
sudo apt update
sudo apt install libmysqlclient-dev
```
## Installing

You can find your MySQL plugin directory by running this MySQL query

```sql
select @@plugin_dir;
```

then replace `/usr/lib/mysql/plugin` below with your MySQL plugin directory.

```shell
cd ~ # or wherever you store your git projects
git clone https://github.com/StirlingMarketingGroup/mysql-bid2.git
cd mysql-bid2
go get -d ./...
go build -buildmode=c-shared -o bid2.so
sudo cp bid2.so /usr/lib/mysql/plugin/ # replace plugin dir here if needed
```

Enable the functions in MySQL by running this MySQL query

```sql
create function`bid2`returns string soname'bid2.so';
```