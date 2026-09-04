---
atlas_id: AML.T0016.004
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-08-31"
description: Злоумышленники могут искать и получать инструменты, расширяющие возможности ИИ-агента. С их помощью агент может взаимодействовать с операционными системами, браузерами, сетями, облачными сервисами, хранилищами данных,...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-08-31"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
    - Enterprise
procedure_count: 1
source_name: AI Agent Tools
subtechnique_count: 0
subtechnique_of: AML.T0016
tactics:
    - AML.TA0003
title: Инструменты ИИ-агента
url: /techniques/AML.T0016.004/
---

Злоумышленники могут искать и получать инструменты, расширяющие возможности ИИ-агента. С их помощью агент может взаимодействовать с операционными системами, браузерами, сетями, облачными сервисами, хранилищами данных, репозиториями ПО, системами управления идентификацией или другими внешними ресурсами.

Инструменты ИИ-агента могут распространяться в виде серверов Model Context Protocol (MCP), плагинов, навыков, коннекторов, библиотек функций, адаптеров для управления компьютером, брокеров выполнения, удалённых API или аналогичных интеграций. Злоумышленники могут получать легитимные инструменты и настраивать их для применения во вредоносных целях, приобретать модифицированные или специально созданные инструменты либо объединять несколько интеграций в набор инструментов для проведения операций.

Инструменты агента могут предоставлять описания, передаваемые модели, а также интерфейсы для выполнения действий. И описания, и интерфейсы влияют на то, какие возможности выбирает агент и как он их задействует. Благодаря таким инструментам агенту могут стать доступны ресурсы, учётные данные или действия, недоступные при использовании одного лишь инференса модели.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0016/"><span class="relation-id">AML.T0016</span><strong>Получение средств для атаки</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0016.000/"><span class="relation-id">AML.T0016.000</span><strong>Готовые реализации состязательных атак на ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.001/"><span class="relation-id">AML.T0016.001</span><strong>Программные инструменты</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.002/"><span class="relation-id">AML.T0016.002</span><strong>Генеративный ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.003/"><span class="relation-id">AML.T0016.003</span><strong>Эксплойты</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0070/"><span class="relation-id">AML.CS0070</span><strong>Threat Actor Uses a DeepSeek-Powered Hermes Agent in Langflow and n8n Exploitation Attempts</strong><span class="relation-meta">Актор: Chinese-speaking threat actor using the aliases knaithe and KnYuan / Тактика: AML.TA0003 Подготовка ресурсов</span><p>The actor obtained agent-specific capabilities, including Hermes&#39;s framework-bundled godmode skill and the open-source FofaMap MCP server. The MCP server exposed FOFA asset search, natural-language query translation, and Nuclei scan generation to DeepSeek. Unit 42 does not establish that godmode was invoked during the recovered session.</p></a>
</div>
