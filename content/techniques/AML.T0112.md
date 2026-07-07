---
atlas_id: AML.T0112
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-03-30"
description: Злоумышленники могут скомпрометировать машину, эксплуатируя компоненты с поддержкой ИИ в системе или манипулируя ими. Компрометация системы жертвы позволяет злоумышленнику выполнять произвольный код, красть учетные...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: Machine Compromise
subtechnique_count: 2
subtechnique_of: ""
tactics:
    - AML.TA0011
title: Компрометация машины
url: /techniques/AML.T0112/
---

Злоумышленники могут скомпрометировать машину, эксплуатируя компоненты с поддержкой ИИ в системе или манипулируя ими. Компрометация системы жертвы позволяет злоумышленнику выполнять произвольный код, красть учетные данные, эксфильтровать данные и сохранять закрепление в системе.

Злоумышленники могут нацеливаться на [локального ИИ-агента](/techniques/AML.T0112.000), компрометация которого дает им возможности и разрешения агента, или на [ИИ-артефакты](/techniques/AML.T0112.001), которые могут содержать встроенное вредоносное ПО.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0011/"><span class="relation-id">AML.TA0011</span><strong>Воздействие</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0112.000/"><span class="relation-id">AML.T0112.000</span><strong>Локальный ИИ-агент</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0112.001/"><span class="relation-id">AML.T0112.001</span><strong>ИИ-артефакты</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0062/"><span class="relation-id">AML.CS0062</span><strong>RCE-уязвимость в Semantic Kernel Search Plugin</strong><span class="relation-meta">Актор: Microsoft Defender Security Research Team / Тактика: AML.TA0011 Воздействие</span><p>Исследователи смогли выполнить произвольный код, что привело бы к компрометации хостовой машины.</p></a>
</div>
