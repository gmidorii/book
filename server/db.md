# DB

## Definition

### Book

```sql
CREATE TABLE book.book_t (
  isbn varchar(13) PRIMARY KEY,
  title varchar(256),
  author varchar(256),
  abstract TEXT
) DEFAULT CHARACTER SET utf8;
```

### Img

```sql
CREATE TABLE book.img_t (
  id int AUTO_INCREMENT PRIMARY KEY,
  isbn varchar(64),
  title varchar(256),
  url varchar(256)
) DEFAULT CHARACTER SET utf8;
```