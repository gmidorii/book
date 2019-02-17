# DB

## Definition

```sql
DROP TABLE book.book_t;
CREATE TABLE book.book_t (
  isbn varchar(13) PRIMARY KEY,
  title varchar(256),
  abstract TEXT,
  registered_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) DEFAULT CHARACTER SET utf8;

DROP TABLE book.author_t;
CREATE TABLE book.author_t (
  id INT AUTO_INCREMENT,
  name VARCHAR(128),
  isbn VARCHAR(13),
  PRIMARY KEY(id, name, isbn)
) DEFAULT CHARACTER SET utf8;

DROP TABLE book.img_t;
CREATE TABLE book.img_t (
  id int AUTO_INCREMENT,
  isbn varchar(64),
  title varchar(256),
  url varchar(256),
  PRIMARY KEY(id, isbn)
) DEFAULT CHARACTER SET utf8;
```


## Sample
```
https://www.googleapis.com/books/v1/volumes?q=%E3%82%A8%E3%83%B3%E3%82%B8%E3%83%8B%E3%82%A2%E3%81%AE%E7%9F%A5%E7%9A%84%E7%94%9F%E7%94%A3%E8%A1%93
```

```sql
select
  books.isbn,
  books.title,
  books.abstract,
  author.name as author_name,
  img.title as img_title,
  img.url as img_url
from
 (
    select
      isbn, title, abstract
    from book.book_t
    LIMIT 0,10
) as books
inner join 
book.author_t as author
on
books.isbn = author.isbn
inner join
book.img_t as img
on
books.isbn = img.isbn
;
```