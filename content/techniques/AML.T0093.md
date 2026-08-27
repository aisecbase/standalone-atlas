---
atlas_id: AML.T0093
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-10-29"
description: Злоумышленник может внедрять вредоносные промпты в систему жертвы через публично доступное приложение, рассчитывая, что в будущем они будут обработаны ИИ и в итоге окажут последующее воздействие. Это может...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 14
source_name: Prompt Infiltration via Public-Facing Application
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0004
    - AML.TA0006
title: Внедрение промпта через публичное приложение
url: /techniques/AML.T0093/
---

Злоумышленник может внедрять вредоносные промпты в систему жертвы через публично доступное приложение, рассчитывая, что в будущем они будут обработаны ИИ и в итоге окажут последующее воздействие. Это может происходить, когда источник данных индексируется системой RAG (генерации, дополненной извлечением), когда правило запускает действие ИИ-агента или когда пользователь использует большую языковую модель (LLM) для взаимодействия с вредоносным содержимым. Вредоносные промпты могут сохраняться в системе жертвы длительное время и влиять на нескольких пользователей и разные ИИ-инструменты внутри организации-жертвы.

Целью может стать любое публично доступное приложение, принимающее текстовый ввод. Сюда входят электронная почта, системы совместной работы с документами, такие как OneDrive или Google Drive, и сервис-дески или тикет-системы, такие как Jira. Также сюда относится внедрение через OCR, когда вредоносные инструкции встроены в изображения, скриншоты и счета, которые попадают в систему.

