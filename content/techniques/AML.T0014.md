---
atlas_id: AML.T0014
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут выявлять общее семейство модели. Общие сведения о модели могут быть раскрыты в документации, либо злоумышленник может использовать специально подготовленные примеры и анализировать ответы модели,...
generated: true
generated_by: atlasgen
maturity: feasible
mitigation_count: 3
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 0
source_name: Discover AI Model Family
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0008
title: Выявление семейства ИИ-модели
url: /techniques/AML.T0014/
---

Злоумышленники могут выявлять общее семейство модели.

Общие сведения о модели могут быть раскрыты в документации, либо злоумышленник может использовать специально подготовленные примеры и анализировать ответы модели, чтобы классифицировать ее.

Знание семейства модели может помочь злоумышленнику определить способы атаки на модель и адаптировать атаку под нее.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0008/"><span class="relation-id">AML.TA0008</span><strong>Выявление</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0002/"><span class="relation-id">AML.M0002</span><strong>Обфускация выходных данных предиктивного ИИ</strong><p>Рекомендуемые подходы:</p><ul><li>ограничить количество показываемых результатов</li><li>ограничить детализацию онтологии выходных классов</li><li>использовать методы рандомизированного сглаживания</li><li>снизить точность числовых выходных данных</li></ul></a>
<a class="relation-item" href="/mitigations/AML.M0004/"><span class="relation-id">AML.M0004</span><strong>Ограничение объёма и частоты запросов к ИИ-сервису</strong><p>Ограничивайте запросы на инференс, чтобы снизить способность злоумышленника определить семейство ИИ-модели.</p></a>
<a class="relation-item" href="/mitigations/AML.M0006/"><span class="relation-id">AML.M0006</span><strong>Ансамбли моделей предиктивного ИИ</strong><p>Используйте несколько разных моделей, чтобы запутать злоумышленников относительно типа используемой модели и способа ее применения.</p></a>
</div>
