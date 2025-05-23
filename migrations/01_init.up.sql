CREATE EXTENSION IF NOT EXISTS pg_trgm;


CREATE TABLE facts (
                       id SERIAL PRIMARY KEY,
                       fact TEXT NOT NULL,
                       normalized_fact TEXT NOT NULL
);

CREATE INDEX facts_normalized_fact_trgm_idx
    ON facts USING gin (normalized_fact gin_trgm_ops);


INSERT INTO facts (fact, normalized_fact) VALUES ('Александр Македонский никогда не проигрывал ни одного сражения.', 'александр македонский никогда не проигрывал ни одного сражения');
INSERT INTO facts (fact, normalized_fact) VALUES ('Ганнибал пересёк Альпы с боевыми слонами.', 'ганнибал пересёк альпы с боевыми слонами');
INSERT INTO facts (fact, normalized_fact) VALUES ('Пётр I ввёл налог на бороды в России.', 'пётр i ввёл налог на бороды в россии');
INSERT INTO facts (fact, normalized_fact) VALUES ('Папа Римский Стефан VI провёл "Кадверный синод", суд над трупом предыдущего папы.', 'папа римский стефан vi провёл кадверный синод суд над трупом предыдущего папы');
INSERT INTO facts (fact, normalized_fact) VALUES ('Великая Китайская стена видна из космоса — миф, её не видно невооружённым глазом.', 'великая китайская стена видна из космоса миф её не видно невооружённым глазом');
INSERT INTO facts (fact, normalized_fact) VALUES ('В Древней Спарте детей с физическими недостатками сбрасывали со скалы.', 'в древней спарте детей с физическими недостатками сбрасывали со скалы');
INSERT INTO facts (fact, normalized_fact) VALUES ('Джордж Вашингтон отказался от короны и стал первым президентом США.', 'джордж вашингтон отказался от короны и стал первым президентом сша');
INSERT INTO facts (fact, normalized_fact) VALUES ('Последняя гильотина во Франции была использована в 1977 году.', 'последняя гильотина во франции была использована в 1977 году');
INSERT INTO facts (fact, normalized_fact) VALUES ('Сталин готовился сбежать из ссылки на лыжах.', 'сталин готовился сбежать из ссылки на лыжах');
INSERT INTO facts (fact, normalized_fact) VALUES ('Византийская империя просуществовала на тысячу лет дольше Западной Римской.', 'византийская империя просуществовала на тысячу лет дольше западной римской');
INSERT INTO facts (fact, normalized_fact) VALUES ('Наполеон однажды был атакован кроликами во время охоты.', 'наполеон однажды был атакован кроликами во время охоты');
INSERT INTO facts (fact, normalized_fact) VALUES ('В Средние века кошек массово уничтожали, что способствовало распространению чумы.', 'в средние века кошек массово уничтожали что способствовало распространению чумы');
INSERT INTO facts (fact, normalized_fact) VALUES ('Мумии в Древнем Египте иногда использовались в медицине как порошок.', 'мумии в древнем египте иногда использовались в медицине как порошок');
INSERT INTO facts (fact, normalized_fact) VALUES ('Царь Иван Грозный убил своего сына.', 'царь иван грозный убил своего сына');
INSERT INTO facts (fact, normalized_fact) VALUES ('Первые Олимпийские игры проводились нагими.', 'первые олимпийские игры проводились нагими');
INSERT INTO facts (fact, normalized_fact) VALUES ('В Англии 19 века была мода есть мумии.', 'в англии 19 века была мода есть мумии');
INSERT INTO facts (fact, normalized_fact) VALUES ('Самый короткий конфликт в истории — Англо-занзибарская война (38 минут).', 'самый короткий конфликт в истории англозанзибарская война 38 минут');
INSERT INTO facts (fact, normalized_fact) VALUES ('Черепаха, участвовавшая в наполеоновских войнах, умерла в 2006 году.', 'черепаха участвовавшая в наполеоновских войнах умерла в 2006 году');
INSERT INTO facts (fact, normalized_fact) VALUES ('Адольф Гитлер не умел водить машину.', 'адольф гитлер не умел водить машину');
INSERT INTO facts (fact, normalized_fact) VALUES ('Банк Англии однажды был спасён женщиной, спрятавшей золото в подвале.', 'банк англии однажды был спасён женщиной спрятавшей золото в подвале');
INSERT INTO facts (fact, normalized_fact) VALUES ('Линкольн, Кеннеди: оба были убиты в пятницу, выстрелом в голову, и сменивший их — Джонсон.', 'линкольн кеннеди оба были убиты в пятницу выстрелом в голову и сменивший их джонсон');
INSERT INTO facts (fact, normalized_fact) VALUES ('Папирусы из Геркуланума до сих пор не расшифрованы.', 'папирусы из геркуланума до сих пор не расшифрованы');
INSERT INTO facts (fact, normalized_fact) VALUES ('Клеопатра жила ближе к изобретению iPhone, чем к строительству пирамид.', 'клеопатра жила ближе к изобретению iphone чем к строительству пирамид');
INSERT INTO facts (fact, normalized_fact) VALUES ('СССР первым отправил женщину в космос.', 'ссср первым отправил женщину в космос');
INSERT INTO facts (fact, normalized_fact) VALUES ('В Японии периода Эдо христианство каралось смертной казнью.', 'в японии периода эдо христианство каралось смертной казнью');
INSERT INTO facts (fact, normalized_fact) VALUES ('Первая бомба в ВМВ упала на немецкую ферму, убив единственную корову.', 'первая бомба в вмв упала на немецкую ферму убив единственную корову');
INSERT INTO facts (fact, normalized_fact) VALUES ('В Третьем рейхе была запрещена джазовая музыка.', 'в третьем рейхе была запрещена джазовая музыка');
INSERT INTO facts (fact, normalized_fact) VALUES ('Римский император Калигула назначил своего коня консулом.', 'римский император калигула назначил своего коня консулом');
INSERT INTO facts (fact, normalized_fact) VALUES ('Первая женщина-доктор в Европе выдавала себя за мужчину.', 'первая женщинадоктор в европе выдавала себя за мужчину');
INSERT INTO facts (fact, normalized_fact) VALUES ('В Исландии в 1973 году извержение вулкана спасло гавань от заиливания.', 'в исландии в 1973 году извержение вулкана спасло гавань от заиливания');
INSERT INTO facts (fact, normalized_fact) VALUES ('Ложки в древнем Китае делались из нефрита.', 'ложки в древнем китае делались из нефрита');
INSERT INTO facts (fact, normalized_fact) VALUES ('Крестовые походы включали детский поход, окончившийся трагедией.', 'крестовые походы включали детский поход окончившийся трагедией');
INSERT INTO facts (fact, normalized_fact) VALUES ('Советский лётчик Гагарин получил кодовую фразу "Поехали" перед стартом.', 'советский лётчик гагарин получил кодовую фразу поехали перед стартом');
INSERT INTO facts (fact, normalized_fact) VALUES ('В эпоху Возрождения моду задавали чумные маски.', 'в эпоху возрождения моду задавали чумные маски');
INSERT INTO facts (fact, normalized_fact) VALUES ('Леонардо да Винчи писал зеркально, справа налево.', 'леонардо да винчи писал зеркально справа налево');
INSERT INTO facts (fact, normalized_fact) VALUES ('Японский солдат Хиро Онода продолжал воевать до 1974 года.', 'японский солдат хиро онода продолжал воевать до 1974 года');
INSERT INTO facts (fact, normalized_fact) VALUES ('Авраам Линкольн держал на рабочем столе череп.', 'авраам линкольн держал на рабочем столе череп');
INSERT INTO facts (fact, normalized_fact) VALUES ('Римляне использовали мочу для чистки одежды.', 'римляне использовали мочу для чистки одежды');
INSERT INTO facts (fact, normalized_fact) VALUES ('Фараоны иногда женились на своих сёстрах.', 'фараоны иногда женились на своих сёстрах');
INSERT INTO facts (fact, normalized_fact) VALUES ('Викинги носили не рогатые шлемы, а простые железные.', 'викинги носили не рогатые шлемы а простые железные');
INSERT INTO facts (fact, normalized_fact) VALUES ('Уинстон Черчилль получил Нобелевскую премию по литературе, не мира.', 'уинстон черчилль получил нобелевскую премию по литературе не мира');
INSERT INTO facts (fact, normalized_fact) VALUES ('Великая депрессия вызвала всплеск самоубийств в США.', 'великая депрессия вызвала всплеск самоубийств в сша');
INSERT INTO facts (fact, normalized_fact) VALUES ('США планировали взорвать Луну в 1950-х годах для демонстрации силы.', 'сша планировали взорвать луну в 1950х годах для демонстрации силы');
INSERT INTO facts (fact, normalized_fact) VALUES ('Эйнштейну предложили пост президента Израиля — он отказался.', 'эйнштейну предложили пост президента израиля он отказался');
INSERT INTO facts (fact, normalized_fact) VALUES ('В СССР существовали "почтовые ящики" — секретные города без названий.', 'в ссср существовали почтовые ящики секретные города без названий');
INSERT INTO facts (fact, normalized_fact) VALUES ('Византийцы использовали греческий огонь — вещество, горящее на воде.', 'византийцы использовали греческий огонь вещество горящее на воде');
INSERT INTO facts (fact, normalized_fact) VALUES ('В Афинах граждан можно было изгнать на 10 лет — остракизм.', 'в афинах граждан можно было изгнать на 10 лет остракизм');
INSERT INTO facts (fact, normalized_fact) VALUES ('Японцы во Вторую мировую применяли воздушные шары с бомбами, перелетавшими океан.', 'японцы во вторую мировую применяли воздушные шары с бомбами перелетавшими океан');
INSERT INTO facts (fact, normalized_fact) VALUES ('В Древнем Риме гладиаторы часто не убивали друг друга — это было невыгодно.', 'в древнем риме гладиаторы часто не убивали друг друга это было невыгодно');
INSERT INTO facts (fact, normalized_fact) VALUES ('Первая бумага в Европе появилась благодаря арабам, захватившим китайскую технологию.', 'первая бумага в европе появилась благодаря арабам захватившим китайскую технологию');

