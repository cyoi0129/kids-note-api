DELETE FROM kids_users;
DELETE FROM kids_schools;
DELETE FROM kids_families;
DELETE FROM kids_kids;
DELETE FROM kids_tasks;
DELETE FROM kids_items;
INSERT INTO kids_users (email, password, name, gender, family)
VALUES (
        'user1@test.com',
        '9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08',
        'パパ',
        'MALE',
        1
    ),
    (
        'user2@test.com',
        '9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08',
        'ママ',
        'FEMALE',
        1
    );
INSERT INTO kids_schools (prefecture, city, type, name)
VALUES (
        '東京都',
        '品川区',
        '公立',
        '◯◯小学校'
    );
INSERT INTO kids_families (name)
VALUES ('◯◯家');
INSERT INTO kids_kids (name, birth, gender, family, school)
VALUES (
        '長女',
        '2015-09-01',
        'FEMALE',
        1,
        1
    ),
    (
        '次女',
        '2018-07-01',
        'FEMALE',
        1,
        1
    );
INSERT INTO kids_task_types (name, family)
VALUES ('宿題', 1),
    ('提出物', 1);
INSERT INTO kids_tasks (
        name,
        detail,
        types,
        status,
        update,
        due,
        items,
        kid,
        userId,
        family
    )
VALUES (
        '夏休みの宿題',
        '夏休みの宿題を終わらせること',
        ARRAY [1,2],
        'TODO',
        '2025-08-12',
        '2025-08-31',
        ARRAY [1],
        1,
        1,
        1
    ),
    (
        '夏休みの観察',
        'アサガオの日記作成',
        ARRAY [1,2],
        'DOING',
        '2025-08-22',
        '2025-08-31',
        ARRAY [2],
        2,
        2,
        1
    );
INSERT INTO kids_items (
        name,
        detail,
        type,
        image,
        kid,
        family
    )
VALUES (
        '問題集',
        'A4サイズの青色のやつ',
        '本',
        '',
        1,
        1
    ),
    (
        '盆栽',
        'アサガオの盆栽',
        'もの',
        '',
        2,
        1
    );
SELECT *
FROM kids_users;
SELECT *
FROM kids_schools;
SELECT *
FROM kids_families;
SELECT *
FROM kids_kids;
SELECT *
FROM kids_tasks;
SELECT *
FROM kids_items;