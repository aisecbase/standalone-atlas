---
atlas_id: AML.T0112.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-03-30"
description: Злоумышленники могут добиться полной компрометации системы, внедряя вредоносные ИИ-артефакты, например модели или данные, содержащие встроенное вредоносное ПО или другие вредоносные команды. ИИ-артефакты часто...
generated: true
generated_by: atlasgen
maturity: feasible
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 0
source_name: AI Artifacts
subtechnique_count: 0
subtechnique_of: AML.T0112
tactics:
    - AML.TA0011
title: ИИ-артефакты
url: /techniques/AML.T0112.001/
---

Злоумышленники могут добиться полной компрометации системы, внедряя вредоносные ИИ-артефакты, например модели или данные, содержащие встроенное вредоносное ПО или другие вредоносные команды. ИИ-артефакты часто хранятся в реестрах моделей или хранилищах данных и могут затрагивать множество систем, которые загружают эти ресурсы.

Вредоносное содержимое, сохраненное в ИИ-артефактах, может выполняться из-за небезопасных форматов сериализации, например Python pickle, либо через другие поставляемые вместе с артефактом скрипты или ноутбуки.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0011/"><span class="relation-id">AML.TA0011</span><strong>Воздействие</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0112/"><span class="relation-id">AML.T0112</span><strong>Компрометация машины</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0112.000/"><span class="relation-id">AML.T0112.000</span><strong>Локальный ИИ-агент</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0005/"><span class="relation-id">AML.M0005</span><strong>Контроль доступа к ИИ-моделям и хранимым данным</strong><p>Restrict write access to model registries and AI artifact stores.</p></a>
</div>
