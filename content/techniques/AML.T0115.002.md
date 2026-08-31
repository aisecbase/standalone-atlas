---
aliases:
    - /techniques/AML.T0104/
atlas_id: AML.T0115.002
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-07-31"
description: Злоумышленники могут создавать и публиковать отравленные инструменты ИИ-агента. Отравленные инструменты могут содержать вредоносные определения или инструкции, доступные модели, скрытое поведение, реализуемое...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 1
modified_date: "2026-07-31"
platforms:
    - Agentic AI
procedure_count: 3
source_name: AI Agent Tools
subtechnique_count: 0
subtechnique_of: AML.T0115
tactics:
    - AML.TA0003
title: Инструменты ИИ-агента
url: /techniques/AML.T0115.002/
---

Злоумышленники могут создавать и публиковать отравленные инструменты ИИ-агента. Отравленные инструменты могут содержать вредоносные определения или инструкции, доступные модели, скрытое поведение, реализуемое исполняемой логикой, либо ответы во время выполнения, предназначенные для манипулирования ИИ-агентом. Инструменты могут публиковаться через репозитории исходного кода, реестры пакетов, реестры инструментов или навыков ИИ-агентов либо удалённые сервисы, контролируемые злоумышленником. Инструмент может быть вновь созданным либо представлять собой модифицированный вариант легитимного инструмента и может проявлять вредоносное поведение при его выборе, установке или вызове ИИ-агентом жертвы.

Злоумышленники могут распространять отравленные инструменты через открытые репозитории систем контроля версий (например, GitHub или GitLab), реестры пакетов (например, npm) или специализированные репозитории для обмена инструментами (например, OpenClaw Hub). Такие реестры могут почти не регулироваться и содержать множество отравленных инструментов [[opensourcemalware]]. Инструменты также могут публиковаться в виде удалённо размещённых серверов [[mcpservers]].


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0115/"><span class="relation-id">AML.T0115</span><strong>Публикация отравленных ИИ-артефактов</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0115.000/"><span class="relation-id">AML.T0115.000</span><strong>Наборы данных</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0115.001/"><span class="relation-id">AML.T0115.001</span><strong>Модели</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0016/"><span class="relation-id">AML.M0016</span><strong>Сканирование уязвимостей</strong><p>Реестры инструментов сканируют загружаемые пакеты инструментов и их зависимости на наличие вредоносного кода и уязвимостей перед добавлением в каталог.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0049/"><span class="relation-id">AML.CS0049</span><strong>Компрометация цепочки поставки через отравленный навык ClawdBot</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователь разработал отравленный навык ClawdBot под названием &#34;What Would Elon Do?&#34; Вредоносный промпт находился в файле `rules/logic.md`, который считывается при активации навыка. Исследователь опубликовал навык в ClawdHub.</p></a>
<a class="relation-item" href="/studies/AML.CS0053/"><span class="relation-id">AML.CS0053</span><strong>Эксфильтрация писем через отравленный MCP-сервер Postmark</strong><span class="relation-meta">Актор: Unknown Bad Actor / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленник опубликовал вредоносную версию `postmark-mcp` в npm.</p></a>
<a class="relation-item" href="/studies/AML.CS0054/"><span class="relation-id">AML.CS0054</span><strong>Эксфильтрация данных через отравленный удалённый MCP-инструмент</strong><span class="relation-meta">Актор: Invariant Labs / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи разместили отравленный MCP-сервер, содержащий вредоносные инструкции, скрытые в строке документации (docstring) одного из предоставляемых инструментов.</p></a>
</div>


## Источники

- [Remote MCP Servers | Awesome MCP Servers](https://mcpservers.org/remote-mcp-servers)
- [ClawdBot Skills Just Ganked Your Crypto | OpenSourceMalware](https://opensourcemalware.com/blog/clawdbot-skills-ganked-your-crypto)
