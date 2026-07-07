---
atlas_id: AML.T0051.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2023-10-25"
description: Злоумышленник может внедрять промпты косвенно через отдельный канал данных, который обрабатывает LLM, например через текст или мультимедийные данные, полученные из баз данных или с веб-сайтов. Такие вредоносные...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 13
source_name: Indirect
subtechnique_count: 0
subtechnique_of: AML.T0051
tactics:
    - AML.TA0005
title: Косвенная промпт-инъекция
url: /techniques/AML.T0051.001/
---

Злоумышленник может внедрять промпты косвенно через отдельный канал данных, который обрабатывает LLM, например через текст или мультимедийные данные, полученные из баз данных или с веб-сайтов.

Такие вредоносные промпты могут быть скрыты или обфусцированы от пользователя. Этот тип инъекции может использоваться злоумышленником для закрепления в системе или для атаки на неосведомленного пользователя системы.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0005/"><span class="relation-id">AML.TA0005</span><strong>Выполнение</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0051/"><span class="relation-id">AML.T0051</span><strong>Промпт-инъекция в LLM</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0051.000/"><span class="relation-id">AML.T0051.000</span><strong>Прямая промпт-инъекция</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0051.002/"><span class="relation-id">AML.T0051.002</span><strong>Триггерная промпт-инъекция</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0024/"><span class="relation-id">AML.M0024</span><strong>Логирование телеметрии ИИ</strong><p>Логирование телеметрии может помочь выявить отправку небезопасных промптов в LLM.</p></a>
<a class="relation-item" href="/mitigations/AML.M0033/"><span class="relation-id">AML.M0033</span><strong>Валидация входных и выходных данных компонентов ИИ-агента</strong><p>Валидация может помешать злоумышленникам выполнять промпт-инъекции, способные повлиять на агентные рабочие процессы.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0020/"><span class="relation-id">AML.CS0020</span><strong>Угрозы косвенной промпт-инъекции: Bing Chat как похититель данных</strong><span class="relation-meta">Актор: Kai Greshake, Saarland University / Тактика: AML.TA0005 Выполнение</span><p>Если пользователь разрешил Bing Chat просматривать открытые сайты, чат может видеть их содержимое. Когда у пользователя открыт сайт злоумышленника, скрытая вредоносная инструкция попадает в Bing Chat и выполняется.</p></a>
<a class="relation-item" href="/studies/AML.CS0021/"><span class="relation-id">AML.CS0021</span><strong>Эксфильтрация разговоров ChatGPT</strong><span class="relation-meta">Актор: Embrace The Red / Тактика: AML.TA0005 Выполнение</span><p>Промпт-инъекция срабатывает и заставляет ChatGPT добавить Markdown-изображение с сервера злоумышленника. История переписки пользователя при этом встраивается в URL как query-параметр.</p></a>
<a class="relation-item" href="/studies/AML.CS0026/"><span class="relation-id">AML.CS0026</span><strong>Перехват финансовой транзакции с использованием M365 Copilot в роли инсайдера</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0005 Выполнение</span><p>Исследователи использовали промпт-инъекцию, чтобы при ответе LLM выполняла другие инструкции. Это происходит каждый раз, когда пользователь выполняет поиск и извлекается отравленная RAG-запись с промпт-инъекцией.</p></a>
<a class="relation-item" href="/studies/AML.CS0029/"><span class="relation-id">AML.CS0029</span><strong>Эксфильтрация разговоров Google Bard</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0005 Выполнение</span><p>Когда пользователь отправляет запрос, приводящий к извлечению документа, встроенный промпт выполняется. Вредоносный промпт заставляет Bard ответить Markdown-разметкой изображения, URL которого указывает на Google Apps Script исследователя и содержит разговор пользователя в параметре запроса.</p></a>
<a class="relation-item" href="/studies/AML.CS0035/"><span class="relation-id">AML.CS0035</span><strong>Эксфильтрация данных из Slack AI через косвенную промпт-инъекцию</strong><span class="relation-meta">Актор: PromptArmor / Тактика: AML.TA0005 Выполнение</span><p>Когда жертва просит Slack AI найти ее «EldritchNexus API key», Slack AI извлекает вредоносное содержимое и выполняет инструкции.</p></a>
<a class="relation-item" href="/studies/AML.CS0038/"><span class="relation-id">AML.CS0038</span><strong>Внедрение инструкций для отложенного автоматического вызова инструмента ИИ-агента</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0005 Выполнение</span><p>Когда пользователь просил Google Gemini кратко изложить письмо или как-либо с ним взаимодействовать, выполнялся вредоносный промпт.</p></a>
<a class="relation-item" href="/studies/AML.CS0039/"><span class="relation-id">AML.CS0039</span><strong>Living Off AI: промпт-инъекция через Jira Service Management</strong><span class="relation-meta">Актор: Cato CTRL / Тактика: AML.TA0005 Выполнение</span><p>В рамках стандартного рабочего процесса инженер поддержки в организации-жертве использовал Claude Sonnet, который может взаимодействовать с Jira через MCP-сервер Atlassian, чтобы помочь обработать вредоносное обращение. В результате инъекция была непреднамеренно выполнена.</p></a>
<a class="relation-item" href="/studies/AML.CS0040/"><span class="relation-id">AML.CS0040</span><strong>Взлом памяти ChatGPT с помощью промпт-инъекции</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0005 Выполнение</span><p>Когда пользователь ссылался на что-либо в общем документе, его содержимое добавлялось в контекст чата, и ChatGPT выполнял промпт.</p></a>
<a class="relation-item" href="/studies/AML.CS0045/"><span class="relation-id">AML.CS0045</span><strong>Эксфильтрация данных через MCP-сервер, используемый Cursor</strong><span class="relation-meta">Актор: Backslash Security Research Team / Тактика: AML.TA0005 Выполнение</span><p>При обработке вредоносного сайта MCP-сервер вернул внедренный промпт MCP-клиенту и отравил контекст LLM в Cursor. После этого Cursor выполнил промпт, встроенный в сайт.</p></a>
<a class="relation-item" href="/studies/AML.CS0046/"><span class="relation-id">AML.CS0046</span><strong>Уничтожение данных через косвенную промпт-инъекцию, нацеленную на Claude Computer Use</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0005 Выполнение</span><p>Когда пользователь попросил Claude взаимодействовать с PDF-файлом, встроенный промпт был выполнен.</p></a>
<a class="relation-item" href="/studies/AML.CS0048/"><span class="relation-id">AML.CS0048</span><strong>Публично доступные интерфейсы управления ClawdBot позволили получить учетные данные и выполнить команды</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0005 Выполнение</span><p>Исследователь смог напрямую отправлять промпты ClawdBot через интерфейс управления.</p></a>
<a class="relation-item" href="/studies/AML.CS0051/"><span class="relation-id">AML.CS0051</span><strong>Использование OpenClaw для командования и управления через промпт-инъекцию</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0005 Выполнение</span><p>OpenClaw выполнил промпт-инъекцию, встроенную во вредоносный сайт.</p></a>
<a class="relation-item" href="/studies/AML.CS0055/"><span class="relation-id">AML.CS0055</span><strong>AI ClickFix: захват управления computer-use-агентами с помощью ClickFix</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0005 Выполнение</span><p>Промпт инструктировал Computer Use Agent выполнить несколько действий: нажать &#34;Please see instructions to confirm&#34;, затем найти и нажать значок терминала, нажать `SHIFT+CTRL+V` и `RETURN`, после чего нажать кнопку &#34;OK&#34;.</p></a>
</div>
