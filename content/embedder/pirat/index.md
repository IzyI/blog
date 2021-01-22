---
title: "Sql получить непрерывную пооследовательных дат"
date: 2020-04-08T19:40:28+03:00
type: embedder
draft: false
linkhabr: false
description: SQL. MySQL POSTGRESQL, непрерывная последовательность дней
image: default_open_graph.jpg
tags: [си]
---
<p>
    Как то раз мне задали интересную задачу.
</p>
<p>
    Нужно получить в виде статистики юзеров разбитых по непрерывности их посещений.
    Например юзер посещал 3 дня подряд сайт каждый день по 2 раза потом день не посещал а потом сново поситил 5
    раз. <br>
    При этом нужно сделать фильтрацию по количеству посещений а также по количеству дней подряд.
</p>
<!--more-->
<p>
    Есть вот такая таблица история в бд MariaDB
</p>
{{< highlight sql  >}}
    create table history_tim
    (
    id        int auto_increment  primary key,
    name_user varchar(255)  null,
    date      datetime  null
    );
{{< /highlight >}}
<p>
    Недолго думаю я полез искать в наш великий интернет чтобы посмотреть если ли какие то решения.
    И наткнулся на вот эту статью:
    <a href="https://habr.com/ru/post/270573/" target="_blank">ТыЦ</a>>
</p>
<p>
    Но для меня это было не до конца решение так как в примере таблица была только для одного юзера у меня же в одной
    таблице вся история всех юзеров и статистику надо получить по всем сразу.

</p>
<p>
    Моим решением здесь было сделать поправку в дату через name_user который я преобразовал для начала в число d_rank с
    помощью команды RANK
</p>
{{< highlight sql  >}}
WITH timCTE
AS
(
    SELECT RANK() OVER (ORDER BY name_user) AS d_rank,
    phw.name_user,
    date(phw.date)       as date_d,
    COUNT(phw.name_user)             as count_user
    FROM history_tim as phw
    group by phw.name_user,  date(phw.date)
    ),
{{< /highlight >}}
<p>
    А потом добавил число d_rank в год даты date_rank . Впоследствии именно по ней и была группировка.
</p>
{{< highlight sql  >}}
    grp AS (
    SELECT ROW_NUMBER() OVER (ORDER BY DATE_ADD(date_d, interval d_rank year))                               AS row_num,
    DATE_ADD(date_d, interval d_rank year)                                                            as date_psi,
    DATE_ADD(date_d, interval -ROW_NUMBER() OVER (ORDER BY DATE_ADD(date_d, interval d_rank year)) day) AS date_rank,
    date_d,
    d_rank,
    count_user,
    name_user
    FROM timCTE
    )
{{< /highlight >}}

<p>
    В последствии я уже добавил в таблицу <b>grp</b> сколько минимум должен был юзер заходить каждый день в цепочке дней
    <em>например минимум 6 раз в день:</em>
</p>

{{< highlight sql  >}}
    FROM timCTE
    WHERE count_user>=6
    )
{{< /highlight >}}
<p>
    Также фильтр по длине цепочки дней <em>например минимум 8 раз подряд:</em>
</p>
{{< highlight sql  >}}
    SELECT name_user, MAX(date_d), MIN(date_d), COUNT(date_rank) as count_row, SUM(count_user) as count_mac
    FROM grp
    group by name_user, date_rank
    HAVING count_row>=8
    order by name_user
;
{{< /highlight >}}
<p>
    На выходе мы получаем статистику по юзерам которые были минимум 6 раз в день,
    в течении минимум 8 дней подряд на всем временном промежутке.
</p>
<p>
    Сразу решил написать чтобы более полно понимать что происходит лучше прочитать статью на хабре так как она более
    подробная:
    <a href="https://habr.com/ru/post/270573/" target="_blank">ТыЦ</a> <br>
    Потом же можно посмотреть полный листинг моего кода:
    <a href="https://gist.github.com/IzyI/076077081bcf9213e299d66396b8a3bd" target="_blank">ТыЦ</a>
</p>