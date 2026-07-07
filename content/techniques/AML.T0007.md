---
atlas_id: AML.T0007
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут искать в закрытых источниках артефакты обучения ИИ, присутствующие в системе, и собирать сведения о них. К таким артефактам могут относиться программный стек, используемый для обучения и...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 3
source_name: Discover AI Artifacts
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0008
title: Выявление ИИ-артефактов
url: /techniques/AML.T0007/
---

Злоумышленники могут искать в закрытых источниках артефакты обучения ИИ, присутствующие в системе, и собирать сведения о них.

К таким артефактам могут относиться программный стек, используемый для обучения и развертывания моделей, системы управления обучающими и тестовыми данными, реестры контейнеров, репозитории программного обеспечения и каталоги моделей.

Эти сведения могут использоваться для выявления целей для дальнейшего сбора данных, эксфильтрации или нарушения работы, а также для адаптации и улучшения атак.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0008/"><span class="relation-id">AML.TA0008</span><strong>Выявление</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0005/"><span class="relation-id">AML.M0005</span><strong>Контроль доступа к ИИ-моделям и хранимым данным</strong><p>Контроль доступа может ограничить способность злоумышленника выявлять ИИ-модели, наборы данных и другие артефакты в системе.</p></a>
<a class="relation-item" href="/mitigations/AML.M0012/"><span class="relation-id">AML.M0012</span><strong>Шифрование чувствительной информации</strong><p>Шифрование ИИ-артефактов может защитить от попыток злоумышленника выявить чувствительную информацию.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0008 Выявление</span><p>Исследователь мог искать ИИ-модели в среде организации-жертвы.</p></a>
<a class="relation-item" href="/studies/AML.CS0028/"><span class="relation-id">AML.CS0028</span><strong>Подмена ИИ-модели через атаку на цепочку поставок</strong><span class="relation-meta">Актор: Trend Micro Nebula Cloud Research Team / Тактика: AML.TA0008 Выявление</span><p>Исследователи обнаружили 1 453 уникальные ИИ-модели, встроенные в приватные контейнерные образы. Около половины из них были в формате Open Neural Network Exchange (ONNX).</p></a>
<a class="relation-item" href="/studies/AML.CS0058/"><span class="relation-id">AML.CS0058</span><strong>Извлечение ИИ-моделей из Google Photos</strong><span class="relation-meta">Актор: Skyld / Тактика: AML.TA0008 Выявление</span><p>Исследователи Skyld проанализировали Android Package (APK) Google Photos и определили, что приложение использует TensorFlow Lite как фреймворк машинного обучения. Они искали артефакты TFLite в пакете приложения и нативных библиотеках с помощью файлового идентификатора TFL3.</p></a>
</div>
