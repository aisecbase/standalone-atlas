---
atlas_id: AML.T0110.002
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-07-31"
description: Злоумышленники могут отравить канал, по которому вредоносный или скомпрометированный инструмент ИИ-агента возвращает ответы во время выполнения, намеренно включая в эти ответы содержимое, предназначенное для...
generated: true
generated_by: atlasgen
maturity: feasible
mitigation_count: 0
modified_date: "2026-07-31"
platforms:
    - Agentic AI
procedure_count: 0
source_name: Runtime Response
subtechnique_count: 0
subtechnique_of: AML.T0110
tactics:
    - AML.TA0006
title: Ответ во время выполнения
url: /techniques/AML.T0110.002/
---

Злоумышленники могут отравить канал, по которому вредоносный или скомпрометированный инструмент ИИ-агента возвращает ответы во время выполнения, намеренно включая в эти ответы содержимое, предназначенное для воздействия на последующие рассуждения, решения или действия модели. Отравленные ответы могут содержать вредоносные инструкции, вводящие в заблуждение данные, сфабрикованные сообщения об ошибках, встроенные ресурсы или иное содержимое, предназначенное для того, чтобы модель воспринимала его как доверенный контекст после одобренного вызова инструмента.

Поскольку ответы инструментов обычно включаются в контекст модели, злоумышленник может использовать их, чтобы побуждать агента вызывать дополнительные инструменты, получать доступ к чувствительной информации, изменять ход выполняемого рабочего процесса или передавать данные в контролируемую злоумышленником точку назначения. Отравленные инструкции могут быть смешаны с легитимными результатами так, чтобы создавалось впечатление штатной работы инструмента. В ответах может использоваться структурированное или неструктурированное содержимое, включая текст, изображения, ссылки на ресурсы, встроенные ресурсы или поля, соответствующие схеме [[mcp-tools]][[owasp-tool-poisoning]].


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
<a class="relation-item" href="/techniques/AML.T0110.000/"><span class="relation-id">AML.T0110.000</span><strong>Определение и инструкции</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0110.001/"><span class="relation-id">AML.T0110.001</span><strong>Реализация</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Источники

- [Model Context Protocol Specification: Tools](https://modelcontextprotocol.io/specification/2025-11-25/server/tools)
- [MCP Tool Poisoning](https://owasp.org/www-community/attacks/MCP_Tool_Poisoning)
- [MCP Security Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/MCP_Security_Cheat_Sheet.html)
