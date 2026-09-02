---
atlas_id: AML.TA0005
atlas_type: tactic
attack_ref_id: TA0002
attack_ref_url: https://attack.mitre.org/tactics/TA0002/
created_date: "2022-01-24"
description: Злоумышленник пытается запустить вредоносный код, встроенный в ИИ-артефакты или программное обеспечение. Выполнение включает техники, которые приводят к запуску кода, контролируемого злоумышленником, в локальной или...
generated: true
generated_by: atlasgen
modified_date: "2025-04-09"
procedure_count: 65
source_name: Execution
technique_count: 20
title: Выполнение
url: /tactics/AML.TA0005/
---

Злоумышленник пытается запустить вредоносный код, встроенный в ИИ-артефакты или программное обеспечение.

Выполнение включает техники, которые приводят к запуску кода, контролируемого злоумышленником, в локальной или удаленной системе. Техники запуска вредоносного кода часто комбинируются с техниками из других тактик для достижения более широких целей, например исследования сети или кражи данных. Например, злоумышленник может использовать средство удаленного доступа, чтобы запустить скрипт PowerShell, выполняющий [выявление удаленных систем](https://attack.mitre.org/techniques/T1018/).


## Техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0011/"><span class="relation-id">AML.T0011</span><strong>Запуск пользователем</strong></a>
<a class="relation-item" href="/techniques/AML.T0011.000/"><span class="relation-id">AML.T0011.000</span><strong>Небезопасные ИИ-артефакты</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0011.000/"><span class="relation-id">AML.T0011.000</span><strong>Небезопасные ИИ-артефакты</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0011.001/"><span class="relation-id">AML.T0011.001</span><strong>Вредоносный пакет</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0011.001/"><span class="relation-id">AML.T0011.001</span><strong>Вредоносный пакет</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0011.002/"><span class="relation-id">AML.T0011.002</span><strong>Отравленный инструмент ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0011.002/"><span class="relation-id">AML.T0011.002</span><strong>Отравленный инструмент ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0011.003/"><span class="relation-id">AML.T0011.003</span><strong>Вредоносная ссылка</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0011.003/"><span class="relation-id">AML.T0011.003</span><strong>Вредоносная ссылка</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0050/"><span class="relation-id">AML.T0050</span><strong>Интерпретатор команд и сценариев</strong></a>
<a class="relation-item" href="/techniques/AML.T0051/"><span class="relation-id">AML.T0051</span><strong>Промпт-инъекция в LLM</strong></a>
<a class="relation-item" href="/techniques/AML.T0051.000/"><span class="relation-id">AML.T0051.000</span><strong>Прямая промпт-инъекция</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0051.000/"><span class="relation-id">AML.T0051.000</span><strong>Прямая промпт-инъекция</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0051.001/"><span class="relation-id">AML.T0051.001</span><strong>Косвенная промпт-инъекция</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0051.001/"><span class="relation-id">AML.T0051.001</span><strong>Косвенная промпт-инъекция</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0051.002/"><span class="relation-id">AML.T0051.002</span><strong>Триггерная промпт-инъекция</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0051.002/"><span class="relation-id">AML.T0051.002</span><strong>Триггерная промпт-инъекция</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0053/"><span class="relation-id">AML.T0053</span><strong>Вызов инструментов ИИ-агента</strong></a>
<a class="relation-item" href="/techniques/AML.T0100/"><span class="relation-id">AML.T0100</span><strong>Кликбейт для ИИ-агента</strong></a>
<a class="relation-item" href="/techniques/AML.T0103/"><span class="relation-id">AML.T0103</span><strong>Развертывание ИИ-агента</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0016/"><span class="relation-id">AML.CS0016</span><strong>Выполнение кода в MathGPT через промпт-инъекцию</strong><span class="relation-meta">Актор: Ludwig-Ferdinand Stumpp / Тактика: AML.TA0005 Выполнение</span><p>Исследователь вручную составлял состязательные промпты, чтобы проверить, уязвима ли модель к промпт-инъекции и действительно ли приложение напрямую выполняет код, сгенерированный GPT-3.</p></a>
<a class="relation-item" href="/studies/AML.CS0016/"><span class="relation-id">AML.CS0016</span><strong>Выполнение кода в MathGPT через промпт-инъекцию</strong><span class="relation-meta">Актор: Ludwig-Ferdinand Stumpp / Тактика: AML.TA0005 Выполнение</span><p>Исследователь получил возможность выполнения кода, поскольку LLM была подключена к интерпретатору Python. Исследователь мог косвенно выполнить произвольный код в Python-интерпретаторе приложения, если специально подготовленными промптами заставлял LLM сгенерировать этот код.</p></a>
<a class="relation-item" href="/studies/AML.CS0018/"><span class="relation-id">AML.CS0018</span><strong>Выполнение произвольного кода через Google Colab</strong><span class="relation-meta">Актор: Tony Piazza / Тактика: AML.TA0005 Выполнение</span><p>Пользователь-жертва может неосознанно выполнить вредоносный код, входящий в состав скомпрометированного Colab notebook. Вредоносный код может быть обфусцирован или скрыт в других файлах, которые notebook загружает при выполнении.</p></a>
<a class="relation-item" href="/studies/AML.CS0020/"><span class="relation-id">AML.CS0020</span><strong>Угрозы косвенной промпт-инъекции: Bing Chat как похититель данных</strong><span class="relation-meta">Актор: Kai Greshake, Saarland University / Тактика: AML.TA0005 Выполнение</span><p>Если пользователь разрешил Bing Chat просматривать открытые сайты, чат может видеть их содержимое. Когда у пользователя открыт сайт злоумышленника, скрытая вредоносная инструкция попадает в Bing Chat и выполняется.</p></a>
<a class="relation-item" href="/studies/AML.CS0021/"><span class="relation-id">AML.CS0021</span><strong>Эксфильтрация разговоров ChatGPT</strong><span class="relation-meta">Актор: Embrace The Red / Тактика: AML.TA0005 Выполнение</span><p>Промпт-инъекция срабатывает и заставляет ChatGPT добавить Markdown-изображение с сервера злоумышленника. История переписки пользователя при этом встраивается в URL как query-параметр.</p></a>
<a class="relation-item" href="/studies/AML.CS0022/"><span class="relation-id">AML.CS0022</span><strong>Галлюцинация пакетов ChatGPT</strong><span class="relation-meta">Актор: Vulcan Cyber, Lasso Security / Тактика: AML.TA0005 Выполнение</span><p>В итоге пользователь загрузит вредоносный пакет, что позволит выполнить произвольный код.</p></a>
<a class="relation-item" href="/studies/AML.CS0024/"><span class="relation-id">AML.CS0024</span><strong>Червь Morris II: атака на основе RAG</strong><span class="relation-meta">Актор: Stav Cohen, Ron Bitton, Ben Nassi / Тактика: AML.TA0005 Выполнение</span><p>Исследователи тестируют промпты через публичные API моделей, чтобы найти работающие промпт-инъекции.</p></a>
<a class="relation-item" href="/studies/AML.CS0024/"><span class="relation-id">AML.CS0024</span><strong>Червь Morris II: атака на основе RAG</strong><span class="relation-meta">Актор: Stav Cohen, Ron Bitton, Ben Nassi / Тактика: AML.TA0005 Выполнение</span><p>Когда почтовый ассистент извлекает письмо с червем при очередной генерации ответа, промпт-инъекция меняет поведение GenAI-почтового ассистента.</p></a>
<a class="relation-item" href="/studies/AML.CS0024/"><span class="relation-id">AML.CS0024</span><strong>Червь Morris II: атака на основе RAG</strong><span class="relation-meta">Актор: Stav Cohen, Ron Bitton, Ben Nassi / Тактика: AML.TA0005 Выполнение</span><p>Исследователи отправляют на адрес в целевой почтовой системе письмо с состязательным самореплицирующимся промптом, или «ИИ-червем». GenAI-почтовый ассистент автоматически обрабатывает это письмо в рамках обычной генерации предлагаемого ответа. Письмо сохраняется в базе данных для RAG (генерации, дополненной извлечением), что компрометирует RAG-систему.</p></a>
<a class="relation-item" href="/studies/AML.CS0026/"><span class="relation-id">AML.CS0026</span><strong>Перехват финансовой транзакции с использованием M365 Copilot в роли инсайдера</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0005 Выполнение</span><p>Исследователи использовали промпт-инъекцию, чтобы при ответе LLM выполняла другие инструкции. Это происходит каждый раз, когда пользователь выполняет поиск и извлекается отравленная RAG-запись с промпт-инъекцией.</p></a>
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0005 Выполнение</span><p>Когда любой пользователь позже загрузит модель, она автоматически выполнит полезную нагрузку злоумышленника.</p></a>
<a class="relation-item" href="/studies/AML.CS0029/"><span class="relation-id">AML.CS0029</span><strong>Эксфильтрация разговоров Google Bard</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0005 Выполнение</span><p>Когда пользователь отправляет запрос, приводящий к извлечению документа, встроенный промпт выполняется. Вредоносный промпт заставляет Bard ответить Markdown-разметкой изображения, URL которого указывает на Google Apps Script исследователя и содержит разговор пользователя в параметре запроса.</p></a>
</div>


Показано 12 из 65 примеров.
