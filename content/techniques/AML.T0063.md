---
atlas_id: AML.T0063
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-03-12"
description: Злоумышленники могут выявлять выходные данные модели, например оценки классов, наличие которых не требуется для работы системы и которые не предназначены для использования конечным пользователем. Выходные данные...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 4
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 2
source_name: Discover AI Model Outputs
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0008
title: Выявление выходных данных ИИ-модели
url: /techniques/AML.T0063/
---

Злоумышленники могут выявлять выходные данные модели, например оценки классов, наличие которых не требуется для работы системы и которые не предназначены для использования конечным пользователем. Выходные данные модели могут находиться в журналах или включаться в ответы API.

Выходные данные модели могут позволить злоумышленнику выявить слабые места модели и разработать атаки.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0008/"><span class="relation-id">AML.TA0008</span><strong>Выявление</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0002/"><span class="relation-id">AML.M0002</span><strong>Обфускация выходных данных предиктивного ИИ</strong><p>Обфускация выходных данных модели может помешать злоумышленникам собирать чувствительную информацию о выходных данных модели.</p></a>
<a class="relation-item" href="/mitigations/AML.M0012/"><span class="relation-id">AML.M0012</span><strong>Шифрование чувствительной информации</strong><p>Шифрование выходных данных модели может помешать злоумышленникам выявлять чувствительную информацию об ИИ-системе или ее операциях.</p></a>
<a class="relation-item" href="/mitigations/AML.M0017/"><span class="relation-id">AML.M0017</span><strong>Методы распространения ИИ-моделей</strong><p>Отказ от развертывания моделей на периферийных устройствах снижает способность злоумышленника собирать чувствительную информацию о выходных данных модели.</p></a>
<a class="relation-item" href="/mitigations/AML.M0019/"><span class="relation-id">AML.M0019</span><strong>Контроль доступа к ИИ-моделям и данным в продакшене</strong><p>Контроль доступа к модели в продакшене может помочь помешать злоумышленникам извлекать информацию из выходных данных модели.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0003/"><span class="relation-id">AML.CS0003</span><strong>Обход ИИ-детектора вредоносного ПО Cylance</strong><span class="relation-meta">Актор: Skylight Cyber / Тактика: AML.TA0008 Выявление</span><p>Исследователи включили подробное логирование, раскрывающее внутреннюю логику работы ML-модели, особенно в части репутационного скоринга и ансамблирования моделей.</p></a>
<a class="relation-item" href="/studies/AML.CS0008/"><span class="relation-id">AML.CS0008</span><strong>Обход ProofPoint</strong><span class="relation-meta">Актор: Researchers at Silent Break Security / Тактика: AML.TA0008 Выявление</span><p>Исследователи обнаружили, что ProofPoint Email Protection оставляла выходные оценки модели в заголовках писем.</p></a>
</div>
