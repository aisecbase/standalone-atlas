# ATLAS на русском

Неофициальный статический русский справочник по MITRE ATLAS. Проект не является копией официального сайта MITRE, а строит собственный HTML на базе [Hugo](gohugo.io) из официального набора данных `mitre-atlas/atlas-data`.

Русский перевод создается и поддерживается проектом [aisecbase.ru](https://aisecbase.ru/).

## Источник данных

- Данные: https://github.com/mitre-atlas/atlas-data
- Официальный сайт: https://atlas.mitre.org/
- Лицензия данных: Apache License 2.0
- Лицензия кода, шаблонов, стилей и локальной документации: Apache License 2.0, см. `LICENSE`
- Текущий загруженный файл: `data/atlas/ATLAS.yaml`
- По умолчанию сборка берет release asset `ATLAS-YYYY.MM.yaml` из последнего release `mitre-atlas/atlas-data`. Для старых release без v6 asset остается fallback на legacy `dist/ATLAS.yaml`.

## Как устроено

- `data/atlas/ATLAS.yaml` - upstream-данные ATLAS. Начиная с ATLAS `v2026.05`, это YAML format v6 с `collection`, keyed maps и first-class `relationships`.
- `data/generated/atlas.yaml` - адаптированная структура для Hugo-шаблонов, которую пишет `cmd/atlasgen`.
- `data/translations/ru.yaml` - русский overlay по ID объектов.
- `cmd/atlasgen` - адаптирует legacy/v6 YAML, создает страницы объектов Hugo и пишет отчет по покрытию переводов.
- `layouts/` - шаблоны матрицы, списков, карточек объектов и графа знаний.

Переводы намеренно отделены от исходных данных. Если MITRE обновит технику или кейс, можно подтянуть новый `ATLAS.yaml`, снова сгенерировать страницы и увидеть, какие ID пока не переведены.

Сгенерированные файлы `data/generated/*.yaml`, `reports/translation-coverage.md` и `content/translation-coverage.md` коммитятся намеренно как snapshot для ревью изменений данных и переводов. Их нельзя править вручную; обновляй их через `go run ./cmd/atlasgen`.

## Для переводчиков

Переводы и исправления принимаются через pull request в этот репозиторий. Обычно вклад выглядит как правка `data/translations/ru.yaml` с последующей регенерацией страниц и отчетов.

Что переводить:

- названия и описания объектов ATLAS в `data/translations/ru.yaml`;
- статические страницы в `content/resources/*.md`;
- текст интерфейса в `layouts/` и `assets/css/site.css`, если он относится к отображению сайта.

Что не нужно менять вручную:

- `data/atlas/ATLAS.yaml` - это upstream-данные;
- `data/generated/atlas.yaml` - это адаптированные данные для Hugo, они пересоздаются генератором;
- сгенерированные страницы объектов в `content/tactics`, `content/techniques`, `content/mitigations`, `content/studies` - они пересоздаются генератором;
- `public/` - это результат сборки.

После правок переводов объектов запусти:

```bash
go run ./cmd/atlasgen
hugo --minify
```

После правок статических Markdown-страниц достаточно:

```bash
hugo --minify
```

Покрытие перевода можно проверить в `reports/translation-coverage.md` или на странице `/translation-coverage/`.

Если у перевода указаны `source.name_sha256`, `source.description_sha256` или `source.summary_sha256`, генератор сравнивает эти хеши с текущими исходными полями ATLAS. Хеши процедур кейсов и отдельных описаний `use` проверяются аналогично через вложенные `source.description_sha256` и `source.use_sha256`. При несовпадении объект попадает в колонку и раздел «Требует проверки». Переводы без `source.*_sha256` считаются legacy-переводами: они засчитываются, но автоматически не помечаются как устаревшие после изменения upstream-текста.

Чтобы посчитать хеш для конкретного объекта, используйте генератор:

```bash
go run ./cmd/atlasgen --source-hash AML.TA0000
```

Команда выведет готовый YAML-фрагмент `source:`. Его нужно добавить к переводу того же ID в `data/translations/ru.yaml`.

Если вы хотите добавить новую технику, кейс, меру защиты или изменить смысл оригинальных данных ATLAS, лучше отправлять вклад в официальный проект: https://atlas.mitre.org/resources/contribute. В этот репозиторий стоит отправлять именно перевод, локализацию и исправления русской версии.

## Другие языки

Проект можно использовать как основу для других локализаций ATLAS. Сейчас основной локальный язык - русский, а пункт `EN` в переключателе языка ведет на официальный англоязычный ATLAS.

Для нового языка нужно добавить отдельный overlay с переводами объектов:

```text
data/translations/<lang>.yaml
```

Формат такой же, как у `data/translations/ru.yaml`: ключи объектов ATLAS остаются оригинальными ID, а внутри переводятся `name`, `description`, `summary` и другие поддержанные поля.

Сгенерировать объектные страницы из другого overlay можно так:

```bash
ATLAS_LANG=es go run ./cmd/atlasgen
hugo --minify
```

Если файл лежит не в `data/translations/<lang>.yaml`, укажи путь явно:

```bash
ATLAS_TRANSLATION_FILE=data/translations/es-MX.yaml go run ./cmd/atlasgen
hugo --minify
```

Для полноценной локализации языка также нужно перевести статические страницы `content/resources/*.md`, меню и UI-строки в `hugo.yaml`, `layouts/` и `assets/css/site.css`. Такой язык можно публиковать отдельным доменом или отдельной веткой сборки. Если позже понадобится один мультиязычный Hugo-сайт с `/ru/`, `/es/` и другими префиксами, нужно будет дополнительно разнести контент и конфиг по языкам Hugo.

## Команды

```bash
make update-data
make generate
hugo server -D
```

Чтобы собрать сайт из последнего release данных ATLAS:

```bash
make build
```

В `hugo.yaml` задан базовый `baseURL`. Для production-домена передавай `HUGO_BASEURL`, чтобы sitemap, RSS, canonical и social URL собирались с правильным абсолютным адресом:

```bash
HUGO_BASEURL=https://example.org/ make build
```

## Аналитика

Google Analytics и Яндекс Метрика подключаются через `params.analytics` в `hugo.yaml`:

```yaml
params:
  analytics:
    google: "G-XXXXXXXXXX"
    yandex: "12345678"
```

Пустые значения отключают соответствующий счетчик. Код аналитики добавляется только в production-сборке Hugo.

Для Cloudflare Pages можно не коммитить реальные ID в `hugo.yaml`, а задать их переменными окружения:

```text
HUGO_PARAMS_ANALYTICS_GOOGLE=G-XXXXXXXXXX
HUGO_PARAMS_ANALYTICS_YANDEX=12345678
```

Чтобы собрать конкретный release:

```bash
ATLAS_RELEASE_TAG=v2026.06 make build
```

Чтобы пересобрать сайт из уже лежащего локально `data/atlas/ATLAS.yaml` без сетевого обновления:

```bash
UPDATE_ATLAS_DATA=0 make build
```

После `make generate` смотрите `reports/translation-coverage.md` или страницу `/translation-coverage/`: там видно, какие объекты полностью переведены, какие переведены частично, а какие еще не покрыты.

Генератор также восстанавливает связи, которые официальный сайт строит в клиентском store:

- тактика -> техники и подтехники;
- техника -> родительская техника и подтехники;
- техника -> меры защиты;
- техника -> примеры процедур из кейсов;
- mitigation -> техники;
- case study -> процедура, источники, reporter и дата с учетом гранулярности;
- ATT&CK references для тактик, техник и mitigations;
- platforms техник из ATLAS v6 для страниц техник и фильтра матрицы.

Готовая статика появится в `public/`.
