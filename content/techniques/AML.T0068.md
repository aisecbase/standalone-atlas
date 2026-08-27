---
atlas_id: AML.T0068
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-03-12"
description: Злоумышленники могут скрывать или иным образом обфусцировать промпт-инъекции либо содержимое для извлечения, чтобы избежать обнаружения людьми, защитными ограничениями большой языковой модели (LLM) или другими...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 10
source_name: LLM Prompt Obfuscation
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0007
title: Обфускация промпта LLM
url: /techniques/AML.T0068/
---

Злоумышленники могут скрывать или иным образом обфусцировать промпт-инъекции либо содержимое для извлечения, чтобы избежать обнаружения людьми, защитными ограничениями большой языковой модели (LLM) или другими механизмами обнаружения.

Для текстовых входных данных это может включать изменение способа отображения инструкций, например мелкий текст, текст того же цвета, что и фон, или скрытые HTML-элементы. Для мультимодальных входных данных вредоносные инструкции могут быть скрыты в самих данных, например в пикселях изображения, или в метаданных файла, например EXIF для изображений, ID3-тегах для аудио или метаданных документа.

Входные данные также могут быть замаскированы с помощью схемы кодирования, например base64 или rot13. Это может позволить обойти защитные ограничения LLM, выявляющие вредоносное содержимое, и затруднить распознавание такого содержимого как вредоносного для человека, участвующего в процессе.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0007/"><span class="relation-id">AML.TA0007</span><strong>Уклонение от защиты</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0020/"><span class="relation-id">AML.M0020</span><strong>Защитные ограничения (Guardrails) для генеративного ИИ</strong><p>Apply input guardrails that decode, normalize, inspect, and block concealed malicious instructions.</p></a>
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>Красная команда по ИИ</strong><p>Тестируйте закодированные, преобразованные, визуально скрытые, многоязычные и мультимодальные инструкции. Используйте случаи успешного обхода для совершенствования нормализации, декодирования, анализа содержимого и средств обнаружения.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0020/"><span class="relation-id">AML.CS0020</span><strong>Угрозы косвенной промпт-инъекции: Bing Chat как похититель данных</strong><span class="relation-meta">Актор: Kai Greshake, Saarland University / Тактика: AML.TA0007 Уклонение от защиты</span><p>Вредоносные инструкции были скрыты за счет нулевого размера шрифта, что затрудняло их обнаружение человеком.</p></a>
<a class="relation-item" href="/studies/AML.CS0026/"><span class="relation-id">AML.CS0026</span><strong>Перехват финансовой транзакции с использованием M365 Copilot в роли инсайдера</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0007 Уклонение от защиты</span><p>Чтобы получатель письма не заметил атаку, исследователи обфусцировали вредоносную часть письма.</p></a>
<a class="relation-item" href="/studies/AML.CS0040/"><span class="relation-id">AML.CS0040</span><strong>Взлом памяти ChatGPT с помощью промпт-инъекции</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0007 Уклонение от защиты</span><p>Исследователь поместил промпт в Google Doc, скрыв его в заголовке мелким шрифтом, цвет которого совпадал с цветом фона документа, чтобы сделать промпт невидимым.</p></a>
<a class="relation-item" href="/studies/AML.CS0041/"><span class="relation-id">AML.CS0041</span><strong>Бэкдор в файле правил: атака на цепочку поставки ИИ-ассистентов для программирования</strong><span class="relation-meta">Актор: Pillar Security / Тактика: AML.TA0007 Уклонение от защиты</span><p>Исследователи скрыли промпт в файле правил ИИ-ассистента для программирования с помощью невидимых Unicode-символов, таких как соединители нулевой ширины и маркеры двунаправленного текста. Промпт остается невидимым в редакторах кода и в процессе одобрения пул-реквестов на GitHub, что позволяет ему избегать обнаружения при ручной проверке. Видимая строка и скрытый</p></a>
<a class="relation-item" href="/studies/AML.CS0045/"><span class="relation-id">AML.CS0045</span><strong>Эксфильтрация данных через MCP-сервер, используемый Cursor</strong><span class="relation-meta">Актор: Backslash Security Research Team / Тактика: AML.TA0007 Уклонение от защиты</span><p>Вредоносный промпт был скрыт в HTML-теге `title` веб-страницы.</p></a>
<a class="relation-item" href="/studies/AML.CS0045/"><span class="relation-id">AML.CS0045</span><strong>Эксфильтрация данных через MCP-сервер, используемый Cursor</strong><span class="relation-meta">Актор: Backslash Security Research Team / Тактика: AML.TA0007 Уклонение от защиты</span><p>При обработке вредоносного сайта MCP-сервер вернул внедренный промпт MCP-клиенту и отравил контекст LLM в Cursor. Команда оболочки в промпте была скрыта с помощью кодирования base64, поэтому пользователю было сложнее понять, что может быть выполнено вредоносное действие.</p></a>
<a class="relation-item" href="/studies/AML.CS0046/"><span class="relation-id">AML.CS0046</span><strong>Уничтожение данных через косвенную промпт-инъекцию, нацеленную на Claude Computer Use</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0007 Уклонение от защиты</span><p>Вредоносная команда была обфусцирована с помощью кодирования base64 и ROT13. Промпт содержал инструкции для Claude по декодированию этой команды.</p></a>
<a class="relation-item" href="/studies/AML.CS0059/"><span class="relation-id">AML.CS0059</span><strong>EchoLeak: промпт-инъекция нулевого клика против M365 Copilot для эксфильтрации данных</strong><span class="relation-meta">Актор: Aim Labs / Тактика: AML.TA0007 Уклонение от защиты</span><p>Промпт был сформулирован как безобидный деловой текст, а не как очевидно вредоносная инструкция, чтобы не вызвать подозрений у пользователя.</p></a>
<a class="relation-item" href="/studies/AML.CS0061/"><span class="relation-id">AML.CS0061</span><strong>AI in the Middle: веб-сервисы ИИ как ретрансляторы C2</strong><span class="relation-meta">Актор: Check Point Research / Тактика: AML.TA0007 Уклонение от защиты</span><p>Когда некоторые промпты блокировались защитными механизмами модели, исследователи кодировали или шифровали данные полезной нагрузки в высокоэнтропийные блоки, чтобы снизить вероятность распознавания содержимого как вредоносного.</p></a>
<a class="relation-item" href="/studies/AML.CS0066/"><span class="relation-id">AML.CS0066</span><strong>ZombieAgent: атака на ChatGPT с эксфильтрацией данных</strong><span class="relation-meta">Актор: Radware Security Researchers / Тактика: AML.TA0007 Уклонение от защиты</span><p>Промпт-инъекция была визуально скрыта в содержимом, контролируемом внешней стороной, с помощью таких приёмов, как белый текст на белом фоне или микроскопический размер шрифта. ChatGPT мог обрабатывать инструкции, хотя они были незаметны пользователю.</p></a>
</div>
