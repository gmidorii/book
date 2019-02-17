INSERT INTO book.book_t
  (isbn, title, abstract)
VALUE
  ('9784774198767', 'エンジニアの知的生産術', '仕事をするうえで、どのように学び、整理し、アウトプットするのか、エンジニアの知的生産の方法を解説した書籍です。知的生産の方法を解説した書籍のほとんどは執筆者の方法の紹介にとどまっており、各自の環境に合うようにどのようにカスタマイズするかまでは書かれていません。本書では、数々の知的生産術の共通点や特徴を知ることで、どこが重要な部分なのかを解説します。これにより、みなさんが自分の環境に合わせて手法を変化させ、取り入れられるようになることを目的とします。筆者が日ごろ行っている具体的な手法や試行錯誤も紹介します。');

INSERT INTO book.author_t
  (name, isbn)
VALUES
  ('西尾泰和', '9784774198767');

INSERT INTO book.img_t
  (isbn, title, url)
VALUES
  ('9784774198767', 'thumbnail', 'http://books.google.com/books/content?id=qFOHugEACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api');
