---
atlas_id: AML.T0078
atlas_type: technique
attack_ref_id: T1189
attack_ref_url: https://attack.mitre.org/techniques/T1189/
created_date: "2025-04-16"
description: Злоумышленники могут получить доступ к ИИ-системе, когда пользователь посещает сайт в ходе обычного веб-серфинга или когда ИИ-агент получает информацию из интернета от имени пользователя. Сайты могут содержать...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 4
source_name: Drive-by Compromise
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0004
title: Компрометация при посещении сайта
url: /techniques/AML.T0078/
---

Злоумышленники могут получить доступ к ИИ-системе, когда пользователь посещает сайт в ходе обычного веб-серфинга или когда ИИ-агент получает информацию из интернета от имени пользователя. Сайты могут содержать [промпт-инъекцию LLM](/techniques/AML.T0051), которая при выполнении способна изменить поведение ИИ-модели.

Тот же подход может использоваться для доставки других типов вредоносного кода, не нацеленных напрямую на ИИ (см. [Drive-by Compromise в ATT&CK](https://attack.mitre.org/techniques/T1189/)).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0004/"><span class="relation-id">AML.TA0004</span><strong>Первичный доступ</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0021/"><span class="relation-id">AML.CS0021</span><strong>Эксфильтрация разговоров ChatGPT</strong><span class="relation-meta">Актор: Embrace The Red / Тактика: AML.TA0004 Первичный доступ</span><p>Когда запрос пользователя заставляет ChatGPT открыть эту веб-страницу через плагин `WebPilot`, модель считывает промпт злоумышленника.</p></a>
<a class="relation-item" href="/studies/AML.CS0045/"><span class="relation-id">AML.CS0045</span><strong>Эксфильтрация данных через MCP-сервер, используемый Cursor</strong><span class="relation-meta">Актор: Backslash Security Research Team / Тактика: AML.TA0004 Первичный доступ</span><p>Когда пользователь попросил Cursor использовать MCP-инструмент для извлечения содержимого вредоносного сайта, промпт был получен и добавлен в контекст Cursor.</p></a>
<a class="relation-item" href="/studies/AML.CS0051/"><span class="relation-id">AML.CS0051</span><strong>Использование OpenClaw для командования и управления через промпт-инъекцию</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0004 Первичный доступ</span><p>Когда жертва попросила OpenClaw кратко пересказать `https://openclaw.aisystem.tech`, промпт-инъекция была получена с сайта через навык OpenClaw `web_fetch`.</p></a>
<a class="relation-item" href="/studies/AML.CS0055/"><span class="relation-id">AML.CS0055</span><strong>AI ClickFix: захват управления computer-use-агентами с помощью ClickFix</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0004 Первичный доступ</span><p>Claude Computer-Use Agent жертвы посетил сайт исследователя и загрузил его содержимое в свой контекст.</p></a>
</div>
