---
atlas_id: AML.T0110.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-07-31"
description: Злоумышленники могут отравить исполняемую реализацию инструмента ИИ-агента, чтобы обычный вызов инструмента приводил к неавторизованному поведению или скрытым побочным эффектам. Инструмент может продолжать...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-07-31"
platforms:
    - Agentic AI
procedure_count: 1
source_name: Implementation
subtechnique_count: 0
subtechnique_of: AML.T0110
tactics:
    - AML.TA0006
title: Реализация
url: /techniques/AML.T0110.001/
---

Злоумышленники могут отравить исполняемую реализацию инструмента ИИ-агента, чтобы обычный вызов инструмента приводил к неавторизованному поведению или скрытым побочным эффектам. Инструмент может продолжать предоставлять заявленную функциональность, одновременно получая доступ к дополнительным данным, изменяя запросы, меняя получателей или места назначения, выполняя неавторизованные команды, ослабляя защитные механизмы либо скрытно эксфильтруя данные.

Отравление реализации не требует, чтобы модель интерпретировала вредоносные инструкции или следовала им. Вредоносный эффект создаётся исполняемой логикой, когда агент или пользователь вызывает инструмент. Например, отравленный почтовый инструмент может отправлять запрошенное сообщение, незаметно добавляя подконтрольного злоумышленнику получателя в поле скрытой копии (BCC). Инструмент обработки файлов может возвращать запрошенный результат, одновременно передавая исходный файл внешнему сервису.

Отравление реализации может быть внедрено до публикации посредством компрометации репозитория исходного кода инструмента или процесса сборки либо через вредоносное обновление после того, как пользователи начали использовать безопасную версию. Установленные экземпляры могут продолжать демонстрировать отравленное поведение даже после удаления вредоносного пакета или записи о нём в удалённом каталоге[[koi-postmark]].


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
<a class="relation-item" href="/techniques/AML.T0110.002/"><span class="relation-id">AML.T0110.002</span><strong>Ответ во время выполнения</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0053/"><span class="relation-id">AML.CS0053</span><strong>Эксфильтрация писем через отравленный MCP-сервер Postmark</strong><span class="relation-meta">Актор: Unknown Bad Actor / Тактика: AML.TA0006 Закрепление</span><p>После включения отравленного MCP-сервера Postmark в конфигурацию ИИ-агентов организации его вредоносное воздействие сохраняется.</p></a>
</div>


## Источники

- [First Malicious MCP in the Wild: The Postmark Backdoor That's Stealing Your Emails](https://www.koi.ai/blog/postmark-mcp-npm-malicious-backdoor-email-theft)
