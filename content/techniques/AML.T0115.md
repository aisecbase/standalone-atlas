---
atlas_id: AML.T0115
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-07-31"
description: Злоумышленники могут создавать или модифицировать ИИ-артефакты и публиковать их через общедоступные или совместно используемые каналы распространения, способствуя компрометации нижестоящих ИИ-систем. К отравленным...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 3
modified_date: "2026-07-31"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 0
source_name: Publish Poisoned AI Artifacts
subtechnique_count: 3
subtechnique_of: ""
tactics:
    - AML.TA0003
title: Публикация отравленных ИИ-артефактов
url: /techniques/AML.T0115/
---

Злоумышленники могут создавать или модифицировать ИИ-артефакты и публиковать их через общедоступные или совместно используемые каналы распространения, способствуя компрометации нижестоящих ИИ-систем. К отравленным ИИ-артефактам могут относиться наборы данных, модели и инструменты ИИ-агентов, содержащие вредоносное содержимое, вредоносные варианты поведения, вредоносный код или вредоносные конфигурации.

Злоумышленники могут публиковать новые артефакты или вредоносные варианты легитимных артефактов через репозитории наборов данных или моделей, реестры пакетов, репозитории исходного кода, хабы инструментов или удалённо размещённые сервисы. Впоследствии жертвы могут получить и интегрировать эти артефакты в ходе [компрометации цепочки поставок ИИ](/techniques/AML.T0010).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0115.000/"><span class="relation-id">AML.T0115.000</span><strong>Наборы данных</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0115.001/"><span class="relation-id">AML.T0115.001</span><strong>Модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0115.002/"><span class="relation-id">AML.T0115.002</span><strong>Инструменты ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0007/"><span class="relation-id">AML.M0007</span><strong>Санитизация обучающих данных</strong><p>Репозитории наборов данных проверяют поступающие материалы и до публикации в каталоге помещают в карантин отравленные образцы данных, метки, аннотации или метаданные.</p></a>
<a class="relation-item" href="/mitigations/AML.M0008/"><span class="relation-id">AML.M0008</span><strong>Валидация ИИ-модели</strong><p>Model repositories evaluate submissions for backdoors, data leakage, adversarial influence, and unexpected behavior before listing.</p></a>
<a class="relation-item" href="/mitigations/AML.M0016/"><span class="relation-id">AML.M0016</span><strong>Сканирование уязвимостей</strong><p>Реестры моделей и инструментов ИИ-агентов сканируют загружаемые артефакты на наличие вредоносного содержимого перед добавлением в каталог.</p></a>
</div>