Злоумышленники могут выполнять [разведку](/tactics/AML.TA0002), чтобы выявить публично доступные приложения, которые, вероятно, отслеживаются ИИ-агентом или индексируются RAG. Они могут выполнять [выявление конфигурации ИИ-агента](/techniques/AML.T0084), чтобы уточнить выбор целей.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0004/"><span class="relation-id">AML.TA0004</span><strong>Первичный доступ</strong></a>
<a class="relation-item" href="/tactics/AML.TA0006/"><span class="relation-id">AML.TA0006</span><strong>Закрепление</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0016/"><span class="relation-id">AML.CS0016</span><strong>Выполнение кода в MathGPT через промпт-инъекцию</strong><span class="relation-meta">Актор: Ludwig-Ferdinand Stumpp / Тактика: AML.TA0004 Первичный доступ</span><p>Это показало, что исследователь может использовать уязвимость к промпт-инъекции в модели GPT-3, применяемой в MathGPT, как вектор первичного доступа.</p></a>
<a class="relation-item" href="/studies/AML.CS0026/"><span class="relation-id">AML.CS0026</span><strong>Перехват финансовой транзакции с использованием M365 Copilot в роли инсайдера</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0004 Первичный доступ</span><p>Исследователи отправили пользователю в организации-жертве письмо с вредоносным пейлоадом, используя знание о том, что все полученные письма попадают в RAG-базу Copilot.</p></a>
<a class="relation-item" href="/studies/AML.CS0029/"><span class="relation-id">AML.CS0029</span><strong>Эксфильтрация разговоров Google Bard</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0004 Первичный доступ</span><p>Исследователь делится с целевым пользователем Google Doc, содержащим вредоносный промпт. В атаке используется то, что расширения Bard позволяют Bard обращаться к документам пользователя.</p></a>
<a class="relation-item" href="/studies/AML.CS0037/"><span class="relation-id">AML.CS0037</span><strong>Эксфильтрация данных через инструменты ИИ-агента в Copilot Studio</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0004 Первичный доступ</span><p>Исследователи отправляют письмо с вредоносным промптом в почтовый ящик, который, по их предположению, может управляться ИИ-агентом.</p></a>
<a class="relation-item" href="/studies/AML.CS0038/"><span class="relation-id">AML.CS0038</span><strong>Внедрение инструкций для отложенного автоматического вызова инструмента ИИ-агента</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0004 Первичный доступ</span><p>Исследователь включил вредоносный промпт в тело длинного письма, отправленного жертве.</p></a>
<a class="relation-item" href="/studies/AML.CS0039/"><span class="relation-id">AML.CS0039</span><strong>Living Off AI: промпт-инъекция через Jira Service Management</strong><span class="relation-meta">Актор: Cato CTRL / Тактика: AML.TA0004 Первичный доступ</span><p>Исследователи создали новое обращение с вредоносным промптом на публичном портале Jira Service Management (JSM) жертвы, выявленном во время разведки.</p></a>
<a class="relation-item" href="/studies/AML.CS0040/"><span class="relation-id">AML.CS0040</span><strong>Взлом памяти ChatGPT с помощью промпт-инъекции</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0004 Первичный доступ</span><p>К Google Doc был предоставлен общий доступ для жертвы, что сделало документ доступным для ChatGPT через функцию Connected App.</p></a>
<a class="relation-item" href="/studies/AML.CS0040/"><span class="relation-id">AML.CS0040</span><strong>Взлом памяти ChatGPT с помощью промпт-инъекции</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0006 Закрепление</span><p>Промпт-инъекция для отравления памяти сохраняется в общем Google Doc, откуда она может распространяться на других пользователей и сессии чата. Это затрудняет отслеживание источников воспоминаний и их удаление.</p></a>
<a class="relation-item" href="/studies/AML.CS0046/"><span class="relation-id">AML.CS0046</span><strong>Уничтожение данных через косвенную промпт-инъекцию, нацеленную на Claude Computer Use</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0004 Первичный доступ</span><p>Исследователи встроили вредоносный промпт в PDF-документ. Такой документ мог попасть в систему жертвы через публично доступное приложение, например электронную почту или общее хранилище документов.</p></a>
<a class="relation-item" href="/studies/AML.CS0059/"><span class="relation-id">AML.CS0059</span><strong>EchoLeak: промпт-инъекция нулевого клика против M365 Copilot для эксфильтрации данных</strong><span class="relation-meta">Актор: Aim Labs / Тактика: AML.TA0004 Первичный доступ</span><p>Исследователи отправили письмо в почтовый ящик пользователя Microsoft 365.</p></a>
<a class="relation-item" href="/studies/AML.CS0060/"><span class="relation-id">AML.CS0060</span><strong>Межсайтовый скриптинг (XSS) через манипуляцию промптом в ИИ-чат-боте Lenovo</strong><span class="relation-meta">Актор: Cybernews Research Team / Тактика: AML.TA0004 Первичный доступ</span><p>Исследователи внедрили подконтрольный злоумышленнику HTML в рабочий процесс поддержки Lenovo, отправив промпт Lena через публичный чат-интерфейс. В результате сгенерированная нагрузка была сохранена в истории чата для последующего отображения.</p></a>
<a class="relation-item" href="/studies/AML.CS0063/"><span class="relation-id">AML.CS0063</span><strong>Атаки на Gemini с помощью промптов в приглашениях Google Calendar</strong><span class="relation-meta">Актор: SafeBreach Research Team / Тактика: AML.TA0004 Первичный доступ</span><p>Исследователи отправили отравленное приглашение Calendar или электронное письмо с вредоносными инструкциями в заголовке приглашения или теме письма.</p></a>
<a class="relation-item" href="/studies/AML.CS0066/"><span class="relation-id">AML.CS0066</span><strong>ZombieAgent: атака на ChatGPT с эксфильтрацией данных</strong><span class="relation-meta">Актор: Radware Security Researchers / Тактика: AML.TA0004 Первичный доступ</span><p>Исследователи отправили вредоносное письмо в подключённый к ChatGPT почтовый ящик либо предоставили вредоносный документ, который можно было загрузить в ChatGPT или получить через подключённый сервис. В результате промпт оказался в источнике данных, к которому имел доступ ИИ-агент жертвы.</p></a>
<a class="relation-item" href="/studies/AML.CS0067/"><span class="relation-id">AML.CS0067</span><strong>Claude Code GitHub Action Secret Exposure</strong><span class="relation-meta">Актор: Microsoft Defender Security Research Team / Тактика: AML.TA0004 Первичный доступ</span><p>The researchers introduced the malicious prompt through attacker-controlled GitHub content processed by the lab workflow, modeling delivery through an issue body, pull request description, or comment handled by Claude Code Action.</p></a>
</div>
