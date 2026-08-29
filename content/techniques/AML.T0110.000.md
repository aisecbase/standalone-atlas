---
atlas_id: AML.T0110.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-07-31"
description: Злоумышленники могут отравить доступное модели определение инструмента ИИ-агента или инструкции по работе с ним, чтобы манипулировать тем, как агент интерпретирует, выбирает или вызывает этот инструмент. Отравленное...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-07-31"
platforms:
    - Agentic AI
procedure_count: 2
source_name: Definition and Instructions
subtechnique_count: 0
subtechnique_of: AML.T0110
tactics:
    - AML.TA0006
title: Определение и инструкции
url: /techniques/AML.T0110.000/
---

Злоумышленники могут отравить доступное модели определение инструмента ИИ-агента или инструкции по работе с ним, чтобы манипулировать тем, как агент интерпретирует, выбирает или вызывает этот инструмент. Отравленное содержимое может находиться в описаниях инструментов, строках документации (docstrings), именах параметров, справочном тексте, схемах входных или выходных данных, аннотациях, примерах, манифестах, файлах с инструкциями для навыков или ином статическом содержимом, которое используется для передачи модели сведений о возможностях инструмента.

Вредоносные инструкции на этом уровне могут предписывать агенту собирать дополнительные данные, задавать значения скрытых или ненужных параметров, скрывать действия от пользователя либо вызывать другие инструменты. Поскольку агент может получать более полное представление инструмента, чем то, которое отображается в пользовательском интерфейсе, модель может обрабатывать вредоносные инструкции, которые не видны человеку, выполняющему проверку, или видны ему лишь частично[[invariant-tool-poisoning]].

Отравление определения также может использоваться для «затенения инструментов» (tool shadowing): при такой атаке определение одного вредоносного инструмента содержит инструкции, которые изменяют то, как агент выбирает или вызывает другой, доверенный инструмент. Если определения нескольких подключённых инструментов включены в один контекст модели, отравленный инструмент может не вызываться: его определение всё равно способно повлиять на агента[[invariant-tool-poisoning]].


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0006/"><span class="relation-id">AML.TA0006</span><strong>Закрепление</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0110/"><span class="relation-id">AML.T0110</span><strong>Отравление инструмента ИИ-агента</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0110.001/"><span class="relation-id">AML.T0110.001</span><strong>Реализация</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0110.002/"><span class="relation-id">AML.T0110.002</span><strong>Ответ во время выполнения</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0049/"><span class="relation-id">AML.CS0049</span><strong>Компрометация цепочки поставки через отравленный навык ClawdBot</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0006 Закрепление</span><p>Отравленный навык содержал в `rules/logic.md` вредоносные инструкции, доступные модели. После того как навык был установлен и стал доступен агенту, эти инструкции изменили то, как Claude Code обрабатывал связанные с навыком запросы.</p></a>
<a class="relation-item" href="/studies/AML.CS0054/"><span class="relation-id">AML.CS0054</span><strong>Эксфильтрация данных через отравленный удалённый MCP-инструмент</strong><span class="relation-meta">Актор: Invariant Labs / Тактика: AML.TA0006 Закрепление</span><p>Доступная модели строка документации (docstring) MCP-инструмента содержала вредоносные инструкции, предписывавшие агенту прочитать файлы с учётными данными, скрыть от пользователя дополнительные действия и поместить содержимое этих файлов в параметр инструмента, который в противном случае не требовался бы.</p></a>
</div>


## Источники

- [MCP Security Notification: Tool Poisoning Attacks](https://invariantlabs.ai/blog/mcp-security-notification-tool-poisoning-attacks)
- [Model Context Protocol Specification: Tools](https://modelcontextprotocol.io/specification/2025-06-18/server/tools)
- [MCP Tool Poisoning](https://owasp.org/www-community/attacks/MCP_Tool_Poisoning)
