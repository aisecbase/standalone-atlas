---
atlas_id: AML.T0110
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-03-30"
description: 'Злоумышленники могут добиться закрепления, отравляя инструменты, используемые ИИ-агентами: встроенные инструменты или инструменты, доступные агенту через подключения Model Context Protocol (MCP). Это предполагает...'
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-07-31"
platforms:
    - Agentic AI
procedure_count: 0
source_name: AI Agent Tool Poisoning
subtechnique_count: 3
subtechnique_of: ""
tactics:
    - AML.TA0006
title: Отравление инструмента ИИ-агента
url: /techniques/AML.T0110/
---

Злоумышленники могут добиться закрепления, отравляя инструменты, используемые ИИ-агентами: встроенные инструменты или инструменты, доступные агенту через подключения Model Context Protocol (MCP). Это предполагает компрометацию легитимных инструментов, уже интегрированных в среду агента.

Меняя поведение инструмента, например изменяя параметры или описания, внедряя скрытую логику или перенаправляя выходные данные, злоумышленники могут сохранять долгосрочное влияние на действия, решения или внешние взаимодействия агента. Отравленные инструменты могут незаметно эксфильтровать данные, выполнять несанкционированные команды или манипулировать последующими процессами, не вызывая подозрений.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0006/"><span class="relation-id">AML.TA0006</span><strong>Закрепление</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0110.000/"><span class="relation-id">AML.T0110.000</span><strong>Определение и инструкции</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0110.001/"><span class="relation-id">AML.T0110.001</span><strong>Имплементация</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0110.002/"><span class="relation-id">AML.T0110.002</span><strong>Runtime Response</strong><span class="relation-meta">Подтехника</span></a>
</div>
