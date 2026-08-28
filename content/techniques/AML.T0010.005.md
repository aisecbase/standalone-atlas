---
atlas_id: AML.T0010.005
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-03-30"
description: Злоумышленники могут нацеливаться на инструменты ИИ-агента, чтобы скомпрометировать цепочку поставок ИИ организации-жертвы. Инструменты расширяют возможности ИИ-агентов, позволяя им взаимодействовать с другими...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 1
modified_date: "2026-07-31"
platforms:
    - Agentic AI
procedure_count: 3
source_name: AI Agent Tool
subtechnique_count: 0
subtechnique_of: AML.T0010
tactics:
    - AML.TA0004
title: Инструмент ИИ-агента
url: /techniques/AML.T0010.005/
---

Злоумышленники могут нацеливаться на инструменты ИИ-агента, чтобы скомпрометировать цепочку поставок ИИ организации-жертвы. Инструменты расширяют возможности ИИ-агентов, позволяя им взаимодействовать с другими сервисами, подключаться к источникам данных, получать доступ к интернет-ресурсам, запускать системные инструменты и выполнять код. Такие инструменты представляют собой привлекательную цель для злоумышленников, поскольку компрометация ИИ-агента может предоставить им обширный доступ к системе жертвы и широкие права в ней через другие инструменты агента.

Отравленные инструменты ИИ-агента могут содержать вредоносные определения или инструкции, скрытую исполняемую логику либо преднамеренно отравленные ответы во время выполнения (см. [Отравление инструмента ИИ-агента](/techniques/AML.T0110)). Такие изменения могут манипулировать поведением агента, изменять способ вызова других инструментов, создавать скрытые побочные эффекты либо предоставлять неавторизованный доступ к данным и сервисам. Злоумышленники уже успешно использовали отравленный MCP-сервер для эксфильтрации приватных данных пользователей [[koi]].

Популярность инструментов ИИ-агента стремительно возросла: тысячи MCP-серверов находятся в открытом доступе [[glama]]. Их часто публикуют в репозиториях ПО с открытым исходным кодом, таких как GitHub, индексируют в хабах, специализирующихся на MCP-серверах [[mcp-hub]][[mcp-server-hub]], и размещают в реестрах пакетов, таких как NPM. ИИ-агенты также могут подключаться к удалённо размещённым инструментам [[remote-mcp]]. В результате возникает среда, в которой вредоносные инструменты могут быстро распространяться, а меры защиты зачастую отсутствуют.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0004/"><span class="relation-id">AML.TA0004</span><strong>Первичный доступ</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0010/"><span class="relation-id">AML.T0010</span><strong>Компрометация цепочки поставок ИИ</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0010.000/"><span class="relation-id">AML.T0010.000</span><strong>Аппаратное обеспечение</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.001/"><span class="relation-id">AML.T0010.001</span><strong>ПО для ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.002/"><span class="relation-id">AML.T0010.002</span><strong>Данные</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.003/"><span class="relation-id">AML.T0010.003</span><strong>Модель</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.004/"><span class="relation-id">AML.T0010.004</span><strong>Реестр контейнеров</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>Красная команда по ИИ</strong><p>Внедряйте недоверенный инструмент ИИ-агента или недоверенное определение инструмента только в контролируемую среду. Убедитесь, что источник входит в число разрешённых; проверьте целостность, факт проведения проверки, границы разрешений и безопасность активации.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0049/"><span class="relation-id">AML.CS0049</span><strong>Компрометация цепочки поставки через отравленный навык ClawdBot</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0004 Первичный доступ</span><p>Пользователи скачивали отравленный навык из ClawdHub. При этом ClawdHub показывает не все файлы навыка, поэтому пользователям трудно проверить навык перед скачиванием.</p></a>
<a class="relation-item" href="/studies/AML.CS0053/"><span class="relation-id">AML.CS0053</span><strong>Эксфильтрация писем через отравленный MCP-сервер Postmark</strong><span class="relation-meta">Актор: Unknown Bad Actor / Тактика: AML.TA0004 Первичный доступ</span><p>Когда организации обновляли `postmark-mcp` до версии `1.0.16`, они получали вредоносную версию инструмента через скомпрометированную цепочку поставок.</p></a>
<a class="relation-item" href="/studies/AML.CS0054/"><span class="relation-id">AML.CS0054</span><strong>Эксфильтрация данных через отравленный удалённый MCP-инструмент</strong><span class="relation-meta">Актор: Invariant Labs / Тактика: AML.TA0004 Первичный доступ</span><p>Жертва настроила своего ИИ-агента на подключение к удалённому MCP-серверу под контролем злоумышленника, в результате чего отравленное определение инструмента попало в набор инструментов, доступных агенту.</p></a>
</div>


## Источники

- [Glama](https://glama.ai/mcp/servers)
- [First Malicious MCP in the Wild: The Postmark Backdoor That's Stealing Your Emails](https://www.koi.ai/blog/postmark-mcp-npm-malicious-backdoor-email-theft)
- [MCP Hub](https://www.mcphub.ai/)
- [MCP Server Hub](https://mcpserverhub.com/)
- [Remote MCP Servers](https://mcpservers.org/remote-mcp-servers)
