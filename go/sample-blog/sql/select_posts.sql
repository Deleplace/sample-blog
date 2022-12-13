SELECT u.username, p.created, p.title, p.body
FROM post p 
JOIN user u ON u.id=p.author_id;