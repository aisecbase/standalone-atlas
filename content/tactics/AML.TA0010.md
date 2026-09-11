---
atlas_id: AML.TA0010
atlas_type: tactic
attack_ref_id: TA0010
attack_ref_url: https://attack.mitre.org/tactics/TA0010/
created_date: "2022-01-24"
description: Злоумышленник пытается украсть ИИ-артефакты или другую информацию об ИИ-системе. Эксфильтрация включает техники, которые злоумышленники могут использовать для кражи данных из вашей сети. Данные могут быть украдены...
generated: true
generated_by: atlasgen
modified_date: "2025-04-09"
procedure_count: 32
source_name: Exfiltration
technique_count: 12
title: Эксфильтрация
url: /tactics/AML.TA0010/
---

Злоумышленник пытается украсть ИИ-артефакты или другую информацию об ИИ-системе.

Эксфильтрация включает техники, которые злоумышленники могут использовать для кражи данных из вашей сети. Данные могут быть украдены из-за их ценности как интеллектуальной собственности или для использования при подготовке будущих операций.

Техники получения данных из целевой сети обычно включают их передачу по каналу командования и управления злоумышленника или по альтернативному каналу, а также могут включать ограничение размера передачи.


## Техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0024/"><span class="relation-id">AML.T0024</span><strong>Эксфильтрация через API инференса ИИ</strong></a>
<a class="relation-item" href="/techniques/AML.T0024.000/"><span class="relation-id">AML.T0024.000</span><strong>Определение принадлежности к обучающей выборке</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0024.000/"><span class="relation-id">AML.T0024.000</span><strong>Определение принадлежности к обучающей выборке</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0024.001/"><span class="relation-id">AML.T0024.001</span><strong>Инверсия ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0024.001/"><span class="relation-id">AML.T0024.001</span><strong>Инверсия ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0024.002/"><span class="relation-id">AML.T0024.002</span><strong>Извлечение ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0024.002/"><span class="relation-id">AML.T0024.002</span><strong>Извлечение ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0025/"><span class="relation-id">AML.T0025</span><strong>Эксфильтрация киберсредствами</strong></a>
<a class="relation-item" href="/techniques/AML.T0056/"><span class="relation-id">AML.T0056</span><strong>Извлечение системного промпта LLM</strong></a>
<a class="relation-item" href="/techniques/AML.T0057/"><span class="relation-id">AML.T0057</span><strong>Утечка данных из LLM</strong></a>
<a class="relation-item" href="/techniques/AML.T0077/"><span class="relation-id">AML.T0077</span><strong>Рендеринг ответа LLM</strong></a>
<a class="relation-item" href="/techniques/AML.T0086/"><span class="relation-id">AML.T0086</span><strong>Эксфильтрация через вызов инструмента ИИ-агента</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0010/"><span class="relation-id">AML.CS0010</span><strong>Нарушение работы сервиса Microsoft Azure</strong><span class="relation-meta">Актор: Microsoft AI Red Team / Тактика: AML.TA0010 Эксфильтрация</span><p>Команда эксфильтровала модель и данные обычными киберсредствами.</p></a>
<a class="relation-item" href="/studies/AML.CS0015/"><span class="relation-id">AML.CS0015</span><strong>Компрометация цепочки зависимостей PyTorch</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0010 Эксфильтрация</span><p>Вся собранная информация, включая содержимое файлов, отправлялась через зашифрованные DNS-запросы на домен `*[dot]h4ck[dot]cfd` с использованием DNS-сервера `wheezy[dot]io`.</p></a>
<a class="relation-item" href="/studies/AML.CS0018/"><span class="relation-id">AML.CS0018</span><strong>Выполнение произвольного кода через Google Colab</strong><span class="relation-meta">Актор: Tony Piazza / Тактика: AML.TA0010 Эксфильтрация</span><p>Получив доступ к Google Drive, злоумышленник может открыть сервер для эксфильтрации частных данных или артефактов ML-моделей. В примере из исходной статьи показаны загрузка, установка и использование `ngrok`, серверного приложения, чтобы открыть доступный злоумышленнику URL к Google Drive жертвы и всем его файлам.</p></a>
<a class="relation-item" href="/studies/AML.CS0021/"><span class="relation-id">AML.CS0021</span><strong>Эксфильтрация разговоров ChatGPT</strong><span class="relation-meta">Актор: Embrace The Red / Тактика: AML.TA0010 Эксфильтрация</span><p>ChatGPT автоматически отображает изображение пользователю, из-за чего отправляет запрос на сервер злоумышленника и передает туда разговор пользователя.</p></a>
<a class="relation-item" href="/studies/AML.CS0023/"><span class="relation-id">AML.CS0023</span><strong>ShadowRay: захват кластеров Ray, доступных из интернета</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0010 Эксфильтрация</span><p>ИИ-артефакты, учетные данные и другая ценная информация могут быть эксфильтрованы киберсредствами. Исследователи обнаружили признаки reverse shell на уязвимых кластерах; такие оболочки могут использоваться для закрепления, продолжения выполнения произвольного кода и эксфильтрации данных.</p></a>
<a class="relation-item" href="/studies/AML.CS0024/"><span class="relation-id">AML.CS0024</span><strong>Червь Morris II: атака на основе RAG</strong><span class="relation-meta">Актор: Stav Cohen, Ron Bitton, Ben Nassi / Тактика: AML.TA0010 Эксфильтрация</span><p>Вредоносные инструкции в промпте заставляют сгенерированный ответ раскрывать чувствительные данные, например электронные письма, адреса и номера телефонов.</p></a>
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0010 Эксфильтрация</span><p>Обнаруженные учетные данные могли быть эксфильтрованы через имплант Sliver.</p></a>
<a class="relation-item" href="/studies/AML.CS0029/"><span class="relation-id">AML.CS0029</span><strong>Эксфильтрация разговоров Google Bard</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0010 Эксфильтрация</span><p>Bard автоматически отображает Markdown-разметку, отправляя запрос к Google Apps Script и тем самым эксфильтруя разговор пользователя. Content Security Policy Bard разрешает такой запрос, потому что URL размещен на домене, принадлежащем Google.</p></a>
<a class="relation-item" href="/studies/AML.CS0035/"><span class="relation-id">AML.CS0035</span><strong>Эксфильтрация данных из Slack AI через косвенную промпт-инъекцию</strong><span class="relation-meta">Актор: PromptArmor / Тактика: AML.TA0010 Эксфильтрация</span><p>В соответствии с вредоносными инструкциями ответ отображается как ссылка для перехода, в URL которой закодирован API-ключ жертвы. Фрагмент ответа: Жертву вводят в заблуждение: она думает, что должна нажать на ссылку для повторной аутентификации, после чего ее API-ключ отправляется на сервер под контролем злоумышленника.</p></a>
<a class="relation-item" href="/studies/AML.CS0037/"><span class="relation-id">AML.CS0037</span><strong>Эксфильтрация данных через инструменты ИИ-агента в Copilot Studio</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0010 Эксфильтрация</span><p>Промпт просит агента отправить результаты письмом на адрес, выбранный исследователями, с помощью почтового инструмента агента. Исследователи успешно эксфильтруют целевые данные через вызов инструмента.</p></a>
<a class="relation-item" href="/studies/AML.CS0039/"><span class="relation-id">AML.CS0039</span><strong>Living Off AI: промпт-инъекция через Jira Service Management</strong><span class="relation-meta">Актор: Cato CTRL / Тактика: AML.TA0010 Эксфильтрация</span><p>Вредоносный промпт предписывал опубликовать собранные сведения о тикетах в ответе к обращению. Это вызывало инструмент Atlassian MCP, который выполнял запрошенное действие и эксфильтровал данные в место, доступное исследователям на портале JSM.</p></a>
<a class="relation-item" href="/studies/AML.CS0043/"><span class="relation-id">AML.CS0043</span><strong>Прототип вредоносного ПО со встроенной промпт-инъекцией</strong><span class="relation-meta">Актор: Unknown Threat Actor / Тактика: AML.TA0010 Эксфильтрация</span><p>Вредоносное ПО Skynet настраивает Tor-прокси для эксфильтрации собранных файлов. Примечание: собранные файлы только выводились в stdout и фактически не были успешно эксфильтрованы.</p></a>
</div>


Показано 12 из 32 примеров.
