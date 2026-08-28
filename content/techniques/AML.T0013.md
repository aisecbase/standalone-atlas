---
atlas_id: AML.T0013
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут выявлять онтологию выходного пространства ИИ-модели, например типы объектов, которые модель способна обнаруживать. Злоумышленник может выявить онтологию с помощью повторяющихся запросов к модели,...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Predictive AI
procedure_count: 1
source_name: Discover AI Model Ontology
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0008
title: Выявление онтологии ИИ-модели
url: /techniques/AML.T0013/
---

Злоумышленники могут выявлять онтологию выходного пространства ИИ-модели, например типы объектов, которые модель способна обнаруживать.

Злоумышленник может выявить онтологию с помощью повторяющихся запросов к модели, вынуждая ее перечислять свое выходное пространство.

Также онтология может быть раскрыта в конфигурационном файле или документации к модели.

Онтология модели помогает злоумышленнику понять, как организация-жертва использует модель.

Эта информация полезна злоумышленнику при создании целевых атак.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0008/"><span class="relation-id">AML.TA0008</span><strong>Выявление</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0002/"><span class="relation-id">AML.M0002</span><strong>Обфускация выходных данных предиктивного ИИ</strong><p>Рекомендуемые подходы:</p><ul><li>ограничить количество показываемых результатов</li><li>ограничить детализацию онтологии выходных классов</li><li>использовать методы рандомизированного сглаживания</li><li>снизить точность числовых выходных данных</li></ul></a>
<a class="relation-item" href="/mitigations/AML.M0004/"><span class="relation-id">AML.M0004</span><strong>Ограничение объёма и частоты запросов к ИИ-сервису</strong><p>Ограничивайте запросы на инференс, чтобы снизить способность злоумышленника выявить полную онтологию выходных данных модели.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0012/"><span class="relation-id">AML.CS0012</span><strong>Обход системы идентификации лиц с помощью физических контрмер</strong><span class="relation-meta">Актор: MITRE AI Red Team / Тактика: AML.TA0008 Выявление</span><p>Команда определила список идентичностей, на которые была нацелена модель, отправляя запросы к API инференса целевой модели.</p></a>
</div>
