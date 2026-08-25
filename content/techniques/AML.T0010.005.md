---
atlas_id: AML.T0010.005
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-03-30"
description: Злоумышленники могут нацеливаться на инструменты ИИ-агентов как на способ компрометации цепочки поставок ИИ организации-жертвы. Инструменты добавляют ИИ-агентам возможности, позволяя им взаимодействовать с другими...
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

Злоумышленники могут нацеливаться на инструменты ИИ-агентов как на способ компрометации цепочки поставок ИИ организации-жертвы. Инструменты добавляют ИИ-агентам возможности, позволяя им взаимодействовать с другими сервисами, подключаться к источникам данных, обращаться к интернет-ресурсам, запускать системные инструменты и выполнять код. Они являются привлекательной целью для злоумышленников, поскольку компрометация ИИ-агента может предоставить им широкие доступы и права в системе жертвы через другие инструменты агента.

Отравленные инструменты агентов (см. [отравление инструментов ИИ-агента](/techniques/AML.T0110)) могут содержать вредоносный код или [промпт-инъекции в LLM](/techniques/AML.T0051), которые манипулируют поведением агента и даже изменяют порядок вызова других инструментов. Злоумышленники уже успешно использовали отравленный MCP-сервер для эксфильтрации приватных пользовательских данных [[koi]].

Инструменты агентов резко выросли в популярности: тысячи MCP-серверов доступны публично [[glama]]. Их часто публикуют в репозиториях ПО с открытым исходным кодом, таких как GitHub, индексируют в специализированных хабах MCP-серверов [[mcp-hub]][[mcp-server-hub]] и размещают в реестрах пакетов, таких как NPM. ИИ-агенты также могут подключаться к удаленно размещенным инструментам [[remote-mcp]]. Это создает среду, в которой вредоносные инструменты могут быстро распространяться, а защитные меры часто отсутствуют.


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
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>Красная команда по ИИ</strong><p>Introduce a controlled untrusted agent tool or tool definition. Verify source authorization, integrity, review, permission boundaries, and safe activation.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0049/"><span class="relation-id">AML.CS0049</span><strong>Компрометация цепочки поставки через отравленный навык ClawdBot</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0004 Первичный доступ</span><p>Пользователи скачивали отравленный навык из ClawdHub. При этом ClawdHub показывает не все файлы навыка, поэтому пользователям трудно проверить навык перед скачиванием.</p></a>
<a class="relation-item" href="/studies/AML.CS0053/"><span class="relation-id">AML.CS0053</span><strong>Эксфильтрация писем через отравленный MCP-сервер Postmark</strong><span class="relation-meta">Актор: Unknown Bad Actor / Тактика: AML.TA0004 Первичный доступ</span><p>Когда организации обновляли `postmark-mcp` до версии `1.0.16`, они получали вредоносную версию инструмента через скомпрометированную цепочку поставок.</p></a>
<a class="relation-item" href="/studies/AML.CS0054/"><span class="relation-id">AML.CS0054</span><strong>Эксфильтрация данных через удаленный отравленный MCP-инструмент</strong><span class="relation-meta">Актор: Invariant Labs / Тактика: AML.TA0004 Первичный доступ</span><p>Исследователи разместили отравленный MCP-инструмент, где вредоносные инструкции были скрыты в docstring-описании этого инструмента.</p></a>
</div>


## Источники

- [Glama](https://glama.ai/mcp/servers)
- [First Malicious MCP in the Wild: The Postmark Backdoor That's Stealing Your Emails](https://www.koi.ai/blog/postmark-mcp-npm-malicious-backdoor-email-theft)
- [MCP Hub](https://www.mcphub.ai/)
- [MCP Server Hub](https://mcpserverhub.com/)
- [Remote MCP Servers](https://mcpservers.org/remote-mcp-servers)
