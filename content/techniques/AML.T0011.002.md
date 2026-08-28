---
atlas_id: AML.T0011.002
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-01-30"
description: Жертва может вызвать отравленный инструмент при взаимодействии со своим ИИ-агентом. Отравленный инструмент может выполнить промпт-инъекцию в LLM или инициировать вызов инструмента ИИ-агента. Отравленные инструменты...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-07-31"
platforms:
    - Agentic AI
procedure_count: 3
source_name: Poisoned AI Agent Tool
subtechnique_count: 0
subtechnique_of: AML.T0011
tactics:
    - AML.TA0005
title: Отравленный инструмент ИИ-агента
url: /techniques/AML.T0011.002/
---

Жертва может вызвать отравленный инструмент при взаимодействии со своим ИИ-агентом. Отравленный инструмент может выполнить [промпт-инъекцию в LLM](/techniques/AML.T0051) или инициировать [вызов инструмента ИИ-агента](/techniques/AML.T0053).

Отравленные инструменты ИИ-агентов могут внедряться в среду жертвы через [ПО для ИИ](/techniques/AML.T0010.001), либо пользователь может настроить своего агента на подключение к удаленным инструментам.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0005/"><span class="relation-id">AML.TA0005</span><strong>Выполнение</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0011/"><span class="relation-id">AML.T0011</span><strong>Запуск пользователем</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0011.000/"><span class="relation-id">AML.T0011.000</span><strong>Небезопасные ИИ-артефакты</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0011.001/"><span class="relation-id">AML.T0011.001</span><strong>Вредоносный пакет</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0011.003/"><span class="relation-id">AML.T0011.003</span><strong>Вредоносная ссылка</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0049/"><span class="relation-id">AML.CS0049</span><strong>Компрометация цепочки поставки через отравленный навык ClawdBot</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0005 Выполнение</span><p>Когда пользователь спрашивал Claude Code: &#34;what would Elon do?&#34;, Claude Code вызывал отравленный навык.</p></a>
<a class="relation-item" href="/studies/AML.CS0053/"><span class="relation-id">AML.CS0053</span><strong>Эксфильтрация писем через отравленный MCP-сервер Postmark</strong><span class="relation-meta">Актор: Unknown Bad Actor / Тактика: AML.TA0005 Выполнение</span><p>Когда пользователи в организации-жертве поручали своему ИИ-агенту использовать инструменты отравленного MCP-сервера Postmark, выполнялся вредоносный код.</p></a>
<a class="relation-item" href="/studies/AML.CS0054/"><span class="relation-id">AML.CS0054</span><strong>Эксфильтрация данных через отравленный удалённый MCP-инструмент</strong><span class="relation-meta">Актор: Invariant Labs / Тактика: AML.TA0005 Выполнение</span><p>Когда пользователь отправил запрос, соответствовавший заявленной функциональности отравленного инструмента, агент вызвал отравленный MCP-инструмент и последовал вредоносным инструкциям из его определения.</p></a>
</div>
